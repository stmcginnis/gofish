//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
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

func TestCollectResourceCollectionDefaultConcurrency(t *testing.T) {
	resources := testCollectionResources(6)
	maxActive := collectResourceCollectionMaxActive(resources)

	if maxActive != defaultCollectionRequestConcurrency {
		t.Errorf("Expected default concurrency %d, got %d", defaultCollectionRequestConcurrency, maxActive)
	}
}

func TestCollectResourceCollectionConfiguredConcurrency(t *testing.T) {
	resources := testCollectionResources(8)
	concurrency := 5
	maxActive := collectResourceCollectionMaxActive(
		resources,
		WithCollectionQueryOpts(WithCollectionRequestConcurrency(concurrency)),
	)

	if maxActive != concurrency {
		t.Errorf("Expected configured concurrency %d, got %d", concurrency, maxActive)
	}
}

func testCollectionResources(count int) []*Resource {
	resources := make([]*Resource, 0, count)
	for i := 0; i < count; i++ {
		resources = append(resources, &Resource{Entity: Entity{ODataID: fmt.Sprintf("/redfish/v1/Systems/%d", i)}})
	}
	return resources
}

func collectResourceCollectionMaxActive(resources []*Resource, queryOpts ...QueryGroupOption) int {
	var mu sync.Mutex
	active := 0
	maxActive := 0

	CollectResourceCollection(func(_ *Resource, _ ...QueryGroupOption) {
		mu.Lock()
		active++
		if active > maxActive {
			maxActive = active
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)

		mu.Lock()
		active--
		mu.Unlock()
	}, resources, queryOpts...)

	return maxActive
}
