//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
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
