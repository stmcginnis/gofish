//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var biosBody = strings.NewReader(
	`{
		"@odata.type": "#Bios.v1_0_6.Bios",
		"@odata.context": "/redfish/v1/$metadata#Bios.Bios",
		"@odata.id": "/redfish/v1/Systems/437XR1138R2/BIOS",
		"Id": "BIOS",
		"Name": "BIOS Configuration Current Settings",
		"Description": "BIOD Settings",
		"AttributeRegistry": "BiosAttributeRegistryP89.v1_0_0",
		"Attributes": {
			"AdminPhone": "",
			"BootMode": "Uefi",
			"EmbeddedSata": "Raid",
			"NicBoot1": "NetworkBoot",
			"NicBoot2": "Disabled",
			"PowerProfile": "MaxPerf",
			"ProcCoreDisable": 0,
			"ProcHyperthreading": "Enabled",
			"ProcTurboMode": "Enabled",
			"UsbControl": "UsbEnabled"
		},
		"Actions": {
			"#Bios.ResetBios": {
				"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ResetBios"
			},
			"#Bios.ChangePassword": {
				"target": "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ChangePassword"
			}
		}
	}`)

// TestBios tests the parsing of Bios objects.
func TestBios(t *testing.T) {
	var result Bios
	err := json.NewDecoder(biosBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "BIOS" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "BIOS Configuration Current Settings" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AttributeRegistry != "BiosAttributeRegistryP89.v1_0_0" {
		t.Errorf("Received incorrect attribute registry: %s", result.AttributeRegistry)
	}
}
