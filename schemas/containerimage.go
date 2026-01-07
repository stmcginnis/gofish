//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #ContainerImage.v1_0_1.ContainerImage

package schemas

import (
	"encoding/json"
)

type ImageTypes string

const (
	// DockerV1ImageTypes shall indicate a Docker Image Specification Version
	// 1-defined Docker image.
	DockerV1ImageTypes ImageTypes = "DockerV1"
	// DockerV2ImageTypes shall indicate a Docker Image Manifest Version 2-defined
	// Docker image.
	DockerV2ImageTypes ImageTypes = "DockerV2"
	// OCIImageTypes shall indicate an Open Container Specification-defined OCI
	// (Open Container Initiative) image.
	OCIImageTypes ImageTypes = "OCI"
)

// ContainerImage shall represent a container image available to a computer
// system.
type ContainerImage struct {
	Entity
	// CreateTime shall indicate the date and time when the container image was
	// created.
	CreateTime string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProgrammaticID shall contain the programmatic identifier for this container
	// image. This is typically a hash string that represents the identifier of
	// this container image.
	ProgrammaticID string `json:"ProgrammaticId"`
	// SizeBytes shall contain the size of this container image in bytes.
	SizeBytes *int `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Type shall contain the supported image types for this container engine.
	Type ImageTypes
	// Version shall contain the version of this application.
	Version string
	// containers are the URIs for Containers.
	containers []string
	// softwareImage is the URI for SoftwareImage.
	softwareImage string
}

// UnmarshalJSON unmarshals a ContainerImage object from the raw JSON.
func (c *ContainerImage) UnmarshalJSON(b []byte) error {
	type temp ContainerImage
	type cLinks struct {
		Containers    Links `json:"Containers"`
		SoftwareImage Link  `json:"SoftwareImage"`
	}
	var tmp struct {
		temp
		Links cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ContainerImage(tmp.temp)

	// Extract the links to other entities for later
	c.containers = tmp.Links.Containers.ToStrings()
	c.softwareImage = tmp.Links.SoftwareImage.String()

	return nil
}

// GetContainerImage will get a ContainerImage instance from the service.
func GetContainerImage(c Client, uri string) (*ContainerImage, error) {
	return GetObject[ContainerImage](c, uri)
}

// ListReferencedContainerImages gets the collection of ContainerImage from
// a provided reference.
func ListReferencedContainerImages(c Client, link string) ([]*ContainerImage, error) {
	return GetCollectionObjects[ContainerImage](c, link)
}

// Containers gets the Containers linked resources.
func (c *ContainerImage) Containers() ([]*Container, error) {
	return GetObjects[Container](c.client, c.containers)
}

// SoftwareImage gets the SoftwareImage linked resource.
func (c *ContainerImage) SoftwareImage() (*SoftwareInventory, error) {
	if c.softwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](c.client, c.softwareImage)
}
