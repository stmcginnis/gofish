//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var environmentMetricsBody = `{
	"@odata.type": "#EnvironmentMetrics.v1_3_1.EnvironmentMetrics",
	"ID": "Metrics1",
	"Name": "Processor Environment Metrics",
	"TemperatureCelsius": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/CPU1Temp",
	  "Reading": 44
	},
	"PowerWatts": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/CPU1Power",
	  "Reading": 12.87
	},
	"FanSpeedsPercent": [
	  {
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/CPUFan1",
		"DeviceName": "CPU #1 Fan Speed",
		"Reading": 80
	  }
	],
	"@odata.id": "/redfish/v1/Systems/437XR1138R2/Processors/1/EnvironmentMetrics"
  }`

// TestEnvironmentMetrics tests the parsing of EnvironmentMetrics objects.
func TestEnvironmentMetrics(t *testing.T) {
	var result EnvironmentMetrics
	err := json.NewDecoder(strings.NewReader(environmentMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics1", result.ID)
	assertEquals(t, "Processor Environment Metrics", result.Name)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/CPU1Temp", result.TemperatureCelsius.DataSourceURI)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/CPU1Power", result.PowerWatts.DataSourceURI)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/CPUFan1", result.FanSpeedsPercent[0].DataSourceURI)
	assertEquals(t, "CPU #1 Fan Speed", result.FanSpeedsPercent[0].DeviceName)

	if *result.PowerWatts.Reading != 12.87 {
		t.Errorf("Unexpected PowerWatts reading: %.2f", *result.PowerWatts.Reading)
	}
}
