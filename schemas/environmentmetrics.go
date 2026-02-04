//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/EnvironmentMetrics.v1_5_0.json
// 2025.2 - #EnvironmentMetrics.v1_5_0.EnvironmentMetrics

package schemas

import (
	"encoding/json"
)

// EnvironmentMetrics shall represent the environmental metrics for a Redfish
// implementation.
type EnvironmentMetrics struct {
	Entity
	// AbsoluteHumidity shall contain the absolute (volumetric) humidity sensor
	// reading, in grams per cubic meter units, for this resource. The value should
	// reflect the humidity measured at the exterior of the containing 'Chassis'
	// instance, or the interior of the containing 'Facility' instance. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Sensor' with the 'ReadingType' property containing the value
	// 'AbsoluteHumidity'.
	//
	// Version added: v1.2.0
	AbsoluteHumidity SensorExcerpt
	// AmbientTemperatureCelsius shall contain the ambient temperature, in degree
	// Celsius units, for this resource. The ambient temperature shall be the
	// temperature measured at a point exterior to the 'Chassis' containing this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'. This property shall only be present, if
	// supported, in resource instances subordinate to a 'Chassis' or 'CoolingUnit'
	// resource.
	//
	// Version added: v1.4.0
	AmbientTemperatureCelsius SensorExcerpt
	// CurrentAmps shall contain the current, in ampere units, for this device. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Current'.
	//
	// Version added: v1.5.0
	CurrentAmps SensorCurrentExcerpt
	// DewPointCelsius shall contain the dew point, in degree Celsius units, based
	// on the temperature and humidity values for this resource. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Temperature'.
	//
	// Version added: v1.1.0
	DewPointCelsius SensorExcerpt
	// EnergyJoules shall contain the total energy, in joule units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergyJoules'. This property is used for reporting
	// device-level energy consumption measurements, while 'EnergykWh' is used for
	// large-scale consumption measurements.
	//
	// Version added: v1.2.0
	EnergyJoules SensorExcerpt
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FanSpeedsPercent shall contain the fan speeds, in percent units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Percent'.
	FanSpeedsPercent []SensorFanArrayExcerpt
	// FanSpeedsPercentCount
	FanSpeedsPercentCount int `json:"FanSpeedsPercent@odata.count"`
	// HumidityPercent shall contain the humidity, in percent units, for this
	// resource. The value should reflect the humidity measured at the exterior of
	// the containing 'Chassis' instance, or the interior of the containing
	// 'Facility' instance. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Humidity'.
	HumidityPercent SensorExcerpt
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerLimitWatts shall contain the power limit control, in watt units, for
	// this resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Control' with the 'ControlType' property
	// containing the value of 'Power'.
	//
	// Version added: v1.1.0
	PowerLimitWatts ControlSingleExcerpt
	// PowerLoadPercent shall contain the power load, in percent units, for this
	// device that represents the 'Total' 'ElectricalContext' for this device. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Percent'.
	//
	// Version added: v1.1.0
	PowerLoadPercent SensorExcerpt
	// PowerWatts shall contain the total power, in watt units, for this resource.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this resource. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// Voltage shall contain the voltage, in volt units, for this device. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Sensor' with the 'ReadingType' property containing the value
	// 'Voltage'.
	//
	// Version added: v1.5.0
	Voltage SensorVoltageExcerpt
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a EnvironmentMetrics object from the raw JSON.
func (e *EnvironmentMetrics) UnmarshalJSON(b []byte) error {
	type temp EnvironmentMetrics
	type eActions struct {
		ResetMetrics    ActionTarget `json:"#EnvironmentMetrics.ResetMetrics"`
		ResetToDefaults ActionTarget `json:"#EnvironmentMetrics.ResetToDefaults"`
	}
	var tmp struct {
		temp
		Actions eActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = EnvironmentMetrics(tmp.temp)

	// Extract the links to other entities for later
	e.resetMetricsTarget = tmp.Actions.ResetMetrics.Target
	e.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target

	// This is a read/write object, so we need to save the raw object data for later
	e.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (e *EnvironmentMetrics) Update() error {
	readWriteFields := []string{
		"PowerLimitWatts",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}

// GetEnvironmentMetrics will get a EnvironmentMetrics instance from the service.
func GetEnvironmentMetrics(c Client, uri string) (*EnvironmentMetrics, error) {
	return GetObject[EnvironmentMetrics](c, uri)
}

// ListReferencedEnvironmentMetricss gets the collection of EnvironmentMetrics from
// a provided reference.
func ListReferencedEnvironmentMetricss(c Client, link string) ([]*EnvironmentMetrics, error) {
	return GetCollectionObjects[EnvironmentMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// equipment.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EnvironmentMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(e.client,
		e.resetMetricsTarget, payload, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (e *EnvironmentMetrics) ResetToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(e.client,
		e.resetToDefaultsTarget, payload, e.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
