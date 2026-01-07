//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var storageGroupBody = `{
		"@odata.context": "/redfish/v1/$metadata#StorageGroup.StorageGroup",
		"@odata.type": "#StorageGroup.v1_2_0.StorageGroup",
		"@odata.id": "/redfish/v1/StorageGroup",
		"Id": "StorageGroup-1",
		"Name": "StorageGroupOne",
		"Description": "StorageGroup One",
		"AccessState": "Optimized",
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
		"VolumesAreExposed": true,
		"Actions": {
			"#StorageGroup.ExposeVolumes": {
				"target": "/redfish/v1/StorageGroup/Actions/StorageGroup.ExposeVolumes"
			},
			"#StorageGroup.HideVolumes": {
				"target": "/redfish/v1/StorageGroup/Actions/StorageGroup.HideVolumes"
			}
		}
	}`

// TestStorageGroup tests the parsing of StorageGroup objects.
func TestStorageGroup(t *testing.T) {
	var result StorageGroup
	err := json.NewDecoder(strings.NewReader(storageGroupBody)).Decode(&result)

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

	if result.MappedVolumes[0].LogicalUnitNumber != "1" {
		t.Errorf("Invalid mapped volume LUN: %s", result.MappedVolumes[0].LogicalUnitNumber)
	}

	if !result.VolumesAreExposed {
		t.Error("VolumesAreExposed should be True")
	}

	if result.exposeVolumesTarget != "/redfish/v1/StorageGroup/Actions/StorageGroup.ExposeVolumes" {
		t.Errorf("Invalid ExposeVolumes target: %s", result.exposeVolumesTarget)
	}

	if result.hideVolumesTarget != "/redfish/v1/StorageGroup/Actions/StorageGroup.HideVolumes" {
		t.Errorf("Invalid HideVolumes target: %s", result.hideVolumesTarget)
	}
}

// TestStorageGroupUpdate tests the Update call.
func TestStorageGroupUpdate(t *testing.T) {
	var result StorageGroup
	err := json.NewDecoder(strings.NewReader(storageGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	result.AccessState = NonOptimizedAccessState
	result.AuthenticationMethod = NoneAuthenticationMethod
	result.VolumesAreExposed = true
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AccessState:NonOptimized") {
		t.Errorf("Unexpected AccessState update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "AuthenticationMethod:None") {
		t.Errorf("Unexpected AuthenticationMethod update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "VolumeAreExposed") {
		t.Errorf("Unexpected VolumeAreExposed update payload: %s", calls[0].Payload)
	}
}
