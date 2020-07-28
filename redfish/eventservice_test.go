//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		"EventTypesForSubscription":[
			"StatusChange",
			"ResourceUpdated",
			"ResourceAdded",
			"ResourceRemoved",
			"Alert"
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

	for _, et := range result.EventTypesForSubscription {
		if !et.IsValidEventType() {
			t.Errorf("invalid event type: %s", et)
		}
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

// OemVendor is the Oem used during create event subscription test
type OemVendor struct {
	Vendor Vendor `json:"Vendor"`
}

// Vendor used by OemVendor during create event subscription test
type Vendor struct {
	FirstVendorSpecificConfiguration  int `json:"FirstVendorSpecificConfiguration"`
	SecondVendorSpecificConfiguration int `json:"SecondVendorSpecificConfiguration"`
}

// TestEventServiceCreateEventSubscription tests the CreateEventSubscription call.
func TestEventServiceCreateEventSubscription(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// define the expected subscription URI that should be
	// returned during create event subscription
	expectedSubscriptionURI := "/redfish/v1/EventService/Subscriptions/SubscriptionId/"

	// define the custom return for the operations
	// during the use of the test client
	customReturnForOperationsForTestClient := map[string]interface{}{
		// custom return for the POST operation
		http.MethodPost: &http.Response{
			Status:        "201 Created",
			StatusCode:    201,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString("")),
			ContentLength: int64(len("")),
			Header: http.Header{
				"Location": []string{
					fmt.Sprintf("https://redfish-server%s", expectedSubscriptionURI),
				},
			},
		},
	}

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForOperations: customReturnForOperationsForTestClient,
	}
	result.SetClient(testClient)

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		[]EventType{SupportedEventTypes["Alert"]},
		map[string]string{
			"Header": "HeaderValue",
		},
		OemVendor{
			Vendor: Vendor{
				FirstVendorSpecificConfiguration:  1,
				SecondVendorSpecificConfiguration: 2,
			},
		},
	)

	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	if subscriptionURI != expectedSubscriptionURI {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			subscriptionURI,
			expectedSubscriptionURI)
	}

}
