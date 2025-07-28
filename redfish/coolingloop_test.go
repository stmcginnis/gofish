//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var coolingLoopBody = `{
	"@odata.type": "#CoolingLoop.v1_0_2.CoolingLoop",
	"Id": "BuildingChiller",
	"Name": "Feed from building chiller",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"UserLabel": "Building Chiller",
	"Coolant": {
	  "CoolantType": "Water",
	  "AdditiveName": "Generic cooling water biocide",
	  "AdditivePercent": 0
	},
	"CoolantLevelStatus": "OK",
	"CoolantQuality": "OK",
	"CoolantLevelPercent": {
	  "Reading": 95
	},
	"SupplyEquipmentNames": [
	  "Chiller"
	],
	"ConsumingEquipmentNames": [
	  "Rack #1 CDU",
	  "Rack #2 CDU",
	  "Rack #3 CDU",
	  "Rack #4 CDU"
	],
	"@odata.id": "/redfish/v1/ThermalEquipment/CoolingLoops/BuildingChiller"
  }`

// TestCoolingLoop tests the parsing of CoolingLoop objects.
func TestCoolingLoop(t *testing.T) {
	var result CoolingLoop
	err := json.NewDecoder(strings.NewReader(coolingLoopBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "BuildingChiller", result.ID)
	assertEquals(t, "Feed from building chiller", result.Name)
	assertEquals(t, "Building Chiller", result.UserLabel)
	assertEquals(t, "Water", string(result.Coolant.CoolantType))
	assertEquals(t, "Generic cooling water biocide", result.Coolant.AdditiveName)
	assertEquals(t, "OK", string(result.CoolantLevelStatus))
	assertEquals(t, "OK", string(result.CoolantQuality))
	assertEquals(t, "Chiller", result.SupplyEquipmentNames[0])
	assertEquals(t, "Rack #3 CDU", result.ConsumingEquipmentNames[2])

	if *result.CoolantLevelPercent.Reading != 95 {
		t.Errorf("Unexpected CoolantLevelPercent, got %.2f", *result.CoolantLevelPercent.Reading)
	}
}
