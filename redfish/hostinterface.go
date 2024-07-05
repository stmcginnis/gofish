//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AuthenticationMode is the method used for authentication.
type AuthenticationMode string

const (

	// AuthNoneAuthenticationMode Requests without any sort of authentication
	// are allowed.
	AuthNoneAuthenticationMode AuthenticationMode = "AuthNone"
	// BasicAuthAuthenticationMode Requests using HTTP Basic Authentication
	// are allowed.
	BasicAuthAuthenticationMode AuthenticationMode = "BasicAuth"
	// RedfishSessionAuthAuthenticationMode Requests using Redfish Session
	// Authentication are allowed.
	RedfishSessionAuthAuthenticationMode AuthenticationMode = "RedfishSessionAuth"
	// OemAuthAuthenticationMode Requests using OEM authentication mechanisms
	// are allowed.
	OemAuthAuthenticationMode AuthenticationMode = "OemAuth"
)

// HostInterfaceType is the type of network interface.
type HostInterfaceType string

const (

	// NetworkHostInterfaceHostInterfaceType This interface is a Network Host
	// Interface.
	NetworkHostInterfaceHostInterfaceType HostInterfaceType = "NetworkHostInterface"
)

type CredentialBootstrapping struct {
	// EnableAfterReset shall indicate whether credential bootstrapping is enabled after a reset for this interface. If
	// 'true', services shall set the Enabled property to 'true' after a reset of the host or the service.
	EnableAfterReset bool
	// Enabled shall indicate whether credential bootstrapping is enabled for this interface.
	Enabled bool
	// RoleID shall contain the ID property of the role resource that is used for the bootstrap account created for
	// this interface.
	RoleID string
}

// HostInterface is used to represent Host Interface resources as part of
// the Redfish specification.
type HostInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AuthNoneRoleID is used when no authentication on this interface is
	// performed. This property shall be absent if AuthNone is not supported
	// by the service for the AuthenticationModes property.
	AuthNoneRoleID string `json:"AuthNoneRoleId"`
	// AuthenticationModes shall be an array consisting of the authentication
	// modes allowed on this interface.
	AuthenticationModes []AuthenticationMode
	// CredentialBootstrapping shall contain settings for the Redfish Host Interface Specification-defined 'credential
	// bootstrapping via IPMI commands' feature for this interface. This property shall be absent if credential
	// bootstrapping is not supported by the service.
	CredentialBootstrapping CredentialBootstrapping
	// Description provides a description of this resource.
	Description string
	// ExternallyAccessible is used by external clients, and this property
	// will have the value set to true.
	ExternallyAccessible bool
	// FirmwareAuthEnabled shall be a boolean
	// indicating whether firmware authentication for this interface is
	// enabled.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	FirmwareAuthEnabled bool
	// FirmwareAuthRoleID shall be the ID of the Role resource that is
	// configured for firmware authentication on this interface.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	FirmwareAuthRoleID string `json:"FirmwareAuthRoleId"`
	// HostEthernetInterfaces shall be a link to a collection of type
	// EthernetInterfaceCollection that Computer Systems use as the Host
	// Interface to this Manager.
	hostEthernetInterfaces string
	// HostInterfaceType shall be an enumeration describing type of the interface.
	HostInterfaceType HostInterfaceType
	// InterfaceEnabled shall be a boolean indicating whether this interface is
	// enabled.
	InterfaceEnabled bool
	// KernelAuthEnabled shall be a boolean indicating whether kernel
	// authentication for this interface is enabled.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	KernelAuthEnabled bool
	// KernelAuthRoleID shall be the ID of the Role resource that is configured
	// for kernel authentication on this interface.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	KernelAuthRoleID string `json:"KernelAuthRoleId"`
	// ManagerEthernetInterface is used by this Manager as the HostInterface.
	managerEthernetInterface string
	// NetworkProtocol shall contain a reference to a resource of type
	// ManagerNetworkProtocol which represents the network services for this
	// Manager.
	networkProtocol string
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status

	// AuthNoneRole shall be a link to a Role object instance, and should
	// reference the object identified by property AuthNoneRoleId. This property
	// shall be absent if AuthNone is not supported by the service for the
	// AuthenticationModes property.
	authNoneRole string
	// ComputerSystems shall be an array of references to resources of type
	// ComputerSystem that are connected to this HostInterface.
	computerSystems []string
	// ComputerSystemsCount is the number of computer systems.
	ComputerSystemsCount int
	// CredentialBootstrappingRole shall contain a link to a resource of type Role, and should
	// link to the resource identified by the RoleId property within CredentialBootstrapping.
	// This property shall be absent if the Redfish Host Interface Specification-defined
	// 'credential bootstrapping via IPMI commands' feature is not supported by the service.
	credentialBootstrappingRole string
	// FirmwareAuthRole shall be a link to a Role object instance, and should
	// reference the object identified by property FirmwareAuthRoleID.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	firmwareAuthRole string
	// KernelAuthRole shall be a link to a Role object instance, and should
	// reference the object identified by property KernelAuthRoleId.
	// This property has been deprecated in favor of newer methods of negotiating credentials.
	kernelAuthRole string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a HostInterface object from the raw JSON.
func (hostinterface *HostInterface) UnmarshalJSON(b []byte) error {
	type temp HostInterface

	type links struct {
		AuthNoneRole                common.Link
		ComputerSystems             common.Links
		ComputerSystemsCount        int `json:"ComputerSystems@odata.count"`
		CredentialBootstrappingRole common.Link
		FirmwareAuthRole            common.Link
		KernelAuthRole              common.Link
	}

	var t struct {
		temp
		Links                    links
		HostEthernetInterfaces   common.Link
		ManagerEthernetInterface common.Link
		NetworkProtocol          common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostinterface = HostInterface(t.temp)

	// Extract the links to other entities for later
	hostinterface.authNoneRole = t.Links.AuthNoneRole.String()
	hostinterface.computerSystems = t.Links.ComputerSystems.ToStrings()
	hostinterface.ComputerSystemsCount = t.Links.ComputerSystemsCount
	hostinterface.credentialBootstrappingRole = t.Links.CredentialBootstrappingRole.String()
	hostinterface.firmwareAuthRole = t.Links.FirmwareAuthRole.String()
	hostinterface.kernelAuthRole = t.Links.KernelAuthRole.String()

	hostinterface.hostEthernetInterfaces = t.HostEthernetInterfaces.String()
	hostinterface.managerEthernetInterface = t.ManagerEthernetInterface.String()
	hostinterface.networkProtocol = t.NetworkProtocol.String()

	// This is a read/write object, so we need to save the raw object data for later
	hostinterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (hostinterface *HostInterface) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(HostInterface)
	err := original.UnmarshalJSON(hostinterface.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AuthNoneRoleId",
		"AuthenticationModes",
		"InterfaceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(hostinterface).Elem()

	return hostinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetHostInterface will get a HostInterface instance from the service.
func GetHostInterface(c common.Client, uri string) (*HostInterface, error) {
	return common.GetObject[HostInterface](c, uri)
}

// ListReferencedHostInterfaces gets the collection of HostInterface from
// a provided reference.
func ListReferencedHostInterfaces(c common.Client, link string) ([]*HostInterface, error) {
	return common.GetCollectionObjects[HostInterface](c, link)
}

// ComputerSystems references the ComputerSystems that this host interface is associated with.
func (hostinterface *HostInterface) ComputerSystems() ([]*ComputerSystem, error) {
	var result []*ComputerSystem

	collectionError := common.NewCollectionError()
	for _, computerSystemLink := range hostinterface.computerSystems {
		computerSystem, err := GetComputerSystem(hostinterface.GetClient(), computerSystemLink)
		if err != nil {
			collectionError.Failures[computerSystemLink] = err
		} else {
			result = append(result, computerSystem)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// HostEthernetInterfaces gets the network interface controllers or cards (NICs)
// that a Computer System uses to communicate with this Host Interface.
func (hostinterface *HostInterface) HostEthernetInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(hostinterface.GetClient(), hostinterface.hostEthernetInterfaces)
}

// ManagerNetworkInterfaces gets the network interface controllers or cards
// (NIC) that this Manager uses for network communication with this Host Interface.
func (hostinterface *HostInterface) ManagerNetworkInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(hostinterface.GetClient(), hostinterface.managerEthernetInterface)
}

// AuthRoleNone gets the role that contains the privileges on this host interface when no authentication is performed.
func (hostinterface *HostInterface) AuthNoneRole() (*Role, error) {
	if hostinterface.authNoneRole == "" {
		return nil, nil
	}
	return GetRole(hostinterface.GetClient(), hostinterface.authNoneRole)
}

// CredentialBootstrappingRole gets the role that contains the privileges for the bootstrap account created for this interface.
func (hostinterface *HostInterface) CredentialBootstrappingRole() (*Role, error) {
	if hostinterface.credentialBootstrappingRole == "" {
		return nil, nil
	}
	return GetRole(hostinterface.GetClient(), hostinterface.credentialBootstrappingRole)
}
