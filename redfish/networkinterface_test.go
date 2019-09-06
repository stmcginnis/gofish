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
