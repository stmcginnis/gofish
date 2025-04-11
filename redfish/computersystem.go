//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/stmcginnis/gofish/common"
)

type AutomaticRetryConfig string

const (
	// DisabledAutomaticRetryConfig shall indicate that automatic retrying of booting is disabled.
	DisabledAutomaticRetryConfig AutomaticRetryConfig = "Disabled"
	// RetryAttemptsAutomaticRetryConfig shall indicate that the number of retries of booting is based on the
	// AutomaticRetryAttempts property, and the RemainingAutomaticRetryAttempts property indicates the number of
	// remaining attempts.
	RetryAttemptsAutomaticRetryConfig AutomaticRetryConfig = "RetryAttempts"
	// RetryAlwaysAutomaticRetryConfig shall indicate that the system will always automatically retry booting.
	RetryAlwaysAutomaticRetryConfig AutomaticRetryConfig = "RetryAlways"
)

// BootOrderTypes is the choice of boot order property to use when controller
// the persistent boot order for this computer system.
type BootOrderTypes string

const (

	// BootOrderBootOrderTypes indicates an ordered array of persistent boot
	// order for this system.
	BootOrderBootOrderTypes BootOrderTypes = "BootOrder"
	// AliasBootOrderBootOrderTypes  indicates an ordered array of aliased
	// persistent boot order for this system.
	AliasBootOrderBootOrderTypes BootOrderTypes = "AliasBootOrder"
)

type BootProgressTypes string

const (
	// NoneBootProgressTypes shall indicate that the system is not booting or running, such as the system is powered
	// off.
	NoneBootProgressTypes BootProgressTypes = "None"
	// PrimaryProcessorInitializationStartedBootProgressTypes shall indicate that the system has started to initialize
	// the primary processor.
	PrimaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "PrimaryProcessorInitializationStarted"
	// BusInitializationStartedBootProgressTypes shall indicate that the system has started to initialize the buses.
	BusInitializationStartedBootProgressTypes BootProgressTypes = "BusInitializationStarted"
	// MemoryInitializationStartedBootProgressTypes shall indicate that the system has started to initialize the
	// memory.
	MemoryInitializationStartedBootProgressTypes BootProgressTypes = "MemoryInitializationStarted"
	// SecondaryProcessorInitializationStartedBootProgressTypes shall indicate that the system has started to
	// initialize the secondary processors.
	SecondaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "SecondaryProcessorInitializationStarted"
	// PCIResourceConfigStartedBootProgressTypes shall indicate that the system has started to initialize the PCI
	// resources.
	PCIResourceConfigStartedBootProgressTypes BootProgressTypes = "PCIResourceConfigStarted"
	// SystemHardwareInitializationCompleteBootProgressTypes shall indicate that the system has completed initializing
	// all hardware.
	SystemHardwareInitializationCompleteBootProgressTypes BootProgressTypes = "SystemHardwareInitializationComplete"
	// SetupEnteredBootProgressTypes shall indicate that the system has entered the setup utility.
	SetupEnteredBootProgressTypes BootProgressTypes = "SetupEntered"
	// OSBootStartedBootProgressTypes shall indicate that the operating system has started to boot.
	OSBootStartedBootProgressTypes BootProgressTypes = "OSBootStarted"
	// OSRunningBootProgressTypes shall indicate that the operating system is running and shall indicate the final boot
	// progress state.
	OSRunningBootProgressTypes BootProgressTypes = "OSRunning"
	// OEMBootProgressTypes shall indicate an OEM-defined boot progress state.
	OEMBootProgressTypes BootProgressTypes = "OEM"
)

// BootSourceOverrideEnabled describes the state of the Boot Source Override feature.
type BootSourceOverrideEnabled string

const (

	// DisabledBootSourceOverrideEnabled the system will boot normally.
	DisabledBootSourceOverrideEnabled BootSourceOverrideEnabled = "Disabled"
	// OnceBootSourceOverrideEnabled on its next boot cycle, the system will
	// boot (one time) to the Boot Source Override Target. The value of
	// BootSourceOverrideEnabled is then reset back to Disabled.
	OnceBootSourceOverrideEnabled BootSourceOverrideEnabled = "Once"
	// ContinuousBootSourceOverrideEnabled the system will boot to the target
	// specified in the BootSourceOverrideTarget until this property is set
	// to Disabled.
	ContinuousBootSourceOverrideEnabled BootSourceOverrideEnabled = "Continuous"
)

type CompositionUseCase string

const (
	// ResourceBlockCapableCompositionUseCase shall indicate the computer system supports being registered as a
	// resource block in order for it to participate in composition requests.
	ResourceBlockCapableCompositionUseCase CompositionUseCase = "ResourceBlockCapable"
	// ExpandableSystemCompositionUseCase shall indicate the computer system supports expandable system composition and
	// is associated with a resource block.
	ExpandableSystemCompositionUseCase CompositionUseCase = "ExpandableSystem"
)

type DecommissionType string

const (
	// AllDecommissionType shall indicate the service removes all the data that it can from the system. This shall
	// include all possible OEM data as well.
	AllDecommissionType DecommissionType = "All"
	// UserDataDecommissionType shall indicate the service removes all the data from block devices or other operating
	// system accessible storage. If the RequireSecureErase parameter contains 'true', this shall be equivalent to
	// performing the SecureErase action on each drive.
	UserDataDecommissionType DecommissionType = "UserData"
	// ManagerConfigDecommissionType shall indicate the service resets all associated managers to factory defaults.
	// This shall be equivalent to performing the ResetToDefaults action on each Manager resource with the ResetType
	// parameter of 'ResetAll'.
	ManagerConfigDecommissionType DecommissionType = "ManagerConfig"
	// BIOSConfigDecommissionType shall indicate the service resets all BIOS settings to factory defaults. This shall
	// be equivalent to performing the ResetBios action on each Bios resource.
	BIOSConfigDecommissionType DecommissionType = "BIOSConfig"
	// NetworkConfigDecommissionType shall indicate the service resets all network settings on all network devices to
	// factory defaults.
	NetworkConfigDecommissionType DecommissionType = "NetworkConfig"
	// StorageConfigDecommissionType shall indicate the service resets all storage controller settings to factory
	// defaults. This shall be equivalent to performing the ResetToDefaults action on each Storage resource with the
	// ResetType parameter of 'PreserveVolumes'.
	StorageConfigDecommissionType DecommissionType = "StorageConfig"
	// LogsDecommissionType shall indicate the service clears all logs. This shall be equivalent to performing the
	// ClearLog action on each LogService resource.
	LogsDecommissionType DecommissionType = "Logs"
)

// BootSourceOverrideTarget the current boot source to be used at next boot instead of the normal boot device, if BootSourceOverrideEnabled is true.
type BootSourceOverrideTarget string

const (
	// NoneBootSourceOverrideTarget boot from the normal boot device
	NoneBootSourceOverrideTarget BootSourceOverrideTarget = "None"
	// PxeBootSourceOverrideTarget boot from the Pre-Boot EXecution (PXE) environment.
	PxeBootSourceOverrideTarget BootSourceOverrideTarget = "Pxe"
	// FloppyBootSourceOverrideTarget boot from the floppy disk drive
	FloppyBootSourceOverrideTarget BootSourceOverrideTarget = "Floppy"
	// CdBootSourceOverrideTarget boot from the CD/DVD disc
	CdBootSourceOverrideTarget BootSourceOverrideTarget = "Cd"
	// UsbBootSourceOverrideTarget boot from a USB device as specified by the system BIOS
	UsbBootSourceOverrideTarget BootSourceOverrideTarget = "Usb"
	// HddBootSourceOverrideTarget boot from a hard drive
	HddBootSourceOverrideTarget BootSourceOverrideTarget = "Hdd"
	// BiosSetupBootSourceOverrideTarget boot to the BIOS Setup Utility
	BiosSetupBootSourceOverrideTarget BootSourceOverrideTarget = "BiosSetup"
	// UtilitiesBootSourceOverrideTarget boot the manufacturer's Utilities program(s)
	UtilitiesBootSourceOverrideTarget BootSourceOverrideTarget = "Utilities"
	// DiagsBootSourceOverrideTarget boot the manufacturer's Diagnostics program
	DiagsBootSourceOverrideTarget BootSourceOverrideTarget = "Diags"
	// UefiShellBootSourceOverrideTarget boot to the UEFI Shell.
	UefiShellBootSourceOverrideTarget BootSourceOverrideTarget = "UefiShell"
	// UefiTargetBootSourceOverrideTarget boot to the UEFI Device specified in the UefiTargetBootSourceOverride property.
	UefiTargetBootSourceOverrideTarget BootSourceOverrideTarget = "UefiTarget"
	// SDCardBootSourceOverrideTarget boot from an SD Card
	SDCardBootSourceOverrideTarget BootSourceOverrideTarget = "SDCard"
	// UefiHTTPBootSourceOverrideTarget boot from a UEFI HTTP network location
	UefiHTTPBootSourceOverrideTarget BootSourceOverrideTarget = "UefiHttp"
	// RemoteDriveBootSourceOverrideTarget boot from a remote drive (e.g. iSCSI).
	RemoteDriveBootSourceOverrideTarget BootSourceOverrideTarget = "RemoteDrive"
	// UefiBootNextBootSourceOverrideTarget boot to the UEFI Device specified in the BootNext property
	UefiBootNextBootSourceOverrideTarget BootSourceOverrideTarget = "UefiBootNext"
)

// BootSourceOverrideMode is the BIOS mode (Legacy or UEFI) to be used.
type BootSourceOverrideMode string

const (

	// LegacyBootSourceOverrideMode the system will boot in non-UEFI boot
	// mode to the Boot Source Override Target.
	LegacyBootSourceOverrideMode BootSourceOverrideMode = "Legacy"
	// UEFIBootSourceOverrideMode the system will boot in UEFI boot mode to
	// the Boot Source Override Target.
	UEFIBootSourceOverrideMode BootSourceOverrideMode = "UEFI"
)

// HostingRole specifies different features that the hosting ComputerSystem supports.
type HostingRole string

const (

	// ApplicationServerHostingRole the system hosts functionality that
	// supports general purpose applications.
	ApplicationServerHostingRole HostingRole = "ApplicationServer"
	// StorageServerHostingRole the system hosts functionality that supports
	// the system acting as a storage server.
	StorageServerHostingRole HostingRole = "StorageServer"
	// SwitchHostingRole the system hosts functionality that supports the
	// system acting as a switch.
	SwitchHostingRole HostingRole = "Switch"
)

// InterfaceType is the Trusted Platform Module type.
type InterfaceType string

const (

	// TPM1_2InterfaceType Trusted Platform Module (TPM) 1.2.
	TPM1_2InterfaceType InterfaceType = "TPM1_2"
	// TPM2_0InterfaceType Trusted Platform Module (TPM) 2.0.
	TPM2_0InterfaceType InterfaceType = "TPM2_0"
	// TCM1_0InterfaceType Trusted Cryptography Module (TCM) 1.0.
	TCM1_0InterfaceType InterfaceType = "TCM1_0"
)

// InterfaceTypeSelection specify the method for switching the TrustedModule
// InterfaceType, for instance between TPM1_2 and TPM2_0, if supported.
type InterfaceTypeSelection string

const (

	// NoneInterfaceTypeSelection the TrustedModule does not support
	// switching the InterfaceType.
	NoneInterfaceTypeSelection InterfaceTypeSelection = "None"
	// FirmwareUpdateInterfaceTypeSelection the TrustedModule supports
	// switching InterfaceType via a firmware update.
	FirmwareUpdateInterfaceTypeSelection InterfaceTypeSelection = "FirmwareUpdate"
	// BiosSettingInterfaceTypeSelection the TrustedModule supports switching
	// InterfaceType via platform software, such as a BIOS configuration
	// Attribute.
	BiosSettingInterfaceTypeSelection InterfaceTypeSelection = "BiosSetting"
	// OemMethodInterfaceTypeSelection the TrustedModule supports switching
	// InterfaceType via an OEM proprietary mechanism.
	OemMethodInterfaceTypeSelection InterfaceTypeSelection = "OemMethod"
)

type KMIPCachePolicy string

const (
	// NoneKMIPCachePolicy The system does not cache KMIP data.
	NoneKMIPCachePolicy KMIPCachePolicy = "None"
	// AfterFirstUseKMIPCachePolicy The system caches KMIP data after first use for the duration specified by the
	// CacheDuration property.
	AfterFirstUseKMIPCachePolicy KMIPCachePolicy = "AfterFirstUse"
)

// MemoryMirroring indicates the memory mirroring setting
type MemoryMirroring string

const (

	// SystemMemoryMirroring the system supports DIMM mirroring at the System
	// level. Individual DIMMs are not paired for mirroring in this mode.
	SystemMemoryMirroring MemoryMirroring = "System"
	// DIMMMemoryMirroring the system supports DIMM mirroring at the DIMM
	// level. Individual DIMMs can be mirrored.
	DIMMMemoryMirroring MemoryMirroring = "DIMM"
	// HybridMemoryMirroring the system supports a hybrid mirroring at the
	// system and DIMM levels. Individual DIMMs can be mirrored.
	HybridMemoryMirroring MemoryMirroring = "Hybrid"
	// NoneMemoryMirroring the system does not support DIMM mirroring.
	NoneMemoryMirroring MemoryMirroring = "None"
)

type PowerMode string

const (
	// MaximumPerformancePowerMode shall indicate the system performs at the highest speeds possible. This mode should
	// be used when performance is the top priority.
	MaximumPerformancePowerMode PowerMode = "MaximumPerformance"
	// BalancedPerformancePowerMode shall indicate the system performs at the highest speeds possible when the
	// utilization is high and performs at reduced speeds when the utilization is low to save power. This mode is a
	// compromise between 'MaximumPerformance' and 'PowerSaving'.
	BalancedPerformancePowerMode PowerMode = "BalancedPerformance"
	// PowerSavingPowerMode shall indicate the system performs at reduced speeds to save power. This mode should be
	// used when power saving is the top priority.
	PowerSavingPowerMode PowerMode = "PowerSaving"
	// StaticPowerMode shall indicate the system performs at a static base speed.
	StaticPowerMode PowerMode = "Static"
	// OSControlledPowerMode shall indicate the system performs at an operating system-controlled power mode.
	OSControlledPowerMode PowerMode = "OSControlled"
	// OEMPowerMode shall indicate the system performs at an OEM-defined power mode.
	OEMPowerMode PowerMode = "OEM"
	// EfficiencyFavorPowerPowerMode shall indicate the system performs at reduced speeds at all utilizations to save
	// power at the cost of performance. This mode differs from 'PowerSaving' in that more performance is retained and
	// less power is saved. This mode differs from 'EfficiencyFavorPerformance' in that less performance is retained
	// but more power is saved. This mode differs from 'BalancedPerformance' in that power saving occurs at all
	// utilizations.
	EfficiencyFavorPowerPowerMode PowerMode = "EfficiencyFavorPower"
	// EfficiencyFavorPerformancePowerMode shall indicate the system performs at reduced speeds at all utilizations to
	// save power while attempting to maintain performance. This mode differs from 'EfficiencyFavorPower' in that more
	// performance is retained but less power is saved. This mode differs from 'MaximumPerformance' in that power is
	// saved at the cost of some performance. This mode differs from 'BalancedPerformance' in that power saving occurs
	// at all utilizations.
	EfficiencyFavorPerformancePowerMode PowerMode = "EfficiencyFavorPerformance"
)

// PowerRestorePolicyTypes specifies the choice of power state for the system
// when power is applied.
type PowerRestorePolicyTypes string

const (

	// AlwaysOnPowerRestorePolicyTypes the system will always power on when
	// power is applied.
	AlwaysOnPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOn"
	// AlwaysOffPowerRestorePolicyTypes the system will always remain powered
	// off when power is applied.
	AlwaysOffPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOff"
	// LastStatePowerRestorePolicyTypes the system will return to its last
	// power state (on or off) when power is applied.
	LastStatePowerRestorePolicyTypes PowerRestorePolicyTypes = "LastState"
)

// PowerState is the power state of the system.
type PowerState string

const (
	// OnPowerState the system is powered on.
	OnPowerState PowerState = "On"
	// OffPowerState the system is powered off, although some components may
	// continue to have AUX power such as management controller.
	OffPowerState PowerState = "Off"
	// PausedPowerState the system is paused.
	PausedPowerState PowerState = "Paused"
	// PoweringOnPowerState A temporary state between Off and On. This
	// temporary state can be very short.
	PoweringOnPowerState PowerState = "PoweringOn"
	// PoweringOffPowerState A temporary state between On and Off. The power
	// off action can take time while the OS is in the shutdown process.
	PoweringOffPowerState PowerState = "PoweringOff"
)

type StopBootOnFault string

const (
	// NeverStopBootOnFault shall indicate the system will continue to attempt to boot if a fault occurs.
	NeverStopBootOnFault StopBootOnFault = "Never"
	// AnyFaultStopBootOnFault shall indicate the system will stop the boot if a fault occurs. This includes, but is
	// not limited to, faults that affect performance, fault tolerance, or capacity.
	AnyFaultStopBootOnFault StopBootOnFault = "AnyFault"
)

// SystemType is the type of system.
type SystemType string

const (

	// PhysicalSystemType is a System Type of Physical is typically used when
	// representing the hardware aspects of a system such as is done by a
	// management controller.
	PhysicalSystemType SystemType = "Physical"
	// VirtualSystemType is a System Type of Virtual is typically used when
	// representing a system that is actually a virtual machine instance.
	VirtualSystemType SystemType = "Virtual"
	// OSSystemType is a System Type of OS is typically used when representing
	// an OS or hypervisor view of the system.
	OSSystemType SystemType = "OS"
	// PhysicallyPartitionedSystemType is a System Type of PhysicallyPartition is
	// typically used when representing a single system constructed from
	// one or more physical systems via a firmware or hardware-based service.
	PhysicallyPartitionedSystemType SystemType = "PhysicallyPartitioned"
	// VirtuallyPartitionedSystemType is a System Type of VirtuallyPartition is
	// typically used when representing a single system constructed from
	// one or more virtual systems via a software-based service.
	VirtuallyPartitionedSystemType SystemType = "VirtuallyPartitioned"
	// ComposedSystemType is a System Type of Composed is typically used when
	// representing a single system constructed from disaggregated resource
	// via the Redfish Composition service.
	ComposedSystemType SystemType = "Composed"
)

type TrustedModuleRequiredToBoot string

const (
	// DisabledTrustedModuleRequiredToBoot shall indicate a Trusted Module is not required to boot.
	DisabledTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Disabled"
	// RequiredTrustedModuleRequiredToBoot shall indicate a functioning Trusted Module is required to boot.
	RequiredTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Required"
)

// WatchdogTimeoutActions specifies the choice of action to take when the Host
// Watchdog Timer reaches its timeout value.
type WatchdogTimeoutActions string

const (

	// NoneWatchdogTimeoutActions means no action taken.
	NoneWatchdogTimeoutActions WatchdogTimeoutActions = "None"
	// ResetSystemWatchdogTimeoutActions means reset the system.
	ResetSystemWatchdogTimeoutActions WatchdogTimeoutActions = "ResetSystem"
	// PowerCycleWatchdogTimeoutActions means power cycle the system.
	PowerCycleWatchdogTimeoutActions WatchdogTimeoutActions = "PowerCycle"
	// PowerDownWatchdogTimeoutActions means power down the system.
	PowerDownWatchdogTimeoutActions WatchdogTimeoutActions = "PowerDown"
	// OEMWatchdogTimeoutActions means perform an OEM-defined action.
	OEMWatchdogTimeoutActions WatchdogTimeoutActions = "OEM"
)

// WatchdogWarningActions specifies the choice of action to take when the Host
// Watchdog Timer is close (typically 3-10 seconds) to reaching its timeout value.
type WatchdogWarningActions string

const (

	// NoneWatchdogWarningActions means no action taken.
	NoneWatchdogWarningActions WatchdogWarningActions = "None"
	// DiagnosticInterruptWatchdogWarningActions means raise a (typically non-
	// maskable) Diagnostic Interrupt.
	DiagnosticInterruptWatchdogWarningActions WatchdogWarningActions = "DiagnosticInterrupt"
	// SMIWatchdogWarningActions means raise a Systems Management Interrupt (SMI).
	SMIWatchdogWarningActions WatchdogWarningActions = "SMI"
	// MessagingInterruptWatchdogWarningActions means raise a legacy IPMI messaging
	// interrupt.
	MessagingInterruptWatchdogWarningActions WatchdogWarningActions = "MessagingInterrupt"
	// SCIWatchdogWarningActions means raise an interrupt using the ACPI System
	// Control Interrupt (SCI).
	SCIWatchdogWarningActions WatchdogWarningActions = "SCI"
	// OEMWatchdogWarningActions means perform an OEM-defined action.
	OEMWatchdogWarningActions WatchdogWarningActions = "OEM"
)

// HostGraphicalConsole shall describe a graphical console service for a computer system.
type HostGraphicalConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. KVMIP shall be included if a vendor-defined
	// KVM-IP protocol is supported.
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions int
	// Port shall contain the port assigned to the service.
	Port int
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled bool
}

// HostedServices shall describe services that a computer system supports.
type HostedServices struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// StorageServices shall contain a link to a resource collection of type HostedStorageServices.
	storageServices string
}

// UnmarshalJSON unmarshals a HostedServices object from the raw JSON.
func (hostedservices *HostedServices) UnmarshalJSON(b []byte) error {
	type temp HostedServices
	var t struct {
		temp
		StorageServices common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostedservices = HostedServices(t.temp)

	// Extract the links to other entities for later
	hostedservices.storageServices = t.StorageServices.String()

	return nil
}

// Boot contains properties which describe boot information for a system.
type Boot struct {
	// AliasBootOrder shall be an ordered array
	// of boot source aliases (of type BootSource) representing the
	// persistent Boot Order of this computer system.
	AliasBootOrder []string `json:",omitempty"`
	// AutomaticRetryAttempts is the number of attempts the system will automatically retry
	// booting in the event the system enters an error state on boot.
	AutomaticRetryAttempts int `json:",omitempty"`
	// AutomaticRetryConfig is how the system retries booting automatically.
	AutomaticRetryConfig AutomaticRetryConfig `json:",omitempty"`
	// BootNext shall contain the BootOptionReference of the UEFI boot option for one time boot, as defined by the UEFI
	// Specification. The valid values for this property are specified in the values of the BootOrder array.
	// BootSourceOverrideEnabled set to 'Continuous' is not supported for BootSourceOverrideTarget set to
	// 'UefiBootNext' because this setting is defined in UEFI as a one-time boot setting.
	BootNext string `json:",omitempty"`
	// bootOptions is a link to the collection of the UEFI boot options
	// associated with this computer system.
	bootOptions string
	// BootOrder shall be an ordered array of
	// BootOptionReference strings representing the persistent Boot Order of
	// this computer system. For UEFI systems, this is the UEFI BootOrder as
	// defined by the UEFI Specification.
	BootOrder []string `json:",omitempty"`
	// BootOrderPropertySelection shall
	// indicate which boot order property the system uses when specifying the
	// persistent boot order.
	BootOrderPropertySelection string `json:",omitempty"`
	// BootSourceOverrideEnabled shall be Once
	// if this is a one time boot override and Continuous if this selection
	// should remain active until cancelled. If the property value is set to
	// Once, the value will be reset back to Disabled after the
	// BootSourceOverrideTarget actions have been completed. Changes to this
	// property do not alter the BIOS persistent boot order configuration.
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:",omitempty"`
	// BootSourceOverrideMode shall be Legacy
	// for non-UEFI BIOS boot or UEFI for UEFI boot from boot source
	// specified in BootSourceOverrideTarget property.
	BootSourceOverrideMode BootSourceOverrideMode `json:",omitempty"`
	// BootSourceOverrideTarget shall contain
	// the source to boot the system from, overriding the normal boot order.
	// The valid values for this property are specified through the
	// Redfish.AllowableValues annotation. Pxe indicates to PXE boot from the
	// primary NIC; Floppy, Cd, Usb, Hdd indicates to boot from their devices
	// respectively. BiosSetup indicates to boot into the native BIOS screen
	// setup. Utilities and Diags indicate to boot from the local utilities
	// or diags partitions. UefiTarget indicates to boot from the UEFI device
	// path found in UefiTargetBootSourceOverride. UefiBootNext indicates to
	// boot from the UEFI BootOptionReference found in BootNext. Changes to
	// this property do not alter the BIOS persistent boot order
	// configuration.
	BootSourceOverrideTarget BootSourceOverrideTarget `json:",omitempty"`
	// The link to a collection of certificates used for booting through HTTPS by this computer system.
	certificates string
	// The URI to boot from when BootSourceOverrideTarget is set to UefiHttp.
	HTTPBootURI string `json:"HttpBootUri,omitempty"`
	// RemainingAutomaticRetryAttempts shall contain the number of attempts remaining the system will retry booting in
	// the event the system enters an error state on boot. If '0', the system has no remaining automatic boot retry
	// attempts and shall not automatically retry booting if the system enters an error state. This property shall be
	// reset to the value of AutomaticRetryAttempts upon a successful boot attempt.
	RemainingAutomaticRetryAttempts int `json:",omitempty"`
	// StopBootOnFault shall contain the setting if the boot should stop on a fault.
	StopBootOnFault StopBootOnFault `json:",omitempty"`
	// TrustedModuleRequiredToBoot shall contain the Trusted Module boot requirement.
	TrustedModuleRequiredToBoot TrustedModuleRequiredToBoot `json:",omitempty"`
	// UefiTargetBootSourceOverride shall be
	// the UEFI device path of the override boot target. The valid values for
	// this property are specified through the Redfish.AllowableValues
	// annotation. BootSourceOverrideEnabled = Continuous is not supported
	// for UEFI Boot Source Override as this setting is defined in UEFI as a
	// one time boot only. Changes to this property do not alter the BIOS
	// persistent boot order configuration.
	UefiTargetBootSourceOverride string `json:",omitempty"`
}

// UnmarshalJSON unmarshals a Boot object from the raw JSON.
func (boot *Boot) UnmarshalJSON(b []byte) error {
	type temp Boot
	var t struct {
		temp
		BootOptions  common.Link
		Certificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boot = Boot(t.temp)

	// Extract the links to other entities for later
	boot.bootOptions = t.BootOptions.String()
	boot.certificates = t.Certificates.String()

	return nil
}

// BootProgress shall contain the last boot progress state and time.
type BootProgress struct {
	// LastBootTimeSeconds shall contain the number of seconds that elapsed between system reset or power on and
	// LastState transitioning to 'OSRunning'. If LastState contains 'OSRunning', this property shall contain the most
	// recent boot time. For other values of LastState, this property shall contain the boot time for the previous
	// boot.
	LastBootTimeSeconds float64
	// LastState shall contain the last boot progress state.
	LastState BootProgressTypes
	// LastStateTime shall contain the date and time when the last boot state was updated.
	LastStateTime string
}

// Composition shall contain information about the composition capabilities and state of a computer system.
type Composition struct {
	// UseCases shall contain the composition use cases in which this computer system can participate.
	UseCases []CompositionUseCase
}

// BootOption represents the properties of a bootable device available in the
// system.
type BootOption struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Alias is the alias of this boot source if one exists.
	Alias BootSourceOverrideTarget
	// BootOptionEnabled is an indication of whether the boot option is
	// enabled. If true , it is enabled. If false, the boot option that the
	// boot order array on the computer system contains is skipped. In the
	// UEFI context, this property shall influence the load option active
	// flag for the boot option.
	BootOptionEnabled bool
	// BootOptionReference is the unique identifier seen in Boot.BootOrder.
	BootOptionReference string
	// DisplayName is the user-readable display name of the boot option
	// that appears in the boot order list in the user interface.
	DisplayName string
	// UefiDevicePath is the UEFI device path to access this UEFI boot
	// option.
	UefiDevicePath string
}

// GetBootOption will get a BootOption instance from the service.
func GetBootOption(c common.Client, uri string) (*BootOption, error) {
	var bootoption BootOption
	return &bootoption, bootoption.Get(c, uri, &bootoption)
}

// ResetType describe the type off reset to be issue by the resource
type ResetType string

const (
	// OnResetType shall be used to power on the machine
	OnResetType ResetType = "On"
	// ForceOnResetType shall be used to power on the machine immediately
	ForceOnResetType ResetType = "ForceOn"
	// ForceOffResetType shall be used to power off the machine without wait the OS to shutdown
	ForceOffResetType ResetType = "ForceOff"
	// ForceRestartResetType shall be used to restart the machine without wait the OS to shutdown
	ForceRestartResetType ResetType = "ForceRestart"
	// FullPowerCycleResetType shall be used to perform an AC power cycle
	FullPowerCycleResetType ResetType = "FullPowerCycle"
	// GracefulRestartResetType shall be used to restart the machine waiting the OS shutdown gracefully
	GracefulRestartResetType ResetType = "GracefulRestart"
	// GracefulShutdownResetType shall be used to restart the machine waiting the OS shutdown gracefully
	GracefulShutdownResetType ResetType = "GracefulShutdown"
	// PushPowerButtonResetType shall be used to emulate pushing the power button
	PushPowerButtonResetType ResetType = "PushPowerButton"
	// PowerCycleResetType shall be used to power cycle the machine
	PowerCycleResetType ResetType = "PowerCycle"
	// NmiResetType shall be used to trigger a crash/core dump file
	NmiResetType ResetType = "Nmi"
	// Pause execution on the unit but do not remove power.
	// This is typically a feature of virtual machine hypervisors
	PauseResetType ResetType = "Pause"
	// Resume execution on the paused unit.
	// This is typically a feature of virtual machine hypervisors
	ResumeResetType ResetType = "Resume"
	// Write the state of the unit to disk before powering off.
	// This allows for the state to be restored when powered back on
	SuspendResetType ResetType = "Suspend"
)

// SerialConsoleProtocol shall describe a serial console service for a computer system.
type SerialConsoleProtocol struct {
	// ConsoleEntryCommand shall contain a command string that can be provided by a client to select or enter the
	// system's serial console, when the console is shared among several systems or a manager CLI.
	ConsoleEntryCommand string
	// HotKeySequenceDisplay shall contain a string that can be provided to a user to describe the hotkey sequence used
	// to exit the serial console session, or, if shared with a manager CLI, to return to the CLI.
	HotKeySequenceDisplay string
	// Port shall contain the port assigned to the protocol.
	Port int
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled bool
	// SharedWithManagerCLI shall indicate whether the serial console service is shared with access to the manager's
	// command-line interface (CLI).
	SharedWithManagerCLI bool
}

// HostSerialConsole shall describe the serial console services for a computer system.
type HostSerialConsole struct {
	// IPMI shall contain connection details for a serial console service that uses the IPMI Serial-over-LAN (SOL)
	// protocol.
	IPMI SerialConsoleProtocol
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions int
	// SSH shall contain connection details for a serial console service that uses the Secure Shell (SSH) protocol.
	SSH SerialConsoleProtocol
	// Telnet shall contain connection details for a serial console service that uses the Telnet protocol.
	Telnet SerialConsoleProtocol
}

// IdlePowerSaver shall contain the idle power saver settings of a computer system.
type IdlePowerSaver struct {
	// Enabled shall indicate if idle power saver is enabled.
	Enabled bool
	// EnterDwellTimeSeconds shall contain the duration in seconds the computer system is below the
	// EnterUtilizationPercent value before the idle power save is activated.
	EnterDwellTimeSeconds int
	// EnterUtilizationPercent shall contain the percentage of utilization, typically '0' to '100', when the computer
	// system enters idle power save. If the computer system's utilization goes below this value for the duration
	// specified by EnterDwellTimeSeconds, it shall enter idle power save.
	EnterUtilizationPercent int
	// ExitDwellTimeSeconds shall contain the duration in seconds the computer system is above the
	// ExitUtilizationPercent value before the idle power save is stopped.
	ExitDwellTimeSeconds int
	// ExitUtilizationPercent shall contain the percentage of utilization, typically '0' to '100', when the computer
	// system exits idle power save. If the computer system's utilization goes above this value for the duration
	// specified by ExitDwellTimeSeconds, it shall exit idle power save.
	ExitUtilizationPercent int
}

// KMIPServer shall contain the KMIP server settings for a computer system.
type KMIPServer struct {
	// Address shall contain the KMIP server address.
	Address string
	// CacheDuration shall contain the duration that the system caches KMIP data.
	CacheDuration string
	// CachePolicy shall contain the cache policy to control how KMIP data is cached.
	CachePolicy KMIPCachePolicy
	// Password shall contain the password to access the KMIP server. The value shall be 'null' in responses.
	Password string
	// Port shall contain the KMIP server port.
	Port int
	// Username shall contain the username to access the KMIP server.
	Username string
}

// KeyManagement shall contain the key management settings of a computer system.
type KeyManagement struct {
	// KMIPCertificates shall contain a link to a resource collection of type CertificateCollection that represents the
	// server certificates for the servers referenced by the KMIPServers property.
	kmipCertificates string
	// KMIPServers shall contain the KMIP servers to which this computer system is subscribed for key management.
	KMIPServers []KMIPServer
}

// UnmarshalJSON unmarshals a KeyManagement object from the raw JSON.
func (keymanagement *KeyManagement) UnmarshalJSON(b []byte) error {
	type temp KeyManagement
	var t struct {
		temp
		KMIPCertificates common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keymanagement = KeyManagement(t.temp)

	// Extract the links to other entities for later
	keymanagement.kmipCertificates = t.KMIPCertificates.String()

	return nil
}

// ComputerSystem is used to represent resources that represent a
// computing system in the Redfish specification.
type ComputerSystem struct {
	common.Entity

	// ODataContext is the @odata.context
	ODataContext string `json:"@odata.context"`
	// ODataType is the @odata.type
	ODataType string `json:"@odata.type"`

	// AssetTag shall contain the value of the asset tag of the system.
	AssetTag string
	// bios shall be a link to a resource of
	// type Bios that lists the Bios settings for this system.
	bios string
	// BIOSVersion shall be the version string
	// of the currently installed and running BIOS (for x86 systems). For
	// other systems, the value may contain a version string representing the
	// primary system firmware.
	BIOSVersion string `json:"BiosVersion"`
	// Boot describes boot information for the current resource.
	Boot Boot
	// BootProgress shall contain the last boot progress state and time.
	BootProgress BootProgress
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates string
	// Composition shall contain information about the composition capabilities and state of the computer system.
	Composition Composition
	// Description is the resource description.
	Description string
	// EthernetInterfaces shall be a link to a collection of type EthernetInterfaceCollection.
	ethernetInterfaces string
	// FabricAdapters shall contain a link to a resource collection of type FabricAdapterCollection.
	fabricAdapters string
	// GraphicalConsole shall contain the information about the graphical console (KVM-IP) service of this system.
	GraphicalConsole HostGraphicalConsole
	// GraphicsControllers shall contain a link to a resource collection of type GraphicsControllerCollection that
	// contains graphics controllers that can output video for this system.
	graphicsControllers string
	// HostName shall be the host name for this
	// system, as reported by the operating system or hypervisor. This value
	// is typically provided to the Manager by a service running in the host
	// operating system.
	HostName string
	// HostWatchdogTimer shall contain properties which
	// describe the host watchdog timer functionality for this
	// ComputerSystem.
	HostWatchdogTimer WatchdogTimer
	// HostedServices shall describe services that this computer system supports.
	HostedServices HostedServices
	// HostingRoles shall be the hosting roles supported by this computer system.
	HostingRoles []string
	// IdlePowerSaver shall contain the idle power saver settings of the computer system.
	IdlePowerSaver IdlePowerSaver
	// KeyManagement shall contain the key management settings of the computer system.
	KeyManagement KeyManagement
	// LastResetTime shall contain the date and time when the system last came out of a reset or was rebooted.
	LastResetTime string
	// IndicatorLED shall contain the indicator
	// light state for the indicator light associated with this system.
	IndicatorLED common.IndicatorLED
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function. Modifying this property may modify the
	// LocationIndicatorActive in the containing Chassis resource.
	LocationIndicatorActive bool
	// logServices shall be a link to a collection of type LogServiceCollection.
	logServices string
	// Manufacturer shall contain a value that represents the manufacturer of the system.
	Manufacturer string
	// Memory shall be a link to a collection of type MemoryCollection.
	memory string
	// memoryDomains shall be a link to a collection of type MemoryDomainCollection.
	memoryDomains string
	// MemorySummary is This object shall contain properties which describe
	// the central memory for the current resource.
	MemorySummary MemorySummary
	// Model shall contain the information
	// about how the manufacturer references this system. This is typically
	// the product name, without the manufacturer name.
	Model string
	// networkInterfaces shall be a link to a collection of type
	// NetworkInterfaceCollection.
	networkInterfaces string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSystem shall contain a link to a resource of type OperatingSystem that contains operating system
	// information for this system.
	operatingSystem string
	// PCIeDevices shall be an array of references of type PCIeDevice.
	pcieDevices []string
	// PCIeDevicesCount is the number of PCIeDevices.
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// PCIeFunctions shall be an array of references of type PCIeFunction.
	pcieFunctions []string
	// PCIeFunctionsCount is the number of PCIeFunctions.
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// PartNumber shall contain the part number
	// for the system as defined by the manufacturer.
	PartNumber string
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a 'Reset' action requesting
	// 'PowerCycle'. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerMode shall contain the computer system power mode setting.
	PowerMode PowerMode
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off during a reset. The value '0' shall
	// indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power on after a power cycle or during a reset.
	// The value '0' shall indicate no delay to power on.
	PowerOnDelaySeconds float64
	// PowerRestorePolicy is the desired
	// PowerState of the system when power is applied to the system. A value
	// of 'LastState' shall return the system to the PowerState it was in
	// when power was lost.
	PowerRestorePolicy PowerState
	// PowerState shall contain the power state of the system.
	PowerState PowerState
	// ProcessorSummary shall contain properties which
	// describe the central processors for the current resource.
	ProcessorSummary ProcessorSummary
	// Processors shall be a link to a collection of type ProcessorCollection.
	processors string
	// Redundancy references a redundancy
	// entity that specifies a kind and level of redundancy and a collection
	// (RedundancySet) of other ComputerSystems that provide the specified
	// redundancy to this ComputerSystem.
	redundancy string
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount string `json:"Redundancy@odata.count"`
	// SKU shall contain the Stock Keeping Unit (SKU) for the system.
	SKU string
	// secureBoot shall be a link to a resource of type SecureBoot.
	secureBoot string
	// SerialConsole shall contain information about the serial console services of this system.
	SerialConsole HostSerialConsole
	// SerialNumber shall contain the serial number for the system.
	SerialNumber string
	// SimpleStorage shall be a link to a collection of type SimpleStorageCollection.
	simpleStorage string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// storage shall be a link to a collection
	// of type StorageCollection.
	storage string
	// SubModel shall contain the information about the sub-model (or configuration) of the system. This shall not
	// include the model/product name or the manufacturer name.
	SubModel string
	// SystemType An enumeration that indicates the kind of system that this resource represents.
	SystemType SystemType
	// USBControllers shall contain a link to a resource collection of type USBControllerCollection that contains USB
	// controllers for this system.
	usbControllers string
	// UUID shall contain the universally unique identifier number for this system. RFC4122 describes methods to create
	// this value. The value should be considered to be opaque. Client software should only treat the overall value as
	// a UUID and should not interpret any subfields within the UUID. If the system supports SMBIOS, the property value
	// should follow the SMBIOS 2.6 and later recommendation for converting the SMBIOS 16-byte UUID structure into the
	// Redfish canonical 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' string format, so that the property value matches the
	// byte order presented by current OS APIs, such as WMI and dmidecode.
	UUID string
	// virtualMedia shall contain a reference to a collection of type
	// VirtualMediaCollection which are for the use of this system.
	virtualMedia string
	// TrustedModules shall contain an array of objects with
	// properties which describe the trusted modules for the current resource.
	// This property has been deprecated in favor of the TrustedComponents property in Links.
	TrustedModules []TrustedModules

	// Chassis is an array of references to the chassis in which this system is contained.
	chassis []string
	// SupportedResetTypes, if provided, is the reset types this system supports.
	SupportedResetTypes []ResetType
	// settingsApplyTimes is a set of allowed settings update apply times. If none
	// are specified, then the system does not provide that information.
	settingsApplyTimes []common.ApplyTime
	managedBy          []string

	// addResourceBlockTarget is the internal URL for the AddResourceBlock action.
	addResourceBlockTarget string
	// decommissionTarget is the URL for the Decommission action.
	decommissionTarget string
	// removeResourceBlockTarget is the URL for the RemoveResourceBlock action.
	removeResourceBlockTarget string
	// resetTarget is the internal URL to send reset targets to.
	resetTarget string
	// setDefaultBootOrderTarget is the URL to send SetDefaultBootOrder actions to.
	setDefaultBootOrderTarget string
	settingsTarget            string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
func (computersystem *ComputerSystem) UnmarshalJSON(b []byte) error {
	type CSActions struct {
		AddResourceBlock    common.ActionTarget `json:"#ComputerSystem.AddResourceBlock"`
		Decommission        common.ActionTarget `json:"#ComputerSystem.Decommission"`
		RemoveResourceBlock common.ActionTarget `json:"#ComputerSystem.RemoveResourceBlock"`
		Reset               struct {
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			Target            string
		} `json:"#ComputerSystem.Reset"`
		SetDefaultBootOrder common.ActionTarget `json:"#ComputerSystem.SetDefaultBootOrder"`
	}

	type temp ComputerSystem
	var t struct {
		temp
		Actions             CSActions
		Bios                common.Link
		Certificates        common.Link
		EthernetInterfaces  common.Link
		FabricAdapters      common.Link
		GraphicsControllers common.Link
		Processors          common.Link
		Redundancy          common.Link
		Memory              common.Link
		OperatingSystem     common.Link
		SimpleStorage       common.Link
		SecureBoot          common.Link
		Storage             common.Link
		NetworkInterfaces   common.Link
		LogServices         common.Link
		MemoryDomains       common.Link
		PCIeDevices         common.Links
		PCIeFunctions       common.Links
		USBControllers      common.Link
		VirtualMedia        common.Link
		Links               CSLinks
		Settings            common.Settings `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*computersystem = ComputerSystem(t.temp)

	// Extract the links to other entities for later
	computersystem.bios = t.Bios.String()
	computersystem.certificates = t.Certificates.String()
	computersystem.ethernetInterfaces = t.EthernetInterfaces.String()
	computersystem.fabricAdapters = t.FabricAdapters.String()
	computersystem.graphicsControllers = t.GraphicsControllers.String()
	computersystem.logServices = t.LogServices.String()
	computersystem.pcieDevices = t.PCIeDevices.ToStrings()
	computersystem.pcieFunctions = t.PCIeFunctions.ToStrings()
	computersystem.processors = t.Processors.String()
	computersystem.memory = t.Memory.String()
	computersystem.memoryDomains = t.MemoryDomains.String()
	computersystem.networkInterfaces = t.NetworkInterfaces.String()
	computersystem.operatingSystem = t.OperatingSystem.String()
	computersystem.redundancy = t.Redundancy.String()
	computersystem.secureBoot = t.SecureBoot.String()
	computersystem.simpleStorage = t.SimpleStorage.String()
	computersystem.storage = t.Storage.String()
	computersystem.usbControllers = t.USBControllers.String()
	computersystem.virtualMedia = t.VirtualMedia.String()

	computersystem.addResourceBlockTarget = t.Actions.AddResourceBlock.Target
	computersystem.decommissionTarget = t.Actions.Decommission.Target
	computersystem.removeResourceBlockTarget = t.Actions.RemoveResourceBlock.Target
	computersystem.resetTarget = t.Actions.Reset.Target
	computersystem.SupportedResetTypes = t.Actions.Reset.AllowedResetTypes
	computersystem.setDefaultBootOrderTarget = t.Actions.SetDefaultBootOrder.Target

	computersystem.chassis = t.Links.Chassis.ToStrings()
	computersystem.managedBy = t.Links.ManagedBy.ToStrings()
	computersystem.settingsApplyTimes = t.Settings.SupportedApplyTimes

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	computersystem.settingsTarget = t.Settings.SettingsObject.String()
	if computersystem.settingsTarget == "" {
		computersystem.settingsTarget = computersystem.ODataID
	}

	// This is a read/write object, so we need to save the raw object data for later
	computersystem.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (computersystem *ComputerSystem) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	cs := new(ComputerSystem)
	err := cs.UnmarshalJSON(computersystem.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"HostName",
		"LocationIndicatorActive",
		"PowerCycleDelaySeconds",
		"PowerMode",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestorePolicy",
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(cs).Elem()
	currentElement := reflect.ValueOf(computersystem).Elem()

	return computersystem.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetComputerSystem will get a ComputerSystem instance from the service.
func GetComputerSystem(c common.Client, uri string) (*ComputerSystem, error) {
	return common.GetObject[ComputerSystem](c, uri)
}

// ListReferencedComputerSystems gets the collection of ComputerSystem from
// a provided reference.
func ListReferencedComputerSystems(c common.Client, link string) ([]*ComputerSystem, error) {
	return common.GetCollectionObjects[ComputerSystem](c, link)
}

// Bios gets the Bios information for this ComputerSystem.
func (computersystem *ComputerSystem) Bios() (*Bios, error) {
	if computersystem.bios == "" {
		return nil, nil
	}

	return GetBios(computersystem.GetClient(), computersystem.bios)
}

// BootOptions gets all BootOption items for this system.
func (computersystem *ComputerSystem) BootOptions() ([]*BootOption, error) {
	return common.GetCollectionObjects[BootOption](
		computersystem.GetClient(),
		computersystem.Boot.bootOptions)
}

func (computersystem *ComputerSystem) BootCertificates() ([]*Certificate, error) {
	return ListReferencedCertificates(computersystem.GetClient(), computersystem.Boot.certificates)
}

// EthernetInterfaces get this system's ethernet interfaces.
func (computersystem *ComputerSystem) EthernetInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(computersystem.GetClient(), computersystem.ethernetInterfaces)
}

// LogServices get this system's log services.
func (computersystem *ComputerSystem) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(computersystem.GetClient(), computersystem.logServices)
}

// ManagedBy gets all Managers for this system.
func (computersystem *ComputerSystem) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](computersystem.GetClient(), computersystem.managedBy)
}

// Memory gets this system's memory.
func (computersystem *ComputerSystem) Memory() ([]*Memory, error) {
	return ListReferencedMemorys(computersystem.GetClient(), computersystem.memory)
}

// MemoryDomains gets this system's memory domains.
func (computersystem *ComputerSystem) MemoryDomains() ([]*MemoryDomain, error) {
	return ListReferencedMemoryDomains(computersystem.GetClient(), computersystem.memoryDomains)
}

// NetworkInterfaces returns a collection of network interfaces in this system.
func (computersystem *ComputerSystem) NetworkInterfaces() ([]*NetworkInterface, error) {
	return ListReferencedNetworkInterfaces(computersystem.GetClient(), computersystem.networkInterfaces)
}

// OperatingSystem gets this system's operating system.
func (computersystem *ComputerSystem) OperatingSystem() (*OperatingSystem, error) {
	return GetOperatingSystem(computersystem.GetClient(), computersystem.operatingSystem)
}

// PCIeDevices gets all PCIeDevices for this system.
func (computersystem *ComputerSystem) PCIeDevices() ([]*PCIeDevice, error) {
	return common.GetObjects[PCIeDevice](computersystem.GetClient(), computersystem.pcieDevices)
}

// PCIeFunctions gets all PCIeFunctions for this system.
func (computersystem *ComputerSystem) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](computersystem.GetClient(), computersystem.pcieFunctions)
}

// Processors returns a collection of processors from this system
func (computersystem *ComputerSystem) Processors() ([]*Processor, error) {
	return ListReferencedProcessors(computersystem.GetClient(), computersystem.processors)
}

// SecureBoot gets the secure boot information for the system.
func (computersystem *ComputerSystem) SecureBoot() (*SecureBoot, error) {
	if computersystem.secureBoot == "" {
		return nil, nil
	}

	return GetSecureBoot(computersystem.GetClient(), computersystem.secureBoot)
}

// SetBoot set a boot object based on a payload request
func (computersystem *ComputerSystem) SetBoot(b Boot) error { //nolint
	t := struct {
		Boot Boot
	}{Boot: b}
	return computersystem.Patch(computersystem.ODataID, t)
}

// Reset shall perform a reset of the ComputerSystem. For systems which implement
// ACPI Power Button functionality, the PushPowerButton value shall perform or
// emulate an ACPI Power Button push. The ForceOff value shall remove power from
// the system or perform an ACPI Power Button Override (commonly known as a
// 4-second hold of the Power Button). The ForceRestart value shall perform a
// ForceOff action followed by a On action.
func (computersystem *ComputerSystem) Reset(resetType ResetType) error {
	// Make sure the requested reset type is supported by the system
	valid := false
	if len(computersystem.SupportedResetTypes) > 0 {
		for _, allowed := range computersystem.SupportedResetTypes {
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
		return fmt.Errorf("reset type '%s' is not supported by this service",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}

	return computersystem.Post(computersystem.resetTarget, t)
}

// UpdateBootAttributesApplyAt is used to update attribute values and set apply time together
func (computersystem *ComputerSystem) UpdateBootAttributesApplyAt(attrs SettingsAttributes, applyTime common.ApplyTime) error { //nolint:dupl
	payload := make(map[string]interface{})

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(computersystem.RawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := computersystem.GetClient().Get(computersystem.settingsTarget)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]interface{}{"Boot": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = computersystem.GetClient().PatchWithHeaders(computersystem.settingsTarget, data, header)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil
}

// UpdateBootAttributes is used to update attribute values.
func (computersystem *ComputerSystem) UpdateBootAttributes(attrs SettingsAttributes) error {
	return computersystem.UpdateBootAttributesApplyAt(attrs, "")
}

// SetDefaultBootOrder shall set the BootOrder array to the default settings.
func (computersystem *ComputerSystem) SetDefaultBootOrder() error {
	// This action wasn't added until 1.5.0, make sure this is supported.
	if computersystem.setDefaultBootOrderTarget == "" {
		return fmt.Errorf("SetDefaultBootOrder is not supported by this system") //nolint:golint
	}

	return computersystem.Post(computersystem.setDefaultBootOrderTarget, nil)
}

// SimpleStorages gets all simple storage services of this system.
func (computersystem *ComputerSystem) SimpleStorages() ([]*SimpleStorage, error) {
	return ListReferencedSimpleStorages(computersystem.GetClient(), computersystem.simpleStorage)
}

// Storage gets the storage associated with this system.
func (computersystem *ComputerSystem) Storage() ([]*Storage, error) {
	return ListReferencedStorages(computersystem.GetClient(), computersystem.storage)
}

// VirtualMedia gets the virtual media associated with this system.
func (computersystem *ComputerSystem) VirtualMedia() ([]*VirtualMedia, error) {
	return ListReferencedVirtualMedias(computersystem.GetClient(), computersystem.virtualMedia)
}

// USBControllers gets the USB controllers associated with this system.
func (computersystem *ComputerSystem) USBControllers() ([]*USBController, error) {
	return ListReferencedUSBControllers(computersystem.GetClient(), computersystem.usbControllers)
}

// CSLinks are references to resources that are related to, but not contained
// by (subordinate to), this resource.
type CSLinks struct {
	// Chassis shall reference a resource of
	// type Chassis that represents the physical container associated with
	// this resource.
	Chassis common.Links
	// ChassisCount is
	ChassisCount int `json:"Chassis@odata.count"`
	// ConsumingComputerSystems shall be an array of references
	// to ComputerSystems that are realized, in whole or in part, from this
	// ComputerSystem.
	ConsumingComputerSystems common.Links
	// ConsumingComputerSystemsCount is
	ConsumingComputerSystemsCount int `json:"ConsumingComputerSystems@odata.count"`
	// CooledBy shall be an array of IDs
	// containing pointers consistent with JSON pointer syntax to the
	// resource that powers this computer system.
	CooledBy common.Links
	// CooledByCount is
	CooledByCount int `json:"CooledBy@odata.count"`
	// Endpoints shall be a reference to the
	// resources that this system is associated with and shall reference a
	// resource of type Endpoint.
	Endpoints common.Links
	// EndpointsCount is
	EndpointsCount int `json:"Endpoints@odata.count"`
	// ManagedBy shall reference a resource of
	// type manager that represents the resource with management
	// responsibility for this resource.
	ManagedBy common.Links
	// ManagedByCount is
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// PoweredBy shall be an array of IDs
	// containing pointers consistent with JSON pointer syntax to the
	// resource that powers this computer system.
	PoweredBy common.Links
	// PoweredByCount is the number of PoweredBy objects.
	PoweredByCount int `json:"PoweredBy@odata.count"`
	// ResourceBlocks is used in this Computer System.
	ResourceBlocks common.Links
	// ResourceBlocksCount is the number of ResourceBlocks.
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
	// SupplyingComputerSystems shall be an array of references
	// to ComputerSystems that contribute, in whole or in part, to the
	// implementation of this ComputerSystem.
	SupplyingComputerSystems common.Links
	// SupplyingComputerSystemsCount is the number of SupplyingComputerSystems.
	SupplyingComputerSystemsCount int `json:"SupplyingComputerSystems@odata.count"`
}

// MemorySummary contains properties which describe the central memory for a system.
type MemorySummary struct {
	// MemoryMirroring is the ability and type of memory mirroring supported by this system.
	MemoryMirroring MemoryMirroring
	// Metrics shall be a reference to the Metrics
	// associated with this MemorySummary.
	metrics string
	// Status is the status or health properties of the resource.
	Status common.Status
	// TotalSystemMemoryGiB is the amount of configured system general purpose
	// volatile (RAM) memory as measured in gibibytes.
	TotalSystemMemoryGiB float32
	// TotalSystemPersistentMemoryGiB is the total amount of configured
	// persistent memory available to the system as measured in gibibytes.
	TotalSystemPersistentMemoryGiB float32
}

func (memorySummary *MemorySummary) UnmarshalJSON(b []byte) error {
	type temp MemorySummary
	type t1 struct {
		temp
		Metrics common.Link
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorySummary = MemorySummary(t.temp)

	memorySummary.metrics = t.Metrics.String()

	return nil
}

// Metrics gets the memory summary metrics
func (memorySummary *MemorySummary) Metrics(c common.Client) (*MemoryMetrics, error) {
	if memorySummary.metrics == "" {
		return nil, nil
	}
	return GetMemoryMetrics(c, memorySummary.metrics)
}

// ProcessorSummary is This type shall contain properties which describe
// the central processors for a system.
type ProcessorSummary struct {
	// Count is the number of physical central processors in the system.
	Count int
	// LogicalProcessorCount is the number of logical central processors in the system.
	LogicalProcessorCount int
	// Metrics shall be a reference to the Metrics
	// associated with this ProcessorSummary.
	metrics string
	// Model is the processor model for the central processors in the system,
	// per the description in the Processor Information - Processor Family
	// section of the SMBIOS Specification DSP0134 2.8 or later.
	Model string
	// Status is any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a ProcessorSummary object from the raw JSON.
func (processorSummary *ProcessorSummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorSummary
	type t1 struct {
		temp
		Metrics common.Link
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorSummary = ProcessorSummary(t.temp)

	processorSummary.metrics = t.Metrics.String()

	return nil
}

// Metrics gets the processor summary metrics
func (processorSummary *ProcessorSummary) Metrics(c common.Client) (*ProcessorMetrics, error) {
	if processorSummary.metrics == "" {
		return nil, nil
	}
	return GetProcessorMetrics(c, processorSummary.metrics)
}

// TrustedModules is This type shall describe a trusted module for a system.
type TrustedModules struct {
	// FirmwareVersion is the firmware version as
	// defined by the manufacturer for the Trusted Module.
	FirmwareVersion string
	// FirmwareVersion2 is the 2nd firmware
	// version, if applicable, as defined by the manufacturer for the Trusted
	// Module.
	FirmwareVersion2 string
	// InterfaceType is the interface type of the installed Trusted Module.
	InterfaceType InterfaceType
	// InterfaceTypeSelection is the Interface
	// Type Selection method (for example to switch between TPM1_2 and
	// TPM2_0) that is supported by this TrustedModule.
	InterfaceTypeSelection InterfaceTypeSelection
	// Status is any status or health properties
	// of the resource.
	Status common.Status
}

// WatchdogTimer contains properties which describe the
// host watchdog timer functionality for this ComputerSystem.
type WatchdogTimer struct {
	// FunctionEnabled shall indicate whether
	// the host watchdog timer functionality has been enabled or not. This
	// property indicates only that the functionality is enabled or disabled
	// by the user, and updates to this property shall not initiate a
	// watchdog timer countdown.
	FunctionEnabled bool
	// Status is any status or health properties of the resource.
	Status struct {
		State common.State
	}
	// TimeoutAction is the action to perform
	// upon the  expiration of the Watchdog Timer.
	TimeoutAction string
	// WarningAction is the action to perform
	// prior to the expiration of the Watchdog Timer. This action typically
	// occurs 3-10 seconds prior to the timeout value, but the exact timing
	// is dependent on the implementation.
	WarningAction string
}
