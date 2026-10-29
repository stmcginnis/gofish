//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

const (
	lcServiceURI                  = "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService"
	getRemoteServicesAPIStatusURI = lcServiceURI + "/Actions/DellLCService.GetRemoteServicesAPIStatus"
)

// lcServiceBody is the DellLCService resource from a Dell iDRAC9 (15G
// Monolithic) running firmware 7.30.10.50, trimmed to two of its actions.
var lcServiceBody = `{
	"@odata.context": "/redfish/v1/$metadata#DellLCService.DellLCService",
	"@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService",
	"@odata.type": "#DellLCService.v1_8_1.DellLCService",
	"Actions": {
		"#DellLCService.GetRSStatus": {
			"target": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService/Actions/DellLCService.GetRSStatus"
		},
		"#DellLCService.GetRemoteServicesAPIStatus": {
			"target": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService/Actions/DellLCService.GetRemoteServicesAPIStatus"
		}
	},
	"Description": "The DellLCService resource provides some actions to support Lifecycle Controller functionality.",
	"Id": "DellLCService",
	"Name": "DellLCService"
}`

// remoteServicesAPIStatusReadyBody is the GetRemoteServicesAPIStatus response
// from a Dell iDRAC9 (15G Monolithic) running firmware 7.30.10.50.
var remoteServicesAPIStatusReadyBody = `{
	"@Message.ExtendedInfo": [
		{
			"Message": "The request completed successfully.",
			"MessageArgs": [],
			"MessageArgs@odata.count": 0,
			"MessageId": "Base.1.12.Success",
			"RelatedProperties": [],
			"RelatedProperties@odata.count": 0,
			"Resolution": "None",
			"Severity": "OK"
		},
		{
			"Message": "Lifecycle controller is ready to take provisioning requests",
			"MessageArgs": [],
			"MessageArgs@odata.count": 0,
			"MessageId": "IDRAC.2.9.LC061",
			"RelatedProperties": [],
			"RelatedProperties@odata.count": 0,
			"Resolution": "No response action is required.",
			"Severity": "Informational"
		}
	],
	"LCStatus": "Ready",
	"RTStatus": "Ready",
	"SEKMServiceStatus": "Ready",
	"ServerStatus": "OutOfPOST",
	"Status": "Ready",
	"TelemetryServiceStatus": "Ready"
}`

func TestLCService(t *testing.T) {
	var result LCService
	if err := json.NewDecoder(strings.NewReader(lcServiceBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	if result.ID != "DellLCService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.ODataID != lcServiceURI {
		t.Errorf("Received invalid ODataID: %s", result.ODataID)
	}

	if result.getRemoteServicesAPIStatusTarget != getRemoteServicesAPIStatusURI {
		t.Errorf("Invalid GetRemoteServicesAPIStatus target: %s", result.getRemoteServicesAPIStatusTarget)
	}
}

func TestManagerGetRemoteServicesAPIStatus(t *testing.T) { //nolint:funlen
	tests := []struct {
		name         string
		status       int
		body         string
		expected     RemoteServicesAPIStatus
		expectReady  bool
		expectInfo   int
		expectError  bool
		errorMessage string
	}{
		{
			name:   "ready",
			status: http.StatusOK,
			body:   remoteServicesAPIStatusReadyBody,
			expected: RemoteServicesAPIStatus{
				Status:                 ReadyRemoteServicesStatus,
				LCStatus:               ReadyLCStatus,
				RTStatus:               ReadyRTStatus,
				ServerStatus:           OutOfPOSTServerStatus,
				SEKMServiceStatus:      ReadyServiceStatus,
				TelemetryServiceStatus: ReadyServiceStatus,
			},
			expectReady: true,
			expectInfo:  2,
		},
		{
			name:   "lifecycle controller in use",
			status: http.StatusOK,
			body: `{
				"LCStatus": "InUse",
				"RTStatus": "NotReady",
				"SEKMServiceStatus": "NotReady",
				"ServerStatus": "HaltedF1/F2Prompt",
				"Status": "InUse",
				"TelemetryServiceStatus": "NotReady"
			}`,
			expected: RemoteServicesAPIStatus{
				Status:                 InUseRemoteServicesStatus,
				LCStatus:               InUseLCStatus,
				RTStatus:               NotReadyRTStatus,
				ServerStatus:           HaltedF1F2PromptServerStatus,
				SEKMServiceStatus:      NotReadyServiceStatus,
				TelemetryServiceStatus: NotReadyServiceStatus,
			},
		},
		{
			name:   "fields not reported by the service are left empty",
			status: http.StatusOK,
			body:   `{"LCStatus": "ReloadingData", "RTStatus": "Ready", "ServerStatus": "InPOST", "Status": "NotReady"}`,
			expected: RemoteServicesAPIStatus{
				Status:       NotReadyRemoteServicesStatus,
				LCStatus:     ReloadingDataLCStatus,
				RTStatus:     ReadyRTStatus,
				ServerStatus: InPOSTServerStatus,
			},
		},
		{
			name:   "unknown values are preserved",
			status: http.StatusOK,
			body:   `{"LCStatus": "SomethingNew", "RTStatus": "NotApplicable", "Status": "Ready"}`,
			expected: RemoteServicesAPIStatus{
				Status:   ReadyRemoteServicesStatus,
				LCStatus: LCStatus("SomethingNew"),
				RTStatus: RTStatus("NotApplicable"),
			},
			expectReady: true,
		},
		{
			name:         "server error is returned",
			status:       http.StatusBadRequest,
			body:         `{"error": {"code": "Base.1.12.GeneralError", "message": "LC060"}}`,
			expectError:  true,
			errorMessage: "400",
		},
		{
			name:         "invalid response body is returned",
			status:       http.StatusOK,
			body:         `not json`,
			expectError:  true,
			errorMessage: "failed to decode",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
				http.MethodGet: {
					&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(lcServiceBody))},
				},
				http.MethodPost: {
					&http.Response{StatusCode: test.status, Body: io.NopCloser(strings.NewReader(test.body))},
				},
			}}

			manager := newDellManagerForLCService(t, client)
			status, err := manager.GetRemoteServicesAPIStatus()

			calls := client.CapturedCalls()
			if len(calls) != 2 ||
				calls[0].Action != http.MethodGet || calls[0].URL != lcServiceURI ||
				calls[1].Action != http.MethodPost || calls[1].URL != getRemoteServicesAPIStatusURI {
				t.Fatalf("expected service GET and status POST, got %#v", calls)
			}
			if calls[1].Payload != "" {
				t.Errorf("expected empty request payload, got %q", calls[1].Payload)
			}

			if test.expectError {
				if err == nil || !strings.Contains(err.Error(), test.errorMessage) {
					t.Fatalf("expected error containing %q, got %v", test.errorMessage, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to get remote services API status: %v", err)
			}

			if len(status.ExtendedInfo) != test.expectInfo {
				t.Errorf("expected %d extended info messages, got %d", test.expectInfo, len(status.ExtendedInfo))
			}
			status.ExtendedInfo = nil
			if !reflect.DeepEqual(*status, test.expected) {
				t.Errorf("expected status %#v, got %#v", test.expected, *status)
			}
			if status.IsReady() != test.expectReady {
				t.Errorf("expected IsReady %t, got %t", test.expectReady, status.IsReady())
			}
		})
	}
}

func TestRemoteServicesAPIStatusExtendedInfo(t *testing.T) {
	var status RemoteServicesAPIStatus
	if err := json.Unmarshal([]byte(remoteServicesAPIStatusReadyBody), &status); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	if len(status.ExtendedInfo) != 2 {
		t.Fatalf("expected 2 extended info messages, got %d", len(status.ExtendedInfo))
	}
	if status.ExtendedInfo[1].MessageID != "IDRAC.2.9.LC061" {
		t.Errorf("unexpected message ID: %s", status.ExtendedInfo[1].MessageID)
	}
}

func TestLCServiceGetRemoteServicesAPIStatusUnsupported(t *testing.T) {
	client := &schemas.TestClient{}
	service := &LCService{}
	service.SetClient(client)

	if _, err := service.GetRemoteServicesAPIStatus(); err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported action error, got %v", err)
	}
	if calls := client.CapturedCalls(); len(calls) != 0 {
		t.Fatalf("expected no API calls, got %d", len(calls))
	}
}

func TestManagerLCServiceUnsupported(t *testing.T) {
	manager := &Manager{}
	if _, err := manager.LCService(); err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported service error, got %v", err)
	}
	if _, err := manager.GetRemoteServicesAPIStatus(); err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported service error, got %v", err)
	}
}

func newDellManagerForLCService(t *testing.T, client schemas.Client) *Manager {
	t.Helper()

	manager := &schemas.Manager{
		RawData: []byte(`{}`),
		OEMLinks: json.RawMessage(fmt.Sprintf(`{
			"Dell": {
				"DellLCService": {"@odata.id": %q}
			}
		}`, lcServiceURI)),
	}
	manager.SetClient(client)

	dellManager, err := FromManager(manager)
	if err != nil {
		t.Fatalf("failed to convert Dell manager: %v", err)
	}
	return dellManager
}
