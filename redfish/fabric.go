//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Fabric shall represent a simple switchable fabric for a Redfish implementation.
type Fabric struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
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
func (fabric *Fabric) AddressPools() ([]*AddressPool, error) {
	return ListReferencedAddressPools(fabric.GetClient(), fabric.addressPools)
}

// Connections gets any connections associated with this fabric.
func (fabric *Fabric) Connections() ([]*Connection, error) {
	return ListReferencedConnections(fabric.GetClient(), fabric.connections)
}

// EndpointGroups gets any endpoint groups associated with this fabric.
func (fabric *Fabric) EndpointGroups() ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(fabric.GetClient(), fabric.endpointGroups)
}

// Endpoints gets any endpoints associated with this fabric.
func (fabric *Fabric) Endpoints() ([]*Endpoint, error) {
	return ListReferencedEndpoints(fabric.GetClient(), fabric.endpoints)
}

// Switches gets any switches associated with this fabric.
func (fabric *Fabric) Switches() ([]*Switch, error) {
	return ListReferencedSwitches(fabric.GetClient(), fabric.switches)
}

// Update commits updates to this object's properties to the running system.
func (fabric *Fabric) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Fabric)
	original.UnmarshalJSON(fabric.rawData)

	readWriteFields := []string{
		"UUID",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fabric).Elem()

	return fabric.Entity.Update(originalElement, currentElement, readWriteFields)
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
