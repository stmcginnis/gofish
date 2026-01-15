//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.3 - #EndpointGroup.v1_3_4.EndpointGroup

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// AccessState is This type shall describe the access to all associated
// resources through all aggregated endpoints.
type AccessState string

const (
	// OptimizedAccessState shall indicate each endpoint is in an active and
	// optimized state.
	OptimizedAccessState AccessState = "Optimized"
	// NonOptimizedAccessState shall indicate each endpoint is in an active and
	// non-optimized state.
	NonOptimizedAccessState AccessState = "NonOptimized"
	// StandbyAccessState shall indicate each endpoint is in a standby state.
	StandbyAccessState AccessState = "Standby"
	// UnavailableAccessState shall indicate each endpoint is in an unavailable
	// state.
	UnavailableAccessState AccessState = "Unavailable"
	// TransitioningAccessState shall indicate each endpoint is transitioning to a
	// new state.
	TransitioningAccessState AccessState = "Transitioning"
)

type GroupType string

const (
	// ClientGroupType shall indicate that the endpoint group contains client
	// (initiator) endpoints. If the associated endpoints contain the 'EntityRole'
	// property, the 'EntityRole' property shall contain the value 'Initiator' or
	// 'Both'.
	ClientGroupType GroupType = "Client"
	// ServerGroupType shall indicate that the endpoint group contains server
	// (target) endpoints. If the associated endpoints contain the 'EntityRole'
	// property, the 'EntityRole' property shall contain the value 'Target' or
	// 'Both'.
	ServerGroupType GroupType = "Server"
	// InitiatorGroupType shall indicate that the endpoint group contains initiator
	// endpoints. If the associated endpoints contain the 'EntityRole' property,
	// the 'EntityRole' property shall contain the value 'Initiator' or 'Both'.
	InitiatorGroupType GroupType = "Initiator"
	// TargetGroupType shall indicate that the endpoint group contains target
	// endpoints. If the associated endpoints contain the 'EntityRole' property,
	// the 'EntityRole' property shall contain the value 'Target' or 'Both'.
	TargetGroupType GroupType = "Target"
)

// EndpointGroup shall represent a group of endpoints that are managed as a unit
// for a Redfish implementation.
type EndpointGroup struct {
	common.Entity
	// AccessState shall contain the access state for all associated resources in
	// this endpoint group.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'AccessState' property in
	// the connection resource.
	AccessState AccessState
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// GroupType shall contain the endpoint group type. If this endpoint group
	// represents a SCSI target group, the value of this property shall contain
	// 'Server' or 'Target'.
	GroupType GroupType
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Preferred shall indicate if access to the resources through the endpoint
	// group is preferred over access through other endpoints. The default value
	// for this property is 'false'.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'AccessState' property in
	// the connection resource.
	Preferred bool
	// TargetEndpointGroupIdentifier shall contain a SCSI-defined identifier for
	// this group that corresponds to the TARGET PORT GROUP field in the REPORT
	// TARGET PORT GROUPS response and the TARGET PORT GROUP field in an INQUIRY
	// VPD page 85 response, type 5h identifier. See the INCITS SAM-5
	// specification. This property may not be present if the endpoint group does
	// not represent a SCSI target group.
	TargetEndpointGroupIdentifier *int `json:",omitempty"`
	// connections are the URIs for Connections.
	connections []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (e *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	type eLinks struct {
		Connections common.Links `json:"Connections"`
	}
	var tmp struct {
		temp
		Links     eLinks
		Endpoints common.Links `json:"Endpoints"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*e = EndpointGroup(tmp.temp)

	// Extract the links to other entities for later
	e.connections = tmp.Links.Connections.ToStrings()
	e.endpoints = tmp.Endpoints.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	e.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (e *EndpointGroup) Update() error {
	readWriteFields := []string{
		"AccessState",
		"Endpoints",
		"Endpoints@odata.count",
		"GroupType",
		"Identifier",
		"Preferred",
		"TargetEndpointGroupIdentifier",
	}

	return e.UpdateFromRawData(e, e.rawData, readWriteFields)
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c common.Client, uri string) (*EndpointGroup, error) {
	return common.GetObject[EndpointGroup](c, uri)
}

// ListReferencedEndpointGroups gets the collection of EndpointGroup from
// a provided reference.
func ListReferencedEndpointGroups(c common.Client, link string) ([]*EndpointGroup, error) {
	return common.GetCollectionObjects[EndpointGroup](c, link)
}

// Connections gets the Connections linked resources.
func (e *EndpointGroup) Connections(client common.Client) ([]*redfish.Connection, error) {
	return common.GetObjects[redfish.Connection](client, e.connections)
}

// Endpoints gets the Endpoints linked resources.
func (e *EndpointGroup) Endpoints(client common.Client) ([]*redfish.Endpoint, error) {
	return common.GetObjects[redfish.Endpoint](client, e.endpoints)
}
