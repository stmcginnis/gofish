//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

var updateServiceBody = `{
  "@odata.type": "#UpdateService.v1_8_4.UpdateService",
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
    "Oem": {
      "#SmcUpdateService.Install": {
        "target": "/redfish/v1/UpdateService/Actions/Oem/SmcUpdateService.Install",
        "@Redfish.ActionInfo": "/redfish/v1/UpdateService/Oem/Supermicro/InstallActionInfo"
      }
    },
    "#UpdateService.SimpleUpdate": {
      "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate",
      "@Redfish.ActionInfo": "/redfish/v1/UpdateService/SimpleUpdateActionInfo"
    },
    "#UpdateService.StartUpdate": {
      "target": "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate"
    }
  },
  "Oem": {
    "Supermicro": {
      "@odata.type": "#SmcUpdateServiceExtensions.v1_0_0.UpdateService",
      "SSLCert": {
        "@odata.id": "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert"
      },
      "IPMIConfig": {
        "@odata.id": "/redfish/v1/UpdateService/Oem/Supermicro/IPMIConfig"
      }
    }
  },
  "@odata.etag": "\"e9b94401dae9992fef2e71ef30cbcfdc\""
}`

// TestSmcUpdateService tests the parsing of the UpdateService oem field
func TestSmcUpdateService(t *testing.T) {
	us := &redfish.UpdateService{}
	if err := json.Unmarshal([]byte(updateServiceBody), us); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	updateService, err := FromUpdateService(us)
	if err != nil {
		t.Fatalf("error getting oem object: %v", err)
	}

	if updateService.ID != "UpdateService" {
		t.Errorf("unexpected ID: %s", updateService.ID)
	}

	if updateService.installTarget != "/redfish/v1/UpdateService/Actions/Oem/SmcUpdateService.Install" {
		t.Errorf("unexpected install target: %s", updateService.installTarget)
	}

	if updateService.sslCert != "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert" {
		t.Errorf("unexpected ssl cert link: %s", updateService.installTarget)
	}

	if updateService.ipmiConfig != "/redfish/v1/UpdateService/Oem/Supermicro/IPMIConfig" {
		t.Errorf("unexpected ipmi config link: %s", updateService.installTarget)
	}
}
