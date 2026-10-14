//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// UnmarshalJSON unmarshals a EventService object from the raw JSON.
func (e *EventService) UnmarshalJSON(b []byte) error {
	type temp EventService
	type eActions struct {
		SubmitTestEvent       ActionTarget `json:"#EventService.SubmitTestEvent"`
		TestEventSubscription ActionTarget `json:"#EventService.TestEventSubscription"`
	}
	var tmp struct {
		temp
		Actions       eActions
		Subscriptions Link `json:"Subscriptions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = EventService(tmp.temp)

	// Extract the links to other entities for later
	e.SubmitTestEventTarget = tmp.Actions.SubmitTestEvent.Target
	e.testEventSubscriptionTarget = tmp.Actions.TestEventSubscription.Target
	e.SubscriptionsLink = strings.TrimSpace(tmp.Subscriptions.String())

	// This is a read/write object, so we need to save the raw object data for later
	e.RawData = b

	return nil
}

// Deprecated: (v1.5) EventType-based eventing is DEPRECATED in the Redfish schema
// in favor of using RegistryPrefix and ResourceTypes
func (e *EventService) CreateEventSubscription(
	destination string,
	eventTypes []EventType,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem any,
) (string, error) {
	if e.SubscriptionsLink == "" {
		return "", errors.New("empty subscription link in the event service")
	}

	return CreateEventDestination(
		e.GetClient(),
		e.SubscriptionsLink,
		destination,
		eventTypes,
		httpHeaders,
		protocol,
		context,
		oem,
	)
}

// For Redfish v1.5+
// CreateEventSubscription creates the subscription using the event service.
// Destination should contain the URL of the destination for events to be sent.
// RegistryPrefixes is the list of the prefixes for the Message Registries
// that contain the MessageIds that are sent to this event destination.
// If RegistryPrefixes is empty on subscription, the client is subscribing to all Message Registries.
// ResourceTypes is the list of Resource Type values (Schema names) that correspond to the OriginOfCondition,
// the version and full namespace should not be specified.
// If ResourceTypes is empty on subscription, the client is subscribing to receive events regardless of ResourceType.
// HttpHeaders is optional and gives the opportunity to specify any arbitrary
// HTTP headers required for the event POST operation.
// Protocol should be the communication protocol of the event endpoint, usually RedfishEventDestinationProtocol.
// Context is a required client-supplied string that is sent with the event notifications.
// DeliveryRetryPolicy is optional, it should contain the subscription delivery retry policy for events,
// where the subscription type is RedfishEvent.
// Oem is optional and gives the opportunity to specify any OEM specific properties,
// it should contain the vendor specific struct that goes inside the Oem session.
// It returns the new subscription URI if the event subscription is created
// with success or any error encountered.
func (e *EventService) CreateEventSubscriptionInstance(
	destination string,
	registryPrefixes []string,
	resourceTypes []string,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	deliveryRetryPolicy DeliveryRetryPolicy,
	oem any,
) (string, error) {
	if e.SubscriptionsLink == "" {
		return "", errors.New("empty subscription link in the event service")
	}

	return CreateEventDestinationInstance(
		e.GetClient(),
		e.SubscriptionsLink,
		destination,
		registryPrefixes,
		resourceTypes,
		httpHeaders,
		protocol,
		context,
		deliveryRetryPolicy,
		oem,
	)
}

// DeleteEventSubscription deletes a specific subscription using the event service.
func (e *EventService) DeleteEventSubscription(uri string) error {
	return DeleteEventDestination(e.GetClient(), uri)
}

// This action shall add a test event to the event service with the event data
// specified in the action parameters. Then, this message should be sent to any
// appropriate event destinations.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EventService) SubmitTestEvent(params *EventServiceSubmitTestEventParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(e.client,
		e.SubmitTestEventTarget, params, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// SubmitTestingEvent shall add a test event to the event service with the event data specified in the action parameters. Then, this message should be sent to any appropriate event destinations.
// message - Event message content.
func (e *EventService) SubmitTestingEvent(message string) error {
	_, err := e.SubmitTestEvent(&EventServiceSubmitTestEventParameters{
		EventGroupID:      123,
		EventID:           "TEST123",
		EventTimestamp:    time.Now().String(),
		EventType:         AlertEventType,
		Message:           message,
		OriginOfCondition: e.ODataID,
		Severity:          "Informational",
	})

	return err
}

// This action shall send an event containing the 'TestMessage' message from
// the Resource Event Message Registry to all appropriate event destinations.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EventService) TestEventSubscription() (*TaskMonitorInfo, error) {
	if e.testEventSubscriptionTarget == "" {
		return nil, errors.New("TestEventSubsciption not supported by this service") //nolint:error-strings
	}

	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(e.client,
		e.testEventSubscriptionTarget, payload, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Subscriptions gets the Subscriptions collection.
func (e *EventService) Subscriptions() ([]*EventDestination, error) {
	if e.SubscriptionsLink == "" {
		return nil, errors.New("empty subscription link in the event service")
	}
	return GetCollectionObjects[EventDestination](e.client, e.SubscriptionsLink)
}

// GetEventSubscription gets a specific subscription using the event service.
func (e *EventService) GetEventSubscription(uri string) (*EventDestination, error) {
	if uri == "" {
		return nil, errors.New("uri should not be empty")
	}
	return GetEventDestination(e.GetClient(), uri)
}
