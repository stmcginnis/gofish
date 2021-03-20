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

// SNMPAuthenticationProtocols is
type SNMPAuthenticationProtocols string

const (

	// NoneSNMPAuthenticationProtocols shall indicate authentication is not
	// required.
	NoneSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "None"
	// CommunityStringSNMPAuthenticationProtocols shall indicate
	// authentication using SNMP community strings and the value of
	// TrapCommunity.
	CommunityStringSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "CommunityString"
	// HMACMD5SNMPAuthenticationProtocols shall indicate authentication
	// conforms to the RFC3414-defined HMAC-MD5-96 authentication protocol.
	HMACMD5SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_MD5"
	// HMACSHA96SNMPAuthenticationProtocols shall indicate authentication
	// conforms to the RFC3414-defined HMAC-SHA-96 authentication protocol.
	HMACSHA96SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_SHA96"
)

// SNMPEncryptionProtocols is
type SNMPEncryptionProtocols string

const (

	// NoneSNMPEncryptionProtocols shall indicate there is no encryption.
	NoneSNMPEncryptionProtocols SNMPEncryptionProtocols = "None"
	// CBCDESSNMPEncryptionProtocols shall indicate encryption conforms to
	// the RFC3414-defined CBC-DES encryption protocol.
	CBCDESSNMPEncryptionProtocols SNMPEncryptionProtocols = "CBC_DES"
	// CFB128AES128SNMPEncryptionProtocols shall indicate encryption
	// conforms to the RFC3826-defined CFB128-AES-128 encryption protocol.
	CFB128AES128SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES128"
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

	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventdestination EventDestination
	err = json.NewDecoder(resp.Body).Decode(&eventdestination)
	if err != nil {
		return nil, err
	}

	eventdestination.SetClient(c)
	return &eventdestination, nil
}

// subscriptionPayload is the payload to create the event subscription
type subscriptionPayload struct {
	Destination string                   `json:"Destination"`
	EventTypes  []EventType              `json:"EventTypes"`
	HTTPHeaders map[string]string        `json:"HttpHeaders,omitempty"`
	Oem         interface{}              `json:"Oem,omitempty"`
	Protocol    EventDestinationProtocol `json:"Protocol,omitempty"`
	Context     string                   `json:"Context,omitempty"`
}

// validateCreateEventDestinationParams will validate
// CreateEventDestination parameters
func validateCreateEventDestinationParams(
	uri string,
	destination string,
	eventTypes []EventType,
) error {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}

	// validate destination
	if strings.TrimSpace(destination) == "" {
		return fmt.Errorf("empty destination is not valid")
	}

	if !strings.HasPrefix(destination, "http") {
		return fmt.Errorf("destination should start with http")
	}

	// validate event types
	if len(eventTypes) == 0 {
		return fmt.Errorf("at least one event type for subscription should be defined")
	}

	for _, et := range eventTypes {
		if !et.IsValidEventType() {
			return fmt.Errorf("invalid event type")
		}
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
	// validate input parameters
	err := validateCreateEventDestinationParams(
		uri,
		destination,
		eventTypes,
	)

	if err != nil {
		return "", err
	}

	// create subscription payload
	s := &subscriptionPayload{
		Destination: destination,
		EventTypes:  eventTypes,
		Protocol:    protocol,
		Context:     context,
	}

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
	if urlParser, err := url.ParseRequestURI(subscriptionLink); err == nil {
		subscriptionLink = urlParser.RequestURI()
	}

	return subscriptionLink, nil
}

// DeleteEventDestination will delete a EventDestination.
func DeleteEventDestination(c common.Client, uri string) (err error) {
	// validate uri
	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri should not be empty")
	}
	_, err = c.Delete(uri)

	return err
}

// ListReferencedEventDestinations gets the collection of EventDestination from
// a provided reference.
func ListReferencedEventDestinations(c common.Client, link string) ([]*EventDestination, error) {
	var result []*EventDestination
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, eventdestinationLink := range links.ItemLinks {
		eventdestination, err := GetEventDestination(c, eventdestinationLink)
		if err != nil {
			return result, err
		}
		result = append(result, eventdestination)
	}

	return result, nil
}

// HTTPHeaderProperty shall a names and value of an HTTP header to be included
// with every event POST to the Event Destination.
type HTTPHeaderProperty map[string][]string

// SNMPSettings is shall contain the settings for an SNMP event
// destination.
type SNMPSettings struct {

	// AuthenticationKey is used for SNMPv3 authentication. The value shall
	// be `null` in responses.
	AuthenticationKey string
	// AuthenticationProtocol is This property shall contain the SNMPv3
	// authentication protocol.
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey is This property shall contain the key for SNMPv3
	// encryption. The value shall be `null` in responses.
	EncryptionKey string
	// EncryptionProtocol is This property shall contain the SNMPv3
	// encryption protocol.
	EncryptionProtocol SNMPEncryptionProtocols
	// TrapCommunity is This property shall contain the SNMP trap community
	// string. The value shall be `null` in responses.
	TrapCommunity string
}
