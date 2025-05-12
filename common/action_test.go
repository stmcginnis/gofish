//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"strings"
	"testing"
)

var actionBody = `
{
  "#Manager.Reset": {
    "ResetType@Redfish.AllowableValues": [
      "GracefulRestart",
      "ForceRestart"
    ],
    "target": "/redfish/v1/Managers/1/Actions/Manager.Reset"
  },
  "#Manager.ResetToDefaults": {
    "@Redfish.ActionInfo": "/redfish/v1/Managers/1/Actions/ResetToDefaultsActionInfo",
    "target": "/redfish/v1/Managers/1/Actions/Manager.ResetToDefaults"
  }
}
`

func TestAction(t *testing.T) {
	var result map[string]Action
	err := json.NewDecoder(strings.NewReader(actionBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	r := result["#Manager.ResetToDefaults"]

	if r.Target != "/redfish/v1/Managers/1/Actions/Manager.ResetToDefaults" {
		t.Errorf("Received invalid ResetToDefaults Action target: %s", r.Target)
	}

	if r.Info != "/redfish/v1/Managers/1/Actions/ResetToDefaultsActionInfo" {
		t.Errorf("Received invalid ResetToDefaultsActionInfo target: %s", r.Info)
	}
}
