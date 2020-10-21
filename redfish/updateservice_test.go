//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
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
    "Name": "Update Service",
    "ServiceEnabled": true,
    "Status": {
        "Health": "OK",
        "State": "Enabled"
    }
}`

func TestUpdateService(t *testing.T) {
	var result UpdateService
	err := json.NewDecoder(strings.NewReader(simpleUpdateBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.FirmwareInventory != "/redfish/v1/UpdateService/FirmwareInventory" {
		t.Errorf("FirmwareInventory was wrong")
	}

	if result.HTTPPushURI != "/redfish/v1/UpdateService/FirmwareInventory" {
		t.Errorf("HTTPPushURI was wrong")
	}

	if result.TransferProtocol[0] != "HTTP" {
		t.Errorf("TransferProtocol was wrong")
	}

	if result.UpdateServiceTarget != "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate" {
		t.Errorf("UpdateServiceTarget was wrong")
	}
}
