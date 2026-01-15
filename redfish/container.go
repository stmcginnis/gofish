//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #Container.v1_0_1.Container

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Container shall represent an instance of a container that is running on a
// computer system.
type Container struct {
	common.Entity
	// EthernetInterfaces shall contain a link to a resource collection of type
	// 'EthernetInterfaceCollection'.
	ethernetInterfaces string
	// Limits shall contain the resource limits allocated to this container.
	Limits Limits
	// MountPoints shall contain the file system mount points configured for this
	// container.
	MountPoints []MountPoint
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProgrammaticId shall contain the programmatic identifier for this container.
	// This is typically a hash string that represents the running instance of this
	// container.
	ProgrammaticID string `json:"ProgrammaticId"`
	// StartTime shall indicate the date and time when the container started
	// running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// containerImage is the URI for ContainerImage.
	containerImage string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Container object from the raw JSON.
func (c *Container) UnmarshalJSON(b []byte) error {
	type temp Container
	type cActions struct {
		Reset common.ActionTarget `json:"#Container.Reset"`
	}
	type cLinks struct {
		ContainerImage common.Link `json:"ContainerImage"`
	}
	var tmp struct {
		temp
		Actions            cActions
		Links              cLinks
		EthernetInterfaces common.Link `json:"ethernetInterfaces"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Container(tmp.temp)

	// Extract the links to other entities for later
	c.resetTarget = tmp.Actions.Reset.Target
	c.containerImage = tmp.Links.ContainerImage.String()
	c.ethernetInterfaces = tmp.EthernetInterfaces.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Container) Update() error {
	readWriteFields := []string{
		"Limits",
		"MountPoints",
		"Status",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
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

// Reset shall reset the container.
// resetType - This parameter shall contain the type of reset.
// 'GracefulRestart' and 'ForceRestart' shall indicate requests to restart the
// container. 'GracefulShutdown' and 'ForceOff' shall indicate requests to stop
// or disable the container. 'On' and 'ForceOn' shall indicate requests to
// start or enable the container. The service can accept a request without the
// parameter and shall perform a 'GracefulRestart'.
func (c *Container) Reset(resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return c.Post(c.resetTarget, payload)
}

// ContainerImage gets the ContainerImage linked resource.
func (c *Container) ContainerImage(client common.Client) (*ContainerImage, error) {
	if c.containerImage == "" {
		return nil, nil
	}
	return common.GetObject[ContainerImage](client, c.containerImage)
}

// EthernetInterfaces gets the EthernetInterfaces collection.
func (c *Container) EthernetInterfaces(client common.Client) ([]*EthernetInterface, error) {
	if c.ethernetInterfaces == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EthernetInterface](client, c.ethernetInterfaces)
}

// Limits shall contain the resource limits allocated to a container.
type Limits struct {
	// CPUCount shall contain the number of processors available to this container.
	CPUCount *float32 `json:",omitempty"`
	// MemoryBytes shall contain the amount of memory available to this container
	// in bytes.
	MemoryBytes *int `json:",omitempty"`
}

// MountPoint shall contain a file system mount point configured for a
// container.
type MountPoint struct {
	// Destination shall contain the file system path in the container that is
	// provided as the mount point to access the files and folders specified by the
	// 'Source' property.
	Destination string
	// Source shall contain the file system path from the hosting system that is
	// provided to this container.
	Source string
}
