//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SessionService This resource contains the session service properties for a Redfish implementation.
type SessionService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled. If 'true', this service is enabled. If 'false',
	// it is disabled, and new sessions shall not be created, old sessions shall not be deleted, and established
	// sessions can continue operating.
	ServiceEnabled bool
	// SessionTimeout shall contain the threshold of time in seconds between requests on a specific session at which
	// point the session service shall close the session due to inactivity. The session service shall support any value
	// between the Validation.Minimum and Validation.Maximum.
	SessionTimeout int
	// Sessions shall contain a link to a resource collection of type SessionCollection.
	sessions []string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SessionService object from the raw JSON.
func (sessionservice *SessionService) UnmarshalJSON(b []byte) error {
	type temp SessionService
	var t struct {
		temp
		Sessions common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sessionservice = SessionService(t.temp)

	// Extract the links to other entities for later
	sessionservice.sessions = t.Sessions.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	sessionservice.rawData = b

	return nil
}

// Sessions gets a collection of sessions.
func (sessionservice *SessionService) Sessions() ([]*Session, error) {
	var result []*Session

	collectionError := common.NewCollectionError()
	for _, ethLink := range sessionservice.sessions {
		eth, err := GetSession(sessionservice.GetClient(), ethLink)
		if err != nil {
			collectionError.Failures[ethLink] = err
		} else {
			result = append(result, eth)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Update commits updates to this object's properties to the running system.
func (sessionservice *SessionService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SessionService)
	original.UnmarshalJSON(sessionservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
		"SessionTimeout",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(sessionservice).Elem()

	return sessionservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSessionService will get a SessionService instance from the service.
func GetSessionService(c common.Client, uri string) (*SessionService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sessionservice SessionService
	err = json.NewDecoder(resp.Body).Decode(&sessionservice)
	if err != nil {
		return nil, err
	}

	sessionservice.SetClient(c)
	return &sessionservice, nil
}

// ListReferencedSessionServices gets the collection of SessionService from
// a provided reference.
func ListReferencedSessionServices(c common.Client, link string) ([]*SessionService, error) {
	var result []*SessionService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *SessionService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		sessionservice, err := GetSessionService(c, link)
		ch <- GetResult{Item: sessionservice, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
