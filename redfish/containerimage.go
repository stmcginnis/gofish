//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type ImageTypes string

const (
	// DockerV1ImageTypes shall indicate a Docker Image Specification Version 1-defined Docker image.
	DockerV1ImageTypes ImageTypes = "DockerV1"
	// DockerV2ImageTypes shall indicate a Docker Image Manifest Version 2-defined Docker image.
	DockerV2ImageTypes ImageTypes = "DockerV2"
	// OCIImageTypes shall indicate an Open Container Specification-defined OCI (Open Container Initiative) image.
	OCIImageTypes ImageTypes = "OCI"
)

// ContainerImage shall represent a container image available to a computer system.
type ContainerImage struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CreateTime shall indicate the date and time when the container image was created.
	CreateTime string
	// Description provides a description of this resource.
	Description string
	// ProgrammaticID shall contain the programmatic identifier for this container image. This is typically a hash
	// string that represents the identifier of this container image.
	ProgrammaticID string
	// SizeBytes shall contain the size of this container image in bytes.
	SizeBytes int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Type shall contain the supported image types for this container engine.
	Type ImageTypes
	// Version shall contain the version of this application.
	Version string

	containers []string
	// ContainersCount is the number of container instances running from this container image.
	ContainersCount int
	softwareImage   string
}

// UnmarshalJSON unmarshals a ContainerImage object from the raw JSON.
func (containerimage *ContainerImage) UnmarshalJSON(b []byte) error {
	type temp ContainerImage
	type Links struct {
		// Containers shall contain an array of links to resources of type Container that represent the container instances
		// running from this container image.
		Containers      common.Links
		ContainersCount int `json:"Containers@odata.count"`
		// SoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the software image
		// for this container image.
		SoftwareImage common.Link
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*containerimage = ContainerImage(t.temp)

	// Extract the links to other entities for later
	containerimage.containers = t.Links.Containers.ToStrings()
	containerimage.ContainersCount = t.Links.ContainersCount
	containerimage.softwareImage = t.Links.SoftwareImage.String()

	return nil
}

// GetContainerImage will get a ContainerImage instance from the service.
func GetContainerImage(c common.Client, uri string) (*ContainerImage, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var containerimage ContainerImage
	err = json.NewDecoder(resp.Body).Decode(&containerimage)
	if err != nil {
		return nil, err
	}

	containerimage.SetClient(c)
	return &containerimage, nil
}

// ListReferencedContainerImages gets the collection of ContainerImage from
// a provided reference.
func ListReferencedContainerImages(c common.Client, link string) ([]*ContainerImage, error) {
	return common.GetCollectionObjects(c, link, GetContainerImage)
}

// Containers get the container instances using this container image.
func (containerimage *ContainerImage) Containers() ([]*Container, error) {
	var result []*Container

	collectionError := common.NewCollectionError()
	for _, uri := range containerimage.containers {
		rb, err := GetContainer(containerimage.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, rb)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
