//
// SPDX-License-Identifier: BSD-3-Clause
//

package hpe

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var hpeThermalBody = `{
    "@odata.context": "/redfish/v1/$metadata#Thermal.Thermal",
    "@odata.etag": "W/\"AEC0B081\"",
    "@odata.id": "/redfish/v1/Chassis/1/Thermal/",
    "@odata.type": "#Thermal.v1_1_0.Thermal",
    "Id": "Thermal",
    "Fans": [
        {
            "@odata.id": "/redfish/v1/Chassis/1/Thermal/#Fans/0",
            "MemberId": "0",
            "Name": "Fan 1",
            "Oem": {
                "Hpe": {
                    "@odata.context": "/redfish/v1/$metadata#HpeServerFan.HpeServerFan",
                    "@odata.type": "#HpeServerFan.v2_0_0.HpeServerFan",
                    "HotPluggable": true,
                    "Location": "System",
                    "Redundant": true
                }
            },
            "Reading": 28,
            "ReadingUnits": "Percent",
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            }
        }
    ],
    "Name": "Thermal",
    "Oem": {
        "Hpe": {
            "@odata.context": "/redfish/v1/$metadata#HpeThermalExt.HpeThermalExt",
            "@odata.type": "#HpeThermalExt.v2_0_0.HpeThermalExt",
            "FanPercentMinimum": 15,
            "ThermalConfiguration": "OptimalCooling"
        }
    },
    "Temperatures": [
        {
            "@odata.id": "/redfish/v1/Chassis/1/Thermal/#Temperatures/0",
            "MemberId": "0",
            "Name": "01-Inlet Ambient",
            "Oem": {
                "Hpe": {
                    "@odata.context": "/redfish/v1/$metadata#HpeSeaOfSensors.HpeSeaOfSensors",
                    "@odata.type": "#HpeSeaOfSensors.v2_0_0.HpeSeaOfSensors",
                    "LocationXmm": 15,
                    "LocationYmm": 0
                }
            },
            "PhysicalContext": "Intake",
            "ReadingCelsius": 25,
            "SensorNumber": 1,
            "Status": {
                "Health": "OK",
                "State": "Enabled"
            },
            "UpperThresholdCritical": 42,
            "UpperThresholdFatal": 47,
            "UpperThresholdUser": 0
        }
    ]
}`

// TestHpeThermalOem tests the parsing of Thermal objects and support oem field.
func TestHpeThermalOem(t *testing.T) {
	var thermal *schemas.Thermal
	err := json.NewDecoder(strings.NewReader(hpeThermalBody)).Decode(&thermal)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	hpeThermal, err := FromThermal(thermal)
	if err != nil {
		t.Errorf("Convert thermal failed: %s", err)
		return
	}

	if hpeThermal.Oem.Hpe.FanPercentMinimum != 15 {
		t.Errorf("Received invalid fan fan percent minimum: %d", hpeThermal.Oem.Hpe.FanPercentMinimum)
	}

	if hpeThermal.Oem.Hpe.ThermalConfiguration != "OptimalCooling" {
		t.Errorf("Received invalid hpeThermal configuration: %s", hpeThermal.Oem.Hpe.ThermalConfiguration)
	}

	if hpeThermal.Fans[0].Oem.Hpe.Location != "System" {
		t.Errorf("Received invalid location: %s", hpeThermal.Fans[0].Oem.Hpe.Location)
	}
}
