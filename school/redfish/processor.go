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

// FpgaInterfaceType is
type FpgaInterfaceType string

const (

	// QPIFpgaInterfaceType The Intel QuickPath Interconnect.
	QPIFpgaInterfaceType FpgaInterfaceType = "QPI"
	// UPIFpgaInterfaceType The Intel UltraPath Interconnect.
	UPIFpgaInterfaceType FpgaInterfaceType = "UPI"
	// PCIeFpgaInterfaceType A PCI Express interface.
	PCIeFpgaInterfaceType FpgaInterfaceType = "PCIe"
	// EthernetFpgaInterfaceType An Ethernet interface.
	EthernetFpgaInterfaceType FpgaInterfaceType = "Ethernet"
	// OEMFpgaInterfaceType An OEM defined interface.
	OEMFpgaInterfaceType FpgaInterfaceType = "OEM"
)

// FpgaType is
type FpgaType string

const (

	// IntegratedFpgaType The FPGA device integrasted with other porcessor in
	// the single chip.
	IntegratedFpgaType FpgaType = "Integrated"
	// DiscreteFpgaType The discrete FPGA device.
	DiscreteFpgaType FpgaType = "Discrete"
)

// InstructionSet is
type InstructionSet string

const (

	// x86InstructionSet x86 32-bit.
	x86InstructionSet InstructionSet = "x86"
	// x86_64InstructionSet x86 64-bit.
	x86_64InstructionSet InstructionSet = "x86-64"
	// IA_64InstructionSet Intel IA-64.
	IA_64InstructionSet InstructionSet = "IA-64"
	// ARM_A32InstructionSet ARM 32-bit.
	ARM_A32InstructionSet InstructionSet = "ARM-A32"
	// ARM_A64InstructionSet ARM 64-bit.
	ARM_A64InstructionSet InstructionSet = "ARM-A64"
	// MIPS32InstructionSet MIPS 32-bit.
	MIPS32InstructionSet InstructionSet = "MIPS32"
	// MIPS64InstructionSet MIPS 64-bit.
	MIPS64InstructionSet InstructionSet = "MIPS64"
	// PowerISAInstructionSet PowerISA-64 or PowerISA-32.
	PowerISAInstructionSet InstructionSet = "PowerISA"
	// OEMInstructionSet OEM-defined.
	OEMInstructionSet InstructionSet = "OEM"
)

// ProcessorArchitecture is
type ProcessorArchitecture string

const (

	// x86ProcessorArchitecture x86 or x86-64.
	x86ProcessorArchitecture ProcessorArchitecture = "x86"
	// IA_64ProcessorArchitecture Intel Itanium.
	IA_64ProcessorArchitecture ProcessorArchitecture = "IA-64"
	// ARMProcessorArchitecture ARM.
	ARMProcessorArchitecture ProcessorArchitecture = "ARM"
	// MIPSProcessorArchitecture MIPS.
	MIPSProcessorArchitecture ProcessorArchitecture = "MIPS"
	// PowerProcessorArchitecture Power.
	PowerProcessorArchitecture ProcessorArchitecture = "Power"
	// OEMProcessorArchitecture OEM-defined.
	OEMProcessorArchitecture ProcessorArchitecture = "OEM"
)

// ProcessorMemoryType is
type ProcessorMemoryType string

const (

	// L1CacheProcessorMemoryType L1 cache.
	L1CacheProcessorMemoryType ProcessorMemoryType = "L1Cache"
	// L2CacheProcessorMemoryType L2 cache.
	L2CacheProcessorMemoryType ProcessorMemoryType = "L2Cache"
	// L3CacheProcessorMemoryType L3 cache.
	L3CacheProcessorMemoryType ProcessorMemoryType = "L3Cache"
	// L4CacheProcessorMemoryType L4 cache.
	L4CacheProcessorMemoryType ProcessorMemoryType = "L4Cache"
	// L5CacheProcessorMemoryType L5 cache.
	L5CacheProcessorMemoryType ProcessorMemoryType = "L5Cache"
	// L6CacheProcessorMemoryType L6 cache.
	L6CacheProcessorMemoryType ProcessorMemoryType = "L6Cache"
	// L7CacheProcessorMemoryType L7 cache.
	L7CacheProcessorMemoryType ProcessorMemoryType = "L7Cache"
	// HBM1ProcessorMemoryType High Bandwidth Memory.
	HBM1ProcessorMemoryType ProcessorMemoryType = "HBM1"
	// HBM2ProcessorMemoryType The second generation of High Bandwidth
	// Memory.
	HBM2ProcessorMemoryType ProcessorMemoryType = "HBM2"
	// HBM3ProcessorMemoryType The third generation of High Bandwidth Memory.
	HBM3ProcessorMemoryType ProcessorMemoryType = "HBM3"
	// SGRAMProcessorMemoryType Synchronous graphics RAM.
	SGRAMProcessorMemoryType ProcessorMemoryType = "SGRAM"
	// GDDRProcessorMemoryType Synchronous graphics random-access memory.
	GDDRProcessorMemoryType ProcessorMemoryType = "GDDR"
	// GDDR2ProcessorMemoryType Double data rate type two synchronous
	// graphics random-access memory.
	GDDR2ProcessorMemoryType ProcessorMemoryType = "GDDR2"
	// GDDR3ProcessorMemoryType Double data rate type three synchronous
	// graphics random-access memory.
	GDDR3ProcessorMemoryType ProcessorMemoryType = "GDDR3"
	// GDDR4ProcessorMemoryType Double data rate type four synchronous
	// graphics random-access memory.
	GDDR4ProcessorMemoryType ProcessorMemoryType = "GDDR4"
	// GDDR5ProcessorMemoryType Double data rate type five synchronous
	// graphics random-access memory.
	GDDR5ProcessorMemoryType ProcessorMemoryType = "GDDR5"
	// GDDR5XProcessorMemoryType Double data rate type five synchronous
	// graphics random-access memory.
	GDDR5XProcessorMemoryType ProcessorMemoryType = "GDDR5X"
	// GDDR6ProcessorMemoryType Double data rate type five synchronous
	// graphics random-access memory.
	GDDR6ProcessorMemoryType ProcessorMemoryType = "GDDR6"
	// DDRProcessorMemoryType Double data rate synchronous dynamic random-
	// access memory.
	DDRProcessorMemoryType ProcessorMemoryType = "DDR"
	// DDR2ProcessorMemoryType Double data rate type two synchronous dynamic
	// random-access memory.
	DDR2ProcessorMemoryType ProcessorMemoryType = "DDR2"
	// DDR3ProcessorMemoryType Double data rate type three synchronous
	// dynamic random-access memory.
	DDR3ProcessorMemoryType ProcessorMemoryType = "DDR3"
	// DDR4ProcessorMemoryType Double data rate type four synchronous dynamic
	// random-access memory.
	DDR4ProcessorMemoryType ProcessorMemoryType = "DDR4"
	// DDR5ProcessorMemoryType Double data rate type five synchronous dynamic
	// random-access memory.
	DDR5ProcessorMemoryType ProcessorMemoryType = "DDR5"
	// SDRAMProcessorMemoryType Synchronous dynamic random-access memory.
	SDRAMProcessorMemoryType ProcessorMemoryType = "SDRAM"
	// SRAMProcessorMemoryType Static random-access memory.
	SRAMProcessorMemoryType ProcessorMemoryType = "SRAM"
	// FlashProcessorMemoryType Flash memory.
	FlashProcessorMemoryType ProcessorMemoryType = "Flash"
	// OEMProcessorMemoryType OEM-defined.
	OEMProcessorMemoryType ProcessorMemoryType = "OEM"
)

// ProcessorType is
type ProcessorType string

const (

	// CPUProcessorType A Central Processing Unit.
	CPUProcessorType ProcessorType = "CPU"
	// GPUProcessorType A Graphics Processing Unit.
	GPUProcessorType ProcessorType = "GPU"
	// FPGAProcessorType A Field Programmable Gate Array.
	FPGAProcessorType ProcessorType = "FPGA"
	// DSPProcessorType A Digital Signal Processor.
	DSPProcessorType ProcessorType = "DSP"
	// AcceleratorProcessorType An Accelerator.
	AcceleratorProcessorType ProcessorType = "Accelerator"
	// CoreProcessorType A Core in a Processor.
	CoreProcessorType ProcessorType = "Core"
	// ThreadProcessorType A Thread in a Processor.
	ThreadProcessorType ProcessorType = "Thread"
	// OEMProcessorType An OEM-defined Processing Unit.
	OEMProcessorType ProcessorType = "OEM"
)

// FPGA is This object shall contain the properties of the FPGA device
// represented by a Processor.
type FPGA struct {
	common.Entity

	// ExternalInterfaces is The value of this property shall be an array of
	// objects that describe the external connectivity of the FPGA.
	ExternalInterfaces []FpgaInterface
	// FirmwareId is The value of this property shall contain a string
	// decsribing the FPGA firmware identifier.
	FirmwareId string
	// FirmwareManufacturer is The value of this property shall contain a
	// string decsribing the FPGA firmware manufacturer.
	FirmwareManufacturer string
	// FirmwareVersion is The value of this property shall contain a string
	// decsribing the FPGA firmware version.
	FirmwareVersion string
	// FpgaType is The value of this property shall be a type of the FPGA
	// device.
	FpgaType string
	// HostInterface is The value of this property shall be an object that
	// describes the connectivity to the host for system software to use.
	HostInterface string
	// Model is The value of this property shall be a model of the FPGA
	// device.
	Model string
	// Oem is This object represents the Oem property.  All values for
	// resources described by this schema shall comply to the requirements as
	// described in the Redfish specification.
	OEM string `json:"Oem"`
	// PCIeVirtualFunctions is The value of this property shall be an integer
	// that describes the number of PCIe Virtual Functions configured within
	// the FPGA.
	PCIeVirtualFunctions string
	// ProgrammableFromHost is The value of this property shall indicate
	// whether the FPGA firmware can be reprogrammed from the host using
	// system software.  If set to false, system software shall not be able
	// to program the FPGA firmware from the host interface.  In either
	// state, a management controller may be able to program the FPGA
	// firmware using the sideband interface.
	ProgrammableFromHost bool
	// ReconfigurationSlots is The value of this property shall be an array
	// of the structures describing the FPGA reconfiguration slots that can
	// be programmed with the acceleration functions.
	ReconfigurationSlots []FpgaReconfigurationSlot
}

// UnmarshalJSON unmarshals a FPGA object from the raw JSON.
func (fpga *FPGA) UnmarshalJSON(b []byte) error {
	type temp FPGA
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fpga = FPGA(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FpgaInterface is This type shall contain information about the
// interface to the FPGA.
type FpgaInterface struct {
	common.Entity

	// Ethernet is The value of this property shall be an object the
	// describes the Ethernet related information about this FPGA interface.
	Ethernet string
	// InterfaceType is The value of this property shall be an enum that
	// describes the type of interface to the FPGA.
	InterfaceType FpgaInterfaceType
	// PCIe is The value of this property shall be an object the describes
	// the PCI-e related information about this FPGA interface.
	PCIe string
}

// UnmarshalJSON unmarshals a FpgaInterface object from the raw JSON.
func (fpgainterface *FpgaInterface) UnmarshalJSON(b []byte) error {
	type temp FpgaInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fpgainterface = FpgaInterface(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FpgaReconfigurationSlot is This type shall contain information about
// the FPGA reconfiguration slot.
type FpgaReconfigurationSlot struct {
	common.Entity

	// AccelerationFunction is The value of this property shall be a
	// reference to the acceleration function resources provided by the code
	// programmed into a reconfiguration slot and shall reference a resource
	// of type AccelerationFunction.
	AccelerationFunction string
	// ProgrammableFromHost is The value of this property shall indicate
	// whether the reconfiguration slot can be reprogrammed from the host
	// using system software.  If set to false, system software shall not be
	// able to program the reconfiguration slot from the host interface.  In
	// either state, a management controller may be able to program the
	// reconfiguration slot using the sideband interface.
	ProgrammableFromHost bool
	// SlotId is The value of this property shall be the FPGA reconfiguration
	// slot identifier.
	SlotId string
	// UUID is used to contain a universal unique identifier number for the
	// reconfiguration slot.
	UUID string
}

// UnmarshalJSON unmarshals a FpgaReconfigurationSlot object from the raw JSON.
func (fpgareconfigurationslot *FpgaReconfigurationSlot) UnmarshalJSON(b []byte) error {
	type temp FpgaReconfigurationSlot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fpgareconfigurationslot = FpgaReconfigurationSlot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OemActions is This type shall contain any additional OEM actions for
// this resource.
type OemActions struct {
	common.Entity
}

// UnmarshalJSON unmarshals a OemActions object from the raw JSON.
func (oemactions *OemActions) UnmarshalJSON(b []byte) error {
	type temp OemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*oemactions = OemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Processor is used to represent a single processor contained within a
// system.
type Processor struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataId string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccelerationFunctions is The value of this property shall be a link to
	// a collection of type AccelerationFunctionCollection.
	AccelerationFunctions string
	// Actions is The Actions property shall contain the available actions
	// for this resource.
	Actions string
	// Assembly is The value of this property shall be a link to a resource
	// of type Assembly.
	assembly string
	// Description provides a description of this resource.
	Description string
	// FPGA is The value of this property shall be an object containing
	// properties specific for Processors of type FPGA.
	FPGA string
	// InstructionSet is This property shall contain the string which
	// identifies the instruction set of the processor contained in this
	// socket.
	InstructionSet InstructionSet
	// Links is The Links property, as described by the Redfish
	// Specification, shall contain references to resources that are related
	// to, but not contained by (subordinate to), this resource.
	Links string
	// Location is This property shall contain location information of the
	// associated processor.
	Location string
	// Manufacturer is This property shall contain a string which identifies
	// the manufacturer of the processor.
	Manufacturer string
	// MaxSpeedMHz is This property shall indicate the maximum rated clock
	// speed of the processor in MHz.
	MaxSpeedMHz int
	// MaxTDPWatts is The value of this property shall be the maximum Thermal
	// Design Power (TDP) in watts.
	MaxTDPWatts int
	// Metrics is This property shall be a reference to the Metrics
	// associated with this Processor.
	Metrics string
	// Model is This property shall indicate the model information as
	// provided by the manufacturer of this processor.
	Model string
	// ProcessorArchitecture is This property shall contain the string which
	// identifies the architecture of the processor contained in this Socket.
	ProcessorArchitecture ProcessorArchitecture
	// ProcessorId is This object shall contain identification information
	// for this processor.
	ProcessorId ProcessorId
	// ProcessorMemory is The value of this property shall be the memory
	// directly attached or integrated witin this Processor.
	processorMemory []string
	// ProcessorType is This property shall contain the string which
	// identifies the type of processor contained in this Socket.
	ProcessorType ProcessorType
	// Socket is This property shall contain the string which identifies the
	// physical location or socket of the processor.
	Socket string
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status
	// SubProcessors is The value of this property shall be a link to a
	// collection of type ProcessorCollection.
	SubProcessors string
	// TDPWatts is The value of this property shall be the nominal Thermal
	// Design Power (TDP) in watts.
	TDPWatts int
	// TotalCores is This property shall indicate the total count of
	// independent processor cores contained within this processor.
	TotalCores int
	// TotalEnabledCores is This property shall indicate the total count of
	// enabled independent processor cores contained within this processor.
	TotalEnabledCores int
	// TotalThreads is This property shall indicate the total count of
	// independent execution threads supported by this processor.
	TotalThreads int
	// UUID is used to contain a universal unique identifier number for the
	// processor.  RFC4122 describes methods that can be used to create the
	// value.  The value should be considered to be opaque.  Client software
	// should only treat the overall value as a universally unique identifier
	// and should not interpret any sub-fields within the UUID.
	UUID string

	// Chassis is The value of this property shall be a reference to a
	// resource of type Chassis that represent the physical container
	// associated with this Processor.
	chassis string
	// ConnectedProcessors is The value of this property shall be an array of
	// references of type Processor that are directly connected to this
	// Processor.
	connectedProcessors []string
	// ConnectedProcessors@odata.count is
	ConnectedProcessorsCount int
	// Endpoints is The value of this property shall be an array of
	// references of type Endpoint that represent Endpoints accociated with
	// this Processor.
	endpoints []string
	// Endpoints@odata.count is
	EndpointsCount int
	// PCIeDevice is The value of this property shall be a reference of type
	// PCIeDevice that represents the PCI-e Device associated with this
	// Processor.
	pcieDevice string
	// PCIeFunctions is The value of this property shall be an array of
	// references of type PCIeFunction that represent the PCI-e Functions
	// associated with this Processor.
	pcieFunctions []string
	// PCIeFunctions@odata.count is
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (processor *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	var t struct {
		temp
		Assembly        common.Link
		ProcessorMemory common.Links
		Links           struct {
			Chassis                  common.Link
			ConnectedProcessors      common.Links
			ConnectedProcessorsCount int `json:"ConnectedProcessors@odata.count"`
			Endpoints                common.Links
			EndpointsCount           int `json:"Endpoints@odata.count"`
			PCIeDevice               common.Link
			PCIeFunctions            common.Links
			PCIeFunctionsCount       int `json:"PCIeFunctions@odata.count"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processor = Processor(t.temp)

	// Extract the links to other entities for later
	processor.assembly = string(t.Assembly)
	processor.chassis = string(t.Links.Chassis)
	processor.processorMemory = t.ProcessorMemory.ToStrings()
	processor.connectedProcessors = t.Links.ConnectedProcessors.ToStrings()
	processor.ConnectedProcessorsCount = t.Links.ConnectedProcessorsCount
	processor.endpoints = t.Links.Endpoints.ToStrings()
	processor.EndpointsCount = t.Links.EndpointsCount
	processor.pcieDevice = string(t.Links.PCIeDevice)
	processor.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	processor.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	return nil
}

// GetProcessor will get a Processor instance from the system
func GetProcessor(c common.Client, uri string) (*Processor, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var processor Processor
	err = json.NewDecoder(resp.Body).Decode(&processor)
	if err != nil {
		return nil, err
	}

	processor.SetClient(c)
	return &processor, nil
}

// ListReferencedProcessors gets the collection of Processor from a provided reference.
func ListReferencedProcessors(c common.Client, link string) ([]*Processor, error) {
	var result []*Processor
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, processorLink := range links.ItemLinks {
		processor, err := GetProcessor(c, processorLink)
		if err != nil {
			return result, err
		}
		result = append(result, processor)
	}

	return result, nil
}

// ProcessorId is This type shall contain identification information for
// a processor.
type ProcessorId struct {
	common.Entity

	// EffectiveFamily is This property shall indicate the effective Family
	// information as provided by the manufacturer of this processor.
	EffectiveFamily string
	// EffectiveModel is This property shall indicate the effective Model
	// information as provided by the manufacturer of this processor.
	EffectiveModel string
	// IdentificationRegisters is This property shall include the raw CPUID
	// instruction output as provided by the manufacturer of this processor.
	IdentificationRegisters string
	// MicrocodeInfo is This property shall indicate the Microcode
	// Information as provided by the manufacturer of this processor.
	MicrocodeInfo string
	// Step is This property shall indicate the Step or revision string
	// information as provided by the manufacturer of this processor.
	Step string
	// VendorId is This property shall indicate the Vendor Identification
	// string information as provided by the manufacturer of this processor.
	VendorId string
}

// UnmarshalJSON unmarshals a ProcessorId object from the raw JSON.
func (processorid *ProcessorId) UnmarshalJSON(b []byte) error {
	type temp ProcessorId
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorid = ProcessorId(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ProcessorMemory is This type shall contain information about memory
// directly attached or integratied within a processor.
type ProcessorMemory struct {
	common.Entity

	// CapacityMiB is The value of this property shall be the memory capacity
	// in MiB.
	CapacityMiB int
	// IntegratedMemory is The value of this property shall be a boolean
	// indicating whether this memory is integrated within the Porcessor.
	// Otherwise it is discrete memory attached to the Processor.
	IntegratedMemory bool
	// MemoryType is The value of this property shall be a type of the
	// processor memory type.
	MemoryType ProcessorMemoryType
	// SpeedMHz is The value of this property shall be the operating speed of
	// the memory in MHz.
	SpeedMHz int
}

// UnmarshalJSON unmarshals a ProcessorMemory object from the raw JSON.
func (processormemory *ProcessorMemory) UnmarshalJSON(b []byte) error {
	type temp ProcessorMemory
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processormemory = ProcessorMemory(t.temp)

	// Extract the links to other entities for later

	return nil
}
