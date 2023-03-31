//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var logEntryCollectionBody = `{
	"@odata.context": "/redfish/v1/$metadata#LogEntryCollection.LogEntryCollection",
	"@odata.etag": "W/\"1649121068\"",
	"@odata.id": "/redfish/v1/Chassis/Self/LogServices/Logs/Entries",
	"@odata.type": "#LogEntryCollection.LogEntryCollection",
	"Description": "Collection of entries",
	"Members": [
	  {
		"@odata.id": "/redfish/v1/Chassis/Self/LogServices/Logs/Entries/1",
		"@odata.type": "#LogEntry.v1_4_3.LogEntry",
		"Created": "2022-05-14T07:23:58Z",
		"Description": "Log Entry",
		"EntryCode": "Upper Non-critical - going high",
		"EntryType": "SEL",
		"EventTimestamp": "2022-05-14T07:23:58Z",
		"EventType": "Alert",
		"Id": "1",
		"Links": {
		  "OriginOfCondition": {
			"@odata.id": "/redfish/v1/Chassis/Self/Power"
		  }
		},
		"Message": "Test message",
		"MessageId": "0x57CECD",
		"Name": "LOG 1",
		"SensorNumber": 24,
		"SensorType": "Voltage",
		"Severity": "Warning"
	  }
	],
	"Members@odata.count": 1,
	"Members@odata.nextLink": "/redfish/v1/Chassis/Self/LogServices/Logs/Entries?$skip=50",
	"Name": "Log Service Entries Collection"
	}`

// TestLogService tests the parsing of LogService objects.
func TestLogEntryCollection(t *testing.T) {
	var result LogEntryCollection
	err := json.NewDecoder(strings.NewReader(logEntryCollectionBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals := func(t testing.TB, expected string, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("\nExpected value: %s \nActual value: %s", expected, actual)
		}
	}

	assertEquals(t, "Log Service Entries Collection", result.Name)
	assertEquals(t, "Collection of entries", result.Description)
	assertEquals(t, "1", fmt.Sprint(result.MembersCount))
	assertEquals(t, "24", fmt.Sprint(result.Members[0].SensorNumber))
	assertEquals(t, "Voltage", fmt.Sprint(result.Members[0].SensorType))
	assertEquals(t, "Warning", fmt.Sprint(result.Members[0].Severity))
	assertEquals(t, "/redfish/v1/Chassis/Self/LogServices/Logs/Entries?$skip=50", result.MembersNextLink)
}
