//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var accelerationFunctionBody = strings.NewReader(
	`{
		"@odata.type": "#AccelerationFunction.v1_0_4.AccelerationFunction",
		"Id": "Compression",
		"Name": "Compression Accelerator",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"FpgaReconfigurationSlots": [
			"AFU0"
		],
		"AccelerationFunctionType": "Compression",
		"Manufacturer": "Intel (R) Corporation",
		"Version": "Green Compression Type 1 v.1.00.86",
		"PowerWatts": 15,
		"Links": {
			"Endpoints": [],
			"PCIeFunctions": []
		},
		"@odata.id": "/redfish/v1/Systems/1/Processors/FPGA1/AccelerationFunctions/Compression"
	}`)

// TestAccerlationFunction tests the parsing of AccerlationFunction objects.
func TestAccerlationFunction(t *testing.T) {
	var result AccelerationFunction
	err := json.NewDecoder(accelerationFunctionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Compression", result.ID)
	assertEquals(t, "Compression Accelerator", result.Name)
	assertEquals(t, "Compression", string(result.AccelerationFunctionType))

	if len(result.FPGAReconfigurationSlots) != 1 {
		t.Errorf("Expected 1 FPGA reconfiguration slots, got %#v", result.FPGAReconfigurationSlots)
	}
}
