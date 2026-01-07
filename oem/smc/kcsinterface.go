//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

// KCSInterface is an instance of a KCSInterface object.
type KCSInterface struct {
	schemas.Entity

	// Privilege shall contain the current KCS privilege setting.
	Privilege string

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a KCSInterface object from the raw JSON.
func (i *KCSInterface) UnmarshalJSON(b []byte) error {
	type temp KCSInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = KCSInterface(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *KCSInterface) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(KCSInterface)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Privilege",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetKCSInterface will get a KCSInterface instance from the service.
func GetKCSInterface(c schemas.Client, uri string) (*KCSInterface, error) {
	return schemas.GetObject[KCSInterface](c, uri)
}
