//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/OutboundConnection.v1_0_2.json
// 2023.2 - #OutboundConnection.v1_0_2.OutboundConnection

package schemas

import (
	"encoding/json"
)

type AuthenticationType string

const (
	// MTLSAuthenticationType shall indicate the service will exchange and verify
	// certificates during TLS handshaking when establishing the outbound
	// connecting.
	MTLSAuthenticationType AuthenticationType = "MTLS"
	// JWTAuthenticationType shall indicate an RFC7519-defined JSON Web Token (JWT)
	// is specified in one of the HTTP headers in the 'PreUpgradeHTTPHeaders'
	// property. This is typically encoded in the 'Authorization' header with the
	// scheme 'Bearer'.
	JWTAuthenticationType AuthenticationType = "JWT"
	// NoneAuthenticationType shall indicate the service does not provide any
	// authentication information to the remote client.
	NoneAuthenticationType AuthenticationType = "None"
	// OEMAuthenticationType shall indicate an OEM-specific authentication
	// mechanism.
	OEMAuthenticationType AuthenticationType = "OEM"
)

type OutboundConnectionRetryPolicyType string

const (
	// NoneOutboundConnectionRetryPolicyType shall indicate the service will not
	// attempt to re-establish the outbound connection if the connection is dropped
	// or not established. If the connection is dropped or not established, the
	// service shall set the 'ConnectionEnabled' property to 'false'.
	NoneOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "None"
	// RetryForeverOutboundConnectionRetryPolicyType shall indicate the service
	// will attempt to re-establish the outbound connection at the interval
	// specified by the 'RetryIntervalMinutes' property regardless of the number of
	// retries.
	RetryForeverOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "RetryForever"
	// RetryCountOutboundConnectionRetryPolicyType shall indicate the service will
	// attempt to re-establish the outbound connection at the interval specified by
	// the 'RetryIntervalMinutes' property until the number of retries reaches the
	// count specified by the 'RetryCount' property. If the limit is reached, the
	// service shall set the 'ConnectionEnabled' property to 'false'. If a
	// connection is established, the service shall reset the count.
	RetryCountOutboundConnectionRetryPolicyType OutboundConnectionRetryPolicyType = "RetryCount"
)

// OutboundConnection shall represent the connection configuration necessary to
// connect to a remote client. Services shall initiate the outbound connection
// over a WebSocket defined in the 'Outbound connections' clause of the Redfish
// Specification.
type OutboundConnection struct {
	Entity
	// Authentication shall contain the authentication mechanism for the WebSocket
	// connection.
	Authentication AuthenticationType
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the server certificates for the
	// remote client referenced by the 'EndpointURI' property. If the
	// 'Authentication' property contains 'MTLS', the service shall compare the
	// certificates in this collection with the certificate obtained during
	// handshaking with the WebSocket service to verify the identity of the remote
	// client prior to completing the connection. If the remote client cannot be
	// verified, the service shall not complete the connection. Regardless of the
	// contents of this collection, services may perform additional verification
	// based on other factors, such as the configuration of the 'SecurityPolicy'
	// resource.
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates for
	// the service. If the 'Authentication' property contains 'MTLS', these
	// certificates are provided to the remote client referenced by the
	// 'EndpointURI' property as part of TLS handshaking.
	clientCertificates string
	// ConnectionEnabled shall indicate if the outbound connection is enabled. If
	// 'true', the service shall attempt to establish an outbound connection to the
	// remote client specified by the 'EndpointURI' property. If 'false', the
	// service shall not attempt to establish a connection to the remote client and
	// shall close the connection if one is already established. When a connection
	// is established, the service shall create a Session resource to represent the
	// active connection. When a connection is closed, the service shall delete the
	// connection's respective Session resource. If the client does not provide
	// this property, the service shall default this value to 'true'.
	ConnectionEnabled bool
	// EndpointURI shall contain the WebSocket URI to the external web service of
	// the remote client. The value shall follow the URI format defined in RFC6455.
	// Services shall reject URIs that do not contain the scheme 'wss'.
	EndpointURI string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PreUpgradeHTTPHeaders shall contain an object consisting of the names and
	// values of HTTP headers to send to the remote client during the initial
	// connection prior to the WebSocket upgrade. This property shall be an empty
	// object in responses.
	PreUpgradeHTTPHeaders HTTPHeaderProperty
	// RetryPolicy shall contain the retry policy for this outbound connection. If
	// not specified by the client in the create request, the service shall assume
	// ConnectionRetryPolicy contains 'None'.
	RetryPolicy RetryPolicyType
	// Roles shall contain the Redfish roles that contain the privileges of the
	// remote client for the outbound connection.
	Roles []string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// WebSocketPingIntervalMinutes shall contain the interval for the service to
	// send the WebSocket ping opcode to the remote client in minutes. If '0', the
	// service shall not send the WebSocket ping opcode to the remote client.
	WebSocketPingIntervalMinutes *uint `json:",omitempty"`
	// session is the URI for Session.
	session string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a OutboundConnection object from the raw JSON.
func (o *OutboundConnection) UnmarshalJSON(b []byte) error {
	type temp OutboundConnection
	type oLinks struct {
		Session Link `json:"Session"`
	}
	var tmp struct {
		temp
		Links              oLinks
		Certificates       Link `json:"Certificates"`
		ClientCertificates Link `json:"ClientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = OutboundConnection(tmp.temp)

	// Extract the links to other entities for later
	o.session = tmp.Links.Session.String()
	o.certificates = tmp.Certificates.String()
	o.clientCertificates = tmp.ClientCertificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	o.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (o *OutboundConnection) Update() error {
	readWriteFields := []string{
		"ConnectionEnabled",
		"WebSocketPingIntervalMinutes",
	}

	return o.UpdateFromRawData(o, o.RawData, readWriteFields)
}

// GetOutboundConnection will get a OutboundConnection instance from the service.
func GetOutboundConnection(c Client, uri string) (*OutboundConnection, error) {
	return GetObject[OutboundConnection](c, uri)
}

// ListReferencedOutboundConnections gets the collection of OutboundConnection from
// a provided reference.
func ListReferencedOutboundConnections(c Client, link string) ([]*OutboundConnection, error) {
	return GetCollectionObjects[OutboundConnection](c, link)
}

// Session gets the Session linked resource.
func (o *OutboundConnection) Session() (*Session, error) {
	if o.session == "" {
		return nil, nil
	}
	return GetObject[Session](o.client, o.session)
}

// Certificates gets the Certificates collection.
func (o *OutboundConnection) Certificates() ([]*Certificate, error) {
	if o.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](o.client, o.certificates)
}

// ClientCertificates gets the ClientCertificates collection.
func (o *OutboundConnection) ClientCertificates() ([]*Certificate, error) {
	if o.clientCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](o.client, o.clientCertificates)
}

// RetryPolicyType shall contain the retry policy for an outbound connection.
type RetryPolicyType struct {
	// ConnectionRetryPolicy shall contain the type of retry policy for this
	// outbound connection.
	ConnectionRetryPolicy OutboundConnectionRetryPolicyType
	// RetryCount shall contain the number of retries to attempt if the retry
	// policy specifies a maximum number of retries.
	RetryCount *int `json:",omitempty"`
	// RetryIntervalMinutes shall contain the interval for the service to retry
	// connecting to remote client in minutes.
	RetryIntervalMinutes *int `json:",omitempty"`
}
