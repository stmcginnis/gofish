//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Mapping shall describe a mapping between a Resource type and the relevant privileges that accesses the Resource.
type Mapping struct {
	// Entity shall contain the Resource name, such as 'Manager'.
	Entity string
	// OperationMap shall list the mapping between HTTP methods and the privilege required for the Resource.
	OperationMap OperationMap
	// PropertyOverrides shall contain the privilege overrides of properties, such as the 'Password' property in the
	// 'ManagerAccount' Resource.
	PropertyOverrides []TargetPrivilegeMap
	// ResourceURIOverrides shall contain the privilege overrides of Resource URIs. The target lists the Resource URI
	// and the new privileges.
	ResourceURIOverrides []TargetPrivilegeMap
	// SubordinateOverrides shall contain the privilege overrides of the subordinate Resource. The target lists are
	// identified by Resource type.
	SubordinateOverrides []TargetPrivilegeMap
}

// OperationMap shall describe the specific privileges required to complete a set of HTTP operations.
type OperationMap struct {
	// DELETE shall contain the privilege required to complete an HTTP DELETE operation.
	DELETE []OperationPrivilege
	// GET shall contain the privilege required to complete an HTTP GET operation.
	GET []OperationPrivilege
	// HEAD shall contain the privilege required to complete an HTTP HEAD operation.
	HEAD []OperationPrivilege
	// PATCH shall contain the privilege required to complete an HTTP PATCH operation.
	PATCH []OperationPrivilege
	// POST shall contain the privilege required to complete an HTTP POST operation.
	POST []OperationPrivilege
	// PUT shall contain the privilege required to complete an HTTP PUT operation.
	PUT []OperationPrivilege
}

// OperationPrivilege shall describe the privileges required to complete a specific HTTP operation.
type OperationPrivilege struct {
	// Privilege shall contain an array of privileges that are required to complete a specific HTTP operation on a
	// Resource. This set of strings match zero or more strings in the PrivilegesUsed and OEMPrivilegesUsed properties.
	Privilege []string
}

// PrivilegeRegistry This Resource contains operation-to-privilege mappings.
type PrivilegeRegistry struct {
	common.Entity
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Mappings shall describe the mappings between entities and the relevant privileges that access those entities.
	Mappings []Mapping
	// OEMPrivilegesUsed shall contain an array of OEM privileges used in this mapping.
	OEMPrivilegesUsed json.RawMessage
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PrivilegesUsed shall contain an array of Redfish standard privileges used in this mapping.
	PrivilegesUsed []PrivilegeType
}

// GetPrivilegeRegistry will get a PrivilegeRegistry instance from the service.
func GetPrivilegeRegistry(c common.Client, uri string) (*PrivilegeRegistry, error) {
	return common.GetObject[PrivilegeRegistry](c, uri)
}

// ListReferencedPrivilegeRegistrys gets the collection of PrivilegeRegistry from
// a provided reference.
func ListReferencedPrivilegeRegistrys(c common.Client, link string) ([]*PrivilegeRegistry, error) {
	return common.GetCollectionObjects[PrivilegeRegistry](c, link)
}

// TargetPrivilegeMap shall describe a mapping between one or more targets and the HTTP operations associated with
// them.
type TargetPrivilegeMap struct {
	// OperationMap shall contain the mapping between the HTTP operation and the privilege required to complete the
	// operation.
	OperationMap string
	// Targets shall contain the array of URIs, Resource types, or properties. For example, '/redfish/v1/Systems/1',
	// 'Manager', or 'Password'. When the Targets property is not present, no override is specified.
	Targets []string
}
