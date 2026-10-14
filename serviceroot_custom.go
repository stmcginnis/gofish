//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/schemas"
)

// ServiceRoot will get a Service instance from the service.
func ServiceRoot(c schemas.Client) (*Service, error) {
	resp, err := c.Get(schemas.DefaultServiceRoot)
	defer schemas.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var serviceroot Service
	err = json.NewDecoder(resp.Body).Decode(&serviceroot)
	if err != nil {
		return nil, err
	}

	serviceroot.SetClient(c)
	return &serviceroot, nil
}

// StorageSystems gets the StorageSystems collection.
func (s *Service) StorageSystems() ([]*schemas.ComputerSystem, error) {
	if s.storageSystems == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.ComputerSystem](s.GetClient(), s.storageSystems)
}

// CreateSession creates a new session and returns the token and id
func (s *Service) CreateSession(username, password string) (*schemas.AuthToken, error) {
	return schemas.CreateSession(s.GetClient(), s.sessions, username, password)
}

// DeleteSession logout the specified session
func (s *Service) DeleteSession(url string) error {
	return schemas.DeleteSession(s.GetClient(), url)
}
