//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/coreweave/gofish/common"
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
	return i.UpdateWithContext(common.ContextOf(i.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (i *MemoryHealthComp) UpdateWithContext(ctx context.Context) error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(MemoryHealthComp)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Next",
		"MemoryHealthCompNext",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.UpdateWithContext(ctx, originalElement, currentElement, readWriteFields)
}

// GetMemoryHealthComp will get a MemoryHealthComp instance from the service.
func GetMemoryHealthComp(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*MemoryHealthComp, error) {
	return GetMemoryHealthCompWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetMemoryHealthCompWithContext will get a MemoryHealthComp instance from the service.
func GetMemoryHealthCompWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*MemoryHealthComp, error) {
	return common.GetObjectWithContext[MemoryHealthComp](ctx, c, uri, queryOpts...)
}
