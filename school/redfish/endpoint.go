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

// EntityRole is the role of the endpoint.
type EntityRole string

const (
	// InitiatorEntityRole means the entity is acting as an initiator.
	InitiatorEntityRole EntityRole = "Initiator"
	// TargetEntityRole means the entity is acting as a target.
	TargetEntityRole EntityRole = "Target"
	// BothEntityRole means the entity is acting as both an initiator and a target.
	BothEntityRole EntityRole = "Both"
)

// EntityType is the type of endpoint.
type EntityType string

const (
	// StorageInitiatorEntityType means the entity is a storage initator. The
	// EntityLink property (if present) should be a Storage.StorageController
	// entity.
	StorageInitiatorEntityType EntityType = "StorageInitiator"
	// RootComplexEntityType means the entity is a PCI(e) root complex. The
	// EntityLink property (if present) should be a
	// ComputerSystem.ComputerSystem entity.
	RootComplexEntityType EntityType = "RootComplex"
	// NetworkControllerEntityType means the entity is a network controller. The
	// EntityLink property (if present) should be an
	// EthernetInterface.EthernetInterface entity.
	NetworkControllerEntityType EntityType = "NetworkController"
	// DriveEntityType means the entity is a disk drive. The EntityLink property
	// (if present) should be a Drive.Drive entity.
	DriveEntityType EntityType = "Drive"
	// StorageExpanderEntityType means the entity is a storage expander. The
	// EntityLink property (if present) should be a Chassis.Chassis entity.
	StorageExpanderEntityType EntityType = "StorageExpander"
	// DisplayControllerEntityType means the entity is a display controller.
	DisplayControllerEntityType EntityType = "DisplayController"
	// BridgeEntityType means the entity is a PCI(e) bridge.
	BridgeEntityType EntityType = "Bridge"
	// ProcessorEntityType means the entity is a processor device.
	ProcessorEntityType EntityType = "Processor"
	// VolumeEntityType means the entity is a volume. The EntityLink property (if
	// present) should be a Volume.Volume entity.
	VolumeEntityType EntityType = "Volume"
	// AccelerationFunctionEntityType means the entity is an acceleration function
	// realized through a device, such as an FPGA. The EntityLink property
	// (if present) should be a AccelerationFunction.AccelerationFunction
	// entity.
	AccelerationFunctionEntityType EntityType = "AccelerationFunction"
)

// ConnectedEntity shall represent a remote resource that is
// connected to a network accessible to an endpoint.
type ConnectedEntity struct {
	// entityLink shall be a reference to an entity of the
	// type specified by the description of the value of the EntityType
	// property.
	entityLink common.Link
	// EntityPciID shall be the PCI ID of the connected PCIe entity.
	EntityPciID PciID `json:"entityPciId"`
	// entityRole shall indicate if the specified entity is an initiator,
	// target, or both.
	EntityRole EntityRole
	// entityType shall indicate if type of connected entity.
	EntityType EntityType
	// Identifiers for the remote entity shall be unique in
	// the context of other resources that can reached over the connected
	// network.
	Identifiers []common.Identifier
	// PciClassCode shall be the PCI Class
	// Code, Subclass code, and Programming Interface code of the PCIe device
	// function.
	PciClassCode string
	// PciFunctionNumber shall be the PCI
	// Function Number of the connected PCIe entity.
	PciFunctionNumber int
}

// Endpoint is used to represent a fabric endpoint for a Redfish
// implementation.
type Endpoint struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConnectedEntities shall contain all the entities which this endpoint
	// allows access to.
	ConnectedEntities string
	// Description provides a description of this resource.
	Description string
	// endpointProtocol shall contain the protocol this endpoint uses to
	// communicate with other endpoints on this fabric.
	endpointProtocol string
	// HostReservationMemoryBytes shall be the amount of memory in Bytes that
	// the Host should allocate to connect to this endpoint.
	HostReservationMemoryBytes int
	// IPTransportDetails shall contain the details for each IP transport
	// supported by this endpoint.
	IPTransportDetails string
	// ID uniquely identifies the resource.
	ID string `json:"Id"`
	// Identifiers shall be unique in the context of other endpoints that can
	// reached over the connected network.
	Identifiers []common.Identifier
	// Name is the name of the resource or array element.
	Name string
	// PciID shall be the PCI ID of the endpoint.
	PciID PciID `json:"PciId"`
	// Redundancy is used to show how this endpoint is grouped with other
	// endpoints to form redundancy sets.
	redundancy []string
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a Endpoint object from the raw JSON.
func (endpoint *Endpoint) UnmarshalJSON(b []byte) error {
	type temp Endpoint
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpoint = Endpoint(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetEndpoint will get a Endpoint instance from the service.
func GetEndpoint(c common.Client, uri string) (*Endpoint, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endpoint Endpoint
	err = json.NewDecoder(resp.Body).Decode(&endpoint)
	if err != nil {
		return nil, err
	}

	endpoint.SetClient(c)
	return &endpoint, nil
}

// ListReferencedEndpoints gets the collection of Endpoint from
// a provided reference.
func ListReferencedEndpoints(c common.Client, link string) ([]*Endpoint, error) {
	var result []*Endpoint
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, endpointLink := range links.ItemLinks {
		endpoint, err := GetEndpoint(c, endpointLink)
		if err != nil {
			return result, err
		}
		result = append(result, endpoint)
	}

	return result, nil
}

// IPTransportDetails shall contain properties which specify
// the details of the transport supported by the endpoint.
type IPTransportDetails struct {
	// IPv4Address shall specify the
	// IPv4Address.
	IPv4Address string
	// IPv6Address shall specify the
	// IPv6Address.
	IPv6Address string
	// Port is used for communication with the Endpoint.
	Port int
	// TransportProtocol is used by the connection entity.
	TransportProtocol string
}

// Links is This type, as described by the Redfish Specification, shall
// contain references to resources that are related to, but not contained
// by (subordinate to), this resource.
type Links struct {
	// MutuallyExclusiveEndpoints is used in a zone if this endpoint is used
	// in a zone.
	MutuallyExclusiveEndpoints common.Link
	// MutuallyExclusiveEndpointsCount is the number of MutuallyExclusiveEndpoints.
	MutuallyExclusiveEndpointsCount int `json:"MutuallyExclusiveEndpoints@odata.count"`
	// NetworkDeviceFunction shall be a reference to a NetworkDeviceFunction
	// resource, with which this endpoint is associated.
	NetworkDeviceFunction common.Link
	// NetworkDeviceFunctionCount is the number of NetworkDeviceFunctions.
	NetworkDeviceFunctionCount int `json:"NetworkDeviceFunction@odata.count"`
	// Ports shall be an array of references of type Port that are utilized by
	// this endpoint.
	Ports []common.Link
	// PortsCount is the number of Ports.
	PortsCount int `json:"Ports@odata.count"`
}

// PciID shall describe a PCI ID.
type PciID struct {
	// ClassCode shall be the PCI Class Code,
	// Subclass code, and Programming Interface code of the PCIe device
	// function.
	ClassCode string
	// DeviceID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	DeviceID string `json:"DeviceId"`
	// FunctionNumber shall be the PCI Function Number of the connected PCIe entity.
	FunctionNumber string
	// SubsystemID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemID string `json:"SubsystemId"`
	// SubsystemVendorID shall be the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemVendorID string `json:"SubsystemVendorId"`
	// VendorID shall be the PCI Vendor ID of the PCIe device function.
	VendorID string `json:"VendorId"`
}
