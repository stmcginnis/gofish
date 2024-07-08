//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
	// WarningEventSeverity A condition requiring attention.
	WarningEventSeverity EventSeverity = "Warning"
	// CriticalEventSeverity A critical condition requiring immediate
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
	// CPERLogDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI
	// property is a complete UEFI Specification-defined Common Platform Error Record. The CPER data shall contain a
	// Record Header and at least one Section as defined by the UEFI Specification.
	CPERLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPER"
	// CPERSectionLogDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI
	// property is a single Section of a UEFI Specification-defined Common Platform Error Record. The CPER data shall
	// contain one Section as defined by the UEFI Specification, with no Record Header.
	CPERSectionLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPERSection"
)

type LogEntryCode string

const (

	// AssertLogEntryCode The condition has been asserted.
	AssertLogEntryCode LogEntryCode = "Assert"
	// DeassertLogEntryCode The condition has been deasserted.
	DeassertLogEntryCode LogEntryCode = "Deassert"
	// LowerNonCriticalGoingLowLogEntryCode The reading crossed the
	// Lower Non-critical threshold while going low.
	LowerNonCriticalGoingLowLogEntryCode LogEntryCode = "Lower Non-critical - going low"
	// LowerNonCriticalGoingHighLogEntryCode The reading crossed the
	// Lower Non-critical threshold while going high.
	LowerNonCriticalGoingHighLogEntryCode LogEntryCode = "Lower Non-critical - going high"
	// LowerCriticalGoingLowLogEntryCode The reading crossed the Lower
	// Critical threshold while going low.
	LowerCriticalGoingLowLogEntryCode LogEntryCode = "Lower Critical - going low"
	// LowerCriticalGoingHighLogEntryCode The reading crossed the Lower
	// Critical threshold while going high.
	LowerCriticalGoingHighLogEntryCode LogEntryCode = "Lower Critical - going high"
	// LowerNonRecoverableGoingLowLogEntryCode The reading crossed the
	// Lower Non-recoverable threshold while going low.
	LowerNonRecoverableGoingLowLogEntryCode LogEntryCode = "Lower Non-recoverable - going low"
	// LowerNonRecoverableGoingHighLogEntryCode The reading crossed the
	// Lower Non-recoverable threshold while going high.
	LowerNonRecoverableGoingHighLogEntryCode LogEntryCode = "Lower Non-recoverable - going high"
	// UpperNonCriticalGoingLowLogEntryCode The reading crossed the
	// Upper Non-critical threshold while going low.
	UpperNonCriticalGoingLowLogEntryCode LogEntryCode = "Upper Non-critical - going low"
	// UpperNonCriticalGoingHighLogEntryCode The reading crossed the
	// Upper Non-critical threshold while going high.
	UpperNonCriticalGoingHighLogEntryCode LogEntryCode = "Upper Non-critical - going high"
	// UpperCriticalGoingLowLogEntryCode The reading crossed the Upper
	// Critical threshold while going low.
	UpperCriticalGoingLowLogEntryCode LogEntryCode = "Upper Critical - going low"
	// UpperCriticalGoingHighLogEntryCode The reading crossed the Upper
	// Critical threshold while going high.
	UpperCriticalGoingHighLogEntryCode LogEntryCode = "Upper Critical - going high"
	// UpperNonRecoverableGoingLowLogEntryCode The reading crossed the
	// Upper Non-recoverable threshold while going low.
	UpperNonRecoverableGoingLowLogEntryCode LogEntryCode = "Upper Non-recoverable - going low"
	// UpperNonRecoverableGoingHighLogEntryCode The reading crossed the
	// Upper Non-recoverable threshold while going high.
	UpperNonRecoverableGoingHighLogEntryCode LogEntryCode = "Upper Non-recoverable - going high"
	// TransitionToIdleLogEntryCode The state transitioned to idle.
	TransitionToIdleLogEntryCode LogEntryCode = "Transition to Idle"
	// TransitionToActiveLogEntryCode The state transitioned to active.
	TransitionToActiveLogEntryCode LogEntryCode = "Transition to Active"
	// TransitionToBusyLogEntryCode The state transitioned to busy.
	TransitionToBusyLogEntryCode LogEntryCode = "Transition to Busy"
	// StateDeassertedLogEntryCode The state has been deasserted.
	StateDeassertedLogEntryCode LogEntryCode = "State Deasserted"
	// StateAssertedLogEntryCode The state has been asserted.
	StateAssertedLogEntryCode LogEntryCode = "State Asserted"
	// PredictiveFailureDeassertedLogEntryCode A Predictive Failure is no
	// longer present.
	PredictiveFailureDeassertedLogEntryCode LogEntryCode = "Predictive Failure deasserted"
	// PredictiveFailureAssertedLogEntryCode A Predictive Failure has been
	// detected.
	PredictiveFailureAssertedLogEntryCode LogEntryCode = "Predictive Failure asserted"
	// LimitNotExceededLogEntryCode A limit has not been exceeded.
	LimitNotExceededLogEntryCode LogEntryCode = "Limit Not Exceeded"
	// LimitExceededLogEntryCode A limit has been exceeded.
	LimitExceededLogEntryCode LogEntryCode = "Limit Exceeded"
	// PerformanceMetLogEntryCode Performance meets expectations.
	PerformanceMetLogEntryCode LogEntryCode = "Performance Met"
	// PerformanceLagsLogEntryCode Performance does not meet expectations.
	PerformanceLagsLogEntryCode LogEntryCode = "Performance Lags"
	// TransitionToOKLogEntryCode A state has changed to OK.
	TransitionToOKLogEntryCode LogEntryCode = "Transition to OK"
	// TransitionToNonCriticalFromOKLogEntryCode A state has changed to
	// Non-Critical from OK.
	TransitionToNonCriticalFromOKLogEntryCode LogEntryCode = "Transition to Non-Critical from OK"
	// TransitionToCriticalFromLessSevereLogEntryCode A state has
	// changed to Critical from less severe.
	TransitionToCriticalFromLessSevereLogEntryCode LogEntryCode = "Transition to Critical from less severe"
	// TransitionToNonrecoverableFromLessSevereLogEntryCode A state has
	// changed to Non-recoverable from less severe.
	TransitionToNonrecoverableFromLessSevereLogEntryCode LogEntryCode = "Transition to Non-recoverable from less severe"
	// TransitionToNonCriticalFromMoreSevereLogEntryCode A state has
	// changed to Non-Critical from more severe.
	TransitionToNonCriticalFromMoreSevereLogEntryCode LogEntryCode = "Transition to Non-Critical from more severe"
	// TransitionToCriticalFromNonrecoverableLogEntryCode A state has
	// changed to Critical from Non-recoverable.
	TransitionToCriticalFromNonrecoverableLogEntryCode LogEntryCode = "Transition to Critical from Non-recoverable"
	// TransitionToNonrecoverableLogEntryCode A state has changed to Non-
	// recoverable.
	TransitionToNonrecoverableLogEntryCode LogEntryCode = "Transition to Non-recoverable"
	// MonitorLogEntryCode A Monitor event.
	MonitorLogEntryCode LogEntryCode = "Monitor"
	// InformationalLogEntryCode An Informational event.
	InformationalLogEntryCode LogEntryCode = "Informational"
	// DeviceRemovedDeviceAbsentLogEntryCode A device has been removed
	// or is now absent.
	DeviceRemovedDeviceAbsentLogEntryCode LogEntryCode = "Device Removed / Device Absent"
	// DeviceInsertedDevicePresentLogEntryCode A device has been
	// inserted or is now present.
	DeviceInsertedDevicePresentLogEntryCode LogEntryCode = "Device Inserted / Device Present"
	// DeviceDisabledLogEntryCode A device has been disabled.
	DeviceDisabledLogEntryCode LogEntryCode = "Device Disabled"
	// DeviceEnabledLogEntryCode A device has been enabled.
	DeviceEnabledLogEntryCode LogEntryCode = "Device Enabled"
	// TransitionToRunningLogEntryCode A state has transitioned to Running.
	TransitionToRunningLogEntryCode LogEntryCode = "Transition to Running"
	// TransitionToInTestLogEntryCode A state has transitioned to In Test.
	TransitionToInTestLogEntryCode LogEntryCode = "Transition to In Test"
	// TransitionToPowerOffLogEntryCode A state has transitioned to Power
	// Off.
	TransitionToPowerOffLogEntryCode LogEntryCode = "Transition to Power Off"
	// TransitionToOnLineLogEntryCode A state has transitioned to On Line.
	TransitionToOnLineLogEntryCode LogEntryCode = "Transition to On Line"
	// TransitionToOffLineLogEntryCode A state has transitioned to Off
	// Line.
	TransitionToOffLineLogEntryCode LogEntryCode = "Transition to Off Line"
	// TransitionToOffDutyLogEntryCode A state has transitioned to Off
	// Duty.
	TransitionToOffDutyLogEntryCode LogEntryCode = "Transition to Off Duty"
	// TransitionToDegradedLogEntryCode A state has transitioned to
	// Degraded.
	TransitionToDegradedLogEntryCode LogEntryCode = "Transition to Degraded"
	// TransitionToPowerSaveLogEntryCode A state has transitioned to Power
	// Save.
	TransitionToPowerSaveLogEntryCode LogEntryCode = "Transition to Power Save"
	// InstallErrorLogEntryCode An Install Error has been detected.
	InstallErrorLogEntryCode LogEntryCode = "Install Error"
	// FullyRedundantLogEntryCode Indicates that full redundancy has been
	// regained.
	FullyRedundantLogEntryCode LogEntryCode = "Fully Redundant"
	// RedundancyLostLogEntryCode Entered any non-redundant state, including
	// Non-redundant: Insufficient Resources.
	RedundancyLostLogEntryCode LogEntryCode = "Redundancy Lost"
	// RedundancyDegradedLogEntryCode Redundancy still exists, but at less
	// than full level.
	RedundancyDegradedLogEntryCode LogEntryCode = "Redundancy Degraded"
	// NonredundantSufficientResourcesFromRedundantLogEntryCode Redundancy has
	// been lost but unit is functioning with minimum resources needed for
	// normal operation.
	NonredundantSufficientResourcesFromRedundantLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"
	// NonredundantSufficientResourcesFromInsufficientResourcesLogEntryCode Unit has regained minimum resources needed for
	// normal operation.
	NonredundantSufficientResourcesFromInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"
	// NonredundantInsufficientResourcesLogEntryCode Unit is non-redundant
	// and has insufficient resource to maintain normal operation.
	NonredundantInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Insufficient Resources"
	// RedundancyDegradedFromFullyRedundantLogEntryCode Unit has lost
	// some redundant resource(s) but is still in a redundant state.
	RedundancyDegradedFromFullyRedundantLogEntryCode LogEntryCode = "Redundancy Degraded from Fully Redundant"
	// RedundancyDegradedFromNonredundantLogEntryCode Unit has regained
	// some resource(s) and is redundant but not fully redundant.
	RedundancyDegradedFromNonredundantLogEntryCode LogEntryCode = "Redundancy Degraded from Non-redundant"
	// D0PowerStateLogEntryCode The ACPI defined D0 Power State.
	D0PowerStateLogEntryCode LogEntryCode = "D0 Power State"
	// D1PowerStateLogEntryCode The ACPI defined D1 Power State.
	D1PowerStateLogEntryCode LogEntryCode = "D1 Power State"
	// D2PowerStateLogEntryCode The ACPI defined D2 Power State.
	D2PowerStateLogEntryCode LogEntryCode = "D2 Power State"
	// D3PowerStateLogEntryCode The ACPI defined D3 Power State.
	D3PowerStateLogEntryCode LogEntryCode = "D3 Power State"
	// OEMLogEntryCode An OEM defined event.
	OEMLogEntryCode LogEntryCode = "OEM"
)

type LogEntryType string

const (

	// EventLogEntryType Contains a Redfish-defined message (event).
	EventLogEntryType LogEntryType = "Event"
	// SELLogEntryType Contains a legacy IPMI System Event Log (SEL) entry.
	SELLogEntryType LogEntryType = "SEL"
	// OemLogEntryType Contains an entry in an OEM-defined format.
	OemLogEntryType LogEntryType = "Oem"
	// CXLLogEntryType is a CXL log entry.
	CXLLogEntryType LogEntryType = "CXL"
)

type OriginatorTypes string

const (
	// ClientOriginatorTypes is a client of the service created this log entry.
	ClientOriginatorTypes OriginatorTypes = "Client"
	// InternalOriginatorTypes is a process running on the service created this log entry.
	InternalOriginatorTypes OriginatorTypes = "Internal"
	// SupportingServiceOriginatorTypes is a process not running on the service but running on a supporting service, such
	// as RDE implementations, UEFI, or host processes, created this log entry.
	SupportingServiceOriginatorTypes OriginatorTypes = "SupportingService"
)

type SensorType string

const (

	// PlatformSecurityViolationAttemptSensorType is a platform security
	// sensor.
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
	// PowerSupplyConverterSensorType is a sensor for a power supply or DC-
	// to-DC converter.
	PowerSupplyConverterSensorType SensorType = "Power Supply / Converter"
	// PowerUnitSensorType is a sensor for a power unit.
	PowerUnitSensorType SensorType = "PowerUnit"
	// CoolingDeviceSensorType is a sensor for a cooling device.
	CoolingDeviceSensorType SensorType = "CoolingDevice"
	// OtherUnitsBasedSensorSensorType is a sensor for a miscellaneous analog
	// sensor.
	OtherUnitsBasedSensorSensorType SensorType = "Other Units-based Sensor"
	// MemorySensorType is a sensor for a memory device.
	MemorySensorType SensorType = "Memory"
	// DriveSlotBaySensorType is a sensor for a drive slot or bay.
	DriveSlotBaySensorType SensorType = "Drive Slot/Bay"
	// POSTMemoryResizeSensorType is a sensor for a POST memory resize event.
	POSTMemoryResizeSensorType SensorType = "POST Memory Resize"
	// SystemFirmwareProgressSensorType is a sensor for a system firmware
	// progress event.
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
	// MicrocontrollerCoprocessorSensorType is a sensor for a microcontroller
	// or coprocessor.
	MicrocontrollerCoprocessorSensorType SensorType = "Microcontroller/Coprocessor"
	// AddinCardSensorType is a sensor for an add-in card.
	AddinCardSensorType SensorType = "Add-in Card"
	// ChassisSensorType is a sensor for a chassis.
	ChassisSensorType SensorType = "Chassis"
	// ChipSetSensorType is a sensor for a chipset.
	ChipSetSensorType SensorType = "ChipSet"
	// OtherFRUSensorType is a sensor for an other type of FRU.
	OtherFRUSensorType SensorType = "Other FRU"
	// CableInterconnectSensorType is a sensor for a cable or interconnect type
	// of device.
	CableInterconnectSensorType SensorType = "Cable/Interconnect"
	// TerminatorSensorType is a sensor for a terminator.
	TerminatorSensorType SensorType = "Terminator"
	// SystemBootRestartSensorType is a sensor for a system boot or restart
	// event.
	SystemBootRestartSensorType SensorType = "SystemBoot/Restart"
	// BootErrorSensorType is a sensor for a boot error event.
	BootErrorSensorType SensorType = "Boot Error"
	// BaseOSBootInstallationStatusSensorType is a sensor for a base OS boot or
	// installation status event.
	BaseOSBootInstallationStatusSensorType SensorType = "BaseOSBoot/InstallationStatus"
	// OSStopShutdownSensorType is a sensor for an OS stop or shutdown event
	OSStopShutdownSensorType SensorType = "OS Stop/Shutdown"
	// SlotConnectorSensorType is a sensor for a slot or connector.
	SlotConnectorSensorType SensorType = "Slot/Connector"
	// SystemACPIPowerStateSensorType is a sensor for an ACPI power state
	// event.
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
	// ManagementSubsystemHealthSensorType is a sensor for a management
	// subsystem health event.
	ManagementSubsystemHealthSensorType SensorType = "Management Subsystem Health"
	// BatterySensorType is a sensor for a battery.
	BatterySensorType SensorType = "Battery"
	// SessionAuditSensorType is a sensor for a session audit event.
	SessionAuditSensorType SensorType = "Session Audit"
	// VersionChangeSensorType is a sensor for a version change event.
	VersionChangeSensorType SensorType = "Version Change"
	// FRUStateSensorType is a sensor for a FRU state event.
	FRUStateSensorType SensorType = "FRUState"
	// OEMSensorType is an OEM defined sensor.
	OEMSensorType SensorType = "OEM"
)

// LogEntry shall represent the log format for log services.
type LogEntry struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AdditionalDataSizeBytes shall contain the size of the additional data retrieved from the URI specified by the
	// AdditionalDataURI property for this log entry.
	AdditionalDataSizeBytes int
	// AdditionalDataURI shall contain the URI at which to access the additional data for this log entry, using the
	// Redfish protocol and authentication methods. If both DiagnosticData and AdditionalDataURI are present,
	// DiagnosticData shall contain the Base64-encoding of the data retrieved from the URI specified by the
	// AdditionalDataURI property.
	AdditionalDataURI string
	// CPER shall contain the details for a CPER section or record that is the source of this log entry.
	CPER CPER
	// CXLEntryType shall contain the specific CXL entry type. This property shall only be present if EntryType
	// contains 'CXL'.
	CXLEntryType CXLEntryType
	// Created shall be the time at which the log entry was created.
	Created string
	// Description provides a description of this resource.
	Description string
	// DiagnosticData shall contain a Base64-encoded string that represents diagnostic data associated with this log
	// entry. The contents shall depend on the value of the DiagnosticDataType property. The length of the value should
	// not exceed 4 KB. Larger diagnostic data payloads should omit this property and use the AdditionalDataURI
	// property to reference the data. If both DiagnosticData and AdditionalDataURI are present, DiagnosticData shall
	// contain the Base64-encoding of the data retrieved from the URI specified by the AdditionalDataURI property.
	DiagnosticData string
	// DiagnosticDataType shall contain the type of data available in the DiagnosticData property or retrieved from the
	// URI specified by the AdditionalDataURI property.
	DiagnosticDataType LogDiagnosticDataTypes
	// EntryCode shall be present if the EntryType value is
	// SEL. These enumerations are the values from tables 42-1 and 42-2 of
	// the IPMI specification.
	EntryCode LogEntryCode
	// EntryType shall represent the type of LogEntry. If
	// the resource represents an IPMI SEL log entry, the value shall be SEL.
	// If the resource represents an Event log, the value shall be Event. If
	// the resource represents an OEM log format, the value shall be Oem.
	EntryType LogEntryType
	// EventGroupID shall indicate that events are related and shall have the
	// same value in the case where multiple Event messages are produced by the
	// same root cause. Implementations shall use separate values for events
	// with separate root cause. There shall not be ordering of events implied
	// by the value of this property.
	EventGroupID int `json:"EventGroupId"`
	// EventID records an Event and the value shall indicate a unique identifier
	// for the event, the format of which is implementation dependent.
	EventID string `json:"EventId"`
	// EventTimestamp records an Event and the value shall be the time the event
	// occurred.
	EventTimestamp string
	// FirstOverflowTimestamp shall contain the timestamp of the first overflow captured after this log entry. If this
	// log entry is the most recent log entry in the log service, this property shall not be present if no overflow
	// errors occurred after the time of this log entry. If this log entry is not the most recent log entry in the log
	// service, this property shall not be present if no overflow errors occurred between the time of this log entry
	// and the time of the next log entry.
	FirstOverflowTimestamp string
	// GeneratorId if EntryType is `SEL`, this property shall contain the
	// 'Generator ID' field of the IPMI SEL Event Record. If EntryType is
	// not `SEL`, this property should not be present.
	GeneratorID string `json:"GeneratorId"`
	// LastOverflowTimestamp shall contain the timestamp of the last overflow captured after this log entry. If this
	// log entry is the most recent log entry in the log service, this property shall not be present if no overflow
	// errors occurred after the time of this log entry. If this log entry is not the most recent log entry in the log
	// service, this property shall not be present if no overflow errors occurred between the time of this log entry
	// and the time of the next log entry.
	LastOverflowTimestamp string
	// Message shall be the Message property of
	// the event if the EntryType is Event, the Description if the EntryType
	// is SEL, and OEM Specific if the EntryType is Oem.
	Message string
	// MessageArgs contains message arguments to be substituted into
	// the message included or in the message looked up via a registry.
	MessageArgs []string
	// MessageId shall the MessageId property
	// of the event if the EntryType is Event, the three IPMI Event Data
	// bytes if the EntryType is SEL, and OEM Specific if the EntryType is
	// Oem. The format of this property shall be as defined in the Redfish
	// specification. If representing the three IPMI Event Data bytes, the
	// format should follow the pattern '^0[xX](([a-fA-F]|[0-9]){2}){3}$',
	// where Event Data 1 is the first byte in the string, Event Data 2 is
	// the second byte in the string, and Event Data 3 is the third byte in
	// the string.
	MessageID string `json:"MessageId"`
	// Modified shall contain the date and time when the log
	// entry was last modified. This property shall not appear if the log
	// entry has not been modified since it was created.
	Modified string
	// OEMDiagnosticDataType shall contain the OEM-defined type of data available in the DiagnosticData property or
	// retrieved from the URI specified by the AdditionalDataURI property. This property shall be present if
	// DiagnosticDataType is 'OEM'.
	OEMDiagnosticDataType string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemLogEntryCode shall represent the OEM
	// specific Log Entry Code type of the Entry. This property shall only
	// be present if the value of EntryType is SEL and the value of
	// LogEntryCode is OEM.
	OemLogEntryCode string
	// OemRecordFormat shall represent the OEM
	// specific format of the Entry. This property shall be required if the
	// value of EntryType is Oem.
	OemRecordFormat string
	// OemSensorType is used if the value of EntryType is SEL and the value
	// of SensorType is OEM.
	OemSensorType string
	// Originator shall contain the source of the log entry.
	Originator string
	// OriginatorType shall contain the type of originator data.
	OriginatorType OriginatorTypes
	// OverflowErrorCount shall contain the count of overflow errors that occurred after this log entry. If this log
	// entry is the most recent log entry in the log service, this property shall not be present if no overflow errors
	// occurred after the time of this log entry. If this log entry is not the most recent log entry in the log
	// service, this property shall not be present if no overflow errors occurred between the time of this log entry
	// and the time of the next log entry.
	OverflowErrorCount int
	// Persistency shall indicate whether the log entry is persistent across a cold reset of the device.
	Persistency bool
	// Resolution shall contain the resolution of the log entry. Services should replace the resolution defined in the
	// message registry with a more specific resolution in a log entry.
	Resolution string
	// ResolutionSteps shall contain an array of recommended steps to resolve the cause of the log entry. This property
	// shall not be present if the Severity property contains 'OK'. A client can stop executing the resolution steps
	// once the Resolved property resource contains 'true' or the Health property in the associated resource referenced
	// by the OriginOfCondition property contains 'OK'.
	ResolutionSteps []ResolutionStep
	// Resolved shall contain an indication if the cause of the log entry has been resolved or repaired. The value
	// 'true' shall indicate if the cause of the log entry has been resolved or repaired. This property shall contain
	// the value 'false' if the log entry is still active. The value 'false' shall be the initial state. Clients should
	// ignore this property if Severity contains 'OK'.
	Resolved bool
	// SensorNumber shall be the IPMI sensor
	// number if the EntryType is SEL, the count of events if the EntryType
	// is Event, and OEM Specific if EntryType is Oem.
	SensorNumber int
	// SensorType shall be present if the EntryType value is
	// SEL. These enumerations are the values from table 42-3 of the IPMI
	// specification.
	SensorType SensorType
	// ServiceProviderNotified shall contain an indication if the log entry has been sent to the service provider.
	ServiceProviderNotified bool
	// Severity shall be the severity of the
	// condition resulting in the log entry, as defined in the Status section
	// of the Redfish specification.
	Severity EventSeverity
	// SpecificEventExistsInGroup shall indicate that this log entry is equivalent to another log entry, with a more
	// specific definition, within the same EventGroupId. For example, the 'DriveFailed' message from the Storage
	// Device Message Registry is more specific than the 'ResourceStatusChangedCritical' message from the Resource
	// Event Message Registry, when both occur with the same EventGroupId. This property shall contain 'true' if a more
	// specific event is available, and shall contain 'false' if no equivalent event exists in the same EventGroupId.
	// If this property is absent, the value shall be assumed to be 'false'.
	SpecificEventExistsInGroup bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// originOfCondition shall be an href that references the resource for which the log is associated.
	OriginOfCondition string
	// RelatedItem shall contain an array of links to resources that are related to this log entry.
	RelatedItem []string
	// RelatedItemCount is the number of related items.
	RelatedItemCount  int
	relatedLogEntries []string
	// RelatedLogEntriesCount is the number of related log entries.
	RelatedLogEntriesCount int
}

// UnmarshalJSON unmarshals a LogEntry object from the raw JSON.
func (logentry *LogEntry) UnmarshalJSON(b []byte) error {
	type temp LogEntry
	var t struct {
		temp
		Links struct {
			// OriginOfCondition shall be an href that
			// references the resource for which the log is associated.
			OriginOfCondition common.Link
			// RelatedItem shall contain an array of links to resources that are related to this log entry. It shall not
			// contain links to LogEntry resources. RelatedLogEntries is used to reference related log entries. This property
			// shall not contain the value of the OriginOfCondition property.
			RelatedItem      []string
			RelatedItemCount int `json:"RelatedItem@odata.count"`
			// RelatedLogEntries shall contain an array of links to resources of type LogEntry in this or other log services
			// that are related to this log entry.
			RelatedLogEntries      common.Links
			RelatedLogEntriesCount int `json:"RelatedLogEntries@odata.count"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*logentry = LogEntry(t.temp)
	logentry.OriginOfCondition = t.Links.OriginOfCondition.String()
	logentry.RelatedItem = t.Links.RelatedItem
	logentry.RelatedItemCount = t.Links.RelatedItemCount
	logentry.relatedLogEntries = t.Links.RelatedLogEntries.ToStrings()
	logentry.RelatedItemCount = t.Links.RelatedLogEntriesCount

	// This is a read/write object, so we need to save the raw object data for later
	logentry.rawData = b

	return nil
}

// RelatedLogEntries gets the set of LogEntry in this or other log services that are related to this log entry.
func (logentry *LogEntry) RelatedLogEntries() ([]*LogEntry, error) {
	return common.GetObjects[LogEntry](logentry.GetClient(), logentry.relatedLogEntries)
}

// Update commits updates to this object's properties to the running system.
func (logentry *LogEntry) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LogEntry)
	original.UnmarshalJSON(logentry.rawData)

	readWriteFields := []string{
		"Resolved",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(logentry).Elem()

	return logentry.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLogEntry will get a LogEntry instance from the service.
func GetLogEntry(c common.Client, uri string) (*LogEntry, error) {
	return common.GetObject[LogEntry](c, uri)
}

// ListReferencedLogEntrys gets the collection of LogEntry from
// a provided reference.
func ListReferencedLogEntrys(c common.Client, link string) ([]*LogEntry, error) {
	return common.GetCollectionObjects[LogEntry](c, link)
}
