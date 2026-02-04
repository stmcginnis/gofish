//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PowerSupplyMetrics.v1_2_0.json
// 2025.4 - #PowerSupplyMetrics.v1_2_0.PowerSupplyMetrics

package schemas

import (
	"encoding/json"
)

// PowerSupplyMetrics shall be used to represent the metrics of a power supply
// unit for a Redfish implementation.
type PowerSupplyMetrics struct {
	Entity
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
	// FanSpeedsPercentCount
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputPowerWatts shall contain the total output power, in watt units, for
	// this power supply. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	OutputPowerWatts SensorPowerExcerpt
	// PolyPhaseCurrentAmps shall contain the input polyphase current sensors for
	// this power supply.
	//
	// Version added: v1.2.0
	PolyPhaseCurrentAmps CurrentSensors
	// PolyPhaseEnergykWh shall contain the input polyphase energy sensors for this
	// power supply.
	//
	// Version added: v1.2.0
	PolyPhaseEnergykWh EnergySensors
	// PolyPhasePowerWatts shall contain the input polyphase power sensors for this
	// power supply.
	//
	// Version added: v1.2.0
	PolyPhasePowerWatts PowerSensors
	// PolyPhaseVoltage shall contain the input polyphase voltage sensors for this
	// power supply.
	//
	// Version added: v1.2.0
	PolyPhaseVoltage VoltageSensors
	// RailCurrentAmps shall contain the output currents, in ampere units, for this
	// power supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailCurrentAmps []SensorCurrentExcerpt
	// RailCurrentAmpsCount
	RailCurrentAmpsCount int `json:"RailCurrentAmps@odata.count"`
	// RailPowerWatts shall contain the output power readings, in watt units, for
	// this power supply. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailPowerWatts []SensorPowerExcerpt
	// RailPowerWattsCount
	RailPowerWattsCount int `json:"RailPowerWatts@odata.count"`
	// RailVoltage shall contain the output voltages, in volt units, for this power
	// supply. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. The sensors shall appear in the same array
	// order as the 'OutputRails' property in the associated 'PowerSupply'
	// resource.
	RailVoltage []SensorVoltageExcerpt
	// RailVoltageCount
	RailVoltageCount int `json:"RailVoltage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this resource. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a PowerSupplyMetrics object from the raw JSON.
func (p *PowerSupplyMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyMetrics
	type pActions struct {
		ResetMetrics ActionTarget `json:"#PowerSupplyMetrics.ResetMetrics"`
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

	return nil
}

// GetPowerSupplyMetrics will get a PowerSupplyMetrics instance from the service.
func GetPowerSupplyMetrics(c Client, uri string) (*PowerSupplyMetrics, error) {
	return GetObject[PowerSupplyMetrics](c, uri)
}

// ListReferencedPowerSupplyMetricss gets the collection of PowerSupplyMetrics from
// a provided reference.
func ListReferencedPowerSupplyMetricss(c Client, link string) ([]*PowerSupplyMetrics, error) {
	return GetCollectionObjects[PowerSupplyMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// equipment.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerSupplyMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetMetricsTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
