//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var coolantConnectorBody = `{
	"@odata.type": "#CoolantConnector.v1_0_1.CoolantConnector",
	"Id": "A",
	"Name": "Rack Cooling Loop A",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"CoolantConnectorType": "Pair",
	"RatedFlowLitersPerMinute": 30,
	"FlowLitersPerMinute": {
	  "Reading": 24.3
	},
	"SupplyTemperatureCelsius": {
	  "Reading": 14.8
	},
	"ReturnTemperatureCelsius": {
	  "Reading": 38.2
	},
	"DeltaTemperatureCelsius": {
	  "Reading": 23.4
	},
	"SupplyPressurekPa": {
	  "Reading": 426.6
	},
	"ReturnPressurekPa": {
	  "Reading": 409.9
	},
	"DeltaPressurekPa": {
	  "Reading": 31.7
	},
	"Links": {
	  "ConnectedCoolingLoop": {
		"@odata.id": "/redfish/v1/ThermalEquipment/CoolingLoops/Rack4"
	  }
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/SecondaryCoolantConnectors/A"
  }`

// TestCoolantConnector tests the parsing of CoolantConnector objects.
func TestCoolantConnector(t *testing.T) {
	var result CoolantConnector
	err := json.NewDecoder(strings.NewReader(coolantConnectorBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "A", result.ID)
	assertEquals(t, "Rack Cooling Loop A", result.Name)
	assertEquals(t, "Pair", string(result.CoolantConnectorType))
	assertEquals(t, "/redfish/v1/ThermalEquipment/CoolingLoops/Rack4", result.connectedCoolingLoop)

	if result.RatedFlowLitersPerMinute != 30 {
		t.Errorf("Unexpected RatedFlowLitersPerMinute, got %.2f", result.RatedFlowLitersPerMinute)
	}

	if *result.FlowLitersPerMinute.Reading != 24.3 {
		t.Errorf("Unexpected FlowLitersPerMinute reading, got %.2f", *result.FlowLitersPerMinute.Reading)
	}

	if *result.SupplyTemperatureCelsius.Reading != 14.8 {
		t.Errorf("Unexpected SupplyTemperatureCelsius reading, got %.2f", *result.SupplyTemperatureCelsius.Reading)
	}

	if *result.ReturnTemperatureCelsius.Reading != 38.2 {
		t.Errorf("Unexpected ReturnTemperatureCelsius reading, got %.2f", *result.ReturnTemperatureCelsius.Reading)
	}

	if *result.ReturnTemperatureCelsius.Reading != 38.2 {
		t.Errorf("Unexpected ReturnTemperatureCelsius reading, got %.2f", *result.ReturnTemperatureCelsius.Reading)
	}

	if *result.DeltaTemperatureCelsius.Reading != 23.4 {
		t.Errorf("Unexpected DeltaTemperatureCelsius reading, got %.2f", *result.DeltaTemperatureCelsius.Reading)
	}

	if *result.SupplyPressurekPa.Reading != 426.6 {
		t.Errorf("Unexpected SupplyPressurekPa reading, got %.2f", *result.SupplyPressurekPa.Reading)
	}

	if *result.ReturnPressurekPa.Reading != 409.9 {
		t.Errorf("Unexpected ReturnPressurekPa reading, got %.2f", *result.ReturnPressurekPa.Reading)
	}

	if *result.DeltaPressurekPa.Reading != 31.7 {
		t.Errorf("Unexpected DeltaPressurekPa reading, got %.2f", *result.DeltaPressurekPa.Reading)
	}
}
