//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
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

// TestGetCollectionObjectsOrder verifies that GetCollectionObjects returns
// members in the order the service returned them, even though each member is
// resolved concurrently.
func TestGetCollectionObjectsOrder(t *testing.T) {
	const count = 25

	var members strings.Builder
	for i := 1; i <= count; i++ {
		if i > 1 {
			members.WriteString(",")
		}
		fmt.Fprintf(&members,
			`{"@odata.id":"/redfish/v1/Systems/System-%d","Id":"%d"}`, i, i)
	}
	body := fmt.Sprintf(
		`{"@odata.id":"/redfish/v1/Systems","Members@odata.count":%d,"Members":[%s]}`,
		count, members.String())

	client := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
				},
			},
		},
	}

	systems, err := GetCollectionObjects[ComputerSystem](client, "/redfish/v1/Systems")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(systems) != count {
		t.Fatalf("expected %d members, got %d", count, len(systems))
	}
	for i, s := range systems {
		want := fmt.Sprintf("%d", i+1)
		if s.ID != want {
			t.Errorf("member %d out of order: expected Id %q, got %q", i, want, s.ID)
		}
	}
}
