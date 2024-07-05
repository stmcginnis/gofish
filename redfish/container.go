//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

// Container shall represent an instance of a container that is running on a computer system.
type Container struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// EthernetInterfaces shall contain a link to a resource collection of type EthernetInterfaceCollection.
	ethernetInterfaces string
	// Limits shall contain the resource limits allocated to this container.
	Limits Limits
	// MountPoints shall contain the file system mount points configured for this container.
	MountPoints []MountPoint
	// ProgrammaticID shall contain the programmatic identifier for this container. This is typically a hash string
	// that represents the running instance of this container.
	ProgrammaticID string
	// StartTime shall indicate the date and time when the container started running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	resetTarget string

	containerImage string
}

// UnmarshalJSON unmarshals a Container object from the raw JSON.
func (container *Container) UnmarshalJSON(b []byte) error {
	type temp Container
	type Actions struct {
		Reset common.ActionTarget `json:"#Container.Reset"`
	}
	type Links struct {
		// ContainerImage shall contain a link to a resource of type ContainerImage that represents the container image for
		// this container.
		ContainerImage common.Link
	}
	var t struct {
		temp
		Actions            Actions
		EthernetInterfaces common.Link
		Links              Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*container = Container(t.temp)

	// Extract the links to other entities for later
	container.resetTarget = t.Actions.Reset.Target
	container.ethernetInterfaces = t.EthernetInterfaces.String()
	container.containerImage = t.Links.ContainerImage.String()

	return nil
}

// GetContainer will get a Container instance from the service.
func GetContainer(c common.Client, uri string) (*Container, error) {
	return common.GetObject[Container](c, uri)
}

// ListReferencedContainers gets the collection of Container from
// a provided reference.
func ListReferencedContainers(c common.Client, link string) ([]*Container, error) {
	return common.GetCollectionObjects[Container](c, link)
}

// Reset resets the container.
func (container *Container) Reset() error {
	if container.resetTarget == "" {
		return fmt.Errorf("Reset is not supported by this system")
	}

	return container.Post(container.resetTarget, nil)
}

// EthernetIntefaces gets the ethernet interfaces associated with this container.
func (container *Container) EthernetInterfaces() ([]*EthernetInterface, error) {
	if container.ethernetInterfaces == "" {
		return nil, nil
	}

	return ListReferencedEthernetInterfaces(container.GetClient(), container.ethernetInterfaces)
}

// ContainerImage gets the image used by this container.
func (container *Container) ContainerImage() (*ContainerImage, error) {
	if container.containerImage == "" {
		return nil, nil
	}

	return GetContainerImage(container.GetClient(), container.containerImage)
}

// Limits shall contain the resource limits allocated to a container.
type Limits struct {
	// CPUCount shall contain the number of processors available to this container.
	CPUCount float32
	// MemoryBytes shall contain the amount of memory available to this container in bytes.
	MemoryBytes int
}

// MountPoint shall contain a file system mount point configured for a container.
type MountPoint struct {
	// Destination shall contain the file system path in the container that is provided as the mount point to access
	// the files and folders specified by the Source property.
	Destination string
	// Source shall contain the file system path from the hosting system that is provided to this container.
	Source string
}
