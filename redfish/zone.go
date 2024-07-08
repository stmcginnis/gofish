//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ExternalAccessibility string

const (
	// GloballyAccessibleExternalAccessibility shall indicate that any external entity with the correct access details,
	// which may include authorization information, can access the endpoints that this zone lists, regardless of zone.
	GloballyAccessibleExternalAccessibility ExternalAccessibility = "GloballyAccessible"
	// NonZonedAccessibleExternalAccessibility shall indicate that any external entity that another zone does not
	// explicitly list can access the endpoints that this zone lists.
	NonZonedAccessibleExternalAccessibility ExternalAccessibility = "NonZonedAccessible"
	// ZoneOnlyExternalAccessibility shall indicate that endpoints in this zone are only accessible by endpoints that
	// this zone explicitly lists.
	ZoneOnlyExternalAccessibility ExternalAccessibility = "ZoneOnly"
	// NoInternalRoutingExternalAccessibility shall indicate that implicit routing within this zone is not defined.
	NoInternalRoutingExternalAccessibility ExternalAccessibility = "NoInternalRouting"
)

type ZoneType string

const (
	// DefaultZoneType shall indicate a zone in which all endpoints are added by default when instantiated. This value
	// shall only be used for zones subordinate to the fabric collection.
	DefaultZoneType ZoneType = "Default"
	// ZoneOfEndpointsZoneType shall indicate a zone that contains resources of type Endpoint. This value shall only be
	// used for zones subordinate to the fabric collection.
	ZoneOfEndpointsZoneType ZoneType = "ZoneOfEndpoints"
	// ZoneOfZonesZoneType shall indicate a zone that contains resources of type Zone. This value shall only be used
	// for zones subordinate to the fabric collection.
	ZoneOfZonesZoneType ZoneType = "ZoneOfZones"
	// ZoneOfResourceBlocksZoneType shall indicate a zone that contains resources of type ResourceBlock. This value
	// shall only be used for zones subordinate to the composition service.
	ZoneOfResourceBlocksZoneType ZoneType = "ZoneOfResourceBlocks"
)

// ZoneLinks shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type zoneLinks struct {
	// AddressPools shall contain an array of links to resources of type AddressPool with which this zone is
	// associated.
	AddressPools common.Links
	// AddressPools@odata.count
	AddressPoolsCount int `json:"AddressPools@odata.count"`
	// ContainedByZones shall contain an array of links to resources of type Zone that represent the zones that contain
	// this zone. The zones referenced by this property shall not be contained by other zones.
	ContainedByZones common.Links
	// ContainedByZones@odata.count
	ContainedByZonesCount int `json:"ContainedByZones@odata.count"`
	// ContainsZones shall contain an array of links to resources of type Zone that represent the zones that are
	// contained by this zone. The zones referenced by this property shall not contain other zones.
	ContainsZones common.Links
	// ContainsZones@odata.count
	ContainsZonesCount int `json:"ContainsZones@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint that this zone contains.
	Endpoints common.Links
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// InvolvedSwitches shall contain an array of links to resources of type Switch in this zone.
	InvolvedSwitches common.Links
	// InvolvedSwitches@odata.count
	InvolvedSwitchesCount int `json:"InvolvedSwitches@odata.count"`
	// ResourceBlocks shall contain an array of links to resources of type ResourceBlock with which this zone is
	// associated.
	ResourceBlocks common.Links
	// ResourceBlocks@odata.count
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
}

// Zone shall represent a simple fabric zone for a Redfish implementation.
type Zone struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// DefaultRoutingEnabled shall indicate whether routing within this zone is enabled.
	DefaultRoutingEnabled bool
	// Description provides a description of this resource.
	Description string
	// ExternalAccessibility shall contain and indication of accessibility of endpoints in this zone to endpoints
	// outside of this zone.
	ExternalAccessibility ExternalAccessibility
	// Identifiers shall contain a list of all known durable names for the associated zone.
	Identifiers []common.Identifier
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ZoneType shall contain the type of zone that this zone represents.
	ZoneType ZoneType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	addEndpointTarget    string
	removeEndpointTarget string

	addressPools []string
	// AddressPoolsCount is the number of address pools associated with this zone.
	AddressPoolsCount int
	containedByZones  []string
	// ContainedByZonesCount is the number of zones that contain this zone.
	ContainedByZonesCount int
	containsZones         []string
	// ContainsZonesCount is the number of zones contained in this zone.
	ContainsZonesCount int
	endpoints          []string
	// EndpointsCount is the number of endpoints that this zone contains.
	EndpointsCount   int
	involvedSwitches []string
	// InvolvedSwitchesCount is the number of switches in this zone.
	InvolvedSwitchesCount int
	resourceBlocks        []string
	// ResourceBlocksCount is the number of resource blocks with which this zone is associated.
	ResourceBlockCount int
}

// UnmarshalJSON unmarshals a Zone object from the raw JSON.
func (zone *Zone) UnmarshalJSON(b []byte) error {
	type temp Zone
	var t struct {
		temp
		Actions struct {
			AddEndpoint    common.ActionTarget `json:"#Zone.AddEndpoint"`
			RemoveEndpoint common.ActionTarget `json:"#Zone.RemoveEndpoint"`
		}
		Links zoneLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*zone = Zone(t.temp)

	// Extract the links to other entities for later
	zone.addEndpointTarget = t.Actions.AddEndpoint.Target
	zone.removeEndpointTarget = t.Actions.RemoveEndpoint.Target

	zone.addressPools = t.Links.AddressPools.ToStrings()
	zone.AddressPoolsCount = t.Links.AddressPoolsCount
	zone.containedByZones = t.Links.ContainedByZones.ToStrings()
	zone.ContainedByZonesCount = t.Links.ContainedByZonesCount
	zone.containsZones = t.Links.ContainsZones.ToStrings()
	zone.ContainsZonesCount = t.Links.ContainsZonesCount
	zone.endpoints = t.Links.Endpoints.ToStrings()
	zone.EndpointsCount = t.Links.EndpointsCount
	zone.involvedSwitches = t.Links.InvolvedSwitches.ToStrings()
	zone.InvolvedSwitchesCount = t.Links.InvolvedSwitchesCount
	zone.resourceBlocks = t.Links.ResourceBlocks.ToStrings()
	zone.ResourceBlockCount = t.Links.ResourceBlocksCount

	// This is a read/write object, so we need to save the raw object data for later
	zone.rawData = b

	return nil
}

// AddEndpoint adds an endpoint to a zone.
//
// `endpointURI` is the URI for the endpoint to add to the zone.
//
// `endpointETag` is the current ETag of the endpoint to add to the zone.
//
// `zoneETag` is the current ETag of the zone. If the client-provided ETag does not
// match the current ETag of the zone, the service shall return the HTTP 428
// (Precondition Required) status code to reject the request.
func (zone *Zone) AddEndpoint(endpointURI, endpointETag, zoneETag string) error {
	if zone.addEndpointTarget == "" {
		return errors.New("addEndpoint not supported by this zone")
	}

	t := struct {
		Endpoint     string
		EndpointETag string
		ZoneETag     string
	}{
		Endpoint:     endpointURI,
		EndpointETag: endpointETag,
		ZoneETag:     zoneETag,
	}
	return zone.Post(zone.addEndpointTarget, t)
}

// RemoveEndpoint removes an endpoint from a zone.
//
// `endpointURI` is the URI for the endpoint to remove from the zone.
//
// `endpointETag` is the current ETag of the endpoint to remove from the zone.
//
// `zoneETag` is the current ETag of the zone. If the client-provided ETag does not
// match the current ETag of the zone, the service shall return the HTTP 428
// (Precondition Required) status code to reject the request.
func (zone *Zone) RemoveEndpoint(endpointURI, endpointETag, zoneETag string) error {
	if zone.removeEndpointTarget == "" {
		return errors.New("removeEndpoint not supported by this zone")
	}

	t := struct {
		Endpoint     string
		EndpointETag string
		ZoneETag     string
	}{
		Endpoint:     endpointURI,
		EndpointETag: endpointETag,
		ZoneETag:     zoneETag,
	}
	return zone.Post(zone.removeEndpointTarget, t)
}

// AddressPools gets the address pools associated with this zone.
func (zone *Zone) AddressPools() ([]*AddressPool, error) {
	return common.GetObjects[AddressPool](zone.GetClient(), zone.addressPools)
}

// ContainedByZones gets the zone that contain this zone.
func (zone *Zone) ContainedByZones() ([]*Zone, error) {
	return common.GetObjects[Zone](zone.GetClient(), zone.containedByZones)
}

// ContainsZones gets the zones that are contained by this zone.
func (zone *Zone) ContainsZones() ([]*Zone, error) {
	return common.GetObjects[Zone](zone.GetClient(), zone.containsZones)
}

// Endpoints gets the endpoints that this zone contains.
func (zone *Zone) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](zone.GetClient(), zone.endpoints)
}

// InvolvedSwitches gets the switches in this zone.
func (zone *Zone) InvolvedSwitches() ([]*Switch, error) {
	return common.GetObjects[Switch](zone.GetClient(), zone.involvedSwitches)
}

// ResourceBlocks gets the resource blocks with which this zone is associated.
func (zone *Zone) ResourceBlocks() ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](zone.GetClient(), zone.resourceBlocks)
}

// Update commits updates to this object's properties to the running system.
func (zone *Zone) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Zone)
	original.UnmarshalJSON(zone.rawData)

	readWriteFields := []string{
		"DefaultRoutingEnabled",
		"ExternalAccessibility",
		"ZoneType",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(zone).Elem()

	return zone.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetZone will get a Zone instance from the service.
func GetZone(c common.Client, uri string) (*Zone, error) {
	return common.GetObject[Zone](c, uri)
}

// ListReferencedZones gets the collection of Zone from
// a provided reference.
func ListReferencedZones(c common.Client, link string) ([]*Zone, error) {
	return common.GetCollectionObjects[Zone](c, link)
}
