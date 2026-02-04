//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Resource.v1_23_0.json
// 2025.3 - #Resource.v1_23_0

package schemas

import (
	"encoding/json"
	"strconv"
)

type ConditionType string

const (
	// AlertConditionType shall indicate a condition that requires correction, such
	// as a fault.
	AlertConditionType ConditionType = "Alert"
	// InformationalConditionType shall indicate a condition that requires
	// attention, maintenance, or some other user intervention, such as performing
	// a reset to activate new firmware.
	InformationalConditionType ConditionType = "Informational"
	// SubsystemConditionType shall indicate a condition that contains the health
	// of a subsystem. If supported by the service, the service shall always
	// provide this condition in responses.
	SubsystemConditionType ConditionType = "Subsystem"
)

type DurableNameFormat string

const (
	// NAADurableNameFormat shall contain a hexadecimal representation of the Name
	// Address Authority structure, as defined in the T11 Fibre Channel - Framing
	// and Signaling - 3 (FC-FS-3) specification. The 'DurableName' property shall
	// follow the regular expression pattern '^(([0-9A-Fa-f]{2}){8}){1,2}$', where
	// the most significant octet is first.
	NAADurableNameFormat DurableNameFormat = "NAA"
	// iQNDurableNameFormat shall be in the iSCSI Qualified Name (iQN) format, as
	// defined in RFC3720 and RFC3721.
	IQNDurableNameFormat DurableNameFormat = "iQN"
	// FCWWNDurableNameFormat shall contain a hexadecimal representation of the
	// World-Wide Name (WWN) format, as defined in the T11 Fibre Channel Physical
	// and Signaling Interface Specification. The 'DurableName' property shall
	// follow the regular expression pattern
	// '^([0-9A-Fa-f]{2}[:-]){7}([0-9A-Fa-f]{2})$', where the most significant
	// octet is first.
	FCWWNDurableNameFormat DurableNameFormat = "FC_WWN"
	// UUIDDurableNameFormat shall contain the hexadecimal representation of the
	// UUID, as defined by RFC4122. The 'DurableName' property shall follow the
	// regular expression pattern
	// '([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})'.
	UUIDDurableNameFormat DurableNameFormat = "UUID"
	// EUIDurableNameFormat shall contain the hexadecimal representation of the
	// IEEE-defined 64-bit Extended Unique Identifier (EUI), as defined in the
	// IEEE's Guidelines for 64-bit Global Identifier (EUI-64) Specification. The
	// 'DurableName' property shall follow the regular expression pattern
	// '^([0-9A-Fa-f]{2}[:-]){7}([0-9A-Fa-f]{2})$', where the most significant
	// octet is first.
	EUIDurableNameFormat DurableNameFormat = "EUI"
	// NQNDurableNameFormat shall be in the NVMe Qualified Name (NQN) format, as
	// defined in the NVN Express over Fabric Specification.
	NQNDurableNameFormat DurableNameFormat = "NQN"
	// NSIDDurableNameFormat shall be in the NVM Namespace Identifier (NSID)
	// format, as defined in the NVN Express Specification.
	NSIDDurableNameFormat DurableNameFormat = "NSID"
	// NGUIDDurableNameFormat shall be in the Namespace Globally Unique Identifier
	// (NGUID), as defined in the NVN Express Specification. The 'DurableName'
	// property shall follow the regular expression pattern
	// '^([0-9A-Fa-f]{2}){16}$', where the most significant octet is first.
	NGUIDDurableNameFormat DurableNameFormat = "NGUID"
	// MACAddressDurableNameFormat shall be a media access control address (MAC
	// address), which is a unique identifier assigned to a network interface
	// controller (NIC) for use as a network address. This value should not be used
	// if a more specific type of identifier is available. The 'DurableName'
	// property shall follow the regular expression pattern
	// '^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$', where the most significant
	// octet is first.
	MACAddressDurableNameFormat DurableNameFormat = "MACAddress"
	// GCXLIDDurableNameFormat shall be in the globally unique CXL logical device
	// identifier (GCXLID). The 'DurableName' property shall follow the regular
	// expression pattern '^([0-9A-Fa-f]{2}-){7}[0-9A-Fa-f]{2}:([0-9A-Fa-f]{4})$',
	// where the first eight hyphen-delimited octets contain the PCIe serial
	// number, where the most significant octet is first, and the remaining 16-bit
	// field contains the CXL Logical Device Identifier, with the most significant
	// byte first.
	GCXLIDDurableNameFormat DurableNameFormat = "GCXLID"
)

type Health string

const (
	// OKHealth Normal.
	OKHealth Health = "OK"
	// WarningHealth is a condition requires attention.
	WarningHealth Health = "Warning"
	// CriticalHealth is a critical condition requires immediate attention.
	CriticalHealth Health = "Critical"
)

type IndicatorLED string

const (
	// UnknownIndicatorLED shall represent that the indicator LED is in an unknown
	// state. The service shall reject 'PATCH' or 'PUT' requests containing this
	// value by returning the HTTP '400 Bad Request' status code.
	UnknownIndicatorLED IndicatorLED = "Unknown"
	// LitIndicatorLED shall represent that the indicator LED is in a solid on
	// state. If the service does not support this value, it shall reject 'PATCH'
	// or 'PUT' requests containing this value by returning the HTTP '400 Bad
	// Request' status code.
	LitIndicatorLED IndicatorLED = "Lit"
	// BlinkingIndicatorLED shall represent that the indicator LED is in a blinking
	// state where the LED is being turned on and off in repetition. If the service
	// does not support this value, it shall reject 'PATCH' or 'PUT' requests
	// containing this value by returning the HTTP '400 Bad Request' status code.
	BlinkingIndicatorLED IndicatorLED = "Blinking"
	// OffIndicatorLED shall represent that the indicator LED is in a solid off
	// state. If the service does not support this value, it shall reject 'PATCH'
	// or 'PUT' requests containing this value by returning the HTTP '400 Bad
	// Request' status code.
	OffIndicatorLED IndicatorLED = "Off"
)

// LocationType is This enumeration shall list the types of locations for a part
// within an enclosure.
type LocationType string

const (
	// SlotLocationType shall indicate the part is located in a slot.
	SlotLocationType LocationType = "Slot"
	// BayLocationType shall indicate the part is located in a bay.
	BayLocationType LocationType = "Bay"
	// ConnectorLocationType shall indicate the part is located in a connector or
	// port.
	ConnectorLocationType LocationType = "Connector"
	// SocketLocationType shall indicate the part is located in a socket.
	SocketLocationType LocationType = "Socket"
	// BackplaneLocationType shall indicate the part is a backplane in an
	// enclosure.
	BackplaneLocationType LocationType = "Backplane"
	// EmbeddedLocationType shall indicate the part is embedded or otherwise
	// permanently incorporated into a larger part or device. This value shall not
	// be used for parts that can be removed by a user or are considered
	// field-replaceable.
	EmbeddedLocationType LocationType = "Embedded"
)

// Orientation is This enumeration shall list the orientations for the ordering
// of the 'LocationOrdinalValue' property.
type Orientation string

const (
	// FrontToBackOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is front to back.
	FrontToBackOrientation Orientation = "FrontToBack"
	// BackToFrontOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is back to front.
	BackToFrontOrientation Orientation = "BackToFront"
	// TopToBottomOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is top to bottom.
	TopToBottomOrientation Orientation = "TopToBottom"
	// BottomToTopOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is bottom to top.
	BottomToTopOrientation Orientation = "BottomToTop"
	// LeftToRightOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is left to right.
	LeftToRightOrientation Orientation = "LeftToRight"
	// RightToLeftOrientation shall indicate the ordering for
	// 'LocationOrdinalValue' is right to left.
	RightToLeftOrientation Orientation = "RightToLeft"
)

type PowerState string

const (
	// OnPowerState The resource is powered on.
	OnPowerState PowerState = "On"
	// OffPowerState The resource is powered off. The components within the
	// resource might continue to have AUX power.
	OffPowerState PowerState = "Off"
	// PoweringOnPowerState is a temporary state between off and on. The components
	// within the resource can take time to process the power on action.
	PoweringOnPowerState PowerState = "PoweringOn"
	// PoweringOffPowerState is a temporary state between on and off. The
	// components within the resource can take time to process the power off
	// action.
	PoweringOffPowerState PowerState = "PoweringOff"
	// PausedPowerState The resource is paused.
	PausedPowerState PowerState = "Paused"
)

// RackUnits is Enumeration literals shall name the type of rack unit in use.
type RackUnits string

const (
	// OpenURackUnits shall be specified in terms of the Open Compute Open Rack
	// Specification.
	OpenURackUnits RackUnits = "OpenU"
	// EIA310RackUnits shall conform to the EIA-310 standard.
	EIA310RackUnits RackUnits = "EIA_310"
)

// Reference is This enumeration shall list the reference areas for the location
// of the part within an enclosure.
type Reference string

const (
	// TopReference shall indicate the part is in the top of the unit.
	TopReference Reference = "Top"
	// BottomReference shall indicate the part is in the bottom of the unit.
	BottomReference Reference = "Bottom"
	// FrontReference shall indicate the part is in the front of the unit.
	FrontReference Reference = "Front"
	// RearReference shall indicate the part is in the rear of the unit.
	RearReference Reference = "Rear"
	// LeftReference shall indicate the part is on the left side of the unit.
	LeftReference Reference = "Left"
	// RightReference shall indicate the part is on the right side of the unit.
	RightReference Reference = "Right"
	// MiddleReference shall indicate the part is in the middle of the unit.
	MiddleReference Reference = "Middle"
)

type ResetType string

const (
	// OnResetType shall indicate the resource will transition to a power on state.
	// Upon successful completion, the 'PowerState' property, if supported, shall
	// contain the value 'On'.
	OnResetType ResetType = "On"
	// ForceOffResetType shall indicate the resource will transition to a power off
	// state. The transition will start immediately. Upon successful completion,
	// the 'PowerState' property, if supported, shall contain the value 'Off'.
	ForceOffResetType ResetType = "ForceOff"
	// GracefulShutdownResetType shall indicate the resource will transition to a
	// power off state. The transition will start after first performing tasks to
	// safely shut down the resource. For example, when shutting down a computer
	// system, the host operating system is allowed to safely shut down processes
	// and close connections. Upon successful completion, the 'PowerState'
	// property, if supported, shall contain the value 'Off'.
	GracefulShutdownResetType ResetType = "GracefulShutdown"
	// GracefulRestartResetType shall indicate the resource will transition to a
	// power on state, after transiting through a restart. The transition will
	// start after first performing tasks to safely shut down the resource. For
	// example, when shutting down a computer system, the host operating system is
	// allowed to safely shut down processes and close connections. Upon successful
	// completion, the 'PowerState' property, if supported, shall contain the value
	// 'On'.
	GracefulRestartResetType ResetType = "GracefulRestart"
	// ForceRestartResetType shall indicate the resource will transition to a power
	// on state, after transiting through a restart. The transition will start
	// immediately. Upon successful completion, the 'PowerState' property, if
	// supported, shall contain the value 'On'.
	ForceRestartResetType ResetType = "ForceRestart"
	// NmiResetType shall indicate the resource will generate a diagnostic
	// interrupt.
	NmiResetType ResetType = "Nmi"
	// ForceOnResetType shall indicate the resource will transition to a power on
	// state. The transition will start immediately. Upon successful completion,
	// the 'PowerState' property shall contain the value 'On'.
	ForceOnResetType ResetType = "ForceOn"
	// PushPowerButtonResetType shall indicate the resource will behave as if the
	// physical power button is pressed. The behavior of pressing the physical
	// power button may be dependent on the state of the unit and the behavior may
	// be configurable.
	PushPowerButtonResetType ResetType = "PushPowerButton"
	// PowerCycleResetType shall indicate the resource will perform a power cycle.
	// The transition will start immediately. This is different from
	// 'FullPowerCycle' in that it typically performs localized power sequencing of
	// the resource while external power is still present. For example, turning DC
	// voltage regulators off and then turning DC voltage regulators back on. If
	// currently in the power on state, the resource will transition to a power off
	// state, then transition to a power on state. If currently in the power off
	// state, the resource will transition to a power on state. Upon successful
	// completion, the 'PowerState' property, if supported, shall contain the value
	// 'On'.
	PowerCycleResetType ResetType = "PowerCycle"
	// SuspendResetType shall indicate the resource will have any state information
	// written to persistent memory and then transition to a power off state. Upon
	// successful completion, the 'PowerState' property, if supported, shall
	// contain the value 'Off'.
	SuspendResetType ResetType = "Suspend"
	// PauseResetType shall indicate the resource will transition to a paused
	// state. Upon successful completion, the 'PowerState' property, if supported,
	// shall contain the value 'Paused'.
	PauseResetType ResetType = "Pause"
	// ResumeResetType shall indicate the resource will transition to a power on
	// state. Upon successful completion, the 'PowerState' property, if supported,
	// shall contain the value 'On'.
	ResumeResetType ResetType = "Resume"
	// FullPowerCycleResetType shall indicate the resource will perform a full
	// power cycle as if utility lines to the resource are removed and restored.
	// The transition will start immediately. This is different from 'PowerCycle'
	// in that it's as close to a true power cycle as possible. For example,
	// removing AC power cables from an enclosure and then restoring the AC power
	// cables. If currently in the power on state, the resource will transition to
	// a power off state, then transition to a power on state. If currently in the
	// power off state, the resource will transition to a power on state. Upon
	// successful completion, the 'PowerState' property, if supported, shall
	// contain the value 'On'. If a service supports this value but there are no
	// other distinct power cycle flows, the service shall support 'PowerCycle' as
	// an alias to 'FullPowerCycle'. This type of reset may cause the manager
	// providing the Redfish service to power cycle. If the manager providing the
	// Redfish service is affected by this type of reset, the service shall send
	// the action response before resetting to prevent client timeouts.
	FullPowerCycleResetType ResetType = "FullPowerCycle"
)

type State string

const (
	// EnabledState shall indicate that a function or resource is capable of
	// operating.
	EnabledState State = "Enabled"
	// DisabledState shall indicate that a function or resource is unavailable.
	DisabledState State = "Disabled"
	// StandbyOfflineState shall indicate that a resource is temporarily
	// unavailable but may become available through an external action.
	StandbyOfflineState State = "StandbyOffline"
	// StandbySpareState shall indicate that a resource is unavailable but may
	// become available automatically as part of a failover, through an external
	// action, or in response to the change in state of another device or resource.
	StandbySpareState State = "StandbySpare"
	// InTestState shall indicate that the component is undergoing testing or is in
	// the process of capturing information for debugging.
	InTestState State = "InTest"
	// StartingState shall indicate that the resource is attempting to transition
	// to 'Enabled'.
	StartingState State = "Starting"
	// AbsentState shall indicate that the function or device is absent as defined
	// in the 'Absent resources' clause of the Redfish Specification.
	AbsentState State = "Absent"
	// UnavailableOfflineState shall indicate that a function or resource is
	// present but not able to be used.
	UnavailableOfflineState State = "UnavailableOffline"
	// DeferringState shall indicate the element does not process any commands but
	// queues new requests.
	DeferringState State = "Deferring"
	// QuiescedState shall indicate the element is enabled but only processes a
	// restricted set of commands.
	QuiescedState State = "Quiesced"
	// UpdatingState shall indicate the element is updating. The element may become
	// unavailable or operate at a degraded level of performance or functionality.
	UpdatingState State = "Updating"
	// QualifiedState shall indicate the element is within the acceptable range of
	// operation.
	QualifiedState State = "Qualified"
	// DegradedState shall indicate the resource is enabled but operating in a
	// degraded mode.
	DegradedState State = "Degraded"
)

// Condition shall contain the description and details of a condition that
// exists within this resource or a related resource that requires attention.
type Condition struct {
	// ConditionType shall contain the type of condition.
	//
	// Version added: v1.22.0
	ConditionType ConditionType
	// LogEntry shall contain a link to a resource of type 'LogEntry' that
	// represents the log entry created for this condition.
	logEntry string
	// Message shall contain a human-readable message describing this condition.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted
	// for the arguments in the message when looked up in the message registry. It
	// has the same semantics as the 'MessageArgs' property in the Redfish
	// 'MessageRegistry' schema.
	MessageArgs []string
	// MessageID shall contain a 'MessageId', as defined in the 'MessageId format'
	// clause of the Redfish Specification.
	MessageID string `json:"MessageId"`
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.21.0
	OEM json.RawMessage `json:"Oem"`
	// OriginOfCondition shall contain a link to the resource or object that
	// originated the condition. This property shall not be present if the
	// condition was caused by this resource.
	originOfCondition string
	// Resolution shall contain the resolution of the condition. Services should
	// replace the resolution defined in the message registry with a more specific
	// resolution.
	//
	// Version added: v1.14.0
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the
	// condition. A client can stop executing the resolution steps once the
	// condition is removed from the resource.
	//
	// Version added: v1.18.0
	ResolutionSteps []ResolutionStep
	// Severity shall contain the severity of the condition. Services can replace
	// the value defined in the message registry with a value more applicable to
	// the implementation.
	Severity Health
	// Timestamp shall indicate the time the condition occurred.
	Timestamp string
	// UserAuthenticationSource shall contain the URL to the authentication service
	// that is associated with the username property. This should be used for
	// conditions that result from a user action.
	//
	// Version added: v1.20.0
	UserAuthenticationSource string
	// Username shall contain the username of the account associated with the
	// condition. This should be used for conditions that result from a user
	// action.
	//
	// Version added: v1.20.0
	Username string
}

// UnmarshalJSON unmarshals a Condition object from the raw JSON.
func (c *Condition) UnmarshalJSON(b []byte) error {
	type temp Condition
	var tmp struct {
		temp
		LogEntry          Link `json:"LogEntry"`
		OriginOfCondition Link `json:"OriginOfCondition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Condition(tmp.temp)

	// Extract the links to other entities for later
	c.logEntry = tmp.LogEntry.String()
	c.originOfCondition = tmp.OriginOfCondition.String()

	return nil
}

// LogEntry gets the LogEntry linked resource.
func (c *Condition) LogEntry(client Client) (*LogEntry, error) {
	if c.logEntry == "" {
		return nil, nil
	}
	return GetObject[LogEntry](client, c.logEntry)
}

// OriginOfCondition gets the OriginOfCondition linked resource.
func (c *Condition) OriginOfCondition(client Client) (*Entity, error) {
	if c.originOfCondition == "" {
		return nil, nil
	}
	return GetObject[Entity](client, c.originOfCondition)
}

// ContactInfo shall contain contact information for an individual or
// organization responsible for this resource.
type ContactInfo struct {
	// ContactName shall contain the name of a person or organization to contact
	// for information about this resource.
	//
	// Version added: v1.7.0
	ContactName string
	// EmailAddress shall contain the email address for a person or organization to
	// contact for information about this resource.
	//
	// Version added: v1.7.0
	EmailAddress string
	// PhoneNumber shall contain the phone number for a person or organization to
	// contact for information about this resource.
	//
	// Version added: v1.7.0
	PhoneNumber string
}

// Identifier shall contain any additional identifiers for a resource.
type Identifier struct {
	// DurableName shall contain the world-wide unique identifier for the resource.
	// The string shall be in the format described by the value in the
	// 'DurableNameFormat' property.
	//
	// Version added: v1.1.0
	DurableName string
	// DurableNameFormat shall represent the format of the 'DurableName' property.
	//
	// Version added: v1.1.0
	DurableNameFormat DurableNameFormat
}

// ImportParameters shall contain the import parameters when passing a
// configuration file when using the URI specified by the
// 'MultipartImportConfigurationPushURI' property to import configuration data.
type ImportParameters struct {
	// EncryptionPassphrase shall contain the encryption passphrase for the import
	// file. This property shall not be provided if the import file is not
	// encrypted.
	//
	// Version added: v1.23.0
	EncryptionPassphrase string
}

// Location shall describe the location of a resource.
type Location struct {
	// AltitudeMeters shall contain the altitude of the resource, in meter units,
	// defined as the elevation above sea level.
	//
	// Version added: v1.6.0
	AltitudeMeters *float64 `json:",omitempty"`
	// Contacts shall contain an array of contact information for an individual or
	// organization responsible for this resource.
	//
	// Version added: v1.7.0
	Contacts []ContactInfo
	// Info shall represent the location of the resource.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the 'PostalAddress',
	// 'Placement', and 'PartLocation' properties.
	Info string
	// InfoFormat shall represent the 'Info' property format.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the 'PostalAddress',
	// 'Placement', and 'PartLocation' properties.
	InfoFormat string
	// Latitude shall contain the latitude of the resource specified in degree
	// units using a decimal format and not minutes or seconds.
	//
	// Version added: v1.6.0
	Latitude *float64 `json:",omitempty"`
	// Longitude shall contain the longitude of the resource specified in degree
	// units using a decimal format and not minutes or seconds.
	//
	// Version added: v1.6.0
	Longitude *float64 `json:",omitempty"`
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.1.0
	OEM json.RawMessage `json:"Oem"`
	// PartLocation shall contain the part location for a resource within an
	// enclosure. This representation shall indicate the location of a part within
	// a location specified by the 'Placement' property.
	//
	// Version added: v1.5.0
	PartLocation PartLocation
	// PartLocationContext shall contain a human-readable string to enable
	// differentiation between 'PartLocation' values for parts in the same
	// enclosure, which may include hierarchical information of containing
	// 'PartLocation' values for the part. The value of this property shall not
	// include values of the 'PartLocation' properties for the part itself. The
	// purpose of this value, in conjunction with the 'PartLocation' of the part
	// itself, is to allow clients to determine the physical location of the part
	// without tracing through the 'PartLocation' of multiple resources.
	//
	// Version added: v1.16.0
	PartLocationContext string
	// PhysicalAddress shall contain a physical address for a resource. This
	// property should be present for resources that represent physical objects
	// that can operate without requiring physical containment by another resource.
	// For example, a server chassis might be contained by a rack, but it might
	// also be deployed individually, while a drive is always contained by a
	// chassis and therefore is described by the containing resource.
	//
	// Version added: v1.17.0
	PhysicalAddress PhysicalAddress
	// Placement shall contain a place within the addressed location.
	//
	// Version added: v1.3.0
	Placement Placement
	// PostalAddress shall contain a postal address of the resource.
	//
	// Version added: v1.3.0
	PostalAddress PostalAddress
}

// PartLocation shall describe a location for a resource within an enclosure.
type PartLocation struct {
	// LocationOrdinalValue shall contain the number that represents the location
	// of the part based on the 'LocationType'. 'LocationOrdinalValue' shall be
	// measured based on the Orientation value starting with '0'.
	//
	// Version added: v1.5.0
	LocationOrdinalValue *int `json:",omitempty"`
	// LocationType shall contain the type of location of the part.
	//
	// Version added: v1.5.0
	LocationType LocationType
	// Orientation shall contain the orientation for the ordering used by the
	// 'LocationOrdinalValue' property.
	//
	// Version added: v1.5.0
	Orientation Orientation
	// Reference shall contain the general location within the unit of the part.
	//
	// Version added: v1.5.0
	Reference Reference
	// ServiceLabel shall contain the label assigned for service at the part
	// location.
	//
	// Version added: v1.5.0
	ServiceLabel string
}

func (p *PartLocation) UnmarshalJSON(b []byte) error {
	type temp PartLocation
	var t struct {
		temp
		// Some implementations incorrectly have this as a string, e.g. "1" instead of 1.
		LocationOrdinalValue any
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	i := 0
	if t.LocationOrdinalValue != nil {
		i, err = strconv.Atoi(parseString(t.LocationOrdinalValue))
		if err != nil {
			return err
		}
	}

	*p = PartLocation(t.temp)
	p.LocationOrdinalValue = &i

	return nil
}

// PhysicalAddress shall contain a physical address for a resource.
type PhysicalAddress struct {
	// City shall contain the city, township, or shi (JP) location for this
	// resource.
	//
	// Version added: v1.17.0
	City string
	// Country shall contain the country location for this resource.
	//
	// Version added: v1.17.0
	Country string
	// ISOCountryCode shall contain the ISO 3166-1-defined alpha-2 or alpha-3
	// country code.
	//
	// Version added: v1.17.0
	ISOCountryCode string
	// ISOSubdivisionCode shall contain the ISO 3166-2-defined state, province, or
	// territory subdivision code for this resource.
	//
	// Version added: v1.17.0
	ISOSubdivisionCode string
	// PostalCode shall contain the postal code for this resource. The value shall
	// conform to the RFC5139-defined requirements of the PC field.
	//
	// Version added: v1.17.0
	PostalCode string
	// StateOrProvince shall contain the state, province, or territory location
	// within the country for this resource.
	//
	// Version added: v1.17.0
	StateOrProvince string
	// StreetAddress shall contain the street-level physical address of the
	// resource, including identifiers such as apartment, room, or building to
	// further locate the resource within a given street address.
	//
	// Version added: v1.17.0
	StreetAddress string
}

// Placement shall describe a location within a resource. Examples include a
// shelf in a rack.
type Placement struct {
	// AdditionalInfo shall contain additional information, such as Tile, Column
	// (Post), Wall, or other designation that describes a location that cannot be
	// conveyed with other properties defined for the Placement object.
	//
	// Version added: v1.7.0
	AdditionalInfo string
	// FacilityName shall contain the name of the facility.
	//
	// Version added: v1.23.0
	FacilityName string
	// Rack shall contain the name of the rack within a row.
	//
	// Version added: v1.3.0
	Rack string
	// RackOffset shall be measured from bottom to top, starting with 0.
	//
	// Version added: v1.3.0
	RackOffset *int `json:",omitempty"`
	// RackOffsetUnits shall contain a RackUnit enumeration literal that indicates
	// the type of rack units in use.
	//
	// Version added: v1.3.0
	RackOffsetUnits RackUnits
	// Room shall contain the name or number of the room.
	//
	// Version added: v1.23.0
	Room string
	// Row shall contain the name of the row.
	//
	// Version added: v1.3.0
	Row string
}

// PostalAddress shall describe a postal address for a resource. For more
// information, see RFC5139. Depending on use, the instance can represent a
// past, current, or future location.
type PostalAddress struct {
	// AdditionalCode shall conform to the RFC5139-defined requirements of the
	// ADDCODE field.
	//
	// Version added: v1.3.0
	AdditionalCode string
	// AdditionalInfo shall conform to the requirements of the LOC field as defined
	// in RFC5139. Provides additional information.
	//
	// Version added: v1.7.0
	AdditionalInfo string
	// Building shall conform to the RFC5139-defined requirements of the BLD field.
	// Names the building.
	//
	// Version added: v1.3.0
	Building string
	// City shall conform to the RFC5139-defined requirements of the A3 field.
	// Names a city, township, or shi (JP).
	//
	// Version added: v1.3.0
	City string
	// Community shall conform to the RFC5139-defined requirements of the PCN
	// field. A postal community name.
	//
	// Version added: v1.3.0
	Community string
	// Country shall conform to the RFC5139-defined requirements of the Country
	// field.
	//
	// Version added: v1.3.0
	Country string
	// District shall conform to the RFC5139-defined requirements of the A2 field.
	// Names a county, parish, gun (JP), or district (IN).
	//
	// Version added: v1.3.0
	District string
	// Division shall conform to the RFC5139-defined requirements of the A4 field.
	// Names a city division, borough, city district, ward, or chou (JP).
	//
	// Version added: v1.3.0
	Division string
	// Floor shall conform to the RFC5139-defined requirements of the FLR field.
	// Provides a floor designation.
	//
	// Version added: v1.3.0
	Floor string
	// GPSCoords shall conform to the RFC5139-defined requirements of the ADDCODE
	// field. Shall contain the GPS coordinates of the location. If furnished,
	// expressed in the '[-][nn]n.nnnnnn, [-][nn]n.nnnnn' format. For example, two
	// comma-separated positive or negative numbers with six decimal places of
	// precision.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the Longitude and Latitude
	// properties.
	GPSCoords string
	// HouseNumber shall conform to the RFC5139-defined requirements of the HNO
	// field. The numeric portion of the house number.
	//
	// Version added: v1.3.0
	HouseNumber *int `json:",omitempty"`
	// HouseNumberSuffix shall conform to the RFC5139-defined requirements of the
	// HNS field. Provides a suffix to a house number, (F, B, or 1/2).
	//
	// Version added: v1.3.0
	HouseNumberSuffix string
	// Landmark shall conform to the RFC5139-defined requirements of the LMK field.
	// Identifies a landmark or vanity address.
	//
	// Version added: v1.3.0
	Landmark string
	// LeadingStreetDirection shall conform to the requirements of the PRD field as
	// defined in RFC5139. Names a leading street direction, (N, W, or SE).
	//
	// Version added: v1.3.0
	LeadingStreetDirection string
	// // Location shall conform to the RFC5139-defined requirements of the LOC field.
	// // Provides additional information.
	// //
	// // Version added: v1.3.0
	// //
	// // Deprecated: v1.7.0
	// // This property has been deprecated in favor of the 'AdditionalInfo' property.
	// Location Location
	// Name is the name of the resource or array element.
	//
	// Version added: v1.3.0
	Name string
	// Neighborhood shall conform to the RFC5139-defined requirements of the A5
	// field. Names a neighborhood or block.
	//
	// Version added: v1.3.0
	Neighborhood string
	// POBox shall conform to the RFC5139-defined requirements of the POBOX field.
	// A post office box (PO box).
	//
	// Version added: v1.3.0
	POBox string
	// PlaceType shall conform to the RFC5139-defined requirements of the PLC
	// field. Examples include office and residence.
	//
	// Version added: v1.3.0
	PlaceType string
	// PostalCode shall conform to the RFC5139-defined requirements of the PC
	// field. A postal code (or zip code).
	//
	// Version added: v1.3.0
	PostalCode string
	// Road shall conform to the RFC5139-defined requirements of the RD field.
	// Designates a primary road or street.
	//
	// Version added: v1.3.0
	Road string
	// RoadBranch shall conform to the RFC5139-defined requirements of the RDBR
	// field. Shall contain a post office box (PO box) road branch.
	//
	// Version added: v1.3.0
	RoadBranch string
	// RoadPostModifier shall conform to the RFC5139-defined requirements of the
	// POM field. For example, Extended.
	//
	// Version added: v1.3.0
	RoadPostModifier string
	// RoadPreModifier shall conform to the RFC5139-defined requirements of the PRM
	// field. For example, Old or New.
	//
	// Version added: v1.3.0
	RoadPreModifier string
	// RoadSection shall conform to the RFC5139-defined requirements of the RDSEC
	// field. A road section.
	//
	// Version added: v1.3.0
	RoadSection string
	// RoadSubBranch shall conform to the RFC5139-defined requirements of the
	// RDSUBBR field.
	//
	// Version added: v1.3.0
	RoadSubBranch string
	// Room shall conform to the RFC5139-defined requirements of the ROOM field. A
	// name or number of a room to locate the resource within the unit.
	//
	// Version added: v1.3.0
	Room string
	// Seat shall conform to the RFC5139-defined requirements of the SEAT field. A
	// name or number of a seat, such as the desk, cubicle, or workstation.
	//
	// Version added: v1.3.0
	Seat string
	// Street shall conform to the RFC5139-defined requirements of the A6 field.
	// Names a street.
	//
	// Version added: v1.3.0
	Street string
	// StreetSuffix shall conform to the RFC5139-defined requirements of the STS
	// field. Names a street suffix.
	//
	// Version added: v1.3.0
	StreetSuffix string
	// Territory shall conform to the RFC5139-defined requirements of the A1 field
	// when it names a territory, state, region, province, or prefecture within a
	// country.
	//
	// Version added: v1.3.0
	Territory string
	// TrailingStreetSuffix shall conform to the RFC5139-defined requirements of
	// the POD field. Names a trailing street suffix.
	//
	// Version added: v1.3.0
	TrailingStreetSuffix string
	// Unit shall conform to the RFC5139-defined requirements of the UNIT field.
	// The name or number of a unit, such as the apartment or suite, to locate the
	// resource.
	//
	// Version added: v1.3.0
	Unit string
}

// ReferenceableMember shall contain the location of this element within an
// item.
type ReferenceableMember struct {
	Entity
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetReferenceableMember will get a ReferenceableMember instance from the service.
func GetReferenceableMember(c Client, uri string) (*ReferenceableMember, error) {
	return GetObject[ReferenceableMember](c, uri)
}

// ListReferencedReferenceableMembers gets the collection of ReferenceableMember from
// a provided reference.
func ListReferencedReferenceableMembers(c Client, link string) ([]*ReferenceableMember, error) {
	return GetCollectionObjects[ReferenceableMember](c, link)
}

// Status shall contain any status or health properties of a resource.
type Status struct {
	// Conditions shall represent the active conditions requiring attention in this
	// or a related resource. The conditions may affect the 'Health' or
	// 'HealthRollup' of this resource. The service may roll up multiple conditions
	// originating from a resource, using the 'ConditionInRelatedResource' message
	// from the Base Message Registry. The array order of conditions may change as
	// new conditions occur or as conditions are resolved by the service.
	//
	// Version added: v1.11.0
	Conditions []Condition
	// Health shall represent the health state of the resource without considering
	// its dependent resources. The values shall conform to those defined in the
	// Redfish Specification.
	Health Health
	// HealthRollup shall represent the health state of the resource and its
	// dependent resources. The values shall conform to those defined in the
	// Redfish Specification. For additional property requirements, see the
	// corresponding definition in the Redfish Data Model Specification.
	HealthRollup Health
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// State shall indicate the state of the resource.
	State State
}
