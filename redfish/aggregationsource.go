//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type AggregationType string

const (
	// NotificationsOnlyAggregationType shall indicate that the aggregator is only aggregating notifications or events
	// from the aggregation source according to the connection method used. This value shall not be used with
	// connection methods that do not include notifications.
	NotificationsOnlyAggregationType AggregationType = "NotificationsOnly"
	// FullAggregationType shall indicate that the aggregator is performing full aggregation according to the
	// connection method without any limitation such as only receiving notifications.
	FullAggregationType AggregationType = "Full"
)

type SNMPAuthenticationProtocols string

const (
	// NoneSNMPAuthenticationProtocols shall indicate authentication is not required.
	NoneSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "None"
	// CommunityStringSNMPAuthenticationProtocols shall indicate authentication using SNMP community strings and the
	// value of TrapCommunity.
	CommunityStringSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "CommunityString"
	// HMACMD5SNMPAuthenticationProtocols shall indicate authentication conforms to the RFC3414-defined HMAC-MD5-96
	// authentication protocol.
	HMACMD5SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_MD5"
	// HMACSHA96SNMPAuthenticationProtocols shall indicate authentication conforms to the RFC3414-defined HMAC-SHA-96
	// authentication protocol.
	HMACSHA96SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_SHA96"
	// HMAC128SHA224SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC128SHA224AuthProtocol.
	HMAC128SHA224SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC128_SHA224"
	// HMAC192SHA256SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC192SHA256AuthProtocol.
	HMAC192SHA256SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC192_SHA256"
	// HMAC256SHA384SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC256SHA384AuthProtocol.
	HMAC256SHA384SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC256_SHA384"
	// HMAC384SHA512SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC384SHA512AuthProtocol.
	HMAC384SHA512SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC384_SHA512"
)

type SNMPEncryptionProtocols string

const (
	// NoneSNMPEncryptionProtocols shall indicate there is no encryption.
	NoneSNMPEncryptionProtocols SNMPEncryptionProtocols = "None"
	// CBCDESSNMPEncryptionProtocols shall indicate encryption conforms to the RFC3414-defined CBC-DES encryption
	// protocol.
	CBCDESSNMPEncryptionProtocols SNMPEncryptionProtocols = "CBC_DES"
	// CFB128AES128SNMPEncryptionProtocols shall indicate encryption conforms to the RFC3826-defined CFB128-AES-128
	// encryption protocol.
	CFB128AES128SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES128"
	// CFB128AES192SNMPEncryptionProtocols shall indicate encryption conforms to the CFB128-AES-192 encryption
	// protocol, extended from RFC3826.
	CFB128AES192SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES192"
	// CFB128AES256SNMPEncryptionProtocols shall indicate encryption conforms to the CFB128-AES-256 encryption
	// protocol, extended from RFC3826.
	CFB128AES256SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES256"
)

type UserAuthenticationMethod string

const (
	// PublicKeyUserAuthenticationMethod shall indicate SSH user authentication with a public key specified by the
	// PublicIdentityKey property in SSHSettings.
	PublicKeyUserAuthenticationMethod UserAuthenticationMethod = "PublicKey"
	// PasswordUserAuthenticationMethod shall indicate SSH user authentication with a password specified by the
	// Password property.
	PasswordUserAuthenticationMethod UserAuthenticationMethod = "Password"
)

// AggregationSource shall represent an aggregation source for a Redfish implementation.
type AggregationSource struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AggregationType shall contain the type of aggregation used for the connection method towards the aggregation
	// source. If this property is not present, the value shall be assumed to be 'Full'.
	AggregationType AggregationType
	// Description provides a description of this resource.
	Description string
	// HostName shall contain the URI of the system to be aggregated. This property shall not be required when the
	// aggregation source is configured to only receive notifications from the aggregated system and the
	// AggregationType property contains the value 'NotificationsOnly'.
	HostName string
	// Password shall contain a password for accessing the aggregation source. The value shall be 'null' in responses.
	Password string
	// SNMP shall contain the SNMP settings of the aggregation source.
	SNMP SNMPSettings
	// SSHSettings shall contain the settings for an aggregation source using SSH as part of the associated connection
	// method.
	SSHSettings SSHSettingsType
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserName shall contain the user name for accessing the aggregation source.
	UserName string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	generateSSHIdentityKeyPairTarget string
	removeSSHIdentityKeyPairTarget   string

	connectionMethod  string
	resourcesAccessed []string
	// ResourcesAccessedCount is the number of the resources added to the service through the aggregation source.
	ResourcesAccessedCount int
}

// UnmarshalJSON unmarshals a AggregationSource object from the raw JSON.
func (aggregationsource *AggregationSource) UnmarshalJSON(b []byte) error {
	type temp AggregationSource
	type Actions struct {
		GenerateSSHIdentityKeyPair common.ActionTarget `json:"#AggregationSource.GenerateSSHIdentityKeyPair"`
		RemoveSSHIdentityKeyPair   common.ActionTarget `json:"#AggregationSource.RemoveSSHIdentityKeyPair"`
	}
	type Links struct {
		// ConnectionMethod shall contain a link to resources of type ConnectionMethod that are used to connect
		// to the aggregation source.
		ConnectionMethod common.Link
		// ResourcesAccessed shall contain an array of links to the resources added to the service through the aggregation
		// source. It is recommended that this be the minimal number of properties needed to find the resources that would
		// be lost when the aggregation source is deleted. For example, this could be the pointers to the members of the
		// root-level collections or the manager of a BMC.
		ResourcesAccessed common.Links
		// ResourcesAccessed@odata.count
		ResourcesAccessedCount int `json:"ResourcesAccessed@odata.count"`
	}
	var t struct {
		temp
		Actions Actions
		Links   Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregationsource = AggregationSource(t.temp)
	aggregationsource.generateSSHIdentityKeyPairTarget = t.Actions.GenerateSSHIdentityKeyPair.Target
	aggregationsource.removeSSHIdentityKeyPairTarget = t.Actions.RemoveSSHIdentityKeyPair.Target

	aggregationsource.connectionMethod = t.Links.ConnectionMethod.String()
	aggregationsource.resourcesAccessed = t.Links.ResourcesAccessed.ToStrings()
	aggregationsource.ResourcesAccessedCount = t.Links.ResourcesAccessedCount

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	aggregationsource.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (aggregationsource *AggregationSource) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AggregationSource)
	original.UnmarshalJSON(aggregationsource.rawData)

	readWriteFields := []string{
		"AggregationType",
		"HostName",
		"Password",
		"UserName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(aggregationsource).Elem()

	return aggregationsource.Entity.Update(originalElement, currentElement, readWriteFields)
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

// SNMPSettings shall contain the settings for an SNMP aggregation source.
type SNMPSettings struct {
	// AuthenticationKey shall contain the key for SNMPv3 authentication. The value shall be 'null' in responses. This
	// property accepts a passphrase or a hex-encoded key. If the string starts with 'Passphrase:', the remainder of
	// the string shall be the passphrase and shall be converted to the key as described in the 'Password to Key
	// Algorithm' section of RFC3414. If the string starts with 'Hex:', then the remainder of the string shall be the
	// key encoded in hexadecimal notation. If the string starts with neither, the full string shall be a passphrase
	// and shall be converted to the key as described in the 'Password to Key Algorithm' section of RFC3414. The
	// passphrase can contain any printable characters except for the double quotation mark.
	AuthenticationKey string
	// AuthenticationKeySet shall contain 'true' if a valid value was provided for the AuthenticationKey property.
	// Otherwise, the property shall contain 'false'.
	AuthenticationKeySet string
	// AuthenticationProtocol shall contain the SNMPv3 authentication protocol.
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey shall contain the key for SNMPv3 encryption. The value shall be 'null' in responses. This property
	// accepts a passphrase or a hex-encoded key. If the string starts with 'Passphrase:', the remainder of the string
	// shall be the passphrase and shall be converted to the key as described in the 'Password to Key Algorithm'
	// section of RFC3414. If the string starts with 'Hex:', then the remainder of the string shall be the key encoded
	// in hexadecimal notation. If the string starts with neither, the full string shall be a passphrase and shall be
	// converted to the key as described in the 'Password to Key Algorithm' section of RFC3414. The passphrase can
	// contain any printable characters except for the double quotation mark.
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the EncryptionKey property. Otherwise,
	// the property shall contain 'false'.
	EncryptionKeySet string
	// EncryptionProtocol shall contain the SNMPv3 encryption protocol.
	EncryptionProtocol SNMPEncryptionProtocols
	// TrapCommunity shall contain the SNMP trap community string. The value shall be 'null' in responses. Services may
	// provide a common trap community if not specified by the client when creating the aggregation source.
	TrapCommunity string
}

// SSHSettingsType shall contain the settings for an aggregation source using SSH as part of the associated
// connection method.
type SSHSettingsType struct {
	// PresentedPublicHostKey shall contain a link to a resource of type Key that represents the last public host key
	// presented by the remote service corresponding to the aggregation source. This property shall not be present if a
	// public host key has not yet been presented by the remote service.
	PresentedPublicHostKey string
	// PresentedPublicHostKeyTimestamp shall contain the date and time when the key referenced by the
	// PresentedPublicHostKey property was last updated.
	PresentedPublicHostKeyTimestamp string
	// PublicIdentityKey shall contain a link to a resource of type Key that represents the public key that is used
	// with the aggregation source when UserAuthenticationMethod contains 'PublicKey'. This property shall not be
	// present if a key-pair is not available. The State property within Status shall contain 'Disabled' if a key-pair
	// is not available and UserAuthenticationMethod contains 'PublicKey'.
	PublicIdentityKey string
	// TrustedPublicHostKeys shall contain a link to a resource collection of type KeyCollection that represents the
	// trusted public host keys of the remote service corresponding to the aggregation source. If the associated
	// connection method specifies SSH tunneling, the service shall compare the public host key presented by the remote
	// service with members of this collection to determine if the remote service can be trusted. If the remote service
	// cannot be trusted, the State property within Status shall contain 'Disabled' and the service shall not connect
	// to the remote service.
	TrustedPublicHostKeys string
	// UserAuthenticationMethod shall contain the client user authentication method.
	UserAuthenticationMethod UserAuthenticationMethod
}
