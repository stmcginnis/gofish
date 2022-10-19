//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"strings"
	"testing"
)

var supermicroNTPBody = `{
    "@odata.type": "#NTP.v1_0_1.NTP",
    "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/NTP",
    "Id": "NTP",
    "Name": "NTP Service",
    "NTPEnable": true,
    "PrimaryNTPServer": "192.168.0.1",
    "SecondaryNTPServer": "192.168.1.1",
    "DaylightSavingTime": false
}`

func TestSupermicroNTPOem(t *testing.T) {
	var ntp *NTP
	err := json.NewDecoder(strings.NewReader(supermicroNTPBody)).Decode(&ntp)

	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	if !ntp.NTPEnable {
		t.Errorf("NTPEnable should be true")
	}

	if ntp.PrimaryNTPServer != "192.168.0.1" {
		t.Errorf("Expected primary NTP server at 192.168.0.1, got %s", ntp.PrimaryNTPServer)
	}

	if ntp.SecondaryNTPServer != "192.168.1.1" {
		t.Errorf("Expected secondary NTP server at 192.168.1.1, got %s", ntp.SecondaryNTPServer)
	}

	if ntp.DaylightSavingTime {
		t.Errorf("DaylightSavingTime should be false")
	}
}
