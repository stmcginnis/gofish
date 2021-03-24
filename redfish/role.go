//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// PrivilegeType is the role privilege type.
type PrivilegeType string

const (

	// LoginPrivilegeType Can log in to the service and read Resources.
	LoginPrivilegeType PrivilegeType = "Login"
	// ConfigureManagerPrivilegeType Can configure managers.
	ConfigureManagerPrivilegeType PrivilegeType = "ConfigureManager"
	// ConfigureUsersPrivilegeType Can configure users and their accounts.
	ConfigureUsersPrivilegeType PrivilegeType = "ConfigureUsers"
	// ConfigureSelfPrivilegeType Can change the password for the current
	// user account and log out of their own sessions.
	ConfigureSelfPrivilegeType PrivilegeType = "ConfigureSelf"
	// ConfigureComponentsPrivilegeType Can configure components that this
	// service manages.
	ConfigureComponentsPrivilegeType PrivilegeType = "ConfigureComponents"
	// NoAuthPrivilegeType shall be used to indicate an operation does not
	// require authentication.  This privilege shall not be used in Redfish
	// Roles.
	NoAuthPrivilegeType PrivilegeType = "NoAuth"
)

// Role represents the Redfish Role for the user account.
type Role struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AssignedPrivileges shall contain the Redfish
	// privileges for this Role. For predefined Roles, this property shall
	// be read-only. For custom Roles, some implementations may not allow
	// writing to this property.
	AssignedPrivileges []PrivilegeType
	// Description provides a description of this resource.
	Description string
	// IsPredefined shall indicate whether the Role is a
	// Redfish-predefined Role rather than a custom Redfish Role.
	IsPredefined bool
	// OemPrivileges shall contain the OEM privileges for
	// this Role. For predefined Roles, this property shall be read-only.
	// For custom Roles, some implementations may not allow writing to this
	// property.
	OemPrivileges []string
	// RoleID shall contain the string name of the Role.
	// This property shall contain the same value as the Id property.
	RoleID string `json:"RoleId"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Role object from the raw JSON.
func (role *Role) UnmarshalJSON(b []byte) error {
	type temp Role
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*role = Role(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	role.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (role *Role) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Role)
	err := original.UnmarshalJSON(role.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssignedPrivileges",
		"OemPrivileges",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(role).Elem()

	return role.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRole will get a Role instance from the service.
func GetRole(c common.Client, uri string) (*Role, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var role Role
	err = json.NewDecoder(resp.Body).Decode(&role)
	if err != nil {
		return nil, err
	}

	role.SetClient(c)
	return &role, nil
}

// ListReferencedRoles gets the collection of Role from
// a provided reference.
func ListReferencedRoles(c common.Client, link string) ([]*Role, error) {
	var result []*Role
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, roleLink := range links.ItemLinks {
		role, err := GetRole(c, roleLink)
		if err != nil {
			return result, err
		}
		result = append(result, role)
	}

	return result, nil
}
