//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

// LogEntryTypes is the type of log entry.
type LogEntryTypes string

const (
	// EventLogEntryTypes contains Redfish-defined messages (events).
	EventLogEntryTypes LogEntryTypes = "Event"
	// SELLogEntryTypes contains legacy IPMI System Event Log (SEL) entries.
	SELLogEntryTypes LogEntryTypes = "SEL"
	// MultipleLogEntryTypes contains multiple Log Entry types or a
	// single entry type cannot be guaranteed by the Log Service.
	MultipleLogEntryTypes LogEntryTypes = "Multiple"
	// OEMLogEntryTypes contains entries in an OEM-defined format.
	OEMLogEntryTypes LogEntryTypes = "OEM"
)

type LogPurpose string

const (
	// DiagnosticLogPurpose provides information for diagnosing hardware or software issues, such as error
	// conditions, sensor threshold trips, or exception cases.
	DiagnosticLogPurpose LogPurpose = "Diagnostic"
	// OperationsLogPurpose provides information about management operations that have a significant impact on
	// the system, such as firmware updates, system resets, and storage volume creation.
	OperationsLogPurpose LogPurpose = "Operations"
	// SecurityLogPurpose provides security-related information such as authentication, authorization, and data
	// access logging required for security audits.
	SecurityLogPurpose LogPurpose = "Security"
	// TelemetryLogPurpose provides telemetry history, typically collected on a regular basis.
	TelemetryLogPurpose LogPurpose = "Telemetry"
	// ExternalEntityLogPurpose The log exposes log entries provided by external entities, such as external users,
	// system firmware, operating systems, or management applications.
	ExternalEntityLogPurpose LogPurpose = "ExternalEntity"
	// OEMLogPurpose The log is used for an OEM-defined purpose.
	OEMLogPurpose LogPurpose = "OEM"
)

// OverWritePolicy is the log overwriting policy.
type OverWritePolicy string

const (

	// UnknownOverWritePolicy means the policy is not known or is undefined.
	UnknownOverWritePolicy OverWritePolicy = "Unknown"
	// WrapsWhenFullOverWritePolicy means when full, new entries to the Log will
	// overwrite previous entries.
	WrapsWhenFullOverWritePolicy OverWritePolicy = "WrapsWhenFull"
	// NeverOverWritesOverWritePolicy means when full, new entries to the Log will
	// be discarded.
	NeverOverWritesOverWritePolicy OverWritePolicy = "NeverOverWrites"
)

// DiagnosticDataDetail is the detailed information for a supported CollectDiagnosticData action.
type DiagnosticDataDetail struct {
	// DiagnosticDataType indicates the type of diagnostic data to collect.
	DiagnosticDataType LogDiagnosticDataTypes
	// EstimatedDuration represents the estimated total time required to generate the data.
	EstimatedDuration string
	// EstimatedSizeBytes represents the estimated size of the data to be collected.
	EstimatedSizeBytes int
	// OEMDiagnosticDataType represents the type of data if DiagnosticDataType is `OEM`.
	OEMDiagnosticDataType string
}

// LogService is used to represent a log service for a Redfish
// implementation.
type LogService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AutoDSTEnabled shall indicate whether the log service is configured for automatic Daylight Saving Time (DST)
	// adjustment. DST adjustment shall not modify the timestamp of existing log entries.
	AutoDSTEnabled bool
	// DateTime shall represent the current DateTime value that the log service
	// is using, with offset from UTC, in Redfish Timestamp format.
	DateTime string
	// DateTimeLocalOffset shall represent the offset from UTC time that the
	// current value of DataTime property contains.
	DateTimeLocalOffset string
	// Description provides a description of this resource.
	Description string
	// DiagnosticDataDetails provides information about the supported CollectDiagnosticData requests
	DiagnosticDataDetails []DiagnosticDataDetail
	// Entries shall reference a collection of resources of type LogEntry.
	entries string
	// LogEntryType shall represent the
	// EntryType of all LogEntry resources contained in the Entries
	// collection. If a single EntryType for all LogEntry resources cannot
	// be determined or guaranteed by the Service, the value of this property
	// shall be 'Multiple'.
	LogEntryType LogEntryTypes
	// MaxNumberOfRecords shall be the maximum numbers of LogEntry resources in
	// the Entries collection for this service.
	MaxNumberOfRecords uint64
	// OEMLogPurpose shall contain the OEM-specified purpose of the log if LogPurposes contains 'OEM'.
	OEMLogPurpose string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OverWritePolicy shall indicate the
	// policy of the log service when the MaxNumberOfRecords has been
	// reached. Unknown indicates the log overwrite policy is unknown.
	// WrapsWhenFull indicates that the log overwrites its entries with new
	// entries when the log has reached its maximum capacity. NeverOverwrites
	// indicates that the log never overwrites its entries by the new entries
	// and ceases logging when the limit has been reached.
	OverWritePolicy OverWritePolicy
	// Overflow shall indicate whether the log service has overflowed and is no longer able to store new logs.
	Overflow bool
	// Persistency shall indicate whether the log service is persistent across a cold reset of the device.
	Persistency bool
	// ServiceEnabled shall be a boolean
	// indicating whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SyslogFilters shall describe all desired syslog messages to be logged locally. If this property contains an
	// empty array, all messages shall be logged.
	SyslogFilters []SyslogFilter
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// clearLogTarget is the URL to send ClearLog actions to.
	clearLogTarget string

	// collectDiagnosticDataTarget is the URL to send CollectDiagnosticData actions to. (v1.2+)
	collectDiagnosticDataTarget string
	// collectDiagnosticInfoTarget is the URL to get ActionInfo about the CollectDiagnosticData action.
	collectDiagnosticInfoTarget string
}

// UnmarshalJSON unmarshals a LogService object from the raw JSON.
func (logservice *LogService) UnmarshalJSON(b []byte) error {
	type temp LogService
	type Actions struct {
		ClearLog              common.ActionTarget `json:"#LogService.ClearLog"`
		CollectDiagnosticData common.ActionTarget `json:"#LogService.CollectDiagnosticData"`
	}
	var t struct {
		temp
		Entries common.Link
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*logservice = LogService(t.temp)
	logservice.entries = t.Entries.String()
	logservice.clearLogTarget = t.Actions.ClearLog.Target
	logservice.collectDiagnosticDataTarget = t.Actions.CollectDiagnosticData.Target
	logservice.collectDiagnosticInfoTarget = t.Actions.CollectDiagnosticData.ActionInfoTarget

	// This is a read/write object, so we need to save the raw object data for later
	logservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (logservice *LogService) Update() error {
	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"ServiceEnabled",
	}

	return logservice.UpdateFromRawData(logservice, logservice.rawData, readWriteFields)
}

// GetLogService will get a LogService instance from the service.
func GetLogService(c common.Client, uri string) (*LogService, error) {
	return common.GetObject[LogService](c, uri)
}

// ListReferencedLogServices gets the collection of LogService from a provided reference.
func ListReferencedLogServices(c common.Client, link string) ([]*LogService, error) {
	return common.GetCollectionObjects[LogService](c, link)
}

// Entries gets the log entries of this service.
func (logservice *LogService) Entries() ([]*LogEntry, error) {
	return ListReferencedLogEntrys(logservice.GetClient(), logservice.entries)
}

// FilteredEntries gets the log entries of this service with filtering applied (e.g. skip, top).
func (logservice *LogService) FilteredEntries(options ...common.FilterOption) ([]*LogEntry, error) {
	var filter common.Filter
	filter.SetFilter(options...)
	return ListReferencedLogEntrys(logservice.GetClient(), fmt.Sprintf("%s%s", logservice.entries, filter))
}

// SupportsClearLog indicates if the ClearLog action is supported.
func (logservice *LogService) SupportsClearLog() bool {
	return logservice.clearLogTarget != ""
}

// ClearLog shall delete all entries found in the Entries collection for this
// Log Service.
func (logservice *LogService) ClearLog() error {
	err := logservice.Post(logservice.clearLogTarget, struct{}{})
	if err == nil {
		return nil
	}

	// As of LogService 1.3.0, need to pass the LogEntryCollection etag. If our first attempt failed, try that.
	entryCollection := &struct {
		ETag string `json:"@odata.etag"`
	}{}

	retryErr := logservice.Get(logservice.GetClient(), logservice.entries, entryCollection)
	if retryErr == nil {
		payload := struct {
			LogEntriesETag string
		}{LogEntriesETag: strings.Trim(entryCollection.ETag, "\"")}

		retryErr = logservice.Post(logservice.clearLogTarget, payload)
		if retryErr == nil {
			return nil
		}
	}

	// Fall back to broken implementation to workaround vendor bug
	t := struct {
		Action string
	}{
		Action: "LogService.ClearLog",
	}

	return logservice.Post(logservice.clearLogTarget, t)
}

type CollectDiagnosticDataParameters struct {
	// DiagnosticDataType (required) shall contain the type of diagnostic data to collect.
	DiagnosticDataType LogDiagnosticDataTypes `json:",omitempty"`
	// OEMDiagnosticDataType (optional) shall contain the OEM-defined type of diagnostic data if the
	// DiagnosticDataType is set to `OEMDiagnosticDataTypes`
	OEMDiagnosticDataType string `json:",omitempty"`
	// TargetDevice (optional) is a link to the device to collect diagnostic data from.
	TargetDevice common.Link `json:",omitempty"`
	// TargetURI (optional) is the URI to access when sending the diagnostic data.
	TargetURI string `json:",omitempty"`
	// UserName (optional) shall contain the username to access the URI specified by `TargetURI`.
	UserName string `json:",omitempty"`
	// Password (optional) shall contain the password to access the URI specified by `TargetURI`.
	Password string `json:",omitempty"`
	// TransferProtocol is the network protocol to be used to send the diagnostic data.
	TransferProtocol TransferProtocolType `json:",omitempty"`
}

// SupportsCollectDiagnosticData indicates if the CollectDiagnosticData action is supported.
func (logservice *LogService) SupportsCollectDiagnosticData() bool {
	return logservice.collectDiagnosticDataTarget != ""
}

// For Redfish v1.2+
// CollectDiagnosticData shall trigger the generation of a diagnostic data dump.
// Returns the URI to a LogEntry that will contain the DiagnosticData in the `AdditionalDataURI` when ready.
// This URI should be polled until the log is generated.
func (logservice *LogService) CollectDiagnosticData(parameters *CollectDiagnosticDataParameters) (string, error) {
	if !logservice.SupportsCollectDiagnosticData() {
		return "", errors.New("CollectDiagnosticsData not supported by this service")
	}

	resp, err := logservice.PostWithResponse(logservice.collectDiagnosticDataTarget, parameters)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if location := resp.Header["Location"]; len(location) > 0 && location[0] != "" {
		return location[0], nil
	}

	return "", nil
}

// For Redfish v1.2+
// CollectDiagnosticDataActionInfo, if supported, provides the ActionInfo for a CollectDiagnosticData action.
func (logservice *LogService) CollectDiagnosticDataActionInfo() (*ActionInfo, error) {
	if logservice.collectDiagnosticInfoTarget == "" {
		return nil, errors.New("CollectDiagnosticData ActionInfo not supported by this service")
	}

	return common.GetObject[ActionInfo](logservice.GetClient(), logservice.collectDiagnosticInfoTarget)
}
