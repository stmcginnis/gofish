//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var oemLinksBody = `
			{
				"Dell": {
					"DellAttributes": [
						{
							"@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"
						},
						{
							"@odata.id": "/redfish/v1/Managers/System.Embedded.1/Attributes"
						},
						{
							"@odata.id": "/redfish/v1/Managers/LifecycleController.Embedded.1/Attributes"
						}
					],
					"DellAttributes@odata.count": 3,
					"DellTimeService": {
						"@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/DellTimeService"
					}
				}
			}
`
var oemDataBody = `
		{
			"Dell": {
				"DelliDRACCard": {
					"@odata.context": "/redfish/v1/$metadata#DelliDRACCard.DelliDRACCard",
					"@odata.id": "/redfish/v1/Managers/iDRAC.Embedded.1/DelliDRACCard/iDRAC.Embedded.1-1_0x23_IDRACinfo",
					"@odata.type": "#DelliDRACCard.v1_1_0.DelliDRACCard",
					"IPMIVersion": "2.0",
					"URLString": "https://10.5.1.83:443"
				}
			}
		}
`

var oemActions = `
{
	"#OemManager.v1_2_0.OemManager#OemManager.ExportSystemConfiguration": {
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
				"RAID"
			]
		},
		"target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ExportSystemConfiguration"
	},
	"#OemManager.v1_2_0.OemManager#OemManager.ImportSystemConfiguration": {
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
				"RAID"
			]
		},
		"ShutdownType@Redfish.AllowableValues": [
			"Graceful",
			"Forced",
			"NoReboot"
		],
		"target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfiguration"
	},
	"#OemManager.v1_2_0.OemManager#OemManager.ImportSystemConfigurationPreview": {
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
	},
	"DellManager.v1_0_0#DellManager.ResetToDefaults": {
		"ResetType@Redfish.AllowableValues": [
			"All",
			"ResetAllWithRootDefaults",
			"Default"
		],
		"target": "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/DellManager.ResetToDefaults"
	}
}
`
var managerBody = `{
		"@Redfish.Copyright": "Copyright 2014-2019 DMTF. All rights reserved.",
		"@odata.context": "/redfish/v1/$metadata#Manager.Manager",
		"@odata.id": "/redfish/v1/Managers/BMC-1",
		"@odata.type": "#Manager.v1_1_0.Manager",
		"Id": "BMC-1",
		"LastResetTime": "2022-11-17T08:46:24+00:00",
		"Name": "Manager",
		"ManagerType": "BMC",
		"Description": "BMC",
		"AutoDSTEnabled": true,
		"ServiceEntryPointUUID": "92384634-2938-2342-8820-489239905423",
		"UUID": "00000000-0000-0000-0000-000000000000",
		"Model": "Joo Janta 200",
		"DateTime": "2015-03-13T04:14:33+06:00",
		"DateTimeLocalOffset": "+06:00",
		"PowerState": "On",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"GraphicalConsole": {
			"ServiceEnabled": true,
			"MaxConcurrentSessions": 2,
			"ConnectTypesSupported": [
				"KVMIP"
			]
		},
		"SerialConsole": {
			"ServiceEnabled": true,
			"MaxConcurrentSessions": 1,
			"ConnectTypesSupported": [
				"Telnet",
				"SSH",
				"IPMI"
			]
		},
		"CommandShell": {
			"ServiceEnabled": true,
			"MaxConcurrentSessions": 4,
			"ConnectTypesSupported": [
				"Telnet",
				"SSH"
			]
		},
		"FirmwareVersion": "1.00",
		"RemoteAccountService": {
			"@odata.id": "/redfish/v1/Managers/AccountService"
		},
		"RemoteRedfishServiceUri": "http://example.com/",
		"NetworkProtocol": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/NetworkProtocol"
		},
		"HostInterfaces": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/HostInterfaces"
		},
		"EthernetInterfaces": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/EthernetInterfaces"
		},
		"SerialInterfaces": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/SerialInterfaces"
		},
		"LogServices": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/LogServices"
		},
		"VirtualMedia": {
			"@odata.id": "/redfish/v1/Managers/BMC-1/VM1"
		},
		"Links": {
			"ManagerForServers": [
				{
					"@odata.id": "/redfish/v1/Systems/System-1"
				}
			],
			"ManagerForChassis": [
				{
					"@odata.id": "/redfish/v1/Chassis/Chassis-1"
				}
			],
			"ManagerInChassis": {
				"@odata.id": "/redfish/v1/Chassis/Chassis-1"
			},
			"Oem":
` + oemLinksBody +
	`		},
		"Actions": {
			"#Manager.Reset": {
				"target": "/redfish/v1/Managers/BMC-1/Actions/Manager.Reset",
				"ResetType@Redfish.AllowableValues": [
					"ForceRestart",
					"GracefulRestart"
				]
			},
			"Oem":
` + oemActions +
	`	},
		"Oem":
` + oemDataBody +
	`	}`

// TestManager tests the parsing of Manager objects.
func TestManager(t *testing.T) {
	var result Manager
	err := json.NewDecoder(strings.NewReader(managerBody)).Decode(&result)

	t.Run("Check fields", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		if result.ID != "BMC-1" {
			t.Errorf("Received invalid ID: %s", result.ID)
		}

		if result.LastResetTime != "2022-11-17T08:46:24+00:00" {
			t.Errorf("Received invalid LastResetTime: %s", result.LastResetTime)
		}

		if result.Name != "Manager" {
			t.Errorf("Received invalid name: %s", result.Name)
		}

		if result.ManagerType != BMCManagerType {
			t.Errorf("Received manager type: %s", result.ManagerType)
		}

		if result.PowerState != OnPowerState {
			t.Errorf("Received power state: %s", result.PowerState)
		}

		if !result.GraphicalConsole.ServiceEnabled {
			t.Error("Graphical console service state should be enabled")
		}

		if len(result.SerialConsole.ConnectTypesSupported) != 3 {
			t.Errorf("Serial console should have 3 connect types, got %d",
				len(result.SerialConsole.ConnectTypesSupported))
		}

		if result.managerForServers[0] != "/redfish/v1/Systems/System-1" {
			t.Errorf("Received manager for servers: %s", result.managerForServers)
		}

		if result.resetTarget != "/redfish/v1/Managers/BMC-1/Actions/Manager.Reset" {
			t.Errorf("Invalid Reset target: %s", result.resetTarget)
		}

		var expectedOEM map[string]interface{}
		if err := json.Unmarshal([]byte(oemLinksBody), &expectedOEM); err != nil {
			t.Errorf("Failed to unmarshall link body: %v", err)
		}
		if err := json.Unmarshal([]byte(oemDataBody), &expectedOEM); err != nil {
			t.Errorf("Failed to unmarshall data body: %v", err)
		}
		// Check OEM fields
		if len(result.Oem) == 0 {
			t.Errorf("Oem field empty, expected not empty")
		}
		if len(result.OemLinks) == 0 {
			t.Errorf("OemLinks field empty, expected not empty")
		}
		if len(result.OemActions) == 0 {
			t.Errorf("OemActions field empty, expected not empty")
		}
	})
}

// TestManagerUpdate tests the Update call.
func TestManagerUpdate(t *testing.T) {
	var result Manager
	err := json.NewDecoder(strings.NewReader(managerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.AutoDSTEnabled = false
	result.DateTimeLocalOffset = "+05:00"
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AutoDSTEnabled:false") {
		t.Errorf("Unexpected AutoDSTEnabled update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "DateTimeLocalOffset:+05:00") {
		t.Errorf("Unexpected DateTimeLocalOffset update payload: %s", calls[0].Payload)
	}
}
