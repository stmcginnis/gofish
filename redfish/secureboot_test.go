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

var secureBootBody = `{
		"@odata.context": "/redfish/v1/$metadata#SecureBoot.SecureBoot",
		"@odata.type": "#SecureBoot.v1_0_5.SecureBoot",
		"@odata.id": "/redfish/v1/SecureBoot",
		"Id": "SecureBoot-1",
		"Name": "SecureBootOne",
		"Description": "SecureBoot One",
		"SecureBootCurrentBoot": "Enabled",
		"SecureBootEnable": true,
		"SecureBootMode": "UserMode",
		"Actions": {
			"#SecureBoot.ResetKeys": {
				"target": "/redfish/v1/SecureBoot/Actions/SecureBoot.ResetKeys"
			}
		}
	}`

// TestSecureBoot tests the parsing of SecureBoot objects.
func TestSecureBoot(t *testing.T) {
	var result SecureBoot
	err := json.NewDecoder(strings.NewReader(secureBootBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SecureBoot-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SecureBootOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.SecureBootCurrentBoot != EnabledSecureBootCurrentBootType {
		t.Errorf("Invalid SecureBootCurrentBoot: %s", result.SecureBootCurrentBoot)
	}

	if !result.SecureBootEnable {
		t.Error("SecureBootEnable should be true")
	}

	if result.SecureBootMode != UserModeSecureBootModeType {
		t.Errorf("Invalid SecureBootMode: %s", result.SecureBootMode)
	}

	if result.resetKeysTarget != "/redfish/v1/SecureBoot/Actions/SecureBoot.ResetKeys" {
		t.Errorf("Invalid ResetKeys target: %s", result.resetKeysTarget)
	}
}

// TestSecureBootUpdate tests the Update call.
func TestSecureBootUpdate(t *testing.T) {
	var result SecureBoot
	err := json.NewDecoder(strings.NewReader(secureBootBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.SecureBootEnable = false
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "SecureBootEnable:false") {
		t.Errorf("Unexpected SecureBootEnable update payload: %s", calls[0].Payload)
	}
}
