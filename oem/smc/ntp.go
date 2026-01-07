//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

// NTP is an instance of an NTP object.
type NTP struct {
	schemas.Entity

	Enabled            bool   `json:"NTPEnable"`
	PrimaryServer      string `json:"PrimaryNTPServer"`
	SecondaryServer    string `json:"SecondaryNTPServer"`
	DaylightSavingTime bool   `json:"DaylightSavingTime"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a NTP object from the raw JSON.
func (r *NTP) UnmarshalJSON(b []byte) error {
	type temp NTP
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*r = NTP(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *NTP) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	rad := new(NTP)
	err := rad.UnmarshalJSON(r.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
		"NTPEnable",
		"PrimaryServer",
		"PrimaryNTPServer",
		"SecondaryServer",
		"SecondaryNTPServer",
		"DaylightSavingTime",
	}

	originalElement := reflect.ValueOf(rad).Elem()
	currentElement := reflect.ValueOf(r).Elem()

	return r.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNTP will get a NTP instance from the service.
func GetNTP(c schemas.Client, uri string) (*NTP, error) {
	return schemas.GetObject[NTP](c, uri)
}
