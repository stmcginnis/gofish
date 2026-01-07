//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #PrivilegeRegistry.v1_2_0.PrivilegeRegistry

package schemas

import (
	"encoding/json"
)

// PrivilegeRegistry shall contain operation-to-privilege mappings.
type PrivilegeRegistry struct {
	Entity
	// Mappings shall describe the mappings between entities and the relevant
	// privileges that access those entities.
	Mappings []Mapping
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMPrivilegeDescriptions shall contain an array of the descriptions of the
	// values of the 'OEMPrivilegesUsed' property. The order of this array shall
	// match the order of the 'OEMPrivilegesUsed' property.
	//
	// Version added: v1.2.0
	OEMPrivilegeDescriptions []string
	// OEMPrivilegesUsed shall contain an array of OEM privileges used in this
	// mapping.
	OEMPrivilegesUsed []string
	// PrivilegesUsed shall contain an array of Redfish standard privileges used in
	// this mapping.
	PrivilegesUsed []PrivilegeType
}

// GetPrivilegeRegistry will get a PrivilegeRegistry instance from the service.
func GetPrivilegeRegistry(c Client, uri string) (*PrivilegeRegistry, error) {
	return GetObject[PrivilegeRegistry](c, uri)
}

// ListReferencedPrivilegeRegistrys gets the collection of PrivilegeRegistry from
// a provided reference.
func ListReferencedPrivilegeRegistrys(c Client, link string) ([]*PrivilegeRegistry, error) {
	return GetCollectionObjects[PrivilegeRegistry](c, link)
}

// Mapping shall describe a mapping between a resource type and the relevant
// privileges that accesses the resource.
type Mapping struct {
	// Entity shall contain the resource name, such as 'Manager'.
	Entity string
	// OperationMap shall list the mapping between HTTP methods and the privilege
	// required for the resource.
	OperationMap OperationMap
	// PropertyOverrides shall contain the privilege overrides of properties, such
	// as the 'Password' property in the 'ManagerAccount' resource.
	PropertyOverrides []TargetPrivilegeMap
	// ResourceURIOverrides shall contain the privilege overrides of resource URIs.
	// The target lists the resource URI and the new privileges.
	ResourceURIOverrides []TargetPrivilegeMap
	// SubordinateOverrides shall contain the privilege overrides of the
	// subordinate resource. The target lists are identified by resource type.
	SubordinateOverrides []TargetPrivilegeMap
}

// OperationMap shall describe the specific privileges required to complete a
// set of HTTP operations.
type OperationMap struct {
	// DELETE shall contain the privilege required to complete an HTTP 'DELETE'
	// operation.
	DELETE []OperationPrivilege
	// GET shall contain the privilege required to complete an HTTP 'GET'
	// operation.
	GET []OperationPrivilege
	// HEAD shall contain the privilege required to complete an HTTP 'HEAD'
	// operation.
	HEAD []OperationPrivilege
	// PATCH shall contain the privilege required to complete an HTTP 'PATCH'
	// operation.
	PATCH []OperationPrivilege
	// POST shall contain the privilege required to complete an HTTP 'POST'
	// operation.
	POST []OperationPrivilege
	// PUT shall contain the privilege required to complete an HTTP 'PUT'
	// operation.
	PUT []OperationPrivilege
}

// OperationPrivilege shall describe the privileges required to complete a
// specific HTTP operation.
type OperationPrivilege struct {
	// Privilege shall contain an array of privileges that are required to complete
	// a specific HTTP operation on a resource. This set of strings match zero or
	// more strings in the 'PrivilegesUsed' and 'OEMPrivilegesUsed' properties.
	Privilege []string
}

// TargetPrivilegeMap shall describe a mapping between one or more targets and
// the HTTP operations associated with them.
type TargetPrivilegeMap struct {
	// OperationMap shall contain the mapping between the HTTP operation and the
	// privilege required to complete the operation.
	OperationMap OperationMap
	// Targets shall contain the array of URIs, resource types, or properties. For
	// example, '/redfish/v1/Systems/1', 'Manager', or 'Password'. When the
	// 'Targets' property is not present, no override is specified.
	Targets []string
}
