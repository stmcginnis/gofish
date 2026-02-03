//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #EventService.v1_12_0.EventService

package schemas

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type SMTPAuthenticationMethods string

const (
	// NoneSMTPAuthenticationMethods shall indicate authentication is not required.
	NoneSMTPAuthenticationMethods SMTPAuthenticationMethods = "None"
	// AutoDetectSMTPAuthenticationMethods shall indicate authentication is
	// auto-detected.
	AutoDetectSMTPAuthenticationMethods SMTPAuthenticationMethods = "AutoDetect"
	// PlainSMTPAuthenticationMethods shall indicate authentication conforms to the
	// RFC4954-defined AUTH PLAIN mechanism.
	PlainSMTPAuthenticationMethods SMTPAuthenticationMethods = "Plain"
	// LoginSMTPAuthenticationMethods shall indicate authentication conforms to the
	// RFC4954-defined AUTH LOGIN mechanism.
	LoginSMTPAuthenticationMethods SMTPAuthenticationMethods = "Login"
	// CRAMMD5SMTPAuthenticationMethods shall indicate authentication conforms to
	// the RFC4954-defined AUTH CRAM-MD5 mechanism.
	CRAMMD5SMTPAuthenticationMethods SMTPAuthenticationMethods = "CRAM_MD5"
)

type SMTPConnectionProtocol string

const (
	// NoneSMTPConnectionProtocol shall indicate the connection is in clear text.
	NoneSMTPConnectionProtocol SMTPConnectionProtocol = "None"
	// AutoDetectSMTPConnectionProtocol shall indicate the connection is
	// auto-detected.
	AutoDetectSMTPConnectionProtocol SMTPConnectionProtocol = "AutoDetect"
	// StartTLSSMTPConnectionProtocol shall indicate the connection conforms to the
	// RFC3207-defined StartTLS extension.
	StartTLSSMTPConnectionProtocol SMTPConnectionProtocol = "StartTLS"
	// TLSSSLSMTPConnectionProtocol shall indicate the connection is TLS/SSL.
	TLSSSLSMTPConnectionProtocol SMTPConnectionProtocol = "TLS_SSL"
)

// EventService shall represent an event service for a Redfish implementation.
type EventService struct {
	Entity
	// DeliveryRetryAttempts shall contain the number of times that the 'POST' of
	// an event is retried before the subscription terminates or is suspended. This
	// retry occurs at the service level, which means that the HTTP 'POST' to the
	// event destination fails with an HTTP '4XX' or '5XX' status code or an HTTP
	// timeout occurs this many times before the event destination subscription
	// terminates or is suspended. The service shall delete the 'EventDestination'
	// resource to terminate the subscription. The service shall set the value of
	// the 'State' property within 'Status' of the 'EventDestination' resource to
	// 'Disabled' for a suspended subscription.
	DeliveryRetryAttempts int
	// DeliveryRetryIntervalSeconds shall contain the interval, in seconds, between
	// the retry attempts for any event sent to the subscription destination.
	DeliveryRetryIntervalSeconds int
	// EventFormatTypes shall contain the content types of the message that this
	// service can send to the event destination. If this property is not present,
	// the 'EventFormatType' shall be assumed to be 'Event'.
	//
	// Version added: v1.2.0
	EventFormatTypes []EventFormatType
	// EventTypesForSubscription shall contain the types of events to which a
	// client can subscribe. The semantics associated with the enumeration values
	// are defined in the Redfish Specification.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated. Starting with Redfish Specification v1.6
	// (Event v1.3), subscriptions are based on the 'RegistryPrefix' and
	// 'ResourceType' properties and not on the 'EventType' property.
	EventTypesForSubscription []EventType
	// ExcludeMessageID shall indicate whether this service supports filtering by
	// the 'ExcludeMessageIds' property.
	//
	// Version added: v1.8.0
	ExcludeMessageID bool `json:"ExcludeMessageId"`
	// ExcludeRegistryPrefix shall indicate whether this service supports filtering
	// by the 'ExcludeRegistryPrefixes' property.
	//
	// Version added: v1.8.0
	ExcludeRegistryPrefix bool
	// IncludeOriginOfConditionSupported shall indicate whether the service
	// supports including the resource payload of the origin of condition in the
	// event payload. If 'true', event subscriptions are allowed to specify the
	// 'IncludeOriginOfCondition' property.
	//
	// Version added: v1.6.0
	IncludeOriginOfConditionSupported bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginResourcesSupported shall indicate whether this service supports
	// filtering by the 'OriginResources' property.
	//
	// Version added: v1.12.0
	OriginResourcesSupported bool
	// RegistryPrefixes shall contain the array of the prefixes of the message
	// registries that shall be allowed or excluded for an event subscription.
	//
	// Version added: v1.2.0
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of the valid '@odata.type' values that
	// can be used for an event subscription.
	//
	// Version added: v1.2.0
	ResourceTypes []string
	// SMTP shall contain settings for SMTP event delivery.
	//
	// Version added: v1.5.0
	SMTP SMTP
	// SSEFilterPropertiesSupported shall contain the properties that are supported
	// in the '$filter' query parameter for the URI indicated by the
	// 'ServerSentEventUri' property, as described by the Redfish Specification.
	//
	// Version added: v1.2.0
	SSEFilterPropertiesSupported SSEFilterPropertiesSupported
	// SSEIncludeOriginOfConditionSupported shall indicate whether the service
	// supports the 'includeoriginofcondition' query parameter for the
	// 'ServerSentEventUri', as described by the Redfish Specification.
	//
	// Version added: v1.11.0
	SSEIncludeOriginOfConditionSupported bool
	// ServerSentEventURI shall contain a URI that specifies an HTML5 Server-Sent
	// Event-conformant endpoint.
	//
	// Version added: v1.1.0
	ServerSentEventURI string `json:"ServerSentEventUri"`
	// ServiceEnabled shall indicate whether this service is enabled. If 'false',
	// events are no longer published, new SSE connections cannot be established,
	// and existing SSE connections are terminated.
	ServiceEnabled bool
	// Severities shall specify an array of the allowable severities that can be
	// used for an event subscription. If this property is absent or contains an
	// empty array, the service does not support severity-based subscriptions.
	//
	// Version added: v1.9.0
	Severities []Health
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SubordinateResourcesSupported shall indicate whether the service supports
	// the 'SubordinateResources' property on both event subscriptions and
	// generated events.
	//
	// Version added: v1.2.0
	SubordinateResourcesSupported bool
	// Subscriptions shall contain the link to a resource collection of type
	// 'EventDestinationCollection'.
	SubscriptionsLink string
	// submitTestEventTarget is the URL to send SubmitTestEvent requests.
	SubmitTestEventTarget string
	// testEventSubscriptionTarget is the URL to send TestEventSubscription requests.
	testEventSubscriptionTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

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

// Update commits updates to this object's properties to the running system.
func (e *EventService) Update() error {
	readWriteFields := []string{
		"DeliveryRetryAttempts",
		"DeliveryRetryIntervalSeconds",
		"ServiceEnabled",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}

// GetEventService will get a EventService instance from the service.
func GetEventService(c Client, uri string) (*EventService, error) {
	return GetObject[EventService](c, uri)
}

// ListReferencedEventServices gets the collection of EventService from
// a provided reference.
func ListReferencedEventServices(c Client, link string) ([]*EventService, error) {
	return GetCollectionObjects[EventService](c, link)
}

// CreateEventSubscription creates the subscription using the event service.
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

// EventServiceSubmitTestEventParameters holds the parameters for the SubmitTestEvent action.
type EventServiceSubmitTestEventParameters struct {
	// EventGroupID shall contain the group identifier for the event. It has the
	// same semantics as the 'EventGroupId' property in the 'Event' schema for
	// Redfish. If not provided by the client, the resulting event should not
	// contain the 'EventGroupId' property.
	EventGroupID int `json:"EventGroupId,omitempty"`
	// EventID shall have the same semantics as the 'EventId' property in the
	// 'Event' schema for Redfish. A service can ignore this value and replace it
	// with its own. If not provided by the client, the resulting event may contain
	// a service-defined 'EventId' property.
	EventID string `json:"EventId,omitempty"`
	// EventTimestamp shall contain the date and time for the event to add and have
	// the same semantics as the 'EventTimestamp' property in the 'Event' schema
	// for Redfish. If not provided by the client, the resulting event should not
	// contain the 'EventTimestamp' property.
	EventTimestamp string `json:"EventTimestamp,omitempty"`
	// EventType shall contain the property name for which the following allowable
	// values apply. If not provided by the client, the resulting event should not
	// contain the 'EventType' property.
	EventType EventType `json:"EventType,omitempty"`
	// Message shall have the same semantics as the 'Message' property in the
	// 'Event' schema for Redfish. If not provided by the client, the resulting
	// event should not contain the 'Message' property.
	Message string `json:"Message,omitempty"`
	// MessageArgs shall have the same semantics as the 'MessageArgs' property in
	// the 'Event' schema for Redfish. If not provided by the client, the resulting
	// event should not contain the 'MessageArgs' property.
	MessageArgs []string `json:"MessageArgs,omitempty"`
	// MessageID shall contain the 'MessageId' for the event to add and have the
	// same semantics as the 'MessageId' property in the 'Event' schema for
	// Redfish. Services should accept arbitrary values for this parameter that
	// match that match the defined pattern.
	MessageID string `json:"MessageId,omitempty"`
	// MessageSeverity shall contain the severity for the event to add and have the
	// same semantics as the 'MessageSeverity' property in the 'Event' schema for
	// Redfish. If not provided by the client, the resulting event should not
	// contain the 'MessageSeverity' property.
	MessageSeverity Health `json:"MessageSeverity,omitempty"`
	// OriginOfCondition shall be a string that represents the URL contained by the
	// 'OriginOfCondition' property in the 'Event' schema for Redfish. If not
	// provided by the client, the resulting event should not contain the
	// 'OriginOfCondition' property.
	OriginOfCondition string `json:"OriginOfCondition,omitempty"`
	// Severity shall contain the severity for the event to add and have the same
	// semantics as the 'Severity' property in the 'Event' schema for Redfish. If
	// not provided by the client, the resulting event should not contain the
	// 'Severity' property.
	Severity string `json:"Severity,omitempty"`
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

// SMTP shall contain settings for SMTP event delivery.
type SMTP struct {
	// Authentication shall contain the authentication method for the SMTP server.
	//
	// Version added: v1.5.0
	Authentication SMTPAuthenticationMethods
	// ConnectionProtocol shall contain the connection type to the outgoing SMTP
	// server.
	//
	// Version added: v1.5.0
	ConnectionProtocol SMTPConnectionProtocol
	// FromAddress shall contain the email address to use for the 'from' field in
	// an outgoing email.
	//
	// Version added: v1.5.0
	FromAddress string
	// Password shall contain the password for authentication with the SMTP server.
	// The value shall be 'null' in responses.
	//
	// Version added: v1.5.0
	Password string
	// PasswordSet shall contain 'true' if a valid value was provided for the
	// 'Password' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.9.0
	PasswordSet bool
	// Port shall contain the destination port for the SMTP server.
	//
	// Version added: v1.5.0
	Port *uint `json:",omitempty"`
	// ServerAddress shall contain the address of the SMTP server for outgoing
	// email.
	//
	// Version added: v1.5.0
	ServerAddress string
	// ServiceEnabled shall indicate if SMTP for event delivery is enabled.
	//
	// Version added: v1.5.0
	ServiceEnabled bool
	// Username shall contain the username for authentication with the SMTP server.
	//
	// Version added: v1.5.0
	Username string
}

// SSEFilterPropertiesSupported shall contain a set of properties that are
// supported in the '$filter' query parameter for the URI indicated by the
// 'ServerSentEventUri' property, as described by the Redfish Specification.
type SSEFilterPropertiesSupported struct {
	// EventFormatType shall indicate whether this service supports filtering by
	// the 'EventFormatType' property.
	//
	// Version added: v1.2.0
	EventFormatType bool
	// EventType shall indicate whether this service supports filtering by the
	// 'EventTypes' property.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.3.0
	// This property has been deprecated. Starting with Redfish Specification v1.6
	// (Event v1.3), subscriptions are based on the 'RegistryPrefix' and
	// 'ResourceType' properties and not on the 'EventType' property.
	EventType bool
	// MessageID shall indicate whether this service supports filtering by the
	// 'MessageIds' property.
	//
	// Version added: v1.2.0
	MessageID bool `json:"MessageId"`
	// MetricReportDefinition shall indicate whether this service supports
	// filtering by the 'MetricReportDefinitions' property.
	//
	// Version added: v1.2.0
	MetricReportDefinition bool
	// OriginResource shall indicate whether this service supports filtering by the
	// 'OriginResources' property.
	//
	// Version added: v1.2.0
	OriginResource bool
	// RegistryPrefix shall indicate whether this service supports filtering by the
	// 'RegistryPrefixes' property.
	//
	// Version added: v1.2.0
	RegistryPrefix bool
	// ResourceType shall indicate whether this service supports filtering by the
	// 'ResourceTypes' property.
	//
	// Version added: v1.2.0
	ResourceType bool
	// SubordinateResources shall indicate whether this service supports filtering
	// by the 'SubordinateResources' property.
	//
	// Version added: v1.4.0
	SubordinateResources bool
}
