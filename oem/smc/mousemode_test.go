//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var mouseModeBody = `{
  "@odata.type": "#MouseMode.v1_0_0.MouseMode",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MouseMode",
  "Name": "MouseMode",
  "Id": "Mouse Mode",
  "Mode": "Absolute",
  "Mode@Redfish.AllowableValues": [
    "Absolute",
    "Relative",
    "Single"
  ],
  "@odata.etag": "\"9dee13296f2c876a0733eec3371ad4f4\""
}`

// TestMouseMode tests the parsing of MouseMode objects.
func TestMouseMode(t *testing.T) {
	var result MouseMode
	err := json.NewDecoder(strings.NewReader(mouseModeBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Mouse Mode" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Mode != "Absolute" {
		t.Errorf("Invalid mode: %s", result.Mode)
	}
}
