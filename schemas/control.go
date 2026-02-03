//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #Control.v1_7_0.Control

package schemas

import (
	"encoding/json"
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
	// TemperatureControlType shall indicate a control used to regulate
	// temperature, in degree Celsius units, either to a single set point or within
	// a range. The 'SetPointUnits' property shall contain 'Cel'.
	TemperatureControlType ControlType = "Temperature"
	// PowerControlType shall indicate a control used to regulate or limit maximum
	// power consumption, in watt units, either to a single set point or within a
	// range. The 'SetPointUnits' property shall contain 'W'.
	PowerControlType ControlType = "Power"
	// FrequencyControlType shall indicate a control used to limit the operating
	// frequency, in hertz units, of a device, either to a single set point or
	// within a range. The 'SetPointUnits' property shall contain 'Hz'.
	FrequencyControlType ControlType = "Frequency"
	// FrequencyMHzControlType shall indicate a control used to limit the operating
	// frequency, in megahertz units, of a device, either to a single set point or
	// within a range. The 'SetPointUnits' property shall contain 'MHz'.
	FrequencyMHzControlType ControlType = "FrequencyMHz"
	// PressureControlType shall indicate a control used to adjust pressure in a
	// system, in kilopascal units. The 'SetPointUnits' property shall contain
	// 'kPa'.
	PressureControlType ControlType = "Pressure"
	// PressurekPaControlType shall indicate a control used to adjust pressure in a
	// system, in kilopascal units. The 'SetPointUnits' property shall contain
	// 'kPa'.
	PressurekPaControlType ControlType = "PressurekPa"
	// ValveControlType shall indicate a control used to adjust a valve in a
	// system, in percent units. The 'SetPointUnits' property shall contain '%'. A
	// value of '100' shall indicate the valve is completely open, and a value of
	// '0' shall indicate the valve is completely closed.
	ValveControlType ControlType = "Valve"
	// PercentControlType shall indicate a percent-based control, in percent units.
	// The 'SetPointUnits' property shall contain '%'.
	PercentControlType ControlType = "Percent"
	// DutyCycleControlType shall indicate a control used to adjust the duty cycle,
	// such as a PWM-based control, in percent units. The 'SetPointUnits' property
	// shall contain '%'.
	DutyCycleControlType ControlType = "DutyCycle"
	// LinearPositionControlType shall indicate a control used to adjust linear
	// position or distance, in meter units. The 'SetPointUnits' property shall
	// contain 'm'.
	LinearPositionControlType ControlType = "LinearPosition"
	// LinearVelocityControlType shall indicate a control used to adjust linear
	// velocity, in meters per second units. The 'SetPointUnits' property shall
	// contain 'm/s'.
	LinearVelocityControlType ControlType = "LinearVelocity"
	// LinearAccelerationControlType shall indicate a control used to adjust linear
	// acceleration, in meters per square second units. The 'SetPointUnits'
	// property shall contain 'm/s2'.
	LinearAccelerationControlType ControlType = "LinearAcceleration"
	// RotationalPositionControlType shall indicate a control used to adjust
	// rotational position, in radian units. The 'SetPointUnits' property shall
	// contain 'rad'.
	RotationalPositionControlType ControlType = "RotationalPosition"
	// RotationalVelocityControlType shall indicate a control used to adjust
	// rotational velocity, in radians per second units. The 'SetPointUnits'
	// property shall contain 'rad/s'.
	RotationalVelocityControlType ControlType = "RotationalVelocity"
	// RotationalAccelerationControlType shall indicate a control used to adjust
	// rotational acceleration, in radians per square second units. The
	// 'SetPointUnits' property shall contain 'rad/s2'.
	RotationalAccelerationControlType ControlType = "RotationalAcceleration"
	// LiquidFlowLPMControlType shall indicate a control used to adjust the volume
	// of liquid per unit of time, in liters per minute units, that flows through a
	// particular junction. The 'SetPointUnits' property shall contain 'L/min'.
	LiquidFlowLPMControlType ControlType = "LiquidFlowLPM"
)

type ControlImplementationType string

const (
	// ProgrammableControlImplementationType The set point can be adjusted through this
	// interface.
	ProgrammableControlImplementationType ControlImplementationType = "Programmable"
	// DirectControlImplementationType The set point directly affects the control value.
	DirectControlImplementationType ControlImplementationType = "Direct"
	// MonitoredControlImplementationType is a physical control that cannot be adjusted
	// through this interface.
	MonitoredControlImplementationType ControlImplementationType = "Monitored"
)

type SetPointType string

const (
	// SingleSetPointType shall indicate the control utilizes a single set point
	// for its operation. The 'SetPoint' property shall be present for this control
	// type. The 'SettingMin' and 'SettingMax' properties shall not be present for
	// this control type.
	SingleSetPointType SetPointType = "Single"
	// RangeSetPointType shall indicate the control utilizes a set point range for
	// its operation. The 'SettingMin' and 'SettingMax' properties shall be present
	// for this control type. The 'SetPoint' property shall not be present for this
	// control type.
	RangeSetPointType SetPointType = "Range"
	// MonitorSetPointType shall indicate the control provides only monitoring of a
	// sensor reading, and does not provide the ability to affect the reading. This
	// value allows for multiple controls to be populated while only some of those
	// controls provide a set point. The 'SetPoint,'SettingMin' and 'SettingMax'
	// properties shall not be present for this control type. For example, a motion
	// controller may provide both position-based and velocity-based control modes,
	// where the selected mode provides the 'SetPoint', and the control for the
	// non-selected mode provides only a sensor reading.
	MonitorSetPointType SetPointType = "Monitor"
)

// Control shall represent a control point for a Redfish implementation.
type Control struct {
	Entity
	// Accuracy shall contain the percent error of the measured versus actual
	// values of the 'SetPoint' property.
	//
	// Deprecated: v1.4.0
	// This property has been deprecated in favor of 'SetPointAccuracy' to provide
	// a range instead of a percentage.
	Accuracy *float64 `json:",omitempty"`
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// AllowableNumericValues shall contain the supported values for this control.
	// The units shall follow the value of 'SetPointUnits'. This property should
	// only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the 'Increment' property.
	AllowableNumericValues []*float64
	// AssociatedSensors shall contain an array of links to resources of type
	// 'Sensor' that represent the sensors related to this control.
	associatedSensors []string
	// AssociatedSensorsCount
	AssociatedSensorsCount int `json:"AssociatedSensors@odata.count"`
	// ControlDelaySeconds shall contain the time in seconds that will elapse after
	// the control value deviates above or below the value of 'SetPoint' before the
	// control will activate.
	ControlDelaySeconds *float64 `json:",omitempty"`
	// ControlLoop shall contain the details for the control loop described by this
	// resource.
	ControlLoop ControlLoop
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// ControlType shall contain the type of the control.
	ControlType ControlType
	// DeadBand shall contain the maximum deviation value allowed above or below
	// the value of 'SetPoint' before the control will activate.
	DeadBand *float64 `json:",omitempty"`
	// DefaultSetPoint shall contain the default set point control value. The units
	// shall follow the value of 'SetPointUnits'. Services apply this value to the
	// 'SetPoint' property under certain conditions, such as a reset of the manager
	// or a 'ResetToDefaults' action.
	//
	// Version added: v1.3.0
	DefaultSetPoint *float64 `json:",omitempty"`
	// Implementation shall contain the implementation of the control.
	Implementation ControlImplementationType
	// Increment shall contain the smallest change allowed to the value of the
	// 'SetPoint', 'SettingMin', or 'SettingMax' properties. The units shall follow
	// the value of 'SetPointUnits'.
	Increment *float64 `json:",omitempty"`
	// Location shall indicate the location information for this control.
	Location Location
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected component or
	// region within the equipment to which this control applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region
	// within the equipment to which this control applies. This property generally
	// differentiates multiple controls within the same 'PhysicalContext' instance.
	PhysicalSubContext PhysicalSubContext
	// RelatedItem shall contain an array of links to resources that this control
	// services.
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Sensor shall contain the 'Sensor' excerpt directly associated with this
	// control. The value of the 'DataSourceUri' property shall reference a
	// resource of type 'Sensor'. This property shall not be present if multiple
	// sensors are associated with a single control.
	Sensor SensorExcerpt
	// SetPoint shall contain the desired set point control value. The units shall
	// follow the value of 'SetPointUnits'. If the 'DefaultSetPoint' property is
	// not supported and if a user-defined set point is not configured, the
	// property may contain 'null' in responses.
	SetPoint *float64 `json:",omitempty"`
	// SetPointAccuracy shall contain the accuracy of the value of the 'SetPoint'
	// for this control. The value shall be the absolute value of the maximum
	// deviation of the 'SetPoint' from its actual value. The value shall be in
	// units that follow the 'SetPointUnits' for this control.
	//
	// Version added: v1.4.0
	SetPointAccuracy *float64 `json:",omitempty"`
	// SetPointError shall contain the error, or difference, of the related Sensor
	// 'Reading' value from the value of the 'SetPoint'. The units shall follow the
	// value of 'SetPointUnits'.
	//
	// Version added: v1.7.0
	SetPointError *float64 `json:",omitempty"`
	// SetPointType shall contain the type of set point definitions used to
	// describe this control.
	SetPointType SetPointType
	// SetPointUnits shall contain the units of the control's set point and related
	// properties. The value shall follow the case-sensitive symbol format defined
	// by the Unified Code for Units of Measure (UCUM), as specified by the 'Units
	// of measure annotation' clause of the Redfish Specification.
	SetPointUnits string
	// SetPointUpdateTime shall contain the date and time that the value of
	// 'SetPoint' was last changed.
	SetPointUpdateTime string
	// SettingMax shall contain the maximum desired set point within the acceptable
	// range. The service shall reject values greater than the value of
	// 'AllowableMax'. The units shall follow the value of 'SetPointUnits'.
	SettingMax *float64 `json:",omitempty"`
	// SettingMin shall contain the minimum desired set point within the acceptable
	// range. The service shall reject values less than the value of
	// 'AllowableMin'. The units shall follow the value of 'SetPointUnits'.
	SettingMin *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Control object from the raw JSON.
func (c *Control) UnmarshalJSON(b []byte) error {
	type temp Control
	type cActions struct {
		ResetToDefaults ActionTarget `json:"#Control.ResetToDefaults"`
	}
	var tmp struct {
		temp
		Actions           cActions
		AssociatedSensors Links `json:"AssociatedSensors"`
		RelatedItem       Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Control(tmp.temp)

	// Extract the links to other entities for later
	c.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	c.associatedSensors = tmp.AssociatedSensors.ToStrings()
	c.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Control) Update() error {
	readWriteFields := []string{
		"ControlDelaySeconds",
		"ControlMode",
		"DeadBand",
		"SetPoint",
		"SettingMax",
		"SettingMin",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetControl will get a Control instance from the service.
func GetControl(c Client, uri string) (*Control, error) {
	return GetObject[Control](c, uri)
}

// ListReferencedControls gets the collection of Control from
// a provided reference.
func ListReferencedControls(c Client, link string) ([]*Control, error) {
	return GetCollectionObjects[Control](c, link)
}

// This action shall reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Control) ResetToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetToDefaultsTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AssociatedSensors gets the AssociatedSensors linked resources.
func (c *Control) AssociatedSensors() ([]*Sensor, error) {
	return GetObjects[Sensor](c.client, c.associatedSensors)
}

// RelatedItem gets the RelatedItem linked resources.
func (c *Control) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.relatedItem)
}

// ControlExcerpt shall represent a control point for a Redfish implementation.
type ControlExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy. If no source resource is
	// implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the value of the 'Reading' property of the 'Sensor'
	// resource directly associated with this control. This property shall not be
	// present if multiple sensors are associated with a single control.
	Reading *float64 `json:",omitempty"`
}

// ControlLoop shall describe the details of a control loop.
type ControlLoop struct {
	// CoefficientUpdateTime shall contain the date and time that any of the
	// coefficients for the control loop were last changed.
	CoefficientUpdateTime string
	// Differential shall contain the coefficient for the differential factor in a
	// control loop.
	Differential *float64 `json:",omitempty"`
	// Integral shall contain the coefficient for the integral factor in a control
	// loop.
	Integral *float64 `json:",omitempty"`
	// Proportional shall contain the coefficient for the proportional factor in a
	// control loop.
	Proportional *float64 `json:",omitempty"`
}

// ControlNodeExcerpt shall represent a control point for a Redfish
// implementation.
type ControlNodeExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy. If no source resource is
	// implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the value of the 'Reading' property of the 'Sensor'
	// resource directly associated with this control. This property shall not be
	// present if multiple sensors are associated with a single control.
	Reading *float64 `json:",omitempty"`
	// ReadingUnits shall contain the units of the sensor's reading and thresholds.
	// This property shall not be present if multiple sensors are associated with a
	// single control.
	ReadingUnits string
	// SetPoint shall contain the desired set point control value. The units shall
	// follow the value of 'SetPointUnits'. If the 'DefaultSetPoint' property is
	// not supported and if a user-defined set point is not configured, the
	// property may contain 'null' in responses.
	SetPoint *float64 `json:",omitempty"`
	// SetPointUnits shall contain the units of the control's set point and related
	// properties. The value shall follow the case-sensitive symbol format defined
	// by the Unified Code for Units of Measure (UCUM), as specified by the 'Units
	// of measure annotation' clause of the Redfish Specification.
	SetPointUnits string
}

// ControlRangeExcerpt shall represent a control point for a Redfish
// implementation.
type ControlRangeExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// AllowableNumericValues shall contain the supported values for this control.
	// The units shall follow the value of 'SetPointUnits'. This property should
	// only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the 'Increment' property.
	AllowableNumericValues []*float64
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy. If no source resource is
	// implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the value of the 'Reading' property of the 'Sensor'
	// resource directly associated with this control. This property shall not be
	// present if multiple sensors are associated with a single control.
	Reading *float64 `json:",omitempty"`
	// SettingMax shall contain the maximum desired set point within the acceptable
	// range. The service shall reject values greater than the value of
	// 'AllowableMax'. The units shall follow the value of 'SetPointUnits'.
	SettingMax *float64 `json:",omitempty"`
	// SettingMin shall contain the minimum desired set point within the acceptable
	// range. The service shall reject values less than the value of
	// 'AllowableMin'. The units shall follow the value of 'SetPointUnits'.
	SettingMin *float64 `json:",omitempty"`
}

// ControlSingleExcerpt shall represent a control point for a Redfish
// implementation.
type ControlSingleExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy. If no source resource is
	// implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string `json:"DataSourceUri"`
	// DefaultSetPoint shall contain the default set point control value. The units
	// shall follow the value of 'SetPointUnits'. Services apply this value to the
	// 'SetPoint' property under certain conditions, such as a reset of the manager
	// or a 'ResetToDefaults' action.
	//
	// Version added: v1.3.0
	DefaultSetPoint *float64 `json:",omitempty"`
	// Reading shall contain the value of the 'Reading' property of the 'Sensor'
	// resource directly associated with this control. This property shall not be
	// present if multiple sensors are associated with a single control.
	Reading *float64 `json:",omitempty"`
	// SetPoint shall contain the desired set point control value. The units shall
	// follow the value of 'SetPointUnits'. If the 'DefaultSetPoint' property is
	// not supported and if a user-defined set point is not configured, the
	// property may contain 'null' in responses.
	SetPoint *float64 `json:",omitempty"`
}

// ControlSingleLoopExcerpt shall represent a control point for a Redfish
// implementation.
type ControlSingleLoopExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the 'SetPoint' or
	// 'SettingMax' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMax' above this value.
	AllowableMax *float64 `json:",omitempty"`
	// AllowableMin shall indicate the minimum possible value of the 'SetPoint' or
	// 'SettingMin' properties for this control. Services shall not accept values
	// for 'SetPoint' or 'SettingMin' below this value.
	AllowableMin *float64 `json:",omitempty"`
	// ControlLoop shall contain the details for the control loop described by this
	// resource.
	ControlLoop ControlLoop
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceURI shall contain a URI to the resource that provides the source
	// of the excerpt contained within this copy. If no source resource is
	// implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceURI string `json:"DataSourceUri"`
	// Reading shall contain the value of the 'Reading' property of the 'Sensor'
	// resource directly associated with this control. This property shall not be
	// present if multiple sensors are associated with a single control.
	Reading *float64 `json:",omitempty"`
	// SetPoint shall contain the desired set point control value. The units shall
	// follow the value of 'SetPointUnits'. If the 'DefaultSetPoint' property is
	// not supported and if a user-defined set point is not configured, the
	// property may contain 'null' in responses.
	SetPoint *float64 `json:",omitempty"`
}
