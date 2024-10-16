//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var fanModeBody = `{
  "@odata.type": "#FanMode.v1_0_1.FanMode",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/FanMode",
  "Name": "FanMode",
  "Id": "Fan Mode",
  "Mode": "HeavyIO",
  "Mode@Redfish.AllowableValues": [
    "FullSpeed",
    "Optimal",
    "HeavyIO"
  ],
  "@odata.etag": "\"753bafbafcb8047326ec2c269ad52111\""
}`

// TestFanMode tests the parsing of FanMode objects.
func TestFanMode(t *testing.T) {
	var result FanMode
	err := json.NewDecoder(strings.NewReader(fanModeBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Fan Mode" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Mode != "HeavyIO" {
		t.Errorf("Invalid fan mode: %s", result.Mode)
	}

	if len(result.AllowableModes) != 3 {
		t.Errorf("Unexpected allowable modes: %v", result.AllowableModes)
	}
}
