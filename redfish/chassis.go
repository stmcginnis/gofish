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
	// HeatExchangerChassisType A heat exchanger.
	HeatExchangerChassisType ChassisType = "HeatExchanger"
	// ImmersionTankChassisType An immersion cooling tank.
	ImmersionTankChassisType ChassisType = "ImmersionTank"
	// IPBasedDriveChassisType is a chassis in a drive form factor with IP-based network connections.
	IPBasedDriveChassisType ChassisType = "IPBasedDrive"
	// ModuleChassisType is a small, typically removable, chassis or card which contains devices
	// for a particular subsystem or function.
	ModuleChassisType ChassisType = "Module"
	// OtherChassisType is a chassis that does not fit any of these definitions.
	OtherChassisType ChassisType = "Other"
	// PodChassisType is a collection of equipment racks in a large, likely transportable, container.
	PodChassisType ChassisType = "Pod"
	// PowerStripChassisType A power strip, typically placed in the zero-U space of a rack.
	PowerStripChassisType ChassisType = "PowerStrip"
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

// Door shall describe a door or access panel on the chassis.
type Door struct {
	// DoorState shall contain the current state of the door.
	DoorState DoorState
	// Locked shall indicate if the door is locked.
	Locked bool
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
}

// Doors shall describe the doors or access panels of the chassis.
type Doors struct {
	// Front shall contain information related to the front door, as defined by the manufacturer, of the chassis.
	Front Door
	// Rear shall contain information related to the rear door, as defined by the manufacturer, of the chassis.
	Rear Door
}

type DoorState string

const (
	// LockedDoorState shall indicate that the door is both closed and locked. In this state, the door cannot be opened
	// unless the value of the Locked property is set to 'false'.
	LockedDoorState DoorState = "Locked"
	// ClosedDoorState shall indicate that the door is closed but unlocked.
	ClosedDoorState DoorState = "Closed"
	// LockedAndOpenDoorState shall indicate that the door is open but the lock has been engaged. It may be possible to
	// close the door while in this state.
	LockedAndOpenDoorState DoorState = "LockedAndOpen"
	// OpenDoorState shall indicate that the door is open.
	OpenDoorState DoorState = "Open"
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

type ThermalDirection string

const (
	// FrontToBackThermalDirection shall indicate a chassis with the air intake generally from the front of the chassis
	// and the air exhaust out the back of the chassis.
	FrontToBackThermalDirection ThermalDirection = "FrontToBack"
	// BackToFrontThermalDirection shall indicate a chassis with the air intake generally from the back of the chassis
	// and the air exhaust out the front of the chassis.
	BackToFrontThermalDirection ThermalDirection = "BackToFront"
	// TopExhaustThermalDirection shall indicate a chassis with the air exhaust out the top of the chassis.
	TopExhaustThermalDirection ThermalDirection = "TopExhaust"
	// SealedThermalDirection shall indicate a sealed chassis with no air pathway through the chassis.
	SealedThermalDirection ThermalDirection = "Sealed"
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
	assembly  string
	// AssetTag shall contain an identifying string that
	// tracks the chassis for inventory purposes.
	AssetTag string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates string
	// ChassisType shall indicate the physical form factor
	// for the type of chassis.
	ChassisType ChassisType
	controls    string
	// DepthMm shall represent the depth (length) of the
	// chassis, in millimeters, as specified by the manufacturer.
	DepthMm float64
	// Description provides a description of this resource.
	Description string
	// linkedDrives shall contain an array of links to resources
	// of type Drive that are in this chassis.
	linkedDrives []string
	// Doors shall contain information about the doors or access panels of the chassis.
	// Added in v1.24.0.
	Doors Doors
	// Drives shall contain a link to a resource collection of type DriveCollection.
	drives string
	// DrivesCount is the number of drives attached to this chassis.
	DrivesCount int `json:"Drives@odata.count"`
	// ElectricalSourceManagerURIs shall contain an array of URIs to the management applications or devices that
	// provide monitoring or control of the external electrical sources that provide power to this chassis.
	ElectricalSourceManagerURIs []string
	// ElectricalSourceNames shall contain an array of strings that identify the external electrical sources, such as
	// the names of circuits or outlets, that provide power to this chassis.
	ElectricalSourceNames []string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this chassis.
	environmentMetrics string
	// EnvironmentalClass shall contain the ASHRAE
	// Environmental Class for this chassis, as defined by ASHRAE Thermal
	// Guidelines for Data Processing Environments. These classes define
	// respective environmental limits that include temperature, relative
	// humidity, dew point, and maximum allowable elevation.
	EnvironmentalClass EnvironmentalClass
	// FabricAdapters shall contain a link to a resource collection of type FabricAdapterCollection that represents
	// fabric adapters in this chassis that provide access to fabric-related resource pools.
	fabricAdapters string
	// HeatingCoolingEquipmentNames shall contain an array of strings that identify the external heating or cooling
	// equipment, such as the names of specific coolant distribution units, that provide thermal management for this
	// chassis.
	HeatingCoolingEquipmentNames []string
	// HeatingCoolingManagerURIs shall contain an array of URIs to the management applications or devices that provide
	// monitoring or control of the external heating or cooling equipment that provide thermal management for this
	// chassis.
	HeatingCoolingManagerURIs []string
	// HeightMm shall represent the height of the chassis,
	// in millimeters, as specified by the manufacturer.
	HeightMm float64
	// HotPluggable shall indicate whether the component can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Components indicated as hot-pluggable shall allow the
	// component to become operable without altering the operational state of the underlying equipment. Components that
	// cannot be inserted or removed from equipment in operation, or components that cannot become operable without
	// affecting the operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the indicator light state
	// for the indicator light associated with this system.
	// Deprecated v1.14+ in favor of LocationIndicatorActive property
	IndicatorLED common.IndicatorLED
	// Location shall contain location information of the
	// associated chassis.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function. Modifying this property may modify the
	// LocationIndicatorActive in the resource that represents the functional view of this Chassis, such as a
	// ComputerSystem resource.
	LocationIndicatorActive bool
	logServices             string
	// Manufacturer shall contain the name of the
	// organization responsible for producing the chassis. This organization
	// might be the entity from whom the chassis is purchased, but this is
	// not necessarily true.
	Manufacturer string
	// MaxPowerWatts shall contain the upper bound of the total power consumed by the chassis.
	MaxPowerWatts float64
	// Memory shall contain a link to a resource collection of type MemoryCollection that represents memory in this
	// chassis that belong to fabric-related resource pools.
	memory string
	// MemoryDomains shall contain a link to a resource collection of type MemoryDomainCollection that represents
	// memory domains in this chassis that belong to fabric-related resource pools.
	memoryDomains string
	// MinPowerWatts shall contain the lower bound of the total power consumed by the chassis.
	MinPowerWatts float64
	// Model shall contain the name by which the
	// manufacturer generally refers to the chassis.
	Model           string
	networkAdapters string
	// PartNumber shall contain a part number assigned by
	// the organization that is responsible for producing or manufacturing
	// the chassis.
	PartNumber string
	// PCIeDevices shall contain a link to a resource collection of type PCIeDeviceCollection.
	pcieDevices string
	// PhysicalSecurity shall contain the sensor state of
	// the physical security.
	PhysicalSecurity PhysicalSecurity
	// PowerState shall contain the power state of the
	// chassis.
	PowerState PowerState
	// PowerSubsystem shall contain a link to a resource of type PowerSubsystem that represents the power subsystem
	// information for this chassis.
	powerSubsystem string
	// PoweredByParent shall indicate whether the chassis receives power from the chassis that contains it. The value
	// 'true' shall indicate that the containing chassis provides power. The value 'false' shall indicate the chassis
	// receives power from its own power subsystem, another chassis instance's power supplies, or outlets.
	PoweredByParent bool
	// Processors shall contain a link to a resource collection of type ProcessorCollection
	// that represents processors in this chassis that belong to fabric-related resource pools.
	processors string
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains 'Embedded', this property shall contain
	// 'false'.
	Replaceable bool
	// SKU shall contain the stock-keeping unit number for
	// this chassis.
	SKU     string
	sensors string
	// SerialNumber shall contain a manufacturer-allocated
	// number that identifies the chassis.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the chassis.
	SparePartNumber string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// ThermalDirection shall indicate the general direction of the thermal management path through the chassis.
	ThermalDirection ThermalDirection
	// ThermalManagedByParent shall indicate whether the chassis relies on the containing chassis to provide thermal
	// management. The value 'true' shall indicate that the chassis relies on the containing chassis to provide thermal
	// management. The value 'false' shall indicate the chassis provides thermal management, and may provide details in
	// a ThermalSubsystem resource, or by populating the Fans property in Links.
	ThermalManagedByParent bool
	// ThermalSubsystem shall contain a link to a resource of type ThermalSubsystem that represents the thermal
	// subsystem information for this chassis.
	thermalSubsystem string
	// TrustedComponents shall contain a link to a resource collection of type TrustedComponentCollection.
	trustedComponents string
	// UUID shall contain the universal unique identifier
	// number for this chassis.
	UUID string
	// WeightKg shall represent the published mass, commonly
	// referred to as weight, of the chassis, in kilograms.
	WeightKg float64
	// WidthMm shall represent the width of the chassis, in
	// millimeters, as specified by the manufacturer.
	WidthMm float64
	// rawData holds the original serialized JSON so we can compare updates.
	RawData []byte

	// Deprecated properties

	power     string
	thermal   string
	pcieSlots string

	// Action links

	// resetTarget is the internal URL to send reset actions to.
	resetTarget string
	// SupportedResetTypes, if provided, is the reset types this chassis supports.
	SupportedResetTypes []ResetType

	// Links

	cables []string
	// CablesCount is the number of connected cables
	CablesCount     int
	computerSystems []string
	// ComputerSystemsCount is the number of computer systems in this chassis.
	ComputerSystemsCount  int
	connectedCoolingLoops []string
	// ConnectedCoolingLoopsCount is the number of cooling loops connected to this chassis.
	ConnectedCoolingLoopsCount int
	containedBy                string
	contains                   []string
	// ContainsCount is the number of chassis this chassis contains.
	ContainsCount int
	coolingUnits  []string
	// CoolingUnitsCount is the number of cooling units connected to this chassis.
	CoolingUnitsCount int
	facility          string
	fans              []string
	// FansCount is the number of fans connected to this chassis.
	FansCount int
	managedBy []string
	// ManagedByCount is the number of managers for this chassis.
	ManagedByCount    int
	managersInChassis []string
	// ManagersInChassisCount is the number of managers contained in this chassis.
	ManagersInChassisCount int
	powerDistribution      string
	powerOutlets           []string
	// PowerOutletsCount is the number of power outlets in this chassis.
	PowerOutletsCount int
	powerSupplies     []string
	// PowerSuppliesCount gets the number of power supplies in the chassis.
	PowerSuppliesCount int
	resourceBlocks     []string
	// ResourceBlocksCount is the number of resource blocks in this chassis.
	ResourceBlockCount int
	storage            []string
	// StorageCount is the number of storage instances connected to this chassis.
	StorageCount int
	switches     []string
	// SwitchesCount is the number of switches in this chassis.
	SwitchesCount int
}

type chassisLinks struct {
	Cables                     common.Links
	CablesCount                int `json:"Cables@odata.count"`
	ComputerSystems            common.Links
	ComputerSystemsCount       int `json:"ComputerSystems@odata.count"`
	ConnectedCoolingLoops      common.Links
	ConnectedCoolingLoopsCount int `json:"ConnectedCoolingLoops@odata.count"`
	ContainedBy                common.Link
	Contains                   common.Links
	ContainsCount              int `json:"Contains@odata.count"`
	CoolingUnits               common.Links
	CoolingUnitsCount          int `json:"CoolingUnits@odata.count"`
	Drives                     common.Links
	DrivesCount                int `json:"Drives@odata.count"`
	Facility                   common.Link
	Fans                       common.Links
	FansCount                  int `json:"Fans@odata.count"`
	ManagedBy                  common.Links
	ManagedByCount             int `json:"ManagedBy@odata.count"`
	ManagersInChassis          common.Links
	ManagersInChassisCount     int `json:"ManagersInChassis@odata.count"`
	PowerDistribution          common.Link
	PowerOutlets               common.Links
	PowerOutletsCount          int `json:"PowerOutlets@odata.count"`
	PowerSupplies              common.Links
	PowerSuppliesCount         int `json:"PowerSupplies@odata.count"`
	Processors                 common.Links
	ProcessorsCount            int `json:"Processors@odata.count"`
	ResourceBlocks             common.Links
	ResourceBlocksCount        int `json:"ResourceBlocks@odata.count"`
	Storage                    common.Links
	StorageCount               int `json:"Storage@odata.count"`
	Switches                   common.Links
	SwitchesCount              int `json:"Switches@odata.count"`
}

type chassisActions struct {
	ChassisReset struct {
		AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		Target            string
	} `json:"#Chassis.Reset"`
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
//
//nolint:funlen
func (chassis *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis

	var t struct {
		temp
		Assembly           common.Link
		Certificates       common.Link
		Controls           common.Link
		Drives             common.Link
		EnvironmentMetrics common.Link
		FabricAdapters     common.Link
		LogServices        common.Link
		Memory             common.Link
		MemoryDomains      common.Link
		NetworkAdapters    common.Link
		PCIeDevices        common.Link
		PCIeSlots          common.Link
		Power              common.Link
		PowerSubsystem     common.Link
		Processors         common.Link
		Sensors            common.Link
		Thermal            common.Link
		ThermalSubsystem   common.Link
		TrustedComponents  common.Link
		Links              chassisLinks
		Actions            chassisActions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*chassis = Chassis(t.temp)

	// Extract the links to other entities for later
	chassis.assembly = t.Assembly.String()
	chassis.certificates = t.Certificates.String()
	chassis.controls = t.Controls.String()
	chassis.drives = t.Drives.String()
	chassis.environmentMetrics = t.EnvironmentMetrics.String()
	chassis.fabricAdapters = t.FabricAdapters.String()
	chassis.logServices = t.LogServices.String()
	chassis.memory = t.Memory.String()
	chassis.memoryDomains = t.MemoryDomains.String()
	chassis.networkAdapters = t.NetworkAdapters.String()
	chassis.pcieDevices = t.PCIeDevices.String()
	chassis.pcieSlots = t.PCIeSlots.String()
	chassis.power = t.Power.String()
	chassis.powerSubsystem = t.PowerSubsystem.String()
	chassis.processors = t.Processors.String()
	chassis.sensors = t.Sensors.String()
	chassis.thermal = t.Thermal.String()
	chassis.thermalSubsystem = t.ThermalSubsystem.String()
	chassis.trustedComponents = t.TrustedComponents.String()

	chassis.resetTarget = t.Actions.ChassisReset.Target
	chassis.SupportedResetTypes = t.Actions.ChassisReset.AllowedResetTypes

	chassis.cables = t.Links.Cables.ToStrings()
	chassis.computerSystems = t.Links.ComputerSystems.ToStrings()
	chassis.ComputerSystemsCount = t.Links.ComputerSystemsCount
	chassis.connectedCoolingLoops = t.Links.ConnectedCoolingLoops.ToStrings()
	chassis.ConnectedCoolingLoopsCount = t.Links.ConnectedCoolingLoopsCount
	chassis.containedBy = t.Links.ContainedBy.String()
	chassis.contains = t.Links.Contains.ToStrings()
	chassis.ContainsCount = t.Links.ContainsCount
	chassis.coolingUnits = t.Links.CoolingUnits.ToStrings()
	chassis.CoolingUnitsCount = t.Links.CoolingUnitsCount
	chassis.facility = t.Links.Facility.String()
	chassis.fans = t.Links.Fans.ToStrings()
	chassis.FansCount = t.Links.FansCount
	chassis.managedBy = t.Links.ManagedBy.ToStrings()
	chassis.ManagedByCount = t.Links.ManagedByCount
	chassis.managersInChassis = t.Links.ManagersInChassis.ToStrings()
	chassis.ManagersInChassisCount = t.Links.ManagersInChassisCount
	chassis.powerDistribution = t.Links.PowerDistribution.String()
	chassis.powerOutlets = t.Links.PowerOutlets.ToStrings()
	chassis.PowerOutletsCount = t.Links.PowerOutletsCount
	chassis.powerSupplies = t.Links.PowerSupplies.ToStrings()
	chassis.PowerSuppliesCount = t.Links.PowerSuppliesCount
	// Links contain processors, but so does the chassis object. Since the chassis
	// object includes a collection link that is a little more efficient, sticking
	// with that until there is a need identified to grab it a different way.
	chassis.resourceBlocks = t.Links.ResourceBlocks.ToStrings()
	chassis.storage = t.Links.Storage.ToStrings()
	chassis.StorageCount = t.Links.StorageCount
	chassis.switches = t.Links.Switches.ToStrings()
	chassis.SwitchesCount = t.Links.SwitchesCount
	chassis.linkedDrives = t.Links.Drives.ToStrings()
	if chassis.DrivesCount == 0 && t.Links.DrivesCount > 0 {
		chassis.DrivesCount = t.Links.DrivesCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	chassis.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (chassis *Chassis) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Chassis)
	err := original.UnmarshalJSON(chassis.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"IndicatorLED",
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"EnvironmentalClass",
		"HeatingCoolingEquipmentNames",
		"HeatingCoolingManagerURIs",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(chassis).Elem()

	return chassis.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetChassis will get a Chassis instance from the Redfish service.
func GetChassis(c common.Client, uri string) (*Chassis, error) {
	return common.GetObject[Chassis](c, uri)
}

// ListReferencedChassis gets the collection of Chassis from a provided reference.
func ListReferencedChassis(c common.Client, link string) ([]*Chassis, error) {
	return common.GetCollectionObjects[Chassis](c, link)
}

// Certificates returns certificates in this Chassis.
func (chassis *Chassis) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(chassis.GetClient(), chassis.certificates)
}

// Controls returns controls in this Chassis.
func (chassis *Chassis) Controls() ([]*Control, error) {
	return ListReferencedControls(chassis.GetClient(), chassis.controls)
}

// ConnectedCoolingLoops returns a collection of CoolingLoops associated with this Chassis.
func (chassis *Chassis) ConnectedCoolingLoops() ([]*CoolingLoop, error) {
	return common.GetObjects[CoolingLoop](chassis.GetClient(), chassis.connectedCoolingLoops)
}

// CoolingUnits returns a collection of CoolingUnits associated with this Chassis.
func (chassis *Chassis) CoolingUnits() ([]*CoolingUnit, error) {
	return common.GetObjects[CoolingUnit](chassis.GetClient(), chassis.coolingUnits)
}

// Drives gets the drives attached to the storage controllers that this
// resource represents.
func (chassis *Chassis) Drives() ([]*Drive, error) {
	// In version v1.2.0 of the spec, Drives were added to the Chassis.Links
	// property. But in v1.14.0 of the spec, Chassis.Drives was added as a
	// direct property.
	if chassis.drives != "" {
		return common.GetCollectionObjects[Drive](chassis.GetClient(), chassis.drives)
	}

	if len(chassis.linkedDrives) > 0 {
		return common.GetObjects[Drive](chassis.GetClient(), chassis.linkedDrives)
	}

	return []*Drive{}, nil
}

// EnvironmentMetrics gets the environment metrics for the chassis
func (chassis *Chassis) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if chassis.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(chassis.GetClient(), chassis.environmentMetrics)
}

// Facility gets the smallest facility that contains the chassis.
func (chassis *Chassis) Facility() (*Facility, error) {
	if chassis.facility == "" {
		return nil, nil
	}

	return GetFacility(chassis.GetClient(), chassis.facility)
}

// Memory gets the memory in the chassis.
func (chassis *Chassis) Memory() ([]*Memory, error) {
	return ListReferencedMemorys(chassis.GetClient(), chassis.memory)
}

// MemoryDomains gets the memory domains in the chassis.
func (chassis *Chassis) MemoryDomains() ([]*MemoryDomain, error) {
	return ListReferencedMemoryDomains(chassis.GetClient(), chassis.memoryDomains)
}

// Thermal gets the thermal temperature and cooling information for the chassis
// This link has been deprecated in favor of the ThermalSubsystem link property.
func (chassis *Chassis) Thermal() (*Thermal, error) {
	if chassis.thermal == "" {
		return nil, nil
	}

	return GetThermal(chassis.GetClient(), chassis.thermal)
}

// ThermalSubsystem gets the thermal subsystem for this chassis.
func (chassis *Chassis) ThermalSubsystem() (*ThermalSubsystem, error) {
	if chassis.thermalSubsystem == "" {
		return nil, nil
	}

	return GetThermalSubsystem(chassis.GetClient(), chassis.thermalSubsystem)
}

// PCIeDevices gets the PCIe devices in the chassis
func (chassis *Chassis) PCIeDevices() ([]*PCIeDevice, error) {
	return ListReferencedPCIeDevices(chassis.GetClient(), chassis.pcieDevices)
}

// PCIeSlots gets the PCIe slots properties for the chassis.
// This property has been deprecated in favor of the PCIeDevices property. The PCIeSlots schema has been
// deprecated in favor of the PCIeDevice schema. Empty PCIe slots are represented by PCIeDevice resources
// using the `Absent` value of the State property within Status.
func (chassis *Chassis) PCIeSlots() (*PCIeSlots, error) {
	if chassis.pcieSlots == "" {
		return nil, nil
	}

	return GetPCIeSlots(chassis.GetClient(), chassis.pcieSlots)
}

// Power gets the power information for the chassis
// This link has been deprecated in favor of the PowerSubsystem link property.
func (chassis *Chassis) Power() (*Power, error) {
	if chassis.power == "" {
		return nil, nil
	}

	return GetPower(chassis.GetClient(), chassis.power)
}

// PowerSubsystem gets the power subsystem for the chassis
// This link has been deprecated in favor of the PowerSubsystem link property.
func (chassis *Chassis) PowerSubsystem() (*PowerSubsystem, error) {
	if chassis.powerSubsystem == "" {
		return nil, nil
	}

	return GetPowerSubsystem(chassis.GetClient(), chassis.powerSubsystem)
}

// Cables gets the connected cables.
func (chassis *Chassis) Cables() ([]*Cable, error) {
	return common.GetObjects[Cable](chassis.GetClient(), chassis.cables)
}

// ComputerSystems returns the collection of systems from this chassis
func (chassis *Chassis) ComputerSystems() ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](chassis.GetClient(), chassis.computerSystems)
}

// ContainedBy gets the chassis that contains this chassis. The result is nil
// if this chassis is not contained by another one.
func (chassis *Chassis) ContainedBy() (*Chassis, error) {
	if chassis.containedBy == "" {
		return nil, nil
	}
	return GetChassis(chassis.GetClient(), chassis.containedBy)
}

// Contains gets the chassis instances that this chassis contains.
func (chassis *Chassis) Contains() ([]*Chassis, error) {
	return common.GetObjects[Chassis](chassis.GetClient(), chassis.contains)
}

// Fans gets the the fans that provide cooling to this chassis. This property shall not be present if the
// ThermalManagedByParent property contains `true` or if the fans are contained in the ThermalSubsystem
// resource for this chassis.
func (chassis *Chassis) Fans() ([]*Fan, error) {
	return common.GetObjects[Fan](chassis.GetClient(), chassis.fans)
}

// ManagedBy gets the collection of managers of this chassis
func (chassis *Chassis) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](chassis.GetClient(), chassis.managedBy)
}

// ManagersInChassis gets the managers contained in this chassis.
func (chassis *Chassis) ManagersInChassis() ([]*Manager, error) {
	return common.GetObjects[Manager](chassis.GetClient(), chassis.managersInChassis)
}

// PowerDistribution gets the power distribution functionality contained within this chassis.
func (chassis *Chassis) PowerDistribution() (*PowerDistribution, error) {
	if chassis.powerDistribution == "" {
		return nil, nil
	}
	return GetPowerDistribution(chassis.GetClient(), chassis.powerDistribution)
}

// PowerOutlets gets the power outlets in this chassis.
func (chassis *Chassis) PowerOutlets() ([]*Outlet, error) {
	return common.GetObjects[Outlet](chassis.GetClient(), chassis.powerOutlets)
}

// PowerSupplies gets the power supplies that provide power to this chassis. This property shall not be
// present if the PoweredByParent property contains 'true' or if the power supplies are contained in the
// PowerSubsystem resource for this chassis.
func (chassis *Chassis) PowerSupplies() ([]*PowerSupply, error) {
	return common.GetObjects[PowerSupply](chassis.GetClient(), chassis.powerSupplies)
}

// Processors returns the collection of systems from this chassis.
// Added in v1.25.0.
func (chassis *Chassis) Processors() ([]*Processor, error) {
	return ListReferencedProcessors(chassis.GetClient(), chassis.processors)
}

// ResourceBlocks gets the resource blocks located in this chassis.
func (chassis *Chassis) ResourceBlocks() ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](chassis.GetClient(), chassis.resourceBlocks)
}

// Storage gets the storage subsystems connected to or inside this chassis.
func (chassis *Chassis) Storage() ([]*Storage, error) {
	return common.GetObjects[Storage](chassis.GetClient(), chassis.storage)
}

// Sensors gets the collection of sensors located in the equipment and sub-components of this chassis
func (chassis *Chassis) Sensors() ([]*Sensor, error) {
	return ListReferencedSensors(chassis.GetClient(), chassis.sensors)
}

// Switches gets the switches in this chassis.
func (chassis *Chassis) Switches() ([]*Switch, error) {
	return common.GetObjects[Switch](chassis.GetClient(), chassis.switches)
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
