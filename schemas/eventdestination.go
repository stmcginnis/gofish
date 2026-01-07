//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #EventDestination.v1_16_0.EventDestination

package schemas

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type DeliveryRetryPolicy string

const (
	// TerminateAfterRetriesDeliveryRetryPolicy shall indicate the subscription is
	// terminated after the maximum number of retries is reached, specified by the
	// 'DeliveryRetryAttempts' property in the event service. The service shall
	// delete this resource to terminate the subscription.
	TerminateAfterRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "TerminateAfterRetries"
	// SuspendRetriesDeliveryRetryPolicy shall indicate the subscription is
	// suspended after the maximum number of retries is reached, specified by the
	// 'DeliveryRetryAttempts' property in the event service. The value of the
	// 'State' property within 'Status' shall contain 'Disabled' for a suspended
	// subscription.
	SuspendRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "SuspendRetries"
	// RetryForeverDeliveryRetryPolicy shall indicate the subscription is not
	// suspended or terminated, and attempts at delivery of future events shall
	// continue regardless of the number of retries. The interval between retries
	// remains constant and is specified by the 'DeliveryRetryIntervalSeconds'
	// property in the event service.
	RetryForeverDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForever"
	// RetryForeverWithBackoffDeliveryRetryPolicy shall indicate the subscription
	// is not suspended or terminated, and attempts at delivery of future events
	// shall continue regardless of the number of retries. Retry attempts are
	// issued over time according to a service-defined backoff algorithm. The
	// backoff algorithm may insert an increasing amount of delay between retry
	// attempts and may reach a maximum.
	RetryForeverWithBackoffDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForeverWithBackoff"
)

type EventDestinationProtocol string

const (
	// RedfishEventDestinationProtocol shall indicate the destination follows the
	// Redfish Specification for event notifications. Destinations requesting
	// 'EventFormatType' of 'Event' shall receive a Redfish resource of type
	// 'Event'. Destinations requesting 'EventFormatType' of 'MetricReport' shall
	// receive a Redfish resource of type 'MetricReport'.
	RedfishEventDestinationProtocol EventDestinationProtocol = "Redfish"
	// KafkaEventDestinationProtocol shall indicate the destination follows the
	// Apache-defined Kafka protocol as defined by the Kafka Protocol Guide. The
	// 'Context' property shall contain the Kafka topic of the destination broker.
	KafkaEventDestinationProtocol EventDestinationProtocol = "Kafka"
	// SNMPv1EventDestinationProtocol shall indicate the destination follows the
	// RFC1157-defined SNMPv1 protocol.
	SNMPv1EventDestinationProtocol EventDestinationProtocol = "SNMPv1"
	// SNMPv2cEventDestinationProtocol shall indicate the destination follows the
	// SNMPv2c protocol as defined by RFC1441 and RFC1452.
	SNMPv2cEventDestinationProtocol EventDestinationProtocol = "SNMPv2c"
	// SNMPv3EventDestinationProtocol shall indicate the destination follows the
	// SNMPv3 protocol as defined by RFC3411 and RFC3418.
	SNMPv3EventDestinationProtocol EventDestinationProtocol = "SNMPv3"
	// SMTPEventDestinationProtocol shall indicate the destination follows the
	// RFC5321-defined SMTP specification.
	SMTPEventDestinationProtocol EventDestinationProtocol = "SMTP"
	// SyslogTLSEventDestinationProtocol shall indicate the destination follows the
	// TLS-based transport for syslog as defined in RFC5424.
	SyslogTLSEventDestinationProtocol EventDestinationProtocol = "SyslogTLS"
	// SyslogTCPEventDestinationProtocol shall indicate the destination follows the
	// TCP-based transport for syslog as defined in RFC6587.
	SyslogTCPEventDestinationProtocol EventDestinationProtocol = "SyslogTCP"
	// SyslogUDPEventDestinationProtocol shall indicate the destination follows the
	// UDP-based transport for syslog as defined in RFC5424.
	SyslogUDPEventDestinationProtocol EventDestinationProtocol = "SyslogUDP"
	// SyslogRELPEventDestinationProtocol shall indicate the destination follows
	// the Reliable Event Logging Protocol (RELP) transport for syslog as defined
	// by www.rsyslog.com.
	SyslogRELPEventDestinationProtocol EventDestinationProtocol = "SyslogRELP"
	// OEMEventDestinationProtocol shall indicate an OEM-specific protocol. The
	// 'OEMProtocol' property shall contain the specific OEM event destination
	// protocol.
	OEMEventDestinationProtocol EventDestinationProtocol = "OEM"
)

type EventFormatType string

const (
	// EventEventFormatType shall receive an event payload as defined by the value
	// of the 'Protocol' property.
	EventEventFormatType EventFormatType = "Event"
	// MetricReportEventFormatType shall receive a metric report payload as defined
	// by the value of the 'Protocol' property.
	MetricReportEventFormatType EventFormatType = "MetricReport"
)

type SNMPAuthenticationProtocols string

const (
	// NoneSNMPAuthenticationProtocols shall indicate authentication is not
	// required.
	NoneSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "None"
	// CommunityStringSNMPAuthenticationProtocols shall indicate authentication
	// using SNMP community strings and the value of TrapCommunity.
	CommunityStringSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "CommunityString"
	// HMACMD5SNMPAuthenticationProtocols shall indicate authentication conforms to
	// the RFC3414-defined HMAC-MD5-96 authentication protocol.
	HMACMD5SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_MD5"
	// HMACSHA96SNMPAuthenticationProtocols shall indicate authentication conforms
	// to the RFC3414-defined HMAC-SHA-96 authentication protocol.
	HMACSHA96SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_SHA96"
	// HMAC128SHA224SNMPAuthenticationProtocols shall indicate authentication for
	// SNMPv3 access conforms to the RFC7860-defined usmHMAC128SHA224AuthProtocol.
	HMAC128SHA224SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC128_SHA224"
	// HMAC192SHA256SNMPAuthenticationProtocols shall indicate authentication for
	// SNMPv3 access conforms to the RFC7860-defined usmHMAC192SHA256AuthProtocol.
	HMAC192SHA256SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC192_SHA256"
	// HMAC256SHA384SNMPAuthenticationProtocols shall indicate authentication for
	// SNMPv3 access conforms to the RFC7860-defined usmHMAC256SHA384AuthProtocol.
	HMAC256SHA384SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC256_SHA384"
	// HMAC384SHA512SNMPAuthenticationProtocols shall indicate authentication for
	// SNMPv3 access conforms to the RFC7860-defined usmHMAC384SHA512AuthProtocol.
	HMAC384SHA512SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC384_SHA512"
	// AccountSNMPAuthenticationProtocols shall indicate authentication for SNMPv3
	// access is determined based on the corresponding account settings.
	AccountSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "Account" // From managernetworkprotocol
)

type SNMPEncryptionProtocols string

const (
	// NoneSNMPEncryptionProtocols shall indicate there is no encryption.
	NoneSNMPEncryptionProtocols SNMPEncryptionProtocols = "None"
	// CBCDESSNMPEncryptionProtocols shall indicate encryption conforms to the
	// RFC3414-defined CBC-DES encryption protocol.
	CBCDESSNMPEncryptionProtocols SNMPEncryptionProtocols = "CBC_DES"
	// CFB128AES128SNMPEncryptionProtocols shall indicate encryption conforms to
	// the RFC3826-defined CFB128-AES-128 encryption protocol.
	CFB128AES128SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES128"
	// CFB128AES192SNMPEncryptionProtocols shall indicate encryption conforms to
	// the CFB128-AES-192 encryption protocol, extended from RFC3826.
	CFB128AES192SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES192"
	// CFB128AES256SNMPEncryptionProtocols shall indicate encryption conforms to
	// the CFB128-AES-256 encryption protocol, extended from RFC3826.
	CFB128AES256SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES256"
	// AccountSNMPEncryptionProtocols shall indicate encryption is determined based
	// on the corresponding account settings.
	AccountSNMPEncryptionProtocols SNMPEncryptionProtocols = "Account" // From managernetworkprotocol
)

type SubscriptionType string

const (
	// RedfishEventSubscriptionType The subscription follows the Redfish
	// Specification for event notifications. To send an event notification, a
	// service sends an HTTP 'POST' to the subscriber's destination URI.
	RedfishEventSubscriptionType SubscriptionType = "RedfishEvent"
	// SSESubscriptionType The subscription follows the HTML5 server-sent event
	// definition for event notifications.
	SSESubscriptionType SubscriptionType = "SSE"
	// SNMPTrapSubscriptionType shall indicate the subscription follows the various
	// versions of SNMP Traps for event notifications. 'Protocol' shall specify the
	// appropriate version of SNMP.
	SNMPTrapSubscriptionType SubscriptionType = "SNMPTrap"
	// SNMPInformSubscriptionType shall indicate the subscription follows versions
	// 2 and 3 of SNMP Inform for event notifications. 'Protocol' shall specify the
	// appropriate version of SNMP.
	SNMPInformSubscriptionType SubscriptionType = "SNMPInform"
	// SyslogSubscriptionType shall indicate the subscription forwards syslog
	// messages to the event destination. 'Protocol' shall specify the appropriate
	// syslog protocol.
	SyslogSubscriptionType SubscriptionType = "Syslog"
	// OEMSubscriptionType shall indicate an OEM subscription type. The
	// 'OEMSubscriptionType' property shall contain the specific OEM subscription
	// type.
	OEMSubscriptionType SubscriptionType = "OEM"
)

// EventDestination shall represent the target of an event subscription,
// including the event types and context to provide to the target in the event
// payload.
type EventDestination struct {
	Entity
	// BackupDestinations shall contain an array of URIs to destination where
	// events are sent if the event receiver specified by 'Destination' is
	// unreachable or returns an error. Events are sent to each of the backup
	// destinations, in array order, until a destination has been reached. An empty
	// array shall indicate that the service supports backup event receivers, but
	// none have been specified by the user.
	//
	// Version added: v1.15.0
	BackupDestinations []string
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represent the server certificates for the
	// server referenced by the 'Destination' property. If 'VerifyCertificate' is
	// 'true', services shall compare the certificates in this collection with the
	// certificate obtained during handshaking with the event destination in order
	// to verify the identity of the event destination prior to sending an event.
	// If the server cannot be verified, the service shall not send the event. If
	// 'VerifyCertificate' is 'false', the service shall not perform certificate
	// verification with certificates in this collection. Regardless of the
	// contents of this collection, services may perform additional verification
	// based on other factors, such as the configuration of the SecurityPolicy
	// resource.
	//
	// Version added: v1.9.0
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates
	// that are provided to the server referenced by the 'Destination' property as
	// part of TLS handshaking.
	//
	// Version added: v1.11.0
	clientCertificates string
	// Context shall contain a client-supplied context that remains with the
	// connection through the connection's lifetime.
	Context string
	// DeliveryRetryPolicy shall indicate the subscription delivery retry policy
	// for events where the subscription type is 'RedfishEvent'.
	//
	// Version added: v1.6.0
	DeliveryRetryPolicy DeliveryRetryPolicy
	// Destination shall contain a URI to the destination where the events are
	// sent. If 'Protocol' is 'SMTP', the URI shall follow the RFC6068-described
	// format. SNMP URIs shall be consistent with RFC4088. Specifically, for
	// SNMPv3, if a username is specified in the SNMP URI, the SNMPv3
	// authentication and encryption configuration associated with that user shall
	// be utilized in the SNMPv3 traps. Syslog URIs shall be consistent with
	// RFC3986 and contain the scheme 'syslog://'. Server-sent event destinations
	// shall be in the form 'redfish-sse://ip:port' where 'ip' and 'port' are the
	// IP address and the port of the client with the open SSE connection. For
	// other URIs, such as HTTP or HTTPS, they shall be consistent with RFC3986.
	Destination string
	// EventFormatType shall indicate the content types of the message that this
	// service sends to the 'EventDestination'. If this property is not present,
	// the 'EventFormatType' shall be assumed to be Event.
	//
	// Version added: v1.4.0
	EventFormatType EventFormatType
	// EventTypes shall contain an array that contains the types of events that
	// shall be sent to the destination. To specify that a client is subscribing
	// for metric reports, the 'EventTypes' property should include 'MetricReport'.
	// If the subscription does not include this property, the service shall use a
	// single element with a default of 'Other'.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated. Starting with Redfish Specification v1.6
	// (Event v1.3), subscriptions are based on the 'RegistryPrefix' and
	// 'ResourceType' properties and not on the 'EventType' property. Use
	// 'EventFormatType' to create subscriptions for metric reports. If the
	// subscription does not include this property, the service shall use a single
	// element with a default of 'Other'.
	EventTypes []EventType
	// ExcludeMessageIDs shall contain an array of excluded 'MessageId' values that
	// are not allowed values for the 'MessageId' property within an event sent to
	// the subscriber. The 'MessageId' shall be in the
	// 'MessageRegistryPrefix.MessageKey' format. If included, the 'MessageId'
	// major and minor version details should be ignored. Events with a 'MessageId'
	// that is contained in this array shall not be sent to the subscriber. If this
	// property is an empty array or is absent, no exclusive filtering based upon
	// the 'MessageId' of an event is performed.
	//
	// Version added: v1.12.0
	ExcludeMessageIDs []string `json:"ExcludeMessageIds"`
	// ExcludeRegistryPrefixes shall contain an array of prefixes of excluded
	// message registries that contain the 'MessageId' values that are not allowed
	// values for the 'MessageId' property within an event sent to the subscriber.
	// Events with a 'MessageId' that is from a message registry contained in this
	// array shall not be sent to the subscriber. If this property is an empty
	// array or is absent, no exclusive filtering based upon message registry of
	// the 'MessageId' of an event is performed.
	//
	// Version added: v1.12.0
	ExcludeRegistryPrefixes []string
	// HTTPHeaders shall contain an array of objects consisting of the names and
	// values of the HTTP headers to include with every event 'POST' to the event
	// destination. This object shall be 'null' or an empty array in responses. An
	// empty array is the preferred return value in responses.
	HTTPHeaders []HTTPHeaderProperty `json:"HttpHeaders"`
	// HeartbeatIntervalMinutes shall indicate the interval for sending periodic
	// heartbeat events to the subscriber. The value shall be the interval, in
	// minutes, between each periodic event. This property shall not be present if
	// the 'SendHeartbeat' property is not present.
	//
	// Version added: v1.11.0
	HeartbeatIntervalMinutes *uint `json:",omitempty"`
	// IncludeOriginOfCondition shall indicate whether the event payload sent to
	// the subscription destination will expand the 'OriginOfCondition' property to
	// include the resource or object referenced by the 'OriginOfCondition'
	// property.
	//
	// Version added: v1.8.0
	IncludeOriginOfCondition bool
	// MessageIDs shall contain an array of 'MessageId' values that are the
	// allowable values for the 'MessageId' property within an event sent to the
	// subscriber. The 'MessageId' should be in the
	// 'MessageRegistryPrefix.MessageKey' format. If included, the 'MessageId'
	// major and minor version details should be ignored. Events with a 'MessageId'
	// that is not contained in this array and is not from a message registry
	// contained in RegistryPrefixes shall not be sent to the subscriber. If this
	// property is an empty array or is absent, no inclusive filtering based upon
	// the 'MessageId' of an event is performed.
	//
	// Version added: v1.1.0
	MessageIDs []string `json:"MessageIds"`
	// MetricReportDefinitions shall specify an array of metric report definitions
	// that are the only allowable generators of metric reports for this
	// subscription. Metric reports originating from metric report definitions not
	// contained in this array shall not be sent to the subscriber. If this
	// property is absent or the array is empty, the service shall send metric
	// reports originating from any metric report definition to the subscriber.
	//
	// Version added: v1.6.0
	MetricReportDefinitions []MetricReportDefinition
	// MetricReportDefinitionsCount
	MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMProtocol shall contain the protocol type that the event uses to send the
	// event to the destination. This property shall be present if 'Protocol' is
	// 'OEM'.
	//
	// Version added: v1.9.0
	OEMProtocol string
	// OEMSubscriptionType shall indicate the OEM-defined type of subscription for
	// events. This property shall be present if 'SubscriptionType' is 'OEM'.
	//
	// Version added: v1.9.0
	OEMSubscriptionType string
	// OriginResources shall specify an array of resources, resource collections,
	// or referenceable members that are the only allowable values for the
	// 'OriginOfCondition' property within an event that the service sends to the
	// subscriber. Events with an 'OriginOfCondition' that is not contained in this
	// array, and is not subordinate to members of this array if
	// 'SubordinateResources' contains the value 'true', shall not be sent to the
	// subscriber. If this property is an empty array or is absent, no filtering
	// based upon the URI of the 'OriginOfCondition' of an event is performed.
	//
	// Version added: v1.1.0
	OriginResources []string
	// OriginResourcesCount
	OriginResourcesCount int `json:"OriginResources@odata.count"`
	// Protocol shall contain the protocol type that the event uses to send the
	// event to the destination. A 'Redfish' value shall indicate that the event
	// type shall adhere to the type defined in the Redfish Specification.
	Protocol EventDestinationProtocol
	// RegistryPrefixes shall contain an array the prefixes of message registries
	// that contain the 'MessageId' values that are the allowable values for the
	// 'MessageId' property within an event sent to the subscriber. Events with a
	// 'MessageId' that is not from a message registry contained in this array and
	// is not contained by 'MessageIds' shall not be sent to the subscriber. If
	// this property is an empty array or is absent, no inclusive filtering based
	// upon message registry of the 'MessageId' of an event is performed.
	//
	// Version added: v1.4.0
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of resource type values that contain
	// the allowable resource types for the resource referenced by the
	// 'OriginOfCondition' property. Events with the resource type of the resource
	// referenced by the 'OriginOfCondition' property that is not contained in this
	// array shall not be sent to the subscriber. If this property is an empty
	// array or is absent, no filtering based upon the resource type of the
	// 'OriginOfCondition' of an event is performed. This property shall contain
	// only the general namespace for the type and not the versioned value. For
	// example, it shall not contain 'Task.v1_2_0.Task' and instead shall contain
	// 'Task'. To specify that a client is subscribing to metric reports, the
	// 'EventTypes' property should include 'MetricReport'.
	//
	// Version added: v1.4.0
	ResourceTypes []string
	// SNMP shall contain the settings for an SNMP event destination.
	//
	// Version added: v1.7.0
	SNMP SNMPSettings
	// SendHeartbeat shall indicate that the service shall periodically send the
	// 'RedfishServiceFunctional' message defined in the Heartbeat Event Message
	// Registry to the subscriber. If this property is not present, no periodic
	// event shall be sent. This property shall not apply to event destinations if
	// the 'SubscriptionType' property contains the value 'SSE'.
	//
	// Version added: v1.11.0
	SendHeartbeat bool
	// Severities shall contain an array of severities that are the allowable
	// values for the 'MessageSeverity' property within an event sent to the
	// subscriber. If this property is an empty array or is absent, no filtering
	// based upon the 'MessageSeverity' of an event is performed.
	//
	// Version added: v1.13.0
	Severities []Health
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.6.0
	Status Status
	// SubordinateResources shall indicate whether the subscription is for events
	// in the 'OriginResources' array and its subordinate resources. If 'true' and
	// the 'OriginResources' array is specified, the subscription is for events in
	// the 'OriginResources' array and its subordinate resources. Note that
	// resources associated through the Links section are not considered
	// subordinate. If 'false' and the 'OriginResources' array is specified, the
	// subscription shall be for events in the 'OriginResources' array only. If the
	// 'OriginResources' array is not present, this property shall have no
	// relevance.
	//
	// Version added: v1.4.0
	SubordinateResources bool
	// SubscriptionType shall indicate the type of subscription for events. If this
	// property is not present, the 'SubscriptionType' shall be assumed to be
	// 'RedfishEvent'.
	//
	// Version added: v1.3.0
	SubscriptionType SubscriptionType
	// SyslogFilters shall describe all desired syslog messages to send to a remote
	// syslog server. If this property contains an empty array or is absent, all
	// messages shall be sent.
	//
	// Version added: v1.9.0
	SyslogFilters []SyslogFilter
	// VerifyCertificate shall indicate whether the service will verify the
	// certificate of the server referenced by the 'Destination' property prior to
	// sending the event with the certificates found in the collection referenced
	// by the 'Certificates' property. If this property is not supported by the
	// service or specified by the client in the create request, it shall be
	// assumed to be 'false'. Regardless of the value of this property, services
	// may perform additional verification based on other factors, such as the
	// configuration of the SecurityPolicy resource.
	//
	// Version added: v1.9.0
	VerifyCertificate bool
	// resumeSubscriptionTarget is the URL to send ResumeSubscription requests.
	resumeSubscriptionTarget string
	// suspendSubscriptionTarget is the URL to send SuspendSubscription requests.
	suspendSubscriptionTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a EventDestination object from the raw JSON.
func (e *EventDestination) UnmarshalJSON(b []byte) error {
	type temp EventDestination
	type eActions struct {
		ResumeSubscription  ActionTarget `json:"#EventDestination.ResumeSubscription"`
		SuspendSubscription ActionTarget `json:"#EventDestination.SuspendSubscription"`
	}
	var tmp struct {
		temp
		Actions            eActions
		Certificates       Link `json:"Certificates"`
		ClientCertificates Link `json:"ClientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		// Workaround for the OriginResources actually being an idref
		var x struct {
			temp
			OriginResources    Links
			Certificates       Link
			ClientCertificates Link
		}

		err2 := json.Unmarshal(b, &x)
		if err2 != nil {
			// That didn't work either, just return the original error
			return err
		}

		tmp.temp = x.temp
		tmp.Certificates = x.Certificates
		tmp.ClientCertificates = x.ClientCertificates
		tmp.OriginResources = x.OriginResources.ToStrings()
	}

	*e = EventDestination(tmp.temp)

	// Extract the links to other entities for later
	e.resumeSubscriptionTarget = tmp.Actions.ResumeSubscription.Target
	e.suspendSubscriptionTarget = tmp.Actions.SuspendSubscription.Target
	e.certificates = tmp.Certificates.String()
	e.clientCertificates = tmp.ClientCertificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	e.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (e *EventDestination) Update() error {
	readWriteFields := []string{
		"BackupDestinations",
		"Context",
		"DeliveryRetryPolicy",
		"VerifyCertificate",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}

// GetEventDestination will get a EventDestination instance from the service.
func GetEventDestination(c Client, uri string) (*EventDestination, error) {
	return GetObject[EventDestination](c, uri)
}

// ListReferencedEventDestinations gets the collection of EventDestination from
// a provided reference.
func ListReferencedEventDestinations(c Client, link string) ([]*EventDestination, error) {
	return GetCollectionObjects[EventDestination](c, link)
}

// This action shall resume a suspended event subscription, which affects the
// subscription status. The service may deliver buffered events when the
// subscription is resumed.
// deliverBufferedEventDuration - This parameter shall indicate the event age
// of any buffered or otherwise undelivered events that shall be delivered to
// this event destination when the subscription is resumed. The service shall
// deliver any available, previously undelivered event that was created within
// the duration specified. A value that equates to zero time, such as 'PT0S',
// shall indicate that no previously undelivered events shall be sent. If
// undelivered events within the duration may have been discarded due to a lack
// of buffer space, the service should send the 'EventBufferExceeded' message
// from the Base Message Registry. If the client does not provide this
// parameter, the service shall apply an implementation-specific duration.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EventDestination) ResumeSubscription(deliverBufferedEventDuration string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["DeliverBufferedEventDuration"] = deliverBufferedEventDuration
	resp, taskInfo, err := PostWithTask(e.client,
		e.resumeSubscriptionTarget, payload, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall suspend an event subscription. No events shall be sent to
// the event destination until invocation of the 'ResumeSubscription' action.
// The value of the 'State' property within 'Status' shall contain 'Disabled'
// for a suspended subscription. The service may buffer events while the
// subscription is suspended.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EventDestination) SuspendSubscription() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(e.client,
		e.suspendSubscriptionTarget, payload, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Certificates gets the Certificates collection.
func (e *EventDestination) Certificates() ([]*Certificate, error) {
	if e.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](e.client, e.certificates)
}

// ClientCertificates gets the ClientCertificates collection.
func (e *EventDestination) ClientCertificates() ([]*Certificate, error) {
	if e.clientCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](e.client, e.clientCertificates)
}

// HTTPHeaderProperty shall contain the HTTP header name and value to include
// with every event 'POST' to the event destination.
type HTTPHeaderProperty map[string][]string

// SNMPSettings shall contain the settings for an SNMP event destination.
type SNMPSettings struct {
	// AuthenticationKey shall contain the key for SNMPv3 authentication. The value
	// shall be 'null' in responses. This property accepts a passphrase or a
	// hex-encoded key. If the string starts with 'Passphrase:', the remainder of
	// the string shall be the passphrase and shall be converted to the key as
	// described in the 'Password to Key Algorithm' section of RFC3414. If the
	// string starts with 'Hex:', then the remainder of the string shall be the key
	// encoded in hexadecimal notation. If the string starts with neither, the full
	// string shall be a passphrase and shall be converted to the key as described
	// in the 'Password to Key Algorithm' section of RFC3414.
	//
	// Version added: v1.7.0
	AuthenticationKey string
	// AuthenticationKeySet shall contain 'true' if a valid value was provided for
	// the 'AuthenticationKey' property. Otherwise, the property shall contain
	// 'false'.
	//
	// Version added: v1.10.0
	AuthenticationKeySet bool
	// AuthenticationProtocol shall contain the SNMPv3 authentication protocol.
	//
	// Version added: v1.7.0
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey shall contain the key for SNMPv3 encryption. The value shall
	// be 'null' in responses. This property accepts a passphrase or a hex-encoded
	// key. If the string starts with 'Passphrase:', the remainder of the string
	// shall be the passphrase and shall be converted to the key as described in
	// the 'Password to Key Algorithm' section of RFC3414. If the string starts
	// with 'Hex:', then the remainder of the string shall be the key encoded in
	// hexadecimal notation. If the string starts with neither, the full string
	// shall be a passphrase and shall be converted to the key as described in the
	// 'Password to Key Algorithm' section of RFC3414.
	//
	// Version added: v1.7.0
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the
	// 'EncryptionKey' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.10.0
	EncryptionKeySet bool
	// EncryptionProtocol shall contain the SNMPv3 encryption protocol.
	//
	// Version added: v1.7.0
	EncryptionProtocol SNMPEncryptionProtocols
	// HideCommunityStrings shall indicate whether the SNMP trap community string
	// is hidden in responses. If 'true', the 'TrapCommunity' property shall be
	// 'null' in responses. If 'false', the 'TrapCommunity' property shall contain
	// the trap community string in responses.
	//
	// Version added: v1.16.0
	HideCommunityStrings bool
	// TrapCommunity shall contain the SNMP trap community string. If
	// 'HideCommunityStrings' contains 'true' or is not supported, the value shall
	// be 'null' in responses. Otherwise the trap community string shall be
	// returned.
	//
	// Version added: v1.7.0
	TrapCommunity string
}

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
