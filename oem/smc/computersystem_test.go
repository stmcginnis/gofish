//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var computerSystemBody = `{
  "@odata.type": "#ComputerSystem.v1_14_0.ComputerSystem",
  "@odata.id": "/redfish/v1/Systems/1",
  "Id": "1",
  "Name": "System",
  "Description": "Description of server",
  "Status": {
    "State": "Enabled",
    "Health": "OK"
  },
  "SerialNumber": "S514359X4916804",
  "PartNumber": "SYS-821GE-200-02-LL014",
  "IndicatorLED": "Off",
  "LocationIndicatorActive": false,
  "SystemType": "Physical",
  "BiosVersion": "2.1",
  "Manufacturer": "Supermicro",
  "Model": "SYS-821GE-TNHR",
  "SKU": "0x1D1415D9",
  "UUID": "D4216600-C32C-11EE-8000-7CC25586E492",
  "ProcessorSummary": {
    "Count": 2,
    "Model": "Intel(R) Xeon(R) processor",
    "Status": {
      "State": "Enabled",
      "Health": "OK",
      "HealthRollup": "OK"
    },
    "Metrics": {
      "@odata.id": "/redfish/v1/Systems/1/ProcessorSummary/ProcessorMetrics"
    }
  },
  "MemorySummary": {
    "TotalSystemMemoryGiB": 2048,
    "MemoryMirroring": "System",
    "Status": {
      "State": "Enabled",
      "Health": "OK",
      "HealthRollup": "OK"
    },
    "Metrics": {
      "@odata.id": "/redfish/v1/Systems/1/MemorySummary/MemoryMetrics"
    }
  },
  "PowerState": "On",
  "PowerOnDelaySeconds": 3,
  "PowerOnDelaySeconds@Redfish.AllowableNumbers": [
    "3:254:1"
  ],
  "PowerOffDelaySeconds": 3,
  "PowerOffDelaySeconds@Redfish.AllowableNumbers": [
    "3:254:1"
  ],
  "PowerCycleDelaySeconds": 5,
  "PowerCycleDelaySeconds@Redfish.AllowableNumbers": [
    "5:254:1"
  ],
  "Boot": {
    "BootSourceOverrideEnabled": "Disabled",
    "BootSourceOverrideMode": "UEFI",
    "BootSourceOverrideTarget": "Pxe",
    "BootSourceOverrideTarget@Redfish.AllowableValues": [
      "None",
      "Pxe",
      "Floppy",
      "Cd",
      "Usb",
      "Hdd",
      "BiosSetup",
      "UsbCd",
      "UefiBootNext",
      "UefiHttp"
    ],
    "BootOptions": {
      "@odata.id": "/redfish/v1/Systems/1/BootOptions"
    },
    "BootNext": null,
    "BootOrder": [
      "Boot0003",
      "Boot0004",
      "Boot0005",
      "Boot0006",
      "Boot0002",
      "Boot0001"
    ]
  },
  "GraphicalConsole": {
    "ServiceEnabled": true,
    "Port": 5900,
    "MaxConcurrentSessions": 4,
    "ConnectTypesSupported": [
      "KVMIP"
    ]
  },
  "SerialConsole": {
    "MaxConcurrentSessions": 1,
    "SSH": {
      "ServiceEnabled": true,
      "Port": 22,
      "SharedWithManagerCLI": true,
      "ConsoleEntryCommand": "cd system1/sol1; start",
      "HotKeySequenceDisplay": "press <Enter>, <Esc>, and then <T> to terminate session"
    },
    "IPMI": {
      "HotKeySequenceDisplay": "Press ~.  - terminate connection",
      "ServiceEnabled": true,
      "Port": 623
    }
  },
  "VirtualMediaConfig": {
    "ServiceEnabled": true,
    "Port": 623
  },
  "BootProgress": {
    "LastState": "SystemHardwareInitializationComplete"
  },
  "Processors": {
    "@odata.id": "/redfish/v1/Systems/1/Processors"
  },
  "Memory": {
    "@odata.id": "/redfish/v1/Systems/1/Memory"
  },
  "EthernetInterfaces": {
    "@odata.id": "/redfish/v1/Systems/1/EthernetInterfaces"
  },
  "NetworkInterfaces": {
    "@odata.id": "/redfish/v1/Systems/1/NetworkInterfaces"
  },
  "SimpleStorage": {
    "@odata.id": "/redfish/v1/Systems/1/SimpleStorage"
  },
  "Storage": {
    "@odata.id": "/redfish/v1/Systems/1/Storage"
  },
  "LogServices": {
    "@odata.id": "/redfish/v1/Systems/1/LogServices"
  },
  "SecureBoot": {
    "@odata.id": "/redfish/v1/Systems/1/SecureBoot"
  },
  "Bios": {
    "@odata.id": "/redfish/v1/Systems/1/Bios"
  },
  "VirtualMedia": {
    "@odata.id": "/redfish/v1/Managers/1/VirtualMedia"
  },
  "Links": {
    "Chassis": [
      {
        "@odata.id": "/redfish/v1/Chassis/1"
      }
    ],
    "ManagedBy": [
      {
        "@odata.id": "/redfish/v1/Managers/1"
      }
    ]
  },
  "Actions": {
    "Oem": {},
    "#ComputerSystem.Reset": {
      "target": "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset",
      "@Redfish.ActionInfo": "/redfish/v1/Systems/1/ResetActionInfo"
    }
  },
  "Oem": {
    "Supermicro": {
      "NodeManager": {
        "@odata.id": "/redfish/v1/Systems/1/Oem/Supermicro/NodeManager"
      },
      "FixedBootOrder": {
        "@odata.id": "/redfish/v1/Systems/1/Oem/Supermicro/FixedBootOrder"
      },
      "@odata.type": "#SmcSystemExtensions.v1_0_0.System"
    }
  },
  "@odata.etag": "\"36f059d60095337115952f02709f65d6\""
}`

// TestComputerSystem tests the parsing of ComputerSystem objects.
func TestComputerSystem(t *testing.T) {
	var cs schemas.ComputerSystem
	err := json.NewDecoder(strings.NewReader(computerSystemBody)).Decode(&cs)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromComputerSystem(&cs)
	if err != nil {
		t.Errorf("Error converting Redfish ComputerSystem to SMC ComputerSystem: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "System" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.nodeManager != "/redfish/v1/Systems/1/Oem/Supermicro/NodeManager" {
		t.Errorf("Invalid node manager link: %s", result.nodeManager)
	}

	if result.fixedBootOrder != "/redfish/v1/Systems/1/Oem/Supermicro/FixedBootOrder" {
		t.Errorf("Invalid fixed boot order link: %s", result.fixedBootOrder)
	}
}
