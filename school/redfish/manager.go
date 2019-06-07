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

// DefaultManagerPath is the default URI for Manager collections.
const DefaultManagerPath = "/redfish/v1/Managers"

// UIConsoleInfo contains information about GUI services.
type UIConsoleInfo struct {
	ServiceEnabled        string
	MaxConcurrentSessions uint
	ConnectTypesSupported []string
}

// Manager is a management subsystem. Examples of managers are BMCs, Enclosure
// Managers, Management Controllers and other subsystems assigned managability
// functions.
type Manager struct {
	common.Entity
	ManagerType           string
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

// ListManagers gets all Manager in the system
func ListManagers(c common.Client) ([]*Manager, error) {
	return ListReferencedManagers(c, DefaultManagerPath)
}
