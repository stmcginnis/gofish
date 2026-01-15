//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0.7a - #SpareResourceSet.v1_0_2.SpareResourceSet

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SpareResourceSet The values define a set of spares of a particular type.
type SpareResourceSet struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OnHandLocation The location where this set of spares is kept.
	OnHandLocation common.Location
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SpareResourceSet object from the raw JSON.
func (s *SpareResourceSet) UnmarshalJSON(b []byte) error {
	type temp SpareResourceSet
	type sLinks struct {
		OnHandSpares         common.Links `json:"OnHandSpares"`
		ReplacementSpareSets common.Links `json:"ReplacementSpareSets"`
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
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SpareResourceSet) Update() error {
	readWriteFields := []string{
		"OnHandLocation",
		"OnLine",
		"ResourceType",
		"TimeToProvision",
		"TimeToReplenish",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSpareResourceSet will get a SpareResourceSet instance from the service.
func GetSpareResourceSet(c common.Client, uri string) (*SpareResourceSet, error) {
	return common.GetObject[SpareResourceSet](c, uri)
}

// ListReferencedSpareResourceSets gets the collection of SpareResourceSet from
// a provided reference.
func ListReferencedSpareResourceSets(c common.Client, link string) ([]*SpareResourceSet, error) {
	return common.GetCollectionObjects[SpareResourceSet](c, link)
}

// OnHandSpares gets the OnHandSpares linked resources.
func (s *SpareResourceSet) OnHandSpares(client common.Client) ([]*common.Entity, error) {
	return common.GetObjects[common.Entity](client, s.onHandSpares)
}

// ReplacementSpareSets gets the ReplacementSpareSets linked resources.
func (s *SpareResourceSet) ReplacementSpareSets(client common.Client) ([]*SpareResourceSet, error) {
	return common.GetObjects[SpareResourceSet](client, s.replacementSpareSets)
}
