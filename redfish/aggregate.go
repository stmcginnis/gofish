//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.2 - #Aggregate.v1_0_3.Aggregate

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Aggregate shall represent an aggregation service grouping method for a
// Redfish implementation.
type Aggregate struct {
	common.Entity
	// Elements shall contain an array of links to the elements of this aggregate.
	elements []string
	// Elements@odata.count
	ElementsCount int `json:"Elements@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// addElementsTarget is the URL to send AddElements requests.
	addElementsTarget string
	// removeElementsTarget is the URL to send RemoveElements requests.
	removeElementsTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// setDefaultBootOrderTarget is the URL to send SetDefaultBootOrder requests.
	setDefaultBootOrderTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Aggregate object from the raw JSON.
func (a *Aggregate) UnmarshalJSON(b []byte) error {
	type temp Aggregate
	type aActions struct {
		AddElements         common.ActionTarget `json:"#Aggregate.AddElements"`
		RemoveElements      common.ActionTarget `json:"#Aggregate.RemoveElements"`
		Reset               common.ActionTarget `json:"#Aggregate.Reset"`
		SetDefaultBootOrder common.ActionTarget `json:"#Aggregate.SetDefaultBootOrder"`
	}
	var tmp struct {
		temp
		Actions  aActions
		Elements common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = Aggregate(tmp.temp)

	// Extract the links to other entities for later
	a.addElementsTarget = tmp.Actions.AddElements.Target
	a.removeElementsTarget = tmp.Actions.RemoveElements.Target
	a.resetTarget = tmp.Actions.Reset.Target
	a.setDefaultBootOrderTarget = tmp.Actions.SetDefaultBootOrder.Target

	a.elements = tmp.Elements.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *Aggregate) Update() error {
	readWriteFields := []string{
		"Elements@odata.count",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetAggregate will get a Aggregate instance from the service.
func GetAggregate(c common.Client, uri string) (*Aggregate, error) {
	return common.GetObject[Aggregate](c, uri)
}

// ListReferencedAggregates gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregates(c common.Client, link string) ([]*Aggregate, error) {
	return common.GetCollectionObjects[Aggregate](c, link)
}

// Elements get the elements of this aggregate.
func (a *Aggregate) Elements() ([]*common.Resource, error) {
	return common.GetObjects[common.Resource](a.GetClient(), a.elements)
}

// AddElements shall add one or more resources to the aggregate, with the
// result that the resources are included in the 'Elements' array of the
// aggregate.
// elements - This parameter shall contain an array of links to the specified
// resources to add to the aggregate's 'Elements' array.
func (a *Aggregate) AddElements(elements []string) error {
	payload := make(map[string]any)
	payload["Elements"] = elements
	return a.Post(a.addElementsTarget, payload)
}

// RemoveElements shall remove one or more resources from the aggregate, with the
// result that the resources are removed from the 'Elements' array of the
// aggregate.
// elements - This parameter shall contain an array of links to the specified
// resources to remove from the aggregate's 'Elements' array.
func (a *Aggregate) RemoveElements(elements []string) error {
	payload := make(map[string]any)
	payload["Elements"] = elements
	return a.Post(a.removeElementsTarget, payload)
}

// Reset shall perform a reset of a collection of resources.
// batchSize - This parameter shall contain the number of elements in each
// batch simultaneously being issued a reset.
// delayBetweenBatchesInSeconds - This parameter shall contain the delay of the
// batches of elements being reset in seconds.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
func (a *Aggregate) Reset(batchSize int, delayBetweenBatchesInSeconds uint, resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["BatchSize"] = batchSize
	payload["DelayBetweenBatchesInSeconds"] = delayBetweenBatchesInSeconds
	payload["ResetType"] = resetType
	return a.Post(a.resetTarget, payload)
}

// SetDefaultBootOrder shall restore the boot order to the default state for the
// computer systems that are members of this aggregate.
func (a *Aggregate) SetDefaultBootOrder() error {
	payload := make(map[string]any)
	return a.Post(a.setDefaultBootOrderTarget, payload)
}
