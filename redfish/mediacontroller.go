//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.1 - #MediaController.v1_3_2.MediaController

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

// MediaController shall represent a media controller in a Redfish
// implementation.
type MediaController struct {
	common.Entity
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this media
	// controller.
	//
	// Version added: v1.2.0
	environmentMetrics string
	// Manufacturer shall contain the manufacturer of the media controller.
	Manufacturer string
	// MediaControllerType shall contain the type of media controller.
	MediaControllerType MediaControllerType
	// Model shall contain the model of the media controller.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall indicate the part number as provided by the manufacturer of
	// this media controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// SerialNumber shall indicate the serial number as provided by the
	// manufacturer of this media controller.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain a universally unique identifier number for the media
	// controller.
	//
	// Version added: v1.1.0
	UUID string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// memoryDomains are the URIs for MemoryDomains.
	memoryDomains []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MediaController object from the raw JSON.
func (m *MediaController) UnmarshalJSON(b []byte) error {
	type temp MediaController
	type mActions struct {
		Reset common.ActionTarget `json:"#MediaController.Reset"`
	}
	type mLinks struct {
		Endpoints     common.Links `json:"Endpoints"`
		MemoryDomains common.Links `json:"MemoryDomains"`
	}
	var tmp struct {
		temp
		Actions            mActions
		Links              mLinks
		EnvironmentMetrics common.Link `json:"environmentMetrics"`
		Ports              common.Link `json:"ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MediaController(tmp.temp)

	// Extract the links to other entities for later
	m.resetTarget = tmp.Actions.Reset.Target
	m.endpoints = tmp.Links.Endpoints.ToStrings()
	m.memoryDomains = tmp.Links.MemoryDomains.ToStrings()
	m.environmentMetrics = tmp.EnvironmentMetrics.String()
	m.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	m.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MediaController) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return m.UpdateFromRawData(m, m.rawData, readWriteFields)
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

// Reset shall reset this media controller.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
func (m *MediaController) Reset(resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return m.Post(m.resetTarget, payload)
}

// Endpoints gets the Endpoints linked resources.
func (m *MediaController) Endpoints(client common.Client) ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](client, m.endpoints)
}

// MemoryDomains gets the MemoryDomains linked resources.
func (m *MediaController) MemoryDomains(client common.Client) ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](client, m.memoryDomains)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (m *MediaController) EnvironmentMetrics(client common.Client) (*EnvironmentMetrics, error) {
	if m.environmentMetrics == "" {
		return nil, nil
	}
	return common.GetObject[EnvironmentMetrics](client, m.environmentMetrics)
}

// Ports gets the Ports collection.
func (m *MediaController) Ports(client common.Client) ([]*Port, error) {
	if m.ports == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, m.ports)
}
