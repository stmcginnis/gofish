//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/AggregationService.v1_0_3.json
// 2020.2 - #AggregationService.v1_0_3.AggregationService

package schemas

import (
	"encoding/json"
)

// AggregationService shall represent an aggregation service for a Redfish
// implementation.
type AggregationService struct {
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether the aggregation service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// setDefaultBootOrderTarget is the URL to send SetDefaultBootOrder requests.
	setDefaultBootOrderTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a AggregationService object from the raw JSON.
func (a *AggregationService) UnmarshalJSON(b []byte) error {
	type temp AggregationService
	type aActions struct {
		Reset               ActionTarget `json:"#AggregationService.Reset"`
		SetDefaultBootOrder ActionTarget `json:"#AggregationService.SetDefaultBootOrder"`
	}
	var tmp struct {
		temp
		Actions            aActions
		Aggregates         Link `json:"Aggregates"`
		AggregationSources Link `json:"AggregationSources"`
		ConnectionMethods  Link `json:"ConnectionMethods"`
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
	a.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AggregationService) Update() error {
	readWriteFields := []string{
		"ServiceEnabled",
	}

	return a.UpdateFromRawData(a, a.RawData, readWriteFields)
}

// GetAggregationService will get a AggregationService instance from the service.
func GetAggregationService(c Client, uri string) (*AggregationService, error) {
	return GetObject[AggregationService](c, uri)
}

// ListReferencedAggregationServices gets the collection of AggregationService from
// a provided reference.
func ListReferencedAggregationServices(c Client, link string) ([]*AggregationService, error) {
	return GetCollectionObjects[AggregationService](c, link)
}

// AggregationServiceResetParameters holds the parameters for the Reset action.
type AggregationServiceResetParameters struct {
	// BatchSize shall contain the number of elements in each batch simultaneously
	// being issued a reset.
	BatchSize uint `json:"BatchSize,omitempty"`
	// DelayBetweenBatchesInSeconds shall contain the delay of the batches of
	// elements being reset in seconds.
	DelayBetweenBatchesInSeconds uint `json:"DelayBetweenBatchesInSeconds,omitempty"`
	// ResetType shall contain the type of reset. The service can accept a request
	// without the parameter and perform an implementation-specific default reset.
	ResetType ResetType `json:"ResetType,omitempty"`
	// TargetURIs shall contain an array of links to the resources being reset.
	TargetURIs []string `json:"TargetURIs,omitempty"`
}

// This action shall perform a reset of a set of resources.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AggregationService) Reset(params *AggregationServiceResetParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(a.client,
		a.resetTarget, params, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall restore the boot order to the default state for the
// specified computer systems.
// systems - This parameter shall contain an array of links to resources of
// type 'ComputerSystem'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AggregationService) SetDefaultBootOrder(systems []string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Systems"] = systems
	resp, taskInfo, err := PostWithTask(a.client,
		a.setDefaultBootOrderTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Aggregates gets the Aggregates collection.
func (a *AggregationService) Aggregates() ([]*Aggregate, error) {
	if a.aggregates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Aggregate](a.client, a.aggregates)
}

// AggregationSources gets the AggregationSources collection.
func (a *AggregationService) AggregationSources() ([]*AggregationSource, error) {
	if a.aggregationSources == "" {
		return nil, nil
	}
	return GetCollectionObjects[AggregationSource](a.client, a.aggregationSources)
}

// ConnectionMethods gets the ConnectionMethods collection.
func (a *AggregationService) ConnectionMethods() ([]*ConnectionMethod, error) {
	if a.connectionMethods == "" {
		return nil, nil
	}
	return GetCollectionObjects[ConnectionMethod](a.client, a.connectionMethods)
}
