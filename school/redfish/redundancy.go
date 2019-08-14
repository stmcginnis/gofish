//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
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

	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// MaxNumSupported is The value of this property shall contain the
	// maximum number of members allowed in the redundancy group.
	MaxNumSupported int
	// MemberID ivalue of this string shall uniquely identify the member within
	// the collection.
	MemberID string `json:"MemberId"`
	// MinNumNeeded is The value of this property shall contain the minimum
	// number of members allowed in the redundancy group for the current
	// redundancy mode to still be fault tolerant.
	MinNumNeeded int
	// Mode is The value of this property shall contain the information about
	// the redundancy mode of this subsystem.
	Mode RedundancyMode
	// RedundancyEnabled is The value of this property shall be a boolean
	// indicating whether the redundancy is enabled.
	RedundancyEnabled bool
	// RedundancySet is The value of this property shall contain the ids of
	// components that are part of this redundancy set. The id values may or
	// may not be dereferenceable.
	RedundancySet []string
	// RedundancySetCount is
	RedundancySetCount int `json:"RedundancySet@odata.count"`
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// GetRedundancy will get a Redundancy instance from the service.
func GetRedundancy(c common.Client, uri string) (*Redundancy, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var redundancy Redundancy
	err = json.NewDecoder(resp.Body).Decode(&redundancy)
	if err != nil {
		return nil, err
	}

	redundancy.SetClient(c)
	return &redundancy, nil
}

// ListReferencedRedundancies gets the collection of Redundancy from
// a provided reference.
func ListReferencedRedundancies(c common.Client, link string) ([]*Redundancy, error) {
	var result []*Redundancy
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, redundancyLink := range links.ItemLinks {
		redundancy, err := GetRedundancy(c, redundancyLink)
		if err != nil {
			return result, err
		}
		result = append(result, redundancy)
	}

	return result, nil
}
