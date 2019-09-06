//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var logEntryBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#LogEntry.LogEntry",
		"@odata.type": "#LogEntry.v1_0_0.LogEntry",
		"@odata.id": "/redfish/v1/LogEntry",
		"Id": "LogEntry-1",
		"Name": "LogEntryOne",
		"Description": "LogEntry One",
		"Created": "2012-03-07T14:44+06:00",
		"EntryCode": "Informational",
		"EntryType": "Event",
		"EventGroupId": 21,
		"EventId": "event_entry_1",
		"EventTimestamp":  "2012-03-07T14:44+06:00",
		"Message": "Sorry folks, the parks closed.",
		"MessageArgs": [],
		"SensorNumber": 1,
		"SensorType": "Processor",
		"Severity": "Warning"
	}`)

// TestLogEntry tests the parsing of LogEntry objects.
func TestLogEntry(t *testing.T) {
	var result LogEntry
	err := json.NewDecoder(logEntryBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "LogEntry-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "LogEntryOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.EntryCode != InformationalLogEntryCode {
		t.Errorf("Received entry code: %s", result.EntryCode)
	}

	if result.EventGroupID != 21 {
		t.Errorf("Expected group ID 21, got %d", result.EventGroupID)
	}

	if result.SensorNumber != 1 {
		t.Errorf("Received sensor number %d", result.SensorNumber)
	}

	if result.SensorType != ProcessorSensorType {
		t.Errorf("Received sensor type %s", result.SensorType)
	}

	if result.Severity != WarningEventSeverity {
		t.Errorf("Received log severity %s", result.Severity)
	}
}
