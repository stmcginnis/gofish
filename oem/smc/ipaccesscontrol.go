//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type FilterRulePolicy string

const (
	FilterRulePolicyAllow FilterRulePolicy = "Allow"
	FilterRulePolicyDeny  FilterRulePolicy = "Deny"
)

// FilterRule represents an individual filter rule.
type FilterRule struct {
	common.Entity
	Address      string
	PrefixLength int
	Policy       FilterRulePolicy
}

// IPAccessControl is an instance of an IPAccessControl object.
type IPAccessControl struct {
	common.Entity

	Enabled     bool `json:"ServiceEnabled"`
	filterRules string

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IPAccessControl object from the raw JSON.
func (i *IPAccessControl) UnmarshalJSON(b []byte) error {
	type temp IPAccessControl
	var t struct {
		temp
		FilterRules common.Link `json:"FilterRules"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = IPAccessControl(t.temp)
	i.filterRules = t.FilterRules.String()

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IPAccessControl) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(IPAccessControl)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

func (i *IPAccessControl) FilterRules() ([]*FilterRule, error) {
	return common.GetCollectionObjects[FilterRule](i.GetClient(), i.filterRules)
}

// GetIPAccessControl will get a IPAccessControl instance from the service.
func GetIPAccessControl(c common.Client, uri string) (*IPAccessControl, error) {
	return common.GetObject[IPAccessControl](c, uri)
}
