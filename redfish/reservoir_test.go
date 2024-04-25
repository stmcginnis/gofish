//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var reservoirBody = `{
	"@odata.type": "#Reservoir.v1_0_1.Reservoir",
	"Id": "1",
	"ReservoirType": "Reserve",
	"Name": "Cooling Loop Reservoir",
	"Manufacturer": "Contoso",
	"Model": "Tarantino",
	"CapacityLiters": 10,
	"PartNumber": "Pink",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Location": {
	  "Placement": {
		"Row": "North 1"
	  }
	},
	"FluidLevelPercent": {
	  "Reading": 64.8
	},
	"InternalPressurekPa": {
	  "Reading": 138.7
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/Reservoirs/1"
  }`

// TestReservoir tests the parsing of Reservoir objects.
func TestReservoir(t *testing.T) {
	var result Reservoir
	err := json.NewDecoder(strings.NewReader(reservoirBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Cooling Loop Reservoir", result.Name)
	assertEquals(t, "Reserve", string(result.ReservoirType))
	assertEquals(t, "North 1", result.Location.Placement.Row)

	if result.CapacityLiters != 10 {
		t.Errorf("Unexpected CapacityLiters: %.2f", result.CapacityLiters)
	}

	if result.FluidLevelPercent.Reading != 64.8 {
		t.Errorf("Unexpected FluidLevelPercent.Reading: %.2f", result.FluidLevelPercent.Reading)
	}

	if result.InternalPressurekPa.Reading != 138.7 {
		t.Errorf("Unexpected InternalPressurekPa.Reading: %.2f", result.InternalPressurekPa.Reading)
	}
}
