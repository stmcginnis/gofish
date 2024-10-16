//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var snoopingBody = `{
  "@odata.type": "#Snooping.v1_0_0.Snooping",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Snooping",
  "Name": "Snooping",
  "Id": "Snooping",
  "PostCode": "e3",
  "@odata.etag": "\"b6b45c5595e17a1a2c878fa9f47b05c6\""
}`

// TestSnooping tests the parsing of Snooping objects.
func TestSnooping(t *testing.T) {
	var result Snooping
	err := json.NewDecoder(strings.NewReader(snoopingBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Snooping" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.PostCode != "e3" {
		t.Errorf("Invalid post code: %s", result.PostCode)
	}
}
