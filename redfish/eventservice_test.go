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

const (
	subscriptionMask   = "/redfish/v1/EventService/Subscriptions/%s/"
	subsctiptionID     = "SubscriptionId"
	eventDestinationID = "EventDestination-1"
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
		"ResourceTypes": ["Chassis"],
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

func assertContains(t testing.TB, expected, actual string) {
	t.Helper()
	if !strings.Contains(actual, expected) {
		t.Errorf("\nExpected payload item: %s \nActual CreateEventSubscription payload: %s", expected, actual)
	}
}

func assertNotContain(t testing.TB, expected, actual string) {
	t.Helper()
	if strings.Contains(actual, expected) {
		t.Errorf("\nExpected payload item: %s \nActual CreateEventSubscription payload: %s", expected, actual)
	}
}

func assertError(t testing.TB, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("\nExpected error: %s \nActual error: %s",
			actual,
			expected)
	}
}

// TestEventService tests the parsing of EventService objects.
func TestEventService(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals := func(t testing.TB, expected string, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("\nExpected value: %s \nActual value: %s", expected, actual)
		}
	}

	assertEquals(t, "EventService", result.ID)
	assertEquals(t, "Event Service", result.Name)
	assertEquals(t, "4", fmt.Sprint(result.DeliveryRetryAttempts))
	assertEquals(t, "30", fmt.Sprint(result.DeliveryRetryIntervalSeconds))
	assertEquals(t, "false", fmt.Sprint(result.SSEFilterPropertiesSupported.MetricReportDefinition))
	assertEquals(t, "true", fmt.Sprint(result.SSEFilterPropertiesSupported.MessageID))
	assertEquals(t, "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent", result.SubmitTestEventTarget)

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
	expectedSubscriptionURI := fmt.Sprintf(subscriptionMask, subsctiptionID)

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

	validDestinationURI := "https://myeventreceiver/eventreceiver"
	validEventTypes := []EventType{AlertEventType}
	validCustomHeader := map[string]string{
		"Header": "HeaderValue"}
	validDestinationProtocol := RedfishEventDestinationProtocol
	validContext := "Public"
	validOem := OemVendor{
		Vendor: Vendor{
			FirstVendorSpecificConfiguration:  1,
			SecondVendorSpecificConfiguration: 2,
		},
	}

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscription(
		validDestinationURI,
		validEventTypes,
		validCustomHeader,
		validDestinationProtocol,
		validContext,
		validOem,
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	// validate CreateEventSubscription call
	assertError(t, expectedSubscriptionURI, subscriptionURI)

	// validate the payload
	calls := testClient.CapturedCalls()

	actual := calls[0].Payload

	propertyName := "Destination"
	expected := fmt.Sprintf("%s:%s", propertyName, validDestinationURI)
	assertContains(t, expected, actual)

	propertyName = "EventTypes"
	expected = fmt.Sprintf("%s:%v", propertyName, validEventTypes)
	assertContains(t, expected, actual)

	propertyName = "HttpHeaders"
	expected = fmt.Sprintf("%s:%v", propertyName, validCustomHeader)
	assertContains(t, expected, actual)

	propertyName = "Protocol"
	expected = fmt.Sprintf("%s:%v", propertyName, validDestinationProtocol)
	assertContains(t, expected, actual)

	propertyName = "Context"
	expected = fmt.Sprintf("%s:%v", propertyName, validContext)
	assertContains(t, expected, actual)

	propertyName = "Oem"
	expected = fmt.Sprintf("%s:map[Vendor:map[FirstVendorSpecificConfiguration:%d SecondVendorSpecificConfiguration:%d]",
		propertyName, validOem.Vendor.FirstVendorSpecificConfiguration, validOem.Vendor.SecondVendorSpecificConfiguration)
	assertContains(t, expected, actual)
}

func TestEventServiceCreateEventSubscriptionInstance(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	// define the expected subscription URI that should be
	// returned during create event subscription
	expectedSubscriptionURI := fmt.Sprintf(subscriptionMask, subsctiptionID)

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

	validDestinationURI := "https://myeventreceiver/eventreceiver"
	validDestinationProtocol := RedfishEventDestinationProtocol
	validContext := "Public"
	validRegistryPrefixes := []string{"EVENT_"}
	validResourceTypes := []string{"Chassis"}
	validDeliveryRetryPolicy := SuspendRetriesDeliveryRetryPolicy

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscriptionInstance(
		validDestinationURI,
		validRegistryPrefixes,
		validResourceTypes,
		nil,
		validDestinationProtocol,
		validContext,
		validDeliveryRetryPolicy,
		nil,
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	// validate CreateEventSubscription call
	assertError(t, expectedSubscriptionURI, subscriptionURI)

	// validate the payload
	calls := testClient.CapturedCalls()

	actual := calls[0].Payload

	propertyName := "Destination"
	expected := fmt.Sprintf("%s:%s", propertyName, validDestinationURI)
	assertContains(t, expected, actual)

	propertyName = "RegistryPrefixes"
	expected = fmt.Sprintf("%s:%v", propertyName, validRegistryPrefixes)
	assertContains(t, expected, actual)

	propertyName = "ResourceTypes"
	expected = fmt.Sprintf("%s:%v", propertyName, validResourceTypes)
	assertContains(t, expected, actual)

	propertyName = "Protocol"
	expected = fmt.Sprintf("%s:%v", propertyName, validDestinationProtocol)
	assertContains(t, expected, actual)

	propertyName = "Context"
	expected = fmt.Sprintf("%s:%v", propertyName, validContext)
	assertContains(t, expected, actual)

	propertyName = "DeliveryRetryPolicy"
	expected = fmt.Sprintf("%s:%s", propertyName, validDeliveryRetryPolicy)
	assertContains(t, expected, actual)
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
	subscriptionToDelete := fmt.Sprintf(subscriptionMask, subsctiptionID)
	err = result.DeleteEventSubscription(subscriptionToDelete)

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
	eventSubscription := fmt.Sprintf(subscriptionMask, eventDestinationID)
	eventDestination, err := result.GetEventSubscription(eventSubscription)

	// validate the return values
	if eventDestination.ID != eventDestinationID {
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
	if eventDestinations[0].ID != eventDestinationID {
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
	expectedSubscriptionURI := fmt.Sprintf(subscriptionMask, subsctiptionID)

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

	validDestinationURI := "https://myeventreceiver/eventreceiver"
	validEventTypes := []EventType{AlertEventType}
	validDestinationProtocol := RedfishEventDestinationProtocol
	validContext := "Public"

	// create event subscription
	subscriptionURI, err := result.CreateEventSubscription(
		validDestinationURI,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the return values
	if err != nil {
		t.Errorf("Error making CreateEventSubscription call: %s", err)
	}

	// validate CreateEventSubscription call
	assertError(t, expectedSubscriptionURI, subscriptionURI)

	// validate the payload
	calls := testClient.CapturedCalls()

	actual := calls[0].Payload

	propertyName := "Destination"
	expected := fmt.Sprintf("%s:%s", propertyName, validDestinationURI)
	assertContains(t, expected, actual)

	propertyName = "EventTypes"
	expected = fmt.Sprintf("%s:%v", propertyName, validEventTypes)
	assertContains(t, expected, actual)

	propertyName = "Protocol"
	expected = fmt.Sprintf("%s:%v", propertyName, validDestinationProtocol)
	assertContains(t, expected, actual)

	propertyName = "Context"
	expected = fmt.Sprintf("%s:%v", propertyName, validContext)
	assertContains(t, expected, actual)

	expected = "HttpHeaders"
	assertNotContain(t, expected, actual)

	expected = "Oem"
	assertNotContain(t, expected, actual)
}

// TestEventServiceCreateEventSubscriptionInputParametersValidation
// tests the validation of input parameters for CreateEventSubscription.
func TestEventServiceCreateEventSubscriptionInputParametersValidation(t *testing.T) { //nolint
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	validDestinationURI := "https://myeventreceiver/eventreceiver"
	validEventTypes := []EventType{AlertEventType}
	validDestinationProtocol := RedfishEventDestinationProtocol
	validContext := "Public"

	// create event subscription invalid destination
	invalidDestination := "myeventreciever/eventreceiver"
	_, err = result.CreateEventSubscription(
		invalidDestination,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError := fmt.Sprintf("parse %q: invalid URI for request", invalidDestination)
	assertError(t, expectedError, err.Error())

	// create event subscription invalid destination
	invalidDestination = "ftp://myeventreciever/eventreceiver"
	_, err = result.CreateEventSubscription(
		invalidDestination,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "destination should start with http"
	assertError(t, expectedError, err.Error())

	// create event subscription invalid destination
	invalidDestination = ""
	_, err = result.CreateEventSubscription(
		invalidDestination,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "empty destination is not valid"
	assertError(t, expectedError, err.Error())

	// create event subscription invalid destination
	invalidDestination = "   "
	_, err = result.CreateEventSubscription(
		invalidDestination,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "empty destination is not valid"
	assertError(t, expectedError, err.Error())

	// create event subscription empty event type
	_, err = result.CreateEventSubscription(
		validDestinationURI,
		[]EventType{},
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "at least one event type for subscription should be defined"
	assertError(t, expectedError, err.Error())

	// create event subscription nil event type
	_, err = result.CreateEventSubscription(
		validDestinationURI,
		nil,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "at least one event type for subscription should be defined"
	assertError(t, expectedError, err.Error())

	// create event subscription empty
	// subscription link in the event service
	result.Subscriptions = ""
	_, err = result.CreateEventSubscription(
		validDestinationURI,
		validEventTypes,
		nil,
		validDestinationProtocol,
		validContext,
		nil,
	)

	// validate the returned error
	expectedError = "empty subscription link in the event service"
	assertError(t, expectedError, err.Error())
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
	assertError(t, expectedError, err.Error())

	// delete event subscription
	err = result.DeleteEventSubscription(" ")

	// validate the returned error
	expectedError = "uri should not be empty"
	assertError(t, expectedError, err.Error())
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
	assertError(t, expectedError, err.Error())

	// get event subscription
	_, err = result.GetEventSubscription(" ")

	// validate the returned error
	expectedError = "uri should not be empty"
	assertError(t, expectedError, err.Error())
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
	assertError(t, expectedError, err.Error())
}
