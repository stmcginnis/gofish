//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

// PowerDistributionMetrics shall be used to represent
// the metrics of a power distribution component or unit for a Redfish implementation.
type PowerDistributionMetrics struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The absolute (volumetric) humidity sensor reading,
	// in grams/cubic meter units.
	AbsoluteHumidity SensorExcerpt
	// The total energy, in kilowatt-hours
	// that represents the Total ElectricalContext sensor
	// when multiple energy sensors exist.
	EnergykWh SensorEnergykWhExcerpt
	// The humidity, in percent units
	HumidityPercent SensorExcerpt
	// The power load, in percent units, for this device
	// that represents the Total ElectricalContext for this device.
	PowerLoadPercent SensorExcerpt
	// The total power, in watt units
	// that represents the Total ElectricalContext sensor
	// when multiple power sensors exist.
	PowerWatts SensorPowerExcerpt
	// The temperature, in degrees Celsius units.
	TemperatureCelsius SensorExcerpt
	Oem                json.RawMessage

	// Actions section
	// // This action resets the summary metrics related to this equipment.
	resetMetricsTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a PowerDistributionMetrics object from the raw JSON.
func (metrics *PowerDistributionMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerDistributionMetrics

	type actions struct {
		ResetMetrics common.ActionTarget `json:"#PowerDistributionMetrics.ResetMetrics"`
		Oem          json.RawMessage     // OEM actions will be stored here
	}
	var t struct {
		temp
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*metrics = PowerDistributionMetrics(t.temp)
	metrics.resetMetricsTarget = t.Actions.ResetMetrics.Target
	metrics.OemActions = t.Actions.Oem

	return nil
}

// GetPowerDistributionMetrics will get a PowerDistributionMetrics instance from the Redfish service.
func GetPowerDistributionMetrics(c common.Client, uri string) (*PowerDistributionMetrics, error) {
	return common.GetObject[PowerDistributionMetrics](c, uri)
}

// This action shall reset any time intervals or counted values for this equipment.
func (metrics *PowerDistributionMetrics) ResetMetrics() error {
	if metrics.resetMetricsTarget == "" {
		return fmt.Errorf("ResetMetrics is not supported") //nolint:golint
	}

	return metrics.Post(metrics.resetMetricsTarget, nil)
}
