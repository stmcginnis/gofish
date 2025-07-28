//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

// ImplementationType specifies the method used to obtain the sensor reading.
type ImplementationType string

const (
	// ImplementationTypePhysicalSensor indicates the reading comes from a physical sensor.
	ImplementationTypePhysicalSensor ImplementationType = "PhysicalSensor"
	// ImplementationTypeSynthesized indicates the reading is calculated from multiple sensors.
	ImplementationTypeSynthesized ImplementationType = "Synthesized"
	// ImplementationTypeReported indicates the reading comes from software or device reporting.
	ImplementationTypeReported ImplementationType = "Reported"
)

// ReadingType specifies the type of physical quantity being measured.
type ReadingType string

const (
	// ReadingTypeAbsoluteHumidity measures absolute humidity in g/m^3.
	ReadingTypeAbsoluteHumidity ReadingType = "AbsoluteHumidity"
	// ReadingTypeAirFlow measures airflow in cu ft/min (deprecated in v1.7.0).
	ReadingTypeAirFlow ReadingType = "AirFlow"
	// ReadingTypeAirFlowCMM measures airflow in m^3/min.
	ReadingTypeAirFlowCMM ReadingType = "AirFlowCMM"
	// ReadingTypeAltitude measures altitude in meters.
	ReadingTypeAltitude ReadingType = "Altitude"
	// ReadingTypeBarometric measures barometric pressure in mmHg.
	ReadingTypeBarometric ReadingType = "Barometric"
	// ReadingTypeChargeAh measures charge in ampere-hours.
	ReadingTypeChargeAh ReadingType = "ChargeAh"
	// ReadingTypeCurrent measures current in amperes.
	ReadingTypeCurrent ReadingType = "Current"
	// ReadingTypeEnergyJoules measures energy in joules.
	ReadingTypeEnergyJoules ReadingType = "EnergyJoules"
	// ReadingTypeEnergykWh measures energy in kilowatt-hours.
	ReadingTypeEnergykWh ReadingType = "EnergykWh"
	// ReadingTypeEnergyWh measures energy in watt-hours.
	ReadingTypeEnergyWh ReadingType = "EnergyWh"
	// ReadingTypeFrequency measures frequency in hertz.
	ReadingTypeFrequency ReadingType = "Frequency"
	// ReadingTypeHeat measures heat in kilowatts.
	ReadingTypeHeat ReadingType = "Heat"
	// ReadingTypeHumidity measures relative humidity in percent.
	ReadingTypeHumidity ReadingType = "Humidity"
	// ReadingTypeLinearAcceleration measures linear acceleration in m/s^2.
	ReadingTypeLinearAcceleration ReadingType = "LinearAcceleration"
	// ReadingTypeLinearPosition measures linear position in meters.
	ReadingTypeLinearPosition ReadingType = "LinearPosition"
	// ReadingTypeLinearVelocity measures linear velocity in m/s.
	ReadingTypeLinearVelocity ReadingType = "LinearVelocity"
	// ReadingTypeLiquidFlow measures liquid flow in L/s (deprecated in v1.7.0).
	ReadingTypeLiquidFlow ReadingType = "LiquidFlow"
	// ReadingTypeLiquidFlowLPM measures liquid flow in L/min.
	ReadingTypeLiquidFlowLPM ReadingType = "LiquidFlowLPM"
	// ReadingTypeLiquidLevel measures liquid level in cm.
	ReadingTypeLiquidLevel ReadingType = "LiquidLevel"
	// ReadingTypePercent measures percentage values.
	ReadingTypePercent ReadingType = "Percent"
	// ReadingTypePower measures power in watts.
	ReadingTypePower ReadingType = "Power"
	// ReadingTypePressure measures pressure in pascals (deprecated in v1.7.0).
	ReadingTypePressure ReadingType = "Pressure"
	// ReadingTypePressurekPa measures pressure in kilopascals.
	ReadingTypePressurekPa ReadingType = "PressurekPa"
	// ReadingTypePressurePa measures pressure in pascals.
	ReadingTypePressurePa ReadingType = "PressurePa"
	// ReadingTypeRotational measures rotational speed in RPM.
	ReadingTypeRotational ReadingType = "Rotational"
	// ReadingTypeRotationalAcceleration measures rotational acceleration in rad/s^2.
	ReadingTypeRotationalAcceleration ReadingType = "RotationalAcceleration"
	// ReadingTypeRotationalPosition measures rotational position in radians.
	ReadingTypeRotationalPosition ReadingType = "RotationalPosition"
	// ReadingTypeRotationalVelocity measures rotational velocity in rad/s.
	ReadingTypeRotationalVelocity ReadingType = "RotationalVelocity"
	// ReadingTypeTemperature measures temperature in Celsius.
	ReadingTypeTemperature ReadingType = "Temperature"
	// ReadingTypeValve measures valve position in percent open.
	ReadingTypeValve ReadingType = "Valve"
	// ReadingTypeVoltage measures voltage in volts.
	ReadingTypeVoltage ReadingType = "Voltage"
)

// ReadingBasisType specifies the reference basis for sensor readings.
type ReadingBasisType string

const (
	// ReadingBasisTypeDelta indicates a difference between measurements.
	ReadingBasisTypeDelta ReadingBasisType = "Delta"
	// ReadingBasisTypeHeadroom indicates decreasing value approaching reference.
	ReadingBasisTypeHeadroom ReadingBasisType = "Headroom"
	// ReadingBasisTypeZero indicates zero-based measurement.
	ReadingBasisTypeZero ReadingBasisType = "Zero"
)

// ThresholdActivation specifies how a threshold is triggered.
type ThresholdActivation string

const (
	// ThresholdActivationIncreasing triggers when value rises above threshold.
	ThresholdActivationIncreasing ThresholdActivation = "Increasing"
	// ThresholdActivationDecreasing triggers when value falls below threshold.
	ThresholdActivationDecreasing ThresholdActivation = "Decreasing"
	// ThresholdActivationEither triggers on any crossing direction.
	ThresholdActivationEither ThresholdActivation = "Either"
	// ThresholdActivationDisabled indicates the threshold is inactive.
	ThresholdActivationDisabled ThresholdActivation = "Disabled"
)

// Threshold defines parameters for sensor threshold monitoring.
type Threshold struct {
	// Activation specifies the crossing direction that triggers this threshold.
	Activation ThresholdActivation `json:"Activation,omitempty"`
	// DwellTime specifies how long the reading must violate the threshold before triggering.
	DwellTime *string `json:"DwellTime,omitempty"`
	// HysteresisDuration specifies how long the reading must comply before clearing the threshold.
	HysteresisDuration *string `json:"HysteresisDuration,omitempty"`
	// HysteresisReading specifies the offset required to clear the threshold.
	HysteresisReading *float32 `json:"HysteresisReading,omitempty"`
	// Reading specifies the threshold value to monitor.
	Reading *float32 `json:"Reading,omitempty"`
}

// Thresholds contains all threshold definitions for a sensor.
type Thresholds struct {
	// LowerCaution defines when reading is below normal range.
	LowerCaution *Threshold `json:"LowerCaution,omitempty"`
	// LowerCautionUser defines user-adjustable below-normal threshold.
	LowerCautionUser *Threshold `json:"LowerCautionUser,omitempty"`
	// LowerCritical defines when reading is below critical range.
	LowerCritical *Threshold `json:"LowerCritical,omitempty"`
	// LowerCriticalUser defines user-adjustable critical threshold.
	LowerCriticalUser *Threshold `json:"LowerCriticalUser,omitempty"`
	// LowerFatal defines when reading is below fatal range.
	LowerFatal *Threshold `json:"LowerFatal,omitempty"`
	// UpperCaution defines when reading is above normal range.
	UpperCaution *Threshold `json:"UpperCaution,omitempty"`
	// UpperCautionUser defines user-adjustable above-normal threshold.
	UpperCautionUser *Threshold `json:"UpperCautionUser,omitempty"`
	// UpperCritical defines when reading is above critical range.
	UpperCritical *Threshold `json:"UpperCritical,omitempty"`
	// UpperCriticalUser defines user-adjustable critical threshold.
	UpperCriticalUser *Threshold `json:"UpperCriticalUser,omitempty"`
	// UpperFatal defines when reading is above fatal range.
	UpperFatal *Threshold `json:"UpperFatal,omitempty"`
}

// SensorArrayExcerpt represents a sensor array summary.
type SensorArrayExcerpt struct {
	// DataSourceURI links to the source resource providing this data.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// DeviceName identifies the associated device.
	DeviceName string `json:"DeviceName,omitempty"`
	// PhysicalContext describes the measurement location.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`
	// PhysicalSubContext describes the sub-region being measured.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`
	// Reading contains the current sensor value.
	Reading *float64 `json:"Reading,omitempty"`
}

// SensorCurrentExcerpt represents current measurement summary.
type SensorCurrentExcerpt struct {
	// CrestFactor indicates peak-to-RMS ratio.
	CrestFactor *float64 `json:"CrestFactor,omitempty"`
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// THDPercent indicates total harmonic distortion.
	THDPercent *float64 `json:"THDPercent,omitempty"`
}

// SensorExcerpt represents basic sensor information.
type SensorExcerpt struct {
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// Reading contains the sensor measurement.
	Reading *float64 `json:"Reading,omitempty"`
}

// SensorFanArrayExcerpt represents fan array sensor summary.
type SensorFanArrayExcerpt struct {
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// DeviceName identifies the fan array.
	DeviceName string `json:"DeviceName,omitempty"`
	// PhysicalContext describes the measurement location.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`
	// PhysicalSubContext describes the sub-region being measured.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// SpeedRPM indicates rotational speed.
	SpeedRPM *float64 `json:"SpeedRPM,omitempty"`
}

// SensorFanExcerpt represents fan sensor summary.
type SensorFanExcerpt struct {
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// SpeedRPM indicates rotational speed.
	SpeedRPM *float64 `json:"SpeedRPM,omitempty"`
}

// SensorEnergykWhExcerpt represents energy consumption summary.
type SensorEnergykWhExcerpt struct {
	// ApparentkVAh indicates apparent energy consumption.
	ApparentkVAh *float64 `json:"ApparentkVAh,omitempty"`
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// LifetimeReading indicates total accumulated energy.
	LifetimeReading *float64 `json:"LifetimeReading,omitempty"`
	// ReactivekVARh indicates reactive energy consumption.
	ReactivekVARh *float64 `json:"ReactivekVARh,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// SensorResetTime indicates when metrics were last reset.
	SensorResetTime string `json:"SensorResetTime,omitempty"`
}

// SensorPowerArrayExcerpt represents power array summary.
type SensorPowerArrayExcerpt struct {
	// ApparentVA indicates apparent power.
	ApparentVA *float64 `json:"ApparentVA,omitempty"`
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// PhaseAngleDegrees indicates current-voltage phase difference.
	PhaseAngleDegrees *float64 `json:"PhaseAngleDegrees,omitempty"`
	// PhysicalContext describes the measurement location.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`
	// PhysicalSubContext describes the sub-region being measured.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`
	// PowerFactor indicates power efficiency.
	PowerFactor *float64 `json:"PowerFactor,omitempty"`
	// ReactiveVAR indicates reactive power.
	ReactiveVAR *float64 `json:"ReactiveVAR,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
}

// SensorPowerExcerpt represents power measurement summary.
type SensorPowerExcerpt struct {
	// ApparentVA indicates apparent power.
	ApparentVA *float64 `json:"ApparentVA,omitempty"`
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// PhaseAngleDegrees indicates current-voltage phase difference.
	PhaseAngleDegrees *float64 `json:"PhaseAngleDegrees,omitempty"`
	// PowerFactor indicates power efficiency.
	PowerFactor *float64 `json:"PowerFactor,omitempty"`
	// ReactiveVAR indicates reactive power.
	ReactiveVAR *float64 `json:"ReactiveVAR,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
}

// SensorPumpExcerpt represents pump sensor summary.
type SensorPumpExcerpt struct {
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// SpeedRPM indicates rotational speed.
	SpeedRPM *float64 `json:"SpeedRPM,omitempty"`
}

// SensorVoltageExcerpt represents voltage measurement summary.
type SensorVoltageExcerpt struct {
	// CrestFactor indicates peak-to-RMS ratio.
	CrestFactor *float64 `json:"CrestFactor,omitempty"`
	// DataSourceURI links to the source resource.
	DataSourceURI string `json:"DataSourceUri,omitempty"`
	// Reading contains the current measurement.
	Reading *float64 `json:"Reading,omitempty"`
	// THDPercent indicates total harmonic distortion.
	THDPercent *float64 `json:"THDPercent,omitempty"`
}

// Sensor represents a hardware or software sensor.
type Sensor struct {
	common.Entity

	// ODataContext is the OData context URL.
	ODataContext string `json:"@odata.context,omitempty"`
	// ODataType is the OData type identifier.
	ODataType string `json:"@odata.type,omitempty"`

	// Accuracy indicates measurement error percentage (deprecated in v1.8.0).
	Accuracy *float64 `json:"Accuracy,omitempty"`
	// AdjustedMaxAllowableOperatingValue accounts for environmental conditions.
	AdjustedMaxAllowableOperatingValue *float64 `json:"AdjustedMaxAllowableOperatingValue,omitempty"`
	// AdjustedMinAllowableOperatingValue accounts for environmental conditions.
	AdjustedMinAllowableOperatingValue *float64 `json:"AdjustedMinAllowableOperatingValue,omitempty"`
	// ApparentkVAh measures apparent energy in kVAh.
	ApparentkVAh *float64 `json:"ApparentkVAh,omitempty"`
	// ApparentVA measures apparent power in volt-amperes.
	ApparentVA *float64 `json:"ApparentVA,omitempty"`
	// AverageReading provides mean value over time.
	AverageReading *float64 `json:"AverageReading,omitempty"`
	// AveragingInterval specifies the averaging period.
	AveragingInterval string `json:"AveragingInterval,omitempty"`
	// AveragingIntervalAchieved indicates sufficient data collected.
	AveragingIntervalAchieved *bool `json:"AveragingIntervalAchieved,omitempty"`
	// Calibration indicates sensor calibration offset.
	Calibration *float64 `json:"Calibration,omitempty"`
	// CalibrationTime records last calibration timestamp.
	CalibrationTime string `json:"CalibrationTime,omitempty"`
	// CrestFactor indicates peak-to-RMS ratio.
	CrestFactor *float64 `json:"CrestFactor,omitempty"`
	// ElectricalContext identifies the circuit being measured.
	ElectricalContext common.ElectricalContext `json:"ElectricalContext,omitempty"`
	// Enabled indicates if the sensor is active.
	Enabled *bool `json:"Enabled,omitempty"`
	// Implementation specifies how readings are obtained.
	Implementation ImplementationType `json:"Implementation,omitempty"`
	// LifetimeReading provides total accumulated value.
	LifetimeReading *float64 `json:"LifetimeReading,omitempty"`
	// LifetimeStartDateTime records when accumulation began.
	LifetimeStartDateTime string `json:"LifetimeStartDateTime,omitempty"`
	// LoadPercent indicates utilization percentage (deprecated in v1.1.0).
	LoadPercent *float64 `json:"LoadPercent,omitempty"`
	// Location describes physical placement.
	Location *common.Location `json:"Location,omitempty"`
	// LowestReading records minimum observed value.
	LowestReading *float64 `json:"LowestReading,omitempty"`
	// LowestReadingTime records when minimum occurred.
	LowestReadingTime string `json:"LowestReadingTime,omitempty"`
	// MaxAllowableOperatingValue specifies upper safe limit.
	MaxAllowableOperatingValue *float64 `json:"MaxAllowableOperatingValue,omitempty"`
	// MinAllowableOperatingValue specifies lower safe limit.
	MinAllowableOperatingValue *float64 `json:"MinAllowableOperatingValue,omitempty"`
	// PeakReading records maximum observed value.
	PeakReading *float64 `json:"PeakReading,omitempty"`
	// PeakReadingTime records when maximum occurred.
	PeakReadingTime string `json:"PeakReadingTime,omitempty"`
	// PhaseAngleDegrees indicates current-voltage phase difference.
	PhaseAngleDegrees *float64 `json:"PhaseAngleDegrees,omitempty"`
	// PhysicalContext describes measurement location.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`
	// PhysicalSubContext describes sub-region being measured.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`
	// PowerFactor indicates power efficiency.
	PowerFactor *float64 `json:"PowerFactor,omitempty"`
	// Precision indicates significant digits.
	Precision *int `json:"Precision,omitempty"`
	// ReactivekVARh measures reactive energy in kVARh.
	ReactivekVARh *float64 `json:"ReactivekVARh,omitempty"`
	// ReactiveVAR measures reactive power in volt-amperes reactive.
	ReactiveVAR *float64 `json:"ReactiveVAR,omitempty"`
	// Reading contains current measurement value.
	Reading *float64 `json:"Reading,omitempty"`
	// ReadingAccuracy indicates measurement uncertainty.
	ReadingAccuracy *float64 `json:"ReadingAccuracy,omitempty"`
	// ReadingBasis specifies measurement reference.
	ReadingBasis ReadingBasisType `json:"ReadingBasis,omitempty"`
	// ReadingRangeMax specifies maximum valid value.
	ReadingRangeMax *float64 `json:"ReadingRangeMax,omitempty"`
	// ReadingRangeMin specifies minimum valid value.
	ReadingRangeMin *float64 `json:"ReadingRangeMin,omitempty"`
	// ReadingTime records measurement timestamp.
	ReadingTime string `json:"ReadingTime,omitempty"`
	// ReadingType specifies what is being measured.
	ReadingType ReadingType `json:"ReadingType,omitempty"`
	// ReadingUnits specifies measurement units.
	ReadingUnits string `json:"ReadingUnits,omitempty"`
	// RelatedItem links to associated resources.
	RelatedItem []string `json:"RelatedItem,omitempty"`
	// RelatedItemCount indicates number of related items.
	RelatedItemCount int `json:"RelatedItem@odata.count,omitempty"`
	// SensingFrequency specifies sampling rate (deprecated in v1.1.0).
	SensingFrequency *float64 `json:"SensingFrequency,omitempty"`
	// SensingInterval specifies sampling period.
	SensingInterval string `json:"SensingInterval,omitempty"`
	// SensorGroup identifies contributing sensors.
	SensorGroup *RedundantGroup `json:"SensorGroup,omitempty"`
	// SensorResetTime records last metrics reset.
	SensorResetTime string `json:"SensorResetTime,omitempty"`
	// SpeedRPM indicates rotational speed.
	SpeedRPM *float64 `json:"SpeedRPM,omitempty"`
	// Status provides health and state information.
	Status common.Status `json:"Status,omitempty"`
	// THDPercent indicates total harmonic distortion.
	THDPercent *float64 `json:"THDPercent,omitempty"`
	// Thresholds defines operational limits.
	Thresholds Thresholds `json:"Thresholds,omitempty"`
	// UserLabel provides customizable identification.
	UserLabel string `json:"UserLabel,omitempty"`
	// VoltageType specifies AC/DC measurement type.
	VoltageType VoltageType `json:"VoltageType,omitempty"`

	// AssociatedControls links to affecting controls.
	AssociatedControls []string `json:"AssociatedControls,omitempty"`
	// AssociatedControlsCount indicates number of controls.
	AssociatedControlsCount int `json:"AssociatedControls@odata.count,omitempty"`

	// Oem contains vendor-specific extensions.
	Oem json.RawMessage `json:"Oem,omitempty"`
	// OemActions contains vendor-specific actions.
	OemActions json.RawMessage `json:"OemActions,omitempty"`
	// OemLinks contains vendor-specific links.
	OemLinks json.RawMessage `json:"OemLinks,omitempty"`

	resetMetricsTarget    string `json:"-"`
	resetToDefaultsTarget string `json:"-"`
}

// UnmarshalJSON converts JSON data to a Sensor object.
func (sensor *Sensor) UnmarshalJSON(b []byte) error {
	type temp Sensor
	type Actions struct {
		ResetMetrics    common.ActionTarget `json:"#Sensor.ResetMetrics"`
		ResetToDefaults common.ActionTarget `json:"#Sensor.ResetToDefaults"`
		Oem             json.RawMessage     `json:"Oem"`
	}
	type Links struct {
		AssociatedControls      common.Links `json:"AssociatedControls"`
		AssociatedControlsCount int          `json:"AssociatedControls@odata.count"`
		Oem                     json.RawMessage
	}

	var t struct {
		temp
		RelatedItem      common.Links `json:"RelatedItem"`
		RelatedItemCount int          `json:"RelatedItem@odata.count"`
		Links            Links        `json:"Links"`
		Actions          Actions      `json:"Actions"`
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*sensor = Sensor(t.temp)
	sensor.RelatedItem = t.RelatedItem.ToStrings()
	sensor.RelatedItemCount = t.RelatedItemCount
	sensor.AssociatedControls = t.Links.AssociatedControls.ToStrings()
	sensor.AssociatedControlsCount = t.Links.AssociatedControlsCount
	sensor.OemLinks = t.Links.Oem
	sensor.resetMetricsTarget = t.Actions.ResetMetrics.Target
	sensor.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	sensor.OemActions = t.Actions.Oem

	return nil
}

// ResetMetrics clears accumulated sensor metrics.
func (sensor *Sensor) ResetMetrics() error {
	if sensor.resetMetricsTarget == "" {
		return fmt.Errorf("ResetMetrics is not supported")
	}
	return sensor.Post(sensor.resetMetricsTarget, nil)
}

// ResetToDefaults restores factory settings (v1.6+).
func (sensor *Sensor) ResetToDefaults() error {
	if sensor.resetToDefaultsTarget == "" {
		return fmt.Errorf("ResetToDefaults is not supported")
	}
	return sensor.Post(sensor.resetToDefaultsTarget, nil)
}

// GetSensor retrieves a specific sensor from the service.
func GetSensor(c common.Client, uri string) (*Sensor, error) {
	return common.GetObject[Sensor](c, uri)
}

// ListReferencedSensors retrieves a collection of sensors.
func ListReferencedSensors(c common.Client, link string) ([]*Sensor, error) {
	return common.GetCollectionObjects[Sensor](c, link)
}
