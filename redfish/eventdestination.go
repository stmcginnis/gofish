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
	// RetryForeverWithBackoffDeliveryRetryPolicy shall indicate the subscription is not suspended or terminated, and
	// attempts at delivery of future events shall continue regardless of the number of retries. Retry attempts are
	// issued over time according to a service-defined backoff algorithm. The backoff algorithm may insert an
	// increasing amount of delay between retry attempts and may reach a maximum.
	RetryForeverWithBackoffDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForeverWithBackoff"
)

// EventDestinationProtocol is the communication protocol of the event destination.
type EventDestinationProtocol string

const (
	// RedfishEventDestinationProtocol shall indicate the destination follows the Redfish Specification for event
	// notifications. Destinations requesting EventFormatType of 'Event' shall receive a Redfish resource of type
	// Event. Destinations requesting EventFormatType of 'MetricReport' shall receive a Redfish resource of type
	// MetricReport.
	RedfishEventDestinationProtocol EventDestinationProtocol = "Redfish"
	// KafkaEventDestinationProtocol shall indicate the destination follows the Apache-defined Kafka protocol as
	// defined by the Kafka Protocol Guide. The Context property shall contain the Kafka topic of the destination
	// broker.
	KafkaEventDestinationProtocol EventDestinationProtocol = "Kafka"
	// SNMPv1EventDestinationProtocol shall indicate the destination follows the RFC1157-defined SNMPv1 protocol.
	SNMPv1EventDestinationProtocol EventDestinationProtocol = "SNMPv1"
	// SNMPv2cEventDestinationProtocol shall indicate the destination follows the SNMPv2c protocol as defined by
	// RFC1441 and RFC1452.
	SNMPv2cEventDestinationProtocol EventDestinationProtocol = "SNMPv2c"
	// SNMPv3EventDestinationProtocol shall indicate the destination follows the SNMPv3 protocol as defined by RFC3411
	// and RFC3418.
	SNMPv3EventDestinationProtocol EventDestinationProtocol = "SNMPv3"
	// SMTPEventDestinationProtocol shall indicate the destination follows the RFC5321-defined SMTP specification.
	SMTPEventDestinationProtocol EventDestinationProtocol = "SMTP"
	// SyslogTLSEventDestinationProtocol shall indicate the destination follows the TLS-based transport for syslog as
	// defined in RFC5424.
	SyslogTLSEventDestinationProtocol EventDestinationProtocol = "SyslogTLS"
	// SyslogTCPEventDestinationProtocol shall indicate the destination follows the TCP-based transport for syslog as
	// defined in RFC6587.
	SyslogTCPEventDestinationProtocol EventDestinationProtocol = "SyslogTCP"
	// SyslogUDPEventDestinationProtocol shall indicate the destination follows the UDP-based transport for syslog as
	// defined in RFC5424.
	SyslogUDPEventDestinationProtocol EventDestinationProtocol = "SyslogUDP"
	// SyslogRELPEventDestinationProtocol shall indicate the destination follows the Reliable Event Logging Protocol
	// (RELP) transport for syslog as defined by www.rsyslog.com.
	SyslogRELPEventDestinationProtocol EventDestinationProtocol = "SyslogRELP"
	// OEMEventDestinationProtocol shall indicate an OEM-specific protocol. The OEMProtocol property shall contain the
	// specific OEM event destination protocol.
	OEMEventDestinationProtocol EventDestinationProtocol = "OEM"
)

// SubscriptionType is the type of subscription used.
type SubscriptionType string

const (
	// RedfishEventSubscriptionType The subscription follows the Redfish Specification for event notifications. To send
	// an event notification, a service sends an HTTP POST to the subscriber's destination URI.
	RedfishEventSubscriptionType SubscriptionType = "RedfishEvent"
	// SSESubscriptionType The subscription follows the HTML5 server-sent event definition for event notifications.
	SSESubscriptionType SubscriptionType = "SSE"
	// SNMPTrapSubscriptionType shall indicate the subscription follows the various versions of SNMP Traps for event
	// notifications. Protocol shall specify the appropriate version of SNMP.
	SNMPTrapSubscriptionType SubscriptionType = "SNMPTrap"
	// SNMPInformSubscriptionType shall indicate the subscription follows versions 2 and 3 of SNMP Inform for event
	// notifications. Protocol shall specify the appropriate version of SNMP.
	SNMPInformSubscriptionType SubscriptionType = "SNMPInform"
	// SyslogSubscriptionType shall indicate the subscription forwards syslog messages to the event destination.
	// Protocol shall specify the appropriate syslog protocol.
	SyslogSubscriptionType SubscriptionType = "Syslog"
	// OEMSubscriptionType shall indicate an OEM subscription type. The OEMSubscriptionType property shall contain the
	// specific OEM subscription type.
	OEMSubscriptionType SubscriptionType = "OEM"
)

// SyslogFacility shall specify the syslog facility codes as program types. Facility values are
// described in the RFC5424.
type SyslogFacility string

const (
	// KernSyslogFacility Kernel messages.
	KernSyslogFacility SyslogFacility = "Kern"
	// UserSyslogFacility User-level messages.
	UserSyslogFacility SyslogFacility = "User"
	// MailSyslogFacility Mail system.
	MailSyslogFacility SyslogFacility = "Mail"
	// DaemonSyslogFacility System daemons.
	DaemonSyslogFacility SyslogFacility = "Daemon"
	// AuthSyslogFacility Security/authentication messages.
	AuthSyslogFacility SyslogFacility = "Auth"
	// SyslogSyslogFacility Messages generated internally by syslogd.
	SyslogSyslogFacility SyslogFacility = "Syslog"
	// LPRSyslogFacility Line printer subsystem.
	LPRSyslogFacility SyslogFacility = "LPR"
	// NewsSyslogFacility Network news subsystem.
	NewsSyslogFacility SyslogFacility = "News"
	// UUCPSyslogFacility UUCP subsystem.
	UUCPSyslogFacility SyslogFacility = "UUCP"
	// CronSyslogFacility Clock daemon.
	CronSyslogFacility SyslogFacility = "Cron"
	// AuthprivSyslogFacility Security/authentication messages.
	AuthprivSyslogFacility SyslogFacility = "Authpriv"
	// FTPSyslogFacility FTP daemon.
	FTPSyslogFacility SyslogFacility = "FTP"
	// NTPSyslogFacility NTP subsystem.
	NTPSyslogFacility SyslogFacility = "NTP"
	// SecuritySyslogFacility Log audit.
	SecuritySyslogFacility SyslogFacility = "Security"
	// ConsoleSyslogFacility Log alert.
	ConsoleSyslogFacility SyslogFacility = "Console"
	// SolarisCronSyslogFacility Scheduling daemon.
	SolarisCronSyslogFacility SyslogFacility = "SolarisCron"
	// Local0SyslogFacility Locally used facility 0.
	Local0SyslogFacility SyslogFacility = "Local0"
	// Local1SyslogFacility Locally used facility 1.
	Local1SyslogFacility SyslogFacility = "Local1"
	// Local2SyslogFacility Locally used facility 2.
	Local2SyslogFacility SyslogFacility = "Local2"
	// Local3SyslogFacility Locally used facility 3.
	Local3SyslogFacility SyslogFacility = "Local3"
	// Local4SyslogFacility Locally used facility 4.
	Local4SyslogFacility SyslogFacility = "Local4"
	// Local5SyslogFacility Locally used facility 5.
	Local5SyslogFacility SyslogFacility = "Local5"
	// Local6SyslogFacility Locally used facility 6.
	Local6SyslogFacility SyslogFacility = "Local6"
	// Local7SyslogFacility Locally used facility 7.
	Local7SyslogFacility SyslogFacility = "Local7"
)

// SyslogSeverity shall specify the syslog severity levels as an application-specific rating used to
// describe the urgency of the message. 'Emergency' should be reserved for messages indicating the system is
// unusable and 'Debug' should only be used when debugging a program. Severity values are described in RFC5424.
type SyslogSeverity string

const (
	// EmergencySyslogSeverity A panic condition.
	EmergencySyslogSeverity SyslogSeverity = "Emergency"
	// AlertSyslogSeverity A condition that should be corrected immediately, such as a corrupted system database.
	AlertSyslogSeverity SyslogSeverity = "Alert"
	// CriticalSyslogSeverity Hard device errors.
	CriticalSyslogSeverity SyslogSeverity = "Critical"
	// ErrorSyslogSeverity An Error.
	ErrorSyslogSeverity SyslogSeverity = "Error"
	// WarningSyslogSeverity A Warning.
	WarningSyslogSeverity SyslogSeverity = "Warning"
	// NoticeSyslogSeverity Conditions that are not error conditions, but that might require special handling.
	NoticeSyslogSeverity SyslogSeverity = "Notice"
	// InformationalSyslogSeverity Informational only.
	InformationalSyslogSeverity SyslogSeverity = "Informational"
	// DebugSyslogSeverity Messages that contain information normally of use only when debugging a program.
	DebugSyslogSeverity SyslogSeverity = "Debug"
	// AllSyslogSeverity A message of any severity.
	AllSyslogSeverity SyslogSeverity = "All"
)

// SyslogFilter shall contain the filter for a syslog message. The filter shall describe the desired syslog message
// to forward to a remote syslog server.
type SyslogFilter struct {
	// LogFacilities shall contain the types of programs that can log messages. If this property contains an empty
	// array or is absent, all facilities shall be indicated.
	LogFacilities []SyslogFacility
	// LowestSeverity shall contain the lowest syslog severity level that will be forwarded. The service shall forward
	// all messages equal to or greater than the value in this property. The value 'All' shall indicate all severities.
	LowestSeverity SyslogSeverity
}

// EventDestination is used to represent the target of an event
// subscription, including the types of events subscribed and context to
// provide to the target in the Event payload.
type EventDestination struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represent the
	// server certificates for the server referenced by the Destination property. If VerifyCertificate is 'true',
	// services shall compare the certificates in this collection with the certificate obtained during handshaking with
	// the event destination in order to verify the identity of the event destination prior to sending an event. If the
	// server cannot be verified, the service shall not send the event. If VerifyCertificate is 'false', the service
	// shall not perform certificate verification with certificates in this collection. Regardless of the contents of
	// this collection, services may perform additional verification based on other factors, such as the configuration
	// of the SecurityPolicy resource.
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the Destination property as part
	// of TLS handshaking.
	clientCertificates string
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
	// EventTypes contains the types of events that will be sent to the destination.
	// This property has been deprecated. Starting with Redfish Specification v1.6 (Event v1.3),
	// subscriptions are based on the RegistryPrefix and ResourceType properties and not on the
	// EventType property. Use EventFormatType to create subscriptions for metric reports.
	EventTypes []EventType
	// ExcludeMessageIDs shall contain an array of excluded MessageIds that are not allowed values for the MessageId
	// property within an event sent to the subscriber. The MessageId shall be in the
	// 'MessageRegistryPrefix.MessageKey' format. If included, the MessageId major and minor version details should be
	// ignored. Events with a MessageId that is contained in this array shall not be sent to the subscriber. If this
	// property is an empty array or is absent, no exclusive filtering based upon the MessageId of an event is
	// performed.
	ExcludeMessageIDs []string
	// ExcludeRegistryPrefixes shall contain an array of prefixes of excluded message registries that contain the
	// MessageIds that are not allowed values for the MessageId property within an event sent to the subscriber. Events
	// with a MessageId that is from a message registry contained in this array shall not be sent to the subscriber. If
	// this property is an empty array or is absent, no exclusive filtering based upon message registry of the
	// MessageId of an event is performed.
	ExcludeRegistryPrefixes []string
	// HeartbeatIntervalMinutes shall indicate the interval for sending periodic heartbeat events to the subscriber.
	// The value shall be the interval, in minutes, between each periodic event. This property shall not be present if
	// the SendHeartbeat property is not present.
	HeartbeatIntervalMinutes int
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
	// MetricReportDefinitions shall specify an array of metric report definitions that are the only allowable
	// generators of metric reports for this subscription. Metric reports originating from metric report definitions
	// not contained in this array shall not be sent to the subscriber. If this property is absent or the array is
	// empty, the service shall send metric reports originating from any metric report definition to the subscriber.
	MetricReportDefinitions []string
	// MetricReportDefinitionsCount is the number of MetricReportDefinitions.
	MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
	// OEMProtocol shall contain the protocol type that the event uses to send the event to the destination. This
	// property shall be present if Protocol is 'OEM'.
	OEMProtocol string
	// OEMSubscriptionType shall indicate the OEM-defined type of subscription for events. This property shall be
	// present if SubscriptionType is 'OEM'.
	OEMSubscriptionType string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginResources shall specify an array of Resources, Resource Collections,
	// or Referenceable Members that are the only allowable values for the
	// OriginOfCondition property within an EventRecord sent to the subscriber.
	// Events originating from Resources, Resource Collections, or Referenceable
	// Members not contained in this array shall not be sent to the subscriber.
	// If this property is absent or the array is empty, the service shall send
	// Events originating from any Resource, Resource Collection, or
	// Referenceable Member to the subscriber.
	OriginResources []string
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
	// SyslogFilters shall describe all desired syslog messages to send to a remote syslog server. If this property
	// contains an empty array or is absent, all messages shall be sent.
	SyslogFilters []SyslogFilter
	// VerifyCertificate shall indicate whether the service will verify the certificate of the server referenced by the
	// Destination property prior to sending the event with the certificates found in the collection referenced by the
	// Certificates property. If this property is not supported by the service or specified by the client in the create
	// request, it shall be assumed to be 'false'. Regardless of the value of this property, services may perform
	// additional verification based on other factors, such as the configuration of the SecurityPolicy resource.
	VerifyCertificate bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EventDestination object from the raw JSON.
func (eventdestination *EventDestination) UnmarshalJSON(b []byte) error {
	type temp EventDestination
	var t struct {
		temp
		Certificates       common.Link
		ClientCertificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*eventdestination = EventDestination(t.temp)
	eventdestination.certificates = t.Certificates.String()
	eventdestination.clientCertificates = t.ClientCertificates.String()

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
		"VerifyCertificate",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(eventdestination).Elem()

	return eventdestination.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEventDestination will get a EventDestination instance from the service.
func GetEventDestination(c common.Client, uri string) (*EventDestination, error) {
	return common.GetObject[EventDestination](c, uri)
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
	return common.GetCollectionObjects[EventDestination](c, link)
}

// HTTPHeaderProperty shall a names and value of an HTTP header to be included
// with every event POST to the Event Destination.
type HTTPHeaderProperty map[string][]string

// Certificates gets the server certificates for the server referenced by the Destination property.
func (eventdestination *EventDestination) Certificates() ([]*Certificate, error) {
	if eventdestination.certificates == "" {
		return []*Certificate{}, nil
	}

	return ListReferencedCertificates(eventdestination.GetClient(), eventdestination.certificates)
}

// ClientCertificates gets the client identity certificates for the server referenced by the Destination property.
func (eventdestination *EventDestination) ClientCertificates() ([]*Certificate, error) {
	if eventdestination.clientCertificates == "" {
		return []*Certificate{}, nil
	}

	return ListReferencedCertificates(eventdestination.GetClient(), eventdestination.clientCertificates)
}
