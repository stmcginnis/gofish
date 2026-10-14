//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"fmt"
	"net/url"
	"strings"
)

// HTTPHeaderProperty shall contain the HTTP header name and value to include
// with every event 'POST' to the event destination.
type HTTPHeaderProperty map[string][]string

// subscriptionPayload is the payload to create the event subscription
type subscriptionPayload struct {
	Destination         string                   `json:"Destination,omitempty"`
	EventTypes          []EventType              `json:"EventTypes,omitempty"`
	RegistryPrefixes    []string                 `json:"RegistryPrefixes,omitempty"`
	ResourceTypes       []string                 `json:"ResourceTypes,omitempty"`
	DeliveryRetryPolicy DeliveryRetryPolicy      `json:"DeliveryRetryPolicy,omitempty"`
	HTTPHeaders         map[string]string        `json:"HttpHeaders,omitempty"`
	Oem                 any                      `json:"Oem,omitempty"`
	Protocol            EventDestinationProtocol `json:"Protocol,omitempty"`
	Context             string                   `json:"Context,omitempty"`
}

// validateCreateEventDestinationParams will validate
// CreateEventDestination parameters
//
// Deprecated: (v1.5) EventType-based eventing is DEPRECATED in the Redfish schema
// in favor of using RegistryPrefix and ResourceTypes
func validateCreateEventDestinationParams(
	uri string,
	destination string,
	protocol EventDestinationProtocol,
	context string,
	eventTypes []EventType,
) error {
	// validate event types
	if len(eventTypes) == 0 {
		return fmt.Errorf("at least one event type for subscription should be defined")
	}

	for _, et := range eventTypes {
		if !et.IsValidEventType() {
			return fmt.Errorf("invalid event type")
		}
	}

	return validateCreateEventDestinationMandatoryParams(uri, destination, protocol, context)
}

// validateCreateEventDestinationMandatoryParams will validate
// mandatory parameters for CreateEventDestination HTTP POST request
func validateCreateEventDestinationMandatoryParams(
	uri string,
	destination string,
	protocol EventDestinationProtocol,
	context string,
) error {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}

	// validate destination
	if strings.TrimSpace(destination) == "" {
		return fmt.Errorf("empty destination is not valid")
	}

	u, err := url.ParseRequestURI(destination)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(u.Scheme, "http") {
		return fmt.Errorf("destination should start with http")
	}

	// validate protocol
	if strings.TrimSpace(string(protocol)) == "" {
		return fmt.Errorf("the required property protocol should be defined")
	}

	// validate context
	if strings.TrimSpace(context) == "" {
		return fmt.Errorf("the required property context should be defined")
	}

	return nil
}

// Deprecated: (v1.5) EventType-based eventing is DEPRECATED in the Redfish schema
// in favor of using RegistryPrefix and ResourceTypes
func CreateEventDestination(
	c Client,
	uri string,
	destination string,
	eventTypes []EventType,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem any,
) (string, error) {
	// validate mandatory input parameters
	if err := validateCreateEventDestinationParams(uri, destination, protocol, context, eventTypes); err != nil {
		return "", err
	}

	// create subscription payload
	s := &subscriptionPayload{
		EventTypes: eventTypes,
	}

	return sendCreateEventDestinationRequest(c, s, uri, destination, httpHeaders, protocol, context, oem)
}

// For Redfish v1.5+
// CreateEventDestination will create a EventDestination instance.
// URI should contain the address of the collection for Event Subscriptions.
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
// Returns the new subscription URI if the event subscription is created with success or any error encountered.
func CreateEventDestinationInstance(
	c Client,
	uri string,
	destination string,
	registryPrefixes []string,
	resourceTypes []string,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	deliveryRetryPolicy DeliveryRetryPolicy,
	oem any,
) (string, error) {
	// validate mandatory input parameters
	if err := validateCreateEventDestinationMandatoryParams(uri, destination, protocol, context); err != nil {
		return "", err
	}

	// create subscription payload
	s := &subscriptionPayload{
		RegistryPrefixes:    registryPrefixes,
		ResourceTypes:       resourceTypes,
		DeliveryRetryPolicy: deliveryRetryPolicy,
	}

	return sendCreateEventDestinationRequest(c, s, uri, destination, httpHeaders, protocol, context, oem)
}

func sendCreateEventDestinationRequest(
	c Client,
	s *subscriptionPayload,
	uri string,
	destination string,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem any,
) (string, error) {
	s.Destination = destination
	s.Protocol = protocol
	s.Context = context

	// HTTP headers
	if len(httpHeaders) > 0 {
		s.HTTPHeaders = httpHeaders
	}

	// Oem
	if oem != nil {
		s.Oem = oem
	}

	resp, err := c.Post(uri, s)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return "", err
	}

	// return subscription link from returned location
	subscriptionLink := resp.Header.Get("Location")
	urlParser, err := url.ParseRequestURI(subscriptionLink)
	if err == nil {
		subscriptionLink = urlParser.RequestURI()
	}

	return subscriptionLink, err
}

// DeleteEventDestination will delete a EventDestination.
func DeleteEventDestination(c Client, uri string) error {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}

	resp, err := c.Delete(uri)
	defer DeferredCleanupHTTPResponse(resp)

	return err
}

// linkURIs converts OriginResources, which the schema defines as URIs but
// some services send as links ({"@odata.id": ...}).
func linkURIs(v any) []string {
	items, ok := v.([]any)
	if !ok {
		return nil
	}
	uris := make([]string, 0, len(items))
	for _, item := range items {
		switch x := item.(type) {
		case string:
			uris = append(uris, x)
		case map[string]any:
			if id, ok := x["@odata.id"].(string); ok {
				uris = append(uris, id)
			}
		}
	}
	return uris
}
