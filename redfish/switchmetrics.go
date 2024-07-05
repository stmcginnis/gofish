//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"

	"github.com/stmcginnis/gofish/common"
)

// SwitchMetricCurrentPeriod shall describe the memory metrics since the last reset or ClearCurrentPeriod action for a switch.
type SwitchMetricCurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of memory since reset.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors of memory since reset.
	UncorrectableECCErrorCount int
}

// InternalMemoryMetrics shall contain properties that describe the memory metrics for a switch.
type InternalMemoryMetrics struct {
	// CurrentPeriod shall contain properties that describe the metrics for the current period of memory for this
	// switch.
	CurrentPeriod SwitchMetricCurrentPeriod
	// LifeTime shall contain properties that describe the metrics for the lifetime of memory for this switch.
	LifeTime SwitchMetricLifeTime
}

// SwitchMetricLifeTime shall describe the memory metrics since manufacturing for a switch.
type SwitchMetricLifeTime struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors for the lifetime of memory.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors for the lifetime of memory.
	UncorrectableECCErrorCount int
}

// SwitchMetrics shall represent the metrics for a switch device in a Redfish implementation.
type SwitchMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// InternalMemoryMetrics shall contain properties that describe the memory metrics for a switch.
	InternalMemoryMetrics InternalMemoryMetrics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this switch.
	PCIeErrors PCIeErrors

	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a SwitchMetrics object from the raw JSON.
func (switchmetrics *SwitchMetrics) UnmarshalJSON(b []byte) error {
	type temp SwitchMetrics
	type Actions struct {
		ClearCurrentPeriod common.ActionTarget `json:"#SwitchMetrics.ClearCurrentPeriod"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*switchmetrics = SwitchMetrics(t.temp)

	// Extract the links to other entities for later
	switchmetrics.clearCurrentPeriodTarget = t.Actions.ClearCurrentPeriod.Target

	return nil
}

// ClearCurrentPeriod sets the CurrentPeriod property's values to 0.
func (switchmetrics *SwitchMetrics) ClearCurrentPeriod() error {
	if switchmetrics.clearCurrentPeriodTarget == "" {
		return errors.New("ClearCurrentPeriod is not supported by this system")
	}
	return switchmetrics.Post(switchmetrics.clearCurrentPeriodTarget, nil)
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
