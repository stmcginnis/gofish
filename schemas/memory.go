//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Memory.v1_23_0.json
// 2025.4 - #Memory.v1_23_0.Memory

package schemas

import (
	"encoding/json"
)

type BaseModuleType string

const (
	// RDIMMBaseModuleType Registered DIMM.
	RDIMMBaseModuleType BaseModuleType = "RDIMM"
	// UDIMMBaseModuleType UDIMM.
	UDIMMBaseModuleType BaseModuleType = "UDIMM"
	// SODIMMBaseModuleType SO_DIMM.
	SODIMMBaseModuleType BaseModuleType = "SO_DIMM"
	// LRDIMMBaseModuleType Load Reduced.
	LRDIMMBaseModuleType BaseModuleType = "LRDIMM"
	// MiniRDIMMBaseModuleType Mini_RDIMM.
	MiniRDIMMBaseModuleType BaseModuleType = "Mini_RDIMM"
	// MiniUDIMMBaseModuleType Mini_UDIMM.
	MiniUDIMMBaseModuleType BaseModuleType = "Mini_UDIMM"
	// SORDIMM72bBaseModuleType SO_RDIMM_72b.
	SORDIMM72bBaseModuleType BaseModuleType = "SO_RDIMM_72b"
	// SOUDIMM72bBaseModuleType SO_UDIMM_72b.
	SOUDIMM72bBaseModuleType BaseModuleType = "SO_UDIMM_72b"
	// SODIMM16bBaseModuleType SO_DIMM_16b.
	SODIMM16bBaseModuleType BaseModuleType = "SO_DIMM_16b"
	// SODIMM32bBaseModuleType SO_DIMM_32b.
	SODIMM32bBaseModuleType BaseModuleType = "SO_DIMM_32b"
	// DieBaseModuleType is a die within a package.
	DieBaseModuleType BaseModuleType = "Die"
	// CAMMBaseModuleType Compression Attached Memory Module.
	CAMMBaseModuleType BaseModuleType = "CAMM"
)

type ErrorCorrection string

const (
	// NoECCErrorCorrection No ECC available.
	NoECCErrorCorrection ErrorCorrection = "NoECC"
	// SingleBitECCErrorCorrection Single bit data errors can be corrected by ECC.
	SingleBitECCErrorCorrection ErrorCorrection = "SingleBitECC"
	// MultiBitECCErrorCorrection Multibit data errors can be corrected by ECC.
	MultiBitECCErrorCorrection ErrorCorrection = "MultiBitECC"
	// AddressParityErrorCorrection Address parity errors can be corrected.
	AddressParityErrorCorrection ErrorCorrection = "AddressParity"
)

type MemoryClassification string

const (
	// VolatileMemoryClassification Volatile memory.
	VolatileMemoryClassification MemoryClassification = "Volatile"
	// ByteAccessiblePersistentMemoryClassification Byte-accessible persistent
	// memory.
	ByteAccessiblePersistentMemoryClassification MemoryClassification = "ByteAccessiblePersistent"
	// BlockMemoryClassification Block-accessible memory.
	BlockMemoryClassification MemoryClassification = "Block"
)

type MemoryDeviceType string

const (
	// DDRMemoryDeviceType DDR.
	DDRMemoryDeviceType MemoryDeviceType = "DDR"
	// DDR2MemoryDeviceType DDR2.
	DDR2MemoryDeviceType MemoryDeviceType = "DDR2"
	// DDR3MemoryDeviceType DDR3.
	DDR3MemoryDeviceType MemoryDeviceType = "DDR3"
	// DDR4MemoryDeviceType DDR4.
	DDR4MemoryDeviceType MemoryDeviceType = "DDR4"
	// DDR4SDRAMMemoryDeviceType DDR4 SDRAM.
	DDR4SDRAMMemoryDeviceType MemoryDeviceType = "DDR4_SDRAM"
	// DDR4ESDRAMMemoryDeviceType DDR4E SDRAM.
	DDR4ESDRAMMemoryDeviceType MemoryDeviceType = "DDR4E_SDRAM"
	// LPDDR4SDRAMMemoryDeviceType LPDDR4 SDRAM.
	LPDDR4SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR4_SDRAM"
	// DDR3SDRAMMemoryDeviceType DDR3 SDRAM.
	DDR3SDRAMMemoryDeviceType MemoryDeviceType = "DDR3_SDRAM"
	// LPDDR3SDRAMMemoryDeviceType LPDDR3 SDRAM.
	LPDDR3SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR3_SDRAM"
	// DDR2SDRAMMemoryDeviceType DDR2 SDRAM.
	DDR2SDRAMMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM"
	// DDR2SDRAMFBDIMMMemoryDeviceType DDR2 SDRAM FB_DIMM.
	DDR2SDRAMFBDIMMMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM_FB_DIMM"
	// DDR2SDRAMFBDIMMPROBEMemoryDeviceType DDR2 SDRAM FB_DIMM PROBE.
	DDR2SDRAMFBDIMMPROBEMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM_FB_DIMM_PROBE"
	// DDRSGRAMMemoryDeviceType DDR SGRAM.
	DDRSGRAMMemoryDeviceType MemoryDeviceType = "DDR_SGRAM"
	// DDRSDRAMMemoryDeviceType DDR SDRAM.
	DDRSDRAMMemoryDeviceType MemoryDeviceType = "DDR_SDRAM"
	// ROMMemoryDeviceType ROM.
	ROMMemoryDeviceType MemoryDeviceType = "ROM"
	// SDRAMMemoryDeviceType SDRAM.
	SDRAMMemoryDeviceType MemoryDeviceType = "SDRAM"
	// EDOMemoryDeviceType EDO.
	EDOMemoryDeviceType MemoryDeviceType = "EDO"
	// FastPageModeMemoryDeviceType Fast Page Mode.
	FastPageModeMemoryDeviceType MemoryDeviceType = "FastPageMode"
	// PipelinedNibbleMemoryDeviceType Pipelined Nibble.
	PipelinedNibbleMemoryDeviceType MemoryDeviceType = "PipelinedNibble"
	// LogicalMemoryDeviceType Logical device, such as when the memory is
	// fabric-attached.
	LogicalMemoryDeviceType MemoryDeviceType = "Logical"
	// HBMMemoryDeviceType High Bandwidth Memory.
	HBMMemoryDeviceType MemoryDeviceType = "HBM"
	// HBM2MemoryDeviceType The second generation of High Bandwidth Memory.
	HBM2MemoryDeviceType MemoryDeviceType = "HBM2"
	// HBM2EMemoryDeviceType is an updated version of the second generation of High
	// Bandwidth Memory.
	HBM2EMemoryDeviceType MemoryDeviceType = "HBM2E"
	// HBM3MemoryDeviceType The third generation of High Bandwidth Memory.
	HBM3MemoryDeviceType MemoryDeviceType = "HBM3"
	// GDDRMemoryDeviceType Synchronous graphics random-access memory.
	GDDRMemoryDeviceType MemoryDeviceType = "GDDR"
	// GDDR2MemoryDeviceType Double data rate type two synchronous graphics
	// random-access memory.
	GDDR2MemoryDeviceType MemoryDeviceType = "GDDR2"
	// GDDR3MemoryDeviceType Double data rate type three synchronous graphics
	// random-access memory.
	GDDR3MemoryDeviceType MemoryDeviceType = "GDDR3"
	// GDDR4MemoryDeviceType Double data rate type four synchronous graphics
	// random-access memory.
	GDDR4MemoryDeviceType MemoryDeviceType = "GDDR4"
	// GDDR5MemoryDeviceType Double data rate type five synchronous graphics
	// random-access memory.
	GDDR5MemoryDeviceType MemoryDeviceType = "GDDR5"
	// GDDR5XMemoryDeviceType Double data rate type five X synchronous graphics
	// random-access memory.
	GDDR5XMemoryDeviceType MemoryDeviceType = "GDDR5X"
	// GDDR6MemoryDeviceType Double data rate type six synchronous graphics
	// random-access memory.
	GDDR6MemoryDeviceType MemoryDeviceType = "GDDR6"
	// GDDR7MemoryDeviceType Double data rate type seven synchronous graphics
	// random-access memory.
	GDDR7MemoryDeviceType MemoryDeviceType = "GDDR7"
	// DDR5MemoryDeviceType Double data rate type five synchronous dynamic
	// random-access memory.
	DDR5MemoryDeviceType MemoryDeviceType = "DDR5"
	// OEMMemoryDeviceType OEM-defined.
	OEMMemoryDeviceType MemoryDeviceType = "OEM"
	// LPDDR5SDRAMMemoryDeviceType LPDDR5 SDRAM.
	LPDDR5SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR5_SDRAM"
	// DDR5MRDIMMMemoryDeviceType DDR5 MRDIMM.
	DDR5MRDIMMMemoryDeviceType MemoryDeviceType = "DDR5_MRDIMM"
)

type MemoryMedia string

const (
	// DRAMMemoryMedia DRAM media.
	DRAMMemoryMedia MemoryMedia = "DRAM"
	// NANDMemoryMedia NAND media.
	NANDMemoryMedia MemoryMedia = "NAND"
	// Intel3DXPointMemoryMedia Intel 3D XPoint media.
	Intel3DXPointMemoryMedia MemoryMedia = "Intel3DXPoint"
	// ProprietaryMemoryMedia Proprietary media.
	ProprietaryMemoryMedia MemoryMedia = "Proprietary"
)

type MemoryType string

const (
	// DRAMMemoryType shall represent a volatile DRAM memory device.
	DRAMMemoryType MemoryType = "DRAM"
	// NVDIMMNMemoryType shall represent an NVDIMM_N memory device as defined by
	// JEDEC.
	NVDIMMNMemoryType MemoryType = "NVDIMM_N"
	// NVDIMMFMemoryType shall represent an NVDIMM_F memory device as defined by
	// JEDEC.
	NVDIMMFMemoryType MemoryType = "NVDIMM_F"
	// NVDIMMPMemoryType shall represent an NVDIMM_P memory device as defined by
	// JEDEC.
	NVDIMMPMemoryType MemoryType = "NVDIMM_P"
	// IntelOptaneMemoryType shall represent an Intel Optane Persistent Memory
	// Module.
	IntelOptaneMemoryType MemoryType = "IntelOptane"
	// CacheMemoryType shall represent cache memory.
	CacheMemoryType MemoryType = "Cache"
)

type OperatingMemoryModes string

const (
	// VolatileOperatingMemoryModes Volatile memory.
	VolatileOperatingMemoryModes OperatingMemoryModes = "Volatile"
	// PMEMOperatingMemoryModes Persistent memory, byte-accessible through system
	// address space.
	PMEMOperatingMemoryModes OperatingMemoryModes = "PMEM"
	// BlockOperatingMemoryModes Block-accessible system memory.
	BlockOperatingMemoryModes OperatingMemoryModes = "Block"
)

type SecurityStates string

const (
	// EnabledSecurityStates Secure mode is enabled and access to the data is
	// allowed.
	EnabledSecurityStates SecurityStates = "Enabled"
	// DisabledSecurityStates Secure mode is disabled.
	DisabledSecurityStates SecurityStates = "Disabled"
	// UnlockedSecurityStates Secure mode is enabled and access to the data is
	// unlocked.
	UnlockedSecurityStates SecurityStates = "Unlocked"
	// LockedSecurityStates Secure mode is enabled and access to the data is
	// locked.
	LockedSecurityStates SecurityStates = "Locked"
	// FrozenSecurityStates Secure state is frozen and cannot be modified until
	// reset.
	FrozenSecurityStates SecurityStates = "Frozen"
	// PassphraselimitSecurityStates Number of attempts to unlock the memory
	// exceeded limit.
	PassphraselimitSecurityStates SecurityStates = "Passphraselimit"
)

// Memory shall represent a memory device in a Redfish implementation. It may
// also represent a location, such as a slot, socket, or bay, where a unit may
// be installed, but the 'State' property within the 'Status' property contains
// 'Absent'.
type Memory struct {
	Entity
	// AllocationAlignmentMiB shall contain the alignment boundary on which memory
	// regions are allocated, measured in MiB.
	//
	// Version added: v1.2.0
	AllocationAlignmentMiB *int `json:",omitempty"`
	// AllocationIncrementMiB shall contain the allocation increment for regions,
	// measured in MiB.
	//
	// Version added: v1.2.0
	AllocationIncrementMiB *int `json:",omitempty"`
	// AllowedSpeedsMHz shall contain the speeds supported by this memory device.
	AllowedSpeedsMHz []int
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.4.0
	assembly string
	// BaseModuleType shall contain the base module type of the memory device.
	BaseModuleType BaseModuleType
	// BusWidthBits shall contain the bus width, in bits.
	BusWidthBits *int `json:",omitempty"`
	// CXL shall contain CXL-specific properties for this memory device.
	//
	// Version added: v1.17.0
	CXL MemoryCXL
	// CacheLevel shall contain the level of the cache memory. This property shall
	// only be present if the 'MemoryType' contains the value 'Cache'.
	//
	// Version added: v1.20.0
	CacheLevel int
	// CacheSizeMiB shall contain the total size of the cache portion memory in
	// MiB. If the 'MemoryType' property contains the value 'Cache', the
	// 'CacheSizeMiB' property shall be absent and the 'CapacityMiB' property shall
	// be present.
	//
	// Version added: v1.4.0
	CacheSizeMiB *int `json:",omitempty"`
	// CapacityMiB shall contain the memory capacity in MiB.
	CapacityMiB *int `json:",omitempty"`
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.11.0
	certificates string
	// ConfigurationLocked shall indicate whether the configuration of this memory
	// device is locked and cannot be altered.
	//
	// Version added: v1.7.0
	ConfigurationLocked bool
	// DataWidthBits shall contain the data width in bits.
	DataWidthBits *int `json:",omitempty"`
	// DeviceID shall contain the device ID of the memory device.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of 'ModuleProductID'.
	DeviceID string
	// DeviceLocator shall contain the location of the memory device in the
	// platform, typically marked in the silk screen.
	//
	// Deprecated: v1.9.0
	// This property has been deprecated in favor of the 'ServiceLabel' property
	// within 'Location'.
	DeviceLocator string
	// Enabled shall indicate if this memory is enabled.
	//
	// Version added: v1.12.0
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this memory.
	//
	// Version added: v1.11.0
	environmentMetrics string
	// ErrorCorrection shall contain the error correction scheme supported for this
	// memory device.
	ErrorCorrection ErrorCorrection
	// FirmwareAPIVersion shall contain the version of API supported by the
	// firmware.
	FirmwareAPIVersion string
	// FirmwareRevision shall contain the revision of firmware on the memory
	// controller.
	FirmwareRevision string
	// FunctionClasses shall contain the function classes by the memory device.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of 'OperatingMemoryModes' at the
	// root of the resource, or 'MemoryClassification' found within 'RegionSet'.
	FunctionClasses []string
	// HealthData shall contain the health data of this memory device.
	//
	// Version added: v1.17.0
	HealthData MemoryHealthData
	// IsRankSpareEnabled shall indicate whether rank spare is enabled for this
	// memory device.
	IsRankSpareEnabled bool
	// IsSpareDeviceEnabled shall indicate whether the spare device is enabled.
	IsSpareDeviceEnabled bool
	// Location shall contain the location information of the associated memory
	// device.
	//
	// Version added: v1.4.0
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.10.0
	LocationIndicatorActive bool
	// Log shall contain a link to a resource of type 'LogService'.
	//
	// Version added: v1.13.0
	log string
	// LogicalSizeMiB shall contain the total size of the logical memory in MiB.
	//
	// Version added: v1.4.0
	LogicalSizeMiB *int `json:",omitempty"`
	// Manufacturer shall contain the manufacturer of the memory device.
	Manufacturer string
	// MaxTDPMilliWatts shall contain an array of maximum power budgets supported
	// by the memory device in milliwatt units.
	MaxTDPMilliWatts []int
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.11.0
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// MemoryDeviceType shall contain the Memory Device Type as defined by SMBIOS.
	MemoryDeviceType MemoryDeviceType
	// MemoryLocation shall contain properties that describe the memory connection
	// information to sockets and memory controllers.
	MemoryLocation MemoryLocation
	// MemoryMedia shall contain the media types of this memory device.
	MemoryMedia []MemoryMedia
	// MemorySubsystemControllerManufacturerID shall contain the two byte
	// manufacturer ID of the memory subsystem controller of this memory device as
	// defined by JEDEC in JEP-106.
	//
	// Version added: v1.3.0
	MemorySubsystemControllerManufacturerID string
	// MemorySubsystemControllerProductID shall contain the two byte product ID of
	// the memory subsystem controller of this memory device as defined by the
	// manufacturer.
	//
	// Version added: v1.3.0
	MemorySubsystemControllerProductID string
	// MemoryType shall contain the type of memory device that this resource
	// represents.
	MemoryType MemoryType
	// Metrics The link to the metrics associated with this memory device.
	metrics string
	// Model shall indicate the model information as provided by the manufacturer
	// of this memory.
	//
	// Version added: v1.11.0
	Model string
	// ModuleManufacturerID shall contain the two byte manufacturer ID of this
	// memory device as defined by JEDEC in JEP-106.
	//
	// Version added: v1.3.0
	ModuleManufacturerID string
	// ModuleProductID shall contain the two byte product ID of this memory device
	// as defined by the manufacturer.
	//
	// Version added: v1.3.0
	ModuleProductID string
	// NonVolatileSizeLimitMiB shall contain the total non-volatile memory capacity
	// in mebibytes (MiB).
	//
	// Version added: v1.17.0
	NonVolatileSizeLimitMiB int
	// NonVolatileSizeMiB shall contain the total size of the non-volatile portion
	// memory in MiB.
	//
	// Version added: v1.4.0
	NonVolatileSizeMiB *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingMemoryModes shall contain the memory modes supported by the memory
	// device.
	OperatingMemoryModes []OperatingMemoryModes
	// OperatingSpeedMhz shall contain the operating speed of the memory device in
	// MHz or MT/s (mega-transfers per second) as reported by the memory device.
	// Memory devices that operate at their bus speed shall report the operating
	// speed in MHz (bus speed), while memory devices that transfer data faster
	// than their bus speed, such as DDR memory, shall report the operating speed
	// in MT/s (mega-transfers/second). The reported value shall match the
	// conventionally reported values for the technology used by the memory device.
	OperatingSpeedMhz *int `json:",omitempty"`
	// OperatingSpeedRangeMHz shall contain the operating speed control, in
	// megahertz units, for this resource. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Control' with the
	// 'ControlType' property containing the value of 'FrequencyMHz'.
	//
	// Version added: v1.13.0
	OperatingSpeedRangeMHz ControlRangeExcerpt
	// PartNumber shall indicate the part number as provided by the manufacturer of
	// this memory device.
	PartNumber string
	// PersistentRegionNumberLimit shall contain the total number of persistent
	// regions this memory device can support.
	//
	// Version added: v1.2.0
	PersistentRegionNumberLimit *int `json:",omitempty"`
	// PersistentRegionSizeLimitMiB shall contain the total size of persistent
	// regions in MiB.
	PersistentRegionSizeLimitMiB *int `json:",omitempty"`
	// PersistentRegionSizeMaxMiB shall contain the maximum size of a single
	// persistent regions in MiB.
	//
	// Version added: v1.2.0
	PersistentRegionSizeMaxMiB *int `json:",omitempty"`
	// PoisonListMaxMediaErrorRecords shall contain the maximum number of media
	// error records this device can track in its poison list.
	//
	// Version added: v1.17.0
	PoisonListMaxMediaErrorRecords int
	// PowerManagementICManufacturerID shall contain the two byte manufacturer ID
	// of the Power Management Integrated Controller on this memory device as
	// defined by JEDEC in JESD301.
	//
	// Version added: v1.20.0
	PowerManagementICManufacturerID string
	// PowerManagementICRevisionID shall contain the two byte revision ID of the
	// Power Management Integrated Controller on this memory device as defined by
	// JEDEC in JESD301.
	//
	// Version added: v1.20.0
	PowerManagementICRevisionID string
	// PowerManagementPolicy shall contain properties that describe the power
	// management policy for this resource.
	PowerManagementPolicy PowerManagementPolicy
	// ProductionDate shall contain the date of production or manufacture for this
	// memory device.
	//
	// Version added: v1.23.0
	ProductionDate string
	// RankCount shall contain the number of ranks available in the memory device.
	// The ranks could be used for spare or interleave.
	RankCount *int `json:",omitempty"`
	// Regions shall contain the memory region information within the memory
	// device.
	Regions []RegionSet
	// SecurityCapabilities shall contain properties that describe the security
	// capabilities of the memory device.
	SecurityCapabilities SecurityCapabilities
	// SecurityState shall contain the current security state of this memory
	// device.
	//
	// Version added: v1.7.0
	SecurityState SecurityStates
	// SecurityStates shall contain the security states of this memory device.
	//
	// Version added: v1.17.0
	SecurityStates SecurityStateInfo
	// SerialNumber shall indicate the serial number as provided by the
	// manufacturer of this memory device.
	SerialNumber string
	// SpareDeviceCount shall contain the number of unused spare devices available
	// in the memory device. If the memory device fails, the spare devices could be
	// used.
	SpareDeviceCount *int `json:",omitempty"`
	// SparePartNumber shall contain the spare part number of the memory.
	//
	// Version added: v1.11.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.1.0
	Status Status
	// SubsystemDeviceID shall contain the subsystem device ID of the memory
	// device.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of
	// 'MemorySubsystemControllerProductID'.
	SubsystemDeviceID string
	// SubsystemVendorID shall contain the subsystem vendor ID of the memory
	// device.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of
	// 'MemorySubsystemControllerManufacturerID'.
	SubsystemVendorID string
	// VendorID shall contain the vendor ID of the memory device.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of 'ModuleManufacturerID'.
	VendorID string
	// VolatileRegionNumberLimit shall contain the total number of volatile regions
	// this memory device can support.
	//
	// Version added: v1.2.0
	VolatileRegionNumberLimit *int `json:",omitempty"`
	// VolatileRegionSizeLimitMiB shall contain the total size of volatile regions
	// in MiB.
	VolatileRegionSizeLimitMiB *int `json:",omitempty"`
	// VolatileRegionSizeMaxMiB shall contain the maximum size of a single volatile
	// regions in MiB.
	//
	// Version added: v1.2.0
	VolatileRegionSizeMaxMiB *int `json:",omitempty"`
	// VolatileSizeLimitMiB shall contain the total volatile memory capacity in
	// mebibytes (MiB).
	//
	// Version added: v1.17.0
	VolatileSizeLimitMiB int
	// VolatileSizeMiB shall contain the total size of the volatile portion memory
	// in MiB.
	//
	// Version added: v1.4.0
	VolatileSizeMiB *int `json:",omitempty"`
	// disableMasterPassphraseTarget is the URL to send DisableMasterPassphrase requests.
	disableMasterPassphraseTarget string
	// disablePassphraseTarget is the URL to send DisablePassphrase requests.
	disablePassphraseTarget string
	// freezeSecurityStateTarget is the URL to send FreezeSecurityState requests.
	freezeSecurityStateTarget string
	// injectPersistentPoisonTarget is the URL to send InjectPersistentPoison requests.
	injectPersistentPoisonTarget string
	// overwriteUnitTarget is the URL to send OverwriteUnit requests.
	overwriteUnitTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// scanMediaTarget is the URL to send ScanMedia requests.
	scanMediaTarget string
	// secureEraseUnitTarget is the URL to send SecureEraseUnit requests.
	secureEraseUnitTarget string
	// setMasterPassphraseTarget is the URL to send SetMasterPassphrase requests.
	setMasterPassphraseTarget string
	// setPassphraseTarget is the URL to send SetPassphrase requests.
	setPassphraseTarget string
	// unlockUnitTarget is the URL to send UnlockUnit requests.
	unlockUnitTarget string
	// batteries are the URIs for Batteries.
	batteries []string
	// chassis is the URI for Chassis.
	chassis string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// memoryMediaSources are the URIs for MemoryMediaSources.
	memoryMediaSources []string
	// memoryRegionMediaSources are the URIs for MemoryRegionMediaSources.
	memoryRegionMediaSources []string
	// processors are the URIs for Processors.
	processors []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Memory object from the raw JSON.
func (m *Memory) UnmarshalJSON(b []byte) error {
	type temp Memory
	type mActions struct {
		DisableMasterPassphrase ActionTarget `json:"#Memory.DisableMasterPassphrase"`
		DisablePassphrase       ActionTarget `json:"#Memory.DisablePassphrase"`
		FreezeSecurityState     ActionTarget `json:"#Memory.FreezeSecurityState"`
		InjectPersistentPoison  ActionTarget `json:"#Memory.InjectPersistentPoison"`
		OverwriteUnit           ActionTarget `json:"#Memory.OverwriteUnit"`
		Reset                   ActionTarget `json:"#Memory.Reset"`
		ResetToDefaults         ActionTarget `json:"#Memory.ResetToDefaults"`
		ScanMedia               ActionTarget `json:"#Memory.ScanMedia"`
		SecureEraseUnit         ActionTarget `json:"#Memory.SecureEraseUnit"`
		SetMasterPassphrase     ActionTarget `json:"#Memory.SetMasterPassphrase"`
		SetPassphrase           ActionTarget `json:"#Memory.SetPassphrase"`
		UnlockUnit              ActionTarget `json:"#Memory.UnlockUnit"`
	}
	type mLinks struct {
		Batteries                Links `json:"Batteries"`
		Chassis                  Link  `json:"Chassis"`
		Endpoints                Links `json:"Endpoints"`
		MemoryMediaSources       Links `json:"MemoryMediaSources"`
		MemoryRegionMediaSources Links `json:"MemoryRegionMediaSources"`
		Processors               Links `json:"Processors"`
	}
	var tmp struct {
		temp
		Actions            mActions
		Links              mLinks
		Assembly           Link `json:"Assembly"`
		Certificates       Link `json:"Certificates"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		Log                Link `json:"Log"`
		Metrics            Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Memory(tmp.temp)

	// Extract the links to other entities for later
	m.disableMasterPassphraseTarget = tmp.Actions.DisableMasterPassphrase.Target
	m.disablePassphraseTarget = tmp.Actions.DisablePassphrase.Target
	m.freezeSecurityStateTarget = tmp.Actions.FreezeSecurityState.Target
	m.injectPersistentPoisonTarget = tmp.Actions.InjectPersistentPoison.Target
	m.overwriteUnitTarget = tmp.Actions.OverwriteUnit.Target
	m.resetTarget = tmp.Actions.Reset.Target
	m.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	m.scanMediaTarget = tmp.Actions.ScanMedia.Target
	m.secureEraseUnitTarget = tmp.Actions.SecureEraseUnit.Target
	m.setMasterPassphraseTarget = tmp.Actions.SetMasterPassphrase.Target
	m.setPassphraseTarget = tmp.Actions.SetPassphrase.Target
	m.unlockUnitTarget = tmp.Actions.UnlockUnit.Target
	m.batteries = tmp.Links.Batteries.ToStrings()
	m.chassis = tmp.Links.Chassis.String()
	m.endpoints = tmp.Links.Endpoints.ToStrings()
	m.memoryMediaSources = tmp.Links.MemoryMediaSources.ToStrings()
	m.memoryRegionMediaSources = tmp.Links.MemoryRegionMediaSources.ToStrings()
	m.processors = tmp.Links.Processors.ToStrings()
	m.assembly = tmp.Assembly.String()
	m.certificates = tmp.Certificates.String()
	m.environmentMetrics = tmp.EnvironmentMetrics.String()
	m.log = tmp.Log.String()
	m.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *Memory) Update() error {
	readWriteFields := []string{
		"Enabled",
		"LocationIndicatorActive",
		"NonVolatileSizeLimitMiB",
		"OperatingSpeedRangeMHz",
		"PoisonListMaxMediaErrorRecords",
		"SecurityState",
		"VolatileSizeLimitMiB",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetMemory will get a Memory instance from the service.
func GetMemory(c Client, uri string) (*Memory, error) {
	return GetObject[Memory](c, uri)
}

// ListReferencedMemorys gets the collection of Memory from
// a provided reference.
func ListReferencedMemorys(c Client, link string) ([]*Memory, error) {
	return GetCollectionObjects[Memory](c, link)
}

// This action shall disable the master passphrase on the supplied region
// provided the supplied master passphrase matches that of the region.
// passphrase - This parameter shall contain the master passphrase for the
// specified region.
// regionID - This parameter shall contain the memory region ID to which to
// disable the master passphrase.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) DisableMasterPassphrase(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.disableMasterPassphraseTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall disable the need for passphrases on the supplied region
// provided the supplied passphrase matches that of the region.
// passphrase - This property shall contain the passphrase used in this action.
// regionID - This property shall contain the memory region ID to which to
// apply this action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) DisablePassphrase(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.disablePassphraseTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall freeze the security state of the memory device.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) FreezeSecurityState() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(m.client,
		m.freezeSecurityStateTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall inject poison to a specific persistent memory address in
// the memory device.
// physicalAddress - This parameter shall contain the device persistent
// physical address in which to perform a poison injection as a hex-encoded
// string.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) InjectPersistentPoison(physicalAddress string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PhysicalAddress"] = physicalAddress
	resp, taskInfo, err := PostWithTask(m.client,
		m.injectPersistentPoisonTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall securely erase the supplied region provided the supplied
// passphrase matches that of the given region using the NIST SP800-88 Purge:
// Overwrite. Use the 'SecureEraseUnit' method to perform NIST SP800-88 Purge:
// Cryptographic Erase.
// passphrase - This property shall contain the passphrase used in this action.
// regionID - This property shall contain the memory region ID to which to
// apply this action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) OverwriteUnit(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.overwriteUnitTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset this memory device.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(m.client,
		m.resetTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) ResetToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(m.client,
		m.resetToDefaultsTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall scan the media of the memory device.
// length - This parameter shall contain the length of the target region to
// scan in bytes from the PhysicalAddress parameter.
// noEventLog - This parameter shall indicate whether events related to the
// media scan are not logged. If not provided by the client, the value shall be
// assumed to be 'false'.
// physicalAddress - This parameter shall contain the starting device physical
// address to scan as a hex-encoded string.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) ScanMedia(length int, noEventLog bool, physicalAddress string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Length"] = length
	payload["NoEventLog"] = noEventLog
	payload["PhysicalAddress"] = physicalAddress
	resp, taskInfo, err := PostWithTask(m.client,
		m.scanMediaTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall securely erase the supplied region provided the supplied
// passphrase matches that of the given region using the NIST SP800-88 Purge:
// Cryptographic Erase. Use the 'OverwriteUnit' method to perform NIST SP800-88
// Purge: Overwrite.
// passphrase - This property shall contain the passphrase used in this action.
// regionID - This property shall contain the memory region ID to which to
// apply this action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) SecureEraseUnit(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.secureEraseUnitTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the supplied master passphrase to the supplied region.
// passphrase - This parameter shall contain the master passphrase to set for
// the specified region.
// regionID - This parameter shall contain the memory region ID to which to
// apply the master passphrase.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) SetMasterPassphrase(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.setMasterPassphraseTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall apply the supplied passphrase to the supplied region.
// passphrase - This property shall contain the passphrase used in this action.
// regionID - This property shall contain the memory region ID to which to
// apply this action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) SetPassphrase(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.setPassphraseTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall apply the supplied passphrase to the supplied region for
// the purpose of unlocking the given regions.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// regionID - This property shall contain the memory region ID to which to
// apply this action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *Memory) UnlockUnit(passphrase string, regionID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["RegionId"] = regionID
	resp, taskInfo, err := PostWithTask(m.client,
		m.unlockUnitTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Batteries gets the Batteries linked resources.
func (m *Memory) Batteries() ([]*Battery, error) {
	return GetObjects[Battery](m.client, m.batteries)
}

// Chassis gets the Chassis linked resource.
func (m *Memory) Chassis() (*Chassis, error) {
	if m.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](m.client, m.chassis)
}

// Endpoints gets the Endpoints linked resources.
func (m *Memory) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](m.client, m.endpoints)
}

// MemoryMediaSources gets the MemoryMediaSources linked resources.
func (m *Memory) MemoryMediaSources() ([]*MemoryChunks, error) {
	return GetObjects[MemoryChunks](m.client, m.memoryMediaSources)
}

// MemoryRegionMediaSources gets the MemoryRegionMediaSources linked resources.
func (m *Memory) MemoryRegionMediaSources() ([]*MemoryRegion, error) {
	return GetObjects[MemoryRegion](m.client, m.memoryRegionMediaSources)
}

// Processors gets the Processors linked resources.
func (m *Memory) Processors() ([]*Processor, error) {
	return GetObjects[Processor](m.client, m.processors)
}

// Assembly gets the Assembly linked resource.
func (m *Memory) Assembly() (*Assembly, error) {
	if m.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](m.client, m.assembly)
}

// Certificates gets the Certificates collection.
func (m *Memory) Certificates() ([]*Certificate, error) {
	if m.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](m.client, m.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (m *Memory) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if m.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](m.client, m.environmentMetrics)
}

// Log gets the Log linked resource.
func (m *Memory) Log() (*LogService, error) {
	if m.log == "" {
		return nil, nil
	}
	return GetObject[LogService](m.client, m.log)
}

// Metrics gets the Metrics linked resource.
func (m *Memory) Metrics() (*MemoryMetrics, error) {
	if m.metrics == "" {
		return nil, nil
	}
	return GetObject[MemoryMetrics](m.client, m.metrics)
}

// MemoryCXL shall contain CXL-specific properties for a memory device.
type MemoryCXL struct {
	// LabelStorageSizeBytes shall contain the size of the label storage area in
	// bytes of this memory device.
	//
	// Version added: v1.17.0
	LabelStorageSizeBytes int
	// StagedNonVolatileSizeMiB shall indicate the total device non-volatile memory
	// capacity in mebibytes. The value shall be in multiples of 256 mebibytes.
	//
	// Version added: v1.17.0
	StagedNonVolatileSizeMiB int
	// StagedVolatileSizeMiB shall indicate the total device volatile memory
	// capacity in mebibytes staged for next activation. This value shall be in
	// multiples of 256 mebibytes.
	//
	// Version added: v1.17.0
	StagedVolatileSizeMiB int
}

// MemoryHealthData shall contain the health data of a memory device.
type MemoryHealthData struct {
	// PredictedMediaLifeLeftPercent shall contain the current health of the memory
	// device as a percentage, '0' to '100'.
	//
	// Version added: v1.17.0
	//
	// Deprecated: v1.19.0
	// This property has been deprecated in favor of
	// 'PredictedMediaLifeLeftPercent' in the 'MemoryMetrics' resource.
	PredictedMediaLifeLeftPercent *float64 `json:",omitempty"`
}

// MemoryLocation shall contain properties that describe the memory connection
// information to sockets and memory controllers.
type MemoryLocation struct {
	// Channel shall contain the channel number to which the memory device is
	// connected.
	Channel *int `json:",omitempty"`
	// MemoryController shall contain the memory controller number to which the
	// memory device is connected.
	MemoryController *int `json:",omitempty"`
	// Slot shall contain the slot number to which the memory device is connected.
	Slot *int `json:",omitempty"`
	// Socket shall contain the socket number to which the memory device is
	// connected.
	Socket *int `json:",omitempty"`
}

// PowerManagementPolicy shall contain properties that describe the power
// management policy for this resource.
type PowerManagementPolicy struct {
	// AveragePowerBudgetMilliWatts shall contain the average power budget, in
	// milliwatt units.
	AveragePowerBudgetMilliWatts *int `json:",omitempty"`
	// MaxTDPMilliWatts shall contain the maximum TDP in milliwatt units.
	MaxTDPMilliWatts *int `json:",omitempty"`
	// PeakPowerBudgetMilliWatts shall contain the peak power budget, in milliwatt
	// units.
	PeakPowerBudgetMilliWatts *int `json:",omitempty"`
	// PolicyEnabled shall indicate whether the power management policy is enabled.
	PolicyEnabled bool
}

// RegionSet shall describe the memory region information within a memory
// device.
type RegionSet struct {
	// MasterPassphraseEnabled shall indicate whether the master passphrase is
	// enabled for this region.
	//
	// Version added: v1.17.0
	MasterPassphraseEnabled bool
	// MemoryClassification shall contain the classification of memory that the
	// memory region occupies.
	MemoryClassification MemoryClassification
	// OffsetMiB shall contain the offset within the memory that corresponds to the
	// start of this memory region in MiB.
	OffsetMiB *int `json:",omitempty"`
	// PassphraseEnabled shall indicate whether the passphrase is enabled for this
	// region.
	//
	// Version added: v1.5.0
	PassphraseEnabled bool
	// PassphraseState shall indicate whether the state of the passphrase for this
	// region is enabled.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of 'PassphraseEnabled' found
	// within 'RegionSet'.
	PassphraseState bool
	// RegionID shall contain the unique region ID representing a specific region
	// within the memory device.
	RegionID string `json:"RegionId"`
	// SizeMiB shall contain the size of this memory region in MiB.
	SizeMiB *int `json:",omitempty"`
}

// SecurityCapabilities shall contain properties that describe the security
// capabilities of a memory device.
type SecurityCapabilities struct {
	// ConfigurationLockCapable shall indicate whether this memory device supports
	// the locking, or freezing, of the configuration.
	//
	// Version added: v1.7.0
	ConfigurationLockCapable bool
	// DataLockCapable shall indicate whether this memory device supports the
	// locking of data access.
	//
	// Version added: v1.7.0
	DataLockCapable bool
	// MaxPassphraseCount shall contain the maximum number of passphrases supported
	// for this memory device.
	MaxPassphraseCount *int `json:",omitempty"`
	// PassphraseCapable shall indicate whether the memory device is passphrase
	// capable.
	PassphraseCapable bool
	// PassphraseLockLimit shall contain the maximum number of incorrect passphrase
	// access attempts allowed before access to data is locked. If 0, the number of
	// attempts is infinite.
	//
	// Version added: v1.7.0
	PassphraseLockLimit *int `json:",omitempty"`
	// SecurityStates shall contain the security states supported by the memory
	// device.
	//
	// Deprecated: v1.7.0
	// This property has been deprecated in favor of using the individual
	// 'PassphraseCapable', 'DataLockCapable', and 'ConfigurationLockCapable'
	// properties.
	SecurityStates []SecurityStates
}

// SecurityStateInfo shall contain the security states of a memory device.
type SecurityStateInfo struct {
	// MasterPassphraseAttemptCountReached shall indicate whether an incorrect
	// master passphrase attempt count has been reached.
	//
	// Version added: v1.17.0
	MasterPassphraseAttemptCountReached bool
	// UserPassphraseAttemptCountReached shall indicate whether an incorrect user
	// passphrase attempt count has been reached.
	//
	// Version added: v1.17.0
	UserPassphraseAttemptCountReached bool
}
