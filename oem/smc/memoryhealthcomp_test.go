//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var memoryHealthCompBody = `{
  "@odata.type": "#MemoryHealthComp.v1_0_0.MemoryHealthComp",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MemoryHealthComp",
  "Id": "MemoryHealthComp",
  "Name": "MemoryHealthComp",
  "MemoryHealthCompInit": "Disable",
  "MemoryHealthCompNext": "Disable",
  "@odata.etag": "\"af1226ca6b365ca917866f09f3a6b846\""
}`

// TestMemoryHealthComp tests the parsing of MemoryHealthComp objects.
func TestMemoryHealthComp(t *testing.T) {
	var result MemoryHealthComp
	err := json.NewDecoder(strings.NewReader(memoryHealthCompBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "MemoryHealthComp" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Init != "Disable" {
		t.Errorf("Invalid init state: %s", result.Init)
	}

	if result.Next != "Disable" {
		t.Errorf("Invalid next state: %s", result.Next)
	}
}
