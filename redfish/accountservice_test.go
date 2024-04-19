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
