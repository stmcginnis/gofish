//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"strings"
	"testing"
)

var supermicroSyslogBody = `{
    "@odata.type": "#Syslog.v1_0_1.Syslog",
    "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog",
    "Id": "Syslog",
    "Name": "Syslog",
    "EnableSyslog": true,
    "SyslogServer": "192.168.0.1",
    "SyslogPortNumber": 514
}`

func TestSupermicroSyslogOem(t *testing.T) {
	var syslog *Syslog
	err := json.NewDecoder(strings.NewReader(supermicroSyslogBody)).Decode(&syslog)

	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	if !syslog.EnableSyslog {
		t.Errorf("EnableSyslog should be true")
	}

	if syslog.SyslogServer != "192.168.0.1" {
		t.Errorf("Expected syslog server at 192.168.0.1, got %s", syslog.SyslogServer)
	}

	if syslog.SyslogPortNumber != 514 {
		t.Errorf("Expected syslog server port 514, got %d", syslog.SyslogPortNumber)
	}
}
