//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

// MemoryPFA is an instance of a MemoryPFA object.
type MemoryPFA struct {
	schemas.Entity

	// Init shall contain the current state from Bios HII value.
	Init string `json:"MemoryPfaInit"`
	// Next shall contain the next status set by tools would like to change the state.
	Next string `json:"MemoryPfaNext"`
	// AlertID shall contain the forward SELs at one specific Alert Id registered by tool.
	AlertID int `json:"AlertId"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MemoryPFA object from the raw JSON.
func (i *MemoryPFA) UnmarshalJSON(b []byte) error {
	type temp MemoryPFA
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = MemoryPFA(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *MemoryPFA) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(MemoryPFA)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Next",
		"MemoryPfaNext",
		"AlertID",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemoryPFA will get a MemoryPFA instance from the service.
func GetMemoryPFA(c schemas.Client, uri string) (*MemoryPFA, error) {
	return schemas.GetObject[MemoryPFA](c, uri)
}
