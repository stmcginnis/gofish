//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var softwareInventoryBody = `{
	"@odata.context": "/redfish/v1/$metadata#SoftwareInventory.SoftwareInventory",
	"@odata.etag": "W/\"918DA6A4\"",
	"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/1/",
	"@odata.type": "#SoftwareInventory.v1_0_0.SoftwareInventory",
	"Id": "1",
	"Description": "SystemBMC",
	"Name": "Bob",
	"Oem": {},
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"RelatedItem": [
		{
		"@odata.id": "/redfish/v1/Managers/1"
		}
	],
	"Version": "1.2.3.4",
	"LowestSupportedVersion": "1.1.1.1",
	"Manufacturer": "BobCo",
	"ReleaseDate": "1/1/2020",
	"SoftwareId": "1234",
	"UefiDevicePaths": [ "/1", "/2" ],
	"Updateable": true,
	"WriteProtected": false
}`

func TestSoftwareInventory(t *testing.T) {
	var result SoftwareInventory
	err := json.NewDecoder(strings.NewReader(softwareInventoryBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.Description != "SystemBMC" {
		t.Errorf("Description was wrong")
	}

	if result.Status.Health != "OK" {
		t.Errorf("Health is wrong")
	}

	if result.Version != "1.2.3.4" {
		t.Errorf("Version is wrong")
	}

	if result.LowestSupportedVersion != "1.1.1.1" {
		t.Errorf("LowestSupportedVersion is wrong")
	}

	if result.Manufacturer != "BobCo" {
		t.Errorf("Manufacturer is wrong")
	}

	if len(result.relatedItem) != 1 || result.relatedItem[0] != "/redfish/v1/Managers/1" {
		t.Errorf("Unexpected related items: %#v", result.relatedItem)
	}

	if result.ReleaseDate != "1/1/2020" {
		t.Errorf("ReleaseDate is wrong")
	}

	if result.SoftwareID != "1234" {
		t.Errorf("SoftwareID is wrong")
	}

	if result.UefiDevicePaths[0] != "/1" {
		t.Errorf("UefiDevicePaths is wrong")
	}

	if result.Updateable != true {
		t.Errorf("Updateable is wrong")
	}

	if result.WriteProtected != false {
		t.Errorf("WriteProtected is wrong")
	}
}
