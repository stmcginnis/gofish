//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// IKVM is an instance of a IKVM object.
type IKVM struct {
	common.Entity

	CurrentInterface string `json:"Current Interface"`
	URI              string

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IKVM object from the raw JSON.
func (i *IKVM) UnmarshalJSON(b []byte) error {
	type temp IKVM
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = IKVM(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IKVM) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(IKVM)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"CurrentInterface",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetIKVM will get a IKVM instance from the service.
func GetIKVM(c common.Client, uri string) (*IKVM, error) {
	return common.GetObject[IKVM](c, uri)
}
