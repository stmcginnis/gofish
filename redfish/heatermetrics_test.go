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

var heaterMetricsBody = `{
	"@odata.type": "#HeaterMetrics.v1_0_1.HeaterMetrics",
	"Id": "HeaterMetrics",
	"Description": "Heater Metrics for CPU1 Heater",
	"Name": "CPU1 Heater Metrics",
	"PrePowerOnHeatingTimeSeconds": 600,
	"RuntimeHeatingTimeSeconds": 3600,
	"PowerWatts": {
	  "Reading": 200.3
	},
	"TemperatureReadingsCelsius": [
	  {
		"DeviceName": "Heater Average Temperature",
		"Reading": 2.5
	  }
	],
	"Actions": {
	  "#HeaterMetrics.ResetMetrics": {
		"target": "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater/Metrics/HeaterMetrics.ResetMetrics"
	  }
	},
	"@odata.id": "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater/Metrics"
  }`

// TestHeaterMetrics tests the parsing of HeaterMetrics objects.
func TestHeaterMetrics(t *testing.T) {
	var result HeaterMetrics
	err := json.NewDecoder(strings.NewReader(heaterMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "HeaterMetrics", result.ID)
	assertEquals(t, "CPU1 Heater Metrics", result.Name)
	assertEquals(t, "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater/Metrics/HeaterMetrics.ResetMetrics", result.resetMetricsTarget)
	assertEquals(t, "Heater Average Temperature", result.TemperatureReadingsCelsius[0].DeviceName)

	if result.PrePowerOnHeatingTimeSeconds != 600 {
		t.Errorf("Unexpected PrePowerOnHeatingTimeSeconds: %d", result.PrePowerOnHeatingTimeSeconds)
	}

	if result.RuntimeHeatingTimeSeconds != 3600 {
		t.Errorf("Unexpected RuntimeHeatingTimeSeconds: %d", result.RuntimeHeatingTimeSeconds)
	}

	if *result.PowerWatts.Reading != 200.3 {
		t.Errorf("Unexpected RuntimeHeatingTimeSeconds: %.2f", *result.PowerWatts.Reading)
	}
}

// TestHeaterMetricsResetMetrics tests the HeaterMetrics ResetMetrics call.
func TestHeaterMetricsResetMetrics(t *testing.T) {
	var result HeaterMetrics
	err := json.NewDecoder(strings.NewReader(heaterMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.ResetMetrics()
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "") {
		t.Errorf("Expected payload to be empty: %s", calls[0].Payload)
	}
}
