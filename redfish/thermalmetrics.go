//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// HeaterSummary shall contain properties that describe the heater metrics summary for the subsystem.
type HeaterSummary struct {
	// TotalPrePowerOnHeatingTimeSeconds shall contain the total number of seconds all the heaters in the thermal
	// subsystem were active while the respective devices they heat were powered off.
	TotalPrePowerOnHeatingTimeSeconds int
	// TotalRuntimeHeatingTimeSeconds shall contain the total number of seconds all the heaters in the thermal
	// subsystem were active while the respective devices they heat were powered on.
	TotalRuntimeHeatingTimeSeconds int
}

// TemperatureSummary shall contain properties that describe temperature sensor for a subsystem.
type TemperatureSummary struct {
	// Ambient shall contain the temperature, in degree Celsius units, for the ambient temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Ambient SensorExcerpt
	// Exhaust shall contain the temperature, in degree Celsius units, for the exhaust temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Exhaust SensorExcerpt
	// Intake shall contain the temperature, in degree Celsius units, for the intake temperature of this subsystem. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'Temperature'.
	Intake SensorExcerpt
	// Internal shall contain the temperature, in degree Celsius units, for the internal temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Internal SensorExcerpt
}

// ThermalMetrics shall represent the thermal metrics of a chassis for a Redfish implementation.
type ThermalMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AirFlowCubicMetersPerMinute shall contain the rate of air flow, in cubic meters per minute units, between the
	// air intake and the air exhaust of this chassis. The value of the DataSourceUri property, if present, shall
	// reference a resource of type Sensor with the ReadingType property containing the value 'AirFlowCMM'.
	AirFlowCubicMetersPerMinute SensorExcerpt
	// DeltaPressurekPa shall contain the pressure, in kilopascal units, for the difference in pressure between the air
	// intake and the air exhaust of this chassis. The value of the DataSourceUri property, if present, shall reference
	// a resource of type Sensor with the ReadingType property containing the value 'PressurekPa'.
	DeltaPressurekPa SensorExcerpt
	// Description provides a description of this resource.
	Description string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for the thermal subsystem. The value shall
	// include the total energy consumption of devices involved in thermal management of the chassis, such as fans,
	// pumps, and heaters. The value of the DataSourceUri property, if present, shall reference a resource of type
	// Sensor with the ReadingType property containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// HeaterSummary shall contain the summary of heater metrics for this subsystem.
	HeaterSummary HeaterSummary
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the power, in watt units, for the thermal subsystem. The value shall include the total
	// power consumption of devices involved in thermal management of the chassis, such as fans, pumps, and heaters.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureReadingsCelsius shall contain the temperatures, in degree Celsius units, for this subsystem. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'Temperature'.
	TemperatureReadingsCelsius []SensorArrayExcerpt
	// TemperatureReadingsCelsius@odata.count
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`
	// TemperatureSummaryCelsius shall contain the temperature sensor readings for this subsystem.
	TemperatureSummaryCelsius TemperatureSummary

	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a ThermalMetrics object from the raw JSON.
func (thermalmetrics *ThermalMetrics) UnmarshalJSON(b []byte) error {
	type temp ThermalMetrics
	var t struct {
		temp
		Actions struct {
			ResetMetrics struct {
				Target string
			} `json:"#ThermalMetrics.ResetMetrics"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalmetrics = ThermalMetrics(t.temp)

	// Extract the links to other entities for later
	thermalmetrics.resetMetricsTarget = t.Actions.ResetMetrics.Target

	return nil
}

// ResetMetrics resets the summary metrics related to this equipment.
func (thermalmetrics *ThermalMetrics) ResetMetrics() error {
	return thermalmetrics.Post(thermalmetrics.resetMetricsTarget, nil)
}

// GetThermalMetrics will get a ThermalMetrics instance from the service.
func GetThermalMetrics(c common.Client, uri string) (*ThermalMetrics, error) {
	return common.GetObject[ThermalMetrics](c, uri)
}

// ListReferencedThermalMetricss gets the collection of ThermalMetrics from
// a provided reference.
func ListReferencedThermalMetrics(c common.Client, link string) ([]*ThermalMetrics, error) {
	return common.GetCollectionObjects[ThermalMetrics](c, link)
}
