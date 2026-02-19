//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ComputerSystem.v1_27_0.json
// 2025.4 - #ComputerSystem.v1_27_0.ComputerSystem

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type AutomaticRetryConfig string

const (
	// DisabledAutomaticRetryConfig shall indicate that automatic retrying of
	// booting is disabled.
	DisabledAutomaticRetryConfig AutomaticRetryConfig = "Disabled"
	// RetryAttemptsAutomaticRetryConfig shall indicate that the number of retries
	// of booting is based on the 'AutomaticRetryAttempts' property, and the
	// 'RemainingAutomaticRetryAttempts' property indicates the number of remaining
	// attempts.
	RetryAttemptsAutomaticRetryConfig AutomaticRetryConfig = "RetryAttempts"
	// RetryAlwaysAutomaticRetryConfig shall indicate that the system will always
	// automatically retry booting.
	RetryAlwaysAutomaticRetryConfig AutomaticRetryConfig = "RetryAlways"
)

// BootOrderTypes is The enumerations of 'BootOrderTypes' specify the choice of
// boot order property to use when controller the persistent boot order for this
// computer system.
type BootOrderTypes string

const (
	// BootOrderBootOrderTypes The system uses the 'BootOrder' property to specify
	// the persistent boot order.
	BootOrderBootOrderTypes BootOrderTypes = "BootOrder"
	// AliasBootOrderBootOrderTypes The system uses the 'AliasBootOrder' property
	// to specify the persistent boot order.
	AliasBootOrderBootOrderTypes BootOrderTypes = "AliasBootOrder"
)

type BootProgressTypes string

const (
	// NoneBootProgressTypes shall indicate that the system is not booting or
	// running, such as the system is powered off.
	NoneBootProgressTypes BootProgressTypes = "None"
	// PrimaryProcessorInitializationStartedBootProgressTypes shall indicate that
	// the system has started to initialize the primary processor.
	PrimaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "PrimaryProcessorInitializationStarted"
	// BusInitializationStartedBootProgressTypes shall indicate that the system has
	// started to initialize the buses.
	BusInitializationStartedBootProgressTypes BootProgressTypes = "BusInitializationStarted"
	// MemoryInitializationStartedBootProgressTypes shall indicate that the system
	// has started to initialize the memory.
	MemoryInitializationStartedBootProgressTypes BootProgressTypes = "MemoryInitializationStarted"
	// SecondaryProcessorInitializationStartedBootProgressTypes shall indicate that
	// the system has started to initialize the secondary processors.
	SecondaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "SecondaryProcessorInitializationStarted"
	// PCIResourceConfigStartedBootProgressTypes shall indicate that the system has
	// started to initialize the PCI resources.
	PCIResourceConfigStartedBootProgressTypes BootProgressTypes = "PCIResourceConfigStarted"
	// SystemHardwareInitializationCompleteBootProgressTypes shall indicate that
	// the system has completed initializing all hardware.
	SystemHardwareInitializationCompleteBootProgressTypes BootProgressTypes = "SystemHardwareInitializationComplete"
	// SetupEnteredBootProgressTypes shall indicate that the system has entered the
	// setup utility.
	SetupEnteredBootProgressTypes BootProgressTypes = "SetupEntered"
	// OSBootStartedBootProgressTypes shall indicate that the operating system has
	// started to boot.
	OSBootStartedBootProgressTypes BootProgressTypes = "OSBootStarted"
	// OSRunningBootProgressTypes shall indicate that the operating system is
	// running and shall indicate the final boot progress state.
	OSRunningBootProgressTypes BootProgressTypes = "OSRunning"
	// OEMBootProgressTypes shall indicate an OEM-defined boot progress state.
	OEMBootProgressTypes BootProgressTypes = "OEM"
)

type BootSource string

const (
	// NoneBootSource Boot from the normal boot device.
	NoneBootSource BootSource = "None"
	// PxeBootSource Boot from the Pre-boot eXecution Environment (PXE).
	PxeBootSource BootSource = "Pxe"
	// FloppyBootSource Boot from the floppy disk drive.
	FloppyBootSource BootSource = "Floppy"
	// CdBootSource Boot from the CD or DVD.
	CdBootSource BootSource = "Cd"
	// UsbBootSource Boot from a system BIOS-specified USB device.
	UsbBootSource BootSource = "Usb"
	// HddBootSource Boot from a hard drive.
	HddBootSource BootSource = "Hdd"
	// BiosSetupBootSource Boot to the BIOS setup utility.
	BiosSetupBootSource BootSource = "BiosSetup"
	// UtilitiesBootSource Boot to the manufacturer's utilities program or
	// programs.
	UtilitiesBootSource BootSource = "Utilities"
	// DiagsBootSource Boot to the manufacturer's diagnostics program.
	DiagsBootSource BootSource = "Diags"
	// UefiShellBootSource Boot to the UEFI Shell.
	UefiShellBootSource BootSource = "UefiShell"
	// UefiTargetBootSource Boot to the UEFI device specified in the
	// 'UefiTargetBootSourceOverride' property.
	UefiTargetBootSource BootSource = "UefiTarget"
	// SDCardBootSource Boot from an SD card.
	SDCardBootSource BootSource = "SDCard"
	// UefiHTTPBootSource Boot from a UEFI HTTP network location.
	UefiHTTPBootSource BootSource = "UefiHttp"
	// RemoteDriveBootSource Boot from a remote drive, such as an iSCSI target.
	RemoteDriveBootSource BootSource = "RemoteDrive"
	// UefiBootNextBootSource Boot to the UEFI device that the 'BootNext' property
	// specifies.
	UefiBootNextBootSource BootSource = "UefiBootNext"
	// RecoveryBootSource Boot to a system-designated recovery process or image.
	RecoveryBootSource BootSource = "Recovery"
)

type BootSourceOverrideEnabled string

const (
	// DisabledBootSourceOverrideEnabled The system boots normally.
	DisabledBootSourceOverrideEnabled BootSourceOverrideEnabled = "Disabled"
	// OnceBootSourceOverrideEnabled On its next boot cycle, the system boots one
	// time to the boot source override target. Then, the
	// 'BootSourceOverrideEnabled' value is reset to 'Disabled'.
	OnceBootSourceOverrideEnabled BootSourceOverrideEnabled = "Once"
	// ContinuousBootSourceOverrideEnabled The system boots to the target specified
	// in the 'BootSourceOverrideTarget' property until this property is
	// 'Disabled'.
	ContinuousBootSourceOverrideEnabled BootSourceOverrideEnabled = "Continuous"
)

type BootSourceOverrideMode string

const (
	// LegacyBootSourceOverrideMode The system boots in non-UEFI boot mode to the
	// boot source override target.
	LegacyBootSourceOverrideMode BootSourceOverrideMode = "Legacy"
	// UEFIBootSourceOverrideMode The system boots in UEFI boot mode to the boot
	// source override target.
	UEFIBootSourceOverrideMode BootSourceOverrideMode = "UEFI"
)

type Component string

const (
	// AllComponent shall export all available configuration data from the system
	// including OEM components.
	AllComponent Component = "All"
	// ManagerComponent shall export configuration data associated with any
	// managers in the system.
	ManagerComponent Component = "Manager"
	// BIOSComponent shall export configuration data associated with the BIOS for
	// the system.
	BIOSComponent Component = "BIOS"
	// NetworkComponent shall export configuration data associated with the network
	// devices for the system.
	NetworkComponent Component = "Network"
	// StorageComponent shall export configuration data associated with the storage
	// devices for the system.
	StorageComponent Component = "Storage"
	// PowerDistributionComponent shall export configuration data associated with
	// the power distribution unit functions and subsystems of this equipment.
	PowerDistributionComponent Component = "PowerDistribution" // from powerdistribution.go
	// ManagerAccountsComponent shall export configuration data associated with any
	// managers for the equipment.
	ManagerAccountsComponent Component = "ManagerAccounts" // from powerdistribution and coolunit
	// CoolingUnitComponent shall export configuration data associated with the
	// cooling unit functions and subsystems of this equipment.
	CoolingUnitComponent Component = "CoolingUnit" // from coolingunit
)

type CompositionUseCase string

const (
	// ResourceBlockCapableCompositionUseCase shall indicate the computer system
	// supports being registered as a resource block in order for it to participate
	// in composition requests.
	ResourceBlockCapableCompositionUseCase CompositionUseCase = "ResourceBlockCapable"
	// ExpandableSystemCompositionUseCase shall indicate the computer system
	// supports expandable system composition and is associated with a resource
	// block.
	ExpandableSystemCompositionUseCase CompositionUseCase = "ExpandableSystem"
)

type DecommissionType string

const (
	// AllDecommissionType shall indicate the service removes all the data that it
	// can from the system. This shall include all possible OEM data as well.
	AllDecommissionType DecommissionType = "All"
	// UserDataDecommissionType shall indicate the service removes all the data
	// from block devices or other operating system accessible storage. If the
	// 'RequireSecureErase' parameter contains 'true', this shall be equivalent to
	// performing the SecureErase action on each drive.
	UserDataDecommissionType DecommissionType = "UserData"
	// ManagerConfigDecommissionType shall indicate the service resets all
	// associated managers to factory defaults. This shall be equivalent to
	// performing the 'ResetToDefaults' action on each 'Manager' resource with the
	// 'ResetType' parameter of 'ResetAll'.
	ManagerConfigDecommissionType DecommissionType = "ManagerConfig"
	// BIOSConfigDecommissionType shall indicate the service resets all BIOS
	// settings to factory defaults. This shall be equivalent to performing the
	// 'ResetBios' action on each 'Bios' resource.
	BIOSConfigDecommissionType DecommissionType = "BIOSConfig"
	// NetworkConfigDecommissionType shall indicate the service resets all network
	// settings on all network devices to factory defaults.
	NetworkConfigDecommissionType DecommissionType = "NetworkConfig"
	// StorageConfigDecommissionType shall indicate the service resets all storage
	// controller settings to factory defaults. This shall be equivalent to
	// performing the 'ResetToDefaults' action on each 'Storage' resource with the
	// 'ResetType' parameter of 'PreserveVolumes'.
	StorageConfigDecommissionType DecommissionType = "StorageConfig"
	// LogsDecommissionType shall indicate the service clears all logs. This shall
	// be equivalent to performing the 'ClearLog' action on each 'LogService'
	// resource.
	LogsDecommissionType DecommissionType = "Logs"
	// TPMDecommissionType shall indicate the service resets all user-accessible
	// TPM device settings to factory defaults. All sensitive data stored within
	// the applicable TPMs shall be erased.
	TPMDecommissionType DecommissionType = "TPM"
)

type ExportSecurity string

const (
	// IncludeSensitiveDataExportSecurity shall export all requested data
	// regardless of the sensitivity.
	IncludeSensitiveDataExportSecurity ExportSecurity = "IncludeSensitiveData"
	// HashedDataOnlyExportSecurity shall export requested data including hashed
	// passwords, but shall exclude other sensitive data.
	HashedDataOnlyExportSecurity ExportSecurity = "HashedDataOnly"
	// ExcludeSensitiveDataExportSecurity shall export only non-sensitive data.
	ExcludeSensitiveDataExportSecurity ExportSecurity = "ExcludeSensitiveData"
)

type ExportType string

const (
	// NonDestructiveExportType shall export only configuration data that would not
	// potentially result in data loss on import.
	NonDestructiveExportType ExportType = "NonDestructive"
	// CloneWithinFabricExportType shall export only configuration data that would
	// not result in network collisions if applied to another system on a shared
	// fabric.
	CloneWithinFabricExportType ExportType = "CloneWithinFabric"
	// ReplacementExportType shall export all configuration data required to
	// replace this system.
	ReplacementExportType ExportType = "Replacement"
	// CloneExportType shall export configuration data for this equipment that
	// allows for configuration to be duplicated on other equipment as defined by
	// the vendor. The service shall export only data which would not result in
	// problems if applied to another instance of this equipment. For example,
	// identifiers such as MAC Addresses and UUIDs will be excluded under this
	// option. Consult the vendor documentation for this equipment for more
	// information about which equipment is able to accept the resulting
	// configuration data.
	CloneExportType ExportType = "Clone" // from powerdistribution
)

type GraphicalConnectTypesSupported string

const (
	// KVMIPGraphicalConnectTypesSupported The controller supports a graphical
	// console connection through a KVM-IP (redirection of Keyboard, Video, Mouse
	// over IP) protocol.
	KVMIPGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "KVMIP"
	// OEMGraphicalConnectTypesSupported The controller supports a graphical
	// console connection through an OEM-specific protocol.
	OEMGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "OEM"
	// OemGraphicalConnectTypesSupported The controller supports a graphical
	// console connection through an OEM-specific protocol.
	OemGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "Oem" // from manager
)

// HostingRole is The enumerations of 'HostingRole' specify different features
// that the hosting computer system supports.
type HostingRole string

const (
	// ApplicationServerHostingRole The system hosts functionality that supports
	// general purpose applications.
	ApplicationServerHostingRole HostingRole = "ApplicationServer"
	// StorageServerHostingRole The system hosts functionality that supports the
	// system acting as a storage server.
	StorageServerHostingRole HostingRole = "StorageServer"
	// SwitchHostingRole The system hosts functionality that supports the system
	// acting as a switch.
	SwitchHostingRole HostingRole = "Switch"
	// ApplianceHostingRole The system hosts functionality that supports the system
	// acting as an appliance.
	ApplianceHostingRole HostingRole = "Appliance"
	// BareMetalServerHostingRole The system hosts functionality that supports the
	// system acting as a bare-metal server.
	BareMetalServerHostingRole HostingRole = "BareMetalServer"
	// VirtualMachineServerHostingRole The system hosts functionality that supports
	// the system acting as a virtual machine server.
	VirtualMachineServerHostingRole HostingRole = "VirtualMachineServer"
	// ContainerServerHostingRole The system hosts functionality that supports the
	// system acting as a container server.
	ContainerServerHostingRole HostingRole = "ContainerServer"
)

type InterfaceType string

const (
	// TPM12InterfaceType Trusted Platform Module (TPM) 1.2.
	TPM12InterfaceType InterfaceType = "TPM1_2"
	// TPM20InterfaceType Trusted Platform Module (TPM) 2.0.
	TPM20InterfaceType InterfaceType = "TPM2_0"
	// TCM10InterfaceType Trusted Cryptography Module (TCM) 1.0.
	TCM10InterfaceType InterfaceType = "TCM1_0"
)

// InterfaceTypeSelection is The enumerations of 'InterfaceTypeSelection'
// specify the method for switching the TrustedModule InterfaceType, for
// instance between TPM1_2 and TPM2_0, if supported.
type InterfaceTypeSelection string

const (
	// NoneInterfaceTypeSelection The TrustedModule does not support switching the
	// InterfaceType.
	NoneInterfaceTypeSelection InterfaceTypeSelection = "None"
	// FirmwareUpdateInterfaceTypeSelection The TrustedModule supports switching
	// InterfaceType through a firmware update.
	FirmwareUpdateInterfaceTypeSelection InterfaceTypeSelection = "FirmwareUpdate"
	// BiosSettingInterfaceTypeSelection The TrustedModule supports switching
	// InterfaceType through platform software, such as a BIOS configuration
	// attribute.
	BiosSettingInterfaceTypeSelection InterfaceTypeSelection = "BiosSetting"
	// OemMethodInterfaceTypeSelection The TrustedModule supports switching
	// InterfaceType through an OEM proprietary mechanism.
	OemMethodInterfaceTypeSelection InterfaceTypeSelection = "OemMethod"
)

type KMIPCachePolicy string

const (
	// NoneKMIPCachePolicy The system does not cache KMIP data.
	NoneKMIPCachePolicy KMIPCachePolicy = "None"
	// AfterFirstUseKMIPCachePolicy The system caches KMIP data after first use for
	// the duration specified by the 'CacheDuration' property.
	AfterFirstUseKMIPCachePolicy KMIPCachePolicy = "AfterFirstUse"
)

type LastResetCauses string

const (
	// PowerButtonPressLastResetCauses shall indicate the system start or restart
	// was due to a power button press.
	PowerButtonPressLastResetCauses LastResetCauses = "PowerButtonPress"
	// ManagementCommandLastResetCauses shall indicate the system start or restart
	// was due to an external command to the management controller or BMC. Examples
	// include the Redfish 'Reset' action for the 'ComputerSystem' resource or the
	// IPMI 'Chassis Control' command.
	ManagementCommandLastResetCauses LastResetCauses = "ManagementCommand"
	// PowerRestorePolicyLastResetCauses shall indicate the system automatically
	// powered-up on AC being applied due the 'PowerRestorePolicy' property
	// containing 'AlwaysOn' or 'LastState'.
	PowerRestorePolicyLastResetCauses LastResetCauses = "PowerRestorePolicy"
	// RTCWakeupLastResetCauses shall indicate the system powered-up via an RTC
	// (system real time clock) wakeup.
	RTCWakeupLastResetCauses LastResetCauses = "RTCWakeup"
	// WatchdogExpirationLastResetCauses shall indicate the system start or restart
	// was caused by a watchdog expiration.
	WatchdogExpirationLastResetCauses LastResetCauses = "WatchdogExpiration"
	// OSSoftRestartLastResetCauses shall indicate the system start or restart was
	// due to an OS soft restart. Examples include 'CTRL-ALT-DEL', 'init 6', or
	// 'reboot'.
	OSSoftRestartLastResetCauses LastResetCauses = "OSSoftRestart"
	// SystemCrashLastResetCauses shall indicate the system start or restart was
	// caused by a system crash. Examples include an OS panic, hardware fault, or
	// firmware fault.
	SystemCrashLastResetCauses LastResetCauses = "SystemCrash"
	// ThermalEventLastResetCauses shall indicate the system start or restart was
	// caused by a thermal event triggering a system shutdown.
	ThermalEventLastResetCauses LastResetCauses = "ThermalEvent"
	// PowerEventLastResetCauses shall indicate the system start or restart was
	// caused by a power event triggering a system shutdown.
	PowerEventLastResetCauses LastResetCauses = "PowerEvent"
	// UnknownLastResetCauses shall indicate the system start or restart cause is
	// unknown.
	UnknownLastResetCauses LastResetCauses = "Unknown"
)

type MemoryMirroring string

const (
	// SystemMemoryMirroring The system supports DIMM mirroring at the system
	// level. Individual DIMMs are not paired for mirroring in this mode.
	SystemMemoryMirroring MemoryMirroring = "System"
	// DIMMMemoryMirroring The system supports DIMM mirroring at the DIMM level.
	// Individual DIMMs can be mirrored.
	DIMMMemoryMirroring MemoryMirroring = "DIMM"
	// HybridMemoryMirroring The system supports a hybrid mirroring at the system
	// and DIMM levels. Individual DIMMs can be mirrored.
	HybridMemoryMirroring MemoryMirroring = "Hybrid"
	// NoneMemoryMirroring The system does not support DIMM mirroring.
	NoneMemoryMirroring MemoryMirroring = "None"
)

type PowerMode string

const (
	// MaximumPerformancePowerMode shall indicate the system performs at the
	// highest speeds possible. This mode should be used when performance is the
	// top priority.
	MaximumPerformancePowerMode PowerMode = "MaximumPerformance"
	// BalancedPerformancePowerMode shall indicate the system performs at the
	// highest speeds possible when the utilization is high and performs at reduced
	// speeds when the utilization is low to save power. This mode is a compromise
	// between 'MaximumPerformance' and 'PowerSaving'.
	BalancedPerformancePowerMode PowerMode = "BalancedPerformance"
	// PowerSavingPowerMode shall indicate the system performs at reduced speeds to
	// save power. This mode should be used when power saving is the top priority.
	PowerSavingPowerMode PowerMode = "PowerSaving"
	// StaticPowerMode shall indicate the system performs at a static base speed.
	StaticPowerMode PowerMode = "Static"
	// OSControlledPowerMode shall indicate the system performs at an operating
	// system-controlled power mode.
	OSControlledPowerMode PowerMode = "OSControlled"
	// OEMPowerMode shall indicate the system performs at an OEM-defined power
	// mode.
	OEMPowerMode PowerMode = "OEM"
	// EfficiencyFavorPowerPowerMode shall indicate the system performs at reduced
	// speeds at all utilizations to save power at the cost of performance. This
	// mode differs from 'PowerSaving' in that more performance is retained and
	// less power is saved. This mode differs from 'EfficiencyFavorPerformance' in
	// that less performance is retained but more power is saved. This mode differs
	// from 'BalancedPerformance' in that power saving occurs at all utilizations.
	EfficiencyFavorPowerPowerMode PowerMode = "EfficiencyFavorPower"
	// EfficiencyFavorPerformancePowerMode shall indicate the system performs at
	// reduced speeds at all utilizations to save power while attempting to
	// maintain performance. This mode differs from 'EfficiencyFavorPower' in that
	// more performance is retained but less power is saved. This mode differs from
	// 'MaximumPerformance' in that power is saved at the cost of some performance.
	// This mode differs from 'BalancedPerformance' in that power saving occurs at
	// all utilizations.
	EfficiencyFavorPerformancePowerMode PowerMode = "EfficiencyFavorPerformance"
)

// PowerRestorePolicyTypes is The enumerations of 'PowerRestorePolicyTypes'
// specify the choice of power state for the system when power is applied.
type PowerRestorePolicyTypes string

const (
	// AlwaysOnPowerRestorePolicyTypes The system always powers on when power is
	// applied.
	AlwaysOnPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOn"
	// AlwaysOffPowerRestorePolicyTypes The system always remains powered off when
	// power is applied.
	AlwaysOffPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOff"
	// LastStatePowerRestorePolicyTypes The system returns to its last on or off
	// power state when power is applied.
	LastStatePowerRestorePolicyTypes PowerRestorePolicyTypes = "LastState"
)

type StopBootOnFault string

const (
	// NeverStopBootOnFault shall indicate the system will continue to attempt to
	// boot if a fault occurs.
	NeverStopBootOnFault StopBootOnFault = "Never"
	// AnyFaultStopBootOnFault shall indicate the system will stop the boot if a
	// fault occurs. This includes, but is not limited to, faults that affect
	// performance, fault tolerance, or capacity.
	AnyFaultStopBootOnFault StopBootOnFault = "AnyFault"
)

type SystemType string

const (
	// PhysicalSystemType is a 'SystemType' of 'Physical' typically represents the
	// hardware aspects of a system, such as a management controller.
	PhysicalSystemType SystemType = "Physical"
	// VirtualSystemType is a 'SystemType' of 'Virtual' typically represents a
	// system that is actually a virtual machine instance. Responses should contain
	// the 'ProcessorSummary' and 'MemorySummary' properties to show the processor
	// and memory resources allocated to the virtual machine.
	VirtualSystemType SystemType = "Virtual"
	// OSSystemType is a 'SystemType' of 'OS' typically represents an OS or
	// hypervisor view of the system.
	OSSystemType SystemType = "OS"
	// PhysicallyPartitionedSystemType is a 'SystemType' of 'PhysicallyPartitioned'
	// typically represents a single system constructed from one or more physical
	// systems through a firmware or hardware-based service.
	PhysicallyPartitionedSystemType SystemType = "PhysicallyPartitioned"
	// VirtuallyPartitionedSystemType is a 'SystemType' of 'VirtuallyPartitioned'
	// typically represents a single system constructed from one or more virtual
	// systems through a software-based service.
	VirtuallyPartitionedSystemType SystemType = "VirtuallyPartitioned"
	// ComposedSystemType is a 'SystemType' of 'Composed' typically represents a
	// single system constructed from disaggregated resources through the Redfish
	// composition service.
	ComposedSystemType SystemType = "Composed"
	// DPUSystemType is a 'SystemType' of 'DPU' typically represents a single
	// system that performs offload computation as a data processing unit, such as
	// a SmartNIC.
	DPUSystemType SystemType = "DPU"
)

type TrustedModuleRequiredToBoot string

const (
	// DisabledTrustedModuleRequiredToBoot shall indicate a Trusted Module is not
	// required to boot.
	DisabledTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Disabled"
	// RequiredTrustedModuleRequiredToBoot shall indicate a functioning Trusted
	// Module is required to boot.
	RequiredTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Required"
)

// WatchdogTimeoutActions is The enumerations of 'WatchdogTimeoutActions'
// specify the choice of action to take when the host watchdog timer reaches its
// timeout value.
type WatchdogTimeoutActions string

const (
	// NoneWatchdogTimeoutActions No action taken.
	NoneWatchdogTimeoutActions WatchdogTimeoutActions = "None"
	// ResetSystemWatchdogTimeoutActions Reset the system.
	ResetSystemWatchdogTimeoutActions WatchdogTimeoutActions = "ResetSystem"
	// PowerCycleWatchdogTimeoutActions Power cycle the system.
	PowerCycleWatchdogTimeoutActions WatchdogTimeoutActions = "PowerCycle"
	// PowerDownWatchdogTimeoutActions Power down the system.
	PowerDownWatchdogTimeoutActions WatchdogTimeoutActions = "PowerDown"
	// OEMWatchdogTimeoutActions Perform an OEM-defined action.
	OEMWatchdogTimeoutActions WatchdogTimeoutActions = "OEM"
)

// WatchdogWarningActions is The enumerations of 'WatchdogWarningActions'
// specify the choice of action to take when the host watchdog timer is close
// (typically 3-10 seconds) to reaching its timeout value.
type WatchdogWarningActions string

const (
	// NoneWatchdogWarningActions No action taken.
	NoneWatchdogWarningActions WatchdogWarningActions = "None"
	// DiagnosticInterruptWatchdogWarningActions Raise a (typically non-maskable)
	// Diagnostic Interrupt.
	DiagnosticInterruptWatchdogWarningActions WatchdogWarningActions = "DiagnosticInterrupt"
	// SMIWatchdogWarningActions Raise a Systems Management Interrupt (SMI).
	SMIWatchdogWarningActions WatchdogWarningActions = "SMI"
	// MessagingInterruptWatchdogWarningActions Raise a legacy IPMI messaging
	// interrupt.
	MessagingInterruptWatchdogWarningActions WatchdogWarningActions = "MessagingInterrupt"
	// SCIWatchdogWarningActions Raise an interrupt using the ACPI System Control
	// Interrupt (SCI).
	SCIWatchdogWarningActions WatchdogWarningActions = "SCI"
	// OEMWatchdogWarningActions Perform an OEM-defined action.
	OEMWatchdogWarningActions WatchdogWarningActions = "OEM"
)

// ComputerSystem shall represent a computing system in the Redfish
// Specification.
type ComputerSystem struct {
	Entity
	// AssetTag shall contain the system asset tag value. Modifying this property
	// may modify the 'AssetTag' in the containing 'Chassis' resource.
	AssetTag string
	// Bios shall contain a link to a resource of type 'Bios' that lists the BIOS
	// settings for this system.
	//
	// Version added: v1.1.0
	bios string
	// BiosVersion shall contain the version string of the currently installed and
	// running BIOS for x86 systems. For other systems, the property may contain a
	// version string that represents the primary system firmware.
	BiosVersion string
	// Boot shall contain the boot settings for this system.
	Boot Boot
	// BootProgress shall contain the last boot progress state and time.
	//
	// Version added: v1.13.0
	BootProgress BootProgress
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.14.0
	certificates string
	// Composition shall contain information about the composition capabilities and
	// state of the computer system.
	//
	// Version added: v1.18.0
	Composition Composition
	// EthernetInterfaces shall contain a link to a resource collection of type
	// 'EthernetInterfaceCollection'.
	ethernetInterfaces string
	// FabricAdapters shall contain a link to a resource collection of type
	// 'FabricAdapterCollection'.
	//
	// Version added: v1.10.0
	fabricAdapters string
	// GraphicalConsole shall contain the information about the graphical console
	// (KVM-IP) service of this system.
	//
	// Version added: v1.13.0
	GraphicalConsole HostGraphicalConsole
	// GraphicsControllers shall contain a link to a resource collection of type
	// 'GraphicsControllerCollection' that contains graphics controllers that can
	// output video for this system.
	//
	// Version added: v1.15.0
	graphicsControllers string
	// HostName shall contain the host name for this system, as reported by the
	// operating system or hypervisor. A service running in the host operating
	// system typically reports this value to the manager. Modifying this property
	// may modify the 'HostName' in one or more 'EthernetInterface' resources
	// contained in this system.
	HostName string
	// HostWatchdogTimer shall contain properties that describe the host watchdog
	// timer functionality for this system.
	//
	// Version added: v1.5.0
	HostWatchdogTimer WatchdogTimer
	// HostedServices shall describe services that this computer system supports.
	//
	// Version added: v1.2.0
	HostedServices HostedServices
	// HostingRoles shall contain the hosting roles that this computer system
	// supports.
	//
	// Version added: v1.2.0
	HostingRoles []HostingRole
	// IPMIHostInterface shall contain the information about the in-band IPMI
	// service of this system.
	//
	// Version added: v1.25.0
	IPMIHostInterface IPMIHostInterface
	// IdlePowerSaver shall contain the idle power saver settings of the computer
	// system.
	//
	// Version added: v1.16.0
	IdlePowerSaver IdlePowerSaver
	// IndicatorLED shall contain the state of the indicator light, which
	// identifies this system.
	//
	// Deprecated: v1.13.0
	// This property has been deprecated in favor of the 'LocationIndicatorActive'
	// property.
	IndicatorLED IndicatorLED
	// KeyManagement shall contain the key management settings of the computer
	// system.
	//
	// Version added: v1.16.0
	KeyManagement KeyManagement
	// LastResetCause shall contain the cause when the system last came out of a
	// reset or was rebooted.
	//
	// Version added: v1.23.0
	LastResetCause LastResetCauses
	// LastResetTime shall contain the date and time when the system last came out
	// of a reset or was rebooted.
	//
	// Version added: v1.12.0
	LastResetTime string
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function. Modifying this property
	// may modify the 'LocationIndicatorActive' in the containing 'Chassis'
	// resource.
	//
	// Version added: v1.13.0
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type
	// 'LogServiceCollection'.
	logServices string
	// Manufacturer shall contain a value that represents the manufacturer of the
	// system.
	Manufacturer string
	// ManufacturingMode shall indicate whether the system is in manufacturing
	// mode. If the system supports SMBIOS, the value shall match the
	// 'Manufacturing mode is enabled' setting from the 'BIOS Characteristics'
	// entry.
	//
	// Version added: v1.18.0
	ManufacturingMode bool
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.14.0
	//
	// Deprecated: v1.17.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// Memory shall contain a link to a resource collection of type
	// 'MemoryCollection'.
	//
	// Version added: v1.1.0
	memory string
	// MemoryDomains shall contain a link to a resource collection of type
	// 'MemoryDomainCollection'.
	//
	// Version added: v1.2.0
	memoryDomains string
	// MemorySummary shall describe the central memory for this resource.
	MemorySummary MemorySummary
	// Model shall describe how the manufacturer refers to this system. Typically,
	// this value is the product name for this system without the manufacturer
	// name.
	Model string
	// MultipartImportConfigurationPushURI shall contain a URI used to perform a
	// multipart HTTP or HTTPS 'POST' of a vendor-specific configuration file for
	// the purpose of importing the configuration contained within the file as
	// defined by the 'Import configuration data' clause of the Redfish
	// Specification. The value of this property should not contain a URI of a
	// Redfish resource. See the 'Redfish-defined URIs and relative reference
	// rules' clause in the Redfish Specification.
	//
	// Version added: v1.26.0
	MultipartImportConfigurationPushURI string
	// NetworkInterfaces shall contain a link to a resource collection of type
	// 'NetworkInterfaceCollection'.
	//
	// Version added: v1.3.0
	networkInterfaces string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSystem shall contain a link to a resource of type 'OperatingSystem'
	// that contains operating system information for this system.
	//
	// Version added: v1.21.0
	operatingSystem string
	// PCIeDevices shall contain an array of links to resources of type
	// 'PCIeDevice'.
	//
	// Version added: v1.2.0
	pCIeDevices []string
	// PCIeDevicesCount
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// PCIeFunctions shall contain an array of links to resources of type
	// 'PCIeFunction'.
	//
	// Version added: v1.2.0
	pCIeFunctions []string
	// PCIeFunctionsCount
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// PartNumber shall contain the manufacturer-defined part number for the
	// system.
	PartNumber string
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on
	// after a 'Reset' action requesting 'PowerCycle' or 'FullPowerCycle'. The
	// value '0' shall indicate no delay to power on.
	//
	// Version added: v1.13.0
	PowerCycleDelaySeconds *float64 `json:",omitempty"`
	// PowerMode shall contain the computer system power mode setting.
	//
	// Version added: v1.15.0
	PowerMode PowerMode
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off
	// during a reset. The value '0' shall indicate no delay to power off.
	//
	// Version added: v1.13.0
	PowerOffDelaySeconds *float64 `json:",omitempty"`
	// PowerOnDelaySeconds shall contain the number of seconds to delay power on
	// after an externally performed power cycle or during a reset. The value '0'
	// shall indicate no delay to power on.
	//
	// Version added: v1.13.0
	PowerOnDelaySeconds *float64 `json:",omitempty"`
	// PowerRestorePolicy shall indicate the desired power state of the system when
	// power is applied to the system. The 'LastState' value shall return the
	// system to the 'PowerState' property value it was in when power was lost.
	//
	// Version added: v1.6.0
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the system.
	PowerState PowerState
	// ProcessorSummary shall describe the central processors for this resource.
	// Processors described by this property shall be limited to the processors
	// that execute system code, and shall not include processors used for offload
	// functionality.
	ProcessorSummary ProcessorSummary
	// Processors shall contain a link to a resource collection of type
	// 'ProcessorCollection'.
	processors string
	// Redundancy shall contain a set of redundancy entities. Each entity specifies
	// a kind and level of redundancy and a collection, or redundancy set, of other
	// computer systems that provide the specified redundancy to this computer
	// system.
	//
	// Version added: v1.5.0
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SKU shall contain the SKU for the system.
	SKU string
	// SecureBoot shall contain a link to a resource of type 'SecureBoot'.
	//
	// Version added: v1.1.0
	secureBoot string
	// SerialConsole shall contain information about the serial console services of
	// this system.
	//
	// Version added: v1.13.0
	SerialConsole HostSerialConsole
	// SerialNumber shall contain the serial number for the system.
	SerialNumber string
	// SimpleStorage shall contain a link to a resource collection of type
	// 'SimpleStorageCollection'.
	simpleStorage string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Storage shall contain a link to a resource collection of type
	// 'StorageCollection'.
	//
	// Version added: v1.1.0
	storage string
	// SubModel shall contain the information about the sub-model (or
	// configuration) of the system. This shall not include the model/product name
	// or the manufacturer name.
	//
	// Version added: v1.5.0
	SubModel string
	// SystemType is an enumeration that indicates the kind of system that this
	// resource represents.
	SystemType SystemType
	// TrustedModules shall contain an array of objects with properties that
	// describe the trusted modules for this resource.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.19.0
	// This property has been deprecated in favor of the 'TrustedComponents'
	// property in 'Links'.
	TrustedModules []TrustedModules
	// USBControllers shall contain a link to a resource collection of type
	// 'USBControllerCollection' that contains USB controllers for this system.
	//
	// Version added: v1.15.0
	uSBControllers string
	// UUID shall contain the universally unique identifier number for this system.
	// RFC4122 describes methods to create this value. The value should be
	// considered to be opaque. Client software should only treat the overall value
	// as a UUID and should not interpret any subfields within the UUID. If the
	// system supports SMBIOS, the property value should follow the SMBIOS 2.6 and
	// later recommendation for converting the SMBIOS 16-byte UUID structure into
	// the Redfish canonical 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' string format,
	// so that the property value matches the byte order presented by current OS
	// APIs, such as WMI and dmidecode. For additional property requirements, see
	// the corresponding definition in the Redfish Data Model Specification.
	UUID string
	// VirtualMedia shall contain a link to a resource collection of type
	// 'VirtualMediaCollection' that this system uses.
	//
	// Version added: v1.13.0
	virtualMedia string
	// VirtualMediaConfig shall contain the information about the virtual media
	// service of this system.
	//
	// Version added: v1.13.0
	VirtualMediaConfig ComputerVirtualMediaConfig
	// addResourceBlockTarget is the URL to send AddResourceBlock requests.
	addResourceBlockTarget string
	// decommissionTarget is the URL to send Decommission requests.
	decommissionTarget string
	// exportConfigurationTarget is the URL to send ExportConfiguration requests.
	exportConfigurationTarget string
	// removeResourceBlockTarget is the URL to send RemoveResourceBlock requests.
	removeResourceBlockTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// setDefaultBootOrderTarget is the URL to send SetDefaultBootOrder requests.
	setDefaultBootOrderTarget string
	// chassis are the URIs for Chassis.
	chassis []string
	// consumingComputerSystems are the URIs for ConsumingComputerSystems.
	consumingComputerSystems []string
	// cooledBy are the URIs for CooledBy.
	cooledBy []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// hostingComputerSystem is the URI for HostingComputerSystem.
	hostingComputerSystem string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// offloadedNetworkDeviceFunctions are the URIs for OffloadedNetworkDeviceFunctions.
	offloadedNetworkDeviceFunctions []string
	// poweredBy are the URIs for PoweredBy.
	poweredBy []string
	// resourceBlocks are the URIs for ResourceBlocks.
	resourceBlocks []string
	// supplyingComputerSystems are the URIs for SupplyingComputerSystems.
	supplyingComputerSystems []string
	// trustedComponents are the URIs for TrustedComponents.
	trustedComponents []string
	// virtualMachines are the URIs for VirtualMachines.
	virtualMachines []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte

	// resetActionInfoTarget is the URL to check what values are supported
	resetActionInfoTarget string
	// settingsApplyTimes is a set of allowed settings update apply times. If none
	// are specified, then the system does not provide that information.
	SettingsApplyTimes []SettingsApplyTime
	settingsTarget     string
	// SupportedResetTypes, if provided, is the reset types this system supports.
	supportedResetTypes []ResetType
}

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
//
//nolint:funlen
func (c *ComputerSystem) UnmarshalJSON(b []byte) error {
	type temp ComputerSystem
	type cActions struct {
		AddResourceBlock    ActionTarget `json:"#ComputerSystem.AddResourceBlock"`
		Decommission        ActionTarget `json:"#ComputerSystem.Decommission"`
		ExportConfiguration ActionTarget `json:"#ComputerSystem.ExportConfiguration"`
		RemoveResourceBlock ActionTarget `json:"#ComputerSystem.RemoveResourceBlock"`
		Reset               struct {
			ActionTarget
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#ComputerSystem.Reset"`
		SetDefaultBootOrder ActionTarget `json:"#ComputerSystem.SetDefaultBootOrder"`
	}
	type cLinks struct {
		Chassis                         Links `json:"Chassis"`
		ConsumingComputerSystems        Links `json:"ConsumingComputerSystems"`
		CooledBy                        Links `json:"CooledBy"`
		Endpoints                       Links `json:"Endpoints"`
		HostingComputerSystem           Link  `json:"HostingComputerSystem"`
		ManagedBy                       Links `json:"ManagedBy"`
		OffloadedNetworkDeviceFunctions Links `json:"OffloadedNetworkDeviceFunctions"`
		PoweredBy                       Links `json:"PoweredBy"`
		ResourceBlocks                  Links `json:"ResourceBlocks"`
		SupplyingComputerSystems        Links `json:"SupplyingComputerSystems"`
		TrustedComponents               Links `json:"TrustedComponents"`
		VirtualMachines                 Links `json:"VirtualMachines"`
	}
	var tmp struct {
		temp
		Actions             cActions
		Links               cLinks
		Bios                Link     `json:"Bios"`
		Certificates        Link     `json:"Certificates"`
		EthernetInterfaces  Link     `json:"EthernetInterfaces"`
		FabricAdapters      Link     `json:"FabricAdapters"`
		GraphicsControllers Link     `json:"GraphicsControllers"`
		LogServices         Link     `json:"LogServices"`
		Memory              Link     `json:"Memory"`
		MemoryDomains       Link     `json:"MemoryDomains"`
		NetworkInterfaces   Link     `json:"NetworkInterfaces"`
		OperatingSystem     Link     `json:"OperatingSystem"`
		PCIeDevices         Links    `json:"PCIeDevices"`
		PCIeFunctions       Links    `json:"PCIeFunctions"`
		Processors          Link     `json:"Processors"`
		Redundancy          Link     `json:"Redundancy"`
		SecureBoot          Link     `json:"SecureBoot"`
		SimpleStorage       Link     `json:"SimpleStorage"`
		Storage             Link     `json:"Storage"`
		USBControllers      Link     `json:"USBControllers"`
		VirtualMedia        Link     `json:"VirtualMedia"`
		Settings            Settings `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ComputerSystem(tmp.temp)

	// Extract the links to other entities for later
	c.addResourceBlockTarget = tmp.Actions.AddResourceBlock.Target
	c.decommissionTarget = tmp.Actions.Decommission.Target
	c.exportConfigurationTarget = tmp.Actions.ExportConfiguration.Target
	c.removeResourceBlockTarget = tmp.Actions.RemoveResourceBlock.Target
	c.resetTarget = tmp.Actions.Reset.Target
	c.resetActionInfoTarget = tmp.Actions.Reset.ActionInfoTarget
	c.supportedResetTypes = tmp.Actions.Reset.AllowedResetTypes

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	c.settingsTarget = tmp.Settings.SettingsObject
	if c.settingsTarget == "" {
		c.settingsTarget = c.ODataID
	}
	c.SettingsApplyTimes = tmp.Settings.SupportedApplyTimes

	c.setDefaultBootOrderTarget = tmp.Actions.SetDefaultBootOrder.Target
	c.chassis = tmp.Links.Chassis.ToStrings()
	c.consumingComputerSystems = tmp.Links.ConsumingComputerSystems.ToStrings()
	c.cooledBy = tmp.Links.CooledBy.ToStrings()
	c.endpoints = tmp.Links.Endpoints.ToStrings()
	c.hostingComputerSystem = tmp.Links.HostingComputerSystem.String()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.offloadedNetworkDeviceFunctions = tmp.Links.OffloadedNetworkDeviceFunctions.ToStrings()
	c.poweredBy = tmp.Links.PoweredBy.ToStrings()
	c.resourceBlocks = tmp.Links.ResourceBlocks.ToStrings()
	c.supplyingComputerSystems = tmp.Links.SupplyingComputerSystems.ToStrings()
	c.trustedComponents = tmp.Links.TrustedComponents.ToStrings()
	c.virtualMachines = tmp.Links.VirtualMachines.ToStrings()
	c.bios = tmp.Bios.String()
	c.certificates = tmp.Certificates.String()
	c.ethernetInterfaces = tmp.EthernetInterfaces.String()
	c.fabricAdapters = tmp.FabricAdapters.String()
	c.graphicsControllers = tmp.GraphicsControllers.String()
	c.logServices = tmp.LogServices.String()
	c.memory = tmp.Memory.String()
	c.memoryDomains = tmp.MemoryDomains.String()
	c.networkInterfaces = tmp.NetworkInterfaces.String()
	c.operatingSystem = tmp.OperatingSystem.String()
	c.pCIeDevices = tmp.PCIeDevices.ToStrings()
	c.pCIeFunctions = tmp.PCIeFunctions.ToStrings()
	c.processors = tmp.Processors.String()
	c.redundancy = tmp.Redundancy.String()
	c.secureBoot = tmp.SecureBoot.String()
	c.simpleStorage = tmp.SimpleStorage.String()
	c.storage = tmp.Storage.String()
	c.uSBControllers = tmp.USBControllers.String()
	c.virtualMedia = tmp.VirtualMedia.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ComputerSystem) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"HostName",
		"IndicatorLED",
		"LocationIndicatorActive",
		"PowerCycleDelaySeconds",
		"PowerMode",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestorePolicy",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetComputerSystem will get a ComputerSystem instance from the service.
func GetComputerSystem(c Client, uri string) (*ComputerSystem, error) {
	return GetObject[ComputerSystem](c, uri)
}

// ListReferencedComputerSystems gets the collection of ComputerSystem from
// a provided reference.
func ListReferencedComputerSystems(c Client, link string) ([]*ComputerSystem, error) {
	return GetCollectionObjects[ComputerSystem](c, link)
}

// SetBoot sets a boot object based on a payload request
func (c *ComputerSystem) SetBoot(b *Boot) error {
	t := struct {
		Boot *Boot
	}{Boot: b}
	return c.Patch(c.ODataID, t)
}

// This action shall add a resource block to a system.
// computerSystemETag - This parameter shall contain the current ETag of the
// system. If the client-provided ETag does not match the current ETag of the
// system, the service shall return the HTTP '428 Precondition Required' status
// code to reject the request.
// resourceBlock - This parameter shall contain a link to the specified
// resource block to add to the system.
// resourceBlockETag - This parameter shall contain the current ETag of the
// resource block to add to the system. If the client-provided ETag does not
// match the current ETag of the resource block that the 'ResourceBlock'
// parameter specifies, the service shall return the HTTP '428 Precondition
// Required' status code to reject the request.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) AddResourceBlock(computerSystemETag string, resourceBlock string, resourceBlockETag string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ComputerSystemETag"] = computerSystemETag
	payload["ResourceBlock"] = resourceBlock
	payload["ResourceBlockETag"] = resourceBlockETag
	resp, taskInfo, err := PostWithTask(c.client,
		c.addResourceBlockTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ComputerSystemDecommissionParameters holds the parameters for the Decommission action.
type ComputerSystemDecommissionParameters struct {
	// ComputerSystemETag shall contain the current ETag of the system. If the
	// client-provided ETag does not match the current ETag of the system, the
	// service shall return the HTTP '428 Precondition Required' status code to
	// reject the request.
	ComputerSystemETag string `json:"ComputerSystemETag,omitempty"`
	// DecommissionTypes shall contain a list of the types of data to remove from
	// the system.
	DecommissionTypes []DecommissionType `json:"DecommissionTypes,omitempty"`
	// OEMDecommissionTypes shall contain any OEM-specific types of data to remove
	// from the system.
	OEMDecommissionTypes []string `json:"OEMDecommissionTypes,omitempty"`
	// RequireSecureErase shall indicate if a secure erase is required. If the
	// parameter contains 'true' and a secure erase to the level of NIST 800-88
	// Clear or Purge for all specified components cannot be performed the service
	// shall return the HTTP '501 Not Implemented' status code. This failure may
	// occur after the process has already started. If not provided by the client,
	// the value shall be assumed to be 'false'.
	RequireSecureErase bool `json:"RequireSecureErase,omitempty"`
}

// This action shall remove all specified data from a system in preparation to
// decommission the system.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) Decommission(params *ComputerSystemDecommissionParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(c.client,
		c.decommissionTarget, params, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ComputerSystemExportConfigurationParameters holds the parameters for the ExportConfiguration action.
type ComputerSystemExportConfigurationParameters struct {
	// Components shall contain an array of components of the system for which to
	// export configuration data.
	Components []Component `json:"Components,omitempty"`
	// EncryptionPassphrase shall contain the encryption passphrase for the
	// exported file. If this parameter is specified and has a non-zero length, the
	// service shall encrypt the exported file with the passphrase. Otherwise, the
	// service shall not encrypt the exported file.
	EncryptionPassphrase string `json:"EncryptionPassphrase,omitempty"`
	// ExportType shall contain the type of export to perform.
	ExportType ExportType `json:"ExportType,omitempty"`
	// OEMComponents shall contain an array of OEM-specific components of the
	// system for which to export configuration data.
	OEMComponents []string `json:"OEMComponents,omitempty"`
	// Security shall contain the policy to apply when exporting secure
	// information.
	Security ExportSecurity `json:"Security,omitempty"`
}

// This action shall export the specified configuration of a system in a
// vendor-specific format. Upon successful completion of the action and any
// asynchronous processing, the 'Location' header in the response shall contain
// a URI to a file that contains the configuration data.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) ExportConfiguration(params *ComputerSystemExportConfigurationParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(c.client,
		c.exportConfigurationTarget, params, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall remove a resource block from a system.
// computerSystemETag - This parameter shall contain the current ETag of the
// system. If the client-provided ETag does not match the current ETag of the
// system, the service shall return the HTTP '428 Precondition Required' status
// code to reject the request.
// resourceBlock - This parameter shall contain a link to the specified
// resource block to remove from the system.
// resourceBlockETag - This parameter shall contain the current ETag of the
// resource block to remove from the system. If the client-provided ETag does
// not match the current ETag of the resource block that the 'ResourceBlock'
// parameter specifies, the service shall return the HTTP '428 Precondition
// Required' status code to reject the request.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) RemoveResourceBlock(computerSystemETag string, resourceBlock string, resourceBlockETag string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ComputerSystemETag"] = computerSystemETag
	payload["ResourceBlock"] = resourceBlock
	payload["ResourceBlockETag"] = resourceBlockETag
	resp, taskInfo, err := PostWithTask(c.client,
		c.removeResourceBlockTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the system represented by the resource. For systems
// that implement ACPI Power Button functionality, the 'PushPowerButton' value
// shall perform or emulate an ACPI Power Button Push, and the 'ForceOff' value
// shall perform an ACPI Power Button Override, commonly known as a four-second
// hold of the power button.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	// Make sure the requested reset type is supported by the system
	valid := false
	if len(c.supportedResetTypes) > 0 {
		for _, allowed := range c.supportedResetTypes {
			if resetType == allowed {
				valid = true
				break
			}
		}
	} else {
		// No allowed values supplied, assume we are OK
		valid = true
	}

	if !valid {
		return nil, fmt.Errorf("reset type '%s' is not supported by this service",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetTarget, t, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// GetSupportedResetTypes returns any reset types that the ComputerSystem declares as supported
// via either ActionInfo or AllowableValues.
func (c *ComputerSystem) GetSupportedResetTypes() ([]ResetType, error) {
	if len(c.supportedResetTypes) > 0 {
		return c.supportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if c.resetActionInfoTarget != "" {
		resetActionInfo, err := c.ResetActionInfo()
		if err != nil {
			return c.supportedResetTypes, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return c.supportedResetTypes, err
		}

		for _, val := range vals {
			c.supportedResetTypes = append(c.supportedResetTypes, ResetType(val))
		}
	}

	return c.supportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the ComputerSystem reset action if supported
func (c *ComputerSystem) ResetActionInfo() (*ActionInfo, error) {
	if c.resetActionInfoTarget == "" {
		return nil, errors.New("ComputerSystem Reset ActionInfo not supported")
	}

	return GetObject[ActionInfo](c.GetClient(), c.resetActionInfoTarget)
}

// UpdateBootAttributesApplyAt is used to update attribute values and set apply time together
func (c *ComputerSystem) UpdateBootAttributesApplyAt(attrs SettingsAttributes, applyTime SettingsApplyTime) error {
	payload := make(map[string]any)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(c.RawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := c.GetClient().Get(c.settingsTarget)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]any{"Boot": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = c.GetClient().PatchWithHeaders(c.settingsTarget, data, header)
		defer DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBootAttributes is used to update attribute values.
func (c *ComputerSystem) UpdateBootAttributes(attrs SettingsAttributes) error {
	return c.UpdateBootAttributesApplyAt(attrs, "")
}

// This action shall set the 'BootOrder' array to the default settings.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) SetDefaultBootOrder() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(c.client,
		c.setDefaultBootOrderTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Chassis gets the Chassis linked resources.
func (c *ComputerSystem) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.chassis)
}

// ConsumingComputerSystems gets the ConsumingComputerSystems linked resources.
func (c *ComputerSystem) ConsumingComputerSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](c.client, c.consumingComputerSystems)
}

// CooledBy gets the CooledBy linked resources.
func (c *ComputerSystem) CooledBy() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.cooledBy)
}

// Endpoints gets the Endpoints linked resources.
func (c *ComputerSystem) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](c.client, c.endpoints)
}

// HostingComputerSystem gets the HostingComputerSystem linked resource.
func (c *ComputerSystem) HostingComputerSystem() (*ComputerSystem, error) {
	if c.hostingComputerSystem == "" {
		return nil, nil
	}
	return GetObject[ComputerSystem](c.client, c.hostingComputerSystem)
}

// ManagedBy gets the ManagedBy linked resources.
func (c *ComputerSystem) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](c.client, c.managedBy)
}

// OffloadedNetworkDeviceFunctions gets the OffloadedNetworkDeviceFunctions linked resources.
func (c *ComputerSystem) OffloadedNetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](c.client, c.offloadedNetworkDeviceFunctions)
}

// PoweredBy gets the PoweredBy linked resources.
func (c *ComputerSystem) PoweredBy() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.poweredBy)
}

// ResourceBlocks gets the ResourceBlocks linked resources.
func (c *ComputerSystem) ResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](c.client, c.resourceBlocks)
}

// SupplyingComputerSystems gets the SupplyingComputerSystems linked resources.
func (c *ComputerSystem) SupplyingComputerSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](c.client, c.supplyingComputerSystems)
}

// TrustedComponents gets the TrustedComponents linked resources.
func (c *ComputerSystem) TrustedComponents() ([]*TrustedComponent, error) {
	return GetObjects[TrustedComponent](c.client, c.trustedComponents)
}

// VirtualMachines gets the VirtualMachines linked resources.
func (c *ComputerSystem) VirtualMachines() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](c.client, c.virtualMachines)
}

// Bios gets the Bios linked resource.
func (c *ComputerSystem) Bios() (*Bios, error) {
	if c.bios == "" {
		return nil, nil
	}
	return GetObject[Bios](c.client, c.bios)
}

// BootOptions gets all BootOption items for this system.
func (c *ComputerSystem) BootOptions() ([]*BootOption, error) {
	return GetCollectionObjects[BootOption](
		c.GetClient(),
		c.Boot.bootOptions)
}

// Certificates gets the Certificates collection.
func (c *ComputerSystem) Certificates() ([]*Certificate, error) {
	if c.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](c.client, c.certificates)
}

// EthernetInterfaces gets the EthernetInterfaces collection.
func (c *ComputerSystem) EthernetInterfaces() ([]*EthernetInterface, error) {
	if c.ethernetInterfaces == "" {
		return nil, nil
	}
	return GetCollectionObjects[EthernetInterface](c.client, c.ethernetInterfaces)
}

// FabricAdapters gets the FabricAdapters collection.
func (c *ComputerSystem) FabricAdapters() ([]*FabricAdapter, error) {
	if c.fabricAdapters == "" {
		return nil, nil
	}
	return GetCollectionObjects[FabricAdapter](c.client, c.fabricAdapters)
}

// GraphicsControllers gets the GraphicsControllers collection.
func (c *ComputerSystem) GraphicsControllers() ([]*GraphicsController, error) {
	if c.graphicsControllers == "" {
		return nil, nil
	}
	return GetCollectionObjects[GraphicsController](c.client, c.graphicsControllers)
}

// LogServices gets the LogServices collection.
func (c *ComputerSystem) LogServices() ([]*LogService, error) {
	if c.logServices == "" {
		return nil, nil
	}
	return GetCollectionObjects[LogService](c.client, c.logServices)
}

// Memory gets the Memory collection.
func (c *ComputerSystem) Memory() ([]*Memory, error) {
	if c.memory == "" {
		return nil, nil
	}
	return GetCollectionObjects[Memory](c.client, c.memory)
}

// MemoryDomains gets the MemoryDomains collection.
func (c *ComputerSystem) MemoryDomains() ([]*MemoryDomain, error) {
	if c.memoryDomains == "" {
		return nil, nil
	}
	return GetCollectionObjects[MemoryDomain](c.client, c.memoryDomains)
}

// NetworkInterfaces gets the NetworkInterfaces collection.
func (c *ComputerSystem) NetworkInterfaces() ([]*NetworkInterface, error) {
	if c.networkInterfaces == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkInterface](c.client, c.networkInterfaces)
}

// OperatingSystem gets the OperatingSystem linked resource.
func (c *ComputerSystem) OperatingSystem() (*OperatingSystem, error) {
	if c.operatingSystem == "" {
		return nil, nil
	}
	return GetObject[OperatingSystem](c.client, c.operatingSystem)
}

// PCIeDevices gets the PCIeDevices linked resources.
func (c *ComputerSystem) PCIeDevices() ([]*PCIeDevice, error) {
	return GetObjects[PCIeDevice](c.client, c.pCIeDevices)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (c *ComputerSystem) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](c.client, c.pCIeFunctions)
}

// Processors gets the Processors collection.
func (c *ComputerSystem) Processors() ([]*Processor, error) {
	if c.processors == "" {
		return nil, nil
	}
	return GetCollectionObjects[Processor](c.client, c.processors)
}

// Redundancy gets the Redundancy linked resource.
func (c *ComputerSystem) Redundancy() (*Redundancy, error) {
	if c.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](c.client, c.redundancy)
}

// SecureBoot gets the SecureBoot linked resource.
func (c *ComputerSystem) SecureBoot() (*SecureBoot, error) {
	if c.secureBoot == "" {
		return nil, nil
	}
	return GetObject[SecureBoot](c.client, c.secureBoot)
}

// SimpleStorage gets the SimpleStorage collection.
func (c *ComputerSystem) SimpleStorage() ([]*SimpleStorage, error) {
	if c.simpleStorage == "" {
		return nil, nil
	}
	return GetCollectionObjects[SimpleStorage](c.client, c.simpleStorage)
}

// Storage gets the Storage collection.
func (c *ComputerSystem) Storage() ([]*Storage, error) {
	if c.storage == "" {
		return nil, nil
	}
	return GetCollectionObjects[Storage](c.client, c.storage)
}

// USBControllers gets the USBControllers collection.
func (c *ComputerSystem) USBControllers() ([]*USBController, error) {
	if c.uSBControllers == "" {
		return nil, nil
	}
	return GetCollectionObjects[USBController](c.client, c.uSBControllers)
}

// VirtualMedia gets the VirtualMedia collection.
func (c *ComputerSystem) VirtualMedia() ([]*VirtualMedia, error) {
	if c.virtualMedia == "" {
		return nil, nil
	}
	return GetCollectionObjects[VirtualMedia](c.client, c.virtualMedia)
}

// Boot shall contain properties that describe boot information for a system.
type Boot struct {
	// AliasBootOrder shall contain an ordered array of boot source aliases of the
	// 'BootSource' type that represents the persistent boot order of this computer
	// system. This array shall not contain duplicate values. Virtual devices for
	// an alias should take precedence over a physical device. Systems may attempt
	// to boot from multiple devices that share an alias.
	//
	// Version added: v1.6.0
	AliasBootOrder []BootSource `json:",omitempty"`
	// AutomaticRetryAttempts shall contain the number of attempts the system will
	// automatically retry booting in the event the system enters an error state on
	// boot.
	//
	// Version added: v1.11.0
	AutomaticRetryAttempts *uint `json:",omitempty"`
	// AutomaticRetryConfig shall contain the configuration of how the system
	// retries booting automatically.
	//
	// Version added: v1.11.0
	AutomaticRetryConfig AutomaticRetryConfig `json:",omitempty"`
	// BootNext shall contain the 'BootOptionReference' of the UEFI boot option for
	// one time boot, as defined by the UEFI Specification. The valid values for
	// this property are specified in the values of the BootOrder array.
	// 'BootSourceOverrideEnabled' set to 'Continuous' is not supported for
	// 'BootSourceOverrideTarget' set to 'UefiBootNext' because this setting is
	// defined in UEFI as a one-time boot setting.
	//
	// Version added: v1.5.0
	BootNext string `json:",omitempty"`
	// BootOptions shall contain a link to a resource collection of type
	// 'BootOptionCollection'.
	//
	// Version added: v1.5.0
	bootOptions string
	// BootOrder shall contain an array of 'BootOptionReference' strings that
	// represent the persistent boot order for this computer system. For UEFI
	// systems, this is the UEFI Specification-defined UEFI BootOrder.
	//
	// Version added: v1.5.0
	BootOrder []string `json:",omitempty"`
	// BootOrderPropertySelection shall indicate which boot order property the
	// system uses for the persistent boot order.
	//
	// Version added: v1.6.0
	BootOrderPropertySelection BootOrderTypes `json:",omitempty"`
	// BootSourceOverrideEnabled shall contain 'Once' for a one-time boot override,
	// and 'Continuous' for a remain-active-until-cancelled override. If set to
	// 'Once', the value is reset to 'Disabled' after the
	// 'BootSourceOverrideTarget' actions have completed successfully. Changes to
	// this property do not alter the BIOS persistent boot order configuration.
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:",omitempty"`
	// BootSourceOverrideMode shall contain the BIOS boot mode to use when the
	// system boots from the 'BootSourceOverrideTarget' boot source.
	//
	// Version added: v1.1.0
	BootSourceOverrideMode BootSourceOverrideMode `json:",omitempty"`
	// BootSourceOverrideTarget shall contain the source to boot the system from,
	// overriding the normal boot order. The '@Redfish.AllowableValues' annotation
	// specifies the valid values for this property. 'UefiTarget' indicates to boot
	// from the UEFI device path found in 'UefiTargetBootSourceOverride'.
	// 'UefiBootNext' indicates to boot from the UEFI 'BootOptionReference' found
	// in 'BootNext'. Virtual devices for a target should take precedence over a
	// physical device. Systems may attempt to boot from multiple devices that
	// share a target identifier. Changes to this property do not alter the BIOS
	// persistent boot order configuration.
	BootSourceOverrideTarget                BootSource   `json:",omitempty"`
	AllowableBootSourceOverrideTargetValues []BootSource `json:"BootSourceOverrideTarget@Redfish.AllowableValues,omitempty"`
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection'.
	//
	// Version added: v1.7.0
	certificates string
	// HTTPBootURI shall contain the URI to perform an HTTP or HTTPS boot when
	// 'BootSourceOverrideTarget' is set to 'UefiHttp'. If this property is not
	// configured or supported, the URI shall be provided by a DHCP server as
	// specified by the UEFI Specification.
	//
	// Version added: v1.9.0
	HTTPBootURI string `json:"HttpBootUri"`
	// RemainingAutomaticRetryAttempts shall contain the number of attempts
	// remaining the system will retry booting in the event the system enters an
	// error state on boot. If '0', the system has no remaining automatic boot
	// retry attempts and shall not automatically retry booting if the system
	// enters an error state. This property shall be reset to the value of
	// 'AutomaticRetryAttempts' upon a successful boot attempt.
	//
	// Version added: v1.11.0
	RemainingAutomaticRetryAttempts *uint `json:",omitempty"`
	// StopBootOnFault shall contain the setting if the boot should stop on a
	// fault.
	//
	// Version added: v1.15.0
	StopBootOnFault StopBootOnFault `json:",omitempty"`
	// TrustedModuleRequiredToBoot shall contain the Trusted Module boot
	// requirement.
	//
	// Version added: v1.14.0
	TrustedModuleRequiredToBoot TrustedModuleRequiredToBoot `json:",omitempty"`
	// UefiTargetBootSourceOverride shall contain the UEFI device path of the
	// override boot target. Changes to this property do not alter the BIOS
	// persistent boot order configuration.
	UefiTargetBootSourceOverride                string   `json:",omitempty"`
	AllowableUefiTargetBootSourceOverrideValues []string `json:"UefiTargetBootSourceOverride@Redfish.AllowableValues,omitempty"`
}

// UnmarshalJSON unmarshals a Boot object from the raw JSON.
func (bo *Boot) UnmarshalJSON(b []byte) error {
	type temp Boot
	var tmp struct {
		temp
		BootOptions  Link `json:"BootOptions"`
		Certificates Link `json:"Certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*bo = Boot(tmp.temp)

	// Extract the links to other entities for later
	bo.bootOptions = tmp.BootOptions.String()
	bo.certificates = tmp.Certificates.String()

	return nil
}

// BootOptions gets the BootOptions collection.
func (bo *Boot) BootOptions(client Client) ([]*BootOption, error) {
	if bo.bootOptions == "" {
		return nil, nil
	}
	return GetCollectionObjects[BootOption](client, bo.bootOptions)
}

// Certificates gets the Certificates collection.
func (bo *Boot) Certificates(client Client) ([]*Certificate, error) {
	if bo.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, bo.certificates)
}

// BootProgress shall contain the last boot progress state and time.
type BootProgress struct {
	// LastBootTimeSeconds shall contain the number of seconds that elapsed between
	// system reset or power on and LastState transitioning to 'OSRunning'. If
	// 'LastState' contains 'OSRunning', this property shall contain the most
	// recent boot time. For other values of 'LastState', this property shall
	// contain the boot time for the previous boot.
	//
	// Version added: v1.18.0
	LastBootTimeSeconds *float64 `json:",omitempty"`
	// LastState shall contain the last boot progress state.
	//
	// Version added: v1.13.0
	LastState BootProgressTypes
	// LastStateTime shall contain the date and time when the last boot state was
	// updated.
	//
	// Version added: v1.13.0
	LastStateTime string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.13.0
	OEM json.RawMessage `json:"Oem"`
	// OEMLastState shall represent the OEM-specific 'LastState' of the
	// 'BootProgress'. This property shall only be present if 'LastState' is 'OEM'.
	//
	// Version added: v1.13.0
	OEMLastState string
}

// Composition shall contain information about the composition capabilities and
// state of a computer system.
type Composition struct {
	// UseCases shall contain the composition use cases in which this computer
	// system can participate.
	//
	// Version added: v1.18.0
	UseCases []CompositionUseCase
}

// HostGraphicalConsole shall describe a graphical console service for a
// computer system.
type HostGraphicalConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. 'KVMIP'
	// shall be included if a vendor-defined KVM-IP protocol is supported.
	//
	// Version added: v1.13.0
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	//
	// Version added: v1.13.0
	MaxConcurrentSessions uint
	// Port shall contain the port assigned to the service.
	//
	// Version added: v1.13.0
	Port *uint `json:",omitempty"`
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	//
	// Version added: v1.13.0
	ServiceEnabled bool
}

// HostSerialConsole shall describe the serial console services for a computer
// system.
type HostSerialConsole struct {
	// IPMI shall contain connection details for a serial console service that uses
	// the IPMI Serial-over-LAN (SOL) protocol.
	//
	// Version added: v1.13.0
	IPMI SerialConsoleProtocol
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	//
	// Version added: v1.13.0
	MaxConcurrentSessions uint
	// SSH shall contain connection details for a serial console service that uses
	// the Secure Shell (SSH) protocol.
	//
	// Version added: v1.13.0
	SSH SerialConsoleProtocol
	// Telnet shall contain connection details for a serial console service that
	// uses the Telnet protocol.
	//
	// Version added: v1.13.0
	Telnet SerialConsoleProtocol
	// WebSocket shall contain connection details for a serial console service that
	// uses WebSockets as defined by the 'WebSocket inbound access' clause of the
	// Redfish Specification. Services shall send WebSocket packetized bytes in a
	// manner that emulates a pty (pseudoterminal).
	//
	// Version added: v1.26.0
	WebSocket WebSocketConsole
}

// HostedServices shall describe services that a computer system supports.
type HostedServices struct {
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.2.0
	OEM json.RawMessage `json:"Oem"`
	// StorageServices shall contain a link to a resource collection of type
	// 'HostedStorageServices'.
	//
	// Version added: v1.2.0
	storageServices string
}

// UnmarshalJSON unmarshals a HostedServices object from the raw JSON.
func (h *HostedServices) UnmarshalJSON(b []byte) error {
	type temp HostedServices
	var tmp struct {
		temp
		StorageServices Link `json:"StorageServices"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = HostedServices(tmp.temp)

	// Extract the links to other entities for later
	h.storageServices = tmp.StorageServices.String()

	return nil
}

// StorageServices gets the StorageServices linked resource.
func (h *HostedServices) StorageServices(client Client) (*StorageService, error) {
	if h.storageServices == "" {
		return nil, nil
	}
	return GetObject[StorageService](client, h.storageServices)
}

// IPMIHostInterface shall describe the in-band IPMI service for a computer
// system.
type IPMIHostInterface struct {
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	//
	// Version added: v1.25.0
	ServiceEnabled bool
}

// IdlePowerSaver shall contain the idle power saver settings of a computer
// system.
type IdlePowerSaver struct {
	// Enabled shall indicate if idle power saver is enabled.
	//
	// Version added: v1.16.0
	Enabled bool
	// EnterDwellTimeSeconds shall contain the duration in seconds the computer
	// system is below the 'EnterUtilizationPercent' value before the idle power
	// save is activated.
	//
	// Version added: v1.16.0
	EnterDwellTimeSeconds *uint `json:",omitempty"`
	// EnterUtilizationPercent shall contain the percentage of utilization,
	// typically '0' to '100', when the computer system enters idle power save. If
	// the computer system's utilization goes below this value for the duration
	// specified by 'EnterDwellTimeSeconds', it shall enter idle power save.
	//
	// Version added: v1.16.0
	EnterUtilizationPercent *float64 `json:",omitempty"`
	// ExitDwellTimeSeconds shall contain the duration in seconds the computer
	// system is above the 'ExitUtilizationPercent' value before the idle power
	// save is stopped.
	//
	// Version added: v1.16.0
	ExitDwellTimeSeconds *uint `json:",omitempty"`
	// ExitUtilizationPercent shall contain the percentage of utilization,
	// typically '0' to '100', when the computer system exits idle power save. If
	// the computer system's utilization goes above this value for the duration
	// specified by 'ExitDwellTimeSeconds', it shall exit idle power save.
	//
	// Version added: v1.16.0
	ExitUtilizationPercent *float64 `json:",omitempty"`
}

// KMIPServer shall contain the KMIP server settings for a computer system.
type KMIPServer struct {
	// Address shall contain the KMIP server address.
	//
	// Version added: v1.16.0
	Address string
	// CacheDuration shall contain the duration that the system caches KMIP data.
	//
	// Version added: v1.20.0
	CacheDuration string
	// CachePolicy shall contain the cache policy to control how KMIP data is
	// cached.
	//
	// Version added: v1.20.0
	CachePolicy KMIPCachePolicy
	// Password shall contain the password to access the KMIP server. The value
	// shall be 'null' in responses.
	//
	// Version added: v1.16.0
	Password string
	// Port shall contain the KMIP server port.
	//
	// Version added: v1.16.0
	Port *int `json:",omitempty"`
	// Username shall contain the username to access the KMIP server.
	//
	// Version added: v1.16.0
	Username string
}

// KeyManagement shall contain the key management settings of a computer system.
type KeyManagement struct {
	// KMIPCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the server certificates for the
	// servers referenced by the 'KMIPServers' property.
	//
	// Version added: v1.16.0
	kMIPCertificates string
	// KMIPClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates
	// provided to the servers referenced by the 'KMIPServers' property.
	//
	// Version added: v1.27.0
	kMIPClientCertificates string
	// KMIPServers shall contain the KMIP servers to which this computer system is
	// subscribed for key management.
	//
	// Version added: v1.16.0
	KMIPServers []KMIPServer
}

// UnmarshalJSON unmarshals a KeyManagement object from the raw JSON.
func (k *KeyManagement) UnmarshalJSON(b []byte) error {
	type temp KeyManagement
	var tmp struct {
		temp
		KMIPCertificates       Link `json:"KMIPCertificates"`
		KMIPClientCertificates Link `json:"KMIPClientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*k = KeyManagement(tmp.temp)

	// Extract the links to other entities for later
	k.kMIPCertificates = tmp.KMIPCertificates.String()
	k.kMIPClientCertificates = tmp.KMIPClientCertificates.String()

	return nil
}

// KMIPCertificates gets the KMIPCertificates collection.
func (k *KeyManagement) KMIPCertificates(client Client) ([]*Certificate, error) {
	if k.kMIPCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, k.kMIPCertificates)
}

// KMIPClientCertificates gets the KMIPClientCertificates collection.
func (k *KeyManagement) KMIPClientCertificates(client Client) ([]*Certificate, error) {
	if k.kMIPClientCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](client, k.kMIPClientCertificates)
}

// MemorySummary shall contain properties that describe the central memory for a
// system.
type MemorySummary struct {
	// MemoryMirroring shall contain the ability and type of memory mirroring that
	// this computer system supports.
	//
	// Version added: v1.1.0
	MemoryMirroring MemoryMirroring
	// Metrics shall contain a link to the metrics associated with all memory in
	// this system.
	//
	// Version added: v1.8.0
	metrics string
	// Status shall contain any status or health properties of the resource.
	//
	// Deprecated: v1.16.0
	// This property has been deprecated in favor of the 'Conditions' property
	// within 'Status' in the root of this resource.
	Status Status
	// TotalSystemMemoryGiB shall contain the amount of configured system general
	// purpose volatile (RAM) memory as measured in gibibytes.
	TotalSystemMemoryGiB *float64 `json:",omitempty"`
	// TotalSystemPersistentMemoryGiB shall contain the total amount of configured
	// persistent memory available to the system as measured in gibibytes.
	//
	// Version added: v1.4.0
	TotalSystemPersistentMemoryGiB *float64 `json:",omitempty"`
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (m *MemorySummary) UnmarshalJSON(b []byte) error {
	type temp MemorySummary
	var tmp struct {
		temp
		Metrics Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemorySummary(tmp.temp)

	// Extract the links to other entities for later
	m.metrics = tmp.Metrics.String()

	return nil
}

// Metrics gets the Metrics linked resource.
func (m *MemorySummary) Metrics(client Client) (*MemoryMetrics, error) {
	if m.metrics == "" {
		return nil, nil
	}
	return GetObject[MemoryMetrics](client, m.metrics)
}

// ProcessorSummary shall contain properties that describe the central
// processors for a system. Processors described by this type shall be limited
// to the processors that execute system code, and shall not include processors
// used for offload functionality.
type ProcessorSummary struct {
	// CoreCount shall contain the total number of central processor cores in the
	// system.
	//
	// Version added: v1.14.0
	CoreCount *uint `json:",omitempty"`
	// Count shall contain the total number of physical central processors in the
	// system.
	Count *uint `json:",omitempty"`
	// LogicalProcessorCount shall contain the total number of logical central
	// processors in the system.
	//
	// Version added: v1.5.0
	LogicalProcessorCount *uint `json:",omitempty"`
	// Metrics shall contain a link to the metrics associated with all processors
	// in this system.
	//
	// Version added: v1.7.0
	metrics string
	// Model shall contain the processor model for the central processors in the
	// system, per the description in the Processor Information - Processor Family
	// section of the SMBIOS Specification DSP0134 2.8 or later.
	Model string
	// Status shall contain any status or health properties of the resource.
	//
	// Deprecated: v1.16.0
	// This property has been deprecated in favor of the 'Conditions' property
	// within 'Status' in the root of this resource.
	Status Status
	// ThreadingEnabled shall indicate that all 'Processor' resources in this
	// system where the 'ProcessorType' property contains 'CPU' have multiple
	// threading support enabled.
	//
	// Version added: v1.15.0
	ThreadingEnabled bool
}

// UnmarshalJSON unmarshals a ProcessorSummary object from the raw JSON.
func (p *ProcessorSummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorSummary
	var tmp struct {
		temp
		Metrics Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = ProcessorSummary(tmp.temp)

	// Extract the links to other entities for later
	p.metrics = tmp.Metrics.String()

	return nil
}

// Metrics gets the Metrics linked resource.
func (p *ProcessorSummary) Metrics(client Client) (*ProcessorMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[ProcessorMetrics](client, p.metrics)
}

// SerialConsoleProtocol shall describe a serial console service for a computer
// system.
type SerialConsoleProtocol struct {
	// ConsoleEntryCommand shall contain a command string that can be provided by a
	// client to select or enter the system's serial console, when the console is
	// shared among several systems or a manager CLI.
	//
	// Version added: v1.13.0
	ConsoleEntryCommand string
	// HotKeySequenceDisplay shall contain a string that can be provided to a user
	// to describe the hotkey sequence used to exit the serial console session, or,
	// if shared with a manager CLI, to return to the CLI.
	//
	// Version added: v1.13.0
	HotKeySequenceDisplay string
	// Port shall contain the port assigned to the protocol.
	//
	// Version added: v1.13.0
	Port *uint `json:",omitempty"`
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	//
	// Version added: v1.13.0
	ServiceEnabled bool
	// SharedWithManagerCLI shall indicate whether the serial console service is
	// shared with access to the manager's command-line interface (CLI).
	//
	// Version added: v1.13.0
	SharedWithManagerCLI bool
}

// TrustedModules shall describe a Trusted Module for a system.
type TrustedModules struct {
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the Trusted Module.
	//
	// Version added: v1.1.0
	FirmwareVersion string
	// FirmwareVersion2 shall contain the 2nd firmware version, if applicable, as
	// defined by the manufacturer for the Trusted Module.
	//
	// Version added: v1.3.0
	FirmwareVersion2 string
	// InterfaceType shall contain the interface type of the installed Trusted
	// Module.
	//
	// Version added: v1.1.0
	InterfaceType InterfaceType
	// InterfaceTypeSelection shall contain the interface type 'Selection' method
	// (for example to switch between TPM1_2 and TPM2_0) that is supported by this
	// Trusted Module.
	//
	// Version added: v1.3.0
	InterfaceTypeSelection InterfaceTypeSelection
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.1.0
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.1.0
	Status Status
}

// ComputerVirtualMediaConfig shall describe a virtual media service for a computer
// system.
type ComputerVirtualMediaConfig struct {
	// Port shall contain the port assigned to the service.
	//
	// Version added: v1.13.0
	Port *uint `json:",omitempty"`
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	//
	// Version added: v1.13.0
	ServiceEnabled bool
}

// WatchdogTimer shall contain properties that describe the host watchdog timer
// functionality for this system.
type WatchdogTimer struct {
	// FunctionEnabled shall indicate whether a user has enabled the host watchdog
	// timer functionality. This property indicates only that a user has enabled
	// the timer. To activate the timer, installation of additional host-based
	// software is necessary; an update to this property does not initiate the
	// timer.
	//
	// Version added: v1.5.0
	FunctionEnabled bool
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.5.0
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.5.0
	Status Status
	// TimeoutAction shall contain the action to perform when the watchdog timer
	// reaches its timeout value.
	//
	// Version added: v1.5.0
	TimeoutAction WatchdogTimeoutActions
	// WarningAction shall contain the action to perform before the watchdog timer
	// expires. This action typically occurs from three to ten seconds before to
	// the timeout value, but the exact timing is dependent on the implementation.
	//
	// Version added: v1.5.0
	WarningAction WatchdogWarningActions
}

// WebSocketConsole shall describe a WebSocket serial console service for a
// computer system.
type WebSocketConsole struct {
	// ConsoleURI shall contain the URI at which to access the WebSocket serial
	// console, using the Redfish protocol and authentication methods. See the
	// 'WebSocket inbound access' clause in the Redfish Specification.
	//
	// Version added: v1.26.0
	ConsoleURI string
	// Interactive shall indicate if the WebSocket serial console allows
	// interactive input. If 'true', the WebSocket is bidirectional. If 'false',
	// the WebSocket only allows console output from the service.
	//
	// Version added: v1.26.0
	Interactive bool
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	//
	// Version added: v1.26.0
	ServiceEnabled bool
}
