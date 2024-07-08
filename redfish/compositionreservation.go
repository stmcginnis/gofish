//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CompositionReservation This resource represents the composition reservation of the composition service for a
// Redfish implementation.
type CompositionReservation struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Client shall contain the client that owns the reservation. The service shall determine this value based on the
	// client that invoked the Compose action that resulted in the creation of this reservation.
	Client string
	// Description provides a description of this resource.
	Description string
	// Manifest shall contain the manifest document processed by the service that resulted in this reservation. This
	// property shall be required if the RequestFormat parameter in the Compose action request contained the value
	// 'Manifest'.
	Manifest Manifest
	// ReservationTime shall indicate the date and time when the reservation was created by the service.
	ReservationTime string
	// ReservedResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// reserved resource blocks for this reservation. Upon deletion of the reservation or when the reservation is
	// applied, the Reserved property in the referenced resource blocks shall change to 'false'.
	reservedResourceBlocks []string
	// ReservedResourceBlocksCount is the number of ResourceBlocks for this reservation.
	ReservedResourceBlocksCount int `json:"ReservedResourceBlocks@odata.count"`
}

// UnmarshalJSON unmarshals a CompositionReservation object from the raw JSON.
func (compositionreservation *CompositionReservation) UnmarshalJSON(b []byte) error {
	type temp CompositionReservation
	var t struct {
		temp
		ReservedResourceBlocks common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*compositionreservation = CompositionReservation(t.temp)

	// Extract the links to other entities for later
	compositionreservation.reservedResourceBlocks = t.ReservedResourceBlocks.ToStrings()

	return nil
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
func (compositionreservation *CompositionReservation) ReservedResourceBlocks() ([]*ResourceBlock, error) {
	return common.GetObjects[ResourceBlock](
		compositionreservation.GetClient(),
		compositionreservation.reservedResourceBlocks)
}
