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
		"Description": "BIOS Settings",
		"AttributeRegistry": "BiosAttributeRegistryP89.v1_0_0",
		"Attributes": {
			"AdminPhone": "",
			"BootMode": "Uefi",
			"EmbeddedSata": "Raid",
			"NicBoot1": "NetworkBoot",
			"NicBoot2": "Disabled",
			"PowerProfile": "MaxPerf",
			"ProcCoreDisable": 3,
			"ProcHyperthreading": "Enabled",
			"ProcTurboMode": "Enabled",
			"UsbControl": "UsbEnabled",
			"BoolTest1": "True",
			"BoolTest2": 1,
			"BoolTest3": "NotBool"
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

	if result.resetBiosTarget != "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ResetBios" {
		t.Errorf("Invalid ResetBios link: %s", result.resetBiosTarget)
	}

	if result.changePasswordTarget != "/redfish/v1/Systems/437XR1138R2/BIOS/Actions/Bios.ChangePassword" {
		t.Errorf("Invalid ChangePassword target: %s", result.changePasswordTarget)
	}

	if result.Attributes.String("AdminPhone") != "" {
		t.Errorf("Invalid 'AdminPhone' attribute: %s", result.Attributes["AdminPhone"])
	}

	if result.Attributes.String("PowerProfile") != "MaxPerf" {
		t.Errorf("Invalid 'PowerProfile' attribute: %s", result.Attributes["PowerProfile"])
	}

	if result.Attributes.Int("ProcCoreDisable") != 3 {
		t.Errorf("Invalid 'ProcCoreDisable' attribute: %v", result.Attributes["ProcCoreDisable"])
	}

	if !result.Attributes.Bool("BoolTest1") {
		t.Errorf("Expected True boolean value for 'BoolTest1': %v", result.Attributes["BoolTest1"])
	}

	if !result.Attributes.Bool("BoolTest2") {
		t.Errorf("Expected True boolean value for 'BoolTest2': %v", result.Attributes["BoolTest1"])
	}

	if result.Attributes.Bool("BoolTest3") {
		t.Errorf("Expected False boolean value for 'BoolTest3': %v", result.Attributes["BoolTest1"])
	}
}
