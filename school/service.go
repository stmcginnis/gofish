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

package school

import (
	"encoding/json"
)

// Service represents the service root of a Redfish enabled device.
type Service struct {
	Entity
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
		Chassis            Link
		Managers           Link
		TaskService        Link
		SessionService     Link
		StorageServices    Link
		StorageSystems     Link
		AccountService     Link
		EventService       Link
		Registries         Link
		Systems            Link
		CompositionService Link
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
func ServiceRoot(c Client) (*Service, error) {

	resp, err := c.Get("/redfish/v1")
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
func (s *Service) Chassis() ([]*Chassis, error) {
	return ListReferencedChassis(s.client, s.chassis)
}

// StorageSystems gets the storage system instances managed by this service.
func (s *Service) StorageSystems() ([]*StorageSystem, error) {
	return ListReferencedStorageSystems(s.client, s.storageSystems)
}

// StorageServices gets the Swordfish storage services
func (s *Service) StorageServices() ([]*StorageService, error) {
	return ListReferencedStorageServices(s.client, s.storageServices)
}

// Tasks gets the system's tasks
func (s *Service) Tasks() ([]*Task, error) {
	return ListReferencedTasks(s.client, s.taskService)
}

// Sessions gets the system's active sessions
func (s *Service) Sessions() ([]*Session, error) {
	return ListReferencedSessions(s.client, s.sessionService)
}

// AccountService gets the Redfish AccountService
func (s *Service) AccountService() (*AccountService, error) {
	return GetAccountService(s.client, s.accountService)
}

// EventService gets the Redfish EventService
func (s *Service) EventService() (*EventService, error) {
	return GetEventService(s.client, s.eventService)
}

// Systems get the system instances from the service
func (s *Service) Systems() ([]*ComputerSystem, error) {
	return ListReferencedComputerSystems(s.client, s.systems)
}

// CompositionService gets the composition service instance
func (s *Service) CompositionService() (*CompositionService, error) {
	return GetCompositionService(s.client, s.compositionService)
}
