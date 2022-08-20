//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	if result.SubmitTestEventTarget != "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent" {
		t.Errorf("Invalid SubmitTestEvent target: %s", result.SubmitTestEventTarget)
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

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodPost: {
				// defining the custom return for the first POST operation
				&http.Response{
					Status:        "201 Created",
					StatusCode:    201,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString("")),
					ContentLength: int64(len("")),
					Header: http.Header{
						"Location": []string{
							fmt.Sprintf("https://redfish-server%s",
								expectedSubscriptionURI),
						},
					},
				},
			},
		},
	}
	result.SetClient(testClient)

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		[]EventType{SupportedEventTypes["Alert"]},
		map[string]string{
			"Header": "HeaderValue",
		},
		RedfishEventDestinationProtocol,
		"Public",
		OemVendor{
			Vendor: Vendor{
				FirstVendorSpecificConfiguration:  1,
				SecondVendorSpecificConfiguration: 2,
			},
		},
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	if subscriptionURI != expectedSubscriptionURI {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			subscriptionURI,
			expectedSubscriptionURI)
	}

	// validate the payload
	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Destination:https://myeventreciever/eventreceiver") {
		t.Errorf("Unexpected Destination CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "EventTypes:[Alert]") {
		t.Errorf("Unexpected EventTypes CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "HttpHeaders:map[Header:HeaderValue]") {
		t.Errorf("Unexpected HttpHeaders CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "Oem:map[Vendor:map[FirstVendorSpecificConfiguration:1 SecondVendorSpecificConfiguration:2]") {
		t.Errorf("Unexpected Oem CreateEventSubscription payload: %s", calls[0].Payload)
	}
}

// TestEventServiceDeleteEventSubscription tests the DeleteEventSubscription call.
func TestEventServiceDeleteEventSubscription(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodDelete: {
				// defining the custom return for the
				// first DELETE operation
				err,
			},
		},
	}
	result.SetClient(testClient)

	// create event subscription
	err = result.DeleteEventSubscription(
		"/redfish/v1/EventService/Subscriptions/SubscriptionId/")

	// validate the return values
	if err != nil {
		t.Errorf("Error making DeleteEventSubscription call: %s", err)
	}
}

// TestEventServiceGetEventSubscription tests the GetEventSubscription call.
func TestEventServiceGetEventSubscription(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				// defining the custom return for the first GET operation
				&http.Response{
					Status:        "200 OK",
					StatusCode:    200,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString(eventDestinationBody)),
					ContentLength: int64(len(eventDestinationBody)),
					Header:        make(http.Header),
				},
			},
		},
	}
	result.SetClient(testClient)

	// create event subscription
	eventDestination, err := result.GetEventSubscription(
		"/redfish/v1/EventService/Subscriptions/EventDestination-1/")

	// validate the return values
	if eventDestination.ID != "EventDestination-1" {
		t.Errorf("Error making GetEventSubscription call: %s", err)
	}
}

// TestEventServiceGetEventSubscriptions tests the GetEventSubscriptions call.
func TestEventServiceGetEventSubscriptions(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				// defining the custom return for the first GET operation
				&http.Response{
					Status:        "200 OK",
					StatusCode:    200,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString(eventDestinationsBody)),
					ContentLength: int64(len(eventDestinationsBody)),
					Header:        make(http.Header),
				},
				// defining the custom return for the second GET operation
				&http.Response{
					Status:        "200 OK",
					StatusCode:    200,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString(eventDestinationBody)),
					ContentLength: int64(len(eventDestinationBody)),
					Header:        make(http.Header),
				},
			},
		},
	}
	result.SetClient(testClient)

	// create event subscription
	eventDestinations, err := result.GetEventSubscriptions()

	// validate the return values
	if eventDestinations[0].ID != "EventDestination-1" {
		t.Errorf("Error making GetEventSubscriptions call: %s", err)
	}
}

// TestEventServiceCreateEventSubscriptionWithoutOptionalParameters
// tests the CreateEventSubscription call without optional parameters.
func TestEventServiceCreateEventSubscriptionWithoutOptionalParameters(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// define the expected subscription URI that should be
	// returned during create event subscription
	expectedSubscriptionURI := "/redfish/v1/EventService/Subscriptions/SubscriptionId/"

	// create the custom test client
	testClient := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodPost: {
				// defining the custom return for the first POST operation
				&http.Response{
					Status:        "201 Created",
					StatusCode:    201,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(bytes.NewBufferString("")),
					ContentLength: int64(len("")),
					Header: http.Header{
						"Location": []string{
							fmt.Sprintf("https://redfish-server%s",
								expectedSubscriptionURI),
						},
					},
				},
			},
		},
	}
	result.SetClient(testClient)

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		[]EventType{SupportedEventTypes["Alert"]},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	if subscriptionURI != expectedSubscriptionURI {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			subscriptionURI,
			expectedSubscriptionURI)
	}

	// validate the payload
	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "Destination:https://myeventreciever/eventreceiver") {
		t.Errorf("Unexpected Destination CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if !strings.Contains(calls[0].Payload, "EventTypes:[Alert]") {
		t.Errorf("Unexpected EventTypes CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "Oem") {
		t.Errorf("Unexpected Oem CreateEventSubscription payload: %s", calls[0].Payload)
	}

	if strings.Contains(calls[0].Payload, "HttpHeaders") {
		t.Errorf("Unexpected HttpHeaders CreateEventSubscription payload: %s", calls[0].Payload)
	}
}

// TestEventServiceCreateEventSubscriptionInputParametersValidation
// tests the validation of input parameters for CreateEventSubscription.
func TestEventServiceCreateEventSubscriptionInputParametersValidation(t *testing.T) { //nolint
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// create event subscription invalid destination
	invalidDestination := "myeventreciever/eventreceiver"
	_, err = result.CreateEventSubscription(
		invalidDestination,
		[]EventType{SupportedEventTypes["Alert"]},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError := "destination should start with http"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// create event subscription invalid destination
	invalidDestination = ""
	_, err = result.CreateEventSubscription(
		invalidDestination,
		[]EventType{SupportedEventTypes["Alert"]},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError = "empty destination is not valid"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// create event subscription invalid destination
	invalidDestination = "   "
	_, err = result.CreateEventSubscription(
		invalidDestination,
		[]EventType{SupportedEventTypes["Alert"]},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError = "empty destination is not valid"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// create event subscription empty event type
	_, err = result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		[]EventType{},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError = "at least one event type for subscription should be defined"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// create event subscription nil event type
	_, err = result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		nil,
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError = "at least one event type for subscription should be defined"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// create event subscription empty
	// subscription link in the event service
	result.Subscriptions = ""
	_, err = result.CreateEventSubscription(
		"https://myeventreciever/eventreceiver",
		[]EventType{SupportedEventTypes["Alert"]},
		nil,
		RedfishEventDestinationProtocol,
		"Public",
		nil,
	)

	// validate the returned error
	expectedError = "empty subscription link in the event service"
	if err.Error() != expectedError {
		t.Errorf("Error CreateEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}
}

// TestEventServiceDeleteEventSubscriptionInputParametersValidation
// tests the validation of input parameters for DeleteEventSubscription.
func TestEventServiceDeleteEventSubscriptionInputParametersValidation(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	// delete event subscription
	err = result.DeleteEventSubscription("")

	// validate the returned error
	expectedError := "uri should not be empty"
	if err.Error() != expectedError {
		t.Errorf("Error DeleteEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// delete event subscription
	err = result.DeleteEventSubscription(" ")

	// validate the returned error
	expectedError = "uri should not be empty"
	if err.Error() != expectedError {
		t.Errorf("Error DeleteEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}
}

// TestEventServiceGetEventSubscriptionInputParametersValidation
// tests the validation of input parameters for GetEventSubscription.
func TestEventServiceGetEventSubscriptionInputParametersValidation(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
	// get event subscription
	_, err = result.GetEventSubscription("")

	// validate the returned error
	expectedError := "uri should not be empty"
	if err.Error() != expectedError {
		t.Errorf("Error GetEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}

	// get event subscription
	_, err = result.GetEventSubscription(" ")

	// validate the returned error
	expectedError = "uri should not be empty"
	if err.Error() != expectedError {
		t.Errorf("Error GetEventSubscription returned: %s expected: %s",
			err,
			expectedError)
	}
}

// TestEventServiceGetEventSubscriptionsEmptySubscriptionsLink
// tests the GetEventSubscriptions when the subscriptions link
// is empty.
func TestEventServiceGetEventSubscriptionsEmptySubscriptionsLink(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// get event subscriptions with empty subscription link
	result.Subscriptions = ""
	_, err = result.GetEventSubscriptions()

	// validate the returned error
	expectedError := "empty subscription link in the event service"
	if err.Error() != expectedError {
		t.Errorf("Error GetEventSubscriptions returned: %s expected: %s",
			err,
			expectedError)
	}
}
