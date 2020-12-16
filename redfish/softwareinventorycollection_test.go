//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var softwareCollectionInventoryBody = `{
	"@odata.context": "/redfish/v1/$metadata#SoftwareInventoryCollection.SoftwareInventoryCollection",
	"@odata.etag": "W/\"A3D58239\"",
	"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/",
	"@odata.type": "#SoftwareInventoryCollection.SoftwareInventoryCollection",
	"Description": "Firmware Inventory Collection",
	"Name": "Firmware Inventory Collection",
	"Members": [
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/1/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/2/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/3/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/4/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/5/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/6/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/7/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/8/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/9/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/10/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/11/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/12/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/13/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/14/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/15/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/16/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/17/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/18/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/19/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/20/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/21/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/22/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/23/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/24/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/25/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/26/"
	  },
	  {
		"@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/27/"
	  }
	],
	"Members@odata.count": 27
  }`

func TestSoftwareInventoryCollection(t *testing.T) {
	var result UpdateService
	err := json.NewDecoder(strings.NewReader(softwareCollectionInventoryBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
}
