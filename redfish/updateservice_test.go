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
    "Oem": {
        "VendorName": {
            "OemInfo1": "The Oem info 1.",
            "OemInfoN": "The Oem info N."
        }
    },
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

	// test oem
	switch oem := result.Oem.(type) {
	case map[string]interface{}:
		for vendor, values := range oem {
			if vendor != "VendorName" {
				t.Errorf("Received invalid Oem vendor: %s", vendor)
			}
			switch val := values.(type) {
			case map[string]interface{}:
				for k, v := range val {
					if k != "OemInfo1" && k != "OemInfoN" {
						t.Errorf("Received invalid Oem key %s for vendor: %s", k, vendor)
					}
					if k == "OemInfo1" && v != "The Oem info 1." {
						t.Errorf("Received invalid OemInfo1: %s", v)
					}
					if v == "OemInfoN" && v != "The Oem info N." {
						t.Errorf("Received invalid OemInfoN: %s", v)
					}
				}
			default:
				t.Error("Invalid Oem values")
			}
		}
	default:
		t.Error("Received invalid Oem")
	}
}
