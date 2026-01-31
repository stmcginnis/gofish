//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var thermalMetricsBody = `{
	"@odata.type": "#ThermalMetrics.v1_3_1.ThermalMetrics",
	"Id": "ThermalMetrics",
	"Name": "Chassis Thermal Metrics",
	"TemperatureSummaryCelsius": {
	  "Internal": {
		"Reading": 39,
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/CPU1Temp"
	  },
	  "Intake": {
		"Reading": 24.8,
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/IntakeTemp"
	  },
	  "Ambient": {
		"Reading": 22.5,
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/AmbientTemp"
	  },
	  "Exhaust": {
		"Reading": 40.5,
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/ExhaustTemp"
	  }
	},
	"PowerWatts": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/FanTotalPower",
	  "Reading": 24.72
	},
	"EnergykWh": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/FanTotalEnergy",
	  "Reading": 38.84
	},
	"TemperatureReadingsCelsius": [
	  {
		"Reading": 40,
		"DeviceName": "SystemBoard",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/SysBrdTemp"
	  },
	  {
		"Reading": 24.8,
		"DeviceName": "Intake",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/IntakeTemp"
	  },
	  {
		"Reading": 39,
		"DeviceName": "CPUSubsystem",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/CPUTemps"
	  },
	  {
		"Reading": 42,
		"DeviceName": "MemorySubsystem",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/MemoryTemp"
	  },
	  {
		"Reading": 33,
		"DeviceName": "PowerSupply",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/PSTemp"
	  },
	  {
		"Reading": 40.5,
		"DeviceName": "Exhaust",
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/ExhaustTemp"
	  }
	],
	"@odata.id": "/redfish/v1/Chassis/1U/ThermalSubsystem/ThermalMetrics"
  }`

// TestThermalMetrics tests the parsing of ThermalMetrics objects.
func TestThermalMetrics(t *testing.T) {
	var result ThermalMetrics
	err := json.NewDecoder(strings.NewReader(thermalMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ThermalMetrics", result.ID)
	assertEquals(t, "Chassis Thermal Metrics", result.Name)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/CPU1Temp", result.TemperatureSummaryCelsius.Internal.DataSourceURI)
	assertEquals(t, "SystemBoard", result.TemperatureReadingsCelsius[0].DeviceName)

	if *result.PowerWatts.Reading != 24.72 {
		t.Errorf("Unexpected PowerWatts.Reading value: %.2f", *result.PowerWatts.Reading)
	}
}
