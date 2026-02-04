//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ManagerAccount.v1_14_1.json
// 2025.2 - #ManagerAccount.v1_14_1.ManagerAccount

package schemas

import (
	"encoding/json"
)

type AccountTypes string

const (
	// RedfishAccountTypes shall indicate the account is allowed to access Redfish
	// services. If the version of the 'ManagerAccount' resource is lower than the
	// schema version when another enumeration value in this list was added, the
	// implementation may include that functionality as part of the 'Redfish'
	// value.
	RedfishAccountTypes AccountTypes = "Redfish"
	// SNMPAccountTypes shall indicate the account is allowed to access SNMP
	// services.
	SNMPAccountTypes AccountTypes = "SNMP"
	// OEMAccountTypes shall indicate the account is allowed to access the services
	// listed in the 'OEMAccountTypes' property.
	OEMAccountTypes AccountTypes = "OEM"
	// HostConsoleAccountTypes shall indicate the account is allowed to access the
	// host console.
	HostConsoleAccountTypes AccountTypes = "HostConsole"
	// ManagerConsoleAccountTypes shall indicate the account is allowed to access
	// the manager console.
	ManagerConsoleAccountTypes AccountTypes = "ManagerConsole"
	// IPMIAccountTypes shall indicate the account is allowed to access the
	// Intelligent Platform Management Interface service.
	IPMIAccountTypes AccountTypes = "IPMI"
	// KVMIPAccountTypes shall indicate the account is allowed to access the
	// Keyboard-Video-Mouse over IP session service.
	KVMIPAccountTypes AccountTypes = "KVMIP"
	// VirtualMediaAccountTypes shall indicate the account is allowed to control
	// virtual media.
	VirtualMediaAccountTypes AccountTypes = "VirtualMedia"
	// WebUIAccountTypes shall indicate the account is allowed to access the web
	// interface.
	WebUIAccountTypes AccountTypes = "WebUI"
	// ControlPanelAccountTypes shall indicate the account is used to allow
	// PIN-based access via an external control panel. If this value is specified,
	// the 'AccountTypes' property should not contain other values for the same
	// 'ManagerAccount' resource. If this value is specified, 'Password' shall
	// contain the PIN to enable access, and may not follow other password-related
	// rules. The 'ChangePassword' action, if supported, may be used to update the
	// PIN.
	ControlPanelAccountTypes AccountTypes = "ControlPanel"
)

// ManagerAccount shall represent a user account for the manager in a Redfish
// implementation. The account shall indicate the allowed access to one of more
// services in the manager.
type ManagerAccount struct {
	Entity
	// AccountExpiration shall contain the date and time when this account expires.
	// The service shall disable or delete an account that has expired. This
	// property shall not apply to accounts created by the Redfish Host Interface
	// Specification-defined credential bootstrapping. If the value is 'null', or
	// the property is not present, the account never expires.
	//
	// Version added: v1.8.0
	AccountExpiration string
	// AccountTypes shall contain an array of the various manager services that the
	// account is allowed to access. This shall not include functionality for
	// receiving events or other notifications. If this property is not provided by
	// the client, the default value shall be an array that contains the value
	// 'Redfish'. The service may add additional values when this property is set
	// or updated if allowed by the value of the 'StrictAccountTypes' property.
	//
	// Version added: v1.4.0
	AccountTypes []AccountTypes
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the user identity certificates for
	// this account.
	//
	// Version added: v1.2.0
	certificates string
	// EmailAddress shall contain the email address associated with this account.
	//
	// Version added: v1.11.0
	EmailAddress string
	// Enabled shall indicate whether an account is enabled. If 'true', the account
	// is enabled and the user can log in. If 'false', the account is disabled and,
	// in the future, the user cannot log in.
	Enabled bool
	// HostBootstrapAccount shall indicate whether this account is a bootstrap
	// account created by the Redfish Host Interface Specification-defined
	// credential bootstrapping.
	//
	// Version added: v1.8.0
	HostBootstrapAccount bool
	// Keys shall contain a link to a resource collection of type 'KeyCollection'
	// that contains the keys that can be used to authenticate this account.
	//
	// Version added: v1.9.0
	keys string
	// Locked shall indicate whether the account service automatically locked the
	// account because the 'AccountLockoutThreshold' was exceeded. To manually
	// unlock the account before the lockout duration period, an administrator
	// shall be able to change the property to 'false' to clear the lockout
	// condition.
	Locked bool
	// MFABypass shall contain the multi-factor authentication bypass settings for
	// this account.
	//
	// Version added: v1.10.0
	MFABypass MFABypass
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMAccountTypes shall contain an array of the OEM account types for this
	// account. This property shall be valid when 'AccountTypes' contains 'OEM'.
	//
	// Version added: v1.4.0
	OEMAccountTypes []string
	// OneTimePasscodeDeliveryAddress shall contain the contact address for
	// receiving one-time passcode messages for multi-factor authentication for
	// this account when the 'Enabled' property in the 'OneTimePasscode' property
	// in 'AccountService' resource contains 'true'. This is typically the contact
	// email address associated with the account, but may be a separate, relay
	// email address for delivery via SMS or other contact method.
	//
	// Version added: v1.11.0
	OneTimePasscodeDeliveryAddress string
	// Password shall contain the password for this account. The value shall be
	// 'null' in responses.
	Password string
	// PasswordChangeRequired shall indicate whether the service requires that the
	// password for this account be changed before further access to the account is
	// allowed. The implementation may deny access to the service if the password
	// has not been changed. A manager account created with an initial
	// 'PasswordChangeRequired' value of 'true' may force a password change before
	// first access of the account. When the 'Password' property for this account
	// is updated, the service shall set this property to 'false'.
	//
	// Version added: v1.3.0
	PasswordChangeRequired bool
	// PasswordExpiration shall contain the date and time when this account
	// password expires. If the value is 'null', the account password never
	// expires. If provided during account creation or password modification, and
	// allowed by the service, this value shall override the value of the
	// 'PasswordExpirationDays' property in the 'AccountService' resource.
	//
	// Version added: v1.6.0
	PasswordExpiration string
	// PhoneNumber shall contain the contact phone number associated with this
	// account.
	//
	// Version added: v1.11.0
	PhoneNumber string
	// RoleID shall contain the 'RoleId' of the role resource configured for this
	// account. The service shall reject 'POST', 'PATCH', or 'PUT' operations that
	// provide a 'RoleId' that does not exist by returning the HTTP '400 Bad
	// Request' status code.
	RoleID string `json:"RoleId"`
	// SNMP shall contain the SNMP settings for this account when 'AccountTypes'
	// contains 'SNMP'.
	//
	// Version added: v1.4.0
	SNMP SNMPUserInfo
	// SecretKeySet shall indicate if the secret key for RFC6238-defined Time-based
	// One-Time Password (TOTP) multi-factor authentication is set.
	//
	// Version added: v1.13.0
	SecretKeySet bool
	// StrictAccountTypes shall indicate if the service needs to use the value of
	// 'AccountTypes' and 'OEMAccountTypes' values exactly as specified. A 'true'
	// value shall indicate the service needs to either accept the value without
	// changes or reject the request. A 'false' value shall indicate the service
	// may add additional 'AccountTypes' and 'OEMAccountTypes' values as needed to
	// support limitations it has in separately controlling access to individual
	// services. If this property is not present, the value shall be assumed to be
	// 'false'. An update of the service can cause account types to be added to or
	// removed from the 'AccountTypes' and 'OEMAccountTypes' properties, regardless
	// of the value of this property. After a service update, clients should
	// inspect all accounts where the value of this property is 'true' and perform
	// maintenance as needed.
	//
	// Version added: v1.7.0
	StrictAccountTypes bool
	// UserName shall contain the username for this account.
	UserName string
	// changePasswordTarget is the URL to send ChangePassword requests.
	changePasswordTarget string
	// clearSecretKeyTarget is the URL to send ClearSecretKey requests.
	clearSecretKeyTarget string
	// generateSecretKeyTarget is the URL to send GenerateSecretKey requests.
	generateSecretKeyTarget string
	// verifyTimeBasedOneTimePasswordTarget is the URL to send VerifyTimeBasedOneTimePassword requests.
	verifyTimeBasedOneTimePasswordTarget string
	// role is the URI for Role.
	role string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ManagerAccount object from the raw JSON.
func (m *ManagerAccount) UnmarshalJSON(b []byte) error {
	type temp ManagerAccount
	type mActions struct {
		ChangePassword                 ActionTarget `json:"#ManagerAccount.ChangePassword"`
		ClearSecretKey                 ActionTarget `json:"#ManagerAccount.ClearSecretKey"`
		GenerateSecretKey              ActionTarget `json:"#ManagerAccount.GenerateSecretKey"`
		VerifyTimeBasedOneTimePassword ActionTarget `json:"#ManagerAccount.VerifyTimeBasedOneTimePassword"`
	}
	type mLinks struct {
		Role Link `json:"Role"`
	}
	var tmp struct {
		temp
		Actions      mActions
		Links        mLinks
		Certificates Link `json:"Certificates"`
		Keys         Link `json:"Keys"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ManagerAccount(tmp.temp)

	// Extract the links to other entities for later
	m.changePasswordTarget = tmp.Actions.ChangePassword.Target
	m.clearSecretKeyTarget = tmp.Actions.ClearSecretKey.Target
	m.generateSecretKeyTarget = tmp.Actions.GenerateSecretKey.Target
	m.verifyTimeBasedOneTimePasswordTarget = tmp.Actions.VerifyTimeBasedOneTimePassword.Target
	m.role = tmp.Links.Role.String()
	m.certificates = tmp.Certificates.String()
	m.keys = tmp.Keys.String()

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *ManagerAccount) Update() error {
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

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetManagerAccount will get a ManagerAccount instance from the service.
func GetManagerAccount(c Client, uri string) (*ManagerAccount, error) {
	return GetObject[ManagerAccount](c, uri)
}

// ListReferencedManagerAccounts gets the collection of ManagerAccount from
// a provided reference.
func ListReferencedManagerAccounts(c Client, link string) ([]*ManagerAccount, error) {
	return GetCollectionObjects[ManagerAccount](c, link)
}

// This action shall change the account password while requiring password for
// the current session. This action prevents session hijacking.
// newPassword - This parameter shall contain the new password.
// sessionAccountPassword - This parameter shall contain the password of the
// current session's account. A user changing their own password shall provide
// their current password for this parameter. An administrator changing the
// password for a different user shall provide their own password for this
// parameter. If the request is performed with HTTP Basic authentication, this
// parameter shall contain the same password encoded in the 'Authorization'
// header.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *ManagerAccount) ChangePassword(newPassword string, sessionAccountPassword string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["NewPassword"] = newPassword
	payload["SessionAccountPassword"] = sessionAccountPassword
	resp, taskInfo, err := PostWithTask(m.client,
		m.changePasswordTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall clear the secret key for RFC6238-defined Time-based
// One-Time Password (TOTP) multi-factor authentication for this account.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *ManagerAccount) ClearSecretKey() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(m.client,
		m.clearSecretKeyTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall randomly generate a new secret key for RFC6238-defined
// Time-based One-Time Password (TOTP) multi-factor authentication for this
// account.
func (m *ManagerAccount) GenerateSecretKey() (*GenerateSecretKeyResponse, error) {
	payload := make(map[string]any)

	resp, err := m.PostWithResponse(m.generateSecretKeyTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result GenerateSecretKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall verify a user-provided RFC6238-defined Time-based One-Time
// Password (TOTP).
// timeBasedOneTimePassword - This parameter shall contain the Time-based
// One-Time Password (TOTP) to verify. If the Time-based One-Time Password
// (TOTP) is not valid, the service shall return the HTTP '400 Bad Request'
// status code.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *ManagerAccount) VerifyTimeBasedOneTimePassword(timeBasedOneTimePassword string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["TimeBasedOneTimePassword"] = timeBasedOneTimePassword
	resp, taskInfo, err := PostWithTask(m.client,
		m.verifyTimeBasedOneTimePasswordTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Role gets the Role linked resource.
func (m *ManagerAccount) Role() (*Role, error) {
	if m.role == "" {
		return nil, nil
	}
	return GetObject[Role](m.client, m.role)
}

// Certificates gets the Certificates collection.
func (m *ManagerAccount) Certificates() ([]*Certificate, error) {
	if m.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](m.client, m.certificates)
}

// Keys gets the Keys collection.
func (m *ManagerAccount) Keys() ([]*Key, error) {
	if m.keys == "" {
		return nil, nil
	}
	return GetCollectionObjects[Key](m.client, m.keys)
}

// GenerateSecretKeyResponse shall contain the properties found in the response
// body for the 'GenerateSecretKey' action.
type GenerateSecretKeyResponse struct {
	// SecretKey shall contain secret key generated for RFC6238-defined Time-based
	// One-Time Password (TOTP) multi-factor authentication. Clients shall retain
	// the value of this property to generate tokens for future session creation
	// requests.
	//
	// Version added: v1.13.0
	SecretKey string
}

// SNMPUserInfo shall contain the SNMP settings for an account.
type SNMPUserInfo struct {
	// AuthenticationKey shall contain the key for SNMPv3 authentication. The value
	// shall be 'null' in responses. This property accepts a passphrase or a
	// hex-encoded key. If the string starts with 'Passphrase:', the remainder of
	// the string shall be the passphrase and shall be converted to the key as
	// described in the 'Password to Key Algorithm' section of RFC3414. If the
	// string starts with 'Hex:', then the remainder of the string shall be the key
	// encoded in hexadecimal notation. If the string starts with neither, the full
	// string shall be a passphrase and shall be converted to the key as described
	// in the 'Password to Key Algorithm' section of RFC3414.
	//
	// Version added: v1.4.0
	AuthenticationKey string
	// AuthenticationKeySet shall contain 'true' if a valid value was provided for
	// the 'AuthenticationKey' property. Otherwise, the property shall contain
	// 'false'.
	//
	// Version added: v1.5.0
	AuthenticationKeySet bool
	// AuthenticationProtocol shall contain the SNMPv3 authentication protocol.
	//
	// Version added: v1.4.0
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey shall contain the key for SNMPv3 encryption. The value shall
	// be 'null' in responses. This property accepts a passphrase or a hex-encoded
	// key. If the string starts with 'Passphrase:', the remainder of the string
	// shall be the passphrase and shall be converted to the key as described in
	// the 'Password to Key Algorithm' section of RFC3414. If the string starts
	// with 'Hex:', then the remainder of the string shall be the key encoded in
	// hexadecimal notation. If the string starts with neither, the full string
	// shall be a passphrase and shall be converted to the key as described in the
	// 'Password to Key Algorithm' section of RFC3414.
	//
	// Version added: v1.4.0
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the
	// 'EncryptionKey' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.5.0
	EncryptionKeySet bool
	// EncryptionProtocol shall contain the SNMPv3 encryption protocol.
	//
	// Version added: v1.4.0
	EncryptionProtocol SNMPEncryptionProtocols
}
