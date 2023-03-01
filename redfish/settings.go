//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"fmt"
	"strings"
)

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
	// Integer values may be interpeted as float64, so get it as that first,
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
