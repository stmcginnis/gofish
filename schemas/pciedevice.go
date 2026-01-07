//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #PCIeDevice.v1_21_0.PCIeDevice

package schemas

import (
	"encoding/json"
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
	// FreeCXLDynamicCapacityPolicies shall indicate the CXL Specification-defined
	// free add capacity policy.
	FreeCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Free"
	// ContiguousCXLDynamicCapacityPolicies shall indicate the CXL
	// Specification-defined contiguous add capacity policy.
	ContiguousCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Contiguous"
	// PrescriptiveCXLDynamicCapacityPolicies shall indicate the CXL
	// Specification-defined prescriptive add or release policy.
	PrescriptiveCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "Prescriptive"
	// TagBasedCXLDynamicCapacityPolicies shall indicate the CXL
	// Specification-defined tag-based release policy.
	TagBasedCXLDynamicCapacityPolicies CXLDynamicCapacityPolicies = "TagBased"
)

type CXLProtocolVersion string

const (
	// CXL11CXLProtocolVersion CXL protocol specification revision 1.1.
	CXL11CXLProtocolVersion CXLProtocolVersion = "CXL1_1"
	// CXL20CXLProtocolVersion CXL protocol specification revision 2.0.
	CXL20CXLProtocolVersion CXLProtocolVersion = "CXL2_0"
	// CXL30CXLProtocolVersion CXL protocol specification revision 3.0.
	CXL30CXLProtocolVersion CXLProtocolVersion = "CXL3_0"
	// CXL31CXLProtocolVersion CXL protocol specification revision 3.1.
	CXL31CXLProtocolVersion CXLProtocolVersion = "CXL3_1"
	// CXL32CXLProtocolVersion CXL protocol specification revision 3.2.
	CXL32CXLProtocolVersion CXLProtocolVersion = "CXL3_2"
)

type DeviceType string

const (
	// SingleFunctionDeviceType is a single-function PCIe device.
	SingleFunctionDeviceType DeviceType = "SingleFunction"
	// MultiFunctionDeviceType is a multi-function PCIe device.
	MultiFunctionDeviceType DeviceType = "MultiFunction"
	// SimulatedDeviceType is a PCIe device that is not currently physically
	// present, but is being simulated by the PCIe infrastructure.
	SimulatedDeviceType DeviceType = "Simulated"
	// RetimerDeviceType is a PCIe retimer device.
	RetimerDeviceType DeviceType = "Retimer"
)

type LaneSplittingType string

const (
	// NoneLaneSplittingType The slot has no lane splitting.
	NoneLaneSplittingType LaneSplittingType = "None"
	// BridgedLaneSplittingType The slot has a bridge to share the lanes with
	// associated devices.
	BridgedLaneSplittingType LaneSplittingType = "Bridged"
	// BifurcatedLaneSplittingType The slot is bifurcated to split the lanes with
	// associated devices.
	BifurcatedLaneSplittingType LaneSplittingType = "Bifurcated"
)

type PCIeTypes string

const (
	// Gen1PCIeTypes is a PCIe v1.0 slot.
	Gen1PCIeTypes PCIeTypes = "Gen1"
	// Gen2PCIeTypes is a PCIe v2.0 slot.
	Gen2PCIeTypes PCIeTypes = "Gen2"
	// Gen3PCIeTypes is a PCIe v3.0 slot.
	Gen3PCIeTypes PCIeTypes = "Gen3"
	// Gen4PCIeTypes is a PCIe v4.0 slot.
	Gen4PCIeTypes PCIeTypes = "Gen4"
	// Gen5PCIeTypes is a PCIe v5.0 slot.
	Gen5PCIeTypes PCIeTypes = "Gen5"
	// Gen6PCIeTypes is a PCIe v6.0 slot.
	Gen6PCIeTypes PCIeTypes = "Gen6"
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
	// OEMSlotType is an OEM-specific slot.
	OEMSlotType SlotType = "OEM"
	// OCP3SmallSlotType Open Compute Project 3.0 small form factor slot.
	OCP3SmallSlotType SlotType = "OCP3Small"
	// OCP3LargeSlotType Open Compute Project 3.0 large form factor slot.
	OCP3LargeSlotType SlotType = "OCP3Large"
	// U2SlotType U.2 / SFF-8639 slot or bay.
	U2SlotType SlotType = "U2"
	// EDSFFSlotType EDSFF slot.
	EDSFFSlotType SlotType = "EDSFF"
)

// PCIeDevice shall represent a PCIe device in a Redfish implementation. It may
// also represent a location, such as a slot, socket, or bay, where a unit may
// be installed, but the 'State' property within the 'Status' property contains
// 'Absent'.
type PCIeDevice struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.2.0
	assembly string
	// AssetTag shall contain an identifying string that tracks the PCIe device for
	// inventory purposes.
	AssetTag string
	// CXLDevice shall contain CXL-specific properties of this PCIe device.
	//
	// Version added: v1.11.0
	CXLDevice CXLDevice
	// CXLLogicalDevices shall contain a link to a resource collection of type
	// 'CXLLogicalDeviceCollection'.
	//
	// Version added: v1.11.0
	cXLLogicalDevices string
	// DeviceType shall contain the device type of the PCIe device such as
	// 'SingleFunction' or 'MultiFunction'.
	DeviceType DeviceType
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this PCIe
	// device.
	//
	// Version added: v1.7.0
	environmentMetrics string
	// FirmwareVersion shall contain the firmware version of the PCIe device.
	FirmwareVersion string
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	//
	// Version added: v1.12.0
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the PCIe device. This organization may be the entity from whom the
	// PCIe device is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to
	// the PCIe device.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall contain a link to a resource collection of type
	// 'PCIeFunctionCollection'. This property should not be present if
	// 'DeviceType' contains 'Retimer'.
	//
	// Version added: v1.4.0
	pCIeFunctions string
	// PCIeInterface shall contain details for the PCIe interface that connects
	// this PCIe device to its host or upstream switch.
	//
	// Version added: v1.3.0
	PCIeInterface PCIeInterface
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the PCIe device.
	PartNumber string
	// ReadyToRemove shall indicate whether the PCIe device is ready for removal.
	// Setting the value to 'true' shall cause the service to perform appropriate
	// actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	//
	// Version added: v1.7.0
	ReadyToRemove bool
	// SKU shall contain the stock-keeping unit number for this PCIe device.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the PCIe device.
	SerialNumber string
	// Slot shall contain information about the PCIe slot for this PCIe device.
	//
	// Version added: v1.9.0
	Slot Slot
	// SparePartNumber shall contain the spare part number of the PCIe device.
	//
	// Version added: v1.6.0
	SparePartNumber string
	// StagedVersion shall contain the staged firmware version for this PCIe
	// device; this firmware is not yet active.
	//
	// Version added: v1.11.0
	StagedVersion string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UUID shall contain the universally unique identifier number for this PCIe
	// device.
	//
	// Version added: v1.5.0
	UUID string
	// chassis are the URIs for Chassis.
	chassis []string
	// connectedPCIePorts are the URIs for ConnectedPCIePorts.
	connectedPCIePorts []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctionsLink []string
	// processors are the URIs for Processors.
	processors []string
	// switch_ is the URI for Switch.
	switch_ string //nolint:revive
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PCIeDevice object from the raw JSON.
func (p *PCIeDevice) UnmarshalJSON(b []byte) error {
	type temp PCIeDevice
	type pLinks struct {
		Chassis            Links `json:"Chassis"`
		ConnectedPCIePorts Links `json:"ConnectedPCIePorts"`
		PCIeFunctions      Links `json:"PCIeFunctions"`
		Processors         Links `json:"Processors"`
		Switch             Link  `json:"Switch"`
	}
	var tmp struct {
		temp
		Links              pLinks
		Assembly           Link `json:"Assembly"`
		CXLLogicalDevices  Link `json:"CXLLogicalDevices"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		PCIeFunctions      Link `json:"PCIeFunctions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PCIeDevice(tmp.temp)

	// Extract the links to other entities for later
	p.chassis = tmp.Links.Chassis.ToStrings()
	p.connectedPCIePorts = tmp.Links.ConnectedPCIePorts.ToStrings()
	p.pCIeFunctionsLink = tmp.Links.PCIeFunctions.ToStrings()
	p.processors = tmp.Links.Processors.ToStrings()
	p.switch_ = tmp.Links.Switch.String()
	p.assembly = tmp.Assembly.String()
	p.cXLLogicalDevices = tmp.CXLLogicalDevices.String()
	p.environmentMetrics = tmp.EnvironmentMetrics.String()
	p.pCIeFunctions = tmp.PCIeFunctions.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PCIeDevice) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"LocationIndicatorActive",
		"ReadyToRemove",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPCIeDevice will get a PCIeDevice instance from the service.
func GetPCIeDevice(c Client, uri string) (*PCIeDevice, error) {
	return GetObject[PCIeDevice](c, uri)
}

// ListReferencedPCIeDevices gets the collection of PCIeDevice from
// a provided reference.
func ListReferencedPCIeDevices(c Client, link string) ([]*PCIeDevice, error) {
	return GetCollectionObjects[PCIeDevice](c, link)
}

// Chassis gets the Chassis linked resources.
func (p *PCIeDevice) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](p.client, p.chassis)
}

// ConnectedPCIePorts gets the ConnectedPCIePorts linked resources.
func (p *PCIeDevice) ConnectedPCIePorts() ([]*Port, error) {
	return GetObjects[Port](p.client, p.connectedPCIePorts)
}

// Processors gets the Processors linked resources.
func (p *PCIeDevice) Processors() ([]*Processor, error) {
	return GetObjects[Processor](p.client, p.processors)
}

// Switch gets the Switch linked resource.
func (p *PCIeDevice) Switch() (*Switch, error) {
	if p.switch_ == "" {
		return nil, nil
	}
	return GetObject[Switch](p.client, p.switch_)
}

// Assembly gets the Assembly linked resource.
func (p *PCIeDevice) Assembly() (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](p.client, p.assembly)
}

// CXLLogicalDevices gets the CXLLogicalDevices collection.
func (p *PCIeDevice) CXLLogicalDevices() ([]*CXLLogicalDevice, error) {
	if p.cXLLogicalDevices == "" {
		return nil, nil
	}
	return GetCollectionObjects[CXLLogicalDevice](p.client, p.cXLLogicalDevices)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (p *PCIeDevice) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if p.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](p.client, p.environmentMetrics)
}

// PCIeFunctions gets the PCIeFunctions collection.
func (p *PCIeDevice) PCIeFunctions() ([]*PCIeFunction, error) {
	if len(p.pCIeFunctionsLink) > 0 {
		return GetObjects[PCIeFunction](p.client, p.pCIeFunctionsLink)
	}

	if p.pCIeFunctions == "" {
		return nil, nil
	}
	return GetCollectionObjects[PCIeFunction](p.client, p.pCIeFunctions)
}

// CXLDevice shall contain CXL-specific properties of a PCIe device.
type CXLDevice struct {
	// CapableProtocolVersions shall contain an array of the CXL specification
	// revisions that this device supports.
	//
	// Version added: v1.18.0
	CapableProtocolVersions []CXLProtocolVersion
	// CurrentProtocolVersion shall contain the negotiated CXL specification
	// revision in use by this device.
	//
	// Version added: v1.18.0
	CurrentProtocolVersion CXLProtocolVersion
	// DeviceType shall contain the CXL device type.
	//
	// Version added: v1.11.0
	DeviceType CXLDeviceType
	// DynamicCapacity shall contain the CXL dynamic capacity device (DCD)
	// information for this CXL device.
	//
	// Version added: v1.12.0
	DynamicCapacity CXLDynamicCapacity
	// EgressPortCongestionSupport shall indicate whether the CXL device supports
	// the CXL Specification-defined 'Egress Port Congestion' mechanism.
	//
	// Version added: v1.11.0
	EgressPortCongestionSupport bool
	// MaxNumberLogicalDevices shall contain the maximum number of logical devices
	// supported by this CXL device.
	//
	// Version added: v1.11.0
	MaxNumberLogicalDevices *int `json:",omitempty"`
	// TemporaryThroughputReductionEnabled shall indicate whether the CXL
	// Specification-defined 'Temporary Throughput Reduction' mechanism is enabled
	// on this device.
	//
	// Version added: v1.14.0
	TemporaryThroughputReductionEnabled bool
	// TemporaryThroughputReductionSupported shall indicate whether the CXL
	// Specification-defined 'Temporary Throughput Reduction' mechanism is
	// supported on this device.
	//
	// Version added: v1.14.0
	TemporaryThroughputReductionSupported bool
	// ThroughputReductionSupport shall indicate whether the CXL device supports
	// the CXL Specification-defined 'Throughput Reduction' mechanism.
	//
	// Version added: v1.11.0
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of
	// 'TemporaryThroughputReductionSupported' to align with the CXL
	// Specification-defined FMAPI command.
	ThroughputReductionSupport bool
	// Timestamp shall contain the timestamp set on the CXL device.
	//
	// Version added: v1.11.0
	Timestamp string
}

// CXLDynamicCapacity shall contain the CXL dynamic capacity device (DCD)
// information for a CXL device.
type CXLDynamicCapacity struct {
	// AddCapacityPoliciesSupported shall contain the CXL Specification-defined
	// dynamic capacity policies that are supported by this CXL device when dynamic
	// capacity is added.
	//
	// Version added: v1.12.0
	AddCapacityPoliciesSupported []CXLDynamicCapacityPolicies
	// MaxDynamicCapacityRegions shall contain the maximum number of dynamic
	// capacity memory regions available per host from this CXL device.
	//
	// Version added: v1.12.0
	MaxDynamicCapacityRegions *int `json:",omitempty"`
	// MaxHosts shall contain the maximum number of hosts supported by this CXL
	// device.
	//
	// Version added: v1.12.0
	MaxHosts *int `json:",omitempty"`
	// MemoryBlockSizesSupported shall contain the set of memory block sizes
	// supported by memory regions in this CXL device.
	//
	// Version added: v1.12.0
	MemoryBlockSizesSupported []CXLRegionBlockSizes
	// ReleaseCapacityPoliciesSupported shall contain the CXL Specification-defined
	// dynamic capacity policies that are supported by this CXL device when dynamic
	// capacity is released.
	//
	// Version added: v1.12.0
	ReleaseCapacityPoliciesSupported []CXLDynamicCapacityPolicies
	// SanitizationOnReleaseSupport shall indicate whether the sanitization on
	// capacity release is configurable for the memory regions in this CXL device.
	//
	// Version added: v1.12.0
	SanitizationOnReleaseSupport []CXLRegionSanitization
	// TotalDynamicCapacityMiB shall contain the total memory media capacity of the
	// CXL device available for dynamic assignment in mebibytes (MiB).
	//
	// Version added: v1.12.0
	TotalDynamicCapacityMiB *int `json:",omitempty"`
}

// CXLRegionBlockSizes shall contain the set of memory block sizes supported by
// memory region in the dynamic capacity device.
type CXLRegionBlockSizes struct {
	// BlockSizeMiB shall contain the set of memory block sizes supported by this
	// memory region, with units in MiB.
	//
	// Version added: v1.12.0
	BlockSizeMiB []*int
	// RegionNumber shall contain the memory region number.
	//
	// Version added: v1.12.0
	RegionNumber *int `json:",omitempty"`
}

// CXLRegionSanitization shall indicate whether the sanitization on capacity
// release is configurable for the memory region.
type CXLRegionSanitization struct {
	// RegionNumber shall contain the memory region number.
	//
	// Version added: v1.12.0
	RegionNumber *int `json:",omitempty"`
	// SanitizationOnReleaseSupported shall indicate whether the sanitization on
	// capacity release is configurable for this memory region.
	//
	// Version added: v1.12.0
	SanitizationOnReleaseSupported bool
}

// PCIeErrors shall contain properties that describe the PCIe errors associated
// with this device.
type PCIeErrors struct {
	// BadDLLPCount shall contain the total number of Bad DLLPs issued on the PCIe
	// link by the receiver. A Bad DLLP in the context of PCIe communication is a
	// packet that has encountered errors at the data link layer. When a DLLP is
	// considered bad, it means it has been corrupted or is incorrectly formatted,
	// potentially due to transmission errors, hardware failures, or other issues
	// that affect its integrity.
	//
	// Version added: v1.15.0
	BadDLLPCount *int `json:",omitempty"`
	// BadTLPCount shall contain the total number of Bad TLPs issued on the PCIe
	// link by the receiver. A Bad TLP in the context of PCIe communication is a
	// packet that cannot be properly processed due to errors at the transaction
	// layer. These errors could include corrupted data, incorrect packet
	// formatting, invalid header information, or a mismatched checksum.
	//
	// Version added: v1.15.0
	BadTLPCount *int `json:",omitempty"`
	// CorrectableErrorCount shall contain the total number of PCIe correctable
	// errors for this device.
	//
	// Version added: v1.8.0
	CorrectableErrorCount *int `json:",omitempty"`
	// FatalErrorCount shall contain the total number of PCIe fatal errors for this
	// device.
	//
	// Version added: v1.8.0
	FatalErrorCount *int `json:",omitempty"`
	// FlowControlTimeoutErrors shall contain the total number of Flow Control
	// Credit timeouts, indicating a link-level problem that typically triggers
	// link recovery.
	//
	// Version added: v1.21.0
	FlowControlTimeoutErrors *int `json:",omitempty"`
	// L0ToRecoveryCount shall contain the total number of times the PCIe link
	// transitioned from L0 to the recovery state for this device.
	//
	// Version added: v1.8.0
	L0ToRecoveryCount *int `json:",omitempty"`
	// NAKReceivedCount shall contain the total number of NAKs issued on the PCIe
	// link by the receiver. A NAK is issued by the receiver when it detects that a
	// TLP from this device was missed. This could be because this device did not
	// transmit it, or because the receiver could not properly decode the packet.
	//
	// Version added: v1.8.0
	NAKReceivedCount *int `json:",omitempty"`
	// NAKSentCount shall contain the total number of NAKs issued on the PCIe link
	// by this device. A NAK is issued by the device when it detects that a TLP
	// from the receiver was missed. This could be because the receiver did not
	// transmit it, or because this device could not properly decode the packet.
	//
	// Version added: v1.8.0
	NAKSentCount *int `json:",omitempty"`
	// NonFatalErrorCount shall contain the total number of PCIe non-fatal errors
	// for this device.
	//
	// Version added: v1.8.0
	NonFatalErrorCount *int `json:",omitempty"`
	// ReplayCount shall contain the total number of replays issued on the PCIe
	// link by this device. A replay is a retransmission of a TLP and occurs
	// because the ACK timer is expired, which means that the receiver did not send
	// the ACK or this device did not properly decode the ACK.
	//
	// Version added: v1.8.0
	ReplayCount *int `json:",omitempty"`
	// ReplayRolloverCount shall contain the total number of replay rollovers
	// issued on the PCIe link by this device. A replay rollover occurs when
	// consecutive replays failed to resolve the errors on the link, which means
	// that this device forced the link into the recovery state.
	//
	// Version added: v1.8.0
	ReplayRolloverCount *int `json:",omitempty"`
	// UnsupportedRequestCount shall contain the total number of PCIe unsupported
	// requests received by this device.
	//
	// Version added: v1.13.0
	UnsupportedRequestCount *int `json:",omitempty"`
}

// PCIeInterface shall contain the definition for a PCIe interface for a Redfish
// implementation.
type PCIeInterface struct {
	// LanesInUse shall contain the number of PCIe lanes in use by this device,
	// which shall be equal to or less than the 'MaxLanes' property value.
	//
	// Version added: v1.3.0
	LanesInUse *int `json:",omitempty"`
	// MaxLanes shall contain the maximum number of PCIe lanes supported by this
	// device.
	//
	// Version added: v1.3.0
	MaxLanes *int `json:",omitempty"`
	// MaxPCIeType shall contain the maximum PCIe specification that this device
	// supports.
	//
	// Version added: v1.3.0
	MaxPCIeType PCIeTypes
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
	// PCIeType shall contain the negotiated PCIe interface version in use by this
	// device.
	//
	// Version added: v1.3.0
	PCIeType PCIeTypes
}

// PCIeMetrics shall contain properties that describe the PCIe metrics
// associated with this device.
type PCIeMetrics struct {
	// CompletionCreditExhaustionDrops shall contain the total number of read
	// requests dropped due to completion credit exhaustion.
	//
	// Version added: v1.20.0
	CompletionCreditExhaustionDrops *int `json:",omitempty"`
	// NPCreditExhaustionDrops shall contain the total number of of read requests
	// dropped due to non-posted credit exhaustion.
	//
	// Version added: v1.20.0
	NPCreditExhaustionDrops *int `json:",omitempty"`
	// OutboundCompletionTLPBytes shall contain the total data payload in bytes
	// transferred through outbound completion transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundCompletionTLPBytes *int `json:",omitempty"`
	// OutboundCompletionTLPCount shall contain the total number of outbound
	// completion transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundCompletionTLPCount *int `json:",omitempty"`
	// OutboundReadTLPBytes shall contain the total data payload in bytes
	// transferred through outbound read transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundReadTLPBytes *int `json:",omitempty"`
	// OutboundReadTLPCount shall contain the total number of outbound read
	// transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundReadTLPCount *int `json:",omitempty"`
	// OutboundWriteTLPBytes shall contain the total data payload in bytes
	// transferred through outbound write transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundWriteTLPBytes *int `json:",omitempty"`
	// OutboundWriteTLPCount shall contain the total number of outbound write
	// transaction layer packets (TLPs).
	//
	// Version added: v1.20.0
	OutboundWriteTLPCount *int `json:",omitempty"`
	// TagUnavailabilityDrops shall contain the total number of read requests
	// dropped due to tag unavailability.
	//
	// Version added: v1.20.0
	TagUnavailabilityDrops *int `json:",omitempty"`
}

// Slot shall contain properties that describe the PCIe slot associated with a
// PCIe device.
type Slot struct {
	// HotPluggable shall indicate whether this PCIe slot supports hotplug.
	//
	// Version added: v1.12.0
	HotPluggable bool
	// LaneSplitting shall contain lane splitting information of the associated
	// PCIe slot.
	//
	// Version added: v1.9.0
	LaneSplitting LaneSplittingType
	// Lanes shall contain the maximum number of PCIe lanes supported by the slot.
	//
	// Version added: v1.9.0
	Lanes *int `json:",omitempty"`
	// Location shall contain part location information, including a 'ServiceLabel'
	// property, of the associated PCIe slot.
	//
	// Version added: v1.9.0
	Location Location
	// PCIeType shall contain the maximum PCIe specification that this slot
	// supports.
	//
	// Version added: v1.9.0
	PCIeType PCIeTypes
	// SlotType shall contain the PCIe slot type.
	//
	// Version added: v1.9.0
	SlotType SlotType
}
