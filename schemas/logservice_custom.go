//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
)

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

// For Redfish v1.2+
// CollectDiagnosticDataActionInfo, if supported, provides the ActionInfo for a CollectDiagnosticData action.
func (l *LogService) CollectDiagnosticDataActionInfo() (*ActionInfo, error) {
	if l.collectDiagnosticInfoTarget == "" {
		return nil, errors.New("CollectDiagnosticData ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](l.GetClient(), l.collectDiagnosticInfoTarget)
}
