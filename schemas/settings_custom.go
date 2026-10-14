//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SettingsAttributes handles the settings attribute values that may be any of several
// types and adds some basic helper methods to make accessing values easier.
type SettingsAttributes map[string]any

// MaintenanceWindowResource gets the MaintenanceWindowResource linked resource.
func (s *Settings) MaintenanceWindowResource(client Client) (*Entity, error) {
	if s.maintenanceWindowResource == "" {
		return nil, nil
	}
	return GetObject[Entity](client, s.maintenanceWindowResource)
}

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

// postUnmarshal records the SettingsObject and MaintenanceWindowResource links.
func (s *Settings) postUnmarshal(b []byte) error {
	var t struct {
		MaintenanceWindowResource Link
		SettingsObject            Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	s.maintenanceWindowResource = t.MaintenanceWindowResource.String()
	s.SettingsObject = t.SettingsObject.String()

	return nil
}
