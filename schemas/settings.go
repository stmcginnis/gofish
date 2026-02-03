//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.1 - #Settings.v1_4_0

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SettingsApplyTime string

const (
	// ImmediateSettingsApplyTime shall indicate the values within the settings resource
	// are applied immediately. This value may result in an immediate host reset,
	// manager reset, or other side effects.
	ImmediateSettingsApplyTime SettingsApplyTime = "Immediate"
	// OnResetSettingsApplyTime shall indicate the values within settings resource are
	// applied when the system or service is reset.
	OnResetSettingsApplyTime SettingsApplyTime = "OnReset"
	// AtMaintenanceWindowStartSettingsApplyTime shall indicate the values within the
	// settings resource are applied during the maintenance window specified by the
	// 'MaintenanceWindowStartTime' and 'MaintenanceWindowDurationInSeconds'
	// properties. A service can perform resets during this maintenance window.
	AtMaintenanceWindowStartSettingsApplyTime SettingsApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetSettingsApplyTime shall indicate the values within the
	// settings resource are applied during the maintenance window specified by the
	// 'MaintenanceWindowStartTime' and 'MaintenanceWindowDurationInSeconds'
	// properties, and if a reset occurs within the maintenance window.
	InMaintenanceWindowOnResetSettingsApplyTime SettingsApplyTime = "InMaintenanceWindowOnReset"
)

type OperationApplyTime string

const (
	// ImmediateOperationApplyTime shall indicate the requested create, delete, or
	// action operation is applied immediately. This value may result in an
	// immediate host reset, manager reset, or other side effects.
	ImmediateOperationApplyTime OperationApplyTime = "Immediate"
	// OnResetOperationApplyTime shall indicate the requested create, delete, or
	// action operation is applied when the system or service is reset.
	OnResetOperationApplyTime OperationApplyTime = "OnReset"
	// AtMaintenanceWindowStartOperationApplyTime shall indicate the requested
	// create, delete, or action operation is applied during the maintenance window
	// that the 'MaintenanceWindowStartTime' and
	// 'MaintenanceWindowDurationInSeconds' properties specify. A service can
	// complete resets during this maintenance window.
	AtMaintenanceWindowStartOperationApplyTime OperationApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetOperationApplyTime shall indicate the requested
	// create, delete, or action operation is applied during the maintenance window
	// that the 'MaintenanceWindowStartTime' and
	// 'MaintenanceWindowDurationInSeconds' properties specify, and if a reset
	// occurs within the maintenance window.
	InMaintenanceWindowOnResetOperationApplyTime OperationApplyTime = "InMaintenanceWindowOnReset"
	// OnStartUpdateRequestOperationApplyTime shall indicate the requested create,
	// delete, or action operation is applied when the 'StartUpdate' action of the
	// update service is invoked.
	OnStartUpdateRequestOperationApplyTime OperationApplyTime = "OnStartUpdateRequest"
	// OnTargetResetOperationApplyTime shall indicate the requested create, delete,
	// or action operation is applied when the target is reset.
	OnTargetResetOperationApplyTime OperationApplyTime = "OnTargetReset"
)

// Settings shall describe any settings of a resource.
type Settings struct {
	// ETag shall contain the entity tag (ETag) of the resource to which the
	// settings were applied, after the application. The client can check this
	// value against the ETag of this resource to determine whether the resource
	// had other changes.
	ETag string
	// MaintenanceWindowResource shall contain a link to a resource that contains
	// the '@Redfish.MaintenanceWindow' property that governs this resource. This
	// property should be supported if the 'SupportedApplyTimes' property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.2.0
	maintenanceWindowResource string
	// Messages shall contain an array of messages associated with the settings.
	Messages []Message
	// SettingsObject shall contain the URI of the resource that the client can
	// 'PUT' or 'PATCH' to modify the resource.
	SettingsObject string
	// SupportedApplyTimes shall contain the supported apply time values a client
	// is allowed to request when configuring the settings apply time. Services
	// that do not support clients configuring the apply time can support this
	// property with a single array member in order to inform the client when the
	// settings will be applied.
	//
	// Version added: v1.1.0
	SupportedApplyTimes []SettingsApplyTime
	// Time shall indicate the time when the settings were applied to the resource.
	Time string
}

// UnmarshalJSON unmarshals a Settings object from the raw JSON.
func (s *Settings) UnmarshalJSON(b []byte) error {
	type temp Settings
	var tmp struct {
		temp
		MaintenanceWindowResource Link `json:"MaintenanceWindowResource"`
		SettingsObject            Link `json:"SettingsObject"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Settings(tmp.temp)

	// Extract the links to other entities for later
	s.maintenanceWindowResource = tmp.MaintenanceWindowResource.String()
	s.SettingsObject = tmp.SettingsObject.String()

	return nil
}

// MaintenanceWindowResource gets the MaintenanceWindowResource linked resource.
func (s *Settings) MaintenanceWindowResource(client Client) (*Entity, error) {
	if s.maintenanceWindowResource == "" {
		return nil, nil
	}
	return GetObject[Entity](client, s.maintenanceWindowResource)
}

// MaintenanceWindow shall indicate that a resource has a maintenance window
// assignment for applying settings or operations. Other resources can link to
// this object to convey a common control surface for the configuration of the
// maintenance window.
type MaintenanceWindow struct {
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance
	// window as the number of seconds after the time specified by the
	// 'MaintenanceWindowStartTime' property.
	//
	// Version added: v1.2.0
	MaintenanceWindowDurationInSeconds uint
	// MaintenanceWindowStartTime shall indicate the date and time when the service
	// can start to apply the requested settings or operation as part of a
	// maintenance window. Services shall provide a default value if not configured
	// by a user.
	//
	// Version added: v1.2.0
	MaintenanceWindowStartTime string
}

// OperationApplyTimeSupport shall indicate that a client can request a specific
// apply time of a create, delete, or action operation of a resource.
type OperationApplyTimeSupport struct {
	// MaintenanceWindowDurationInSeconds shall contain the same as the
	// 'MaintenanceWindowDurationInSeconds' property found in the
	// 'MaintenanceWindow' structure on the 'MaintenanceWindowResource'. This
	// property shall be required if the 'SupportedValues' property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.2.0
	MaintenanceWindowDurationInSeconds uint
	// MaintenanceWindowResource shall contain a link to a resource that contains
	// the '@Redfish.MaintenanceWindow' property that governs this resource. This
	// property shall be required if the 'SupportedValues' property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.2.0
	maintenanceWindowResource string
	// MaintenanceWindowStartTime shall contain the same as the
	// 'MaintenanceWindowStartTime' property found in the 'MaintenanceWindow'
	// structure on the 'MaintenanceWindowResource'. Services shall provide a
	// default value if not configured by a user. This property shall be required
	// if the 'SupportedValues' property contains 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.2.0
	MaintenanceWindowStartTime string
	// SupportedValues shall indicate the types of apply times the client can
	// request when performing a create, delete, or action operation.
	//
	// Version added: v1.2.0
	SupportedValues []OperationApplyTime
}

// UnmarshalJSON unmarshals a OperationApplyTimeSupport object from the raw JSON.
func (o *OperationApplyTimeSupport) UnmarshalJSON(b []byte) error {
	type temp OperationApplyTimeSupport
	var tmp struct {
		temp
		MaintenanceWindowResource Link `json:"MaintenanceWindowResource"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = OperationApplyTimeSupport(tmp.temp)

	// Extract the links to other entities for later
	o.maintenanceWindowResource = tmp.MaintenanceWindowResource.String()

	return nil
}

// MaintenanceWindowResource gets the MaintenanceWindowResource linked resource.
func (o *OperationApplyTimeSupport) MaintenanceWindowResource(client Client) (*Entity, error) {
	if o.maintenanceWindowResource == "" {
		return nil, nil
	}
	return GetObject[Entity](client, o.maintenanceWindowResource)
}

// PreferredApplyTime shall be specified by client to indicate the preferred
// time to apply the configuration settings.
type PreferredApplyTime struct {
	// ApplyTime shall indicate when to apply the values in this settings resource.
	//
	// Version added: v1.1.0
	ApplyTime SettingsApplyTime
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance
	// window as the number of seconds after the time specified by the
	// 'MaintenanceWindowStartTime' property. This property shall be required if
	// the 'ApplyTime' property is 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.1.0
	MaintenanceWindowDurationInSeconds uint
	// MaintenanceWindowStartTime shall indicate the date and time when the service
	// can start to apply the future configuration as part of a maintenance window.
	// Services shall provide a default value if not configured by a user. This
	// property shall be required if the 'ApplyTime' property is
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.1.0
	MaintenanceWindowStartTime string
}

// SettingsAttributes handles the settings attribute values that may be any of several
// types and adds some basic helper methods to make accessing values easier.
type SettingsAttributes map[string]any

// String gets the string representation of the attribute value.
func (sa SettingsAttributes) String(name string) string {
	if val, ok := sa[name]; ok {
		return fmt.Sprintf("%v", val)
	}

	return ""
}

// Float64 gets the value as a float64 or 0 if that is not possible.
func (sa SettingsAttributes) Float64(name string) float64 {
	if val, ok := sa[name]; ok {
		return val.(float64)
	}

	return 0
}

// Int gets the value as an integer or 0 if that is not possible.
func (sa SettingsAttributes) Int(name string) int {
	// Integer values may be interpreted as float64, so get it as that first,
	// then coerce down to int.
	floatVal := int(sa.Float64(name))
	return (floatVal)
}

// Bool gets the value as a boolean or returns false.
func (sa SettingsAttributes) Bool(name string) bool {
	maybeBool := sa.String(name)
	maybeBool = strings.ToLower(maybeBool)
	return (maybeBool == "true" ||
		maybeBool == "1" ||
		maybeBool == "enabled")
}
