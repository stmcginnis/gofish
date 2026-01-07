//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var accountServiceBody = `{
  "@odata.type": "#AccountService.v1_7_2.AccountService",
  "@odata.id": "/redfish/v1/AccountService",
  "Id": "AccountService",
  "Name": "Account Service",
  "Description": "Account Service",
  "Status": {
    "State": "Enabled",
    "Health": "OK"
  },
  "ServiceEnabled": true,
  "MinPasswordLength": 8,
  "MaxPasswordLength": 20,
  "AuthFailureLoggingThreshold": 3,
  "AccountLockoutThreshold": 3,
  "AccountLockoutDuration": 30,
  "AccountLockoutCounterResetAfter": 30,
  "Accounts": {
    "@odata.id": "/redfish/v1/AccountService/Accounts"
  },
  "Roles": {
    "@odata.id": "/redfish/v1/AccountService/Roles"
  },
  "LDAP": {
    "AccountProviderType": "LDAPService",
    "ServiceEnabled": false,
    "ServiceAddresses": [],
    "Authentication": {
      "AuthenticationType": "UsernameAndPassword",
      "Username": "",
      "Password": null,
      "Oem": {}
    },
    "PasswordSet": false,
    "RemoteRoleMapping": [],
    "LDAPService": {
      "SearchSettings": {
        "BaseDistinguishedNames": []
      },
      "Oem": {}
    }
  },
  "ActiveDirectory": {
    "AccountProviderType": "ActiveDirectoryService",
    "ServiceEnabled": false,
    "ServiceAddresses": [],
    "Authentication": {
      "AuthenticationType": "UsernameAndPassword",
      "Username": "",
      "Password": null,
      "Oem": {}
    },
    "PasswordSet": false,
    "RemoteRoleMapping": []
  },
  "Oem": {
    "Supermicro": {
      "@odata.type": "#SmcAccountServiceExtensions.v1_0_1.AccountService",
      "LDAP": {
        "StartTLSEnabled": true
      },
      "ActiveDirectory": {
        "DNSLookupEnable": true,
        "Prefix": "ldap",
        "Port": 389,
        "UserDomainNames": ["example.com"],
        "DynamicServerAddresses": []
      }
    }
  },
  "@odata.etag": "\"01dc844f1c2c3fae22b77263291f161b\""
}`

// TestSmcAccountServiceOem tests the parsing of the AccountService oem field
func TestSmcAccountServiceOem(t *testing.T) {
	drive := &schemas.AccountService{}
	if err := json.Unmarshal([]byte(accountServiceBody), drive); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	accountService, err := FromAccountService(drive)
	if err != nil {
		t.Fatalf("error getting oem info from drive: %v", err)
	}

	if accountService.ID != "AccountService" {
		t.Errorf("unexpected ID: %s", accountService.ID)
	}

	if !accountService.SMCLDAP.StartTLSEnabled {
		t.Errorf("unexpected StartTLSEnabled state: %t", accountService.SMCLDAP.StartTLSEnabled)
	}

	if !accountService.SMCActiveDirectory.DNSLookupEnable ||
		accountService.SMCActiveDirectory.Prefix != "ldap" ||
		accountService.SMCActiveDirectory.Port != 389 ||
		len(accountService.SMCActiveDirectory.UserDomainNames) != 1 ||
		accountService.SMCActiveDirectory.UserDomainNames[0] != "example.com" {
		t.Errorf("unexpected ActiveDirectory settings: %+v", accountService.SMCActiveDirectory)
	}
}
