//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

type MouseModeSetting string

const (
	MouseModeSettingSingle   MouseModeSetting = "Single"
	MouseModeSettingRelative MouseModeSetting = "Relative"
	MouseModeSettingAbsolute MouseModeSetting = "Absolute"
)

// MouseMode is an instance of a MouseMode object.
type MouseMode struct {
	schemas.Entity

	Mode MouseModeSetting `json:"Mode"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MouseMode object from the raw JSON.
func (r *MouseMode) UnmarshalJSON(b []byte) error {
	type temp MouseMode
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*r = MouseMode(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *MouseMode) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	rad := new(MouseMode)
	err := rad.UnmarshalJSON(r.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Mode",
	}

	originalElement := reflect.ValueOf(rad).Elem()
	currentElement := reflect.ValueOf(r).Elem()

	return r.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMouseMode will get a MouseMode instance from the service.
func GetMouseMode(c schemas.Client, uri string) (*MouseMode, error) {
	return schemas.GetObject[MouseMode](c, uri)
}
