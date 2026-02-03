//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.1 - #PowerDistributionMetrics.v1_4_0.PowerDistributionMetrics

package schemas

import (
	"encoding/json"
)

// PowerDistributionMetrics shall represent the metrics of a power distribution
// component or unit for a Redfish implementation.
type PowerDistributionMetrics struct {
	Entity
	// AbsoluteHumidity shall contain the absolute (volumetric) humidity sensor
	// reading, in grams per cubic meter units, for this resource. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value
	// 'AbsoluteHumidity'.
	//
	// Version added: v1.3.0
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
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// resource that represents the 'Total' 'ElectricalContext' sensor when
	// multiple energy sensors exist. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// HumidityPercent shall contain the humidity, in percent units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Humidity'.
	//
	// Version added: v1.1.0
	HumidityPercent SensorExcerpt
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerLoadPercent shall contain the power load, in percent units, for this
	// device that represents the 'Total' 'ElectricalContext' for this device. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Percent'.
	//
	// Version added: v1.2.0
	PowerLoadPercent SensorExcerpt
	// PowerWatts shall contain the total power, in watt units, for this resource
	// that represents the 'Total' 'ElectricalContext' sensor when multiple power
	// sensors exist. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this resource. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	//
	// Version added: v1.1.0
	TemperatureCelsius SensorExcerpt
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a PowerDistributionMetrics object from the raw JSON.
func (p *PowerDistributionMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerDistributionMetrics
	type pActions struct {
		ResetMetrics ActionTarget `json:"#PowerDistributionMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions pActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerDistributionMetrics(tmp.temp)

	// Extract the links to other entities for later
	p.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetPowerDistributionMetrics will get a PowerDistributionMetrics instance from the service.
func GetPowerDistributionMetrics(c Client, uri string) (*PowerDistributionMetrics, error) {
	return GetObject[PowerDistributionMetrics](c, uri)
}

// ListReferencedPowerDistributionMetricss gets the collection of PowerDistributionMetrics from
// a provided reference.
func ListReferencedPowerDistributionMetricss(c Client, link string) ([]*PowerDistributionMetrics, error) {
	return GetCollectionObjects[PowerDistributionMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// equipment.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerDistributionMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetMetricsTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
