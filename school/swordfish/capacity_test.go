//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var capacityBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Capacity.CapacitySource",
		"@odata.type": "#Capacity.v1_1_1.CapacitySource",
		"@odata.id": "/redfish/v1/CapacitySource",
		"Id": "Capacity-1",
		"Name": "CapacityOne",
		"Description": "Capacity One",
		"ProvidedCapacity": {
			"Data": {
				"AllocatedBytes": 2199023255600,
				"ConsumedBytes": 2199023255600,
				"GuaranteedBytes": 2199023255600,
				"ProvisionedBytes": 2199023255600
			},
			"IsThinProvisioned": false,
			"Metadata": {
				"AllocatedBytes": 209715200,
				"ConsumedBytes": 2048,
				"GuaranteedBytes": 209715200,
				"ProvisionedBytes": 209715200
			},
			"Snapshot": {
				"AllocatedBytes": 0,
				"ConsumedBytes": 0,
				"GuaranteedBytes": 0,
				"ProvisionedBytes": 0
			}
		},
		"ProvidingDrives": {
			"@odata.id": "/redfish/v1/System/System-1/Drives"
		}
	}`)

// TestCapacity tests the parsing of CapacitySource objects.
func TestCapacity(t *testing.T) {
	var result CapacitySource
	err := json.NewDecoder(capacityBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Capacity-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "CapacityOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.ProvidedCapacity.Data.AllocatedBytes != 2199023255600 {
		t.Errorf("Invalid data allocated bytes: %d", result.ProvidedCapacity.Data.AllocatedBytes)
	}

	if result.ProvidedCapacity.IsThinProvisioned {
		t.Error("Thin provisioning should be false")
	}

	if result.providingDrives != "/redfish/v1/System/System-1/Drives" {
		t.Errorf("Invalid providing drives link: %s", result.providingDrives)
	}

	if result.providingPools != "" {
		t.Errorf("Invalid providing pools link: %s", result.providingPools)
	}
}
