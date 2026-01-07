//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2019.4 - #RouteEntry.v1_0_2.RouteEntry

package schemas

import (
	"encoding/json"
)

// RouteEntry shall represent the content of route entry rows in the Redfish
// Specification.
type RouteEntry struct {
	Entity
	// MinimumHopCount shall indicate the minimum hop count used to calculate the
	// computed hop count.
	MinimumHopCount int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawEntryHex shall contain a binary data that represents the content of route
	// entry rows.
	RawEntryHex string
	// RouteSet shall contain a link to a resource collection of type
	// 'RouteSetEntryCollection'.
	routeSet string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a RouteEntry object from the raw JSON.
func (r *RouteEntry) UnmarshalJSON(b []byte) error {
	type temp RouteEntry
	var tmp struct {
		temp
		RouteSet Link `json:"RouteSet"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = RouteEntry(tmp.temp)

	// Extract the links to other entities for later
	r.routeSet = tmp.RouteSet.String()

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *RouteEntry) Update() error {
	readWriteFields := []string{
		"MinimumHopCount",
		"RawEntryHex",
	}

	return r.UpdateFromRawData(r, r.RawData, readWriteFields)
}

// GetRouteEntry will get a RouteEntry instance from the service.
func GetRouteEntry(c Client, uri string) (*RouteEntry, error) {
	return GetObject[RouteEntry](c, uri)
}

// ListReferencedRouteEntrys gets the collection of RouteEntry from
// a provided reference.
func ListReferencedRouteEntrys(c Client, link string) ([]*RouteEntry, error) {
	return GetCollectionObjects[RouteEntry](c, link)
}

// RouteSet gets the RouteSet collection.
func (r *RouteEntry) RouteSet() ([]*RouteSetEntry, error) {
	if r.routeSet == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteSetEntry](r.client, r.routeSet)
}
