//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var cxlLogicalDeviceBody = `{
	"@odata.type": "#CXLLogicalDevice.v1_1_1.CXLLogicalDevice",
	"Id": "1",
	"Name": "CXL Logical Device Type 1",
	"Description": "Locally attached CXL Logical Device Type 1",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK",
	  "HealthRollup": "OK"
	},
	"Identifiers": [
	  {
		"DurableName": "4C-1D-96-FF-FE-DD-D8-35:0001",
		"DurableNameFormat": "GCXLID"
	  }
	],
	"SemanticsSupported": [
	  "CXLio",
	  "CXLcache"
	],
	"Links": {
	  "PCIeFunctions": [
		{
		  "@odata.id": "/redfish/v1/Chassis/CXL1/PCIeDevices/1/PCIeFunctions/1"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Chassis/CXL1/PCIeDevices/1/CXLLogicalDevices/1"
  }`

// TestCxlLogicalDevice tests the parsing of CxlLogicalDevice objects.
func TestCxlLogicalDevice(t *testing.T) {
	var result CXLLogicalDevice
	err := json.NewDecoder(strings.NewReader(cxlLogicalDeviceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "CXL Logical Device Type 1", result.Name)
	assertEquals(t, "Locally attached CXL Logical Device Type 1", result.Description)
	assertEquals(t, "4C-1D-96-FF-FE-DD-D8-35:0001", result.Identifiers[0].DurableName)
	assertEquals(t, "GCXLID", string(result.Identifiers[0].DurableNameFormat))
	assertEquals(t, "CXLio", string(result.SemanticsSupported[0]))
	assertEquals(t, "CXLcache", string(result.SemanticsSupported[1]))
	assertEquals(t, "/redfish/v1/Chassis/CXL1/PCIeDevices/1/PCIeFunctions/1", result.pcieFunctions[0])
}
