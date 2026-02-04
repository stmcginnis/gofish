//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Zone.v1_6_3.json
// 2020.4 - #Zone.v1_6_3.Zone

package schemas

import (
	"encoding/json"
)

type ExternalAccessibility string

const (
	// GloballyAccessibleExternalAccessibility shall indicate that any external
	// entity with the correct access details, which may include authorization
	// information, can access the endpoints that this zone lists, regardless of
	// zone.
	GloballyAccessibleExternalAccessibility ExternalAccessibility = "GloballyAccessible"
	// NonZonedAccessibleExternalAccessibility shall indicate that any external
	// entity that another zone does not explicitly list can access the endpoints
	// that this zone lists.
	NonZonedAccessibleExternalAccessibility ExternalAccessibility = "NonZonedAccessible"
	// ZoneOnlyExternalAccessibility shall indicate that endpoints in this zone are
	// only accessible by endpoints that this zone explicitly lists.
	ZoneOnlyExternalAccessibility ExternalAccessibility = "ZoneOnly"
	// NoInternalRoutingExternalAccessibility shall indicate that implicit routing
	// within this zone is not defined.
	NoInternalRoutingExternalAccessibility ExternalAccessibility = "NoInternalRouting"
)

type ZoneType string

const (
	// DefaultZoneType shall indicate a zone in which all endpoints are added by
	// default when instantiated. This value shall only be used for zones
	// subordinate to the fabric collection.
	DefaultZoneType ZoneType = "Default"
	// ZoneOfEndpointsZoneType shall indicate a zone that contains resources of
	// type 'Endpoint'. This value shall only be used for zones subordinate to the
	// fabric collection.
	ZoneOfEndpointsZoneType ZoneType = "ZoneOfEndpoints"
	// ZoneOfZonesZoneType shall indicate a zone that contains resources of type
	// 'Zone'. This value shall only be used for zones subordinate to the fabric
	// collection.
	ZoneOfZonesZoneType ZoneType = "ZoneOfZones"
	// ZoneOfResourceBlocksZoneType shall indicate a zone that contains resources
	// of type 'ResourceBlock'. This value shall only be used for zones subordinate
	// to the composition service.
	ZoneOfResourceBlocksZoneType ZoneType = "ZoneOfResourceBlocks"
)

// Zone shall represent a simple fabric zone for a Redfish implementation.
type Zone struct {
	Entity
	// DefaultRoutingEnabled shall indicate whether routing within this zone is
	// enabled.
	//
	// Version added: v1.4.0
	DefaultRoutingEnabled bool
	// ExternalAccessibility shall contain and indication of accessibility of
	// endpoints in this zone to endpoints outside of this zone.
	//
	// Version added: v1.3.0
	ExternalAccessibility ExternalAccessibility
	// Identifiers shall contain a list of all known durable names for the
	// associated zone.
	//
	// Version added: v1.2.0
	Identifiers []Identifier
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
	// ZoneType shall contain the type of zone that this zone represents.
	//
	// Version added: v1.4.0
	ZoneType ZoneType
	// addEndpointTarget is the URL to send AddEndpoint requests.
	addEndpointTarget string
	// removeEndpointTarget is the URL to send RemoveEndpoint requests.
	removeEndpointTarget string
	// addressPools are the URIs for AddressPools.
	addressPools []string
	// containedByZones are the URIs for ContainedByZones.
	containedByZones []string
	// containsZones are the URIs for ContainsZones.
	containsZones []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// involvedSwitches are the URIs for InvolvedSwitches.
	involvedSwitches []string
	// resourceBlocks are the URIs for ResourceBlocks.
	resourceBlocks []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Zone object from the raw JSON.
func (z *Zone) UnmarshalJSON(b []byte) error {
	type temp Zone
	type zActions struct {
		AddEndpoint    ActionTarget `json:"#Zone.AddEndpoint"`
		RemoveEndpoint ActionTarget `json:"#Zone.RemoveEndpoint"`
	}
	type zLinks struct {
		AddressPools     Links `json:"AddressPools"`
		ContainedByZones Links `json:"ContainedByZones"`
		ContainsZones    Links `json:"ContainsZones"`
		Endpoints        Links `json:"Endpoints"`
		InvolvedSwitches Links `json:"InvolvedSwitches"`
		ResourceBlocks   Links `json:"ResourceBlocks"`
	}
	var tmp struct {
		temp
		Actions zActions
		Links   zLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*z = Zone(tmp.temp)

	// Extract the links to other entities for later
	z.addEndpointTarget = tmp.Actions.AddEndpoint.Target
	z.removeEndpointTarget = tmp.Actions.RemoveEndpoint.Target
	z.addressPools = tmp.Links.AddressPools.ToStrings()
	z.containedByZones = tmp.Links.ContainedByZones.ToStrings()
	z.containsZones = tmp.Links.ContainsZones.ToStrings()
	z.endpoints = tmp.Links.Endpoints.ToStrings()
	z.involvedSwitches = tmp.Links.InvolvedSwitches.ToStrings()
	z.resourceBlocks = tmp.Links.ResourceBlocks.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	z.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (z *Zone) Update() error {
	readWriteFields := []string{
		"DefaultRoutingEnabled",
		"ExternalAccessibility",
		"ZoneType",
	}

	return z.UpdateFromRawData(z, z.RawData, readWriteFields)
}

// GetZone will get a Zone instance from the service.
func GetZone(c Client, uri string) (*Zone, error) {
	return GetObject[Zone](c, uri)
}

// ListReferencedZones gets the collection of Zone from
// a provided reference.
func ListReferencedZones(c Client, link string) ([]*Zone, error) {
	return GetCollectionObjects[Zone](c, link)
}

// This action shall add an endpoint to a zone.
// endpoint - This parameter shall contain a link to the specified endpoint to
// add to the zone.
// endpointETag - This parameter shall contain the current ETag of the endpoint
// to add to the zone. If the client-provided ETag does not match the current
// ETag of the endpoint that the 'Endpoint' parameter specifies, the service
// shall return the HTTP '428 Precondition Required' status code to reject the
// request.
// zoneETag - This parameter shall contain the current ETag of the zone. If the
// client-provided ETag does not match the current ETag of the zone, the
// service shall return the HTTP '428 Precondition Required' status code to
// reject the request.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (z *Zone) AddEndpoint(endpoint string, endpointETag string, zoneETag string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Endpoint"] = endpoint
	payload["EndpointETag"] = endpointETag
	payload["ZoneETag"] = zoneETag
	resp, taskInfo, err := PostWithTask(z.client,
		z.addEndpointTarget, payload, z.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall remove an endpoint from a zone.
// endpoint - This parameter shall contain a link to the specified endpoint to
// remove from the zone.
// endpointETag - This parameter shall contain the current ETag of the endpoint
// to remove from the system. If the client-provided ETag does not match the
// current ETag of the endpoint that the 'Endpoint' parameter specifies, the
// service shall return the HTTP '428 Precondition Required' status code to
// reject the request.
// zoneETag - This parameter shall contain the current ETag of the zone. If the
// client-provided ETag does not match the current ETag of the zone, the
// service shall return the HTTP '428 Precondition Required' status code to
// reject the request.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (z *Zone) RemoveEndpoint(endpoint string, endpointETag string, zoneETag string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Endpoint"] = endpoint
	payload["EndpointETag"] = endpointETag
	payload["ZoneETag"] = zoneETag
	resp, taskInfo, err := PostWithTask(z.client,
		z.removeEndpointTarget, payload, z.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AddressPools gets the AddressPools linked resources.
func (z *Zone) AddressPools() ([]*AddressPool, error) {
	return GetObjects[AddressPool](z.client, z.addressPools)
}

// ContainedByZones gets the ContainedByZones linked resources.
func (z *Zone) ContainedByZones() ([]*Zone, error) {
	return GetObjects[Zone](z.client, z.containedByZones)
}

// ContainsZones gets the ContainsZones linked resources.
func (z *Zone) ContainsZones() ([]*Zone, error) {
	return GetObjects[Zone](z.client, z.containsZones)
}

// Endpoints gets the Endpoints linked resources.
func (z *Zone) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](z.client, z.endpoints)
}

// InvolvedSwitches gets the InvolvedSwitches linked resources.
func (z *Zone) InvolvedSwitches() ([]*Switch, error) {
	return GetObjects[Switch](z.client, z.involvedSwitches)
}

// ResourceBlocks gets the ResourceBlocks linked resources.
func (z *Zone) ResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](z.client, z.resourceBlocks)
}
