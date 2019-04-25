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

// DefaultSystemPath is the default URI for System collections.
const DefaultSystemPath = "/redfish/v1/Systems"

// ProcessorSummaryInfo represents information about the system processors
type ProcessorSummaryInfo struct {
	Status common.Status
	Count  uint
	Model  string
}

// MemorySummaryInfo represents information about the system memory
type MemorySummaryInfo struct {
	Status                         common.Status
	TotalSystemMemoryGiB           uint
	TotalSystemPersistentMemoryGiB uint
}

// TrustedModuleInfo represents information about system TPMs
type TrustedModuleInfo struct {
	Status                common.Status
	ModuleType            string
	FirmwareVersion       string
	FirmwareVersion2      string
	IntefaceTypeSelection string
}

// ComputerSystem represents a machine (physical or virtual) and the local resources
// such as memory, cpu and other devices that can be accessed from that machine.
type ComputerSystem struct {
	common.Entity
	SystemType         string
	AssetTag           string
	Manufacturer       string
	Model              string
	SKU                string
	SerialNumber       string
	PartNumber         string
	Description        string
	UUID               string
	HostName           string
	Status             common.Status
	IndicatorLED       common.IndicatorLED
	PowerState         string
	Boot               common.BootSettings
	BiosVersion        string
	ProcessorSummary   ProcessorSummaryInfo
	MemorySummary      MemorySummaryInfo
	TrustedModules     []TrustedModuleInfo
	processors         string
	memory             string
	ethernetInterfaces string
	simpleStorage      string
	chassis            []string
	managedBy          []string
	oem                string
}

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
func (s *ComputerSystem) UnmarshalJSON(b []byte) error {
	type temp ComputerSystem
	type linkReference struct {
		Chassis   common.Links
		ManagedBy common.Links
		OEM       common.Link `json:"Oem"`
	}
	var t struct {
		temp
		Processors         common.Link
		Memory             common.Link
		EthernetInterfaces common.Link
		SimpleStorage      common.Link
		Links              linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = ComputerSystem(t.temp)

	// Extract the links to other entities
	s.processors = string(t.Processors)
	s.memory = string(t.Memory)
	s.ethernetInterfaces = string(t.EthernetInterfaces)
	s.simpleStorage = string(t.SimpleStorage)
	s.chassis = t.Links.Chassis.ToStrings()
	s.managedBy = t.Links.ManagedBy.ToStrings()
	s.oem = string(t.Links.OEM)

	return nil
}

// GetComputerSystem will get a ComputerSystem instance from the Swordfish service.
func GetComputerSystem(c common.Client, uri string) (*ComputerSystem, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var cs ComputerSystem
	err = json.NewDecoder(resp.Body).Decode(&cs)
	if err != nil {
		return nil, err
	}

	cs.SetClient(c)
	return &cs, nil
}

// ListReferencedComputerSystems gets the collection of ComputerSystems
func ListReferencedComputerSystems(c common.Client, link string) ([]*ComputerSystem, error) {
	var result []*ComputerSystem
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, csLink := range links.ItemLinks {
		cs, err := GetComputerSystem(c, csLink)
		if err != nil {
			return result, err
		}
		result = append(result, cs)
	}

	return result, nil
}

// ListComputerSystems gets all ComputerSystem in the system
func ListComputerSystems(c common.Client) ([]*ComputerSystem, error) {
	return ListReferencedComputerSystems(c, DefaultSystemPath)
}
