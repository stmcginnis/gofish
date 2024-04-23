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

var aggregateBody = `{
		"@odata.type": "#Aggregate.v1_0_2.Aggregate",
		"Id": "Aggregate1",
		"Name": "Aggregate One",
		"ElementsCount": 2,
		"Elements": [
			{
				"@odata.id": "/redfish/v1/Systems/cluster-node3"
			},
			{
				"@odata.id": "/redfish/v1/Systems/cluster-node4"
			}
		],
		"Actions": {
			"#Aggregate.Reset": {
				"target": "/redfish/v1/AggregationService/Aggregates/Aggregate1/Actions/Aggregate.Reset",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/Aggregates/Aggregate1/ResetActionInfo"
			},
			"#Aggregate.SetDefaultBootOrder": {
				"target": "/redfish/v1/AggregationService/Aggregates/Aggregate1/Actions/Aggregate.SetDefaultBootOrder",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/Aggregates/Aggregate1/SetDefaultBootOrderActionInfo"
			},
			"#Aggregate.AddElements": {
				"target": "/redfish/v1/AggregationService/Aggregates/Aggregate1/Actions/Aggregate.AddElements",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/Aggregates/Aggregate1/AddElementsActionInfo"
			},
			"#Aggregate.RemoveElements": {
				"target": "/redfish/v1/AggregationService/Aggregates/Aggregate1/Actions/Aggregate.RemoveElements",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/Aggregates/Aggregate1/RemoveElementsActionInfo"
			}
		},
		"@odata.id": "/redfish/v1/AggregationService/Aggregates/Aggregate1"
	}`

// TestAggregate tests the parsing of Aggregate objects.
func TestAggregate(t *testing.T) {
	var result Aggregate
	err := json.NewDecoder(strings.NewReader(aggregateBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Aggregate1", result.ID)
	assertEquals(t, "Aggregate One", result.Name)

	if len(result.elements) != 2 {
		t.Errorf("Expected 2 elements, got %#v", result.elements)
	}
}

// TestAggregateReset tests the Aggregate reset call.
func TestAggregateReset(t *testing.T) {
	var result Aggregate
	err := json.NewDecoder(strings.NewReader(aggregateBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.Reset(5, 10, PowerCycleResetType)
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "ResetType:PowerCycle") {
		t.Errorf("Unexpected reset payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "BatchSize:5") {
		t.Errorf("Expected 'BatchSize' to be present and set to 5 %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "DelayBetweenBatchesInSeconds:10") {
		t.Errorf("Expected 'DelayBetweenBatchesInSeconds' to be present and set to 10 %s", calls[0].Payload)
	}
}

// TestAggregateSetDefaultBootOrder tests the Aggregate SetDefaultBootOrder call.
func TestAggregateSetDefaultBootOrder(t *testing.T) {
	var result Aggregate
	err := json.NewDecoder(strings.NewReader(aggregateBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	err = result.SetDefaultBootOrder()
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if calls[0].Payload != "" {
		t.Errorf("Expected empty payload: %#v", calls[0].Payload)
	}
}
