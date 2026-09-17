//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type collectionPageClient struct {
	TestClient
	get func(string) (*http.Response, error)
}

func (c *collectionPageClient) Get(uri string) (*http.Response, error) { return c.get(uri) }

func collectionPageResponse(body string) *http.Response {
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body))}
}

func readCollectionPages(c Client, callback bool, opts ...QueryGroupOption) ([]string, error) {
	var ids []string
	if callback {
		err := CollectListGeneric(func(entry *LogEntry, _ ...QueryGroupOption) {
			ids = append(ids, entry.ID)
		}, c, "/logs", opts...)
		return ids, err
	}
	entries, err := GetCollectionObjects[LogEntry](c, "/logs", opts...)
	for _, entry := range entries {
		ids = append(ids, entry.ID)
	}
	return ids, err
}

func TestCollectionPagination(t *testing.T) {
	for _, callback := range []bool{false, true} {
		t.Run(fmt.Sprintf("callback=%v", callback), func(t *testing.T) {
			for _, tc := range []struct {
				name, lastLink, wantError string
				pages                     int
			}{
				{"self loop", "/logs", "pagination loop", 1},
				{"cycle to first page", "/logs", "pagination loop", 3},
				{"cycle to later page", "/logs?page=1", "pagination loop", 3},
				{"complete long collection", "", "", 257},
				{"failed next page", "/unavailable", "source unavailable", 2},
			} {
				t.Run(tc.name, func(t *testing.T) {
					pages := make(map[string]string)
					var wantIDs []string
					for i := 0; i < tc.pages; i++ {
						uri := "/logs"
						if i != 0 {
							uri = fmt.Sprintf("/logs?page=%d", i)
						}
						next := tc.lastLink
						if i+1 < tc.pages {
							next = fmt.Sprintf("/logs?page=%d", i+1)
						}
						id := fmt.Sprint(i)
						wantIDs = append(wantIDs, id)
						pages[uri] = fmt.Sprintf(`{"Members":[{"Id":%q,"@odata.id":%q}],"Members@odata.nextLink":%q}`,
							id, "/logs/"+id, next)
					}
					visits := make(map[string]int)
					client := &collectionPageClient{get: func(uri string) (*http.Response, error) {
						visits[uri]++
						if visits[uri] > 1 {
							// Bound the unfixed traversal without relying on a timeout.
							return nil, fmt.Errorf("fixture stopped a repeated request")
						}
						body, ok := pages[uri]
						if !ok {
							return nil, fmt.Errorf("source unavailable")
						}
						return collectionPageResponse(body), nil
					}}
					ids, err := readCollectionPages(client, callback)
					if tc.wantError == "" && err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if tc.wantError != "" && (err == nil || !strings.Contains(err.Error(), tc.wantError)) {
						t.Errorf("expected %q error, got %v", tc.wantError, err)
					}
					if !reflect.DeepEqual(ids, wantIDs) {
						t.Errorf("expected members %v in source order, got %v", wantIDs, ids)
					}
					for uri, count := range visits {
						if count != 1 {
							t.Errorf("page %q fetched %d times; expected once", uri, count)
						}
					}
				})
			}
		})
	}
}

func TestCollectionPaginationExpandFallback(t *testing.T) {
	for _, callback := range []bool{false, true} {
		t.Run(fmt.Sprintf("callback=%v", callback), func(t *testing.T) {
			var calls []string
			client := &collectionPageClient{get: func(uri string) (*http.Response, error) {
				calls = append(calls, uri)
				switch uri {
				case "/logs?$expand=*":
					return nil, fmt.Errorf("expansion unsupported")
				case "/logs":
					return collectionPageResponse(`{"Members":[{"Id":"first"}],"Members@odata.nextLink":"/logs?page=2"}`), nil
				case "/logs?page=2":
					return collectionPageResponse(`{"Members":[{"Id":"second"}]}`), nil
				default:
					return nil, fmt.Errorf("unexpected request %q", uri)
				}
			}}
			ids, err := readCollectionPages(client, callback,
				WithCollectionQueryOpts(WithExpand(ExpandOptionAsterisk), WithExpandFallback(true)))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if want := []string{"first", "second"}; !reflect.DeepEqual(ids, want) {
				t.Errorf("expected %v, got %v", want, ids)
			}
			if want := []string{"/logs?$expand=*", "/logs", "/logs?page=2"}; !reflect.DeepEqual(calls, want) {
				t.Errorf("expected requests %v, got %v", want, calls)
			}
		})
	}
}
