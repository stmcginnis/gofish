//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #ThermalMetrics.v1_3_2.ThermalMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ThermalMetrics shall represent the thermal metrics of a chassis for a Redfish
// implementation.
type ThermalMetrics struct {
	common.Entity
	// AirFlowCubicMetersPerMinute shall contain the rate of air flow, in cubic
	// meters per minute units, between the air intake and the air exhaust of this
	// chassis. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'AirFlowCMM'.
	//
	// Version added: v1.2.0
	AirFlowCubicMetersPerMinute SensorExcerpt
	// DeltaPressurekPa shall contain the pressure, in kilopascal units, for the
	// difference in pressure between the air intake and the air exhaust of this
	// chassis. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'PressurekPa'.
	//
	// Version added: v1.2.0
	DeltaPressurekPa SensorExcerpt
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for the
	// thermal subsystem. The value shall include the total energy consumption of
	// devices involved in thermal management of the chassis, such as fans, pumps,
	// and heaters. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'.
	//
	// Version added: v1.3.0
	EnergykWh SensorEnergykWhExcerpt
	// HeaterSummary shall contain the summary of heater metrics for this
	// subsystem.
	//
	// Version added: v1.1.0
	HeaterSummary HeaterSummary
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the power, in watt units, for the thermal
	// subsystem. The value shall include the total power consumption of devices
	// involved in thermal management of the chassis, such as fans, pumps, and
	// heaters. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	//
	// Version added: v1.3.0
	PowerWatts SensorPowerExcerpt
	// TemperatureReadingsCelsius shall contain the temperatures, in degree Celsius
	// units, for this subsystem. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Temperature'.
	TemperatureReadingsCelsius []SensorArrayExcerpt
	// TemperatureReadingsCelsius@odata.count
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`
	// TemperatureSummaryCelsius shall contain the temperature sensor readings for
	// this subsystem.
	TemperatureSummaryCelsius TemperatureSummary
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ThermalMetrics object from the raw JSON.
func (t *ThermalMetrics) UnmarshalJSON(b []byte) error {
	type temp ThermalMetrics
	type tActions struct {
		ResetMetrics common.ActionTarget `json:"#ThermalMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions tActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = ThermalMetrics(tmp.temp)

	// Extract the links to other entities for later
	t.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *ThermalMetrics) Update() error {
	readWriteFields := []string{
		"EnergykWh",
		"HeaterSummary",
		"PowerWatts",
		"TemperatureReadingsCelsius",
		"TemperatureReadingsCelsius@odata.count",
		"TemperatureSummaryCelsius",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
}

// GetThermalMetrics will get a ThermalMetrics instance from the service.
func GetThermalMetrics(c common.Client, uri string) (*ThermalMetrics, error) {
	return common.GetObject[ThermalMetrics](c, uri)
}

// ListReferencedThermalMetricss gets the collection of ThermalMetrics from
// a provided reference.
func ListReferencedThermalMetricss(c common.Client, link string) ([]*ThermalMetrics, error) {
	return common.GetCollectionObjects[ThermalMetrics](c, link)
}

// ResetMetrics shall reset any time intervals or counted values for this
// equipment.
func (t *ThermalMetrics) ResetMetrics() error {
	payload := make(map[string]any)
	return t.Post(t.resetMetricsTarget, payload)
}

// HeaterSummary shall contain properties that describe the heater metrics
// summary for the subsystem.
type HeaterSummary struct {
	// TotalPrePowerOnHeatingTimeSeconds shall contain the total number of seconds
	// all the heaters in the thermal subsystem were active while the respective
	// devices they heat were powered off.
	//
	// Version added: v1.1.0
	TotalPrePowerOnHeatingTimeSeconds *int `json:",omitempty"`
	// TotalRuntimeHeatingTimeSeconds shall contain the total number of seconds all
	// the heaters in the thermal subsystem were active while the respective
	// devices they heat were powered on.
	//
	// Version added: v1.1.0
	TotalRuntimeHeatingTimeSeconds *int `json:",omitempty"`
}

// TemperatureSummary shall contain properties that describe temperature sensor
// for a subsystem.
type TemperatureSummary struct {
	// Ambient shall contain the temperature, in degree Celsius units, for the
	// ambient temperature of this subsystem. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Temperature'.
	Ambient SensorExcerpt
	// Exhaust shall contain the temperature, in degree Celsius units, for the
	// exhaust temperature of this subsystem. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Temperature'.
	Exhaust SensorExcerpt
	// Intake shall contain the temperature, in degree Celsius units, for the
	// intake temperature of this subsystem. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Temperature'.
	Intake SensorExcerpt
	// Internal shall contain the temperature, in degree Celsius units, for the
	// internal temperature of this subsystem. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Temperature'.
	Internal SensorExcerpt
}
