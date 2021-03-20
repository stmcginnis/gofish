//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AccountProviderTypes is
type AccountProviderTypes string

const (

	// RedfishServiceAccountProviderTypes shall be a DMTF Redfish
	// Specification-comformant service.  The ServiceAddresses format shall
	// contain a set of URIs that correspond to a Redfish Account Service.
	RedfishServiceAccountProviderTypes AccountProviderTypes = "RedfishService"
	// ActiveDirectoryServiceAccountProviderTypes shall be a Microsoft Active
	// Directory Technical Specification-comformant service.  The
	// ServiceAddresses format shall contain a set of fully qualified domain
	// names (FQDN) or NetBIOS names that links to the set of domain servers
	// for the Active Directory Service.
	ActiveDirectoryServiceAccountProviderTypes AccountProviderTypes = "ActiveDirectoryService"
	// LDAPServiceAccountProviderTypes shall be an RFC4511-conformant
	// service.  The ServiceAddresses format shall contain a set of fully
	// qualified domain names (FQDN) that links to the set of LDAP servers
	// for the Service.
	LDAPServiceAccountProviderTypes AccountProviderTypes = "LDAPService"
	// OEMAccountProviderTypes An OEM-specific external authentication or
	// directory service.
	OEMAccountProviderTypes AccountProviderTypes = "OEM"
)

// AuthenticationTypes is
type AuthenticationTypes string

const (

	// TokenAuthenticationTypes An opaque authentication token.
	TokenAuthenticationTypes AuthenticationTypes = "Token"
	// KerberosKeytabAuthenticationTypes A Kerberos keytab.
	KerberosKeytabAuthenticationTypes AuthenticationTypes = "KerberosKeytab"
	// UsernameAndPasswordAuthenticationTypes A user name and password
	// combination.
	UsernameAndPasswordAuthenticationTypes AuthenticationTypes = "UsernameAndPassword"
	// OEMAuthenticationTypes An OEM-specific authentication mechanism.
	OEMAuthenticationTypes AuthenticationTypes = "OEM"
)

// LocalAccountAuth is
type LocalAccountAuth string

const (

	// EnabledLocalAccountAuth shall authenticate users based on the Account
	// Service-defined Accounts Resource Collection.
	EnabledLocalAccountAuth LocalAccountAuth = "Enabled"
	// DisabledLocalAccountAuth shall never authenticate users based on the
	// Account Service-defined Accounts Resource Collection.
	DisabledLocalAccountAuth LocalAccountAuth = "Disabled"
	// FallbackLocalAccountAuth shall authenticate users based on the Account
	// Service-defined Accounts Resource Collection only if any external
	// account providers are currently unreachable.
	FallbackLocalAccountAuth LocalAccountAuth = "Fallback"
	// LocalFirstLocalAccountAuth shall first authenticate users based on the
	// Account Service-defined Accounts Resource Collection.  If
	// authentication fails, the Service shall authenticate by using external
	// account providers.
	LocalFirstLocalAccountAuth LocalAccountAuth = "LocalFirst"
)

// RoleMapping shall contain mapping rules that are used to convert
// the external account providers account information to the local
// Redfish Role.
type RoleMapping struct {

	// LocalRole shall contain the RoleId property value
	// within a Role Resource on this Redfish Service to which to map the
	// remote user or group.
	LocalRole string
	// RemoteGroup shall contain the name of the remote
	// group, or the remote role in the case of a Redfish Service, that maps
	// to the local Redfish Role to which this entity links.
	RemoteGroup string
	// RemoteUser shall contain the name of the remote user
	// that maps to the local Redfish Role to which this entity links.
	RemoteUser string
}

// Authentication is shall contain the information required to
// authenticate to the external service.
type Authentication struct {

	// AuthenticationType is used to connect to the external account
	// provider.
	AuthenticationType AuthenticationTypes
	// KerberosKeytab shall contain a Base64-encoded version
	// of the Kerberos keytab for this Service.  A PATCH or PUT operation
	// writes the keytab.  The value shall be `null` in responses.
	KerberosKeytab string
	// Password shall contain the password for this Service.
	// A PATCH or PUT operation writes the password.  The value shall be
	// `null` in responses.
	Password string
	// Token shall contain the token for this Service.  A
	// PATCH or PUT operation writes the token.  The value shall be `null` in
	// responses.
	Token string
	// Username shall contain the user name for this
	// Service.
	Username string
}

// ExternalAccountProvider shall contain properties that represent
// external account provider services that can provide accounts for this
// manager to use for authentication.
type ExternalAccountProvider struct {

	// Authentication shall contain the authentication
	// information for the external account provider.
	Authentication Authentication
	// certificates shall contain a link to a Resource
	// Collection of certificates of the CertificateCollection type that the
	// external account provider uses.
	// certificates string
	// ldapService shall contain any additional mapping
	// information needed to parse a generic LDAP service.  This property
	// should only be present inside the LDAP property.
	// ldapService string
	// PasswordSet shall contain `true` if a valid value was
	// provided for the Password property.  Otherwise, the property shall
	// contain `false`.
	PasswordSet bool
	// RemoteRoleMapping is used to convert the external account providers
	// account information to the local Redfish Role.
	RemoteRoleMapping []RoleMapping
	// ServiceAddresses shall contain the addresses of the
	// account providers to which this external account provider links.  The
	// format of this field depends on the type of external account provider.
	// Each item in the array shall contain a single address.  Services may
	// define their own behavior for managing multiple addresses.
	ServiceAddresses []string
	// ServiceEnabled shall indicate whether this service is
	// enabled.
	ServiceEnabled bool
}

// AccountService contains properties for managing user accounts. The
// properties are common to all user accounts, such as password requirements,
// and control features such as account lockout. The schema also contains links
// to the collections of Manager Accounts and Roles.
type AccountService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountLockoutCounterResetAfter shall contain the
	// period of time, in seconds, from the last failed login attempt when
	// the AccountLockoutThreshold counter, which counts the number of failed
	// login attempts, is reset to `0`.  Then, AccountLockoutThreshold
	// failures are required before the account is locked.  This value shall
	// be less than or equal to the AccountLockoutDuration value.  The
	// threshold counter also resets to `0` after each successful login.  If
	// the AccountLockoutCounterResetEnabled value is `false`, this property
	// shall be ignored.
	AccountLockoutCounterResetAfter int
	// AccountLockoutCounterResetEnabled shall indicate
	// whether the threshold counter is reset after the
	// AccountLockoutCounterResetAfter expires.  If `true`, it is reset.  If
	// `false`, only a successful login resets the threshold counter and if
	// the user reaches the AccountLockoutThreshold limit, the account shall
	// be locked out indefinitely and only an administrator-issued reset
	// clears the threshold counter.  If this property is absent, the default
	// is `true`.
	AccountLockoutCounterResetEnabled bool
	// AccountLockoutDuration shall contain the period of
	// time, in seconds, that an account is locked after the number of failed
	// login attempts reaches the AccountLockoutThreshold value, within the
	// AccountLockoutCounterResetAfter window of time.  The value shall be
	// greater than or equal to the AccountLockoutResetAfter value.  If this
	// value is `0`, no lockout shall occur.  If
	// AccountLockoutCounterResetEnabled value is `false`, this property
	// shall be ignored.
	AccountLockoutDuration int
	// AccountLockoutThreshold shall contain the threshold
	// of failed login attempts before a user account is locked.  If `0`, the
	// account shall never be locked.
	AccountLockoutThreshold int
	// accounts shall contain a link to a Resource Collection of type
	// ManagerAccountCollection.
	accounts string
	// activeDirectory shall contain the first Active
	// Directory external account provider that this Account Service
	// supports.  If the Account Service supports one or more Active
	// Directory services as an external account provider, this entity shall
	// be populated by default.  This entity shall not be present in the
	// AdditionalExternalAccountProviders Resource Collection.
	// activeDirectory string
	// additionalExternalAccountProviders shall contain the
	// additional external account providers that this Account Service uses.
	// additionalExternalAccountProviders string
	// AuthFailureLoggingThreshold shall contain the
	// threshold for when an authorization failure is logged.  This value
	// represents a modulo function.  The failure shall be logged every `n`th
	// occurrence, where `n` represents this property.
	AuthFailureLoggingThreshold int
	// Description provides a description of this resource.
	Description string
	// LDAP shall contain the first LDAP external account
	// provider that this Account Service supports.  If the Account Service
	// supports one or more LDAP services as an external account provider,
	// this entity shall be populated by default.  This entity shall not be
	// present in the AdditionalExternalAccountProviders Resource Collection.
	LDAP ExternalAccountProvider
	// LocalAccountAuth shall govern how the Service uses
	// the Accounts Resource Collection within this Account Service as part
	// of authentication.  The enumerated values describe the details for
	// each mode.
	LocalAccountAuth LocalAccountAuth
	// MaxPasswordLength shall contain the maximum password
	// length that the implementation allows for this Account Service.
	MaxPasswordLength int
	// MinPasswordLength shall contain the minimum password
	// length that the implementation allows for this Account Service.
	MinPasswordLength int
	// privilegeMap shall contain a link to a Resource of
	// type PrivilegeMapping that contains the privileges that are required
	// for a user context to complete a requested operation on a URI
	// associated with this Service.
	// privilegeMap string
	// roles shall contain a link to a Resource Collection
	// of type RoleCollection.
	roles string
	// ServiceEnabled shall indicate whether the Account
	// Service is enabled.  If `true`, it is enabled.  If `false`, it is
	// disabled and users cannot be created, deleted, or modified, and new
	// sessions cannot be started.  However, established sessions may still
	// continue to run.  Any service, such as the Session Service, that
	// attempts to access the disabled Account Service fails.  However, this
	// does not affect HTTP Basic Authentication connections.
	ServiceEnabled bool
	// Status shall contain any status or health properties
	// of the Resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an AccountService object from the raw JSON.
func (accountservice *AccountService) UnmarshalJSON(b []byte) error {
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
	err = json.Unmarshal(b, &t.Links)
	if err != nil {
		return err
	}

	*accountservice = AccountService(t.temp)

	// Extract the links to other entities for later
	accountservice.accounts = string(t.Links.Accounts)
	accountservice.roles = string(t.Links.Roles)

	// This is a read/write object, so we need to save the raw object data for later
	accountservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (accountservice *AccountService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AccountService)
	err := original.UnmarshalJSON(accountservice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccountLockoutCounterResetAfter",
		"AccountLockoutCounterResetEnabled",
		"AccountLockoutDuration",
		"AccountLockoutThreshold",
		"AuthFailureLoggingThreshold",
		"LocalAccountAuth",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(accountservice).Elem()

	return accountservice.Entity.Update(originalElement, currentElement, readWriteFields)
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
func (accountservice *AccountService) Accounts() ([]*ManagerAccount, error) {
	return ListReferencedManagerAccounts(accountservice.Client, accountservice.accounts)
}

// Roles gets the roles from the account service
func (accountservice *AccountService) Roles() ([]*Role, error) {
	return ListReferencedRoles(accountservice.Client, accountservice.roles)
}
