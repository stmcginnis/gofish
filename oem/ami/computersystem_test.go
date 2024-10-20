//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

var computerSystemBody = `{
  "@Redfish.Settings": {
    "@odata.type": "#Settings.v1_2_2.Settings",
    "SettingsObject": {
      "@odata.id": "/redfish/v1/Systems/System_0/SD"
    }
  },
  "@odata.context": "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
  "@odata.etag": "\"1729105654\"",
  "@odata.id": "/redfish/v1/Systems/System_0",
  "@odata.type": "#ComputerSystem.v1_16_0.ComputerSystem",
  "AssetTag": "---",
  "Bios": {
    "@odata.id": "/redfish/v1/Systems/System_0/Bios"
  },
  "Description": "System Self",
  "EthernetInterfaces": {
    "@odata.id": "/redfish/v1/Systems/System_0/EthernetInterfaces"
  },
  "GraphicsControllers": {
    "@odata.id": "/redfish/v1/Systems/System_0/GraphicsControllers"
  },
  "HostName": null,
  "HostingRoles": [
    "ApplicationServer"
  ],
  "Id": "System_0",
  "IndicatorLED": "Off",
  "IndicatorLED@Redfish.AllowableValues": [
    "Lit",
    "Blinking",
    "Off"
  ],
  "MemorySummary": {
    "MemoryMirroring": null,
    "Metrics": {
      "@odata.id": "/redfish/v1/Systems/System_0/MemorySummary/MemoryMetrics"
    },
    "Status": {
      "Health": "OK",
      "HealthRollup": null,
      "State": "Enabled"
    },
    "TotalSystemMemoryGiB": 0,
    "TotalSystemPersistentMemoryGiB": 0
  },
  "Model": "QuantaGrid S74G-2U 1S7GZ9Z0001",
  "Name": "System",
  "NetworkInterfaces": {
    "@odata.id": "/redfish/v1/Systems/System_0/NetworkInterfaces"
  },
  "Oem": {
    "Ami": {
      "@odata.type": "#AMIBIOSInventoryCRC.v1_0_0.AMIBIOSInventoryCRC",
      "Bios": {
        "Inventory": {
          "Crc": {
            "@odata.id": "/redfish/v1/Systems/System_0/Oem/Ami/Inventory/Crc",
            "GroupCrcList": [
              {
                "PCIE": 17555861
              },
              {
                "CERTIFICATE": 0
              },
              {
                "CPU": 2772866038
              },
              {
                "DIMM": 3469537117
              },
              {
                "SECUREBOOT": 2614701783
              }
            ]
          }
        },
        "RedfishVersion": "1.15.1",
        "RtpVersion": "RB_1.0.17"
      },
      "ManagerBootConfiguration": {
        "ManagerBootMode": "None",
        "ManagerBootMode@Redfish.AllowableValues": [
          "None",
          "SoftReset",
          "ResetTimeout"
        ]
      },
      "SSIFMode": "Enabled"
    }
  },
  "PCIeFunctions@odata.count": 21,
  "PartNumber": "1S7GZ9Z0001",
  "PowerRestorePolicy": "LastState",
  "PowerState": "On",
  "ProcessorSummary": {
    "CoreCount": 72,
    "Count": 1,
    "Model": null,
    "Status": {
      "Health": "OK",
      "HealthRollup": null,
      "State": "Enabled"
    }
  },
  "Processors": {
    "@odata.id": "/redfish/v1/Systems/System_0/Processors"
  },
  "SKU": "Default string",
  "SecureBoot": {
    "@odata.id": "/redfish/v1/Systems/System_0/SecureBoot"
  },
  "SerialNumber": "QTWS7G0234700030",
  "SimpleStorage": {
    "@odata.id": "/redfish/v1/Systems/System_0/SimpleStorage"
  },
  "Status": {
    "Health": "OK",
    "HealthRollup": "OK",
    "State": "Enabled"
  },
  "Storage": {
    "@odata.id": "/redfish/v1/Systems/System_0/Storage"
  },
  "SubModel": null,
  "SystemType": "Physical",
  "TrustedModules": [
    {
      "FirmwareVersion": "7.85",
      "FirmwareVersion2": "17.51968",
      "InterfaceType": "TPM2_0",
      "InterfaceTypeSelection": "BiosSetting",
      "Status": {
        "Health": null,
        "HealthRollup": null,
        "State": "Enabled"
      }
    }
  ],
  "USBControllers": {
    "@odata.id": "/redfish/v1/Systems/System_0/USBControllers"
  },
  "UUID": "C7CEB99A-0133-11EE-B226-74D4DD2E8868"
}`

// TestComputerSystem tests the parsing of ComputerSystem objects.
func TestComputerSystem(t *testing.T) {
	var cs redfish.ComputerSystem
	err := json.NewDecoder(strings.NewReader(computerSystemBody)).Decode(&cs)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromComputerSystem(&cs)
	if err != nil {
		t.Errorf("Error converting Redfish ComputerSystem to AMI ComputerSystem: %s", err)
	}

	if result.ID != "System_0" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "System" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.BIOS.Inventory.Crc.GroupCrcList[0]["PCIE"] != 17555861 {
		t.Errorf("Received invalid PCIe value: %d", result.BIOS.Inventory.Crc.GroupCrcList[0]["PCIE"])
	}

	if result.ManagerBootConfiguration.ManagerBootMode != NoneManagerBootMode {
		t.Errorf("Received invalid ManagerBootMode: %s", result.ManagerBootConfiguration.ManagerBootMode)
	}

	if result.SSIFMode != "Enabled" {
		t.Errorf("Received invalid SSIFMode: %s", result.SSIFMode)
	}
}
