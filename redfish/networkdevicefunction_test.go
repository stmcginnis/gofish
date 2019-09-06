//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkDeviceFunctionBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#NetworkDeviceFunction.NetworkDeviceFunction",
		"@odata.type": "#NetworkDeviceFunction.v1_3_2.NetworkDeviceFunction",
		"@odata.id": "/redfish/v1/NetworkDeviceFunction",
		"Id": "NetworkDeviceFunction-1",
		"Name": "NetworkDeviceFunctionOne",
		"Description": "NetworkDeviceFunction One",
		"AssignablePhysicalPorts": [{
				"@odata.id": "/redfish/v1/Port/1"
			},
			{
				"@odata.id": "/redfish/v1/Port/2"
			}
		],
		"AssignablePhysicalPorts@odata.count": 2,
		"BootMode": "Disabled",
		"DeviceEnabled": true,
		"Ethernet": {
			"MACAddress": "98:E7:43:00:01:0A",
			"MTUSize": 9000,
			"PermanentMACAddress": "98:E7:43:00:01:0A",
			"VLAN": {
				"@odata.id": "/redfish/v1/VLAN/1"
			},
			"VLANs": [{
				"@odata.id": "/redfish/v1/Port/1"
			}]
		},
		"FibreChannel": {
			"AllowFIPVLANDiscovery": true,
			"BootTargets": [],
			"FCoEActiveVLANId": 500,
			"FCoELocalVLANId": 500,
			"FibreChannelId": "10",
			"PermanentWWNN": "5764839588724069681",
			"PermanentWWPN": "5764839588724069680",
			"WWNN": "5764839588724069681",
			"WWNSource": "ConfiguredLocally",
			"WWPN": "5764839588724069680"
		},
		"Links": {
			"Endpoints": [{
				"@odata.id": "/redfish/v1/Endpoints/Endpoint-1"
			}],
			"Endpoints@odata.count": 1,
			"PCIeFunction": {
				"@odata.id": "/redfish/v1/Functions/1"
			},
			"PhysicalPortsAssignment": {
				"@odata.id": "/redfish/v1/Ports/1"
			}
		},
		"MaxVirtualFunctions": 5,
		"NetDevFuncCapabilities": ["FibreChannelOverEthernet"],
		"NetDevFuncType": "FibreChannelOverEthernet",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"VirtualFunctionsEnabled": true,
		"iSCSIBoot": {}
	}`)

// TestNetworkDeviceFunction tests the parsing of NetworkDeviceFunction objects.
func TestNetworkDeviceFunction(t *testing.T) {
	var result NetworkDeviceFunction
	err := json.NewDecoder(networkDeviceFunctionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NetworkDeviceFunction-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "NetworkDeviceFunctionOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.BootMode != DisabledBootMode {
		t.Errorf("Invalid boot mode: %s", result.BootMode)
	}

	if !result.DeviceEnabled {
		t.Error("Device should be enabled")
	}

	if result.Ethernet.MACAddress != "98:E7:43:00:01:0A" {
		t.Errorf("Invalid ethernet MAC address: %s", result.Ethernet.MACAddress)
	}

	if result.FibreChannel.FCoEActiveVLANId != 500 {
		t.Errorf("Invalid active VLAN: %d", result.FibreChannel.FCoEActiveVLANId)
	}

	if result.FibreChannel.WWNSource != ConfiguredLocallyWWNSource {
		t.Errorf("Invalid WWN source: %s", result.FibreChannel.WWNSource)
	}

	if result.NetDevFuncType != FibreChannelOverEthernetNetworkDeviceTechnology {
		t.Errorf("Invalid network device function type: %s", result.NetDevFuncType)
	}
}
