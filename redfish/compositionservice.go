//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CompositionService is used to represent the Composition Service
// Properties for a Redfish implementation.
type CompositionService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowOverprovisioning shall be a boolean indicating whether this service
	// is allowed to overprovision a composition relative to the composition request.
	AllowOverprovisioning bool
	// AllowZoneAffinity shall be a boolean indicating whether a client is
	// allowed to request that given composition request is fulfilled by a
	// specified Resource Zone.
	AllowZoneAffinity bool
	// Description provides a description of this resource.
	Description string
	// resourceBlocks shall contain the link to a collection of type ResourceBlockCollection.
	resourceBlocks string
	// resourceZones shall contain the link to a collection of type ZoneCollection.
	resourceZones string
	// ServiceEnabled shall be a boolean indicating whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals CompositionService object from the raw JSON.
func (compositionservice *CompositionService) UnmarshalJSON(b []byte) error {
	type temp CompositionService
	var t struct {
		temp
		ResourceBlocks common.Link
		ResourceZones  common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*compositionservice = CompositionService(t.temp)
	compositionservice.resourceBlocks = string(t.ResourceBlocks)
	compositionservice.resourceZones = string(t.ResourceZones)

	// This is a read/write object, so we need to save the raw object data for later
	compositionservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (compositionservice *CompositionService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(CompositionService)
	err := original.UnmarshalJSON(compositionservice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AllowOverprovisioning",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(compositionservice).Elem()

	return compositionservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCompositionService will get a CompositionService instance from the service.
func GetCompositionService(c common.Client, uri string) (*CompositionService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var compositionservice CompositionService
	err = json.NewDecoder(resp.Body).Decode(&compositionservice)
	if err != nil {
		return nil, err
	}

	compositionservice.SetClient(c)
	return &compositionservice, nil
}

// ListReferencedCompositionServices gets the collection of CompositionService from
// a provided reference.
func ListReferencedCompositionServices(c common.Client, link string) ([]*CompositionService, error) {
	var result []*CompositionService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, compositionserviceLink := range links.ItemLinks {
		compositionservice, err := GetCompositionService(c, compositionserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, compositionservice)
	}

	return result, nil
}
