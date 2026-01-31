//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var operatingConfigBody = `{
	"@odata.type": "#OperatingConfig.v1_0_3.OperatingConfig",
	"Id": "0",
	"Name": "Processor Profile",
	"TotalAvailableCoreCount": 28,
	"TDPWatts": 150,
	"BaseSpeedMHz": 2500,
	"MaxSpeedMHz": 4100,
	"MaxJunctionTemperatureCelsius": 90,
	"TurboProfile": [
	  {
		"ActiveCoreCount": 2,
		"MaxSpeedMHz": 4100
	  },
	  {
		"ActiveCoreCount": 4,
		"MaxSpeedMHz": 4000
	  },
	  {
		"ActiveCoreCount": 8,
		"MaxSpeedMHz": 3800
	  },
	  {
		"ActiveCoreCount": 28,
		"MaxSpeedMHz": 3200
	  }
	],
	"BaseSpeedPrioritySettings": [
	  {
		"CoreCount": 8,
		"CoreIDs": [
		  0,
		  2,
		  3,
		  4,
		  5,
		  6,
		  7,
		  8
		],
		"BaseSpeedMHz": 2900
	  },
	  {
		"CoreCount": 20,
		"BaseSpeedMHz": 2200
	  }
	],
	"@odata.id": "/redfish/v1/Systems/operating-config-example/Processors/CPU1/OperatingConfigs/0"
  }`

// TestOperatingConfig tests the parsing of OperatingConfig objects.
func TestOperatingConfig(t *testing.T) {
	var result OperatingConfig
	err := json.NewDecoder(strings.NewReader(operatingConfigBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "Processor Profile", result.Name)

	if *result.TotalAvailableCoreCount != 28 {
		t.Errorf("Unexpected TotalAvailableCoreCount value: %d", result.TotalAvailableCoreCount)
	}

	if *result.MaxJunctionTemperatureCelsius != 90 {
		t.Errorf("Unexpected MaxJunctionTemperatureCelsius value: %d", result.MaxJunctionTemperatureCelsius)
	}

	if *result.TurboProfile[0].ActiveCoreCount != 2 {
		t.Errorf("Unexpected ActiveCoreCount value: %d", result.TurboProfile[0].ActiveCoreCount)
	}

	if *result.TurboProfile[2].MaxSpeedMHz != 3800 {
		t.Errorf("Unexpected MaxSpeedMHz value: %d", result.TurboProfile[2].MaxSpeedMHz)
	}

	if *result.BaseSpeedPrioritySettings[0].BaseSpeedMHz != 2900 {
		t.Errorf("Unexpected BaseSpeedMHz value: %d", result.BaseSpeedPrioritySettings[0].BaseSpeedMHz)
	}
}
