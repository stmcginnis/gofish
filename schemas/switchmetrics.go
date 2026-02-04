//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/SwitchMetrics.v1_1_0.json
// 2025.3 - #SwitchMetrics.v1_1_0.SwitchMetrics

package schemas

import (
	"encoding/json"
)

// SwitchMetrics shall represent the metrics for a switch device in a Redfish
// implementation.
type SwitchMetrics struct {
	Entity
	// InternalMemoryMetrics shall contain properties that describe the memory
	// metrics for a switch.
	InternalMemoryMetrics InternalMemoryMetrics
	// LifetimeStartDateTime shall contain the date and time when the switch
	// started accumulating data for the 'LifeTime' property. This might contain
	// the same value as the production date of the switch.
	//
	// Version added: v1.1.0
	LifetimeStartDateTime string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this switch.
	PCIeErrors PCIeErrors
	// clearCurrentPeriodTarget is the URL to send ClearCurrentPeriod requests.
	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a SwitchMetrics object from the raw JSON.
func (s *SwitchMetrics) UnmarshalJSON(b []byte) error {
	type temp SwitchMetrics
	type sActions struct {
		ClearCurrentPeriod ActionTarget `json:"#SwitchMetrics.ClearCurrentPeriod"`
	}
	var tmp struct {
		temp
		Actions sActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SwitchMetrics(tmp.temp)

	// Extract the links to other entities for later
	s.clearCurrentPeriodTarget = tmp.Actions.ClearCurrentPeriod.Target

	return nil
}

// GetSwitchMetrics will get a SwitchMetrics instance from the service.
func GetSwitchMetrics(c Client, uri string) (*SwitchMetrics, error) {
	return GetObject[SwitchMetrics](c, uri)
}

// ListReferencedSwitchMetricss gets the collection of SwitchMetrics from
// a provided reference.
func ListReferencedSwitchMetricss(c Client, link string) ([]*SwitchMetrics, error) {
	return GetCollectionObjects[SwitchMetrics](c, link)
}

// This action shall set the 'CurrentPeriod' property's values to 0.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *SwitchMetrics) ClearCurrentPeriod() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.clearCurrentPeriodTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// SwitchMetricsCurrentPeriod shall describe the memory metrics since the last reset or
// 'ClearCurrentPeriod' action for a switch.
type SwitchMetricsCurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of
	// memory since reset.
	CorrectableECCErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// of memory since reset.
	UncorrectableECCErrorCount *int `json:",omitempty"`
}

// InternalMemoryMetrics shall contain properties that describe the memory
// metrics for a switch.
type InternalMemoryMetrics struct {
	// CurrentPeriod shall contain properties that describe the metrics for the
	// current period of memory for this switch.
	CurrentPeriod SwitchMetricsCurrentPeriod
	// LifeTime shall contain properties that describe the metrics for the lifetime
	// of memory for this switch.
	LifeTime SwitchMetricsLifeTime
}

// SwitchMetricsLifeTime shall describe the memory metrics since manufacturing for a switch.
type SwitchMetricsLifeTime struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors for
	// the lifetime of memory.
	CorrectableECCErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// for the lifetime of memory.
	UncorrectableECCErrorCount *int `json:",omitempty"`
}
