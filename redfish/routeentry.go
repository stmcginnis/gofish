//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// RouteEntry shall represent the content of route entry rows in the Redfish Specification.
type RouteEntry struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// MinimumHopCount shall indicate the minimum hop count used to calculate the computed hop count.
	MinimumHopCount int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawEntryHex shall contain a binary data that represents the content of route entry rows.
	RawEntryHex string
	// RouteSet shall contain a link to a Resource Collection of type RouteSetEntryCollection.
	routeSet string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a RouteEntry object from the raw JSON.
func (routeentry *RouteEntry) UnmarshalJSON(b []byte) error {
	type temp RouteEntry
	var t struct {
		temp
		RouteSet common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*routeentry = RouteEntry(t.temp)

	// Extract the links to other entities for later
	routeentry.routeSet = t.RouteSet.String()

	// This is a read/write object, so we need to save the raw object data for later
	routeentry.rawData = b

	return nil
}

// RouteSet gets the associated route set.
func (routeentry *RouteEntry) RouteSet(queryOpts ...common.QueryGroupOption) ([]*RouteSetEntry, error) {
	return routeentry.RouteSetWithContext(common.ContextOf(routeentry.GetClient()), queryOpts...)
}

// RouteSetWithContext gets the associated route set.
func (routeentry *RouteEntry) RouteSetWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*RouteSetEntry, error) {
	return ListReferencedRouteSetEntrysWithContext(ctx, routeentry.GetClient(), routeentry.routeSet, queryOpts...)
}

// Update commits updates to this object's properties to the running system.
func (routeentry *RouteEntry) Update() error {
	return routeentry.UpdateWithContext(common.ContextOf(routeentry.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (routeentry *RouteEntry) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{"MinimumHopCount",
		"RawEntryHex"}

	return routeentry.UpdateFromRawDataWithContext(ctx, routeentry, routeentry.rawData, readWriteFields)
}

// GetRouteEntry will get a RouteEntry instance from the service.
func GetRouteEntry(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*RouteEntry, error) {
	return GetRouteEntryWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetRouteEntryWithContext will get a RouteEntry instance from the service.
func GetRouteEntryWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*RouteEntry, error) {
	return common.GetObjectWithContext[RouteEntry](ctx, c, uri, queryOpts...)
}

// ListReferencedRouteEntrys gets the collection of RouteEntry from
// a provided reference.
func ListReferencedRouteEntrys(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*RouteEntry, error) {
	return ListReferencedRouteEntrysWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedRouteEntrysWithContext gets the collection of RouteEntry from
// a provided reference.
func ListReferencedRouteEntrysWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*RouteEntry, error) {
	return common.GetCollectionObjectsWithContext[RouteEntry](ctx, c, link, queryOpts...)
}
