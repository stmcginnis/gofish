//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PolicyService.v1_0_0.json
// 2026.2 - #PolicyService.v1_0_0.PolicyService

package schemas

import (
	"encoding/json"
)

type OperatingMode string

const (
	// DisabledOperatingMode shall indicate that the policy service is disabled and
	// shall take no automatic action. The underlying service implementation may
	// still perform some actions as determined by the manufacturer, such as
	// throttling on thermal events, in order to preserve the system.
	DisabledOperatingMode OperatingMode = "Disabled"
	// AlertOnlyOperatingMode shall indicate that the policy service is operating
	// in alert only mode and shall only perform responses of type 'Event' or
	// 'Recheck'. The underlying service implementation may still perform some
	// actions as determined by the manufacturer, such as throttling on thermal
	// events, in order to preserve the system.
	AlertOnlyOperatingMode OperatingMode = "AlertOnly"
	// EnabledOperatingMode shall indicate that the policy service is operating in
	// autonomous mode where it performs all responses.
	EnabledOperatingMode OperatingMode = "Enabled"
)

// PolicyService This resource contains the policy service properties for a
// Redfish implementation.
type PolicyService struct {
	Entity
	// ActionsSupported shall contain a list of all action types that this policy
	// service can perform.
	ActionsSupported []ResponseAction
	// ConditionTypesSupported shall contain a list of all condition types that
	// this policy service can check.
	ConditionTypesSupported []ConditionType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingMode shall contain the mode in which this service is operating.
	OperatingMode OperatingMode
	// Policies shall contain a link to a resource collection of type
	// 'PolicyCollection'.
	policies string
	// ResourceTypesSupported shall contain a list of all resource types that this
	// policy service can reference in a response. Values shall contain the schema
	// name and shall not contain the complete '@odata.type' value with version
	// information.
	ResourceTypesSupported []string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PolicyService object from the raw JSON.
func (p *PolicyService) UnmarshalJSON(b []byte) error {
	type temp PolicyService
	var tmp struct {
		temp
		Policies Link `json:"Policies"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PolicyService(tmp.temp)

	// Extract the links to other entities for later
	p.policies = tmp.Policies.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PolicyService) Update() error {
	readWriteFields := []string{
		"OperatingMode",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPolicyService will get a PolicyService instance from the service.
func GetPolicyService(c Client, uri string) (*PolicyService, error) {
	return GetObject[PolicyService](c, uri)
}

// ListReferencedPolicyServices gets the collection of PolicyService from
// a provided reference.
func ListReferencedPolicyServices(c Client, link string) ([]*PolicyService, error) {
	return GetCollectionObjects[PolicyService](c, link)
}

// Policies gets the Policies collection.
func (p *PolicyService) Policies() ([]*Policy, error) {
	if p.policies == "" {
		return nil, nil
	}
	return GetCollectionObjects[Policy](p.client, p.policies)
}
