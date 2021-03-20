//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strconv"

	"github.com/stmcginnis/gofish/common"
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

// FpgaType is The FPGA type.
type FpgaType string

const (

	// IntegratedFpgaType The FPGA device integrated with other processor in
	// the single chip.
	IntegratedFpgaType FpgaType = "Integrated"
	// DiscreteFpgaType The discrete FPGA device.
	DiscreteFpgaType FpgaType = "Discrete"
)

// InstructionSet is the process instruction set used.
type InstructionSet string

const (

	// X86InstructionSet x86 32-bit.
	X86InstructionSet InstructionSet = "x86"
	// X86_64InstructionSet x86 64-bit.
	X8664InstructionSet InstructionSet = "x86-64"
	// IA64InstructionSet Intel IA-64.
	IA64InstructionSet InstructionSet = "IA-64"
	// ARMA32InstructionSet ARM 32-bit.
	ARMA32InstructionSet InstructionSet = "ARM-A32"
	// ARMA64InstructionSet ARM 64-bit.
	ARMA64InstructionSet InstructionSet = "ARM-A64"
	// MIPS32InstructionSet MIPS 32-bit.
	MIPS32InstructionSet InstructionSet = "MIPS32"
	// MIPS64InstructionSet MIPS 64-bit.
	MIPS64InstructionSet InstructionSet = "MIPS64"
	// PowerISAInstructionSet PowerISA-64 or PowerISA-32.
	PowerISAInstructionSet InstructionSet = "PowerISA"
	// OEMInstructionSet OEM-defined.
	OEMInstructionSet InstructionSet = "OEM"
)

// ProcessorArchitecture is processor architecture type.
type ProcessorArchitecture string

const (

	// X86ProcessorArchitecture x86 or x86-64.
	X86ProcessorArchitecture ProcessorArchitecture = "x86"
	// IA64ProcessorArchitecture Intel Itanium.
	IA64ProcessorArchitecture ProcessorArchitecture = "IA-64"
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

// ProcessorType is the processor type.
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

// FPGA shall contain the properties of the FPGA device represented by a
// Processor.
type FPGA struct {
	// ExternalInterfaces shall be an array of objects that describe the
	// external connectivity of the FPGA.
	ExternalInterfaces []FpgaInterface
	// FirmwareID shall contain a string describing the FPGA firmware
	// identifier.
	FirmwareID string `json:"FirmwareId"`
	// FirmwareManufacturer shall contain a string describing the FPGA firmware
	// manufacturer.
	FirmwareManufacturer string
	// FirmwareVersion shall contain a string describing the FPGA firmware
	// version.
	FirmwareVersion string
	// FpgaType shall be a type of the FPGA device.
	FpgaType FpgaType
	// HostInterface shall be an object that describes the connectivity to the
	// host for system software to use.
	HostInterface FpgaInterface
	// Model shall be a model of the FPGA device.
	Model string
	// PCIeVirtualFunctions shall be an integer that describes the number of
	// PCIe Virtual Functions configured within the FPGA.
	PCIeVirtualFunctions int
	// ProgrammableFromHost shall indicate
	// whether the FPGA firmware can be reprogrammed from the host using
	// system software. If set to false, system software shall not be able
	// to program the FPGA firmware from the host interface. In either
	// state, a management controller may be able to program the FPGA
	// firmware using the sideband interface.
	ProgrammableFromHost bool
	// ReconfigurationSlots shall be an array
	// of the structures describing the FPGA reconfiguration slots that can
	// be programmed with the acceleration functions.
	ReconfigurationSlots []FpgaReconfigurationSlot
}

// FpgaInterface shall contain information about the interface to the FPGA.
type FpgaInterface struct {
	// Ethernet shall be an object the
	// describes the Ethernet related information about this FPGA interface.
	Ethernet EthernetInterface
	// InterfaceType shall be an enum that
	// describes the type of interface to the FPGA.
	InterfaceType FpgaInterfaceType
	// pcie shall be an object the describes the PCI-e related information about
	// this FPGA interface. TODO: Get link to PCIeInterface.
	// pcie string
}

// FpgaReconfigurationSlot shall contain information about the FPGA
// reconfiguration slot.
type FpgaReconfigurationSlot struct {
	// AccelerationFunction shall be a reference to the acceleration function
	// resources provided by the code programmed into a reconfiguration slot and
	// shall reference a resource of type AccelerationFunction.
	// TODO: Get link to resource.
	// accelerationFunction string
	// ProgrammableFromHost shall indicate
	// whether the reconfiguration slot can be reprogrammed from the host
	// using system software. If set to false, system software shall not be
	// able to program the reconfiguration slot from the host interface. In
	// either state, a management controller may be able to program the
	// reconfiguration slot using the sideband interface.
	ProgrammableFromHost bool
	// SlotID shall be the FPGA reconfiguration slot identifier.
	SlotID string `json:"SlotId"`
	// UUID is used to contain a universal unique identifier number for the
	// reconfiguration slot.
	UUID string
}

// Processor is used to represent a single processor contained within a
// system.
type Processor struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// accelerationFunctions shall be a link to
	// a collection of type AccelerationFunctionCollection.
	accelerationFunctions string
	// Actions is The Actions property shall contain the available actions
	// for this resource.
	Actions string
	// Assembly shall be a link to a resource
	// of type Assembly.
	assembly string
	// Description provides a description of this resource.
	Description string
	// FPGA shall be an object containing
	// properties specific for Processors of type FPGA.
	FPGA FPGA
	// InstructionSet shall contain the string which
	// identifies the instruction set of the processor contained in this
	// socket.
	InstructionSet InstructionSet
	// Links is The Links property, as described by the Redfish
	// Specification, shall contain references to resources that are related
	// to, but not contained by (subordinate to), this resource.
	Links string
	// Location shall contain location information of the
	// associated processor.
	Location common.Location
	// Manufacturer shall contain a string which identifies
	// the manufacturer of the processor.
	Manufacturer string
	// MaxSpeedMHz shall indicate the maximum rated clock
	// speed of the processor in MHz.
	MaxSpeedMHz float32
	// MaxTDPWatts shall be the maximum Thermal
	// Design Power (TDP) in watts.
	MaxTDPWatts int
	// Metrics shall be a reference to the Metrics
	// associated with this Processor.
	metrics string
	// Model shall indicate the model information as
	// provided by the manufacturer of this processor.
	Model string
	// ProcessorArchitecture shall contain the string which
	// identifies the architecture of the processor contained in this Socket.
	ProcessorArchitecture ProcessorArchitecture
	// ProcessorID shall contain identification information for this processor.
	ProcessorID ProcessorID `json:"ProcessorId"`
	// ProcessorMemory shall be the memory
	// directly attached or integrated within this Processor.
	processorMemory []string
	// ProcessorType shall contain the string which
	// identifies the type of processor contained in this Socket.
	ProcessorType ProcessorType
	// Socket shall contain the string which identifies the
	// physical location or socket of the processor.
	Socket string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// SubProcessors shall be a link to a
	// collection of type ProcessorCollection.
	SubProcessors string
	// TDPWatts shall be the nominal Thermal
	// Design Power (TDP) in watts.
	TDPWatts int
	// TotalCores shall indicate the total count of
	// independent processor cores contained within this processor.
	TotalCores int
	// TotalEnabledCores shall indicate the total count of
	// enabled independent processor cores contained within this processor.
	TotalEnabledCores int
	// TotalThreads shall indicate the total count of
	// independent execution threads supported by this processor.
	TotalThreads int
	// UUID is used to contain a universal unique identifier number for the
	// processor. RFC4122 describes methods that can be used to create the
	// value. The value should be considered to be opaque. Client software
	// should only treat the overall value as a universally unique identifier
	// and should not interpret any sub-fields within the UUID.
	UUID string
	// Chassis shall be a reference to a
	// resource of type Chassis that represent the physical container
	// associated with this Processor.
	chassis string
	// ConnectedProcessors shall be an array of
	// references of type Processor that are directly connected to this
	// Processor.
	connectedProcessors []string
	// ConnectedProcessors@odata.count is
	ConnectedProcessorsCount int
	// Endpoints shall be an array of
	// references of type Endpoint that represent Endpoints associated with
	// this Processor.
	endpoints []string
	// Endpoints@odata.count is
	EndpointsCount int
	// PCIeDevice shall be a reference of type
	// PCIeDevice that represents the PCI-e Device associated with this
	// Processor.
	pcieDevice string
	// PCIeFunctions shall be an array of
	// references of type PCIeFunction that represent the PCI-e Functions
	// associated with this Processor.
	pcieFunctions []string
	// PCIeFunctions@odata.count is
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (processor *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	type t1 struct {
		temp
		AccelerationFunctions common.Link
		Assembly              common.Link
		Metrics               common.Link
		ProcessorMemory       common.Links
		Links                 struct {
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
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		// Handle invalid data type returned for MaxSpeedMHz
		var t2 struct {
			t1
			MaxSpeedMHz string
		}
		err2 := json.Unmarshal(b, &t2)

		if err2 != nil {
			// Return the original error
			return err
		}

		// Extract the real Processor struct and replace its MaxSpeedMHz with
		// the parsed string version
		t = t2.t1
		if t2.MaxSpeedMHz != "" {
			mhz, err := strconv.ParseFloat(t2.MaxSpeedMHz, 32)
			if err != nil {
				t.MaxSpeedMHz = float32(mhz)
			}
		}
	}

	*processor = Processor(t.temp)

	// Extract the links to other entities for later
	processor.accelerationFunctions = string(t.AccelerationFunctions)
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
	processor.metrics = string(t.Metrics)

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

// ProcessorID shall contain identification information for a processor.
type ProcessorID struct {
	// EffectiveFamily shall indicate the effective Family
	// information as provided by the manufacturer of this processor.
	EffectiveFamily string
	// EffectiveModel shall indicate the effective Model
	// information as provided by the manufacturer of this processor.
	EffectiveModel string
	// IdentificationRegisters shall include the raw CPUID
	// instruction output as provided by the manufacturer of this processor.
	IdentificationRegisters string
	// MicrocodeInfo shall indicate the Microcode
	// Information as provided by the manufacturer of this processor.
	MicrocodeInfo string
	// Step shall indicate the Step or revision string
	// information as provided by the manufacturer of this processor.
	Step string
	// VendorID shall indicate the Vendor Identification
	// string information as provided by the manufacturer of this processor.
	VendorID string `json:"VendorId"`
}

// ProcessorMemory shall contain information about memory
// directly attached or integrated within a processor.
type ProcessorMemory struct {
	// CapacityMiB shall be the memory capacity in MiB.
	CapacityMiB int
	// IntegratedMemory shall be a boolean
	// indicating whether this memory is integrated within the Processor.
	// Otherwise it is discrete memory attached to the Processor.
	IntegratedMemory bool
	// MemoryType shall be a type of the processor memory type.
	MemoryType ProcessorMemoryType
	// SpeedMHz shall be the operating speed of the memory in MHz.
	SpeedMHz int
}
