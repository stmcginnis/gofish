//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

// The implementation of the sensor.
type SensorImplementation string

const (
	// The reading is acquired from a physical sensor.
	PhysicalSensorImplementation SensorImplementation = "PhysicalSensor"
	// The reading is obtained from software or a device.
	ReportedImplementation SensorImplementation = "Reported"
	// The reading is obtained by applying a calculation
	// on one or more properties or multiple sensors.
	// The calculation is not provided.
	SynthesizedImplementation SensorImplementation = "Synthesized"
)

// The type of sensor.
type ReadingType string

const (
	// Absolute humidity (g/cu m).
	AbsoluteHumidityReadingType ReadingType = "AbsoluteHumidity"
	// Airflow (cu ft/min).
	AirFlowReadingType ReadingType = "AirFlow"
	// Air flow (m^3/min).
	AirFlowCMMReadingType ReadingType = "AirFlowCMM"
	// Altitude (m).
	AltitudeReadingType ReadingType = "Altitude"
	// Barometric pressure (mm).
	BarometricReadingType ReadingType = "Barometric"
	// Charge (Ah).
	ChargeAhReadingType ReadingType = "ChargeAh"
	// Current (A).
	CurrentReadingType ReadingType = "Current"
	// Energy (J).
	EnergyJoulesReadingType ReadingType = "EnergyJoules"
	// Energy (kWh).
	EnergykWhReadingType ReadingType = "EnergykWh"
	// Energy (Wh).
	EnergyWhReadingType ReadingType = "EnergyWh"
	// Frequency (Hz).
	FrequencyReadingType ReadingType = "Frequency"
	// Heat (kW).
	HeatReadingType ReadingType = "Heat"
	// Relative humidity (percent).
	HumidityReadingType ReadingType = "Humidity"
	// Deprecated: (v1.7) Liquid flow (L/s).
	LiquidFlowReadingType ReadingType = "LiquidFlow"
	// Liquid flow (L/min).
	LiquidFlowLPMReadingType ReadingType = "LiquidFlowLPM"
	// Liquid level (cm).
	LiquidLevelReadingType ReadingType = "LiquidLevel"
	// Percent (%).
	PercentReadingType ReadingType = "Percent"
	// Power (W).
	PowerReadingType ReadingType = "Power"
	// Deprecated: (v1.7) Pressure (Pa).
	PressureReadingType ReadingType = "Pressure"
	// Pressure (kPa).
	PressurekPaReadingType ReadingType = "PressurekPa"
	// Pressure (Pa).
	PressurePaReadingType ReadingType = "PressurePa"
	// Rotational (RPM).
	RotationalReadingType ReadingType = "Rotational"
	// Temperature (C).
	TemperatureReadingType ReadingType = "Temperature"
	// Voltage (VAC or VDC).
	VoltageReadingType ReadingType = "Voltage"
)

type ReadingBasisType string

const (
	// A reading that reports the difference between two measurements.
	DeltaReadingBasisType ReadingBasisType = "Delta"
	// A reading that decreases as it approaches a defined reference point.
	HeadroomReadingBasisType ReadingBasisType = "Headroom"
	// A zero-based reading.
	ZeroReadingBasisType ReadingBasisType = "Zero"
)

type ThresholdActivation string

const (
	// Value decreases below the threshold.
	DecreasingThresholdActivation ThresholdActivation = "Decreasing"
	// The threshold is disabled.
	DisabledThresholdActivation ThresholdActivation = "Disabled"
	// Value crosses the threshold in either direction.
	EitherThresholdActivation ThresholdActivation = "Either"
	// Value increases above the threshold.
	IncreasingThresholdActivation ThresholdActivation = "Increasing"
)

type Threshold struct {
	// The direction of crossing that activates this threshold.
	Activation ThresholdActivation
	// The duration the sensor value must violate the threshold before the threshold is activated.
	DwellTime string
	// The duration the sensor value must not violate the threshold before the threshold is deactivated.
	HysteresisDuration string
	// The reading offset from the threshold value required to clear the threshold.
	HysteresisReading float32
	// The threshold value.
	Reading float32
}

type Thresholds struct {
	// The value at which the reading is below normal range.
	LowerCaution Threshold
	// 	A user-defined value at which the reading is considered below normal range.
	LowerCautionUser Threshold
	// The value at which the reading is below normal range but not yet fatal.
	LowerCritical Threshold
	// A user-defined value at which the reading is considered below normal range but not yet fatal.
	LowerCriticalUser Threshold
	// The value at which the reading is below normal range and fatal.
	LowerFatal Threshold
	// The value at which the reading is above normal range.
	UpperCaution Threshold
	// A user-defined value at which the reading is considered above normal range.
	UpperCautionUser Threshold
	// The value at which the reading is above normal range but not yet fatal.
	UpperCritical Threshold
	// A user-defined value at which the reading is considered above normal range but not yet fatal.
	UpperCriticalUser Threshold
	// The value at which the reading is above normal range and fatal.
	UpperFatal Threshold
}

// Sensor represents the sensors located in the chassis and sub-components. (v1.9+)
type Sensor struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// The estimated percent error of measured versus actual values.
	Accuracy float32
	// The adjusted maximum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMaxAllowableOperatingValue float32
	// The adjusted minimum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMinAllowableOperatingValue float32
	// Apparent energy (kVAh).
	ApparentkVAh float32
	// The product of voltage and current for an AC circuit, in volt-ampere units.
	ApparentVA float32
	// The average sensor value.
	AverageReading float32
	// The interval over which the average sensor value is calculated.
	AveragingInterval string
	// Indicates that enough readings were collected to calculate the average sensor reading over the averaging interval time.
	AveragingIntervalAchieved bool
	// The calibration offset applied to the Reading.
	Calibration float32
	// The date and time that the sensor was last calibrated.
	CalibrationTime string
	// The crest factor for this sensor.
	CrestFactor float32
	// The combination of current-carrying conductors.
	ElectricalContext common.ElectricalContext
	// The implementation of the sensor.
	Implementation SensorImplementation
	// The total accumulation value for this sensor.
	LifetimeReading float32
	// Deprecated: (v1.1) The power load utilization for this sensor.
	LoadPercent float32
	// The location information for this sensor.
	Location common.Location
	// The lowest sensor value.
	LowestReading float32
	// The time when the lowest sensor value occurred.
	LowestReadingTime string
	// The maximum allowable operating value for this equipment.
	MaxAllowableOperatingValue float32
	// The minimum allowable operating value for this equipment.
	MinAllowableOperatingValue float32
	// The peak sensor value.
	PeakReading float32
	// The time when the peak sensor value occurred.
	PeakReadingTime string
	// The phase angle (degrees) between the current and voltage waveforms.
	PhaseAngleDegrees float32
	// The area or device to which this sensor measurement applies.
	PhysicalContext common.PhysicalContext
	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext common.PhysicalSubContext
	// The power factor for this sensor.
	PowerFactor float32
	// The number of significant digits in the reading.
	Precision float32
	// Reactive energy (kVARh).
	ReactivekVARh float32
	// The square root of the difference term of squared apparent VA and squared power (Reading) for a circuit, in VAR units.
	ReactiveVAR float32
	// The sensor value.
	Reading float32
	// The basis for the reading of this sensor.
	ReadingBasis ReadingBasisType
	// The maximum possible value for this sensor.
	ReadingRangeMax float32
	// The minimum possible value for this sensor.
	ReadingRangeMin float32
	// The date and time that the reading was acquired from the sensor.
	ReadingTime string
	// The type of sensor.
	ReadingType ReadingType
	// The units of the reading and thresholds.
	ReadingUnits string
	// Deprecated: (v1.1) The time interval between readings of the physical sensor.
	SensingFrequency float32
	// The time interval between readings of the sensor.
	SensingInterval string
	// The group of sensors that provide readings for this sensor.
	SensorGroup RedundantGroup
	// The date and time when the time-based properties were last reset.
	SensorResetTime string
	// The rotational speed.
	SpeedRPM float32
	// 	The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// The total harmonic distortion (THD).
	THDPercent float32
	// The set of thresholds defined for this sensor.
	Thresholds Thresholds
	// This property shall represent the type of input voltage the sensor monitors.
	VoltageType InputType

	// An array of links to resources or objects that this sensor services.
	relatedItem      []string
	RelatedItemCount int
	// An array of links to the controls that can affect this sensor.
	associatedControls      []string
	AssociatedControlsCount int
	// The manufacturer- or provider-specific data.
	Oem json.RawMessage
	// OEMLinks are all OEM data under link section
	OemLinks json.RawMessage

	resetMetricsTarget    string
	resetToDefaultsTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (sensor *Sensor) UnmarshalJSON(b []byte) error {
	type temp Sensor
	type linkReference struct {
		AssociatedControls      common.Links
		AssociatedControlsCount int `json:"AssociatedControls@odata.count"`
		Oem                     json.RawMessage
	}
	type actions struct {
		ResetMetrics struct {
			Target string
		} `json:"#Sensor.ResetMetrics"`
		ResetToDefaults struct {
			Target string
		} `json:"#Sensor.ResetToDefaults"`
		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		RelatedItem      common.Links
		RelatedItemCount int `json:"RelatedItem@odata.count"`
		Links            linkReference
		Actions          actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*sensor = Sensor(t.temp)
	sensor.relatedItem = t.RelatedItem.ToStrings()
	sensor.RelatedItemCount = t.RelatedItemCount
	sensor.associatedControls = t.Links.AssociatedControls.ToStrings()
	sensor.AssociatedControlsCount = t.Links.AssociatedControlsCount
	sensor.OemLinks = t.Links.Oem
	sensor.resetMetricsTarget = t.Actions.ResetMetrics.Target
	sensor.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	sensor.OemActions = t.Actions.Oem

	return nil
}

// GetSensor will get a Sensor instance from the Redfish service.
func GetSensor(c common.Client, uri string) (*Sensor, error) {
	var sensor Sensor
	return &sensor, sensor.Get(c, uri, &sensor)
}

// ListReferencedSensor gets the Sensor collection.
func ListReferencedSensors(c common.Client, link string) ([]*Sensor, error) { //nolint:dupl
	var result []*Sensor
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Sensor
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		sensor, err := GetSensor(c, link)
		ch <- GetResult{Item: sensor, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

func (sensor *Sensor) ResetMetrics() error {
	if sensor.resetMetricsTarget == "" {
		return fmt.Errorf("ResetMetrics is not supported") //nolint:golint
	}

	return sensor.Post(sensor.resetMetricsTarget, nil)
}

// Available for redfish v1.6+
func (sensor *Sensor) ResetToDefaults() error {
	if sensor.resetToDefaultsTarget == "" {
		return fmt.Errorf("ResetToDefaults is not supported") //nolint:golint
	}

	return sensor.Post(sensor.resetToDefaultsTarget, nil)
}
