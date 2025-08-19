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

var powerSupplyUnitMetricsBody = strings.NewReader(
	`{
		"@odata.type": "#PowerSupplyMetrics.v1_1_0.PowerSupplyMetrics",
		"Id": "Metrics",
		"Name": "Metrics for Power Supply 1",
		"Status": {
			"State": "Enabled",
			"Health": "Warning"
		},
		"InputVoltage": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1InputVoltage",
			"Reading": 230.2
		},
		"InputCurrentAmps": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1InputCurrent",
			"Reading": 5.19
		},
		"InputPowerWatts": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1InputPower",
			"Reading": 937.4
		},
		"RailVoltage": [
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_3VOutput",
				"Reading": 3.31
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_5VOutput",
				"Reading": 5.03
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_12VOutput",
				"Reading": 12.06
			}
		],
		"RailCurrentAmps": [
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_3VCurrent",
				"Reading": 9.84
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_5VCurrent",
				"Reading": 1.25
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_12Current",
				"Reading": 2.58
			}
		],
		"OutputPowerWatts": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1OutputPower",
			"Reading": 937.4
		},
		"RailPowerWatts": [
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_3VPower",
				"Reading": 79.84
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_5VPower",
				"Reading": 26.25
			},
			{
				"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1_12VPower",
				"Reading": 91.58
			}
		],
		"EnergykWh": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1Energy",
			"Reading": 325675
		},
		"FrequencyHz": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1InputFrequency",
			"Reading": 60
		},
		"TemperatureCelsius": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1Temp",
			"Reading": 43.9
		},
		"FanSpeedPercent": {
			"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PS1Fan",
			"Reading": 68,
			"SpeedRPM": 3290
		},
		"Actions": {
			"#PowerSupplyMetrics.ResetMetrics": {
				"target": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Metrics/PowerSupplyMetrics.ResetMetrics"
			}
		},
		"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/PowerSupplies/Bay1/Metrics"
	}`)

// TestPowerDistributionMetrics tests the parsing of PowerDistributionMetrics objects.
func TestPowerSupplyUnitMetrics(t *testing.T) {
	var result PowerSupplyUnitMetrics
	err := json.NewDecoder(powerSupplyUnitMetricsBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Metrics for Power Supply 1", result.Name)
	assertEquals(t, "5.19", fmt.Sprint(*result.InputCurrentAmps.Reading))
	assertEquals(t, "937.4", fmt.Sprint(*result.InputPowerWatts.Reading))
	assertEquals(t, "230.2", fmt.Sprint(*result.InputVoltage.Reading))
	assertEquals(t, "60", fmt.Sprint(*result.FrequencyHz.Reading))
	assertEquals(t, "68", fmt.Sprint(result.FanSpeedPercent.Reading))
	assertEquals(t, "3290", fmt.Sprint(result.FanSpeedPercent.SpeedRPM))
	assertEquals(t, "79.84", fmt.Sprint(*result.RailPowerWatts[0].Reading))
	assertEquals(t, "1.25", fmt.Sprint(*result.RailCurrentAmps[1].Reading))
}
