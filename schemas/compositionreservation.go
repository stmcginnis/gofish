//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #CompositionReservation.v1_0_2.CompositionReservation

package schemas

import (
	"encoding/json"
)

// CompositionReservation This resource represents the composition reservation
// of the composition service for a Redfish implementation.
type CompositionReservation struct {
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
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
	// ReservedResourceBlocksCount
	ReservedResourceBlocksCount int `json:"ReservedResourceBlocks@odata.count"`
}

// UnmarshalJSON unmarshals a CompositionReservation object from the raw JSON.
func (c *CompositionReservation) UnmarshalJSON(b []byte) error {
	type temp CompositionReservation
	var tmp struct {
		temp
		ReservedResourceBlocks Links `json:"ReservedResourceBlocks"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CompositionReservation(tmp.temp)

	// Extract the links to other entities for later
	c.reservedResourceBlocks = tmp.ReservedResourceBlocks.ToStrings()

	return nil
}

// GetCompositionReservation will get a CompositionReservation instance from the service.
func GetCompositionReservation(c Client, uri string) (*CompositionReservation, error) {
	return GetObject[CompositionReservation](c, uri)
}

// ListReferencedCompositionReservations gets the collection of CompositionReservation from
// a provided reference.
func ListReferencedCompositionReservations(c Client, link string) ([]*CompositionReservation, error) {
	return GetCollectionObjects[CompositionReservation](c, link)
}

// ReservedResourceBlocks gets the ReservedResourceBlocks linked resources.
func (c *CompositionReservation) ReservedResourceBlocks() ([]*ResourceBlock, error) {
	return GetObjects[ResourceBlock](c.client, c.reservedResourceBlocks)
}
