//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreweave/gofish/common"
)

// Container shall represent an instance of a container that is running on a computer system.
type Container struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
func GetContainer(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Container, error) {
	return GetContainerWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetContainerWithContext will get a Container instance from the service.
func GetContainerWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Container, error) {
	return common.GetObjectWithContext[Container](ctx, c, uri, queryOpts...)
}

// ListReferencedContainers gets the collection of Container from
// a provided reference.
func ListReferencedContainers(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Container, error) {
	return ListReferencedContainersWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedContainersWithContext gets the collection of Container from
// a provided reference.
func ListReferencedContainersWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*Container, error) {
	return common.GetCollectionObjectsWithContext[Container](ctx, c, link, queryOpts...)
}

// Reset resets the container.
func (container *Container) Reset() error {
	return container.ResetWithContext(common.ContextOf(container.GetClient()))
}

// ResetWithContext resets the container.
func (container *Container) ResetWithContext(ctx context.Context) error {
	if container.resetTarget == "" {
		return fmt.Errorf("Reset is not supported by this system")
	}

	return container.PostWithContext(ctx, container.resetTarget, nil)
}

// EthernetIntefaces gets the ethernet interfaces associated with this container.
func (container *Container) EthernetInterfaces(queryOpts ...common.QueryGroupOption) ([]*EthernetInterface, error) {
	return container.EthernetInterfacesWithContext(common.ContextOf(container.GetClient()), queryOpts...)
}

// EthernetIntefaces gets the ethernet interfaces associated with this container.
func (container *Container) EthernetInterfacesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*EthernetInterface, error) {
	if container.ethernetInterfaces == "" {
		return nil, nil
	}

	return ListReferencedEthernetInterfacesWithContext(ctx, container.GetClient(), container.ethernetInterfaces, queryOpts...)
}

// ContainerImage gets the image used by this container.
func (container *Container) ContainerImage(queryOpts ...common.QueryGroupOption) (*ContainerImage, error) {
	return container.ContainerImageWithContext(common.ContextOf(container.GetClient()), queryOpts...)
}

// ContainerImageWithContext gets the image used by this container.
func (container *Container) ContainerImageWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*ContainerImage, error) {
	if container.containerImage == "" {
		return nil, nil
	}

	return GetContainerImageWithContext(ctx, container.GetClient(), container.containerImage, queryOpts...)
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
