//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/stmcginnis/gofish/common"
)

// EventFormatType is
type EventFormatType string

const (

	// EventEventFormatType The subscription destination will receive JSON
	// Bodies of the Resource Type Event.
	EventEventFormatType EventFormatType = "Event"
	// MetricReportEventFormatType The subscription destination will receive
	// JSON Bodies of the Resource Type MetricReport.
	MetricReportEventFormatType EventFormatType = "MetricReport"
)

// EventType is
type EventType string

const (
	// AlertEventType indicates a condition exists which requires attention.
	AlertEventType EventType = "Alert"
	// ResourceAddedEventType indicates a resource has been added.
	ResourceAddedEventType EventType = "ResourceAdded"
	// ResourceRemovedEventType indicates a resource has been removed.
	ResourceRemovedEventType EventType = "ResourceRemoved"
	// ResourceUpdatedEventType indicates a resource has been updated.
	ResourceUpdatedEventType EventType = "ResourceUpdated"
	// StatusChangeEventType indicates the status of this resource has changed.
	StatusChangeEventType EventType = "StatusChange"
)

// IsValidEventType will check if it is a valid EventType
func (et EventType) IsValidEventType() bool {
	switch et {
	case AlertEventType, ResourceAddedEventType,
		ResourceRemovedEventType, ResourceUpdatedEventType,
		StatusChangeEventType:
		return true
	}
	return false
}

// SupportedEventTypes contains a map of supported EventType
var SupportedEventTypes = map[string]EventType{
	"Alert":                    AlertEventType,
	"ResourceAdded":            ResourceAddedEventType,
	"ResourceRemovedEventType": ResourceRemovedEventType,
	"ResourceUpdated":          ResourceUpdatedEventType,
	"StatusChange":             StatusChangeEventType,
}

// SMTPAuthenticationMethods is
type SMTPAuthenticationMethods string

const (

	// NoneSMTPAuthenticationMethods shall indicate authentication is not
	// required.
	NoneSMTPAuthenticationMethods SMTPAuthenticationMethods = "None"
	// AutoDetectSMTPAuthenticationMethods shall indicate authentication is
	// auto-detected.
	AutoDetectSMTPAuthenticationMethods SMTPAuthenticationMethods = "AutoDetect"
	// PlainSMTPAuthenticationMethods shall indicate authentication conforms
	// to the RFC4954-defined AUTH PLAIN mechanism.
	PlainSMTPAuthenticationMethods SMTPAuthenticationMethods = "Plain"
	// LoginSMTPAuthenticationMethods shall indicate authentication conforms
	// to the RFC4954-defined AUTH LOGIN mechanism.
	LoginSMTPAuthenticationMethods SMTPAuthenticationMethods = "Login"
	// CRAMMD5SMTPAuthenticationMethods shall indicate authentication
	// conforms to the RFC4954-defined AUTH CRAM-MD5 mechanism.
	CRAMMD5SMTPAuthenticationMethods SMTPAuthenticationMethods = "CRAM_MD5"
)

// SMTPConnectionProtocol is
type SMTPConnectionProtocol string

const (

	// NoneSMTPConnectionProtocol shall indicate the connection is in clear
	// text.
	NoneSMTPConnectionProtocol SMTPConnectionProtocol = "None"
	// AutoDetectSMTPConnectionProtocol shall indicate the connection is
	// auto-detected.
	AutoDetectSMTPConnectionProtocol SMTPConnectionProtocol = "AutoDetect"
	// StartTLSSMTPConnectionProtocol shall indicate the connection conforms
	// to the RFC3207-defined StartTLS extension.
	StartTLSSMTPConnectionProtocol SMTPConnectionProtocol = "StartTLS"
	// TLSSSLSMTPConnectionProtocol shall indicate the connection is
	// TLS/SSL.
	TLSSSLSMTPConnectionProtocol SMTPConnectionProtocol = "TLS_SSL"
)

// EventService is used to represent an event service for a Redfish
// implementation.
type EventService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// DeliveryRetryAttempts shall be the
	// number of retries attempted for any given event to the subscription
	// destination before the subscription is terminated.  This retry is at
	// the service level, meaning the HTTP POST to the Event Destination was
	// returned by the HTTP operation as unsuccessful (4xx or 5xx return
	// code) or an HTTP timeout occurred this many times before the Event
	// Destination subscription is terminated.
	DeliveryRetryAttempts int
	// DeliveryRetryIntervalSeconds shall be the interval in seconds between the
	// retry attempts for any given event
	// to the subscription destination.
	DeliveryRetryIntervalSeconds int
	// Description provides a description of this resource.
	Description string
	// EventFormatTypes shall indicate the
	// content types of the message that this service can send to the event
	// destination.  If this property is not present, the EventFormatType
	// shall be assumed to be Event.
	EventFormatTypes []EventFormatType
	// EventTypesForSubscription is the types of Events
	// that can be subscribed to.
	EventTypesForSubscription []EventType
	// IncludeOriginOfConditionSupported shall indicate
	// whether the service supports including the resource payload of the
	// origin of condition in the event payload.  If `true`, event
	// subscriptions are allowed to specify the IncludeOriginOfCondition
	// property.
	IncludeOriginOfConditionSupported bool
	// RegistryPrefixes is the array of the Prefixes of the Message Registries
	// that shall be allowed for an Event Subscription.
	RegistryPrefixes []string
	// ResourceTypes is used for an Event Subscription.
	ResourceTypes []string
	// SMTP shall contain settings for SMTP event delivery.
	SMTP SMTP
	// SSEFilterPropertiesSupported shall contain a set of properties that
	// indicate which properties are supported in the $filter query parameter
	// for the URI indicated by the ServerSentEventUri property.
	SSEFilterPropertiesSupported SSEFilterPropertiesSupported
	// ServerSentEventURI shall be a URI that specifies an HTML5 Server-Sent
	// Event conformant endpoint.
	ServerSentEventURI string `json:"ServerSentEventUri"`
	// ServiceEnabled shall be a boolean indicating whether this service is enabled.
	ServiceEnabled bool
	// Status is This property shall contain any status or health properties of
	// the resource.
	Status common.Status
	// SubordinateResourcesSupported is When set to true, the service is
	// indicating that it supports the SubordinateResource property on Event
	// Subscriptions and on generated Events.
	SubordinateResourcesSupported bool
	// Subscriptions shall contain the link to a collection of type
	// EventDestination.
	subscriptions string
	// submitTestEventTarget is the URL to send SubmitTestEvent actions.
	submitTestEventTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EventService object from the raw JSON.
func (eventservice *EventService) UnmarshalJSON(b []byte) error { // nolint:dupl
	type temp EventService
	type Actions struct {
		SubmitTestEvent struct {
			Target string
		} `json:"#EventService.SubmitTestEvent"`
	}
	var t struct {
		temp
		Subscriptions common.Link
		Actions       Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*eventservice = EventService(t.temp)
	eventservice.subscriptions = string(t.Subscriptions)
	eventservice.submitTestEventTarget = t.Actions.SubmitTestEvent.Target

	// This is a read/write object, so we need to save the raw object data for later
	eventservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (eventservice *EventService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EventService)
	err := original.UnmarshalJSON(eventservice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"DeliveryRetryAttempts",
		"DeliveryRetryIntervalSeconds",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(eventservice).Elem()

	return eventservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEventService will get a EventService instance from the service.
func GetEventService(c common.Client, uri string) (*EventService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventservice EventService
	err = json.NewDecoder(resp.Body).Decode(&eventservice)
	if err != nil {
		return nil, err
	}

	eventservice.SetClient(c)
	return &eventservice, nil
}

// ListReferencedEventServices gets the collection of EventService from
// a provided reference.
func ListReferencedEventServices(c common.Client, link string) ([]*EventService, error) {
	var result []*EventService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, eventserviceLink := range links.ItemLinks {
		eventservice, err := GetEventService(c, eventserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, eventservice)
	}

	return result, nil
}

// GetEventSubscriptions gets all the subscriptions using the event service.
func (eventservice *EventService) GetEventSubscriptions() ([]*EventDestination, error) {
	if strings.TrimSpace(eventservice.subscriptions) == "" {
		return nil, fmt.Errorf("empty subscription link in the event service")
	}

	return ListReferencedEventDestinations(eventservice.Client, eventservice.subscriptions)
}

// GetEventSubscription gets a specific subscription using the event service.
func (eventservice *EventService) GetEventSubscription(uri string) (*EventDestination, error) {
	return GetEventDestination(eventservice.Client, uri)
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
func (eventservice *EventService) CreateEventSubscription(
	destination string,
	eventTypes []EventType,
	httpHeaders map[string]string,
	protocol EventDestinationProtocol,
	context string,
	oem interface{},
) (string, error) {
	if strings.TrimSpace(eventservice.subscriptions) == "" {
		return "", fmt.Errorf("empty subscription link in the event service")
	}

	return CreateEventDestination(
		eventservice.Client,
		eventservice.subscriptions,
		destination,
		eventTypes,
		httpHeaders,
		protocol,
		context,
		oem,
	)
}

// DeleteEventSubscription deletes a specific subscription using the event service.
func (eventservice *EventService) DeleteEventSubscription(uri string) error {
	return DeleteEventDestination(eventservice.Client, uri)
}

// SubmitTestEvent shall add a test event to the event service with the event
// data specified in the action parameters. This message should then be sent to
// any appropriate ListenerDestination targets.
func (eventservice *EventService) SubmitTestEvent(message string) error {
	type temp struct {
		EventGroupID      string `json:"EventGroupId"`
		EventID           string `json:"EventId"`
		EventTimestamp    string
		EventType         string
		Message           string
		MessageArgs       []string
		MessageID         string `json:"MessageId"`
		OriginOfCondition string
		Severity          string
	}
	t := temp{
		EventGroupID:      "TESTING123",
		EventID:           "TEST123",
		EventTimestamp:    time.Now().String(),
		EventType:         "Alert",
		Message:           message,
		MessageID:         "test123",
		OriginOfCondition: eventservice.ODataID,
		Severity:          "Informational",
	}

	_, err := eventservice.Client.Post(eventservice.submitTestEventTarget, t)
	return err
}

// SSEFilterPropertiesSupported shall contain a set of properties that indicate
// which properties are supported in the $filter query parameter for the URI
// indicated by the ServerSentEventUri property.
type SSEFilterPropertiesSupported struct {
	// EventFormatType shall be a boolean indicating if this service supports
	// the use of the EventFormatType property in the $filter query parameter as
	// described by the specification.
	EventFormatType bool
	// MessageID shall be a boolean indicating if this service supports the use
	// of the MessageId property in the $filter query parameter as described by
	// the specification.
	MessageID bool `json:"MessageId"`
	// MetricReportDefinition shall be a boolean indicating if this service
	// supports the use of the MetricReportDefinition property in the $filter
	// query parameter as described by the specification.
	MetricReportDefinition bool
	// OriginResource shall be a boolean indicating if this service supports the
	// use of the OriginResource property in the $filter query parameter as
	// described by the specification.
	OriginResource bool
	// RegistryPrefix shall be a boolean indicating if this service supports the
	// use of the RegistryPrefix property in the $filter query parameter as
	// described by the specification.
	RegistryPrefix bool
	// ResourceType shall be a boolean indicating if this service supports the
	// use of the ResourceType property in the $filter query parameter as
	// described by the specification.
	ResourceType bool
}

// SMTP is shall contain settings for SMTP event delivery.
type SMTP struct {

	// Authentication shall contain the authentication
	// method for the SMTP server.
	Authentication SMTPAuthenticationMethods
	// ConnectionProtocol shall contain the connection type
	// to the outgoing SMTP server.
	ConnectionProtocol SMTPConnectionProtocol
	// FromAddress shall contain the email address to use
	// for the 'from' field in an outgoing email.
	FromAddress string
	// Password shall contain the password for
	// authentication with the SMTP server. The value shall be `null` in
	// responses.
	Password string
	// Port shall contain the destination port for the SMTP
	// server.
	Port int
	// ServerAddress shall contain the address of the SMTP
	// server for outgoing email.
	ServerAddress string
	// ServiceEnabled shall indicate if SMTP for event
	// delivery is enabled.
	ServiceEnabled bool
	// Username shall contain the username for
	// authentication with the SMTP server.
	Username string
}
