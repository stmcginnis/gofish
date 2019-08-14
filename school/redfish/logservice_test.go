//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var logServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#LogService.LogService",
		"@odata.type": "#LogService.v1_0_0.LogService",
		"@odata.id": "/redfish/v1/LogService",
		"Id": "LogService-1",
		"Name": "LogServiceOne",
		"Description": "LogService One",
		"DateTime": "2012-03-07T14:44+06:00",
		"Entries": {
			"@odata.id": "/redfish/v1/LogEntryCollection"
		},
		"LogEntryType": "Event",
		"MaxNumberOfRecords": 1000,
		"OverWritePolicy": "WrapsWhenFull",
		"ServiceEnabled": true,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`)

// TestLogService tests the parsing of LogService objects.
func TestLogService(t *testing.T) {
	var result LogService
	err := json.NewDecoder(logServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "LogService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "LogServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.entries != "/redfish/v1/LogEntryCollection" {
		t.Errorf("Received invalid log entry collection: %s", result.entries)
	}

	if result.LogEntryType != EventLogEntryTypes {
		t.Errorf("Received %s log entry type", result.LogEntryType)
	}

	if result.MaxNumberOfRecords != 1000 {
		t.Errorf("Received %d max number of records", result.MaxNumberOfRecords)
	}

	if result.OverWritePolicy != WrapsWhenFullOverWritePolicy {
		t.Errorf("Received %s overwrite policy", result.OverWritePolicy)
	}

	if !result.ServiceEnabled {
		t.Error("Service should be enabled")
	}
}
