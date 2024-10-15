//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

var chassisBody = `{
  "@odata.type": "#Chassis.v1_14_0.Chassis",
  "@odata.id": "/redfish/v1/Chassis/1",
  "Id": "1",
  "Name": "Computer System Chassis",
  "ChassisType": "RackMount",
  "Manufacturer": "Supermicro",
  "Model": "X13DEG-OAD",
  "SerialNumber": "C8000MN18AC0000",
  "PartNumber": "CSE-GP801TS-R000NPF",
  "AssetTag": null,
  "IndicatorLED": "Off",
  "LocationIndicatorActive": false,
  "MaxPowerWatts": 18000,
  "PowerState": "On",
  "Status": {
    "State": "Enabled",
    "Health": "OK",
    "HealthRollup": "OK"
  },
  "Power": {
    "@odata.id": "/redfish/v1/Chassis/1/Power"
  },
  "PCIeDevices": {
    "@odata.id": "/redfish/v1/Chassis/1/PCIeDevices"
  },
  "Thermal": {
    "@odata.id": "/redfish/v1/Chassis/1/Thermal"
  },
  "NetworkAdapters": {
    "@odata.id": "/redfish/v1/Chassis/1/NetworkAdapters"
  },
  "PCIeSlots": {
    "@odata.id": "/redfish/v1/Chassis/1/PCIeSlots"
  },
  "Sensors": {
    "@odata.id": "/redfish/v1/Chassis/1/Sensors"
  },
  "Memory": {
    "@odata.id": "/redfish/v1/Systems/1/Memory"
  },
  "Links": {
    "ComputerSystems": [
      {
        "@odata.id": "/redfish/v1/Systems/1"
      }
    ],
    "ManagedBy": [
      {
        "@odata.id": "/redfish/v1/Managers/1"
      }
    ],
    "ManagersInChassis": [
      {
        "@odata.id": "/redfish/v1/Managers/1"
      }
    ]
  },
  "Oem": {
    "Supermicro": {
      "@odata.type": "#SmcChassisExtensions.v1_0_0.Chassis",
      "BoardSerialNumber": "HM241S008000",
      "GUID": "35353031-4D53-0000-0000-E49200000000",
      "BoardID": "0x1d14"
    }
  },
  "@odata.etag": "\"41b7a1190aa686f7ac4cf7abad27a8a2\""
}`

// TestChassis tests the parsing of Chassis objects.
func TestChassis(t *testing.T) {
	var cs redfish.Chassis
	err := json.NewDecoder(strings.NewReader(chassisBody)).Decode(&cs)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromChassis(&cs)
	if err != nil {
		t.Errorf("Error converting Redfish Chassis to SMC Chassis: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Computer System Chassis" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.BoardSerialNumber != "HM241S008000" {
		t.Errorf("Unexpected board serial number: %s", result.BoardSerialNumber)
	}

	if result.GUID != "35353031-4D53-0000-0000-E49200000000" {
		t.Errorf("Unexpected GUID: %s", result.GUID)
	}

	if result.BoardID != "0x1d14" {
		t.Errorf("Unexpected board ID: %s", result.BoardID)
	}
}
