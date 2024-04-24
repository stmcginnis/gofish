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

var endpointGroupBody = `{
	"@odata.type": "#EndpointGroup.v1_3_3.EndpointGroup",
	"Id": "1",
	"Name": "Endpoint group for all initiators",
	"GroupType": "Initiator",
	"Links": {
	  "Endpoints": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator1"
		},
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator2"
		}
	  ],
	  "Connections": [
		{
		  "@odata.id": "/redfish/v1/Fabrics/NVMeoF/Connections/3"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Fabrics/NVMeoF/EndpointGroups/1"
  }`

// TestEndpointGroup tests the parsing of EndpointGroup objects.
func TestEndpointGroup(t *testing.T) {
	var result EndpointGroup
	err := json.NewDecoder(strings.NewReader(endpointGroupBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "1", result.ID)
	assertEquals(t, "Endpoint group for all initiators", result.Name)
	assertEquals(t, "Initiator", string(result.GroupType))
	assertEquals(t, "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator1", result.endpoints[0])
	assertEquals(t, "/redfish/v1/Fabrics/NVMeoF/Endpoints/Initiator2", result.endpoints[1])
	assertEquals(t, "/redfish/v1/Fabrics/NVMeoF/Connections/3", result.connections[0])
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

	result.GroupType = TargetGroupType
	result.TargetEndpointGroupIdentifier = 3
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "GroupType:Target") {
		t.Errorf("Unexpected GroupType update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "TargetEndpointGroupIdentifier:3") {
		t.Errorf("Unexpected TargetEndpointGroupIdentifier update payload: %s", calls[0].Payload)
	}
}
