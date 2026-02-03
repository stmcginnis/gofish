//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0.7a - #SpareResourceSet.v1_0_2.SpareResourceSet

package schemas

import (
	"encoding/json"
)

// SpareResourceSet The values define a set of spares of a particular type.
type SpareResourceSet struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OnHandLocation The location where this set of spares is kept.
	OnHandLocation Location
	// OnLine shall be available online.
	OnLine bool
	// ResourceType The type of resources in the set.
	ResourceType string
	// TimeToProvision Amount of time needed to make an on-hand resource available
	// as a spare.
	TimeToProvision string
	// TimeToReplenish Amount of time needed to replenish consumed on-hand
	// resources.
	TimeToReplenish string
	// onHandSpares are the URIs for OnHandSpares.
	onHandSpares []string
	// replacementSpareSets are the URIs for ReplacementSpareSets.
	replacementSpareSets []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SpareResourceSet object from the raw JSON.
func (s *SpareResourceSet) UnmarshalJSON(b []byte) error {
	type temp SpareResourceSet
	type sLinks struct {
		OnHandSpares         Links `json:"OnHandSpares"`
		ReplacementSpareSets Links `json:"ReplacementSpareSets"`
	}
	var tmp struct {
		temp
		Links sLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SpareResourceSet(tmp.temp)

	// Extract the links to other entities for later
	s.onHandSpares = tmp.Links.OnHandSpares.ToStrings()
	s.replacementSpareSets = tmp.Links.ReplacementSpareSets.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SpareResourceSet) Update() error {
	readWriteFields := []string{
		"OnLine",
		"ResourceType",
		"TimeToProvision",
		"TimeToReplenish",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSpareResourceSet will get a SpareResourceSet instance from the service.
func GetSpareResourceSet(c Client, uri string) (*SpareResourceSet, error) {
	return GetObject[SpareResourceSet](c, uri)
}

// ListReferencedSpareResourceSets gets the collection of SpareResourceSet from
// a provided reference.
func ListReferencedSpareResourceSets(c Client, link string) ([]*SpareResourceSet, error) {
	return GetCollectionObjects[SpareResourceSet](c, link)
}

// OnHandSpares gets the OnHandSpares linked resources.
func (s *SpareResourceSet) OnHandSpares() ([]*Entity, error) {
	return GetObjects[Entity](s.client, s.onHandSpares)
}

// ReplacementSpareSets gets the ReplacementSpareSets linked resources.
func (s *SpareResourceSet) ReplacementSpareSets() ([]*SpareResourceSet, error) {
	return GetObjects[SpareResourceSet](s.client, s.replacementSpareSets)
}
