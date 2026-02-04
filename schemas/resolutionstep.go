//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ResolutionStep.v1_0_1.json
// 2023.3 - #ResolutionStep.v1_0_1

package schemas

import (
	"encoding/json"
)

type ResolutionType string

const (
	// ContactVendorResolutionType Contact vendor for service.
	ContactVendorResolutionType ResolutionType = "ContactVendor"
	// ReplaceComponentResolutionType Replace a component.
	ReplaceComponentResolutionType ResolutionType = "ReplaceComponent"
	// FirmwareUpdateResolutionType Perform a firmware update operation.
	FirmwareUpdateResolutionType ResolutionType = "FirmwareUpdate"
	// ResetResolutionType Perform a reset operation.
	ResetResolutionType ResolutionType = "Reset"
	// PowerCycleResolutionType Perform a power cycle operation.
	PowerCycleResolutionType ResolutionType = "PowerCycle"
	// ResetToDefaultsResolutionType Reset the settings to factory defaults.
	ResetToDefaultsResolutionType ResolutionType = "ResetToDefaults"
	// CollectDiagnosticDataResolutionType Collect diagnostic data.
	CollectDiagnosticDataResolutionType ResolutionType = "CollectDiagnosticData"
	// OEMResolutionType Perform an OEM-defined resolution step.
	OEMResolutionType ResolutionType = "OEM"
)

// ResolutionStep shall describe a recommended step of the service-defined
// resolution. The set of recommended steps are used to resolve the cause of a
// log entry, an event, a condition, or an error message.
type ResolutionStep struct {
	// ActionParameters shall contain the parameters of the action URI for a
	// resolution step.
	ActionParameters []ActionInfoParameter
	// ActionURI shall contain the action URI for a resolution step.
	ActionURI string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Priority shall contain the priority in the set of resolution steps. The
	// value '0' shall indicate the highest priority. Increasing values shall
	// represent decreasing priority. If two or more resolution steps have the same
	// priority, the execution order of the resolution steps shall be in array
	// order. If a resolution step does not have a priority assignment, the default
	// is 0. The priority is used to determine the execution order of the
	// resolution steps.
	Priority *uint `json:",omitempty"`
	// ResolutionType shall contain the type of the resolution step.
	ResolutionType ResolutionType
	// RetryCount shall contain the number of the retries for a resolution step.
	RetryCount *uint `json:",omitempty"`
	// RetryIntervalSeconds shall contain the interval, in seconds, between the
	// retries for a resolution step.
	RetryIntervalSeconds *uint `json:",omitempty"`
	// TargetComponentURI shall contain the target URI of the component for a
	// resolution step. This property shall be present if the 'ActionURI' property
	// is not supported.
	TargetComponentURI string
}
