//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #Manager.v1_23_0.Manager

package redfish

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/stmcginnis/gofish/common"
)

type CommandConnectTypesSupported string

const (
	// SSHCommandConnectTypesSupported The controller supports a command shell
	// connection through the SSH protocol.
	SSHCommandConnectTypesSupported CommandConnectTypesSupported = "SSH"
	// TelnetCommandConnectTypesSupported The controller supports a command shell
	// connection through the Telnet protocol.
	TelnetCommandConnectTypesSupported CommandConnectTypesSupported = "Telnet"
	// IPMICommandConnectTypesSupported The controller supports a command shell
	// connection through the IPMI Serial Over LAN (SOL) protocol.
	IPMICommandConnectTypesSupported CommandConnectTypesSupported = "IPMI"
	// OemCommandConnectTypesSupported The controller supports a command shell
	// connection through an OEM-specific protocol.
	OemCommandConnectTypesSupported CommandConnectTypesSupported = "Oem"
)

type DateTimeSource string

const (
	// RTCDateTimeSource The date and time is retrieved from the manager's real
	// time clock (RTC).
	RTCDateTimeSource DateTimeSource = "RTC"
	// FirmwareDateTimeSource The date and time is set and held by firmware.
	FirmwareDateTimeSource DateTimeSource = "Firmware"
	// HostDateTimeSource The date and time is retrieved from the host.
	HostDateTimeSource DateTimeSource = "Host"
	// NTPDateTimeSource The date and time source is a Network Time Protocol (NTP)
	// server.
	NTPDateTimeSource DateTimeSource = "NTP"
	// PTPDateTimeSource The date and time source is a Precision Time Protocol
	// (PTP) server.
	PTPDateTimeSource DateTimeSource = "PTP"
)

type ManagerType string

const (
	// ManagementControllerManagerType is a controller that primarily monitors or
	// manages the operation of a device or system.
	ManagementControllerManagerType ManagerType = "ManagementController"
	// EnclosureManagerManagerType is a controller that provides management
	// functions for a chassis, group of devices, or group of systems with their
	// own BMCs (baseboard management controllers). An example of this is a manager
	// that aggregates and orchestrates management functions across multiple BMCs
	// in an enclosure.
	EnclosureManagerManagerType ManagerType = "EnclosureManager"
	// BMCManagerType is a controller that provides management functions for one or
	// more computer systems. Commonly known as a BMC (baseboard management
	// controller). Examples of this include a BMC dedicated to one system or a
	// multi-host manager providing BMC capabilities to multiple systems.
	BMCManagerType ManagerType = "BMC"
	// RackManagerManagerType is a controller that provides management functions
	// for a whole or part of a rack. An example of this is a manager that
	// aggregates and orchestrates management functions across multiple managers,
	// such as enclosure managers and BMCs (baseboard management controllers), in a
	// rack.
	RackManagerManagerType ManagerType = "RackManager"
	// AuxiliaryControllerManagerType is a controller that provides management
	// functions for a particular subsystem or group of devices as part of a larger
	// system.
	AuxiliaryControllerManagerType ManagerType = "AuxiliaryController"
	// ServiceManagerType is a software-based service that provides management
	// functions.
	ServiceManagerType ManagerType = "Service"
	// FabricManagerManagerType is a controller that primarily monitors or manages
	// the operation of a group of fabric attached nodes and switches.
	FabricManagerManagerType ManagerType = "FabricManager"
)

type ResetToDefaultsType string

const (
	// ResetAllResetToDefaultsType Reset all settings to factory defaults.
	ResetAllResetToDefaultsType ResetToDefaultsType = "ResetAll"
	// PreserveNetworkAndUsersResetToDefaultsType Reset all settings except network
	// and local usernames/passwords to factory defaults.
	PreserveNetworkAndUsersResetToDefaultsType ResetToDefaultsType = "PreserveNetworkAndUsers"
	// PreserveNetworkResetToDefaultsType Reset all settings except network
	// settings to factory defaults.
	PreserveNetworkResetToDefaultsType ResetToDefaultsType = "PreserveNetwork"
)

type SecurityModeTypes string

const (
	// FIPS1402SecurityModeTypes shall indicate that the implementation complies
	// with FIPS 140-2.
	FIPS1402SecurityModeTypes SecurityModeTypes = "FIPS_140_2"
	// FIPS1403SecurityModeTypes shall indicate that the implementation complies
	// with FIPS 140-3.
	FIPS1403SecurityModeTypes SecurityModeTypes = "FIPS_140_3"
	// CNSA10SecurityModeTypes shall indicate that the implementation meets NSA
	// Commercial National Security Algorithm Suite 1.0 requirements and FIPS 140-2
	// compliance.
	CNSA10SecurityModeTypes SecurityModeTypes = "CNSA_1_0"
	// CNSA20SecurityModeTypes shall indicate that the implementation meets NSA
	// Commercial National Security Algorithm Suite 2.0 requirements and FIPS 140-3
	// compliance.
	CNSA20SecurityModeTypes SecurityModeTypes = "CNSA_2_0"
	// SuiteBSecurityModeTypes shall indicate that the implementation meets NSA
	// Suite B cryptographic standards for Top Secret installations and FIPS 140-2
	// compliance.
	SuiteBSecurityModeTypes SecurityModeTypes = "SuiteB"
	// OEMSecurityModeTypes shall indicate that the implementation is in an
	// OEM-specific security state.
	OEMSecurityModeTypes SecurityModeTypes = "OEM"
	// DefaultSecurityModeTypes shall indicate that the implementation is in a
	// vendor-specific default security state that does not match any other value.
	DefaultSecurityModeTypes SecurityModeTypes = "Default"
)

type SerialConnectTypesSupported string

const (
	// SSHSerialConnectTypesSupported The controller supports a serial console
	// connection through the SSH protocol.
	SSHSerialConnectTypesSupported SerialConnectTypesSupported = "SSH"
	// TelnetSerialConnectTypesSupported The controller supports a serial console
	// connection through the Telnet protocol.
	TelnetSerialConnectTypesSupported SerialConnectTypesSupported = "Telnet"
	// IPMISerialConnectTypesSupported The controller supports a serial console
	// connection through the IPMI Serial Over LAN (SOL) protocol.
	IPMISerialConnectTypesSupported SerialConnectTypesSupported = "IPMI"
	// OemSerialConnectTypesSupported The controller supports a serial console
	// connection through an OEM-specific protocol.
	OemSerialConnectTypesSupported SerialConnectTypesSupported = "Oem"
)

// Manager shall represent a management subsystem for a Redfish implementation.
type Manager struct {
	common.Entity
	// AdditionalFirmwareVersions shall contain the additional firmware versions of
	// the manager.
	//
	// Version added: v1.15.0
	AdditionalFirmwareVersions AdditionalVersions
	// AutoDSTEnabled shall indicate whether the manager is configured for
	// automatic Daylight Saving Time (DST) adjustment.
	//
	// Version added: v1.4.0
	AutoDSTEnabled bool
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.13.0
	certificates string
	// CommandShell shall describe a command line user interface or command shell
	// service provided by this manager. The command shell refers to an interface
	// used to interact with the manager itself, not a dedicated console session
	// redirected from a host operating system. For redirected serial or host
	// operating system consoles, see the 'SerialConsole' property in the
	// 'ComputerSystem' resource.
	CommandShell CommandShell
	// DateTime shall contain the current date and time with UTC offset of the
	// manager.
	DateTime string
	// DateTimeLocalOffset shall contain the offset from UTC time that the
	// 'DateTime' property contains. If both 'DateTime' and 'DateTimeLocalOffset'
	// are provided in modification requests, services shall apply
	// 'DateTimeLocalOffset' after 'DateTime' is applied.
	DateTimeLocalOffset string
	// DateTimeSource shall contain the source of the 'DateTime' property of this
	// manager. The service shall update this property if the source changes
	// internally, for example if an NTP server is unavailable and the source falls
	// back to the time stored by the RTC.
	//
	// Version added: v1.20.0
	DateTimeSource DateTimeSource
	// DaylightSavingTime shall contain the daylight saving time settings for this
	// manager.
	//
	// Version added: v1.19.0
	DaylightSavingTime DaylightSavingTime
	// DedicatedNetworkPorts shall contain a link to a resource collection of type
	// 'PortCollection' that represent the dedicated network ports of the manager.
	//
	// Version added: v1.16.0
	dedicatedNetworkPorts string
	// EthernetInterfaces shall contain a link to a resource collection of type
	// 'EthernetInterfaceCollection'.
	ethernetInterfaces string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated manager.
	FirmwareVersion string
	// GraphicalConsole shall contain the information about the graphical console
	// (KVM-IP) service of this manager. This property should be used to describe a
	// service for the manager's console or operating system, not a service
	// provided on behalf of a host operating system. Implementations representing
	// host OS consoles, known generally as a KVM-IP feature, should use the
	// 'GraphicalConsole' property in the 'ComputerSystem' resource.
	GraphicalConsole GraphicalConsole
	// HostInterfaces shall contain a link to a resource collection of type
	// 'HostInterfaceCollection'.
	//
	// Version added: v1.3.0
	hostInterfaces string
	// LastResetTime shall contain the date and time when the manager last came out
	// of a reset or was rebooted.
	//
	// Version added: v1.9.0
	LastResetTime string
	// Location shall contain the location information of the associated manager.
	//
	// Version added: v1.11.0
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.11.0
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type
	// 'LogServiceCollection' that this manager uses.
	logServices string
	// ManagerDiagnosticData shall contain a link to a resource of type
	// 'ManagerDiagnosticData' that represents the diagnostic data for this
	// manager.
	//
	// Version added: v1.14.0
	managerDiagnosticData string
	// ManagerType shall describe the function of this manager. The
	// 'ManagementController' value shall be used if none of the other enumerations
	// apply.
	ManagerType ManagerType
	// Manufacturer shall contain the name of the organization responsible for
	// producing the manager. This organization may be the entity from whom the
	// manager is purchased, but this is not necessarily true.
	//
	// Version added: v1.7.0
	Manufacturer string
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.13.0
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// Model shall contain the information about how the manufacturer refers to
	// this manager.
	Model string
	// NetworkProtocol shall contain a link to a resource of type
	// 'ManagerNetworkProtocol', which represents the network services for this
	// manager.
	networkProtocol string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMSecurityMode shall the OEM-specific security compliance mode(s) that the
	// manager is currently configured to enforce. This property shall only be
	// present if 'SecurityMode' contains 'OEM'.
	//
	// Version added: v1.21.0
	OEMSecurityMode string
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the manager.
	//
	// Version added: v1.7.0
	PartNumber string
	// PowerState shall contain the power state of the manager.
	//
	// Version added: v1.2.0
	PowerState common.PowerState
	// ReadyToRemove shall indicate whether the manager is ready for removal.
	// Setting the value to 'true' shall cause the service to perform appropriate
	// actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	//
	// Version added: v1.23.0
	ReadyToRemove bool
	// Redundancy shall show how this manager is grouped with other managers for
	// form redundancy sets.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RemoteAccountService shall contain a link to the account service resource
	// for the remote manager that this resource represents. This property shall
	// only be present when providing aggregation of a remote manager.
	//
	// Version added: v1.5.0
	remoteAccountService string
	// RemoteRedfishServiceUri shall contain the URI of the Redfish service root
	// for the remote manager that this resource represents. This property shall
	// only be present when providing aggregation of Redfish services.
	//
	// Version added: v1.5.0
	RemoteRedfishServiceURI string
	// SecurityMode shall contain the security compliance mode that the manager is
	// currently configured to enforce.
	//
	// Version added: v1.21.0
	SecurityMode SecurityModeTypes
	// SecurityPolicy shall contain a link to a resource of type 'SecurityPolicy'
	// that contains the security policy settings for this manager.
	//
	// Version added: v1.16.0
	securityPolicy string
	// SerialConsole shall contain information about the serial console service of
	// this manager.
	//
	// Deprecated: v1.10.0
	// This property has been deprecated in favor of the 'SerialConsole' property
	// in the 'ComputerSystem' resource.
	SerialConsole SerialConsole
	// SerialInterfaces shall contain a link to a resource collection of type
	// 'SerialInterfaceCollection', which this manager uses.
	serialInterfaces string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the manager.
	//
	// Version added: v1.7.0
	SerialNumber string
	// ServiceEntryPointUUID shall contain the UUID of the Redfish service that is
	// hosted by this manager. Each manager providing an entry point to the same
	// Redfish service shall report the same UUID value, even though the name of
	// the property may imply otherwise. This property shall not be present if this
	// manager does not provide a Redfish service entry point.
	ServiceEntryPointUUID string
	// ServiceIdentification shall contain a vendor-provided or user-provided value
	// that identifies and associates a discovered Redfish service with a
	// particular product instance. If this manager provides the Redfish service,
	// the 'ServiceIdentification' property in the 'ServiceRoot' resource shall
	// contain the value of this property. This property shall only be present if
	// the manager provides the Redfish service. The value of this property is used
	// in conjunction with the 'Product' and 'Vendor' properties in 'ServiceRoot'
	// to match user credentials or other a priori product instance information
	// necessary for initial deployment to the correct, matching Redfish service.
	//
	// Version added: v1.15.0
	ServiceIdentification string
	// SharedNetworkPorts shall contain a link to a resource collection of type
	// 'PortCollection' that represent the shared network ports of the manager. The
	// members of this collection shall reference Port resources subordinate to
	// NetworkAdapter resources.
	//
	// Version added: v1.16.0
	sharedNetworkPorts string
	// SparePartNumber shall contain the spare part number of the manager.
	//
	// Version added: v1.11.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TimeZoneName shall contain the time zone of the manager. The time zone shall
	// be either the 'Name' or the 'Format' for the zone as defined in the IANA
	// Time Zone Database. The value of this property is used for display purposes,
	// especially to enhance the display of time. A Redfish service may not be able
	// to ensure accuracy and consistency between the 'DateTimeOffset' property and
	// this property. Therefore, to specify the correct time zone offset, see the
	// 'DateTimeOffset' property.
	//
	// Version added: v1.10.0
	TimeZoneName string
	// USBPorts shall contain a link to a resource collection of type
	// 'PortCollection' that represent the USB ports of the manager.
	//
	// Version added: v1.12.0
	uSBPorts string
	// UUID shall contain the UUID for the manager.
	UUID string
	// Version shall contain the hardware version of this manager as determined by
	// the vendor or supplier.
	//
	// Version added: v1.17.0
	Version string
	// VirtualMedia shall contain a link to a resource collection of type
	// 'VirtualMediaCollection', which this manager uses.
	//
	// Deprecated: v1.10.0
	// This property has been deprecated in favor of the 'VirtualMedia' property in
	// the 'ComputerSystem' resource.
	virtualMedia string
	// forceFailoverTarget is the URL to send ForceFailover requests.
	forceFailoverTarget string
	// modifyRedundancySetTarget is the URL to send ModifyRedundancySet requests.
	modifyRedundancySetTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// resetInfo contains URI for an ActionInfo Resource that describes this action.
	resetActionInfo string
	// SupportedResetTypes, if provided, is the reset types this system supports.
	SupportedResetTypes []common.ResetType
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget           string
	resetToDefaultsActionInfoTarget string
	SupportedResetToDefaultsTypes   []ResetToDefaultsType
	// updateSecurityModeTarget is the URL to send UpdateSecurityMode requests.
	updateSecurityModeTarget string
	// activeSoftwareImage is the URI for ActiveSoftwareImage.
	activeSoftwareImage string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// managerForChassis are the URIs for ManagerForChassis.
	managerForChassis []string
	// managerForFabrics are the URIs for ManagerForFabrics.
	managerForFabrics []string
	// managerForManagers are the URIs for ManagerForManagers.
	managerForManagers []string
	// managerForServers are the URIs for ManagerForServers.
	managerForServers []string
	// managerForSwitches are the URIs for ManagerForSwitches.
	managerForSwitches []string
	// managerInChassis is the URI for ManagerInChassis.
	managerInChassis string
	// selectedNetworkPort is the URI for SelectedNetworkPort.
	selectedNetworkPort string
	// softwareImages are the URIs for SoftwareImages.
	softwareImages []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (m *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	type mActions struct {
		ForceFailover       common.ActionTarget `json:"#Manager.ForceFailover"`
		ModifyRedundancySet common.ActionTarget `json:"#Manager.ModifyRedundancySet"`
		Reset               struct {
			common.ActionTarget
			AllowedResetTypes []common.ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Manager.Reset"`
		ResetToDefaults struct {
			common.ActionTarget
			AllowedResetTypes []ResetToDefaultsType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Manager.ResetToDefaults"`
		UpdateSecurityMode common.ActionTarget `json:"#Manager.UpdateSecurityMode"`
	}
	type mLinks struct {
		ActiveSoftwareImage common.Link  `json:"ActiveSoftwareImage"`
		ManagedBy           common.Links `json:"ManagedBy"`
		ManagerForChassis   common.Links `json:"ManagerForChassis"`
		ManagerForFabrics   common.Links `json:"ManagerForFabrics"`
		ManagerForManagers  common.Links `json:"ManagerForManagers"`
		ManagerForServers   common.Links `json:"ManagerForServers"`
		ManagerForSwitches  common.Links `json:"ManagerForSwitches"`
		ManagerInChassis    common.Link  `json:"ManagerInChassis"`
		SelectedNetworkPort common.Link  `json:"SelectedNetworkPort"`
		SoftwareImages      common.Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions               mActions
		Links                 mLinks
		Certificates          common.Link `json:"certificates"`
		DedicatedNetworkPorts common.Link `json:"dedicatedNetworkPorts"`
		EthernetInterfaces    common.Link `json:"ethernetInterfaces"`
		HostInterfaces        common.Link `json:"hostInterfaces"`
		LogServices           common.Link `json:"logServices"`
		ManagerDiagnosticData common.Link `json:"managerDiagnosticData"`
		NetworkProtocol       common.Link `json:"networkProtocol"`
		RemoteAccountService  common.Link `json:"remoteAccountService"`
		SecurityPolicy        common.Link `json:"securityPolicy"`
		SerialInterfaces      common.Link `json:"serialInterfaces"`
		SharedNetworkPorts    common.Link `json:"sharedNetworkPorts"`
		USBPorts              common.Link `json:"uSBPorts"`
		VirtualMedia          common.Link `json:"virtualMedia"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Manager(tmp.temp)

	// Extract the links to other entities for later
	m.forceFailoverTarget = tmp.Actions.ForceFailover.Target
	m.modifyRedundancySetTarget = tmp.Actions.ModifyRedundancySet.Target
	m.resetTarget = tmp.Actions.Reset.Target
	m.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	m.updateSecurityModeTarget = tmp.Actions.UpdateSecurityMode.Target
	m.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	m.managedBy = tmp.Links.ManagedBy.ToStrings()
	m.managerForChassis = tmp.Links.ManagerForChassis.ToStrings()
	m.managerForFabrics = tmp.Links.ManagerForFabrics.ToStrings()
	m.managerForManagers = tmp.Links.ManagerForManagers.ToStrings()
	m.managerForServers = tmp.Links.ManagerForServers.ToStrings()
	m.managerForSwitches = tmp.Links.ManagerForSwitches.ToStrings()
	m.managerInChassis = tmp.Links.ManagerInChassis.String()
	m.selectedNetworkPort = tmp.Links.SelectedNetworkPort.String()
	m.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	m.certificates = tmp.Certificates.String()
	m.dedicatedNetworkPorts = tmp.DedicatedNetworkPorts.String()
	m.ethernetInterfaces = tmp.EthernetInterfaces.String()
	m.hostInterfaces = tmp.HostInterfaces.String()
	m.logServices = tmp.LogServices.String()
	m.managerDiagnosticData = tmp.ManagerDiagnosticData.String()
	m.networkProtocol = tmp.NetworkProtocol.String()
	m.remoteAccountService = tmp.RemoteAccountService.String()
	m.securityPolicy = tmp.SecurityPolicy.String()
	m.serialInterfaces = tmp.SerialInterfaces.String()
	m.sharedNetworkPorts = tmp.SharedNetworkPorts.String()
	m.uSBPorts = tmp.USBPorts.String()
	m.virtualMedia = tmp.VirtualMedia.String()

	m.SupportedResetTypes = tmp.Actions.Reset.AllowedResetTypes
	m.resetActionInfo = tmp.Actions.Reset.ActionInfoTarget
	m.SupportedResetToDefaultsTypes = tmp.Actions.ResetToDefaults.AllowedResetTypes
	m.resetToDefaultsActionInfoTarget = tmp.Actions.ResetToDefaults.ActionInfoTarget

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *Manager) Update() error {
	readWriteFields := []string{
		"AdditionalFirmwareVersions",
		"AutoDSTEnabled",
		"CommandShell",
		"DateTime",
		"DateTimeLocalOffset",
		"DateTimeSource",
		"DaylightSavingTime",
		"DedicatedNetworkPorts",
		"GraphicalConsole",
		"Location",
		"LocationIndicatorActive",
		"Measurements",
		"ReadyToRemove",
		"Redundancy",
		"Redundancy@odata.count",
		"SerialConsole",
		"ServiceIdentification",
		"SharedNetworkPorts",
		"Status",
		"TimeZoneName",
		"USBPorts",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetManager will get a Manager instance from the service.
func GetManager(c common.Client, uri string) (*Manager, error) {
	return common.GetObject[Manager](c, uri)
}

// ListReferencedManagers gets the collection of Manager from
// a provided reference.
func ListReferencedManagers(c common.Client, link string) ([]*Manager, error) {
	return common.GetCollectionObjects[Manager](c, link)
}

// ForceFailover shall perform a forced failover of the manager's redundancy to
// the manager supplied as a parameter.
// newManager - This parameter shall contain the manager to which to fail over.
func (m *Manager) ForceFailover(newManager string) error {
	payload := make(map[string]any)
	payload["NewManager"] = newManager
	return m.Post(m.forceFailoverTarget, payload)
}

// ModifyRedundancySet operation shall add members to or remove members
// from a redundant group of managers.
// add - This parameter shall contain an array of managers to add to the
// redundancy set.
// remove - This parameter shall contain an array of managers to remove from
// the redundancy set.
func (m *Manager) ModifyRedundancySet(add []string, remove []string) error {
	payload := make(map[string]any)
	payload["Add"] = add
	payload["Remove"] = remove
	return m.Post(m.modifyRedundancySetTarget, payload)
}

// GetSupportedResetTypes returns any reset types that the Manager declares as supported
// via either ActionInfo or AllowableValues.
func (m *Manager) GetSupportedResetTypes() ([]common.ResetType, error) {
	if len(m.SupportedResetTypes) > 0 {
		return m.SupportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if m.resetActionInfo != "" {
		resetActionInfo, err := m.ResetActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringActionInfoDataTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			m.SupportedResetTypes = append(m.SupportedResetTypes, common.ResetType(val))
		}
	}

	return m.SupportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the Manager reset action if supported
func (m *Manager) ResetActionInfo() (*ActionInfo, error) {
	if m.resetActionInfo == "" {
		return nil, errors.New("Manager Reset ActionInfo not supported")
	}

	return common.GetObject[ActionInfo](m.GetClient(), m.resetActionInfo)
}

// Reset shall reset the manager. If this manager provides the Redfish
// service, the service shall send the action response before resetting to
// prevent client timeouts.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
func (m *Manager) Reset(resetType common.ResetType) error {
	if len(m.SupportedResetTypes) == 0 {
		if m.resetActionInfo != "" {
			// reset without confirming the type is supported by the manager.
			// done to minimize overhead though technically not as correct as first checking the supported reset types
			t := struct {
				ResetType common.ResetType
			}{ResetType: resetType}
			return m.Post(m.resetTarget, t)
		}
		// reset directly without reset type. HPE server has the behavior
		return m.Post(m.resetTarget, struct{}{})
	}
	// Make sure the requested reset type is supported by the manager.
	valid := slices.Contains(m.SupportedResetTypes, resetType)
	if !valid {
		return fmt.Errorf("reset type '%s' is not supported by this manager",
			resetType)
	}

	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return m.Post(m.resetTarget, payload)
}

// GetSupportedResetToDefaultsTypes returns any reset to defaults
// types that the Manager declares as supported via either ActionInfo or AllowableValues.
func (m *Manager) GetSupportedResetToDefaultsTypes() ([]ResetToDefaultsType, error) {
	if len(m.SupportedResetToDefaultsTypes) > 0 {
		return m.SupportedResetToDefaultsTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if m.resetToDefaultsActionInfoTarget != "" {
		resetActionInfo, err := m.ResetToDefaultsActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringActionInfoDataTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			m.SupportedResetToDefaultsTypes = append(m.SupportedResetToDefaultsTypes, ResetToDefaultsType(val))
		}
	}

	return m.SupportedResetToDefaultsTypes, nil
}

// ResetToDefaultsActionInfo returns the ActionInfo for the Manager ResetToDefaults action if supported
func (m *Manager) ResetToDefaultsActionInfo() (*ActionInfo, error) {
	if m.resetToDefaultsActionInfoTarget == "" {
		return nil, errors.New("Manager ResetToDefaults ActionInfo not supported")
	}

	return common.GetObject[ActionInfo](m.GetClient(), m.resetToDefaultsActionInfoTarget)
}

// ResetToDefaults shall reset the manager settings. This action may impact other
// resources.
// resetType - This parameter shall contain the type of reset to defaults.
func (m *Manager) ResetToDefaults(resetType ResetToDefaultsType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return m.Post(m.resetToDefaultsTarget, payload)
}

// UpdateSecurityMode shall update the security mode for the manager. Services may
// reset other settings to factory defaults. Services may require the
// 'ResetToDefaults' action to clear security settings. This action may impact
// other resources.
// oEMSecurityMode - This parameter shall contain the OEM-specific security
// mode to apply to the manager. This parameter shall be required if
// 'SecurityMode' is 'OEM'.
// securityMode - This parameter shall contain the security mode to apply to
// the manager.
func (m *Manager) UpdateSecurityMode(oEMSecurityMode string, securityMode SecurityModeTypes) error {
	payload := make(map[string]any)
	payload["OEMSecurityMode"] = oEMSecurityMode
	payload["SecurityMode"] = securityMode
	return m.Post(m.updateSecurityModeTarget, payload)
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (m *Manager) ActiveSoftwareImage(client common.Client) (*SoftwareInventory, error) {
	if m.activeSoftwareImage == "" {
		return nil, nil
	}
	return common.GetObject[SoftwareInventory](client, m.activeSoftwareImage)
}

// ManagedBy gets the ManagedBy linked resources.
func (m *Manager) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, m.managedBy)
}

// ManagerForChassis gets the ManagerForChassis linked resources.
func (m *Manager) ManagerForChassis(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, m.managerForChassis)
}

// ManagerForFabrics gets the ManagerForFabrics linked resources.
func (m *Manager) ManagerForFabrics(client common.Client) ([]*Fabric, error) {
	return common.GetObjects[Fabric](client, m.managerForFabrics)
}

// ManagerForManagers gets the ManagerForManagers linked resources.
func (m *Manager) ManagerForManagers(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, m.managerForManagers)
}

// ManagerForServers gets the ManagerForServers linked resources.
func (m *Manager) ManagerForServers(client common.Client) ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](client, m.managerForServers)
}

// ManagerForSwitches gets the ManagerForSwitches linked resources.
func (m *Manager) ManagerForSwitches(client common.Client) ([]*Switch, error) {
	return common.GetObjects[Switch](client, m.managerForSwitches)
}

// ManagerInChassis gets the ManagerInChassis linked resource.
func (m *Manager) ManagerInChassis(client common.Client) (*Chassis, error) {
	if m.managerInChassis == "" {
		return nil, nil
	}
	return common.GetObject[Chassis](client, m.managerInChassis)
}

// SelectedNetworkPort gets the SelectedNetworkPort linked resource.
func (m *Manager) SelectedNetworkPort(client common.Client) (*NetworkPort, error) {
	if m.selectedNetworkPort == "" {
		return nil, nil
	}
	return common.GetObject[NetworkPort](client, m.selectedNetworkPort)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (m *Manager) SoftwareImages(client common.Client) ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](client, m.softwareImages)
}

// Certificates gets the Certificates collection.
func (m *Manager) Certificates(client common.Client) ([]*Certificate, error) {
	if m.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, m.certificates)
}

// DedicatedNetworkPorts gets the DedicatedNetworkPorts collection.
func (m *Manager) DedicatedNetworkPorts(client common.Client) ([]*Port, error) {
	if m.dedicatedNetworkPorts == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, m.dedicatedNetworkPorts)
}

// EthernetInterfaces gets the EthernetInterfaces collection.
func (m *Manager) EthernetInterfaces(client common.Client) ([]*EthernetInterface, error) {
	if m.ethernetInterfaces == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EthernetInterface](client, m.ethernetInterfaces)
}

// HostInterfaces gets the HostInterfaces collection.
func (m *Manager) HostInterfaces(client common.Client) ([]*HostInterface, error) {
	if m.hostInterfaces == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[HostInterface](client, m.hostInterfaces)
}

// LogServices gets the LogServices collection.
func (m *Manager) LogServices(client common.Client) ([]*LogService, error) {
	if m.logServices == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[LogService](client, m.logServices)
}

// ManagerDiagnosticData gets the ManagerDiagnosticData linked resource.
func (m *Manager) ManagerDiagnosticData(client common.Client) (*ManagerDiagnosticData, error) {
	if m.managerDiagnosticData == "" {
		return nil, nil
	}
	return common.GetObject[ManagerDiagnosticData](client, m.managerDiagnosticData)
}

// NetworkProtocol gets the NetworkProtocol linked resource.
func (m *Manager) NetworkProtocol(client common.Client) (*ManagerNetworkProtocol, error) {
	if m.networkProtocol == "" {
		return nil, nil
	}
	return common.GetObject[ManagerNetworkProtocol](client, m.networkProtocol)
}

// RemoteAccountService gets the RemoteAccountService linked resource.
func (m *Manager) RemoteAccountService(client common.Client) (*AccountService, error) {
	if m.remoteAccountService == "" {
		return nil, nil
	}
	return common.GetObject[AccountService](client, m.remoteAccountService)
}

// SecurityPolicy gets the SecurityPolicy linked resource.
func (m *Manager) SecurityPolicy(client common.Client) (*SecurityPolicy, error) {
	if m.securityPolicy == "" {
		return nil, nil
	}
	return common.GetObject[SecurityPolicy](client, m.securityPolicy)
}

// SerialInterfaces gets the SerialInterfaces collection.
func (m *Manager) SerialInterfaces(client common.Client) ([]*SerialInterface, error) {
	if m.serialInterfaces == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[SerialInterface](client, m.serialInterfaces)
}

// SharedNetworkPorts gets the SharedNetworkPorts collection.
func (m *Manager) SharedNetworkPorts(client common.Client) ([]*Port, error) {
	if m.sharedNetworkPorts == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, m.sharedNetworkPorts)
}

// USBPorts gets the USBPorts collection.
func (m *Manager) USBPorts(client common.Client) ([]*Port, error) {
	if m.uSBPorts == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, m.uSBPorts)
}

// VirtualMedia gets the VirtualMedia collection.
func (m *Manager) VirtualMedia(client common.Client) ([]*VirtualMedia, error) {
	if m.virtualMedia == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[VirtualMedia](client, m.virtualMedia)
}

// CommandShell shall describe a command shell service for a manager.
type CommandShell struct {
	// ConnectTypesSupported shall contain an array of the enumerations. SSH shall
	// be included if the Secure Shell (SSH) protocol is supported. Telnet shall be
	// included if the Telnet protocol is supported. IPMI shall be included if the
	// IPMI Serial Over LAN (SOL) protocol is supported.
	ConnectTypesSupported []CommandConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	MaxConcurrentSessions uint
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	ServiceEnabled bool
}

// DaylightSavingTime shall contain the daylight saving time settings for a
// manager.
type DaylightSavingTime struct {
	// EndDateTime shall contain the end date and time with UTC offset of daylight
	// saving time for this manager. If daylight saving time is permanent, specify
	// a sufficiently distant end date and time. Services shall update the UTC
	// offset based on changes made to 'DateTimeLocalOffset'. This property shall
	// be read-only if the service contains time zone databases.
	//
	// Version added: v1.19.0
	EndDateTime string
	// OffsetMinutes shall contain the number of minutes added to the 'DateTime'
	// value when the 'DateTime' value is between the values of StartDateTime and
	// EndDateTime. This offset shall be applied only if AutoDSTEnabled is 'true'.
	// This property shall be read-only if the service contains time zone
	// databases.
	//
	// Version added: v1.19.0
	OffsetMinutes int
	// StartDateTime shall contain the start date and time with UTC offset of
	// daylight saving time for this manager. Services shall update the UTC offset
	// based on changes made to 'DateTimeLocalOffset'. This property shall be
	// read-only if the service contains time zone databases.
	//
	// Version added: v1.19.0
	StartDateTime string
	// TimeZoneName shall contain the time zone of the manager when daylight saving
	// time is in effect. When daylight saving time is in effect, the service shall
	// update the 'TimeZoneName' property in the root of the resource. When
	// daylight saving time is no longer in effect, the service shall restore the
	// original value of the 'TimeZoneName' property in the root of the resource.
	// The time zone shall be either the 'Name' or the 'Format' for the zone as
	// defined in the IANA Time Zone Database. The value of this property is used
	// for display purposes, especially to enhance the display of time. This
	// property shall be read-only if the service contains time zone databases.
	//
	// Version added: v1.19.0
	TimeZoneName string
}

// GraphicalConsole shall describe a graphical console service for a manager.
type GraphicalConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. RDP shall
	// be included if the Remote Desktop (RDP) protocol is supported. KVMIP shall
	// be included if a vendor-defined KVM-IP protocol is supported.
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	MaxConcurrentSessions uint
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	ServiceEnabled bool
}

// ManagerService The manager services, such as serial console, command shell,
// or graphical console service.
type ManagerService struct {
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	MaxConcurrentSessions uint
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	ServiceEnabled bool
}

// SerialConsole shall describe a serial console service for a manager.
type SerialConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. SSH shall
	// be included if the Secure Shell (SSH) protocol is supported. Telnet shall be
	// included if the Telnet protocol is supported. IPMI shall be included if the
	// IPMI Serial Over LAN (SOL) protocol is supported.
	ConnectTypesSupported []SerialConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service
	// sessions that this implementation supports.
	MaxConcurrentSessions uint
	// ServiceEnabled shall indicate whether the protocol for the service is
	// enabled.
	ServiceEnabled bool
}
