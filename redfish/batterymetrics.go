//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #BatteryMetrics.v1_1_0.BatteryMetrics

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// BatteryMetrics shall be used to represent the metrics of a battery unit for a
// Redfish implementation.
type BatteryMetrics struct {
	common.Entity
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
	// CellVoltages@odata.count
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputCurrentAmps shall contain the output currents, in ampere units, for
	// this battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Current'. The sensors shall appear in the same array
	// order as the 'OutputVoltages' property.
	OutputCurrentAmps []SensorCurrentExcerpt
	// OutputCurrentAmps@odata.count
	OutputCurrentAmpsCount int `json:"OutputCurrentAmps@odata.count"`
	// OutputVoltages shall contain the output voltages, in volt units, for this
	// battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. The sensors shall appear in the same array
	// order as the 'OutputCurrentAmps' property.
	OutputVoltages []SensorVoltageExcerpt
	// OutputVoltages@odata.count
	OutputVoltagesCount int `json:"OutputVoltages@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a BatteryMetrics object from the raw JSON.
func (ba *BatteryMetrics) UnmarshalJSON(b []byte) error {
	type temp BatteryMetrics
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*ba = BatteryMetrics(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	ba.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ba *BatteryMetrics) Update() error {
	readWriteFields := []string{
		"CellVoltages",
		"CellVoltages@odata.count",
		"ChargePercent",
		"InputCurrentAmps",
		"InputVoltage",
		"OutputCurrentAmps",
		"OutputCurrentAmps@odata.count",
		"OutputVoltages",
		"OutputVoltages@odata.count",
		"Status",
		"StoredChargeAmpHours",
		"StoredEnergyWattHours",
		"TemperatureCelsius",
	}

	return ba.UpdateFromRawData(ba, ba.rawData, readWriteFields)
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
