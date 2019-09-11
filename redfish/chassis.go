//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
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

// Chassis represents the physical components of a system. This
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
	networkAdapters string
	computerSystems []string
	resourceBlocks  []string
	managedBy       []string
	// resetTarget is the internal URL to send reset actions to.
	resetTarget string
	// SupportedResetTypes, if provided, is the reset types this chassis supports.
	SupportedResetTypes []ResetType
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
func (c *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	type linkReference struct {
		ComputerSystems common.Links
		ResourceBlocks  common.Links
		ManagedBy       common.Links
	}
	type Actions struct {
		ChassisReset struct {
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			Target            string
		} `json:"#Chassis.Reset"`
	}

	var t struct {
		temp
		Thermal         common.Link
		Power           common.Link
		NetworkAdapters common.Link
		Links           linkReference
		Actions         Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*c = Chassis(t.temp)

	// Extract the links to other entities for later
	c.thermal = string(t.Thermal)
	c.power = string(t.Power)
	c.networkAdapters = string(t.NetworkAdapters)
	c.computerSystems = t.Links.ComputerSystems.ToStrings()
	c.resourceBlocks = t.Links.ResourceBlocks.ToStrings()
	c.managedBy = t.Links.ManagedBy.ToStrings()
	c.resetTarget = t.Actions.ChassisReset.Target
	c.SupportedResetTypes = t.Actions.ChassisReset.AllowedResetTypes

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

// Thermal gets the thermal temperature and cooling information for the chassis
func (c *Chassis) Thermal() (*Thermal, error) {
	if c.thermal == "" {
		return nil, nil
	}

	resp, err := c.Client.Get(c.thermal)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermal Thermal
	err = json.NewDecoder(resp.Body).Decode(&thermal)
	if err != nil {
		return nil, err
	}

	return &thermal, nil
}

// Power gets the power information for the chassis
func (c *Chassis) Power() (*Power, error) {
	if c.power == "" {
		return nil, nil
	}

	resp, err := c.Client.Get(c.power)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var power Power
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

// NetworkAdapters gets the collection of network adapters of this chassis
func (c *Chassis) NetworkAdapters() ([]*NetworkAdapter, error) {
	return ListReferencedNetworkAdapter(c.Client, c.networkAdapters)
}

// Reset shall reset the chassis. This action shall not reset Systems or other
// contained resource, although side effects may occur which affect those resources.
func (c *Chassis) Reset(resetType ResetType) error {
	// Make sure the requested reset type is supported by the chassis
	valid := false
	if len(c.SupportedResetTypes) > 0 {
		for _, allowed := range c.SupportedResetTypes {
			if resetType == allowed {
				valid = true
				break
			}
		}
	} else {
		// No allowed values supplied, assume we are OK
		valid = true
	}

	if !valid {
		return fmt.Errorf("Reset type '%s' is not supported by this chassis",
			resetType)
	}

	type temp struct {
		ResetType ResetType
	}
	t := temp{
		ResetType: resetType,
	}

	payload, err := json.Marshal(t)
	if err != nil {
		return err
	}

	_, err = c.Client.Post(c.resetTarget, payload)
	return err
}
