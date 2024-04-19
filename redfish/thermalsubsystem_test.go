//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var thermalSubsystemBody = `{
	"@odata.context": "/redfish/v1/$metadata#ThermalSubsystem.ThermalSubsystem",
	"@odata.id": "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem",
	"@odata.type": "#ThermalSubsystem.v1_2_0.ThermalSubsystem",
	"Description": "Represents the properties of a Thermal Subsystem of the Chassis",
	"Name": "Thermal Subsystem for Chassis",
	"FanRedundancy": [
	  {
		"RedundancyType": "NPlusM",
		"MaxSupportedInGroup": null,
		"MinNeededInGroup": null,
		"RedundancyGroup": [],
		"RedundancyGroup@odata.count": 0,
		"Status": {
		  "State": "Enabled",
		  "Health": "OK"
		}
	  }
	],
	"Fans": {
	  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/Fans"
	},
	"ThermalMetrics": {
	  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/ThermalMetrics"
	},
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Id": "ThermalSubsystem"
  }`

// TestThermalSubsystem tests the parsing of ThermalSubsystem objects.
func TestThermalSubsystem(t *testing.T) {
	var result ThermalSubsystem
	err := json.NewDecoder(strings.NewReader(thermalSubsystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "ThermalSubsystem" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Thermal Subsystem for Chassis" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.fans != "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/Fans" {
		t.Errorf("Invalid fans link: %s", result.fans)
	}
}
