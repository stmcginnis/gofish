//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var fileSystemBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#FileSystem.FileSystem",
		"@odata.type": "#FileSystem.v1_2_1.FileSystem",
		"@odata.id": "/redfish/v1/FileSystem",
		"Id": "FileSystem-1",
		"Name": "FileSystemOne",
		"Description": "FileSystem One",
		"AccessCapabilities": [
			"Read",
			"Write",
			"Append",
			"Streaming"
		],
		"BlockSizeBytes": 2147483648000,
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
		"CapacitySources": [{
			"@odata.context": "/redfish/v1/$metadata#Capacity.CapacitySource",
			"@odata.type": "#Capacity.v1_1_1.CapacitySource",
			"@odata.id": "/redfish/v1/CapacitySource",
			"Id": "Capacity-1",
			"Name": "CapacityOne",
			"Description": "Capacity One",
			"ProvidedCapacity": {
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
			"ProvidingDrives": {
				"@odata.id": "/redfish/v1/System/System-1/Drives"
			}
		}],
		"CapacitySources@odata.count": 1,
		"CasePreserved": true,
		"CaseSensitive": true,
		"CharacterCodeSet": [
			"Unicode",
			"UTF_8",
			"UTF_16"
		],
		"ClusterSizeBytes": 512,
		"ExportedShares": {
			"@odata.id": "/redfish/v1/Shares/1"
		},
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
			"ClassOfService": {
				"@odata.id": "/redfish/v1/ClassOfService/1"
			},
			"SpareResourceSets": [{
				"@odata.id": "/redfish/v1/Spares/1"
			}]
		},
		"LowSpaceWarningThresholdPercents": [
			25,
			10,
			15,
			5
		],
		"MaxFileNameLengthBytes": 8,
		"RecoverableCapacitySourceCount": 2,
		"RemainingCapacityPercent": 43,
		"ReplicaTargets": [],
		"ReplicaTargets@odata.count": 0
	}`)

// TestFileSystem tests the parsing of FileSystem objects.
func TestFileSystem(t *testing.T) {
	var result FileSystem
	err := json.NewDecoder(fileSystemBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "FileSystem-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "FileSystemOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessCapabilities[3] != StreamingStorageAccessCapability {
		t.Errorf("Invalid access capability: %s", result.AccessCapabilities[3])
	}

	if result.BlockSizeBytes != 2147483648000 {
		t.Errorf("Invalid BlockSizeBytes: %d", result.BlockSizeBytes)
	}

	if result.Capacity.Data.AllocatedBytes != 2199023255600 {
		t.Errorf("Capacity allocated bytes was %d", result.Capacity.Data.AllocatedBytes)
	}

	if len(result.CapacitySources) != 1 {
		t.Errorf("Expected 1 CapacitySource, got %d", len(result.CapacitySources))
	}

	if !result.CasePreserved {
		t.Error("CasePreserved should be true")
	}

	if !result.CaseSensitive {
		t.Error("CaseSensitive should be true")
	}

	if result.CharacterCodeSet[1] != UTF8CharacterCodeSet {
		t.Errorf("Invalid character code set: %s", result.CharacterCodeSet[1])
	}

	if result.RemainingCapacityPercent != 43 {
		t.Errorf("Invalid RemainingCapacityPercent: %d", result.RemainingCapacityPercent)
	}
}
