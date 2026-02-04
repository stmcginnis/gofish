//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Container.v1_0_1.json
// 2023.2 - #Container.v1_0_1.Container

package schemas

import (
	"encoding/json"
)

// Container shall represent an instance of a container that is running on a
// computer system.
type Container struct {
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProgrammaticID shall contain the programmatic identifier for this container.
	// This is typically a hash string that represents the running instance of this
	// container.
	ProgrammaticID string `json:"ProgrammaticId"`
	// StartTime shall indicate the date and time when the container started
	// running.
	StartTime string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// containerImage is the URI for ContainerImage.
	containerImage string
}

// UnmarshalJSON unmarshals a Container object from the raw JSON.
func (c *Container) UnmarshalJSON(b []byte) error {
	type temp Container
	type cActions struct {
		Reset ActionTarget `json:"#Container.Reset"`
	}
	type cLinks struct {
		ContainerImage Link `json:"ContainerImage"`
	}
	var tmp struct {
		temp
		Actions            cActions
		Links              cLinks
		EthernetInterfaces Link `json:"EthernetInterfaces"`
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

	return nil
}

// GetContainer will get a Container instance from the service.
func GetContainer(c Client, uri string) (*Container, error) {
	return GetObject[Container](c, uri)
}

// ListReferencedContainers gets the collection of Container from
// a provided reference.
func ListReferencedContainers(c Client, link string) ([]*Container, error) {
	return GetCollectionObjects[Container](c, link)
}

// This action shall reset the container.
// resetType - This parameter shall contain the type of reset.
// 'GracefulRestart' and 'ForceRestart' shall indicate requests to restart the
// container. 'GracefulShutdown' and 'ForceOff' shall indicate requests to stop
// or disable the container. 'On' and 'ForceOn' shall indicate requests to
// start or enable the container. The service can accept a request without the
// parameter and shall perform a 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Container) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ContainerImage gets the ContainerImage linked resource.
func (c *Container) ContainerImage() (*ContainerImage, error) {
	if c.containerImage == "" {
		return nil, nil
	}
	return GetObject[ContainerImage](c.client, c.containerImage)
}

// EthernetInterfaces gets the EthernetInterfaces collection.
func (c *Container) EthernetInterfaces() ([]*EthernetInterface, error) {
	if c.ethernetInterfaces == "" {
		return nil, nil
	}
	return GetCollectionObjects[EthernetInterface](c.client, c.ethernetInterfaces)
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
