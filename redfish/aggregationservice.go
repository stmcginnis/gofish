//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// AggregationService shall represent an aggregation service for a Redfish implementation.
type AggregationService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
	return aggregationservice.ResetWithContext(common.ContextOf(aggregationservice.GetClient()), batchSize, delayBetweenBatchesInSeconds, resetType, targetURIs)
}

// ResetWithContext performs a reset of a set of resources.
// `batchSize` is the number of elements in each batch being reset.
// `delayBetweenBatchesInSeconds` is the delay of the batches of elements being reset.
// `resetType` is the type of reset to perform.
// `targetURIs` is an array of links to the resources being reset.
func (aggregationservice *AggregationService) ResetWithContext(ctx context.Context, batchSize, delayBetweenBatchesInSeconds int, resetType ResetType, targetURIs []string) error {
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
	return aggregationservice.PostWithContext(ctx, aggregationservice.resetTarget, t)
}

// SetDefaultBootOrder is used to restore the boot order to the default state
// for the specified computer systems.
// `systems` is an array of links to the ComputerSystems to be reset.
func (aggregationservice *AggregationService) SetDefaultBootOrder(systems []string) error {
	return aggregationservice.SetDefaultBootOrderWithContext(common.ContextOf(aggregationservice.GetClient()), systems)
}

// SetDefaultBootOrderWithContext is used to restore the boot order to the default state
// for the specified computer systems.
// `systems` is an array of links to the ComputerSystems to be reset.
func (aggregationservice *AggregationService) SetDefaultBootOrderWithContext(ctx context.Context, systems []string) error {
	t := struct {
		Systems []string
	}{
		Systems: systems,
	}
	return aggregationservice.PostWithContext(ctx, aggregationservice.setDefaultBootOrderTarget, t)
}

// Aggregates gets the aggregates associated with this service.
func (aggregationservice *AggregationService) Aggregates(queryOpts ...common.QueryGroupOption) ([]*Aggregate, error) {
	return aggregationservice.AggregatesWithContext(common.ContextOf(aggregationservice.GetClient()), queryOpts...)
}

// AggregatesWithContext gets the aggregates associated with this service.
func (aggregationservice *AggregationService) AggregatesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Aggregate, error) {
	return ListReferencedAggregatesWithContext(ctx, aggregationservice.GetClient(), aggregationservice.aggregates, queryOpts...)
}

// AggregationSources gets the aggregation sources associated with this service.
func (aggregationservice *AggregationService) AggregationSources(queryOpts ...common.QueryGroupOption) ([]*AggregationSource, error) {
	return aggregationservice.AggregationSourcesWithContext(common.ContextOf(aggregationservice.GetClient()), queryOpts...)
}

// AggregationSourcesWithContext gets the aggregation sources associated with this service.
func (aggregationservice *AggregationService) AggregationSourcesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*AggregationSource, error) {
	return ListReferencedAggregationSourcesWithContext(ctx, aggregationservice.GetClient(), aggregationservice.aggregationSources, queryOpts...)
}

// ConnectionMethods gets the connection methods associated with this service.
func (aggregationservice *AggregationService) ConnectionMethods(queryOpts ...common.QueryGroupOption) ([]*ConnectionMethod, error) {
	return aggregationservice.ConnectionMethodsWithContext(common.ContextOf(aggregationservice.GetClient()), queryOpts...)
}

// ConnectionMethodsWithContext gets the connection methods associated with this service.
func (aggregationservice *AggregationService) ConnectionMethodsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*ConnectionMethod, error) {
	return ListReferencedConnectionMethodsWithContext(ctx, aggregationservice.GetClient(), aggregationservice.connectionMethods, queryOpts...)
}

// Update commits updates to this object's properties to the running system.
func (aggregationservice *AggregationService) Update() error {
	return aggregationservice.UpdateWithContext(common.ContextOf(aggregationservice.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (aggregationservice *AggregationService) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"ServiceEnabled",
	}

	return aggregationservice.UpdateFromRawDataWithContext(ctx, aggregationservice, aggregationservice.rawData, readWriteFields)
}

// GetAggregationService will get a AggregationService instance from the service.
func GetAggregationService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*AggregationService, error) {
	return GetAggregationServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetAggregationServiceWithContext will get a AggregationService instance from the service.
func GetAggregationServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*AggregationService, error) {
	return common.GetObjectWithContext[AggregationService](ctx, c, uri, queryOpts...)
}

// ListReferencedAggregationServices gets the collection of AggregationService from
// a provided reference.
func ListReferencedAggregationServices(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*AggregationService, error) {
	return ListReferencedAggregationServicesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedAggregationServicesWithContext gets the collection of AggregationService from
// a provided reference.
func ListReferencedAggregationServicesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*AggregationService, error) {
	return common.GetCollectionObjectsWithContext[AggregationService](ctx, c, link, queryOpts...)
}
