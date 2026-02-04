//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ExternalAccountProvider.v1_8_2.json
// 2024.3 - #ExternalAccountProvider.v1_8_2.ExternalAccountProvider

package schemas

import (
	"encoding/json"
)

// ExternalAccountProvider shall represent a remote authentication service in
// the Redfish Specification.
type ExternalAccountProvider struct {
	Entity
	// AccountProviderType shall contain the type of external account provider to
	// which this service connects.
	AccountProviderType AccountProviderTypes
	// Authentication shall contain the authentication information for the external
	// account provider.
	Authentication Authentication
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates the external account
	// provider uses.
	//
	// Version added: v1.1.0
	certificates string
	// LDAPService shall contain any additional mapping information needed to parse
	// a generic LDAP service. This property should only be present if
	// 'AccountProviderType' is 'LDAPService'.
	LDAPService LDAPService
	// OAuth2Service shall contain additional information needed to parse an OAuth
	// 2.0 service. This property should only be present inside an 'OAuth2'
	// property.
	//
	// Version added: v1.3.0
	OAuth2Service OAuth2Service
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Priority shall contain the assigned priority for the specified external
	// account provider. The value '0' shall indicate the highest priority.
	// Increasing values shall represent decreasing priority. If an external
	// provider does not have a priority assignment or two or more external
	// providers have the same priority, the behavior shall be determined by the
	// Redfish service. The priority is used to determine the order of
	// authentication and authorization for each external account provider.
	//
	// Version added: v1.2.0
	Priority *uint `json:",omitempty"`
	// RemoteRoleMapping shall contain a set of the mapping rules that are used to
	// convert the external account providers account information to the local
	// Redfish role. The service shall return the HTTP '401 Unauthorized' status
	// code to requests from accounts that do not map to a Redfish role in this
	// property.
	RemoteRoleMapping []RoleMapping
	// Retries shall contain the number of retries to attempt a connection to an
	// address in the 'ServiceAddresses' property before attempting a connection to
	// the next address in the array or giving up. If this property is not present,
	// the service has internal policies for handling retries.
	//
	// Version added: v1.6.0
	Retries *int `json:",omitempty"`
	// ServiceAddresses shall contain the addresses of the account providers to
	// which this external account provider links. The format of this field depends
	// on the type of external account provider. Each item in the array shall
	// contain a single address. Services can define their own behavior for
	// managing multiple addresses.
	ServiceAddresses []string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// TACACSplusService shall contain additional information needed to parse a
	// TACACS+ services. This property should only be present inside a 'TACACSplus'
	// property.
	//
	// Version added: v1.2.0
	TACACSplusService TACACSplusService
	// TimeoutSeconds shall contain the period of time, in seconds, this account
	// service will wait for a response from an address of a user account provider
	// before timing out. If this property is not present, the service has internal
	// policies for handling timeouts.
	//
	// Version added: v1.6.0
	TimeoutSeconds *int `json:",omitempty"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ExternalAccountProvider object from the raw JSON.
func (e *ExternalAccountProvider) UnmarshalJSON(b []byte) error {
	type temp ExternalAccountProvider
	var tmp struct {
		temp
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = ExternalAccountProvider(tmp.temp)

	// Extract the links to other entities for later
	e.certificates = tmp.Certificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	e.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (e *ExternalAccountProvider) Update() error {
	readWriteFields := []string{
		"Priority",
		"Retries",
		"ServiceAddresses",
		"ServiceEnabled",
		"TimeoutSeconds",
	}

	return e.UpdateFromRawData(e, e.RawData, readWriteFields)
}

// GetExternalAccountProvider will get a ExternalAccountProvider instance from the service.
func GetExternalAccountProvider(c Client, uri string) (*ExternalAccountProvider, error) {
	return GetObject[ExternalAccountProvider](c, uri)
}

// ListReferencedExternalAccountProviders gets the collection of ExternalAccountProvider from
// a provided reference.
func ListReferencedExternalAccountProviders(c Client, link string) ([]*ExternalAccountProvider, error) {
	return GetCollectionObjects[ExternalAccountProvider](c, link)
}

// Certificates gets the Certificates collection.
func (e *ExternalAccountProvider) Certificates() ([]*Certificate, error) {
	if e.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](e.client, e.certificates)
}
