//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AccountTypes is the type of the account.
type AccountTypes string

const (

	// RedfishAccountTypes Allow access to the Redfish Service.
	RedfishAccountTypes AccountTypes = "Redfish"
	// SNMPAccountTypes Allow access to SNMP services.
	SNMPAccountTypes AccountTypes = "SNMP"
	// OEMAccountTypes OEM account type.
	OEMAccountTypes AccountTypes = "OEM"
)

// ManagerAccount shall represent Resources that represent the user
// accounts for the manager.
type ManagerAccount struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountTypes shall contain an array of the various
	// account types that apply to the account. If this property is not
	// provided by the client, the default value shall be an array with the
	// single value `Redfish`.
	AccountTypes []AccountTypes
	// Certificates shall contain a link to a Resource
	// Collection of type CertificateCollection.
	certificates string
	// Description provides a description of this resource.
	Description string
	// Enabled shall indicate whether an account is enabled.
	// If `true`, the account is enabled and the user can log in. If
	// `false`, the account is disabled and, in the future, the user cannot
	// log in.
	Enabled bool
	// Locked shall indicate whether the Account Service
	// automatically locked the account because the AccountLockoutThreshold
	// was exceeded. To manually unlock the account before the lockout
	// duration period, an administrator shall be able to change the property
	// to `false` to clear the lockout condition.
	Locked bool
	// Password shall contain the password for this account.
	// The value shall be `null` in responses.
	Password string
	// PasswordChangeRequired shall indicate whether the
	// service requires that the password for this account be changed before
	// further access to the account is allowed. The implementation may deny
	// access to the service if the password has not been changed. A manager
	// account created with an initial PasswordChangeRequired value of `true`
	// may force a password change before first access of the account. When
	// the Password property for this account is updated, the service shall
	// set this property to `false`.
	PasswordChangeRequired bool
	// PasswordExpiration shall contain the date and time
	// when this account password expires. If the value is `null`, the
	// account password never expires.
	PasswordExpiration string
	// RoleID shall contain the RoleId of the Role Resource
	// configured for this account. The Service shall reject POST, PATCH, or
	// PUT operations that provide a RoleId that does not exist by returning
	// the HTTP 400 (Bad Request) status code.
	RoleID string `json:"RoleId"`
	// SNMP shall contain the SNMP settings for this account
	// when AccountTypes contains `SNMP`.
	SNMP SNMPUserInfo
	// UserName shall contain the user name for this account.
	UserName string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
	// role is a link the the user roles.
	role string
}

// UnmarshalJSON unmarshals a ManagerAccount object from the raw JSON.
func (manageraccount *ManagerAccount) UnmarshalJSON(b []byte) error {
	type temp ManagerAccount
	type AccountLinks struct {
		Role common.Link
	}
	var t struct {
		temp
		Links        AccountLinks
		Certificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*manageraccount = ManagerAccount(t.temp)

	// Extract the links to other entities for later
	manageraccount.role = string(t.Links.Role)
	manageraccount.certificates = string(t.Certificates)

	// This is a read/write object, so we need to save the raw object data for later
	manageraccount.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manageraccount *ManagerAccount) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ManagerAccount)
	err := original.UnmarshalJSON(manageraccount.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccountTypes",
		"Enabled",
		"Locked",
		// "OEMAccountTypes",
		"Password",
		"PasswordChangeRequired",
		"PasswordExpiration",
		"RoleId",
		"UserName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manageraccount).Elem()

	return manageraccount.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManagerAccount will get a ManagerAccount instance from the service.
func GetManagerAccount(c common.Client, uri string) (*ManagerAccount, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manageraccount ManagerAccount
	err = json.NewDecoder(resp.Body).Decode(&manageraccount)
	if err != nil {
		return nil, err
	}

	manageraccount.SetClient(c)
	return &manageraccount, nil
}

// ListReferencedManagerAccounts gets the collection of ManagerAccount from
// a provided reference.
func ListReferencedManagerAccounts(c common.Client, link string) ([]*ManagerAccount, error) {
	var result []*ManagerAccount
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, manageraccountLink := range links.ItemLinks {
		manageraccount, err := GetManagerAccount(c, manageraccountLink)
		if err != nil {
			return result, err
		}
		result = append(result, manageraccount)
	}

	return result, nil
}

// SNMPUserInfo is shall contain the SNMP settings for an account.
type SNMPUserInfo struct {

	// AuthenticationKey shall contain the key for SNMPv3
	// authentication. The value shall be `null` in responses. This
	// property accepts a passphrase or a hex-encoded key. If the string
	// starts with `Passphrase:`, the remainder of the string shall be the
	// passphrase and shall be converted to the key as described in the
	// 'Password to Key Algorithm' section of RFC3414. If the string starts
	// with `Hex:`, then the remainder of the string shall be the key encoded
	// in hexadecimal notation. If the string starts with neither, the full
	// string shall be a passphrase and shall be converted to the key as
	// described in the 'Password to Key Algorithm' section of RFC3414. The
	// passphrase may contain any printable characters except for the double
	// quotation mark.
	AuthenticationKey string
	// AuthenticationKeySet shall contain `true` if a valid
	// value was provided for the AuthenticationKey property. Otherwise, the
	// property shall contain `false`.
	AuthenticationKeySet bool
	// AuthenticationProtocol shall contain the SNMPv3
	// authentication protocol.
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey shall contain the key for SNMPv3
	// encryption. The value shall be `null` in responses. This property
	// accepts a passphrase or a hex-encoded key. If the string starts with
	// `Passphrase:`, the remainder of the string shall be the passphrase and
	// shall be converted to the key as described in the 'Password to Key
	// Algorithm' section of RFC3414. If the string starts with `Hex:`, then
	// the remainder of the string shall be the key encoded in hexadecimal
	// notation. If the string starts with neither, the full string shall be
	// a passphrase and shall be converted to the key as described in the
	// 'Password to Key Algorithm' section of RFC3414. The passphrase may
	// contain any printable characters except for the double quotation mark.
	EncryptionKey string
	// EncryptionKeySet shall contain `true` if a valid
	// value was provided for the EncryptionKey property. Otherwise, the
	// property shall contain `false`.
	EncryptionKeySet bool
	// EncryptionProtocol shall contain the SNMPv3
	// encryption protocol.
	EncryptionProtocol SNMPEncryptionProtocols
}
