//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	// CPERDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI property
	// is a complete UEFI Specification-defined Common Platform Error Record. The CPER data shall contain a Record
	// Header and at least one Section as defined by the UEFI Specification.
	CPERDiagnosticDataTypes DiagnosticDataTypes = "CPER"
	// CPERSectionDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI
	// property is a single Section of a UEFI Specification-defined Common Platform Error Record. The CPER data shall
	// contain one Section as defined by the UEFI Specification, with no Record Header.
	CPERSectionDiagnosticDataTypes DiagnosticDataTypes = "CPERSection"
)

// CPER shall contain the details for a CPER section or record that is the source of an event.
type CPER struct {
	// NotificationType shall contain the CPER Notification Type for a CPER record that corresponds to the contents of
	// the DiagnosticData property or data retrieved from the URI specified by the AdditionalDataURI property. This
	// property shall only be present if DiagnosticDataType contains 'CPER'.
	NotificationType string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SectionType shall contain the CPER Section Type for a CPER section that corresponds to the contents of the
	// DiagnosticData property or data retrieved from the URI specified by the AdditionalDataURI property. This
	// property shall only be present if DiagnosticDataType contains 'CPERSection'.
	SectionType string
}

// Event This resource contains an event for a Redfish implementation.
type Event struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Context shall contain a client supplied context for the event destination to which this event is being sent.
	Context string
	// Description provides a description of this resource.
	Description string
	// Events shall contain an array of objects that represent the occurrence of one or more events.
	Events []EventRecord
	// EventsCount is the number of Event records.
	EventsCount int `json:"Events@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetEvent will get a Event instance from the service.
func GetEvent(c common.Client, uri string) (*Event, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var event Event
	err = json.NewDecoder(resp.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	event.SetClient(c)
	return &event, nil
}

// ListReferencedEvents gets the collection of Event from
// a provided reference.
func ListReferencedEvents(c common.Client, link string) ([]*Event, error) {
	var result []*Event
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Event
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		event, err := GetEvent(c, link)
		ch <- GetResult{Item: event, Link: link, Error: err}
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

// EventRecord
type EventRecord struct {
	// AdditionalDataSizeBytes shall contain the size of the additional data retrieved from the URI specified by the
	// AdditionalDataURI property for this event.
	AdditionalDataSizeBytes int
	// AdditionalDataURI shall contain the URI at which to access the additional data for the event, using the Redfish
	// protocol and authentication methods. If both DiagnosticData and AdditionalDataURI are present, DiagnosticData
	// shall contain the Base64-encoding of the data retrieved from the URI specified by the AdditionalDataURI
	// property.
	AdditionalDataURI string
	// CPER shall contain the details for a CPER section or record that is the source of this event.
	CPER CPER
	// DiagnosticData shall contain a Base64-encoded string that represents diagnostic data associated with this event.
	// The contents shall depend on the value of the DiagnosticDataType property. The length of the value should not
	// exceed 4 KB. Larger diagnostic data payloads should omit this property and use the AdditionalDataURI property to
	// reference the data. If both DiagnosticData and AdditionalDataURI are present, DiagnosticData shall contain the
	// Base64-encoding of the data retrieved from the URI specified by the AdditionalDataURI property.
	DiagnosticData string
	// DiagnosticDataType shall contain the type of data available in the DiagnosticData property or retrieved from the
	// URI specified by the AdditionalDataURI property.
	DiagnosticDataType DiagnosticDataTypes
	// EventGroupID shall indicate that events are related and shall have the same value when multiple event messages
	// are produced by the same root cause. Implementations shall use separate values for events with a separate root
	// cause. This property value shall not imply an ordering of events. The '0' value shall indicate that this event
	// is not grouped with any other event.
	EventGroupID int
	// EventID shall contain a service-defined unique identifier for the event.
	EventID string
	// EventTimestamp shall indicate the time the event occurred where the value shall be consistent with the Redfish
	// service time that is also used for the values of the Modified property.
	EventTimestamp string
	// LogEntry shall contain a link to a resource of type LogEntry that represents the log entry created for this
	// event.
	logEntry string
	// MemberID shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberID string
	// Message shall contain a human-readable event message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted for the arguments in the message
	// when looked up in the message registry. It has the same semantics as the MessageArgs property in the Redfish
	// MessageRegistry schema. If the corresponding ParamType value contains 'number', the service shall convert the
	// number to a string representation of the number.
	MessageArgs []string
	// MessageID shall contain a MessageID, as defined in the 'MessageId format' clause of the Redfish Specification.
	MessageID string
	// MessageSeverity shall contain the severity of the message in this event. Services can replace the value defined
	// in the message registry with a value more applicable to the implementation.
	MessageSeverity common.Health
	// OEMDiagnosticDataType shall contain the OEM-defined type of data available in the DiagnosticData property or
	// retrieved from the URI specified by the AdditionalDataURI property. This property shall be present if
	// DiagnosticDataType is 'OEM'.
	OEMDiagnosticDataType string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginOfCondition shall contain a link to the resource or object that originated the condition that caused the
	// event to be generated. If the event subscription has the IncludeOriginOfCondition property set to 'true', it
	// shall include the entire resource or object referenced by the link. For events that represent the creation or
	// deletion of a resource, this property should reference the created or deleted resource and not the collection
	// that contains the resource.
	OriginOfCondition string
	// Resolution shall contain the resolution of the event. Services should replace the resolution defined in the
	// message registry with a more specific resolution in the event.
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the cause of the event. This property
	// shall not be present if the MessageSeverity or Severity properties contain 'OK'. A client can stop executing the
	// resolution steps once the Resolved property in the associated LogEntry resource contains 'true' or the Health
	// property in the associated resource referenced by the OriginOfCondition property contains 'OK'.
	ResolutionSteps []ResolutionStep
	// SpecificEventExistsInGroup shall indicate that the event is equivalent to another event, with a more specific
	// definition, within the same EventGroupId. For example, the 'DriveFailed' message from the Storage Device Message
	// Registry is more specific than the 'ResourceStatusChangedCritical' message from the Resource Event Message
	// Registry, when both occur with the same EventGroupId. This property shall contain 'true' if a more specific
	// event is available, and shall contain 'false' if no equivalent event exists in the same EventGroupId. If this
	// property is absent, the value shall be assumed to be 'false'.
	SpecificEventExistsInGroup bool
}

// UnmarshalJSON unmarshals a EventRecord object from the raw JSON.
func (eventrecord *EventRecord) UnmarshalJSON(b []byte) error {
	type temp EventRecord
	var t struct {
		temp
		LogEntry common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventrecord = EventRecord(t.temp)

	// Extract the links to other entities for later
	eventrecord.logEntry = t.LogEntry.String()

	return nil
}

// EventRecordActions shall contain the available actions for this resource.
type EventRecordActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}
