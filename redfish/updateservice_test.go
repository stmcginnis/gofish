//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish/dell"
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
		c := &common.TestClient{Vendor: "XXX"}
		result.Client = c

		err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}
		assertMessage(t, result.FirmwareInventory, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.HTTPPushURI, "/redfish/v1/UpdateService/FirmwareInventory")
		assertMessage(t, result.TransferProtocol[0], "HTTP")
		assertMessage(t, result.UpdateServiceTarget, "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate")
	})

	t.Run("Check oem redfish fields - dell", func(t *testing.T) {
		c := &common.TestClient{Vendor: "Dell"}
		result.Client = c
		var installUpon []string = []string{"Now", "NowAndReboot", "NextReboot"}

		err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		assertMessage(t, result.Oem.(dell.DellUpdateService).Actions.Target, "/redfish/v1/UpdateService/Actions/Oem/DellUpdateService.Install")

		if !reflect.DeepEqual(result.Oem.(dell.DellUpdateService).Actions.InstallUpon, installUpon) {
			t.Errorf("got %v, want %v", result.Oem.(dell.DellUpdateService).Actions.InstallUpon, installUpon)
		}
	})

}
