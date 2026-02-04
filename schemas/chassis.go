//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Chassis.v1_28_0.json
// 2025.3 - #Chassis.v1_28_0.Chassis

package schemas

import (
	"encoding/json"
	"errors"
)

type ChassisType string

const (
	// RackChassisType is an equipment rack, typically a 19-inch wide freestanding
	// unit.
	RackChassisType ChassisType = "Rack"
	// BladeChassisType is an enclosed or semi-enclosed, typically
	// vertically-oriented, system chassis that must be plugged into a multi-system
	// chassis to function normally.
	BladeChassisType ChassisType = "Blade"
	// EnclosureChassisType is a generic term for a chassis that does not fit any
	// other description.
	EnclosureChassisType ChassisType = "Enclosure"
	// StandAloneChassisType is a single, free-standing system, commonly called a
	// tower or desktop chassis.
	StandAloneChassisType ChassisType = "StandAlone"
	// RackMountChassisType is a single-system chassis designed specifically for
	// mounting in an equipment rack.
	RackMountChassisType ChassisType = "RackMount"
	// CardChassisType is a loose device or circuit board intended to be installed
	// in a system or other enclosure.
	CardChassisType ChassisType = "Card"
	// CartridgeChassisType is a small self-contained system intended to be plugged
	// into a multi-system chassis.
	CartridgeChassisType ChassisType = "Cartridge"
	// RowChassisType is a collection of equipment racks.
	RowChassisType ChassisType = "Row"
	// PodChassisType is a collection of equipment racks in a large, likely
	// transportable, container.
	PodChassisType ChassisType = "Pod"
	// ExpansionChassisType is a chassis that expands the capabilities or capacity
	// of another chassis.
	ExpansionChassisType ChassisType = "Expansion"
	// SidecarChassisType is a chassis that mates mechanically with another chassis
	// to expand its capabilities or capacity.
	SidecarChassisType ChassisType = "Sidecar"
	// ZoneChassisType is a logical division or portion of a physical chassis that
	// contains multiple devices or systems that cannot be physically separated.
	ZoneChassisType ChassisType = "Zone"
	// SledChassisType is an enclosed or semi-enclosed, system chassis that must be
	// plugged into a multi-system chassis to function normally similar to a blade
	// type chassis.
	SledChassisType ChassisType = "Sled"
	// ShelfChassisType is an enclosed or semi-enclosed, typically
	// horizontally-oriented, system chassis that must be plugged into a
	// multi-system chassis to function normally.
	ShelfChassisType ChassisType = "Shelf"
	// DrawerChassisType is an enclosed or semi-enclosed, typically
	// horizontally-oriented, system chassis that can be slid into a multi-system
	// chassis.
	DrawerChassisType ChassisType = "Drawer"
	// ModuleChassisType is a small, typically removable, chassis or card that
	// contains devices for a particular subsystem or function.
	ModuleChassisType ChassisType = "Module"
	// ComponentChassisType is a small chassis, card, or device that contains
	// devices for a particular subsystem or function.
	ComponentChassisType ChassisType = "Component"
	// IPBasedDriveChassisType is a chassis in a drive form factor with IP-based
	// network connections.
	IPBasedDriveChassisType ChassisType = "IPBasedDrive"
	// RackGroupChassisType is a group of racks that form a single entity or share
	// infrastructure.
	RackGroupChassisType ChassisType = "RackGroup"
	// StorageEnclosureChassisType is a chassis that encloses storage.
	StorageEnclosureChassisType ChassisType = "StorageEnclosure"
	// ImmersionTankChassisType is an immersion cooling tank.
	ImmersionTankChassisType ChassisType = "ImmersionTank"
	// HeatExchangerChassisType is a heat exchanger.
	HeatExchangerChassisType ChassisType = "HeatExchanger"
	// PowerStripChassisType is a power strip, typically placed in the zero-U space
	// of a rack.
	PowerStripChassisType ChassisType = "PowerStrip"
	// OtherChassisType is a chassis that does not fit any of these definitions.
	OtherChassisType ChassisType = "Other"
)

type DoorState string

const (
	// LockedDoorState shall indicate that the door is both closed and locked. In
	// this state, the door cannot be opened unless the value of the 'Locked'
	// property is set to 'false'.
	LockedDoorState DoorState = "Locked"
	// ClosedDoorState shall indicate that the door is closed but unlocked.
	ClosedDoorState DoorState = "Closed"
	// LockedAndOpenDoorState shall indicate that the door is open but the lock has
	// been engaged. It may be possible to close the door while in this state.
	LockedAndOpenDoorState DoorState = "LockedAndOpen"
	// OpenDoorState shall indicate that the door is open.
	OpenDoorState DoorState = "Open"
)

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
	// NormalIntrusionSensor No physical security condition is detected at this
	// time.
	NormalIntrusionSensor IntrusionSensor = "Normal"
	// HardwareIntrusionIntrusionSensor is a door, lock, or other mechanism
	// protecting the internal system hardware from being accessed is detected to
	// be in an insecure state.
	HardwareIntrusionIntrusionSensor IntrusionSensor = "HardwareIntrusion"
	// TamperingDetectedIntrusionSensor Physical tampering of the monitored entity
	// is detected.
	TamperingDetectedIntrusionSensor IntrusionSensor = "TamperingDetected"
)

type IntrusionSensorReArm string

const (
	// ManualIntrusionSensorReArm shall indicate a user is required to set the
	// 'IntrusionSensor' property to 'Normal' to restore the sensor to its normal
	// state.
	ManualIntrusionSensorReArm IntrusionSensorReArm = "Manual"
	// AutomaticIntrusionSensorReArm shall indicate the service sets the
	// 'IntrusionSensor' property to 'Normal' when no security condition is
	// detected.
	AutomaticIntrusionSensorReArm IntrusionSensorReArm = "Automatic"
)

type RackMountWidth string

const (
	// OpenURackMountWidth shall conform to the Open Rack standard for 21-inch
	// racks.
	OpenURackMountWidth RackMountWidth = "OpenU"
	// EIA310RackMountWidth shall conform to the EIA-310 standard for 19-inch
	// racks.
	EIA310RackMountWidth RackMountWidth = "EIA_310"
	// EIA310TelcoRackMountWidth shall conform to the EIA-310 standard for 23-inch
	// telecommunications equipment racks.
	EIA310TelcoRackMountWidth RackMountWidth = "EIA_310_Telco"
	// HalfWidthRackMountWidth shall be approximately 9.5 inches in width,
	// following de facto industry practice for a rack half the width of a 19-inch
	// EIA-310 equipment racks.
	HalfWidthRackMountWidth RackMountWidth = "HalfWidth"
)

type ThermalDirection string

const (
	// FrontToBackThermalDirection shall indicate a chassis with the air intake
	// generally from the front of the chassis and the air exhaust out the back of
	// the chassis.
	FrontToBackThermalDirection ThermalDirection = "FrontToBack"
	// BackToFrontThermalDirection shall indicate a chassis with the air intake
	// generally from the back of the chassis and the air exhaust out the front of
	// the chassis.
	BackToFrontThermalDirection ThermalDirection = "BackToFront"
	// TopExhaustThermalDirection shall indicate a chassis with the air exhaust out
	// the top of the chassis.
	TopExhaustThermalDirection ThermalDirection = "TopExhaust"
	// SealedThermalDirection shall indicate a sealed chassis with no air pathway
	// through the chassis.
	SealedThermalDirection ThermalDirection = "Sealed"
)

// Chassis shall represent a chassis or other physical enclosure for a Redfish
// implementation. It may also represent a location, such as a slot, socket, or
// bay, where a unit may be installed, but the 'State' property within the
// 'Status' property contains 'Absent'.
type Chassis struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.6.0
	assembly string
	// AssetTag shall contain an identifying string that tracks the chassis for
	// inventory purposes. Modifying this property may modify the 'AssetTag' in the
	// resource that represents the functional view of this chassis, such as a
	// 'ComputerSystem' resource.
	AssetTag string
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.15.0
	certificates string
	// ChassisType shall indicate the physical form factor for the type of chassis.
	ChassisType ChassisType
	// Controls shall contain a link to a resource collection of type
	// 'ControlCollection'.
	//
	// Version added: v1.17.0
	controls string
	// DepthMm shall represent the depth (length) of the chassis, in millimeter
	// units, as specified by the manufacturer.
	//
	// Version added: v1.4.0
	DepthMm *float64 `json:",omitempty"`
	// Doors shall contain information about the doors or access panels of the
	// chassis.
	//
	// Version added: v1.24.0
	Doors Doors
	// ElectricalSourceManagerURIs shall contain an array of URIs to the management
	// applications or devices that provide monitoring or control of the external
	// electrical sources that provide power to this chassis.
	//
	// Version added: v1.18.0
	ElectricalSourceManagerURIs []string
	// ElectricalSourceNames shall contain an array of strings that identify the
	// external electrical sources, such as the names of circuits or outlets, that
	// provide power to this chassis.
	//
	// Version added: v1.18.0
	ElectricalSourceNames []string
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this
	// chassis, all containing chassis, and devices contained by any of these
	// chassis instances. When determining power and energy readings, care should
	// be taken to ensure any reported values do not overlap or result in
	// double-counting.
	//
	// Version added: v1.15.0
	environmentMetrics string
	// EnvironmentalClass shall contain the ASHRAE Environmental Class for this
	// chassis, as defined by ASHRAE Thermal Guidelines for Data Processing
	// Environments. These classes define respective environmental limits that
	// include temperature, relative humidity, dew point, and maximum allowable
	// elevation.
	//
	// Version added: v1.9.0
	EnvironmentalClass EnvironmentalClass
	// FabricAdapters shall contain a link to a resource collection of type
	// 'FabricAdapterCollection' that represents fabric adapters in this chassis
	// that provide access to fabric-related resource pools.
	//
	// Version added: v1.20.0
	fabricAdapters string
	// HeatingCoolingEquipmentNames shall contain an array of strings that identify
	// the external heating or cooling equipment, such as the names of specific
	// coolant distribution units, that provide thermal management for this
	// chassis.
	//
	// Version added: v1.25.0
	HeatingCoolingEquipmentNames []string
	// HeatingCoolingManagerURIs shall contain an array of URIs to the management
	// applications or devices that provide monitoring or control of the external
	// heating or cooling equipment that provide thermal management for this
	// chassis.
	//
	// Version added: v1.25.0
	HeatingCoolingManagerURIs []string
	// HeightMm shall represent the height of the chassis, in millimeter units, as
	// specified by the manufacturer.
	//
	// Version added: v1.4.0
	HeightMm *float64 `json:",omitempty"`
	// HeightRackUnits shall contain the height of the rack-mountable chassis, in
	// rack units specified by the value of 'RackUnits'.
	//
	// Version added: v1.28.0
	HeightRackUnits *float64 `json:",omitempty"`
	// HotPluggable shall indicate whether the component can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Components indicated as hot-pluggable shall allow the component to
	// become operable without altering the operational state of the underlying
	// equipment. Components that cannot be inserted or removed from equipment in
	// operation, or components that cannot become operable without affecting the
	// operational state of that equipment, shall be indicated as not
	// hot-pluggable.
	//
	// Version added: v1.21.0
	HotPluggable bool
	// IndicatorLED shall contain the indicator light state for the indicator light
	// associated with this system.
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of the 'LocationIndicatorActive'
	// property.
	IndicatorLED IndicatorLED
	// LeakDetectors shall contain a link to a resource collection of type
	// 'LeakDetectorCollection'.
	//
	// Version added: v1.26.0
	leakDetectors string
	// Location shall contain the location information of the associated chassis.
	//
	// Version added: v1.2.0
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function. Modifying this property
	// may modify the 'LocationIndicatorActive' in the resource that represents the
	// functional view of this chassis, such as a 'ComputerSystem' resource.
	//
	// Version added: v1.14.0
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type
	// 'LogServiceCollection'.
	logServices string
	// Manufacturer shall contain the name of the organization responsible for
	// producing the chassis. This organization may be the entity from whom the
	// chassis is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxPowerWatts shall contain the upper bound of the total power consumed by
	// the chassis.
	//
	// Version added: v1.12.0
	MaxPowerWatts *float64 `json:",omitempty"`
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.15.0
	//
	// Deprecated: v1.19.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// MediaControllers shall contain a link to a resource collection of type
	// 'MediaControllerCollection'.
	//
	// Version added: v1.11.0
	//
	// Deprecated: v1.20.0
	// This property has been deprecated in favor of 'FabricAdapters'.
	mediaControllers string
	// Memory shall contain a link to a resource collection of type
	// 'MemoryCollection' that represents memory in this chassis that belong to
	// fabric-related resource pools.
	//
	// Version added: v1.11.0
	memory string
	// MemoryDomains shall contain a link to a resource collection of type
	// 'MemoryDomainCollection' that represents memory domains in this chassis that
	// belong to fabric-related resource pools.
	//
	// Version added: v1.11.0
	memoryDomains string
	// MinPowerWatts shall contain the lower bound of the total power consumed by
	// the chassis.
	//
	// Version added: v1.12.0
	MinPowerWatts *float64 `json:",omitempty"`
	// Model shall contain the name by which the manufacturer generally refers to
	// the chassis.
	Model string
	// NetworkAdapters shall contain a link to a resource collection of type
	// 'NetworkAdapterCollection'.
	//
	// Version added: v1.4.0
	networkAdapters string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain a link to a resource collection of type
	// 'PCIeDeviceCollection'.
	//
	// Version added: v1.10.0
	pCIeDevices string
	// PCIeSlots shall contain a link to the resource of type 'PCIeSlots' that
	// represents the PCIe slot information for this chassis.
	//
	// Version added: v1.8.0
	//
	// Deprecated: v1.24.0
	// This property has been deprecated in favor of the 'PCIeDevices' property.
	// The 'PCIeSlots' schema has been deprecated in favor of the 'PCIeDevice'
	// schema. Empty PCIe slots are represented by 'PCIeDevice' resources using the
	// 'Absent' value of the 'State' property within 'Status'.
	pCIeSlots string
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the chassis.
	PartNumber string
	// PhysicalSecurity shall contain the physical security state of the chassis.
	// Services may construct this state from multiple physical sensors.
	//
	// Version added: v1.1.0
	PhysicalSecurity PhysicalSecurity
	// Power shall contain a link to a resource of type 'Power' that represents the
	// power characteristics of this chassis.
	//
	// Deprecated: v1.15.0
	// This link has been deprecated in favor of the 'PowerSubsystem' link
	// property.
	power string
	// PowerState shall contain the power state of the chassis.
	//
	// Version added: v1.0.1
	PowerState PowerState
	// PowerSubsystem shall contain a link to a resource of type 'PowerSubsystem'
	// that represents the power subsystem information for this chassis.
	//
	// Version added: v1.15.0
	powerSubsystem string
	// PoweredByParent shall indicate whether the chassis receives power from the
	// chassis that contains it. The value 'true' shall indicate that the
	// containing chassis provides power. The value 'false' shall indicate the
	// chassis receives power from its own power subsystem, another chassis
	// instance's power supplies, or outlets.
	//
	// Version added: v1.20.0
	PoweredByParent bool
	// RackMountCapacityUnits shall contain the amount of space, in 'RackUnits' and
	// including fractional units, contained within this chassis that are usable to
	// hold rack-mountable equipment. This property shall not be present if
	// 'ChassisType' does not contain 'Rack'.
	//
	// Version added: v1.28.0
	RackMountCapacityUnits *float64 `json:",omitempty"`
	// RackMountDepthMm shall represent the depth (length) of the chassis, in
	// millimeter units, that is available to contain rack-mounted equipment.
	//
	// Version added: v1.28.0
	RackMountDepthMm *float64 `json:",omitempty"`
	// RackMountWidth shall contain the type of width that describes the
	// rack-mountable equipment space.
	//
	// Version added: v1.28.0
	RackMountWidth RackMountWidth
	// RackUnits shall contain the type of units used to describe rack-mountable
	// equipment.
	//
	// Version added: v1.28.0
	RackUnits RackUnits
	// ReadyToRemove shall indicate whether the chassis is ready for removal.
	// Setting the value to 'true' shall cause the service to perform appropriate
	// actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	//
	// Version added: v1.28.0
	ReadyToRemove bool
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.21.0
	Replaceable bool
	// SKU shall contain the stock-keeping unit number for this chassis.
	SKU string
	// Sensors shall contain a link to a resource collection of type
	// 'SensorCollection' that contains the sensors located in the chassis and
	// sub-components.
	//
	// Version added: v1.9.0
	sensors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the chassis.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the chassis.
	//
	// Version added: v1.16.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Thermal shall contain a link to a resource of type 'Thermal' that represents
	// the thermal characteristics of this chassis.
	//
	// Deprecated: v1.15.0
	// This link has been deprecated in favor of the 'ThermalSubsystem' link
	// property.
	thermal string
	// ThermalDirection shall indicate the general direction of the thermal
	// management path through the chassis.
	//
	// Version added: v1.20.0
	ThermalDirection ThermalDirection
	// ThermalManagedByParent shall indicate whether the chassis relies on the
	// containing chassis to provide thermal management. The value 'true' shall
	// indicate that the chassis relies on the containing chassis to provide
	// thermal management. The value 'false' shall indicate the chassis provides
	// thermal management, and may provide details in a 'ThermalSubsystem'
	// resource, or by populating the 'Fans' property in Links.
	//
	// Version added: v1.20.0
	ThermalManagedByParent bool
	// ThermalSubsystem shall contain a link to a resource of type
	// 'ThermalSubsystem' that represents the thermal subsystem information for
	// this chassis.
	//
	// Version added: v1.15.0
	thermalSubsystem string
	// TrustedComponents shall contain a link to a resource collection of type
	// 'TrustedComponentCollection'.
	//
	// Version added: v1.21.0
	trustedComponents string
	// UUID shall contain the universally unique identifier number for this
	// chassis.
	//
	// Version added: v1.7.0
	UUID string
	// Version shall contain the hardware version of this chassis as determined by
	// the vendor or supplier.
	//
	// Version added: v1.21.0
	Version string
	// WeightKg shall represent the published mass, commonly referred to as weight,
	// of the chassis, in kilogram units.
	//
	// Version added: v1.4.0
	WeightKg *float64 `json:",omitempty"`
	// WidthMm shall represent the width of the chassis, in millimeter units, as
	// specified by the manufacturer.
	//
	// Version added: v1.4.0
	WidthMm *float64 `json:",omitempty"`
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// automationNodes are the URIs for AutomationNodes.
	automationNodes []string
	// cables are the URIs for Cables.
	cables []string
	// computerSystems are the URIs for ComputerSystems.
	computerSystems []string
	// connectedCoolingLoops are the URIs for ConnectedCoolingLoops.
	connectedCoolingLoops []string
	// containedBy is the URI for ContainedBy.
	containedBy string
	// contains are the URIs for Contains.
	contains []string
	// cooledBy are the URIs for CooledBy.
	cooledBy []string
	// coolingUnits are the URIs for CoolingUnits.
	coolingUnits []string
	// linkedDrives shall contain an array of links to resources
	// of type Drive that are in this chassis.
	linkedDrives []string
	// Drives shall contain a link to a resource collection of type DriveCollection.
	drives string
	// DrivesCount is the number of drives attached to this chassis.
	DrivesCount int `json:"Drives@odata.count"`
	// facility is the URI for Facility.
	facility string
	// fans are the URIs for Fans.
	fans []string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// managersInChassis are the URIs for ManagersInChassis.
	managersInChassis []string
	linkedPCIeDevices []string
	// powerDistribution is the URI for PowerDistribution.
	powerDistribution string
	// powerOutlets are the URIs for PowerOutlets.
	powerOutlets []string
	// powerSupplies are the URIs for PowerSupplies.
	powerSupplies []string
	// poweredBy are the URIs for PoweredBy.
	poweredBy []string
	// processors are the URIs for Processors.
	processors []string
	// resourceBlocks are the URIs for ResourceBlocks.
	resourceBlocks []string
	// storage are the URIs for Storage.
	storage []string
	// switches are the URIs for Switches.
	switches []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte

	// resetActionInfoTarget is the URL to check what values are supported
	resetActionInfoTarget string
	// SupportedResetTypes, if provided, is the reset types this chassis supports.
	SupportedResetTypes []ResetType
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
//
//nolint:funlen
func (c *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	type cActions struct {
		Reset struct {
			ActionTarget
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Chassis.Reset"`
	}
	type cLinks struct {
		AutomationNodes       Links `json:"AutomationNodes"`
		Cables                Links `json:"Cables"`
		ComputerSystems       Links `json:"ComputerSystems"`
		ConnectedCoolingLoops Links `json:"ConnectedCoolingLoops"`
		ContainedBy           Link  `json:"ContainedBy"`
		Contains              Links `json:"Contains"`
		CooledBy              Links `json:"CooledBy"`
		CoolingUnits          Links `json:"CoolingUnits"`
		Drives                Links `json:"Drives"`
		DrivesCount           int   `json:"Drives@odata.count"`
		Facility              Link  `json:"Facility"`
		Fans                  Links `json:"Fans"`
		ManagedBy             Links `json:"ManagedBy"`
		ManagersInChassis     Links `json:"ManagersInChassis"`
		PCIeDevices           Links `json:"PCIeDevices"`
		PowerDistribution     Link  `json:"PowerDistribution"`
		PowerOutlets          Links `json:"PowerOutlets"`
		PowerSupplies         Links `json:"PowerSupplies"`
		PoweredBy             Links `json:"PoweredBy"`
		Processors            Links `json:"Processors"`
		ResourceBlocks        Links `json:"ResourceBlocks"`
		Storage               Links `json:"Storage"`
		Switches              Links `json:"Switches"`
	}
	var tmp struct {
		temp
		Actions            cActions
		Links              cLinks
		Assembly           Link `json:"Assembly"`
		Certificates       Link `json:"Certificates"`
		Controls           Link `json:"Controls"`
		Drives             Link
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		FabricAdapters     Link `json:"FabricAdapters"`
		LeakDetectors      Link `json:"LeakDetectors"`
		LogServices        Link `json:"LogServices"`
		MediaControllers   Link `json:"MediaControllers"`
		Memory             Link `json:"Memory"`
		MemoryDomains      Link `json:"MemoryDomains"`
		NetworkAdapters    Link `json:"NetworkAdapters"`
		PCIeDevices        Link `json:"PCIeDevices"`
		PCIeSlots          Link `json:"PCIeSlots"`
		Power              Link `json:"Power"`
		PowerSubsystem     Link `json:"PowerSubsystem"`
		Sensors            Link `json:"Sensors"`
		Thermal            Link `json:"Thermal"`
		ThermalSubsystem   Link `json:"ThermalSubsystem"`
		TrustedComponents  Link `json:"TrustedComponents"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Chassis(tmp.temp)

	// Extract the links to other entities for later
	c.resetTarget = tmp.Actions.Reset.Target
	c.resetActionInfoTarget = tmp.Actions.Reset.ActionInfoTarget
	c.SupportedResetTypes = tmp.Actions.Reset.AllowedResetTypes

	c.automationNodes = tmp.Links.AutomationNodes.ToStrings()
	c.cables = tmp.Links.Cables.ToStrings()
	c.computerSystems = tmp.Links.ComputerSystems.ToStrings()
	c.connectedCoolingLoops = tmp.Links.ConnectedCoolingLoops.ToStrings()
	c.containedBy = tmp.Links.ContainedBy.String()
	c.contains = tmp.Links.Contains.ToStrings()
	c.cooledBy = tmp.Links.CooledBy.ToStrings()
	c.coolingUnits = tmp.Links.CoolingUnits.ToStrings()
	c.drives = tmp.Drives.String()
	c.linkedDrives = tmp.Links.Drives.ToStrings()
	c.facility = tmp.Links.Facility.String()
	c.fans = tmp.Links.Fans.ToStrings()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.managersInChassis = tmp.Links.ManagersInChassis.ToStrings()
	c.powerDistribution = tmp.Links.PowerDistribution.String()
	c.powerOutlets = tmp.Links.PowerOutlets.ToStrings()
	c.powerSupplies = tmp.Links.PowerSupplies.ToStrings()
	c.poweredBy = tmp.Links.PoweredBy.ToStrings()
	c.processors = tmp.Links.Processors.ToStrings()
	c.resourceBlocks = tmp.Links.ResourceBlocks.ToStrings()
	c.storage = tmp.Links.Storage.ToStrings()
	c.switches = tmp.Links.Switches.ToStrings()
	c.assembly = tmp.Assembly.String()
	c.certificates = tmp.Certificates.String()
	c.controls = tmp.Controls.String()
	c.environmentMetrics = tmp.EnvironmentMetrics.String()
	c.fabricAdapters = tmp.FabricAdapters.String()
	c.leakDetectors = tmp.LeakDetectors.String()
	c.logServices = tmp.LogServices.String()
	c.mediaControllers = tmp.MediaControllers.String()
	c.memory = tmp.Memory.String()
	c.memoryDomains = tmp.MemoryDomains.String()
	c.networkAdapters = tmp.NetworkAdapters.String()
	c.pCIeDevices = tmp.PCIeDevices.String()
	c.linkedPCIeDevices = tmp.Links.PCIeDevices.ToStrings()
	c.pCIeSlots = tmp.PCIeSlots.String()
	c.power = tmp.Power.String()
	c.powerSubsystem = tmp.PowerSubsystem.String()
	c.sensors = tmp.Sensors.String()
	c.thermal = tmp.Thermal.String()
	c.thermalSubsystem = tmp.ThermalSubsystem.String()
	c.trustedComponents = tmp.TrustedComponents.String()
	if c.DrivesCount == 0 && tmp.Links.DrivesCount > 0 {
		c.DrivesCount = tmp.Links.DrivesCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Chassis) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"EnvironmentalClass",
		"HeatingCoolingEquipmentNames",
		"HeatingCoolingManagerURIs",
		"IndicatorLED",
		"LocationIndicatorActive",
		"RackMountCapacityUnits",
		"RackMountDepthMm",
		"RackMountWidth",
		"RackUnits",
		"ReadyToRemove",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetChassis will get a Chassis instance from the service.
func GetChassis(c Client, uri string) (*Chassis, error) {
	return GetObject[Chassis](c, uri)
}

// ListReferencedChassiss gets the collection of Chassis from
// a provided reference.
func ListReferencedChassiss(c Client, link string) ([]*Chassis, error) {
	return GetCollectionObjects[Chassis](c, link)
}

// GetSupportedResetTypes returns any reset types that the Chassis declares as supported
// via either ActionInfo or AllowableValues.
func (c *Chassis) GetSupportedResetTypes() ([]ResetType, error) {
	if len(c.SupportedResetTypes) > 0 {
		return c.SupportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if c.resetActionInfoTarget != "" {
		resetActionInfo, err := c.ResetActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			c.SupportedResetTypes = append(c.SupportedResetTypes, ResetType(val))
		}
	}

	return c.SupportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the Chassis reset action if supported
func (c *Chassis) ResetActionInfo() (*ActionInfo, error) {
	if c.resetActionInfoTarget == "" {
		return nil, errors.New("Chassis Reset resetActionInfoTarget not supported")
	}

	return GetObject[ActionInfo](c.GetClient(), c.resetActionInfoTarget)
}

// This action shall reset the chassis. Additionally, it may reset systems or
// other contained resources depending on the 'ResetType' used to invoke this
// action.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and complete an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Chassis) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AutomationNodes gets the AutomationNodes linked resources.
func (c *Chassis) AutomationNodes() ([]*AutomationNode, error) {
	return GetObjects[AutomationNode](c.client, c.automationNodes)
}

// Cables gets the Cables linked resources.
func (c *Chassis) Cables() ([]*Cable, error) {
	return GetObjects[Cable](c.client, c.cables)
}

// ComputerSystems gets the ComputerSystems linked resources.
func (c *Chassis) ComputerSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](c.client, c.computerSystems)
}

// ConnectedCoolingLoops gets the ConnectedCoolingLoops linked resources.
func (c *Chassis) ConnectedCoolingLoops() ([]*CoolingLoop, error) {
	return GetObjects[CoolingLoop](c.client, c.connectedCoolingLoops)
}

// ContainedBy gets the ContainedBy linked resource.
func (c *Chassis) ContainedBy() (*Chassis, error) {
	if c.containedBy == "" {
		return nil, nil
	}
	return GetObject[Chassis](c.client, c.containedBy)
}

// Contains gets the Contains linked resources.
func (c *Chassis) Contains() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.contains)
}

// CooledBy gets the CooledBy linked resources.
func (c *Chassis) CooledBy() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.cooledBy)
}

// CoolingUnits gets the CoolingUnits linked resources.
func (c *Chassis) CoolingUnits() ([]*CoolingUnit, error) {
	return GetObjects[CoolingUnit](c.client, c.coolingUnits)
}

// Drives gets the Drives linked resources.
func (c *Chassis) Drives() ([]*Drive, error) {
	// In version v1.2.0 of the spec, Drives were added to the Chassis.Links
	// property. But in v1.14.0 of the spec, Chassis.Drives was added as a
	// direct property.
	if c.drives != "" {
		return GetCollectionObjects[Drive](c.GetClient(), c.drives)
	}

	if len(c.linkedDrives) > 0 {
		return GetObjects[Drive](c.GetClient(), c.linkedDrives)
	}

	return []*Drive{}, nil
}

// Facility gets the Facility linked resource.
func (c *Chassis) Facility() (*Facility, error) {
	if c.facility == "" {
		return nil, nil
	}
	return GetObject[Facility](c.client, c.facility)
}

// Fans gets the Fans linked resources.
func (c *Chassis) Fans() ([]*Fan, error) {
	return GetObjects[Fan](c.client, c.fans)
}

// ManagedBy gets the ManagedBy linked resources.
func (c *Chassis) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](c.client, c.managedBy)
}

// ManagersInChassis gets the ManagersInChassis linked resources.
func (c *Chassis) ManagersInChassis() ([]*Manager, error) {
	return GetObjects[Manager](c.client, c.managersInChassis)
}

// PowerDistribution gets the PowerDistribution linked resource.
func (c *Chassis) PowerDistribution() (*PowerDistribution, error) {
	if c.powerDistribution == "" {
		return nil, nil
	}
	return GetObject[PowerDistribution](c.client, c.powerDistribution)
}

// PowerOutlets gets the PowerOutlets linked resources.
func (c *Chassis) PowerOutlets() ([]*Outlet, error) {
	return GetObjects[Outlet](c.client, c.powerOutlets)
}

// PowerSupplies gets the PowerSupplies linked resources.
func (c *Chassis) PowerSupplies() ([]*PowerSupply, error) {
	return GetObjects[PowerSupply](c.client, c.powerSupplies)
}

// PoweredBy gets the PoweredBy linked resources.
func (c *Chassis) PoweredBy() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.poweredBy)
}

// Processors gets the Processors linked resources.
func (c *Chassis) Processors() ([]*Processor, error) {
	return GetObjects[Processor](c.client, c.processors)
}

// ResourceBlocks gets the ResourceBlocks linked resources.
func (c *Chassis) ResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](c.client, c.resourceBlocks)
}

// Storage gets the Storage linked resources.
func (c *Chassis) Storage() ([]*Storage, error) {
	return GetObjects[Storage](c.client, c.storage)
}

// Switches gets the Switches linked resources.
func (c *Chassis) Switches() ([]*Switch, error) {
	return GetObjects[Switch](c.client, c.switches)
}

// Assembly gets the Assembly linked resource.
func (c *Chassis) Assembly() (*Assembly, error) {
	if c.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](c.client, c.assembly)
}

// Certificates gets the Certificates collection.
func (c *Chassis) Certificates() ([]*Certificate, error) {
	if c.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](c.client, c.certificates)
}

// Controls gets the Controls collection.
func (c *Chassis) Controls() ([]*Control, error) {
	if c.controls == "" {
		return nil, nil
	}
	return GetCollectionObjects[Control](c.client, c.controls)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (c *Chassis) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if c.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](c.client, c.environmentMetrics)
}

// FabricAdapters gets the FabricAdapters collection.
func (c *Chassis) FabricAdapters() ([]*FabricAdapter, error) {
	if c.fabricAdapters == "" {
		return nil, nil
	}
	return GetCollectionObjects[FabricAdapter](c.client, c.fabricAdapters)
}

// LeakDetectors gets the LeakDetectors collection.
func (c *Chassis) LeakDetectors() ([]*LeakDetector, error) {
	if c.leakDetectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[LeakDetector](c.client, c.leakDetectors)
}

// LogServices gets the LogServices collection.
func (c *Chassis) LogServices() ([]*LogService, error) {
	if c.logServices == "" {
		return nil, nil
	}
	return GetCollectionObjects[LogService](c.client, c.logServices)
}

// MediaControllers gets the MediaControllers collection.
func (c *Chassis) MediaControllers() ([]*MediaController, error) {
	if c.mediaControllers == "" {
		return nil, nil
	}
	return GetCollectionObjects[MediaController](c.client, c.mediaControllers)
}

// Memory gets the Memory collection.
func (c *Chassis) Memory() ([]*Memory, error) {
	if c.memory == "" {
		return nil, nil
	}
	return GetCollectionObjects[Memory](c.client, c.memory)
}

// MemoryDomains gets the MemoryDomains collection.
func (c *Chassis) MemoryDomains() ([]*MemoryDomain, error) {
	if c.memoryDomains == "" {
		return nil, nil
	}
	return GetCollectionObjects[MemoryDomain](c.client, c.memoryDomains)
}

// NetworkAdapters gets the NetworkAdapters collection.
func (c *Chassis) NetworkAdapters() ([]*NetworkAdapter, error) {
	if c.networkAdapters == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkAdapter](c.client, c.networkAdapters)
}

// PCIeDevices gets the PCIeDevices collection.
func (c *Chassis) PCIeDevices() ([]*PCIeDevice, error) {
	if c.pCIeDevices != "" {
		return GetCollectionObjects[PCIeDevice](c.GetClient(), c.pCIeDevices)
	}

	if len(c.linkedPCIeDevices) > 0 {
		return GetObjects[PCIeDevice](c.GetClient(), c.linkedPCIeDevices)
	}

	return []*PCIeDevice{}, nil
}

// PCIeSlots gets the PCIeSlots linked resource.
func (c *Chassis) PCIeSlots() (*PCIeSlots, error) {
	if c.pCIeSlots == "" {
		return nil, nil
	}
	return GetObject[PCIeSlots](c.client, c.pCIeSlots)
}

// Power gets the Power linked resource.
func (c *Chassis) Power() (*Power, error) {
	if c.power == "" {
		return nil, nil
	}
	return GetObject[Power](c.client, c.power)
}

// PowerSubsystem gets the PowerSubsystem linked resource.
func (c *Chassis) PowerSubsystem() (*PowerSubsystem, error) {
	if c.powerSubsystem == "" {
		return nil, nil
	}
	return GetObject[PowerSubsystem](c.client, c.powerSubsystem)
}

// Sensors gets the Sensors collection.
func (c *Chassis) Sensors() ([]*Sensor, error) {
	if c.sensors == "" {
		return nil, nil
	}
	return GetCollectionObjects[Sensor](c.client, c.sensors)
}

// Thermal gets the Thermal linked resource.
func (c *Chassis) Thermal() (*Thermal, error) {
	if c.thermal == "" {
		return nil, nil
	}
	return GetObject[Thermal](c.client, c.thermal)
}

// ThermalSubsystem gets the ThermalSubsystem linked resource.
func (c *Chassis) ThermalSubsystem() (*ThermalSubsystem, error) {
	if c.thermalSubsystem == "" {
		return nil, nil
	}
	return GetObject[ThermalSubsystem](c.client, c.thermalSubsystem)
}

// TrustedComponents gets the TrustedComponents collection.
func (c *Chassis) TrustedComponents() ([]*TrustedComponent, error) {
	if c.trustedComponents == "" {
		return nil, nil
	}
	return GetCollectionObjects[TrustedComponent](c.client, c.trustedComponents)
}

// Door shall describe a door or access panel on the chassis.
type Door struct {
	// DoorState shall contain the current state of the door.
	//
	// Version added: v1.24.0
	DoorState DoorState
	// Locked shall indicate if the door is locked.
	//
	// Version added: v1.24.0
	Locked bool
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.24.0
	UserLabel string
}

// Doors shall describe the doors or access panels of the chassis.
type Doors struct {
	// Front shall contain information related to the front door, as defined by the
	// manufacturer, of the chassis.
	//
	// Version added: v1.24.0
	Front Door
	// Rear shall contain information related to the rear door, as defined by the
	// manufacturer, of the chassis.
	//
	// Version added: v1.24.0
	Rear Door
}

// PhysicalSecurity shall describe the physical security state of the chassis.
type PhysicalSecurity struct {
	// IntrusionSensor shall contain the physical security state of the chassis. If
	// the 'IntrusionSensorReArm' property contains 'Manual', a client may set this
	// property to 'Normal' to reset the physical security state.
	//
	// Version added: v1.1.0
	IntrusionSensor IntrusionSensor
	// IntrusionSensorNumber shall contain a numerical identifier for this physical
	// security sensor that is unique within this resource.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.22.0
	// This property has been deprecated in order to allow for multiple physical
	// sensors to construct this object.
	IntrusionSensorNumber *int `json:",omitempty"`
	// IntrusionSensorReArm shall contain the policy that describes how the
	// 'IntrusionSensor' property returns to the 'Normal' value.
	//
	// Version added: v1.1.0
	IntrusionSensorReArm IntrusionSensorReArm
}
