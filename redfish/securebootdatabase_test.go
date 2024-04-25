//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var secureBootDatabaseBody = `{
	"@odata.type": "#SecureBootDatabase.v1_0_2.SecureBootDatabase",
	"Id": "PK",
	"Name": "PK - Platform Key",
	"Description": "UEFI PK Secure Boot Database",
	"DatabaseId": "PK",
	"Certificates": {
	  "@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/"
	},
	"Actions": {
	  "#SecureBootDatabase.ResetKeys": {
		"target": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Actions/SecureBootDatabase.ResetKeys",
		"ResetKeysType@Redfish.AllowableValues": [
		  "ResetAllKeysToDefault",
		  "DeleteAllKeys"
		]
	  }
	},
	"@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK"
  }`

// TestSecureBootDatabase tests the parsing of SecureBootDatabase objects.
func TestSecureBootDatabase(t *testing.T) {
	var result SecureBootDatabase
	err := json.NewDecoder(strings.NewReader(secureBootDatabaseBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PK", result.ID)
	assertEquals(t, "PK - Platform Key", result.Name)
	assertEquals(t, "PK", result.DatabaseID)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/", result.certificates)
}

// TestSecureBootDatabaseResetKeys tests the SecureBootDatabase ResetKeys call.
func TestSecureBootDatabaseResetKeys(t *testing.T) {
	var result SecureBootDatabase
	err := json.NewDecoder(strings.NewReader(applicationBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.ResetKeys(DeleteAllKeysResetKeysType)
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "ResetKeysType:DeleteAllKeys") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}
}
