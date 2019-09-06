//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var simpleStorageBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#SimpleStorage.SimpleStorage",
		"@odata.id": "/redfish/v1/Systems/System-1/SimpleStorage/SAS-CTRL0",
		"@odata.type": "#SimpleStorage.v1_2_2.SimpleStorage",
		"Description": "Simple Storage",
		"Devices": [
			{
				"CapacityBytes": 549755813888,
				"Manufacturer": "Generic",
				"Model": "Generic SATA Disk",
				"Name": "Disk 0",
				"Status": {
					"Health": "OK",
					"State": "Enabled"
				}
			}
		],
		"Id": "SAS-CTRL0",
		"Links": {
			"Chassis": {
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			}
		},
		"Name": "Simple Storage Controller",
		"Status": {
			"Health": "OK",
			"State": "Enabled"
		},
		"UefiDevicePath": "ACPI(PnP)/PCI(1,0)/SAS(0x31000004CF13F6BD,0, SATA)"
	}`)

// TestSimpleStorage tests the parsing of SimpleStorage objects.
func TestSimpleStorage(t *testing.T) {
	var result SimpleStorage
	err := json.NewDecoder(simpleStorageBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SAS-CTRL0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Simple Storage Controller" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Devices[0].CapacityBytes != 549755813888 {
		t.Errorf("Invalid device capacity: %d", result.Devices[0].CapacityBytes)
	}

	if result.UefiDevicePath != "ACPI(PnP)/PCI(1,0)/SAS(0x31000004CF13F6BD,0, SATA)" {
		t.Errorf("Invalid UefiDevicePath: %s", result.UefiDevicePath)
	}
}
