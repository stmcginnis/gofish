//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CommandConnectTypesSupported is the command connection type.
type CommandConnectTypesSupported string

const (

	// SSHCommandConnectTypesSupported The controller supports a Command
	// Shell connection using the SSH protocol.
	SSHCommandConnectTypesSupported CommandConnectTypesSupported = "SSH"
	// TelnetCommandConnectTypesSupported The controller supports a Command
	// Shell connection using the Telnet protocol.
	TelnetCommandConnectTypesSupported CommandConnectTypesSupported = "Telnet"
	// IPMICommandConnectTypesSupported The controller supports a Command
	// Shell connection using the IPMI Serial-over-LAN (SOL) protocol.
	IPMICommandConnectTypesSupported CommandConnectTypesSupported = "IPMI"
	// OemCommandConnectTypesSupported The controller supports a Command
	// Shell connection using an OEM-specific protocol.
	OemCommandConnectTypesSupported CommandConnectTypesSupported = "Oem"
)

// GraphicalConnectTypesSupported is graphical connection type.
type GraphicalConnectTypesSupported string

const (

	// KVMIPGraphicalConnectTypesSupported The controller supports a
	// Graphical Console connection using a KVM-IP (redirection of Keyboard,
	// Video, Mouse over IP) protocol.
	KVMIPGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "KVMIP"
	// OemGraphicalConnectTypesSupported The controller supports a Graphical
	// Console connection using an OEM-specific protocol.
	OemGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "Oem"
)

// UIConsoleInfo contains information about GUI services.
type UIConsoleInfo struct {
	ServiceEnabled        bool
	MaxConcurrentSessions uint
	ConnectTypesSupported []string
}

// DaylightSavingTime shall contain the daylight saving time settings for a manager.
type DaylightSavingTime struct {
	// EndDateTime shall contain the end date and time with UTC offset of daylight saving time for this manager. If
	// daylight saving time is permanent, specify a sufficiently distant end date and time. Services shall update the
	// UTC offset based on changes made to DateTimeLocalOffset. This property shall be read-only if the service
	// contains time zone databases.
	EndDateTime string
	// OffsetMinutes shall contain the number of minutes added to the DateTime value when the DateTime value is between
	// the values of StartDateTime and EndDateTime. This offset shall be applied only if AutoDSTEnabled is 'true'. This
	// property shall be read-only if the service contains time zone databases.
	OffsetMinutes int
	// StartDateTime shall contain the start date and time with UTC offset of daylight saving time for this manager.
	// Services shall update the UTC offset based on changes made to DateTimeLocalOffset. This property shall be read-
	// only if the service contains time zone databases.
	StartDateTime string
	// TimeZoneName shall contain the time zone of the manager when daylight saving time is in effect. When daylight
	// saving time is in effect, the service shall update the TimeZoneName property in the root of the resource. When
	// daylight saving time is no longer in effect, the service shall restore the original value of the TimeZoneName
	// property in the root of the resource. The time zone shall be either the 'Name' or the 'Format' for the zone as
	// defined in the IANA Time Zone Database. The value of this property is used for display purposes, especially to
	// enhance the display of time. This property shall be read-only if the service contains time zone databases.
	TimeZoneName string
}

// SerialConsole shall describe a Serial Console service of a manager.
type SerialConsole struct {
	// ConnectTypesSupported shall be an array of the enumerations provided
	// here. SSH shall be included if the Secure Shell (SSH) protocol is
	// supported. Telnet shall be included if the Telnet protocol is supported.
	// IPMI shall be included if the IPMI (Serial-over-LAN) protocol is supported.
	ConnectTypesSupported []SerialConnectTypesSupported
	// MaxConcurrentSessions shall contain the
	// maximum number of concurrent service sessions supported by the
	// implementation.
	MaxConcurrentSessions int
	// ServiceEnabled is used for the service. The value shall be true if
	// enabled and false if disabled.
	ServiceEnabled bool
}

// ManagerType shall describe the function of this manager. The value
// EnclosureManager shall be used if this manager controls one or more services
// through aggregation. The value BMC shall be used if this manager represents a
// traditional server management controller. The value ManagementController
// shall be used if none of the other enumerations apply.
type ManagerType string

const (

	// ManagementControllerManagerType A controller used primarily to monitor
	// or manage the operation of a device or system.
	ManagementControllerManagerType ManagerType = "ManagementController"
	// EnclosureManagerManagerType A controller which provides management
	// functions for a chassis or group of devices or systems.
	EnclosureManagerManagerType ManagerType = "EnclosureManager"
	// BMCManagerType A controller which provides management functions for a
	// single computer system.
	BMCManagerType ManagerType = "BMC"
	// RackManagerManagerType A controller which provides management
	// functions for a whole or part of a rack.
	RackManagerManagerType ManagerType = "RackManager"
	// AuxiliaryControllerManagerType A controller which provides management
	// functions for a particular subsystem or group of devices.
	AuxiliaryControllerManagerType ManagerType = "AuxiliaryController"
	// ServiceManagerType A software-based service which provides management
	// functions.
	ServiceManagerType ManagerType = "Service"
)

// ResetToDefaultsType is the default to set on reset.
type ResetToDefaultsType string

const (

	// ResetAllResetToDefaultsType Reset all settings to factory defaults.
	ResetAllResetToDefaultsType ResetToDefaultsType = "ResetAll"
	// PreserveNetworkAndUsersResetToDefaultsType Reset all settings except
	// network and local user names/passwords to factory defaults.
	PreserveNetworkAndUsersResetToDefaultsType ResetToDefaultsType = "PreserveNetworkAndUsers"
	// PreserveNetworkResetToDefaultsType Reset all settings except network
	// settings to factory defaults.
	PreserveNetworkResetToDefaultsType ResetToDefaultsType = "PreserveNetwork"
)

// SerialConnectTypesSupported is serial connection type.
type SerialConnectTypesSupported string

const (

	// SSHSerialConnectTypesSupported The controller supports a Serial
	// Console connection using the SSH protocol.
	SSHSerialConnectTypesSupported SerialConnectTypesSupported = "SSH"
	// TelnetSerialConnectTypesSupported The controller supports a Serial
	// Console connection using the Telnet protocol.
	TelnetSerialConnectTypesSupported SerialConnectTypesSupported = "Telnet"
	// IPMISerialConnectTypesSupported The controller supports a Serial
	// Console connection using the IPMI Serial-over-LAN (SOL) protocol.
	IPMISerialConnectTypesSupported SerialConnectTypesSupported = "IPMI"
	// OemSerialConnectTypesSupported The controller supports a Serial
	// Console connection using an OEM-specific protocol.
	OemSerialConnectTypesSupported SerialConnectTypesSupported = "Oem"
)

// CommandShell shall describe a Command Shell service of a manager.
type CommandShell struct {
	// ConnectTypesSupported shall be an array of the enumerations provided here.
	// SSH shall be included if the Secure Shell (SSH) protocol is supported.
	// Telnet shall be included if the Telnet protocol is supported. IPMI shall
	// be included if the IPMI (Serial-over-LAN) protocol is supported.
	ConnectTypesSupported []CommandConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent
	// service sessions supported by the implementation.
	MaxConcurrentSessions uint32
	// ServiceEnabled is used for the service. The value shall be true if
	// enabled and false if disabled.
	ServiceEnabled bool
}

// GraphicalConsole shall describe a Graphical Console service of a manager.
type GraphicalConsole struct {
	// ConnectTypesSupported shall be an array of the enumerations provided here.
	// RDP shall be included if the Remote Desktop (RDP) protocol is supported.
	// KVMIP shall be included if a vendor-define KVM-IP protocol is supported.
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent
	// service sessions supported by the implementation.
	MaxConcurrentSessions uint32
	// ServiceEnabled is used for the service. The value shall be true if
	// enabled and false if disabled.
	ServiceEnabled bool
}

// Manager is a management subsystem. Examples of managers are BMCs, Enclosure
// Managers, Management Controllers and other subsystems assigned manageability
// functions.
type Manager struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AutoDSTEnabled shall contain the enabled status of the automatic Daylight
	// Saving Time (DST) adjustment of the manager's DateTime. It shall be true
	// if Automatic DST adjustment is enabled and false if disabled.
	AutoDSTEnabled bool
	// CommandShell shall contain information
	// about the Command Shell service of this manager.
	CommandShell CommandShell
	// DateTime shall represent the current DateTime value for the manager, with
	// offset from UTC, in Redfish Timestamp format.
	DateTime string
	// DateTimeLocalOffset is The value is property shall represent the offset
	// from UTC time that the current value of DataTime property contains.
	DateTimeLocalOffset string
	// DaylightSavingTime shall contain the daylight saving time settings for this manager.
	DaylightSavingTime DaylightSavingTime
	// DedicatedNetworkPorts shall contain a link to a resource collection of type PortCollection that represent the
	// dedicated network ports of the manager.
	dedicatedNetworkPorts string
	// Description provides a description of this resource.
	Description string
	// ethernetInterfaces shall be a link to a collection of type
	// EthernetInterfaceCollection.
	ethernetInterfaces string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated manager.
	FirmwareVersion string
	// GraphicalConsole shall contain the information about the Graphical
	// Console (KVM-IP) service of this manager.
	GraphicalConsole GraphicalConsole
	// hostInterfaces shall be a link to a collection of type
	// HostInterfaceCollection.
	hostInterfaces string
	// LastResetTime last BMC reset time
	LastResetTime string `json:"LastResetTime,omitempty"`
	// Location shall contain the location information of the associated manager.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// logServices shall contain a reference to a collection of type
	// LogServiceCollection which are for the use of this manager.
	logServices string
	// ManagerDiagnosticData shall contain a link to a resource of type ManagerDiagnosticData that represents the
	// diagnostic data for this manager.
	managerDiagnosticData string
	// ManagerType shall describe the function of this manager. The 'ManagementController' value shall be used if none
	// of the other enumerations apply.
	// ManagerType is used if this manager controls one or more services
	// through aggregation. The value BMC shall be used if this manager
	// represents a traditional server management controller. The value
	// ManagementController shall be used if none of the other enumerations
	// apply.
	ManagerType ManagerType
	// Manufacturer shall contain the name of the organization responsible for
	// producing the manager. This organization might be the entity from whom
	// the manager is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the information about how the manufacturer references
	// this manager.
	Model string
	// networkProtocol shall contain a reference to a resource of type
	// ManagerNetworkProtocol which represents the network services for this
	// manager.
	networkProtocol string
	// OemActions contains all the vendor specific actions. It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
	// Oem are all OEM data under top level manager section
	Oem json.RawMessage
	// OEMLinks are all OEM data under link section
	OemLinks json.RawMessage
	// PartNumber shall contain a part number assigned by the organization that
	// is responsible for producing or manufacturing the manager.
	PartNumber string
	// PowerState shall contain the power state of the Manager.
	PowerState PowerState
	// Redundancy is used to show how this manager is grouped with other
	// managers for form redundancy sets.
	Redundancy []Redundancy
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// remoteAccountService shall contain a reference to the
	// AccountService resource for the remote Manager represented by this
	// resource. This property shall only be present when providing
	// aggregation of Redfish services.
	remoteAccountService string
	// RemoteRedfishServiceURI shall contain the URI of the
	// Redfish Service Root for the remote Manager represented by this
	// resource. This property shall only be present when providing
	// aggregation of Redfish services.
	RemoteRedfishServiceURI string `json:"RemoteRedfishServiceUri"`
	// SerialConsole shall contain information about the Serial Console service
	// of this manager.
	SerialConsole SerialConsole
	// serialInterfaces shall be a link to a collection of type
	// SerialInterfaceCollection which are for the use of this manager.
	serialInterfaces string
	// SerialNumber shall contain a manufacturer-allocated number that
	// identifies the manager.
	SerialNumber string
	// ServiceEntryPointUUID shall contain the UUID of the Redfish Service
	// provided by this manager. Each Manager providing an Entry Point to the
	// same Redfish Service shall report the same UUID value (even though the
	// name of the property may imply otherwise). This property shall not be
	// present if this manager does not provide a Redfish Service Entry Point.
	ServiceEntryPointUUID string
	// ServiceIdentification shall contain a vendor-provided or user-provided value that identifies and associates a
	// discovered Redfish service with a particular product instance. If this manager provides the Redfish service, the
	// ServiceIdentification property in the ServiceRoot resource shall contain the value of this property. This
	// property shall only be present if the manager provides the Redfish service. The value of this property is used
	// in conjunction with the Product and Vendor properties in ServiceRoot to match user credentials or other a priori
	// product instance information necessary for initial deployment to the correct, matching Redfish service.
	ServiceIdentification string
	// SharedNetworkPorts shall contain a link to a resource collection of type PortCollection that represent the
	// shared network ports of the manager. The members of this collection shall reference Port resources subordinate
	// to NetworkAdapter resources.
	sharedNetworkPorts string
	// SparePartNumber shall contain the spare part number of the manager.
	SparePartNumber string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// TimeZoneName shall contain the time zone of the manager. The time zone shall be either the 'Name' or the
	// 'Format' for the zone as defined in the IANA Time Zone Database. The value of this property is used for display
	// purposes, especially to enhance the display of time. A Redfish service may not be able to ensure accuracy and
	// consistency between the DateTimeOffset property and this property. Therefore, to specify the correct time zone
	// offset, see the DateTimeOffset property.
	TimeZoneName string
	// USBPorts shall contain a link to a resource collection of type PortCollection that represent the USB ports of
	// the manager.
	usbPorts string
	// UUID shall contain the universal unique
	// identifier number for the manager.
	UUID string
	// Version shall contain the hardware version of this manager as determined by the vendor or supplier.
	Version string
	// virtualMedia shall contain a reference to a collection of type
	// VirtualMediaCollection which are for the use of this manager.
	// This property has been deprecated in favor of the VirtualMedia property in the ComputerSystem resource.
	virtualMedia string

	// ActiveSoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the active
	// firmware image for this manager.
	activeSoftwareImage string
	// ManagedBy shall contain an array of links to resources of type Manager that represent the managers for this
	// manager.
	managedBy []string
	// ManagedByCount is the number of managers for this manager.
	ManagedByCount int
	// managerForChassis shall contain an array of references to Chassis
	// resources of which this Manager instance has control.
	managerForChassis []string
	// ManagerForChassisCount is the number of Chassis being managed.
	ManagerForChassisCount int
	// ManagerForManagers shall contain an array of links to resources of type Manager that represent the managers
	// being managed by this manager.
	managerForManagers []string
	// ManagerForManagersCount is the number of managers being managed by this manager.
	ManagerForManagersCount int
	// managerForServers shall contain an array of references to ComputerSystem
	// resources of which this Manager instance has control.
	managerForServers []string
	// ManagerForServersCount is the number of Servers being managed.
	ManagerForServersCount int
	// managerForSwitches shall contain an array of references to Switch
	// resources of which this Manager instance has control.
	managerForSwitches []string
	// ManagerForSwitchesCount is the number of Switches being managed.
	ManagerForSwitchesCount int
	// managerInChassis shall contain a reference to the chassis that this
	// manager is located in.
	managerInChassis string
	// SelectedNetworkPort shall contain a link to a resource of type Port that represents the current network port
	// used by this manager.
	selectedNetworkPort string
	// SoftwareImages shall contain an array of links to resources of type SoftwareInventory that represent the
	// firmware images that apply to this manager.
	softwareImages []string
	// SoftwareImagesCount is the number of firmware images that apply to this manager.
	SoftwareImagesCount int

	forceFailoverTarget       string
	modifyRedundancySetTarget string
	// resetTarget is the internal URL to send reset targets to.
	resetTarget string
	// resetInfo contains URI for an ActionInfo Resource that describes this action.
	actionInfo string
	// SupportedResetTypes, if provided, is the reset types this system supports.
	SupportedResetTypes   []ResetType
	resetToDefaultsTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (manager *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	type actions struct {
		ForceFailover       common.ActionTarget `json:"#Manager.ForceFailover"`
		ModifyRedundancySet common.ActionTarget `json:"#Manager.ModifyRedundancySet"`
		Reset               struct {
			ActionInfo        string      `json:"@Redfish.ActionInfo"`
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			Target            string
		} `json:"#Manager.Reset"`
		ResetToDefaults common.ActionTarget `json:"#Manager.ResetToDefaults"`

		Oem json.RawMessage
	}
	type linkReference struct {
		ActiveSoftwareImage     common.Link
		ManagedBy               common.Links
		ManagedByCount          int `json:"ManagedBy@odata.count"`
		ManagerForChassis       common.Links
		ManagerForChassisCount  int `json:"ManagerForChassis@odata.count"`
		ManagerForManagers      common.Links
		ManagerForManagersCount int `json:"ManagerForManagers@odata.count"`
		ManagerForServers       common.Links
		ManagerForServersCount  int `json:"ManagerForServers@odata.count"`
		ManagerForSwitches      common.Links
		ManagerForSwitchesCount int `json:"ManagerForSwitches@odata.count"`
		ManagerInChassis        common.Link
		OEM                     json.RawMessage `json:"Oem"`
		SelectedNetworkPort     common.Link
		SoftwareImages          common.Links
		SoftwareImagesCount     int `json:"SoftwareImages@odata.count"`
	}
	var t struct {
		temp
		DedicatedNetworkPorts common.Link
		EthernetInterfaces    common.Link
		HostInterfaces        common.Link
		LogServices           common.Link
		NetworkProtocol       common.Link
		ManagerDiagnosticData common.Link
		RemoteAccountService  common.Link
		SharedNetworkPorts    common.Link
		SerialInterfaces      common.Link
		USBPorts              common.Link
		VirtualMedia          common.Link
		Links                 linkReference
		Actions               actions
		Oem                   json.RawMessage
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities
	*manager = Manager(t.temp)
	manager.dedicatedNetworkPorts = t.DedicatedNetworkPorts.String()
	manager.ethernetInterfaces = t.EthernetInterfaces.String()
	manager.hostInterfaces = t.HostInterfaces.String()
	manager.logServices = t.LogServices.String()
	manager.managerDiagnosticData = t.ManagerDiagnosticData.String()
	manager.networkProtocol = t.NetworkProtocol.String()
	manager.OemActions = t.Actions.Oem
	manager.Oem = t.Oem
	manager.remoteAccountService = t.RemoteAccountService.String()
	manager.sharedNetworkPorts = t.SharedNetworkPorts.String()
	manager.serialInterfaces = t.SerialInterfaces.String()
	manager.usbPorts = t.USBPorts.String()
	manager.virtualMedia = t.VirtualMedia.String()

	manager.activeSoftwareImage = t.Links.ActiveSoftwareImage.String()
	manager.managedBy = t.Links.ManagedBy.ToStrings()
	manager.ManagedByCount = t.Links.ManagedByCount
	manager.managerForChassis = t.Links.ManagerForChassis.ToStrings()
	manager.ManagerForChassisCount = t.Links.ManagerForChassisCount
	manager.managerForManagers = t.Links.ManagerForManagers.ToStrings()
	manager.ManagerForManagersCount = t.Links.ManagerForManagersCount
	manager.managerForServers = t.Links.ManagerForServers.ToStrings()
	manager.ManagerForServersCount = t.Links.ManagerForServersCount
	manager.ManagerForSwitchesCount = t.Links.ManagerForSwitchesCount
	manager.managerForSwitches = t.Links.ManagerForSwitches.ToStrings()
	manager.managerInChassis = t.Links.ManagerInChassis.String()
	manager.OemLinks = t.Links.OEM
	manager.selectedNetworkPort = t.Links.SelectedNetworkPort.String()
	manager.softwareImages = t.Links.SoftwareImages.ToStrings()
	manager.SoftwareImagesCount = t.Links.SoftwareImagesCount

	manager.forceFailoverTarget = t.Actions.ForceFailover.Target
	manager.modifyRedundancySetTarget = t.Actions.ModifyRedundancySet.Target
	manager.SupportedResetTypes = t.Actions.Reset.AllowedResetTypes
	manager.resetTarget = t.Actions.Reset.Target
	manager.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	manager.actionInfo = t.Actions.Reset.ActionInfo

	// This is a read/write object, so we need to save the raw object data for later
	manager.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manager *Manager) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Manager)
	err := original.UnmarshalJSON(manager.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"LocationIndicatorActive",
		"ServiceIdentification",
		"TimeZoneName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manager).Elem()

	return manager.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManager will get a Manager instance from the Swordfish service.
func GetManager(c common.Client, uri string) (*Manager, error) {
	return common.GetObject[Manager](c, uri)
}

// ListReferencedManagers gets the collection of Managers
func ListReferencedManagers(c common.Client, link string) ([]*Manager, error) {
	return common.GetCollectionObjects[Manager](c, link)
}

// ForceFailover forces a failover to the specified manager.
// **Need to test.** Spec calls for the Manager as a parameter, but it may actually
// be the Manager.ODataID.
func (manager *Manager) ForceFailover(newManager *Manager) error {
	return manager.Post(manager.forceFailoverTarget, newManager)
}

// ModifyRedundancySet adds members to or removes members from a redundant group of managers.
// **Need to test.** Spec calls for the Manager as a parameter, but it may actually
// be the Manager.ODataID.
func (manager *Manager) ModifyRedundancySet(addManagers, removeManagers []*Manager) error {
	parameters := struct {
		Add    []*Manager
		Remove []*Manager
	}{
		Add:    addManagers,
		Remove: removeManagers,
	}
	return manager.Post(manager.modifyRedundancySetTarget, parameters)
}

// Reset shall perform a reset of the manager.
func (manager *Manager) Reset(resetType ResetType) error {
	if len(manager.SupportedResetTypes) == 0 {
		if manager.actionInfo != "" {
			// reset without confirming the type is supported by the manager.
			// done to minimize overhead though technically not as correct as first checking the supported reset types
			t := struct {
				ResetType ResetType
			}{ResetType: resetType}
			return manager.Post(manager.resetTarget, t)
		}
		// reset directly without reset type. HPE server has the behavior
		return manager.Post(manager.resetTarget, struct{}{})
	}
	// Make sure the requested reset type is supported by the manager.
	valid := false
	for _, allowed := range manager.SupportedResetTypes {
		if resetType == allowed {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("reset type '%s' is not supported by this manager",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	return manager.Post(manager.resetTarget, t)
}

// ResetToDefaults resets the manager settings to factory defaults. This can cause the manager to reset.
func (manager *Manager) ResetToDefaults(resetType ResetToDefaultsType) error {
	t := struct {
		ResetType ResetToDefaultsType
	}{ResetType: resetType}
	return manager.Post(manager.resetToDefaultsTarget, t)
}

// DedicatedNetworkPorts gets the dedicated network ports of the manager.
func (manager *Manager) DedicatedNetworkPorts() ([]*Port, error) {
	return ListReferencedPorts(manager.GetClient(), manager.dedicatedNetworkPorts)
}

// EthernetInterfaces get this manager's ethernet interfaces.
func (manager *Manager) EthernetInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(manager.GetClient(), manager.ethernetInterfaces)
}

// HostInterfaces get this manager's host interfaces.
func (manager *Manager) HostInterfaces() ([]*HostInterface, error) {
	return ListReferencedHostInterfaces(manager.GetClient(), manager.hostInterfaces)
}

// LogServices get this manager's log services on this system.
func (manager *Manager) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(manager.GetClient(), manager.logServices)
}

// ManagerDiagnosticData gets the diagnostic data for this manager.
func (manager *Manager) ManagerDiagnosticData() (*ManagerDiagnosticData, error) {
	return GetManagerDiagnosticData(manager.GetClient(), manager.managerDiagnosticData)
}

// NetworkProtocol get this manager's network protocol settings.
func (manager *Manager) NetworkProtocol() (*NetworkProtocolSettings, error) {
	return GetNetworkProtocol(manager.GetClient(), manager.networkProtocol)
}

// RemoteAccountService gets the account service resource for the remote manager that this resource represents.
// This property shall only be present when providing aggregation of a remote manager.
func (manager *Manager) RemoteAccountService() (*AccountService, error) {
	return GetAccountService(manager.GetClient(), manager.remoteAccountService)
}

// SharedNetworkPorts gets the shared network ports of the manager.
func (manager *Manager) SharedNetworkPorts() ([]*Port, error) {
	return ListReferencedPorts(manager.GetClient(), manager.sharedNetworkPorts)
}

// SerialInterfaces get this manager's serial interfaces.
func (manager *Manager) SerialInterfaces() ([]*SerialInterface, error) {
	return ListReferencedSerialInterfaces(manager.GetClient(), manager.serialInterfaces)
}

// USBPorts get the USB ports of the manager.
func (manager *Manager) USBPorts() ([]*Port, error) {
	return ListReferencedPorts(manager.GetClient(), manager.usbPorts)
}

// VirtualMedia gets the virtual media associated with this manager.
// This property has been deprecated in favor of the VirtualMedia property in the ComputerSystem resource.
func (manager *Manager) VirtualMedia() ([]*VirtualMedia, error) {
	return ListReferencedVirtualMedias(manager.GetClient(), manager.virtualMedia)
}

// ActiveSoftwareImage gets the software inventory resource that represents the active firmware image for this manager.
func (manager *Manager) ActiveSoftwareImage() (*SoftwareInventory, error) {
	if manager.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetSoftwareInventory(manager.GetClient(), manager.activeSoftwareImage)
}

// ManagedBy gets the managers responsible for managing this manager.
func (manager *Manager) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](manager.GetClient(), manager.managedBy)
}

// ManagedForChassis gets the the chassis this manager controls.
func (manager *Manager) ManagedForChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](manager.GetClient(), manager.managerForChassis)
}

// ManagerForManagers gets the managers that are managed by this manager.
func (manager *Manager) ManagerForManagers() ([]*Manager, error) {
	return common.GetObjects[Manager](manager.GetClient(), manager.managerForManagers)
}

// ManagerForServers gets the systems that this manager controls.
func (manager *Manager) ManagerForServers() ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](manager.GetClient(), manager.managerForServers)
}

// ManagerForSwitches gets the switches that this manager controls.
func (manager *Manager) ManagerForSwitches() ([]*Switch, error) {
	return common.GetObjects[Switch](manager.GetClient(), manager.managerForSwitches)
}

// SelectedNetworkPort gets the current network port used by this manager.
func (manager *Manager) SelectedNetworkPort() (*Port, error) {
	if manager.selectedNetworkPort == "" {
		return nil, nil
	}
	return GetPort(manager.GetClient(), manager.selectedNetworkPort)
}

// SoftwareImages gets the firmware images that apply to this manager.
func (manager *Manager) SoftwareImages() ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](manager.GetClient(), manager.softwareImages)
}
