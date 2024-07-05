//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ControlMode string

const (
	// AutomaticControlMode Automatically adjust control to meet the set point.
	AutomaticControlMode ControlMode = "Automatic"
	// OverrideControlMode User override of the automatic set point value.
	OverrideControlMode ControlMode = "Override"
	// ManualControlMode No automatic adjustments are made to the control.
	ManualControlMode ControlMode = "Manual"
	// DisabledControlMode The control has been disabled.
	DisabledControlMode ControlMode = "Disabled"
)

type ControlType string

const (
	// TemperatureControlType shall indicate a control used to regulate temperature, in degree Celsius units, either to
	// a single set point or within a range. The SetPointUnits property shall contain 'Cel'.
	TemperatureControlType ControlType = "Temperature"
	// PowerControlType shall indicate a control used to regulate or limit maximum power consumption, in watt units,
	// either to a single set point or within a range. The SetPointUnits property shall contain 'W'.
	PowerControlType ControlType = "Power"
	// FrequencyControlType shall indicate a control used to limit the operating frequency, in hertz units, of a
	// device, either to a single set point or within a range. The SetPointUnits property shall contain 'Hz'.
	FrequencyControlType ControlType = "Frequency"
	// FrequencyMHzControlType shall indicate a control used to limit the operating frequency, in megahertz units, of a
	// device, either to a single set point or within a range. The SetPointUnits property shall contain 'MHz'.
	FrequencyMHzControlType ControlType = "FrequencyMHz"
	// PressureControlType shall indicate a control used to adjust pressure in a system, in kilopascal units. The
	// SetPointUnits property shall contain 'kPa'.
	PressureControlType ControlType = "Pressure"
	// PressurekPaControlType shall indicate a control used to adjust pressure in a system, in kilopascal units. The
	// SetPointUnits property shall contain 'kPa'.
	PressurekPaControlType ControlType = "PressurekPa"
	// ValveControlType shall indicate a control used to adjust a valve in a system, in percent units. The
	// SetPointUnits property shall contain '%'. A value of '100' shall indicate the valve is completely open, and a
	// value of '0' shall indicate the valve is completely closed.
	ValveControlType ControlType = "Valve"
	// PercentControlType shall indicate a percent-based control, in percent units. The SetPointUnits property shall
	// contain '%'.
	PercentControlType ControlType = "Percent"
	// DutyCycleControlType shall indicate a control used to adjust the duty cycle, such as a PWM-based control, in
	// percent units. The SetPointUnits property shall contain '%'.
	DutyCycleControlType ControlType = "DutyCycle"
)

type ControlImplementationType string

const (
	// ProgrammableImplementationType The set point can be adjusted through this interface.
	ProgrammableControlImplementationType ControlImplementationType = "Programmable"
	// DirectImplementationType The set point directly affects the control value.
	DirectControlImplementationType ControlImplementationType = "Direct"
	// MonitoredImplementationType A physical control that cannot be adjusted through this interface.
	MonitoredControlImplementationType ControlImplementationType = "Monitored"
)

type SetPointType string

const (
	// SingleSetPointType shall indicate the control utilizes a single set point for its operation. The SetPoint
	// property shall be present for this control type. The SettingMin and SettingMax properties shall not be present
	// for this control type.
	SingleSetPointType SetPointType = "Single"
	// RangeSetPointType shall indicate the control utilizes a set point range for its operation. The SettingMin and
	// SettingMax properties shall be present for this control type. The SetPoint property shall not be present for
	// this control type.
	RangeSetPointType SetPointType = "Range"
)

// Control shall represent a control point for a Redfish implementation.
type Control struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// AllowableNumericValues shall contain the supported values for this control. The units shall follow the value of
	// SetPointUnits. This property should only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the Increment property.
	AllowableNumericValues []float64
	// AssociatedSensors shall contain an array of links to resources of type Sensor that represent the sensors related
	// to this control.
	associatedSensors []string
	// AssociatedSensorsCount is the number of sensors related to this control.
	AssociatedSensorsCount int `json:"AssociatedSensors@odata.count"`
	// ControlDelaySeconds shall contain the time in seconds that will elapse after the control value deviates above or
	// below the value of SetPoint before the control will activate.
	ControlDelaySeconds float64
	// ControlLoop shall contain the details for the control loop described by this resource.
	ControlLoop ControlLoop
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// ControlType shall contain the type of the control.
	ControlType ControlType
	// DeadBand shall contain the maximum deviation value allowed above or below the value of SetPoint before the
	// control will activate.
	DeadBand float64
	// DefaultSetPoint shall contain the default set point control value. The units shall follow the value of
	// SetPointUnits. Services apply this value to the SetPoint property under certain conditions, such as a reset of
	// the manager or a ResetToDefaults action.
	DefaultSetPoint float64
	// Description provides a description of this resource.
	Description string
	// Implementation shall contain the implementation of the control.
	Implementation ControlImplementationType
	// Increment shall contain the smallest change allowed to the value of the SetPoint, SettingMin, or SettingMax
	// properties. The units shall follow the value of SetPointUnits.
	Increment float64
	// Location shall indicate the location information for this control.
	Location common.Location
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this control applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// control applies. This property generally differentiates multiple controls within the same PhysicalContext
	// instance.
	PhysicalSubContext PhysicalSubContext
	// RelatedItem shall contain an array of links to resources that this control services.
	relatedItem []string
	// RelatedItemCount is the number of resources that this control services.
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Sensor shall contain the Sensor excerpt directly associated with this control. The value of the DataSourceUri
	// property shall reference a resource of type Sensor. This property shall not be present if multiple sensors are
	// associated with a single control.
	Sensor SensorExcerpt
	// SetPoint shall contain the desired set point control value. The units shall follow the value of SetPointUnits.
	// If the DefaultSetPoint property is not supported and if a user-defined set point is not configured, the property
	// may contain 'null' in responses.
	SetPoint float64
	// SetPointAccuracy shall contain the accuracy of the value of the SetPoint for this control. The value shall be
	// the absolute value of the maximum deviation of the SetPoint from its actual value. The value shall be in units
	// that follow the SetPointUnits for this control.
	SetPointAccuracy float64
	// SetPointType shall contain the type of set point definitions used to describe this control.
	SetPointType SetPointType
	// SetPointUnits shall contain the units of the control's set point.
	SetPointUnits string
	// SetPointUpdateTime shall contain the date and time that the value of SetPoint was last changed.
	SetPointUpdateTime string
	// SettingMax shall contain the maximum desired set point within the acceptable range. The service shall reject
	// values greater than the value of AllowableMax. The units shall follow the value of SetPointUnits.
	SettingMax float64
	// SettingMin shall contain the minimum desired set point within the acceptable range. The service shall reject
	// values less than the value of AllowableMin. The units shall follow the value of SetPointUnits.
	SettingMin float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	resetToDefaultsTarget string
}

// UnmarshalJSON unmarshals a Control object from the raw JSON.
func (control *Control) UnmarshalJSON(b []byte) error {
	type temp Control
	type Actions struct {
		ResetToDefaults common.ActionTarget `json:"#Control.ResetToDefaults"`
	}
	var t struct {
		temp
		Actions           Actions
		AssociatedSensors common.Links
		RelatedItem       common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*control = Control(t.temp)

	// Extract the links to other entities for later
	control.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	control.associatedSensors = t.AssociatedSensors.ToStrings()
	control.relatedItem = t.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	control.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (control *Control) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Control)
	original.UnmarshalJSON(control.rawData)

	readWriteFields := []string{
		"ControlDelaySeconds",
		"ControlMode",
		"DeadBand",
		"SetPoint",
		"SettingMax",
		"SettingMin",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(control).Elem()

	return control.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetControl will get a Control instance from the service.
func GetControl(c common.Client, uri string) (*Control, error) {
	return common.GetObject[Control](c, uri)
}

// ListReferencedControls gets the collection of Control from
// a provided reference.
func ListReferencedControls(c common.Client, link string) ([]*Control, error) {
	return common.GetCollectionObjects[Control](c, link)
}

// ResetToDefault resets the values of writable properties to factory defaults.
func (control *Control) ResetToDefault() error {
	if control.resetToDefaultsTarget == "" {
		return fmt.Errorf("ResetToDefault is not supported by this system")
	}

	return control.Post(control.resetToDefaultsTarget, nil)
}

// ControlExcerpt shall represent a control point for a Redfish implementation.
type ControlExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
}

// ControlLoop shall describe the details of a control loop.
type ControlLoop struct {
	// CoefficientUpdateTime shall contain the date and time that any of the coefficients for the control loop were
	// last changed.
	CoefficientUpdateTime string
	// Differential shall contain the coefficient for the differential factor in a control loop.
	Differential float64
	// Integral shall contain the coefficient for the integral factor in a control loop.
	Integral float64
	// Proportional shall contain the coefficient for the proportional factor in a control loop.
	Proportional float64
}

// ControlRangeExcerpt shall represent a control point for a Redfish implementation.
type ControlRangeExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// AllowableNumericValues shall contain the supported values for this control. The units shall follow the value of
	// SetPointUnits. This property should only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the Increment property.
	AllowableNumericValues []string
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
	// SettingMax shall contain the maximum desired set point within the acceptable range. The service shall reject
	// values greater than the value of AllowableMax. The units shall follow the value of SetPointUnits.
	SettingMax float64
	// SettingMin shall contain the minimum desired set point within the acceptable range. The service shall reject
	// values less than the value of AllowableMin. The units shall follow the value of SetPointUnits.
	SettingMin float64
}

// ControlSingleExcerpt shall represent a control point for a Redfish implementation.
type ControlSingleExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string
	// DefaultSetPoint shall contain the default set point control value. The units shall follow the value of
	// SetPointUnits. Services apply this value to the SetPoint property under certain conditions, such as a reset of
	// the manager or a ResetToDefaults action.
	DefaultSetPoint float64
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
	// SetPoint shall contain the desired set point control value. The units shall follow the value of SetPointUnits.
	// If the DefaultSetPoint property is not supported and if a user-defined set point is not configured, the property
	// may contain 'null' in responses.
	SetPoint float64
}
