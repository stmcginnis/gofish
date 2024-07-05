//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"

	"github.com/stmcginnis/gofish/common"
)

// HeaterMetrics shall be used to represent the metrics of a heater unit for a Redfish implementation.
type HeaterMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the total power consumption, in watt units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// PrePowerOnHeatingTimeSeconds shall contain the total number of seconds the heater was active while the device it
	// heats was powered off.
	PrePowerOnHeatingTimeSeconds int
	// RuntimeHeatingTimeSeconds shall contain the total number of seconds the heater was active while the device it
	// heats was powered on.
	RuntimeHeatingTimeSeconds int
	// TemperatureReadingsCelsius shall contain the temperatures, in degree Celsius units, for this subsystem. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'Temperature'.
	TemperatureReadingsCelsius []SensorArrayExcerpt
	// TemperatureReadingsCelsiusCount is the number of TemperatureReadingCelsius entries.
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`

	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a HeaterMetrics object from the raw JSON.
func (heatermetrics *HeaterMetrics) UnmarshalJSON(b []byte) error {
	type temp HeaterMetrics
	type Actions struct {
		ResetMetrics common.ActionTarget `json:"#HeaterMetrics.ResetMetrics"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*heatermetrics = HeaterMetrics(t.temp)

	// Extract the links to other entities for later
	heatermetrics.resetMetricsTarget = t.Actions.ResetMetrics.Target

	return nil
}

// This action shall reset any time intervals or counted values for this circuit.
func (heatermetrics *HeaterMetrics) ResetMetrics() error {
	if heatermetrics.resetMetricsTarget == "" {
		return errors.New("ResetMetrics is not supported")
	}

	return heatermetrics.Post(heatermetrics.resetMetricsTarget, nil)
}

// GetHeaterMetrics will get a HeaterMetrics instance from the service.
func GetHeaterMetrics(c common.Client, uri string) (*HeaterMetrics, error) {
	return common.GetObject[HeaterMetrics](c, uri)
}

// ListReferencedHeaterMetrics gets the collection of HeaterMetrics from
// a provided reference.
func ListReferencedHeaterMetrics(c common.Client, link string) ([]*HeaterMetrics, error) {
	return common.GetCollectionObjects[HeaterMetrics](c, link)
}
