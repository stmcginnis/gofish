//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var circuitBody = strings.NewReader(
	`{
		"@odata.type": "#Circuit.v1_7_0.Circuit",
		"Id": "A",
		"Name": "Branch Circuit A",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"CircuitType": "Branch",
		"PhaseWiringType": "TwoPhase3Wire",
		"NominalVoltage": "AC200To240V",
		"RatedCurrentAmps": 16,
		"BreakerState": "Normal",
		"PolyPhaseVoltage": {
			"Line1ToNeutral": {
				"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageAL1N",
				"Reading": 118.2
			},
			"Line1ToLine2": {
				"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageAL1L2",
				"Reading": 203.5
			}
		},
		"CurrentAmps": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/CurrentA",
			"Reading": 5.19
		},
		"PolyPhaseCurrentAmps": {
			"Line1": {
				"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/CurrentA",
				"Reading": 5.19
			}
		},
		"PowerWatts": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PowerA",
			"Reading": 937.4,
			"ApparentVA": 937.4,
			"ReactiveVAR": 0,
			"PowerFactor": 1
		},
		"PolyPhasePowerWatts": {
			"Line1ToNeutral": {
				"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PowerA1",
				"Reading": 937.4,
				"ApparentVA": 937.4,
				"ReactiveVAR": 0,
				"PowerFactor": 1
			}
		},
		"FrequencyHz": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/FrequencyA",
			"Reading": 60
		},
		"EnergykWh": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/EnergyA",
			"Reading": 325675
		},
		"Links": {
			"Outlets": [
				{
					"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1"
				},
				{
					"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A2"
				},
				{
					"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A3"
				}
			]
		},
		"Actions": {
			"#Circuit.BreakerControl": {
				"target": "/redfish/v1/PowerEquipment/RackPDUs/1/Branches/A/Circuit.BreakerControl"
			},
			"#Outlet.ResetMetrics": {
				"target": "/redfish/v1/PowerEquipment/RackPDUs/1/Branches/A/Circuit.ResetMetrics"
			}
		},
		"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Branches/A"
	}`)

// TestCircuit tests the parsing of Circuit objects.
func TestCircuit(t *testing.T) {
	var result Circuit
	err := json.NewDecoder(circuitBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "A", result.ID)
	assertEquals(t, "Branch Circuit A", result.Name)
	assertEquals(t, "Branch", fmt.Sprint(result.CircuitType))
	assertEquals(t, "TwoPhase3Wire", fmt.Sprint(result.PhaseWiringType))
	assertEquals(t, "AC200To240V", fmt.Sprint(result.NominalVoltage))
	assertEquals(t, "16", fmt.Sprint(result.RatedCurrentAmps))
	assertEquals(t, "Normal", fmt.Sprint(result.BreakerState))
	assertEquals(t, "118.2", fmt.Sprint(result.PolyPhaseVoltage.Line1ToNeutral.Reading))
	assertEquals(t, "325675", fmt.Sprint(result.EnergykWh.Reading))
}
