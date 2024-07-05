//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

type FanSpeedPercent struct {
	// The link to the resource that provides the data for this sensor.
	DataSourceURI string `json:"DataSourceUri"`
	// The name of the device.
	DeviceName string
	// The area or device to which this sensor measurement applies.
	PhysicalContext common.PhysicalContext
	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext common.PhysicalSubContext
	// The sensor value.
	Reading float32
	// The rotational speed.
	SpeedRPM float32
}

// PowerSupplyUnitMetrics shall be used to represent the metrics
// of a power supply unit for a Redfish implementation.
type PowerSupplyUnitMetrics struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The total energy, in kilowatt-hours
	// that represents the Total ElectricalContext sensor
	// when multiple energy sensors exist.
	EnergykWh SensorEnergykWhExcerpt
	// The fan speed (percent) for this power supply.
	// Deprecated: (v1.1) This property has been deprecated
	// in favor of FanSpeedsPercent to support multiple fans
	// within a power supply.
	FanSpeedPercent FanSpeedPercent
	// Fan speeds (percent).
	FanSpeedsPercent []FanSpeedPercent
	// The frequency (Hz) for this power supply.
	FrequencyHz SensorExcerpt
	// The input current (A) for this power supply.
	InputCurrentAmps SensorVoltageExcerpt
	// The input power (W) for this power supply.
	InputPowerWatts SensorPowerExcerpt
	// The input voltage (V) for this power supply.
	InputVoltage SensorVoltageExcerpt
	// The total power output (W) for this power supply.
	OutputPowerWatts SensorPowerExcerpt
	// The output currents (A) for this power supply.
	RailCurrentAmps []SensorVoltageExcerpt
	// The output power readings (W) for this power supply.
	RailPowerWatts []SensorPowerExcerpt
	// The output voltages (V) for this power supply.
	RailVoltage []SensorVoltageExcerpt
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// The temperature (C) for this power supply.
	TemperatureCelsius SensorExcerpt
	Oem                json.RawMessage

	// Actions section
	// // This action resets the summary metrics related to this equipment.
	resetMetricsTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a PowerSupplyMetrics object from the raw JSON.
func (metrics *PowerSupplyUnitMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyUnitMetrics

	type actions struct {
		ResetMetrics common.ActionTarget `json:"#PowerSupplyMetrics.ResetMetrics"`
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
	*metrics = PowerSupplyUnitMetrics(t.temp)
	metrics.resetMetricsTarget = t.Actions.ResetMetrics.Target
	metrics.OemActions = t.Actions.Oem

	return nil
}

// GetPowerSupplyUnitMetrics will get a PowerSupplyMetrics instance from the Redfish service.
func GetPowerSupplyUnitMetrics(c common.Client, uri string) (*PowerSupplyUnitMetrics, error) {
	return common.GetObject[PowerSupplyUnitMetrics](c, uri)
}

// This action resets the summary metrics related to this equipment.
func (metrics *PowerSupplyUnitMetrics) ResetMetrics() error {
	if metrics.resetMetricsTarget == "" {
		return fmt.Errorf("ResetMetrics is not supported") //nolint:golint
	}

	return metrics.Post(metrics.resetMetricsTarget, nil)
}
