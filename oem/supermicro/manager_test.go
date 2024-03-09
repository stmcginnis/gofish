//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

const managerBody = `{
    "@odata.type": "#Manager.v1_5_2.Manager",
    "@odata.id": "/redfish/v1/Managers/1",
    "Id": "1",
    "Name": "Manager",
    "Description": "BMC",
    "ManagerType": "BMC",
    "UUID": "00000000-0000-0000-0000-AC1F6B3FF75A",
    "Model": "ASPEED",
    "FirmwareVersion": "01.00.17",
    "DateTime": "2022-10-19T14:29:05Z",
    "DateTimeLocalOffset": "+00:00",
    "Status": {
        "State": "Enabled",
        "Health": "OK"
    },
    "GraphicalConsole": {
        "ServiceEnabled": true,
        "MaxConcurrentSessions": 4,
        "ConnectTypesSupported": [
            "KVMIP"
        ]
    },
    "SerialConsole": {
        "ServiceEnabled": true,
        "MaxConcurrentSessions": 1,
        "ConnectTypesSupported": [
            "SSH",
            "IPMI"
        ]
    },
    "CommandShell": {
        "ServiceEnabled": true,
        "MaxConcurrentSessions": 0,
        "ConnectTypesSupported": [
            "SSH"
        ]
    },
    "EthernetInterfaces": {
        "@odata.id": "/redfish/v1/Managers/1/EthernetInterfaces"
    },
    "HostInterfaces": {
        "@odata.id": "/redfish/v1/Managers/1/HostInterfaces"
    },
    "SerialInterfaces": {
        "@odata.id": "/redfish/v1/Managers/1/SerialInterfaces"
    },
    "NetworkProtocol": {
        "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol"
    },
    "LogServices": {
        "@odata.id": "/redfish/v1/Managers/1/LogServices"
    },
    "VirtualMedia": {
        "@odata.id": "/redfish/v1/Managers/1/VirtualMedia"
    },
    "Links": {
        "ManagerForServers": [
            {
                "@odata.id": "/redfish/v1/Systems/1"
            }
        ],
        "ManagerForChassis": [
            {
                "@odata.id": "/redfish/v1/Chassis/1"
            }
        ],
        "Oem": {}
    },
    "Actions": {
        "Oem": {
            "#SmcManagerConfig.Reset": {
                "target": "/redfish/v1/Managers/1/Actions/Oem/SmcManagerConfig.Reset",
                "@Redfish.ActionInfo": "/redfish/v1/Managers/1/Oem/Supermicro/ResetActionInfo"
            }
        },
        "#Manager.Reset": {
            "target": "/redfish/v1/Managers/1/Actions/Manager.Reset"
        }
    },
    "Oem": {
        "Supermicro": {
            "@odata.type": "#SmcManagerExtensions.v1_0_0.Manager",
            "SMTP": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SMTP"
            },
            "RADIUS": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/RADIUS"
            },
            "MouseMode": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MouseMode"
            },
            "NTP": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/NTP"
            },
            "IPAccessControl": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IPAccessControl"
            },
            "SMCRAKP": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SMCRAKP"
            },
            "SNMP": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SNMP"
            },
            "Syslog": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog"
            },
            "SysLockdown": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SysLockdown"
            },
            "Snooping": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Snooping"
            },
            "FanMode": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/FanMode"
            },
            "IKVM": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IKVM"
            },
            "KCSInterface": {
                "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/KCSInterface"
            },
            "LicenseManager": {
                "@odata.id": "/redfish/v1/Managers/1/LicenseManager"
            }
        }
    }
}`

func TestManager(t *testing.T) {
	var manager *redfish.Manager
	err := json.NewDecoder((strings.NewReader(managerBody))).Decode(&manager)

	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	supermicroManager, err := FromManager(manager)
	if err != nil {
		t.Errorf("Convert manager failed: %v", err)
		return
	}

	var emptyManager ManagerOem
	if supermicroManager.Oem == emptyManager {
		t.Errorf("Manager is empty")
	}
}
