//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/AccountService.v1_18_1.json
// 2025.1 - #AccountService.v1_18_1.AccountService

package schemas

import (
	"encoding/json"
	"fmt"
)

type AccountProviderTypes string

const (
	// RedfishServiceAccountProviderTypes shall be a DMTF Redfish
	// Specification-conformant service. The 'ServiceAddresses' property shall
	// contain URIs to 'AccountService' resources that correspond to Redfish
	// services. For example, 'https://192.168.1.50/redfish/v1/AccountService'.
	RedfishServiceAccountProviderTypes AccountProviderTypes = "RedfishService"
	// ActiveDirectoryServiceAccountProviderTypes shall be a Microsoft Active
	// Directory Technical Specification-conformant service. The 'ServiceAddresses'
	// property shall contain fully qualified domain names (FQDN) or NetBIOS names
	// that link to the domain servers for the Active Directory service.
	ActiveDirectoryServiceAccountProviderTypes AccountProviderTypes = "ActiveDirectoryService"
	// LDAPServiceAccountProviderTypes shall be an RFC4511-conformant service. The
	// 'ServiceAddresses' property shall contain RFC3986-defined URIs in the format
	// 'scheme://host:port', where 'scheme://' and ':port' are optional, that link
	// to the LDAP servers for the service. If the scheme is not specified,
	// services shall assume it is 'ldaps://'. If the port is not specified,
	// services shall assume it is '636'. For example, 'ldaps://contoso.com:636' or
	// 'contoso.com'.
	LDAPServiceAccountProviderTypes AccountProviderTypes = "LDAPService"
	// OEMAccountProviderTypes is an OEM-specific external authentication or
	// directory service.
	OEMAccountProviderTypes AccountProviderTypes = "OEM"
	// TACACSplusAccountProviderTypes shall be an RFC8907-conformant service. The
	// 'ServiceAddresses' property shall contain RFC3986-defined URIs in the format
	// 'host:port' that correspond to the TACACS+ services.
	TACACSplusAccountProviderTypes AccountProviderTypes = "TACACSplus"
	// OAuth2AccountProviderTypes shall be an RFC6749-conformant service. The
	// 'ServiceAddresses' property shall contain RFC3986-defined URIs that
	// correspond to the RFC8414-defined metadata for the OAuth 2.0 service. For
	// example, 'https://contoso.org/.well-known/oauth-authorization-server'.
	OAuth2AccountProviderTypes AccountProviderTypes = "OAuth2"
)

type AuthenticationTypes string

const (
	// TokenAuthenticationTypes is an opaque authentication token.
	TokenAuthenticationTypes AuthenticationTypes = "Token"
	// KerberosKeytabAuthenticationTypes is a Kerberos keytab.
	KerberosKeytabAuthenticationTypes AuthenticationTypes = "KerberosKeytab"
	// UsernameAndPasswordAuthenticationTypes is a username and password
	// combination.
	UsernameAndPasswordAuthenticationTypes AuthenticationTypes = "UsernameAndPassword"
	// OEMAuthenticationTypes is an OEM-specific authentication mechanism.
	OEMAuthenticationTypes AuthenticationTypes = "OEM"
)

type BasicAuthState string

const (
	// EnabledBasicAuthState shall indicate that HTTP Basic authentication is
	// enabled for the service. The service shall include the 'WWW-Authenticate'
	// HTTP response header with the value including 'Basic' when returning the
	// HTTP '401 Unauthorized' status code.
	EnabledBasicAuthState BasicAuthState = "Enabled"
	// UnadvertisedBasicAuthState shall indicate that HTTP Basic authentication is
	// enabled for the service. The service shall not include 'Basic' in the value
	// of the 'WWW-Authenticate' HTTP response header and may omit the header
	// entirely from responses. The lack of advertisement prevents some clients
	// from accessing the service with HTTP Basic authentication, such as web
	// browsers.
	UnadvertisedBasicAuthState BasicAuthState = "Unadvertised"
	// DisabledBasicAuthState shall indicate that HTTP Basic authentication is
	// disabled for the service.
	DisabledBasicAuthState BasicAuthState = "Disabled"
)

type CertificateMappingAttribute string

const (
	// WholeCertificateMappingAttribute shall indicate the service matches the
	// entire certificate with a 'Certificate' resource subordinate to a
	// 'ManagerAccount' resource or the entire certificate matches the appropriate
	// field from an external account provider.
	WholeCertificateMappingAttribute CertificateMappingAttribute = "Whole"
	// CommonNameCertificateMappingAttribute shall indicate the service matches the
	// RFC5280-defined 'commonName' attribute in the provided certificate to the
	// 'UserName' property in a 'ManagerAccount' resource or the appropriate field
	// from an external account provider.
	CommonNameCertificateMappingAttribute CertificateMappingAttribute = "CommonName"
	// UserPrincipalNameCertificateMappingAttribute shall indicate the service
	// matches the User Principal Name (UPN) field in the provided certificate to
	// the 'UserName' property in a 'ManagerAccount' resource or the appropriate
	// field from an external account provider.
	UserPrincipalNameCertificateMappingAttribute CertificateMappingAttribute = "UserPrincipalName"
)

type LocalAccountAuth string

const (
	// EnabledLocalAccountAuth shall authenticate users based on the account
	// service-defined manager accounts resource collection.
	EnabledLocalAccountAuth LocalAccountAuth = "Enabled"
	// DisabledLocalAccountAuth shall never authenticate users based on the account
	// service-defined manager accounts resource collection.
	DisabledLocalAccountAuth LocalAccountAuth = "Disabled"
	// FallbackLocalAccountAuth shall authenticate users based on the account
	// service-defined manager accounts resource collection only if any external
	// account providers are currently unreachable.
	FallbackLocalAccountAuth LocalAccountAuth = "Fallback"
	// LocalFirstLocalAccountAuth shall first authenticate users based on the
	// account service-defined manager accounts resource collection. If
	// authentication fails, the service shall authenticate by using external
	// account providers.
	LocalFirstLocalAccountAuth LocalAccountAuth = "LocalFirst"
)

type MFABypassType string

const (
	// AllMFABypassType shall indicate an account or role mapping can bypass all
	// multi-factor authentication types including OEM-defined types.
	AllMFABypassType MFABypassType = "All"
	// SecurIDMFABypassType shall indicate an account or role mapping can bypass
	// RSA SecurID. Authentication with RSA SecurID is configured with the
	// 'SecurID' property.
	SecurIDMFABypassType MFABypassType = "SecurID"
	// GoogleAuthenticatorMFABypassType shall indicate an account or role mapping
	// can bypass Google Authenticator. Authentication with Google Authenticator is
	// configured with the 'GoogleAuthenticator' property.
	GoogleAuthenticatorMFABypassType MFABypassType = "GoogleAuthenticator"
	// MicrosoftAuthenticatorMFABypassType shall indicate an account or role
	// mapping can bypass Microsoft Authenticator. Authentication with Microsoft
	// Authenticator is configured with the 'MicrosoftAuthenticator' property.
	MicrosoftAuthenticatorMFABypassType MFABypassType = "MicrosoftAuthenticator"
	// ClientCertificateMFABypassType shall indicate an account or role mapping can
	// bypass client certificate authentication. Authentication with client
	// certificates is configured with the 'ClientCertificate' property.
	ClientCertificateMFABypassType MFABypassType = "ClientCertificate"
	// OneTimePasscodeMFABypassType shall indicate an account or role mapping can
	// bypass one-time passcode authentication. Authentication with a one-time
	// passcode is configured with the 'OneTimePasscode' property.
	OneTimePasscodeMFABypassType MFABypassType = "OneTimePasscode"
	// TimeBasedOneTimePasswordMFABypassType shall indicate an account or role
	// mapping can bypass RFC6238-defined Time-based One-Time Password (TOTP)
	// authentication. Authentication with a Time-based One-Time Password is
	// configured with the 'TimeBasedOneTimePassword' property.
	TimeBasedOneTimePasswordMFABypassType MFABypassType = "TimeBasedOneTimePassword"
	// OEMMFABypassType shall indicate an account or role mapping can bypass
	// OEM-defined multi-factor authentication.
	OEMMFABypassType MFABypassType = "OEM"
)

type OAuth2Mode string

const (
	// DiscoveryOAuth2Mode shall indicate the service performs token validation
	// from information found at the URIs specified by the 'ServiceAddresses'
	// property. Services shall implement a caching method of this information so
	// it's not necessary to retrieve metadata and key information for every
	// request containing a token.
	DiscoveryOAuth2Mode OAuth2Mode = "Discovery"
	// OfflineOAuth2Mode shall indicate the service performs token validation from
	// properties configured by a client. Clients should configure the 'Issuer' and
	// 'OAuthServiceSigningKeys' properties for this mode.
	OfflineOAuth2Mode OAuth2Mode = "Offline"
)

type TACACSplusPasswordExchangeProtocol string

const (
	// ASCIITACACSplusPasswordExchangeProtocol shall indicate the ASCII Login flow
	// as described under section 5.4.2 of RFC8907.
	ASCIITACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "ASCII"
	// PAPTACACSplusPasswordExchangeProtocol shall indicate the PAP Login flow as
	// described under section 5.4.2 of RFC8907.
	PAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "PAP"
	// CHAPTACACSplusPasswordExchangeProtocol shall indicate the CHAP Login flow as
	// described under section 5.4.2 of RFC8907.
	CHAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "CHAP"
	// MSCHAPv1TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v1
	// Login flow as described under section 5.4.2 of RFC8907.
	MSCHAPv1TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv1"
	// MSCHAPv2TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v2
	// Login flow as described under section 5.4.2 of RFC8907.
	MSCHAPv2TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv2"
)

// AccountService shall represent an account service for a Redfish
// implementation. The properties are common to, and enable management of, all
// user accounts. The properties include the password requirements and control
// features, such as account lockout. Properties and actions in this service
// specify general behavior that should be followed for typical accounts,
// however implementations may override these behaviors for special accounts or
// situations to avoid denial of service or other deadlock situations.
type AccountService struct {
	Entity
	// AccountLockoutCounterResetAfter shall contain the period of time, in
	// seconds, from the last failed login attempt when the
	// 'AccountLockoutThreshold' counter, which counts the number of failed login
	// attempts, is reset to '0'. Then, 'AccountLockoutThreshold' failures are
	// required before the account is locked. This value shall be less than or
	// equal to the 'AccountLockoutDuration' value. The threshold counter also
	// resets to '0' after each successful login. If the
	// 'AccountLockoutCounterResetEnabled' value is 'false', this property shall be
	// ignored.
	AccountLockoutCounterResetAfter uint
	// AccountLockoutCounterResetEnabled shall indicate whether the threshold
	// counter is reset after the 'AccountLockoutCounterResetAfter' expires. If
	// 'true', it is reset. If 'false', only a successful login resets the
	// threshold counter and if the user reaches the 'AccountLockoutThreshold'
	// limit, the account shall be locked out indefinitely and only an
	// administrator-issued reset clears the threshold counter. If this property is
	// absent, the default is 'true'.
	//
	// Version added: v1.5.0
	AccountLockoutCounterResetEnabled bool
	// AccountLockoutDuration shall contain the period of time, in seconds, that an
	// account is locked after the number of failed login attempts reaches the
	// 'AccountLockoutThreshold' value, within the
	// 'AccountLockoutCounterResetAfter' window of time. The value shall be greater
	// than or equal to the 'AccountLockoutCounterResetAfter' value. If this value
	// is '0', no lockout shall occur. If 'AccountLockoutCounterResetEnabled' value
	// is 'false', this property shall be ignored.
	AccountLockoutDuration *uint `json:",omitempty"`
	// AccountLockoutThreshold shall contain the threshold of failed login attempts
	// before a user account is locked. If '0', the account shall never be locked.
	AccountLockoutThreshold *uint `json:",omitempty"`
	// Accounts shall contain a link to a resource collection of type
	// 'ManagerAccountCollection'.
	accounts string
	// ActiveDirectory shall contain the first Active Directory external account
	// provider that this account service supports. If the account service supports
	// one or more Active Directory services as an external account provider, this
	// entity shall be populated by default. This entity shall not be present in
	// the additional external account providers resource collection.
	//
	// Version added: v1.3.0
	ActiveDirectory ExternalAccountProvider
	// AdditionalExternalAccountProviders shall contain a link to a resource
	// collection of type 'ExternalAccountProviderCollection' that represents the
	// additional external account providers that this account service uses.
	//
	// Version added: v1.3.0
	additionalExternalAccountProviders string
	// AuthFailureLoggingThreshold shall contain the threshold for when an
	// authorization failure is logged. Logging shall occur after every 'n'
	// occurrences of an authorization failure on the same account, where 'n'
	// represents the value of this property. If the value is '0', logging of
	// authorization failures shall be disabled.
	AuthFailureLoggingThreshold uint
	// EnforcePasswordHistoryCount shall contain the number of unique new passwords
	// that need to be associated with a user account before a previous password is
	// accepted when modifying the password. If not '0', services shall reject
	// modification requests of the 'Password' property and 'ChangePassword'
	// actions that contain a previously used password in the specified count. If
	// '0', services shall not require the user to provide a unique new password.
	// This property does not apply to accounts from external account providers.
	//
	// Version added: v1.17.0
	EnforcePasswordHistoryCount uint
	// HTTPBasicAuth shall indicate whether clients are able to authenticate to the
	// Redfish service with HTTP Basic authentication. This property should default
	// to 'Enabled' for client compatibility. If this property is not present in
	// responses, the value shall be assumed to be 'Enabled'.
	//
	// Version added: v1.15.0
	HTTPBasicAuth BasicAuthState
	// LDAP shall contain the first LDAP external account provider that this
	// account service supports. If the account service supports one or more LDAP
	// services as an external account provider, this entity shall be populated by
	// default. This entity shall not be present in the additional external account
	// providers resource collection.
	//
	// Version added: v1.3.0
	LDAP ExternalAccountProvider
	// LocalAccountAuth shall govern how the service uses the manager accounts
	// resource collection within this account service as part of authentication.
	// The enumerated values describe the details for each mode.
	//
	// Version added: v1.3.0
	LocalAccountAuth LocalAccountAuth
	// MaxPasswordLength shall contain the maximum password length that the
	// implementation allows for this account service. This property does not apply
	// to accounts from external account providers.
	MaxPasswordLength uint
	// MinPasswordLength shall contain the minimum password length that the
	// implementation allows for this account service. This property does not apply
	// to accounts from external account providers.
	MinPasswordLength uint
	// MultiFactorAuth shall contain the multi-factor authentication settings that
	// this account service supports.
	//
	// Version added: v1.12.0
	MultiFactorAuth MultiFactorAuth
	// OAuth2 shall contain the first OAuth 2.0 external account provider that this
	// account service supports. If the account service supports one or more OAuth
	// 2.0 services as an external account provider, this entity shall be populated
	// by default. This entity shall not be present in the additional external
	// account providers resource collection.
	//
	// Version added: v1.10.0
	OAuth2 ExternalAccountProvider
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutboundConnections shall contain a resource collection of type
	// 'OutboundConnectionCollection'.
	//
	// Version added: v1.14.0
	OutboundConnections OutboundConnection
	// PasswordExpirationDays shall contain the number of days before account
	// passwords in this account service will expire. The value shall be applied
	// during account creation and password modification unless the
	// 'PasswordExpiration' property is provided. The value 'null' shall indicate
	// that account passwords never expire. This property does not apply to
	// accounts from external account providers.
	//
	// Version added: v1.9.0
	PasswordExpirationDays *int `json:",omitempty"`
	// PasswordGuidanceMessage shall contain guidance for creating passwords that
	// meet the password complexity or other related requirements for this service.
	//
	// Version added: v1.18.0
	PasswordGuidanceMessage string
	// PasswordGuidanceMessageID shall contain a 'MessageId' value that contains
	// guidance for creating passwords that meet the password complexity or other
	// related requirements for this service. The value shall contain a
	// 'MessageId', as defined in the 'MessageId format' clause of the Redfish
	// Specification.
	//
	// Version added: v1.18.0
	PasswordGuidanceMessageID string `json:"PasswordGuidanceMessageId"`
	// PrivilegeMap shall contain a link to a resource of type 'PrivilegeMapping'
	// that contains the privileges that are required for a user context to
	// complete a requested operation on a URI associated with this service.
	//
	// Version added: v1.1.0
	privilegeMap string
	// RequireChangePasswordAction shall indicate whether clients are required to
	// invoke the 'ChangePassword' action to modify the 'Password' property in
	// 'ManagerAccount' resources. If 'true', services shall reject 'PATCH' and
	// 'PUT' requests to modify the 'Password' property in 'ManagerAccount'
	// resources.
	//
	// Version added: v1.14.0
	RequireChangePasswordAction bool
	// RestrictedOemPrivileges shall contain an array of OEM privileges that are
	// restricted by the service.
	//
	// Version added: v1.8.0
	RestrictedOemPrivileges []string
	// RestrictedPrivileges shall contain an array of Redfish privileges that are
	// restricted by the service.
	//
	// Version added: v1.8.0
	RestrictedPrivileges []PrivilegeType
	// Roles shall contain a link to a resource collection of type
	// 'RoleCollection'.
	roles string
	// ServiceEnabled shall indicate whether the account service is enabled. If
	// 'true', it is enabled. If 'false', it is disabled and users cannot be
	// created, deleted, or modified, and new sessions cannot be started. However,
	// established sessions may still continue to run. Any service, such as the
	// session service, that attempts to access the disabled account service fails.
	// However, this does not affect HTTP Basic Authentication connections.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedAccountTypes shall contain an array of the account types supported
	// by the service.
	//
	// Version added: v1.8.0
	SupportedAccountTypes []AccountTypes
	// SupportedOEMAccountTypes shall contain an array of the OEM account types
	// supported by the service.
	//
	// Version added: v1.8.0
	SupportedOEMAccountTypes []string
	// TACACSplus shall contain the first TACACS+ external account provider that
	// this account service supports. If the account service supports one or more
	// TACACS+ services as an external account provider, this entity shall be
	// populated by default. This entity shall not be present in the additional
	// external account providers resource collection.
	//
	// Version added: v1.8.0
	TACACSplus ExternalAccountProvider
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a AccountService object from the raw JSON.
func (a *AccountService) UnmarshalJSON(b []byte) error {
	type temp AccountService
	var tmp struct {
		temp
		Accounts                           Link `json:"Accounts"`
		AdditionalExternalAccountProviders Link `json:"AdditionalExternalAccountProviders"`
		PrivilegeMap                       Link `json:"PrivilegeMap"`
		Roles                              Link `json:"Roles"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AccountService(tmp.temp)

	// Extract the links to other entities for later
	a.accounts = tmp.Accounts.String()
	a.additionalExternalAccountProviders = tmp.AdditionalExternalAccountProviders.String()
	a.privilegeMap = tmp.PrivilegeMap.String()
	a.roles = tmp.Roles.String()

	// This is a read/write object, so we need to save the raw object data for later
	a.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AccountService) Update() error {
	readWriteFields := []string{
		"AccountLockoutCounterResetAfter",
		"AccountLockoutCounterResetEnabled",
		"AccountLockoutDuration",
		"AccountLockoutThreshold",
		"AuthFailureLoggingThreshold",
		"EnforcePasswordHistoryCount",
		"HTTPBasicAuth",
		"LocalAccountAuth",
		"MaxPasswordLength",
		"MinPasswordLength",
		"PasswordExpirationDays",
		"RequireChangePasswordAction",
		"ServiceEnabled",
	}

	return a.UpdateFromRawData(a, a.RawData, readWriteFields)
}

// GetAccountService will get a AccountService instance from the service.
func GetAccountService(c Client, uri string) (*AccountService, error) {
	return GetObject[AccountService](c, uri)
}

// ListReferencedAccountServices gets the collection of AccountService from
// a provided reference.
func ListReferencedAccountServices(c Client, link string) ([]*AccountService, error) {
	return GetCollectionObjects[AccountService](c, link)
}

// Accounts gets the Accounts collection.
func (a *AccountService) Accounts() ([]*ManagerAccount, error) {
	if a.accounts == "" {
		return nil, nil
	}
	return GetCollectionObjects[ManagerAccount](a.client, a.accounts)
}

// AdditionalExternalAccountProviders gets the AdditionalExternalAccountProviders collection.
func (a *AccountService) AdditionalExternalAccountProviders() ([]*ExternalAccountProvider, error) {
	if a.additionalExternalAccountProviders == "" {
		return nil, nil
	}
	return GetCollectionObjects[ExternalAccountProvider](a.client, a.additionalExternalAccountProviders)
}

// PrivilegeMap gets the PrivilegeMap linked resource.
func (a *AccountService) PrivilegeMap() (*PrivilegeRegistry, error) {
	if a.privilegeMap == "" {
		return nil, nil
	}
	return GetObject[PrivilegeRegistry](a.client, a.privilegeMap)
}

// Roles gets the Roles collection.
func (a *AccountService) Roles() ([]*Role, error) {
	if a.roles == "" {
		return nil, nil
	}
	return GetCollectionObjects[Role](a.client, a.roles)
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

	resp, err := baseEntity.PostWithResponse(accountservice.accounts, t)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var result ManagerAccount
	err = json.NewDecoder(resp.Body).Decode(&result)
	result.SetClient(accountservice.GetClient())

	return &result, err
}

// Authentication shall contain the information required to authenticate to the
// external service.
type Authentication struct {
	// AuthenticationType shall contain the type of authentication used to connect
	// to the external account provider.
	//
	// Version added: v1.3.0
	AuthenticationType AuthenticationTypes
	// EncryptionKey shall contain the value of a symmetric encryption key for
	// account services that support some form of encryption, obfuscation, or
	// authentication such as TACACS+. The value shall be 'null' in responses. The
	// property shall accept a hexadecimal string whose length depends on the
	// external account service, such as TACACS+. A TACACS+ service shall use this
	// property to specify the secret key as defined in RFC8907.
	//
	// Version added: v1.8.0
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the
	// 'EncryptionKey' property. Otherwise, the property shall contain 'false'. For
	// a TACACS+ service, the value 'false' shall indicate data obfuscation, as
	// defined in section 4.5 of RFC8907, is disabled.
	//
	// Version added: v1.8.0
	EncryptionKeySet bool
	// KerberosKeytab shall contain a Base64-encoded string, with padding
	// characters, of the Kerberos keytab for this service. A 'PATCH' or 'PUT'
	// operation writes the keytab. The value shall be 'null' in responses.
	//
	// Version added: v1.3.0
	KerberosKeytab string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password for this service. A 'PATCH' or 'PUT'
	// operation writes the password. The value shall be 'null' in responses.
	//
	// Version added: v1.3.0
	Password string
	// Token shall contain the token for this service. A 'PATCH' or 'PUT' operation
	// writes the token. The value shall be 'null' in responses.
	//
	// Version added: v1.3.0
	Token string
	// Username shall contain the username for this service.
	//
	// Version added: v1.3.0
	Username string
}

// ClientCertificate shall contain settings for client certificate
// authentication.
type ClientCertificate struct {
	// CertificateMappingAttribute shall contain the client certificate attribute
	// to map to a user.
	//
	// Version added: v1.12.0
	CertificateMappingAttribute CertificateMappingAttribute
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the CA certificates used to validate
	// client certificates during TLS handshaking. Regardless of the contents of
	// this collection, services may perform additional verification based on other
	// factors, such as the configuration of the 'SecurityPolicy' resource. If the
	// service supports the 'RevokedCertificates' or 'TrustedCertificates'
	// properties within the 'Server' property within the 'TLS' property of the
	// 'SecurityPolicy' resource, the service shall verify the provided client
	// certificate with the 'SecurityPolicy' resource prior to verifying it with
	// this collection.
	//
	// Version added: v1.12.0
	certificates string
	// Enabled shall indicate whether client certificate authentication is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// RespondToUnauthenticatedClients shall indicate whether the service responds
	// to clients that do not successfully authenticate. If this property is not
	// supported by the service, it shall be assumed to be 'true'. See the 'Client
	// certificate authentication' clause in the Redfish Specification.
	//
	// Version added: v1.12.0
	RespondToUnauthenticatedClients bool
}

// UnmarshalJSON unmarshals a ClientCertificate object from the raw JSON.
func (c *ClientCertificate) UnmarshalJSON(b []byte) error {
	type temp ClientCertificate
	var tmp struct {
		temp
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ClientCertificate(tmp.temp)

	// Extract the links to other entities for later
	c.certificates = tmp.Certificates.String()

	return nil
}

// Certificates gets the Certificates collection.
func (c *ClientCertificate) Certificates(client Client) ([]*Certificate, error) {
	if c.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, c.certificates)
}

// GoogleAuthenticator shall contain settings for Google Authenticator
// multi-factor authentication.
type GoogleAuthenticator struct {
	// Enabled shall indicate whether multi-factor authentication with Google
	// Authenticator is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// SecretKey shall contain the client key to use when communicating with the
	// Google Authenticator Server. The value shall be 'null' in responses.
	//
	// Version added: v1.12.0
	SecretKey string
	// SecretKeySet shall contain 'true' if a valid value was provided for the
	// 'SecretKey' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.12.0
	SecretKeySet bool
}

// LDAPSearchSettings shall contain all required settings to search a generic
// LDAP service.
type LDAPSearchSettings struct {
	// BaseDistinguishedNames shall contain an array of base distinguished names to
	// use to search an external LDAP service.
	//
	// Version added: v1.3.0
	BaseDistinguishedNames []string
	// EmailAttribute shall contain the attribute name that contains the LDAP
	// user's email address. If this value is not set by the user, or the property
	// is not present, the value shall be 'mail'.
	//
	// Version added: v1.14.0
	EmailAttribute string
	// GroupNameAttribute shall contain the attribute name that contains the LDAP
	// group name.
	//
	// Version added: v1.3.0
	GroupNameAttribute string
	// GroupsAttribute shall contain the attribute name that contains the groups
	// for an LDAP user entry.
	//
	// Version added: v1.3.0
	GroupsAttribute string
	// SSHKeyAttribute shall contain the attribute name that contains the LDAP
	// user's SSH public key.
	//
	// Version added: v1.11.0
	SSHKeyAttribute string
	// UsernameAttribute shall contain the attribute name that contains the LDAP
	// username.
	//
	// Version added: v1.3.0
	UsernameAttribute string
}

// LDAPService shall contain all required settings to parse a generic LDAP
// service.
type LDAPService struct {
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
	// SearchSettings shall contain the required settings to search an external
	// LDAP service.
	//
	// Version added: v1.3.0
	SearchSettings LDAPSearchSettings
}

// MFABypass shall contain multi-factor authentication bypass settings.
type MFABypass struct {
	// BypassTypes shall contain the types of multi-factor authentication this
	// account or role mapping is allowed to bypass. An empty array shall indicate
	// this account or role mapping cannot bypass any multi-factor authentication
	// types that are currently enabled.
	//
	// Version added: v1.12.0
	BypassTypes []MFABypassType
}

// MicrosoftAuthenticator shall contain settings for Microsoft Authenticator
// multi-factor authentication.
type MicrosoftAuthenticator struct {
	// Enabled shall indicate whether multi-factor authentication with Microsoft
	// Authenticator is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// SecretKey shall contain the client key to use when communicating with the
	// Microsoft Authenticator server. The value shall be 'null' in responses.
	//
	// Version added: v1.12.0
	SecretKey string
	// SecretKeySet shall contain 'true' if a valid value was provided for the
	// 'SecretKey' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.12.0
	SecretKeySet bool
}

// MultiFactorAuth shall contain multi-factor authentication settings.
type MultiFactorAuth struct {
	// ClientCertificate shall contain the settings related to client certificate
	// authentication.
	//
	// Version added: v1.12.0
	ClientCertificate ClientCertificate
	// GoogleAuthenticator shall contain the settings related to Google
	// Authenticator multi-factor authentication.
	//
	// Version added: v1.12.0
	GoogleAuthenticator GoogleAuthenticator
	// MicrosoftAuthenticator shall contain the settings related to Microsoft
	// Authenticator multi-factor authentication.
	//
	// Version added: v1.12.0
	MicrosoftAuthenticator MicrosoftAuthenticator
	// OneTimePasscode shall contain the settings related to one-time passcode
	// multi-factor authentication.
	//
	// Version added: v1.14.0
	OneTimePasscode OneTimePasscode
	// SecurID shall contain the settings related to RSA SecurID multi-factor
	// authentication.
	//
	// Version added: v1.12.0
	SecurID SecurID
	// TimeBasedOneTimePassword shall contain the settings related to
	// RFC6238-defined Time-based One-Time Password (TOTP) multi-factor
	// authentication.
	//
	// Version added: v1.16.0
	TimeBasedOneTimePassword TimeBasedOneTimePassword
}

// OAuth2Service shall contain settings for parsing an OAuth 2.0 service.
type OAuth2Service struct {
	// Audience shall contain an array of allowable RFC7519-defined audience
	// strings of the Redfish service. The values shall uniquely identify the
	// Redfish service. For example, a MAC address or UUID for the manager can
	// uniquely identify the service.
	//
	// Version added: v1.10.0
	Audience []string
	// Issuer shall contain the RFC8414-defined issuer string of the OAuth 2.0
	// service. If the 'Mode' property contains the value 'Discovery', this
	// property shall contain the value of the 'issuer' string from the OAuth 2.0
	// service's metadata and this property shall be read-only. Clients should
	// configure this property if 'Mode' contains 'Offline'.
	//
	// Version added: v1.10.0
	Issuer string
	// Mode shall contain the mode of operation for token validation.
	//
	// Version added: v1.10.0
	Mode OAuth2Mode
	// OAuthServiceSigningKeys shall contain a Base64-encoded string, with padding
	// characters, of the RFC7517-defined signing keys of the issuer of the OAuth
	// 2.0 service. Services shall verify the token provided in the 'Authorization'
	// header of the request with the value of this property. If the 'Mode'
	// property contains the value 'Discovery', this property shall contain the
	// keys found at the URI specified by the 'jwks_uri' string from the OAuth 2.0
	// service's metadata and this property shall be read-only. Clients should
	// configure this property if 'Mode' contains 'Offline'.
	//
	// Version added: v1.10.0
	OAuthServiceSigningKeys string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.13.0
	OEM json.RawMessage `json:"Oem"`
}

// OneTimePasscode shall contain settings for one-time passcode (OTP)
// multi-factor authentication.
type OneTimePasscode struct {
	// Enabled shall indicate whether multi-factor authentication using a one-time
	// passcode is enabled. The passcode is sent to the delivery address associated
	// with the account credentials provided in the request. If the credentials are
	// associated with a 'ManagerAccount' resource, the delivery address is
	// specified by the 'OneTimePasscodeDeliveryAddress' property. If the
	// credentials are associated with a user from an LDAP account provider, the
	// delivery address is contained in the LDAP attribute specified by the
	// 'EmailAttribute' property. An attempt to create a session when the 'Token'
	// property is not included in the request shall generate a message sent to the
	// delivery address, using the SMTP settings from the Redfish event service,
	// containing a one-time passcode. The service shall accept the one-time
	// passcode as the valid value for the 'Token' property in the next 'POST'
	// operation to create a session for the respective account.
	//
	// Version added: v1.14.0
	Enabled bool
}

// RoleMapping shall contain mapping rules that are used to convert the external
// account providers account information to the local Redfish role.
type RoleMapping struct {
	// LocalAccountTypes shall contain an array of the various local manager
	// services that the remote user or group is allowed to access. This shall not
	// include functionality for receiving events or other notifications. If this
	// property is not supported, the value shall be assumed to be an array that
	// contains the value 'Redfish'.
	//
	// Version added: v1.16.0
	LocalAccountTypes []AccountTypes
	// LocalOEMAccountTypes shall contain an array of the OEM account types for the
	// remote user or group when 'LocalAccountTypes' contains 'OEM'.
	//
	// Version added: v1.16.0
	LocalOEMAccountTypes []string
	// LocalRole shall contain the 'RoleId' property value within a role resource
	// on this Redfish service to which to map the remote user or group.
	//
	// Version added: v1.3.0
	LocalRole string
	// MFABypass shall contain the multi-factor authentication bypass settings.
	//
	// Version added: v1.12.0
	MFABypass MFABypass
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
	// RemoteGroup shall contain the name of the remote group, or the remote role
	// in the case of a Redfish service, that maps to the local Redfish role to
	// which this entity links. If within the 'ActiveDirectory' property, this
	// property shall contain the value of the 'Group Name' attribute from the
	// Active Directory service. If within the 'LDAP' property, this property shall
	// contain the distinguished name for the group or the value of the attribute
	// specified by the 'GroupNameAttribute' property from the LDAP service.
	//
	// Version added: v1.3.0
	RemoteGroup string
	// RemoteUser shall contain the name of the remote user that maps to the local
	// Redfish role to which this entity links. If within the 'ActiveDirectory'
	// property, this property shall contain the value of the 'Username' attribute,
	// with optional domain, from the Active Directory service. If within the
	// 'LDAP' property, this property shall contain the distinguished name for the
	// user or the value of the attribute specified by the 'UsernameAttribute'
	// property from the LDAP service.
	//
	// Version added: v1.3.0
	RemoteUser string
}

// SecurID shall contain settings for RSA SecurID multi-factor authentication.
type SecurID struct {
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represent the server certificates for the RSA
	// SecurID server referenced by the 'ServerURI' property. Regardless of the
	// contents of this collection, services may perform additional verification
	// based on other factors, such as the configuration of the 'SecurityPolicy'
	// resource.
	//
	// Version added: v1.12.0
	certificates string
	// ClientID shall contain the client ID to use when communicating with the RSA
	// SecurID server.
	//
	// Version added: v1.12.0
	ClientID string `json:"ClientId"`
	// ClientSecret shall contain the client secret to use when communicating with
	// the RSA SecurID server. The value shall be 'null' in responses.
	//
	// Version added: v1.12.0
	ClientSecret string
	// ClientSecretSet shall contain 'true' if a valid value was provided for the
	// 'ClientSecret' property. Otherwise, the property shall contain 'false'.
	//
	// Version added: v1.12.0
	ClientSecretSet bool
	// Enabled shall indicate whether multi-factor authentication with RSA SecurID
	// is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// ServerURI shall contain the URI of the RSA SecurID server.
	//
	// Version added: v1.12.0
	ServerURI string
}

// UnmarshalJSON unmarshals a SecurID object from the raw JSON.
func (s *SecurID) UnmarshalJSON(b []byte) error {
	type temp SecurID
	var tmp struct {
		temp
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SecurID(tmp.temp)

	// Extract the links to other entities for later
	s.certificates = tmp.Certificates.String()

	return nil
}

// Certificates gets the Certificates collection.
func (s *SecurID) Certificates(client Client) ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, s.certificates)
}

// TACACSplusService shall contain settings for parsing a TACACS+ service.
type TACACSplusService struct {
	// AuthorizationService shall contain the TACACS+ service authorization
	// argument as defined by section 8.2 of RFC8907. If this property is not
	// present, the service defines the value to provide to the TACACS+ server.
	//
	// Version added: v1.13.0
	AuthorizationService string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.13.0
	OEM json.RawMessage `json:"Oem"`
	// PasswordExchangeProtocols shall indicate all the allowed TACACS+ password
	// exchange protocol described under section 5.4.2 of RFC8907.
	//
	// Version added: v1.8.0
	PasswordExchangeProtocols []TACACSplusPasswordExchangeProtocol
	// PrivilegeLevelArgument shall specify the name of the argument in a TACACS+
	// Authorization REPLY packet body, as defined in RFC8907, that contains the
	// user's privilege level.
	//
	// Version added: v1.8.0
	PrivilegeLevelArgument string
}

// TimeBasedOneTimePassword shall contain settings for RFC6238-defined
// Time-based One-Time Password (TOTP) multi-factor authentication.
type TimeBasedOneTimePassword struct {
	// Enabled shall indicate whether multi-factor authentication with an
	// RFC6238-defined Time-based One-Time Password (TOTP) is enabled.
	//
	// Version added: v1.16.0
	Enabled bool
	// TimeStepSeconds shall contain the RFC6238-defined time step, in seconds, for
	// calculating the one-time password. If this property is not supported by the
	// service, it shall be assumed to be '30'.
	//
	// Version added: v1.16.0
	TimeStepSeconds *uint `json:",omitempty"`
}
