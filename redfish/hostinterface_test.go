//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var hostInterfaceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#HostInterface.HostInterface",
		"@odata.type": "#HostInterface.v1_0_0.HostInterface",
		"@odata.id": "/redfish/v1/HostInterface",
		"Id": "HostInterface-1",
		"Name": "HostInterfaceOne",
		"Description": "HostInterface One",
		"AuthNoneRoleId": "role-1",
		"AuthenticationModes": [
			"BasicAuth",
			"RedfishSessionAuth"
		],
		"ExternallyAccessible": true,
		"FirmwareAuthEnabled": false,
		"FirmwareAuthRoleId": "role-1",
		"HostEthernetInterfaces": {
			"@odata.id": "/redfish/v1/Host/1/EthernetInterfaceCollection"
		},
		"HostInterfaceType": "NetworkHostInterface",
		"InterfaceEnabled": true,
		"KernelAuthEnabled": false,
		"KernelAuthRoleId": "role-2",
		"Links": {
			"AuthNoneRole": {
				"@odata.id": "/redfish/v1/Roles/role-1"
			},
			"ComputerSystems": [{
				"@odata.id": "/redfish/v1/Systems/System-1"
				}
			],
			"ComputerSystems@odata.count": 1,
			"FirmwareAuthRole": {
				"@odata.id": "/redfish/v1/Roles/role-1"
			},
			"KernelAuthRole": {
				"@odata.id": "/redfish/v1/Roles/role-2"
			}
		},
		"ManagerEthernetInterface": {
			"@odata.id": "/redfish/v1/Host/1/EthernetInterface/1"
		},
		"NetworkProtocol": {
			"@odata.id": "/redfish/v1/Host/1/ManagerNetworkProtocol/1"
		},
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestHostInterface tests the parsing of HostInterface objects.
func TestHostInterface(t *testing.T) {
	var result HostInterface
	err := json.NewDecoder(hostInterfaceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "HostInterface-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "HostInterfaceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.AuthenticationModes) != 2 {
		t.Errorf("Expected 2 auth modes, got: %d", len(result.AuthenticationModes))
	}

	if !result.ExternallyAccessible {
		t.Error("Should be externally accessible")
	}

	if result.FirmwareAuthEnabled {
		t.Error("Firmware auth should not be enabled")
	}

	if result.authNoneRole != "/redfish/v1/Roles/role-1" {
		t.Errorf("Received incorrect auth role none link: %s", result.authNoneRole)
	}

	if len(result.computerSystems) != 1 {
		t.Errorf("Should be 1 computer system, got %d", len(result.computerSystems))
	}
}
