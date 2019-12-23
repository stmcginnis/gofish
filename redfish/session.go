//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SessionTypes is the type of session.
type SessionTypes string

const (

	// HostConsoleSessionTypes The Host's console, which could be connected
	// via Telnet, SSH, or other protocol.
	HostConsoleSessionTypes SessionTypes = "HostConsole"
	// ManagerConsoleSessionTypes The Manager's console, which could be
	// connected via Telnet, SSH, SM CLP, or other protocol.
	ManagerConsoleSessionTypes SessionTypes = "ManagerConsole"
	// IPMISessionTypes Intelligent Platform Management Interface.
	IPMISessionTypes SessionTypes = "IPMI"
	// KVMIPSessionTypes Keyboard-Video-Mouse over IP Session.
	KVMIPSessionTypes SessionTypes = "KVMIP"
	// OEMSessionTypes OEM Type. Please look at OemSessionType for OEM
	// session type(s).
	OEMSessionTypes SessionTypes = "OEM"
	// RedfishSessionTypes A Redfish Session.
	RedfishSessionTypes SessionTypes = "Redfish"
	// VirtualMediaSessionTypes Virtual Media.
	VirtualMediaSessionTypes SessionTypes = "VirtualMedia"
	// WebUISessionTypes A non-Redfish Web User Interface session such as a
	// graphical interface or other kinds of web-based protocols.
	WebUISessionTypes SessionTypes = "WebUI"
)

// Session describes a single connection (session) between a client and a
// Redfish service instance.
type Session struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// OemSessionType is used to report the OEM-specific session type. Thus,
	// this property shall represent the type of OEM session that is
	// currently active.
	OemSessionType string
	// Password shall be the password for this session. The value shall be null
	// for GET requests.
	Password string
	// SessionType shall represent the type of session that is currently active.
	SessionType SessionTypes
	// UserName shall be the UserName that matches a registered account
	// identified by a ManagerAccount resource registered with the Account
	// Service.
	UserName string
}

// AuthToken contains the authentication and session information.
type AuthToken struct {
	Token   string
	Session string
}

type authPayload struct {
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

// CreateSession creates a new session and returns the token and id
func CreateSession(c common.Client, uri string, username string, password string) (auth *AuthToken, err error) {
	a := &authPayload{
		Username: username,
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

	return auth, err
}

// DeleteSession deletes a session using the location as argument
func DeleteSession(c common.Client, url string) (err error) {
	return c.Delete(url)
}

// GetSession will get a Session instance from the Redfish service.
func GetSession(c common.Client, uri string) (*Session, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t Session
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedSessions gets the collection of Sessions
func ListReferencedSessions(c common.Client, link string) ([]*Session, error) {
	var result []*Session
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, sLink := range links.ItemLinks {
		s, err := GetSession(c, sLink)
		if err != nil {
			return result, err
		}
		result = append(result, s)
	}

	return result, nil
}
