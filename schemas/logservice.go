//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #LogService.v1_9_0.LogService

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AutoClearResolvedEntries string

const (
	// ClearEventGroupAutoClearResolvedEntries shall indicate this log service
	// automatically clears all log entries that contain the value 'true' for the
	// 'Resolved' property and other entries within the same 'EventGroupId'.
	ClearEventGroupAutoClearResolvedEntries AutoClearResolvedEntries = "ClearEventGroup"
	// RetainCauseResolutionEntriesAutoClearResolvedEntries shall indicate this log
	// service retains the entries containing the original cause and the final
	// resolution, but automatically clears other entries containing the
	// intermediate results within the same 'EventGroupId'. For example, the
	// original cause of a fan failure is indicated by an entry containing the
	// 'FanFailed' message key in the 'MessageId' property, followed by the entries
	// containing 'FanRemoved' and 'FanInserted' message keys corresponding to user
	// actions. Finally, an entry showing the fan failure is repaired with
	// 'FanRestored' message key in the 'MessageId' property. In this case, the
	// entries with 'FanFailed' and 'FanRestored' message keys are retained, but
	// other entries within the same 'EventGroupId' are automatically cleared.
	RetainCauseResolutionEntriesAutoClearResolvedEntries AutoClearResolvedEntries = "RetainCauseResolutionEntries"
	// UpdateCauseEntryAutoClearResolvedEntries shall indicate this log service
	// updates the entry containing the original cause for the repaired status, but
	// automatically clears other entries within the same 'EventGroupId'. For
	// example, the original cause of a fan failure is indicated by an entry
	// containing 'FanFailed' message key in the 'MessageId' property, followed by
	// the entries containing 'FanRemoved' and 'FanInserted' message keys
	// corresponding to user actions. Finally, an entry showing the fan failure is
	// repaired with 'FanRestored' message key in the 'MessageId' property. In this
	// case, the entry with 'FanFailed' message key is updated to show the repaired
	// status, such as updating the value of Resolved property to 'true' and the
	// timestamp contained by the 'Modified' property, but other entries within the
	// same 'EventGroupId' are automatically cleared.
	UpdateCauseEntryAutoClearResolvedEntries AutoClearResolvedEntries = "UpdateCauseEntry"
	// NoneAutoClearResolvedEntries shall indicate this log service does not
	// automatically clear the resolved log entries.
	NoneAutoClearResolvedEntries AutoClearResolvedEntries = "None"
)

type LogEntryTypes string

const (
	// EventLogEntryTypes The log contains Redfish-defined messages.
	EventLogEntryTypes LogEntryTypes = "Event"
	// SELLogEntryTypes The log contains legacy IPMI System Event Log (SEL)
	// entries.
	SELLogEntryTypes LogEntryTypes = "SEL"
	// MultipleLogEntryTypes The log contains multiple log entry types and,
	// therefore, the log service cannot guarantee a single entry type.
	MultipleLogEntryTypes LogEntryTypes = "Multiple"
	// OEMLogEntryTypes The log contains entries in an OEM-defined format.
	OEMLogEntryTypes LogEntryTypes = "OEM"
	// CXLLogEntryTypes The log contains CXL log entries.
	CXLLogEntryTypes LogEntryTypes = "CXL"
)

type LogPurpose string

const (
	// DiagnosticLogPurpose The log provides information for diagnosing hardware or
	// software issues, such as error conditions, sensor threshold trips, or
	// exception cases.
	DiagnosticLogPurpose LogPurpose = "Diagnostic"
	// OperationsLogPurpose The log provides information about management
	// operations that have a significant impact on the system, such as firmware
	// updates, system resets, and storage volume creation.
	OperationsLogPurpose LogPurpose = "Operations"
	// SecurityLogPurpose The log provides security-related information such as
	// authentication, authorization, and data access logging required for security
	// audits.
	SecurityLogPurpose LogPurpose = "Security"
	// TelemetryLogPurpose The log provides telemetry history, typically collected
	// on a regular basis.
	TelemetryLogPurpose LogPurpose = "Telemetry"
	// ExternalEntityLogPurpose The log exposes log entries provided by external
	// entities, such as external users, system firmware, operating systems, or
	// management applications.
	ExternalEntityLogPurpose LogPurpose = "ExternalEntity"
	// OEMLogPurpose The log is used for an OEM-defined purpose.
	OEMLogPurpose LogPurpose = "OEM"
)

type LogServiceOverWritePolicy string

const (
	// UnknownLogServiceOverWritePolicy The overwrite policy is not known or is undefined.
	UnknownLogServiceOverWritePolicy LogServiceOverWritePolicy = "Unknown"
	// WrapsWhenFullLogServiceOverWritePolicy When full, new entries to the log overwrite
	// earlier entries.
	WrapsWhenFullLogServiceOverWritePolicy LogServiceOverWritePolicy = "WrapsWhenFull"
	// NeverOverWritesLogServiceOverWritePolicy When full, new entries to the log are
	// discarded.
	NeverOverWritesLogServiceOverWritePolicy LogServiceOverWritePolicy = "NeverOverWrites"
)

// SyslogFacility is This type shall specify the syslog facility codes as
// program types. Facility values are described in the RFC5424.
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

// SyslogSeverity is This type shall specify the syslog severity levels as an
// application-specific rating used to describe the urgency of the message.
// 'Emergency' should be reserved for messages indicating the system is unusable
// and 'Debug' should only be used when debugging a program. Severity values are
// described in RFC5424.
type SyslogSeverity string

const (
	// EmergencySyslogSeverity is a panic condition.
	EmergencySyslogSeverity SyslogSeverity = "Emergency"
	// AlertSyslogSeverity is a condition that should be corrected immediately,
	// such as a corrupted system database.
	AlertSyslogSeverity SyslogSeverity = "Alert"
	// CriticalSyslogSeverity Hard device errors.
	CriticalSyslogSeverity SyslogSeverity = "Critical"
	// ErrorSyslogSeverity is an Error.
	ErrorSyslogSeverity SyslogSeverity = "Error"
	// WarningSyslogSeverity is a Warning.
	WarningSyslogSeverity SyslogSeverity = "Warning"
	// NoticeSyslogSeverity Conditions that are not error conditions, but that
	// might require special handling.
	NoticeSyslogSeverity SyslogSeverity = "Notice"
	// InformationalSyslogSeverity Informational only.
	InformationalSyslogSeverity SyslogSeverity = "Informational"
	// DebugSyslogSeverity Messages that contain information normally of use only
	// when debugging a program.
	DebugSyslogSeverity SyslogSeverity = "Debug"
	// AllSyslogSeverity is a message of any severity.
	AllSyslogSeverity SyslogSeverity = "All"
)

type LogServiceTransferProtocolType string

const (
	// CIFSLogServiceTransferProtocolType Common Internet File System (CIFS).
	CIFSLogServiceTransferProtocolType LogServiceTransferProtocolType = "CIFS"
	// FTPLogServiceTransferProtocolType File Transfer Protocol (FTP).
	FTPLogServiceTransferProtocolType LogServiceTransferProtocolType = "FTP"
	// SFTPLogServiceTransferProtocolType SSH File Transfer Protocol (SFTP).
	SFTPLogServiceTransferProtocolType LogServiceTransferProtocolType = "SFTP"
	// HTTPLogServiceTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPLogServiceTransferProtocolType LogServiceTransferProtocolType = "HTTP"
	// HTTPSLogServiceTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSLogServiceTransferProtocolType LogServiceTransferProtocolType = "HTTPS"
	// NFSLogServiceTransferProtocolType Network File System (NFS).
	NFSLogServiceTransferProtocolType LogServiceTransferProtocolType = "NFS"
	// SCPLogServiceTransferProtocolType Secure Copy Protocol (SCP).
	SCPLogServiceTransferProtocolType LogServiceTransferProtocolType = "SCP"
	// TFTPLogServiceTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPLogServiceTransferProtocolType LogServiceTransferProtocolType = "TFTP"
	// OEMLogServiceTransferProtocolType is a manufacturer-defined protocol.
	OEMLogServiceTransferProtocolType LogServiceTransferProtocolType = "OEM"
)

// LogService shall represent a log service for a Redfish implementation. When
// the 'Id' property contains 'DeviceLog', the log shall contain log entries
// that migrate with the device.
type LogService struct {
	Entity
	// AutoClearResolvedEntries shall indicate if this log service automatically
	// clears the resolved log entries found in the 'LogEntryCollection' resource.
	// If this property is not present, the value shall be assumed to be 'None'.
	//
	// Version added: v1.7.0
	AutoClearResolvedEntries AutoClearResolvedEntries
	// AutoDSTEnabled shall indicate whether the log service is configured for
	// automatic Daylight Saving Time (DST) adjustment. DST adjustment shall not
	// modify the timestamp of existing log entries.
	//
	// Version added: v1.3.0
	AutoDSTEnabled bool
	// DateTime shall contain the current date and time with UTC offset of the log
	// service.
	DateTime string
	// DateTimeLocalOffset shall contain the offset from UTC time that the
	// 'DateTime' property contains. If both 'DateTime' and 'DateTimeLocalOffset'
	// are provided in modification requests, services shall apply
	// DateTimeLocalOffset after DateTime is applied.
	DateTimeLocalOffset string
	// DiagnosticDataDetails shall contain the detailed information for the data
	// collected with the 'CollectDiagnosticData' action.
	//
	// Version added: v1.7.0
	DiagnosticDataDetails []DiagnosticDataDetails
	// Entries shall contain a link to a resource collection of type
	// 'LogEntryCollection'.
	entries string
	// LogEntryType shall contain the value for the 'EntryType' property of all
	// 'LogEntry' resources contained in the 'LogEntryCollection' resource for this
	// log service. If the service cannot determine or guarantee a single EntryType
	// value for all 'LogEntry' resources, this property shall contain the value
	// 'Multiple'.
	//
	// Version added: v1.1.0
	LogEntryType LogEntryTypes
	// LogPurposes shall contain the purposes of the log.
	//
	// Version added: v1.4.0
	LogPurposes []LogPurpose
	// MaxNumberOfRecords shall contain the maximum number of 'LogEntry' resources
	// in the 'LogEntryCollection' resource for this service.
	MaxNumberOfRecords uint
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMLogPurpose shall contain the OEM-specified purpose of the log if
	// 'LogPurposes' contains 'OEM'.
	//
	// Version added: v1.4.0
	OEMLogPurpose string
	// OverWritePolicy shall indicate the policy of the log service when the
	// 'MaxNumberOfRecords' has been reached.
	OverWritePolicy LogServiceOverWritePolicy
	// Overflow shall indicate whether the log service has overflowed and is no
	// longer able to store new logs.
	//
	// Version added: v1.4.0
	Overflow bool
	// Persistency shall indicate whether the log service is persistent across a
	// cold reset of the device.
	//
	// Version added: v1.4.0
	Persistency bool
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SyslogFilters shall describe all desired syslog messages to be logged
	// locally. If this property contains an empty array, all messages shall be
	// logged.
	//
	// Version added: v1.2.0
	SyslogFilters []SyslogFilter
	// clearLogTarget is the URL to send ClearLog requests.
	clearLogTarget string
	// collectDiagnosticDataTarget is the URL to send CollectDiagnosticData requests.
	collectDiagnosticDataTarget string
	// collectDiagnosticInfoTarget is the URL to get ActionInfo about the CollectDiagnosticData action.
	collectDiagnosticInfoTarget string
	// downloadRawLogTarget is the URL to send DownloadRawLog requests.
	downloadRawLogTarget string
	// pushDiagnosticDataTarget is the URL to send PushDiagnosticData requests.
	pushDiagnosticDataTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a LogService object from the raw JSON.
func (l *LogService) UnmarshalJSON(b []byte) error {
	type temp LogService
	type lActions struct {
		ClearLog              ActionTarget `json:"#LogService.ClearLog"`
		CollectDiagnosticData ActionTarget `json:"#LogService.CollectDiagnosticData"`
		DownloadRawLog        ActionTarget `json:"#LogService.DownloadRawLog"`
		PushDiagnosticData    ActionTarget `json:"#LogService.PushDiagnosticData"`
	}
	var tmp struct {
		temp
		Actions lActions
		Entries Link `json:"Entries"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LogService(tmp.temp)

	// Extract the links to other entities for later
	l.clearLogTarget = tmp.Actions.ClearLog.Target
	l.collectDiagnosticDataTarget = tmp.Actions.CollectDiagnosticData.Target
	l.collectDiagnosticInfoTarget = tmp.Actions.CollectDiagnosticData.ActionInfoTarget
	l.downloadRawLogTarget = tmp.Actions.DownloadRawLog.Target
	l.pushDiagnosticDataTarget = tmp.Actions.PushDiagnosticData.Target
	l.entries = tmp.Entries.String()

	// This is a read/write object, so we need to save the raw object data for later
	l.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LogService) Update() error {
	readWriteFields := []string{
		"AutoClearResolvedEntries",
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"ServiceEnabled",
	}

	return l.UpdateFromRawData(l, l.RawData, readWriteFields)
}

// GetLogService will get a LogService instance from the service.
func GetLogService(c Client, uri string) (*LogService, error) {
	return GetObject[LogService](c, uri)
}

// ListReferencedLogServices gets the collection of LogService from
// a provided reference.
func ListReferencedLogServices(c Client, link string) ([]*LogService, error) {
	return GetCollectionObjects[LogService](c, link)
}

// SupportsCollectDiagnosticData indicates if the CollectDiagnosticData action is supported.
func (l *LogService) SupportsCollectDiagnosticData() bool {
	return l.collectDiagnosticDataTarget != ""
}

// SupportsClearLog indicates if the ClearLog action is supported.
func (l *LogService) SupportsClearLog() bool {
	return l.clearLogTarget != ""
}

// FilteredEntries gets the log entries of this service with filtering applied
// (e.g. skip, top).
func (l *LogService) FilteredEntries(options ...APIFilterOption) ([]*LogEntry, error) {
	var filter APIFilter
	filter.SetFilter(options...)
	return ListReferencedLogEntrys(l.client, fmt.Sprintf("%s%s", l.entries, filter))
}

// This action shall delete all entries found in the 'LogEntryCollection'
// resource for this log service.
// logEntriesETag - This parameter shall contain the ETag of the
// 'LogEntryCollection' resource for this log service. If the client-provided
// ETag does not match the current ETag of the 'LogEntryCollection' resource
// for this log service, the service shall return the HTTP '428 Precondition
// Required' status code to reject the request.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (l *LogService) ClearLog(logEntriesETag string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	if logEntriesETag != "" {
		payload["LogEntriesETag"] = logEntriesETag
	}
	resp, taskInfo, err := PostWithTask(l.client,
		l.clearLogTarget, payload, l.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// LogServiceCollectDiagnosticDataParameters holds the parameters for the CollectDiagnosticData action.
type LogServiceCollectDiagnosticDataParameters struct {
	// DiagnosticDataType shall contain the type of diagnostic data to collect.
	DiagnosticDataType LogDiagnosticDataTypes `json:"DiagnosticDataType,omitempty"`
	// OEMDiagnosticDataType shall contain the OEM-defined type of diagnostic data
	// to collect. This parameter shall be required if 'DiagnosticDataType' is
	// 'OEM'.
	OEMDiagnosticDataType string `json:"OEMDiagnosticDataType,omitempty"`
	// Password shall contain the password to access the URI specified by the
	// 'TargetURI' parameter.
	Password string `json:"Password,omitempty"`
	// TargetDevice shall contain a link to the resource that represents the device
	// to collect diagnostic data. This parameter shall be required if
	// 'DiagnosticDataType' contains 'Device'.
	TargetDevice string `json:"TargetDevice,omitempty"`
	// TargetURI shall contain the URI to access when sending the diagnostic data.
	// If this parameter is not provided by the client, the service shall not send
	// the diagnostic data.
	TargetURI string `json:"TargetURI,omitempty"`
	// TransferProtocol shall contain the network protocol that the service uses to
	// send the diagnostic data.
	TransferProtocol LogServiceTransferProtocolType `json:"TransferProtocol,omitempty"`
	// UserName shall contain the username to access the URI specified by the
	// 'TargetURI' parameter.
	UserName string `json:"UserName,omitempty"`
}

// This action shall collect the diagnostic data for the given type. Upon
// successful completion of the action and any asynchronous processing, the
// 'Location' header in the response shall contain a URI to a resource of type
// 'LogEntry' that contains the diagnostic data. The 'AdditionalDataURI'
// property in the referenced 'LogEntry' resource shall contain the URI to
// download the diagnostic data.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (l *LogService) CollectDiagnosticData(params *LogServiceCollectDiagnosticDataParameters) (string, *TaskMonitorInfo, error) {
	if !l.SupportsCollectDiagnosticData() {
		return "", nil, errors.New("CollectDiagnosticsData not supported by this service")
	}

	resp, taskInfo, err := PostWithTask(l.GetClient(),
		l.collectDiagnosticDataTarget, params, l.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return "", taskInfo, err
	}

	// CollectDiagnosticData can return either:
	// Location pointing to a TaskMonitor if the action is async
	// Location pointing to a LogEntry if the action is sync
	if taskInfo != nil {
		return "", taskInfo, err
	}

	if location := resp.Header["Location"]; len(location) > 0 && location[0] != "" {
		return location[0], nil, nil
	}

	return "", nil, nil
}

// This action shall download a raw log for this log service.
func (l *LogService) DownloadRawLog() (*DownloadRawLogResponse, error) {
	payload := make(map[string]any)

	resp, err := l.PostWithResponse(l.downloadRawLogTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result DownloadRawLogResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// LogServicePushDiagnosticDataParameters holds the parameters for the PushDiagnosticData action.
type LogServicePushDiagnosticDataParameters struct {
	// AdditionalDataURI shall contain the URI of the diagnostic data to transfer
	// to the URI specified by the 'TargetURI' parameter.
	AdditionalDataURI string `json:"AdditionalDataURI,omitempty"`
	// Password shall contain the password to access the URI specified by the
	// 'TargetURI' parameter.
	Password string `json:"Password,omitempty"`
	// TargetURI shall contain the URI to access when sending the diagnostic data.
	TargetURI string `json:"TargetURI,omitempty"`
	// TransferProtocol shall contain the network protocol that the service uses to
	// send the diagnostic data.
	TransferProtocol TransferProtocolType `json:"TransferProtocol,omitempty"`
	// UserName shall contain the username to access the URI specified by the
	// 'TargetURI' parameter.
	UserName string `json:"UserName,omitempty"`
}

// This action shall send an existing diagnostic data to a target URI.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (l *LogService) PushDiagnosticData(params *LogServicePushDiagnosticDataParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(l.client,
		l.pushDiagnosticDataTarget, params, l.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Entries gets the Entries collection.
func (l *LogService) Entries() ([]*LogEntry, error) {
	if l.entries == "" {
		return nil, nil
	}
	return GetCollectionObjects[LogEntry](l.client, l.entries)
}

// DiagnosticDataDetails shall contain the detailed information for the data
// collected with the 'CollectDiagnosticData' action for a type of diagnostic
// data.
type DiagnosticDataDetails struct {
	// DiagnosticDataType shall contain the type of diagnostic data to collect with
	// the 'CollectDiagnosticData' action.
	//
	// Version added: v1.7.0
	DiagnosticDataType LogDiagnosticDataTypes
	// EstimatedDuration shall contain the estimated total time required to
	// generate the data with the 'CollectDiagnosticData' action. This value shall
	// not include the duration that it takes the data to transfer to a remote
	// server.
	//
	// Version added: v1.7.0
	EstimatedDuration string
	// EstimatedSizeBytes shall contain the estimated size of the data collected by
	// 'CollectDiagnosticData' action.
	//
	// Version added: v1.7.0
	EstimatedSizeBytes *int `json:",omitempty"`
	// OEMDiagnosticDataType shall contain the OEM-defined type of diagnostic data
	// to collect with the 'CollectDiagnosticData' action. This property is
	// required if 'DiagnosticDataType' is 'OEM'.
	//
	// Version added: v1.7.0
	OEMDiagnosticDataType string
}

// DownloadRawLogResponse shall contain the response for the 'DownloadRawLog'
// action.
type DownloadRawLogResponse struct {
	// DownloadURI shall contain the URI from which to download the raw log file.
	// The value of this property should not contain a URI of a Redfish resource.
	// See the 'Redfish-defined URIs and relative reference rules' clause in the
	// Redfish Specification.
	//
	// Version added: v1.9.0
	DownloadURI string
}

// SyslogFilter shall contain the filter for a syslog message. The filter shall
// describe the desired syslog message to be enabled locally.
type SyslogFilter struct {
	// LogFacilities shall contain the types of programs that can log messages. If
	// this property contains an empty array or is absent, all facilities shall be
	// indicated.
	//
	// Version added: v1.2.0
	LogFacilities []SyslogFacility
	// LowestSeverity shall contain the lowest syslog severity level that will be
	// logged. The service shall log all messages equal to or greater than the
	// value in this property. The value 'All' shall indicate all severities.
	//
	// Version added: v1.2.0
	LowestSeverity SyslogSeverity
}

// For Redfish v1.2+
// CollectDiagnosticDataActionInfo, if supported, provides the ActionInfo for a CollectDiagnosticData action.
func (l *LogService) CollectDiagnosticDataActionInfo() (*ActionInfo, error) {
	if l.collectDiagnosticInfoTarget == "" {
		return nil, errors.New("CollectDiagnosticData ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](l.GetClient(), l.collectDiagnosticInfoTarget)
}
