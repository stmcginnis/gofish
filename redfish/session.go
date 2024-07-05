//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"net/url"

	"github.com/stmcginnis/gofish/common"
)

// SessionTypes is the type of session.
type SessionTypes string

const (
	// HostConsoleSessionTypes shall indicate the session is the host's console, which could be connected through
	// Telnet, SSH, or another protocol. If this session is terminated or deleted, the service shall close the
	// connection for the respective host console session.
	HostConsoleSessionTypes SessionTypes = "HostConsole"
	// ManagerConsoleSessionTypes shall indicate the session is the manager's console, which could be connected through
	// Telnet, SSH, SM CLP, or another protocol. If this session is terminated or deleted, the service shall close the
	// connection for the respective manager console session.
	ManagerConsoleSessionTypes SessionTypes = "ManagerConsole"
	// IPMISessionTypes shall indicate the session is an Intelligent Platform Management Interface session. If this
	// session is terminated or deleted, the service shall close the connection for the respective IPMI session.
	IPMISessionTypes SessionTypes = "IPMI"
	// KVMIPSessionTypes shall indicate the session is a Keyboard-Video-Mouse over IP session. If this session is
	// terminated or deleted, the service shall close the connection for the respective KVM-IP session.
	KVMIPSessionTypes SessionTypes = "KVMIP"
	// OEMSessionTypes shall indicate the session is an OEM-specific session and is further described by the
	// OemSessionType property.
	OEMSessionTypes SessionTypes = "OEM"
	// RedfishSessionTypes shall indicate the session is a Redfish session defined by the 'Redfish session login
	// authentication' clause of the Redfish Specification. If this session is terminated or deleted, the service shall
	// invalidate the respective session token.
	RedfishSessionTypes SessionTypes = "Redfish"
	// VirtualMediaSessionTypes shall indicate the session is a virtual media session. If this session is terminated or
	// deleted, the service shall close the connection for the respective virtual media session and make the media
	// inaccessible to the host.
	VirtualMediaSessionTypes SessionTypes = "VirtualMedia"
	// WebUISessionTypes shall indicate the session is a non-Redfish web user interface session. If this session is
	// terminated or deleted, the service shall invalidate the respective session token.
	WebUISessionTypes SessionTypes = "WebUI"
	// OutboundConnectionSessionTypes shall indicate the session is an outbound connection defined by the 'Outbound
	// connections' clause of the Redfish Specification. The 'OutboundConnection' property inside the 'Links' property
	// shall contain the link to the outbound connection configuration. If this session is terminated or deleted, the
	// service shall disable the associated 'OutboundConnection' resource.
	OutboundConnectionSessionTypes SessionTypes = "OutboundConnection"
)

// Session describes a single connection (session) between a client and a
// Redfish service instance.
type Session struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClientOriginIPAddress shall contain the IP address of the client that created the session.
	ClientOriginIPAddress string
	// Context shall contain a client-supplied context that remains with the session through the session's lifetime.
	Context string
	// CreatedTime shall contain the date and time when the session was created.
	CreatedTime string
	// Description provides a description of this resource.
	Description string
	// OemSessionType is used to report the OEM-specific session type. Thus,
	// this property shall represent the type of OEM session that is
	// currently active.
	OemSessionType string
	// Password shall be the password for this session. The value shall be null
	// for GET requests.
	Password string
	// Roles shall contain the Redfish roles that contain the privileges of this session.
	Roles []string
	// SessionType shall represent the type of session that is currently active.
	SessionType SessionTypes
	// Token shall contain the multi-factor authentication token for this session. The value shall be 'null' in
	// responses.
	Token string
	// UserName shall be the UserName that matches a registered account
	// identified by a ManagerAccount resource registered with the Account
	// Service.
	UserName string

	outboundConnection string
}

// UnmarshalJSON unmarshals a Session object from the raw JSON.
func (session *Session) UnmarshalJSON(b []byte) error {
	type temp Session
	var t struct {
		temp
		Links struct {
			OutboundConnection common.Link
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*session = Session(t.temp)

	// Extract the links to other entities for later
	session.outboundConnection = t.Links.OutboundConnection.String()

	return nil
}

// OutboundConnection gets the outbound connection associated with this session.
func (session *Session) OutboundConnection() (*OutboundConnection, error) {
	if session.outboundConnection == "" {
		return nil, nil
	}
	return GetOutboundConnection(session.GetClient(), session.outboundConnection)
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
func CreateSession(c common.Client, uri, username, password string) (auth *AuthToken, err error) {
	a := &authPayload{
		UserName: username,
		Password: password,
	}

	resp, err := c.Post(uri, a)
	if err != nil {
		return auth, err
	}
	defer resp.Body.Close()

	auth = &AuthToken{}
	auth.Token = resp.Header.Get("X-Auth-Token")
	auth.Session = resp.Header.Get("Location")

	if urlParser, err := url.ParseRequestURI(auth.Session); err == nil {
		auth.Session = urlParser.RequestURI()
	}

	return auth, err
}

// DeleteSession deletes a session using the location as argument
func DeleteSession(c common.Client, sessionURL string) (err error) {
	resp, err := c.Delete(sessionURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// GetSession will get a Session instance from the Redfish service.
func GetSession(c common.Client, uri string) (*Session, error) {
	return common.GetObject[Session](c, uri)
}

// ListReferencedSessions gets the collection of Sessions
func ListReferencedSessions(c common.Client, link string) ([]*Session, error) {
	return common.GetCollectionObjects[Session](c, link)
}
