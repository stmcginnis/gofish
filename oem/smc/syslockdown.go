//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SysLockdown is an instance of a SysLockdown object.
type SysLockdown struct {
	common.Entity

	Enabled bool `json:"SysLockdownEnabled"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SysLockdown object from the raw JSON.
func (i *SysLockdown) UnmarshalJSON(b []byte) error {
	type temp SysLockdown
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = SysLockdown(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *SysLockdown) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(SysLockdown)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
		"SysLockdownEnabled",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSysLockdown will get a SysLockdown instance from the service.
func GetSysLockdown(c common.Client, uri string) (*SysLockdown, error) {
	return common.GetObject[SysLockdown](c, uri)
}
