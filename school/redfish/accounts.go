// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// AccountService contains properties for managing user accounts. The
// properties are common to all user accounts, such as password requirements,
// and control features such as account lockout. The schema also contains links
// to the collections of Manager Accounts and Roles.
type AccountService struct {
	common.Entity
	Description                 string
	Modified                    string
	AuthFailureLoggingThreshold int
	MinPasswordLength           int
	accounts                    string
	roles                       string
}

// UnmarshalJSON unmarshals an AccountService object from the raw JSON.
func (as *AccountService) UnmarshalJSON(b []byte) error {
	type temp AccountService
	type AccountLinks struct {
		Accounts common.Link
		Roles    common.Link
	}
	var t struct {
		temp
		Links AccountLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*as = AccountService(t.temp)

	// Extract the links to other entities for later
	as.accounts = string(t.Links.Accounts)
	as.roles = string(t.Links.Roles)

	return nil
}

// GetAccountService will get the AccountService instance from the Redfish
// service.
func GetAccountService(c common.Client, uri string) (*AccountService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t AccountService
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	t.SetClient(c)
	return &t, nil
}

// Accounts get the accounts from the account service
func (as *AccountService) Accounts() ([]*Account, error) {
	return ListReferencedAccounts(as.Client, as.accounts)
}

// Roles gets the roles from the account service
func (as *AccountService) Roles() ([]*Role, error) {
	return ListReferencedRoles(as.Client, as.roles)
}

// Account is a Redfish account
type Account struct {
	common.Entity
	Modified    string
	Description string
	Password    string
	UserName    string
	Locked      bool
	Enabled     bool
	RoleID      string `json:"RoleId"`
	role        string
}

// UnmarshalJSON unmarshals an Account object from the raw JSON.
func (s *Account) UnmarshalJSON(b []byte) error {
	type temp Account
	type AccountLinks struct {
		Role common.Link
	}
	var t struct {
		temp
		Links AccountLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = Account(t.temp)

	// Extract the links to other entities for later
	s.role = string(t.Links.Role)

	return nil
}

// GetAccount will get an account instance from the Redfish service.
func GetAccount(c common.Client, uri string) (*Account, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t Account
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedAccounts gets the collection of Accounts
func ListReferencedAccounts(c common.Client, link string) ([]*Account, error) {
	var result []*Account
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, aLink := range links.ItemLinks {
		a, err := GetAccount(c, aLink)
		if err != nil {
			return result, err
		}
		result = append(result, a)
	}

	return result, nil
}

// Role is a Redfish role
type Role struct {
	common.Entity
	Modified           string
	Description        string
	IsPredefined       bool
	AssignedPrivileges []string
	OEMPrivileges      []string
}

// GetRole will get a role instance from the Redfish service.
func GetRole(c common.Client, uri string) (*Role, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t Role
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedRoles gets the collection of Roles
func ListReferencedRoles(c common.Client, link string) ([]*Role, error) {
	var result []*Role
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, aLink := range links.ItemLinks {
		a, err := GetRole(c, aLink)
		if err != nil {
			return result, err
		}
		result = append(result, a)
	}

	return result, nil
}
