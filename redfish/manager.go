//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/coreweave/gofish/common"
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

type ManagerActions struct {
	ForceFailover       common.ActionTarget `json:"#Manager.ForceFailover"`
	ModifyRedundancySet common.ActionTarget `json:"#Manager.ModifyRedundancySet"`
	Reset               struct {
		common.ActionTarget
		AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
	} `json:"#Manager.Reset"`
	ResetToDefaults struct {
		common.ActionTarget
		AllowedResetTypes []ResetToDefaultsType `json:"ResetType@Redfish.AllowableValues"`
	} `json:"#Manager.ResetToDefaults"`

	Oem json.RawMessage
}

type ManagerLinks struct {
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

// Manager is a management subsystem. Examples of managers are BMCs, Enclosure
// Managers, Management Controllers and other subsystems assigned manageability
// functions.
type Manager struct {
	common.Entity

	Actions ManagerActions `json:"Actions"`
	Links   ManagerLinks   `json:"Links"`

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
	// DateTimeSource shall contain the source of the current DateTime value for the manager.
	DateTimeSource string
	// DaylightSavingTime shall contain the daylight saving time settings for this manager.
	DaylightSavingTime DaylightSavingTime
	// DedicatedNetworkPortsLink shall contain a link to a resource collection of type PortCollection that represent the
	// dedicated network ports of the manager.
	DedicatedNetworkPortsLink common.Link `json:"DedicatedNetworkPorts"`
	// Description provides a description of this resource.
	Description string
	// EthernetInterfacesLink shall be a link to a collection of type
	// EthernetInterfaceCollection.
	EthernetInterfacesLink common.Link `json:"EthernetInterfaces"`
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated manager.
	FirmwareVersion string
	// GraphicalConsole shall contain the information about the Graphical
	// Console (KVM-IP) service of this manager.
	GraphicalConsole GraphicalConsole
	// HostInterfacesLink shall be a link to a collection of type
	// HostInterfaceCollection.
	HostInterfacesLink common.Link `json:"HostInterfaces"`
	// LastResetTime last BMC reset time
	LastResetTime string `json:"LastResetTime,omitempty"`
	// Location shall contain the location information of the associated manager.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// LogServicesLink shall contain a reference to a collection of type
	// LogServiceCollection which are for the use of this manager.
	LogServicesLink common.Link `json:"LogServices"`
	// ManagerDiagnosticDataLink shall contain a link to a resource of type ManagerDiagnosticData that represents the
	// diagnostic data for this manager.
	ManagerDiagnosticDataLink common.Link `json:"ManagerDiagnosticData"`
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
	// NetworkProtocolLink shall contain a reference to a resource of type
	// ManagerNetworkProtocol which represents the network services for this
	// manager.
	NetworkProtocolLink common.Link `json:"NetworkProtocol"`
	// Oem are all OEM data under top level manager section
	Oem json.RawMessage
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
	// RemoteAccountServiceLink shall contain a reference to the
	// AccountService resource for the remote Manager represented by this
	// resource. This property shall only be present when providing
	// aggregation of Redfish services.
	RemoteAccountServiceLink common.Link `json:"RemoteAccountService"`
	// RemoteRedfishServiceURI shall contain the URI of the
	// Redfish Service Root for the remote Manager represented by this
	// resource. This property shall only be present when providing
	// aggregation of Redfish services.
	RemoteRedfishServiceURI string `json:"RemoteRedfishServiceUri"`
	// SerialConsole shall contain information about the Serial Console service
	// of this manager.
	SerialConsole SerialConsole
	// SerialInterfacesLink shall be a link to a collection of type
	// SerialInterfaceCollection which are for the use of this manager.
	SerialInterfacesLink common.Link `json:"SerialInterfaces"`
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
	// SharedNetworkPortsLink shall contain a link to a resource collection of type PortCollection that represent the
	// shared network ports of the manager. The members of this collection shall reference Port resources subordinate
	// to NetworkAdapter resources.
	SharedNetworkPortsLink common.Link `json:"SharedNetworkPorts"`
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
	// USBPortsLink shall contain a link to a resource collection of type PortCollection that represent the USB ports of
	// the manager.
	USBPortsLink common.Link `json:"USBPorts"`
	// UUID shall contain the universal unique
	// identifier number for the manager.
	UUID string
	// Version shall contain the hardware version of this manager as determined by the vendor or supplier.
	Version string
	// VirtualMediaLink shall contain a reference to a collection of type
	// VirtualMediaCollection which are for the use of this manager.
	// This property has been deprecated in favor of the VirtualMedia property in the ComputerSystem resource.
	VirtualMediaLink common.Link `json:"VirtualMedia"`

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (manager *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	var t temp

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities
	*manager = Manager(t)
	// This is a read/write object, so we need to save the raw object data for later
	manager.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manager *Manager) Update() error {
	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"DateTimeSource",
		"LocationIndicatorActive",
		"ServiceIdentification",
		"TimeZoneName",
	}

	return manager.UpdateFromRawData(manager, manager.RawData, readWriteFields)
}

// GetManager will get a Manager instance from the Swordfish service.
func GetManager(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Manager, error) {
	return common.GetObject[Manager](c, uri, queryOpts...)
}

// ListReferencedManagers gets the collection of Managers
func ListReferencedManagers(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Manager, error) {
	return common.GetCollectionObjects[Manager](c, link, queryOpts...)
}

// ForceFailover forces a failover to the specified manager.
// **Need to test.** Spec calls for the Manager as a parameter, but it may actually
// be the Manager.ODataID.
func (manager *Manager) ForceFailover(newManager *Manager) error {
	if manager.Actions.ForceFailover.Target == "" {
		return ErrActionNotSupported
	}
	return manager.Post(manager.Actions.ForceFailover.Target, newManager)
}

// ModifyRedundancySet adds members to or removes members from a redundant group of managers.
// **Need to test.** Spec calls for the Manager as a parameter, but it may actually
// be the Manager.ODataID.
func (manager *Manager) ModifyRedundancySet(addManagers, removeManagers []*Manager) error {
	if manager.Actions.ModifyRedundancySet.Target == "" {
		return ErrActionNotSupported
	}
	parameters := struct {
		Add    []*Manager
		Remove []*Manager
	}{
		Add:    addManagers,
		Remove: removeManagers,
	}
	return manager.Post(manager.Actions.ModifyRedundancySet.Target, parameters)
}

// GetSupportedResetTypes returns any reset types that the Manager declares as supported
// via either ActionInfo or AllowableValues.
func (manager *Manager) GetSupportedResetTypes() ([]ResetType, error) {
	if len(manager.Actions.Reset.AllowedResetTypes) > 0 {
		return manager.Actions.Reset.AllowedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if manager.Actions.Reset.ActionInfoTarget != "" {
		resetActionInfo, err := manager.ResetActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringActionInfoDataTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			manager.Actions.Reset.AllowedResetTypes = append(manager.Actions.Reset.AllowedResetTypes, ResetType(val))
		}
	}

	return manager.Actions.Reset.AllowedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the Manager reset action if supported
func (manager *Manager) ResetActionInfo(queryOpts ...common.QueryGroupOption) (*ActionInfo, error) {
	if manager.Actions.Reset.ActionInfoTarget == "" {
		return nil, ErrActionNotSupported
	}

	return common.GetObject[ActionInfo](manager.GetClient(), manager.Actions.Reset.ActionInfoTarget, queryOpts...)
}

// Reset shall perform a reset of the manager.
func (manager *Manager) Reset(resetType ResetType) error {
	resetTarget := manager.Actions.Reset.Target
	supportedResetTypes := manager.Actions.Reset.AllowedResetTypes

	if len(supportedResetTypes) == 0 {
		if manager.Actions.Reset.ActionInfoTarget != "" {
			// reset without confirming the type is supported by the manager.
			// done to minimize overhead though technically not as correct as first checking the supported reset types
			t := struct {
				ResetType ResetType
			}{ResetType: resetType}
			return manager.Post(resetTarget, t)
		}
		// reset directly without reset type. HPE server has the behavior
		return manager.Post(resetTarget, struct{}{})
	}
	// Make sure the requested reset type is supported by the manager.
	valid := false
	for _, allowed := range supportedResetTypes {
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
	return manager.Post(resetTarget, t)
}

// GetSupportedResetToDefaultsTypes returns any reset to defaults
// types that the Manager declares as supported via either ActionInfo or AllowableValues.
func (manager *Manager) GetSupportedResetToDefaultsTypes() ([]ResetToDefaultsType, error) {
	if len(manager.Actions.ResetToDefaults.AllowedResetTypes) > 0 {
		return manager.Actions.ResetToDefaults.AllowedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if manager.Actions.ResetToDefaults.ActionInfoTarget != "" {
		resetActionInfo, err := manager.ResetToDefaultsActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringActionInfoDataTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			manager.Actions.ResetToDefaults.AllowedResetTypes = append(
				manager.Actions.ResetToDefaults.AllowedResetTypes, ResetToDefaultsType(val),
			)
		}
	}

	return manager.Actions.ResetToDefaults.AllowedResetTypes, nil
}

// ResetToDefaultsActionInfo returns the ActionInfo for the Manager ResetToDefaults action if supported
func (manager *Manager) ResetToDefaultsActionInfo(queryOpts ...common.QueryGroupOption) (*ActionInfo, error) {
	if manager.Actions.ResetToDefaults.ActionInfoTarget == "" {
		return nil, ErrActionNotSupported
	}

	return common.GetObject[ActionInfo](manager.GetClient(), manager.Actions.ResetToDefaults.ActionInfoTarget, queryOpts...)
}

// ResetToDefaults resets the manager settings to factory defaults. This can cause the manager to reset.
func (manager *Manager) ResetToDefaults(resetType ResetToDefaultsType) error {
	if manager.Actions.ResetToDefaults.Target == "" {
		return ErrActionNotSupported
	}

	t := struct {
		ResetType ResetToDefaultsType
	}{ResetType: resetType}
	return manager.Post(manager.Actions.ResetToDefaults.Target, t)
}

// DedicatedNetworkPorts gets the dedicated network ports of the manager.
func (manager *Manager) DedicatedNetworkPorts(queryOpts ...common.QueryGroupOption) ([]*Port, error) {
	if manager.DedicatedNetworkPortsLink.IsZero() {
		return nil, nil
	}
	return ListReferencedPorts(manager.GetClient(), manager.DedicatedNetworkPortsLink.String(), queryOpts...)
}

// EthernetInterfaces get this manager's ethernet interfaces.
func (manager *Manager) EthernetInterfaces(queryOpts ...common.QueryGroupOption) ([]*EthernetInterface, error) {
	if manager.EthernetInterfacesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedEthernetInterfaces(manager.GetClient(), manager.EthernetInterfacesLink.String(), queryOpts...)
}

// HostInterfaces get this manager's host interfaces.
func (manager *Manager) HostInterfaces(queryOpts ...common.QueryGroupOption) ([]*HostInterface, error) {
	if manager.HostInterfacesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedHostInterfaces(manager.GetClient(), manager.HostInterfacesLink.String(), queryOpts...)
}

// LogServices get this manager's log services on this system.
func (manager *Manager) LogServices(queryOpts ...common.QueryGroupOption) ([]*LogService, error) {
	if manager.LogServicesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedLogServices(manager.GetClient(), manager.LogServicesLink.String(), queryOpts...)
}

// ManagerDiagnosticData gets the diagnostic data for this manager.
func (manager *Manager) ManagerDiagnosticData(queryOpts ...common.QueryGroupOption) (*ManagerDiagnosticData, error) {
	if manager.ManagerDiagnosticDataLink.IsZero() {
		return nil, nil
	}
	return GetManagerDiagnosticData(manager.GetClient(), manager.ManagerDiagnosticDataLink.String(), queryOpts...)
}

// NetworkProtocol get this manager's network protocol settings.
func (manager *Manager) NetworkProtocol(queryOpts ...common.QueryGroupOption) (*NetworkProtocolSettings, error) {
	if manager.NetworkProtocolLink.IsZero() {
		return nil, nil
	}
	return GetNetworkProtocol(manager.GetClient(), manager.NetworkProtocolLink.String(), queryOpts...)
}

// RemoteAccountService gets the account service resource for the remote manager that this resource represents.
// This property shall only be present when providing aggregation of a remote manager.
func (manager *Manager) RemoteAccountService(queryOpts ...common.QueryGroupOption) (*AccountService, error) {
	if manager.RemoteAccountServiceLink.IsZero() {
		return nil, nil
	}
	return GetAccountService(manager.GetClient(), manager.RemoteAccountServiceLink.String(), queryOpts...)
}

// SharedNetworkPorts gets the shared network ports of the manager.
func (manager *Manager) SharedNetworkPorts(queryOpts ...common.QueryGroupOption) ([]*Port, error) {
	if manager.SharedNetworkPortsLink.IsZero() {
		return nil, nil
	}
	return ListReferencedPorts(manager.GetClient(), manager.SharedNetworkPortsLink.String(), queryOpts...)
}

// SerialInterfaces get this manager's serial interfaces.
func (manager *Manager) SerialInterfaces(queryOpts ...common.QueryGroupOption) ([]*SerialInterface, error) {
	if manager.SerialInterfacesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedSerialInterfaces(manager.GetClient(), manager.SerialInterfacesLink.String(), queryOpts...)
}

// USBPorts get the USB ports of the manager.
func (manager *Manager) USBPorts(queryOpts ...common.QueryGroupOption) ([]*Port, error) {
	if manager.USBPortsLink.IsZero() {
		return nil, nil
	}
	return ListReferencedPorts(manager.GetClient(), manager.USBPortsLink.String(), queryOpts...)
}

// VirtualMedia gets the virtual media associated with this manager.
// This property has been deprecated in favor of the VirtualMedia property in the ComputerSystem resource.
func (manager *Manager) VirtualMedia(queryOpts ...common.QueryGroupOption) ([]*VirtualMedia, error) {
	if manager.VirtualMediaLink.IsZero() {
		return nil, nil
	}
	return ListReferencedVirtualMedias(manager.GetClient(), manager.VirtualMediaLink.String(), queryOpts...)
}

// ActiveSoftwareImage gets the software inventory resource that represents the active firmware image for this manager.
func (manager *Manager) ActiveSoftwareImage(queryOpts ...common.QueryGroupOption) (*SoftwareInventory, error) {
	if manager.Links.ActiveSoftwareImage.IsZero() {
		return nil, nil
	}
	return GetSoftwareInventory(manager.GetClient(), manager.Links.ActiveSoftwareImage.String(), queryOpts...)
}

// ManagedBy gets the managers responsible for managing this manager.
func (manager *Manager) ManagedBy(queryOpts ...common.QueryGroupOption) ([]*Manager, error) {
	return common.GetObjects[Manager](manager.GetClient(), manager.Links.ManagedBy.ToStrings(), queryOpts...)
}

// ManagedForChassis gets the the chassis this manager controls.
func (manager *Manager) ManagedForChassis(queryOpts ...common.QueryGroupOption) ([]*Chassis, error) {
	return common.GetObjects[Chassis](manager.GetClient(), manager.Links.ManagerForChassis.ToStrings(), queryOpts...)
}

// ManagerForManagers gets the managers that are managed by this manager.
func (manager *Manager) ManagerForManagers(queryOpts ...common.QueryGroupOption) ([]*Manager, error) {
	return common.GetObjects[Manager](manager.GetClient(), manager.Links.ManagerForManagers.ToStrings(), queryOpts...)
}

// ManagerForServers gets the systems that this manager controls.
func (manager *Manager) ManagerForServers(queryOpts ...common.QueryGroupOption) ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](manager.GetClient(), manager.Links.ManagerForServers.ToStrings(), queryOpts...)
}

// ManagerForSwitches gets the switches that this manager controls.
func (manager *Manager) ManagerForSwitches(queryOpts ...common.QueryGroupOption) ([]*Switch, error) {
	return common.GetObjects[Switch](manager.GetClient(), manager.Links.ManagerForSwitches.ToStrings(), queryOpts...)
}

// SelectedNetworkPort gets the current network port used by this manager.
func (manager *Manager) SelectedNetworkPort(queryOpts ...common.QueryGroupOption) (*Port, error) {
	if manager.Links.SelectedNetworkPort.IsZero() {
		return nil, nil
	}
	return GetPort(manager.GetClient(), manager.Links.SelectedNetworkPort.String(), queryOpts...)
}

// SoftwareImages gets the firmware images that apply to this manager.
func (manager *Manager) SoftwareImages(queryOpts ...common.QueryGroupOption) ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](manager.GetClient(), manager.Links.SoftwareImages.ToStrings(), queryOpts...)
}
