//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"net/http"
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

var thermalSubsystemFanCollection = `{
	"@odata.id": "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/Fans",
	"@odata.type": "#FanCollection.FanCollection",
	"Name": "Fan Collection",
	"Members": [
	  {
		"@odata.id": "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/Fans/Bay1"
	  }
	],
	"Members@odata.count": 1
  }`

// TestThermalSubsystemFans tests that the Fans link resolves to Fan resources
// (FanCollection), not to the legacy Thermal.Fan type.
func TestThermalSubsystemFans(t *testing.T) {
	var result ThermalSubsystem
	err := json.NewDecoder(strings.NewReader(thermalSubsystemBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				getCall(thermalSubsystemFanCollection),
				getCall(fanBody),
			},
		},
	}
	result.SetClient(testClient)

	fans, err := result.Fans()
	if err != nil {
		t.Fatalf("Error getting fans: %s", err)
	}

	if len(fans) != 1 {
		t.Fatalf("Expected 1 fan, got %d", len(fans))
	}

	fan := fans[0]
	if fan.ID != "Bay1" {
		t.Errorf("Received invalid fan ID: %s", fan.ID)
	}

	if fan.SpeedPercent.Reading == nil || *fan.SpeedPercent.Reading != 45 {
		t.Errorf("Received invalid fan speed reading: %v", fan.SpeedPercent.Reading)
	}

	if fan.SpeedPercent.SpeedRPM == nil || *fan.SpeedPercent.SpeedRPM != 2200 {
		t.Errorf("Received invalid fan speed RPM: %v", fan.SpeedPercent.SpeedRPM)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 || calls[0].URL != "/redfish/v1/Chassis/System.Embedded.1/ThermalSubsystem/Fans" {
		t.Errorf("Unexpected calls: %+v", calls)
	}
}

// TestThermalSubsystemNoFans tests that a subsystem without a Fans link
// returns no fans and makes no request.
func TestThermalSubsystemNoFans(t *testing.T) {
	var result ThermalSubsystem
	err := json.NewDecoder(strings.NewReader(`{"Id": "ThermalSubsystem", "Name": "Thermal Subsystem"}`)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	fans, err := result.Fans()
	if err != nil {
		t.Fatalf("Error getting fans: %s", err)
	}

	if fans != nil {
		t.Errorf("Expected no fans, got %d", len(fans))
	}

	if calls := testClient.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no requests, got %+v", calls)
	}
}
