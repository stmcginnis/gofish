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

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// UIConsoleInfo contains information about GUI services.
type UIConsoleInfo struct {
	ServiceEnabled        bool
	MaxConcurrentSessions uint
	ConnectTypesSupported []string
}

// ManagerType shall describe the function of this manager. The value
// EnclosureManager shall be used if this manager controls one or more services
// through aggregation. The value BMC shall be used if this manager represents a
// traditional server management controller. The value ManagementController
// shall be used if none of the other enumerations apply.
type ManagerType string

const (
	// AuxiliaryControllerManagerType a controller which provides management functions for a particular subsystem or group of devices
	AuxiliaryControllerManagerType ManagerType = "AuxiliaryController"
	// BMCManagerType a controller which provides management functions for a single computer system
	BMCManagerType ManagerType = "BMC"
	// EnclosureManagerManagerType a controller which provides management functions for a chassis or group of devices or systems
	EnclosureManagerManagerType ManagerType = "EnclosureManager"
	// ManagementControllerManagerType a controller used primarily to monitor or manage the operation of a device or system
	ManagementControllerManagerType ManagerType = "ManagementController"
	// RackManagerManagerType a controller which provides management functions for a whole or part of a rack
	RackManagerManagerType ManagerType = "RackManager"
	// ServiceManagerType a software-based service which provides management functions
	ServiceManagerType ManagerType = "Service"
)

// Manager is a management subsystem. Examples of managers are BMCs, Enclosure
// Managers, Management Controllers and other subsystems assigned managability
// functions.
type Manager struct {
	common.Entity
	ManagerType           ManagerType
	Description           string
	ServiceEntryPointUUID string
	UUID                  string
	Model                 string
	DateTime              string
	DateTimeLocalOffset   string
	Status                common.Status
	GraphicalConsole      UIConsoleInfo
	SerialConsole         UIConsoleInfo
	CommandShell          UIConsoleInfo
	FirmwareVersion       string
	networkProtocol       string
	ethernetInterfaces    string
	serialInterfaces      string
	logServices           string
	virtualMedia          string
	managerForServers     []string
	managerForChassis     []string
	managerInChassis      string
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (s *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	type linkReference struct {
		ManagerForServers common.Links
		ManagerForChassis common.Links
		ManagerInChassis  common.Link
	}
	var t struct {
		temp
		NetworkProtocol    common.Link
		EthernetInterfaces common.Link
		SerialInterfaces   common.Link
		LogServices        common.Link
		VirtualMedia       common.Link
		Links              linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = Manager(t.temp)

	// Extract the links to other entities
	s.networkProtocol = string(t.NetworkProtocol)
	s.ethernetInterfaces = string(t.EthernetInterfaces)
	s.serialInterfaces = string(t.SerialInterfaces)
	s.logServices = string(t.LogServices)
	s.virtualMedia = string(t.VirtualMedia)
	s.managerForServers = t.Links.ManagerForServers.ToStrings()
	s.managerForChassis = t.Links.ManagerForChassis.ToStrings()
	s.managerInChassis = string(t.Links.ManagerInChassis)

	return nil
}

// GetManager will get a Manager instance from the Swordfish service.
func GetManager(c common.Client, uri string) (*Manager, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manager Manager
	err = json.NewDecoder(resp.Body).Decode(&manager)
	if err != nil {
		return nil, err
	}

	manager.SetClient(c)
	return &manager, nil
}

// ListReferencedManagers gets the collection of Managers
func ListReferencedManagers(c common.Client, link string) ([]*Manager, error) {
	var result []*Manager
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, managerLink := range links.ItemLinks {
		manager, err := GetManager(c, managerLink)
		if err != nil {
			return result, err
		}
		result = append(result, manager)
	}

	return result, nil
}
