//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var syslogBody = `{
  "@odata.type": "#Syslog.v1_0_1.Syslog",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog",
  "Id": "Syslog",
  "Name": "Syslog",
  "EnableSyslog": true,
  "SyslogServer": "localhost",
  "SyslogPortNumber": 514,
  "@odata.etag": "\"b27af6393687bb1810b00fe52874e053\""
}`

// TestSyslog tests the parsing of Syslog objects.
func TestSyslog(t *testing.T) {
	var result Syslog
	err := json.NewDecoder(strings.NewReader(syslogBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Syslog" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Errorf("Invalid enable state: %t", result.Enabled)
	}

	if result.Server != "localhost" {
		t.Errorf("Invalid server: %s", result.Server)
	}

	if result.Port != 514 {
		t.Errorf("Invalid port: %d", result.Port)
	}
}
