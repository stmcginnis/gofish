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

var eventServiceBody = `{
		"@Redfish.Copyright": "Copyright 2014-2019 DMTF. All rights reserved.",
		"@odata.context": "/redfish/v1/$metadata#EventService.EventService",
		"@odata.type": "#EventService.v1_0_0.EventService",
		"@odata.id": "/redfish/v1/EventService",
		"Id": "EventService",
		"Name": "Event Service",
		"DeliveryRetryAttempts": 4,
		"DeliveryRetryIntervalSeconds": 30,
		"Description": "Service for events",
		"EventFormatTypes": [
			"Event",
			"MetricReport"
		],
		"RegistryPrefixes": ["EVENT_"],
		"ResourceTypes": [],
		"SSEFilterPropertiesSupported": {
			"EventFormatType": true,
			"MessageId": true,
			"MetricReportDefinition": false,
			"OriginResource": true,
			"RegistryPrefix": true,
			"ResourceType": true
		},
		"ServerSentEventUri": "http://example.com/events",
		"ServiceEnabled": true,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Subscriptions": {
			"@odata.id": "/redfish/v1/EventService/Subscriptions"
		},
		"Actions": {
			"#EventService.SubmitTestEvent": {
				"target": "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent",
				"EventType@Redfish.AllowableValues": [
					"StatusChange",
					"ResourceUpdated",
					"ResourceAdded",
					"ResourceRemoved",
					"Alert"
				]
			}
		}
	}`

// TestEventService tests the parsing of EventService objects.
func TestEventService(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "EventService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Event Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.DeliveryRetryAttempts != 4 {
		t.Errorf("Expected 4 retry attempts, got: %d", result.DeliveryRetryAttempts)
	}

	if result.DeliveryRetryIntervalSeconds != 30 {
		t.Errorf("Expected 30 second retry interval, got: %d", result.DeliveryRetryIntervalSeconds)
	}

	if result.SSEFilterPropertiesSupported.MetricReportDefinition {
		t.Error("MetricReportDefinition filter should be false")
	}

	if !result.SSEFilterPropertiesSupported.MessageID {
		t.Error("Message ID filter should be true")
	}

	if result.submitTestEventTarget != "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent" {
		t.Errorf("Invalid SubmitTestEvent target: %s", result.submitTestEventTarget)
	}
}

// TestEventServiceUpdate tests the Update call.
func TestEventServiceUpdate(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.DeliveryRetryAttempts = 20
	result.DeliveryRetryIntervalSeconds = 60
	result.ServiceEnabled = true
	err = result.Update()

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "DeliveryRetryAttempts:20") {
		t.Errorf("Unexpected DeliveryRetryAttempts update payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "DeliveryRetryIntervalSeconds:60") {
		t.Errorf("Unexpected DeliveryRetryIntervalSeconds update payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "ServiceEnabled") {
		t.Errorf("Unexpected DeliveryRetryIntervalSeconds update payload: %s", calls[0].Payload)
	}
}
