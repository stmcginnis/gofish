//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var accountBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#AccountService/Links/Members/Accounts/Links/Members/$entity",
		"@odata.id": "/redfish/v1/AccountService/Accounts/1",
		"@odata.type": "#AccountService.0.94.0.ManagerAccount",
		"Id": "1",
		"Name": "User Account",
		"Modified": "2013-09-11T17:03:55+00:00",
		"Description": "User Account",
		"Password": "Password",
		"UserName": "Administrator",
		"Locked": false,
		"Enabled": true,
		"RoleId": "Admin",
		"Links": {
			"Role": {
				"@odata.id": "/redfish/v1/AccountService/Roles/Admin"
			}
		}
	}`)

var accountServiceBody = strings.NewReader(
	`{
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
		}
	}`)

var roleBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#AccountService/Links/Members/Roles/Links/Members/$entity",
		"@odata.id": "/redfish/v1/AccountService/Roles/ReadOnlyUser",
		"@odata.type": "#Role.1.0.0.Role",
		"Id": "ReadOnlyUser",
		"Name": "User Role",
		"Modified": "2013-09-11T17:03:55+00:00",
		"Description": "ReadOnlyUser User Role",
		"IsPredefined": true,
		"AssignedPrivileges": [
			"Login"
		],
		"OEMPrivileges": []
	}`)

// TestAccount tests the parsing of Account objects.
func TestAccount(t *testing.T) {
	var result Account
	err := json.NewDecoder(accountBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "User Account" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.RoleID != "Admin" {
		t.Errorf("Received invalid Role ID: %s", result.RoleID)
	}

	if result.role != "/redfish/v1/AccountService/Roles/Admin" {
		t.Errorf("Received invalid Role: %s", result.role)
	}
}

// TestAccountService tests the parsing of AccountService objects.
func TestAccountService(t *testing.T) {
	var result AccountService
	err := json.NewDecoder(accountServiceBody).Decode(&result)

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
}

// TestRole tests the parsing of Role objects.
func TestRole(t *testing.T) {
	var result Role
	err := json.NewDecoder(roleBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "ReadOnlyUser" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "User Role" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.IsPredefined {
		t.Errorf("IsPredefined incorrect for role.")
	}

	if len(result.AssignedPrivileges) != 1 {
		t.Errorf("Expected 1 assigned privilege, found: %d", len(result.AssignedPrivileges))
	}

	if result.AssignedPrivileges[0] != "Login" {
		t.Errorf("Expected 'Login' assigned privilege, got: %s", result.AssignedPrivileges[0])
	}
}
