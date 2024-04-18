//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkInterfaceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#NetworkInterface.NetworkInterface",
		"@odata.type": "#NetworkInterface.v1_1_2.NetworkInterface",
		"@odata.id": "/redfish/v1/NetworkInterface",
		"Id": "NetworkInterface-1",
		"Name": "NetworkInterfaceOne",
		"Description": "NetworkInterface One",
		"InterfaceEnabled": false,
		"PermanentMACAddress": "BE:3A:F2:B6:05:9F",
		"MACAddress": "00:00:00:00:00:00",
		"SpeedMbps": 100,
		"AutoNeg": true,
		"FullDuplex": true,
		"MTUSize": 0,
		"IPv4Addresses": [
			{
			"Address": "169.254.3.1",
			"SubnetMask": "255.255.255.0",
			"AddressOrigin": "Static",
			"Gateway": "169.254.3.254"
			}
		],
		"IPv4StaticAddresses": [
			{
			"Address": "169.254.3.1",
			"SubnetMask": "255.255.255.0",
			"Gateway": "169.254.3.254"
			}
		],
		"IPv6StaticAddresses": [
			{
			"Address": "::",
			"PrefixLength": 64
			},
			{
			"Address": "::",
			"PrefixLength": 64
			},
			{
			"Address": "::",
			"PrefixLength": 64
			},
			{
			"Address": "::",
			"PrefixLength": 64
			},
			{
			"Address": "::",
			"PrefixLength": 64
			}
		],
		"Links": {
			"NetworkAdapter": {
					"@odata.id": "/redfish/v1/NetworkAdapters/1"
			}
		},
		"NetworkDeviceFunctions": {
			"@odata.id": "/redfish/v1/NetworkFunctions"
		},
		"NetworkPorts": {
			"@odata.id": "/redfish/v1/NetworkPorts"
		},
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestNetworkInterface tests the parsing of NetworkInterface objects.
func TestNetworkInterface(t *testing.T) {
	var result NetworkInterface
	err := json.NewDecoder(networkInterfaceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "NetworkInterface-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "NetworkInterfaceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}
}
