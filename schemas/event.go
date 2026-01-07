//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #Event.v1_13_0.Event

package schemas

import (
	"encoding/json"
)

type DiagnosticDataTypes string

const (
	// ManagerDiagnosticDataTypes Manager diagnostic data.
	ManagerDiagnosticDataTypes DiagnosticDataTypes = "Manager"
	// PreOSDiagnosticDataTypes Pre-OS diagnostic data.
	PreOSDiagnosticDataTypes DiagnosticDataTypes = "PreOS"
	// OSDiagnosticDataTypes Operating system (OS) diagnostic data.
	OSDiagnosticDataTypes DiagnosticDataTypes = "OS"
	// OEMDiagnosticDataTypes OEM diagnostic data.
	OEMDiagnosticDataTypes DiagnosticDataTypes = "OEM"
	// CPERDiagnosticDataTypes shall indicate the data provided at the URI
	// specified by the 'AdditionalDataURI' property is a complete UEFI
	// Specification-defined Common Platform Error Record. The CPER data shall
	// contain a Record Header and at least one Section as defined by the UEFI
	// Specification.
	CPERDiagnosticDataTypes DiagnosticDataTypes = "CPER"
	// CPERSectionDiagnosticDataTypes shall indicate the data provided at the URI
	// specified by the 'AdditionalDataURI' property is a single Section of a UEFI
	// Specification-defined Common Platform Error Record. The CPER data shall
	// contain one Section as defined by the UEFI Specification, with no Record
	// Header.
	CPERSectionDiagnosticDataTypes DiagnosticDataTypes = "CPERSection"
)

type EventType string

const (
	// StatusChangeEventType The status of a resource has changed.
	StatusChangeEventType EventType = "StatusChange"
	// ResourceUpdatedEventType is a resource has been updated.
	ResourceUpdatedEventType EventType = "ResourceUpdated"
	// ResourceAddedEventType is a resource has been added.
	ResourceAddedEventType EventType = "ResourceAdded"
	// ResourceRemovedEventType is a resource has been removed.
	ResourceRemovedEventType EventType = "ResourceRemoved"
	// AlertEventType is a condition requires attention.
	AlertEventType EventType = "Alert"
	// MetricReportEventType shall be sent to a client in accordance with the
	// 'MetricReport' schema definition.
	MetricReportEventType EventType = "MetricReport"
	// OtherEventType shall be sent to a client in accordance with subscriptions to
	// 'RegistryPrefixes' or 'ResourceTypes'.
	OtherEventType EventType = "Other"
)

// IsValidEventType will check if it is a valid EventType.
// Should remove and leave it to the service to decide if it's valid, but since
// this is deprecated leaving it in for now.
func (e EventType) IsValidEventType() bool {
	switch e {
	case AlertEventType, ResourceAddedEventType,
		ResourceRemovedEventType, ResourceUpdatedEventType,
		StatusChangeEventType, MetricReportEventType, OtherEventType:
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

// Event This resource contains an event for a Redfish implementation.
type Event struct {
	Entity
	// Context shall contain a client supplied context for the event destination to
	// which this event is being sent.
	//
	// Version added: v1.1.0
	Context string
	// Events shall contain an array of objects that represent the occurrence of
	// one or more events.
	Events []EventRecord
	// EventsCount
	EventsCount int `json:"Events@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetEvent will get a Event instance from the service.
func GetEvent(c Client, uri string) (*Event, error) {
	return GetObject[Event](c, uri)
}

// ListReferencedEvents gets the collection of Event from
// a provided reference.
func ListReferencedEvents(c Client, link string) ([]*Event, error) {
	return GetCollectionObjects[Event](c, link)
}

// CPER shall contain the details for a CPER section or record that is the
// source of an event.
type CPER struct {
	// NotificationType shall contain the CPER Notification Type for a CPER record
	// that corresponds to the contents of the 'DiagnosticData' property or data
	// retrieved from the URI specified by the 'AdditionalDataURI' property. This
	// property shall only be present if 'DiagnosticDataType' contains 'CPER'.
	//
	// Version added: v1.8.0
	NotificationType string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.8.0
	OEM json.RawMessage `json:"Oem"`
	// SectionType shall contain the CPER Section Type for a CPER section that
	// corresponds to the contents of the 'DiagnosticData' property or data
	// retrieved from the URI specified by the 'AdditionalDataURI' property. This
	// property shall only be present if 'DiagnosticDataType' contains
	// 'CPERSection'.
	//
	// Version added: v1.8.0
	SectionType string
}

// EventRecord represents the EventRecord type.
type EventRecord struct {
	// AdditionalDataSizeBytes shall contain the size of the additional data
	// retrieved from the URI specified by the 'AdditionalDataURI' property for
	// this event.
	//
	// Version added: v1.8.0
	AdditionalDataSizeBytes *int `json:",omitempty"`
	// AdditionalDataURI shall contain the URI at which to access the additional
	// data for the event, using the Redfish protocol and authentication methods.
	// If both 'DiagnosticData' and 'AdditionalDataURI' are present,
	// 'DiagnosticData' shall contain a Base64-encoded string, with padding
	// characters, of the data retrieved from the URI specified by the
	// 'AdditionalDataURI' property.
	//
	// Version added: v1.8.0
	AdditionalDataURI string
	// CPER shall contain the details for a CPER section or record that is the
	// source of this event.
	//
	// Version added: v1.8.0
	CPER CPER
	// Context shall contain a client supplied context for the event destination to
	// which this event is being sent.
	//
	// Deprecated: v1.1.0
	// Events are triggered independently from subscriptions to those events. This
	// property has been deprecated in favor of the 'Context' property found at the
	// root level of the object.
	Context string
	// DiagnosticData shall contain a Base64-encoded string, with padding
	// characters, that represents the diagnostic data associated with this event.
	// The contents shall depend on the value of the 'DiagnosticDataType' property.
	// The length of the value should not exceed 4 KB. Larger diagnostic data
	// payloads should omit this property and use the 'AdditionalDataURI' property
	// to reference the data. If both 'DiagnosticData' and 'AdditionalDataURI' are
	// present, 'DiagnosticData' shall contain the Base64-encoding of the data
	// retrieved from the URI specified by the 'AdditionalDataURI' property.
	//
	// Version added: v1.8.0
	DiagnosticData string
	// DiagnosticDataType shall contain the type of data available in the
	// 'DiagnosticData' property or retrieved from the URI specified by the
	// 'AdditionalDataURI' property.
	//
	// Version added: v1.8.0
	DiagnosticDataType DiagnosticDataTypes
	// EventGroupID shall indicate that events are related and shall have the same
	// value when multiple event messages are produced by the same root cause.
	// Implementations shall use separate values for events with a separate root
	// cause. This property value shall not imply an ordering of events. The '0'
	// value shall indicate that this event is not grouped with any other event.
	//
	// Version added: v1.3.0
	EventGroupID int `json:"EventGroupId"`
	// EventID shall contain a service-defined unique identifier for the event.
	EventID string `json:"EventId"`
	// EventTimestamp shall indicate the time the event occurred where the value
	// shall be consistent with the Redfish service time that is also used for the
	// values of the 'Modified' property.
	EventTimestamp string
	// EventType shall indicate the type of event.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated. Starting with Redfish Specification v1.6
	// (Event v1.3), subscriptions are based on the 'RegistryPrefix' and
	// 'ResourceType' properties and not on the 'EventType' property.
	EventType EventType
	// LogEntry shall contain a link to a resource of type 'LogEntry' that
	// represents the log entry created for this event.
	//
	// Version added: v1.7.0
	logEntry string
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Message shall contain a human-readable event message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted
	// for the arguments in the message when looked up in the message registry. It
	// has the same semantics as the 'MessageArgs' property in the Redfish
	// 'MessageRegistry' schema. If the corresponding 'ParamType' value contains
	// 'number', the service shall convert the number to a string representation of
	// the number.
	MessageArgs []string
	// MessageID shall contain a 'MessageId', as defined in the 'MessageId format'
	// clause of the Redfish Specification.
	MessageID string `json:"MessageId"`
	// MessageSeverity shall contain the severity of the message in this event.
	// Services can replace the value defined in the message registry with a value
	// more applicable to the implementation.
	//
	// Version added: v1.5.0
	MessageSeverity Health
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMDiagnosticDataType shall contain the OEM-defined type of data available
	// in the 'DiagnosticData' property or retrieved from the URI specified by the
	// 'AdditionalDataURI' property. This property shall be present if
	// 'DiagnosticDataType' is 'OEM'.
	//
	// Version added: v1.9.0
	OEMDiagnosticDataType string
	// OriginAddress shall contain the IP address, with scheme, of the user
	// associated with the log entry. This should be used for audit logs that
	// result from a user action.
	//
	// Version added: v1.13.0
	OriginAddress string
	// OriginOfCondition shall contain a link to the resource or object that
	// originated the condition that caused the event to be generated. If the event
	// subscription has the 'IncludeOriginOfCondition' property set to 'true', it
	// shall include the entire resource or object referenced by the link. For
	// events that represent the creation or deletion of a resource, this property
	// should reference the created or deleted resource and not the collection that
	// contains the resource.
	originOfCondition string
	// OriginOfConditionUnavailable shall indicate whether the 'OriginOfCondition'
	// link is unavailable. If 'true', services shall not expand the
	// 'OriginOfCondition' link. If this property is not present, the value shall
	// be assumed to be 'false'.
	//
	// Version added: v1.12.0
	OriginOfConditionUnavailable bool
	// Resolution shall contain the resolution of the event. Services should
	// replace the resolution defined in the message registry with a more specific
	// resolution in the event.
	//
	// Version added: v1.9.0
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the
	// cause of the event. This property shall not be present if the
	// 'MessageSeverity' or 'Severity' properties contain 'OK'. A client can stop
	// executing the resolution steps once the 'Resolved' property in the
	// associated 'LogEntry' resource contains 'true' or the 'Health' property in
	// the associated resource referenced by the 'OriginOfCondition' property
	// contains 'OK'.
	//
	// Version added: v1.10.0
	ResolutionSteps []ResolutionStep
	// Severity shall contain the severity of the event, as defined in the 'Status'
	// clause of the Redfish Specification. Services can replace the value defined
	// in the message registry with a value more applicable to the implementation.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of 'MessageSeverity', which ties
	// the values to the enumerations defined for the 'Health' property within
	// 'Status'.
	Severity string
	// SpecificEventExistsInGroup shall indicate that the event is equivalent to
	// another event, with a more specific definition, within the same
	// 'EventGroupId'. For example, the 'DriveFailed' message from the Storage
	// Device Message Registry is more specific than the
	// 'ResourceStatusChangedCritical' message from the Resource Event Message
	// Registry, when both occur with the same 'EventGroupId'. This property shall
	// contain 'true' if a more specific event is available, and shall contain
	// 'false' if no equivalent event exists in the same 'EventGroupId'. If this
	// property is absent, the value shall be assumed to be 'false'.
	//
	// Version added: v1.6.0
	SpecificEventExistsInGroup bool
	// UserAuthenticationSource shall contain the URL to the authentication service
	// that is associated with the username property. This should be used for
	// events that result from a user action.
	//
	// Version added: v1.11.0
	UserAuthenticationSource string
	// Username shall contain the username of the account associated with the event
	// record. This should be used for events that result from a user action.
	//
	// Version added: v1.11.0
	Username string
}

// UnmarshalJSON unmarshals a EventRecord object from the raw JSON.
func (e *EventRecord) UnmarshalJSON(b []byte) error {
	type temp EventRecord
	var tmp struct {
		temp
		LogEntry          Link `json:"LogEntry"`
		OriginOfCondition Link `json:"OriginOfCondition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = EventRecord(tmp.temp)

	// Extract the links to other entities for later
	e.logEntry = tmp.LogEntry.String()
	e.originOfCondition = tmp.OriginOfCondition.String()

	return nil
}

// LogEntry gets the LogEntry linked resource.
func (e *EventRecord) LogEntry(client Client) (*LogEntry, error) {
	if e.logEntry == "" {
		return nil, nil
	}
	return GetObject[LogEntry](client, e.logEntry)
}

// OriginOfCondition gets the OriginOfCondition linked resource.
func (e *EventRecord) OriginOfCondition(client Client) (*Entity, error) {
	if e.originOfCondition == "" {
		return nil, nil
	}
	return GetObject[Entity](client, e.originOfCondition)
}
