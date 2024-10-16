//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Syslog is an instance of a Syslog object.
type Syslog struct {
	common.Entity

	Enabled bool   `json:"EnableSyslog"`
	Server  string `json:"SyslogServer"`
	Port    int    `json:"SyslogPortNumber"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Syslog object from the raw JSON.
func (i *Syslog) UnmarshalJSON(b []byte) error {
	type temp Syslog
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = Syslog(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *Syslog) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(Syslog)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
		"Port",
		"Server",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSyslog will get a Syslog instance from the service.
func GetSyslog(c common.Client, uri string) (*Syslog, error) {
	return common.GetObject[Syslog](c, uri)
}
