//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/SessionService.v1_2_0.json
// 2024.4 - #SessionService.v1_2_0.SessionService

package schemas

import (
	"encoding/json"
)

// SessionService This resource contains the session service properties for a
// Redfish implementation.
type SessionService struct {
	Entity
	// AbsoluteSessionTimeout shall contain the maximum number of seconds that a
	// session is open before the service closes the session regardless of
	// activity.
	//
	// Version added: v1.2.0
	AbsoluteSessionTimeout uint
	// AbsoluteSessionTimeoutEnabled shall indicate whether an absolute session
	// timeout is applied to sessions. If 'true', the service shall close sessions
	// that are open for the number of seconds specified by the
	// 'AbsoluteSessionTimeout' property regardless of session activity. If 'false'
	// or if this property is not present, the service shall not apply an absolute
	// session timeout.
	//
	// Version added: v1.2.0
	AbsoluteSessionTimeoutEnabled bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled. If 'true',
	// this service is enabled. If 'false', it is disabled, and new sessions shall
	// not be created, old sessions shall not be deleted, and established sessions
	// can continue operating.
	ServiceEnabled bool
	// SessionTimeout shall contain the threshold of time in seconds between
	// requests on a specific session at which point the session service shall
	// close the session due to inactivity. The session service shall support any
	// value between the schema-specified minimum and maximum terms.
	SessionTimeout uint
	// Sessions shall contain a link to a resource collection of type
	// 'SessionCollection'.
	sessions string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SessionService object from the raw JSON.
func (s *SessionService) UnmarshalJSON(b []byte) error {
	type temp SessionService
	var tmp struct {
		temp
		Sessions Link `json:"Sessions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SessionService(tmp.temp)

	// Extract the links to other entities for later
	s.sessions = tmp.Sessions.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SessionService) Update() error {
	readWriteFields := []string{
		"AbsoluteSessionTimeout",
		"AbsoluteSessionTimeoutEnabled",
		"ServiceEnabled",
		"SessionTimeout",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSessionService will get a SessionService instance from the service.
func GetSessionService(c Client, uri string) (*SessionService, error) {
	return GetObject[SessionService](c, uri)
}

// ListReferencedSessionServices gets the collection of SessionService from
// a provided reference.
func ListReferencedSessionServices(c Client, link string) ([]*SessionService, error) {
	return GetCollectionObjects[SessionService](c, link)
}

// Sessions gets the Sessions collection.
func (s *SessionService) Sessions() ([]*Session, error) {
	if s.sessions == "" {
		return nil, nil
	}
	return GetCollectionObjects[Session](s.client, s.sessions)
}
