//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SwitchCXL shall contain CXL-specific properties for a switch.
type SwitchCXL struct {
	// MaxVCSsSupported shall contain the maximum number of Virtual CXL Switches (VCSs) supported in this switch.
	MaxVCSsSupported string
	// TotalNumbervPPBs shall contain the total number of virtual PCI-to-PCI bridges (vPPBs) supported in this switch.
	TotalNumbervPPBs string
	// VCS shall contain Virtual CXL Switch (VCS) properties for this switch.
	VCS string
}

// Switch This resource contains a switch for a Redfish implementation.
type Switch struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the drive for
	// inventory purposes.
	AssetTag string
	// CXL shall contain CXL-specific properties for this switch.
	CXL SwitchCXL
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates string
	// CurrentBandwidthGbps shall contain the internal unidirectional bandwidth of this switch currently negotiated and
	// running.
	CurrentBandwidthGbps float64
	// Description provides a description of this resource.
	Description string
	// DomainID shall contain the domain ID for this switch. This property has a scope of uniqueness within the fabric
	// of which the switch is a member.
	DomainID int
	// Enabled shall indicate if this switch is enabled.
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this switch.
	EnvironmentMetrics EnvironmentMetrics
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for the associated switch.
	FirmwareVersion string
	// IsManaged shall indicate whether this switch is in a managed or unmanaged state.
	IsManaged bool
	// Location shall contain the location information of the associated switch.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type LogServiceCollection.
	logServices string
	// Manufacturer shall contain the name of the organization responsible for producing the switch. This organization
	// may be the entity from which the switch is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxBandwidthGbps shall contain the maximum internal unidirectional bandwidth this switch is capable of being
	// configured. If capable of autonegotiation, the switch shall attempt to negotiate to the specified maximum
	// bandwidth.
	MaxBandwidthGbps float64
	// Metrics shall contain a link to the metrics associated with this switch.
	metrics string
	// Model shall contain the manufacturer-provided model information of this switch.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the switch.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// PowerState shall contain the power state of the switch.
	PowerState PowerState
	// Redundancy shall contain an array that shows how this switch is grouped with other switches for form redundancy
	// sets.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SKU shall contain the SKU number for this switch.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the switch.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedProtocols shall contain an array of protocols this switch supports. If the value of SwitchType is
	// 'MultiProtocol', this property shall be required.
	SupportedProtocols []common.Protocol
	// SwitchType shall contain the protocol being sent over this switch. For a switch that supports multiple
	// protocols, the value should be 'MultiProtocol' and the SupportedProtocols property should be used to describe
	// the supported protocols.
	SwitchType common.Protocol
	// TotalSwitchWidth shall contain the number of physical transport lanes, phys, or other physical transport links
	// that this switch contains. For PCIe, this value shall be the lane count.
	TotalSwitchWidth int
	// UUID shall contain a universally unique identifier number for the switch.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	chassis        string
	endpoints      []string
	EndpointsCount int
	managedBy      []string
	ManagedByCount int
	pcieDevice     string

	resetTarget string
}

// UnmarshalJSON unmarshals a Switch object from the raw JSON.
func (sw *Switch) UnmarshalJSON(b []byte) error {
	type temp Switch
	type Actions struct {
		Reset common.ActionTarget `json:"#Switch.Reset"`
	}
	type Links struct {
		// Chassis shall contain a link to a resource of type Chassis with which this switch is associated.
		Chassis common.Link
		// Endpoints shall contain an array of links to resources of type Endpoint with which this switch is associated.
		Endpoints      common.Links
		EndpointsCount int `json:"Endpoints@odata.count"`
		// ManagedBy shall contain an array of links to resources of type Manager with which this switch is associated.
		ManagedBy      common.Links
		ManagedByCount int `json:"ManagedBy@odata.count"`
		// PCIeDevice shall contain a link to a resource of type PCIeDevice that represents the PCIe device providing this
		// switch.
		PCIeDevice common.Link
	}
	var t struct {
		temp
		Actions      Actions
		Links        Links
		Certificates common.Link
		LogServices  common.Link
		Metrics      common.Link
		Ports        common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sw = Switch(t.temp)

	// Extract the links to other entities for later
	sw.certificates = t.Certificates.String()
	sw.logServices = t.LogServices.String()
	sw.metrics = t.Metrics.String()
	sw.ports = t.Ports.String()

	sw.chassis = t.Links.Chassis.String()
	sw.endpoints = t.Links.Endpoints.ToStrings()
	sw.EndpointsCount = t.Links.EndpointsCount
	sw.managedBy = t.Links.ManagedBy.ToStrings()
	sw.ManagedByCount = t.Links.ManagedByCount
	sw.pcieDevice = t.Links.PCIeDevice.String()

	sw.resetTarget = t.Actions.Reset.Target

	// This is a read/write object, so we need to save the raw object data for later
	sw.rawData = b

	return nil
}

// Certificates returns certificates related to this device.
func (sw *Switch) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(sw.GetClient(), sw.certificates)
}

// LogServices gets the log services related to this device.
func (sw *Switch) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(sw.GetClient(), sw.logServices)
}

// Metrics gets the switch metrics related to this device.
func (sw *Switch) Metrics() (*SwitchMetrics, error) {
	return GetSwitchMetrics(sw.GetClient(), sw.metrics)
}

// Ports gets the ports related to this device.
func (sw *Switch) Ports() ([]*Port, error) {
	return ListReferencedPorts(sw.GetClient(), sw.ports)
}

// Chassis gets the containing chassis of this device.
func (sw *Switch) Chassis() (*Chassis, error) {
	return GetChassis(sw.GetClient(), sw.chassis)
}

// Endpoints gets any endpoints associated with this fabric.
func (sw *Switch) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](sw.GetClient(), sw.endpoints)
}

// ManagedBy gets the managers of this fabric.
func (sw *Switch) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](sw.GetClient(), sw.managedBy)
}

// PCIeDevice gets the PCIe device providing this switch.
func (sw *Switch) PCIeDevice() (*PCIeDevice, error) {
	return GetPCIeDevice(sw.GetClient(), sw.pcieDevice)
}

// Reset resets this switch.
func (sw *Switch) Reset(resetType ResetType) error {
	if sw.resetTarget == "" {
		return errors.New("Reset is not supported by this system")
	}

	parameters := struct {
		ResetType ResetType
	}{
		ResetType: resetType,
	}
	return sw.Post(sw.resetTarget, parameters)
}

// Update commits updates to this object's properties to the running system.
func (sw *Switch) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Switch)
	original.UnmarshalJSON(sw.rawData)

	readWriteFields := []string{
		"AssetTag",
		"Enabled",
		"IsManaged",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(sw).Elem()

	return sw.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSwitch will get a Switch instance from the service.
func GetSwitch(c common.Client, uri string) (*Switch, error) {
	return common.GetObject[Switch](c, uri)
}

// ListReferencedSwitches gets the collection of Switch from
// a provided reference.
func ListReferencedSwitches(c common.Client, link string) ([]*Switch, error) {
	return common.GetCollectionObjects[Switch](c, link)
}

// VCSSwitch shall contain Virtual CXL Switch (VCS) properties for a switch.
type VCSSwitch struct {
	// HDMDecoders shall contain the number of Host Device Memory (HDM) Decoders supported by this switch.
	HDMDecoders string
}
