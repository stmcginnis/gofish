//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CompositionService is used to represent the Composition Service
// Properties for a Redfish implementation.
type CompositionService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
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
}

// UnmarshalJSON unmarshals CompositionService object from the raw JSON.
func (cs *CompositionService) UnmarshalJSON(b []byte) error {
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
	*cs = CompositionService(t.temp)
	cs.resourceBlocks = string(t.ResourceBlocks)
	cs.resourceZones = string(t.ResourceZones)

	return nil
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
