//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"encoding/json"
	"fmt"
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

var actionInfoBody = `
{
  "@odata.etag": "\"0515ed61d44ec38b5e73b1c6d5b0ea8d\"",
  "@odata.id": "/redfish/v1/Managers/1/ResetToDefaultsActionInfo",
  "@odata.type": "#ActionInfo.v1_1_2.ActionInfo",
  "Id": "ResetToDefaultsActionInfo",
  "Name": "Factory Default Action Info",
  "Parameters": [%s]
}
`

func TestActionInfo(t *testing.T) {
	for _, test := range []struct {
		name   string
		body   string
		params string
		err    string
	}{
		{body: "{}", err: "ActionInfo required field Name is missing"},
		{body: `{"ID":"ResetToDefaultsActionInfo"}`, err: "ActionInfo required field Name is missing"},
		{body: `{"ID":"ResetToDefaultsActionInfo","Name":"Factory Default Action Info"}`},
		{params: `{}`, err: "Parameter required field Name is missing"},
		{params: `{"Name":"foo","DataType":"NonExistentDataType"}`, err: "unknown DataType specified: NonExistentDataType"},
		{params: `{"Name":"foo","DataType":"String"}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			body := test.body
			if test.params != "" {
				body = fmt.Sprintf(actionInfoBody, test.params)
			}

			var result ActionInfo
			err := json.NewDecoder(strings.NewReader(body)).Decode(&result)
			if test.err != "" {
				if err.Error() != test.err {
					t.Errorf("Did not get expected error, want: %s, got: %s", test.err, err.Error())
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error decoding JSON: %s", err)
				return
			}

			if result.ID != "ResetToDefaultsActionInfo" {
				t.Errorf("Received invalid ID: %s", result.ID)
			}

			if result.Name != "Factory Default Action Info" {
				t.Errorf("Received invalid name: %s", result.Name)
			}
		})
	}
}
