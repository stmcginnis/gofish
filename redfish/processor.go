//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/stmcginnis/gofish/common"
)

// The state of the base frequency settings of
// the operation configuration applied to this processor.
type BaseSpeedPriorityState string

const (
	// Base speed priority is disabled.
	DisabledBaseSpeedPriority BaseSpeedPriorityState = "Disabled"
	// Base speed priority is enabled.
	EnabledBaseSpeedPriority BaseSpeedPriorityState = "Enabled"
)

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
	// X8664InstructionSet x86 64-bit.
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
	// RV32InstructionSet RISC-V 32-bit.
	RV32InstructionSet InstructionSet = "RV32"
	// RV64InstructionSet RISC-V 64-bit.
	RV64InstructionSet InstructionSet = "RV64"
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
	// RISCVProcessorArchitecture RISC-V.
	RISCVProcessorArchitecture ProcessorArchitecture = "RISC-V"
	// OEMProcessorArchitecture OEM-defined.
	OEMProcessorArchitecture ProcessorArchitecture = "OEM"
)

type ProcessorMemoryType string

const (
	// CacheProcessorMemoryType Processor cache, but no level is determined.
	CacheProcessorMemoryType ProcessorMemoryType = "Cache"
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
	// HBM2ProcessorMemoryType The second generation of High Bandwidth Memory.
	HBM2ProcessorMemoryType ProcessorMemoryType = "HBM2"
	// HBM2EProcessorMemoryType An updated version of the second generation of High Bandwidth Memory.
	HBM2EProcessorMemoryType ProcessorMemoryType = "HBM2E"
	// HBM3ProcessorMemoryType The third generation of High Bandwidth Memory.
	HBM3ProcessorMemoryType ProcessorMemoryType = "HBM3"
	// SGRAMProcessorMemoryType Synchronous graphics RAM.
	SGRAMProcessorMemoryType ProcessorMemoryType = "SGRAM"
	// GDDRProcessorMemoryType Synchronous graphics random-access memory.
	GDDRProcessorMemoryType ProcessorMemoryType = "GDDR"
	// GDDR2ProcessorMemoryType Double data rate type two synchronous graphics random-access memory.
	GDDR2ProcessorMemoryType ProcessorMemoryType = "GDDR2"
	// GDDR3ProcessorMemoryType Double data rate type three synchronous graphics random-access memory.
	GDDR3ProcessorMemoryType ProcessorMemoryType = "GDDR3"
	// GDDR4ProcessorMemoryType Double data rate type four synchronous graphics random-access memory.
	GDDR4ProcessorMemoryType ProcessorMemoryType = "GDDR4"
	// GDDR5ProcessorMemoryType Double data rate type five synchronous graphics random-access memory.
	GDDR5ProcessorMemoryType ProcessorMemoryType = "GDDR5"
	// GDDR5XProcessorMemoryType Double data rate type five X synchronous graphics random-access memory.
	GDDR5XProcessorMemoryType ProcessorMemoryType = "GDDR5X"
	// GDDR6ProcessorMemoryType Double data rate type six synchronous graphics random-access memory.
	GDDR6ProcessorMemoryType ProcessorMemoryType = "GDDR6"
	// DDRProcessorMemoryType Double data rate synchronous dynamic random-access memory.
	DDRProcessorMemoryType ProcessorMemoryType = "DDR"
	// DDR2ProcessorMemoryType Double data rate type two synchronous dynamic random-access memory.
	DDR2ProcessorMemoryType ProcessorMemoryType = "DDR2"
	// DDR3ProcessorMemoryType Double data rate type three synchronous dynamic random-access memory.
	DDR3ProcessorMemoryType ProcessorMemoryType = "DDR3"
	// DDR4ProcessorMemoryType Double data rate type four synchronous dynamic random-access memory.
	DDR4ProcessorMemoryType ProcessorMemoryType = "DDR4"
	// DDR5ProcessorMemoryType Double data rate type five synchronous dynamic random-access memory.
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
	// CPUProcessorType A CPU.
	CPUProcessorType ProcessorType = "CPU"
	// GPUProcessorType A GPU.
	GPUProcessorType ProcessorType = "GPU"
	// FPGAProcessorType An FPGA.
	FPGAProcessorType ProcessorType = "FPGA"
	// DSPProcessorType A DSP.
	DSPProcessorType ProcessorType = "DSP"
	// AcceleratorProcessorType An accelerator.
	AcceleratorProcessorType ProcessorType = "Accelerator"
	// CoreProcessorType A core in a processor.
	CoreProcessorType ProcessorType = "Core"
	// ThreadProcessorType A thread in a processor.
	ThreadProcessorType ProcessorType = "Thread"
	// PartitionProcessorType shall indicate a partition in a processor that is instantiated from a user configuration
	// to carve out resources in a single processor. An example of this is assigning memory to a set of cores in a GPU.
	PartitionProcessorType ProcessorType = "Partition"
	// OEMProcessorType An OEM-defined processing unit.
	OEMProcessorType ProcessorType = "OEM"
)

type SystemInterfaceType string

const (
	// QPISystemInterfaceType The Intel QuickPath Interconnect.
	QPISystemInterfaceType SystemInterfaceType = "QPI"
	// UPISystemInterfaceType The Intel UltraPath Interconnect.
	UPISystemInterfaceType SystemInterfaceType = "UPI"
	// PCIeSystemInterfaceType A PCI Express interface.
	PCIeSystemInterfaceType SystemInterfaceType = "PCIe"
	// EthernetSystemInterfaceType An Ethernet interface.
	EthernetSystemInterfaceType SystemInterfaceType = "Ethernet"
	// AMBASystemInterfaceType The Arm Advanced Microcontroller Bus Architecture interface.
	AMBASystemInterfaceType SystemInterfaceType = "AMBA"
	// CCIXSystemInterfaceType The Cache Coherent Interconnect for Accelerators interface.
	CCIXSystemInterfaceType SystemInterfaceType = "CCIX"
	// CXLSystemInterfaceType The Compute Express Link interface.
	CXLSystemInterfaceType SystemInterfaceType = "CXL"
	// OEMSystemInterfaceType An OEM-defined interface.
	OEMSystemInterfaceType SystemInterfaceType = "OEM"
)

// The causes of the processor being throttled.
type ThrottleCause string

const (
	// PowerLimitThrottleCause The cause of the processor being throttled is a power limit.
	PowerLimitThrottleCause ThrottleCause = "PowerLimit"
	// ThermalLimitThrottleCause The cause of the processor being throttled is a thermal limit.
	ThermalLimitThrottleCause ThrottleCause = "ThermalLimit"
	// ClockLimitThrottleCause The cause of the processor being throttled is a clock limit.
	ClockLimitThrottleCause ThrottleCause = "ClockLimit"
	// ManagementDetectedFaultThrottleCause The cause of the processor being throttled is a fault detected by
	// management hardware or firmware.
	ManagementDetectedFaultThrottleCause ThrottleCause = "ManagementDetectedFault"
	// UnknownThrottleCause The cause of the processor being throttled is not known.
	UnknownThrottleCause ThrottleCause = "Unknown"
	// OEMThrottleCause The cause of the processor being throttled is OEM-specific.
	OEMThrottleCause ThrottleCause = "OEM"
)

// The state of the turbo for this processor.
type TurboState string

const (
	// Turbo is disabled.
	DisabledTurboState TurboState = "Disabled"
	// Turbo is enabled.
	EnabledTurboState TurboState = "Enabled"
)

// AdditionalFirmwareVersions shall contain the additional firmware versions of the processor..
type AdditionalFirmwareVersions struct {
	// (v1.7+) The bootloader version contained in this software, such as U-Boot or UEFI.
	Bootloader string
	// (v1.7+) The kernel version contained in this software.
	// For strict POSIX software, the value shall contain the output of uname -srm.
	// For Microsoft Windows, the value shall contain the output of ver.
	Kernel string
	// (v1.7+)The microcode version contained in this software, such as processor microcode.
	Microcode string
	// (v1.7+) Oem contains vendor-specific data.
	Oem json.RawMessage
	// (v1.8+) The operating system name of this software.
	OSDistribution string
}

// ProcessorEthernetInterface shall contain the definition for an Ethernet interface for a Redfish implementation.
type ProcessorEthernetInterface struct {
	// MaxLanes shall contain the maximum number of lanes supported by this interface.
	MaxLanes int
	// MaxSpeedMbps shall contain the maximum speed supported by this interface.
	MaxSpeedMbps int
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// FPGA shall contain the properties of the FPGA device represented by a
// Processor.
type FPGA struct {
	// ExternalInterfaces shall be an array of objects that describe the
	// external connectivity of the FPGA.
	ExternalInterfaces []ProcessorInterface
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
	// This property has been deprecated in favor of the SystemInterface property in the root of this resource.
	HostInterface ProcessorInterface
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

// ProcessorInterface shall contain information about the system interface, or external connection, to the
// processor.
type ProcessorInterface struct {
	// Ethernet shall contain an object the describes the Ethernet-related information for this interface.
	Ethernet ProcessorEthernetInterface
	// InterfaceType shall contain an enumerated value that describes the type of interface between the system, or
	// external connection, and the processor.
	InterfaceType SystemInterfaceType
	// PCIe shall contain an object the describes the PCIe-related information for this interface.
	PCIe PCIeInterface
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

// ProcessorMemorySummary shall contain properties that describe the summary of all memory that is associated with a
// processor.
type ProcessorMemorySummary struct {
	// ECCModeEnabled shall indicate if memory ECC mode is enabled for this processor. This value shall not affect
	// system memory ECC mode.
	ECCModeEnabled bool
	// Metrics shall contain a link to a resource of type MemoryMetrics that contains the metrics associated with all
	// memory of this processor.
	metrics string
	// TotalCacheSizeMiB shall contain the total size of cache memory of this processor.
	TotalCacheSizeMiB int
	// TotalMemorySizeMiB shall contain the total size of non-cache volatile or non-volatile memory attached to this
	// processor. Examples include DRAMs and NV-DIMMs that are not configured as block storage. This value indicates
	// the size of memory directly attached or with strong affinity to this processor, not the total memory accessible
	// by the processor. This property shall not be present for implementations where all processors have equal memory
	// performance or access characteristics, such as hop count, for all system memory.
	TotalMemorySizeMiB int
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (memorysummary *ProcessorMemorySummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorMemorySummary
	var t struct {
		temp
		Metrics common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorysummary = ProcessorMemorySummary(t.temp)

	// Extract the links to other entities for later
	memorysummary.metrics = t.Metrics.String()

	return nil
}

// Metrics gets the memory metrics for this processor memory summary.
func (memorysummary *ProcessorMemorySummary) Metrics(c common.Client) (*MemoryMetrics, error) {
	if memorysummary.metrics == "" {
		return nil, nil
	}
	return GetMemoryMetrics(c, memorysummary.metrics)
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
	accelerationFunctions []string
	// The additional firmware versions of the processor.
	AdditionalFirmwareVersions AdditionalFirmwareVersions
	// The link to the operating configuration that is applied to this processor.
	appliedOperatingConfig string
	// Assembly shall be a link to a resource
	// of type Assembly.
	assembly string
	// (v1.10+) The base (nominal) clock speed of the processor in MHz.
	BaseSpeedMHz int
	// (v1.9+) The state of the base frequency settings of
	// the operation configuration applied to this processor.
	BaseSpeedPriorityState BaseSpeedPriorityState
	cacheMemory            string
	certificates           []string
	// Description provides a description of this resource.
	Description string
	// (v1.12+) An indication of whether this processor is enabled.
	Enabled            bool
	environmentMetrics string
	// FPGA shall be an object containing
	// properties specific for Processors of type FPGA.
	FPGA FPGA
	// (v1.16+) The processor family, as specified by the combination of
	// the EffectiveFamily and EffectiveModel properties.
	Family string
	// (v1.7+) This property shall contain a string describing the firmware version of
	// the processor as provided by the manufacturer.
	FirmwareVersion string
	// (v1.9+) The list of core identifiers corresponding to the cores that have been configured with
	// the higher clock speed from the operating configuration applied to this processor.
	HighSpeedCoreIDs []int
	// InstructionSet shall contain the string which
	// identifies the instruction set of the processor contained in this
	// socket.
	InstructionSet InstructionSet
	// Location shall contain location information of the
	// associated processor.
	Location common.Location
	// (v1.10+) An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain a string which identifies
	// the manufacturer of the processor.
	Manufacturer string
	// MaxSpeedMHz shall indicate the maximum rated clock
	// speed of the processor in MHz.
	MaxSpeedMHz float32
	// MaxTDPWatts shall be the maximum Thermal
	// Design Power (TDP) in watts.
	MaxTDPWatts int
	// MemorySummary is a summary of all memory associated with this processor.
	MemorySummary MemorySummary
	// Metrics shall be a reference to the Metrics
	// associated with this Processor.
	metrics string
	// (v1.8+) The minimum clock speed of the processor in MHz.
	MinSpeedMHz int
	// Model shall indicate the model information as
	// provided by the manufacturer of this processor.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// (v1.9+) The link to the collection operating configurations
	// that can be applied to this processor.
	operatingConfigs []string
	// (v1.8+) This property shall contain the operating speed of the processor in MHz.
	// The operating speed of the processor may change more frequently
	// than the manager is able to monitor.
	OperatingSpeedMHz int
	// OperatingSpeedRangeMHz is the operating speed control, measured in megahertz units,
	// for this resource. The value of the DataSourceUri property, if present, shall
	// reference a resource of type Control with the ControlType property containing the
	// value of `FrequencyMHz`.
	OperatingSpeedRangeMHz ControlRangeExcerpt
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the processor.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection. It shall contain the interconnect
	// and fabric ports of this processor. It shall not contain ports for GraphicsController resources, USBController
	// resources, or other local adapter-related types of resources.
	ports string
	// PowerState shall contain the power state of the processor. If the PowerState property in the associated Chassis
	// resource contains the value 'Off', this property shall contain 'Off'.
	PowerState PowerState
	// ProcessorArchitecture shall contain the string which
	// identifies the architecture of the processor contained in this Socket.
	ProcessorArchitecture ProcessorArchitecture
	// ProcessorID shall contain identification information for this processor.
	ProcessorID ProcessorID `json:"ProcessorId"`
	// (v1.16+) This property shall contain the zero-based index of the processor,
	// indexed within the next unit of containment.
	ProcessorIndex int
	// ProcessorMemory shall be the memory directly attached or integrated within this Processor.
	ProcessorMemory []ProcessorMemory
	// ProcessorType shall contain the string which
	// identifies the type of processor contained in this Socket.
	ProcessorType ProcessorType
	// (v1.16+) An indication of whether this component can be independently replaced
	// as allowed by the vendor's replacement policy.
	Replaceable bool
	// (v1.7+) The serial number of the processor.
	SerialNumber string
	// Socket shall contain the string which identifies the
	// physical location or socket of the processor.
	Socket string
	// (v1.11+) The spare part number of the processor.
	SparePartNumber string
	// (v1.10+) The clock limit of the processor in MHz.
	SpeedLimitMHz int
	// (v1.10+)	Indicates whether the clock speed of the processor is fixed at the value specified in the SpeedLimitMHz property.
	SpeedLocked bool
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// SubProcessors shall be a link to a
	// collection of type ProcessorCollection.
	subProcessors []string
	// SystemInterface shall contain an object that describes the connectivity between the host system and the
	// processor.
	SystemInterface ProcessorInterface
	// TDPWatts shall be the nominal Thermal
	// Design Power (TDP) in watts.
	TDPWatts int
	// (v1.16+) The causes of the processor being throttled.
	ThrottleCauses []ThrottleCause
	// (v1.16+) An indication of whether the processor is throttled.
	Throttled bool
	// TotalCores shall indicate the total count of
	// independent processor cores contained within this processor.
	TotalCores int
	// TotalEnabledCores shall indicate the total count of
	// enabled independent processor cores contained within this processor.
	TotalEnabledCores int
	// TotalThreads shall indicate the total count of
	// independent execution threads supported by this processor.
	TotalThreads int
	// (v1.9+) The state of the turbo for this processor.
	TurboState TurboState
	// UUID is used to contain a universal unique identifier number for the
	// processor. RFC4122 describes methods that can be used to create the
	// value. The value should be considered to be opaque. Client software
	// should only treat the overall value as a universally unique identifier
	// and should not interpret any sub-fields within the UUID.
	UUID string
	// Version shall contain the hardware version of the processor as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

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
	fabricAdapters []string
	// FabricAdapters@odata.count
	FabricAdaptersCount int
	graphicsController  string
	memory              []string
	// Memory@odata.count
	MemoryCount            int
	networkDeviceFunctions []string
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int
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

	resetTarget           string
	resetToDefaultsTarget string
}

type processorLinks struct {
	Chassis                     common.Link
	ConnectedProcessors         common.Links
	ConnectedProcessorsCount    int `json:"ConnectedProcessors@odata.count"`
	Endpoints                   common.Links
	EndpointsCount              int `json:"Endpoints@odata.count"`
	FabricAdapters              common.Links
	FabricAdaptersCount         int `json:"FabricAdapters@odata.count"`
	GraphicsController          common.Link
	Memory                      common.Links
	MemoryCount                 int `json:"Memory@odata.count"`
	NetworkDeviceFunctions      common.Links
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	PCIeDevice                  common.Link
	PCIeFunctions               common.Links
	PCIeFunctionsCount          int `json:"PCIeFunctions@odata.count"`
}

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (processor *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	type t1 struct {
		temp
		AccelerationFunctions  common.LinksCollection
		AppliedOperatingConfig common.Link
		Assembly               common.Link
		Certificates           common.LinksCollection
		CacheMemory            common.Link
		EnvironmentMetrics     common.Link
		Metrics                common.Link
		OperatingConfigs       common.LinksCollection
		Ports                  common.Link
		SubProcessors          common.LinksCollection
		ProcessorMemory        common.Links
		Links                  processorLinks
		Actions                struct {
			Reset struct {
				Target string
			} `json:"#Processor.Reset"`
			ResetToDefaults struct {
				Target string
			} `json:"#Processor.ResetToDefaults"`
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
			bitSize := 32
			mhz, err := strconv.ParseFloat(t2.MaxSpeedMHz, bitSize)
			if err == nil {
				t.MaxSpeedMHz = float32(mhz)
			}
		}
	}

	*processor = Processor(t.temp)

	// Extract the links to other entities for later
	processor.accelerationFunctions = t.AccelerationFunctions.ToStrings()
	processor.appliedOperatingConfig = t.AppliedOperatingConfig.String()
	processor.assembly = t.Assembly.String()
	processor.cacheMemory = t.CacheMemory.String()
	processor.certificates = t.Certificates.ToStrings()
	processor.environmentMetrics = t.EnvironmentMetrics.String()
	processor.metrics = t.Metrics.String()
	processor.operatingConfigs = t.OperatingConfigs.ToStrings()
	processor.ports = t.Ports.String()
	processor.subProcessors = t.SubProcessors.ToStrings()

	processor.chassis = t.Links.Chassis.String()
	processor.connectedProcessors = t.Links.ConnectedProcessors.ToStrings()
	processor.ConnectedProcessorsCount = t.Links.ConnectedProcessorsCount
	processor.endpoints = t.Links.Endpoints.ToStrings()
	processor.EndpointsCount = t.Links.EndpointsCount
	processor.fabricAdapters = t.Links.FabricAdapters.ToStrings()
	processor.FabricAdaptersCount = t.Links.FabricAdaptersCount
	processor.graphicsController = t.Links.GraphicsController.String()
	processor.memory = t.Links.Memory.ToStrings()
	processor.MemoryCount = t.Links.MemoryCount
	processor.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	processor.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	processor.pcieDevice = t.Links.PCIeDevice.String()
	processor.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	processor.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	processor.resetTarget = t.Actions.Reset.Target
	processor.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target

	// This is a read/write object, so we need to save the raw object data for later
	processor.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (processor *Processor) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Processor)
	original.UnmarshalJSON(processor.rawData)

	readWriteFields := []string{
		"AppliedOperatingConfig",
		"Enabled",
		"LocationIndicatorActive",
		"OperatingSpeedRangeMHz",
		"SpeedLimitMHz",
		"SpeedLocked",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(processor).Elem()

	return processor.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Reset resets the processor.
func (processor *Processor) Reset(resetType ResetType) error {
	t := struct {
		ResetType ResetType
	}{
		ResetType: resetType,
	}
	return processor.Post(processor.resetTarget, t)
}

// ResetToDefaults resets the values of writable properties to factory defaults.
func (processor *Processor) ResetToDefaults() error {
	return processor.Post(processor.resetToDefaultsTarget, nil)
}

// AccelerationFunctions gets acceleration functions associated with this processor.
func (processor *Processor) AcclerationFunctions() ([]*AccelerationFunction, error) {
	return common.GetObjects[AccelerationFunction](processor.GetClient(), processor.accelerationFunctions)
}

// AppliedOperatingConfig gets the operating configuration that is applied to this processor.
func (processor *Processor) AppliedOperatingConfig() (*OperatingConfig, error) {
	if processor.appliedOperatingConfig == "" {
		return nil, nil
	}
	return GetOperatingConfig(processor.GetClient(), processor.appliedOperatingConfig)
}

// Assembly gets the containing assembly for this processor.
func (processor *Processor) Assembly() (*Assembly, error) {
	if processor.assembly == "" {
		return nil, nil
	}
	return GetAssembly(processor.GetClient(), processor.assembly)
}

func (processor *Processor) CacheMemory() ([]*Memory, error) {
	if processor.cacheMemory == "" {
		return nil, nil
	}
	return ListReferencedMemorys(processor.GetClient(), processor.cacheMemory)
}

// Certificates gets the certificates for device identity and attestation.
func (processor *Processor) Certificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](processor.GetClient(), processor.certificates)
}

// EnvironmentMetrics gets the environment metrics for this processor.
func (processor *Processor) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if processor.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(processor.GetClient(), processor.environmentMetrics)
}

// Metrics gets the metrics associated with this processor.
func (processor *Processor) Metrics() (*ProcessorMetrics, error) {
	if processor.metrics == "" {
		return nil, nil
	}
	return GetProcessorMetrics(processor.GetClient(), processor.metrics)
}

// OperatingConfigs gets acceleration functions associated with this processor.
func (processor *Processor) OperatingConfigs() ([]*OperatingConfig, error) {
	return common.GetObjects[OperatingConfig](processor.GetClient(), processor.operatingConfigs)
}

// Ports gets the interconnect and fabric ports of this processor. It shall not
// contain ports for GraphicsController resources, USBController resources, or
// other local adapter-related types of resources.
func (processor *Processor) Ports() ([]*Port, error) {
	return ListReferencedPorts(processor.GetClient(), processor.ports)
}

// SubProcessors gets the sub-processors associated with this processor, such as
// cores or threads, that are part of a processor.
func (processor *Processor) SubProcessors() ([]*Processor, error) {
	return common.GetObjects[Processor](processor.GetClient(), processor.subProcessors)
}

// Chassis gets the physical container associated with this processor.
func (processor *Processor) Chassis() (*Chassis, error) {
	if processor.chassis == "" {
		return nil, nil
	}
	return GetChassis(processor.GetClient(), processor.chassis)
}

// ConnectedProcessors gets the processors that are directly connected to this processor.
func (processor *Processor) ConnectedProcessors() ([]*Processor, error) {
	return common.GetObjects[Processor](processor.GetClient(), processor.connectedProcessors)
}

// Endpoints gets the endpoints associated with this processor.
func (processor *Processor) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](processor.GetClient(), processor.endpoints)
}

// FabricAdapters gets the fabric adapters that present this processor to a fabric.
func (processor *Processor) FabricAdapters() ([]*FabricAdapter, error) {
	return common.GetObjects[FabricAdapter](processor.GetClient(), processor.fabricAdapters)
}

// GraphicsController gets a graphics controller associated with this processor.
func (processor *Processor) GraphicsController() (*GraphicsController, error) {
	if processor.graphicsController == "" {
		return nil, nil
	}
	return GetGraphicsController(processor.GetClient(), processor.graphicsController)
}

// Memory gets the memory objects that are associated with this processor.
func (processor *Processor) Memory() ([]*Memory, error) {
	return common.GetObjects[Memory](processor.GetClient(), processor.memory)
}

// NetworkDeviceFunctions gets the memory objects that are associated with this processor.
func (processor *Processor) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return common.GetObjects[NetworkDeviceFunction](processor.GetClient(), processor.networkDeviceFunctions)
}

// PCIeDevice gets the PCIe device associated with this processor.
func (processor *Processor) PCIeDevice() (*PCIeDevice, error) {
	if processor.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(processor.GetClient(), processor.pcieDevice)
}

// PCIeFunctions gets the PCIeFunctions associated with this processor.
func (processor *Processor) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](processor.GetClient(), processor.pcieFunctions)
}

// GetProcessor will get a Processor instance from the system
func GetProcessor(c common.Client, uri string) (*Processor, error) {
	return common.GetObject[Processor](c, uri)
}

// ListReferencedProcessors gets the collection of Processor from a provided reference.
func ListReferencedProcessors(c common.Client, link string) ([]*Processor, error) {
	return common.GetCollectionObjects[Processor](c, link)
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
	// (v1.10+) The Protected Processor Identification Number (PPIN) for this processor.
	ProtectedIdentificationNumber string
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
