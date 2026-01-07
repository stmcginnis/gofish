//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #Switch.v1_10_0.Switch

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Switch This resource contains a switch for a Redfish implementation.
type Switch struct {
	common.Entity
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the drive for inventory purposes.
	AssetTag string
	// CXL shall contain CXL-specific properties for this switch.
	//
	// Version added: v1.9.0
	CXL SwitchCXL
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.5.0
	certificates string
	// CurrentBandwidthGbps shall contain the internal unidirectional bandwidth of
	// this switch currently negotiated and running.
	//
	// Version added: v1.4.0
	CurrentBandwidthGbps *float64 `json:",omitempty"`
	// DomainID shall contain The domain ID for this switch. This property has a
	// scope of uniqueness within the fabric of which the switch is a member.
	DomainID *int `json:",omitempty"`
	// Enabled shall indicate if this switch is enabled.
	//
	// Version added: v1.6.0
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this switch.
	//
	// Version added: v1.6.0
	environmentMetrics string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated switch.
	//
	// Version added: v1.2.0
	FirmwareVersion string
	// IndicatorLED shall contain the state of the indicator light associated with
	// this switch.
	//
	// Deprecated: v1.4.0
	// This property has been deprecated in favor of the 'LocationIndicatorActive'
	// property.
	IndicatorLED common.IndicatorLED
	// IsManaged shall indicate whether this switch is in a managed or unmanaged
	// state.
	IsManaged bool
	// Location shall contain the location information of the associated switch.
	//
	// Version added: v1.1.0
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.4.0
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type
	// 'LogServiceCollection'.
	logServices string
	// Manufacturer shall contain the name of the organization responsible for
	// producing the switch. This organization may be the entity from which the
	// switch is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxBandwidthGbps shall contain the maximum internal unidirectional bandwidth
	// this switch is capable of being configured. If capable of autonegotiation,
	// the switch shall attempt to negotiate to the specified maximum bandwidth.
	//
	// Version added: v1.4.0
	MaxBandwidthGbps *float64 `json:",omitempty"`
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.5.0
	//
	// Deprecated: v1.8.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// Metrics shall contain a link to the metrics associated with this switch.
	//
	// Version added: v1.7.0
	metrics string
	// Model shall contain the manufacturer-provided model information of this
	// switch.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the
	// switch.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// PowerState shall contain the power state of the switch.
	PowerState common.PowerState
	// Redundancy shall contain an array that shows how this switch is grouped with
	// other switches for form redundancy sets.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SKU shall contain the SKU number for this switch.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the switch.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedProtocols shall contain an array of protocols this switch supports.
	// If the value of 'SwitchType' is 'MultiProtocol', this property shall be
	// required.
	//
	// Version added: v1.3.0
	SupportedProtocols []common.Protocol
	// SwitchType shall contain the protocol being sent over this switch. For a
	// switch that supports multiple protocols, the value should be 'MultiProtocol'
	// and the 'SupportedProtocols' property should be used to describe the
	// supported protocols.
	SwitchType common.Protocol
	// TotalSwitchWidth shall contain the number of physical transport lanes, phys,
	// or other physical transport links that this switch contains. For PCIe, this
	// value shall be the lane count.
	TotalSwitchWidth *int `json:",omitempty"`
	// UUID shall contain a universally unique identifier number for the switch.
	//
	// Version added: v1.3.0
	UUID string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// chassis is the URI for Chassis.
	chassis string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// pCIeDevice is the URI for PCIeDevice.
	pCIeDevice string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Switch object from the raw JSON.
func (s *Switch) UnmarshalJSON(b []byte) error {
	type temp Switch
	type sActions struct {
		Reset common.ActionTarget `json:"#Switch.Reset"`
	}
	type sLinks struct {
		Chassis    common.Link  `json:"Chassis"`
		Endpoints  common.Links `json:"Endpoints"`
		ManagedBy  common.Links `json:"ManagedBy"`
		PCIeDevice common.Link  `json:"PCIeDevice"`
	}
	var tmp struct {
		temp
		Actions            sActions
		Links              sLinks
		Certificates       common.Link `json:"certificates"`
		EnvironmentMetrics common.Link `json:"environmentMetrics"`
		LogServices        common.Link `json:"logServices"`
		Metrics            common.Link `json:"metrics"`
		Ports              common.Link `json:"ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Switch(tmp.temp)

	// Extract the links to other entities for later
	s.resetTarget = tmp.Actions.Reset.Target
	s.chassis = tmp.Links.Chassis.String()
	s.endpoints = tmp.Links.Endpoints.ToStrings()
	s.managedBy = tmp.Links.ManagedBy.ToStrings()
	s.pCIeDevice = tmp.Links.PCIeDevice.String()
	s.certificates = tmp.Certificates.String()
	s.environmentMetrics = tmp.EnvironmentMetrics.String()
	s.logServices = tmp.LogServices.String()
	s.metrics = tmp.Metrics.String()
	s.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *Switch) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"CXL",
		"Enabled",
		"IndicatorLED",
		"IsManaged",
		"Location",
		"LocationIndicatorActive",
		"Measurements",
		"Redundancy",
		"Redundancy@odata.count",
		"Status",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSwitch will get a Switch instance from the service.
func GetSwitch(c common.Client, uri string) (*Switch, error) {
	return common.GetObject[Switch](c, uri)
}

// ListReferencedSwitchs gets the collection of Switch from
// a provided reference.
func ListReferencedSwitchs(c common.Client, link string) ([]*Switch, error) {
	return common.GetCollectionObjects[Switch](c, link)
}

// Reset shall reset this switch.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without this parameter and can complete an
// implementation-specific default reset.
func (s *Switch) Reset(resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return s.Post(s.resetTarget, payload)
}

// Chassis gets the Chassis linked resource.
func (s *Switch) Chassis(client common.Client) (*Chassis, error) {
	if s.chassis == "" {
		return nil, nil
	}
	return common.GetObject[Chassis](client, s.chassis)
}

// Endpoints gets the Endpoints linked resources.
func (s *Switch) Endpoints(client common.Client) ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](client, s.endpoints)
}

// ManagedBy gets the ManagedBy linked resources.
func (s *Switch) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, s.managedBy)
}

// PCIeDevice gets the PCIeDevice linked resource.
func (s *Switch) PCIeDevice(client common.Client) (*PCIeDevice, error) {
	if s.pCIeDevice == "" {
		return nil, nil
	}
	return common.GetObject[PCIeDevice](client, s.pCIeDevice)
}

// Certificates gets the Certificates collection.
func (s *Switch) Certificates(client common.Client) ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, s.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (s *Switch) EnvironmentMetrics(client common.Client) (*EnvironmentMetrics, error) {
	if s.environmentMetrics == "" {
		return nil, nil
	}
	return common.GetObject[EnvironmentMetrics](client, s.environmentMetrics)
}

// LogServices gets the LogServices collection.
func (s *Switch) LogServices(client common.Client) ([]*LogService, error) {
	if s.logServices == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[LogService](client, s.logServices)
}

// Metrics gets the Metrics linked resource.
func (s *Switch) Metrics(client common.Client) (*SwitchMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return common.GetObject[SwitchMetrics](client, s.metrics)
}

// Ports gets the Ports collection.
func (s *Switch) Ports(client common.Client) ([]*Port, error) {
	if s.ports == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, s.ports)
}

// SwitchCXL shall contain CXL-specific properties for a switch.
type SwitchCXL struct {
	// MaxVCSsSupported shall contain the maximum number of Virtual CXL Switches
	// (VCSs) supported in this switch.
	//
	// Version added: v1.9.0
	MaxVCSsSupported int
	// NumberOfBoundvPPBs shall contain the total number of vPPBs (Virtual
	// PCI-to-PCI Bridges) that are currently bound on this switch.
	//
	// Version added: v1.10.0
	NumberOfBoundvPPBs int
	// PBRCapable shall indicate whether the switch is capable of performing
	// port-based routing.
	//
	// Version added: v1.10.0
	PBRCapable bool
	// TotalHDMDecoders shall contain the total number of HDM (Host Device Memory)
	// decoders available per upstream port.
	//
	// Version added: v1.10.0
	TotalHDMDecoders int
	// TotalNumbervPPBs shall contain the total number of virtual PCI-to-PCI
	// bridges (vPPBs) supported in this switch.
	//
	// Version added: v1.9.0
	TotalNumbervPPBs int
	// VCS shall contain Virtual CXL Switch (VCS) properties for this switch.
	//
	// Version added: v1.9.0
	//
	// Deprecated: v1.9.0
	// This property has been deprecated in favor of 'VirtualCXLSwitches' in 'CXL'.
	VCS VCSSwitch
	// VirtualCXLSwitches shall contain a link to a resource collection of type
	// 'VirtualCXLSwitchCollection'.
	//
	// Version added: v1.10.0
	virtualCXLSwitches string
}

// UnmarshalJSON unmarshals a SwitchCXL object from the raw JSON.
func (c *SwitchCXL) UnmarshalJSON(b []byte) error {
	type temp SwitchCXL
	var tmp struct {
		temp
		VirtualCXLSwitches common.Link `json:"virtualCXLSwitches"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = SwitchCXL(tmp.temp)

	// Extract the links to other entities for later
	c.virtualCXLSwitches = tmp.VirtualCXLSwitches.String()

	return nil
}

// VirtualCXLSwitches gets the VirtualCXLSwitches collection.
func (c *SwitchCXL) VirtualCXLSwitches(client common.Client) ([]*VirtualCXLSwitch, error) {
	if c.virtualCXLSwitches == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[VirtualCXLSwitch](client, c.virtualCXLSwitches)
}

// VCSSwitch shall contain Virtual CXL Switch (VCS) properties for a switch.
type VCSSwitch struct {
	// HDMDecoders shall contain the number of Host Device Memory (HDM) Decoders
	// supported by this switch.
	//
	// Version added: v1.9.0
	HDMDecoders int
}
