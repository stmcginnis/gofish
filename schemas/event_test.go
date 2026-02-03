//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var eventBody = `{
	"@odata.type": "#Event.v1_7_0.Event",
	"Id": "1",
	"Name": "Event Array",
	"Context": "ContosoWebClient",
	"Events": [
	  {
		"EventType": "Other",
		"EventId": "4593",
		"Severity": "Warning",
		"Message": "A cable has been removed from network adapter '1' port '1'.",
		"MessageId": "NetworkDevice.1.0.CableRemoved",
		"MessageArgs": [
		  "1",
		  "1"
		],
		"OriginOfCondition": {
		  "@odata.id": "/redfish/v1/Systems/1/EthernetInterfaces/1"
		},
		"LogEntry": {
		  "@odata.id": "/redfish/v1/Managers/BMC/LogServices/EventLog/Entries/532"
		}
	  }
	]
  }`

// TestEvent tests the parsing of Event objects.
func TestEvent(t *testing.T) {
	var result Event
	err := json.NewDecoder(strings.NewReader(eventBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Event Array", result.Name)
	assertEquals(t, "ContosoWebClient", result.Context)
	assertEquals(t, "4593", result.Events[0].EventID)
	assertEquals(t, "Other", string(result.Events[0].EventType))
	assertEquals(t, "Warning", result.Events[0].Severity)
	assertEquals(t, "A cable has been removed from network adapter '1' port '1'.", result.Events[0].Message)
	assertEquals(t, "NetworkDevice.1.0.CableRemoved", result.Events[0].MessageID)
	assertEquals(t, "1", result.Events[0].MessageArgs[0])
	assertEquals(t, "/redfish/v1/Systems/1/EthernetInterfaces/1", result.Events[0].originOfCondition)
	assertEquals(t, "/redfish/v1/Managers/BMC/LogServices/EventLog/Entries/532", result.Events[0].logEntry)
}
