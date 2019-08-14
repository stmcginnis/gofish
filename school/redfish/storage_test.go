//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Storage.Storage",
		"@odata.type": "#Storage.v1_7_0.Storage",
		"@odata.id": "/redfish/v1/Storage",
		"Id": "Storage-1",
		"Name": "StorageOne",
		"Description": "Storage One",
		"Drives": [{
				"@odata.id": "/redfish/v1/Drive/1"
			},
			{
				"@odata.id": "/redfish/v1/Drive/2"
			},
			{
				"@odata.id": "/redfish/v1/Drive/3"
			},
			{
				"@odata.id": "/redfish/v1/Drive/4"
			},
			{
				"@odata.id": "/redfish/v1/Drive/5"
			},
			{
				"@odata.id": "/redfish/v1/Drive/6"
			}
		],
		"Drives@odata.count": 6,
		"Links": {
			"Enclosures": [{
				"@odata.id": "/redfish/v1/Enclosures/1"
			}],
			"Enclosures@odata.count": 1
		},
		"Redundancy": [],
		"Redundancy@odata.count": 0,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"StorageControllers": [{
			"@odata.id": "/redfish/v1/StorageController/1",
			"Assembly": {
				"@odata.id": "/redfish/v1/Assembly/1"
			},
			"AssetTag": "ABC123",
			"CacheSummary": {
				"PersistentCacheSizeMiB": 1024,
				"Status": {
					"State": "Enabled",
					"Health": "OK"
				},
				"TotalCacheSizeMiB": 1024
			},
			"ControllerRates": {
				"ConsistencyCheckRatePercent": 5,
				"RebuildRatePercent": 5,
				"TransformationRatePercent": 5
			},
			"FirmwareVersion": "1.0",
			"Identifiers": [],
			"Links": {
				"Endpoints": [{
					"@odata.id": "/redfish/v1/Endpoints/1"
				}],
				"Endpoints@odata.count": 1,
				"PCIeFunctions": [{
					"@odata.id": "/redfish/v1/Functions/1"
				}],
				"PCIeFunctions@odata.count": 1,
				"StorageServices": [{
					"@odata.id": "/redfish/v1/StorageServices/1"
				}],
				"StorageServices@odata.count": 1
			},
			"Location": {},
			"Manufacturer": "Acme Storage",
			"MemberId": "SS1",
			"Model": "Model One",
			"Name": "Storage Controller One",
			"PCIeInterface": {
				"LanesInUse": 32,
				"MaxLanes": 32,
				"MaxPCIeType": "Gen4",
				"PCIeType": "Gen4"
			},
			"PartNumber": "A123",
			"SKU": "12324",
			"SerialNumber": "12345",
			"SpeedGbps": 10,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"SupportedControllerProtocols": [],
			"SupportedRAIDTypes": [
				"RAID0",
				"RAID1",
				"RAID5",
				"RAID6",
				"RAID10"
			]
		}],
		"Volumes": {
			"@odata.id": "/redfish/v1/Volumes/1"
		}
	}`)

// TestStorage tests the parsing of Storage objects.
func TestStorage(t *testing.T) {
	var result Storage
	err := json.NewDecoder(storageBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Storage-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "StorageOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.drives) != 6 {
		t.Errorf("Unexpected number of drives: %d", len(result.drives))
	}

	if result.StorageControllers[0].CacheSummary.PersistentCacheSizeMiB != 1024 {
		t.Errorf("Invalid PersistenCacheSize: %d",
			result.StorageControllers[0].CacheSummary.PersistentCacheSizeMiB)
	}

	if result.StorageControllers[0].PCIeInterface.MaxPCIeType != Gen4PCIeTypes {
		t.Errorf("Invalid MaxPCIeType: %s", result.StorageControllers[0].PCIeInterface.MaxPCIeType)
	}
}
