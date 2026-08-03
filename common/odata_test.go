//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

// jsonResponse builds a canned 200 response for TestClient to return.
func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     http.Header{},
	}
}

// getReturns scripts the responses TestClient hands back to successive GETs.
func getReturns(responses ...*http.Response) map[string][]any {
	returns := make([]any, 0, len(responses))
	for _, resp := range responses {
		returns = append(returns, resp)
	}
	return map[string][]any{http.MethodGet: returns}
}

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		collection bool
		defaults   []QueryGroupOption
		opts       []QueryGroupOption
		want       string
	}{
		{
			name:       "no options at all",
			url:        "/redfish/v1/Chassis/1/Drives",
			collection: true,
			want:       "/redfish/v1/Chassis/1/Drives",
		},
		{
			name:       "client default applies",
			url:        "/redfish/v1/Chassis/1/Drives",
			collection: true,
			defaults:   []QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod))},
			want:       "/redfish/v1/Chassis/1/Drives?$expand=.",
		},
		{
			name:       "per-call WithoutExpand overrides an enabled client default",
			url:        "/redfish/v1/Chassis/1/Drives",
			collection: true,
			defaults:   []QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod))},
			opts:       []QueryGroupOption{WithoutExpand()},
			want:       "/redfish/v1/Chassis/1/Drives",
		},
		{
			name:       "per-call expand applies with no client default",
			url:        "/redfish/v1/Chassis/1/Drives",
			collection: true,
			opts:       []QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionAsterisk))},
			want:       "/redfish/v1/Chassis/1/Drives?$expand=*",
		},
		{
			name:       "expand levels",
			url:        "/redfish/v1/Chassis/1/Drives",
			collection: true,
			opts: []QueryGroupOption{WithCollectionQueryOpts(
				WithExpand(ExpandOptionPeriod), WithExpandLevel(2))},
			want: "/redfish/v1/Chassis/1/Drives?$expand=.($levels=2)",
		},
		{
			name:       "collection options do not affect resource requests",
			url:        "/redfish/v1/Chassis/1/Drives/0",
			collection: false,
			defaults:   []QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod))},
			want:       "/redfish/v1/Chassis/1/Drives/0",
		},
		{
			name:       "resource options apply to resource requests",
			url:        "/redfish/v1/Chassis/1/Drives/0",
			collection: false,
			opts:       []QueryGroupOption{WithResourceQueryOpts(WithExpand(ExpandOptionPeriod))},
			want:       "/redfish/v1/Chassis/1/Drives/0?$expand=.",
		},
		{
			name:       "WithoutExpand clears the resource half too",
			url:        "/redfish/v1/Chassis/1/Drives/0",
			collection: false,
			defaults:   []QueryGroupOption{WithResourceQueryOpts(WithExpand(ExpandOptionPeriod))},
			opts:       []QueryGroupOption{WithoutExpand()},
			want:       "/redfish/v1/Chassis/1/Drives/0",
		},
		{
			// Members@odata.nextLink carries the page offset in a query
			// string, so a second "?" would make the URL unparseable.
			name:       "url that already has a query string gets an ampersand",
			url:        "/redfish/v1/Chassis/1/Drives?$skip=50",
			collection: true,
			defaults:   []QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod))},
			want:       "/redfish/v1/Chassis/1/Drives?$skip=50&$expand=.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &TestClient{Settings: ClientSettings{DefaultQueryOptions: test.defaults}}

			got := BuildQuery(c, test.url, test.collection, test.opts...)
			if got != test.want {
				t.Errorf("expected %q, got %q", test.want, got)
			}
		})
	}
}

// TestBuildQueryGroupDoesNotAliasClientDefaults guards against merging per-call
// options into the client's own slice. Appending in place would write into any
// spare capacity there, letting concurrent calls overwrite each other's
// options.
func TestBuildQueryGroupDoesNotAliasClientDefaults(t *testing.T) {
	// spare capacity is what makes an in-place append observable
	defaults := make([]QueryGroupOption, 0, 8)
	defaults = append(defaults, WithCollectionQueryOpts(WithExpandFallback(true)))

	c := &TestClient{Settings: ClientSettings{DefaultQueryOptions: defaults}}

	const callers = 64
	got := make([]ExpandOption, callers)
	want := make([]ExpandOption, callers)

	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		want[i] = ExpandOptionPeriod
		if i%2 == 0 {
			want[i] = ExpandOptionTilde
		}

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			group := BuildQueryGroup(c, WithCollectionQueryOpts(WithExpand(want[i])))
			got[i] = group.QueryCollection.expand
		}(i)
	}
	wg.Wait()

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("caller %d got expand %q, expected %q: per-call options leaked between callers",
				i, got[i], want[i])
		}
	}

	if len(c.GetSettings().DefaultQueryOptions) != 1 {
		t.Errorf("client defaults grew to %d entries, expected 1",
			len(c.GetSettings().DefaultQueryOptions))
	}
}

// TestCollectionMemberRequestsReceiveOptions checks that query options survive
// the fan-out from a collection GET to the per-member GETs it triggers.
func TestCollectionMemberRequestsReceiveOptions(t *testing.T) {
	c := &TestClient{}
	c.CustomReturnForActions = getReturns(
		// the collection: two members, links only, so each needs fetching
		jsonResponse(`{
			"Members@odata.count": 2,
			"Members": [
				{"@odata.id": "/redfish/v1/Chassis/1/Drives/0"},
				{"@odata.id": "/redfish/v1/Chassis/1/Drives/1"}
			]
		}`),
		jsonResponse(`{"@odata.id": "/redfish/v1/Chassis/1/Drives/0", "Id": "0"}`),
		jsonResponse(`{"@odata.id": "/redfish/v1/Chassis/1/Drives/1", "Id": "1"}`),
	)

	_, err := GetCollectionObjects[Resource](c, "/redfish/v1/Chassis/1/Drives",
		WithResourceQueryOpts(WithExpand(ExpandOptionPeriod)))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	calls := c.CapturedCalls()
	if len(calls) != 3 {
		t.Fatalf("expected 3 GETs, got %d", len(calls))
	}

	// the option was resource-scoped, so the collection GET must be untouched
	if calls[0].URL != "/redfish/v1/Chassis/1/Drives" {
		t.Errorf("collection GET picked up a resource-scoped option: %s", calls[0].URL)
	}

	for _, call := range calls[1:] {
		if !strings.Contains(call.URL, "$expand=.") {
			t.Errorf("member GET did not receive the query option: %s", call.URL)
		}
	}
}

// TestCollectionNextPageReceivesOptions checks that query options still apply
// once a collection paginates, and that appending them to a next-link that
// already has a query string produces a valid URL.
func TestCollectionNextPageReceivesOptions(t *testing.T) {
	c := &TestClient{
		Settings: ClientSettings{
			DefaultQueryOptions: []QueryGroupOption{
				WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod)),
			},
		},
	}
	// Members arrive with an Id, meaning the service expanded them inline, so
	// no per-member GETs follow and the only calls are the two pages.
	c.CustomReturnForActions = getReturns(
		jsonResponse(`{
			"Members@odata.count": 2,
			"Members": [{"@odata.id": "/redfish/v1/Chassis/1/Drives/0", "Id": "0"}],
			"Members@odata.nextLink": "/redfish/v1/Chassis/1/Drives?$skip=1"
		}`),
		jsonResponse(`{
			"Members@odata.count": 2,
			"Members": [{"@odata.id": "/redfish/v1/Chassis/1/Drives/1", "Id": "1"}]
		}`),
	)

	drives, err := GetCollectionObjects[Resource](c, "/redfish/v1/Chassis/1/Drives")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(drives) != 2 {
		t.Errorf("expected 2 members across both pages, got %d", len(drives))
	}

	calls := c.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("expected 2 GETs, got %d", len(calls))
	}

	if calls[0].URL != "/redfish/v1/Chassis/1/Drives?$expand=." {
		t.Errorf("unexpected first page URL: %s", calls[0].URL)
	}

	if calls[1].URL != "/redfish/v1/Chassis/1/Drives?$skip=1&$expand=." {
		t.Errorf("unexpected next page URL: %s", calls[1].URL)
	}
}

// TestGetObjectsReceivesOptions checks the path taken by accessors that hold a
// set of links rather than a collection URI.
func TestGetObjectsReceivesOptions(t *testing.T) {
	c := &TestClient{}
	c.CustomReturnForActions = getReturns(
		jsonResponse(`{"@odata.id": "/redfish/v1/Chassis/1/Drives/0", "Id": "0"}`),
		jsonResponse(`{"@odata.id": "/redfish/v1/Chassis/1/Drives/1", "Id": "1"}`),
	)

	links := []string{
		"/redfish/v1/Chassis/1/Drives/0",
		"/redfish/v1/Chassis/1/Drives/1",
	}

	_, err := GetObjects[Resource](c, links,
		WithResourceQueryOpts(WithExpand(ExpandOptionPeriod)))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	calls := c.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("expected 2 GETs, got %d", len(calls))
	}

	for _, call := range calls {
		if !strings.Contains(call.URL, "$expand=.") {
			t.Errorf("GET did not receive the query option: %s", call.URL)
		}
	}
}
