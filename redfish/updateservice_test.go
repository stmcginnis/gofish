//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

// makeHTTPResponse creates an HTTP response for testing, ensuring the body can be read multiple times.
func makeHTTPResponse(statusCode int, headers http.Header, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Header:     headers,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

var simpleUpdateBody = `{
    "@odata.context": "/redfish/v1/$metadata#UpdateService.UpdateService",
    "@odata.id": "/redfish/v1/UpdateService",
    "@odata.type": "#UpdateService.v1_6_0.UpdateService",
    "Actions": {
        "#UpdateService.SimpleUpdate": {
            "TransferProtocol@Redfish.AllowableValues": [
                "HTTP"
            ],
            "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate"
        },
        "Oem": {
            "DellUpdateService.v1_0_0#DellUpdateService.Install": {
                "InstallUpon@Redfish.AllowableValues": [
                    "Now",
                    "NowAndReboot",
                    "NextReboot"
                ],
                "target": "/redfish/v1/UpdateService/Actions/Oem/DellUpdateService.Install"
            }
        }
    },
    "Description": "Represents the properties for the Update Service",
    "FirmwareInventory": {
        "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory"
    },
    "HttpPushUri": "/redfish/v1/UpdateService/FirmwareInventory",
    "Id": "UpdateService",
    "Name": "Update Service"
}`

func TestUpdateService(t *testing.T) {
	var result UpdateService
	assertMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Check default redfish fields", func(t *testing.T) {
		c := &common.TestClient{}
		result.SetClient(c)

		err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.firmwareInventory, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.HTTPPushURI, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.TransferProtocol[0], "HTTP")
		assertMessage(t, result.simpleUpdateTarget, "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate")
	})
}

var startUpdateBody = `{
    "@odata.type": "#UpdateService.v1_8_0.UpdateService",
    "@odata.id": "/redfish/v1/UpdateService",
    "Id": "UpdateService",
    "Name": "Update Service",
    "Description": "Service for updating firmware and includes inventory of firmware",
    "Status": {
      "State": "Enabled",
      "Health": "OK",
      "HealthRollup": "OK"
    },
    "ServiceEnabled": true,
    "MultipartHttpPushUri": "/redfish/v1/UpdateService/upload",
    "FirmwareInventory": {
      "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory"
    },
    "Actions": {
      "Oem": {},
      "#UpdateService.SimpleUpdate": {
        "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate",
        "@Redfish.ActionInfo": "/redfish/v1/UpdateService/SimpleUpdateActionInfo"
      },
      "#UpdateService.StartUpdate": {
        "target": "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate"
      }
    },
    "Oem": {}
    }
  }`

func TestUpdateServiceStartUpdate(t *testing.T) {
	var result UpdateService
	assertMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Check UpdateService.StartUpdate field", func(t *testing.T) {
		c := &common.TestClient{}
		result.SetClient(c)

		err := json.NewDecoder(strings.NewReader(startUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.startUpdateTarget, "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate")
	})
}

func TestSimpleUpdateWithLocationHeader(t *testing.T) {
	var updateService UpdateService
	err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&updateService)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	// Create response with Location header (standard Redfish async pattern)
	taskMonitorURI := "/redfish/v1/TaskService/Tasks/123"
	resp := makeHTTPResponse(http.StatusAccepted, http.Header{"Location": []string{taskMonitorURI}}, "")

	c := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodPost: {resp},
		},
	}
	updateService.SetClient(c)

	result, err := updateService.SimpleUpdate(&SimpleUpdateParameters{
		ImageURI: "http://example.com/firmware.bin",
	})
	if err != nil {
		t.Fatalf("SimpleUpdate failed: %s", err)
	}

	if result == nil {
		t.Fatal("expected TaskMonitorInfo, got nil")
	}

	if result.TaskMonitor != taskMonitorURI {
		t.Errorf("expected TaskMonitor %q, got %q", taskMonitorURI, result.TaskMonitor)
	}
}

func TestSimpleUpdateWithTaskInBody(t *testing.T) {
	var updateService UpdateService
	err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&updateService)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	// Create response with Task object in body (some BMCs do this)
	taskBody := `{
		"@odata.id": "/redfish/v1/TaskService/Tasks/456",
		"Id": "456",
		"Name": "Firmware Update Task",
		"TaskState": "Running",
		"PercentComplete": 0
	}`
	resp := makeHTTPResponse(http.StatusAccepted, http.Header{}, taskBody)

	c := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodPost: {resp},
		},
	}
	updateService.SetClient(c)

	result, err := updateService.SimpleUpdate(&SimpleUpdateParameters{
		ImageURI: "http://example.com/firmware.bin",
	})
	if err != nil {
		t.Fatalf("SimpleUpdate failed: %s", err)
	}

	if result == nil {
		t.Fatal("expected TaskMonitorInfo, got nil")
	}

	if result.Task == nil {
		t.Fatal("expected Task object, got nil")
	}

	if result.Task.ID != "456" {
		t.Errorf("expected Task ID %q, got %q", "456", result.Task.ID)
	}

	if result.Task.TaskState != RunningTaskState {
		t.Errorf("expected TaskState %q, got %q", RunningTaskState, result.Task.TaskState)
	}
}

func TestSimpleUpdateWithLocationHeaderAndTaskInBody(t *testing.T) {
	var updateService UpdateService
	err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&updateService)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	// Create response with both Location header and Task in body
	taskMonitorURI := "/redfish/v1/TaskMonitor/abc"
	taskBody := `{
		"@odata.id": "/redfish/v1/TaskService/Tasks/789",
		"Id": "789",
		"Name": "Firmware Update Task",
		"TaskState": "Starting"
	}`
	resp := makeHTTPResponse(http.StatusAccepted, http.Header{"Location": []string{taskMonitorURI}}, taskBody)

	c := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodPost: {resp},
		},
	}
	updateService.SetClient(c)

	result, err := updateService.SimpleUpdate(&SimpleUpdateParameters{
		ImageURI: "http://example.com/firmware.bin",
	})
	if err != nil {
		t.Fatalf("SimpleUpdate failed: %s", err)
	}

	if result == nil {
		t.Fatal("expected TaskMonitorInfo, got nil")
	}

	// TaskMonitor should be from Location header
	if result.TaskMonitor != taskMonitorURI {
		t.Errorf("expected TaskMonitor %q (from Location header), got %q", taskMonitorURI, result.TaskMonitor)
	}

	// Task object should still be parsed
	if result.Task == nil {
		t.Fatal("expected Task object, got nil")
	}

	if result.Task.ID != "789" {
		t.Errorf("expected Task ID %q, got %q", "789", result.Task.ID)
	}
}
