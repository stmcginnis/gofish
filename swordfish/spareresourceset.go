//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SpareResourceSet define a set of spares of a particular type.
type SpareResourceSet struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// OnHandLocation is the location where this set of spares is kept.
	OnHandLocation common.Location
	// OnLine indicates if the set is online.
	OnLine bool
	// ResourceType is the type of resources in the set.
	ResourceType string
	// TimeToProvision is the amount of time needed to make an on-hand resource
	// available as a spare.
	TimeToProvision string
	// TimeToReplenish is the amount of time to needed replenish consumed on-hand
	// resources.
	TimeToReplenish string
	// OnHandSparesCount is the number of on hand spares.
	OnHandSparesCount int `json:"OnHandSpares@odata.count"`
	// ReplacementSpareSetsCount is the number of replacement spare sets.
	ReplacementSpareSetsCount int `json:"ReplacementSpareSets@odata.count"`
	// onHandSpares are links to available spares.
	onHandSpares []string
	// ReplacementSpareSets are other spare sets that can be utilized to
	// replenish this spare set.
	replacementSpareSets string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SpareResourceSet object from the raw JSON.
func (spareresourceset *SpareResourceSet) UnmarshalJSON(b []byte) error {
	type temp SpareResourceSet
	type links struct {
		OnHandSpares common.Links
		// OnHandSparesCount is the number of on hand spares.
		OnHandSparesCount int `json:"OnHandSpares@odata.count"`
		// ReplacementSpareSets are other spare sets that can be utilized to
		// replenish this spare set.
		ReplacementSpareSets common.Link
		// ReplacementSpareSets@odata.count is the number of replacement spare sets.
		ReplacementSpareSetsCount int `json:"ReplacementSpareSets@odata.count"`
	}
	var t struct {
		temp
		Links links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spareresourceset = SpareResourceSet(t.temp)

	// Extract the links to other entities for later
	spareresourceset.OnHandSparesCount = t.OnHandSparesCount
	spareresourceset.onHandSpares = t.Links.OnHandSpares.ToStrings()
	spareresourceset.ReplacementSpareSetsCount = t.ReplacementSpareSetsCount
	spareresourceset.replacementSpareSets = t.Links.ReplacementSpareSets.String()

	// This is a read/write object, so we need to save the raw object data for later
	spareresourceset.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (spareresourceset *SpareResourceSet) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SpareResourceSet)
	err := original.UnmarshalJSON(spareresourceset.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"OnLine",
		"ResourceType",
		"TimeToProvision",
		"TimeToReplenish",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(spareresourceset).Elem()

	return spareresourceset.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSpareResourceSet will get a SpareResourceSet instance from the service.
func GetSpareResourceSet(c common.Client, uri string) (*SpareResourceSet, error) {
	var spareResourceSet SpareResourceSet
	return &spareResourceSet, spareResourceSet.Get(c, uri, &spareResourceSet)
}

// ListReferencedSpareResourceSets gets the collection of SpareResourceSet from
// a provided reference.
func ListReferencedSpareResourceSets(c common.Client, link string) ([]*SpareResourceSet, error) { //nolint:dupl
	var result []*SpareResourceSet
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *SpareResourceSet
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		spareresourceset, err := GetSpareResourceSet(c, link)
		ch <- GetResult{Item: spareresourceset, Link: link, Error: err}
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

// ReplacementSpareSets gets other spare sets that can be utilized to replenish
// this spare set.
func (spareresourceset *SpareResourceSet) ReplacementSpareSets() ([]*SpareResourceSet, error) {
	return ListReferencedSpareResourceSets(spareresourceset.Client, spareresourceset.replacementSpareSets)
}
