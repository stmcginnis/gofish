//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var routeSetEntryBody = `{
	"@odata.type": "#RouteSetEntry.v1_0_1.RouteSetEntry",
	"Id": "0",
	"Name": "RouteSet0",
	"Description": "Gen-Z Port 1 LPRT Entry 0 Route 0",
	"Valid": false,
	"VCAction": 1,
	"HopCount": 2,
	"EgressIdentifier": 0,
	"@odata.id": "/redfish/v1/Fabrics/GenZ/Switches/Switch1/Ports/1/LPRT/0/RouteSet/0"
  }`

// TestRouteSetEntry tests the parsing of RouteSetEntry objects.
func TestRouteSetEntry(t *testing.T) {
	var result RouteSetEntry
	err := json.NewDecoder(strings.NewReader(routeSetEntryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "RouteSet0", result.Name)

	if result.Valid {
		t.Error("Expected Valid to be false")
	}

	if result.VCAction != 1 {
		t.Errorf("Unexpected VCAction: %d", result.VCAction)
	}

	if result.HopCount != 2 {
		t.Errorf("Unexpected HopCount: %d", result.HopCount)
	}
}
