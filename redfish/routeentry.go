//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2019.4 - #RouteEntry.v1_0_2.RouteEntry

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// RouteEntry shall represent the content of route entry rows in the Redfish
// Specification.
type RouteEntry struct {
	common.Entity
	// MinimumHopCount shall indicate the minimum hop count used to calculate the
	// computed hop count.
	MinimumHopCount int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawEntryHex shall contain a binary data that represents the content of route
	// entry rows.
	RawEntryHex string
	// RouteSet shall contain a link to a resource collection of type
	// 'RouteSetEntryCollection'.
	routeSet string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a RouteEntry object from the raw JSON.
func (r *RouteEntry) UnmarshalJSON(b []byte) error {
	type temp RouteEntry
	var tmp struct {
		temp
		RouteSet common.Link `json:"routeSet"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = RouteEntry(tmp.temp)

	// Extract the links to other entities for later
	r.routeSet = tmp.RouteSet.String()

	// This is a read/write object, so we need to save the raw object data for later
	r.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *RouteEntry) Update() error {
	readWriteFields := []string{
		"MinimumHopCount",
		"RawEntryHex",
	}

	return r.UpdateFromRawData(r, r.rawData, readWriteFields)
}

// GetRouteEntry will get a RouteEntry instance from the service.
func GetRouteEntry(c common.Client, uri string) (*RouteEntry, error) {
	return common.GetObject[RouteEntry](c, uri)
}

// ListReferencedRouteEntrys gets the collection of RouteEntry from
// a provided reference.
func ListReferencedRouteEntrys(c common.Client, link string) ([]*RouteEntry, error) {
	return common.GetCollectionObjects[RouteEntry](c, link)
}

// RouteSet gets the RouteSet collection.
func (r *RouteEntry) RouteSet(client common.Client) ([]*RouteSetEntry, error) {
	if r.routeSet == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[RouteSetEntry](client, r.routeSet)
}
