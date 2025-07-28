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

var powerDistributionMetricsBody = strings.NewReader(
	`{
		"@odata.type": "#PowerDistributionMetrics.v1_3_0.PowerDistributionMetrics",
		"Id": "Metrics",
		"Name": "Summary Metrics",
		"PowerWatts": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PDUPower",
			"Reading": 6438,
			"ApparentVA": 6300,
			"ReactiveVAR": 100,
			"PowerFactor": 0.93
		},
		"EnergykWh": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PDUEnergy",
			"Reading": 56438
		},
		"TemperatureCelsius": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PDUTemp",
			"Reading": 26.3
		},
		"HumidityPercent": {
			"DataSourceUri": "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PDUHumidity",
			"Reading": 52.7
		},
		"Actions": {
			"#PowerDistributionMetrics.ResetMetrics": {
				"target": "/redfish/v1/PowerEquipment/RackPDUs/1/Metrics/PowerDistributionMetrics.ResetMetrics"
			}
		},
		"@odata.id": "/redfish/v1/PowerEquipment/RackPDUs/1/Metrics"
	}`)

// TestPowerDistributionMetrics tests the parsing of PowerDistributionMetrics objects.
func TestPowerDistributionMetrics(t *testing.T) {
	var result PowerDistributionMetrics
	err := json.NewDecoder(powerDistributionMetricsBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Metrics", result.ID)
	assertEquals(t, "Summary Metrics", result.Name)
	assertEquals(t, "/redfish/v1/PowerEquipment/RackPDUs/1/Sensors/PDUPower", result.PowerWatts.DataSourceURI)
	assertEquals(t, "6438", fmt.Sprint(*result.PowerWatts.Reading))
	assertEquals(t, "0.93", fmt.Sprint(*result.PowerWatts.PowerFactor))
	assertEquals(t, "52.7", fmt.Sprint(*result.HumidityPercent.Reading))
}
