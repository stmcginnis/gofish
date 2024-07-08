//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type MediaControllerType string

const (
	// MemoryMediaControllerType shall indicate the media controller is for memory.
	MemoryMediaControllerType MediaControllerType = "Memory"
)

// MediaController This resource contains the media controller in a Redfish implementation.
type MediaController struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this media controller.
	environmentMetrics string
	// Manufacturer shall contain the manufacturer of the media controller.
	Manufacturer string
	// MediaControllerType shall contain the type of media controller.
	MediaControllerType MediaControllerType
	// Model shall contain the model of the media controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall indicate the part number as provided by the manufacturer of this media controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// SerialNumber shall indicate the serial number as provided by the manufacturer of this media controller.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain a universally unique identifier number for the media controller.
	UUID string

	endpoints []string
	// EndpointsCount is the number of Endpoints with which this media controller is associated.
	EndpointsCount int
	memoryDomains  []string
	// MemoryDomainsCount get the number of memory domains associated with this memory controller.
	MemoryDomainsCount int

	resetTarget string
}

// UnmarshalJSON unmarshals a MediaController object from the raw JSON.
func (mediacontroller *MediaController) UnmarshalJSON(b []byte) error {
	type temp MediaController
	type Actions struct {
		Reset common.ActionTarget `json:"#MediaController.Reset"`
	}
	type Links struct {
		// Endpoints shall contain an array of links to resources of type Endpoint with which this media controller is
		// associated.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
		// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
		// domains associated with this memory controller.
		MemoryDomains common.Links
		// MemoryDomains@odata.count
		MemoryDomainsCount int `json:"MemoryDomains@odata.count"`
	}
	var t struct {
		temp
		Actions            Actions
		Links              Links
		EnvironmentMetrics common.Link
		Ports              common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*mediacontroller = MediaController(t.temp)

	// Extract the links to other entities for later
	mediacontroller.resetTarget = t.Actions.Reset.Target

	mediacontroller.endpoints = t.Links.Endpoints.ToStrings()
	mediacontroller.EndpointsCount = t.Links.EndpointsCount
	mediacontroller.memoryDomains = t.Links.MemoryDomains.ToStrings()
	mediacontroller.MemoryDomainsCount = t.Links.MemoryDomainsCount

	mediacontroller.environmentMetrics = t.EnvironmentMetrics.String()
	mediacontroller.ports = t.Ports.String()

	return nil
}

// Reset resets this media controller.
func (mediacontroller *MediaController) Reset(resetType ResetType) error {
	parameters := struct {
		ResetType ResetType
	}{
		ResetType: resetType,
	}
	return mediacontroller.Post(mediacontroller.resetTarget, parameters)
}

// EnvironmentMetrics gets the environment metrics for this media controller.
func (mediacontroller *MediaController) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if mediacontroller.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(mediacontroller.GetClient(), mediacontroller.environmentMetrics)
}

// Ports gets the ports associated with this media controller.
func (mediacontroller *MediaController) Ports() ([]*Port, error) {
	return ListReferencedPorts(mediacontroller.GetClient(), mediacontroller.ports)
}

// Endpoints get the Endpoints with which this media controller is associated.
func (mediacontroller *MediaController) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](mediacontroller.GetClient(), mediacontroller.endpoints)
}

// MemoryDomains get the memory domains associated with this memory controller.
func (mediacontroller *MediaController) MemoryDomains() ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](mediacontroller.GetClient(), mediacontroller.memoryDomains)
}

// GetMediaController will get a MediaController instance from the service.
func GetMediaController(c common.Client, uri string) (*MediaController, error) {
	return common.GetObject[MediaController](c, uri)
}

// ListReferencedMediaControllers gets the collection of MediaController from
// a provided reference.
func ListReferencedMediaControllers(c common.Client, link string) ([]*MediaController, error) {
	return common.GetCollectionObjects[MediaController](c, link)
}
