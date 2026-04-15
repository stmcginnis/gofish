//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
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
	var m redfish.Manager
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
}
