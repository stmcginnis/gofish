//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/stmcginnis/gofish/schemas"
)

// RemoteServicesStatus is the overall status of the iDRAC remote services API.
type RemoteServicesStatus string

const (
	// ReadyRemoteServicesStatus indicates the remote services API is ready to
	// accept configuration and update requests.
	ReadyRemoteServicesStatus RemoteServicesStatus = "Ready"
	// NotReadyRemoteServicesStatus indicates the remote services API is not ready.
	NotReadyRemoteServicesStatus RemoteServicesStatus = "NotReady"
	// InUseRemoteServicesStatus indicates the Lifecycle Controller is in use.
	InUseRemoteServicesStatus RemoteServicesStatus = "InUse"
	// ReloadingDataRemoteServicesStatus indicates the Lifecycle Controller is
	// reloading data.
	ReloadingDataRemoteServicesStatus RemoteServicesStatus = "ReloadingData"
	// DisabledRemoteServicesStatus indicates the Lifecycle Controller is disabled.
	DisabledRemoteServicesStatus RemoteServicesStatus = "Disabled"
	// InRecoveryRemoteServicesStatus indicates the Lifecycle Controller is in
	// recovery mode.
	InRecoveryRemoteServicesStatus RemoteServicesStatus = "InRecovery"
)

// LCStatus is the status of the Lifecycle Controller, including the Data
// Manager status.
type LCStatus string

const (
	// ReadyLCStatus indicates the Lifecycle Controller is ready.
	ReadyLCStatus LCStatus = "Ready"
	// InUseLCStatus indicates the Lifecycle Controller is in use.
	InUseLCStatus LCStatus = "InUse"
	// ReloadingDataLCStatus indicates the Lifecycle Controller is reloading data.
	ReloadingDataLCStatus LCStatus = "ReloadingData"
	// NotInitializedLCStatus indicates the Lifecycle Controller is not initialized.
	NotInitializedLCStatus LCStatus = "NotInitialized"
	// DisabledLCStatus indicates the Lifecycle Controller is disabled.
	DisabledLCStatus LCStatus = "Disabled"
	// InRecoveryLCStatus indicates the Lifecycle Controller is in recovery mode.
	InRecoveryLCStatus LCStatus = "InRecovery"
)

// RTStatus is the status of real-time configuration.
type RTStatus string

const (
	// ReadyRTStatus indicates real-time configuration is ready.
	ReadyRTStatus RTStatus = "Ready"
	// NotReadyRTStatus indicates real-time configuration is not ready.
	NotReadyRTStatus RTStatus = "NotReady"
)

// ServerStatus is the status of the host system.
type ServerStatus string

const (
	// PoweredOffServerStatus indicates the server is powered off.
	PoweredOffServerStatus ServerStatus = "PoweredOff"
	// InPOSTServerStatus indicates the server is in POST.
	InPOSTServerStatus ServerStatus = "InPOST"
	// OutOfPOSTServerStatus indicates the server has completed POST.
	OutOfPOSTServerStatus ServerStatus = "OutOfPOST"
	// CollectingSystemInventoryServerStatus indicates the server is collecting
	// system inventory.
	CollectingSystemInventoryServerStatus ServerStatus = "CollectingSystemInventory"
	// AutomatedTaskExecutionServerStatus indicates the server is executing an
	// automated task.
	AutomatedTaskExecutionServerStatus ServerStatus = "AutomatedTaskExecution"
	// LCUSCServerStatus indicates the server is in the Lifecycle Controller
	// Unified Server Configurator.
	LCUSCServerStatus ServerStatus = "LCUSC"
	// EnteredF2SetupServerStatus indicates the server has entered the F2 setup
	// menu.
	EnteredF2SetupServerStatus ServerStatus = "EnteredF2Setup"
	// EnteredF11BootManagerServerStatus indicates the server has entered the F11
	// Boot Manager menu.
	EnteredF11BootManagerServerStatus ServerStatus = "EnteredF11BootManager"
	// HaltedF1F2PromptServerStatus indicates the server has halted at the F1/F2
	// error prompt because of a POST error.
	HaltedF1F2PromptServerStatus ServerStatus = "HaltedF1/F2Prompt"
	// HaltedF1F2F11PromptServerStatus indicates the server has halted at the
	// F1/F2/F11 prompt because there are no bootable devices available.
	HaltedF1F2F11PromptServerStatus ServerStatus = "HaltedF1/F2/F11Prompt"
)

// ServiceStatus is the status of an iDRAC service, such as SEKM or telemetry.
type ServiceStatus string

const (
	// ReadyServiceStatus indicates the service is ready.
	ReadyServiceStatus ServiceStatus = "Ready"
	// NotReadyServiceStatus indicates the service is not ready.
	NotReadyServiceStatus ServiceStatus = "NotReady"
)

// RemoteServicesAPIStatus is the result of the GetRemoteServicesAPIStatus
// action. Not all fields are reported by every iDRAC version; fields that are
// not reported are left empty.
type RemoteServicesAPIStatus struct {
	// Status is the overall remote services API status. It reflects the host
	// system and Lifecycle Controller status, but not RTStatus.
	Status RemoteServicesStatus
	// LCStatus is the Lifecycle Controller status, including the Data Manager
	// status.
	LCStatus LCStatus
	// RTStatus is the real-time configuration status.
	RTStatus RTStatus
	// ServerStatus is the host system status.
	ServerStatus ServerStatus
	// SEKMServiceStatus is the Secure Enterprise Key Manager service status.
	SEKMServiceStatus ServiceStatus
	// TelemetryServiceStatus is the telemetry service status.
	TelemetryServiceStatus ServiceStatus
	// ExtendedInfo contains the messages returned with the status.
	ExtendedInfo []schemas.Message `json:"@Message.ExtendedInfo"`
}

// IsReady reports whether the iDRAC is ready to accept configuration and
// update requests, based on the overall Status.
func (s *RemoteServicesAPIStatus) IsReady() bool {
	return s.Status == ReadyRemoteServicesStatus
}

// LCService represents Dell's DellLCService OEM resource, which provides
// actions to support Lifecycle Controller functionality.
type LCService struct {
	schemas.Entity

	getRemoteServicesAPIStatusTarget string
}

// UnmarshalJSON unmarshals an LCService and its action targets.
func (s *LCService) UnmarshalJSON(b []byte) error {
	type temp LCService
	var payload struct {
		temp
		Actions struct {
			GetRemoteServicesAPIStatus schemas.ActionTarget `json:"#DellLCService.GetRemoteServicesAPIStatus"`
		}
	}

	if err := json.Unmarshal(b, &payload); err != nil {
		return err
	}

	*s = LCService(payload.temp)
	s.getRemoteServicesAPIStatusTarget = payload.Actions.GetRemoteServicesAPIStatus.Target
	return nil
}

// GetRemoteServicesAPIStatus gets the overall remote services API status,
// including the host system, Lifecycle Controller, and real-time status. Use
// it to check whether the iDRAC is ready before submitting configuration or
// update jobs.
func (s *LCService) GetRemoteServicesAPIStatus() (*RemoteServicesAPIStatus, error) {
	if s.getRemoteServicesAPIStatusTarget == "" {
		return nil, errors.New("GetRemoteServicesAPIStatus is not supported by this service")
	}

	resp, err := s.PostWithResponse(s.getRemoteServicesAPIStatusTarget, struct{}{})
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote services API status: %w", err)
	}

	var status RemoteServicesAPIStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, fmt.Errorf("failed to decode remote services API status: %w", err)
	}

	return &status, nil
}

// GetLCService gets an LCService from the service.
func GetLCService(c schemas.Client, uri string) (*LCService, error) {
	return schemas.GetObject[LCService](c, uri)
}

// LCService gets the Dell Lifecycle Controller service linked by this manager.
func (m *Manager) LCService() (*LCService, error) {
	if m.lcService == "" {
		return nil, errors.New("lifecycle controller service is not supported by this manager")
	}

	return GetLCService(m.GetClient(), m.lcService)
}

// GetRemoteServicesAPIStatus gets the remote services API status using the
// Lifecycle Controller service linked by this manager.
func (m *Manager) GetRemoteServicesAPIStatus() (*RemoteServicesAPIStatus, error) {
	service, err := m.LCService()
	if err != nil {
		return nil, err
	}

	return service.GetRemoteServicesAPIStatus()
}
