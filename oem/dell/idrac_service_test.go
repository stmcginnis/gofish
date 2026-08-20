//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/schemas"
)

const (
	iDRACCardServiceURI = "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DelliDRACCardService"
	iDRACResetURI       = iDRACCardServiceURI + "/Actions/DelliDRACCardService.iDRACReset"
)

func TestManagerResetIDRAC(t *testing.T) { //nolint:funlen
	tests := []struct {
		name             string
		resetType        IDRACResetType
		status           int
		useCompatibility bool
		expectTask       bool
		expectError      bool
	}{
		{
			name:             "synchronous forced reset through compatibility method",
			resetType:        ForceiDRACReset,
			status:           http.StatusOK,
			useCompatibility: true,
		},
		{
			name:      "created response completes synchronously",
			resetType: ForceiDRACReset,
			status:    http.StatusCreated,
		},
		{
			name:       "asynchronous graceful reset returns task monitor",
			resetType:  GracefuliDRACReset,
			status:     http.StatusAccepted,
			expectTask: true,
		},
		{
			name:        "server error is returned",
			resetType:   GracefuliDRACReset,
			status:      http.StatusInternalServerError,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var postedResetType IDRACResetType
			var serviceRequests int
			var resetRequests int

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch {
				case r.Method == http.MethodGet && r.URL.Path == schemas.DefaultServiceRoot:
					_, _ = w.Write([]byte(`{"@odata.id":"/redfish/v1/","Id":"RootService","Name":"Root Service"}`))
				case r.Method == http.MethodGet && r.URL.Path == iDRACCardServiceURI:
					serviceRequests++
					_, _ = fmt.Fprintf(w, `{
						"@odata.id": %q,
						"Id": "DelliDRACCardService",
						"Actions": {
							"#DelliDRACCardService.iDRACReset": {"target": %q}
						}
					}`, iDRACCardServiceURI, iDRACResetURI)
				case r.Method == http.MethodPost && r.URL.Path == iDRACResetURI:
					resetRequests++
					var request iDRACResetRequest
					if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
						t.Errorf("failed to decode reset request: %v", err)
					}
					postedResetType = request.Force
					if test.status == http.StatusAccepted {
						w.Header().Set("Location", "/redfish/v1/TaskService/TaskMonitors/1")
						w.WriteHeader(test.status)
						_, _ = w.Write([]byte(`{"@odata.id":"/redfish/v1/TaskService/Tasks/1","TaskState":"Running"}`))
						return
					}
					w.WriteHeader(test.status)
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer server.Close()

			client, err := gofish.Connect(gofish.ClientConfig{
				Endpoint:   server.URL,
				HTTPClient: server.Client(),
			})
			if err != nil {
				t.Fatalf("failed to connect test client: %v", err)
			}

			manager := newDellManagerForIDRACReset(t, client)
			var taskInfo *schemas.TaskMonitorInfo
			if test.useCompatibility {
				err = manager.ResetiDRAC(test.resetType)
			} else {
				taskInfo, err = manager.ResetIDRAC(test.resetType)
			}
			if test.expectError {
				if err == nil || !strings.Contains(err.Error(), "500") {
					t.Fatalf("expected server error, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to reset iDRAC: %v", err)
			}

			if serviceRequests != 1 || resetRequests != 1 {
				t.Fatalf("expected one service GET and one reset POST, got %d and %d", serviceRequests, resetRequests)
			}
			if postedResetType != test.resetType {
				t.Errorf("expected reset type %q, got %q", test.resetType, postedResetType)
			}
			if test.expectTask {
				if taskInfo == nil {
					t.Fatal("expected task monitor, got nil")
				}
				if taskInfo.TaskMonitor != "/redfish/v1/TaskService/TaskMonitors/1" {
					t.Errorf("unexpected task monitor URI: %s", taskInfo.TaskMonitor)
				}
				if taskInfo.Task == nil || taskInfo.Task.TaskState != schemas.RunningTaskState {
					t.Errorf("unexpected task representation: %#v", taskInfo.Task)
				}
			} else if taskInfo != nil {
				t.Errorf("expected synchronous response, got task: %#v", taskInfo)
			}
		})
	}
}

func TestIDRACCardServiceResetValidation(t *testing.T) {
	service := &IDRACCardService{}

	if _, err := service.Reset(IDRACResetType("Invalid")); err == nil || !strings.Contains(err.Error(), "invalid iDRAC reset type") {
		t.Fatalf("expected invalid reset type error, got %v", err)
	}

	if _, err := service.Reset(GracefuliDRACReset); err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported reset error, got %v", err)
	}
}

func TestManagerIDRACCardServiceUnsupported(t *testing.T) {
	manager := &Manager{}
	if _, err := manager.IDRACCardService(); err == nil || !strings.Contains(err.Error(), "not supported") {
		t.Fatalf("expected unsupported service error, got %v", err)
	}
}

func newDellManagerForIDRACReset(t *testing.T, client schemas.Client) *Manager {
	t.Helper()

	manager := &schemas.Manager{
		RawData: []byte(`{}`),
		OEMLinks: json.RawMessage(fmt.Sprintf(`{
			"Dell": {
				"DelliDRACCardService": {"@odata.id": %q}
			}
		}`, iDRACCardServiceURI)),
	}
	manager.SetClient(client)

	dellManager, err := FromManager(manager)
	if err != nil {
		t.Fatalf("failed to convert Dell manager: %v", err)
	}
	return dellManager
}
