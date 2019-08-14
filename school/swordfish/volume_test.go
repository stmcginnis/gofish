//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var volumeBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Volume.Volume",
		"@odata.type": "#Volume.v1_3_1.Volume",
		"@odata.id": "/redfish/v1/Volume",
		"Id": "Volume-1",
		"Name": "VolumeOne",
		"Description": "Volume One",
		"AccessCapabilities": [
			"Read",
			"Write",
			"Append",
			"Streaming"
		],
		"AllocatedPools": [{
			"@odata.id": "/redfish/v1/Pools/1"
		}],
		"BlockSizeBytes": 512,
		"Capacity": {
			"Data": {
				"AllocatedBytes": 2199023255600,
				"ConsumedBytes": 2199023255600,
				"GuaranteedBytes": 2199023255600,
				"ProvisionedBytes": 2199023255600
			},
			"IsThinProvisioned": false,
			"Metadata": {
				"AllocatedBytes": 209715200,
				"ConsumedBytes": 2048,
				"GuaranteedBytes": 209715200,
				"ProvisionedBytes": 209715200
			},
			"Snapshot": {
				"AllocatedBytes": 0,
				"ConsumedBytes": 0,
				"GuaranteedBytes": 0,
				"ProvisionedBytes": 0
			}
		},
		"CapacityBytes": 2199023255600,
		"CapacitySources": [{
			"@odata.id": "/redfish/v1/CapacitySource/1"
		}],
		"CapacitySources@odata.count": 1,
		"Encrypted": true,
		"EncryptionTypes": [
			"ControllerAssisted",
			"SoftwareAssisted"
		],
		"IOStatistics": {
			"NonIORequestTime": "P0Y0M0DT0H0M5S",
			"NonIORequests": 1000,
			"ReadHitIORequests": 500,
			"ReadIOKiBytes": 1024,
			"ReadIORequestTime": "P0Y0M0DT0H0M5S",
			"ReadIORequests": 5000,
			"WriteHitIORequests": 100,
			"WriteIOKiBytes": 1024,
			"WriteIORequestTime": "P0Y0M0DT0H0M5S",
			"WriteIORequests": 5000
		},
		"Links": {
			"DedicatedSpareDrives": [{
				"@odata.id": "/redfish/v1/Drives/1"
			}],
			"DedicatedSpareDrives@odata.count": 1,
			"SpareResourceSets": [{
				"@odata.id": "/redfish/v1/Spares/1"
			}],
			"SpareResourceSets@odata.count": 1
		},
		"LowSpaceWarningThresholdPercents": [
			25,
			10,
			15,
			5
		],
		"Manufacturer": "Acme Storage",
		"MaxBlockSizeBytes": 2199023255600,
		"Model": "Kate Moss",
		"Operations": [{
				"OperationName": "Overlord",
				"PercentageComplete": 92
			},
			{
				"OperationName": "Rolling Thunder",
				"PercentageComplete": 55
			}
		],
		"OptimumIOSizeBytes": 1024,
		"RAIDType": "RAID10",
		"RecoverableCapacitySourceCount": 0,
		"RemainingCapacityPercent": 24,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"StorageGroups": [{
			"@odata.id": "/redfish/v1/StorageGroups/1"
		}]
	}`)

// TestVolume tests the parsing of Volume objects.
func TestVolume(t *testing.T) {
	var result Volume
	err := json.NewDecoder(volumeBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Volume-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "VolumeOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.allocatedPools[0] != "/redfish/v1/Pools/1" {
		t.Errorf("Invalid allocated pool link: %s", result.allocatedPools[0])
	}

	if result.MaxBlockSizeBytes != 2199023255600 {
		t.Errorf("Invalid max block size: %d", result.MaxBlockSizeBytes)
	}

	if result.BlockSizeBytes != 512 {
		t.Errorf("Invalid BlockSizeBytes: %d", result.BlockSizeBytes)
	}

	if result.MaxBlockSizeBytes != 2199023255600 {
		t.Errorf("Invalid MaxBlockSizeBytes: %d", result.MaxBlockSizeBytes)
	}

	if result.OptimumIOSizeBytes != 1024 {
		t.Errorf("Invalid OptimumIOSizeBytes: %d", result.OptimumIOSizeBytes)
	}

	if result.storageGroups[0] != "/redfish/v1/StorageGroups/1" {
		t.Errorf("Invalid StorageGroup link: %s", result.storageGroups[0])
	}
}
