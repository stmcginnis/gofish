//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageGroupBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#StorageGroup.StorageGroup",
		"@odata.type": "#StorageGroup.v1_2_0.StorageGroup",
		"@odata.id": "/redfish/v1/StorageGroup",
		"Id": "StorageGroup-1",
		"Name": "StorageGroupOne",
		"Description": "StorageGroup One",
		"AuthenticationMethod": "MutualCHAP",
		"ChapInfo": [{
			"InitiatorCHAPPassword": "abc123",
			"InitiatorCHAPUser": "root",
			"TargetCHAPUser": "root",
			"TargetPassword": "password1"
		}],
		"ClientEndpointGroups": [{
			"@odata.id": "/redfish/v1/StorageGroup/Endpoints/1"
		}],
		"ClientEndpointGroups@odata.count": 1,
		"Links": {
			"ChildStorageGroups": [{
				"@odata.id": "/redfish/v1/StorageGroup/1"
			}],
			"ChildStorageGroups@odata.count": 1,
			"ClassOfService": {
				"@odata.id": "/redfish/v1/ClassOfService/1"
			},
			"ParentStorageGroups": [],
			"ParentStorageGroups@odata.count": 0
		},
		"MappedVolumes": [{
			"LogicalUnitNumber": 1,
			"Volume": {
				"@odata.id": "/redfish/v1/Volume/1"
			}
		}],
		"MembersAreConsistent": true,
		"ReplicaTargets": [],
		"ReplicaTargets@odata.count": 0,
		"ServerEndpointGroups": [{
			"@odata.id": "/redfish/v1/Server/1/Endpoints"
		}],
		"ServerEndpointGroups@odata.count": 1,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"VolumesAreExposed": true
	}`)

// TestStorageGroup tests the parsing of StorageGroup objects.
func TestStorageGroup(t *testing.T) {
	var result StorageGroup
	err := json.NewDecoder(storageGroupBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "StorageGroup-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "StorageGroupOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AuthenticationMethod != MutualCHAPAuthenticationMethod {
		t.Errorf("Invalid auth method: %s", result.AuthenticationMethod)
	}

	if result.ChapInfo[0].InitiatorCHAPPassword != "abc123" {
		t.Errorf("Invalid init CHAP password: %s", result.ChapInfo[0].InitiatorCHAPPassword)
	}

	if result.childStorageGroups[0] != "/redfish/v1/StorageGroup/1" {
		t.Errorf("Incorrect child storage groups link: %s", result.childStorageGroups[0])
	}

	if result.MappedVolumes[0].LogicalUnitNumber != 1 {
		t.Errorf("Invalid mapped volume LUN: %d", result.MappedVolumes[0].LogicalUnitNumber)
	}

	if !result.VolumesAreExposed {
		t.Error("VolumesAreExposed should be True")
	}
}
