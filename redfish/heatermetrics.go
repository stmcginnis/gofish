//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #HeaterMetrics.v1_0_2.HeaterMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// HeaterMetrics shall be used to represent the metrics of a heater unit for a
// Redfish implementation.
type HeaterMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
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
	// TemperatureReadingsCelsius@odata.count
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a HeaterMetrics object from the raw JSON.
func (h *HeaterMetrics) UnmarshalJSON(b []byte) error {
	type temp HeaterMetrics
	type hActions struct {
		ResetMetrics common.ActionTarget `json:"#HeaterMetrics.ResetMetrics"`
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

	// This is a read/write object, so we need to save the raw object data for later
	h.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (h *HeaterMetrics) Update() error {
	readWriteFields := []string{
		"PowerWatts",
		"TemperatureReadingsCelsius",
		"TemperatureReadingsCelsius@odata.count",
	}

	return h.UpdateFromRawData(h, h.rawData, readWriteFields)
}

// GetHeaterMetrics will get a HeaterMetrics instance from the service.
func GetHeaterMetrics(c common.Client, uri string) (*HeaterMetrics, error) {
	return common.GetObject[HeaterMetrics](c, uri)
}

// ListReferencedHeaterMetricss gets the collection of HeaterMetrics from
// a provided reference.
func ListReferencedHeaterMetricss(c common.Client, link string) ([]*HeaterMetrics, error) {
	return common.GetCollectionObjects[HeaterMetrics](c, link)
}

// ResetMetrics shall reset any time intervals or counted values for this
// equipment.
func (h *HeaterMetrics) ResetMetrics() error {
	payload := make(map[string]any)
	return h.Post(h.resetMetricsTarget, payload)
}
