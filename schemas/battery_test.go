//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var batteryBody = `{
	"@odata.type": "#Battery.v1_2_2.Battery",
	"Id": "Module1",
	"Name": "Battery 1",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Actions": {
	  "#Battery.SelfTest": {
		"target": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Actions/Battery.SelfTest"
	  },
	  "#Battery.Calibrate": {
		"target": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Actions/Battery.Calibrate"
	  }
	},
	"Location": {
	  "PartLocation": {
		"ServiceLabel": "Battery 1",
		"LocationType": "Bay",
		"LocationOrdinalValue": 0
	  }
	},
	"Model": "RKS-440DC",
	"Manufacturer": "Contoso Power",
	"FirmwareVersion": "1.00",
	"Version": "A05",
	"ProductionDate": "2019-10-01T06:00:00Z",
	"SerialNumber": "3488247",
	"PartNumber": "23456-133",
	"SparePartNumber": "93284-133",
	"LocationIndicatorActive": false,
	"HotPluggable": true,
	"CapacityRatedWattHours": 20,
	"CapacityActualWattHours": 19.41,
	"MaxDischargeRateAmps": 10,
	"StateOfHealthPercent": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1StateOfHealth",
	  "Reading": 91
	},
	"ChargeState": "Idle",
	"Metrics": {
	  "@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Metrics"
	},
	"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1"
  }`

// TestBattery tests the parsing of Battery objects.
func TestBattery(t *testing.T) {
	var result Battery
	err := json.NewDecoder(strings.NewReader(batteryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Module1", result.ID)
	assertEquals(t, "Battery 1", result.Name)
	assertEquals(t, "RKS-440DC", result.Model)
	assertEquals(t, "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Metrics", result.metrics)
	assertEquals(t, "Idle", string(result.ChargeState))

	if !result.HotPluggable {
		t.Error("Expected Hot Pluggable to be true")
	}

	if *result.StateOfHealthPercent.Reading != 91 {
		t.Errorf("Expected state of health percent reading to be 91, got %.0f", *result.StateOfHealthPercent.Reading)
	}
}

// TestBatterySelfTest tests the Battery SelfTest call.
func TestBatterySelfTest(t *testing.T) {
	var result Battery
	err := json.NewDecoder(strings.NewReader(batteryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.SelfTest()

	if err != nil {
		t.Errorf("Error making Battery SelfTest call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[0].Payload, "") {
		t.Errorf("Unexpected self test payload: %s", calls[0].Payload)
	}
}

// TestBatteryCalibrate tests the Battery Calibrate call.
func TestBatteryCalibrate(t *testing.T) {
	var result Battery
	err := json.NewDecoder(strings.NewReader(batteryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.Calibrate()

	if err != nil {
		t.Errorf("Error making Calibbrate call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[0].Payload, "") {
		t.Errorf("Unexpected calibrate payload: %s", calls[0].Payload)
	}
}
