//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var aggregationServiceBody = `{
		"@odata.type": "#AggregationService.v1_0_2.AggregationService",
		"Id": "AggregationService",
		"Description": "Aggregation Service",
		"Name": "Aggregation Service",
		"ServiceEnabled": true,
		"Status": {
			"Health": "OK",
			"HealthRollup": "OK",
			"State": "Enabled"
		},
		"Aggregates": {
			"@odata.id": "/redfish/v1/AggregationService/Aggregates"
		},
		"AggregationSources": {
			"@odata.id": "/redfish/v1/AggregationService/AggregationSources"
		},
		"ConnectionMethods": {
			"@odata.id": "/redfish/v1/AggregationService/ConnectionMethods"
		},
		"Actions": {
			"#AggregationService.Reset": {
				"target": "/redfish/v1/AggregationService/Actions/AggregationService.Reset",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/ResetActionInfo"
			},
			"#AggregationService.SetDefaultBootOrder": {
				"target": "/redfish/v1/AggregationService/Actions/AggregationService.SetDefaultBootOrder",
				"@Redfish.ActionInfo": "/redfish/v1/AggregationService/SetDefaultBootOrderActionInfo"
			}
		},
		"@odata.id": "/redfish/v1/AggregationService/"
	}`

// TestAggregationService tests the parsing of AggregationService objects.
func TestAggregationService(t *testing.T) {
	var result AggregationService
	err := json.NewDecoder(strings.NewReader(aggregationServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "AggregationService", result.ID)
	assertEquals(t, "Aggregation Service", result.Name)

	if !result.ServiceEnabled {
		t.Error("Expected ServiceEnabled to be true")
	}

	if result.aggregates != "/redfish/v1/AggregationService/Aggregates" {
		t.Errorf("Unexpected aggregates link, got %#v", result.aggregates)
	}

	if result.aggregationSources != "/redfish/v1/AggregationService/AggregationSources" {
		t.Errorf("Unexpected aggregation sources link, got %#v", result.aggregationSources)
	}

	if result.connectionMethods != "/redfish/v1/AggregationService/ConnectionMethods" {
		t.Errorf("Unexpected connection methods link, got %#v", result.connectionMethods)
	}
}

// TestAggregationServiceReset tests the AggregationService reset call.
func TestAggregationServiceReset(t *testing.T) {
	var result AggregationService
	err := json.NewDecoder(strings.NewReader(aggregationServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.Reset(&AggregationServiceResetParameters{
		BatchSize:                    5,
		DelayBetweenBatchesInSeconds: 10,
		ResetType:                    PowerCycleResetType,
		TargetURIs:                   []string{"/link/1", "/link/2"},
	})
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

	if !strings.Contains(calls[0].Payload, "/link/2") {
		t.Errorf("Unexpected 'target links': %s", calls[0].Payload)
	}
}

// TestAggregationServiceSetDefaultBootOrder tests the AggregationService SetDefaultBootOrder call.
func TestAggregationServiceSetDefaultBootOrder(t *testing.T) {
	var result AggregationService
	err := json.NewDecoder(strings.NewReader(aggregationServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.SetDefaultBootOrder([]string{"/link/system/1"})
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "/link/system/1") {
		t.Errorf("Target system not found in payload: %s", calls[0].Payload)
	}
}
