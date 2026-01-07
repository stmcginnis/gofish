//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var iKVMBody = `{
  "@odata.type": "#IKVM.v1_0_2.IKVM",
  "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IKVM",
  "Id": "IKVM",
  "Name": "IKVM",
  "Current interface": "HTML 5",
  "URI": "/schemas.VVGBkzp32dNpSOI.IKVM",
  "@odata.etag": "\"916e58e6235aa0579147a3114560a596\""
}`

// TestIKVM tests the parsing of IKVM objects.
func TestIKVM(t *testing.T) {
	var result IKVM
	err := json.NewDecoder(strings.NewReader(iKVMBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IKVM" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.CurrentInterface != "HTML 5" {
		t.Errorf("Invalid current interface: %s", result.CurrentInterface)
	}

	if result.URI != "/schemas.VVGBkzp32dNpSOI.IKVM" {
		t.Errorf("Invalid URI: %s", result.URI)
	}
}
