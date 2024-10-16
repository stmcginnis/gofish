//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryPFABody = `{
  "@odata.type": "#MemoryPFA.v1_0_0.MemoryPFA",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MemoryPFA",
  "Id": "MemoryPFA",
  "Name": "MemoryPFA",
  "MemoryPfaInit": "Disabled",
  "MemoryPfaNext": "Disabled",
  "AlertId": 10,
  "@odata.etag": "\"bb528176e3ac4b740b445ae8aa43f1f1\""
}`

// TestMemoryPFA tests the parsing of MemoryPFA objects.
func TestMemoryPFA(t *testing.T) {
	var result MemoryPFA
	err := json.NewDecoder(strings.NewReader(memoryPFABody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "MemoryPFA" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Init != "Disabled" {
		t.Errorf("Invalid init state: %s", result.Init)
	}

	if result.Next != "Disabled" {
		t.Errorf("Invalid next state: %s", result.Next)
	}

	if result.AlertID != 10 {
		t.Errorf("Invalid alert ID: %d", result.AlertID)
	}
}
