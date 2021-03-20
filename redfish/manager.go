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
	// hostInterfaces string
	// logServices shall contain a reference to a collection of type
	// LogServiceCollection which are for the use of this manager.
	logServices string
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
	// OEMData are all OEM data under top level manager section
	OEMData map[string]interface{}
	// OEMLinks are all OEM data under link section
	OEMLinks map[string]interface{}
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
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// UUID shall contain the universal unique
	// identifier number for the manager.
	UUID string
	// virtualMedia shall contain a reference to a collection of type
	// VirtualMediaCollection which are for the use of this manager.
	virtualMedia string
	// managerForChassis shall contain an array of references to Chassis
	// resources of which this Manager instance has control.
	managerForChassis []string
	// ManagerForChassisCount is the number of Chassis being managed.
	ManagerForChassisCount int
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
	// resetTarget is the internal URL to send reset targets to.
	resetTarget string
	// SupportedResetTypes, if provided, is the reset types this system supports.
	SupportedResetTypes []ResetType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (manager *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	type actions struct {
		Reset struct {
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			Target            string
		} `json:"#Manager.Reset"`
	}
	type linkReference struct {
		ManagerForChassis       common.Links
		ManagerForChassisCount  int `json:"ManagerForChassis@odata.count"`
		ManagerForServers       common.Links
		ManagerForServersCount  int `json:"ManagerForServers@odata.count"`
		ManagerForSwitches      common.Links
		ManagerForSwitchesCount int `json:"ManagerForSwitches@odata.count"`
		ManagerInChassis        common.Link
		OEM                     map[string]interface{} `json:"Oem"`
	}
	var t struct {
		temp
		EthernetInterfaces   common.Link
		LogServices          common.Link
		NetworkProtocol      common.Link
		RemoteAccountService common.Link
		SerialInterfaces     common.Link
		VirtualMedia         common.Link
		Links                linkReference
		Actions              actions
		OEM                  map[string]interface{} `json:"Oem"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities
	*manager = Manager(t.temp)
	manager.ethernetInterfaces = string(t.EthernetInterfaces)
	manager.logServices = string(t.LogServices)
	manager.networkProtocol = string(t.NetworkProtocol)
	manager.OEMData = t.OEM
	manager.OEMLinks = t.Links.OEM
	manager.remoteAccountService = string(t.RemoteAccountService)
	manager.serialInterfaces = string(t.SerialInterfaces)
	manager.virtualMedia = string(t.VirtualMedia)
	manager.managerForServers = t.Links.ManagerForServers.ToStrings()
	manager.ManagerForServersCount = t.Links.ManagerForServersCount
	manager.managerForChassis = t.Links.ManagerForChassis.ToStrings()
	manager.ManagerForChassisCount = t.Links.ManagerForChassisCount
	manager.ManagerForSwitchesCount = t.Links.ManagerForSwitchesCount
	manager.managerForSwitches = t.Links.ManagerForSwitches.ToStrings()
	manager.managerInChassis = string(t.Links.ManagerInChassis)
	manager.SupportedResetTypes = t.Actions.Reset.AllowedResetTypes
	manager.resetTarget = t.Actions.Reset.Target

	// This is a read/write object, so we need to save the raw object data for later
	manager.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manager *Manager) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Manager)
	err := original.UnmarshalJSON(manager.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manager).Elem()

	return manager.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManager will get a Manager instance from the Swordfish service.
func GetManager(c common.Client, uri string) (*Manager, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manager Manager
	err = json.NewDecoder(resp.Body).Decode(&manager)
	if err != nil {
		return nil, err
	}

	manager.SetClient(c)
	return &manager, nil
}

// ListReferencedManagers gets the collection of Managers
func ListReferencedManagers(c common.Client, link string) ([]*Manager, error) {
	var result []*Manager
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, managerLink := range links.ItemLinks {
		manager, err := GetManager(c, managerLink)
		if err != nil {
			return result, err
		}
		result = append(result, manager)
	}

	return result, nil
}

// Reset shall perform a reset of the manager.
func (manager *Manager) Reset(resetType ResetType) error {
	if len(manager.SupportedResetTypes) == 0 {
		// reset directly without reset type. HPE server has the behavior
		type temp struct {
			Action string
		}
		t := temp{
			Action: "Manager.Reset",
		}

		_, err := manager.Client.Post(manager.resetTarget, t)
		return err
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

	type temp struct {
		ResetType ResetType
	}
	t := temp{
		ResetType: resetType,
	}

	_, err := manager.Client.Post(manager.resetTarget, t)
	return err
}

// EthernetInterfaces get this system's ethernet interfaces.
func (manager *Manager) EthernetInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(manager.Client, manager.ethernetInterfaces)
}

// LogServices get this manager's log services on this system.
func (manager *Manager) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(manager.Client, manager.logServices)
}

// VirtualMedia gets the virtual media associated with this manager.
func (manager *Manager) VirtualMedia() ([]*VirtualMedia, error) {
	return ListReferencedVirtualMedias(manager.Client, manager.virtualMedia)
}
