//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/coreweave/gofish/common"
)

// IOConnectivityLoSCapabilities describes capabilities of the system to
// support various IO Connectivity service options.
type IOConnectivityLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// MaxSupportedBytesPerSecond shall be the maximum bytes per second that a
	// connection can support.
	MaxSupportedBytesPerSecond int64
	// MaxSupportedIOPS shall be the maximum IOPS that a connection can support.
	MaxSupportedIOPS int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupportedAccessProtocols is Access protocols supported by this service
	// option. NOTE: SMB+NFS* requires that SMB and at least one of NFSv3 or
	// NFXv4 are also selected, (i.e. {'SMB', 'NFSv4', 'SMB+NFS*'}).
	SupportedAccessProtocols []common.Protocol
	// SupportedLinesOfService shall contain known and
	// supported IOConnectivityLinesOfService.
	SupportedLinesOfService []IOConnectivityLineOfService
	// SupportedLinesOfServiceCount is the number of IOConnectivityLineOfServices.
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a IOConnectivityLoSCapabilities object from the raw JSON.
func (ioconnectivityloscapabilities *IOConnectivityLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp IOConnectivityLoSCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ioconnectivityloscapabilities = IOConnectivityLoSCapabilities(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	ioconnectivityloscapabilities.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ioconnectivityloscapabilities *IOConnectivityLoSCapabilities) Update() error {
	return ioconnectivityloscapabilities.UpdateWithContext(common.ContextOf(ioconnectivityloscapabilities.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (ioconnectivityloscapabilities *IOConnectivityLoSCapabilities) UpdateWithContext(ctx context.Context) error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(IOConnectivityLoSCapabilities)
	err := original.UnmarshalJSON(ioconnectivityloscapabilities.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"MaxSupportedBytesPerSecond",
		"MaxSupportedIOPS",
		"SupportedAccessProtocols",
		"SupportedLinesOfService",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(ioconnectivityloscapabilities).Elem()

	return ioconnectivityloscapabilities.Entity.UpdateWithContext(ctx, originalElement, currentElement, readWriteFields)
}

// GetIOConnectivityLoSCapabilities will get a IOConnectivityLoSCapabilities
// instance from the service.
func GetIOConnectivityLoSCapabilities(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IOConnectivityLoSCapabilities, error) {
	return GetIOConnectivityLoSCapabilitiesWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetIOConnectivityLoSCapabilitiesWithContext will get a IOConnectivityLoSCapabilities
// instance from the service.
func GetIOConnectivityLoSCapabilitiesWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IOConnectivityLoSCapabilities, error) {
	return common.GetObjectWithContext[IOConnectivityLoSCapabilities](ctx, c, uri, queryOpts...)
}

// ListReferencedIOConnectivityLoSCapabilitiess gets the collection of
// IOConnectivityLoSCapabilities from a provided reference.
func ListReferencedIOConnectivityLoSCapabilitiess(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*IOConnectivityLoSCapabilities, error) {
	return ListReferencedIOConnectivityLoSCapabilitiessWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedIOConnectivityLoSCapabilitiessWithContext gets the collection of
// IOConnectivityLoSCapabilities from a provided reference.
func ListReferencedIOConnectivityLoSCapabilitiessWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*IOConnectivityLoSCapabilities, error) {
	return common.GetCollectionObjectsWithContext[IOConnectivityLoSCapabilities](ctx, c, link, queryOpts...)
}
