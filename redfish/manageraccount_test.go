//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var managerAccountBody = strings.NewReader(
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

// TestAccount tests the parsing of Account objects.
func TestManagerAccount(t *testing.T) {
	var result ManagerAccount
	err := json.NewDecoder(managerAccountBody).Decode(&result)

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
