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

var sensorBody = strings.NewReader(
	`{
		"@odata.type": "#Sensor.v1_7_0.Sensor",
		"Id": "CabinetTemp",
		"Name": "Rack Temperature",
		"ReadingType": "Temperature",
		"ReadingTime": "2019-12-25T04:14:33+06:00",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Reading": 31.6,
		"ReadingUnits": "C",
		"ReadingRangeMin": -1.7976931348623157e+308,
		"ReadingRangeMax": 1.7976931348623157e+308,
		"Accuracy": 0.25,
		"Precision": 1,
		"SensingInterval": "PT3S",
		"PhysicalContext": "Chassis",
		"Thresholds": {
			"UpperCritical": {
				"Reading": 40,
				"Activation": "Increasing"
			},
			"UpperCaution": {
				"Reading": 35,
				"Activation": "Decreasing"
			},
			"LowerCaution": {
				"Reading": 10,
				"Activation": "Increasing"
			}
		},
		"@odata.id": "/redfish/v1/Chassis/1/Sensors/CabinetTemp"
	}`)

// TestSensor tests the parsing of Sensor objects.
func TestSensor(t *testing.T) {
	var result Sensor

	if err := json.NewDecoder(sensorBody).Decode(&result); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "CabinetTemp", result.ID)
	assertEquals(t, "Rack Temperature", result.Name)
	assertEquals(t, "Temperature", string(result.ReadingType))
	assertEquals(t, "2019-12-25T04:14:33+06:00", result.ReadingTime)
	assertEquals(t, "31.6", fmt.Sprint(result.Reading))
	assertEquals(t, "C", result.ReadingUnits)
	assertEquals(t, "Decreasing", string(result.Thresholds.UpperCaution.Activation))
}
