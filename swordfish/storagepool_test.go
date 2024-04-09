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

var storagePoolBody = `{
		"@odata.context": "/redfish/v1/$metadata#StoragePool.StoragePool",
		"@odata.type": "#StoragePool.v1_2_0.StoragePool",
		"@odata.id": "/redfish/v1/StoragePool",
		"Id": "StoragePool-1",
		"Name": "StoragePoolOne",
		"Description": "StoragePool One",
		"AllocatedPools": {
			"@odata.id": "/redfish/v1/StoragePool/1"
		},
		"AllocatedVolumes": {
			"@odata.id": "/redfish/v1/Volume/1"
		},
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
		"ClassesOfService": {
			"@odata.id": "/redfish/v1/ClassesOfService"
		},
		"Compressed": true,
		"Deduplicated": true,
		"DefaultClassOfService": {
			"@odata.id": "/redfish/v1/ClassesOfService/1"
		},
		"Encrypted": false,
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
		"MaxBlockSizeBytes": 2199023255600,
		"RecoverableCapacitySourceCount": 0,
		"RemainingCapacityPercent": 24,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		}
	}`

// TestStoragePool tests the parsing of StoragePool objects.
func TestStoragePool(t *testing.T) {
	var result StoragePool
	err := json.NewDecoder(strings.NewReader(storagePoolBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "StoragePool-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "StoragePoolOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.allocatedPools != "/redfish/v1/StoragePool/1" {
		t.Errorf("Invalid allocated pool link: %s", result.allocatedPools)
	}

	if result.MaxBlockSizeBytes != 2199023255600 {
		t.Errorf("Invalid max block size: %d", result.MaxBlockSizeBytes)
	}
}

// TestStoragePoolUpdate tests the Update call.
func TestStoragePoolUpdate(t *testing.T) {
	var result StoragePool
	err := json.NewDecoder(strings.NewReader(storagePoolBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.Compressed = false
	result.Deduplicated = false
	result.Encrypted = false
	result.RecoverableCapacitySourceCount = 2
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Compressed:false") {
		t.Errorf("Unexpected Compressed update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Deduplicated:false") {
		t.Errorf("Unexpected Deduplicated update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Encrypted") {
		t.Errorf("Unexpected Encrypted update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "RecoverableCapacitySourceCount:2") {
		t.Errorf("Unexpected RecoverableCapacitySourceCount update payload: %s", calls[0].Payload)
	}
}
