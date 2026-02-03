//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var allowDenyBody = `{
		"@odata.type": "#AllowDeny.v1_0_2.AllowDeny",
		"Id": "AllowDeny Rule 1",
		"Name": "Allow Rule 1",
		"Direction": "Ingress",
		"AllowType": "Allow",
		"StatefulSession": true,
		"IPAddressType": "IPv4",
		"IPAddressLower": "192.168.1.1",
		"IPAddressUpper": "192.168.1.100",
		"IANAProtocolNumber": 6,
		"SourcePortLower": 5,
		"SourcePortUpper": 65535,
		"DestinationPortLower": 5,
		"DestinationPortUpper": 65535,
		"@odata.id": "/redfish/v1/Chassis/Card1/NetworkAdapters/Slot1/NetworkDeviceFunctions/SC2KP1F0/AllowDeny/Rule1"
	}`

// TestAllowDeny tests the parsing of AllowDeny objects.
func TestAllowDeny(t *testing.T) {
	var result AllowDeny
	err := json.NewDecoder(strings.NewReader(allowDenyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "AllowDeny Rule 1", result.ID)
	assertEquals(t, "Allow Rule 1", result.Name)
	assertEquals(t, "Ingress", string(result.Direction))
	assertEquals(t, "Allow", string(result.AllowType))
	assertEquals(t, "IPv4", string(result.IPAddressType))
	assertEquals(t, "192.168.1.100", result.IPAddressUpper)

	if !result.StatefulSession {
		t.Error("Expected StatefulSession to be true")
	}

	if *result.SourcePortLower != 5 {
		t.Errorf("Expected source lower port to be 5, got %d", result.SourcePortLower)
	}
}
