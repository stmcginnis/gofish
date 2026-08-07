//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// Aggregate shall represent an aggregation service grouping method for a Redfish implementation.
type Aggregate struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Elements shall contain an array of links to the elements of this aggregate.
	elements []string
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
		AddElements         common.ActionTarget `json:"#Aggregate.AddElements"`
		RemoveElements      common.ActionTarget `json:"#Aggregate.RemoveElements"`
		Reset               common.ActionTarget `json:"#Aggregate.Reset"`
		SetDefaultBootOrder common.ActionTarget `json:"#Aggregate.SetDefaultBootOrder"`
	}
	var t struct {
		temp
		Actions  Actions
		Elements common.Links
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

	aggregate.elements = t.Elements.ToStrings()

	return nil
}

// Elements get the elements of this aggregate.
func (aggregate *Aggregate) Elements(queryOpts ...common.QueryGroupOption) ([]*Resource, error) {
	return aggregate.ElementsWithContext(common.ContextOf(aggregate.GetClient()), queryOpts...)
}

// ElementsWithContext get the elements of this aggregate.
func (aggregate *Aggregate) ElementsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Resource, error) {
	return common.GetObjectsWithContext[Resource](ctx, aggregate.GetClient(), aggregate.elements, queryOpts...)
}

// AddElements adds one or more resources to the aggregate.
func (aggregate *Aggregate) AddElements(elements []*Resource) error {
	return aggregate.AddElementsWithContext(common.ContextOf(aggregate.GetClient()), elements)
}

// AddElementsWithContext adds one or more resources to the aggregate.
func (aggregate *Aggregate) AddElementsWithContext(ctx context.Context, elements []*Resource) error {
	t := struct {
		Elements []*Resource
	}{
		Elements: elements,
	}
	return aggregate.PostWithContext(ctx, aggregate.addElementsTarget, t)
}

// RemoveElements removes one or more resources from the aggregate.
func (aggregate *Aggregate) RemoveElements(elements []*Resource) error {
	return aggregate.RemoveElementsWithContext(common.ContextOf(aggregate.GetClient()), elements)
}

// RemoveElementsWithContext removes one or more resources from the aggregate.
func (aggregate *Aggregate) RemoveElementsWithContext(ctx context.Context, elements []*Resource) error {
	t := struct {
		Elements []*Resource
	}{
		Elements: elements,
	}
	return aggregate.PostWithContext(ctx, aggregate.removeElementsTarget, t)
}

// Reset performs a reset of a collection of resources.
// `batchSize` is the number of elements in each batch being reset.
// `delayBetweenBatchesInSeconds` is the delay of the batches of elements being reset.
// `resetType` is the type of reset to perform.
func (aggregate *Aggregate) Reset(batchSize, delayBetweenBatchesInSeconds int, resetType ResetType) error {
	return aggregate.ResetWithContext(common.ContextOf(aggregate.GetClient()), batchSize, delayBetweenBatchesInSeconds, resetType)
}

// ResetWithContext performs a reset of a collection of resources.
// `batchSize` is the number of elements in each batch being reset.
// `delayBetweenBatchesInSeconds` is the delay of the batches of elements being reset.
// `resetType` is the type of reset to perform.
func (aggregate *Aggregate) ResetWithContext(ctx context.Context, batchSize, delayBetweenBatchesInSeconds int, resetType ResetType) error {
	t := struct {
		BatchSize                    int
		DelayBetweenBatchesInSeconds int
		ResetType                    ResetType
	}{
		BatchSize:                    batchSize,
		DelayBetweenBatchesInSeconds: delayBetweenBatchesInSeconds,
		ResetType:                    resetType,
	}
	return aggregate.PostWithContext(ctx, aggregate.resetTarget, t)
}

// SetDefaultBootOrder is used to restore the boot order to the default state for the
// computer systems that are members of this aggregate.
func (aggregate *Aggregate) SetDefaultBootOrder() error {
	return aggregate.SetDefaultBootOrderWithContext(common.ContextOf(aggregate.GetClient()))
}

// SetDefaultBootOrderWithContext is used to restore the boot order to the default state for the
// computer systems that are members of this aggregate.
func (aggregate *Aggregate) SetDefaultBootOrderWithContext(ctx context.Context) error {
	return aggregate.PostWithContext(ctx, aggregate.setDefaultBootOrderTarget, nil)
}

// GetAggregate will get a Aggregate instance from the service.
func GetAggregate(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Aggregate, error) {
	return GetAggregateWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetAggregateWithContext will get a Aggregate instance from the service.
func GetAggregateWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Aggregate, error) {
	return common.GetObjectWithContext[Aggregate](ctx, c, uri, queryOpts...)
}

// ListReferencedAggregates gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregates(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Aggregate, error) {
	return ListReferencedAggregatesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedAggregatesWithContext gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregatesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Aggregate, error) {
	return common.GetCollectionObjectsWithContext[Aggregate](ctx, c, link, queryOpts...)
}
