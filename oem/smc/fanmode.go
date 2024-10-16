//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FanMode is an instance of a FanMode object.
type FanMode struct {
	common.Entity

	Mode           string
	AllowableModes []string `json:"Mode@Redfish.AllowableValues"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a FanMode object from the raw JSON.
func (i *FanMode) UnmarshalJSON(b []byte) error {
	type temp FanMode
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = FanMode(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *FanMode) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(FanMode)
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

// GetFanMode will get a FanMode instance from the service.
func GetFanMode(c common.Client, uri string) (*FanMode, error) {
	return common.GetObject[FanMode](c, uri)
}
