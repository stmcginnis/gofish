//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.2 - #Aggregate.v1_0_3.Aggregate

package schemas

import (
	"encoding/json"
)

// Aggregate shall represent an aggregation service grouping method for a
// Redfish implementation.
type Aggregate struct {
	Entity
	// Elements shall contain an array of links to the elements of this aggregate.
	elements []string
	// ElementsCount shall contain the number of entries in the 'Elements' array.
	ElementsCount *uint `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
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
}

// UnmarshalJSON unmarshals a Aggregate object from the raw JSON.
func (a *Aggregate) UnmarshalJSON(b []byte) error {
	type temp Aggregate
	type aActions struct {
		AddElements         ActionTarget `json:"#Aggregate.AddElements"`
		RemoveElements      ActionTarget `json:"#Aggregate.RemoveElements"`
		Reset               ActionTarget `json:"#Aggregate.Reset"`
		SetDefaultBootOrder ActionTarget `json:"#Aggregate.SetDefaultBootOrder"`
	}
	var tmp struct {
		temp
		Actions  aActions
		Elements Links `json:"Elements"`
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

	return nil
}

// GetAggregate will get a Aggregate instance from the service.
func GetAggregate(c Client, uri string) (*Aggregate, error) {
	return GetObject[Aggregate](c, uri)
}

// ListReferencedAggregates gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregates(c Client, link string) ([]*Aggregate, error) {
	return GetCollectionObjects[Aggregate](c, link)
}

// This action shall add one or more resources to the aggregate, with the
// result that the resources are included in the 'Elements' array of the
// aggregate.
// elements - This parameter shall contain an array of links to the specified
// resources to add to the aggregate's 'Elements' array.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *Aggregate) AddElements(elements []string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Elements"] = elements
	resp, taskInfo, err := PostWithTask(a.client,
		a.addElementsTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall remove one or more resources from the aggregate, with the
// result that the resources are removed from the 'Elements' array of the
// aggregate.
// elements - This parameter shall contain an array of links to the specified
// resources to remove from the aggregate's 'Elements' array.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *Aggregate) RemoveElements(elements []string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Elements"] = elements
	resp, taskInfo, err := PostWithTask(a.client,
		a.removeElementsTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall perform a reset of a collection of resources.
// batchSize - This parameter shall contain the number of elements in each
// batch simultaneously being issued a reset.
// delayBetweenBatchesInSeconds - This parameter shall contain the delay of the
// batches of elements being reset in seconds.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *Aggregate) Reset(batchSize uint, delayBetweenBatchesInSeconds uint, resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["BatchSize"] = batchSize
	payload["DelayBetweenBatchesInSeconds"] = delayBetweenBatchesInSeconds
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(a.client,
		a.resetTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall restore the boot order to the default state for the
// computer systems that are members of this aggregate.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *Aggregate) SetDefaultBootOrder() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.setDefaultBootOrderTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Elements gets the Elements linked resources.
func (a *Aggregate) Elements() ([]*Resource, error) {
	return GetObjects[Resource](a.client, a.elements)
}
