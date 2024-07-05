//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ComposeRequestFormat string

const (
	// ManifestComposeRequestFormat shall indicate that the request contains a manifest as defined by the Redfish
	// Manifest schema.
	ManifestComposeRequestFormat ComposeRequestFormat = "Manifest"
)

type ComposeRequestType string

const (
	// PreviewComposeRequestType shall indicate that the request is to preview the outcome of the operations specified
	// by the manifest to show what the service will do based on the contents of the request, and not affect any
	// resources within the service.
	PreviewComposeRequestType ComposeRequestType = "Preview"
	// PreviewReserveComposeRequestType shall indicate that the request is to preview the outcome of the operations
	// specified by the manifest to show what the service will do based on the contents of the request. Resources that
	// would have been affected by this request shall be marked as reserved but otherwise shall not be affected.
	PreviewReserveComposeRequestType ComposeRequestType = "PreviewReserve"
	// ApplyComposeRequestType shall indicate that the request is to apply the requested operations specified by the
	// manifest and modify resources as needed.
	ApplyComposeRequestType ComposeRequestType = "Apply"
)

// ComposeRequest contains the parameters for the `Compose` action.
type ComposeRequest struct {
	// Manifest is the manifest containing the compose operation request. This parameter shall be required if
	// RequestFormat contains the value `Manifest`.
	Manifest Manifest `json:",omitempty"`
	// RequestFormat is the format of the request.
	RequestFormat ComposeRequestFormat
	// RequestFormat is the type of request.
	RequestType ComposeRequestType
	// ReservationID is the identifier of the composition reservation if applying a reservation. The value for this
	// parameter is obtained from the response of a Compose action where the RequestType parameter contains the value
	// `PreviewReserve`.
	ReservationID string `json:"ReservationID,omitempty"`
}

// ComposeResponse shall contain the properties found in the response body for the Compose action.
type ComposeResponse struct {
	// Manifest shall contain the manifest containing the compose operation response. This property shall be required
	// if RequestFormat contains the value 'Manifest'.
	Manifest Manifest
	// RequestFormat shall contain the format of the request.
	RequestFormat ComposeRequestFormat
	// RequestType shall contain the type of request.
	RequestType ComposeRequestType
	// ReservationID shall contain the value of the ID property of the CompositionReservation resource that was
	// created. This property shall be required if RequestType contains the value 'PreviewReserve'.
	ReservationID string `json:"ReservationId"`
}

// CompositionService is used to represent the Composition Service
// Properties for a Redfish implementation.
type CompositionService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ActivePool shall contain a link to a resource collection of type ResourceBlockCollection. The members of this
	// collection shall represent the resource blocks in the active pool.
	activePool string
	// AllowOverprovisioning shall be a boolean indicating whether this service
	// is allowed to overprovision a composition relative to the composition request.
	AllowOverprovisioning bool
	// AllowZoneAffinity shall be a boolean indicating whether a client is
	// allowed to request that given composition request is fulfilled by a
	// specified Resource Zone.
	AllowZoneAffinity bool
	// CompositionReservations shall contain a link to a resource collection of type CompositionReservationCollection.
	// The members of this collection shall contain links to reserved resource blocks and the related document that
	// caused the reservations.
	compositionReservations string
	// Description provides a description of this resource.
	Description string
	// FreePool shall contain a link to a resource collection of type ResourceBlockCollection. The members of this
	// collection shall represent the resource blocks in the free pool. Services shall filter members of this
	// collection based on the requesting client.
	freePool string
	// ReservationDuration shall contain the length of time a composition reservation is held before the service
	// deletes the reservation and marks any related resource blocks as no longer reserved.
	ReservationDuration string
	// resourceBlocks shall contain the link to a collection of type ResourceBlockCollection.
	resourceBlocks string
	// resourceZones shall contain the link to a collection of type ZoneCollection.
	resourceZones string
	// ServiceEnabled shall be a boolean indicating whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	composeTarget string
}

// UnmarshalJSON unmarshals CompositionService object from the raw JSON.
func (compositionservice *CompositionService) UnmarshalJSON(b []byte) error {
	type temp CompositionService
	type Actions struct {
		Compose struct {
			Target string `json:"target"`
		} `json:"#CompositionService.Compose"`
	}
	var t struct {
		temp
		Actions                 Actions
		ActivePool              common.Link
		CompositionReservations common.Link
		FreePool                common.Link
		ResourceBlocks          common.Link
		ResourceZones           common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*compositionservice = CompositionService(t.temp)
	compositionservice.composeTarget = t.Actions.Compose.Target
	compositionservice.activePool = t.ActivePool.String()
	compositionservice.compositionReservations = t.CompositionReservations.String()
	compositionservice.freePool = t.FreePool.String()
	compositionservice.resourceBlocks = t.ResourceBlocks.String()
	compositionservice.resourceZones = t.ResourceZones.String()

	// This is a read/write object, so we need to save the raw object data for later
	compositionservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (compositionservice *CompositionService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(CompositionService)
	err := original.UnmarshalJSON(compositionservice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AllowOverprovisioning",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(compositionservice).Elem()

	return compositionservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCompositionService will get a CompositionService instance from the service.
func GetCompositionService(c common.Client, uri string) (*CompositionService, error) {
	return common.GetObject[CompositionService](c, uri)
}

// ListReferencedCompositionServices gets the collection of CompositionService from
// a provided reference.
func ListReferencedCompositionServices(c common.Client, link string) ([]*CompositionService, error) {
	return common.GetCollectionObjects[CompositionService](c, link)
}

// Compose performs a set of operations specified by a manifest.
func (compositionservice *CompositionService) Compose(request *ComposeRequest) (*ComposeResponse, error) {
	resp, err := compositionservice.PostWithResponse(compositionservice.composeTarget, request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ComposeResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ActivePool gets a resource collection whose members represent the resource blocks in the active pool.
func (compositionservice *CompositionService) ActivePool() ([]*ResourceBlock, error) {
	if compositionservice.activePool == "" {
		return nil, nil
	}

	return ListReferencedResourceBlocks(compositionservice.GetClient(), compositionservice.activePool)
}

// CompositionReservations gets a resource collection whose members represent the reserved resource blocks and the
// related document that caused the reservations.
func (compositionservice *CompositionService) CompositionReservations() ([]*ResourceBlock, error) {
	if compositionservice.compositionReservations == "" {
		return nil, nil
	}

	return ListReferencedResourceBlocks(compositionservice.GetClient(), compositionservice.compositionReservations)
}

// FreePool gets a resource collection whose members represent the reserved resource blocks in the free pool.
func (compositionservice *CompositionService) FreePool() ([]*ResourceBlock, error) {
	if compositionservice.freePool == "" {
		return nil, nil
	}

	return ListReferencedResourceBlocks(compositionservice.GetClient(), compositionservice.freePool)
}
