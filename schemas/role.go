//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #Role.v1_3_3.Role

package schemas

import (
	"encoding/json"
)

// Role shall represent the Redfish role for the user account.
type Role struct {
	Entity
	// AlternateRoleID shall contain a non-restricted 'RoleId' intended to be used
	// in its place when the 'Restricted' property contains the value 'true'.
	//
	// Version added: v1.3.0
	AlternateRoleID string `json:"AlternateRoleId"`
	// AssignedPrivileges shall contain the Redfish privileges for this role. For
	// predefined roles, this property shall be read-only. For custom roles, some
	// implementations may prevent writing to this property.
	AssignedPrivileges []PrivilegeType
	// IsPredefined shall indicate whether the role is predefined by Redfish or an
	// OEM as contrasted with a client-defined role. If this property is not
	// present, the value should be assumed to be 'false'.
	IsPredefined bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemPrivileges shall contain the OEM privileges for this role. For predefined
	// roles, this property shall be read-only. For custom roles, some
	// implementations may prevent writing to this property.
	OemPrivileges []string
	// Restricted shall indicate whether use of the role is restricted by a service
	// as defined by the 'Restricted roles and restricted privileges' clause of the
	// Redfish Specification. If this property is not present, the value shall be
	// assumed to be 'false'.
	//
	// Version added: v1.3.0
	Restricted bool
	// RoleID shall contain the string name of the role. This property shall
	// contain the same value as the 'Id' property.
	//
	// Version added: v1.2.0
	RoleID string `json:"RoleId"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Role object from the raw JSON.
func (r *Role) UnmarshalJSON(b []byte) error {
	type temp Role
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = Role(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *Role) Update() error {
	readWriteFields := []string{
		"AssignedPrivileges",
		"OemPrivileges",
	}

	return r.UpdateFromRawData(r, r.RawData, readWriteFields)
}

// GetRole will get a Role instance from the service.
func GetRole(c Client, uri string) (*Role, error) {
	return GetObject[Role](c, uri)
}

// ListReferencedRoles gets the collection of Role from
// a provided reference.
func ListReferencedRoles(c Client, link string) ([]*Role, error) {
	return GetCollectionObjects[Role](c, link)
}
