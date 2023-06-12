//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// RedundancyMode is the redundancy mode.
type RedundancyMode string

const (
	// FailoverRedundancyMode Failure of one unit will automatically cause
	// its functions to be taken over by a standby or offline unit in the
	// redundancy set.
	FailoverRedundancyMode RedundancyMode = "Failover"
	// NMRedundancyMode Multiple units are available and active such that
	// normal operation will continue if one or more units fail.
	NMRedundancyMode RedundancyMode = "N+m"
	// SharingRedundancyMode Multiple units contribute or share such that
	// operation will continue, but at a reduced capacity, if one or more
	// units fail.
	SharingRedundancyMode RedundancyMode = "Sharing"
	// SparingRedundancyMode One or more spare units are available to take
	// over the function of a failed unit, but takeover is not automatic.
	SparingRedundancyMode RedundancyMode = "Sparing"
	// NotRedundantRedundancyMode The subsystem is not configured in a
	// redundancy mode, either due to configuration or the functionality has
	// been disabled by the user.
	NotRedundantRedundancyMode RedundancyMode = "NotRedundant"
)

// Redundancy represents the Redundancy element property.
// All values for resources described by this schema shall comply to the
// requirements as described in the Redfish specification.  The value of
// this string shall be of the format for the reserved word *Redundancy*.
type Redundancy struct {
	common.Entity

	// MaxNumSupported shall contain the maximum number of members allowed in
	// the redundancy group.
	MaxNumSupported int
	// MemberID value of this string shall uniquely identify the member within
	// the collection.
	MemberID string `json:"MemberId"`
	// MinNumNeeded shall contain the minimum
	// number of members allowed in the redundancy group for the current
	// redundancy mode to still be fault tolerant.
	MinNumNeeded int
	// Mode shall contain the information about the redundancy mode of this
	// subsystem.
	Mode RedundancyMode
	// RedundancyEnabled shall be a boolean indicating whether the redundancy is
	// enabled.
	RedundancyEnabled bool
	// RedundancySet shall contain the ids of components that are part of this
	// redundancy set. The id values may or may not be dereferenceable.
	redundancySet []string
	// RedundancySetCount is the number of RedundancySets.
	RedundancySetCount int `json:"RedundancySet@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Redundancy object from the raw JSON.
func (redundancy *Redundancy) UnmarshalJSON(b []byte) error {
	type temp Redundancy
	var t struct {
		temp
		RedundancySet common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*redundancy = Redundancy(t.temp)
	redundancy.redundancySet = t.RedundancySet.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	redundancy.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (redundancy *Redundancy) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Redundancy)
	err := original.UnmarshalJSON(redundancy.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Mode",
		"RedundancyEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(redundancy).Elem()

	return redundancy.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRedundancy will get a Redundancy instance from the service.
func GetRedundancy(c common.Client, uri string) (*Redundancy, error) {
	var redundancy Redundancy
	return &redundancy, redundancy.Get(c, uri, &redundancy)
}

// ListReferencedRedundancies gets the collection of Redundancy from
// a provided reference.
func ListReferencedRedundancies(c common.Client, link string) ([]*Redundancy, error) { //nolint:dupl
	var result []*Redundancy
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Redundancy
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		redundancy, err := GetRedundancy(c, link)
		ch <- GetResult{Item: redundancy, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// The redundancy mode of the group.
type RedundancyType string

const (
	// Failure of one unit automatically causes a standby or offline unit in the redundancy set to take over its functions.
	FailoverRedundancyType RedundancyType = "Failover"
	// Multiple units are available and active such that normal operation will continue if one or more units fail.
	NPlusMRedundancyType RedundancyType = "NPlusM"
	// The subsystem is not configured in a redundancy mode, either due to configuration or the functionality has been disabled by the user.
	NotRedundantRedundancyType RedundancyType = "NotRedundant"
	//  Multiple units contribute or share such that operation will continue, but at a reduced capacity, if one or more units fail.
	SharingRedundancyType RedundancyType = "Sharing"
	// One or more spare units are available to take over the function of a failed unit, but takeover is not automatic.
	SparingRedundancyType RedundancyType = "Sparing"
)

// The redundancy information for the devices in a redundancy group.
type RedundantGroup struct {
	// The maximum number of devices supported in this redundancy group.
	MaxSupportedInGroup int64
	// The minimum number of devices needed for this group to be redundant.
	MinNeededInGroup int64
	// The links to the devices included in this redundancy group.
	redundancyGroup      []string
	RedundancyGroupCount int
	// The redundancy mode of the group.
	RedundancyType RedundancyType
	// The status and health of the resource and its subordinate or dependent resources
	Status common.Status
}

// UnmarshalJSON unmarshals a RedundancyGroup object from the raw JSON.
func (redundantGroup *RedundantGroup) UnmarshalJSON(b []byte) error {
	type temp RedundantGroup
	type groupReference struct {
		RedundancyGroup      common.Links
		RedundancyGroupCount int `json:"RedundancyGroup@odata.count"`
	}

	var t struct {
		temp
		Group groupReference
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*redundantGroup = RedundantGroup(t.temp)
	redundantGroup.redundancyGroup = t.Group.RedundancyGroup.ToStrings()
	redundantGroup.RedundancyGroupCount = t.Group.RedundancyGroupCount

	return nil
}
