//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/HeaterMetrics.v1_0_2.json
// 2022.3 - #HeaterMetrics.v1_0_2.HeaterMetrics

package schemas

import (
	"encoding/json"
)

// HeaterMetrics shall be used to represent the metrics of a heater unit for a
// Redfish implementation.
type HeaterMetrics struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the total power consumption, in watt units, for
	// this resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// PrePowerOnHeatingTimeSeconds shall contain the total number of seconds the
	// heater was active while the device it heats was powered off.
	PrePowerOnHeatingTimeSeconds *int `json:",omitempty"`
	// RuntimeHeatingTimeSeconds shall contain the total number of seconds the
	// heater was active while the device it heats was powered on.
	RuntimeHeatingTimeSeconds *int `json:",omitempty"`
	// TemperatureReadingsCelsius shall contain the temperatures, in degree Celsius
	// units, for this subsystem. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Temperature'.
	TemperatureReadingsCelsius []SensorArrayExcerpt
	// TemperatureReadingsCelsiusCount
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a HeaterMetrics object from the raw JSON.
func (h *HeaterMetrics) UnmarshalJSON(b []byte) error {
	type temp HeaterMetrics
	type hActions struct {
		ResetMetrics ActionTarget `json:"#HeaterMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions hActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = HeaterMetrics(tmp.temp)

	// Extract the links to other entities for later
	h.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetHeaterMetrics will get a HeaterMetrics instance from the service.
func GetHeaterMetrics(c Client, uri string) (*HeaterMetrics, error) {
	return GetObject[HeaterMetrics](c, uri)
}

// ListReferencedHeaterMetricss gets the collection of HeaterMetrics from
// a provided reference.
func ListReferencedHeaterMetricss(c Client, link string) ([]*HeaterMetrics, error) {
	return GetCollectionObjects[HeaterMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// equipment.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (h *HeaterMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(h.client,
		h.resetMetricsTarget, payload, h.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
