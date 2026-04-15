//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type Manager struct {
	redfish.Manager

	importSystemConfigTarget string
}

type ISCExecutionMode string

const (
	DefaultISCExecutionMode       ISCExecutionMode = "Default"
	DeployOnSledISCExecutionMode  ISCExecutionMode = "DeployOnSledInsert"
	InstantDeployISCExecutionMode ISCExecutionMode = "InstantDeploy"
)

type ISCHostPowerState string

const (
	OnISCHostPowerState  ISCHostPowerState = "On"
	OffISCHostPowerState ISCHostPowerState = "Off"
)

type ISCIgnoreCertificateWarning string

const (
	DisabledISCIgnoreCertificateWarning ISCIgnoreCertificateWarning = "Disabled"
	EnabledISCIgnoreCertificateWarning  ISCIgnoreCertificateWarning = "Enabled"
)

type ISCShareType string

const (
	LocalISCShareType ISCShareType = "LOCAL"
	NFSISCShareType   ISCShareType = "NFS"
	CIFSISCShareType  ISCShareType = "CIFS"
	HTTPISCShareType  ISCShareType = "HTTP"
	HTTPSISCShareType ISCShareType = "HTTPS"
)

type ISCShutdownType string

const (
	GracefulISCShutdownType ISCShutdownType = "Graceful"
	ForcedISCShutdownType   ISCShutdownType = "Forced"
	NoRebootISCShutdownType ISCShutdownType = "NoReboot"
)

type ShareParameters struct {
	// IP address for the remote share.
	IPAddress string `json:",omitempty"`
	// Specifies if certificate warning should be ignored when HTTPS is used. If IgnoreCertWarning is On, warnings are ignored. Default is 2 (On).
	IgnoreCertificateWarning ISCIgnoreCertificateWarning `json:",omitempty"`
	// Name of the CIFS share or full path to the NFS share. Optional for HTTP/HTTPS share, this may be treated as the path of the directory containing the file.
	ShareName string `json:",omitempty"`
	// File name on share
	FileName string `json:",omitempty"`
	// Type of the network share.
	ShareType ISCShareType `json:",omitempty"`
	// User name for the remote share. This parameter must be provided for CIFS.
	UserName string `json:",omitempty"`
	// Workgroup for the CIFS share - optional.
	Workgroup string `json:",omitempty"`
	// Specify a device's Fully Qualified Device Descriptor (FQDD) in the Target parameter to return attributes for that specific device. For example, using NIC returns attributes for all NICs, while NIC.Slot.1-1-1 limits the result to the NIC in slot 1, port 1.
	//
	// The following generic targets can also be used: "ALL", "IDRAC", "BIOS", "NIC", "RAID", "FC", "InfiniBand", "SupportAssist", "EventFilters", "System", "LifecycleController", "AHCI", "PCIeSSD"
	Target string `json:",omitempty"`
}

// Body of the POST request for ImportSystemConfiguration Action
type ImportSystemConfigurationBody struct {
	ExecutionMode ISCExecutionMode `json:",omitempty"`
	// Power state node should be in after applying the configuration.
	HostPowerState ISCHostPowerState `json:",omitempty"`
	// Buffer content to perform import. Required only for LOCAL and not required for CIFS, NFS, HTTP, or HTTPS.
	ImportBuffer    string `json:",omitempty"`
	ShareParameters ShareParameters
	// Shutdown type when applying configuration. NoReboot will queue the job and wait until the next boot.
	ShutdownType ISCShutdownType `json:",omitempty"`
}

// Unmarshals a redfish.Manager into a dell.Manager
func FromManager(manager *redfish.Manager) (*Manager, error) {
	m := Manager{
		Manager: *manager,
	}

	var t struct {
		ImportSystemConfiguration struct {
			Target string `json:"target,omitempty"`
		} `json:"#OemManager.ImportSystemConfiguration,omitempty"`
	}

	err := json.Unmarshal(manager.OemActions, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Dell OEM manager actions: %w", err)
	}

	m.importSystemConfigTarget = t.ImportSystemConfiguration.Target

	return &m, nil
}

// validateImportSystemConfigurationBody validates required fields in ImportSystemConfigurationBody
func validateImportSystemConfigurationBody(b *ImportSystemConfigurationBody) error {
	if b.ShareParameters.Target == "" {
		return errors.New("ShareParameters.Target is required")
	}

	// Validate ShareType is one of the allowed values
	validShareTypes := []ISCShareType{
		LocalISCShareType,
		NFSISCShareType,
		CIFSISCShareType,
		HTTPISCShareType,
		HTTPSISCShareType,
	}

	valid := false
	for _, validType := range validShareTypes {
		if b.ShareParameters.ShareType == validType {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid ShareType: %s, must be one of LOCAL, NFS, CIFS, HTTP, HTTPS", b.ShareParameters.ShareType)
	}

	// For non-LOCAL shares, IPAddress is required
	if b.ShareParameters.ShareType != LocalISCShareType && b.ShareParameters.IPAddress == "" {
		return errors.New("IPAddress is required for non-LOCAL shares")
	}

	// Username is required for CIFS shares
	if b.ShareParameters.ShareType == CIFSISCShareType && b.ShareParameters.UserName == "" {
		return errors.New("UserName is required for CIFS shares")
	}

	// For LOCAL shares, ImportBuffer is required
	if b.ShareParameters.ShareType == LocalISCShareType && b.ImportBuffer == "" {
		return errors.New("ImportBuffer is required for LOCAL shares")
	}

	// For non-LOCAL shares, ShareName is required
	if b.ShareParameters.ShareType != LocalISCShareType && b.ShareParameters.ShareName == "" {
		return errors.New("ShareName is required for non-LOCAL shares")
	}

	return nil
}

// Import a system configuration in JSON format.
//
// This can be used to set BIOS, iDRAC, and device settings automatically.
func (m *Manager) ImportSystemConfiguration(b *ImportSystemConfigurationBody) (*redfish.Task, error) {
	if err := validateImportSystemConfigurationBody(b); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	if m.importSystemConfigTarget == "" {
		return nil, errors.New("import system config is not supported by this system")
	}

	res, err := m.PostWithResponse(m.importSystemConfigTarget, b)
	defer common.DeferredCleanupHTTPResponse(res)
	if err != nil {
		return nil, err
	}

	return redfish.GetTask(m.GetClient(), res.Header.Get("Location"))
}
