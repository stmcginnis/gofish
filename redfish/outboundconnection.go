//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type AuthenticationType string

const (
	// MTLSAuthenticationType shall indicate the service will exchange and verify certificates during TLS handshaking
	// when establishing the outbound connecting.
	MTLSAuthenticationType AuthenticationType = "MTLS"
	// JWTAuthenticationType shall indicate an RFC7519-defined JSON Web Token (JWT) is specified in one of the HTTP
	// headers in the PreUpgradeHTTPHeaders property. This is typically encoded in the 'Authorization' header with the
	// scheme 'Bearer'.
	JWTAuthenticationType AuthenticationType = "JWT"
	// NoneAuthenticationType shall indicate the service does not provide any authentication information to the remote
	// client.
	NoneAuthenticationType AuthenticationType = "None"
	// OEMAuthenticationType shall indicate an OEM-specific authentication mechanism.
	OEMAuthenticationType AuthenticationType = "OEM"
)

type OutboundConnectionRetryPolicyType string

const (
	// NoneOutboundConnectionRetryPolicyType shall indicate the service will not attempt to re-establish the outbound
	// connection if the connection is dropped or not established. If the connection is dropped or not established, the
	// service shall set the ConnectionEnabled property to 'false'.
	NoneOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "None"
	// RetryForeverOutboundConnectionRetryPolicyType shall indicate the service will attempt to re-establish the
	// outbound connection at the interval specified by the RetryIntervalMinutes property regardless of the number of
	// retries.
	RetryForeverOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "RetryForever"
	// RetryCountOutboundConnectionRetryPolicyType shall indicate the service will attempt to re-establish the outbound
	// connection at the interval specified by the RetryIntervalMinutes property until the number of retries reaches
	// the count specified by the RetryCount property. If the limit is reached, the service shall set the
	// ConnectionEnabled property to 'false'. If a connection is established, the service shall reset the count.
	RetryCountOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "RetryCount"
)

// OutboundConnection shall represent the connection configuration necessary to connect to a remote client.
// Services shall initiate the outbound connection over a WebSocket defined in the 'Outbound connections' clause of
// the Redfish Specification.
type OutboundConnection struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Authentication shall contain the authentication mechanism for the WebSocket connection.
	Authentication AuthenticationType
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represents the
	// server certificates for the remote client referenced by the EndpointURI property. If the Authentication property
	// contains 'MTLS', the service shall compare the certificates in this collection with the certificate obtained
	// during handshaking with the WebSocket service to verify the identity of the remote client prior to completing
	// the connection. If the remote client cannot be verified, the service shall not complete the connection.
	// Regardless of the contents of this collection, services may perform additional verification based on other
	// factors, such as the configuration of the SecurityPolicy resource.
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates for the service. If the Authentication property contains 'MTLS', these
	// certificates are provided to the remote client referenced by the EndpointURI property as part of TLS
	// handshaking.
	clientCertificates string
	// ConnectionEnabled shall indicate if the outbound connection is enabled. If 'true', the service shall attempt to
	// establish an outbound connection to the remote client specified by the EndpointURI property. If 'false', the
	// service shall not attempt to establish a connection to the remote client and shall close the connection if one
	// is already established. When a connection is established, the service shall create a Session resource to
	// represent the active connection. When a connection is closed, the service shall delete the connection's
	// respective Session resource. If the client does not provide this property, the service shall default this value
	// to 'true'.
	ConnectionEnabled bool
	// Description provides a description of this resource.
	Description string
	// EndpointURI shall contain the WebSocket URI to the external web service of the remote client. The value shall
	// follow the URI format defined in RFC6455. Services shall reject URIs that do not contain the scheme 'wss'.
	EndpointURI string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PreUpgradeHTTPHeaders shall contain an object consisting of the names and values of HTTP headers to send to the
	// remote client during the initial connection prior to the WebSocket upgrade. This property shall be an empty
	// object in responses.
	PreUpgradeHTTPHeaders map[string]string
	// RetryPolicy shall contain the retry policy for this outbound connection. If not specified by the client in the
	// create request, the service shall assume ConnectionRetryPolicy contains 'None'.
	RetryPolicy RetryPolicyType
	// Roles shall contain the Redfish roles that contain the privileges of the remote client for the outbound
	// connection.
	Roles []string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// WebSocketPingIntervalMinutes shall contain the interval for the service to send the WebSocket ping opcode to the
	// remote client in minutes. If '0', the service shall not send the WebSocket ping opcode to the remote client.
	WebSocketPingIntervalMinutes int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	session string
}

// UnmarshalJSON unmarshals a OutboundConnection object from the raw JSON.
func (outboundconnection *OutboundConnection) UnmarshalJSON(b []byte) error {
	type temp OutboundConnection
	type Links struct {
		// Session shall contain the link to a resource of type Session that represents the active connection for this
		// outbound connection.
		Session common.Link
	}
	var t struct {
		temp
		Links              Links
		Certificates       common.Link
		ClientCertificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*outboundconnection = OutboundConnection(t.temp)

	// Extract the links to other entities for later
	outboundconnection.session = t.Links.Session.String()

	outboundconnection.certificates = t.Certificates.String()
	outboundconnection.clientCertificates = t.ClientCertificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	outboundconnection.rawData = b

	return nil
}

// Session gets the the active connection for this outbound connection.
func (outboundconnection *OutboundConnection) Session() (*Session, error) {
	if outboundconnection.session == "" {
		return nil, nil
	}
	return GetSession(outboundconnection.GetClient(), outboundconnection.session)
}

// Certificates gets the server certificates for the remote client referenced by the EndpointURI property.
func (outboundconnection *OutboundConnection) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(outboundconnection.GetClient(), outboundconnection.certificates)
}

// ClientCertificates gets the client identity certificates for the service.
func (outboundconnection *OutboundConnection) ClientCertificates() ([]*Certificate, error) {
	return ListReferencedCertificates(outboundconnection.GetClient(), outboundconnection.clientCertificates)
}

// Update commits updates to this object's properties to the running system.
func (outboundconnection *OutboundConnection) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(OutboundConnection)
	original.UnmarshalJSON(outboundconnection.rawData)

	readWriteFields := []string{
		"ConnectionEnabled",
		"WebSocketPingIntervalMinutes",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(outboundconnection).Elem()

	return outboundconnection.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetOutboundConnection will get a OutboundConnection instance from the service.
func GetOutboundConnection(c common.Client, uri string) (*OutboundConnection, error) {
	return common.GetObject[OutboundConnection](c, uri)
}

// ListReferencedOutboundConnections gets the collection of OutboundConnection from
// a provided reference.
func ListReferencedOutboundConnections(c common.Client, link string) ([]*OutboundConnection, error) {
	return common.GetCollectionObjects[OutboundConnection](c, link)
}

// RetryPolicyType shall contain the retry policy for an outbound connection.
type RetryPolicyType struct {
	// ConnectionRetryPolicy shall contain the type of retry policy for this outbound connection.
	ConnectionRetryPolicy OutboundConnectionRetryPolicyType
	// RetryCount shall contain the number of retries to attempt if the retry policy specifies a maximum number of
	// retries.
	RetryCount int
	// RetryIntervalMinutes shall contain the interval for the service to retry connecting to remote client in minutes.
	RetryIntervalMinutes int
}

// UnmarshalJSON unmarshals a RetryPolicyType object from the raw JSON.
func (retrypolicytype *RetryPolicyType) UnmarshalJSON(b []byte) error {
	type temp RetryPolicyType
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*retrypolicytype = RetryPolicyType(t.temp)

	// Extract the links to other entities for later

	return nil
}
