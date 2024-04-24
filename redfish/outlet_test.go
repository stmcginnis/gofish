//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var outletBody = `{
	"@odata.type": "#Outlet.v1_4_2.Outlet",
	"Id": "A1",
	"Name": "Outlet A1, Branch Circuit A",
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"PhaseWiringType": "OnePhase3Wire",
	"VoltageType": "AC",
	"OutletType": "NEMA_5_20R",
	"RatedCurrentAmps": 20,
	"NominalVoltage": "AC120V",
	"LocationIndicatorActive": true,
	"PowerOnDelaySeconds": 4,
	"PowerOffDelaySeconds": 0,
	"PowerState": "On",
	"PowerEnabled": true,
	"Voltage": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageA1",
	  "Reading": 117.5
	},
	"PolyPhaseVoltage": {
	  "Line1ToNeutral": {
		"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageA1",
		"Reading": 117.5
	  }
	},
	"CurrentAmps": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/CurrentA1",
	  "Reading": 1.68
	},
	"PolyPhaseCurrentAmps": {
	  "Line1": {
		"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/CurrentA1",
		"Reading": 1.68
	  }
	},
	"PowerWatts": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PowerA1",
	  "Reading": 197.4,
	  "ApparentVA": 197.4,
	  "ReactiveVAR": 0,
	  "PowerFactor": 1
	},
	"FrequencyHz": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/FrequencyA1",
	  "Reading": 60
	},
	"EnergykWh": {
	  "DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/EnergyA1",
	  "Reading": 36166
	},
	"Actions": {
	  "#Outlet.PowerControl": {
		"target": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1/Outlet.PowerControl"
	  },
	  "#Outlet.ResetMetrics": {
		"target": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1/Outlet.ResetMetrics"
	  }
	},
	"Links": {
	  "BranchCircuit": {
		"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Branches/A"
	  }
	},
	"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1"
  }`

// TestOutlet tests the parsing of Outlet objects.
func TestOutlet(t *testing.T) {
	var result Outlet
	err := json.NewDecoder(strings.NewReader(outletBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "A1", result.ID)
	assertEquals(t, "Outlet A1, Branch Circuit A", result.Name)
	assertEquals(t, "OnePhase3Wire", string(result.PhaseWiringType))
	assertEquals(t, "AC", string(result.VoltageType))
	assertEquals(t, "NEMA_5_20R", string(result.OutletType))
	assertEquals(t, "AC120V", string(result.NominalVoltage))
	assertEquals(t, "On", string(result.PowerState))
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageA1", result.Voltage.DataSourceURI)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/VoltageA1", result.PolyPhaseVoltage.Line1ToNeutral.DataSourceURI)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Branches/A", result.branchCircuit)

	if !result.LocationIndicatorActive {
		t.Error("Expected location indicator to be active")
	}

	if !result.PowerEnabled {
		t.Error("Expected power to be enabled")
	}

	if result.PowerOnDelaySeconds != 4 {
		t.Errorf("Unexpected PowerOnDelaySeconds value: %.2f", result.PowerOnDelaySeconds)
	}

	if result.PowerOffDelaySeconds != 0 {
		t.Errorf("Unexpected PowerOffDelaySeconds value: %.2f", result.PowerOffDelaySeconds)
	}

	if result.RatedCurrentAmps != 20 {
		t.Errorf("Unexpected RatedCurrentAmps value: %.2f", result.RatedCurrentAmps)
	}

	if result.CurrentAmps.Reading != 1.68 {
		t.Errorf("Unexpected current amps reading: %.2f", result.CurrentAmps.Reading)
	}

	if result.PolyPhaseCurrentAmps.Line1.Reading != 1.68 {
		t.Errorf("Unexpected PolyPhaseCurrentAmps line1 reading: %.2f", result.PolyPhaseCurrentAmps.Line1.Reading)
	}

	if result.PowerWatts.ReactiveVAR != 0 {
		t.Errorf("Unexpected PowerWatts.ReactiveVAR value: %.2f", result.PowerWatts.ReactiveVAR)
	}
}

// TestOutletResetMetrics tests the Outlet ResetMetrics call.
func TestOutletResetMetrics(t *testing.T) {
	var result Outlet
	err := json.NewDecoder(strings.NewReader(outletBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.ResetMetrics()
	if err != nil {
		t.Errorf("Error making ResetMetrics call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if calls[0].Payload != "" {
		t.Errorf("Expected payload: %s", calls[0].Payload)
	}

	if calls[0].URL != "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1/Outlet.ResetMetrics" {
		t.Errorf("Expected target URL: %s", calls[0].URL)
	}
}

// TestOutletPowerControl tests the Outlet PowerControl call.
func TestOutletPowerControl(t *testing.T) {
	var result Outlet
	err := json.NewDecoder(strings.NewReader(outletBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.PowerControl(OffPowerState)
	if err != nil {
		t.Errorf("Error making PowerControl call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "PowerState:Off") {
		t.Errorf("Expected payload: %s", calls[0].Payload)
	}

	if calls[0].URL != "/redfish/v1/PowerEquipment/RackPDUs/1/Outlets/A1/Outlet.PowerControl" {
		t.Errorf("Expected target URL: %s", calls[0].URL)
	}
}
