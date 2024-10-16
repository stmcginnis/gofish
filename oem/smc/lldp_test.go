//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var lldpBody = `{
  "@odata.type": "#LldpService.v1_0_0.LldpService",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/LLDP",
  "Id": "LLDP",
  "Name": "LLDP",
  "Description": "LLDP Service",
  "LLDPEnabled": true,
  "LLDPReceive": {},
  "@odata.etag": "\"1f8f5d90dca8588b586dbe4ba3657003\""
}`

// TestLLDP tests the parsing of LLDP objects.
func TestLLDP(t *testing.T) {
	var result LLDP
	err := json.NewDecoder(strings.NewReader(lldpBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "LLDP" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if !result.Enabled {
		t.Error("LLDP is not enabled")
	}
}
