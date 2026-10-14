//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"fmt"
)

// CreateAccountParameters holds the properties to set when creating a new
// Redfish user account. Any property left at its zero value is omitted from
// the create request, so the service applies its own default.
type CreateAccountParameters struct {
	// UserName shall contain the username for this account.
	UserName string `json:"UserName"`
	// Password shall contain the password for this account, and must conform to
	// the configured password requirements.
	Password string `json:"Password"`
	// RoleID shall contain the 'RoleId' of the role resource configured for this
	// account, typically one of 'Administrator', 'Operator', or 'ReadOnly'.
	RoleID string `json:"RoleId,omitempty"`
	// Enabled shall indicate whether the account is enabled. If 'true', the user
	// can log in. If 'false', the account is disabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// AccountExpiration shall contain the date and time when this account
	// expires. The service shall disable or delete an account that has expired.
	//
	// Version added: v1.8.0
	AccountExpiration string `json:"AccountExpiration,omitempty"`
	// AccountTypes shall contain the manager services that the account is allowed
	// to access. If not provided, the service typically defaults to 'Redfish'.
	//
	// Version added: v1.4.0
	AccountTypes []AccountTypes `json:"AccountTypes,omitempty"`
	// OEMAccountTypes shall contain the OEM account types for this account. It is
	// only valid when AccountTypes contains 'OEM'.
	//
	// Version added: v1.4.0
	OEMAccountTypes []string `json:"OEMAccountTypes,omitempty"`
	// StrictAccountTypes shall indicate if the service needs to use the
	// AccountTypes and OEMAccountTypes values exactly as specified, rejecting the
	// request if it cannot. If not provided, the service assumes 'false'.
	//
	// Version added: v1.7.0
	StrictAccountTypes *bool `json:"StrictAccountTypes,omitempty"`
	// EmailAddress shall contain the email address associated with this account.
	//
	// Version added: v1.11.0
	EmailAddress string `json:"EmailAddress,omitempty"`
	// PhoneNumber shall contain the contact phone number associated with this
	// account.
	//
	// Version added: v1.11.0
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	// OneTimePasscodeDeliveryAddress shall contain the contact address for
	// receiving one-time passcode messages for multi-factor authentication.
	//
	// Version added: v1.11.0
	OneTimePasscodeDeliveryAddress string `json:"OneTimePasscodeDeliveryAddress,omitempty"`
	// PasswordChangeRequired shall indicate whether the password must be changed
	// before further access to the account is allowed. Setting this to 'true' can
	// force a password change before first access of the account.
	//
	// Version added: v1.3.0
	PasswordChangeRequired *bool `json:"PasswordChangeRequired,omitempty"`
	// PasswordExpiration shall contain the date and time when this account
	// password expires. If allowed by the service, this overrides the
	// 'PasswordExpirationDays' property of the account service.
	//
	// Version added: v1.6.0
	PasswordExpiration string `json:"PasswordExpiration,omitempty"`
	// MFABypass shall contain the multi-factor authentication bypass settings for
	// this account.
	//
	// Version added: v1.10.0
	MFABypass *MFABypass `json:"MFABypass,omitempty"`
	// SNMP shall contain the SNMP settings for this account when AccountTypes
	// contains 'SNMP'.
	//
	// Version added: v1.4.0
	SNMP *SNMPUserInfo `json:"SNMP,omitempty"`
	// OEM gives the opportunity to specify any OEM specific properties. It should
	// contain the vendor specific struct that goes inside the Oem object.
	OEM any `json:"Oem,omitempty"`
}

// CreateAccount creates a new Redfish user account.
//
// `userName` is the new username to use.
//
// `password` is the initial password, must conform to configured password requirements.
//
// `roleID` is the role to assign to the user, typically one of `Administrator`, `Operator`, or `ReadOnly`.
//
// Returns the created user account that can then be updated for things like setting `passwordChangeRequried`, etc.
//
// Use CreateAccountParams to set any of the other account properties, such as
// OEM extensions, at creation time.
func (accountservice *AccountService) CreateAccount(userName, password, roleID string) (*ManagerAccount, error) {
	enabled := true
	return accountservice.CreateAccountParams(&CreateAccountParameters{
		UserName: userName,
		Password: password,
		RoleID:   roleID,
		Enabled:  &enabled,
	})
}

// CreateAccountParams creates a new Redfish user account from the given
// parameters. Properties left unset in `params` are not sent to the service.
//
// Returns the created user account.
func (accountservice *AccountService) CreateAccountParams(params *CreateAccountParameters) (*ManagerAccount, error) {
	baseEntity := &accountservice.Entity

	// The proper ETag for creating an account is found at `/AccountService/Accounts`
	// If ETag matching is disabled, we don't care and don't need to waste a request,
	// but otherwise, we need to load /Accounts to get the ETag then issue the request against that Entity
	if !accountservice.IsEtagMatchDisabled() {
		accountsEntity, err := GetObject[Entity](accountservice.GetClient(), accountservice.accounts)
		if err != nil {
			return nil, fmt.Errorf("failed to get accounts entity: %w", err)
		}

		baseEntity = accountsEntity
	}

	resp, err := baseEntity.PostWithResponse(accountservice.accounts, params)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var result ManagerAccount
	err = json.NewDecoder(resp.Body).Decode(&result)
	result.SetClient(accountservice.GetClient())

	return &result, err
}
