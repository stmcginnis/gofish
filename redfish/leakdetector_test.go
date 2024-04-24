//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var leakDetectorBody = `{
	"@odata.type": "#LeakDetector.v1_0_1.LeakDetector",
	"Id": "Moisture",
	"Name": "Moisture-type Leak Detector",
	"LeakDetectorType": "Moisture",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"DetectorState": "OK",
	"PartNumber": "3493-A44",
	"SerialNumber": "916239",
	"Manufacturer": "Contoso Water Detection Systems",
	"Model": "Depends 3000",
	"Location": {
	  "PartLocation": {
		"Reference": "Bottom",
		"ServiceLabel": "Leak Detector"
	  }
	},
	"PhysicalContext": "Chassis",
	"@odata.id": "/redfish/v1/ThermalEquipment/CDUs/1/LeakDetection/LeakDetectors/Moisture"
  }`

// TestLeakDetector tests the parsing of LeakDetector objects.
func TestLeakDetector(t *testing.T) {
	var result LeakDetector
	err := json.NewDecoder(strings.NewReader(leakDetectorBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Moisture", result.ID)
	assertEquals(t, "Moisture-type Leak Detector", result.Name)
	assertEquals(t, "Moisture", string(result.LeakDetectorType))
	assertEquals(t, "OK", string(result.DetectorState))
	assertEquals(t, "Bottom", string(result.Location.PartLocation.Reference))
	assertEquals(t, "Leak Detector", result.Location.PartLocation.ServiceLabel)
	assertEquals(t, "Chassis", string(result.PhysicalContext))
}
