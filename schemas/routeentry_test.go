//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var routeEntryBody = `{
	"@odata.type": "#RouteEntry.v1_0_1.RouteEntry",
	"Id": "0",
	"Name": "LPRT0",
	"Description": "Gen-Z Port 1 LPRT Entry 0",
	"RawEntryHex": "0x34EF124500000000",
	"RouteSet": {
	  "@odata.id": "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/LPRT/0/RouteSet"
	},
	"MinimumHopCount": 1,
	"@odata.id": "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/LPRT/0"
  }`

// TestRouteEntry tests the parsing of RouteEntry objects.
func TestRouteEntry(t *testing.T) {
	var result RouteEntry
	err := json.NewDecoder(strings.NewReader(routeEntryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "LPRT0", result.Name)
	assertEquals(t, "0x34EF124500000000", result.RawEntryHex)
	assertEquals(t, "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/LPRT/0/RouteSet", result.routeSet)

	if result.MinimumHopCount != 1 {
		t.Errorf("Unexpected MinimumHopCount: %d", result.MinimumHopCount)
	}
}
