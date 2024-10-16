//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var smcRAKPBody = `{
  "@odata.type": "#SMCRAKP.v1_0_0.SMCRAKP",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SMCRAKP",
  "Name": "SMCRAKP",
  "Id": "SMC RAKP",
  "Mode": "Disabled",
  "@odata.etag": "\"2e678283a52f99ccf06bc13d8434178a\""
}`

// TestSMCRAKP tests the parsing of SMCRAKP objects.
func TestSMCRAKP(t *testing.T) {
	var result SMCRAKP
	err := json.NewDecoder(strings.NewReader(smcRAKPBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SMC RAKP" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Mode != "Disabled" {
		t.Errorf("Invalid mode: %s", result.Mode)
	}
}
