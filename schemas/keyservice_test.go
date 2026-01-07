//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var keyServiceBody = `{
	"@odata.type": "#KeyService.v1_0_0.KeyService",
	"Id": "KeyService",
	"Name": "Key Service",
	"NVMeoFSecrets": {
	  "@odata.id": "/redfish/v1/KeyService/NVMeoFSecrets"
	},
	"NVMeoFKeyPolicies": {
	  "@odata.id": "/redfish/v1/KeyService/NVMeoFKeyPolicies"
	},
	"@odata.id": "/redfish/v1/KeyService"
  }`

// TestKeyService tests the parsing of KeyService objects.
func TestKeyService(t *testing.T) {
	var result KeyService
	err := json.NewDecoder(strings.NewReader(keyServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "KeyService", result.ID)
	assertEquals(t, "Key Service", result.Name)
	assertEquals(t, "/redfish/v1/KeyService/NVMeoFSecrets", result.nVMeoFSecrets)
	assertEquals(t, "/redfish/v1/KeyService/NVMeoFKeyPolicies", result.nVMeoFKeyPolicies)
}
