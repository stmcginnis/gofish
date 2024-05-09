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

var storageBody = `{
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
		},
		"Actions": {
			"#Storage.SetEncryptionKey": {
				"target": "/redfish/v1/Storage/Actions/Storage.SetEncryptionKey"
			}
		}
	}`

var storageBodyDell = `{
	"@odata.context": "/redfish/v1/$metadata#Storage.Storage",
	"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1",
	"@odata.type": "#Storage.v1_15_0.Storage",
	"Controllers": {
	  "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1/Controllers"
	},
	"Description": "Embedded AHCI 1",
	"Drives": [],
	"Drives@odata.count": 0,
	"Id": "AHCI.Embedded.1-1",
	"Identifiers": [
	  {
		"DurableName": null,
		"DurableNameFormat": null
	  }
	],
	"Identifiers@odata.count": 1,
	"Links": {
	  "Enclosures": [
		{
		  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1"
		}
	  ],
	  "Enclosures@odata.count": 1,
	  "Oem": {
		"Dell": {
		  "@odata.type": "#DellOem.v1_3_0.DellOemLinks",
		  "CPUAffinity": [],
		  "CPUAffinity@odata.count": 0
		}
	  },
	  "SimpleStorage": {
		"@odata.id": "/redfish/v1/Systems/System.Embedded.1/SimpleStorage/AHCI.Embedded.1-1"
	  }
	},
	"Name": "Sapphire Rapids SATA AHCI Controller",
	"Oem": {
	  "Dell": {
		"@odata.type": "#DellOem.v1_3_0.DellOemResources",
		"DellController": {
		  "@odata.context": "/redfish/v1/$metadata#DellController.DellController",
		  "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1/Oem/Dell/DellControllers/AHCI.Embedded.1-1",
		  "@odata.type": "#DellController.v1_4_1.DellController",
		  "AlarmState": "AlarmNotPresent",
		  "AutoConfigBehavior": "NotApplicable",
		  "BootVirtualDiskFQDD": null,
		  "CacheSizeInMB": 0,
		  "CachecadeCapability": "NotSupported",
		  "ConnectorCount": 0,
		  "ControllerFirmwareVersion": null,
		  "CurrentControllerMode": "NotSupported",
		  "Description": "An instance of DellController will have RAID Controller specific data.",
		  "Device": "0",
		  "DeviceCardDataBusWidth": "Unknown",
		  "DeviceCardSlotLength": "Unknown",
		  "DeviceCardSlotType": "Unknown",
		  "DriverVersion": null,
		  "EncryptionCapability": "None",
		  "EncryptionMode": "None",
		  "Id": "AHCI.Embedded.1-1",
		  "KeyID": null,
		  "LastSystemInventoryTime": "2024-04-10T22:38:06+00:00",
		  "LastUpdateTime": "2024-02-20T05:22:02+00:00",
		  "MaxAvailablePCILinkSpeed": null,
		  "MaxPossiblePCILinkSpeed": null,
		  "Name": "DellController",
		  "PCISlot": null,
		  "PatrolReadState": "Unknown",
		  "PersistentHotspare": "NotApplicable",
		  "RealtimeCapability": "Incapable",
		  "RollupStatus": "Unknown",
		  "SASAddress": "0",
		  "SecurityStatus": "EncryptionNotCapable",
		  "SharedSlotAssignmentAllowed": "NotApplicable",
		  "SlicedVDCapability": "NotSupported",
		  "SupportControllerBootMode": "NotSupported",
		  "SupportEnhancedAutoForeignImport": "NotSupported",
		  "SupportRAID10UnevenSpans": "NotSupported",
		  "SupportsLKMtoSEKMTransition": "No",
		  "T10PICapability": "NotSupported"
		}
	  }
	},
	"Status": {
	  "Health": null,
	  "HealthRollup": null,
	  "State": "Enabled"
	},
	"StorageControllers": [
	  {
		"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1#/StorageControllers/0",
		"Assembly": {
		  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/Assembly"
		},
		"CacheSummary": {
		  "TotalCacheSizeMiB": 0
		},
		"ControllerRates": {
		  "ConsistencyCheckRatePercent": null,
		  "RebuildRatePercent": null
		},
		"FirmwareVersion": "",
		"Identifiers": [
		  {
			"DurableName": null,
			"DurableNameFormat": null
		  }
		],
		"Identifiers@odata.count": 1,
		"Links": {
		  "PCIeFunctions": [
			{
			  "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/PCIeDevices/0-24/PCIeFunctions/0-24-0"
			}
		  ],
		  "PCIeFunctions@odata.count": 1
		},
		"Manufacturer": "DELL",
		"MemberId": "0",
		"Model": "Sapphire Rapids SATA AHCI Controller",
		"Name": "Sapphire Rapids SATA AHCI Controller",
		"SpeedGbps": null,
		"Status": {
		  "Health": null,
		  "HealthRollup": null,
		  "State": "Enabled"
		},
		"SupportedControllerProtocols": [
		  "PCIe"
		],
		"SupportedControllerProtocols@odata.count": 1,
		"SupportedDeviceProtocols": [],
		"SupportedDeviceProtocols@odata.count": 0,
		"SupportedRAIDTypes": [],
		"SupportedRAIDTypes@odata.count": 0
	  }
	],
	"StorageControllers@Redfish.Deprecated": "Please migrate to use /redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1/Controllers",
	"StorageControllers@odata.count": 1,
	"Volumes": {
	  "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1/Volumes"
	}
  }`

// TestStorage tests the parsing of Storage objects.
func TestStorage(t *testing.T) {
	var result Storage
	err := json.NewDecoder(strings.NewReader(storageBody)).Decode(&result)

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

	if result.setEncryptionKeyTarget != "/redfish/v1/Storage/Actions/Storage.SetEncryptionKey" {
		t.Errorf("Invalid SetEncryptionKey target: %s", result.setEncryptionKeyTarget)
	}
}

func TestStorageDell(t *testing.T) {
	var result Storage
	err := json.NewDecoder(strings.NewReader(storageBodyDell)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
}

// TestStorageControllerUpdate tests the Update call.
func TestStorageControllerUpdate(t *testing.T) {
	var result Storage
	err := json.NewDecoder(strings.NewReader(storageBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	scResult := result.StorageControllers[0]
	scResult.AssetTag = TestAssetTag

	// TODO: This highlights an issue that child objects of an object do not
	// get their client set. Need to review objects like Storage that include
	// the full objects rather than just links to them.
	testClient := &common.TestClient{}
	scResult.SetClient(testClient)
	err = scResult.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AssetTag:TestAssetTag") {
		t.Errorf("Unexpected AssetTag update payload: %s", calls[0].Payload)
	}
}
