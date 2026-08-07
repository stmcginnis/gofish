//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreweave/gofish/common"
)

// EnvironmentMetrics shall represent the environmental metrics for a Redfish implementation.
type EnvironmentMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AbsoluteHumidity shall contain the absolute (volumetric) humidity sensor reading, in grams per cubic meter
	// units, for this resource. The value of the DataSourceUri property, if present, shall reference a resource of
	// type Sensor with the ReadingType property containing the value 'AbsoluteHumidity'.
	AbsoluteHumidity SensorExcerpt
	// Description provides a description of this resource.
	Description string
	// DewPointCelsius shall contain the dew point, in degree Celsius units, based on the temperature and humidity
	// values for this resource. The value of the DataSourceUri property, if present, shall reference a resource of
	// type Sensor with the ReadingType property containing the value 'Temperature'.
	DewPointCelsius SensorExcerpt
	// EnergyJoules shall contain the total energy, in joule units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergyJoules'. This property is used for reporting device-level energy consumption measurements, while
	// EnergykWh is used for large-scale consumption measurements.
	EnergyJoules SensorExcerpt
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FanSpeedsPercent shall contain the fan speeds, in percent units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Percent'.
	FanSpeedsPercent []SensorFanArrayExcerpt
	// FanSpeedsPercent@odata.count
	FanSpeedsPercentCount int `json:"FanSpeedsPercent@odata.count"`
	// HumidityPercent shall contain the humidity, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Humidity'.
	HumidityPercent SensorExcerpt
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerLimitWatts shall contain the power limit control, in watt units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Control with the ControlType property
	// containing the value of 'Power'.
	PowerLimitWatts ControlSingleExcerpt
	// PowerLoadPercent shall contain the power load, in percent units, for this device that represents the 'Total'
	// ElectricalContext for this device. The value of the DataSourceUri property, if present, shall reference a
	// resource of type Sensor with the ReadingType property containing the value 'Percent'.
	PowerLoadPercent SensorExcerpt
	// PowerWatts shall contain the total power, in watt units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureCelsius shall contain the temperature, in degree Celsius units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	resetMetricsTarget    string
	resetToDefaultsTarget string
}

// UnmarshalJSON unmarshals a EnvironmentMetrics object from the raw JSON.
func (environmentmetrics *EnvironmentMetrics) UnmarshalJSON(b []byte) error {
	type temp EnvironmentMetrics
	type Actions struct {
		ResetMetrics    common.ActionTarget `json:"#EnvironmentMetrics.ResetMetrics"`
		ResetToDefaults common.ActionTarget `json:"#EnvironmentMetrics.ResetToDefaults"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*environmentmetrics = EnvironmentMetrics(t.temp)

	// Extract the links to other entities for later
	environmentmetrics.resetMetricsTarget = t.Actions.ResetMetrics.Target
	environmentmetrics.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target

	// This is a read/write object, so we need to save the raw object data for later
	environmentmetrics.rawData = b

	return nil
}

// ResetMetrics resets the summary metrics related to this equipment.
func (environmentmetrics *EnvironmentMetrics) ResetMetrics() error {
	return environmentmetrics.ResetMetricsWithContext(common.ContextOf(environmentmetrics.GetClient()))
}

// ResetMetricsWithContext resets the summary metrics related to this equipment.
func (environmentmetrics *EnvironmentMetrics) ResetMetricsWithContext(ctx context.Context) error {
	if environmentmetrics.resetMetricsTarget == "" {
		return fmt.Errorf("ResetMetrics is not supported by this system")
	}

	return environmentmetrics.PostWithContext(ctx, environmentmetrics.resetMetricsTarget, nil)
}

// ResetToDefaults resets the values of writable properties to factory defaults.
func (environmentmetrics *EnvironmentMetrics) ResetToDefaults() error {
	return environmentmetrics.ResetToDefaultsWithContext(common.ContextOf(environmentmetrics.GetClient()))
}

// ResetToDefaultsWithContext resets the values of writable properties to factory defaults.
func (environmentmetrics *EnvironmentMetrics) ResetToDefaultsWithContext(ctx context.Context) error {
	if environmentmetrics.resetToDefaultsTarget == "" {
		return fmt.Errorf("ResetToDefaults is not supported by this system")
	}

	return environmentmetrics.PostWithContext(ctx, environmentmetrics.resetToDefaultsTarget, nil)
}

// Update commits updates to this object's properties to the running system.
func (environmentmetrics *EnvironmentMetrics) Update() error {
	return environmentmetrics.UpdateWithContext(common.ContextOf(environmentmetrics.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (environmentmetrics *EnvironmentMetrics) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"PowerLimitWatts",
	}

	return environmentmetrics.UpdateFromRawDataWithContext(ctx, environmentmetrics, environmentmetrics.rawData, readWriteFields)
}

// GetEnvironmentMetrics will get a EnvironmentMetrics instance from the service.
func GetEnvironmentMetrics(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*EnvironmentMetrics, error) {
	return GetEnvironmentMetricsWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetEnvironmentMetricsWithContext will get a EnvironmentMetrics instance from the service.
func GetEnvironmentMetricsWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*EnvironmentMetrics, error) {
	return common.GetObjectWithContext[EnvironmentMetrics](ctx, c, uri, queryOpts...)
}

// ListReferencedEnvironmentMetrics gets the collection of EnvironmentMetrics from
// a provided reference.
func ListReferencedEnvironmentMetrics(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*EnvironmentMetrics, error) {
	return ListReferencedEnvironmentMetricsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedEnvironmentMetricsWithContext gets the collection of EnvironmentMetrics from
// a provided reference.
func ListReferencedEnvironmentMetricsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*EnvironmentMetrics, error) {
	return common.GetCollectionObjectsWithContext[EnvironmentMetrics](ctx, c, link, queryOpts...)
}
