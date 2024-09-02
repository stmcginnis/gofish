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
	// LogicalMemoryDeviceType Logical device, such as when the memory is fabric-attached.
	LogicalMemoryDeviceType MemoryDeviceType = "Logical"
	// HBMMemoryDeviceType High Bandwidth Memory.
	HBMMemoryDeviceType MemoryDeviceType = "HBM"
	// HBM2MemoryDeviceType The second generation of High Bandwidth Memory.
	HBM2MemoryDeviceType MemoryDeviceType = "HBM2"
	// HBM2EMemoryDeviceType An updated version of the second generation of High Bandwidth Memory.
	HBM2EMemoryDeviceType MemoryDeviceType = "HBM2E"
	// HBM3MemoryDeviceType The third generation of High Bandwidth Memory.
	HBM3MemoryDeviceType MemoryDeviceType = "HBM3"
	// GDDRMemoryDeviceType Synchronous graphics random-access memory.
	GDDRMemoryDeviceType MemoryDeviceType = "GDDR"
	// GDDR2MemoryDeviceType Double data rate type two synchronous graphics random-access memory.
	GDDR2MemoryDeviceType MemoryDeviceType = "GDDR2"
	// GDDR3MemoryDeviceType Double data rate type three synchronous graphics random-access memory.
	GDDR3MemoryDeviceType MemoryDeviceType = "GDDR3"
	// GDDR4MemoryDeviceType Double data rate type four synchronous graphics random-access memory.
	GDDR4MemoryDeviceType MemoryDeviceType = "GDDR4"
	// GDDR5MemoryDeviceType Double data rate type five synchronous graphics random-access memory.
	GDDR5MemoryDeviceType MemoryDeviceType = "GDDR5"
	// GDDR5XMemoryDeviceType Double data rate type five X synchronous graphics random-access memory.
	GDDR5XMemoryDeviceType MemoryDeviceType = "GDDR5X"
	// GDDR6MemoryDeviceType Double data rate type six synchronous graphics random-access memory.
	GDDR6MemoryDeviceType MemoryDeviceType = "GDDR6"
	// DDR5MemoryDeviceType Double data rate type five synchronous dynamic random-access memory.
	DDR5MemoryDeviceType MemoryDeviceType = "DDR5"
	// OEMMemoryDeviceType OEM-defined.
	OEMMemoryDeviceType MemoryDeviceType = "OEM"
	// LPDDR5SDRAMMemoryDeviceType LPDDR5 SDRAM.
	LPDDR5SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR5_SDRAM"
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

// MemoryCXL shall contain CXL-specific properties for a memory device.
type MemoryCXL struct {
	// LabelStorageSizeBytes shall contain the size of the label storage area in bytes of this memory device.
	LabelStorageSizeBytes int
	// StagedNonVolatileSizeMiB shall indicate the total device non-volatile memory capacity in mebibytes. The value
	// shall be in multiples of 256 mebibytes.
	StagedNonVolatileSizeMiB int
	// StagedVolatileSizeMiB shall indicate the total device volatile memory capacity in mebibytes staged for next
	// activation. This value shall be in multiples of 256 mebibytes.
	StagedVolatileSizeMiB int
}

// MemoryHealthData is the health data of a memory device.
type MemoryHealthData struct {
	// PredictedMediaLifeLeftPercent is the current health of the memory device as a percentage.
	// This property has been deprecated in favor of PredictedMediaLifeLeftPercent in the MemoryMetrics resource.
	PredictedMediaLifeLeftPercent int
}

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
	// CXL shall contain CXL-specific properties for this memory device.
	CXL MemoryCXL
	// CacheSizeMiB shall be the total size of the cache portion memory in MiB.
	CacheSizeMiB int
	// CapacityMiB shall be the Memory capacity in MiB.
	CapacityMiB int
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates []string
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
	// This property has been deprecated in favor of the ServiceLabel property within Location.
	DeviceLocator string
	// Enabled shall indicate if this memory is enabled.
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this memory.
	environmentMetrics string
	// ErrorCorrection shall be the error correction scheme supported for this memory.
	ErrorCorrection ErrorCorrection
	// FirmwareAPIVersion shall be the version of API supported by the firmware.
	FirmwareAPIVersion string `json:"FirmwareApiVersion"`
	// FirmwareRevision shall be the revision of firmware on the Memory controller.
	FirmwareRevision string
	// HealthData shall contain the health data of this memory device.
	HealthData string
	// IsRankSpareEnabled shall be true if a rank spare is enabled for this Memory.
	IsRankSpareEnabled bool
	// IsSpareDeviceEnabled shall be true if a spare device is enabled for this Memory.
	IsSpareDeviceEnabled bool
	// Location shall contain location information of the associated memory.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Log shall contain a link to a resource of type LogService.
	log string
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
	// Model shall indicate the model information as provided by the manufacturer of this memory.
	Model string
	// ModuleManufacturerID shall be the two byte manufacturer ID of this memory
	// module as defined by JEDEC in JEP-106.
	ModuleManufacturerID string
	// ModuleProductID shall be the two byte
	// product ID of this memory module as defined by the manufacturer.
	ModuleProductID string
	// NonVolatileSizeLimitMiB shall contain the total non-volatile memory capacity in mebibytes (MiB).
	NonVolatileSizeLimitMiB int
	// NonVolatileSizeMiB shall contain the total size of the non-volatile portion memory in MiB.
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
	// PoisonListMaxMediaErrorRecords shall contain the maximum number of media error records this device can track in
	// its poison list.
	PoisonListMaxMediaErrorRecords int
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
	// SecurityStates shall contain the security states of this memory device.
	SecurityStates SecurityStateInfo
	// SerialNumber shall indicate the serial number as provided by the
	// manufacturer of this Memory.
	SerialNumber string
	// SpareDeviceCount is used spare devices available in the Memory. If
	// memory devices fails, the spare device could be used.
	SpareDeviceCount int
	// SparePartNumber shall contain the spare part number of the memory.
	SparePartNumber string
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
	// VolatileSizeLimitMiB shall contain the total volatile memory capacity in mebibytes (MiB).
	VolatileSizeLimitMiB int
	// VolatileSizeMiB shall be the total size of the volatile portion memory in MiB.
	VolatileSizeMiB int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	batteries []string
	// BatteriesCount gets the number of batteries that provide power to this memory device during
	// a power-loss event, such as with battery-backed NVDIMMs.
	BatteriesCount int
	chassis        string
	endpoints      []string
	// EndpointsCount gets the number of endpoints associated with this memory.
	EndpointsCount     int
	memoryMediaSources []string
	// MemoryMediaSourcesCount gets the number of memory chunk instances providing media for this memory.
	MemoryMediaSourcesCount  int
	memoryRegionMediaSources []string
	// MemoryMediaRegionSourcesCount gets the number of memory region instances providing media for this memory.
	MemoryRegionMediaSourcesCount int
	processors                    []string
	// ProcessorsCount gets the number of processors associated with this memory device.
	ProcessorsCount int

	disableMasterPassphraseTarget string
	disablePassphraseTarget       string
	freezeSecurityStateTarget     string
	injectPersistentPoisonTarget  string
	overwriteUnitTarget           string
	resetTarget                   string
	resetToDefaultsTarget         string
	scanMediaTarget               string
	secureEraseUnitTarget         string
	setMasterPassphraseTarget     string
	setPassphraseTarget           string
	unlockUnitTarget              string
}

type memoryActions struct {
	DisableMasterPassphrase struct {
		Target string
	} `json:"#Memory.DisableMasterPassphrase"`
	DisablePassphrase struct {
		Target string
	} `json:"#Memory.DisablePassphrase"`
	FreezeSecurityState struct {
		Target string
	} `json:"#Memory.FreezeSecurityState"`
	InjectPersistentPoison struct {
		Target string
	} `json:"#Memory.InjectPersistentPoison"`
	OverwriteUnit struct {
		Target string
	} `json:"#Memory.OverwriteUnit"`
	Reset struct {
		Target string
	} `json:"#Memory.Reset"`
	ResetToDefaults struct {
		Target string
	} `json:"#Memory.ResetToDefaults"`
	ScanMedia struct {
		Target string
	} `json:"#Memory.ScanMedia"`
	SecureEraseUnit struct {
		Target string
	} `json:"#Memory.SecureEraseUnit"`
	SetMasterPassphrase struct {
		Target string
	} `json:"#Memory.SetMasterPassphrase"`
	SetPassphrase struct {
		Target string
	} `json:"#Memory.SetPassphrase"`
	UnlockUnit struct {
		Target string
	} `json:"#Memory.UnlockUnit"`
}

type memoryLinks struct {
	// Batteries shall contain an array of links to resources of type Battery that represent the batteries that provide
	// power to this memory device during a power-loss event, such as with battery-backed NVDIMMs. This property shall
	// not be present if the batteries power the containing chassis as a whole rather than the individual memory
	// device.
	Batteries      common.Links
	BatteriesCount int `json:"Batteries@odata.count"`
	// Chassis shall contain a link to a resource of type Chassis that represents the physical container associated
	// with this memory device.
	Chassis common.Link
	// Endpoints shall contain an array of links to resources of type Endpoint that represent the endpoints associated
	// with this memory.
	Endpoints      common.Links
	EndpointsCount int `json:"Endpoints@odata.count"`
	// MemoryMediaSources shall contain an array of links to resources of type MemoryChunks that represent the memory
	// chunk instances providing media for this memory.
	MemoryMediaSources      common.Links
	MemoryMediaSourcesCount int `json:"MemoryMediaSources@odata.count"`
	// MemoryRegionMediaSources shall contain an array of links to resources of type MemoryRegion that represent the
	// memory region instances providing media for this memory.
	MemoryRegionMediaSources      common.Links
	MemoryRegionMediaSourcesCount int `json:"MemoryRegionMediaSources@odata.count"`
	// Processors shall contain an array of links to resources of type Processor that are associated with this memory
	// device.
	Processors      common.Links
	ProcessorsCount int `json:"Processors@odata.count"`
}

// UnmarshalJSON unmarshals a Memory object from the raw JSON.
func (memory *Memory) UnmarshalJSON(b []byte) error {
	type temp Memory
	var t struct {
		temp
		Actions            memoryActions
		Links              memoryLinks
		Assembly           common.Link
		Certificates       common.LinksCollection
		EnvironmentMetrics common.Link
		Log                common.Link
		Metrics            common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memory = Memory(t.temp)

	// Extract the links to other entities for later
	memory.assembly = t.Assembly.String()
	memory.certificates = t.Certificates.ToStrings()
	memory.environmentMetrics = t.EnvironmentMetrics.String()
	memory.log = t.Log.String()
	memory.metrics = t.Metrics.String()

	memory.batteries = t.Links.Batteries.ToStrings()
	memory.BatteriesCount = t.Links.BatteriesCount
	memory.chassis = t.Links.Chassis.String()
	memory.endpoints = t.Links.Endpoints.ToStrings()
	memory.EndpointsCount = t.Links.EndpointsCount
	memory.memoryMediaSources = t.Links.MemoryMediaSources.ToStrings()
	memory.MemoryMediaSourcesCount = t.Links.MemoryMediaSourcesCount
	memory.memoryRegionMediaSources = t.Links.MemoryRegionMediaSources.ToStrings()
	memory.MemoryRegionMediaSourcesCount = t.Links.MemoryRegionMediaSourcesCount
	memory.processors = t.Links.Processors.ToStrings()
	memory.ProcessorsCount = t.Links.ProcessorsCount

	memory.disableMasterPassphraseTarget = t.Actions.DisableMasterPassphrase.Target
	memory.disablePassphraseTarget = t.Actions.DisablePassphrase.Target
	memory.freezeSecurityStateTarget = t.Actions.FreezeSecurityState.Target
	memory.injectPersistentPoisonTarget = t.Actions.InjectPersistentPoison.Target
	memory.overwriteUnitTarget = t.Actions.OverwriteUnit.Target
	memory.resetTarget = t.Actions.Reset.Target
	memory.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	memory.scanMediaTarget = t.Actions.ScanMedia.Target
	memory.secureEraseUnitTarget = t.Actions.SecureEraseUnit.Target
	memory.setMasterPassphraseTarget = t.Actions.SetMasterPassphrase.Target
	memory.setPassphraseTarget = t.Actions.SetPassphrase.Target
	memory.unlockUnitTarget = t.Actions.UnlockUnit.Target

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
		"Enabled",
		"LocationIndicatorActive",
		"NonVolatileSizeLimitMiB",
		"OperatingSpeedRangeMHz",
		"PoisonListMaxMediaErrorRecords",
		"SecurityState",
		"VolatileSizeLimitMiB",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memory).Elem()

	return memory.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemory will get a Memory instance from the service.
func GetMemory(c common.Client, uri string) (*Memory, error) {
	return common.GetObject[Memory](c, uri)
}

// ListReferencedMemorys gets the collection of Memory from
// a provided reference.
func ListReferencedMemorys(c common.Client, link string) ([]*Memory, error) {
	return common.GetCollectionObjects[Memory](c, link)
}

// Assembly gets this memory's assembly.
func (memory *Memory) Assembly() (*Assembly, error) {
	if memory.assembly == "" {
		return nil, nil
	}
	return GetAssembly(memory.GetClient(), memory.assembly)
}

// Certificates gets certificates for device identity and attestation.
func (memory *Memory) Certificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](memory.GetClient(), memory.certificates)
}

// EnvironmentMetrics gets the environment metrics for this memory.
func (memory *Memory) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if memory.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(memory.GetClient(), memory.environmentMetrics)
}

// Log gets the log service for this memory.
func (memory *Memory) Log() (*LogService, error) {
	if memory.log == "" {
		return nil, nil
	}
	return GetLogService(memory.GetClient(), memory.log)
}

// Metrics gets the memory metrics.
func (memory *Memory) Metrics() (*MemoryMetrics, error) {
	if memory.metrics == "" {
		return nil, nil
	}
	return GetMemoryMetrics(memory.GetClient(), memory.metrics)
}

// Batteries gets the batteries that provide power to this memory device during
// a power-loss event, such as with battery-backed NVDIMMs.
func (memory *Memory) Batteries() ([]*Battery, error) {
	return common.GetObjects[Battery](memory.GetClient(), memory.batteries)
}

// Chassis gets the containing chassis of this memory.
func (memory *Memory) Chassis() (*Chassis, error) {
	if memory.chassis == "" {
		return nil, nil
	}
	return GetChassis(memory.GetClient(), memory.chassis)
}

// Endpoints gets the endpoints associated with this memory.
func (memory *Memory) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](memory.GetClient(), memory.endpoints)
}

// MemoryMediaSources gets the memory chunks providing media for this memory.
func (memory *Memory) MemoryMediaSources() ([]*MemoryChunks, error) {
	return common.GetObjects[MemoryChunks](memory.GetClient(), memory.memoryMediaSources)
}

// MemoryRegionMediaSources gets the memory regions providing media for this memory.
func (memory *Memory) MemoryRegionMediaSources() ([]*MemoryRegion, error) {
	return common.GetObjects[MemoryRegion](memory.GetClient(), memory.memoryRegionMediaSources)
}

// Processors gets the processors associated with this memory device.
func (memory *Memory) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](memory.GetClient(), memory.processors)
}

// DisalbeMasterPassphrase will disable the master passphrase on the supplied
// region provided the supplied master passphrase matches that of the region.
func (memory *Memory) DisableMasterPassphrase(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.disableMasterPassphraseTarget, param)
}

// DisablePassphrase will disable the need for passphrases on the supplied
// region provided the supplied passphrase matches that of the region.
func (memory *Memory) DisablePassphrase(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.disablePassphraseTarget, param)
}

// FreezeSecurityState will freeze the security state of the memory device.
func (memory *Memory) FreezeSecurityState() error {
	return memory.Post(memory.freezeSecurityStateTarget, nil)
}

// InjectPersistentPoison will inject poison to a specific persistent memory address
// in the memory device.
func (memory *Memory) InjectPersistentPoison(physicalAddress string) error {
	param := struct {
		PhysicalAddress string
	}{
		PhysicalAddress: physicalAddress,
	}
	return memory.Post(memory.injectPersistentPoisonTarget, param)
}

// OverwriteUnit will securely erase the supplied region provided the supplied
// passphrase matches that of the given region using the NIST SP800-88 Purge: Overwrite.
// Use the SecureEraseUnit method to perform NIST SP800-88 Purge: Cryptographic Erase.
func (memory *Memory) OverwriteUnit(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.overwriteUnitTarget, param)
}

// Reset resets this memory device.
func (memory *Memory) Reset(resetType ResetType) error {
	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	return memory.Post(memory.resetTarget, t)
}

// ResetToDefaults will reset the values of writable properties in this resource
// to their default values as specified by the manufacturer.
func (memory *Memory) ResetToDefaults() error {
	return memory.Post(memory.resetToDefaultsTarget, nil)
}

// ScanMedia will scan the media of the memory device.
// `length` is the length of the target region to scan in bytes from the physical address.
// `noEventLog` is used to indicate whether events related to the media scan are not logged.
// `physicalAddress` is the starting device physical address to scan as a hex-encoded string.
func (memory *Memory) ScanMedia(length int, noEventLog bool, physicalAddress string) error {
	param := struct {
		Length          int
		NoEventLog      bool
		PhysicalAddress string
	}{
		Length:          length,
		NoEventLog:      noEventLog,
		PhysicalAddress: physicalAddress,
	}
	return memory.Post(memory.scanMediaTarget, param)
}

// SecureEraseUnit will securely erase the supplied region provided the supplied passphrase
// matches that of the given region using the NIST SP800-88 Purge: Cryptographic Erase.
// Use the OverwriteUnit method to perform NIST SP800-88 Purge: Overwrite.
func (memory *Memory) SecureEraseUnit(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.secureEraseUnitTarget, param)
}

// SetMasterPassphrase will set the supplied master passphrase to the supplied region.
func (memory *Memory) SetMasterPassphrase(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.setMasterPassphraseTarget, param)
}

// SetPassphrase will apply the supplied passphrase to the supplied region.
func (memory *Memory) SetPassphrase(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.setPassphraseTarget, param)
}

// UnlockUnit will apply the supplied passphrase to the supplied region for the purpose
// of unlocking the given regions.
func (memory *Memory) UnlockUnit(passphrase, regionID string) error {
	param := memoryParams{
		Passphrase: passphrase,
		RegionID:   regionID,
	}
	return memory.Post(memory.unlockUnitTarget, param)
}

type memoryParams struct {
	Passphrase string
	RegionID   string `json:"RegionId"`
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
	// MasterPassphraseEnabled shall indicate whether the master passphrase is enabled for this region.
	MasterPassphraseEnabled bool
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

// SecurityStateInfo shall contain the security states of a memory device.
type SecurityStateInfo struct {
	// MasterPassphraseAttemptCountReached shall indicate whether an incorrect master passphrase attempt count has been
	// reached.
	MasterPassphraseAttemptCountReached bool
	// UserPassphraseAttemptCountReached shall indicate whether an incorrect user passphrase attempt count has been
	// reached.
	UserPassphraseAttemptCountReached bool
}
