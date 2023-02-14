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

var accountServiceBody = `{
		"@odata.context": "/redfish/v1/$metadata#AccountService",
		"@odata.id": "/redfish/v1/AccountService",
		"@odata.type": "#AccountService.0.94.0.AccountService",
		"Id": "AccountService",
		"Name": "Account Service",
		"Description": "BMC User Accounts",
		"Modified": "2036-09-11T14:17:21+00:00",
		"AuthFailureLoggingThreshold": 3,
		"MinPasswordLength": 8,
		"Links": {
			"Accounts": {
				"@odata.id": "/redfish/v1/AccountService/Accounts"
			},
			"Roles": {
				"@odata.id": "/redfish/v1/AccountService/Roles"
			}
		},
		"Oem": {
			"VendorName": {
				"PasswordChangeOnNextLogin": false,
				"PasswordChangeOnFirstAccess": false
			}
		}
	}`

// TestAccountService tests the parsing of AccountService objects.
func TestAccountService(t *testing.T) {
	var result AccountService
	err := json.NewDecoder(strings.NewReader(accountServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "AccountService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Account Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AuthFailureLoggingThreshold != 3 {
		t.Errorf("Received invalid authentication failure logging threshold: %d",
			result.AuthFailureLoggingThreshold)
	}

	if result.MinPasswordLength != 8 {
		t.Errorf("Received invalid minimum password length: %d", result.MinPasswordLength)
	}

	if result.accounts != "/redfish/v1/AccountService/Accounts" {
		t.Errorf("Received invalid Accounts: %s", result.accounts)
	}

	if result.roles != "/redfish/v1/AccountService/Roles" {
		t.Errorf("Received invalid Roles: %s", result.roles)
	}

	switch oem := result.Oem.(type) {
	case map[string]interface{}:
		for vendor, values := range oem {
			if vendor != "VendorName" {
				t.Errorf("Received invalid Oem vendor: %s", vendor)
			}
			switch val := values.(type) {
			case map[string]interface{}:
				for k, v := range val {
					if k != "PasswordChangeOnNextLogin" && k != "PasswordChangeOnFirstAccess" {
						t.Errorf("Received invalid Oem key %s for vendor: %s", k, vendor)
					}
					if k == "PasswordChangeOnNextLogin" && v != false {
						t.Errorf("Received invalid OemInfo1: %s", v)
					}
					if v == "PasswordChangeOnFirstAccess" && v != false {
						t.Errorf("Received invalid OemInfoN: %s", v)
					}
				}
			}
		}
	default:
		t.Errorf("Received invalid Oem")
	}
}

// TestAccountServiceUpdate tests the Update call for the account service.
func TestAccountServiceUpdate(t *testing.T) {
	var result AccountService
	err := json.NewDecoder(strings.NewReader(accountServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	orginalValue := result.AccountLockoutCounterResetEnabled
	result.AccountLockoutCounterResetEnabled = !orginalValue
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[0].Payload, "AccountLockoutCounterResetEnabled") {
		t.Errorf("Unexpected update payload: %s", calls[0].Payload)
	}
}
