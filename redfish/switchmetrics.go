//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #SwitchMetrics.v1_1_0.SwitchMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SwitchMetrics shall represent the metrics for a switch device in a Redfish
// implementation.
type SwitchMetrics struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this switch.
	PCIeErrors PCIeErrors
	// clearCurrentPeriodTarget is the URL to send ClearCurrentPeriod requests.
	clearCurrentPeriodTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SwitchMetrics object from the raw JSON.
func (s *SwitchMetrics) UnmarshalJSON(b []byte) error {
	type temp SwitchMetrics
	type sActions struct {
		ClearCurrentPeriod common.ActionTarget `json:"#SwitchMetrics.ClearCurrentPeriod"`
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

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SwitchMetrics) Update() error {
	readWriteFields := []string{
		"InternalMemoryMetrics",
		"PCIeErrors",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSwitchMetrics will get a SwitchMetrics instance from the service.
func GetSwitchMetrics(c common.Client, uri string) (*SwitchMetrics, error) {
	return common.GetObject[SwitchMetrics](c, uri)
}

// ListReferencedSwitchMetricss gets the collection of SwitchMetrics from
// a provided reference.
func ListReferencedSwitchMetricss(c common.Client, link string) ([]*SwitchMetrics, error) {
	return common.GetCollectionObjects[SwitchMetrics](c, link)
}

// ClearCurrentPeriod shall set the 'CurrentPeriod' property's values to 0.
func (s *SwitchMetrics) ClearCurrentPeriod() error {
	payload := make(map[string]any)
	return s.Post(s.clearCurrentPeriodTarget, payload)
}

// InternalMemoryMetrics shall contain properties that describe the memory
// metrics for a switch.
type InternalMemoryMetrics struct {
	// CurrentPeriod shall contain properties that describe the metrics for the
	// current period of memory for this switch.
	CurrentPeriod CurrentPeriod
	// LifeTime shall contain properties that describe the metrics for the lifetime
	// of memory for this switch.
	LifeTime LifeTime
}
