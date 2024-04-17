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

var simpleUpdateBody = `{
    "@odata.context": "/redfish/v1/$metadata#UpdateService.UpdateService",
    "@odata.id": "/redfish/v1/UpdateService",
    "@odata.type": "#UpdateService.v1_6_0.UpdateService",
    "Actions": {
        "#UpdateService.SimpleUpdate": {
            "TransferProtocol@Redfish.AllowableValues": [
                "HTTP"
            ],
            "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate"
        },
        "Oem": {
            "DellUpdateService.v1_0_0#DellUpdateService.Install": {
                "InstallUpon@Redfish.AllowableValues": [
                    "Now",
                    "NowAndReboot",
                    "NextReboot"
                ],
                "target": "/redfish/v1/UpdateService/Actions/Oem/DellUpdateService.Install"
            }
        }
    },
    "Description": "Represents the properties for the Update Service",
    "FirmwareInventory": {
        "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory"
    },
    "HttpPushUri": "/redfish/v1/UpdateService/FirmwareInventory",
    "Id": "UpdateService",
    "Name": "Update Service"
}`

func TestUpdateService(t *testing.T) {
	var result UpdateService
	assertMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Check default redfish fields", func(t *testing.T) {
		c := &common.TestClient{}
		result.SetClient(c)

		err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.firmwareInventory, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.HTTPPushURI, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.TransferProtocol[0], "HTTP")
		assertMessage(t, result.simpleUpdateTarget, "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate")
	})
}

var startUpdateBody = `{
    "@odata.type": "#UpdateService.v1_8_0.UpdateService",
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
      "Oem": {},
      "#UpdateService.SimpleUpdate": {
        "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate",
        "@Redfish.ActionInfo": "/redfish/v1/UpdateService/SimpleUpdateActionInfo"
      },
      "#UpdateService.StartUpdate": {
        "target": "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate"
      }
    },
    "Oem": {}
    }
  }`

func TestUpdateServiceStartUpdate(t *testing.T) {
	var result UpdateService
	assertMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Check UpdateService.StartUpdate field", func(t *testing.T) {
		c := &common.TestClient{}
		result.SetClient(c)

		err := json.NewDecoder(strings.NewReader(startUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.startUpdateTarget, "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate")
	})
}
