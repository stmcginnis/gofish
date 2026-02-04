//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/CompositionService.v1_2_3.json
// 2021.1 - #CompositionService.v1_2_3.CompositionService

package schemas

import (
	"encoding/json"
)

type ComposeRequestFormat string

const (
	// ManifestComposeRequestFormat shall indicate that the request contains a
	// manifest as defined by the Redfish 'Manifest' schema.
	ManifestComposeRequestFormat ComposeRequestFormat = "Manifest"
)

type ComposeRequestType string

const (
	// PreviewComposeRequestType shall indicate that the request is to preview the
	// outcome of the operations specified by the manifest to show what the service
	// will do based on the contents of the request, and not affect any resources
	// within the service.
	PreviewComposeRequestType ComposeRequestType = "Preview"
	// PreviewReserveComposeRequestType shall indicate that the request is to
	// preview the outcome of the operations specified by the manifest to show what
	// the service will do based on the contents of the request. Resources that
	// would have been affected by this request shall be marked as reserved but
	// otherwise shall not be affected.
	PreviewReserveComposeRequestType ComposeRequestType = "PreviewReserve"
	// ApplyComposeRequestType shall indicate that the request is to apply the
	// requested operations specified by the manifest and modify resources as
	// needed.
	ApplyComposeRequestType ComposeRequestType = "Apply"
)

// CompositionService shall represent the composition service and its properties
// for a Redfish implementation.
type CompositionService struct {
	Entity
	// ActivePool shall contain a link to a resource collection of type
	// 'ResourceBlockCollection'. The members of this collection shall represent
	// the resource blocks in the active pool. Services shall filter members of
	// this collection based on the requesting client.
	//
	// Version added: v1.2.0
	activePool string
	// AllowOverprovisioning shall indicate whether this service is allowed to
	// overprovision a composition relative to the composition request.
	//
	// Version added: v1.1.0
	AllowOverprovisioning bool
	// AllowZoneAffinity shall indicate whether a client can request that a
	// specific resource zone fulfill a composition request.
	//
	// Version added: v1.1.0
	AllowZoneAffinity bool
	// CompositionReservations shall contain a link to a resource collection of
	// type 'CompositionReservationCollection'. The members of this collection
	// shall contain links to reserved resource blocks and the related document
	// that caused the reservations. Services shall filter members of this
	// collection based on the requesting client.
	//
	// Version added: v1.2.0
	compositionReservations string
	// FreePool shall contain a link to a resource collection of type
	// 'ResourceBlockCollection'. The members of this collection shall represent
	// the resource blocks in the free pool. Services shall filter members of this
	// collection based on the requesting client.
	//
	// Version added: v1.2.0
	freePool string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReservationDuration shall contain the length of time a composition
	// reservation is held before the service deletes the reservation and marks any
	// related resource blocks as no longer reserved.
	//
	// Version added: v1.2.0
	ReservationDuration string
	// ResourceBlocks shall contain a link to a resource collection of type
	// 'ResourceBlockCollection'.
	resourceBlocks string
	// ResourceZones shall contain a link to a resource collection of type
	// 'ZoneCollection'.
	resourceZones string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// composeTarget is the URL to send Compose requests.
	composeTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a CompositionService object from the raw JSON.
func (c *CompositionService) UnmarshalJSON(b []byte) error {
	type temp CompositionService
	type cActions struct {
		Compose ActionTarget `json:"#CompositionService.Compose"`
	}
	var tmp struct {
		temp
		Actions                 cActions
		ActivePool              Link `json:"ActivePool"`
		CompositionReservations Link `json:"CompositionReservations"`
		FreePool                Link `json:"FreePool"`
		ResourceBlocks          Link `json:"ResourceBlocks"`
		ResourceZones           Link `json:"ResourceZones"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CompositionService(tmp.temp)

	// Extract the links to other entities for later
	c.composeTarget = tmp.Actions.Compose.Target
	c.activePool = tmp.ActivePool.String()
	c.compositionReservations = tmp.CompositionReservations.String()
	c.freePool = tmp.FreePool.String()
	c.resourceBlocks = tmp.ResourceBlocks.String()
	c.resourceZones = tmp.ResourceZones.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CompositionService) Update() error {
	readWriteFields := []string{
		"AllowOverprovisioning",
		"ReservationDuration",
		"ServiceEnabled",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCompositionService will get a CompositionService instance from the service.
func GetCompositionService(c Client, uri string) (*CompositionService, error) {
	return GetObject[CompositionService](c, uri)
}

// ListReferencedCompositionServices gets the collection of CompositionService from
// a provided reference.
func ListReferencedCompositionServices(c Client, link string) ([]*CompositionService, error) {
	return GetCollectionObjects[CompositionService](c, link)
}

// CompositionServiceComposeParameters holds the parameters for the Compose action.
type CompositionServiceComposeParameters struct {
	// Manifest shall contain the manifest containing the compose operation
	// request. This parameter shall be required if 'RequestFormat' contains the
	// value 'Manifest'.
	Manifest Manifest `json:"Manifest,omitempty"`
	// RequestFormat shall contain the format of the request.
	RequestFormat ComposeRequestFormat `json:"RequestFormat,omitempty"`
	// RequestType shall contain the type of request.
	RequestType ComposeRequestType `json:"RequestType,omitempty"`
	// ReservationID shall contain the value of the 'Id' property of the
	// 'CompositionReservation' resource for applying a reservation.
	ReservationID string `json:"ReservationId,omitempty"`
}

// This action shall perform a set of operations specified by a manifest.
// Services shall not apply any part of the manifest unless all operations
// specified by the manifest are successful.
func (c *CompositionService) Compose(params *CompositionServiceComposeParameters) (*ComposeResponse, error) {
	resp, err := c.PostWithResponse(c.composeTarget, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result ComposeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ActivePool gets the ActivePool collection.
func (c *CompositionService) ActivePool() ([]*ResourceBlock, error) {
	if c.activePool == "" {
		return nil, nil
	}
	return GetCollectionObjects[ResourceBlock](c.client, c.activePool)
}

// CompositionReservations gets the CompositionReservations collection.
func (c *CompositionService) CompositionReservations() ([]*CompositionReservation, error) {
	if c.compositionReservations == "" {
		return nil, nil
	}
	return GetCollectionObjects[CompositionReservation](c.client, c.compositionReservations)
}

// FreePool gets the FreePool collection.
func (c *CompositionService) FreePool() ([]*ResourceBlock, error) {
	if c.freePool == "" {
		return nil, nil
	}
	return GetCollectionObjects[ResourceBlock](c.client, c.freePool)
}

// ResourceBlocks gets the ResourceBlocks collection.
func (c *CompositionService) ResourceBlocks() ([]*ResourceBlock, error) {
	if c.resourceBlocks == "" {
		return nil, nil
	}
	return GetCollectionObjects[ResourceBlock](c.client, c.resourceBlocks)
}

// ResourceZones gets the ResourceZones collection.
func (c *CompositionService) ResourceZones() ([]*Zone, error) {
	if c.resourceZones == "" {
		return nil, nil
	}
	return GetCollectionObjects[Zone](c.client, c.resourceZones)
}

// ComposeResponse shall contain the properties found in the response body for
// the 'Compose' action.
type ComposeResponse struct {
	// Manifest shall contain the manifest containing the compose operation
	// response. This property shall be required if 'RequestFormat' contains the
	// value 'Manifest'.
	//
	// Version added: v1.2.0
	Manifest Manifest
	// RequestFormat shall contain the format of the request.
	//
	// Version added: v1.2.0
	RequestFormat ComposeRequestFormat
	// RequestType shall contain the type of request.
	//
	// Version added: v1.2.0
	RequestType ComposeRequestType
	// ReservationID shall contain the value of the 'Id' property of the
	// 'CompositionReservation' resource that was created. This property shall be
	// required if 'RequestType' contains the value 'PreviewReserve'.
	//
	// Version added: v1.2.0
	ReservationID string `json:"ReservationId"`
}
