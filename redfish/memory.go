//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BaseModuleType is the type of module.
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
	// DieBaseModuleType A die within a package.
	DieBaseModuleType BaseModuleType = "Die"
)

// ErrorCorrection is the type of error correction used.
type ErrorCorrection string

const (
	// NoECCErrorCorrection No ECC available.
	NoECCErrorCorrection ErrorCorrection = "NoECC"
	// SingleBitECCErrorCorrection Single bit Data error can be corrected by
	// ECC.
	SingleBitECCErrorCorrection ErrorCorrection = "SingleBitECC"
	// MultiBitECCErrorCorrection Multi-bit Data errors can be corrected by
	// ECC.
	MultiBitECCErrorCorrection ErrorCorrection = "MultiBitECC"
	// AddressParityErrorCorrection Address Parity errors can be corrected.
	AddressParityErrorCorrection ErrorCorrection = "AddressParity"
)

// MemoryClassification is the memory classification.
type MemoryClassification string

const (
	// VolatileMemoryClassification Volatile memory.
	VolatileMemoryClassification MemoryClassification = "Volatile"
	// ByteAccessiblePersistentMemoryClassification Byte accessible
	// persistent memory.
	ByteAccessiblePersistentMemoryClassification MemoryClassification = "ByteAccessiblePersistent"
	// BlockMemoryClassification Block accessible memory.
	BlockMemoryClassification MemoryClassification = "Block"
)

// MemoryDeviceType is the type of memory.
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
	// DDR2SDRAMFBDIMMPROBEMemoryDeviceType DDR2 SDRAM FBDIMM PROBE.
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
	// LogicalMemoryDeviceType Logical Non-volatile device.
	LogicalMemoryDeviceType MemoryDeviceType = "Logical"
	// HBMMemoryDeviceType High Bandwidth Memory.
	HBMMemoryDeviceType MemoryDeviceType = "HBM"
	// HBM2MemoryDeviceType High Bandwidth Memory 2.
	HBM2MemoryDeviceType MemoryDeviceType = "HBM2"
)

// MemoryMedia is media type.
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

// MemoryType is the type of memory.
type MemoryType string

const (
	// DRAMMemoryType shall represent volatile DRAM.
	DRAMMemoryType MemoryType = "DRAM"
	// NVDIMMNMemoryType shall represent NVDIMMN as defined by JEDEC.
	NVDIMMNMemoryType MemoryType = "NVDIMM_N"
	// NVDIMMFMemoryType shall represent NVDIMMF as defined by JEDEC.
	NVDIMMFMemoryType MemoryType = "NVDIMM_F"
	// NVDIMMPMemoryType shall represent NVDIMMP as defined by JEDEC.
	NVDIMMPMemoryType MemoryType = "NVDIMM_P"
	// IntelOptaneMemoryType shall represent Intel Optane DC Persistent
	// Memory.
	IntelOptaneMemoryType MemoryType = "IntelOptane"
)

// OperatingMemoryModes is is the memory operating mode.
type OperatingMemoryModes string

const (
	// VolatileOperatingMemoryModes Volatile memory.
	VolatileOperatingMemoryModes OperatingMemoryModes = "Volatile"
	// PMEMOperatingMemoryModes Persistent memory, byte accessible through
	// system address space.
	PMEMOperatingMemoryModes OperatingMemoryModes = "PMEM"
	// BlockOperatingMemoryModes Block accessible system memory.
	BlockOperatingMemoryModes OperatingMemoryModes = "Block"
)

// SecurityStates is memory security state.
type SecurityStates string

const (
	// EnabledSecurityStates Secure mode is enabled and access to the data is
	// allowed.
	EnabledSecurityStates SecurityStates = "Enabled"
	// DisabledSecurityStates Secure mode is disabled.
	DisabledSecurityStates SecurityStates = "Disabled"
	// UnlockedSecurityStates Secure mode is enabled and access to the data
	// is unlocked.
	UnlockedSecurityStates SecurityStates = "Unlocked"
	// LockedSecurityStates Secure mode is enabled and access to the data is
	// locked.
	LockedSecurityStates SecurityStates = "Locked"
	// FrozenSecurityStates Secure state is frozen and can not be modified
	// until reset.
	FrozenSecurityStates SecurityStates = "Frozen"
	// PassphraselimitSecurityStates Number of attempts to unlock the Memory
	// exceeded limit.
	PassphraselimitSecurityStates SecurityStates = "Passphraselimit"
)

// Memory is used to represent the Memory in a Redfish implementation.
type Memory struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllocationAlignmentMiB shall be the alignment boundary on which memory
	// regions are allocated, measured in MiB.
	AllocationAlignmentMiB int
	// AllocationIncrementMiB shall be the allocation increment for regions,
	// measured in MiB.
	AllocationIncrementMiB int
	// AllowedSpeedsMHz shall be the speed supported by this Memory (in MHz).
	AllowedSpeedsMHz []int
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// BaseModuleType shall be the base module type of Memory.
	BaseModuleType BaseModuleType
	// BusWidthBits shall be the bus width in bits.
	BusWidthBits int
	// CacheSizeMiB shall be the total size of the cache portion memory in MiB.
	CacheSizeMiB int
	// CapacityMiB shall be the Memory capacity in MiB.
	CapacityMiB int
	// ConfigurationLocked shall be the current configuration lock state of this
	// memory. True shall indicate that the configuration is locked and cannot
	// be altered. False shall indicate that the configuration is not locked and
	// may be altered.
	ConfigurationLocked bool
	// DataWidthBits shall be the data width in bits.
	DataWidthBits int
	// Description provides a description of this resource.
	Description string
	// DeviceLocator shall be location of the Memory in the platform, typically
	// marked in the silk screen.
	DeviceLocator string
	// ErrorCorrection shall be the error correction scheme supported for this memory.
	ErrorCorrection ErrorCorrection
	// FirmwareAPIVersion shall be the version of API supported by the firmware.
	FirmwareAPIVersion string `json:"FirmwareApiVersion"`
	// FirmwareRevision shall be the revision of firmware on the Memory controller.
	FirmwareRevision string
	// IsRankSpareEnabled shall be true if a rank spare is enabled for this Memory.
	IsRankSpareEnabled bool
	// IsSpareDeviceEnabled shall be true if a spare device is enabled for this Memory.
	IsSpareDeviceEnabled bool
	// Location shall contain location information of the associated memory.
	Location common.Location
	// LogicalSizeMiB shall be the total size of the logical memory in MiB.
	LogicalSizeMiB int
	// Manufacturer shall contain a string which identifies the manufacturer of the Memory.
	Manufacturer string
	// MaxTDPMilliWatts shall be the maximum power budgets supported by the
	// Memory in milli Watts.
	MaxTDPMilliWatts []int
	// MemoryDeviceType shall be the Memory Device Type as defined by SMBIOS.
	MemoryDeviceType MemoryDeviceType
	// MemoryLocation shall contain properties which describe the Memory
	// connection information to sockets and memory controllers.
	MemoryLocation MemoryLocation
	// MemoryMedia shall be the media types of this Memory.
	MemoryMedia []MemoryMedia
	// MemorySubsystemControllerManufacturerID shall be the two byte
	// manufacturer ID of the memory subsystem controller of this memory module
	// as defined by JEDEC in JEP-106.
	MemorySubsystemControllerManufacturerID string
	// MemorySubsystemControllerProductID shall
	// be the two byte product ID of the memory subsystem controller of this
	// memory module as defined by the manufacturer.
	MemorySubsystemControllerProductID string
	// MemoryType shall be the type of Memory
	// represented by this resource.
	MemoryType MemoryType
	// Metrics is a reference to the Metrics associated with this Memory.
	metrics string
	// ModuleManufacturerID shall be the two byte manufacturer ID of this memory
	// module as defined by JEDEC in JEP-106.
	ModuleManufacturerID string
	// ModuleProductID shall be the two byte
	// product ID of this memory module as defined by the manufacturer.
	ModuleProductID string
	// NonVolatileSizeMiB shall be the total
	// size of the non-volatile portion memory in MiB.
	NonVolatileSizeMiB int
	// OperatingMemoryModes shall be the memory
	// modes supported by the Memory.
	OperatingMemoryModes []OperatingMemoryModes
	// OperatingSpeedMhz shall be the operating
	// speed of Memory in MHz or MT/s (mega-transfers per second) as reported
	// by the memory device. Memory devices which operate at their bus speed
	// shall report the operating speed in MHz (bus speed), while memory
	// device which transfer data faster than their bus speed (e.g. DDR
	// memory) shall report the operating speed in MT/s (mega-
	// transfers/second). In any case, the reported value shall match the
	// conventionally reported values for the technology utilized by the
	// memory device.
	OperatingSpeedMhz int
	// PartNumber shall indicate the part number as provided by the manufacturer
	// of this Memory.
	PartNumber string
	// PersistentRegionNumberLimit shall be the total number of persistent
	// regions this Memory can support.
	PersistentRegionNumberLimit int
	// PersistentRegionSizeLimitMiB shall be the total size of persistent regions in MiB.
	PersistentRegionSizeLimitMiB int
	// PersistentRegionSizeMaxMiB shall be the maximum size of a single persistent
	// regions in MiB.
	PersistentRegionSizeMaxMiB int
	// PowerManagementPolicy shall contain properties which describe the power
	// management policy for the current resource.
	PowerManagementPolicy PowerManagementPolicy
	// RankCount is used for spare or interleave.
	RankCount int
	// Regions shall be the memory region information within the Memory.
	Regions []RegionSet
	// SecurityCapabilities shall contain properties which describe the security
	// capabilities of the Memory.
	SecurityCapabilities SecurityCapabilities
	// SecurityState shall be the current security state of this memory.
	SecurityState SecurityStates
	// SerialNumber shall indicate the serial number as provided by the
	// manufacturer of this Memory.
	SerialNumber string
	// SpareDeviceCount is used spare devices available in the Memory. If
	// memory devices fails, the spare device could be used.
	SpareDeviceCount int
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// VolatileRegionNumberLimit shall be the total number of volatile regions
	// this Memory can support.
	VolatileRegionNumberLimit int
	// VolatileRegionSizeLimitMiB shall be the total size of volatile regions in MiB.
	VolatileRegionSizeLimitMiB int
	// VolatileRegionSizeMaxMiB shall be the maximum size of a single volatile
	// regions in MiB.
	VolatileRegionSizeMaxMiB int
	// VolatileSizeMiB shall be the total size of the volatile portion memory in MiB.
	VolatileSizeMiB int
	// Chassis shall be a reference to a resource of type Chassis that represent
	// the physical container associated with this Memory.
	chassis string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Memory object from the raw JSON.
func (memory *Memory) UnmarshalJSON(b []byte) error {
	type temp Memory
	type links struct {
		Chassis common.Link
	}
	var t struct {
		temp
		Links    links
		Assembly common.Link
		Metrics  common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memory = Memory(t.temp)

	// Extract the links to other entities for later
	memory.assembly = string(t.Assembly)
	memory.metrics = string(t.Metrics)
	memory.chassis = string(t.Links.Chassis)

	// This is a read/write object, so we need to save the raw object data for later
	memory.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (memory *Memory) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Memory)
	err := original.UnmarshalJSON(memory.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"SecurityState",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memory).Elem()

	return memory.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemory will get a Memory instance from the service.
func GetMemory(c common.Client, uri string) (*Memory, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memory Memory
	err = json.NewDecoder(resp.Body).Decode(&memory)
	if err != nil {
		return nil, err
	}

	memory.SetClient(c)
	return &memory, nil
}

// ListReferencedMemorys gets the collection of Memory from
// a provided reference.
func ListReferencedMemorys(c common.Client, link string) ([]*Memory, error) {
	var result []*Memory
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, memoryLink := range links.ItemLinks {
		memory, err := GetMemory(c, memoryLink)
		if err != nil {
			return result, err
		}
		result = append(result, memory)
	}

	return result, nil
}

// Assembly gets this memory's assembly.
func (memory *Memory) Assembly() (*Assembly, error) {
	if memory.assembly == "" {
		return nil, nil
	}
	return GetAssembly(memory.Client, memory.assembly)
}

// Metrics gets the memory metrics.
func (memory *Memory) Metrics() (*MemoryMetrics, error) {
	if memory.metrics == "" {
		return nil, nil
	}
	return GetMemoryMetrics(memory.Client, memory.metrics)
}

// Chassis gets the containing chassis of this memory.
func (memory *Memory) Chassis() (*Chassis, error) {
	if memory.chassis == "" {
		return nil, nil
	}
	return GetChassis(memory.Client, memory.chassis)
}

// MemoryLocation shall contain properties which describe the Memory connection
// information to sockets and memory controllers.
type MemoryLocation struct {
	// Channel is Channel number in which Memory is connected.
	Channel int
	// MemoryController is Memory controller number in which Memory is
	// connected.
	MemoryController int
	// Slot is Slot number in which Memory is connected.
	Slot int
	// Socket is Socket number in which Memory is connected.
	Socket int
}

// PowerManagementPolicy shall contain properties which describe the power
// management policy for the current resource.
type PowerManagementPolicy struct {
	// AveragePowerBudgetMilliWatts is Average power budget in milli watts.
	AveragePowerBudgetMilliWatts int
	// MaxTDPMilliWatts is Maximum TDP in milli watts.
	MaxTDPMilliWatts int
	// PeakPowerBudgetMilliWatts is Peak power budget in milli watts.
	PeakPowerBudgetMilliWatts int
	// PolicyEnabled is Power management policy enabled status.
	PolicyEnabled bool
}

// RegionSet shall describe the memory region information within a Memory entity.
type RegionSet struct {
	// MemoryClassification is Classification of memory occupied by the given
	// memory region.
	MemoryClassification MemoryClassification
	// OffsetMiB is Offset with in the Memory that corresponds to the
	// starting of this memory region in MiB.
	OffsetMiB int
	// PassphraseEnabled shall be a boolean
	// indicating if the passphrase is enabled for this region.
	PassphraseEnabled bool
	// RegionID is Unique region ID representing a specific region within the
	// Memory.
	RegionID string `json:"RegionId"`
	// SizeMiB is Size of this memory region in MiB.
	SizeMiB int
}

// SecurityCapabilities shall contain properties which describe the security
// capabilities of a Memory entity.
type SecurityCapabilities struct {
	// ConfigurationLockCapable shall indicate whether this memory supports the
	// locking (freezing) of the configuration.
	ConfigurationLockCapable bool
	// DataLockCapable shall indicate whether this memory supports the locking
	// of data access.
	DataLockCapable bool
	// MaxPassphraseCount is Maximum number of passphrases supported for this
	// Memory.
	MaxPassphraseCount int
	// PassphraseCapable is Memory passphrase set capability.
	PassphraseCapable bool
	// PassphraseLockLimit shall be the maximum number of incorrect passphrase
	// access attempts allowed before access to data is locked. A value of zero
	// shall indicate that there is no limit to the number of attempts.
	PassphraseLockLimit int
}
