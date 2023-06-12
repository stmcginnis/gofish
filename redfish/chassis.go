//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ChassisType is a physical form of the chassis
type ChassisType string

const (
	// BladeChassisType is an enclosed or semi-enclosed, typically vertically-oriented, system
	// chassis which must be plugged into a multi-system chassis to function normally.
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

// EnvironmentalClass is
type EnvironmentalClass string

const (

	// A1EnvironmentalClass ASHRAE Environmental Class 'A1'.
	A1EnvironmentalClass EnvironmentalClass = "A1"
	// A2EnvironmentalClass ASHRAE Environmental Class 'A2'.
	A2EnvironmentalClass EnvironmentalClass = "A2"
	// A3EnvironmentalClass ASHRAE Environmental Class 'A3'.
	A3EnvironmentalClass EnvironmentalClass = "A3"
	// A4EnvironmentalClass ASHRAE Environmental Class 'A4'.
	A4EnvironmentalClass EnvironmentalClass = "A4"
)

// IntrusionSensor is
type IntrusionSensor string

const (

	// NormalIntrusionSensor No abnormal physical security condition is
	// detected at this time.
	NormalIntrusionSensor IntrusionSensor = "Normal"
	// HardwareIntrusionIntrusionSensor A door, lock, or other mechanism
	// protecting the internal system hardware from being accessed is
	// detected to be in an insecure state.
	HardwareIntrusionIntrusionSensor IntrusionSensor = "HardwareIntrusion"
	// TamperingDetectedIntrusionSensor Physical tampering of the monitored
	// entity is detected.
	TamperingDetectedIntrusionSensor IntrusionSensor = "TamperingDetected"
)

// IntrusionSensorReArm is
type IntrusionSensorReArm string

const (

	// ManualIntrusionSensorReArm A manual re-arm of this sensor restores it
	// to the normal state.
	ManualIntrusionSensorReArm IntrusionSensorReArm = "Manual"
	// AutomaticIntrusionSensorReArm Because no abnormal physical security
	// condition is detected, this sensor is automatically restored to the
	// normal state.
	AutomaticIntrusionSensorReArm IntrusionSensorReArm = "Automatic"
)

// PhysicalSecurity shall describe the sensor state of the physical
// security.
type PhysicalSecurity struct {

	// IntrusionSensor is This property shall represent the state of this
	// physical security sensor.  Hardware intrusion indicates the internal
	// hardware is detected as being accessed in an insecure state.
	// Tampering detected indicates the physical tampering of the monitored
	// entity is detected.
	IntrusionSensor IntrusionSensor
	// IntrusionSensorNumber is This property shall contain a numerical
	// identifier for this physical security sensor that is unique within
	// this resource.
	IntrusionSensorNumber int
	// IntrusionSensorReArm is This property shall represent the method that
	// restores this physical security sensor to the normal state.  Manual
	// indicates manual re-arm is needed.  Automatic indicates the state is
	// restored automatically because no abnormal physical security
	// conditions are detected.
	IntrusionSensorReArm IntrusionSensorReArm
}

// Chassis represents the physical components of a system. This
// resource represents the sheet-metal confined spaces and logical zones such
// as racks, enclosures, chassis and all other containers. Subsystems (like sensors)
// that operate outside of a system's data plane (meaning the resources are not
// accessible to software running on the system) are linked either directly or
// indirectly through this resource.
type Chassis struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AssetTag shall contain an identifying string that
	// tracks the chassis for inventory purposes.
	AssetTag string
	// ChassisType shall indicate the physical form factor
	// for the type of chassis.
	ChassisType ChassisType
	// DepthMm shall represent the depth (length) of the
	// chassis, in millimeters, as specified by the manufacturer.
	DepthMm float64
	// Description provides a description of this resource.
	Description string
	// linkedDrives shall contain an array of links to resources
	// of type Drive that are in this chassis.
	linkedDrives []string
	// Drives shall contain a link to a resource collection of type DriveCollection.
	drives string
	// DrivesCount is the number of drives attached to this chassis.
	DrivesCount int `json:"Drives@odata.count"`
	// EnvironmentalClass shall contain the ASHRAE
	// Environmental Class for this chassis, as defined by ASHRAE Thermal
	// Guidelines for Data Processing Environments. These classes define
	// respective environmental limits that include temperature, relative
	// humidity, dew point, and maximum allowable elevation.
	EnvironmentalClass EnvironmentalClass
	// HeightMm shall represent the height of the chassis,
	// in millimeters, as specified by the manufacturer.
	HeightMm float64
	// IndicatorLED shall contain the indicator light state
	// for the indicator light associated with this system.
	IndicatorLED common.IndicatorLED
	// Location shall contain location information of the
	// associated chassis.
	Location common.Location
	// Manufacturer shall contain the name of the
	// organization responsible for producing the chassis. This organization
	// might be the entity from whom the chassis is purchased, but this is
	// not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the
	// manufacturer generally refers to the chassis.
	Model string
	// PartNumber shall contain a part number assigned by
	// the organization that is responsible for producing or manufacturing
	// the chassis.
	PartNumber string
	// PhysicalSecurity shall contain the sensor state of
	// the physical security.
	PhysicalSecurity PhysicalSecurity
	// PowerState shall contain the power state of the
	// chassis.
	PowerState PowerState
	// SKU shall contain the stock-keeping unit number for
	// this chassis.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated
	// number that identifies the chassis.
	SerialNumber string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// UUID shall contain the universal unique identifier
	// number for this chassis.
	UUID string
	// WeightKg shall represent the published mass, commonly
	// referred to as weight, of the chassis, in kilograms.
	WeightKg float64
	// WidthMm shall represent the width of the chassis, in
	// millimeters, as specified by the manufacturer.
	WidthMm float64
	thermal string
	power   string
	// pcieSlots shall be a link to the PCIe slot properties for this chassis
	pcieSlots string
	// sensors shall be a link to to the collection of sensors
	// located in the equipment and sub-components
	sensors         string
	networkAdapters string
	// logServices shall be a link to a collection of type LogServiceCollection.
	logServices     string
	computerSystems []string
	resourceBlocks  []string
	managedBy       []string
	assembly        string
	// resetTarget is the internal URL to send reset actions to.
	resetTarget string
	// SupportedResetTypes, if provided, is the reset types this chassis supports.
	SupportedResetTypes []ResetType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
func (chassis *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	type linkReference struct {
		ComputerSystems common.Links
		ResourceBlocks  common.Links
		ManagedBy       common.Links
		Drives          common.Links
		DrivesCount     int `json:"Drives@odata.count"`
	}
	type Actions struct {
		ChassisReset struct {
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			Target            string
		} `json:"#Chassis.Reset"`
	}

	var t struct {
		temp
		Assembly        common.Link
		Drives          common.Link
		Thermal         common.Link
		Power           common.Link
		PCIeSlots       common.Link
		Sensors         common.Link
		NetworkAdapters common.Link
		LogServices     common.Link
		Links           linkReference
		Actions         Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*chassis = Chassis(t.temp)

	// Extract the links to other entities for later
	chassis.assembly = t.Assembly.String()
	chassis.drives = t.Drives.String()
	chassis.linkedDrives = t.Links.Drives.ToStrings()
	if chassis.DrivesCount == 0 && t.Links.DrivesCount > 0 {
		chassis.DrivesCount = t.Links.DrivesCount
	}
	chassis.thermal = t.Thermal.String()
	chassis.power = t.Power.String()
	chassis.pcieSlots = t.PCIeSlots.String()
	chassis.sensors = t.Sensors.String()
	chassis.networkAdapters = t.NetworkAdapters.String()
	chassis.logServices = t.LogServices.String()
	chassis.computerSystems = t.Links.ComputerSystems.ToStrings()
	chassis.resourceBlocks = t.Links.ResourceBlocks.ToStrings()
	chassis.managedBy = t.Links.ManagedBy.ToStrings()
	chassis.resetTarget = t.Actions.ChassisReset.Target
	chassis.SupportedResetTypes = t.Actions.ChassisReset.AllowedResetTypes

	// This is a read/write object, so we need to save the raw object data for later
	chassis.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (chassis *Chassis) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Chassis)
	err := original.UnmarshalJSON(chassis.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(chassis).Elem()

	return chassis.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetChassis will get a Chassis instance from the Redfish service.
func GetChassis(c common.Client, uri string) (*Chassis, error) {
	var chassis Chassis
	return &chassis, chassis.Get(c, uri, &chassis)
}

// ListReferencedChassis gets the collection of Chassis from a provided reference.
func ListReferencedChassis(c common.Client, link string) ([]*Chassis, error) { //nolint:dupl
	var result []*Chassis
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Chassis
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		chassis, err := GetChassis(c, link)
		ch <- GetResult{Item: chassis, Link: link, Error: err}
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

// Drives gets the drives attached to the storage controllers that this
// resource represents.
func (chassis *Chassis) Drives() ([]*Drive, error) {
	var result []*Drive
	if chassis.drives == "" && len(chassis.linkedDrives) == 0 {
		return result, nil
	}

	// In version v1.2.0 of the spec, Drives were added to the Chassis.Links
	// property. But in v1.14.0 of the spec, Chassis.Drives was added as a
	// direct property.
	// TODO: Update this to use the concurrent collection method
	collectionError := common.NewCollectionError()
	driveLinks := chassis.linkedDrives
	if chassis.drives != "" {
		drives, err := common.GetCollection(chassis.GetClient(), chassis.drives)
		if err != nil {
			collectionError.Failures[chassis.drives] = err
			return nil, collectionError
		}
		driveLinks = drives.ItemLinks
	}

	type GetResult struct {
		Item  *Drive
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	get := func(link string) {
		drive, err := GetDrive(chassis.GetClient(), link)
		ch <- GetResult{Item: drive, Link: link, Error: err}
	}

	go func() {
		common.CollectCollection(get, driveLinks)
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

// Thermal gets the thermal temperature and cooling information for the chassis
func (chassis *Chassis) Thermal() (*Thermal, error) {
	if chassis.thermal == "" {
		return nil, nil
	}

	return GetThermal(chassis.GetClient(), chassis.thermal)
}

// Power gets the power information for the chassis
func (chassis *Chassis) Power() (*Power, error) {
	if chassis.power == "" {
		return nil, nil
	}

	return GetPower(chassis.GetClient(), chassis.power)
}

// PCIeSlots gets the PCIe slots properties for the chassis
func (chassis *Chassis) PCIeSlots() (*PCIeSlots, error) {
	if chassis.pcieSlots == "" {
		return nil, nil
	}

	return GetPCIeSlots(chassis.GetClient(), chassis.pcieSlots)
}

// ComputerSystems returns the collection of systems from this chassis
func (chassis *Chassis) ComputerSystems() ([]*ComputerSystem, error) {
	var result []*ComputerSystem

	collectionError := common.NewCollectionError()
	for _, uri := range chassis.computerSystems {
		cs, err := GetComputerSystem(chassis.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, cs)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ManagedBy gets the collection of managers of this chassis
func (chassis *Chassis) ManagedBy() ([]*Manager, error) {
	var result []*Manager

	collectionError := common.NewCollectionError()
	for _, uri := range chassis.managedBy {
		manager, err := GetManager(chassis.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, manager)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Sensors gets the collection of sensors located in the equipment and sub-components of this chassis
func (chassis *Chassis) Sensors() ([]*Sensor, error) {
	return ListReferencedSensors(chassis.GetClient(), chassis.sensors)
}

// NetworkAdapters gets the collection of network adapters of this chassis
func (chassis *Chassis) NetworkAdapters() ([]*NetworkAdapter, error) {
	return ListReferencedNetworkAdapter(chassis.GetClient(), chassis.networkAdapters)
}

// LogServices get this chassis's log services.
func (chassis *Chassis) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(chassis.GetClient(), chassis.logServices)
}

// The Assembly schema defines an assembly.
// Assembly information contains details about a device, such as part number, serial number, manufacturer, and production date.
// It also provides access to the original data for the assembly.
func (chassis *Chassis) Assembly() (*Assembly, error) {
	return GetAssembly(chassis.GetClient(), chassis.assembly)
}

// Reset shall reset the chassis. This action shall not reset Systems or other
// contained resource, although side effects may occur which affect those resources.
func (chassis *Chassis) Reset(resetType ResetType) error {
	// Make sure the requested reset type is supported by the chassis
	valid := false
	if len(chassis.SupportedResetTypes) > 0 {
		for _, allowed := range chassis.SupportedResetTypes {
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
		return fmt.Errorf("reset type '%s' is not supported by this chassis",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	return chassis.Post(chassis.resetTarget, t)
}
