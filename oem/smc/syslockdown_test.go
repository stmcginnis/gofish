//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var sysLockdownBody = `{
  "@odata.type": "#SysLockdown.v1_0_0.SysLockdown",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SysLockdown",
  "Id": "SysLockdown",
  "Name": "SysLockdown",
  "SysLockdownEnabled": true,
  "@odata.etag": "\"8374ed52bcd7c7e92902143a75345193\""
}`

// TestSysLockdown tests the parsing of SysLockdown objects.
func TestSysLockdown(t *testing.T) {
	var result SysLockdown
	err := json.NewDecoder(strings.NewReader(sysLockdownBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SysLockdown" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}
}
