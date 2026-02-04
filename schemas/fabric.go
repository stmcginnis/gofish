//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Fabric.v1_4_0.json
// 2025.2 - #Fabric.v1_4_0.Fabric

package schemas

import (
	"encoding/json"
)

// Fabric shall represent a simple switchable fabric for a Redfish
// implementation.
type Fabric struct {
	Entity
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
	FabricType Protocol
	// MaxZones shall contain the maximum number of zones the switch can currently
	// configure. Changes in the logical or physical configuration of the system
	// can change this value.
	MaxZones *int `json:",omitempty"`
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Fabric object from the raw JSON.
func (f *Fabric) UnmarshalJSON(b []byte) error {
	type temp Fabric
	type fLinks struct {
		ManagedBy Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Links          fLinks
		AddressPools   Link `json:"AddressPools"`
		Connections    Link `json:"Connections"`
		EndpointGroups Link `json:"EndpointGroups"`
		Endpoints      Link `json:"Endpoints"`
		Switches       Link `json:"Switches"`
		Zones          Link `json:"Zones"`
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
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *Fabric) Update() error {
	readWriteFields := []string{
		"UUID",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetFabric will get a Fabric instance from the service.
func GetFabric(c Client, uri string) (*Fabric, error) {
	return GetObject[Fabric](c, uri)
}

// ListReferencedFabrics gets the collection of Fabric from
// a provided reference.
func ListReferencedFabrics(c Client, link string) ([]*Fabric, error) {
	return GetCollectionObjects[Fabric](c, link)
}

// ManagedBy gets the ManagedBy linked resources.
func (f *Fabric) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](f.client, f.managedBy)
}

// AddressPools gets the AddressPools collection.
func (f *Fabric) AddressPools() ([]*AddressPool, error) {
	if f.addressPools == "" {
		return nil, nil
	}
	return GetCollectionObjects[AddressPool](f.client, f.addressPools)
}

// Connections gets the Connections collection.
func (f *Fabric) Connections() ([]*Connection, error) {
	if f.connections == "" {
		return nil, nil
	}
	return GetCollectionObjects[Connection](f.client, f.connections)
}

// EndpointGroups gets the EndpointGroups collection.
func (f *Fabric) EndpointGroups() ([]*EndpointGroup, error) {
	if f.endpointGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[EndpointGroup](f.client, f.endpointGroups)
}

// Endpoints gets the Endpoints collection.
func (f *Fabric) Endpoints() ([]*Endpoint, error) {
	if f.endpoints == "" {
		return nil, nil
	}
	return GetCollectionObjects[Endpoint](f.client, f.endpoints)
}

// Switches gets the Switches collection.
func (f *Fabric) Switches() ([]*Switch, error) {
	if f.switches == "" {
		return nil, nil
	}
	return GetCollectionObjects[Switch](f.client, f.switches)
}

// Zones gets the Zones collection.
func (f *Fabric) Zones() ([]*Zone, error) {
	if f.zones == "" {
		return nil, nil
	}
	return GetCollectionObjects[Zone](f.client, f.zones)
}
