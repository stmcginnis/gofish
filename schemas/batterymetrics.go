//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/BatteryMetrics.v1_1_0.json
// 2025.3 - #BatteryMetrics.v1_1_0.BatteryMetrics

package schemas

import (
	"encoding/json"
)

// BatteryMetrics shall be used to represent the metrics of a battery unit for a
// Redfish implementation.
type BatteryMetrics struct {
	Entity
	// CRate shall contain the rate at which the battery is charging or
	// discharging, based on current in ampere units, relative to its rated maximum
	// capacity in ampere-hour units. If the battery is discharging, services shall
	// calculate the value as 'SUM(OutputCurrentAmps) / CapacityRatedAmpHours'.
	// Otherwise, services shall calculate the value as 'InputCurrentAmps /
	// CapacityRatedAmpHours'.
	//
	// Version added: v1.1.0
	CRate *float64 `json:",omitempty"`
	// CellVoltages shall contain the cell voltages, in volt units, for this
	// battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'.
	CellVoltages []SensorVoltageExcerpt
	// CellVoltagesCount
	CellVoltagesCount int `json:"CellVoltages@odata.count"`
	// ChargePercent shall contain the amount of charge available, in percent
	// units, typically '0' to '100', in this battery. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Percent'.
	ChargePercent SensorExcerpt
	// DischargeCycles shall contain the number of discharges this battery has
	// sustained.
	DischargeCycles *float64 `json:",omitempty"`
	// ERate shall contain the rate at which the battery is charging or
	// discharging, based on power in watt units, relative to its rated maximum
	// capacity in watt-hour units. If the battery is discharging, services shall
	// calculate the value as 'SUM(OutputCurrentAmps[x] * OutputVoltages[x]) /
	// CapacityRatedWattHours'. Otherwise, services shall calculate the value as
	// '(InputCurrentAmps * InputVoltage) / CapacityRatedWattHours'.
	//
	// Version added: v1.1.0
	ERate *float64 `json:",omitempty"`
	// InputCurrentAmps shall contain the input current, in ampere units, for this
	// battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'.
	InputCurrentAmps SensorCurrentExcerpt
	// InputVoltage shall contain the input voltage, in volt units, for this
	// battery. The value of the 'DataSourceUri' property, if present, shall
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
	// OutputCurrentAmps shall contain the output currents, in ampere units, for
	// this battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'. The sensors shall appear in the same array
	// order as the 'OutputVoltages' property.
	OutputCurrentAmps []SensorCurrentExcerpt
	// OutputCurrentAmpsCount
	OutputCurrentAmpsCount int `json:"OutputCurrentAmps@odata.count"`
	// OutputVoltages shall contain the output voltages, in volt units, for this
	// battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. The sensors shall appear in the same array
	// order as the 'OutputCurrentAmps' property.
	OutputVoltages []SensorVoltageExcerpt
	// OutputVoltagesCount
	OutputVoltagesCount int `json:"OutputVoltages@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// StoredChargeAmpHours shall contain the stored charge, in ampere-hour units,
	// for this battery. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'ChargeAh'.
	StoredChargeAmpHours SensorExcerpt
	// StoredEnergyWattHours shall contain the stored energy, in watt-hour units,
	// for this battery. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergyWh'.
	StoredEnergyWattHours SensorExcerpt
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this battery. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
}

// GetBatteryMetrics will get a BatteryMetrics instance from the service.
func GetBatteryMetrics(c Client, uri string) (*BatteryMetrics, error) {
	return GetObject[BatteryMetrics](c, uri)
}

// ListReferencedBatteryMetricss gets the collection of BatteryMetrics from
// a provided reference.
func ListReferencedBatteryMetricss(c Client, link string) ([]*BatteryMetrics, error) {
	return GetCollectionObjects[BatteryMetrics](c, link)
}
