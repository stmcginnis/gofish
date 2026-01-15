//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #CompositionReservation.v1_0_2.CompositionReservation

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CompositionReservation This resource represents the composition reservation
// of the composition service for a Redfish implementation.
type CompositionReservation struct {
	common.Entity
	// Client shall contain the client that owns the reservation. The service shall
	// determine this value based on the client that invoked the 'Compose' action
	// that resulted in the creation of this reservation.
	Client string
	// Manifest shall contain the manifest document processed by the service that
	// resulted in this reservation. This property shall be required if the
	// 'RequestFormat' parameter in the 'Compose' action request contained the
	// value 'Manifest'.
	Manifest Manifest
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReservationTime shall indicate the date and time when the reservation was
	// created by the service.
	ReservationTime string
	// ReservedResourceBlocks shall contain an array of links to resources of type
	// 'ResourceBlock' that represent the reserved resource blocks for this
	// reservation. Upon deletion of the reservation or when the reservation is
	// applied, the 'Reserved' property in the referenced resource blocks shall
	// change to 'false'.
	reservedResourceBlocks []string
	// ReservedResourceBlocks@odata.count
	ReservedResourceBlocksCount int `json:"ReservedResourceBlocks@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a CompositionReservation object from the raw JSON.
func (c *CompositionReservation) UnmarshalJSON(b []byte) error {
	type temp CompositionReservation
	var tmp struct {
		temp
		ReservedResourceBlocks common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CompositionReservation(tmp.temp)

	// Extract the links to other entities for later
	c.reservedResourceBlocks = tmp.ReservedResourceBlocks.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CompositionReservation) Update() error {
	readWriteFields := []string{
		"Manifest",
		"ReservedResourceBlocks@odata.count",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
}

// GetCompositionReservation will get a CompositionReservation instance from the service.
func GetCompositionReservation(c common.Client, uri string) (*CompositionReservation, error) {
	return common.GetObject[CompositionReservation](c, uri)
}

// ListReferencedCompositionReservations gets the collection of CompositionReservation from
// a provided reference.
func ListReferencedCompositionReservations(c common.Client, link string) ([]*CompositionReservation, error) {
	return common.GetCollectionObjects[CompositionReservation](c, link)
}

// ReservedResourceBlocks gets reserved resource blocks for this reservation.
// Upon deletion of the reservation or when the reservation is applied, the
// Reserved property in the referenced resource blocks shall change to 'false'.
func (c *CompositionReservation) ReservedResourceBlocks() ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](
		c.GetClient(),
		c.reservedResourceBlocks)
}
