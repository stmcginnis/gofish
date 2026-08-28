//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

var accountServiceBody = `{
	"@odata.context": "/redfish/v1/$metadata#AccountService.AccountService",
	"@odata.type": "#AccountService.v1_13_0.AccountService",
	"@odata.id": "/redfish/v1/AccountService",
	"@odata.etag": "\"126793801710\"",
	"AccountLockoutCounterResetAfter": 0,
	"AccountLockoutDuration": 0,
	"AccountLockoutThreshold": 0,
	"Accounts": {
	  "@odata.id": "/redfish/v1/AccountService/Accounts"
	},
	"ActiveDirectory": {
	  "Certificates": {
		"@odata.id": "/redfish/v1/AccountService/ActiveDirectory/Certificates"
	  },
	  "AccountProviderType": "ActiveDirectoryService",
	  "Authentication": {
		"AuthenticationType": "UsernameAndPassword",
		"KerberosKeytab": null
	  },
	  "RemoteRoleMapping": [
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		}
	  ],
	  "RemoteRoleMapping@odata.count": 15,
	  "ServiceAddresses": [
		"",
		"",
		""
	  ],
	  "ServiceAddresses@odata.count": 3,
	  "ServiceEnabled": false
	},
	"AdditionalExternalAccountProviders": {
	  "@odata.id": "/redfish/v1/AccountService/ExternalAccountProviders"
	},
	"AuthFailureLoggingThreshold": 2,
	"Description": "BMC User Accounts",
	"Id": "RemoteAccountService",
	"LDAP": {
	  "Certificates": {
		"@odata.id": "/redfish/v1/AccountService/LDAP/Certificates"
	  },
	  "AccountProviderType": "LDAPService",
	  "Authentication": {
		"AuthenticationType": "UsernameAndPassword"
	  },
	  "LDAPService": {
		"SearchSettings": {
		  "BaseDistinguishedNames": [
			""
		  ],
		  "BaseDistinguishedNames@odata.count": 1,
		  "GroupNameAttribute": "",
		  "UsernameAttribute": ""
		}
	  },
	  "RemoteRoleMapping": [
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		},
		{
		  "RemoteGroup": "",
		  "LocalRole": "None"
		}
	  ],
	  "RemoteRoleMapping@odata.count": 15,
	  "ServiceAddresses": [
		""
	  ],
	  "ServiceAddresses@odata.count": 1,
	  "ServiceEnabled": false
	},
	"LocalAccountAuth": "Fallback",
	"MaxPasswordLength": 127,
	"MinPasswordLength": 0,
	"OAuth2": {
	  "Certificates": {
		"@odata.id": "/redfish/v1/AccountService/ExternalAccountProviders/1/Certificates"
	  },
	  "ServiceEnabled": true,
	  "ServiceAddresses": [
		""
	  ],
	  "ServiceAddresses@odata.count": 1,
	  "OAuth2Service": {
		"Issuer": "",
		"OAuthServiceSigningKeys": "",
		"Audience": [
		  "c4:cb:e1:b4:bc:46"
		],
		"Mode": "Discovery"
	  }
	},
	"Name": "Account Service",
	"PasswordExpirationDays": null,
	"PrivilegeMap": {
	  "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/PrivilegeRegistry"
	},
	"Roles": {
	  "@odata.id": "/redfish/v1/AccountService/Roles"
	},
	"ServiceEnabled": true,
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"SupportedAccountTypes": [
	  "Redfish",
	  "SNMP",
	  "OEM",
	  "HostConsole",
	  "ManagerConsole",
	  "IPMI",
	  "KVMIP",
	  "VirtualMedia",
	  "WebUI"
	],
	"SupportedOEMAccountTypes": [
	  "IPMI",
	  "SOL",
	  "WSMAN",
	  "UI",
	  "Racadm"
	]
  }`

// TestAccountService tests the parsing of AccountService objects.
func TestAccountService(t *testing.T) {
	var result AccountService
	err := json.NewDecoder(strings.NewReader(accountServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "RemoteAccountService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Account Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AuthFailureLoggingThreshold != 2 {
		t.Errorf("Received invalid authentication failure logging threshold: %d",
			result.AuthFailureLoggingThreshold)
	}

	if result.MinPasswordLength != 0 {
		t.Errorf("Received invalid minimum password length: %d", result.MinPasswordLength)
	}

	if result.accounts != "/redfish/v1/AccountService/Accounts" {
		t.Errorf("Received invalid Accounts: %s", result.accounts)
	}

	if result.roles != "/redfish/v1/AccountService/Roles" {
		t.Errorf("Received invalid Roles: %s", result.roles)
	}

	if result.ActiveDirectory.certificates != "/redfish/v1/AccountService/ActiveDirectory/Certificates" {
		t.Errorf("Received invalid ActiveDirectory certificates link: %s", result.ActiveDirectory.certificates)
	}

	if result.LDAP.AccountProviderType != LDAPServiceAccountProviderTypes {
		t.Errorf("Received invalid LDAP account provider type: %s", result.LDAP.AccountProviderType)
	}

	if len(result.SupportedAccountTypes) != 9 {
		t.Errorf("Received invalid number of supported account types: %#v", result.SupportedAccountTypes)
	}
}

// TestAccountServiceUpdate tests the Update call for the account service.
func TestAccountServiceUpdate(t *testing.T) {
	var result AccountService
	err := json.NewDecoder(strings.NewReader(accountServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
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

// TestAccountServiceCreateAccount checks the payload sent for both the simple
// and the parameter-based account creation calls.
func TestAccountServiceCreateAccount(t *testing.T) {
	var result AccountService
	if err := json.NewDecoder(strings.NewReader(accountServiceBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	createdAccount := func() *http.Response {
		return &http.Response{
			Status:     "201 Created",
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(`{"@odata.id": "/redfish/v1/AccountService/Accounts/3", "UserName": "gofish"}`)),
		}
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {createdAccount()},
		},
	}
	result.SetClient(testClient)
	result.disableEtagMatch = true

	if _, err := result.CreateAccount("gofish", "n0tmypassword", "Administrator"); err != nil {
		t.Errorf("Error making CreateAccount call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call to be made, captured: %v", calls)
	}
	if calls[0].URL != result.accounts {
		t.Errorf("Unexpected create URL: %s", calls[0].URL)
	}
	for _, expected := range []string{"UserName:gofish", "Password:n0tmypassword", "RoleId:Administrator", "Enabled:true"} {
		if !strings.Contains(calls[0].Payload, expected) {
			t.Errorf("Expected %q in create payload: %s", expected, calls[0].Payload)
		}
	}
	// Unset optional properties should not be sent.
	if strings.Contains(calls[0].Payload, "PasswordChangeRequired") {
		t.Errorf("Unset properties should be omitted from create payload: %s", calls[0].Payload)
	}

	testClient.Reset()
	testClient.CustomReturnForActions[http.MethodPost] = []any{createdAccount()}
	passwordChangeRequired := true
	_, err := result.CreateAccountParams(&CreateAccountParameters{
		UserName:               "gofish",
		Password:               "n0tmypassword",
		RoleID:                 "Administrator",
		PasswordChangeRequired: &passwordChangeRequired,
		AccountTypes:           []AccountTypes{RedfishAccountTypes},
		OEM:                    map[string]any{"Contoso": map[string]any{"Widget": true}},
	})
	if err != nil {
		t.Errorf("Error making CreateAccountParams call: %s", err)
	}

	calls = testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call to be made, captured: %v", calls)
	}
	for _, expected := range []string{"PasswordChangeRequired:true", "AccountTypes:[Redfish]", "Oem:map[Contoso:map[Widget:true]]"} {
		if !strings.Contains(calls[0].Payload, expected) {
			t.Errorf("Expected %q in create payload: %s", expected, calls[0].Payload)
		}
	}
	if strings.Contains(calls[0].Payload, "Enabled") {
		t.Errorf("Unset properties should be omitted from create payload: %s", calls[0].Payload)
	}
}
