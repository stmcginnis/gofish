//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Sensor.v1_12_0.json
// 2025.4 - #Sensor.v1_12_0.Sensor

package schemas

import (
	"encoding/json"
)

type ElectricalContext string

const (
	// Line1ElectricalContext shall represent a circuit that shares the L1
	// current-carrying conductor, such as circuits with phase wiring types of
	// Two-phase / 3-Wire or 4-Wire, or Three-phase / 4-Wire or 5-Wire.
	Line1ElectricalContext ElectricalContext = "Line1"
	// Line2ElectricalContext shall represent a circuit that shares the L2
	// current-carrying conductor, such as circuits with phase wiring types of
	// Two-phase / 4-Wire or Three-phase / 4-Wire or 5-Wire.
	Line2ElectricalContext ElectricalContext = "Line2"
	// Line3ElectricalContext shall represent a circuit that shares the L3
	// current-carrying conductor, such as circuits with phase wiring types of
	// Three-phase / 4-Wire or 5-Wire.
	Line3ElectricalContext ElectricalContext = "Line3"
	// NeutralElectricalContext shall represent the grounded current-carrying
	// return circuit of current-carrying conductors, such as circuits with phase
	// wiring types of Single-phase / 3-Wire, Two-phase / 4-Wire, or Three-phase /
	// 5-Wire.
	NeutralElectricalContext ElectricalContext = "Neutral"
	// LineToLineElectricalContext shall represent a circuit formed by two
	// current-carrying conductors, such as circuits with phase wiring types of
	// Two-phase / 3-Wire or 4-Wire, or Three-phase / 4-Wire or 5-Wire.
	LineToLineElectricalContext ElectricalContext = "LineToLine"
	// Line1ToLine2ElectricalContext shall represent a circuit formed by L1 and L2
	// current-carrying conductors, such as circuits with phase wiring types of
	// Two-phase / 3-Wire or 4-Wire, or Three-phase / 4-Wire or 5-Wire.
	Line1ToLine2ElectricalContext ElectricalContext = "Line1ToLine2"
	// Line2ToLine3ElectricalContext shall represent a circuit formed by L2 and L3
	// current-carrying conductors, such as circuits with phase wiring types of
	// Three-phase / 4-Wire or 5-Wire.
	Line2ToLine3ElectricalContext ElectricalContext = "Line2ToLine3"
	// Line3ToLine1ElectricalContext shall represent a circuit formed by L3 and L1
	// current-carrying conductors, such as circuits with phase wiring types of
	// Three-phase / 4-Wire or 5-Wire.
	Line3ToLine1ElectricalContext ElectricalContext = "Line3ToLine1"
	// LineToNeutralElectricalContext shall represent a circuit formed by a line
	// and neutral current-carrying conductor, such as circuits with phase wiring
	// types of Single-phase / 3-Wire, Two-phase / 4-Wire, or Three-phase / 4-Wire
	// or 5-Wire.
	LineToNeutralElectricalContext ElectricalContext = "LineToNeutral"
	// Line1ToNeutralElectricalContext shall represent a circuit formed by L1 and
	// neutral current-carrying conductors, such as circuits with phase wiring
	// types of Single-phase / 3-Wire, Two-phase / 3-Wire or 4-Wire, or Three-phase
	// / 4-Wire or 5-Wire.
	Line1ToNeutralElectricalContext ElectricalContext = "Line1ToNeutral"
	// Line2ToNeutralElectricalContext shall represent a circuit formed by L2 and
	// neutral current-carrying conductors, such as circuits with phase wiring
	// types of Two-phase / 4-Wire or Three-phase / 5-Wire.
	Line2ToNeutralElectricalContext ElectricalContext = "Line2ToNeutral"
	// Line3ToNeutralElectricalContext shall represent a circuit formed by L3 and
	// neutral current-carrying conductors, such as circuits with a phase wiring
	// type of Three-phase / 5-Wire.
	Line3ToNeutralElectricalContext ElectricalContext = "Line3ToNeutral"
	// Line1ToNeutralAndL1L2ElectricalContext shall represent a circuit formed by
	// L1, L2, and neutral current-carrying conductors, such as circuits with phase
	// wiring types of Two-phase/ 4-Wire or Three-phase / 5-Wire.
	Line1ToNeutralAndL1L2ElectricalContext ElectricalContext = "Line1ToNeutralAndL1L2"
	// Line2ToNeutralAndL1L2ElectricalContext shall represent a circuit formed by
	// L1, L2, and neutral current-carrying conductors, such as circuits with phase
	// wiring types of Two-phase/ 4-Wire or Three-phase / 5-Wire.
	Line2ToNeutralAndL1L2ElectricalContext ElectricalContext = "Line2ToNeutralAndL1L2"
	// Line2ToNeutralAndL2L3ElectricalContext shall represent a circuit formed by
	// L2, L3, and neutral current-carrying conductors, such as circuits with a
	// phase wiring type of Three-phase / 5-Wire.
	Line2ToNeutralAndL2L3ElectricalContext ElectricalContext = "Line2ToNeutralAndL2L3"
	// Line3ToNeutralAndL3L1ElectricalContext shall represent a circuit formed by
	// L3, L1, and neutral current-carrying conductors, such as circuits with a
	// phase wiring type of Three-phase / 5-Wire.
	Line3ToNeutralAndL3L1ElectricalContext ElectricalContext = "Line3ToNeutralAndL3L1"
	// TotalElectricalContext shall represent the circuits formed by all
	// current-carrying conductors for any phase wiring type.
	TotalElectricalContext ElectricalContext = "Total"
)

type ImplementationType string

const (
	// PhysicalSensorImplementationType The reading is acquired from a physical
	// sensor.
	PhysicalSensorImplementationType ImplementationType = "PhysicalSensor"
	// SynthesizedImplementationType The reading is obtained by applying a
	// calculation on one or more properties or multiple sensors. The calculation
	// is not provided.
	SynthesizedImplementationType ImplementationType = "Synthesized"
	// ReportedImplementationType The reading is obtained from software or a
	// device.
	ReportedImplementationType ImplementationType = "Reported"
)

type ReadingBasisType string

const (
	// ZeroReadingBasisType shall indicate a reading with zero as its reference
	// point.
	ZeroReadingBasisType ReadingBasisType = "Zero"
	// DeltaReadingBasisType shall indicate a reading that reports the difference
	// between two measurements.
	DeltaReadingBasisType ReadingBasisType = "Delta"
	// HeadroomReadingBasisType shall indicate a reading that decreases in value as
	// it approaches the reference point. If the value crosses the reference point,
	// the value may be reported as a negative number or may report a value of
	// zero.
	HeadroomReadingBasisType ReadingBasisType = "Headroom"
)

type ReadingType string

const (
	// TemperatureReadingType shall indicate a temperature measurement, in degree
	// Celsius units. The 'ReadingUnits' property shall contain 'Cel'.
	TemperatureReadingType ReadingType = "Temperature"
	// HumidityReadingType shall indicate a relative humidity measurement, in
	// percent units. The 'ReadingUnits' property shall contain '%'.
	HumidityReadingType ReadingType = "Humidity"
	// PowerReadingType shall indicate the arithmetic mean of product terms of
	// instantaneous voltage and current values measured over integer number of
	// line cycles for a circuit, in watt units. The 'ReadingUnits' property shall
	// contain 'W'.
	PowerReadingType ReadingType = "Power"
	// EnergykWhReadingType shall indicate the energy, integral of real power over
	// time, of the monitored item. If representing metered power consumption the
	// value shall reflect the power consumption since the sensor metrics were last
	// reset. The value of the 'Reading' property shall be in kilowatt-hour units
	// and the 'ReadingUnits' property shall contain 'kW.h'. This value is used for
	// large-scale energy consumption measurements, while 'EnergyJoules' and
	// 'EnergyWh' are used for device-level consumption measurements.
	EnergykWhReadingType ReadingType = "EnergykWh"
	// EnergyJoulesReadingType shall indicate the energy, integral of real power
	// over time, of the monitored item. If representing metered power consumption
	// the value shall reflect the power consumption since the sensor metrics were
	// last reset. The value of the 'Reading' property shall be in joule units and
	// the 'ReadingUnits' property shall contain 'J'. This value is used for
	// device-level energy consumption measurements, while 'EnergykWh' is used for
	// large-scale consumption measurements.
	EnergyJoulesReadingType ReadingType = "EnergyJoules"
	// EnergyWhReadingType shall indicate the energy, integral of real power over
	// time, of the monitored item. If representing metered power consumption the
	// value shall reflect the power consumption since the sensor metrics were last
	// reset. The value of the 'Reading' property shall be in watt-hour units and
	// the 'ReadingUnits' property shall contain 'W.h'. This value is used for
	// device-level energy consumption measurements, while 'EnergykWh' is used for
	// large-scale consumption measurements.
	EnergyWhReadingType ReadingType = "EnergyWh"
	// ChargeAhReadingType shall indicate the amount of charge, integral of current
	// over time, of the monitored item. If representing metered charge consumption
	// the value shall reflect the charge consumption since the sensor metrics were
	// last reset. The value of the 'Reading' property shall be in ampere-hour
	// units and the 'ReadingUnits' property shall contain 'A.h'.
	ChargeAhReadingType ReadingType = "ChargeAh"
	// VoltageReadingType shall indicate a measurement of the root mean square
	// (RMS) of instantaneous voltage calculated over an integer number of line
	// cycles for a circuit. Voltage is expressed in volt units and the
	// 'ReadingUnits' property shall contain 'V'.
	VoltageReadingType ReadingType = "Voltage"
	// CurrentReadingType shall indicate a measurement of the root mean square
	// (RMS) of instantaneous current calculated over an integer number of line
	// cycles for a circuit. Current is expressed in ampere units and the
	// 'ReadingUnits' property shall contain 'A'.
	CurrentReadingType ReadingType = "Current"
	// FrequencyReadingType shall indicate a frequency measurement, in hertz units.
	// The 'ReadingUnits' property shall contain 'Hz'.
	FrequencyReadingType ReadingType = "Frequency"
	// PressureReadingType shall indicate a measurement of force, in pascal units,
	// applied perpendicular to the surface of an object per unit area over which
	// that force is distributed. The 'ReadingUnits' property shall contain 'Pa'.
	PressureReadingType ReadingType = "Pressure"
	// PressurekPaReadingType shall indicate a measurement of pressure, in
	// kilopascal units, relative to atmospheric pressure. The 'ReadingUnits'
	// property shall contain 'kPa'.
	PressurekPaReadingType ReadingType = "PressurekPa"
	// PressurePaReadingType shall indicate a measurement of pressure, in pascal
	// units, relative to atmospheric pressure. The 'ReadingUnits' property shall
	// contain 'Pa'.
	PressurePaReadingType ReadingType = "PressurePa"
	// LiquidLevelReadingType shall indicate a measurement of fluid height, in
	// centimeter units, relative to a specified vertical datum and the
	// 'ReadingUnits' property shall contain 'cm'.
	LiquidLevelReadingType ReadingType = "LiquidLevel"
	// RotationalReadingType shall indicate a measurement of rotational frequency,
	// in revolutions per minute units. The 'ReadingUnits' property shall contain
	// either '{rev}/min', which is preferred, or 'RPM', which is a deprecated
	// value. Services should represent fan speed and pump speed sensors with the
	// 'ReadingType' value 'Percent'.
	RotationalReadingType ReadingType = "Rotational"
	// AirFlowReadingType shall indicate a measurement of a volume of gas per unit
	// of time, in cubic feet per minute units, that flows through a particular
	// junction. The 'ReadingUnits' property shall contain '[ft_i]3/min'.
	AirFlowReadingType ReadingType = "AirFlow"
	// AirFlowCMMReadingType shall indicate a measurement of a volume of gas per
	// unit of time, in cubic meters per minute units, that flows through a
	// particular junction. The 'ReadingUnits' property shall contain 'm3/min'.
	AirFlowCMMReadingType ReadingType = "AirFlowCMM"
	// LiquidFlowReadingType shall indicate a measurement of a volume of liquid per
	// unit of time, in liters per second units, that flows through a particular
	// junction. The 'ReadingUnits' property shall contain 'L/s'.
	LiquidFlowReadingType ReadingType = "LiquidFlow"
	// LiquidFlowLPMReadingType shall indicate a measurement of a volume of liquid
	// per unit of time, in liters per minute units, that flows through a
	// particular junction. The 'ReadingUnits' property shall contain 'L/min'.
	LiquidFlowLPMReadingType ReadingType = "LiquidFlowLPM"
	// BarometricReadingType shall indicate a measurement of barometric pressure,
	// in millimeters of a mercury column. The 'ReadingUnits' property shall
	// contain 'mm[Hg]'.
	BarometricReadingType ReadingType = "Barometric"
	// AltitudeReadingType shall indicate a measurement of altitude, in meter
	// units, defined as the elevation above sea level. The 'ReadingUnits' property
	// shall contain 'm'.
	AltitudeReadingType ReadingType = "Altitude"
	// PercentReadingType shall indicate a percentage measurement, in percent
	// units. The 'Reading' value, while typically '0' to '100', may exceed '100'
	// for rate-of-change or similar readings. The 'ReadingUnits' property shall
	// contain '%'.
	PercentReadingType ReadingType = "Percent"
	// AbsoluteHumidityReadingType shall indicate an absolute (volumetric) humidity
	// measurement, in grams per cubic meter units. The 'ReadingUnits' property
	// shall contain 'g/m3'.
	AbsoluteHumidityReadingType ReadingType = "AbsoluteHumidity"
	// HeatReadingType shall indicate a heat measurement, in kilowatt units. The
	// 'ReadingUnits' property shall contain 'kW'.
	HeatReadingType ReadingType = "Heat"
	// LinearPositionReadingType shall indicate a linear position or distance, in
	// meter units. The 'ReadingUnits' property shall contain 'm'.
	LinearPositionReadingType ReadingType = "LinearPosition"
	// LinearVelocityReadingType shall indicate a linear velocity, in meters per
	// second units. The 'ReadingUnits' property shall contain 'm/s'.
	LinearVelocityReadingType ReadingType = "LinearVelocity"
	// LinearAccelerationReadingType shall indicate a linear acceleration, in
	// meters per square second units. The 'ReadingUnits' property shall contain
	// 'm/s2'.
	LinearAccelerationReadingType ReadingType = "LinearAcceleration"
	// RotationalPositionReadingType shall indicate a rotational position, in
	// radian units. The 'ReadingUnits' property shall contain 'rad'.
	RotationalPositionReadingType ReadingType = "RotationalPosition"
	// RotationalVelocityReadingType shall indicate a rotational velocity, in
	// radians per second units. The 'ReadingUnits' property shall contain 'rad/s'.
	RotationalVelocityReadingType ReadingType = "RotationalVelocity"
	// RotationalAccelerationReadingType shall indicate a rotational acceleration,
	// in radians per square second units. The 'ReadingUnits' property shall
	// contain 'rad/s2'.
	RotationalAccelerationReadingType ReadingType = "RotationalAcceleration"
	// ValveReadingType shall indicate a valve position, in percent units. The
	// 'ReadingUnits' property shall contain '%'. A value of '100' shall indicate
	// the valve is completely open, and a value of '0' shall indicate the valve is
	// completely closed.
	ValveReadingType ReadingType = "Valve"
)

type ThresholdActivation string

const (
	// IncreasingThresholdActivation This threshold is activated when the reading
	// changes from a value lower than the threshold to a value higher than the
	// threshold.
	IncreasingThresholdActivation ThresholdActivation = "Increasing"
	// DecreasingThresholdActivation This threshold is activated when the reading
	// changes from a value higher than the threshold to a value lower than the
	// threshold.
	DecreasingThresholdActivation ThresholdActivation = "Decreasing"
	// EitherThresholdActivation This threshold is activated when either the
	// increasing or decreasing conditions are met.
	EitherThresholdActivation ThresholdActivation = "Either"
	// DisabledThresholdActivation shall indicate the threshold is disabled and no
	// actions shall be taken as a result of the reading crossing the threshold
	// value.
	DisabledThresholdActivation ThresholdActivation = "Disabled"
)

type SensorVoltageType string

const (
	// ACSensorVoltageType Alternating current.
	ACSensorVoltageType SensorVoltageType = "AC"
	// DCSensorVoltageType Direct current.
	DCSensorVoltageType SensorVoltageType = "DC"
)

// Sensor shall represent a sensor for a Redfish implementation.
type Sensor struct {
	Entity
	// Accuracy shall contain the percent error +/- of the measured versus actual
	// values of the 'Reading' property.
	//
	// Deprecated: v1.8.0
	// This property has been deprecated in favor of ReadingAccuracy.
	Accuracy *float64 `json:",omitempty"`
	// AdjustedMaxAllowableOperatingValue shall contain the adjusted maximum
	// allowable operating value for the equipment that this sensor monitors, as
	// specified by a standards body, manufacturer, or both. The value is adjusted
	// based on environmental conditions. For example, liquid inlet temperature can
	// be adjusted based on the available liquid pressure.
	AdjustedMaxAllowableOperatingValue *float64 `json:",omitempty"`
	// AdjustedMinAllowableOperatingValue shall contain the adjusted minimum
	// allowable operating value for the equipment that this sensor monitors, as
	// specified by a standards body, manufacturer, or both. This value is adjusted
	// based on environmental conditions. For example, liquid inlet temperature can
	// be adjusted based on the available liquid pressure.
	AdjustedMinAllowableOperatingValue *float64 `json:",omitempty"`
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current
	// (RMS) for a circuit. This property can appear in sensors of the 'Power'
	// 'ReadingType', and shall not appear in sensors of other 'ReadingType'
	// values.
	ApparentVA *float64 `json:",omitempty"`
	// ApparentkVAh shall contain the apparent energy, in kilovolt-ampere-hour
	// units, for an electrical energy measurement. This property can appear in
	// sensors with a 'ReadingType' containing 'EnergykWh', and shall not appear in
	// sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	ApparentkVAh *float64 `json:",omitempty"`
	// AverageReading shall contain the average sensor value over the time
	// specified by the value of the 'AveragingInterval' property. The value shall
	// be reset by the 'ResetMetrics' action or by a service reset of time-based
	// property values.
	//
	// Version added: v1.4.0
	AverageReading *float64 `json:",omitempty"`
	// AveragingInterval shall contain the interval over which the sensor value is
	// averaged to produce the value of the 'AverageReading' property. This
	// property shall only be present if the 'AverageReading' property is present.
	//
	// Version added: v1.4.0
	AveragingInterval string
	// AveragingIntervalAchieved shall indicate that enough readings were collected
	// to calculate the 'AverageReading' value over the interval specified by the
	// 'AveragingInterval' property. The value shall be reset by the 'ResetMetrics'
	// action. This property shall only be present if the 'AveragingInterval'
	// property is present.
	//
	// Version added: v1.4.0
	AveragingIntervalAchieved bool
	// Calibration shall contain the offset applied to the raw sensor value to
	// provide a calibrated value for the sensor as returned by the 'Reading'
	// property. The value of this property shall follow the units of the 'Reading'
	// property for this sensor instance. Updating the value of this property shall
	// not affect the value of the 'CalibrationTime' property.
	//
	// Version added: v1.4.0
	Calibration *float64 `json:",omitempty"`
	// CalibrationTime shall contain the date and time that the sensor was last
	// calibrated. This property is intended to reflect the actual time the
	// calibration occurred.
	//
	// Version added: v1.4.0
	CalibrationTime string
	// CrestFactor shall contain the ratio of the peak measurement divided by the
	// RMS measurement and calculated over same N line cycles. A sine wave would
	// have a value of 1.414.
	//
	// Version added: v1.1.0
	CrestFactor *float64 `json:",omitempty"`
	// ElectricalContext shall represent the combination of current-carrying
	// conductors that distribute power.
	ElectricalContext ElectricalContext
	// Enabled shall indicate whether the sensor is enabled and provides a
	// 'Reading'. The value 'true' shall indicate the sensor is enabled and returns
	// the 'Reading' property with a valid value. The value 'false' shall indicate
	// the sensor is disabled, shall not return the 'Reading' property, and shall
	// not trigger events, logging, or other functionality. This property allows a
	// user to disable a faulty sensor or to otherwise remove it from use.
	//
	// Version added: v1.10.0
	Enabled bool
	// Implementation shall contain the implementation of the sensor.
	//
	// Version added: v1.1.0
	Implementation ImplementationType
	// LifetimeReading shall contain the total accumulation of the 'Reading'
	// property over the sensor's lifetime. This value shall not be reset by the
	// 'ResetMetrics' action.
	//
	// Version added: v1.1.0
	LifetimeReading *float64 `json:",omitempty"`
	// LifetimeStartDateTime shall contain the date and time when the sensor
	// started accumulating readings for the 'LifetimeReading' property. This might
	// contain the same value as the production date of the device that contains
	// this sensor.
	//
	// Version added: v1.9.0
	LifetimeStartDateTime string
	// LoadPercent shall indicate the power load utilization percent for this
	// sensor. This property can appear in sensors of the 'Power' 'ReadingType',
	// and shall not appear in sensors of other 'ReadingType' values.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of using a sensor instance with a
	// 'ReadingType' of 'Percent' to show utilization values when needed.
	LoadPercent *float64 `json:",omitempty"`
	// Location shall indicate the location information for this sensor.
	Location Location
	// LowestIntervalReading shall contain the lowest sensor value over the time
	// specified by the value of the 'AveragingInterval' property. The value shall
	// be reset by the 'ResetMetrics' action or by a service reset of time-based
	// property values.
	//
	// Version added: v1.12.0
	LowestIntervalReading *float64 `json:",omitempty"`
	// LowestReading shall contain the lowest sensor value since the last
	// 'ResetMetrics' action was performed or since the service last reset the
	// time-based property values.
	//
	// Version added: v1.4.0
	LowestReading *float64 `json:",omitempty"`
	// LowestReadingTime shall contain the date and time when the lowest sensor
	// value was observed, as reported as the value of 'LowestReading'.
	//
	// Version added: v1.4.0
	LowestReadingTime string
	// Manufacturer shall contain the name of the organization responsible for
	// producing the sensor. This organization may be the entity from whom the
	// sensor is purchased, but this is not necessarily true. This property is
	// generally used only for replaceable or user-configurable sensors.
	//
	// Version added: v1.9.0
	Manufacturer string
	// MaxAllowableOperatingValue shall contain the maximum allowable operating
	// value for the equipment that this sensor monitors, as specified by a
	// standards body, manufacturer, or both.
	MaxAllowableOperatingValue *float64 `json:",omitempty"`
	// MinAllowableOperatingValue shall contain the minimum allowable operating
	// value for the equipment that this sensor monitors, as specified by a
	// standards body, manufacturer, or both.
	MinAllowableOperatingValue *float64 `json:",omitempty"`
	// Model shall contain the name by which the manufacturer generally refers to
	// the sensor. This property is generally used only for replaceable or
	// user-configurable sensors.
	//
	// Version added: v1.9.0
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the sensor. This property is
	// generally used only for replaceable or user-configurable sensors.
	//
	// Version added: v1.9.0
	PartNumber string
	// PeakIntervalReading shall contain the peak sensor value over the time
	// specified by the value of the 'AveragingInterval' property. The value shall
	// be reset by the 'ResetMetrics' action or by a service reset of time-based
	// property values.
	//
	// Version added: v1.12.0
	PeakIntervalReading *float64 `json:",omitempty"`
	// PeakReading shall contain the peak sensor value since the last
	// 'ResetMetrics' action was performed or since the service last reset the
	// time-based property values.
	PeakReading *float64 `json:",omitempty"`
	// PeakReadingTime shall contain the date and time when the peak sensor value
	// was observed, as reported as the value of 'PeakReading'.
	PeakReadingTime string
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between
	// the current and voltage waveforms for an electrical measurement. This
	// property can appear in sensors with a 'ReadingType' containing 'Power', and
	// shall not appear in sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	PhaseAngleDegrees *float64 `json:",omitempty"`
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this sensor measurement applies. This property
	// generally differentiates multiple sensors within the same 'PhysicalContext'
	// instance.
	PhysicalSubContext PhysicalSubContext
	// PowerFactor shall identify the quotient of real power (W) and apparent power
	// (VA) for a circuit. 'PowerFactor' is expressed in unit-less 1/100ths. This
	// property can appear in sensors containing a 'ReadingType' value of 'Power',
	// and shall not appear in sensors of other 'ReadingType' values.
	PowerFactor *float64 `json:",omitempty"`
	// Precision shall contain the number of significant digits in the 'Reading'
	// property.
	Precision *float64 `json:",omitempty"`
	// ReactiveVAR shall contain the arithmetic mean of product terms of
	// instantaneous voltage and quadrature current measurements calculated over an
	// integer number of line cycles for a circuit. This property can appear in
	// sensors of the 'Power' 'ReadingType', and shall not appear in sensors of
	// other 'ReadingType' values.
	ReactiveVAR *float64 `json:",omitempty"`
	// ReactivekVARh shall contain the reactive energy, in kilovolt-ampere-hours
	// (reactive) units, for an electrical energy measurement. This property can
	// appear in sensors with a 'ReadingType' containing 'EnergykWh', and shall not
	// appear in sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	ReactivekVARh *float64 `json:",omitempty"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// ReadingAccuracy shall contain the accuracy of the value of the 'Reading'
	// property for this sensor. The value shall be the absolute value of the
	// maximum deviation of the 'Reading' from its actual value. The value shall be
	// in units that follow the 'ReadingUnits' for this sensor.
	//
	// Version added: v1.8.0
	ReadingAccuracy *float64 `json:",omitempty"`
	// ReadingBasis shall indicate the basis or frame of reference for the value of
	// the 'Reading' property. If this property is not present, the value shall be
	// assumed to be 'Zero'.
	//
	// Version added: v1.7.0
	ReadingBasis ReadingBasisType
	// ReadingRangeMax shall indicate the maximum possible value of the 'Reading'
	// property for this sensor. This value is the range of valid readings for this
	// sensor. Values outside this range are discarded as reading errors.
	ReadingRangeMax *float64 `json:",omitempty"`
	// ReadingRangeMin shall indicate the minimum possible value of the 'Reading'
	// property for this sensor. This value is the range of valid readings for this
	// sensor. Values outside this range are discarded as reading errors.
	ReadingRangeMin *float64 `json:",omitempty"`
	// ReadingTime shall contain the date and time that the reading data was
	// acquired from the sensor. This value is used to synchronize readings from
	// multiple sensors and does not represent the time at which the resource was
	// accessed.
	//
	// Version added: v1.1.0
	ReadingTime string
	// ReadingType shall contain the type of the sensor.
	ReadingType ReadingType
	// ReadingUnits shall contain the units of the sensor's reading, thresholds,
	// and other reading-related properties. The value shall follow the
	// case-sensitive symbol format defined by the Unified Code for Units of
	// Measure (UCUM), as specified by the 'Units of measure annotation' clause of
	// the Redfish Specification.
	ReadingUnits string
	// RelatedItem shall contain an array of links to resources or objects that
	// this sensor services.
	//
	// Version added: v1.2.0
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SKU shall contain the stock-keeping unit number for this sensor. This
	// property is generally used only for replaceable or user-configurable
	// sensors.
	//
	// Version added: v1.9.0
	SKU string
	// SensingFrequency shall contain the time interval between readings of the
	// physical sensor.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of the 'SensingInterval'
	// property, which uses the duration time format for interoperability.
	SensingFrequency *float64 `json:",omitempty"`
	// SensingInterval shall contain the time interval between readings of data
	// from the sensor.
	//
	// Version added: v1.1.0
	SensingInterval string
	// SensorGroup shall contain information for a group of sensors that provide
	// input for the value of this sensor's reading. If this property is present,
	// the 'Implementation' property shall contain the value 'Synthesized'. The
	// group may be created for redundancy or to improve the accuracy of the
	// reading through multiple sensor inputs.
	//
	// Version added: v1.4.0
	SensorGroup RedundantGroup
	// SensorResetTime shall contain the date and time when the 'ResetMetrics'
	// action was last performed or when the service last reset the time-based
	// property values.
	SensorResetTime string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the sensor. This property is generally used only for replaceable or
	// user-configurable sensors.
	//
	// Version added: v1.9.0
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the sensor. This
	// property is generally used only for replaceable or user-configurable
	// sensors.
	//
	// Version added: v1.9.0
	SparePartNumber string
	// SpeedRPM shall contain a reading of the rotational speed of the device in
	// revolutions per minute (RPM) units.
	//
	// Version added: v1.2.0
	SpeedRPM *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// THDPercent shall contain the total harmonic distortion of the 'Reading'
	// property in percent units, typically '0' to '100'.
	//
	// Version added: v1.1.0
	THDPercent *float64 `json:",omitempty"`
	// Thresholds shall contain the set of thresholds that derive a sensor's health
	// and operational range.
	Thresholds Thresholds
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. This property shall only be present if the sensor can be
	// configured for different purposes, or is dependent on configuration or
	// end-user settings. This property shall not be present for embedded sensors
	// with defined functions that cannot be altered. If a value has not been
	// assigned by a user, the value of this property shall be an empty string.
	//
	// Version added: v1.9.0
	UserLabel string
	// VoltageType shall represent the type of input voltage the sensor monitors.
	VoltageType SensorVoltageType
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// associatedControls are the URIs for AssociatedControls.
	associatedControls []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Sensor object from the raw JSON.
func (s *Sensor) UnmarshalJSON(b []byte) error {
	type temp Sensor
	type sActions struct {
		ResetMetrics    ActionTarget `json:"#Sensor.ResetMetrics"`
		ResetToDefaults ActionTarget `json:"#Sensor.ResetToDefaults"`
	}
	type sLinks struct {
		AssociatedControls Links `json:"AssociatedControls"`
	}
	var tmp struct {
		temp
		Actions     sActions
		Links       sLinks
		RelatedItem Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Sensor(tmp.temp)

	// Extract the links to other entities for later
	s.resetMetricsTarget = tmp.Actions.ResetMetrics.Target
	s.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	s.associatedControls = tmp.Links.AssociatedControls.ToStrings()
	s.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *Sensor) Update() error {
	readWriteFields := []string{
		"AveragingInterval",
		"Calibration",
		"CalibrationTime",
		"Enabled",
		"PhysicalContext",
		"PhysicalSubContext",
		"RelatedItem",
		"UserLabel",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSensor will get a Sensor instance from the service.
func GetSensor(c Client, uri string) (*Sensor, error) {
	return GetObject[Sensor](c, uri)
}

// ListReferencedSensors gets the collection of Sensor from
// a provided reference.
func ListReferencedSensors(c Client, link string) ([]*Sensor, error) {
	return GetCollectionObjects[Sensor](c, link)
}

// This action shall reset any time intervals or counted values for this
// sensor. The 'SensorResetTime' property shall be updated to reflect the time
// that this action was performed.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Sensor) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetMetricsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Sensor) ResetToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetToDefaultsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AssociatedControls gets the AssociatedControls linked resources.
func (s *Sensor) AssociatedControls() ([]*Control, error) {
	return GetObjects[Control](s.client, s.associatedControls)
}

// RelatedItem gets the RelatedItem linked resources.
func (s *Sensor) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](s.client, s.relatedItem)
}

// SensorArrayExcerpt shall represent a sensor for a Redfish implementation.
type SensorArrayExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// DeviceName shall contain the name of the device associated with this sensor.
	// If the device is represented by a resource, the value shall contain the
	// value of the 'Name' property of the associated resource.
	//
	// Version added: v1.2.0
	DeviceName string
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this sensor measurement applies. This property
	// generally differentiates multiple sensors within the same 'PhysicalContext'
	// instance.
	PhysicalSubContext PhysicalSubContext
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
}

// SensorCurrentExcerpt shall represent a sensor for a Redfish implementation.
type SensorCurrentExcerpt struct {
	// CrestFactor shall contain the ratio of the peak measurement divided by the
	// RMS measurement and calculated over same N line cycles. A sine wave would
	// have a value of 1.414.
	//
	// Version added: v1.1.0
	CrestFactor *float64 `json:",omitempty"`
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// THDPercent shall contain the total harmonic distortion of the 'Reading'
	// property in percent units, typically '0' to '100'.
	//
	// Version added: v1.1.0
	THDPercent *float64 `json:",omitempty"`
}

// SensorEnergykWhExcerpt shall represent a sensor for a Redfish implementation.
type SensorEnergykWhExcerpt struct {
	// ApparentkVAh shall contain the apparent energy, in kilovolt-ampere-hour
	// units, for an electrical energy measurement. This property can appear in
	// sensors with a 'ReadingType' containing 'EnergykWh', and shall not appear in
	// sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	ApparentkVAh *float64 `json:",omitempty"`
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// LifetimeReading shall contain the total accumulation of the 'Reading'
	// property over the sensor's lifetime. This value shall not be reset by the
	// 'ResetMetrics' action.
	//
	// Version added: v1.1.0
	LifetimeReading *float64 `json:",omitempty"`
	// ReactivekVARh shall contain the reactive energy, in kilovolt-ampere-hours
	// (reactive) units, for an electrical energy measurement. This property can
	// appear in sensors with a 'ReadingType' containing 'EnergykWh', and shall not
	// appear in sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	ReactivekVARh *float64 `json:",omitempty"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// SensorResetTime shall contain the date and time when the 'ResetMetrics'
	// action was last performed or when the service last reset the time-based
	// property values.
	SensorResetTime string
}

// SensorExcerpt shall represent a sensor for a Redfish implementation.
type SensorExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
}

// SensorFanArrayExcerpt shall represent a sensor for a Redfish implementation.
type SensorFanArrayExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// DeviceName shall contain the name of the device associated with this sensor.
	// If the device is represented by a resource, the value shall contain the
	// value of the 'Name' property of the associated resource.
	//
	// Version added: v1.2.0
	DeviceName string
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this sensor measurement applies. This property
	// generally differentiates multiple sensors within the same 'PhysicalContext'
	// instance.
	PhysicalSubContext PhysicalSubContext
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// SpeedRPM shall contain a reading of the rotational speed of the device in
	// revolutions per minute (RPM) units.
	//
	// Version added: v1.2.0
	SpeedRPM *float64 `json:",omitempty"`
}

// SensorFanExcerpt shall represent a sensor for a Redfish implementation.
type SensorFanExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// SpeedRPM shall contain a reading of the rotational speed of the device in
	// revolutions per minute (RPM) units.
	//
	// Version added: v1.2.0
	SpeedRPM *float64 `json:",omitempty"`
}

// SensorPowerArrayExcerpt shall represent a sensor for a Redfish
// implementation.
type SensorPowerArrayExcerpt struct {
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current
	// (RMS) for a circuit. This property can appear in sensors of the 'Power'
	// 'ReadingType', and shall not appear in sensors of other 'ReadingType'
	// values.
	ApparentVA *float64 `json:",omitempty"`
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between
	// the current and voltage waveforms for an electrical measurement. This
	// property can appear in sensors with a 'ReadingType' containing 'Power', and
	// shall not appear in sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	PhaseAngleDegrees *float64 `json:",omitempty"`
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this sensor measurement applies. This property
	// generally differentiates multiple sensors within the same 'PhysicalContext'
	// instance.
	PhysicalSubContext PhysicalSubContext
	// PowerFactor shall identify the quotient of real power (W) and apparent power
	// (VA) for a circuit. 'PowerFactor' is expressed in unit-less 1/100ths. This
	// property can appear in sensors containing a 'ReadingType' value of 'Power',
	// and shall not appear in sensors of other 'ReadingType' values.
	PowerFactor *float64 `json:",omitempty"`
	// ReactiveVAR shall contain the arithmetic mean of product terms of
	// instantaneous voltage and quadrature current measurements calculated over an
	// integer number of line cycles for a circuit. This property can appear in
	// sensors of the 'Power' 'ReadingType', and shall not appear in sensors of
	// other 'ReadingType' values.
	ReactiveVAR *float64 `json:",omitempty"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
}

// SensorPowerExcerpt shall represent a sensor for a Redfish implementation.
type SensorPowerExcerpt struct {
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current
	// (RMS) for a circuit. This property can appear in sensors of the 'Power'
	// 'ReadingType', and shall not appear in sensors of other 'ReadingType'
	// values.
	ApparentVA *float64 `json:",omitempty"`
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between
	// the current and voltage waveforms for an electrical measurement. This
	// property can appear in sensors with a 'ReadingType' containing 'Power', and
	// shall not appear in sensors with other 'ReadingType' values.
	//
	// Version added: v1.5.0
	PhaseAngleDegrees *float64 `json:",omitempty"`
	// PowerFactor shall identify the quotient of real power (W) and apparent power
	// (VA) for a circuit. 'PowerFactor' is expressed in unit-less 1/100ths. This
	// property can appear in sensors containing a 'ReadingType' value of 'Power',
	// and shall not appear in sensors of other 'ReadingType' values.
	PowerFactor *float64 `json:",omitempty"`
	// ReactiveVAR shall contain the arithmetic mean of product terms of
	// instantaneous voltage and quadrature current measurements calculated over an
	// integer number of line cycles for a circuit. This property can appear in
	// sensors of the 'Power' 'ReadingType', and shall not appear in sensors of
	// other 'ReadingType' values.
	ReactiveVAR *float64 `json:",omitempty"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
}

// SensorPumpExcerpt shall represent a sensor for a Redfish implementation.
type SensorPumpExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// SpeedRPM shall contain a reading of the rotational speed of the device in
	// revolutions per minute (RPM) units.
	//
	// Version added: v1.2.0
	SpeedRPM *float64 `json:",omitempty"`
}

// SensorVoltageExcerpt shall represent a sensor for a Redfish implementation.
type SensorVoltageExcerpt struct {
	// CrestFactor shall contain the ratio of the peak measurement divided by the
	// RMS measurement and calculated over same N line cycles. A sine wave would
	// have a value of 1.414.
	//
	// Version added: v1.1.0
	CrestFactor *float64 `json:",omitempty"`
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the sensor value. This property shall not be returned
	// if the 'Enabled' property is supported and contains 'false'.
	Reading *float64 `json:",omitempty"`
	// THDPercent shall contain the total harmonic distortion of the 'Reading'
	// property in percent units, typically '0' to '100'.
	//
	// Version added: v1.1.0
	THDPercent *float64 `json:",omitempty"`
}

// Threshold shall contain the properties for an individual threshold for this
// sensor.
type Threshold struct {
	// Activation shall indicate the direction of crossing of the reading for this
	// sensor that activates the threshold.
	Activation ThresholdActivation
	// DwellTime shall indicate the duration the sensor value violates the
	// threshold before the threshold is activated.
	DwellTime string
	// HysteresisDuration shall indicate the duration the sensor value no longer
	// violates the threshold before the threshold is deactivated. A duration of
	// zero seconds, or if the property is not present in the resource, shall
	// indicate the threshold is deactivated immediately once the sensor value no
	// longer violates the threshold. The threshold shall not deactivate until the
	// conditions of both 'HysteresisReading' and 'HysteresisDuration' are met.
	//
	// Version added: v1.7.0
	HysteresisDuration string
	// HysteresisReading shall indicate the offset from the reading for this sensor
	// and the threshold value that deactivates the threshold. For example, a value
	// of '-2' indicates the sensor reading shall fall 2 units below an upper
	// threshold value to deactivate the threshold. The value of the property shall
	// use the same units as the 'Reading' property. A value of '0', or if the
	// property is not present in the resource, shall indicate the threshold is
	// deactivated when the sensor value no longer violates the threshold. The
	// threshold shall not deactivate until the conditions of both
	// 'HysteresisReading' and 'HysteresisDuration' are met.
	//
	// Version added: v1.7.0
	HysteresisReading *float64 `json:",omitempty"`
	// Reading shall indicate the reading for this sensor that activates the
	// threshold. The value of the property shall use the same units as the
	// 'Reading' property.
	Reading *float64 `json:",omitempty"`
}

// Thresholds shall contain the set of thresholds that derive a sensor's health
// and operational range.
type Thresholds struct {
	// LowerCaution shall contain the value at which the 'Reading' property is
	// below normal range. The value of the property shall use the same units as
	// the 'Reading' property.
	LowerCaution Threshold
	// LowerCautionUser shall contain a user-defined value at which the 'Reading'
	// property is considered below the normal range. The value of the property
	// shall use the same units as the 'Reading' property. The 'Reading' property
	// shall be considered below normal range if either the 'LowerCaution' or
	// 'LowerCautionUser' threshold has been violated. This property is used to
	// provide an additional, user-defined threshold value when the 'LowerCaution'
	// threshold is implemented as read-only to reflect a service-defined value
	// that cannot be changed.
	//
	// Version added: v1.2.0
	LowerCautionUser Threshold
	// LowerCritical shall contain the value at which the 'Reading' property is
	// below the normal range but is not yet fatal. The value of the property shall
	// use the same units as the 'Reading' property.
	LowerCritical Threshold
	// LowerCriticalUser shall contain a user-defined value at which the 'Reading'
	// property is considered below the normal range but is not yet fatal. The
	// value of the property shall use the same units as the 'Reading' property.
	// The 'Reading' property shall be considered below normal range if either the
	// 'LowerCritical' or 'LowerCriticalUser' threshold has been violated. This
	// property is used to provide an additional, user-defined threshold value when
	// the 'LowerCritical' threshold is implemented as read-only to reflect a
	// service-defined value that cannot be changed.
	//
	// Version added: v1.2.0
	LowerCriticalUser Threshold
	// LowerFatal shall contain the value at which the 'Reading' property is below
	// the normal range and is fatal. The value of the property shall use the same
	// units as the 'Reading' property.
	LowerFatal Threshold
	// UpperCaution shall contain the value at which the 'Reading' property is
	// above the normal range. The value of the property shall use the same units
	// as the 'Reading' property.
	UpperCaution Threshold
	// UpperCautionUser shall contain a user-defined value at which the 'Reading'
	// property is considered above the normal range. The value of the property
	// shall use the same units as the 'Reading' property. The 'Reading' property
	// shall be considered above normal range if either the 'UpperCaution' or
	// 'UpperCautionUser' threshold has been violated. This property is used to
	// provide an additional, user-defined threshold value when the 'UpperCaution'
	// threshold is implemented as read-only to reflect a service-defined value
	// that cannot be changed.
	//
	// Version added: v1.2.0
	UpperCautionUser Threshold
	// UpperCritical shall contain the value at which the 'Reading' property is
	// above the normal range but is not yet fatal. The value of the property shall
	// use the same units as the 'Reading' property.
	UpperCritical Threshold
	// UpperCriticalUser shall contain a user-defined value at which the 'Reading'
	// property is considered above the normal range but is not yet fatal. The
	// value of the property shall use the same units as the 'Reading' property.
	// The 'Reading' property shall be considered above normal range if either the
	// 'UpperCritical' or 'UpperCriticalUser' threshold has been violated. This
	// property is used to provide an additional, user-defined threshold value when
	// the 'UpperCritical' threshold is implemented as read-only to reflect a
	// service-defined value that cannot be changed.
	//
	// Version added: v1.2.0
	UpperCriticalUser Threshold
	// UpperFatal shall contain the value at which the 'Reading' property is above
	// the normal range and is fatal. The value of the property shall use the same
	// units as the 'Reading' property.
	UpperFatal Threshold
}
