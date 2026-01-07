//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.1 - #PowerSupplyMetrics.v1_1_2.PowerSupplyMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerSupplyMetrics shall be used to represent the metrics of a power supply
// unit for a Redfish implementation.
type PowerSupplyMetrics struct {
	common.Entity
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// unit that represents the 'Total' 'ElectricalContext' sensor when multiple
	// energy sensors exist. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FanSpeedPercent shall contain the fan speed, in percent units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Percent'.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of 'FanSpeedsPercent' to support
	// multiple fans within a power supply.
	FanSpeedPercent SensorFanExcerpt
	// FanSpeedsPercent shall contain the fan speeds, in percent units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Percent'.
	//
	// Version added: v1.1.0
	FanSpeedsPercent []SensorFanArrayExcerpt
	// FanSpeedsPercent@odata.count
	FanSpeedsPercentCount int `json:"FanSpeedsPercent@odata.count"`
	// FrequencyHz shall contain the frequency, in hertz units, for this power
	// supply.
	FrequencyHz SensorExcerpt
	// InputCurrentAmps shall contain the input current, in ampere units, for this
	// power supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'.
	InputCurrentAmps SensorCurrentExcerpt
	// InputPowerWatts shall contain the input power, in watt units, for this power
	// supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	InputPowerWatts SensorPowerExcerpt
	// InputVoltage shall contain the input voltage, in volt units, for this power
	// supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'.
	InputVoltage SensorVoltageExcerpt
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputPowerWatts shall contain the total output power, in watt units, for
	// this power supply. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	OutputPowerWatts SensorPowerExcerpt
	// RailCurrentAmps shall contain the output currents, in ampere units, for this
	// power supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailCurrentAmps []SensorCurrentExcerpt
	// RailCurrentAmps@odata.count
	RailCurrentAmpsCount int `json:"RailCurrentAmps@odata.count"`
	// RailPowerWatts shall contain the output power readings, in watt units, for
	// this power supply. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailPowerWatts []SensorPowerExcerpt
	// RailPowerWatts@odata.count
	RailPowerWattsCount int `json:"RailPowerWatts@odata.count"`
	// RailVoltage shall contain the output voltages, in volt units, for this power
	// supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailVoltage []SensorVoltageExcerpt
	// RailVoltage@odata.count
	RailVoltageCount int `json:"RailVoltage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this resource. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSupplyMetrics object from the raw JSON.
func (p *PowerSupplyMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyMetrics
	type pActions struct {
		ResetMetrics common.ActionTarget `json:"#PowerSupplyMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions pActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerSupplyMetrics(tmp.temp)

	// Extract the links to other entities for later
	p.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerSupplyMetrics) Update() error {
	readWriteFields := []string{
		"EnergykWh",
		"FanSpeedPercent",
		"FanSpeedsPercent",
		"FanSpeedsPercent@odata.count",
		"FrequencyHz",
		"InputCurrentAmps",
		"InputPowerWatts",
		"InputVoltage",
		"OutputPowerWatts",
		"RailCurrentAmps",
		"RailCurrentAmps@odata.count",
		"RailPowerWatts",
		"RailPowerWatts@odata.count",
		"RailVoltage",
		"RailVoltage@odata.count",
		"Status",
		"TemperatureCelsius",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
}

// GetPowerSupplyMetrics will get a PowerSupplyMetrics instance from the service.
func GetPowerSupplyMetrics(c common.Client, uri string) (*PowerSupplyMetrics, error) {
	return common.GetObject[PowerSupplyMetrics](c, uri)
}

// ListReferencedPowerSupplyMetricss gets the collection of PowerSupplyMetrics from
// a provided reference.
func ListReferencedPowerSupplyMetricss(c common.Client, link string) ([]*PowerSupplyMetrics, error) {
	return common.GetCollectionObjects[PowerSupplyMetrics](c, link)
}

// ResetMetrics shall reset any time intervals or counted values for this
// equipment.
func (p *PowerSupplyMetrics) ResetMetrics() error {
	payload := make(map[string]any)
	return p.Post(p.resetMetricsTarget, payload)
}
