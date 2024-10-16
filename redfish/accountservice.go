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

// AuthenticationTypes are the different authentication types.
type AuthenticationTypes string

const (

	// TokenAuthenticationTypes is an opaque authentication token.
	TokenAuthenticationTypes AuthenticationTypes = "Token"
	// KerberosKeytabAuthenticationTypes is a Kerberos keytab.
	KerberosKeytabAuthenticationTypes AuthenticationTypes = "KerberosKeytab"
	// UsernameAndPasswordAuthenticationTypes is a user name and password
	// combination.
	UsernameAndPasswordAuthenticationTypes AuthenticationTypes = "UsernameAndPassword"
	// OEMAuthenticationTypes is an OEM-specific authentication mechanism.
	OEMAuthenticationTypes AuthenticationTypes = "OEM"
)

// BasicAuthState is the state of basic authentication support.
type BasicAuthState string

const (
	// EnabledBasicAuthState shall indicate that HTTP Basic authentication is enabled for the service. The service
	// shall include the 'WWW-Authenticate' HTTP response header with the value including 'Basic' when returning the
	// HTTP 401 (Unauthorized) status code.
	EnabledBasicAuthState BasicAuthState = "Enabled"
	// UnadvertisedBasicAuthState shall indicate that HTTP Basic authentication is enabled for the service. The service
	// shall not include 'Basic' in the value of the 'WWW-Authenticate' HTTP response header and may omit the header
	// entirely from responses. The lack of advertisement prevents some clients from accessing the service with HTTP
	// Basic authentication, such as web browsers.
	UnadvertisedBasicAuthState BasicAuthState = "Unadvertised"
	// DisabledBasicAuthState shall indicate that HTTP Basic authentication is disabled for the service.
	DisabledBasicAuthState BasicAuthState = "Disabled"
)

// CertificateMappingAttribute is how the certificate details are mapped to a user.
type CertificateMappingAttribute string

const (
	// WholeCertificateMappingAttribute shall indicate the service matches the entire certificate with a Certificate
	// resource subordinate to a ManagerAccount resource or the entire certificate matches the appropriate field from
	// an external account provider.
	WholeCertificateMappingAttribute CertificateMappingAttribute = "Whole"
	// CommonNameCertificateMappingAttribute shall indicate the service matches the RFC5280-defined 'commonName'
	// attribute in the provided certificate to the UserName property in a ManagerAccount resource or the appropriate
	// field from an external account provider.
	CommonNameCertificateMappingAttribute CertificateMappingAttribute = "CommonName"
	// UserPrincipalNameCertificateMappingAttribute shall indicate the service matches the User Principal Name (UPN)
	// field in the provided certificate to the UserName property in a ManagerAccount resource or the appropriate field
	// from an external account provider.
	UserPrincipalNameCertificateMappingAttribute CertificateMappingAttribute = "UserPrincipalName"
)

// LocalAccountAuth is the status of local account authentication support.
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

// MFABypass shall contain multi-factor authentication bypass settings.
type MFABypass struct {
	// BypassTypes shall contain the types of multi-factor authentication this account or role mapping is allowed to
	// bypass. An empty array shall indicate this account or role mapping cannot bypass any multi-factor authentication
	// types that are currently enabled.
	BypassTypes []string
}

// MicrosoftAuthenticator shall contain settings for Microsoft Authenticator multi-factor authentication.
type MicrosoftAuthenticator struct {
	// Enabled shall indicate whether multi-factor authentication with Microsoft Authenticator is enabled.
	Enabled bool
	// SecretKey shall contain the client key to use when communicating with the Microsoft Authenticator server. The
	// value shall be 'null' in responses.
	SecretKey string
	// SecretKeySet shall contain 'true' if a valid value was provided for the SecretKey property. Otherwise, the
	// property shall contain 'false'.
	SecretKeySet string
}

// OneTimePasscode shall contain settings for one-time passcode (OTP) multi-factor authentication.
type OneTimePasscode struct {
	// Enabled shall indicate whether multi-factor authentication using a one-time passcode is enabled. The passcode is
	// sent to the delivery address associated with the account credentials provided in the request. If the credentials
	// are associated with a ManagerAccount resource, the delivery address is specified by the
	// OneTimePasscodeDeliveryAddress property. If the credentials are associated with a user from an LDAP account
	// provider, the delivery address is contained in the LDAP attribute specified by the EmailAttribute property. An
	// attempt to create a session when the Token property is not included in the request shall generate a message sent
	// to the delivery address, using the SMTP settings from the Redfish event service, containing a one-time passcode.
	// The service shall accept the one-time passcode as the valid value for the Token property in the next POST
	// operation to create a session for the respective account.
	Enabled bool
}

// SecurID shall contain settings for RSA SecurID multi-factor authentication.
type SecurID struct {
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represent the
	// server certificates for the RSA SecurID server referenced by the ServerURI property. Regardless of the contents
	// of this collection, services may perform additional verification based on other factors, such as the
	// configuration of the SecurityPolicy resource.
	Certificates string
	// ClientId shall contain the client ID to use when communicating with the RSA SecurID server.
	ClientID string `json:"ClientId,omitempty"`
	// ClientSecret shall contain the client secret to use when communicating with the RSA SecurID server. The value
	// shall be 'null' in responses.
	ClientSecret string
	// ClientSecretSet shall contain 'true' if a valid value was provided for the ClientSecret property. Otherwise, the
	// property shall contain 'false'.
	ClientSecretSet string
	// Enabled shall indicate whether multi-factor authentication with RSA SecurID is enabled.
	Enabled bool
	// ServerURI shall contain the URI of the RSA SecurID server.
	ServerURI string
}

// MultiFactorAuth shall contain multi-factor authentication settings.
type MultiFactorAuth struct {
	// ClientCertificate shall contain the settings related to client certificate authentication.
	ClientCertificate ClientCertificate
	// GoogleAuthenticator shall contain the settings related to Google Authenticator multi-factor authentication.
	GoogleAuthenticator GoogleAuthenticator
	// MicrosoftAuthenticator shall contain the settings related to Microsoft Authenticator multi-factor
	// authentication.
	MicrosoftAuthenticator MicrosoftAuthenticator
	// OneTimePasscode shall contain the settings related to one-time passcode multi-factor authentication.
	OneTimePasscode OneTimePasscode
	// SecurID shall contain the settings related to RSA SecurID multi-factor authentication.
	SecurID SecurID
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

// GoogleAuthenticator shall contain settings for Google Authenticator multi-factor authentication.
type GoogleAuthenticator struct {
	// Enabled shall indicate whether multi-factor authentication with Google Authenticator is enabled.
	Enabled bool
	// SecretKey shall contain the client key to use when communicating with the Google Authenticator Server. The value
	// shall be 'null' in responses.
	SecretKey string
	// SecretKeySet shall contain 'true' if a valid value was provided for the SecretKey property. Otherwise, the
	// property shall contain 'false'.
	SecretKeySet string
}

// ClientCertificate shall contain settings for client certificate authentication.
type ClientCertificate struct {
	// CertificateMappingAttribute shall contain the client certificate attribute to map to a user.
	CertificateMappingAttribute CertificateMappingAttribute
	// certificates shall contain a link to a resource collection of type CertificateCollection that represents the CA
	// certificates used to validate client certificates during TLS handshaking. Regardless of the contents of this
	// collection, services may perform additional verification based on other factors, such as the configuration of
	// the SecurityPolicy resource. If the service supports the RevokedCertificates or TrustedCertificates properties
	// within the Client property within TLS property of the SecurityPolicy resource, the service shall verify the
	// provided client certificate with the SecurityPolicy resource prior to verifying it with this collection.
	certificates string
	// Enabled shall indicate whether client certificate authentication is enabled.
	Enabled bool
	// RespondToUnauthenticatedClients shall indicate whether the service responds to clients that do not successfully
	// authenticate. If this property is not supported by the service, it shall be assumed to be 'true'. See the
	// 'Client certificate authentication' clause in the Redfish Specification.
	RespondToUnauthenticatedClients bool
}

// UnmarshalJSON unmarshals a ClientCertificate object from the raw JSON.
func (clientcertificate *ClientCertificate) UnmarshalJSON(b []byte) error {
	type temp ClientCertificate
	var t struct {
		temp
		Certificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*clientcertificate = ClientCertificate(t.temp)

	// Extract the links to other entities for later
	clientcertificate.certificates = t.Certificates.String()

	return nil
}

// ClientCertificates gets the client certificates.
func (clientcertificate *ClientCertificate) ClientCertificates(c common.Client) ([]*Certificate, error) {
	return ListReferencedCertificates(c, clientcertificate.certificates)
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
	ActiveDirectory ExternalAccountProvider
	// additionalExternalAccountProviders shall contain the
	// additional external account providers that this Account Service uses.
	additionalExternalAccountProviders string
	// AuthFailureLoggingThreshold shall contain the
	// threshold for when an authorization failure is logged.  This value
	// represents a modulo function.  The failure shall be logged every `n`th
	// occurrence, where `n` represents this property.
	AuthFailureLoggingThreshold int
	// Description provides a description of this resource.
	Description string
	// HTTPBasicAuth shall indicate whether clients are able to authenticate to the Redfish service with HTTP Basic
	// authentication. This property should default to 'Enabled' for client compatibility. If this property is not
	// present in responses, the value shall be assumed to be 'Enabled'.
	// (Added in schema 1.15.0)
	HTTPBasicAuth BasicAuthState
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
	// MultiFactorAuth shall contain the multi-factor authentication settings that this account service supports.
	MultiFactorAuth MultiFactorAuth
	// OAuth2 is the first OAuth 2.0 external account provider that this account service supports.
	// If the account service supports one or more OAuth 2.0 services as an external account provider,
	// this entity shall be populated by default. This entity shall not be present in the additional
	// external account providers resource collection.
	OAuth2 ExternalAccountProvider
	// outboundConnections shall contain a link to a collection of type OutboundConnection.
	outboundConnections string
	// PasswordExpirationDays is the number of days before account passwords in this account service will expire.
	PasswordExpirationDays int
	// PrivilegeMap shall contain a link to a resource of type PrivilegeMapping that contains the privileges that are
	// required for a user context to complete a requested operation on a URI associated with this service.
	privilegeMap string
	// RequireChangePasswordAction shall indicate whether clients are required to invoke the ChangePassword action to
	// modify the password property in ManagerAccount resources. If 'true', services shall reject PATCH and PUT
	// requests to modify the Password property in ManagerAccount resources.
	RequireChangePasswordAction bool
	// RestrictedOemPrivileges shall contain an array of OEM privileges that are restricted by the service.
	RestrictedOemPrivileges []string
	// RestrictedPrivileges shall contain an array of Redfish privileges that are restricted by the service.
	RestrictedPrivileges []PrivilegeType
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
	// SupportedAccountTypes shall contain an array of the account types supported by the service.
	SupportedAccountTypes []AccountTypes
	// SupportedOEMAccountTypes shall contain an array of the OEM account types supported by the service.
	SupportedOEMAccountTypes []string
	// TACACSplus shall contain the first TACACS+ external account provider that this account service supports. If the
	// account service supports one or more TACACS+ services as an external account provider, this entity shall be
	// populated by default. This entity shall not be present in the additional external account providers resource
	// collection.
	TACACSplus ExternalAccountProvider
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
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
		Links                              AccountLinks
		AdditionalExternalAccountProviders common.Link
		OutboundConnections                common.Link
		PrivilegeMap                       common.Link
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
	accountservice.accounts = t.Links.Accounts.String()
	accountservice.roles = t.Links.Roles.String()

	accountservice.additionalExternalAccountProviders = t.AdditionalExternalAccountProviders.String()
	accountservice.outboundConnections = t.OutboundConnections.String()
	accountservice.privilegeMap = t.PrivilegeMap.String()

	// This is a read/write object, so we need to save the raw object data for later
	accountservice.RawData = b

	return nil
}

// AdditionalExternalAccountProviders gets additional external account providers that this account service uses.
func (accountservice *AccountService) AdditionalExternalAccountProviders() ([]*ExternalAccountProvider, error) {
	return ListReferencedExternalAccountProviders(accountservice.GetClient(), accountservice.additionalExternalAccountProviders)
}

// PrivilegeMap gets the mapping of the privileges required to complete a requested operation on a URI associated with this service.
func (accountservice *AccountService) PrivilegeMap() (*PrivilegeRegistry, error) {
	if accountservice.privilegeMap == "" {
		return nil, nil
	}
	return GetPrivilegeRegistry(accountservice.GetClient(), accountservice.privilegeMap)
}

// Update commits updates to this object's properties to the running system.
func (accountservice *AccountService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AccountService)
	err := original.UnmarshalJSON(accountservice.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccountLockoutCounterResetAfter",
		"AccountLockoutCounterResetEnabled",
		"AccountLockoutDuration",
		"AccountLockoutThreshold",
		"AuthFailureLoggingThreshold",
		"HTTPBasicAuth",
		"LocalAccountAuth",
		"RequireChangePasswordAction",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(accountservice).Elem()

	return accountservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAccountService will get the AccountService instance from the Redfish
// service.
func GetAccountService(c common.Client, uri string) (*AccountService, error) {
	return common.GetObject[AccountService](c, uri)
}

// Accounts get the accounts from the account service
func (accountservice *AccountService) Accounts() ([]*ManagerAccount, error) {
	return ListReferencedManagerAccounts(accountservice.GetClient(), accountservice.accounts)
}

// Roles gets the roles from the account service
func (accountservice *AccountService) Roles() ([]*Role, error) {
	return ListReferencedRoles(accountservice.GetClient(), accountservice.roles)
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
func (accountservice *AccountService) CreateAccount(userName, password, roleID string) (*ManagerAccount, error) {
	t := struct {
		UserName string
		Enabled  bool
		Password string
		RoleID   string `json:"RoleId"`
	}{
		UserName: userName,
		Enabled:  true,
		Password: password,
		RoleID:   roleID,
	}
	resp, err := accountservice.PostWithResponse(accountservice.accounts, t)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ManagerAccount
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
