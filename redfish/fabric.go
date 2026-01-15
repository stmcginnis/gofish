//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #Fabric.v1_4_0.Fabric

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Fabric shall represent a simple switchable fabric for a Redfish
// implementation.
type Fabric struct {
	common.Entity
	// AddressPools shall contain a link to a resource collection of type
	// 'AddressPoolCollection'.
	//
	// Version added: v1.1.0
	addressPools string
	// Connections shall contain a link to a resource collection of type
	// 'ConnectionCollection'.
	//
	// Version added: v1.2.0
	connections string
	// EndpointGroups shall contain a link to a resource collection of type
	// 'EndpointGroupCollection'.
	//
	// Version added: v1.2.0
	endpointGroups string
	// Endpoints shall contain a link to a resource collection of type
	// 'EndpointCollection'.
	endpoints string
	// FabricType shall contain the type of fabric being represented by this simple
	// fabric.
	FabricType common.Protocol
	// MaxZones shall contain the maximum number of zones the switch can currently
	// configure. Changes in the logical or physical configuration of the system
	// can change this value.
	MaxZones *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Switches shall contain a link to a resource collection of type
	// 'SwitchCollection'.
	switches string
	// UUID shall contain a universally unique identifier number for the fabric.
	//
	// Version added: v1.3.0
	UUID string
	// Zones shall contain a link to a resource collection of type
	// 'ZoneCollection'.
	zones string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Fabric object from the raw JSON.
func (f *Fabric) UnmarshalJSON(b []byte) error {
	type temp Fabric
	type fLinks struct {
		ManagedBy common.Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Links          fLinks
		AddressPools   common.Link `json:"addressPools"`
		Connections    common.Link `json:"connections"`
		EndpointGroups common.Link `json:"endpointGroups"`
		Endpoints      common.Link `json:"endpoints"`
		Switches       common.Link `json:"switches"`
		Zones          common.Link `json:"zones"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = Fabric(tmp.temp)

	// Extract the links to other entities for later
	f.managedBy = tmp.Links.ManagedBy.ToStrings()
	f.addressPools = tmp.AddressPools.String()
	f.connections = tmp.Connections.String()
	f.endpointGroups = tmp.EndpointGroups.String()
	f.endpoints = tmp.Endpoints.String()
	f.switches = tmp.Switches.String()
	f.zones = tmp.Zones.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *Fabric) Update() error {
	readWriteFields := []string{
		"Status",
		"UUID",
	}

	return f.UpdateFromRawData(f, f.rawData, readWriteFields)
}

// GetFabric will get a Fabric instance from the service.
func GetFabric(c common.Client, uri string) (*Fabric, error) {
	return common.GetObject[Fabric](c, uri)
}

// ListReferencedFabrics gets the collection of Fabric from
// a provided reference.
func ListReferencedFabrics(c common.Client, link string) ([]*Fabric, error) {
	return common.GetCollectionObjects[Fabric](c, link)
}

// ManagedBy gets the ManagedBy linked resources.
func (f *Fabric) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, f.managedBy)
}

// AddressPools gets the AddressPools collection.
func (f *Fabric) AddressPools(client common.Client) ([]*AddressPool, error) {
	if f.addressPools == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[AddressPool](client, f.addressPools)
}

// Connections gets the Connections collection.
func (f *Fabric) Connections(client common.Client) ([]*Connection, error) {
	if f.connections == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Connection](client, f.connections)
}

// EndpointGroups gets the EndpointGroups collection.
func (f *Fabric) EndpointGroups(client common.Client) ([]*EndpointGroup, error) {
	if f.endpointGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EndpointGroup](client, f.endpointGroups)
}

// Endpoints gets the Endpoints collection.
func (f *Fabric) Endpoints(client common.Client) ([]*Endpoint, error) {
	if f.endpoints == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Endpoint](client, f.endpoints)
}

// Switches gets the Switches collection.
func (f *Fabric) Switches(client common.Client) ([]*Switch, error) {
	if f.switches == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Switch](client, f.switches)
}

// Zones gets the Zones collection.
func (f *Fabric) Zones(client common.Client) ([]*Zone, error) {
	if f.zones == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Zone](client, f.zones)
}
