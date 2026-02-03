//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #Processor.v1_22_0.Processor

package schemas

import (
	"encoding/json"
)

type BaseSpeedPriorityState string

const (
	// EnabledBaseSpeedPriorityState Base speed priority is enabled.
	EnabledBaseSpeedPriorityState BaseSpeedPriorityState = "Enabled"
	// DisabledBaseSpeedPriorityState Base speed priority is disabled.
	DisabledBaseSpeedPriorityState BaseSpeedPriorityState = "Disabled"
)

type FPGAType string

const (
	// IntegratedFPGAType The FPGA device integrated with other processor in the
	// single chip.
	IntegratedFPGAType FPGAType = "Integrated"
	// DiscreteFPGAType The discrete FPGA device.
	DiscreteFPGAType FPGAType = "Discrete"
)

type InstructionSet string

const (
	// x86InstructionSet x86 32-bit.
	X86InstructionSet InstructionSet = "x86"
	// x8664InstructionSet x86 64-bit.
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

type ProcessorArchitecture string

const (
	// x86ProcessorArchitecture x86 or x86-64.
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
	// HBM2EProcessorMemoryType is an updated version of the second generation of
	// High Bandwidth Memory.
	HBM2EProcessorMemoryType ProcessorMemoryType = "HBM2E"
	// HBM3ProcessorMemoryType The third generation of High Bandwidth Memory.
	HBM3ProcessorMemoryType ProcessorMemoryType = "HBM3"
	// SGRAMProcessorMemoryType Synchronous graphics RAM.
	SGRAMProcessorMemoryType ProcessorMemoryType = "SGRAM"
	// GDDRProcessorMemoryType Synchronous graphics random-access memory.
	GDDRProcessorMemoryType ProcessorMemoryType = "GDDR"
	// GDDR2ProcessorMemoryType Double data rate type two synchronous graphics
	// random-access memory.
	GDDR2ProcessorMemoryType ProcessorMemoryType = "GDDR2"
	// GDDR3ProcessorMemoryType Double data rate type three synchronous graphics
	// random-access memory.
	GDDR3ProcessorMemoryType ProcessorMemoryType = "GDDR3"
	// GDDR4ProcessorMemoryType Double data rate type four synchronous graphics
	// random-access memory.
	GDDR4ProcessorMemoryType ProcessorMemoryType = "GDDR4"
	// GDDR5ProcessorMemoryType Double data rate type five synchronous graphics
	// random-access memory.
	GDDR5ProcessorMemoryType ProcessorMemoryType = "GDDR5"
	// GDDR5XProcessorMemoryType Double data rate type five X synchronous graphics
	// random-access memory.
	GDDR5XProcessorMemoryType ProcessorMemoryType = "GDDR5X"
	// GDDR6ProcessorMemoryType Double data rate type six synchronous graphics
	// random-access memory.
	GDDR6ProcessorMemoryType ProcessorMemoryType = "GDDR6"
	// GDDR7ProcessorMemoryType Double data rate type seven synchronous graphics
	// random-access memory.
	GDDR7ProcessorMemoryType ProcessorMemoryType = "GDDR7"
	// DDRProcessorMemoryType Double data rate synchronous dynamic random-access
	// memory.
	DDRProcessorMemoryType ProcessorMemoryType = "DDR"
	// DDR2ProcessorMemoryType Double data rate type two synchronous dynamic
	// random-access memory.
	DDR2ProcessorMemoryType ProcessorMemoryType = "DDR2"
	// DDR3ProcessorMemoryType Double data rate type three synchronous dynamic
	// random-access memory.
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

type ProcessorType string

const (
	// CPUProcessorType is a CPU.
	CPUProcessorType ProcessorType = "CPU"
	// GPUProcessorType is a GPU.
	GPUProcessorType ProcessorType = "GPU"
	// FPGAProcessorType is an FPGA.
	FPGAProcessorType ProcessorType = "FPGA"
	// DSPProcessorType is a DSP.
	DSPProcessorType ProcessorType = "DSP"
	// AcceleratorProcessorType is an accelerator.
	AcceleratorProcessorType ProcessorType = "Accelerator"
	// CoreProcessorType is a core in a processor.
	CoreProcessorType ProcessorType = "Core"
	// ThreadProcessorType is a thread in a processor.
	ThreadProcessorType ProcessorType = "Thread"
	// PartitionProcessorType shall indicate a partition in a processor that is
	// instantiated from a user configuration to carve out resources in a single
	// processor. An example of this is assigning memory to a set of cores in a
	// GPU.
	PartitionProcessorType ProcessorType = "Partition"
	// OEMProcessorType is an OEM-defined processing unit.
	OEMProcessorType ProcessorType = "OEM"
)

type SystemInterfaceType string

const (
	// QPISystemInterfaceType The Intel QuickPath Interconnect.
	QPISystemInterfaceType SystemInterfaceType = "QPI"
	// UPISystemInterfaceType The Intel UltraPath Interconnect.
	UPISystemInterfaceType SystemInterfaceType = "UPI"
	// PCIeSystemInterfaceType is a PCI Express interface.
	PCIeSystemInterfaceType SystemInterfaceType = "PCIe"
	// EthernetSystemInterfaceType is an Ethernet interface.
	EthernetSystemInterfaceType SystemInterfaceType = "Ethernet"
	// AMBASystemInterfaceType The Arm Advanced Microcontroller Bus Architecture
	// interface.
	AMBASystemInterfaceType SystemInterfaceType = "AMBA"
	// CCIXSystemInterfaceType The Cache Coherent Interconnect for Accelerators
	// interface.
	CCIXSystemInterfaceType SystemInterfaceType = "CCIX"
	// CXLSystemInterfaceType The Compute Express Link interface.
	CXLSystemInterfaceType SystemInterfaceType = "CXL"
	// OEMSystemInterfaceType is an OEM-defined interface.
	OEMSystemInterfaceType SystemInterfaceType = "OEM"
)

type ThrottleCause string

const (
	// PowerLimitThrottleCause The cause of the processor being throttled is a
	// power limit.
	PowerLimitThrottleCause ThrottleCause = "PowerLimit"
	// ThermalLimitThrottleCause The cause of the processor being throttled is a
	// thermal limit.
	ThermalLimitThrottleCause ThrottleCause = "ThermalLimit"
	// ClockLimitThrottleCause The cause of the processor being throttled is a
	// clock limit.
	ClockLimitThrottleCause ThrottleCause = "ClockLimit"
	// ManagementDetectedFaultThrottleCause The cause of the processor being
	// throttled is a fault detected by management hardware or firmware.
	ManagementDetectedFaultThrottleCause ThrottleCause = "ManagementDetectedFault"
	// UnknownThrottleCause The cause of the processor being throttled is not
	// known.
	UnknownThrottleCause ThrottleCause = "Unknown"
	// OEMThrottleCause The cause of the processor being throttled is OEM-specific.
	OEMThrottleCause ThrottleCause = "OEM"
)

type TurboState string

const (
	// EnabledTurboState Turbo is enabled.
	EnabledTurboState TurboState = "Enabled"
	// DisabledTurboState Turbo is disabled.
	DisabledTurboState TurboState = "Disabled"
)

// Processor shall represent a single processor that a system contains. A
// processor includes both performance characteristics, clock speed,
// architecture, core count, and so on, and compatibility, such as the CPU ID
// instruction results. It may also represent a location, such as a slot,
// socket, or bay, where a unit may be installed, but the 'State' property
// within the 'Status' property contains 'Absent'.
type Processor struct {
	Entity
	// AccelerationFunctions shall contain a link to a resource collection of type
	// 'AccelerationFunctionCollection'.
	//
	// Version added: v1.4.0
	accelerationFunctions string
	// AdditionalFirmwareVersions shall contain the additional firmware versions of
	// the processor.
	//
	// Version added: v1.15.0
	AdditionalFirmwareVersions AdditionalVersions
	// AppliedOperatingConfig shall contain a link to a resource of type
	// 'OperatingConfig' that specifies the configuration is applied to this
	// processor.
	//
	// Version added: v1.9.0
	appliedOperatingConfig string
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.2.0
	assembly string
	// BaseSpeedMHz shall contain the base (nominal) clock speed of the processor
	// in MHz.
	//
	// Version added: v1.10.0
	BaseSpeedMHz *uint `json:",omitempty"`
	// BaseSpeedPriorityState shall contain the state of the base frequency
	// settings of the operating configuration applied to this processor.
	//
	// Version added: v1.9.0
	BaseSpeedPriorityState BaseSpeedPriorityState
	// CacheMemory shall contain a link to a resource collection of type
	// 'MemoryCollection' that represents the cache memory of this processor.
	//
	// Version added: v1.20.0
	cacheMemory string
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.11.0
	certificates string
	// Enabled shall indicate if this processor is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this
	// processor.
	//
	// Version added: v1.11.0
	environmentMetrics string
	// FPGA shall contain an object containing properties for processors of type
	// 'FPGA'.
	//
	// Version added: v1.4.0
	FPGA FPGA
	// Family shall contain a string that identifies the processor family, as
	// specified by the combination of the 'EffectiveFamily' and 'EffectiveModel'
	// properties.
	//
	// Version added: v1.16.0
	Family string
	// FirmwareVersion shall contain a string describing the firmware version of
	// the processor as provided by the manufacturer.
	//
	// Version added: v1.7.0
	FirmwareVersion string
	// HighSpeedCoreIDs shall contain an array of core identifiers corresponding to
	// the cores that have been configured with the higher clock speed from the
	// operating configuration applied to this processor.
	//
	// Version added: v1.9.0
	HighSpeedCoreIDs []*int
	// InstructionSet shall contain the string that identifies the instruction set
	// of the processor contained in this socket.
	InstructionSet InstructionSet
	// Location shall contain the location information of the associated processor.
	//
	// Version added: v1.2.0
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.10.0
	LocationIndicatorActive bool
	// Manufacturer shall contain a string that identifies the manufacturer of the
	// processor.
	Manufacturer string
	// MaxSpeedMHz shall indicate the maximum rated clock speed of the processor in
	// MHz.
	MaxSpeedMHz *int `json:",omitempty"`
	// MaxTDPWatts shall contain the maximum Thermal Design Power (TDP) in watt
	// units.
	//
	// Version added: v1.4.0
	MaxTDPWatts *int `json:",omitempty"`
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.11.0
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// MemorySummary shall contain properties that describe the summary of all
	// memory that is associated with this processor.
	//
	// Version added: v1.11.0
	MemorySummary ProcessorMemorySummary
	// Metrics shall contain a link to a resource of type 'ProcessorMetrics' that
	// contains the metrics associated with this processor.
	//
	// Version added: v1.4.0
	metrics string
	// MinSpeedMHz shall indicate the minimum rated clock speed of the processor in
	// MHz.
	//
	// Version added: v1.8.0
	MinSpeedMHz *int `json:",omitempty"`
	// Model shall indicate the model information as provided by the manufacturer
	// of this processor.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingConfigs shall contain a link to a resource collection of type
	// 'OperatingConfigCollection'.
	//
	// Version added: v1.9.0
	operatingConfigs string
	// OperatingSpeedMHz shall contain the operating speed of the processor in MHz.
	// The operating speed of the processor may change more frequently than the
	// manager is able to monitor.
	//
	// Version added: v1.8.0
	OperatingSpeedMHz *int `json:",omitempty"`
	// OperatingSpeedRangeMHz shall contain the operating speed control, measured
	// in megahertz units, for this resource. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Control' with the
	// 'ControlType' property containing the value of 'FrequencyMHz'.
	//
	// Version added: v1.13.0
	OperatingSpeedRangeMHz ControlRangeExcerpt
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the processor.
	//
	// Version added: v1.7.0
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'. It shall contain the interconnect and fabric ports of this
	// processor. It shall not contain ports for 'GraphicsController' resources,
	// 'USBController' resources, or other local adapter-related types of
	// resources.
	//
	// Version added: v1.13.0
	ports string
	// PowerState shall contain the power state of the processor. If the
	// 'PowerState' property in the associated 'Chassis' resource contains the
	// value 'Off', this property shall contain 'Off'.
	//
	// Version added: v1.17.0
	PowerState PowerState
	// ProcessorArchitecture shall contain the string that identifies the
	// architecture of the processor contained in this socket.
	ProcessorArchitecture ProcessorArchitecture
	// ProcessorID shall contain identification information for this processor. For
	// additional property requirements, see the corresponding definition in the
	// Redfish Data Model Specification.
	ProcessorID ProcessorID `json:"ProcessorId"`
	// ProcessorIndex shall contain the zero-based index of the processor, indexed
	// within the next unit of containment. The value of this property shall match
	// the ordering in the operating system topology interfaces, with offset
	// adjustments, if needed.
	//
	// Version added: v1.16.0
	ProcessorIndex *int `json:",omitempty"`
	// ProcessorMemory shall contain the memory directly attached or integrated
	// within this processor.
	//
	// Version added: v1.4.0
	ProcessorMemory []ProcessorMemory
	// ProcessorType shall contain the string that identifies the type of processor
	// contained in this socket.
	ProcessorType ProcessorType
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.16.0
	Replaceable bool
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the processor.
	//
	// Version added: v1.7.0
	SerialNumber string
	// Socket shall contain the string that identifies the physical location or
	// socket of the processor.
	Socket string
	// SparePartNumber shall contain the spare part number of the processor.
	//
	// Version added: v1.11.0
	SparePartNumber string
	// SpeedLimitMHz shall contain the clock limit of the processor in MHz. This
	// value shall be within the range of 'MinSpeedMHz' and 'MaxSpeedMHz' as
	// provided by the manufacturer of this processor.
	//
	// Version added: v1.10.0
	SpeedLimitMHz *uint `json:",omitempty"`
	// SpeedLocked shall indicate whether the clock speed of the processor is
	// fixed, where a value 'true' shall indicate that the clock speed is fixed at
	// the value specified in the 'SpeedLimitMHz' property.
	//
	// Version added: v1.10.0
	SpeedLocked bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SubProcessors shall contain a link to a resource collection of type
	// 'ProcessorCollection'.
	//
	// Version added: v1.3.0
	subProcessors string
	// SystemInterface shall contain an object that describes the connectivity
	// between the host system and the processor.
	//
	// Version added: v1.8.0
	SystemInterface ProcessorInterface
	// TDPWatts shall contain the nominal Thermal Design Power (TDP) in watt units.
	//
	// Version added: v1.4.0
	TDPWatts *int `json:",omitempty"`
	// ThrottleCauses shall contain the causes of the processor being throttled. If
	// 'Throttled' contains 'false', this property shall contain an empty array.
	//
	// Version added: v1.16.0
	ThrottleCauses []ThrottleCause
	// Throttled shall indicate whether the processor is throttled.
	//
	// Version added: v1.16.0
	Throttled bool
	// TotalCores shall indicate the total count of independent processor cores,
	// including disabled cores, contained within this processor.
	TotalCores *int `json:",omitempty"`
	// TotalEnabledCores shall indicate the total count of enabled independent
	// processor cores contained within this processor.
	//
	// Version added: v1.5.0
	TotalEnabledCores *int `json:",omitempty"`
	// TotalEnabledThreads shall indicate the total count of enabled independent
	// execution threads contained within this processor.
	//
	// Version added: v1.21.0
	TotalEnabledThreads *int `json:",omitempty"`
	// TotalThreads shall indicate the total count of independent execution
	// threads, including disabled threads, that this processor supports.
	TotalThreads *int `json:",omitempty"`
	// TurboState shall contain the state of turbo for this processor.
	//
	// Version added: v1.9.0
	TurboState TurboState
	// UALink shall contain UALink attributes of the processor.
	//
	// Version added: v1.22.0
	UALink ProcessorUALink
	// UUID shall contain a universally unique identifier number for the processor.
	// RFC4122 describes methods to use to create the value. The value should be
	// considered to be opaque. Client software should only treat the overall value
	// as a universally unique identifier and should not interpret any subfields
	// within the UUID.
	//
	// Version added: v1.4.0
	UUID string
	// Version shall contain the hardware version of the processor as determined by
	// the vendor or supplier.
	//
	// Version added: v1.7.0
	Version string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// chassis is the URI for Chassis.
	chassis string
	// connectedProcessors are the URIs for ConnectedProcessors.
	connectedProcessors []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// fabricAdapters are the URIs for FabricAdapters.
	fabricAdapters []string
	// graphicsController is the URI for GraphicsController.
	graphicsController string
	// memory are the URIs for Memory.
	memory []string
	// networkDeviceFunctions are the URIs for NetworkDeviceFunctions.
	networkDeviceFunctions []string
	// pCIeDevice is the URI for PCIeDevice.
	pCIeDevice string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (p *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	type pActions struct {
		Reset           ActionTarget `json:"#Processor.Reset"`
		ResetToDefaults ActionTarget `json:"#Processor.ResetToDefaults"`
	}
	type pLinks struct {
		Chassis                Link  `json:"Chassis"`
		ConnectedProcessors    Links `json:"ConnectedProcessors"`
		Endpoints              Links `json:"Endpoints"`
		FabricAdapters         Links `json:"FabricAdapters"`
		GraphicsController     Link  `json:"GraphicsController"`
		Memory                 Links `json:"Memory"`
		NetworkDeviceFunctions Links `json:"NetworkDeviceFunctions"`
		PCIeDevice             Link  `json:"PCIeDevice"`
		PCIeFunctions          Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Actions                pActions
		Links                  pLinks
		AccelerationFunctions  Link `json:"AccelerationFunctions"`
		AppliedOperatingConfig Link `json:"AppliedOperatingConfig"`
		Assembly               Link `json:"Assembly"`
		CacheMemory            Link `json:"CacheMemory"`
		Certificates           Link `json:"Certificates"`
		EnvironmentMetrics     Link `json:"EnvironmentMetrics"`
		Metrics                Link `json:"Metrics"`
		OperatingConfigs       Link `json:"OperatingConfigs"`
		Ports                  Link `json:"Ports"`
		SubProcessors          Link `json:"SubProcessors"`

		MaxSpeedMHz any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Processor(tmp.temp)

	// Extract the links to other entities for later
	p.resetTarget = tmp.Actions.Reset.Target
	p.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	p.chassis = tmp.Links.Chassis.String()
	p.connectedProcessors = tmp.Links.ConnectedProcessors.ToStrings()
	p.endpoints = tmp.Links.Endpoints.ToStrings()
	p.fabricAdapters = tmp.Links.FabricAdapters.ToStrings()
	p.graphicsController = tmp.Links.GraphicsController.String()
	p.memory = tmp.Links.Memory.ToStrings()
	p.networkDeviceFunctions = tmp.Links.NetworkDeviceFunctions.ToStrings()
	p.pCIeDevice = tmp.Links.PCIeDevice.String()
	p.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	p.accelerationFunctions = tmp.AccelerationFunctions.String()
	p.appliedOperatingConfig = tmp.AppliedOperatingConfig.String()
	p.assembly = tmp.Assembly.String()
	p.cacheMemory = tmp.CacheMemory.String()
	p.certificates = tmp.Certificates.String()
	p.environmentMetrics = tmp.EnvironmentMetrics.String()
	p.metrics = tmp.Metrics.String()
	p.operatingConfigs = tmp.OperatingConfigs.String()
	p.ports = tmp.Ports.String()
	p.subProcessors = tmp.SubProcessors.String()

	p.MaxSpeedMHz = toInt(tmp.MaxSpeedMHz)

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *Processor) Update() error {
	readWriteFields := []string{
		"AppliedOperatingConfig",
		"Enabled",
		"LocationIndicatorActive",
		"OperatingSpeedRangeMHz",
		"SpeedLimitMHz",
		"SpeedLocked",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetProcessor will get a Processor instance from the service.
func GetProcessor(c Client, uri string) (*Processor, error) {
	return GetObject[Processor](c, uri)
}

// ListReferencedProcessors gets the collection of Processor from
// a provided reference.
func ListReferencedProcessors(c Client, link string) ([]*Processor, error) {
	return GetCollectionObjects[Processor](c, link)
}

// This action shall reset the processor.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Processor) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Processor) ResetToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetToDefaultsTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Chassis gets the Chassis linked resource.
func (p *Processor) Chassis() (*Chassis, error) {
	if p.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](p.client, p.chassis)
}

// ConnectedProcessors gets the ConnectedProcessors linked resources.
func (p *Processor) ConnectedProcessors() ([]*Processor, error) {
	return GetObjects[Processor](p.client, p.connectedProcessors)
}

// Endpoints gets the Endpoints linked resources.
func (p *Processor) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](p.client, p.endpoints)
}

// FabricAdapters gets the FabricAdapters linked resources.
func (p *Processor) FabricAdapters() ([]*FabricAdapter, error) {
	return GetObjects[FabricAdapter](p.client, p.fabricAdapters)
}

// GraphicsController gets the GraphicsController linked resource.
func (p *Processor) GraphicsController() (*GraphicsController, error) {
	if p.graphicsController == "" {
		return nil, nil
	}
	return GetObject[GraphicsController](p.client, p.graphicsController)
}

// Memory gets the Memory linked resources.
func (p *Processor) Memory() ([]*Memory, error) {
	return GetObjects[Memory](p.client, p.memory)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions linked resources.
func (p *Processor) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](p.client, p.networkDeviceFunctions)
}

// PCIeDevice gets the PCIeDevice linked resource.
func (p *Processor) PCIeDevice() (*PCIeDevice, error) {
	if p.pCIeDevice == "" {
		return nil, nil
	}
	return GetObject[PCIeDevice](p.client, p.pCIeDevice)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (p *Processor) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](p.client, p.pCIeFunctions)
}

// AccelerationFunctions gets the AccelerationFunctions collection.
func (p *Processor) AccelerationFunctions() ([]*AccelerationFunction, error) {
	if p.accelerationFunctions == "" {
		return nil, nil
	}
	return GetCollectionObjects[AccelerationFunction](p.client, p.accelerationFunctions)
}

// AppliedOperatingConfig gets the AppliedOperatingConfig linked resource.
func (p *Processor) AppliedOperatingConfig() (*OperatingConfig, error) {
	if p.appliedOperatingConfig == "" {
		return nil, nil
	}
	return GetObject[OperatingConfig](p.client, p.appliedOperatingConfig)
}

// Assembly gets the Assembly linked resource.
func (p *Processor) Assembly() (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](p.client, p.assembly)
}

// CacheMemory gets the CacheMemory collection.
func (p *Processor) CacheMemory() ([]*Memory, error) {
	if p.cacheMemory == "" {
		return nil, nil
	}
	return GetCollectionObjects[Memory](p.client, p.cacheMemory)
}

// Certificates gets the Certificates collection.
func (p *Processor) Certificates() ([]*Certificate, error) {
	if p.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](p.client, p.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (p *Processor) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if p.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](p.client, p.environmentMetrics)
}

// Metrics gets the Metrics linked resource.
func (p *Processor) Metrics() (*ProcessorMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[ProcessorMetrics](p.client, p.metrics)
}

// OperatingConfigs gets the OperatingConfigs collection.
func (p *Processor) OperatingConfigs() ([]*OperatingConfig, error) {
	if p.operatingConfigs == "" {
		return nil, nil
	}
	return GetCollectionObjects[OperatingConfig](p.client, p.operatingConfigs)
}

// Ports gets the Ports collection.
func (p *Processor) Ports() ([]*Port, error) {
	if p.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](p.client, p.ports)
}

// SubProcessors gets the SubProcessors collection.
func (p *Processor) SubProcessors() ([]*Processor, error) {
	if p.subProcessors == "" {
		return nil, nil
	}
	return GetCollectionObjects[Processor](p.client, p.subProcessors)
}

// ProcessorEthernetInterface shall contain the definition for an Ethernet interface for
// a Redfish implementation.
type ProcessorEthernetInterface struct {
	// MaxLanes shall contain the maximum number of lanes supported by this
	// interface.
	//
	// Version added: v1.4.0
	MaxLanes *int `json:",omitempty"`
	// MaxSpeedMbps shall contain the maximum speed supported by this interface.
	//
	// Version added: v1.4.0
	MaxSpeedMbps *int `json:",omitempty"`
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.4.0
	OEM json.RawMessage `json:"Oem"`
}

// FPGA shall contain the properties of the FPGA device represented by a
// processor.
type FPGA struct {
	// ExternalInterfaces shall contain an array of objects that describe the
	// external connectivity of the FPGA.
	//
	// Version added: v1.4.0
	ExternalInterfaces []ProcessorInterface
	// FPGAType shall contain a type of the FPGA device.
	//
	// Version added: v1.4.0
	FPGAType FPGAType `json:"FpgaType"`
	// FirmwareID shall contain a string describing the FPGA firmware identifier.
	//
	// Version added: v1.4.0
	FirmwareID string `json:"FirmwareId"`
	// FirmwareManufacturer shall contain a string describing the FPGA firmware
	// manufacturer.
	//
	// Version added: v1.4.0
	FirmwareManufacturer string
	// FirmwareVersion shall contain a string describing the FPGA firmware version.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.9.0
	// This property has been deprecated in favor of the 'FirmwareVersion' property
	// in the root of this resource.
	FirmwareVersion string
	// HostInterface shall contain an object that describes the connectivity to the
	// host for system software to use.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.8.0
	// This property has been deprecated in favor of the 'SystemInterface' property
	// in the root of this resource.
	HostInterface ProcessorInterface
	// Model shall contain a model of the FPGA device.
	//
	// Version added: v1.4.0
	Model string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.4.0
	OEM json.RawMessage `json:"Oem"`
	// PCIeVirtualFunctions shall contain an integer that describes the number of
	// PCIe Virtual Functions configured within the FPGA.
	//
	// Version added: v1.4.0
	PCIeVirtualFunctions int
	// ProgrammableFromHost shall indicate whether the FPGA firmware can be
	// reprogrammed from the host by using system software. If 'false', system
	// software shall not be able to program the FPGA firmware from the system
	// interface. In either state, a management controller may be able to program
	// the FPGA firmware by using the sideband interface.
	//
	// Version added: v1.4.0
	ProgrammableFromHost bool
	// ReconfigurationSlots shall contain an array of the structures that describe
	// the FPGA reconfiguration slots that the acceleration functions can program.
	//
	// Version added: v1.4.0
	ReconfigurationSlots []FPGAReconfigurationSlot
}

// FPGAReconfigurationSlot shall contain information about the FPGA
// reconfiguration slot.
type FPGAReconfigurationSlot struct {
	// AccelerationFunction shall contain a link to a resource of type
	// 'AccelerationFunction' that represents the code programmed into this
	// reconfiguration slot.
	//
	// Version added: v1.4.0
	accelerationFunction string
	// ProgrammableFromHost shall indicate whether the reconfiguration slot can be
	// reprogrammed from the host by using system software. If 'false', system
	// software shall not be able to program the reconfiguration slot from the
	// system interface. In either state, a management controller may be able to
	// program the reconfiguration slot by using the sideband interface.
	//
	// Version added: v1.4.0
	ProgrammableFromHost bool
	// SlotID shall contain the FPGA reconfiguration slot identifier.
	//
	// Version added: v1.4.0
	SlotID string `json:"SlotId"`
	// UUID shall contain a universally unique identifier number for the
	// reconfiguration slot.
	//
	// Version added: v1.4.0
	UUID string
}

// UnmarshalJSON unmarshals a FPGAReconfigurationSlot object from the raw JSON.
func (f *FPGAReconfigurationSlot) UnmarshalJSON(b []byte) error {
	type temp FPGAReconfigurationSlot
	var tmp struct {
		temp
		AccelerationFunction Link `json:"AccelerationFunction"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FPGAReconfigurationSlot(tmp.temp)

	// Extract the links to other entities for later
	f.accelerationFunction = tmp.AccelerationFunction.String()

	return nil
}

// AccelerationFunction gets the AccelerationFunction linked resource.
func (f *FPGAReconfigurationSlot) AccelerationFunction(client Client) (*AccelerationFunction, error) {
	if f.accelerationFunction == "" {
		return nil, nil
	}
	return GetObject[AccelerationFunction](client, f.accelerationFunction)
}

// ProcessorMemorySummary shall contain properties that describe the summary of all
// memory that is associated with a processor.
type ProcessorMemorySummary struct {
	// ECCModeEnabled shall indicate if memory ECC mode is enabled for this
	// processor. This value shall not affect system memory ECC mode.
	//
	// Version added: v1.13.0
	ECCModeEnabled bool
	// Metrics shall contain a link to a resource of type 'MemoryMetrics' that
	// contains the metrics associated with all memory of this processor.
	//
	// Version added: v1.11.0
	metrics string
	// TotalCacheSizeMiB shall contain the total size of cache memory of this
	// processor.
	//
	// Version added: v1.11.0
	TotalCacheSizeMiB *int `json:",omitempty"`
	// TotalMemorySizeMiB shall contain the total size of non-cache volatile or
	// non-volatile memory attached to this processor. Examples include DRAMs and
	// NV-DIMMs that are not configured as block storage. This value indicates the
	// size of memory directly attached or with strong affinity to this processor,
	// not the total memory accessible by the processor. This property shall not be
	// present for implementations where all processors have equal memory
	// performance or access characteristics, such as hop count, for all system
	// memory.
	//
	// Version added: v1.11.0
	TotalMemorySizeMiB *int `json:",omitempty"`
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (m *ProcessorMemorySummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorMemorySummary
	var tmp struct {
		temp
		Metrics Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ProcessorMemorySummary(tmp.temp)

	// Extract the links to other entities for later
	m.metrics = tmp.Metrics.String()

	return nil
}

// Metrics gets the Metrics linked resource.
func (m *ProcessorMemorySummary) Metrics(client Client) (*MemoryMetrics, error) {
	if m.metrics == "" {
		return nil, nil
	}
	return GetObject[MemoryMetrics](client, m.metrics)
}

// ProcessorID shall contain identification information for a processor.
type ProcessorID struct {
	// EffectiveFamily shall contain the effective family information as provided
	// by the manufacturer of this processor. If this property represents raw
	// register data, as determined by the value of the 'ProcessorArchitecture'
	// property, the service shall encode the value as a hex-encoded string
	// following the regular expression pattern '^0x[0-9A-Fa-f]+$' or a
	// decimal-encoded string following the regular expression pattern '^\d+$'. For
	// additional property requirements, see the corresponding definition in the
	// Redfish Data Model Specification.
	EffectiveFamily string
	// EffectiveModel shall contain the effective model information as provided by
	// the manufacturer of this processor. If this property represents raw register
	// data, as determined by the value of the 'ProcessorArchitecture' property,
	// the service shall encode the value as a hex-encoded string following the
	// regular expression pattern '^0x[0-9A-Fa-f]+$' or a decimal-encoded string
	// following the regular expression pattern '^\d+$'. For additional property
	// requirements, see the corresponding definition in the Redfish Data Model
	// Specification.
	EffectiveModel string
	// IdentificationRegisters shall contain the raw manufacturer-provided
	// processor-specific identification registers of this processor's features.
	// For additional property requirements, see the corresponding definition in
	// the Redfish Data Model Specification.
	IdentificationRegisters string
	// MicrocodeInfo shall contain the microcode information as provided by the
	// manufacturer of this processor. If this property represents raw register
	// data, as determined by the value of the 'ProcessorArchitecture' property,
	// the service shall encode the value as a hex-encoded string following the
	// regular expression pattern '^0x[0-9A-Fa-f]+$' or a decimal-encoded string
	// following the regular expression pattern '^\d+$'. For additional property
	// requirements, see the corresponding definition in the Redfish Data Model
	// Specification.
	MicrocodeInfo string
	// ProtectedIdentificationNumber shall contain the Protected Processor
	// Identification Number (PPIN) for this processor.
	//
	// Version added: v1.10.0
	ProtectedIdentificationNumber string
	// Step shall contain the step or revision information as provided by the
	// manufacturer of this processor. If this property represents raw register
	// data, as determined by the value of the 'ProcessorArchitecture' property,
	// the service shall encode the value as a hex-encoded string following the
	// regular expression pattern '^0x[0-9A-Fa-f]+$' or a decimal-encoded string
	// following the regular expression pattern '^\d+$'. For additional property
	// requirements, see the corresponding definition in the Redfish Data Model
	// Specification.
	Step string
	// VendorID shall contain the vendor identification information as provided by
	// the manufacturer of this processor. If this property represents raw register
	// data, as determined by the value of the 'ProcessorArchitecture' property,
	// the service shall encode the value as a hex-encoded string following the
	// regular expression pattern '^0x[0-9A-Fa-f]+$' or a decimal-encoded string
	// following the regular expression pattern '^\d+$'. For additional property
	// requirements, see the corresponding definition in the Redfish Data Model
	// Specification.
	VendorID string `json:"VendorId"`
}

// ProcessorInterface shall contain information about the system interface, or
// external connection, to the processor.
type ProcessorInterface struct {
	// Ethernet shall contain an object the describes the Ethernet-related
	// information for this interface.
	//
	// Version added: v1.4.0
	Ethernet ProcessorEthernetInterface
	// InterfaceType shall contain an enumerated value that describes the type of
	// interface between the system, or external connection, and the processor.
	//
	// Version added: v1.4.0
	InterfaceType SystemInterfaceType
	// PCIe shall contain an object the describes the PCIe-related information for
	// this interface.
	//
	// Version added: v1.4.0
	PCIe PCIeInterface
}

// ProcessorMemory shall contain information about memory directly attached or
// integrated within a processor.
type ProcessorMemory struct {
	// CapacityMiB shall contain the memory capacity in MiB.
	//
	// Version added: v1.4.0
	CapacityMiB *int `json:",omitempty"`
	// IntegratedMemory shall indicate whether this memory is integrated within the
	// processor. Otherwise, it is discrete memory attached to the processor.
	//
	// Version added: v1.4.0
	IntegratedMemory bool
	// MemoryType shall contain a type of the processor memory type.
	//
	// Version added: v1.4.0
	MemoryType ProcessorMemoryType
	// SpeedMHz shall contain the operating speed of the memory in MHz.
	//
	// Version added: v1.4.0
	SpeedMHz *int `json:",omitempty"`
}

// ProcessorUALink shall contain information about UALink fabric and port management
// within the processor.
type ProcessorUALink struct {
	// AcceleratorID shall contain an integer that is unique within a UALink Pod
	// and that serves as the Source and Destination Accelerator ID on the UALink
	// network.
	//
	// Version added: v1.22.0
	AcceleratorID uint
}
