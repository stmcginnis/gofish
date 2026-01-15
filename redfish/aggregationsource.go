//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #AggregationSource.v1_5_0.AggregationSource

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type AggregationType string

const (
	// NotificationsOnlyAggregationType shall indicate that the aggregator is only
	// aggregating notifications or events from the aggregation source according to
	// the connection method used. This value shall not be used with connection
	// methods that do not include notifications.
	NotificationsOnlyAggregationType AggregationType = "NotificationsOnly"
	// FullAggregationType shall indicate that the aggregator is performing full
	// aggregation according to the connection method without any limitation such
	// as only receiving notifications.
	FullAggregationType AggregationType = "Full"
)

type UserAuthenticationMethod string

const (
	// PublicKeyUserAuthenticationMethod shall indicate SSH user authentication
	// with a public key specified by the 'PublicIdentityKey' property in
	// 'SSHSettings'.
	PublicKeyUserAuthenticationMethod UserAuthenticationMethod = "PublicKey"
	// PasswordUserAuthenticationMethod shall indicate SSH user authentication with
	// a password specified by the 'Password' property.
	PasswordUserAuthenticationMethod UserAuthenticationMethod = "Password"
)

// AggregationSource shall represent an aggregation source for a Redfish
// implementation.
type AggregationSource struct {
	common.Entity
	// AggregationType shall contain the type of aggregation used for the
	// connection method towards the aggregation source. If this property is not
	// present, the value shall be assumed to be 'Full'.
	//
	// Version added: v1.2.0
	AggregationType AggregationType
	// HostName shall contain the URI of the system to be aggregated. This property
	// shall not be required when the aggregation source is configured to only
	// receive notifications from the aggregated system and the 'AggregationType'
	// property contains the value 'NotificationsOnly'.
	HostName string
	// ModbusTargetServerId shall contain the Modbus identifier of this aggregation
	// source.
	//
	// Version added: v1.5.0
	ModbusTargetServerID *int `json:"ModbusTargetServerId,omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain a password for accessing the aggregation source. The
	// value shall be 'null' in responses.
	Password string
	// Port shall contain the network port used to connect to this aggregation
	// source.
	//
	// Version added: v1.5.0
	Port *uint `json:",omitempty"`
	// SNMP shall contain the SNMP settings of the aggregation source.
	//
	// Version added: v1.1.0
	SNMP SNMPSettings
	// SSHSettings shall contain the settings for an aggregation source using SSH
	// as part of the associated connection method.
	//
	// Version added: v1.3.0
	SSHSettings SSHSettingsType
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.3.0
	Status common.Status
	// UserName shall contain the username for accessing the aggregation source.
	UserName string
	// generateSSHIdentityKeyPairTarget is the URL to send GenerateSSHIdentityKeyPair requests.
	generateSSHIdentityKeyPairTarget string
	// removeSSHIdentityKeyPairTarget is the URL to send RemoveSSHIdentityKeyPair requests.
	removeSSHIdentityKeyPairTarget string
	// connectionMethod is the URI for ConnectionMethod.
	connectionMethod string
	// resourcesAccessed are the URIs for ResourcesAccessed.
	resourcesAccessed []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AggregationSource object from the raw JSON.
func (a *AggregationSource) UnmarshalJSON(b []byte) error {
	type temp AggregationSource
	type aActions struct {
		GenerateSSHIdentityKeyPair common.ActionTarget `json:"#AggregationSource.GenerateSSHIdentityKeyPair"`
		RemoveSSHIdentityKeyPair   common.ActionTarget `json:"#AggregationSource.RemoveSSHIdentityKeyPair"`
	}
	type aLinks struct {
		ConnectionMethod  common.Link  `json:"ConnectionMethod"`
		ResourcesAccessed common.Links `json:"ResourcesAccessed"`
	}
	var tmp struct {
		temp
		Actions aActions
		Links   aLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AggregationSource(tmp.temp)

	// Extract the links to other entities for later
	a.generateSSHIdentityKeyPairTarget = tmp.Actions.GenerateSSHIdentityKeyPair.Target
	a.removeSSHIdentityKeyPairTarget = tmp.Actions.RemoveSSHIdentityKeyPair.Target
	a.connectionMethod = tmp.Links.ConnectionMethod.String()
	a.resourcesAccessed = tmp.Links.ResourcesAccessed.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AggregationSource) Update() error {
	readWriteFields := []string{
		"AggregationType",
		"HostName",
		"ModbusTargetServerId",
		"Password",
		"Port",
		"SNMP",
		"SSHSettings",
		"Status",
		"UserName",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetAggregationSource will get a AggregationSource instance from the service.
func GetAggregationSource(c common.Client, uri string) (*AggregationSource, error) {
	return common.GetObject[AggregationSource](c, uri)
}

// ListReferencedAggregationSources gets the collection of AggregationSource from
// a provided reference.
func ListReferencedAggregationSources(c common.Client, link string) ([]*AggregationSource, error) {
	return common.GetCollectionObjects[AggregationSource](c, link)
}

// GenerateSSHIdentityKeyPair shall generate a new SSH identity key-pair to be used with this
// aggregation source. The service shall store the generated public key in the
// 'Key' resource referenced by the 'PublicIdentityKey' property in
// 'SSHSettings'. If the aggregation source already has an associated SSH
// identity key-pair, the service shall delete the key-pair and replace it with
// the new key-pair.
// curve - This parameter shall contain the curve to use with the SSH key. This
// parameter shall be required if the 'KeyType' parameter contains 'ECDSA' and
// shall be rejected for other values.
// keyLength - This parameter shall contain the length of the SSH key, in bits.
// This parameter shall be required if the 'KeyType' parameter contains 'RSA'
// and shall be rejected for other values.
// keyType - This parameter shall contain the type of SSH key.
func (a *AggregationSource) GenerateSSHIdentityKeyPair(curve ECDSACurveType, keyLength int, keyType KeyType) error {
	payload := make(map[string]any)
	payload["Curve"] = curve
	payload["KeyLength"] = keyLength
	payload["KeyType"] = keyType
	return a.Post(a.generateSSHIdentityKeyPairTarget, payload)
}

// RemoveSSHIdentityKeyPair shall remove the private SSH identity key-pair used with this
// aggregation source.
func (a *AggregationSource) RemoveSSHIdentityKeyPair() error {
	payload := make(map[string]any)
	return a.Post(a.removeSSHIdentityKeyPairTarget, payload)
}

// ConnectionMethod gets the ConnectionMethod linked resource.
func (a *AggregationSource) ConnectionMethod(client common.Client) (*ConnectionMethod, error) {
	if a.connectionMethod == "" {
		return nil, nil
	}
	return common.GetObject[ConnectionMethod](client, a.connectionMethod)
}

// ResourcesAccessed gets the ResourcesAccessed linked resources.
func (a *AggregationSource) ResourcesAccessed(client common.Client) ([]*common.Resource, error) {
	return common.GetObjects[common.Resource](client, a.resourcesAccessed)
}

// SSHSettingsType shall contain the settings for an aggregation source using
// SSH as part of the associated connection method.
type SSHSettingsType struct {
	// PresentedPublicHostKey shall contain a link to a resource of type 'Key' that
	// represents the last public host key presented by the remote service
	// corresponding to the aggregation source. This property shall not be present
	// if a public host key has not yet been presented by the remote service.
	//
	// Version added: v1.3.0
	presentedPublicHostKey string
	// PresentedPublicHostKeyTimestamp shall contain the date and time when the key
	// referenced by the 'PresentedPublicHostKey' property was last updated.
	//
	// Version added: v1.3.0
	PresentedPublicHostKeyTimestamp string
	// PublicIdentityKey shall contain a link to a resource of type 'Key' that
	// represents the public key that is used with the aggregation source when
	// 'UserAuthenticationMethod' contains 'PublicKey'. This property shall not be
	// present if a key-pair is not available. The 'State' property within 'Status'
	// shall contain 'Disabled' if a key-pair is not available and
	// 'UserAuthenticationMethod' contains 'PublicKey'.
	//
	// Version added: v1.3.0
	publicIdentityKey string
	// TrustedPublicHostKeys shall contain a link to a resource collection of type
	// 'KeyCollection' that represents the trusted public host keys of the remote
	// service corresponding to the aggregation source. If the associated
	// connection method specifies SSH tunneling, the service shall compare the
	// public host key presented by the remote service with members of this
	// collection to determine if the remote service can be trusted. If the remote
	// service cannot be trusted, the 'State' property within 'Status' shall
	// contain 'Disabled' and the service shall not connect to the remote service.
	//
	// Version added: v1.3.0
	trustedPublicHostKeys string
	// UserAuthenticationMethod shall contain the client user authentication
	// method.
	//
	// Version added: v1.3.0
	UserAuthenticationMethod UserAuthenticationMethod
}

// UnmarshalJSON unmarshals a SSHSettingsType object from the raw JSON.
func (s *SSHSettingsType) UnmarshalJSON(b []byte) error {
	type temp SSHSettingsType
	var tmp struct {
		temp
		PresentedPublicHostKey common.Link `json:"presentedPublicHostKey"`
		PublicIdentityKey      common.Link `json:"publicIdentityKey"`
		TrustedPublicHostKeys  common.Link `json:"trustedPublicHostKeys"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SSHSettingsType(tmp.temp)

	// Extract the links to other entities for later
	s.presentedPublicHostKey = tmp.PresentedPublicHostKey.String()
	s.publicIdentityKey = tmp.PublicIdentityKey.String()
	s.trustedPublicHostKeys = tmp.TrustedPublicHostKeys.String()

	return nil
}

// PresentedPublicHostKey gets the PresentedPublicHostKey linked resource.
func (s *SSHSettingsType) PresentedPublicHostKey(client common.Client) (*Key, error) {
	if s.presentedPublicHostKey == "" {
		return nil, nil
	}
	return common.GetObject[Key](client, s.presentedPublicHostKey)
}

// PublicIdentityKey gets the PublicIdentityKey linked resource.
func (s *SSHSettingsType) PublicIdentityKey(client common.Client) (*Key, error) {
	if s.publicIdentityKey == "" {
		return nil, nil
	}
	return common.GetObject[Key](client, s.publicIdentityKey)
}

// TrustedPublicHostKeys gets the TrustedPublicHostKeys collection.
func (s *SSHSettingsType) TrustedPublicHostKeys(client common.Client) ([]*Key, error) {
	if s.trustedPublicHostKeys == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Key](client, s.trustedPublicHostKeys)
}
