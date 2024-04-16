//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ServiceConditions shall be used to represent the overall conditions present in a service for a Redfish
// implementation.
type ServiceConditions struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Conditions shall represent a roll-up of the active conditions requiring attention in resources of this Redfish
	// service. The service may roll up any number of conditions originating from resources in the service, using the
	// 'ConditionInRelatedResource' message from Base Message Registry.
	Conditions []Condition
	// Description provides a description of this resource.
	Description string
	// HealthRollup shall contain the highest severity of any messages included in the Conditions property.
	HealthRollup common.Health
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetServiceConditions will get a ServiceConditions instance from the service.
func GetServiceConditions(c common.Client, uri string) (*ServiceConditions, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var serviceconditions ServiceConditions
	err = json.NewDecoder(resp.Body).Decode(&serviceconditions)
	if err != nil {
		return nil, err
	}

	serviceconditions.SetClient(c)
	return &serviceconditions, nil
}

// ListReferencedServiceConditionss gets the collection of ServiceConditions from
// a provided reference.
func ListReferencedServiceConditionss(c common.Client, link string) ([]*ServiceConditions, error) {
	var result []*ServiceConditions
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *ServiceConditions
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		serviceconditions, err := GetServiceConditions(c, link)
		ch <- GetResult{Item: serviceconditions, Link: link, Error: err}
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
