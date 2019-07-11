// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gofish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
	"github.com/stmcginnis/gofish/school/redfish"
	"github.com/stmcginnis/gofish/school/swordfish"
)

// Service represents the service root of a Redfish enabled device.
type Service struct {
	common.Entity
	Version            string `json:"ServiceVersion"`
	UUID               string `json:"UUID"`
	chassis            string
	managers           string
	taskService        string
	sessions           string
	storageServices    string
	storageSystems     string
	accountService     string
	eventService       string
	registries         string
	systems            string
	compositionService string
}

// UnmarshalJSON unmarshals a Service object from the raw JSON.
func (s *Service) UnmarshalJSON(b []byte) error {
	type temp Service
	var t struct {
		temp
		// Link references
		Chassis            common.Link
		Managers           common.Link
		TaskService        common.Link
		StorageServices    common.Link
		StorageSystems     common.Link
		AccountService     common.Link
		EventService       common.Link
		Registries         common.Link
		Systems            common.Link
		CompositionService common.Link
		Links              struct {
			Sessions common.Link
		} `json:"Links"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = Service(t.temp)

	// Extract the links to other entities for later
	s.chassis = string(t.Chassis)
	s.managers = string(t.Managers)
	s.taskService = string(t.TaskService)
	s.sessions = string(t.Links.Sessions)
	s.storageServices = string(t.StorageServices)
	s.storageSystems = string(t.StorageSystems)
	s.accountService = string(t.AccountService)
	s.eventService = string(t.EventService)
	s.registries = string(t.Registries)
	s.systems = string(t.Systems)
	s.compositionService = string(t.CompositionService)

	return nil
}

// ServiceRoot gets the root service of the Redfish service.
func ServiceRoot(c common.Client) (*Service, error) {
	resp, err := c.Get(common.DefaultServiceRoot)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var service Service
	err = json.NewDecoder(resp.Body).Decode(&service)
	if err != nil {
		return nil, err
	}

	service.SetClient(c)
	return &service, nil
}

// Chassis gets the chassis instances managed by this service.
func (s *Service) Chassis() ([]*redfish.Chassis, error) {
	return redfish.ListReferencedChassis(s.Client, s.chassis)
}

// Managers gets the manager instances of this service.
func (s *Service) Managers() ([]*redfish.Manager, error) {
	return redfish.ListReferencedManagers(s.Client, s.managers)
}

// StorageSystems gets the storage system instances managed by this service.
func (s *Service) StorageSystems() ([]*swordfish.StorageSystem, error) {
	return swordfish.ListReferencedStorageSystems(s.Client, s.storageSystems)
}

// StorageServices gets the Swordfish storage services
func (s *Service) StorageServices() ([]*swordfish.StorageService, error) {
	return swordfish.ListReferencedStorageServices(s.Client, s.storageServices)
}

// Tasks gets the system's tasks
func (s *Service) Tasks() ([]*redfish.Task, error) {
	return redfish.ListReferencedTasks(s.Client, s.taskService)
}

// CreateSession creates a new session and returns the token and id
func (s *Service) CreateSession(username string, password string) (*redfish.AuthToken, error) {
	return redfish.CreateSession(s.Client, s.sessions, username, password)
}

// Sessions gets the system's active sessions
func (s *Service) Sessions() ([]*redfish.Session, error) {
	return redfish.ListReferencedSessions(s.Client, s.sessions)
}

// DeleteSession logout the specified session
func (s *Service) DeleteSession(url string) error {
	return redfish.DeleteSession(s.Client, url)
}

// AccountService gets the Redfish AccountService
func (s *Service) AccountService() (*redfish.AccountService, error) {
	return redfish.GetAccountService(s.Client, s.accountService)
}

// EventService gets the Redfish EventService
func (s *Service) EventService() (*redfish.EventService, error) {
	return redfish.GetEventService(s.Client, s.eventService)
}

// Systems get the system instances from the service
func (s *Service) Systems() ([]*redfish.ComputerSystem, error) {
	return redfish.ListReferencedComputerSystems(s.Client, s.systems)
}

// CompositionService gets the composition service instance
func (s *Service) CompositionService() (*redfish.CompositionService, error) {
	return redfish.GetCompositionService(s.Client, s.compositionService)
}
