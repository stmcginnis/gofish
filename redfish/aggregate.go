//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Aggregate shall represent an aggregation service grouping method for a Redfish implementation.
type Aggregate struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Elements shall contain an array of links to the elements of this aggregate.
	Elements []Resource
	// ElementsCount shall contain the number of entries in the Elements array.
	ElementsCount int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`

	addElementsTarget         string
	removeElementsTarget      string
	resetTarget               string
	setDefaultBootOrderTarget string
}

// UnmarshalJSON unmarshals a Aggregate object from the raw JSON.
func (aggregate *Aggregate) UnmarshalJSON(b []byte) error {
	type temp Aggregate
	type Actions struct {
		AddElements struct {
			Target string
		} `json:"#Aggregate.AddElements"`
		RemoveElements struct {
			Target string
		} `json:"#Aggregate.RemoveElements"`
		Reset struct {
			Target string
		} `json:"#Aggregate.Reset"`
		SetDefaultBootOrder struct {
			Target string
		} `json:"#Aggregate.SetDefaultBootOrder"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregate = Aggregate(t.temp)

	// Extract the links to other entities for later
	aggregate.addElementsTarget = t.Actions.AddElements.Target
	aggregate.removeElementsTarget = t.Actions.RemoveElements.Target
	aggregate.resetTarget = t.Actions.Reset.Target
	aggregate.setDefaultBootOrderTarget = t.Actions.SetDefaultBootOrder.Target

	return nil
}

// AddElements adds one or more resources to the aggregate.
func (aggregate *Aggregate) AddElements(elements []*Resource) error {
	t := struct {
		Elements []*Resource
	}{
		Elements: elements,
	}
	return aggregate.Post(aggregate.addElementsTarget, t)
}

// RemoveElements removes one or more resources from the aggregate.
func (aggregate *Aggregate) RemoveElements(elements []*Resource) error {
	t := struct {
		Elements []*Resource
	}{
		Elements: elements,
	}
	return aggregate.Post(aggregate.removeElementsTarget, t)
}

// Reset performs a reset of a collection of resources.
// `batchSize` is the number of elements in each batch being reset.
// `delayBetweenBatchesInSeconds` is the delay of the batches of elements being reset.
// `resetType` is the type of reset to perform.
func (aggregate *Aggregate) Reset(batchSize, delayBetweenBatchesInSeconds int, resetType ResetType) error {
	t := struct {
		BatchSize                    int
		DelayBetweenBatchesInSeconds int
		ResetType                    ResetType
	}{
		BatchSize:                    batchSize,
		DelayBetweenBatchesInSeconds: delayBetweenBatchesInSeconds,
		ResetType:                    resetType,
	}
	return aggregate.Post(aggregate.resetTarget, t)
}

// SetDefaultBootOrder is used to restore the boot order to the default state for the
// computer systems that are members of this aggregate.
func (aggregate *Aggregate) SetDefaultBootOrder() error {
	return aggregate.Post(aggregate.setDefaultBootOrderTarget, nil)
}

// GetAggregate will get a Aggregate instance from the service.
func GetAggregate(c common.Client, uri string) (*Aggregate, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aggregate Aggregate
	err = json.NewDecoder(resp.Body).Decode(&aggregate)
	if err != nil {
		return nil, err
	}

	aggregate.SetClient(c)
	return &aggregate, nil
}

// ListReferencedAggregates gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregates(c common.Client, link string) ([]*Aggregate, error) {
	var result []*Aggregate
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Aggregate
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		aggregate, err := GetAggregate(c, link)
		ch <- GetResult{Item: aggregate, Link: link, Error: err}
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
