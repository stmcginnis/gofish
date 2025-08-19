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

var leakDetectionBody = `{
	"@odata.type": "#LeakDetection.v1_0_0.LeakDetection",
	"Id": "LeakDetection",
	"Name": "Leak Detection Systems",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK",
	  "Conditions": []
	},
	"LeakDetectorGroups": [
	  {
		"GroupName": "Detectors under and around the CDU",
		"HumidityPercent": {
		  "Reading": 45
		},
		"Detectors": [
		  {
			"DataSourceUri": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors/Moisture",
			"DeviceName": "Moisture-type Leak Detector",
			"DetectorState": "OK"
		  },
		  {
			"DeviceName": "Leak Detection Rope 1",
			"DetectorState": "OK"
		  },
		  {
			"DataSourceUri": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors/Overflow",
			"DeviceName": "Overflow Float Switch",
			"DetectorState": "OK"
		  }
		]
	  }
	],
	"LeakDetectors": {
	  "@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors"
	},
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection"
  }`

// TestLeakDetection tests the parsing of LeakDetection objects.
func TestLeakDetection(t *testing.T) {
	var result LeakDetection
	err := json.NewDecoder(strings.NewReader(leakDetectionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "LeakDetection", result.ID)
	assertEquals(t, "Leak Detection Systems", result.Name)
	assertEquals(t, "Detectors under and around the CDU", result.LeakDetectorGroups[0].GroupName)
	assertEquals(t, "45", fmt.Sprintf("%.0f", *result.LeakDetectorGroups[0].HumidityPercent.Reading))
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors/Overflow", result.LeakDetectorGroups[0].Detectors[2].DataSourceURI)
	assertEquals(t, "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors", result.leakDetectors)
}
