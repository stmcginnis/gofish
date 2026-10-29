//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/stmcginnis/gofish/schemas"
)

// ESCExportFormat is the format of an exported system configuration.
type ESCExportFormat string

const (
	// XMLESCExportFormat exports the system configuration as XML.
	XMLESCExportFormat ESCExportFormat = "XML"
	// JSONESCExportFormat exports the system configuration as JSON.
	JSONESCExportFormat ESCExportFormat = "JSON"
)

// ESCExportUse is the intended use of an exported system configuration.
type ESCExportUse string

const (
	// DefaultESCExportUse exports the configuration for general use.
	DefaultESCExportUse ESCExportUse = "Default"
	// CloneESCExportUse exports the configuration for cloning to another
	// system.
	CloneESCExportUse ESCExportUse = "Clone"
	// ReplaceESCExportUse exports the configuration for replacing this
	// system's configuration.
	ReplaceESCExportUse ESCExportUse = "Replace"
)

// ESCIncludeInExport selects additional content for an exported system
// configuration.
type ESCIncludeInExport string

const (
	// DefaultESCIncludeInExport exports the default set of attributes.
	DefaultESCIncludeInExport ESCIncludeInExport = "Default"
	// IncludeReadOnlyESCIncludeInExport also exports read-only attributes.
	IncludeReadOnlyESCIncludeInExport ESCIncludeInExport = "IncludeReadOnly"
	// IncludePasswordHashValuesESCIncludeInExport also exports password hash
	// values.
	IncludePasswordHashValuesESCIncludeInExport ESCIncludeInExport = "IncludePasswordHashValues"
	// IncludeCustomTelemetryESCIncludeInExport also exports custom telemetry
	// configuration.
	IncludeCustomTelemetryESCIncludeInExport ESCIncludeInExport = "IncludeCustomTelemetry"
)

// ExportSystemConfigurationBody is the body of the POST request for the
// ExportSystemConfiguration action.
type ExportSystemConfigurationBody struct {
	// Format of the exported configuration. Default is XML.
	ExportFormat ESCExportFormat `json:",omitempty"`
	// Intended use of the exported configuration. Default is Default.
	ExportUse ESCExportUse `json:",omitempty"`
	// Additional content to include in the exported configuration.
	IncludeInExport ESCIncludeInExport `json:",omitempty"`
	// Where to export to and which components to export. ShareParameters.Target
	// is required. When ShareType is empty or LOCAL, the configuration is
	// returned to the caller; use WaitForExportedSystemConfiguration to get it.
	ShareParameters ShareParameters
}

// validateExportSystemConfigurationBody validates required fields in ExportSystemConfigurationBody
func validateExportSystemConfigurationBody(b *ExportSystemConfigurationBody) error {
	if b == nil {
		return errors.New("request body is required")
	}

	if b.ShareParameters.Target == "" {
		return errors.New("ShareParameters.Target is required")
	}

	switch b.ExportFormat {
	case "", XMLESCExportFormat, JSONESCExportFormat:
	default:
		return fmt.Errorf("invalid ExportFormat: %s, must be one of XML, JSON", b.ExportFormat)
	}

	shareType := b.ShareParameters.ShareType
	switch shareType {
	case "", LocalISCShareType:
		return nil
	case NFSISCShareType, CIFSISCShareType, HTTPISCShareType, HTTPSISCShareType:
	default:
		return fmt.Errorf("invalid ShareType: %s, must be one of LOCAL, NFS, CIFS, HTTP, HTTPS", shareType)
	}

	if b.ShareParameters.IPAddress == "" {
		return errors.New("IPAddress is required for non-LOCAL shares")
	}
	if b.ShareParameters.ShareName == "" {
		return errors.New("ShareName is required for non-LOCAL shares")
	}
	if shareType == CIFSISCShareType && b.ShareParameters.UserName == "" {
		return errors.New("UserName is required for CIFS shares")
	}

	return nil
}

// ExportSystemConfiguration starts an export of the system configuration
// (Server Configuration Profile). The export runs asynchronously; the returned
// task monitor can be passed to WaitForExportedSystemConfiguration for a LOCAL
// export, or to schemas.WaitForTaskMonitor to track an export to a network
// share.
func (m *Manager) ExportSystemConfiguration(b *ExportSystemConfigurationBody) (*schemas.TaskMonitorInfo, error) {
	if err := validateExportSystemConfigurationBody(b); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	if m.exportSystemConfigTarget == "" {
		return nil, errors.New("export system configuration is not supported by this system")
	}

	resp, taskInfo, err := schemas.PostWithTask(
		m.GetClient(), m.exportSystemConfigTarget, b, m.Headers(), false,
	)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to export system configuration: %w", err)
	}
	if taskInfo == nil {
		return nil, errors.New("export system configuration response did not start a task")
	}
	if taskInfo.TaskMonitor == "" {
		return nil, errors.New("export system configuration response is missing Location header")
	}

	return taskInfo, nil
}

// WaitForExportedSystemConfiguration waits for a LOCAL system configuration
// export started by ExportSystemConfiguration to complete and returns the
// exported configuration in the requested format. A pollRate of zero uses the
// default of schemas.WaitForTaskMonitor.
//
// When the export completes, iDRAC returns the configuration itself as the
// body of the task monitor response instead of a Task.
func (m *Manager) WaitForExportedSystemConfiguration(
	ctx context.Context, taskInfo *schemas.TaskMonitorInfo, pollRate time.Duration,
) ([]byte, error) {
	resp, err := schemas.WaitForTaskMonitor(ctx, m.GetClient(), pollRate, taskInfo, nil)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for system configuration export: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read exported system configuration: %w", err)
	}

	if isSystemConfiguration(body) {
		return body, nil
	}

	// A failed export, or an export to a network share, completes with a Task.
	var task schemas.Task
	if err := json.Unmarshal(body, &task); err == nil && task.TaskState != "" {
		return nil, fmt.Errorf("system configuration export did not return a configuration: task %s %s%s",
			task.ID, task.TaskState, taskMessages(&task))
	}

	return nil, errors.New("system configuration export response does not contain a system configuration")
}

// isSystemConfiguration reports whether b is an exported system configuration
// in XML or JSON format.
func isSystemConfiguration(b []byte) bool {
	trimmed := bytes.TrimSpace(b)
	if bytes.HasPrefix(trimmed, []byte("<")) {
		return bytes.Contains(trimmed, []byte("<SystemConfiguration"))
	}

	var scp struct {
		SystemConfiguration json.RawMessage
	}
	return json.Unmarshal(trimmed, &scp) == nil &&
		len(scp.SystemConfiguration) > 0 && string(scp.SystemConfiguration) != "null"
}

// taskMessages formats the messages of a task for an error.
func taskMessages(task *schemas.Task) string {
	var buf bytes.Buffer
	for i := range task.Messages {
		if i == 0 {
			buf.WriteString(": ")
		} else {
			buf.WriteString("; ")
		}
		buf.WriteString(task.Messages[i].Message)
	}
	return buf.String()
}
