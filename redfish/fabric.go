//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// Fabric shall represent a simple switchable fabric for a Redfish implementation.
type Fabric struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AddressPools shall contain a link to a resource collection of type AddressPoolCollection.
	addressPools string
	// Connections shall contain a link to a resource collection of type ConnectionCollection.
	connections string
	// Description provides a description of this resource.
	Description string
	// EndpointGroups shall contain a link to a resource collection of type EndpointGroupCollection.
	endpointGroups string
	// Endpoints shall contain a link to a resource collection of type EndpointCollection.
	endpoints string
	// FabricType shall contain the type of fabric being represented by this simple fabric.
	FabricType common.Protocol
	// MaxZones shall contain the maximum number of zones the switch can currently configure. Changes in the logical or
	// physical configuration of the system can change this value.
	MaxZones int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Switches shall contain a link to a resource collection of type SwitchCollection.
	switches string
	// UUID shall contain a universally unique identifier number for the fabric.
	UUID string
	// Zones shall contain a link to a resource collection of type ZoneCollection.
	zones string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Fabric object from the raw JSON.
func (fabric *Fabric) UnmarshalJSON(b []byte) error {
	type temp Fabric
	var t struct {
		temp
		AddressPools   common.Link
		Connections    common.Link
		EndpointGroups common.Link
		Endpoints      common.Link
		Switches       common.Link
		Zones          common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fabric = Fabric(t.temp)

	// Extract the links to other entities for later
	fabric.addressPools = t.AddressPools.String()
	fabric.connections = t.Connections.String()
	fabric.endpointGroups = t.EndpointGroups.String()
	fabric.endpoints = t.Endpoints.String()
	fabric.switches = t.Switches.String()
	fabric.zones = t.Zones.String()

	// This is a read/write object, so we need to save the raw object data for later
	fabric.rawData = b

	return nil
}

// AddressPools gets any address pools associated with this fabric.
func (fabric *Fabric) AddressPools(queryOpts ...common.QueryGroupOption) ([]*AddressPool, error) {
	return fabric.AddressPoolsWithContext(common.ContextOf(fabric.GetClient()), queryOpts...)
}

// AddressPoolsWithContext gets any address pools associated with this fabric.
func (fabric *Fabric) AddressPoolsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*AddressPool, error) {
	return ListReferencedAddressPoolsWithContext(ctx, fabric.GetClient(), fabric.addressPools, queryOpts...)
}

// Connections gets any connections associated with this fabric.
func (fabric *Fabric) Connections(queryOpts ...common.QueryGroupOption) ([]*Connection, error) {
	return fabric.ConnectionsWithContext(common.ContextOf(fabric.GetClient()), queryOpts...)
}

// ConnectionsWithContext gets any connections associated with this fabric.
func (fabric *Fabric) ConnectionsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Connection, error) {
	return ListReferencedConnectionsWithContext(ctx, fabric.GetClient(), fabric.connections, queryOpts...)
}

// EndpointGroups gets any endpoint groups associated with this fabric.
func (fabric *Fabric) EndpointGroups(queryOpts ...common.QueryGroupOption) ([]*EndpointGroup, error) {
	return fabric.EndpointGroupsWithContext(common.ContextOf(fabric.GetClient()), queryOpts...)
}

// EndpointGroupsWithContext gets any endpoint groups associated with this fabric.
func (fabric *Fabric) EndpointGroupsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroupsWithContext(ctx, fabric.GetClient(), fabric.endpointGroups, queryOpts...)
}

// Endpoints gets any endpoints associated with this fabric.
func (fabric *Fabric) Endpoints(queryOpts ...common.QueryGroupOption) ([]*Endpoint, error) {
	return fabric.EndpointsWithContext(common.ContextOf(fabric.GetClient()), queryOpts...)
}

// EndpointsWithContext gets any endpoints associated with this fabric.
func (fabric *Fabric) EndpointsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Endpoint, error) {
	return ListReferencedEndpointsWithContext(ctx, fabric.GetClient(), fabric.endpoints, queryOpts...)
}

// Switches gets any switches associated with this fabric.
func (fabric *Fabric) Switches(queryOpts ...common.QueryGroupOption) ([]*Switch, error) {
	return fabric.SwitchesWithContext(common.ContextOf(fabric.GetClient()), queryOpts...)
}

// SwitchesWithContext gets any switches associated with this fabric.
func (fabric *Fabric) SwitchesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Switch, error) {
	return ListReferencedSwitchesWithContext(ctx, fabric.GetClient(), fabric.switches, queryOpts...)
}

// Update commits updates to this object's properties to the running system.
func (fabric *Fabric) Update() error {
	return fabric.UpdateWithContext(common.ContextOf(fabric.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (fabric *Fabric) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"UUID",
	}

	return fabric.UpdateFromRawDataWithContext(ctx, fabric, fabric.rawData, readWriteFields)
}

// GetFabric will get a Fabric instance from the service.
func GetFabric(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Fabric, error) {
	return GetFabricWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetFabricWithContext will get a Fabric instance from the service.
func GetFabricWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Fabric, error) {
	return common.GetObjectWithContext[Fabric](ctx, c, uri, queryOpts...)
}

// ListReferencedFabrics gets the collection of Fabric from
// a provided reference.
func ListReferencedFabrics(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Fabric, error) {
	return ListReferencedFabricsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedFabricsWithContext gets the collection of Fabric from
// a provided reference.
func ListReferencedFabricsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Fabric, error) {
	return common.GetCollectionObjectsWithContext[Fabric](ctx, c, link, queryOpts...)
}
