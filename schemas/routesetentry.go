//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2019.4 - #RouteSetEntry.v1_0_2.RouteSetEntry

package schemas

import (
	"encoding/json"
)

// RouteSetEntry shall represent the content of a route set in the Redfish
// Specification.
type RouteSetEntry struct {
	Entity
	// EgressIdentifier shall contain the interface identifier corresponding to
	// this route.
	EgressIdentifier int
	// HopCount shall contain the number of hops to the destination component from
	// the indicated egress interface.
	HopCount int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// VCAction shall contain the index to the VCAT entry corresponding to this
	// route.
	VCAction int
	// Valid shall indicate whether the entry is valid.
	Valid bool
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a RouteSetEntry object from the raw JSON.
func (r *RouteSetEntry) UnmarshalJSON(b []byte) error {
	type temp RouteSetEntry
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = RouteSetEntry(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *RouteSetEntry) Update() error {
	readWriteFields := []string{
		"EgressIdentifier",
		"HopCount",
		"VCAction",
		"Valid",
	}

	return r.UpdateFromRawData(r, r.RawData, readWriteFields)
}

// GetRouteSetEntry will get a RouteSetEntry instance from the service.
func GetRouteSetEntry(c Client, uri string) (*RouteSetEntry, error) {
	return GetObject[RouteSetEntry](c, uri)
}

// ListReferencedRouteSetEntrys gets the collection of RouteSetEntry from
// a provided reference.
func ListReferencedRouteSetEntrys(c Client, link string) ([]*RouteSetEntry, error) {
	return GetCollectionObjects[RouteSetEntry](c, link)
}
