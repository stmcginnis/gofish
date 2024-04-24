//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var heaterBody = `{
	"@odata.type": "#Heater.v1_0_1.Heater",
	"Id": "CPU1Heater",
	"Description": "Heater for CPU1",
	"Name": "Heater 1",
	"PhysicalContext": "CPU",
	"Manufacturer": "Contoso Heaters",
	"Model": "CPUHeater",
	"SerialNumber": "SNDHM0123456789",
	"PartNumber": "12345-123",
	"SparePartNumber": "54321-321",
	"LocationIndicatorActive": false,
	"HotPluggable": true,
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"Links": {
	  "Processors": [
		{
		  "@odata.id": "/redfish/v1/Systems/437XR1138R2/Processors/CPU1"
		}
	  ]
	},
	"Metrics": {
	  "@odata.id": "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater/Metrics"
	},
	"@odata.id": "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater"
  }`

// TestHeater tests the parsing of Heater objects.
func TestHeater(t *testing.T) {
	var result Heater
	err := json.NewDecoder(strings.NewReader(heaterBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "CPU1Heater", result.ID)
	assertEquals(t, "Heater 1", result.Name)
	assertEquals(t, "CPU", string(result.PhysicalContext))
	assertEquals(t, "/redfish/v1/Systems/437XR1138R2/Processors/CPU1", result.processors[0])
	assertEquals(t, "/redfish/v1/Chassis/1U/ThermalSubsystem/Heaters/CPU1Heater/Metrics", result.metrics)

	if result.LocationIndicatorActive {
		t.Error("Expected location indicator not to be active")
	}

	if !result.HotPluggable {
		t.Error("Expected to be hot pluggable")
	}
}
