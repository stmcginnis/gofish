//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type OAuth2Mode string

const (
	// DiscoveryOAuth2Mode shall indicate the service performs token validation from information found at the URIs
	// specified by the ServiceAddresses property. Services shall implement a caching method of this information so
	// it's not necessary to retrieve metadata and key information for every request containing a token.
	DiscoveryOAuth2Mode OAuth2Mode = "Discovery"
	// OfflineOAuth2Mode shall indicate the service performs token validation from properties configured by a client.
	// Clients should configure the Issuer and OAuthServiceSigningKeys properties for this mode.
	OfflineOAuth2Mode OAuth2Mode = "Offline"
)

type TACACSplusPasswordExchangeProtocol string

const (
	// ASCIITACACSplusPasswordExchangeProtocol shall indicate the ASCII Login flow as described under section 5.4.2 of
	// RFC8907.
	ASCIITACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "ASCII"
	// PAPTACACSplusPasswordExchangeProtocol shall indicate the PAP Login flow as described under section 5.4.2 of
	// RFC8907.
	PAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "PAP"
	// CHAPTACACSplusPasswordExchangeProtocol shall indicate the CHAP Login flow as described under section 5.4.2 of
	// RFC8907.
	CHAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "CHAP"
	// MSCHAPv1TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v1 Login flow as described under section
	// 5.4.2 of RFC8907.
	MSCHAPv1TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv1"
	// MSCHAPv2TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v2 Login flow as described under section
	// 5.4.2 of RFC8907.
	MSCHAPv2TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv2"
)

// ExternalAccountProvider shall represent a remote authentication service in the Redfish Specification.
type ExternalAccountProvider struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountProviderType shall contain the type of external account provider to which this service connects.
	AccountProviderType AccountProviderTypes
	// Authentication shall contain the authentication information for the external account provider.
	Authentication Authentication
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates the external account provider uses.
	certificates string
	// Description provides a description of this resource.
	Description string
	// LDAPService shall contain any additional mapping information needed to parse a generic LDAP service. This
	// property should only be present if AccountProviderType is 'LDAPService'.
	LDAPService LDAPService
	// OAuth2Service shall contain additional information needed to parse an OAuth 2.0 service. This property should
	// only be present inside an OAuth2 property.
	OAuth2Service OAuth2Service
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Priority shall contain the assigned priority for the specified external account provider. The value '0' shall
	// indicate the highest priority. Increasing values shall represent decreasing priority. If an external provider
	// does not have a priority assignment or two or more external providers have the same priority, the behavior shall
	// be determined by the Redfish service. The priority is used to determine the order of authentication and
	// authorization for each external account provider.
	Priority int
	// RemoteRoleMapping shall contain a set of the mapping rules that are used to convert the external account
	// providers account information to the local Redfish role.
	RemoteRoleMapping []RoleMapping
	// Retries shall contain the number of retries to attempt a connection to an address in the ServiceAddresses
	// property before attempting a connection to the next address in the array or giving up. If this property is not
	// present, the service has internal policies for handling retries.
	Retries int
	// ServiceAddresses shall contain the addresses of the account providers to which this external account provider
	// links. The format of this field depends on the type of external account provider. Each item in the array shall
	// contain a single address. Services can define their own behavior for managing multiple addresses.
	ServiceAddresses []string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// TACACSplusService shall contain additional information needed to parse a TACACS+ services. This property should
	// only be present inside a TACACSplus property.
	TACACSplusService TACACSplusService
	// TimeoutSeconds shall contain the period of time, in seconds, this account service will wait for a response from
	// an address of a user account provider before timing out. If this property is not present, the service has
	// internal policies for handling timeouts.
	TimeoutSeconds int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ExternalAccountProvider object from the raw JSON.
func (externalaccountprovider *ExternalAccountProvider) UnmarshalJSON(b []byte) error {
	type temp ExternalAccountProvider
	var t struct {
		temp
		Certificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*externalaccountprovider = ExternalAccountProvider(t.temp)

	// Extract the links to other entities for later
	externalaccountprovider.certificates = t.Certificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	externalaccountprovider.rawData = b

	return nil
}

// Certificates returns certificates in this external account provider.
func (externalaccountprovider *ExternalAccountProvider) Certificates() ([]*Certificate, error) {
	if externalaccountprovider.certificates == "" {
		return []*Certificate{}, nil
	}
	return ListReferencedCertificates(externalaccountprovider.GetClient(), externalaccountprovider.certificates)
}

// Update commits updates to this object's properties to the running system.
func (externalaccountprovider *ExternalAccountProvider) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ExternalAccountProvider)
	original.UnmarshalJSON(externalaccountprovider.rawData)

	readWriteFields := []string{
		"Priority",
		"Retries",
		"ServiceAddresses",
		"ServiceEnabled",
		"TimeoutSeconds",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(externalaccountprovider).Elem()

	return externalaccountprovider.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetExternalAccountProvider will get a ExternalAccountProvider instance from the service.
func GetExternalAccountProvider(c common.Client, uri string) (*ExternalAccountProvider, error) {
	return common.GetObject[ExternalAccountProvider](c, uri)
}

// ListReferencedExternalAccountProviders gets the collection of ExternalAccountProvider from
// a provided reference.
func ListReferencedExternalAccountProviders(c common.Client, link string) ([]*ExternalAccountProvider, error) {
	return common.GetCollectionObjects[ExternalAccountProvider](c, link)
}

// LDAPSearchSettings shall contain all required settings to search a generic LDAP service.
type LDAPSearchSettings struct {
	// BaseDistinguishedNames shall contain an array of base distinguished names to use to search an external LDAP
	// service.
	BaseDistinguishedNames []string
	// EmailAttribute shall contain the attribute name that contains the LDAP user's email address. If this value is
	// not set by the user, or the property is not present, the value shall be 'mail'.
	EmailAttribute string
	// GroupNameAttribute shall contain the attribute name that contains the LDAP group name.
	GroupNameAttribute string
	// GroupsAttribute shall contain the attribute name that contains the groups for an LDAP user entry.
	GroupsAttribute string
	// SSHKeyAttribute shall contain the attribute name that contains the LDAP user's SSH public key.
	SSHKeyAttribute string
	// UsernameAttribute shall contain the attribute name that contains the LDAP user name.
	UsernameAttribute string
}

// LDAPService shall contain all required settings to parse a generic LDAP service.
type LDAPService struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SearchSettings shall contain the required settings to search an external LDAP service.
	SearchSettings LDAPSearchSettings
}

// OAuth2Service shall contain settings for parsing an OAuth 2.0 service.
type OAuth2Service struct {
	// Audience shall contain an array of allowable RFC7519-defined audience strings of the Redfish service. The values
	// shall uniquely identify the Redfish service. For example, a MAC address or UUID for the manager can uniquely
	// identify the service.
	Audience []string
	// Issuer shall contain the RFC8414-defined issuer string of the OAuth 2.0 service. If the Mode property contains
	// the value 'Discovery', this property shall contain the value of the 'issuer' string from the OAuth 2.0 service's
	// metadata and this property shall be read-only. Clients should configure this property if Mode contains
	// 'Offline'.
	Issuer string
	// Mode shall contain the mode of operation for token validation.
	Mode OAuth2Mode
	// OAuthServiceSigningKeys shall contain a Base64-encoded string of the RFC7517-defined signing keys of the issuer
	// of the OAuth 2.0 service. Services shall verify the token provided in the 'Authorization' header of the request
	// with the value of this property. If the Mode property contains the value 'Discovery', this property shall
	// contain the keys found at the URI specified by the 'jwks_uri' string from the OAuth 2.0 service's metadata and
	// this property shall be read-only. Clients should configure this property if Mode contains 'Offline'.
	OAuthServiceSigningKeys string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// RoleMapping shall contain mapping rules that are used to convert the external account providers account
// information to the local Redfish role.
type RoleMapping struct {
	// LocalRole shall contain the RoleId property value within a role resource on this Redfish service to which to map
	// the remote user or group.
	LocalRole string
	// MFABypass shall contain the multi-factor authentication bypass settings.
	MFABypass MFABypass
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteGroup shall contain the name of the remote group, or the remote role in the case of a Redfish service,
	// that maps to the local Redfish role to which this entity links.
	RemoteGroup string
	// RemoteUser shall contain the name of the remote user that maps to the local Redfish role to which this entity
	// links.
	RemoteUser string
}

// TACACSplusService shall contain settings for parsing a TACACS+ service.
type TACACSplusService struct {
	// AuthorizationService shall contain the TACACS+ service authorization argument as defined by section 8.2 of
	// RFC8907. If this property is not present, the service defines the value to provide to the TACACS+ server.
	AuthorizationService string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PasswordExchangeProtocols shall indicate all the allowed TACACS+ password exchange protocol described under
	// section 5.4.2 of RFC8907.
	PasswordExchangeProtocols []TACACSplusPasswordExchangeProtocol
	// PrivilegeLevelArgument shall specify the name of the argument in a TACACS+ Authorization REPLY packet body, as
	// defined in RFC8907, that contains the user's privilege level.
	PrivilegeLevelArgument string
}
