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

var managerAccountBody = `{
		"@odata.context": "/redfish/v1/$metadata#AccountService/Links/Members/Accounts/Links/Members/$entity",
		"@odata.id": "/redfish/v1/AccountService/Accounts/1",
		"@odata.type": "#AccountService.0.94.0.ManagerAccount",
		"Id": "1",
		"Name": "User Account",
		"Modified": "2013-09-11T17:03:55+00:00",
		"Description": "User Account",
		"Password": null,
		"UserName": "Administrator",
		"Locked": false,
		"Enabled": true,
		"RoleId": "Admin",
		"Links": {
			"Role": {
				"@odata.id": "/redfish/v1/AccountService/Roles/Admin"
			}
		}
	}`

// TestAccount tests the parsing of Account objects.
func TestManagerAccount(t *testing.T) {
	var result ManagerAccount
	err := json.NewDecoder(strings.NewReader(managerAccountBody)).Decode(&result)

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

// TestAccountUpdate tests the Update call.
func TestManagerAccountUpdate(t *testing.T) {
	var result ManagerAccount
	err := json.NewDecoder(strings.NewReader(managerAccountBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.ODataEtag = "aaa" // etag might come from the HTTP header as something different
	result.Enabled = false
	result.Locked = false
	result.Password = "Test"
	result.RoleID = "Administrator"
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if strings.Contains(calls[0].Payload, "@odata.etag") {
		t.Errorf("Unexpected etag payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Enabled:false") {
		t.Errorf("Unexpected Enabled update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Locked") {
		t.Errorf("Unexpected Locked in update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Password:Test") {
		t.Errorf("Unexpected Password update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "RoleId:Administrator") {
		t.Errorf("Unexpected Role ID update payload: %s", calls[0].Payload)
	}
}
