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
	sessionService     string
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
	type linkReference struct {
		Chassis            common.Link
		Managers           common.Link
		TaskService        common.Link
		SessionService     common.Link
		StorageServices    common.Link
		StorageSystems     common.Link
		AccountService     common.Link
		EventService       common.Link
		Registries         common.Link
		Systems            common.Link
		CompositionService common.Link
	}
	var t struct {
		temp
		Links linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = Service(t.temp)

	// Extract the links to other entities for later
	s.chassis = string(t.Links.Chassis)
	s.managers = string(t.Links.Managers)
	s.taskService = string(t.Links.TaskService)
	s.sessionService = string(t.Links.SessionService)
	s.storageServices = string(t.Links.StorageServices)
	s.storageSystems = string(t.Links.StorageSystems)
	s.accountService = string(t.Links.AccountService)
	s.eventService = string(t.Links.EventService)
	s.registries = string(t.Links.Registries)
	s.systems = string(t.Links.Systems)
	s.compositionService = string(t.Links.CompositionService)

	return nil
}

// ServiceRoot gets the root service of the Redfish service.
func ServiceRoot(c common.Client) (*Service, error) {

	resp, err := c.Get(common.DefaultServiceRoot)
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

// Sessions gets the system's active sessions
func (s *Service) Sessions() ([]*redfish.Session, error) {
	return redfish.ListReferencedSessions(s.Client, s.sessionService)
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
