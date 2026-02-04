//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/LogEntry.v1_21_0.json
// 2025.4 - #LogEntry.v1_21_0.LogEntry

package schemas

import (
	"encoding/json"
)

type CXLEntryType string

const (
	// DynamicCapacityCXLEntryType is a CXL Dynamic Capacity log entry.
	DynamicCapacityCXLEntryType CXLEntryType = "DynamicCapacity"
	// InformationalCXLEntryType is a CXL informational log entry.
	InformationalCXLEntryType CXLEntryType = "Informational"
	// WarningCXLEntryType is a CXL warning log entry.
	WarningCXLEntryType CXLEntryType = "Warning"
	// FailureCXLEntryType is a CXL failure log entry.
	FailureCXLEntryType CXLEntryType = "Failure"
	// FatalCXLEntryType is a CXL fatal log entry.
	FatalCXLEntryType CXLEntryType = "Fatal"
)

type EventSeverity string

const (
	// OKEventSeverity Informational or operating normally.
	OKEventSeverity EventSeverity = "OK"
	// WarningEventSeverity is a condition that requires attention.
	WarningEventSeverity EventSeverity = "Warning"
	// CriticalEventSeverity is a critical condition that requires immediate
	// attention.
	CriticalEventSeverity EventSeverity = "Critical"
)

type LogDiagnosticDataTypes string

const (
	// ManagerLogDiagnosticDataTypes Manager diagnostic data.
	ManagerLogDiagnosticDataTypes LogDiagnosticDataTypes = "Manager"
	// PreOSLogDiagnosticDataTypes Pre-OS diagnostic data.
	PreOSLogDiagnosticDataTypes LogDiagnosticDataTypes = "PreOS"
	// OSLogDiagnosticDataTypes Operating system (OS) diagnostic data.
	OSLogDiagnosticDataTypes LogDiagnosticDataTypes = "OS"
	// OEMLogDiagnosticDataTypes OEM diagnostic data.
	OEMLogDiagnosticDataTypes LogDiagnosticDataTypes = "OEM"
	// CPERLogDiagnosticDataTypes shall indicate the data provided at the URI
	// specified by the 'AdditionalDataURI' property is a complete UEFI
	// Specification-defined Common Platform Error Record. The CPER data shall
	// contain a Record Header and at least one Section as defined by the UEFI
	// Specification.
	CPERLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPER"
	// CPERSectionLogDiagnosticDataTypes shall indicate the data provided at the
	// URI specified by the 'AdditionalDataURI' property is a single Section of a
	// UEFI Specification-defined Common Platform Error Record. The CPER data shall
	// contain one Section as defined by the UEFI Specification, with no Record
	// Header.
	CPERSectionLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPERSection"
	// DeviceLogDiagnosticDataTypes shall indicate the data provided at the URI
	// specified by the 'AdditionalDataURI' property contains vendor-defined
	// diagnostic data for a device.
	DeviceLogDiagnosticDataTypes LogDiagnosticDataTypes = "Device"
)

type LogEntryCode string

const (
	// AssertLogEntryCode The condition has been asserted.
	AssertLogEntryCode LogEntryCode = "Assert"
	// DeassertLogEntryCode The condition has been deasserted.
	DeassertLogEntryCode LogEntryCode = "Deassert"
	// LowerNoncriticalgoinglowLogEntryCode The reading crossed the Lower
	// Non-critical threshold while going low.
	LowerNoncriticalgoinglowLogEntryCode LogEntryCode = "Lower Non-critical - going low"
	// LowerNoncriticalgoinghighLogEntryCode The reading crossed the Lower
	// Non-critical threshold while going high.
	LowerNoncriticalgoinghighLogEntryCode LogEntryCode = "Lower Non-critical - going high"
	// LowerCriticalgoinglowLogEntryCode The reading crossed the Lower Critical
	// threshold while going low.
	LowerCriticalgoinglowLogEntryCode LogEntryCode = "Lower Critical - going low"
	// LowerCriticalgoinghighLogEntryCode The reading crossed the Lower Critical
	// threshold while going high.
	LowerCriticalgoinghighLogEntryCode LogEntryCode = "Lower Critical - going high"
	// LowerNonrecoverablegoinglowLogEntryCode The reading crossed the Lower
	// Non-recoverable threshold while going low.
	LowerNonrecoverablegoinglowLogEntryCode LogEntryCode = "Lower Non-recoverable - going low"
	// LowerNonrecoverablegoinghighLogEntryCode The reading crossed the Lower
	// Non-recoverable threshold while going high.
	LowerNonrecoverablegoinghighLogEntryCode LogEntryCode = "Lower Non-recoverable - going high"
	// UpperNoncriticalgoinglowLogEntryCode The reading crossed the Upper
	// Non-critical threshold while going low.
	UpperNoncriticalgoinglowLogEntryCode LogEntryCode = "Upper Non-critical - going low"
	// UpperNoncriticalgoinghighLogEntryCode The reading crossed the Upper
	// Non-critical threshold while going high.
	UpperNoncriticalgoinghighLogEntryCode LogEntryCode = "Upper Non-critical - going high"
	// UpperCriticalgoinglowLogEntryCode The reading crossed the Upper Critical
	// threshold while going low.
	UpperCriticalgoinglowLogEntryCode LogEntryCode = "Upper Critical - going low"
	// UpperCriticalgoinghighLogEntryCode The reading crossed the Upper Critical
	// threshold while going high.
	UpperCriticalgoinghighLogEntryCode LogEntryCode = "Upper Critical - going high"
	// UpperNonrecoverablegoinglowLogEntryCode The reading crossed the Upper
	// Non-recoverable threshold while going low.
	UpperNonrecoverablegoinglowLogEntryCode LogEntryCode = "Upper Non-recoverable - going low"
	// UpperNonrecoverablegoinghighLogEntryCode The reading crossed the Upper
	// Non-recoverable threshold while going high.
	UpperNonrecoverablegoinghighLogEntryCode LogEntryCode = "Upper Non-recoverable - going high"
	// TransitiontoIdleLogEntryCode The state transitioned to idle.
	TransitiontoIdleLogEntryCode LogEntryCode = "Transition to Idle"
	// TransitiontoActiveLogEntryCode The state transitioned to active.
	TransitiontoActiveLogEntryCode LogEntryCode = "Transition to Active"
	// TransitiontoBusyLogEntryCode The state transitioned to busy.
	TransitiontoBusyLogEntryCode LogEntryCode = "Transition to Busy"
	// StateDeassertedLogEntryCode The state has been deasserted.
	StateDeassertedLogEntryCode LogEntryCode = "State Deasserted"
	// StateAssertedLogEntryCode The state has been asserted.
	StateAssertedLogEntryCode LogEntryCode = "State Asserted"
	// PredictiveFailuredeassertedLogEntryCode is a Predictive Failure is no longer
	// present.
	PredictiveFailuredeassertedLogEntryCode LogEntryCode = "Predictive Failure deasserted"
	// PredictiveFailureassertedLogEntryCode is a Predictive Failure has been
	// detected.
	PredictiveFailureassertedLogEntryCode LogEntryCode = "Predictive Failure asserted"
	// LimitNotExceededLogEntryCode is a limit has not been exceeded.
	LimitNotExceededLogEntryCode LogEntryCode = "Limit Not Exceeded"
	// LimitExceededLogEntryCode is a limit has been exceeded.
	LimitExceededLogEntryCode LogEntryCode = "Limit Exceeded"
	// PerformanceMetLogEntryCode Performance meets expectations.
	PerformanceMetLogEntryCode LogEntryCode = "Performance Met"
	// PerformanceLagsLogEntryCode Performance does not meet expectations.
	PerformanceLagsLogEntryCode LogEntryCode = "Performance Lags"
	// TransitiontoOKLogEntryCode is a state has changed to OK.
	TransitiontoOKLogEntryCode LogEntryCode = "Transition to OK"
	// TransitiontoNonCriticalfromOKLogEntryCode is a state has changed to
	// Non-Critical from OK.
	TransitiontoNonCriticalfromOKLogEntryCode LogEntryCode = "Transition to Non-Critical from OK"
	// TransitiontoCriticalfromlesssevereLogEntryCode is a state has changed to
	// Critical from less severe.
	TransitiontoCriticalfromlesssevereLogEntryCode LogEntryCode = "Transition to Critical from less severe"
	// TransitiontoNonrecoverablefromlesssevereLogEntryCode is a state has changed
	// to Non-recoverable from less severe.
	TransitiontoNonrecoverablefromlesssevereLogEntryCode LogEntryCode = "Transition to Non-recoverable from less severe"
	// TransitiontoNonCriticalfrommoresevereLogEntryCode is a state has changed to
	// Non-Critical from more severe.
	TransitiontoNonCriticalfrommoresevereLogEntryCode LogEntryCode = "Transition to Non-Critical from more severe"
	// TransitiontoCriticalfromNonrecoverableLogEntryCode is a state has changed to
	// Critical from Non-recoverable.
	TransitiontoCriticalfromNonrecoverableLogEntryCode LogEntryCode = "Transition to Critical from Non-recoverable"
	// TransitiontoNonrecoverableLogEntryCode is a state has changed to
	// Non-recoverable.
	TransitiontoNonrecoverableLogEntryCode LogEntryCode = "Transition to Non-recoverable"
	// MonitorLogEntryCode is a monitor event.
	MonitorLogEntryCode LogEntryCode = "Monitor"
	// InformationalLogEntryCode is an informational event.
	InformationalLogEntryCode LogEntryCode = "Informational"
	// DeviceRemovedDeviceAbsentLogEntryCode is a device has been removed or is
	// absent.
	DeviceRemovedDeviceAbsentLogEntryCode LogEntryCode = "Device Removed / Device Absent"
	// DeviceInsertedDevicePresentLogEntryCode is a device has been inserted or is
	// present.
	DeviceInsertedDevicePresentLogEntryCode LogEntryCode = "Device Inserted / Device Present"
	// DeviceDisabledLogEntryCode is a device has been disabled.
	DeviceDisabledLogEntryCode LogEntryCode = "Device Disabled"
	// DeviceEnabledLogEntryCode is a device has been enabled.
	DeviceEnabledLogEntryCode LogEntryCode = "Device Enabled"
	// TransitiontoRunningLogEntryCode is a state has transitioned to Running.
	TransitiontoRunningLogEntryCode LogEntryCode = "Transition to Running"
	// TransitiontoInTestLogEntryCode is a state has transitioned to In Test.
	TransitiontoInTestLogEntryCode LogEntryCode = "Transition to In Test"
	// TransitiontoPowerOffLogEntryCode is a state has transitioned to Power Off.
	TransitiontoPowerOffLogEntryCode LogEntryCode = "Transition to Power Off"
	// TransitiontoOnLineLogEntryCode is a state has transitioned to On Line.
	TransitiontoOnLineLogEntryCode LogEntryCode = "Transition to On Line"
	// TransitiontoOffLineLogEntryCode is a state has transitioned to Off Line.
	TransitiontoOffLineLogEntryCode LogEntryCode = "Transition to Off Line"
	// TransitiontoOffDutyLogEntryCode is a state has transitioned to Off Duty.
	TransitiontoOffDutyLogEntryCode LogEntryCode = "Transition to Off Duty"
	// TransitiontoDegradedLogEntryCode is a state has transitioned to Degraded.
	TransitiontoDegradedLogEntryCode LogEntryCode = "Transition to Degraded"
	// TransitiontoPowerSaveLogEntryCode is a state has transitioned to Power Save.
	TransitiontoPowerSaveLogEntryCode LogEntryCode = "Transition to Power Save"
	// InstallErrorLogEntryCode is an install error has been detected.
	InstallErrorLogEntryCode LogEntryCode = "Install Error"
	// FullyRedundantLogEntryCode Indicates that full redundancy has been regained.
	FullyRedundantLogEntryCode LogEntryCode = "Fully Redundant"
	// RedundancyLostLogEntryCode Entered any non-redundant state, including
	// Non-redundant: Insufficient Resources.
	RedundancyLostLogEntryCode LogEntryCode = "Redundancy Lost"
	// RedundancyDegradedLogEntryCode Redundancy still exists, but at less than
	// full level.
	RedundancyDegradedLogEntryCode LogEntryCode = "Redundancy Degraded"
	// NonredundantSufficientResourcesfromRedundantLogEntryCode Redundancy has been
	// lost but unit is functioning with minimum resources needed for normal
	// operation.
	NonredundantSufficientResourcesfromRedundantLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"
	// NonredundantSufficientResourcesfromInsufficientResourcesLogEntryCode Unit
	// has regained minimum resources needed for normal operation.
	NonredundantSufficientResourcesfromInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"
	// NonredundantInsufficientResourcesLogEntryCode Unit is non-redundant and has
	// insufficient resources to maintain normal operation.
	NonredundantInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Insufficient Resources"
	// RedundancyDegradedfromFullyRedundantLogEntryCode Unit has lost some
	// redundant resource(s) but is still in a redundant state.
	RedundancyDegradedfromFullyRedundantLogEntryCode LogEntryCode = "Redundancy Degraded from Fully Redundant"
	// RedundancyDegradedfromNonredundantLogEntryCode Unit has regained some
	// resource(s) and is redundant but not fully redundant.
	RedundancyDegradedfromNonredundantLogEntryCode LogEntryCode = "Redundancy Degraded from Non-redundant"
	// D0PowerStateLogEntryCode The ACPI-defined D0 power state.
	D0PowerStateLogEntryCode LogEntryCode = "D0 Power State"
	// D1PowerStateLogEntryCode The ACPI-defined D1 power state.
	D1PowerStateLogEntryCode LogEntryCode = "D1 Power State"
	// D2PowerStateLogEntryCode The ACPI-defined D2 power state.
	D2PowerStateLogEntryCode LogEntryCode = "D2 Power State"
	// D3PowerStateLogEntryCode The ACPI-defined D3 power state.
	D3PowerStateLogEntryCode LogEntryCode = "D3 Power State"
	// OEMLogEntryCode is an OEM-defined event.
	OEMLogEntryCode LogEntryCode = "OEM"
)

type LogEntryType string

const (
	// EventLogEntryType is a Redfish-defined message.
	EventLogEntryType LogEntryType = "Event"
	// SELLogEntryType is a legacy IPMI System Event Log (SEL) entry.
	SELLogEntryType LogEntryType = "SEL"
	// OemLogEntryType is an entry in an OEM-defined format.
	OemLogEntryType LogEntryType = "Oem"
	// CXLLogEntryType is a CXL log entry.
	CXLLogEntryType LogEntryType = "CXL"
)

type OriginatorTypes string

const (
	// ClientOriginatorTypes is a client of the service created this log entry.
	ClientOriginatorTypes OriginatorTypes = "Client"
	// InternalOriginatorTypes is a process running on the service created this log
	// entry.
	InternalOriginatorTypes OriginatorTypes = "Internal"
	// SupportingServiceOriginatorTypes is a process not running on the service but
	// running on a supporting service, such as RDE implementations, UEFI, or host
	// processes, created this log entry.
	SupportingServiceOriginatorTypes OriginatorTypes = "SupportingService"
)

type SensorType string

const (
	// PlatformSecurityViolationAttemptSensorType is a platform security sensor.
	PlatformSecurityViolationAttemptSensorType SensorType = "Platform Security Violation Attempt"
	// TemperatureSensorType is a temperature sensor.
	TemperatureSensorType SensorType = "Temperature"
	// VoltageSensorType is a voltage sensor.
	VoltageSensorType SensorType = "Voltage"
	// CurrentSensorType is a current sensor.
	CurrentSensorType SensorType = "Current"
	// FanSensorType is a fan sensor.
	FanSensorType SensorType = "Fan"
	// PhysicalChassisSecuritySensorType is a physical security sensor.
	PhysicalChassisSecuritySensorType SensorType = "Physical Chassis Security"
	// ProcessorSensorType is a sensor for a processor.
	ProcessorSensorType SensorType = "Processor"
	// PowerSupplyConverterSensorType is a sensor for a power supply or DC-to-DC
	// converter.
	PowerSupplyConverterSensorType SensorType = "Power Supply / Converter"
	// PowerUnitSensorType is a sensor for a power unit.
	PowerUnitSensorType SensorType = "PowerUnit"
	// CoolingDeviceSensorType is a sensor for a cooling device.
	CoolingDeviceSensorType SensorType = "CoolingDevice"
	// OtherUnitsbasedSensorSensorType is a sensor for a miscellaneous analog
	// sensor.
	OtherUnitsbasedSensorSensorType SensorType = "Other Units-based Sensor"
	// MemorySensorType is a sensor for a memory device.
	MemorySensorType SensorType = "Memory"
	// DriveSlotBaySensorType is a sensor for a drive slot or bay.
	DriveSlotBaySensorType SensorType = "Drive Slot/Bay"
	// POSTMemoryResizeSensorType is a sensor for a POST memory resize event.
	POSTMemoryResizeSensorType SensorType = "POST Memory Resize"
	// SystemFirmwareProgressSensorType is a sensor for a system firmware progress
	// event.
	SystemFirmwareProgressSensorType SensorType = "System Firmware Progress"
	// EventLoggingDisabledSensorType is a sensor for the event log.
	EventLoggingDisabledSensorType SensorType = "Event Logging Disabled"
	// SystemEventSensorType is a sensor for a system event.
	SystemEventSensorType SensorType = "System Event"
	// CriticalInterruptSensorType is a sensor for a critical interrupt event.
	CriticalInterruptSensorType SensorType = "Critical Interrupt"
	// ButtonSwitchSensorType is a sensor for a button or switch.
	ButtonSwitchSensorType SensorType = "Button/Switch"
	// ModuleBoardSensorType is a sensor for a module or board.
	ModuleBoardSensorType SensorType = "Module/Board"
	// MicrocontrollerCoprocessorSensorType is a sensor for a microcontroller or
	// coprocessor.
	MicrocontrollerCoprocessorSensorType SensorType = "Microcontroller/Coprocessor"
	// AddinCardSensorType is a sensor for an add-in card.
	AddinCardSensorType SensorType = "Add-in Card"
	// ChassisSensorType is a sensor for a chassis.
	ChassisSensorType SensorType = "Chassis"
	// ChipSetSensorType is a sensor for a chipset.
	ChipSetSensorType SensorType = "ChipSet"
	// OtherFRUSensorType is a sensor for another type of FRU.
	OtherFRUSensorType SensorType = "Other FRU"
	// CableInterconnectSensorType is a sensor for a cable or interconnect device
	// type.
	CableInterconnectSensorType SensorType = "Cable/Interconnect"
	// TerminatorSensorType is a sensor for a terminator.
	TerminatorSensorType SensorType = "Terminator"
	// SystemBootRestartSensorType is a sensor for a system boot or restart event.
	SystemBootRestartSensorType SensorType = "SystemBoot/Restart"
	// BootErrorSensorType is a sensor for a boot error event.
	BootErrorSensorType SensorType = "Boot Error"
	// BaseOSBootInstallationStatusSensorType is a sensor for a base OS boot or
	// installation status event.
	BaseOSBootInstallationStatusSensorType SensorType = "BaseOSBoot/InstallationStatus"
	// OSStopShutdownSensorType is a sensor for an OS stop or shutdown event.
	OSStopShutdownSensorType SensorType = "OS Stop/Shutdown"
	// SlotConnectorSensorType is a sensor for a slot or connector.
	SlotConnectorSensorType SensorType = "Slot/Connector"
	// SystemACPIPowerStateSensorType is a sensor for an ACPI power state event.
	SystemACPIPowerStateSensorType SensorType = "System ACPI PowerState"
	// WatchdogSensorType is a sensor for a watchdog event.
	WatchdogSensorType SensorType = "Watchdog"
	// PlatformAlertSensorType is a sensor for a platform alert event.
	PlatformAlertSensorType SensorType = "Platform Alert"
	// EntityPresenceSensorType is a sensor for an entity presence event.
	EntityPresenceSensorType SensorType = "Entity Presence"
	// MonitorASICICSensorType is a sensor for a monitor ASIC or IC.
	MonitorASICICSensorType SensorType = "Monitor ASIC/IC"
	// LANSensorType is a sensor for a LAN device.
	LANSensorType SensorType = "LAN"
	// ManagementSubsystemHealthSensorType is a sensor for a management subsystem
	// health event.
	ManagementSubsystemHealthSensorType SensorType = "Management Subsystem Health"
	// BatterySensorType is a sensor for a battery.
	BatterySensorType SensorType = "Battery"
	// SessionAuditSensorType is a sensor for a session audit event.
	SessionAuditSensorType SensorType = "Session Audit"
	// VersionChangeSensorType is a sensor for a version change event.
	VersionChangeSensorType SensorType = "Version Change"
	// FRUStateSensorType is a sensor for a FRU state event.
	FRUStateSensorType SensorType = "FRUState"
	// OEMSensorType is an OEM-defined sensor.
	OEMSensorType SensorType = "OEM"
)

// LogEntry shall represent the log format for log services in a Redfish
// implementation.
type LogEntry struct {
	Entity
	// AdditionalDataSizeBytes shall contain the size of the additional data
	// retrieved from the URI specified by the 'AdditionalDataURI' property for
	// this log entry.
	//
	// Version added: v1.7.0
	AdditionalDataSizeBytes *int `json:",omitempty"`
	// AdditionalDataURI shall contain the URI at which to access the additional
	// data for this log entry, using the Redfish protocol and authentication
	// methods. If both 'DiagnosticData' and 'AdditionalDataURI' are present,
	// 'DiagnosticData' shall contain a Base64-encoded string, with padding
	// characters, of the data retrieved from the URI specified by the
	// 'AdditionalDataURI' property.
	//
	// Version added: v1.7.0
	AdditionalDataURI string
	// CPER shall contain the details for a CPER section or record that is the
	// source of this log entry.
	//
	// Version added: v1.15.0
	CPER CPER
	// CXLEntryType shall contain the specific CXL entry type. This property shall
	// only be present if 'EntryType' contains 'CXL'.
	//
	// Version added: v1.14.0
	CXLEntryType CXLEntryType
	// Created shall contain the date and time when the log entry was created.
	Created string
	// DiagnosticData shall contain a Base64-encoded string, with padding
	// characters, that represents diagnostic data associated with this log entry.
	// The contents shall depend on the value of the 'DiagnosticDataType' property.
	// The length of the value should not exceed 4 KB. Larger diagnostic data
	// payloads should omit this property and use the 'AdditionalDataURI' property
	// to reference the data. If both 'DiagnosticData' and 'AdditionalDataURI' are
	// present, 'DiagnosticData' shall contain the Base64-encoding of the data
	// retrieved from the URI specified by the 'AdditionalDataURI' property.
	//
	// Version added: v1.15.0
	DiagnosticData string
	// DiagnosticDataType shall contain the type of data available in the
	// 'DiagnosticData' property or retrieved from the URI specified by the
	// 'AdditionalDataURI' property. The 'OriginOfCondition' property, if present,
	// shall contain a link to the resource that represents the device from which
	// the diagnostic data was collected.
	//
	// Version added: v1.7.0
	DiagnosticDataType LogDiagnosticDataTypes
	// EntryCode shall contain the entry code for the log entry if the 'EntryType'
	// is 'SEL'. Tables 42-1 and 42-2 of the IPMI Specification v2.0 revision 1.1
	// describe these enumerations.
	EntryCode LogEntryCode
	// EntryType shall represent the type of log entry. If the resource represents
	// an IPMI SEL entry, the value shall contain 'SEL'. If the resource represents
	// a Redfish event log entry, the value shall contain 'Event'. If the resource
	// represents a CXL event record, the value shall contain 'CXL'. If the
	// resource represents an OEM log entry format, the value shall contain 'Oem'.
	EntryType LogEntryType
	// EventGroupID shall indicate that events are related and shall have the same
	// value in the case where multiple event messages are produced by the same
	// root cause. Implementations shall use separate values for events with
	// separate root cause. There shall not be ordering of events implied by this
	// property's value.
	//
	// Version added: v1.4.0
	EventGroupID *int `json:"EventGroupId,omitempty"`
	// EventID shall indicate a unique identifier for the event, the format of
	// which is implementation dependent.
	//
	// Version added: v1.1.0
	EventID string `json:"EventId"`
	// EventTimestamp shall contain the date and time when the event occurred.
	//
	// Version added: v1.1.0
	EventTimestamp string
	// EventType shall indicate the type of event.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.4.0
	// This property has been deprecated. Starting with Redfish Specification v1.6
	// (Event v1.3), subscriptions are based on the 'RegistryPrefix' and
	// 'ResourceType' properties and not on the 'EventType' property.
	EventType EventType
	// FirstOverflowTimestamp shall contain the timestamp of the first overflow
	// captured after this log entry. If this log entry is the most recent log
	// entry in the log service, this property shall not be present if no overflow
	// errors occurred after the time of this log entry. If this log entry is not
	// the most recent log entry in the log service, this property shall not be
	// present if no overflow errors occurred between the time of this log entry
	// and the time of the next log entry.
	//
	// Version added: v1.14.0
	FirstOverflowTimestamp string
	// GeneratorID shall contain the 'Generator ID' field of the IPMI SEL Event
	// Record. If 'EntryType' is not 'SEL', this property should not be present.
	//
	// Version added: v1.5.0
	GeneratorID string `json:"GeneratorId"`
	// LastOverflowTimestamp shall contain the timestamp of the last overflow
	// captured after this log entry. If this log entry is the most recent log
	// entry in the log service, this property shall not be present if no overflow
	// errors occurred after the time of this log entry. If this log entry is not
	// the most recent log entry in the log service, this property shall not be
	// present if no overflow errors occurred between the time of this log entry
	// and the time of the next log entry.
	//
	// Version added: v1.14.0
	LastOverflowTimestamp string
	// Message shall contain the message of the log entry. This property decodes
	// from the entry type. If the entry type is 'Event', this property contains a
	// message. If the entry type is 'SEL', this property contains an SEL-specific
	// message, following the format specified in Table 32-1, SEL Event Records, in
	// the IPMI Specification v2.0 revision 1.1. If the entry type is 'CXL', this
	// property contains the CXL event record as a string of hex bytes in the
	// pattern '^([a-fA-F0-9]{2})+$'. Otherwise, this property contains an
	// OEM-specific log entry. In most cases, this property contains the actual log
	// entry.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted
	// for the arguments in the message when looked up in the message registry. It
	// has the same semantics as the 'MessageArgs' property in the Redfish
	// 'MessageRegistry' schema. If the corresponding 'ParamType' value contains
	// 'number', the service shall convert the number to a string representation of
	// the number.
	MessageArgs []string
	// MessageID shall contain the 'MessageId', event data, or OEM-specific
	// information. This property decodes from the entry type. If the entry type is
	// 'Event', this property contains a Redfish Specification-defined 'MessageId'
	// property of the event. If the entry type is 'SEL', the format should follow
	// the pattern '^0[xX](([a-fA-F]|[0-9]){2}){4}$', which results in a string in
	// the form '0xNNaabbcc', where 'NN' is the EventDir/EventType byte, 'aa' is
	// the Event Data 1 byte, 'bb' is Event Data 2 byte, 'cc' is Event Data 3 byte,
	// corresponding with bytes 13-16 in the IPMI SEL Event Record. If the entry
	// type is 'CXL', this property shall not be present. Otherwise, this property
	// contains OEM-specific information.
	MessageID string `json:"MessageId"`
	// Modified shall contain the date and time when the log entry was last
	// modified. This property shall not appear if the log entry has not been
	// modified since it was created.
	//
	// Version added: v1.6.0
	Modified string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMDiagnosticDataType shall contain the OEM-defined type of data available
	// in the 'DiagnosticData' property or retrieved from the URI specified by the
	// 'AdditionalDataURI' property. This property shall be present if
	// 'DiagnosticDataType' is 'OEM'.
	//
	// Version added: v1.7.0
	OEMDiagnosticDataType string
	// OemLogEntryCode shall represent the OEM-specific Log Entry Code type of the
	// Entry. This property shall only be present if 'EntryType' is 'SEL' and
	// 'LogEntryCode' is 'OEM'.
	//
	// Version added: v1.3.0
	OemLogEntryCode string
	// OemRecordFormat shall represent the OEM-specific format of the entry. This
	// property shall be required if the 'EntryType' value is 'Oem'.
	OemRecordFormat string
	// OemSensorType shall represent the OEM-specific sensor type of the entry.
	// This property shall only be used if 'EntryType' is 'SEL' and 'SensorType' is
	// 'OEM'.
	//
	// Version added: v1.3.0
	OemSensorType string
	// OriginAddress shall contain the IP address, with scheme, of the user
	// associated with the log entry. This should be used for audit logs that
	// result from a user action.
	//
	// Version added: v1.21.0
	OriginAddress string
	// OriginOfConditionUnavailable shall indicate whether the 'OriginOfCondition'
	// link is unavailable. If 'true', services shall not expand the
	// 'OriginOfCondition' link. If this property is not present, the value shall
	// be assumed to be 'false'.
	//
	// Version added: v1.20.0
	OriginOfConditionUnavailable bool
	// Originator shall contain the source of the log entry.
	//
	// Version added: v1.11.0
	Originator string
	// OriginatorType shall contain the type of originator data.
	//
	// Version added: v1.11.0
	OriginatorType OriginatorTypes
	// OverflowErrorCount shall contain the count of overflow errors that occurred
	// after this log entry. If this log entry is the most recent log entry in the
	// log service, this property shall not be present if no overflow errors
	// occurred after the time of this log entry. If this log entry is not the most
	// recent log entry in the log service, this property shall not be present if
	// no overflow errors occurred between the time of this log entry and the time
	// of the next log entry.
	//
	// Version added: v1.14.0
	OverflowErrorCount int
	// PartNumber shall contain the manufacturer-provided part number for the
	// source of this log entry.
	//
	// Version added: v1.19.0
	PartNumber string
	// Persistency shall indicate whether the log entry is persistent across a cold
	// reset of the device.
	//
	// Version added: v1.14.0
	Persistency bool
	// Resolution shall contain the resolution of the log entry. Services should
	// replace the resolution defined in the message registry with a more specific
	// resolution in a log entry.
	//
	// Version added: v1.9.0
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the
	// cause of the log entry. This property shall not be present if the 'Severity'
	// property contains 'OK'. A client can stop executing the resolution steps
	// once the 'Resolved' property resource contains 'true' or the 'Health'
	// property in the associated resource referenced by the 'OriginOfCondition'
	// property contains 'OK'.
	//
	// Version added: v1.16.0
	ResolutionSteps []ResolutionStep
	// Resolved shall contain an indication if the cause of the log entry has been
	// resolved or repaired. The value 'true' shall indicate if the cause of the
	// log entry has been resolved or repaired. This property shall contain the
	// value 'false' if the log entry is still active. The value 'false' shall be
	// the initial state. Clients should ignore this property if 'Severity'
	// contains 'OK'.
	//
	// Version added: v1.8.0
	Resolved bool
	// SensorNumber shall contain the IPMI sensor number if the value of the
	// 'EntryType' property is 'SEL'. This property should not appear in the
	// resource for other values of 'EntryType'.
	SensorNumber *int `json:",omitempty"`
	// SensorType shall contain the sensor type to which the log entry pertains if
	// the entry type is 'SEL'. Table 42-3, Sensor Type Codes, in the IPMI
	// Specification v2.0 revision 1.1 describes these enumerations.
	SensorType SensorType
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the source of this log entry.
	//
	// Version added: v1.19.0
	SerialNumber string
	// ServiceProviderNotified shall contain an indication if the log entry has
	// been sent to the service provider.
	//
	// Version added: v1.9.0
	ServiceProviderNotified bool
	// Severity shall contain the severity of the condition that created the log
	// entry. If 'EntryType' contains 'Event', services can replace the value
	// defined in the message registry with a value more applicable to the
	// implementation.
	Severity EventSeverity
	// SpecificEventExistsInGroup shall indicate that this log entry is equivalent
	// to another log entry, with a more specific definition, within the same
	// 'EventGroupId'. For example, the 'DriveFailed' message from the Storage
	// Device Message Registry is more specific than the
	// 'ResourceStatusChangedCritical' message from the Resource Event Message
	// Registry, when both occur with the same 'EventGroupId'. This property shall
	// contain 'true' if a more specific event is available, and shall contain
	// 'false' if no equivalent event exists in the same 'EventGroupId'. If this
	// property is absent, the value shall be assumed to be 'false'.
	//
	// Version added: v1.13.0
	SpecificEventExistsInGroup bool
	// UserAuthenticationSource shall contain the URL to the authentication service
	// that is associated with the username property. This should be used for audit
	// logs that result from a user action.
	//
	// Version added: v1.17.0
	UserAuthenticationSource string
	// Username shall contain the username of the account associated with the log
	// entry. This should be used for audit logs that result from a user action.
	//
	// Version added: v1.17.0
	Username string
	// originOfCondition is the URI for OriginOfCondition.
	originOfCondition string
	// relatedItem are the URIs for RelatedItem.
	relatedItem []string
	// relatedLogEntries are the URIs for RelatedLogEntries.
	relatedLogEntries []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a LogEntry object from the raw JSON.
func (l *LogEntry) UnmarshalJSON(b []byte) error {
	type temp LogEntry
	type lLinks struct {
		OriginOfCondition Link  `json:"OriginOfCondition"`
		RelatedItem       Links `json:"RelatedItem"`
		RelatedLogEntries Links `json:"RelatedLogEntries"`
	}
	var tmp struct {
		temp
		Links lLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LogEntry(tmp.temp)

	// Extract the links to other entities for later
	l.originOfCondition = tmp.Links.OriginOfCondition.String()
	l.relatedItem = tmp.Links.RelatedItem.ToStrings()
	l.relatedLogEntries = tmp.Links.RelatedLogEntries.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	l.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LogEntry) Update() error {
	readWriteFields := []string{
		"Resolved",
	}

	return l.UpdateFromRawData(l, l.RawData, readWriteFields)
}

// GetLogEntry will get a LogEntry instance from the service.
func GetLogEntry(c Client, uri string) (*LogEntry, error) {
	return GetObject[LogEntry](c, uri)
}

// ListReferencedLogEntrys gets the collection of LogEntry from
// a provided reference.
func ListReferencedLogEntrys(c Client, link string) ([]*LogEntry, error) {
	return GetCollectionObjects[LogEntry](c, link)
}

// OriginOfCondition gets the OriginOfCondition linked resource.
func (l *LogEntry) OriginOfCondition() (*Entity, error) {
	if l.originOfCondition == "" {
		return nil, nil
	}
	return GetObject[Entity](l.client, l.originOfCondition)
}

// RelatedItem gets the RelatedItem linked resources.
func (l *LogEntry) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](l.client, l.relatedItem)
}

// RelatedLogEntries gets the RelatedLogEntries linked resources.
func (l *LogEntry) RelatedLogEntries() ([]*LogEntry, error) {
	return GetObjects[LogEntry](l.client, l.relatedLogEntries)
}
