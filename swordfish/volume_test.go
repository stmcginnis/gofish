//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/common"
)

var volumeBody = `{
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
		"Compressed": false,
		"Deduplicated": false,
		"DisplayName": "Test Volume 1",
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
		"MediaSpanCount": 1,
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
		"ProvisioningPolicy": "Thin",
		"RAIDType": "RAID10",
		"ReadCachePolicy": "ReadAhead",
		"RecoverableCapacitySourceCount": 2,
		"StripSizeBytes": 1024,
		"VolumeUsage": "Data",
		"WriteCachePolicy": "WriteThrough",
		"WriteCacheState": "Protected",
		"WriteHoleProtectionPolicy": "Off",
		"RemainingCapacityPercent": 24,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"StorageGroups": [{
			"@odata.id": "/redfish/v1/StorageGroups/1"
		}],
		"Actions": {
			"#Volume.AssignReplicaTarget": {
				"target": "/redfish/v1/Volume/Actions/Volume.AssignReplicaTarget"
			},
			"#Volume.CheckConsistency": {
				"target": "/redfish/v1/Volume/Actions/Volume.CheckConsistency"
			},
			"#Volume.CreateReplicaTarget": {
				"target": "/redfish/v1/Volume/Actions/Volume.CreateReplicaTarget"
			},
			"#Volume.Initialize": {
				"target": "/redfish/v1/Volume/Actions/Volume.Initialize"
			},
			"#Volume.RemoveReplicaRelationship": {
				"target": "/redfish/v1/Volume/Actions/Volume.RemoveReplicaRelationship"
			},
			"#Volume.ResumeReplication": {
				"target": "/redfish/v1/Volume/Actions/Volume.ResumeReplication"
			},
			"#Volume.ReverseReplicationRelationship": {
				"target": "/redfish/v1/Volume/Actions/Volume.ReverseReplicationRelationship"
			},
			"#Volume.SplitReplication": {
				"target": "/redfish/v1/Volume/Actions/Volume.SplitReplication"
			},
			"#Volume.SuspendReplication": {
				"target": "/redfish/v1/Volume/Actions/Volume.SuspendReplication"
			}
		}
	}`

// TestVolume tests the parsing of Volume objects.
func TestVolume(t *testing.T) {
	var result Volume
	err := json.NewDecoder(strings.NewReader(volumeBody)).Decode(&result)

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

	if result.assignReplicaTargetTarget != "/redfish/v1/Volume/Actions/Volume.AssignReplicaTarget" {
		t.Errorf("Invalid AssignReplicaTarget target: %s", result.assignReplicaTargetTarget)
	}

	if result.checkConsistencyTarget != "/redfish/v1/Volume/Actions/Volume.CheckConsistency" {
		t.Errorf("Invalid CheckConsistency target: %s", result.checkConsistencyTarget)
	}

	if result.createReplicaTargetTarget != "/redfish/v1/Volume/Actions/Volume.CreateReplicaTarget" {
		t.Errorf("Invalid CreateReplicaTarget target: %s", result.createReplicaTargetTarget)
	}

	if result.initializeTarget != "/redfish/v1/Volume/Actions/Volume.Initialize" {
		t.Errorf("Invalid Initialize target: %s", result.initializeTarget)
	}

	if result.removeReplicaRelationshipTarget != "/redfish/v1/Volume/Actions/Volume.RemoveReplicaRelationship" {
		t.Errorf("Invalid RemoveReplicaRelationship target: %s", result.removeReplicaRelationshipTarget)
	}

	if result.resumeReplicationTarget != "/redfish/v1/Volume/Actions/Volume.ResumeReplication" {
		t.Errorf("Invalid ResumeReplication target: %s", result.resumeReplicationTarget)
	}

	if result.reverseReplicationRelationshipTarget != "/redfish/v1/Volume/Actions/Volume.ReverseReplicationRelationship" {
		t.Errorf("Invalid ReverseReplicationRelationship target: %s", result.reverseReplicationRelationshipTarget)
	}

	if result.splitReplicationTarget != "/redfish/v1/Volume/Actions/Volume.SplitReplication" {
		t.Errorf("Invalid SplitReplication target: %s", result.splitReplicationTarget)
	}

	if result.suspendReplicationTarget != "/redfish/v1/Volume/Actions/Volume.SuspendReplication" {
		t.Errorf("Invalid SuspendReplication target: %s", result.suspendReplicationTarget)
	}
}

// TestVolumeUpdate tests the Update call.
func TestVolumeUpdate(t *testing.T) {
	var result Volume
	err := json.NewDecoder(strings.NewReader(volumeBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.CapacityBytes = 2199023255600
	result.Compressed = false
	result.Deduplicated = false
	result.DisplayName = "Testing123"
	result.Encrypted = false
	result.ProvisioningPolicy = FixedProvisioningPolicy
	result.ReadCachePolicy = OffReadCachePolicyType
	result.RecoverableCapacitySourceCount = 2
	result.StripSizeBytes = 1024
	result.WriteCachePolicy = OffWriteCachePolicyType
	result.WriteHoleProtectionPolicy = OEMWriteHoleProtectionPolicyType
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if strings.Contains(calls[0].Payload, "CapacityBytes") {
		t.Errorf("Unexpected CapacityBytes update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Compressed") {
		t.Errorf("Unexpected Compressed update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Deduplicated") {
		t.Errorf("Unexpected Deduplicated update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "DisplayName:Testing123") {
		t.Errorf("Unexpected DisplayName update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Encrypted:false") {
		t.Errorf("Unexpected Encrypted update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "ProvisioningPolicy:Fixed") {
		t.Errorf("Unexpected ProvisioningPolicy update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "ReadCachePolicy:Off") {
		t.Errorf("Unexpected ReadCachePolicy update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "RecoverableCapacitySourceCount") {
		t.Errorf("Unexpected RecoverableCapacitySourceCount update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "StripSizeBytes") {
		t.Errorf("Unexpected StripSizeBytes update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "WriteCachePolicy:Off") {
		t.Errorf("Unexpected WriteCachePolicy update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "WriteHoleProtectionPolicy:Oem") {
		t.Errorf("Unexpected WriteHoleProtectionPolicy update payload: %s", calls[0].Payload)
	}
}
