//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var vcatEntryBody = `{
	"@odata.type": "#VCATEntry.v1_0_2.VCATEntry",
	"Id": "0",
	"Name": "VCAT Entry 0",
	"Description": "Gen-Z Port 1 Virtual Channel Action Table Entry 0",
	"RawEntryHex": "0x123456",
	"VCEntries": [
	  {
		"VCMask": "0x00000034",
		"Threshold": "0x12"
	  },
	  {
		"VCMask": "0x00000034",
		"Threshold": "0x12"
	  },
	  {
		"VCMask": "0x00000034",
		"Threshold": "0x12"
	  },
	  {
		"VCMask": "0x00000034",
		"Threshold": "0x12"
	  }
	],
	"@odata.id": "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/VCAT/0"
  }`

// TestVCATEntry tests the parsing of VCATEntry objects.
func TestVCATEntry(t *testing.T) {
	var result VCATEntry
	err := json.NewDecoder(strings.NewReader(vcatEntryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "0x123456", result.RawEntryHex)
	assertEquals(t, "0x12", result.VCEntries[1].Threshold)
	assertEquals(t, "0x00000034", result.VCEntries[3].VCMask)
}
