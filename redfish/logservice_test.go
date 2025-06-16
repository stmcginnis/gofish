//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var logServiceBodyTmpl = `{
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
		},
		"Actions": {
			"#LogService.ClearLog": {
				"target": "/redfish/v1/Managers/BMC/LogServices/Log/Actions/LogService.ClearLog"
			}%s
		}
	}`

var logServiceBody = fmt.Sprintf(logServiceBodyTmpl, `
	, "#LogService.CollectDiagnosticData": {
		"target": "/redfish/v1/Managers/BMC/LogServices/Log/Actions/LogService.CollectDiagnosticData"
	}
`)
var logServiceBodyNoDiag = fmt.Sprintf(logServiceBodyTmpl, "")

// TestLogService tests the parsing of LogService objects.
func TestLogService(t *testing.T) {
	var result LogService
	err := json.NewDecoder(strings.NewReader(logServiceBody)).Decode(&result)

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

	if result.clearLogTarget != "/redfish/v1/Managers/BMC/LogServices/Log/Actions/LogService.ClearLog" {
		t.Errorf("Invalid ClearLog target: %s", result.clearLogTarget)
	}

	if result.collectDiagnosticDataTarget != "/redfish/v1/Managers/BMC/LogServices/Log/Actions/LogService.CollectDiagnosticData" {
		t.Errorf("Invalid CollectDiagnosticData target: %s", result.collectDiagnosticDataTarget)
	}
}

func initLogServiceClient(t *testing.T, template string) (*LogService, *common.TestClient) {
	var result LogService
	err := json.NewDecoder(strings.NewReader(template)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)
	return &result, testClient
}

// TestLogServiceUpdate tests the Update call.
func TestLogServiceUpdate(t *testing.T) {
	result, testClient := initLogServiceClient(t, logServiceBody)

	result.ServiceEnabled = false
	err := result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "ServiceEnabled:false") {
		t.Errorf("Unexpected ServiceEnabled update payload: %s", calls[0].Payload)
	}
}

// TestLogServiceCollectDiagnosticsDataSuccess
func TestLogServiceCollectDiagnosticsDataSuccess(t *testing.T) {
	logSvc, testClient := initLogServiceClient(t, logServiceBody)

	if !logSvc.SupportsCollectDiagnosticData() {
		t.Errorf("Log service doesn't support diagnostic data")
	}

	diagnosticLocation := "/redfish/v1/Managers/BMC/LogServices/Log/Entries/10"

	testClient.CustomReturnForActions = map[string][]interface{}{
		http.MethodPost: []interface{}{
			&http.Response{
				StatusCode: http.StatusCreated,
				Header: http.Header{
					"Location": []string{diagnosticLocation},
				},
				Body: io.NopCloser(strings.NewReader(`
				{
					"error": {
						"code": "Base.1.12.Success",
						"message": "Successfully Completed Request",
						"@Message.ExtendedInfo": [
							{
								"MessageId": "Base.1.12.Success",
								"Message": "Successfully Completed Request",
								"MessageSeverity": "OK",
								"Resolution": "None",
								"@odata.type": "#Message.v1_1_2.Message"
							}
						]
					}
				}
			`)),
			},
		}}

	location, err := logSvc.CollectDiagnosticData(&CollectDiagnosticDataParameters{
		DiagnosticDataType: ManagerDiagnosticDataTypes,
	})
	if err != nil {
		t.Errorf("Error triggering diagnostic data: %s", err)
	}

	assertEquals(t, diagnosticLocation, location)
}

// TestLogServiceCollectDiagnosticsDataUnsupported
func TestLogServiceCollectionDiagnosticsDataUnsupported(t *testing.T) {
	logSvc, _ := initLogServiceClient(t, logServiceBodyNoDiag)

	if logSvc.SupportsCollectDiagnosticData() {
		t.Errorf("log service unexpectedly supports diagnostic data")
	}
}
