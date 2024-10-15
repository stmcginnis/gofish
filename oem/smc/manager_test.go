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

var managerBody = `{
  "@odata.type": "#Manager.v1_11_0.Manager",
  "@odata.id": "/redfish/v1/Managers/1",
  "Id": "1",
  "Name": "Manager",
  "Description": "BMC",
  "ManagerType": "BMC",
  "UUID": "00000000-0000-0000-0000-7CC25586E000",
  "Model": "ASPEED",
  "FirmwareVersion": "01.01.06",
  "DateTime": "2024-10-15T21:35:01Z",
  "DateTimeLocalOffset": "+00:00",
  "Actions": {
    "Oem": {
      "#SmcManagerConfig.Reset": {
        "target": "/redfish/v1/Managers/1/Actions/Oem/SmcManagerConfig.Reset",
        "@Redfish.ActionInfo": "/redfish/v1/Managers/1/Oem/Supermicro/ResetActionInfo"
      }
    },
    "#Manager.ResetToDefaults": {
      "target": "/redfish/v1/Managers/1/Actions/Manager.ResetToDefaults",
      "@Redfish.ActionInfo": "/redfish/v1/Managers/1/ResetToDefaultsActionInfo"
    },
    "#Manager.Reset": {
      "target": "/redfish/v1/Managers/1/Actions/Manager.Reset",
      "ResetType@Redfish.AllowableValues": [
        "GracefulRestart"
      ]
    }
  },
  "Oem": {
    "Supermicro": {
      "@odata.type": "#SmcManagerExtensions.v1_0_0.Manager",
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
      "Syslog": {
        "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/Syslog"
      },
      "SysLockdown": {
        "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/SysLockdown"
      },
      "MemoryPFA": {
        "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MemoryPFA"
      },
      "MemoryHealthComp": {
        "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/MemoryHealthComp"
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
      "LLDP": {
        "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/LLDP"
      },
      "LicenseManager": {
        "@odata.id": "/redfish/v1/Managers/1/LicenseManager"
      }
    }
  },
  "@odata.etag": "\"70889f859f8f0399a6f71ffb167c1dc1\""
}`

// TestManager tests the parsing of Manager objects.
func TestManager(t *testing.T) {
	var m redfish.Manager
	err := json.NewDecoder(strings.NewReader(managerBody)).Decode(&m)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromManager(&m)
	if err != nil {
		t.Errorf("Error converting Redfish Manager to SMC Manager: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Manager" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.radius != "/redfish/v1/Managers/1/Oem/Supermicro/RADIUS" {
		t.Errorf("Invalid RADIUS link: %s", result.radius)
	}

	if result.mouseMode != "/redfish/v1/Managers/1/Oem/Supermicro/MouseMode" {
		t.Errorf("Invalid MouseMode link: %s", result.mouseMode)
	}

	if result.managerConfigResetTarget != "/redfish/v1/Managers/1/Actions/Oem/SmcManagerConfig.Reset" {
		t.Errorf("Invalid ManagerConfigResetTarget link: %s", result.managerConfigResetTarget)
	}
}
