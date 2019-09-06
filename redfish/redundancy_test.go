//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var redundancyBody = strings.NewReader(
	`{
		"@odata.id": "/redfish/v1/Redundancy",
		"Id": "Redundancy-1",
		"Name": "RedundancyOne",
		"MaxNumSupported": 2,
		"MemberId": "Redundancy1",
		"MinNumNeeded": 2,
		"Mode": "Sparing",
		"RedundancyEnabled": true,
		"RedundancySet": [
			"One",
			"Two"
		],
		"RedundancySet@odata.count": 2,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestRedundancy tests the parsing of Redundancy objects.
func TestRedundancy(t *testing.T) {
	var result Redundancy
	err := json.NewDecoder(redundancyBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Redundancy-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "RedundancyOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaxNumSupported != 2 {
		t.Errorf("Invalid MaxNumSupported: %d", result.MaxNumSupported)
	}

	if result.Mode != SparingRedundancyMode {
		t.Errorf("Invalid mode: %s", result.Mode)
	}
}
