//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// RouteEntry shall represent the content of route entry rows in the Redfish Specification.
type RouteEntry struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
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
func (routeentry *RouteEntry) RouteSet() ([]*RouteSetEntry, error) {
	return ListReferencedRouteSetEntrys(routeentry.GetClient(), routeentry.routeSet)
}

// Update commits updates to this object's properties to the running system.
func (routeentry *RouteEntry) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(RouteEntry)
	original.UnmarshalJSON(routeentry.rawData)

	readWriteFields := []string{
		"MinimumHopCount",
		"RawEntryHex",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(routeentry).Elem()

	return routeentry.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRouteEntry will get a RouteEntry instance from the service.
func GetRouteEntry(c common.Client, uri string) (*RouteEntry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var routeentry RouteEntry
	err = json.NewDecoder(resp.Body).Decode(&routeentry)
	if err != nil {
		return nil, err
	}

	routeentry.SetClient(c)
	return &routeentry, nil
}

// ListReferencedRouteEntrys gets the collection of RouteEntry from
// a provided reference.
func ListReferencedRouteEntrys(c common.Client, link string) ([]*RouteEntry, error) {
	var result []*RouteEntry
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *RouteEntry
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		routeentry, err := GetRouteEntry(c, link)
		ch <- GetResult{Item: routeentry, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
