//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// MemoryHealthComp is an instance of a MemoryHealthComp object.
type MemoryHealthComp struct {
	common.Entity

	// Init shall contain the current state from Bios HII value.
	Init string `json:"MemoryHealthCompInit"`
	// Next shall contain the next status set by tools would like to change the state.
	Next string `json:"MemoryHealthCompNext"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MemoryHealthComp object from the raw JSON.
func (i *MemoryHealthComp) UnmarshalJSON(b []byte) error {
	type temp MemoryHealthComp
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = MemoryHealthComp(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *MemoryHealthComp) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(MemoryHealthComp)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Next",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemoryHealthComp will get a MemoryHealthComp instance from the service.
func GetMemoryHealthComp(c common.Client, uri string) (*MemoryHealthComp, error) {
	return common.GetObject[MemoryHealthComp](c, uri)
}
