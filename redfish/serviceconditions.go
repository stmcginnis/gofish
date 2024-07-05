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
	return common.GetObject[ServiceConditions](c, uri)
}

// ListReferencedServiceConditionss gets the collection of ServiceConditions from
// a provided reference.
func ListReferencedServiceConditionss(c common.Client, link string) ([]*ServiceConditions, error) {
	return common.GetCollectionObjects[ServiceConditions](c, link)
}
