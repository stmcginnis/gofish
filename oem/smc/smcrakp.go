//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

type SMCRAKPType string

const (
	SMCRAKPTypeEnabled  SMCRAKPType = "Enabled"
	SMCRAKPTypeDisabled SMCRAKPType = "Disabled"
)

// SMCRAKP is an instance of an SMCRAKP object.
type SMCRAKP struct {
	schemas.Entity

	Mode SMCRAKPType

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SMCRAKP object from the raw JSON.
func (i *SMCRAKP) UnmarshalJSON(b []byte) error {
	type temp SMCRAKP
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = SMCRAKP(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *SMCRAKP) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(SMCRAKP)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Mode",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSMCRAKP will get a SMCRAKP instance from the service.
func GetSMCRAKP(c schemas.Client, uri string) (*SMCRAKP, error) {
	return schemas.GetObject[SMCRAKP](c, uri)
}
