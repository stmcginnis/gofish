//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var fanBody = `{
	"@odata.type": "#Fan.v1_5_1.Fan",
	"Id": "Bay1",
	"Name": "Fan Bay 1",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"PhysicalContext": "Chassis",
	"Model": "RKS-440DC",
	"Manufacturer": "Contoso Fans",
	"PartNumber": "23456-133",
	"SparePartNumber": "93284-133",
	"LocationIndicatorActive": true,
	"HotPluggable": true,
	"SpeedPercent": {
	  "Reading": 45,
	  "SpeedRPM": 2200,
	  "DataSourceUri": "/redfish/v1/Chassis/1U/Sensors/FanBay1"
	},
	"Location": {
	  "PartLocation": {
		"ServiceLabel": "Chassis Fan Bay 1",
		"LocationType": "Bay",
		"LocationOrdinalValue": 0
	  }
	},
	"@odata.id": "/redfish/v1/Chassis/1U/ThermalSubsystem/Fans/Bay1"
  }`

// TestFan tests the parsing of Fan objects.
func TestFan(t *testing.T) {
	var result Fan
	err := json.NewDecoder(strings.NewReader(fanBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Bay1", result.ID)
	assertEquals(t, "Fan Bay 1", result.Name)
	assertEquals(t, "Chassis", string(result.PhysicalContext))
	assertEquals(t, "RKS-440DC", result.Model)
	assertEquals(t, "Chassis Fan Bay 1", result.Location.PartLocation.ServiceLabel)
	assertEquals(t, "/redfish/v1/Chassis/1U/Sensors/FanBay1", result.SpeedPercent.DataSourceURI)

	if !result.LocationIndicatorActive {
		t.Error("Expected location indicator to be active")
	}

	if !result.HotPluggable {
		t.Error("Expected to be hot pluggable")
	}

	if *result.SpeedPercent.SpeedRPM != 2200 {
		t.Errorf("Unexpected SpeedPercent.SpeedRPM: %.2f", *result.SpeedPercent.SpeedRPM)
	}
}
