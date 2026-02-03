//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

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

	if result.AssignedPrivileges[0] != LoginPrivilegeType {
		t.Errorf("Expected 'Login' assigned privilege, got: %s", result.AssignedPrivileges[0])
	}
}
