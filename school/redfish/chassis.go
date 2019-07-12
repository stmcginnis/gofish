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

// ChassisType is a physical form of the chassis
type ChassisType string

const (
	// BladeChassisType is an enclosed or semi-enclosed, typically vertically-oriented, system
	//mchassis which must be plugged into a multi-system chassis to function normally.
	BladeChassisType ChassisType = "Blade"
	// CardChassisType is a loose device or circuit board intended to be installed in a system
	// or other enclosure.
	CardChassisType ChassisType = "Card"
	// CartridgeChassisType is a small self-contained system intended to be plugged into a multi-system
	// chassis.
	CartridgeChassisType ChassisType = "Cartridge"
	// ComponentChassisType is a small chassis, card, or device which contains devices for a particular
	// subsystem or function.
	ComponentChassisType ChassisType = "Component"
	// DrawerChassisType is an enclosed or semi-enclosed, typically horizontally-oriented, system
	// chassis which may be slid into a multi-system chassis.
	DrawerChassisType ChassisType = "Drawer"
	// EnclosureChassisType is a generic term for a chassis that does not fit any other description.
	EnclosureChassisType ChassisType = "Enclosure"
	// ExpansionChassisType is a chassis which expands the capabilities or capacity of another
	// chassis.
	ExpansionChassisType ChassisType = "Expansion"
	// IPBasedDriveChassisType is a chassis in a drive form factor with IP-based network connections.
	IPBasedDriveChassisType ChassisType = "IPBasedDrive"
	// ModuleChassisType is a small, typically removable, chassis or card which contains devices
	// for a particular subsystem or function.
	ModuleChassisType ChassisType = "Module"
	// OtherChassisType is a chassis that does not fit any of these definitions.
	OtherChassisType ChassisType = "Other"
	// PodChassisType is a collection of equipment racks in a large, likely transportable, container.
	PodChassisType ChassisType = "Pod"
	// RackChassisType is an equipment rack, typically a 19-inch wide freestanding unit.
	RackChassisType ChassisType = "Rack"
	// RackGroupChassisType is a group of racks which form a single entity or share infrastructure.
	RackGroupChassisType ChassisType = "RackGroup"
	// RackMountChassisType is a single system chassis designed specifically for mounting in an
	// equipment rack.
	RackMountChassisType ChassisType = "RackMount"
	// RowChassisType is a collection of equipment racks.
	RowChassisType ChassisType = "Row"
	// ShelfChassisType is an enclosed or semi-enclosed, typically horizontally-oriented, system
	// chassis which must be plugged into a multi-system chassis to function normally.
	ShelfChassisType ChassisType = "Shelf"
	// SidecarChassisType is a chassis that mates mechanically with another chassis to expand
	// its capabilities or capacity.
	SidecarChassisType ChassisType = "Sidecar"
	// SledChassisType is an enclosed or semi-enclosed, system chassis which must be plugged into
	// a multi-system chassis to function normally similar to a blade type chassis.
	SledChassisType ChassisType = "Sled"
	// StandAloneChassisType is a single, free-standing system, commonly called a tower or desktop
	// chassis.
	StandAloneChassisType ChassisType = "StandAlone"
	// StorageEnclosureChassisType is a chassis which encloses storage.
	StorageEnclosureChassisType ChassisType = "StorageEnclosure"
	// ZoneChassisType is a logical division or portion of a physical chassis that contains multiple
	// devices or systems that cannot be physically separated.
	ZoneChassisType ChassisType = "Zone"
)

// Chassis represents the physical components of a system.  This
// resource represents the sheet-metal confined spaces and logical zones such
// as racks, enclosures, chassis and all other containers. Subsystems (like sensors)
// that operate outside of a system's data plane (meaning the resources are not
// accessible to software running on the system) are linked either directly or
// indirectly through this resource.
type Chassis struct {
	common.Entity
	ChassisType     ChassisType   `json:"ChassisType"`
	Manufacturer    string        `json:"Manufacturer"`
	Model           string        `json:"Model"`
	SKU             string        `json:"SKU"`
	SerialNumber    string        `json:"SerialNumber"`
	Version         string        `json:"Version"`
	PartNumber      string        `json:"PartNumber"`
	AssetTag        string        `json:"AssetTag"`
	Status          common.Status `json:"Status"`
	thermal         string
	power           string
	computerSystems []string
	resourceBlocks  []string
	managedBy       []string
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
func (c *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	type linkReference struct {
		ComputerSystems common.Links
		ResourceBlocks  common.Links
		ManagedBy       common.Links
	}
	var t struct {
		temp
		Thermal common.Link
		Power   common.Link
		Links   linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*c = Chassis(t.temp)

	// Extract the links to other entities for later
	c.thermal = string(t.Thermal)
	c.power = string(t.Power)
	c.computerSystems = t.Links.ComputerSystems.ToStrings()
	c.resourceBlocks = t.Links.ResourceBlocks.ToStrings()
	c.managedBy = t.Links.ManagedBy.ToStrings()

	return nil
}

// GetChassis will get a Chassis instance from the Redfish service.
func GetChassis(c common.Client, uri string) (*Chassis, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var chassis Chassis
	err = json.NewDecoder(resp.Body).Decode(&chassis)
	if err != nil {
		return nil, err
	}

	chassis.SetClient(c)
	return &chassis, nil
}

// ListReferencedChassis gets the collection of Chassis from a provided reference.
func ListReferencedChassis(c common.Client, link string) ([]*Chassis, error) {
	var result []*Chassis
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, chassisLink := range links.ItemLinks {
		chassis, err := GetChassis(c, chassisLink)
		if err != nil {
			return result, err
		}
		result = append(result, chassis)
	}

	return result, nil
}

// ThermalInfo is reference to the thermal properties (fans, cooling, sensors)
// of this chassis.
type ThermalInfo struct {
	common.Entity
	Temperatures []struct {
		MemberID                  string `json:"MemberId"`
		Name                      string
		SensorNumber              int
		Status                    common.Status
		ReadingCelsius            int
		UpperThresholdNonCritical int
		UpperThresholdCritical    int
		UpperThresholdFatal       int
		LowerThresholdNonCritical int
		LowerThresholdCritical    int
		LowerThresholdFatal       int
		MinimumValue              int
		MaximumValue              int
		PhysicalContext           string
		RelatedItem               []common.Link
	}
	Fans []struct {
		MemberID                  string `json:"MemberId"`
		FanName                   string
		PhysicalContext           string
		Status                    common.Status
		ReadingRPM                int
		UpperThresholdNonCritical int
		UpperThresholdCritical    int
		UpperThresholdFatal       int
		LowerThresholdNonCritical int
		LowerThresholdCritical    int
		LowerThresholdFatal       int
		MinReadingRange           int
		MaxReadingRange           int
		Redundancy                []common.Link
		RelatedItem               []common.Link
	}
	Redundancy []struct {
		MemberID        string `json:"MemberId"`
		Name            string
		RedundancySet   []common.Link
		Mode            string
		Status          common.Status
		MinNumNeeded    int
		MaxNumSupported int
	}
}

// Thermal gets the thermal temperature and cooling information for the chassis
func (c *Chassis) Thermal() (*ThermalInfo, error) {
	if c.thermal == "" {
		return nil, nil
	}

	resp, err := c.Client.Get(c.thermal)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermal ThermalInfo
	err = json.NewDecoder(resp.Body).Decode(&thermal)
	if err != nil {
		return nil, err
	}

	return &thermal, nil
}

// PowerInfo provides the power properties (power supplies, power
// policies, sensors) of the chassis.
type PowerInfo struct {
	common.Entity
	PowerControl []struct {
		MemberID            string `json:"MemberId"`
		Name                string
		PowerConsumedWatts  int
		PowerRequestedWatts int
		PowerAvailableWatts int
		PowerCapacityWatts  int
		PowerAllocatedWatts int
		PowerMetrics        struct {
			IntervalInMin        int
			MinConsumedWatts     int
			MaxConsumedWatts     int
			AverageConsumedWatts int
		}
		PowerLimit struct {
			LimitInWatts   int
			LimitException string
			CorrectionInMS int `json:"CorrectionInMs"`
		}
		RelatedItem []common.Link
		Status      common.Status
	}
	Voltages []struct {
		MemberID                  string `json:"MemberId"`
		Name                      string
		SensorNumber              int
		Status                    common.Status
		ReadingVolts              int
		UpperThresholdNonCritical float32
		UpperThresholdCritical    float32
		UpperThresholdFatal       float32
		LowerThresholdNonCritical float32
		LowerThresholdCritical    float32
		LowerThresholdFatal       float32
		MinReadingRange           int
		MaxReadingRange           int
		PhysicalContext           string
		RelatedItem               []common.Link
	}
	PowerSupplies []struct {
		MemberID             string `json:"MemberId"`
		Name                 string
		Status               common.Status
		PowerSupplyType      string
		LineInputVoltageType string
		LineInputVoltage     int
		PowerCapacityWatts   int
		LastPowerOutputWatts int
		Model                string
		FirmwareVersion      string
		SerialNumber         string
		PartNumber           string
		SparePartNumber      string
		RelatedItem          []common.Link
	}
}

// Power gets the power information for the chassis
func (c *Chassis) Power() (*PowerInfo, error) {
	if c.power == "" {
		return nil, nil
	}

	resp, err := c.Client.Get(c.power)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var power PowerInfo
	err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	return &power, nil
}

// ComputerSystems returns the collection of systems from this chassis
func (c *Chassis) ComputerSystems() ([]*ComputerSystem, error) {
	var result []*ComputerSystem
	for _, uri := range c.computerSystems {
		cs, err := GetComputerSystem(c.Client, uri)
		if err != nil {
			return nil, err
		}

		result = append(result, cs)
	}

	return result, nil
}

// ManagedBy gets the collection of managers of this chassis
func (c *Chassis) ManagedBy() ([]*Manager, error) {
	var result []*Manager
	for _, uri := range c.managedBy {
		manager, err := GetManager(c.Client, uri)
		if err != nil {
			return nil, err
		}

		result = append(result, manager)
	}

	return result, nil
}
