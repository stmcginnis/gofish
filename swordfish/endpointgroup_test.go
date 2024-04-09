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

var endpointGroupBody = `{
		"@odata.context": "/redfish/v1/$metadata#EndpointGroup.EndpointGroup",
		"@odata.type": "#EndpointGroup.v1_1_2.EndpointGroup",
		"@odata.id": "/redfish/v1/EndpointGroup",
		"Id": "EndpointGroup-1",
		"Name": "EndpointGroupOne",
		"Description": "EndpointGroup One",
		"AccessState": "Optimized",
		"Endpoints": {
			"@odata.id": "/redfish/v1/Endpoints"
		},
		"GroupType": "Server",
		"Preferred": true,
		"TargetEndpointGroupIdentifier": 5
	}`

// TestEndpointGroup tests the parsing of EndpointGroup objects.
func TestEndpointGroup(t *testing.T) {
	var result EndpointGroup
	err := json.NewDecoder(strings.NewReader(endpointGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "EndpointGroup-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "EndpointGroupOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessState != OptimizedAccessState {
		t.Errorf("Access state is %s", result.AccessState)
	}

	if result.endpoints != "/redfish/v1/Endpoints" {
		t.Errorf("Invalid endpoints: %s", result.endpoints)
	}

	if result.GroupType != ServerGroupType {
		t.Errorf("Invalid group type: %s", result.GroupType)
	}

	if !result.Preferred {
		t.Error("Preferred should be true")
	}

	if result.TargetEndpointGroupIdentifier != 5 {
		t.Errorf("Invalid target endpoint group id: %d", result.TargetEndpointGroupIdentifier)
	}
}

// TestEndpointGroupUpdate tests the Update call.
func TestEndpointGroupUpdate(t *testing.T) {
	var result EndpointGroup
	err := json.NewDecoder(strings.NewReader(endpointGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.AccessState = StandbyAccessState
	result.GroupType = ClientGroupType
	result.Preferred = true
	result.TargetEndpointGroupIdentifier = 9
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "AccessState:Standby") {
		t.Errorf("Unexpected AccessState update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "GroupType:Client") {
		t.Errorf("Unexpected GroupType update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Preferred") {
		t.Errorf("Unexpected Preferred update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "TargetEndpointGroupIdentifier:9") {
		t.Errorf("Unexpected TargetEndpointGroupIdentifier update payload: %s", calls[0].Payload)
	}
}
