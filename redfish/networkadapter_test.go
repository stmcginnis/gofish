//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkAdapterBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#NetworkAdapter.NetworkAdapter",
		"@odata.type": "#NetworkAdapter.v1_0_0.NetworkAdapter",
		"@odata.id": "/redfish/v1/NetworkAdapter",
		"Id": "NetworkAdapter-1",
		"Name": "NetworkAdapterOne",
		"Description": "NetworkAdapter One",
		"Controllers": [{
			"ControllerCapabilities": {
				"DataCenterBridging": {
					"Capable": true
				},
				"NPAR": {
					"NparCapable": false,
					"NparEnabled": false
				},
				"NPIV": {
					"MaxDeviceLogins": 1024,
					"MaxPortLogins": 1024
				},
				"NetworkDeviceFunctionCount": 1,
				"NetworkPortCount": 2,
				"VirtualizationOffload": {
					"SRIOV": {
						"SRIOVVEPACapable": true
					},
					"VirtualFunction": {
						"DeviceMaxCount": 1024,
						"MinAssignmentGroupSize": 2,
						"NetworkPortMaxCount": 2
					}
				}
			},
			"FirmwarePackageVersion": "1.2.3",
			"Links": {
				"NetworkDeviceFunctions": [{
					"@odata.id": "/redfish/v1/NetworkAdapters/DeviceFunction-1"
				}],
				"NetworkDeviceFunctions@odata.count": 1,
				"NetworkPorts": [{
						"@odata.id": "/redfish/v1/NetworkAdapters/Port-1"
					},
					{
						"@odata.id": "/redfish/v1/NetworkAdapters/Port-2"
					}
				],
				"NetworkPorts@odata.count": 2,
				"PCIeDevices": [{
					"@odata.id": "/redfish/v1/NetworkAdapters/PCIeDevice-1"
				}],
				"PCIeDevices@odata.count": 1
			},
			"PCIeInterface": {
				"LanesInUse": 32,
				"MaxLanes": 32,
				"MaxPCIeType": "Gen4",
				"PCIeType": "Gen4"
			}
		}],
		"Manufacturer": "Acme Storage",
		"Model": "Anvil 3000",
		"NetworkDeviceFunctions": {
			"@odata.id": "/redfish/v1/NetworkDevice/Functions"
		},
		"NetworkPorts": {
			"@odata.id": "/redfish/v1/NetworkPorts"
		},
		"ParNumber": "3000",
		"SKU": "1234",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestNetworkAdapter tests the parsing of NetworkAdapter objects.
func TestNetworkAdapter(t *testing.T) {
	var result NetworkAdapter
	err := json.NewDecoder(networkAdapterBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NetworkAdapter-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "NetworkAdapterOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.Controllers[0].ControllerCapabilities.DataCenterBridging.Capable {
		t.Error("DCB should be enabled")
	}

	if result.Controllers[0].ControllerCapabilities.NPIV.MaxDeviceLogins != 1024 {
		t.Errorf("Received incorrect Controller NPIC max device logins: %d",
			result.Controllers[0].ControllerCapabilities.NPIV.MaxDeviceLogins)
	}

	if result.Controllers[0].PCIeInterface.MaxPCIeType != Gen4PCIeTypes {
		t.Errorf("Received incorrect max PCIe type: %s", result.Controllers[0].PCIeInterface.MaxPCIeType)
	}

	if result.Controllers[0].PCIeInterface.PCIeType != Gen4PCIeTypes {
		t.Errorf("Received incorrect PCIe type: %s", result.Controllers[0].PCIeInterface.PCIeType)
	}
}
