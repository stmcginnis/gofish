// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// DefaultHostInterfacePath is the default URI for the HostInterface
// object.
const DefaultHostInterfacePath = "/redfish/v1/HostInterface"

// AuthenticationMode is
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

// HostInterfaceType is
type HostInterfaceType string

const (

	// NetworkHostInterfaceHostInterfaceType This interface is a Network Host
	// Interface.
	NetworkHostInterfaceHostInterfaceType HostInterfaceType = "NetworkHostInterface"
)

// HostInterface is used to represent Host Interface resources as part of
// the Redfish specification.
type HostInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AuthNoneRoleID is used when no authentication on this interface is
	// performed. This property shall be absent if AuthNone is not supported
	// by the service for the AuthenticationModes property.
	AuthNoneRoleID string `json:"AuthNoneRoleId"`
	// AuthenticationModes shall be an array consisting of the authentication
	// modes allowed on this interface.
	AuthenticationModes []AuthenticationMode
	// Description provides a description of this resource.
	Description string
	// ExternallyAccessible is used by external clients, and this property
	// will have the value set to true.
	ExternallyAccessible bool
	// FirmwareAuthEnabled shall be a boolean
	// indicating whether firmware authentication for this interface is
	// enabled.
	FirmwareAuthEnabled bool
	// FirmwareAuthRoleID shall be the ID of the Role resource that is
	// configured for firmware authentication on this interface.
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
	KernelAuthEnabled bool
	// KernelAuthRoleID shall be the ID of the Role resource that is configured
	// for kernel authentication on this interface.
	KernelAuthRoleID string `json:"KernelAuthRoleId"`
	// ManagerEthernetInterface is used by this Manager as the HostInterface.
	managerEthernetInterface string
	// NetworkProtocol shall contain a
	// reference to a resource of type ManagerNetworkProtocol which
	// represents the network services for this Manager.
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
	// FirmwareAuthRole shall be a link to a Role object instance, and should
	// reference the object identified by property FirmwareAuthRoleID.
	firmwareAuthRole string
	// KernelAuthRole shall be a link to a Role object instance, and should
	// reference the object identified by property KernelAuthRoleId.
	kernelAuthRole string
}

// UnmarshalJSON unmarshals a HostInterface object from the raw JSON.
func (hostinterface *HostInterface) UnmarshalJSON(b []byte) error {
	type temp HostInterface

	type links struct {
		AuthNoneRole         common.Link
		ComputerSystems      common.Links
		ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
		FirmwareAuthRole     common.Link
		KernelAuthRole       common.Link
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
	hostinterface.authNoneRole = string(t.Links.AuthNoneRole)
	hostinterface.computerSystems = t.Links.ComputerSystems.ToStrings()
	hostinterface.ComputerSystemsCount = t.Links.ComputerSystemsCount
	hostinterface.firmwareAuthRole = string(t.Links.FirmwareAuthRole)
	hostinterface.kernelAuthRole = string(t.Links.KernelAuthRole)
	hostinterface.hostEthernetInterfaces = string(t.HostEthernetInterfaces)
	hostinterface.managerEthernetInterface = string(t.ManagerEthernetInterface)
	hostinterface.networkProtocol = string(t.NetworkProtocol)

	return nil
}

// GetHostInterface will get a HostInterface instance from the service.
func GetHostInterface(c common.Client, uri string) (*HostInterface, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var hostinterface HostInterface
	err = json.NewDecoder(resp.Body).Decode(&hostinterface)
	if err != nil {
		return nil, err
	}

	hostinterface.SetClient(c)
	return &hostinterface, nil
}

// ListReferencedHostInterfaces gets the collection of HostInterface from
// a provided reference.
func ListReferencedHostInterfaces(c common.Client, link string) ([]*HostInterface, error) {
	var result []*HostInterface
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, hostinterfaceLink := range links.ItemLinks {
		hostinterface, err := GetHostInterface(c, hostinterfaceLink)
		if err != nil {
			return result, err
		}
		result = append(result, hostinterface)
	}

	return result, nil
}

// TODO: Add access functions for linked objects
