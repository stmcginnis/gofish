//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"fmt"
	"strings"
)

type ApplyTime string

const (
	// ImmediateApplyTime shall indicate the values within the settings resource are applied immediately. This value
	// may result in an immediate host reset, manager reset, or other side effects.
	ImmediateApplyTime ApplyTime = "Immediate"
	// OnResetApplyTime shall indicate the values within settings resource are applied when the system or service is
	// reset.
	OnResetApplyTime ApplyTime = "OnReset"
	// AtMaintenanceWindowStartApplyTime shall indicate the values within the settings resource are applied during the
	// maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties. A service can perform resets during this maintenance window.
	AtMaintenanceWindowStartApplyTime ApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetApplyTime shall indicate the values within the settings resource are applied during
	// the maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties, and if a reset occurs within the maintenance window.
	InMaintenanceWindowOnResetApplyTime ApplyTime = "InMaintenanceWindowOnReset"
)

// MaintenanceWindow shall indicate that a resource has a maintenance window assignment for applying settings or
// operations. Other resources can link to this object to convey a common control surface for the configuration of
// the maintenance window.
type MaintenanceWindow struct {
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the requested
	// settings or operation as part of a maintenance window. Services shall provide a default value if not configured
	// by a user.
	MaintenanceWindowStartTime string
}

type OperationApplyTime string

const (
	// AtMaintenanceWindowStartOperationApplyTime indicates the requested operation is applied within
	// the administrator-specified maintenance window.
	AtMaintenanceWindowStartOperationApplyTime OperationApplyTime = "AtMaintenanceWindowStart"
	// ImmediateOperationApplyTime indicates the requested operation is applied immediately.
	// This value might result in an immediate host reset, manager reset, or other side effects.
	ImmediateOperationApplyTime OperationApplyTime = "Immediate"
	// InMaintenanceWindowOnResetOperationApplyTime indicates the requested operation is applied after a reset but
	// within the administrator-specified maintenance window.
	InMaintenanceWindowOnResetOperationApplyTime OperationApplyTime = "InMaintenanceWindowOnReset"
	// OnResetOperationApplyTime indicates the requested operation is applied on a reset.
	OnResetOperationApplyTime OperationApplyTime = "OnReset"
	// OnStartUpdateRequestOperationApplyTime indicates the requested operation is applied when the
	// StartUpdate action of the update service is invoked.
	OnStartUpdateRequestOperationApplyTime OperationApplyTime = "OnStartUpdateRequest"
)

type OperationApplyTimeSupport struct {
	// MaintenanceWindowDurationInSeconds shall contain the same as the MaintenanceWindowDurationInSeconds property
	// found in the MaintenanceWindow structure on the MaintenanceWindowResource. This property shall be required if
	// the SupportedValues property contains 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowResource shall contain a link to a resource that contains the @Redfish.MaintenanceWindow
	// property that governs this resource. This property shall be required if the SupportedValues property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowResource string
	// MaintenanceWindowStartTime shall contain the same as the MaintenanceWindowStartTime property found in the
	// MaintenanceWindow structure on the MaintenanceWindowResource. Services shall provide a default value if not
	// configured by a user. This property shall be required if the SupportedValues property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
	// SupportedValues shall indicate the types of apply times the client can request when performing a create, delete,
	// or action operation.
	SupportedValues []OperationApplyTime
}

type PreferredApplyTime struct {
	// ApplyTime shall indicate when to apply the values in this settings resource.
	ApplyTime ApplyTime
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property. This property shall be required if the
	// ApplyTime property is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the future
	// configuration as part of a maintenance window. Services shall provide a default value if not configured by a
	// user. This property shall be required if the ApplyTime property is 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
}

// Settings shall describe any settings of a resource.
type Settings struct {
	// ETag shall contain the entity tag (ETag) of the resource to which the settings were applied, after the
	// application. The client can check this value against the ETag of this resource to determine whether the resource
	// had other changes.
	ETag string
	// MaintenanceWindowResource shall contain a link to a resource that contains the @Redfish.MaintenanceWindow
	// property that governs this resource. This property should be supported if the SupportedApplyTimes property
	// contains 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowResource string
	// Messages shall contain an array of messages associated with the settings.
	Messages []Message
	// SettingsObject shall contain the URI of the resource that the client can PUT or PATCH to modify the resource.
	SettingsObject string
	// SupportedApplyTimes shall contain the supported apply time values a client is allowed to request when
	// configuring the settings apply time. Services that do not support clients configuring the apply time can support
	// this property with a single array member in order to inform the client when the settings will be applied.
	SupportedApplyTimes []ApplyTime
	// Time shall indicate the time when the settings were applied to the resource.
	Time string
}

// SettingsAttributes handles the settings attribute values that may be any of several
// types and adds some basic helper methods to make accessing values easier.
type SettingsAttributes map[string]interface{}

// String gets the string representation of the attribute value.
func (ba SettingsAttributes) String(name string) string {
	if val, ok := ba[name]; ok {
		return fmt.Sprintf("%v", val)
	}

	return ""
}

// Float64 gets the value as a float64 or 0 if that is not possible.
func (ba SettingsAttributes) Float64(name string) float64 {
	if val, ok := ba[name]; ok {
		return val.(float64)
	}

	return 0
}

// Int gets the value as an integer or 0 if that is not possible.
func (ba SettingsAttributes) Int(name string) int {
	// Integer values may be interpreted as float64, so get it as that first,
	// then coerce down to int.
	floatVal := int(ba.Float64(name))
	return (floatVal)
}

// Bool gets the value as a boolean or returns false.
func (ba SettingsAttributes) Bool(name string) bool {
	maybeBool := ba.String(name)
	maybeBool = strings.ToLower(maybeBool)
	return (maybeBool == "true" ||
		maybeBool == "1" ||
		maybeBool == "enabled")
}
