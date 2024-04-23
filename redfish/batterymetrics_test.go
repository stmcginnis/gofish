//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var batteryMetricsBody = `{
	"@odata.type": "#BatteryMetrics.v1_0_3.BatteryMetrics",
	"Id": "Metrics",
	"Name": "Metrics for Battery 1",
	"DischargeCycles": 8.67,
	"InputVoltage": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1InputVoltage",
	  "Reading": 12.22
	},
	"InputCurrentAmps": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1InputCurrent",
	  "Reading": 0
	},
	"OutputVoltages": [
	  {
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1OutputVoltage",
		"Reading": 12.22
	  }
	],
	"OutputCurrentAmps": [
	  {
		"DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1OutputCurrent",
		"Reading": 0
	  }
	],
	"StoredEnergyWattHours": {
	  "Reading": 19.41
	},
	"TemperatureCelsius": {
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/Battery1Temp",
	  "Reading": 33
	},
	"ChargePercent": {
	  "Reading": 100
	},
	"CellVoltages": [
	  {
		"Reading": 3.44
	  },
	  {
		"Reading": 3.45
	  },
	  {
		"Reading": 3.43
	  },
	  {
		"Reading": 3.43
	  },
	  {
		"Reading": 3.45
	  },
	  {
		"Reading": 3.44
	  },
	  {
		"Reading": 3.43
	  },
	  {
		"Reading": 3.44
	  }
	],
	"@odata.id": "/redfish/v1/Chassis/1U/PowerSubsystem/Batteries/Module1/Metrics"
  }`

// TestBatteryMetrics tests the parsing of BatteryMetrics objects.
func TestBatteryMetrics(t *testing.T) {
	var result BatteryMetrics
	err := json.NewDecoder(strings.NewReader(batteryMetricsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Metrics for Battery 1", result.Name)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/Battery1InputVoltage", result.InputVoltage.DataSourceURI)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/Battery1InputCurrent", result.InputCurrentAmps.DataSourceURI)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/Battery1OutputVoltage", result.OutputVoltages[0].DataSourceURI)

	if result.DischargeCycles != 8.67 {
		t.Errorf("Unexpected discharge cycles result: %.2f", result.DischargeCycles)
	}

	if result.OutputVoltages[0].Reading != 12.22 {
		t.Errorf("Unexpected output voltage reading: %.2f", result.OutputVoltages[0].Reading)
	}

	if result.StoredEnergyWattHours.Reading != 19.41 {
		t.Errorf("Unexpected stored energy watt hours: %.2f", result.StoredEnergyWattHours.Reading)
	}

	if len(result.CellVoltages) != 8 {
		t.Errorf("Expected 8 cell voltage readings, got %#v", result.CellVoltages)
	}
}
