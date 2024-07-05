//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AggregationService shall represent an aggregation service for a Redfish implementation.
type AggregationService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Aggregates shall contain a link to a resource collection of type AggregateCollection.
	aggregates string
	// AggregationSources shall contain a link to a resource collection of type AggregationSourceCollection.
	aggregationSources string
	// ConnectionMethods shall contain a link to a resource collection of type ConnectionMethodCollection.
	connectionMethods string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether the aggregation service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	resetTarget               string
	setDefaultBootOrderTarget string
}

// UnmarshalJSON unmarshals a AggregationService object from the raw JSON.
func (aggregationservice *AggregationService) UnmarshalJSON(b []byte) error {
	type temp AggregationService
	type Actions struct {
		Reset               common.ActionTarget `json:"#AggregationService.Reset"`
		SetDefaultBootOrder common.ActionTarget `json:"#AggregationService.SetDefaultBootOrder"`
	}
	var t struct {
		temp
		Actions            Actions
		Aggregates         common.Link
		AggregationSources common.Link
		ConnectionMethods  common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregationservice = AggregationService(t.temp)

	// Extract the links to other entities for later
	aggregationservice.resetTarget = t.Actions.Reset.Target
	aggregationservice.setDefaultBootOrderTarget = t.Actions.SetDefaultBootOrder.Target

	aggregationservice.aggregates = t.Aggregates.String()
	aggregationservice.aggregationSources = t.AggregationSources.String()
	aggregationservice.connectionMethods = t.ConnectionMethods.String()

	// This is a read/write object, so we need to save the raw object data for later
	aggregationservice.rawData = b

	return nil
}

// Reset performs a reset of a set of resources.
// `batchSize` is the number of elements in each batch being reset.
// `delayBetweenBatchesInSeconds` is the delay of the batches of elements being reset.
// `resetType` is the type of reset to perform.
// `targetURIs` is an array of links to the resources being reset.
func (aggregationservice *AggregationService) Reset(batchSize, delayBetweenBatchesInSeconds int, resetType ResetType, targetURIs []string) error {
	t := struct {
		BatchSize                    int
		DelayBetweenBatchesInSeconds int
		ResetType                    ResetType
		TargetURIs                   []string
	}{
		BatchSize:                    batchSize,
		DelayBetweenBatchesInSeconds: delayBetweenBatchesInSeconds,
		ResetType:                    resetType,
		TargetURIs:                   targetURIs,
	}
	return aggregationservice.Post(aggregationservice.resetTarget, t)
}

// SetDefaultBootOrder is used to restore the boot order to the default state
// for the specified computer systems.
// `systems` is an array of links to the ComputerSystems to be reset.
func (aggregationservice *AggregationService) SetDefaultBootOrder(systems []string) error {
	t := struct {
		Systems []string
	}{
		Systems: systems,
	}
	return aggregationservice.Post(aggregationservice.setDefaultBootOrderTarget, t)
}

// Aggregates gets the aggregates associated with this service.
func (aggregationservice *AggregationService) Aggregates() ([]*Aggregate, error) {
	return ListReferencedAggregates(aggregationservice.GetClient(), aggregationservice.aggregates)
}

// AggregationSources gets the aggregation sources associated with this service.
func (aggregationservice *AggregationService) AggregationSources() ([]*AggregationSource, error) {
	return ListReferencedAggregationSources(aggregationservice.GetClient(), aggregationservice.aggregationSources)
}

// ConnectionMethods gets the connection methods associated with this service.
func (aggregationservice *AggregationService) ConnectionMethods() ([]*ConnectionMethod, error) {
	return ListReferencedConnectionMethods(aggregationservice.GetClient(), aggregationservice.connectionMethods)
}

// Update commits updates to this object's properties to the running system.
func (aggregationservice *AggregationService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AggregationService)
	original.UnmarshalJSON(aggregationservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(aggregationservice).Elem()

	return aggregationservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAggregationService will get a AggregationService instance from the service.
func GetAggregationService(c common.Client, uri string) (*AggregationService, error) {
	return common.GetObject[AggregationService](c, uri)
}

// ListReferencedAggregationServices gets the collection of AggregationService from
// a provided reference.
func ListReferencedAggregationServices(c common.Client, link string) ([]*AggregationService, error) {
	return common.GetCollectionObjects[AggregationService](c, link)
}
