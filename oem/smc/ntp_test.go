//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var ntpBody = `{
  "@odata.type": "#NTP.v1_0_3.NTP",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/NTP",
  "Id": "NTP",
  "Name": "NTP Service",
  "NTPEnable": true,
  "PrimaryNTPServer": "localhost",
  "SecondaryNTPServer": "127.0.0.1",
  "DaylightSavingTime": true,
  "@odata.etag": "\"6002d9d6874d76983f5cfb025da6fd57\""
}`

// TestNTP tests the parsing of NTP objects.
func TestNTP(t *testing.T) {
	var result NTP
	err := json.NewDecoder(strings.NewReader(ntpBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.Name != "NTP Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.Enabled {
		t.Errorf("Invalid enabled state: %t", result.Enabled)
	}

	if !result.DaylightSavingTime {
		t.Errorf("Invalid daylight savings time state: %t", result.DaylightSavingTime)
	}

	if result.PrimaryServer != "localhost" {
		t.Errorf("Invalid server: %s", result.PrimaryServer)
	}

	if result.SecondaryServer != "127.0.0.1" {
		t.Errorf("Invalid server: %s", result.SecondaryServer)
	}
}
