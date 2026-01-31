//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.3 - #HostInterface.v1_3_3.HostInterface

package schemas

import (
	"encoding/json"
)

type AuthenticationMode string

const (
	// AuthNoneAuthenticationMode Requests without any sort of authentication are
	// allowed.
	AuthNoneAuthenticationMode AuthenticationMode = "AuthNone"
	// BasicAuthAuthenticationMode Requests using HTTP Basic authentication are
	// allowed.
	BasicAuthAuthenticationMode AuthenticationMode = "BasicAuth"
	// RedfishSessionAuthAuthenticationMode Requests using Redfish session
	// authentication are allowed.
	RedfishSessionAuthAuthenticationMode AuthenticationMode = "RedfishSessionAuth"
	// OemAuthAuthenticationMode Requests using OEM authentication mechanisms are
	// allowed.
	OemAuthAuthenticationMode AuthenticationMode = "OemAuth"
)

type HostInterfaceType string

const (
	// NetworkHostInterfaceHostInterfaceType This interface is a network host
	// interface.
	NetworkHostInterfaceHostInterfaceType HostInterfaceType = "NetworkHostInterface"
)

// HostInterface shall represent a Redfish host interface as part of the Redfish
// Specification.
type HostInterface struct {
	Entity
	// AuthNoneRoleID shall contain the 'Id' property of the 'Role' resource that
	// is used when no authentication on this interface is performed. This property
	// shall be absent if 'AuthNone' is not supported by the service for the
	// 'AuthenticationModes' property.
	//
	// Version added: v1.2.0
	AuthNoneRoleID string `json:"AuthNoneRoleId"`
	// AuthenticationModes shall contain an array consisting of the authentication
	// modes allowed on this interface.
	AuthenticationModes []AuthenticationMode
	// CredentialBootstrapping shall contain settings for the Redfish Host
	// Interface Specification-defined 'credential bootstrapping via IPMI commands'
	// feature for this interface. This property shall be absent if credential
	// bootstrapping is not supported by the service.
	//
	// Version added: v1.3.0
	CredentialBootstrapping CredentialBootstrapping
	// ExternallyAccessible shall indicate whether external entities can access
	// this interface. External entities are non-host entities. For example, if the
	// host and manager are connected through a switch and the switch also exposes
	// an external port on the system, external clients can also use the interface,
	// and this property value is 'true'.
	ExternallyAccessible bool
	// FirmwareAuthEnabled shall indicate whether firmware authentication is
	// enabled for this interface.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of newer methods of negotiating
	// credentials.
	FirmwareAuthEnabled bool
	// FirmwareAuthRoleID shall contain the 'Id' property of the 'Role' resource
	// that is configured for firmware authentication on this interface.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of newer methods of negotiating
	// credentials.
	FirmwareAuthRoleID string `json:"FirmwareAuthRoleId"`
	// HostEthernetInterfaces shall contain a link to a resource collection of type
	// 'EthernetInterfaceCollection' that computer systems use as the host
	// interface to this manager.
	hostEthernetInterfaces string
	// HostInterfaceType shall contain the host interface type for this interface.
	HostInterfaceType HostInterfaceType
	// InterfaceEnabled shall indicate whether this interface is enabled. Modifying
	// this property may modify the 'InterfaceEnabled' property in the
	// 'EthernetInterface' resource for this host interface.
	InterfaceEnabled bool
	// KernelAuthEnabled shall indicate whether kernel authentication is enabled
	// for this interface.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of newer methods of negotiating
	// credentials.
	KernelAuthEnabled bool
	// KernelAuthRoleID shall contain the 'Id' property of the 'Role' resource that
	// is configured for kernel authentication on this interface.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of newer methods of negotiating
	// credentials.
	KernelAuthRoleID string `json:"KernelAuthRoleId"`
	// ManagerEthernetInterface shall contain a link to a resource of type
	// 'EthernetInterface' that represents the network interface that this manager
	// uses as the host interface.
	managerEthernetInterface string
	// NetworkProtocol shall contain a link to a resource of type
	// 'ManagerNetworkProtocol' that represents the network services for this
	// manager.
	networkProtocol string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// authNoneRole is the URI for AuthNoneRole.
	authNoneRole string
	// computerSystems are the URIs for ComputerSystems.
	computerSystems []string
	// credentialBootstrappingRole is the URI for CredentialBootstrappingRole.
	credentialBootstrappingRole string
	// firmwareAuthRole is the URI for FirmwareAuthRole.
	firmwareAuthRole string
	// kernelAuthRole is the URI for KernelAuthRole.
	kernelAuthRole string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a HostInterface object from the raw JSON.
func (h *HostInterface) UnmarshalJSON(b []byte) error {
	type temp HostInterface
	type hLinks struct {
		AuthNoneRole                Link  `json:"AuthNoneRole"`
		ComputerSystems             Links `json:"ComputerSystems"`
		CredentialBootstrappingRole Link  `json:"CredentialBootstrappingRole"`
		FirmwareAuthRole            Link  `json:"FirmwareAuthRole"`
		KernelAuthRole              Link  `json:"KernelAuthRole"`
	}
	var tmp struct {
		temp
		Links                    hLinks
		HostEthernetInterfaces   Link `json:"HostEthernetInterfaces"`
		ManagerEthernetInterface Link `json:"ManagerEthernetInterface"`
		NetworkProtocol          Link `json:"NetworkProtocol"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = HostInterface(tmp.temp)

	// Extract the links to other entities for later
	h.authNoneRole = tmp.Links.AuthNoneRole.String()
	h.computerSystems = tmp.Links.ComputerSystems.ToStrings()
	h.credentialBootstrappingRole = tmp.Links.CredentialBootstrappingRole.String()
	h.firmwareAuthRole = tmp.Links.FirmwareAuthRole.String()
	h.kernelAuthRole = tmp.Links.KernelAuthRole.String()
	h.hostEthernetInterfaces = tmp.HostEthernetInterfaces.String()
	h.managerEthernetInterface = tmp.ManagerEthernetInterface.String()
	h.networkProtocol = tmp.NetworkProtocol.String()

	// This is a read/write object, so we need to save the raw object data for later
	h.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (h *HostInterface) Update() error {
	readWriteFields := []string{
		"AuthNoneRoleId",
		"AuthenticationModes",
		"FirmwareAuthEnabled",
		"FirmwareAuthRoleId",
		"InterfaceEnabled",
		"KernelAuthEnabled",
		"KernelAuthRoleId",
	}

	return h.UpdateFromRawData(h, h.RawData, readWriteFields)
}

// GetHostInterface will get a HostInterface instance from the service.
func GetHostInterface(c Client, uri string) (*HostInterface, error) {
	return GetObject[HostInterface](c, uri)
}

// ListReferencedHostInterfaces gets the collection of HostInterface from
// a provided reference.
func ListReferencedHostInterfaces(c Client, link string) ([]*HostInterface, error) {
	return GetCollectionObjects[HostInterface](c, link)
}

// AuthNoneRole gets the AuthNoneRole linked resource.
func (h *HostInterface) AuthNoneRole() (*Role, error) {
	if h.authNoneRole == "" {
		return nil, nil
	}
	return GetObject[Role](h.client, h.authNoneRole)
}

// ComputerSystems gets the ComputerSystems linked resources.
func (h *HostInterface) ComputerSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](h.client, h.computerSystems)
}

// CredentialBootstrappingRole gets the CredentialBootstrappingRole linked resource.
func (h *HostInterface) CredentialBootstrappingRole() (*Role, error) {
	if h.credentialBootstrappingRole == "" {
		return nil, nil
	}
	return GetObject[Role](h.client, h.credentialBootstrappingRole)
}

// FirmwareAuthRole gets the FirmwareAuthRole linked resource.
func (h *HostInterface) FirmwareAuthRole() (*Role, error) {
	if h.firmwareAuthRole == "" {
		return nil, nil
	}
	return GetObject[Role](h.client, h.firmwareAuthRole)
}

// KernelAuthRole gets the KernelAuthRole linked resource.
func (h *HostInterface) KernelAuthRole() (*Role, error) {
	if h.kernelAuthRole == "" {
		return nil, nil
	}
	return GetObject[Role](h.client, h.kernelAuthRole)
}

// HostEthernetInterfaces gets the HostEthernetInterfaces collection.
func (h *HostInterface) HostEthernetInterfaces() ([]*EthernetInterface, error) {
	if h.hostEthernetInterfaces == "" {
		return nil, nil
	}
	return GetCollectionObjects[EthernetInterface](h.client, h.hostEthernetInterfaces)
}

// ManagerEthernetInterface gets the ManagerEthernetInterface linked resource.
func (h *HostInterface) ManagerEthernetInterface() (*EthernetInterface, error) {
	if h.managerEthernetInterface == "" {
		return nil, nil
	}
	return GetObject[EthernetInterface](h.client, h.managerEthernetInterface)
}

// NetworkProtocol gets the NetworkProtocol linked resource.
func (h *HostInterface) NetworkProtocol() (*ManagerNetworkProtocol, error) {
	if h.networkProtocol == "" {
		return nil, nil
	}
	return GetObject[ManagerNetworkProtocol](h.client, h.networkProtocol)
}

// CredentialBootstrapping shall contain settings for the Redfish Host Interface
// Specification-defined 'credential bootstrapping via IPMI commands' feature
// for this interface.
type CredentialBootstrapping struct {
	// EnableAfterReset shall indicate whether credential bootstrapping is enabled
	// after a reset for this interface. If 'true', services shall set the
	// 'Enabled' property to 'true' after a reset of the host or the service.
	//
	// Version added: v1.3.0
	EnableAfterReset bool
	// Enabled shall indicate whether credential bootstrapping is enabled for this
	// interface.
	//
	// Version added: v1.3.0
	Enabled bool
	// RoleID shall contain the 'Id' property of the 'Role' resource that is used
	// for the bootstrap account created for this interface.
	//
	// Version added: v1.3.0
	RoleID string `json:"RoleId"`
}
