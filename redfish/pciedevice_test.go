//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var pcieDeviceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#PCIeDevice.PCIeDevice",
		"@odata.type": "#PCIeDevice.v1_3_1.PCIeDevice",
		"@odata.id": "/redfish/v1/PCIeDevice",
		"Id": "PCIeDevice-1",
		"Name": "PCIeDeviceOne",
		"Description": "PCIeDevice One",
		"Assembly": {
			"@odata.id": "/redfish/v1/Assembly/1"
		},
		"AssetTag": "Tag-1",
		"DeviceType": "Simulated",
		"FirmwareVersion": "1.2",
		"Links": {
			"Chassis": [{
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			}],
			"Chassis@odata.count": 1,
			"PCIeFunctions": [{
					"@odata.id": "/redfish/v1/Functions/1"
				},
				{
					"@odata.id": "/redfish/v1/Functions/2"
				}
			],
			"PCIeFunctions@odata.count": 1
		},
		"Manufacturer": "Acme Inc",
		"Model": "A1",
		"PCIeInterface": {
			"LanesInUse": 32,
			"MaxLanes": 32,
			"MaxPCIeType": "Gen4",
			"PCIeType": "Gen4"
		},
		"PartNumber": "1234",
		"SKU": "4321",
		"SerialNumber": "A1111",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestPCIeDevice tests the parsing of PCIeDevice objects.
func TestPCIeDevice(t *testing.T) {
	var result PCIeDevice
	err := json.NewDecoder(pcieDeviceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "PCIeDevice-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "PCIeDeviceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.DeviceType != SimulatedDeviceType {
		t.Errorf("Invalid device type: %s", result.DeviceType)
	}

	if result.PCIeInterface.MaxLanes != 32 {
		t.Errorf("Invalid max lanes: %d", result.PCIeInterface.MaxLanes)
	}

	if result.PCIeInterface.PCIeType != Gen4PCIeTypes {
		t.Errorf("Invalid PCIe type: %s", result.PCIeInterface.PCIeType)
	}
}
