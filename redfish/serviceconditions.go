//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.4 - #ServiceConditions.v1_0_1.ServiceConditions

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ServiceConditions shall be used to represent the overall conditions present
// in a service for a Redfish implementation.
type ServiceConditions struct {
	common.Entity
	// Conditions shall represent a roll-up of the active conditions requiring
	// attention in resources of this Redfish service. The service may roll up any
	// number of conditions originating from resources in the service, using the
	// 'ConditionInRelatedResource' message from Base Message Registry. The array
	// order of conditions may change as new conditions occur or as conditions are
	// resolved by the service.
	Conditions []common.Condition
	// HealthRollup shall contain the highest severity of any messages included in
	// the 'Conditions' property.
	HealthRollup common.Health
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ServiceConditions object from the raw JSON.
func (s *ServiceConditions) UnmarshalJSON(b []byte) error {
	type temp ServiceConditions
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = ServiceConditions(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *ServiceConditions) Update() error {
	readWriteFields := []string{
		"Conditions",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
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
