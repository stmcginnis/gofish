//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var sessionServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#SessionService.SessionService",
		"@odata.id": "/redfish/v1/SessionService",
		"@odata.type": "#SessionService.v1_1_8.SessionService",
		"Description": "Session Service",
		"Id": "SessionService",
		"Name": "Session Service",
		"ServiceEnabled": true,
		"SessionTimeout": 1800,
		"Sessions": {
		  "@odata.id": "/redfish/v1/SessionService/Sessions"
		}
	  }`)

// TestSessionService tests the parsing of SessionService objects.
func TestSessionService(t *testing.T) {
	var result SessionService
	err := json.NewDecoder(sessionServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	assertEquals(t, "SessionService", result.ID)
	assertEquals(t, "Session Service", result.Name)
	assertEquals(t, "1800", fmt.Sprintf("%d", result.SessionTimeout))
	assertEquals(t, "true", fmt.Sprintf("%t", result.ServiceEnabled))
	assertEquals(t, "/redfish/v1/SessionService/Sessions", result.sessions)
}
