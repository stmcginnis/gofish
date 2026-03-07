//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

var collectionBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ComputerSystemCollection.ComputerSystemCollection",
		"@odata.id": "/redfish/v1/ComputerSystemCollection",
		"@odata.type": "#ComputerSystemCollection.1.0.0.ComputerSystemCollection",
		"Name": "Test Collection",
		"Links": {
			"Members@odata.count": 2,
			"Members": [
				{
					"@odata.id": "/redfish/v1/Systems/System-1"
				},
				{
					"@odata.id": "/redfish/v1/Systems/System-2"
				}
			]
		}
	}`)

// TestCollection tests the parsing of Collections.
func TestCollection(t *testing.T) {
	var result Collection
	err := json.NewDecoder(collectionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.Name != "Test Collection" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.ItemLinks) != 2 {
		t.Errorf("Expected 2 items in collection, got %d", len(result.ItemLinks))
	}

	linkRoot := "/redfish/v1/Systems/System-%d"
	for i, item := range result.ItemLinks {
		endpoint := fmt.Sprintf(linkRoot, i+1)
		if item != endpoint {
			t.Errorf("Expected link to '%s', got '%s'", endpoint, item)
		}
	}
}

// TestGetCollectionObjectsOrder verifies that GetCollectionObjects preserves
// the server-provided Members order regardless of concurrent fetch completion
// order. The collection page lists members as /3, /1, /2 (deliberately
// non-sequential) and the test asserts the returned slice matches that order.
func TestGetCollectionObjectsOrder(t *testing.T) {
	// Collection page: members listed as 3, 1, 2 to prove ordering is driven
	// by the collection page, not by HTTP response arrival order.
	collectionPage := `{
		"@odata.id": "/redfish/v1/Items",
		"@odata.type": "#ItemCollection.ItemCollection",
		"Name": "Items",
		"Members@odata.count": 3,
		"Members": [
			{"@odata.id": "/redfish/v1/Items/3"},
			{"@odata.id": "/redfish/v1/Items/1"},
			{"@odata.id": "/redfish/v1/Items/2"}
		]
	}`

	makeItemBody := func(id string) string {
		return `{"@odata.id": "/redfish/v1/Items/` + id + `", "Id": "` + id + `", "Name": "Item ` + id + `"}`
	}

	makeResp := func(body string) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		}
	}

	// GET responses in call order:
	//   0: collection page
	//   1: /Items/3
	//   2: /Items/1
	//   3: /Items/2
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				makeResp(collectionPage),
				makeResp(makeItemBody("3")),
				makeResp(makeItemBody("1")),
				makeResp(makeItemBody("2")),
			},
		},
	}

	results, err := GetCollectionObjects[Resource, *Resource](testClient, "/redfish/v1/Items")
	if err != nil {
		t.Fatalf("GetCollectionObjects returned unexpected error: %v", err)
	}

	wantOrder := []string{
		"/redfish/v1/Items/3",
		"/redfish/v1/Items/1",
		"/redfish/v1/Items/2",
	}

	if len(results) != len(wantOrder) {
		t.Fatalf("Expected %d results, got %d", len(wantOrder), len(results))
	}

	for i, want := range wantOrder {
		if got := results[i].ODataID; got != want {
			t.Errorf("results[%d]: want ODataID %q, got %q", i, want, got)
		}
	}
}
