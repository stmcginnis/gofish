//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var pumpBody = `{
	"@odata.type": "#Pump.v1_0_1.Pump",
	"Id": "1",
	"PumpType": "Liquid",
	"Name": "Immersion Unit Pump",
	"Version": "1.03b",
	"ProductionDate": "2021-06-24T08:00:00Z",
	"Manufacturer": "Contoso",
	"Model": "UP-JAM",
	"SerialNumber": "29347ZT599",
	"PartNumber": "MAARS",
	"AssetTag": "PDX5-92399",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"PumpSpeedPercent": {
	  "Reading": 62,
	  "SpeedRPM": 1800
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/ImmersionUnits/1/Pumps/1"
  }`

// TestPump tests the parsing of Pump objects.
func TestPump(t *testing.T) {
	var result Pump
	err := json.NewDecoder(strings.NewReader(pumpBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Immersion Unit Pump", result.Name)
	assertEquals(t, "Liquid", string(result.PumpType))
	assertEquals(t, "1.03b", result.Version)

	if *result.PumpSpeedPercent.Reading != 62 {
		t.Errorf("Unexpected PumpSpeedPercent.Reading: %.2f", *result.PumpSpeedPercent.Reading)
	}

	if *result.PumpSpeedPercent.SpeedRPM != 1800 {
		t.Errorf("Unexpected PumpSpeedPercent.SpeedRPM: %.2f", *result.PumpSpeedPercent.SpeedRPM)
	}
}
