//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var kcsInterfaceBody = `{
  "@odata.type": "#KCSInterface.v1_1_0.KCSInterface",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/KCSInterface",
  "Id": "KCS Interface",
  "Name": "KCS Interface",
  "Privilege": "Administrator",
  "@odata.etag": "\"baa7b14122a605d1a202bd7fa12b8125\""
}`

// TestKCSInterface tests the parsing of KCSInterface objects.
func TestKCSInterface(t *testing.T) {
	var result KCSInterface
	err := json.NewDecoder(strings.NewReader(kcsInterfaceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "KCS Interface" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Privilege != "Administrator" {
		t.Errorf("Invalid privilege: %s", result.Privilege)
	}
}
