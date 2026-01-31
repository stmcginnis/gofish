//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var addressPoolBody = strings.NewReader(
	`{
		"@odata.type": "#AddressPool.v1_2_4.AddressPool",
		"Id": "AP1",
		"Name": "Address Pool 1",
		"Description": "Address Pool 1",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"GenZ": {
			"MinCID": 1,
			"MaxCID": 4096,
			"MinSID": 100,
			"MaxSID": 8192,
			"AccessKey": "0x1A"
		},
		"Links": {
			"Endpoints": [
				{
				"@odata.id": "/redfish/v1/Fabrics/GenZ/Endpoints/1"
				}
			]
		},
		"@odata.id": "/redfish/v1/Fabrics/GenZ/AddressPools/AP1"
		}
	}`)

// TestAddressPool tests the parsing of AddressPool objects.
func TestAddressPool(t *testing.T) {
	var result AddressPool
	err := json.NewDecoder(addressPoolBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "AP1", result.ID)
	assertEquals(t, "Address Pool 1", result.Name)
	assertEquals(t, "0x1A", result.GenZ.AccessKey)

	if *result.GenZ.MinCID != 1 {
		t.Errorf("Expected Genz.MinCID to be 1, got %d", result.GenZ.MinCID)
	}

	if *result.GenZ.MinSID != 100 {
		t.Errorf("Expected Genz.MinSID to be 100, got %d", result.GenZ.MinSID)
	}

	if *result.GenZ.MaxCID != 4096 {
		t.Errorf("Expected Genz.MaxCID to be 4096, got %d", result.GenZ.MaxCID)
	}

	if *result.GenZ.MaxSID != 8192 {
		t.Errorf("Expected Genz.MinSID to be 8192, got %d", result.GenZ.MaxSID)
	}

	if len(result.endpoints) != 1 {
		t.Errorf("Expected 1 endpoint, got %#v", result.endpoints)
	}
}
