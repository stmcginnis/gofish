//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.2 - #AggregationService.v1_0_3.AggregationService

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// AggregationService shall represent an aggregation service for a Redfish
// implementation.
type AggregationService struct {
	common.Entity
	// Aggregates shall contain a link to a resource collection of type
	// 'AggregateCollection'.
	aggregates string
	// AggregationSources shall contain a link to a resource collection of type
	// 'AggregationSourceCollection'.
	aggregationSources string
	// ConnectionMethods shall contain a link to a resource collection of type
	// 'ConnectionMethodCollection'.
	connectionMethods string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether the aggregation service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// setDefaultBootOrderTarget is the URL to send SetDefaultBootOrder requests.
	setDefaultBootOrderTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AggregationService object from the raw JSON.
func (a *AggregationService) UnmarshalJSON(b []byte) error {
	type temp AggregationService
	type aActions struct {
		Reset               common.ActionTarget `json:"#AggregationService.Reset"`
		SetDefaultBootOrder common.ActionTarget `json:"#AggregationService.SetDefaultBootOrder"`
	}
	var tmp struct {
		temp
		Actions            aActions
		Aggregates         common.Link `json:"aggregates"`
		AggregationSources common.Link `json:"aggregationSources"`
		ConnectionMethods  common.Link `json:"connectionMethods"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AggregationService(tmp.temp)

	// Extract the links to other entities for later
	a.resetTarget = tmp.Actions.Reset.Target
	a.setDefaultBootOrderTarget = tmp.Actions.SetDefaultBootOrder.Target
	a.aggregates = tmp.Aggregates.String()
	a.aggregationSources = tmp.AggregationSources.String()
	a.connectionMethods = tmp.ConnectionMethods.String()

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AggregationService) Update() error {
	readWriteFields := []string{
		"ServiceEnabled",
		"Status",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
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

// Reset shall perform a reset of a set of resources.
// batchSize - This parameter shall contain the number of elements in each
// batch simultaneously being issued a reset.
// delayBetweenBatchesInSeconds - This parameter shall contain the delay of the
// batches of elements being reset in seconds.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
// targetURIs - This parameter shall contain an array of links to the resources
// being reset.
func (a *AggregationService) Reset(batchSize uint, delayBetweenBatchesInSeconds uint, resetType common.ResetType, targetURIs []string) error {
	payload := make(map[string]any)
	payload["BatchSize"] = batchSize
	payload["DelayBetweenBatchesInSeconds"] = delayBetweenBatchesInSeconds
	payload["ResetType"] = resetType
	payload["TargetURIs"] = targetURIs
	return a.Post(a.resetTarget, payload)
}

// SetDefaultBootOrder shall restore the boot order to the default state for the
// specified computer systems.
// systems - This parameter shall contain an array of links to resources of
// type 'ComputerSystem'.
func (a *AggregationService) SetDefaultBootOrder(systems []string) error {
	payload := make(map[string]any)
	payload["Systems"] = systems
	return a.Post(a.setDefaultBootOrderTarget, payload)
}

// Aggregates gets the Aggregates collection.
func (a *AggregationService) Aggregates(client common.Client) ([]*Aggregate, error) {
	if a.aggregates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Aggregate](client, a.aggregates)
}

// AggregationSources gets the AggregationSources collection.
func (a *AggregationService) AggregationSources(client common.Client) ([]*AggregationSource, error) {
	if a.aggregationSources == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[AggregationSource](client, a.aggregationSources)
}

// ConnectionMethods gets the ConnectionMethods collection.
func (a *AggregationService) ConnectionMethods(client common.Client) ([]*ConnectionMethod, error) {
	if a.connectionMethods == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[ConnectionMethod](client, a.connectionMethods)
}
