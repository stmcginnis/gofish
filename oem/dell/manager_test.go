//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

const managerBody = `{
    "Actions": {
        "Oem": {
            "#DellManager.ResetToDefaults": {
                "ResetType@Redfish.AllowableValues": [
                    "All",
                    "Default",
                    "ResetAllWithRootDefaults"
                ],
                "target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/DellManager.ResetToDefaults"
            },
            "#DellManager.SetCustomDefaults": {
                "target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/DellManager.SetCustomDefaults"
            },
            "#OemManager.ExportSystemConfiguration": {
                "ExportFormat@Redfish.AllowableValues": [
                    "XML",
                    "JSON"
                ],
                "ExportUse@Redfish.AllowableValues": [
                    "Default",
                    "Clone",
                    "Replace"
                ],
                "IncludeInExport@Redfish.AllowableValues": [
                    "Default",
                    "IncludeReadOnly",
                    "IncludePasswordHashValues",
                    "IncludeCustomTelemetry"
                ],
                "ShareParameters": {
                    "IgnoreCertificateWarning@Redfish.AllowableValues": [
                        "Disabled",
                        "Enabled"
                    ],
                    "ProxySupport@Redfish.AllowableValues": [
                        "Disabled",
                        "EnabledProxyDefault",
                        "Enabled"
                    ],
                    "ProxyType@Redfish.AllowableValues": [
                        "HTTP",
                        "SOCKS4"
                    ],
                    "ShareType@Redfish.AllowableValues": [
                        "LOCAL",
                        "NFS",
                        "CIFS",
                        "HTTP",
                        "HTTPS"
                    ],
                    "Target@Redfish.AllowableValues": [
                        "ALL",
                        "IDRAC",
                        "BIOS",
                        "NIC",
                        "RAID",
                        "FC",
                        "InfiniBand",
                        "SupportAssist",
                        "EventFilters",
                        "System",
                        "LifecycleController",
                        "AHCI",
                        "PCIeSSD"
                    ]
                },
                "target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ExportSystemConfiguration"
            },
            "#OemManager.ImportSystemConfiguration": {
                "ExecutionMode@Redfish.AllowableValues": [
                    "Default",
                    "DeployOnSledInsert",
                    "InstantDeploy"
                ],
                "HostPowerState@Redfish.AllowableValues": [
                    "On",
                    "Off"
                ],
                "ImportSystemConfiguration@Redfish.AllowableValues": [
                    "TimeToWait",
                    "ImportBuffer"
                ],
                "ShareParameters": {
                    "IgnoreCertificateWarning@Redfish.AllowableValues": [
                        "Disabled",
                        "Enabled"
                    ],
                    "ProxySupport@Redfish.AllowableValues": [
                        "Disabled",
                        "EnabledProxyDefault",
                        "Enabled"
                    ],
                    "ProxyType@Redfish.AllowableValues": [
                        "HTTP",
                        "SOCKS4"
                    ],
                    "ShareType@Redfish.AllowableValues": [
                        "LOCAL",
                        "NFS",
                        "CIFS",
                        "HTTP",
                        "HTTPS"
                    ],
                    "Target@Redfish.AllowableValues": [
                        "ALL",
                        "IDRAC",
                        "BIOS",
                        "NIC",
                        "RAID",
                        "FC",
                        "InfiniBand",
                        "SupportAssist",
                        "EventFilters",
                        "System",
                        "LifecycleController",
                        "AHCI",
                        "PCIeSSD"
                    ]
                },
                "ShutdownType@Redfish.AllowableValues": [
                    "Graceful",
                    "Forced",
                    "NoReboot"
                ],
                "target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfiguration"
            },
            "#OemManager.ImportSystemConfigurationPreview": {
                "ImportSystemConfigurationPreview@Redfish.AllowableValues": [
                    "ImportBuffer"
                ],
                "ShareParameters": {
                    "IgnoreCertificateWarning@Redfish.AllowableValues": [
                        "Disabled",
                        "Enabled"
                    ],
                    "ProxySupport@Redfish.AllowableValues": [
                        "Disabled",
                        "EnabledProxyDefault",
                        "Enabled"
                    ],
                    "ProxyType@Redfish.AllowableValues": [
                        "HTTP",
                        "SOCKS4"
                    ],
                    "ShareType@Redfish.AllowableValues": [
                        "LOCAL",
                        "NFS",
                        "CIFS",
                        "HTTP",
                        "HTTPS"
                    ],
                    "Target@Redfish.AllowableValues": [
                        "ALL"
                    ]
                },
                "target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfigurationPreview"
            }
        }
    },
    "Id": "iDRAC.Embedded.1",
    "Links": {
        "Oem": {
            "Dell": {
                "@odata.type": "#DellOem.v1_3_0.DellOemLinks",
                "DellAttributes": [
                    {
                        "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/iDRAC.Embedded.1"
                    },
                    {
                        "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/System.Embedded.1"
                    },
                    {
                        "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/LifecycleController.Embedded.1"
                    }
                ],
                "DellAttributes@odata.count": 3,
                "DellJobService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellJobService"
                },
                "DellLCService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService"
                },
                "DellLicensableDeviceCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLicensableDevices"
                },
                "DellLicenseCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLicenses"
                },
                "DellLicenseManagementService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLicenseManagementService"
                },
                "DellOpaqueManagementDataCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellOpaqueManagementData"
                },
                "DellPersistentStorageService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellPersistentStorageService"
                },
                "DellSwitchConnectionCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections"
                },
                "DellSwitchConnectionService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnectionService"
                },
                "DellSystemManagementService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSystemManagementService"
                },
                "DellSystemQuickSyncCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellSystemQuickSync"
                },
                "DellTimeService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellTimeService"
                },
                "DellUSBDeviceCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellUSBDevices"
                },
                "DelliDRACCardService": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DelliDRACCardService"
                },
                "DellvFlashCollection": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellvFlash"
                },
                "Jobs": {
                    "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/Jobs"
                }
            }
        }
    },
    "Name": "Manager",
    "Oem": {
        "Dell": {
            "@odata.type": "#DellManager.v1_4_0.DellManager",
            "DelliDRACCard": {
                "@odata.context": "/redfish/v1/$metadata#DelliDRACCard.DelliDRACCard",
                "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DelliDRACCard/iDRAC.Embedded.1-1_0x23_IDRACinfo",
                "@odata.type": "#DelliDRACCard.v1_1_0.DelliDRACCard",
                "Description": "An instance of DelliDRACCard will have data specific to the Integrated Dell Remote Access Controller (iDRAC) in the managed system.",
                "IPMIVersion": "2.0",
                "Id": "iDRAC.Embedded.1-1_0x23_IDRACinfo",
                "LastSystemInventoryTime": "2024-08-27T06:34:09+00:00",
                "LastUpdateTime": "2024-09-03T23:23:55+00:00",
                "Name": "DelliDRACCard",
                "URLString": "https://10.17.31.5:443"
            },
            "RemoteSystemLogs": {
                "CA": {
                    "Certificates": {
                        "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/RemoteSystemLogs/CA/Certificates"
                    }
                },
                "HTTPS": {
                    "Certificates": {
                        "@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/RemoteSystemLogs/HTTPS/Certificates"
                    },
                    "SecureClientAuth": "Anonymous",
                    "SecurePort": 6514,
                    "SecureServers": [
                        ""
                    ],
                    "SecureServers@odata.count": 1,
                    "SecureSysLogEnable": "Disabled"
                }
            }
        }
    }
}`

func TestDellManager(t *testing.T) {
	var m schemas.Manager
	err := json.NewDecoder(strings.NewReader(managerBody)).Decode(&m)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromManager(&m)
	if err != nil {
		t.Errorf("Error converting Redfish Manager to SMC Manager: %s", err)
	}

	if result.ID != "iDRAC.Embedded.1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Manager" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.importSystemConfigTarget != "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfiguration" {
		t.Errorf("Invalid ImportSystemConfig link: %s", result.importSystemConfigTarget)
	}

	if result.iDRACCardService != "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DelliDRACCardService" {
		t.Errorf("Invalid iDRAC card service link: %s", result.iDRACCardService)
	}

	expectedAttributes := []string{
		"/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/iDRAC.Embedded.1",
		"/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/System.Embedded.1",
		"/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/LifecycleController.Embedded.1",
	}
	if len(result.DellServiceLinks) != len(expectedAttributes) {
		t.Fatalf("Expected %d DellServiceLinks, got %d", len(expectedAttributes), len(result.DellServiceLinks))
	}
	for i, want := range expectedAttributes {
		if got := result.DellServiceLinks[i]; got != want {
			t.Errorf("DellServiceLinks[%d]: expected %s, got %s", i, want, got)
		}
	}
}

func TestImportSystemConfiguration(t *testing.T) {
	const (
		actionURI = "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfiguration"
		taskURI   = "/redfish/v1/TaskService/Tasks/JID_123"
	)

	validBody := &ImportSystemConfigurationBody{
		ImportBuffer: "<SystemConfiguration/>",
		ShareParameters: ShareParameters{
			ShareType: LocalISCShareType,
			Target:    "ALL",
		},
	}

	t.Run("missing request body", func(t *testing.T) {
		client := &schemas.TestClient{}
		manager := newDellManagerForImport(client, actionURI)

		_, err := manager.ImportSystemConfiguration(nil)
		if err == nil || !strings.Contains(err.Error(), "request body is required") {
			t.Fatalf("expected request body validation error, got %v", err)
		}
		if calls := client.CapturedCalls(); len(calls) != 0 {
			t.Fatalf("expected no API calls, got %d", len(calls))
		}
	})

	t.Run("missing Location header", func(t *testing.T) {
		client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				&http.Response{StatusCode: http.StatusAccepted, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(""))},
			},
		}}
		manager := newDellManagerForImport(client, actionURI)

		_, err := manager.ImportSystemConfiguration(validBody)
		if err == nil || !strings.Contains(err.Error(), "missing Location header") {
			t.Fatalf("expected missing Location header error, got %v", err)
		}
		if calls := client.CapturedCalls(); len(calls) != 1 || calls[0].Action != http.MethodPost {
			t.Fatalf("expected only the action POST, got %#v", calls)
		}
	})

	t.Run("follows task Location", func(t *testing.T) {
		client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				&http.Response{StatusCode: http.StatusAccepted, Header: http.Header{"Location": []string{taskURI}}, Body: io.NopCloser(strings.NewReader(""))},
			},
			http.MethodGet: {
				&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"Id":"JID_123","TaskState":"Running"}`))},
			},
		}}
		manager := newDellManagerForImport(client, actionURI)

		task, err := manager.ImportSystemConfiguration(validBody)
		if err != nil {
			t.Fatalf("unexpected import error: %v", err)
		}
		if task.ID != "JID_123" || task.TaskState != schemas.RunningTaskState {
			t.Fatalf("unexpected task: %#v", task)
		}
		calls := client.CapturedCalls()
		if len(calls) != 2 || calls[0].URL != actionURI || calls[1].URL != taskURI {
			t.Fatalf("unexpected API calls: %#v", calls)
		}
	})
}

func newDellManagerForImport(client schemas.Client, actionURI string) *Manager {
	manager := schemas.Manager{}
	manager.SetClient(client)
	return &Manager{Manager: manager, importSystemConfigTarget: actionURI}
}
