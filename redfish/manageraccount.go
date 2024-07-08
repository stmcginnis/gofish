//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
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
	// AccountExpiration shall contain the date and time when this account expires. The service shall disable or delete
	// an account that has expired. This property shall not apply to accounts created by the Redfish Host Interface
	// Specification-defined credential bootstrapping. If the value is 'null', or the property is not present, the
	// account never expires.
	AccountExpiration string
	// AccountTypes shall contain an array of the various
	// account types that apply to the account. If this property is not
	// provided by the client, the default value shall be an array with the
	// single value `Redfish`.
	AccountTypes []AccountTypes
	// Certificates shall contain a link to a Resource
	// Collection of type CertificateCollection.
	certificates []string
	// Description provides a description of this resource.
	Description string
	// EmailAddress shall contain the email address associated with this account.
	EmailAddress string
	// Enabled shall indicate whether an account is enabled.
	// If `true`, the account is enabled and the user can log in. If
	// `false`, the account is disabled and, in the future, the user cannot
	// log in.
	Enabled bool
	// HostBootstrapAccount shall indicate whether this account is a bootstrap account created by the Redfish Host
	// Interface Specification-defined credential bootstrapping.
	HostBootstrapAccount bool
	// Keys shall contain a link to a resource collection of type KeyCollection that contains the keys that can be used
	// to authenticate this account.
	keys []string
	// Locked shall indicate whether the Account Service
	// automatically locked the account because the AccountLockoutThreshold
	// was exceeded. To manually unlock the account before the lockout
	// duration period, an administrator shall be able to change the property
	// to `false` to clear the lockout condition.
	Locked bool
	// MFABypass shall contain the multi-factor authentication bypass settings for this account.
	MFABypass MFABypass
	// OEMAccountTypes shall contain an array of the OEM account types for this account. This property shall be valid
	// when AccountTypes contains 'OEM'.
	OEMAccountTypes []string
	// OEM object used on this account.
	OEM json.RawMessage
	// OneTimePasscodeDeliveryAddress shall contain the contact address for receiving one-time passcode messages for
	// multi-factor authentication for this account when the Enabled property in the OneTimePasscode property in
	// AccountService resource contains 'true'. This is typically the contact email address associated with the
	// account, but may be a separate, relay email address for delivery via SMS or other contact method.
	OneTimePasscodeDeliveryAddress string
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
	// PhoneNumber shall contain the contact phone number associated with this account.
	PhoneNumber string
	// RoleID shall contain the RoleId of the Role Resource
	// configured for this account. The Service shall reject POST, PATCH, or
	// PUT operations that provide a RoleId that does not exist by returning
	// the HTTP 400 (Bad Request) status code.
	RoleID string `json:"RoleId"`
	// SNMP shall contain the SNMP settings for this account
	// when AccountTypes contains `SNMP`.
	SNMP SNMPUserInfo
	// StrictAccountTypes shall indicate if the service needs to use the value of AccountTypes and OEMAccountTypes
	// values exactly as specified. A 'true' value shall indicate the service needs to either accept the value without
	// changes or reject the request. A 'false' value shall indicate the service may add additional 'AccountTypes' and
	// 'OEMAccountTypes' values as needed to support limitations it has in separately controlling access to individual
	// services. If this property is not present, the value shall be assumed to be 'false'. An update of the service
	// can cause account types to be added to or removed from the AccountTypes and OEMAccountTypes properties,
	// regardless of the value of this property. After a service update, clients should inspect all accounts where the
	// value of this property is 'true' and perform maintenance as needed.
	StrictAccountTypes bool
	// UserName shall contain the user name for this account.
	UserName string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
	// role is a link the the user roles.
	role string

	changePasswordTarget string
}

// UnmarshalJSON unmarshals a ManagerAccount object from the raw JSON.
func (manageraccount *ManagerAccount) UnmarshalJSON(b []byte) error {
	type temp ManagerAccount
	type Actions struct {
		ChangePassword common.ActionTarget `json:"#ManagerAccount.ChangePassword"`
	}
	type AccountLinks struct {
		Role common.Link
	}
	var t struct {
		temp
		Actions      Actions
		Links        AccountLinks
		Certificates common.LinksCollection
		Keys         common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*manageraccount = ManagerAccount(t.temp)

	// Extract the links to other entities for later
	manageraccount.changePasswordTarget = t.Actions.ChangePassword.Target
	manageraccount.role = t.Links.Role.String()
	manageraccount.certificates = t.Certificates.ToStrings()
	manageraccount.keys = t.Keys.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	manageraccount.rawData = b

	return nil
}

// Role gets the role of this ManagerAccount.
func (manageraccount *ManagerAccount) Role() (*Role, error) {
	if manageraccount.role == "" {
		return nil, nil
	}
	return GetRole(manageraccount.GetClient(), manageraccount.role)
}

// Certificates gets the user identity certificates for this account.
func (manageraccount *ManagerAccount) Certificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](manageraccount.GetClient(), manageraccount.certificates)
}

// Keys gets the keys that can be used to authenticate this account.
func (manageraccount *ManagerAccount) Keys() ([]*Key, error) {
	return common.GetObjects[Key](manageraccount.GetClient(), manageraccount.keys)
}

// ChangePassword changes the account password while requiring password for the current session.
// `newPassword` is the new password.
// `sessionAccountPassword` is the current session's account. A user changing their own password shall
// provide their current password for this parameter. An administrator changing the password for a
// different user shall provide their own password for this parameter. If the request is performed
// with HTTP Basic authentication, this parameter shall contain the same password encoded in the
// `Authorization` header.
func (manageraccount *ManagerAccount) ChangePassword(newPassword, sessionAccountPassword string) error {
	if manageraccount.changePasswordTarget == "" {
		return errors.New("ChangePassword is not supported by this service") //nolint:error-strings
	}
	parameters := struct {
		NewPassword            string
		SessionAccountPassword string
	}{
		NewPassword:            newPassword,
		SessionAccountPassword: sessionAccountPassword,
	}
	return manageraccount.Post(manageraccount.changePasswordTarget, parameters)
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
		"AccountExpiration",
		"AccountTypes",
		"EmailAddress",
		"Enabled",
		"Locked",
		"OEMAccountTypes",
		"OneTimePasscodeDeliveryAddress",
		"Password",
		"PasswordChangeRequired",
		"PasswordExpiration",
		"PhoneNumber",
		"RoleId",
		"StrictAccountTypes",
		"UserName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manageraccount).Elem()

	return manageraccount.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManagerAccount will get a ManagerAccount instance from the service.
func GetManagerAccount(c common.Client, uri string) (*ManagerAccount, error) {
	return common.GetObject[ManagerAccount](c, uri)
}

// ListReferencedManagerAccounts gets the collection of ManagerAccount from
// a provided reference.
func ListReferencedManagerAccounts(c common.Client, link string) ([]*ManagerAccount, error) {
	return common.GetCollectionObjects[ManagerAccount](c, link)
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
