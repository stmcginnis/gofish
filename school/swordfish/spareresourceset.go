// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// SpareResourceSet is The values define a set of spares of a particular
// type.
type SpareResourceSet struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// OnHandLocation is The location where this set of spares is kept.
	OnHandLocation common.Location
	// OnLine is This set shall be online.
	OnLine bool
	// ResourceType is The type of resources in the set.
	ResourceType string
	// TimeToProvision is Amount of time  needed to make an on-hand resource
	// available as a spare.
	TimeToProvision string
	// TimeToReplenish is Amount of time to needed replenish consumed on-hand
	// resources.
	TimeToReplenish string
	// OnHandSparesCount is the number of on hand spares.
	OnHandSparesCount int `json:"OnHandSpares@odata.count"`
	// ReplacementSpareSets@odata.count is the number of replacement spare sets.
	ReplacementSpareSetsCount int `json:"ReplacementSpareSets@odata.count"`
	// onHandSpares are links to available spares.
	onHandSpares []string
	// ReplacementSpareSets are other spare sets that can be utilized to
	// replenish this spare set.
	replacementSpareSets []string
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
		ReplacementSpareSets common.Links
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
	spareresourceset.replacementSpareSets = t.Links.ReplacementSpareSets.ToStrings()

	return nil
}

// GetSpareResourceSet will get a SpareResourceSet instance from the service.
func GetSpareResourceSet(c common.Client, uri string) (*SpareResourceSet, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var spareresourceset SpareResourceSet
	err = json.NewDecoder(resp.Body).Decode(&spareresourceset)
	if err != nil {
		return nil, err
	}

	spareresourceset.SetClient(c)
	return &spareresourceset, nil
}

// ListReferencedSpareResourceSets gets the collection of SpareResourceSet from
// a provided reference.
func ListReferencedSpareResourceSets(c common.Client, link string) ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, spareresourcesetLink := range links.ItemLinks {
		spareresourceset, err := GetSpareResourceSet(c, spareresourcesetLink)
		if err != nil {
			return result, err
		}
		result = append(result, spareresourceset)
	}

	return result, nil
}
