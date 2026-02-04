//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Session.v1_8_0.json
// 2024.4 - #Session.v1_8_0.Session

package schemas

import (
	"encoding/json"
	"net/url"
)

type SessionTypes string

const (
	// HostConsoleSessionTypes shall indicate the session is the host's console,
	// which could be connected through Telnet, SSH, or another protocol. If this
	// session is terminated or deleted, the service shall close the connection for
	// the respective host console session.
	HostConsoleSessionTypes SessionTypes = "HostConsole"
	// ManagerConsoleSessionTypes shall indicate the session is the manager's
	// console, which could be connected through Telnet, SSH, SM CLP, or another
	// protocol. If this session is terminated or deleted, the service shall close
	// the connection for the respective manager console session.
	ManagerConsoleSessionTypes SessionTypes = "ManagerConsole"
	// IPMISessionTypes shall indicate the session is an Intelligent Platform
	// Management Interface session. If this session is terminated or deleted, the
	// service shall close the connection for the respective IPMI session.
	IPMISessionTypes SessionTypes = "IPMI"
	// KVMIPSessionTypes shall indicate the session is a Keyboard-Video-Mouse over
	// IP session. If this session is terminated or deleted, the service shall
	// close the connection for the respective KVM-IP session.
	KVMIPSessionTypes SessionTypes = "KVMIP"
	// OEMSessionTypes shall indicate the session is an OEM-specific session and is
	// further described by the 'OemSessionType' property.
	OEMSessionTypes SessionTypes = "OEM"
	// RedfishSessionTypes shall indicate the session is a Redfish session defined
	// by the 'Redfish session login authentication' clause of the Redfish
	// Specification. If this session is terminated or deleted, the service shall
	// invalidate the respective session token.
	RedfishSessionTypes SessionTypes = "Redfish"
	// VirtualMediaSessionTypes shall indicate the session is a virtual media
	// session. If this session is terminated or deleted, the service shall close
	// the connection for the respective virtual media session and make the media
	// inaccessible to the host.
	VirtualMediaSessionTypes SessionTypes = "VirtualMedia"
	// WebUISessionTypes shall indicate the session is a non-Redfish web user
	// interface session. If this session is terminated or deleted, the service
	// shall invalidate the respective session token.
	WebUISessionTypes SessionTypes = "WebUI"
	// OutboundConnectionSessionTypes shall indicate the session is an outbound
	// connection defined by the 'Outbound connections' clause of the Redfish
	// Specification. The 'OutboundConnection' property inside the 'Links' property
	// shall contain the link to the outbound connection configuration. If this
	// session is terminated or deleted, the service shall disable the associated
	// 'OutboundConnection' resource.
	OutboundConnectionSessionTypes SessionTypes = "OutboundConnection"
)

// Session shall represent a session for a Redfish implementation.
type Session struct {
	Entity
	// ClientOriginIPAddress shall contain the IP address of the client that
	// created the session.
	//
	// Version added: v1.3.0
	ClientOriginIPAddress string
	// Context shall contain a client-supplied context that remains with the
	// session through the session's lifetime.
	//
	// Version added: v1.5.0
	Context string
	// CreatedTime shall contain the date and time when the session was created.
	//
	// Version added: v1.4.0
	CreatedTime string
	// ExpirationTime shall contain the date and time when the session expires
	// regardless of session activity. The service shall delete this resource when
	// the expiration time is reached. If this property is not present, the session
	// does not expire based on an absolute time.
	//
	// Version added: v1.8.0
	ExpirationTime string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemSessionType shall contain the OEM-specific session type that is currently
	// active if 'SessionType' contains 'OEM'.
	//
	// Version added: v1.2.0
	OemSessionType string
	// Password shall contain the password for this session. The value shall be
	// 'null' in responses. When creating a session through a Redfish host
	// interface using an 'AuthNone' role, the property shall contain an empty
	// string in the request body.
	Password string
	// Roles shall contain the Redfish roles that contain the privileges of this
	// session.
	//
	// Version added: v1.7.0
	Roles []string
	// SessionType shall represent the type of session that is currently active.
	//
	// Version added: v1.2.0
	SessionType SessionTypes
	// Token shall contain the multi-factor authentication token for this session.
	// The value shall be 'null' in responses.
	//
	// Version added: v1.6.0
	Token string
	// UserName shall contain the username that matches an account recognized by
	// the account service. When a creating a session through a Redfish host
	// interface using an 'AuthNone' role, the property shall contain an empty
	// string in the request body.
	UserName string
	// outboundConnection is the URI for OutboundConnection.
	outboundConnection string
}

// UnmarshalJSON unmarshals a Session object from the raw JSON.
func (s *Session) UnmarshalJSON(b []byte) error {
	type temp Session
	type sLinks struct {
		OutboundConnection Link `json:"OutboundConnection"`
	}
	var tmp struct {
		temp
		Links sLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Session(tmp.temp)

	// Extract the links to other entities for later
	s.outboundConnection = tmp.Links.OutboundConnection.String()

	return nil
}

// GetSession will get a Session instance from the service.
func GetSession(c Client, uri string) (*Session, error) {
	return GetObject[Session](c, uri)
}

// ListReferencedSessions gets the collection of Session from
// a provided reference.
func ListReferencedSessions(c Client, link string) ([]*Session, error) {
	return GetCollectionObjects[Session](c, link)
}

// OutboundConnection gets the OutboundConnection linked resource.
func (s *Session) OutboundConnection() (*OutboundConnection, error) {
	if s.outboundConnection == "" {
		return nil, nil
	}
	return GetObject[OutboundConnection](s.client, s.outboundConnection)
}

// AuthToken contains the authentication and session information.
type AuthToken struct {
	Token     string
	Session   string
	Username  string
	Password  string
	BasicAuth bool
}

type authPayload struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

// CreateSession creates a new session and returns the token and id
func CreateSession(c Client, uri, username, password string) (auth *AuthToken, err error) {
	a := &authPayload{
		UserName: username,
		Password: password,
	}

	resp, err := c.Post(uri, a)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return auth, err
	}

	auth = &AuthToken{}
	auth.Token = resp.Header.Get("X-Auth-Token")
	auth.Session = resp.Header.Get("Location")

	if urlParser, err := url.ParseRequestURI(auth.Session); err == nil {
		auth.Session = urlParser.RequestURI()
	}

	return auth, err
}

// DeleteSession deletes a session using the location as argument
func DeleteSession(c Client, sessionURL string) (err error) {
	resp, err := c.Delete(sessionURL)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}
	return nil
}
