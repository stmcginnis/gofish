//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// BatteryMetrics shall be used to represent the metrics of a battery unit for a Redfish implementation.
type BatteryMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CellVoltages shall contain the cell voltages, in volt units, for this battery. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'.
	CellVoltages []SensorVoltageExcerpt
	// CellVoltages@odata.count
	CellVoltagesCount int `json:"CellVoltages@odata.count"`
	// ChargePercent shall contain the amount of charge available, in percent units, typically '0' to '100', in this
	// battery. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Percent'.
	ChargePercent SensorExcerpt
	// Description provides a description of this resource.
	Description string
	// DischargeCycles shall contain the number of discharges this battery has sustained.
	DischargeCycles float64
	// InputCurrentAmps shall contain the input current, in ampere units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'.
	InputCurrentAmps SensorCurrentExcerpt
	// InputVoltage shall contain the input voltage, in volt units, for this battery. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'.
	InputVoltage SensorVoltageExcerpt
	// OutputCurrentAmps shall contain the output currents, in ampere units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'. The sensors shall appear in the same array order as the OutputVoltages property.
	OutputCurrentAmps []SensorCurrentExcerpt
	// OutputCurrentAmps@odata.count
	OutputCurrentAmpsCount int `json:"OutputCurrentAmps@odata.count"`
	// OutputVoltages shall contain the output voltages, in volt units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. The sensors shall appear in the same array order as the OutputCurrentAmps
	// property.
	OutputVoltages []SensorVoltageExcerpt
	// OutputVoltages@odata.count
	OutputVoltagesCount int `json:"OutputVoltages@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StoredChargeAmpHours shall contain the stored charge, in ampere-hour units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'ChargeAh'.
	StoredChargeAmpHours SensorExcerpt
	// StoredEnergyWattHours shall contain the stored energy, in watt-hour units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergyWh'.
	StoredEnergyWattHours SensorExcerpt
	// TemperatureCelsius shall contain the temperature, in degree Celsius units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
}

// GetBatteryMetrics will get a BatteryMetrics instance from the service.
func GetBatteryMetrics(c common.Client, uri string) (*BatteryMetrics, error) {
	return common.GetObject[BatteryMetrics](c, uri)
}

// ListReferencedBatteryMetricss gets the collection of BatteryMetrics from
// a provided reference.
func ListReferencedBatteryMetricss(c common.Client, link string) ([]*BatteryMetrics, error) {
	return common.GetCollectionObjects[BatteryMetrics](c, link)
}
