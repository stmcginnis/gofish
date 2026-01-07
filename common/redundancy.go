//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #Redundancy.v1_6_0

package common

import "encoding/json"

type RedundancyMode string

const (
	// FailoverRedundancyMode Failure of one unit automatically causes a standby or
	// offline unit in the redundancy set to take over its functions.
	FailoverRedundancyMode RedundancyMode = "Failover"
	// NmRedundancyMode Multiple units are available and active such that normal
	// operation will continue if one or more units fail.
	NmRedundancyMode RedundancyMode = "N+m"
	// SharingRedundancyMode Multiple units contribute or share such that operation
	// will continue, but at a reduced capacity, if one or more units fail.
	SharingRedundancyMode RedundancyMode = "Sharing"
	// SparingRedundancyMode One or more spare units are available to take over the
	// function of a failed unit, but takeover is not automatic.
	SparingRedundancyMode RedundancyMode = "Sparing"
	// NotRedundantRedundancyMode The subsystem is not configured in a redundancy
	// mode, either due to configuration or the functionality has been disabled by
	// the user.
	NotRedundantRedundancyMode RedundancyMode = "NotRedundant"
)

type RedundancyType string

const (
	// FailoverRedundancyType shall indicate that a failure of one unit
	// automatically causes a standby or offline unit in the redundancy set to take
	// over its functions.
	FailoverRedundancyType RedundancyType = "Failover"
	// NPlusMRedundancyType shall indicate that the capacity or services provided
	// by the set of N+M devices can withstand failure of up to M units, with all
	// units in the group normally providing capacity or service.
	NPlusMRedundancyType RedundancyType = "NPlusM"
	// SharingRedundancyType Multiple units contribute or share such that operation
	// will continue, but at a reduced capacity, if one or more units fail.
	SharingRedundancyType RedundancyType = "Sharing"
	// SparingRedundancyType One or more spare units are available to take over the
	// function of a failed unit, but takeover is not automatic.
	SparingRedundancyType RedundancyType = "Sparing"
	// NotRedundantRedundancyType The subsystem is not configured in a redundancy
	// mode, either due to configuration or the functionality has been disabled by
	// the user.
	NotRedundantRedundancyType RedundancyType = "NotRedundant"
)

// Redundancy This object represents the redundancy element property.
type Redundancy struct {
	Entity
	// ActiveRedundancySet shall contain the links to the active resources that
	// represent the active devices that are part of this redundancy set. When
	// 'Mode' contains 'Failover', the failure of an active device shall cause a
	// member of this redundancy set to take over its function. When 'Mode'
	// contains 'N+m' or 'Sharing', all devices in the redundancy set in a
	// non-failed state should be considered active. When 'Mode' contains
	// 'Sparing', the failure of an active device shall cause one or more spares
	// that are available to take over the function.
	//
	// Version added: v1.6.0
	ActiveRedundancySet []Resource
	// ActiveRedundancySet@odata.count
	ActiveRedundancySetCount int `json:"ActiveRedundancySet@odata.count"`
	// MaxNumSupported shall contain the maximum number of members allowed in the
	// redundancy group.
	MaxNumSupported *int `json:",omitempty"`
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// MinNumNeeded shall contain the minimum number of members allowed in the
	// redundancy group for the current redundancy mode to still be fault tolerant.
	MinNumNeeded *int `json:",omitempty"`
	// Mode shall contain the information about the redundancy mode of this
	// subsystem.
	Mode RedundancyMode
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RedundancyEnabled shall indicate whether the redundancy is enabled.
	//
	// Version added: v1.1.0
	RedundancyEnabled bool
	// RedundancySet shall contain the links to components that are part of this
	// redundancy set.
	redundancySet []string
	// RedundancySet@odata.count
	RedundancySetCount int `json:"RedundancySet@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Redundancy object from the raw JSON.
func (redundancy *Redundancy) UnmarshalJSON(b []byte) error {
	type temp Redundancy
	var t struct {
		temp
		RedundancySet Links
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
	readWriteFields := []string{
		"Mode",
		"RedundancyEnabled",
	}

	return redundancy.UpdateFromRawData(redundancy, redundancy.rawData, readWriteFields)
}

// GetRedundancy will get a Redundancy instance from the service.
func GetRedundancy(c Client, uri string) (*Redundancy, error) {
	return GetObject[Redundancy](c, uri)
}

// ListReferencedRedundancies gets the collection of Redundancy from
// a provided reference.
func ListReferencedRedundancies(c Client, link string) ([]*Redundancy, error) {
	return GetCollectionObjects[Redundancy](c, link)
}

// RedundantGroup shall contain redundancy information for the set of devices in
// this redundancy group.
type RedundantGroup struct {
	// ActiveRedundancyGroup shall contain the links to the active resources that
	// represent the active devices that are part of this redundancy group. When
	// 'RedundancyType' contains 'Failover', the failure of an active device shall
	// cause a member of this redundancy group to take over its function. When
	// 'RedundancyType' contains 'NPlusM' or 'Sharing', all devices in the
	// redundancy set in a non-failed state should be considered active. When
	// 'RedundancyType' contains 'Sparing', the failure of an active device shall
	// cause one or more spares that are available to take over the function.
	//
	// Version added: v1.6.0
	ActiveRedundancyGroup []Resource
	// ActiveRedundancyGroup@odata.count
	ActiveRedundancyGroupCount int `json:"ActiveRedundancyGroup@odata.count"`
	// GroupName shall contain the name of the redundant group used to identify the
	// particular group of redundant resources. The value shall conform with the
	// 'Name' clause of the Redfish Specification.
	//
	// Version added: v1.5.0
	GroupName string
	// MaxSupportedInGroup shall contain the maximum number of devices allowed in
	// the redundancy group.
	//
	// Version added: v1.4.0
	MaxSupportedInGroup *int `json:",omitempty"`
	// MinNeededInGroup shall contain the minimum number of functional devices
	// needed in the redundancy group for the current redundancy mode to be fault
	// tolerant.
	//
	// Version added: v1.4.0
	MinNeededInGroup *int `json:",omitempty"`
	// RedundancyGroup shall contain the links to the resources that represent the
	// devices that are part of this redundancy group.
	//
	// Version added: v1.4.0
	RedundancyGroup []Resource
	// RedundancyGroup@odata.count
	RedundancyGroupCount int `json:"RedundancyGroup@odata.count"`
	// RedundancyType shall contain the information about the redundancy mode of
	// this redundancy group.
	//
	// Version added: v1.4.0
	RedundancyType RedundancyType
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.4.0
	Status Status
}
