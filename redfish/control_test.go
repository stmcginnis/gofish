//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var controlBody = `{
	"@odata.type": "#Control.v1_5_0.Control",
	"Id": "PowerLimit",
	"Name": "System Power Limit",
	"PhysicalContext": "Chassis",
	"ControlType": "Power",
	"ControlMode": "Automatic",
	"SetPoint": 500,
	"SetPointUnits": "W",
	"AllowableMax": 1000,
	"AllowableMin": 150,
	"Sensor": {
	  "Reading": 374,
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/TotalPower"
	},
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"@odata.id": "/redfish/v1/Chassis/1U/Controls/PowerLimit"
  }`

// TestControl tests the parsing of Control objects.
func TestControl(t *testing.T) {
	var result Control
	err := json.NewDecoder(strings.NewReader(controlBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PowerLimit", result.ID)
	assertEquals(t, "System Power Limit", result.Name)
	assertEquals(t, "Chassis", string(result.PhysicalContext))
	assertEquals(t, "Power", string(result.ControlType))
	assertEquals(t, "Automatic", string(result.ControlMode))
	assertEquals(t, "W", result.SetPointUnits)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/TotalPower", result.Sensor.DataSourceURI)

	if result.SetPoint != 500 {
		t.Errorf("Unexpected set point, got %.2f", result.SetPoint)
	}

	if result.AllowableMax != 1000 {
		t.Errorf("Unexpected allowable max, got %.2f", result.AllowableMax)
	}

	if result.AllowableMin != 150 {
		t.Errorf("Unexpected allowable min, got %.2f", result.AllowableMin)
	}

	if result.Sensor.Reading != 374 {
		t.Errorf("Unexpected sensor reading, got %.2f", result.Sensor.Reading)
	}
}
