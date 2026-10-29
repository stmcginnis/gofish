//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stmcginnis/gofish/schemas"
)

const (
	exportSystemConfigURI = "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ExportSystemConfiguration"
	exportTaskURI         = "/redfish/v1/TaskService/Tasks/JID_123456789012"
)

// exportedSystemConfigurationBody is the final task monitor response of a LOCAL
// JSON export of the BIOS target from a Dell PowerEdge R450 (iDRAC9 firmware
// 7.30.10.50), trimmed to three attributes and with the service tag replaced.
var exportedSystemConfigurationBody = `{ "SystemConfiguration": {
  "Comments": [
    { "Comment": "Export type is Normal,JSON,Selective" },
    { "Comment": "Exported configuration may contain commented attributes. Attributes may be commented due to dependency, destructive nature, preserving server identity or for security reasons." }
  ],
  "Model": "PowerEdge R450",
  "ServiceTag": "SVCTAG1",
  "TimeStamp": "Thu Oct 29 09:47:52 2026",
  "Components": [
    {
      "FQDD": "BIOS.Setup.1-1",
      "Attributes": [
        { "Name": "WorkloadProfile", "Value": "NotConfigured", "Set On Import": "True", "Comment": "Read and Write" },
        { "Name": "LogicalProc", "Value": "Enabled", "Set On Import": "True", "Comment": "Read and Write" },
        { "Name": "ProcVirtualization", "Value": "Enabled", "Set On Import": "True", "Comment": "Read and Write" }
      ]
    }
  ]
}}`

var exportedSystemConfigurationXMLBody = `<SystemConfiguration Model="PowerEdge R450" ServiceTag="SVCTAG1" TimeStamp="Thu Oct 29 09:47:52 2026">
<Component FQDD="BIOS.Setup.1-1">
<Attribute Name="LogicalProc">Enabled</Attribute>
</Component>
</SystemConfiguration>`

// exportRunningTaskBody is a Task returned while the export is running.
var exportRunningTaskBody = `{
	"@odata.id": "/redfish/v1/TaskService/Tasks/JID_123456789012",
	"Id": "JID_123456789012",
	"TaskState": "Running",
	"PercentComplete": 0
}`

// exportFailedTaskBody is a Task returned when the export fails. It follows
// the Redfish Task schema; a failed export was not captured from hardware.
var exportFailedTaskBody = `{
	"@odata.id": "/redfish/v1/TaskService/Tasks/JID_123456789012",
	"Id": "JID_123456789012",
	"TaskState": "Exception",
	"TaskStatus": "Critical",
	"Messages": [
		{ "Message": "Unable to export the system configuration.", "MessageId": "IDRAC.2.9.SYS043" }
	]
}`

func newDellManagerForExport(client schemas.Client, actionURI string) *Manager {
	manager := schemas.Manager{}
	manager.SetClient(client)
	return &Manager{Manager: manager, exportSystemConfigTarget: actionURI}
}

func newHTTPResponse(status int, header http.Header, body string) *http.Response {
	if header == nil {
		header = http.Header{}
	}
	return &http.Response{StatusCode: status, Header: header, Body: io.NopCloser(strings.NewReader(body))}
}

func TestExportSystemConfigurationValidation(t *testing.T) {
	tests := []struct {
		name         string
		body         *ExportSystemConfigurationBody
		errorMessage string
	}{
		{"missing request body", nil, "request body is required"},
		{"missing target", &ExportSystemConfigurationBody{}, "ShareParameters.Target is required"},
		{
			"invalid export format",
			&ExportSystemConfigurationBody{ExportFormat: "YAML", ShareParameters: ShareParameters{Target: "ALL"}},
			"invalid ExportFormat",
		},
		{
			"invalid share type",
			&ExportSystemConfigurationBody{ShareParameters: ShareParameters{Target: "ALL", ShareType: "FTP"}},
			"invalid ShareType",
		},
		{
			"network share without IP address",
			&ExportSystemConfigurationBody{ShareParameters: ShareParameters{Target: "ALL", ShareType: NFSISCShareType, ShareName: "/exports"}},
			"IPAddress is required",
		},
		{
			"network share without share name",
			&ExportSystemConfigurationBody{ShareParameters: ShareParameters{Target: "ALL", ShareType: NFSISCShareType, IPAddress: "192.0.2.10"}},
			"ShareName is required",
		},
		{
			"CIFS share without user name",
			&ExportSystemConfigurationBody{ShareParameters: ShareParameters{
				Target: "ALL", ShareType: CIFSISCShareType, IPAddress: "192.0.2.10", ShareName: "share",
			}},
			"UserName is required",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &schemas.TestClient{}
			manager := newDellManagerForExport(client, exportSystemConfigURI)

			_, err := manager.ExportSystemConfiguration(test.body)
			if err == nil || !strings.Contains(err.Error(), test.errorMessage) {
				t.Fatalf("expected error containing %q, got %v", test.errorMessage, err)
			}
			if calls := client.CapturedCalls(); len(calls) != 0 {
				t.Fatalf("expected no API calls, got %d", len(calls))
			}
		})
	}
}

func TestExportSystemConfigurationValidBodies(t *testing.T) {
	bodies := []*ExportSystemConfigurationBody{
		{ShareParameters: ShareParameters{Target: "BIOS"}},
		{ExportFormat: XMLESCExportFormat, ShareParameters: ShareParameters{Target: "ALL", ShareType: LocalISCShareType}},
		{ShareParameters: ShareParameters{Target: "ALL", ShareType: NFSISCShareType, IPAddress: "192.0.2.10", ShareName: "/exports"}},
		{ShareParameters: ShareParameters{
			Target: "ALL", ShareType: CIFSISCShareType, IPAddress: "192.0.2.10", ShareName: "share", UserName: "user",
		}},
	}

	for i, body := range bodies {
		if err := validateExportSystemConfigurationBody(body); err != nil {
			t.Errorf("body %d: unexpected validation error: %v", i, err)
		}
	}
}

func TestExportSystemConfigurationUnsupported(t *testing.T) {
	client := &schemas.TestClient{}
	manager := newDellManagerForExport(client, "")

	_, err := manager.ExportSystemConfiguration(&ExportSystemConfigurationBody{ShareParameters: ShareParameters{Target: "ALL"}})
	if err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported action error, got %v", err)
	}
	if calls := client.CapturedCalls(); len(calls) != 0 {
		t.Fatalf("expected no API calls, got %d", len(calls))
	}
}

func TestExportSystemConfiguration(t *testing.T) {
	body := &ExportSystemConfigurationBody{
		ExportFormat:    JSONESCExportFormat,
		IncludeInExport: IncludeReadOnlyESCIncludeInExport,
		ShareParameters: ShareParameters{Target: "BIOS"},
	}

	tests := []struct {
		name         string
		response     *http.Response
		errorMessage string
	}{
		{
			// The iDRAC9 response is a 202 with an empty body and a Location
			// header pointing at the task.
			name:     "accepted export returns task monitor",
			response: newHTTPResponse(http.StatusAccepted, http.Header{"Location": []string{exportTaskURI}}, ""),
		},
		{
			name:         "missing Location header",
			response:     newHTTPResponse(http.StatusAccepted, nil, ""),
			errorMessage: "missing Location header",
		},
		{
			name:         "synchronous response",
			response:     newHTTPResponse(http.StatusOK, nil, "{}"),
			errorMessage: "did not start a task",
		},
		{
			name:         "server error is returned",
			response:     newHTTPResponse(http.StatusBadRequest, nil, `{"error": {"code": "Base.1.12.GeneralError"}}`),
			errorMessage: "400",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
				http.MethodPost: {test.response},
			}}
			manager := newDellManagerForExport(client, exportSystemConfigURI)

			taskInfo, err := manager.ExportSystemConfiguration(body)

			calls := client.CapturedCalls()
			if len(calls) != 1 || calls[0].Action != http.MethodPost || calls[0].URL != exportSystemConfigURI {
				t.Fatalf("expected only the export POST, got %#v", calls)
			}
			expectedPayload := "map[ExportFormat:JSON IncludeInExport:IncludeReadOnly ShareParameters:map[Target:BIOS]]"
			if calls[0].Payload != expectedPayload {
				t.Errorf("expected payload %q, got %q", expectedPayload, calls[0].Payload)
			}

			if test.errorMessage != "" {
				if err == nil || !strings.Contains(err.Error(), test.errorMessage) {
					t.Fatalf("expected error containing %q, got %v", test.errorMessage, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected export error: %v", err)
			}
			if taskInfo.TaskMonitor != exportTaskURI {
				t.Errorf("expected task monitor %s, got %s", exportTaskURI, taskInfo.TaskMonitor)
			}
		})
	}
}

func TestWaitForExportedSystemConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		responses    []any
		expected     string
		errorMessage string
	}{
		{
			name: "running export then JSON configuration",
			responses: []any{
				newHTTPResponse(http.StatusAccepted, nil, exportRunningTaskBody),
				newHTTPResponse(http.StatusOK, http.Header{
					"Content-Disposition": []string{"attachment;filename=guiexportscplocal.json"},
				}, exportedSystemConfigurationBody),
			},
			expected: exportedSystemConfigurationBody,
		},
		{
			name:      "XML configuration",
			responses: []any{newHTTPResponse(http.StatusOK, nil, exportedSystemConfigurationXMLBody)},
			expected:  exportedSystemConfigurationXMLBody,
		},
		{
			name:         "failed export returns task error",
			responses:    []any{newHTTPResponse(http.StatusOK, nil, exportFailedTaskBody)},
			errorMessage: "task JID_123456789012 Exception: Unable to export the system configuration.",
		},
		{
			name:         "unexpected response body",
			responses:    []any{newHTTPResponse(http.StatusOK, nil, `{"Name": "Something else"}`)},
			errorMessage: "does not contain a system configuration",
		},
		{
			name:         "server error is returned",
			responses:    []any{newHTTPResponse(http.StatusInternalServerError, nil, "")},
			errorMessage: "500",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
				http.MethodGet: test.responses,
			}}
			manager := newDellManagerForExport(client, exportSystemConfigURI)
			taskInfo := &schemas.TaskMonitorInfo{TaskMonitor: exportTaskURI}

			content, err := manager.WaitForExportedSystemConfiguration(context.Background(), taskInfo, time.Millisecond)

			calls := client.CapturedCalls()
			if len(calls) != len(test.responses) {
				t.Fatalf("expected %d task monitor GETs, got %#v", len(test.responses), calls)
			}
			for _, call := range calls {
				if call.Action != http.MethodGet || call.URL != exportTaskURI {
					t.Errorf("unexpected call: %#v", call)
				}
			}

			if test.errorMessage != "" {
				if err == nil || !strings.Contains(err.Error(), test.errorMessage) {
					t.Fatalf("expected error containing %q, got %v", test.errorMessage, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected wait error: %v", err)
			}
			if string(content) != test.expected {
				t.Errorf("unexpected exported configuration: %s", content)
			}
		})
	}
}

func TestWaitForExportedSystemConfigurationNilTask(t *testing.T) {
	manager := newDellManagerForExport(&schemas.TestClient{}, exportSystemConfigURI)

	if _, err := manager.WaitForExportedSystemConfiguration(context.Background(), nil, 0); err == nil {
		t.Fatal("expected error for nil task monitor")
	}
}
