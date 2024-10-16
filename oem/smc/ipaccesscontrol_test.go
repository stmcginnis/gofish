//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var ipAccessControlBody = `{
  "@odata.type": "#IPAccessControl.v1_0_1.IPAccessControl",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IPAccessControl",
  "Id": "IP Access Control",
  "Name": "IP Access Control",
  "ServiceEnabled": true,
  "FilterRules": {
    "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IPAccessControl/FilterRules"
  },
  "@odata.etag": "\"ecbaa3f8ca55261d32edce302b8e3ddd\""
}`

// TestIPAccessControl tests the parsing of IPAccessControl objects.
func TestIPAccessControl(t *testing.T) {
	var result IPAccessControl
	err := json.NewDecoder(strings.NewReader(ipAccessControlBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IP Access Control" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}

	if result.filterRules != "/redfish/v1/Managers/1/Oem/Supermicro/IPAccessControl/FilterRules" {
		t.Errorf("Invalid filter rule link: %s", result.filterRules)
	}
}
