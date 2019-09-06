//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var sessionBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Session.Session",
		"@odata.type": "#Session.v1_2_0.Session",
		"@odata.id": "/redfish/v1/Session",
		"Id": "Session-1",
		"Name": "SessionOne",
		"Description": "Session One",
		"OemSessionType": "Ticket",
		"SessionType": "OEM",
		"UserName": "mfreeman"
	}`)

// TestSession tests the parsing of Session objects.
func TestSession(t *testing.T) {
	var result Session
	err := json.NewDecoder(sessionBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Session-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SessionOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Password != "" {
		t.Error("Password should be nil")
	}

	if result.SessionType != OEMSessionTypes {
		t.Errorf("Invalid session type: %s", result.SessionType)
	}

	if result.UserName != "mfreeman" {
		t.Errorf("Invalid user name: %s", result.UserName)
	}
}
