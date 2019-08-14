//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageReplicaInfoBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#StorageReplicaInfo.StorageReplicaInfo",
		"@odata.type": "#StorageReplicaInfo.v1_1_1.StorageReplicaInfo",
		"@odata.id": "/redfish/v1/StorageReplicaInfo",
		"Id": "StorageReplicaInfo-1",
		"Name": "StorageReplicaInfoOne",
		"Description": "StorageReplicaInfo One"
	}`)

var replicaInfoBody = strings.NewReader(
	`{
		"ConsistencyEnabled": true,
		"ConsistencyState": "Consistent",
		"ConsistencyStatus": "Consistent",
		"ConsistencyType": "SequentiallyConsistent",
		"DataProtectionLineOfService": [{
			"@odata.id": "/redfish/v1/DataProtectionLineOfService/1"
		}],
		"FailedCopyStopsHostIO": false,
		"PercentSynced": 21,
		"Replica": {
			"@odata.id": "/redfish/v1/Replica/1"
		},
		"ReplicaPriority": "High",
		"ReplicaProgressStatus": "Completed",
		"ReplicaReadOnlyAccess": "ReplicaElement",
		"ReplicaRecoveryMode": "Manual",
		"ReplicaSkewBytes": 1000,
		"ReplicaState": "Synchronized",
		"ReplicaType": "Clone",
		"ReplicaUpdateMode": "Asynchronous",
		"RequestedReplicaState": "Synchronized",
		"SyncMaintained": true,
		"UndiscoveredElement": "ReplicaElement",
		"WhenActivated": "2007-04-06T00:00",
		"WhenEstablished": "2007-04-06T01:00",
		"WhenSynced": "2007-04-06T01:01"
	}`)

// TestStorageReplicaInfo tests the parsing of StorageReplicaInfo objects.
func TestStorageReplicaInfo(t *testing.T) {
	var result StorageReplicaInfo
	err := json.NewDecoder(storageReplicaInfoBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "StorageReplicaInfo-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "StorageReplicaInfoOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}
}

// TestReplicaInfo tests the parsing of StorageReplicaInfo objects.
func TestReplicaInfo(t *testing.T) {
	var result ReplicaInfo
	err := json.NewDecoder(replicaInfoBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if !result.ConsistencyEnabled {
		t.Error("ConsistencyEnabled should be true")
	}

	if result.ConsistencyState != ConsistentConsistencyState {
		t.Errorf("Invalid ConsistencyState: %s", result.ConsistencyState)
	}

	if result.ConsistencyStatus != ConsistentConsistencyStatus {
		t.Errorf("Invalid ConsistencyStatus: %s", result.ConsistencyStatus)
	}

	if result.ConsistencyType != SequentiallyConsistentConsistencyType {
		t.Errorf("Invalid ConsistencyType: %s", result.ConsistencyType)
	}

	if result.PercentSynced != 21 {
		t.Errorf("Invalid percent synced: %d", result.PercentSynced)
	}

	if result.ReplicaPriority != HighReplicaPriority {
		t.Errorf("Invalid replica priority: %s", result.ReplicaPriority)
	}

	if result.ReplicaProgressStatus != CompletedReplicaProgressStatus {
		t.Errorf("Invalid ReplicaProgressStatus: %s", result.ReplicaProgressStatus)
	}

	if result.ReplicaReadOnlyAccess != ReplicaElementReplicaReadOnlyAccess {
		t.Errorf("Invalid ReplicaReadOnlyAccess: %s", result.ReplicaReadOnlyAccess)
	}

	if result.ReplicaRecoveryMode != ManualReplicaRecoveryMode {
		t.Errorf("Invalid ReplicaRecoverMode: %s", result.ReplicaRecoveryMode)
	}

	if result.ReplicaSkewBytes != 1000 {
		t.Errorf("Invalid ReplicaSkewBytes: %d", result.ReplicaSkewBytes)
	}

	if result.ReplicaState != SynchronizedReplicaState {
		t.Errorf("Invalid ReplicaState: %s", result.ReplicaState)
	}

	if result.ReplicaType != CloneReplicaType {
		t.Errorf("Invalid ReplicaType: %s", result.ReplicaState)
	}

	if result.ReplicaUpdateMode != AsynchronousReplicaUpdateMode {
		t.Errorf("Invalid ReplicaUpdateMode: %s", result.ReplicaUpdateMode)
	}

	if result.RequestedReplicaState != SynchronizedReplicaState {
		t.Errorf("Invalid RequestedReplicaState: %s", result.RequestedReplicaState)
	}

	if !result.SyncMaintained {
		t.Error("SyncMaintained should be true")
	}

	if result.UndiscoveredElement != ReplicaElementUndiscoveredElement {
		t.Errorf("Invalid UndiscoveredElement: %s", result.UndiscoveredElement)
	}
}
