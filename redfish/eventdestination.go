//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

// DeliveryRetryPolicy is the retry policy for delivery failure.
type DeliveryRetryPolicy string

const (

	// TerminateAfterRetriesDeliveryRetryPolicy The subscription is
	// terminated after the maximum number of retries is reached.
	TerminateAfterRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "TerminateAfterRetries"
	// SuspendRetriesDeliveryRetryPolicy The subscription is suspended after
	// the maximum number of retries is reached.
	SuspendRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "SuspendRetries"
	// RetryForeverDeliveryRetryPolicy shall continue even after the after
	// the maximum number of retries is reached.
	RetryForeverDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForever"
)

// EventDestinationProtocol is the communication protocol of the event destination.
type EventDestinationProtocol string

const (

	// RedfishEventDestinationProtocol The destination follows the Redfish
	// specification for event notifications.
	RedfishEventDestinationProtocol EventDestinationProtocol = "Redfish"
	// SNMPv1EventDestinationProtocol shall indicate the destination follows
	// the RFC1157-defined SNMPv1 protocol.
	SNMPv1EventDestinationProtocol EventDestinationProtocol = "SNMPv1"
	// SNMPv2cEventDestinationProtocol shall indicate the destination follows
	// the SNMPv2c protocol as defined by RFC1441 and RFC1452.
	SNMPv2cEventDestinationProtocol EventDestinationProtocol = "SNMPv2c"
	// SNMPv3EventDestinationProtocol shall indicate the destination follows
	// the SNMPv3 protocol as defined by RFC3411 and RFC3418.
	SNMPv3EventDestinationProtocol EventDestinationProtocol = "SNMPv3"
	// SMTPEventDestinationProtocol shall indicate the destination follows
	// the RFC5321-defined SMTP specification.
	SMTPEventDestinationProtocol EventDestinationProtocol = "SMTP"
)

// SubscriptionType is the type of subscription used.
type SubscriptionType string

const (

	// RedfishEventSubscriptionType The subscription follows the Redfish
	// specification for event notifications, which is done by a service
	// sending an HTTP POST to the subscriber's destination URI.
	RedfishEventSubscriptionType SubscriptionType = "RedfishEvent"
	// SSESubscriptionType The subscription follows the HTML5 Server-Sent
	// Event definition for event notifications.
	SSESubscriptionType SubscriptionType = "SSE"
	// SNMPTrapSubscriptionType shall indicate the subscription follows the
	// various versions of SNMP Traps for event notifications.
	// EventDestinationProtocol shall specify the appropriate version of
	// SNMP.
	SNMPTrapSubscriptionType SubscriptionType = "SNMPTrap"
	// SNMPInformSubscriptionType shall indicate the subscription follows
	// versions 2 and 3 of SNMP Inform for event notifications.
	// EventDestinationProtocol shall specify the appropriate version of
	// SNMP.
	SNMPInformSubscriptionType SubscriptionType = "SNMPInform"
)

// EventDestination is used to represent the target of an event
// subscription, including the types of events subscribed and context to
// provide to the target in the Event payload.
type EventDestination struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Context shall contain a client supplied context that will remain with the
	// connection through the connections lifetime.
	Context string
	// DeliveryRetryPolicy shall indicate the subscription
	// delivery retry policy for events where the subscription type is
	// RedfishEvent. If this property is not present, the policy shall be
	// assumed to be TerminateAfterRetries.
	DeliveryRetryPolicy DeliveryRetryPolicy
	// Description provides a description of this resource.
	Description string
	// Destination shall contain a URI to the destination where the events will
	// be sent.
	Destination string
	// EventFormatType shall indicate the the content types of the message that
	// this service will send to the EventDestination. If this property is not
	// present, the EventFormatType shall be assumed to be Event.
	EventFormatType EventFormatType
	// EventTypes contains the types of events
	// that will be sent to the destination.
	EventTypes []EventType
	// HTTPHeaders shall contain an object consisting of the names and values of
	// of HTTP header to be included with every event POST to the Event
	// Destination. This property shall be null or an empty array on a GET. An
	// empty array is the preferred return value on GET.
	HTTPHeaders []HTTPHeaderProperty `json:"HttpHeaders"`
	// IncludeOriginOfCondition shall indicate whether the
	// event payload sent to the subscription destination will expand the
	// OriginOfCondition property to include the resource or object
	// referenced by the OriginOfCondition property.
	IncludeOriginOfCondition bool
	// MessageIDs shall specify an array of MessageIds that are the only
	// allowable values for the MessageId property within an EventRecord sent to
	// the subscriber. Events with MessageIds not contained in this array shall
	// not be sent to the subscriber. If this property is absent or the array is
	// empty, the service shall send Events with any MessageId to the subscriber.
	MessageIDs []string `json:"MessageIds"`
	// metricReportDefinitions shall specify an array of
	// metric report definitions that are the only allowable generators of
	// metric reports for this subscription. Metric reports originating from
	// metric report definitions not contained in this array shall not be
	// sent to the subscriber. If this property is absent or the array is
	// empty, the service shall send metric reports originating from any
	// metric report definition to the subscriber.
	// metricReportDefinitions []string
	// MetricReportDefinitions@odata.count is
	MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
	// originResources shall specify an array of Resources, Resource Collections,
	// or Referenceable Members that are the only allowable values for the
	// OriginOfCondition property within an EventRecord sent to the subscriber.
	// Events originating from Resources, Resource Collections, or Referenceable
	// Members not contained in this array shall not be sent to the subscriber.
	// If this property is absent or the array is empty, the service shall send
	// Events originating from any Resource, Resource Collection, or
	// Referenceable Member to the subscriber.
	// originResources []string
	// OriginResourcesCount is the number of OriginResources.
	OriginResourcesCount int `json:"OriginResources@odata.count"`
	// Protocol is used to indicate that the event type shall adhere to that
	// defined in the Redfish specification.
	Protocol EventDestinationProtocol
	// RegistryPrefixes is The value of this property is the array of the
	// Prefixes of the Message Registries that contain the MessageIds in the
	// Events that shall be sent to the EventDestination. If this property
	// is absent or the array is empty, the service shall send Events with
	// MessageIds from any Message Registry.
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of Resource Type values. When an
	// event is generated, if the OriginOfCondition's Resource Type matches a
	// value in this array, the event shall be sent to the event destination
	// (unless it would be filtered by other property conditions such as
	// RegistryPrefix). If this property is absent or the array is empty, the
	// service shall send Events from any Resource Type to the subscriber. The
	// value of this property shall be only the general namespace for the type
	// and not the versioned value. For example, it shall not be Task.v1_2_0.Task
	// and instead shall just be Task. To specify that a client is subscribing
	// for Metric Reports, the EventTypes property should include 'MetricReport'.
	ResourceTypes []string
	// SNMP shall contain the settings for an SNMP event destination.
	SNMP SNMPSettings
	// Status shall contain the status of the subscription.
	Status common.Status
	// SubordinateResources is When set to true and OriginResources is
	// specified, indicates the subscription shall be for events from the
	// OriginsResources specified and all subordinate resources. When set to
	// false and OriginResources is specified, indicates subscription shall
	// be for events only from the OriginResources. If OriginResources is
	// not specified, it has no relevance.
	SubordinateResources bool
	// SubscriptionType shall indicate the type of subscription for events. If
	// this property is not present, the SubscriptionType shall be assumed to be
	// RedfishEvent.
	SubscriptionType SubscriptionType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EventDestination object from the raw JSON.
func (eventdestination *EventDestination) UnmarshalJSON(b []byte) error {
	type temp EventDestination
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*eventdestination = EventDestination(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	eventdestination.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (eventdestination *EventDestination) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EventDestination)
	err := original.UnmarshalJSON(eventdestination.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Context",
		"DeliveryRetryPolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(eventdestination).Elem()

	return eventdestination.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEventDestination will get a EventDestination instance from the service.
func GetEventDestination(c common.Client, uri string) (*EventDestination, error) {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return nil, fmt.Errorf("uri should not be empty")
	}

	var eventdestination EventDestination
	return &eventdestination, eventdestination.Get(c, uri, &eventdestination)
}

// subscriptionPayload is the payload to create the event subscription
type subscriptionPayload struct {
	Destination         string                   `json:"Destination,omitempty"`
	EventTypes          []EventType              `json:"EventTypes,omitempty"`
	RegistryPrefixes    []string                 `json:"RegistryPrefixes,omitempty"`
	ResourceTypes       []string                 `json:"ResourceTypes,omitempty"`
	DeliveryRetryPolicy DeliveryRetryPolicy      `json:"DeliveryRetryPolicy,omitempty"`
	HTTPHeaders         map[string]string        `json:"HttpHeaders,omitempty"`
	Oem                 interface{}              `json:"Oem,omitempty"`
	Protocol            EventDestinationProtocol `json:"Protocol,omitempty"`
	Context             string                   `json:"Context,omitempty"`
}

// validateCreateEventDestinationParams will validate
// CreateEventDestination parameters

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

// CreateEventDestination will create a EventDestination instance.
// destination should contain the URL of the destination for events to be sent.
// eventTypes is a list of EventType to subscribe to.
// httpHeaders is optional and gives the opportunity to specify any arbitrary
// HTTP headers required for the event POST operation.
// protocol should be the communication protocol of the event endpoint,
// usually RedfishEventDestinationProtocol
// context is a required client-supplied string that is sent with the event notifications
// oem is optional and gives the opportunity to specify any OEM specific properties,
// it should contain the vendor specific struct that goes inside the Oem session.
// It returns the new subscription URI if the event subscription is created
// with success or any error encountered.

// Deprecated: (v1.5) EventType-based eventing is DEPRECATED in the Redfish schema
// in favor of using RegistryPrefix and ResourceTypes
func CreateEventDestination(
	c common.Client,
	uri string,
	destination string,
	eventTypes []EventType,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem interface{},
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
	c common.Client,
	uri string,
	destination string,
	registryPrefixes []string,
	resourceTypes []string,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	deliveryRetryPolicy DeliveryRetryPolicy,
	oem interface{},
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
	c common.Client,
	s *subscriptionPayload,
	uri string,
	destination string,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem interface{},
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
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// return subscription link from returned location
	subscriptionLink := resp.Header.Get("Location")
	urlParser, err := url.ParseRequestURI(subscriptionLink)
	if err == nil {
		subscriptionLink = urlParser.RequestURI()
	}

	return subscriptionLink, err
}

// DeleteEventDestination will delete a EventDestination.
func DeleteEventDestination(c common.Client, uri string) error {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}

	resp, err := c.Delete(uri)
	if err == nil {
		defer resp.Body.Close()
	}

	return err
}

// ListReferencedEventDestinations gets the collection of EventDestination from
// a provided reference.
func ListReferencedEventDestinations(c common.Client, link string) ([]*EventDestination, error) {
	var result []*EventDestination
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *EventDestination
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()

	get := func(link string) {
		eventdestination, err := GetEventDestination(c, link)
		ch <- GetResult{Item: eventdestination, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// HTTPHeaderProperty shall a names and value of an HTTP header to be included
// with every event POST to the Event Destination.
type HTTPHeaderProperty map[string][]string
