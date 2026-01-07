//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var compositionServiceBody = `{
		"@odata.context": "/redfish/v1/$metadata#CompositionService.CompositionService",
		"@odata.type": "#CompositionService.v1_0_0.CompositionService",
		"@odata.id": "/redfish/v1/CompositionService",
		"Id": "CompositionService-1",
		"Name": "Composition Service",
		"Description": "Composition Service",
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"AllowOverprovisioning": true,
		"AllowZoneAffinity": false,
		"ServiceEnabled": true,
		"ResourceBlocks": {
			"@odata.id": "/redfish/v1/CompositionService/ResourceBlocks"
		},
		"ResourceZones": {
			"@odata.id": "/redfish/v1/CompositionService/ResourceZones"
		}
	}`

// TestCompositionService tests the parsing of CompositionService objects.
func TestCompositionService(t *testing.T) {
	var result CompositionService
	err := json.NewDecoder(strings.NewReader(compositionServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CompositionService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Composition Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.AllowOverprovisioning {
		t.Error("Expected AllowOverprovisioning to be true")
	}

	if result.AllowZoneAffinity {
		t.Error("Expected AllowZoneAffinity to be false")
	}

	if !result.ServiceEnabled {
		t.Error("Expected ServiceEnabled to be true")
	}

	if result.Status.Health != OKHealth {
		t.Errorf("Received invalid health status: %s", result.Status.Health)
	}

	if result.resourceBlocks != "/redfish/v1/CompositionService/ResourceBlocks" {
		t.Errorf("Received invalid resource blocks reference: %s", result.resourceBlocks)
	}

	if result.resourceZones != "/redfish/v1/CompositionService/ResourceZones" {
		t.Errorf("Received invalid resource zones reference: %s", result.resourceZones)
	}
}

// TestCompositionServiceUpdate tests the Update call.
func TestCompositionServiceUpdate(t *testing.T) {
	var result CompositionService
	err := json.NewDecoder(strings.NewReader(compositionServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	result.ServiceEnabled = false
	result.AllowOverprovisioning = false
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %v", calls)
	}

	if !strings.Contains(calls[0].Payload, "AllowOverprovisioning:false") {
		t.Errorf("Unexpected AllowOverprovisioning update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "ServiceEnabled:false") {
		t.Errorf("Unexpected ServiceEnabled update payload: %s", calls[0].Payload)
	}
}
