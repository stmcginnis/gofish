//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var thermalEquipmentBody = `{
	"@odata.type": "#ThermalEquipment.v1_1_1.ThermalEquipment",
	"Id": "ThermalEquipment",
	"Name": "Cooling Equipment",
	"Status": {
	  "State": "Enabled",
	  "HealthRollup": "OK"
	},
	"CDUs": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs"
	},
	"CoolingLoops": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CoolingLoops"
	},
	"@odata.id": "/redfish/v1/ThermalEquipment"
  }`

// TestThermalEquipment tests the parsing of ThermalEquipment objects.
func TestThermalEquipment(t *testing.T) {
	var result ThermalEquipment
	err := json.NewDecoder(strings.NewReader(thermalEquipmentBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ThermalEquipment", result.ID)
	assertEquals(t, "Cooling Equipment", result.Name)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs", result.cdus)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CoolingLoops", result.coolingLoops)
}
