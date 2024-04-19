//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var fabricBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Fabric.Fabric",
		"@odata.id": "/redfish/v1/Fabrics/PCIe",
		"@odata.type": "#Fabric.v1_3_0.Fabric",
		"Description": "PCIe Fabric",
		"FabricType": "PCIe",
		"Id": "PCIe",
		"Links": {},
		"Name": "PCIe Fabric",
		"Status": {
		  "Health": "OK",
		  "HealthRollup": "OK",
		  "State": "Enabled"
		},
		"Switches": {
		  "@odata.id": "/redfish/v1/Fabrics/PCIe/Switches"
		}
	  }`)

// TestFabric tests the parsing of Fabric objects.
func TestFabric(t *testing.T) {
	var result Fabric
	err := json.NewDecoder(fabricBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "PCIe" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "PCIe Fabric" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	assertEquals(t, "OK", string(result.Status.Health))
	assertEquals(t, "/redfish/v1/Fabrics/PCIe/Switches", result.switches)
}
