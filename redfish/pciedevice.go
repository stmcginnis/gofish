//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type CXLDeviceType string

const (
	// Type1CXLDeviceType shall indicate a CXL Specification-defined Type 1 device.
	Type1CXLDeviceType CXLDeviceType = "Type1"
	// Type2CXLDeviceType shall indicate a CXL Specification-defined Type 2 device.
	Type2CXLDeviceType CXLDeviceType = "Type2"
	// Type3CXLDeviceType shall indicate a CXL Specification-defined Type 3 device.
	Type3CXLDeviceType CXLDeviceType = "Type3"
)

type CXLDynamicCapacityPolicies string

const (
	// FreeCXLDynamicCapacityPolicies shall indicate the CXL Specification-defined free add capacity policy.
	FreeCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Free"
	// ContiguousCXLDynamicCapacityPolicies shall indicate the CXL Specification-defined contiguous add capacity
	// policy.
	ContiguousCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Contiguous"
	// PrescriptiveCXLDynamicCapacityPolicies shall indicate the CXL Specification-defined prescriptive add or release
	// policy.
	PrescriptiveCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Prescriptive"
	// TagBasedCXLDynamicCapacityPolicies shall indicate the CXL Specification-defined tag-based release policy.
	TagBasedCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "TagBased"
)

type DeviceType string

const (
	// SingleFunctionDeviceType A single-function PCIe device.
	SingleFunctionDeviceType DeviceType = "SingleFunction"
	// MultiFunctionDeviceType A multi-function PCIe device.
	MultiFunctionDeviceType DeviceType = "MultiFunction"
	// SimulatedDeviceType A PCIe device which is not currently physically
	// present, but is being simulated by the PCIe infrastructure.
	SimulatedDeviceType DeviceType = "Simulated"
)

type LaneSplittingType string

const (
	// NoneLaneSplittingType The slot has no lane splitting.
	NoneLaneSplittingType LaneSplittingType = "None"
	// BridgedLaneSplittingType The slot has a bridge to share the lanes with associated devices.
	BridgedLaneSplittingType LaneSplittingType = "Bridged"
	// BifurcatedLaneSplittingType The slot is bifurcated to split the lanes with associated devices.
	BifurcatedLaneSplittingType LaneSplittingType = "Bifurcated"
)

type SlotType string

const (
	// FullLengthSlotType Full-Length PCIe slot.
	FullLengthSlotType SlotType = "FullLength"
	// HalfLengthSlotType Half-Length PCIe slot.
	HalfLengthSlotType SlotType = "HalfLength"
	// LowProfileSlotType Low-Profile or Slim PCIe slot.
	LowProfileSlotType SlotType = "LowProfile"
	// MiniSlotType Mini PCIe slot.
	MiniSlotType SlotType = "Mini"
	// M2SlotType PCIe M.2 slot.
	M2SlotType SlotType = "M2"
	// OEMSlotType An OEM-specific slot.
	OEMSlotType SlotType = "OEM"
	// OCP3SmallSlotType Open Compute Project 3.0 small form factor slot.
	OCP3SmallSlotType SlotType = "OCP3Small"
	// OCP3LargeSlotType Open Compute Project 3.0 large form factor slot.
	OCP3LargeSlotType SlotType = "OCP3Large"
	// U2SlotType U.2 / SFF-8639 slot or bay.
	U2SlotType SlotType = "U2"
)

// CXLDevice shall contain CXL-specific properties of a PCIe device.
type CXLDevice struct {
	// DeviceType shall contain the CXL device type.
	DeviceType CXLDeviceType
	// DynamicCapacity shall contain the CXL dynamic capacity device (DCD) information for this CXL device.
	DynamicCapacity CXLDynamicCapacity
	// EgressPortCongestionSupport shall indicate whether the CXL device supports the CXL Specification-defined 'Egress
	// Port Congestion' mechanism.
	EgressPortCongestionSupport bool
	// MaxNumberLogicalDevices shall contain the maximum number of logical devices supported by this CXL device.
	MaxNumberLogicalDevices int
	// ThroughputReductionSupport shall indicate whether the CXL device supports the CXL Specification-defined
	// 'Throughput Reduction' mechanism.
	ThroughputReductionSupport bool
	// Timestamp shall contain the timestamp set on the CXL device.
	Timestamp string
}

// CXLDynamicCapacity shall contain the CXL dynamic capacity device (DCD) information for a CXL device.
type CXLDynamicCapacity struct {
	// AddCapacityPoliciesSupported shall contain the CXL Specification-defined dynamic capacity policies that are
	// supported by this CXL device when dynamic capacity is added.
	AddCapacityPoliciesSupported []CXLDynamicCapacityPolicies
	// MaxDynamicCapacityRegions shall contain the maximum number of dynamic capacity memory regions available per host
	// from this CXL device.
	MaxDynamicCapacityRegions int
	// MaxHosts shall contain the maximum number of hosts supported by this CXL device.
	MaxHosts int
	// MemoryBlockSizesSupported shall contain the set of memory block sizes supported by memory regions in this CXL
	// device.
	MemoryBlockSizesSupported []CXLRegionBlockSizes
	// ReleaseCapacityPoliciesSupported shall contain the CXL Specification-defined dynamic capacity policies that are
	// supported by this CXL device when dynamic capacity is released.
	ReleaseCapacityPoliciesSupported []CXLDynamicCapacityPolicies
	// SanitizationOnReleaseSupport shall indicate whether the sanitization on capacity release is configurable for the
	// memory regions in this CXL device.
	SanitizationOnReleaseSupport []CXLRegionSanitization
	// TotalDynamicCapacityMiB shall contain the total memory media capacity of the CXL device available for dynamic
	// assignment in mebibytes (MiB).
	TotalDynamicCapacityMiB int
}

// CXLRegionBlockSizes shall contain the set of memory block sizes supported by memory region in the dynamic
// capacity device.
type CXLRegionBlockSizes struct {
	// BlockSizeMiB shall contain the set of memory block sizes supported by this memory region, with units in MiB.
	BlockSizeMiB []string
	// RegionNumber shall contain the memory region number.
	RegionNumber int
}

// CXLRegionSanitization shall indicate whether the sanitization on capacity release is configurable for the memory
// region.
type CXLRegionSanitization struct {
	// RegionNumber shall contain the memory region number.
	RegionNumber int
	// SanitizationOnReleaseSupported shall indicate whether the sanitization on capacity release is configurable for
	// this memory region.
	SanitizationOnReleaseSupported bool
}

// PCIeTypes is the type of PCIe device.
type PCIeTypes string

const (
	// Gen1PCIeTypes A PCIe v1.0 slot.
	Gen1PCIeTypes PCIeTypes = "Gen1"
	// Gen2PCIeTypes A PCIe v2.0 slot.
	Gen2PCIeTypes PCIeTypes = "Gen2"
	// Gen3PCIeTypes A PCIe v3.0 slot.
	Gen3PCIeTypes PCIeTypes = "Gen3"
	// Gen4PCIeTypes A PCIe v4.0 slot.
	Gen4PCIeTypes PCIeTypes = "Gen4"
	// Gen5PCIeTypes A PCIe v5.0 slot.
	Gen5PCIeTypes PCIeTypes = "Gen5"
)

// PCIeErrors shall contain properties that describe the PCIe errors associated with this device.
type PCIeErrors struct {
	// CorrectableErrorCount shall contain the total number of PCIe correctable errors for this device.
	CorrectableErrorCount int
	// FatalErrorCount shall contain the total number of PCIe fatal errors for this device.
	FatalErrorCount int
	// L0ToRecoveryCount shall contain the total number of times the PCIe link transitioned from L0 to the recovery
	// state for this device.
	L0ToRecoveryCount int
	// NAKReceivedCount shall contain the total number of NAKs issued on the PCIe link by the receiver. A NAK is issued
	// by the receiver when it detects that a TLP from this device was missed. This could be because this device did
	// not transmit it, or because the receiver could not properly decode the packet.
	NAKReceivedCount int
	// NAKSentCount shall contain the total number of NAKs issued on the PCIe link by this device. A NAK is issued by
	// the device when it detects that a TLP from the receiver was missed. This could be because the receiver did not
	// transmit it, or because this device could not properly decode the packet.
	NAKSentCount int
	// NonFatalErrorCount shall contain the total number of PCIe non-fatal errors for this device.
	NonFatalErrorCount int
	// ReplayCount shall contain the total number of replays issued on the PCIe link by this device. A replay is a
	// retransmission of a TLP and occurs because the ACK timer is expired, which means that the receiver did not send
	// the ACK or this device did not properly decode the ACK.
	ReplayCount int
	// ReplayRolloverCount shall contain the total number of replay rollovers issued on the PCIe link by this device. A
	// replay rollover occurs when consecutive replays failed to resolve the errors on the link, which means that this
	// device forced the link into the recovery state.
	ReplayRolloverCount int
	// UnsupportedRequestCount shall contain the total number of PCIe unsupported requests received by this device.
	UnsupportedRequestCount int
}

// PCIeDevice shall represent a PCIe device in a Redfish implementation. It may also represent a location, such as
// a slot, socket, or bay, where a unit may be installed, but the State property within the Status property
// contains 'Absent'.
type PCIeDevice struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the PCIe device for inventory purposes.
	AssetTag string
	// CXLDevice shall contain CXL-specific properties of this PCIe device.
	CXLDevice CXLDevice
	// CXLLogicalDevices shall contain a link to a resource collection of type CXLLogicalDeviceCollection.
	cxlLogicalDevices string
	// Description provides a description of this resource.
	Description string
	// DeviceType shall be the device type of the PCIe device such as
	// SingleFunction or MultiFunction.
	DeviceType DeviceType
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this PCIe device.
	EnvironmentMetrics EnvironmentMetrics
	// FirmwareVersion shall be the firmware version of the PCIe device.
	FirmwareVersion string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall be the name of the organization responsible for
	// producing the PCIe device. This organization might be the entity from
	// whom the PCIe device is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall be the name by which the manufacturer generally refers to the
	// PCIe device.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// PCIeInterface is used to connect this PCIe Device to its host or
	// upstream switch.
	PCIeInterface PCIeInterface
	// PartNumber shall be a part number assigned by the organization that is
	// responsible for producing or manufacturing the PCIe device.
	PartNumber string
	// ReadyToRemove shall indicate whether the PCIe device is ready for removal. Setting the value to 'true' shall
	// cause the service to perform appropriate actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	ReadyToRemove bool
	// SKU shall be the stock-keeping unit number for this PCIe device.
	SKU string
	// SerialNumber is used to identify the PCIe device.
	SerialNumber string
	// Slot shall contain information about the PCIe slot for this PCIe device.
	Slot PCIeSlot
	// SparePartNumber shall contain the spare part number of the PCIe device.
	SparePartNumber string
	// StagedVersion shall contain the staged firmware version for this PCIe device; this firmware is not yet active.
	StagedVersion string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain the universally unique identifier number for this PCIe device.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// Chassis shall reference a resource of type Chassis that represents the
	// physical container associated with this resource.
	chassis []string
	// ChassisCount is the number of number of associated chassis.
	ChassisCount int
	// PCIeFunctions shall be a reference to the resources that this device
	// exposes and shall reference a resource of type PCIeFunction.
	pcieFunctions      string
	pcieFunctionsArray []string
	// PCIeFunctionsCount is the number of PCIeFunctions.
	PCIeFunctionsCount int
	processors         []string
	// ProcessorsCount is the number of processors that are directly connected or directly bridged
	// to this PCIeDevice.
	ProcessorsCount int
	sw              string
}

// UnmarshalJSON unmarshals a PCIeDevice object from the raw JSON.
func (pciedevice *PCIeDevice) UnmarshalJSON(b []byte) error {
	type temp PCIeDevice
	type links struct {
		Chassis            common.Links
		ChassisCount       int `json:"Chassis@odata.count"`
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
		// Processors shall contain an array of links to resources of type Processor that represent processors that are
		// directly connected or directly bridged to this PCIe device.
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
		// Switch shall contain a link to a resource of type Switch that is associated with this PCIe device.
		Switch common.Link
	}
	var t struct {
		temp
		Assembly          common.Link
		CXLLogicalDevices common.Link
		PCIeFunctions     common.Link
		Links             links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciedevice = PCIeDevice(t.temp)

	// Extract the links to other entities for later
	pciedevice.assembly = t.Assembly.String()
	pciedevice.cxlLogicalDevices = t.CXLLogicalDevices.String()
	pciedevice.pcieFunctions = t.PCIeFunctions.String()

	pciedevice.chassis = t.Links.Chassis.ToStrings()
	pciedevice.ChassisCount = t.Links.ChassisCount
	pciedevice.processors = t.Links.Processors.ToStrings()
	pciedevice.ProcessorsCount = t.Links.ProcessorsCount
	pciedevice.sw = t.Links.Switch.String()

	if t.Links.PCIeFunctionsCount != 0 {
		pciedevice.PCIeFunctionsCount = t.Links.PCIeFunctionsCount
		pciedevice.pcieFunctionsArray = t.Links.PCIeFunctions.ToStrings()
	} else {
		pciedevice.pcieFunctions = t.PCIeFunctions.String()
	}

	// This is a read/write object, so we need to save the raw object data for later
	pciedevice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (pciedevice *PCIeDevice) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PCIeDevice)
	err := original.UnmarshalJSON(pciedevice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"LocationIndicatorActive",
		"ReadyToRemove",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(pciedevice).Elem()

	return pciedevice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPCIeDevice will get a PCIeDevice instance from the service.
func GetPCIeDevice(c common.Client, uri string) (*PCIeDevice, error) {
	return common.GetObject[PCIeDevice](c, uri)
}

// ListReferencedPCIeDevices gets the collection of PCIeDevice from
// a provided reference.
func ListReferencedPCIeDevices(c common.Client, link string) ([]*PCIeDevice, error) {
	return common.GetCollectionObjects[PCIeDevice](c, link)
}

// PCIeInterface properties shall be the definition for a PCIe Interface for a
// Redfish implementation.
type PCIeInterface struct {
	// LanesInUse shall be the number of PCIe lanes in use by this device, which
	// shall be equal or less than the value of MaxLanes.
	LanesInUse int
	// MaxLanes shall be the maximum number of PCIe lanes supported by this device.
	MaxLanes int
	// MaxPCIeType shall be the maximum PCIe specification that this device supports.
	MaxPCIeType PCIeTypes
	// PCIeType shall be the negotiated PCIe interface version in use by this device.
	PCIeType PCIeTypes
}

// Assembly gets the assembly for this device.
func (pciedevice *PCIeDevice) Assembly() (*Assembly, error) {
	if pciedevice.assembly == "" {
		return nil, nil
	}
	return GetAssembly(pciedevice.GetClient(), pciedevice.assembly)
}

// CXLLogicalDevices gets the associated CXLLogicalDevices for this device.
func (pciedevice *PCIeDevice) CXLLogicalDevices() ([]*CXLLogicalDevice, error) {
	if pciedevice.cxlLogicalDevices == "" {
		return []*CXLLogicalDevice{}, nil
	}
	return ListReferencedCXLLogicalDevices(pciedevice.GetClient(), pciedevice.cxlLogicalDevices)
}

// Chassis gets the chassis in which the PCIe device is contained.
func (pciedevice *PCIeDevice) Chassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](pciedevice.GetClient(), pciedevice.chassis)
}

// PCIeFunctions get the PCIe functions that this device exposes.
func (pciedevice *PCIeDevice) PCIeFunctions() ([]*PCIeFunction, error) {
	if len(pciedevice.pcieFunctionsArray) > 0 {
		return common.GetObjects[PCIeFunction](pciedevice.GetClient(), pciedevice.pcieFunctionsArray)
	}
	return ListReferencedPCIeFunctions(pciedevice.GetClient(), pciedevice.pcieFunctions)
}

// Switch gets the switch for this device.
func (pciedevice *PCIeDevice) Switch() (*Switch, error) {
	if pciedevice.sw == "" {
		return nil, nil
	}
	return GetSwitch(pciedevice.GetClient(), pciedevice.sw)
}

// Processors gets the processors that are directly connected or directly bridged to this PCIe device.
func (pciedevice *PCIeDevice) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](pciedevice.GetClient(), pciedevice.processors)
}
