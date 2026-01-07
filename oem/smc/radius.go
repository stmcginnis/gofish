//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

// RADIUS is an instance of a RADIUS object.
type RADIUS struct {
	schemas.Entity

	Enabled bool   `json:"RadiusEnabled"`
	Server  string `json:"RadiusServer"`
	Port    int    `json:"RadiusPortNumber"`
	Secret  string `json:"RadiusSecret"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a RADIUS object from the raw JSON.
func (r *RADIUS) UnmarshalJSON(b []byte) error {
	type temp RADIUS
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*r = RADIUS(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *RADIUS) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	rad := new(RADIUS)
	err := rad.UnmarshalJSON(r.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
		"RadiusEnabled",
		"ServerIP",
		"RadiusServerIP",
		"PortNumber",
		"RadiusPortNumber",
		"Secret",
		"RadiusSecret",
	}

	originalElement := reflect.ValueOf(rad).Elem()
	currentElement := reflect.ValueOf(r).Elem()

	return r.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRADIUS will get a RADIUS instance from the service.
func GetRADIUS(c schemas.Client, uri string) (*RADIUS, error) {
	return schemas.GetObject[RADIUS](c, uri)
}
