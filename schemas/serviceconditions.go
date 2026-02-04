//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ServiceConditions.v1_0_1.json
// 2021.4 - #ServiceConditions.v1_0_1.ServiceConditions

package schemas

import (
	"encoding/json"
)

// ServiceConditions shall be used to represent the overall conditions present
// in a service for a Redfish implementation.
type ServiceConditions struct {
	Entity
	// Conditions shall represent a roll-up of the active conditions requiring
	// attention in resources of this Redfish service. The service may roll up any
	// number of conditions originating from resources in the service, using the
	// 'ConditionInRelatedResource' message from Base Message Registry. The array
	// order of conditions may change as new conditions occur or as conditions are
	// resolved by the service.
	Conditions []Condition
	// HealthRollup shall contain the highest severity of any messages included in
	// the 'Conditions' property.
	HealthRollup Health
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// GetServiceConditions will get a ServiceConditions instance from the service.
func GetServiceConditions(c Client, uri string) (*ServiceConditions, error) {
	return GetObject[ServiceConditions](c, uri)
}

// ListReferencedServiceConditionss gets the collection of ServiceConditions from
// a provided reference.
func ListReferencedServiceConditionss(c Client, link string) ([]*ServiceConditions, error) {
	return GetCollectionObjects[ServiceConditions](c, link)
}
