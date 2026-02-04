//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Switch.v1_11_0.json
// 2025.4 - #Switch.v1_11_0.Switch

package schemas

import (
	"encoding/json"
)

type TargetType string

const (
	// FabricPortTargetType is a fabric port.
	FabricPortTargetType TargetType = "FabricPort"
	// HostEdgePortTargetType is a host edge port (USP/GAE).
	HostEdgePortTargetType TargetType = "HostEdgePort"
	// DownstreamEdgePortTargetType is a downstream edge port type.
	DownstreamEdgePortTargetType TargetType = "DownstreamEdgePort"
)

// Switch This resource contains a switch for a Redfish implementation.
type Switch struct {
	Entity
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the drive for inventory purposes.
	AssetTag string
	// CXL shall contain CXL-specific properties for this switch.
	//
	// Version added: v1.9.0
	CXL CXL
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
	IndicatorLED IndicatorLED
	// IsManaged shall indicate whether this switch is in a managed or unmanaged
	// state.
	IsManaged bool
	// Location shall contain the location information of the associated switch.
	//
	// Version added: v1.1.0
	Location Location
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
	// OEM shall contain the OEM extensions. All values for properties that this
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
	PowerState PowerState
	// Redundancy shall contain an array that shows how this switch is grouped with
	// other switches for form redundancy sets.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SKU shall contain the SKU number for this switch.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the switch.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedProtocols shall contain an array of protocols this switch supports.
	// If the value of 'SwitchType' is 'MultiProtocol', this property shall be
	// required.
	//
	// Version added: v1.3.0
	SupportedProtocols []Protocol
	// SwitchType shall contain the protocol being sent over this switch. For a
	// switch that supports multiple protocols, the value should be 'MultiProtocol'
	// and the 'SupportedProtocols' property should be used to describe the
	// supported protocols.
	SwitchType Protocol
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Switch object from the raw JSON.
func (s *Switch) UnmarshalJSON(b []byte) error {
	type temp Switch
	type sActions struct {
		Reset ActionTarget `json:"#Switch.Reset"`
	}
	type sLinks struct {
		Chassis    Link  `json:"Chassis"`
		Endpoints  Links `json:"Endpoints"`
		ManagedBy  Links `json:"ManagedBy"`
		PCIeDevice Link  `json:"PCIeDevice"`
	}
	var tmp struct {
		temp
		Actions            sActions
		Links              sLinks
		Certificates       Link `json:"Certificates"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		LogServices        Link `json:"LogServices"`
		Metrics            Link `json:"Metrics"`
		Ports              Link `json:"Ports"`
		Redundancy         Link `json:"Redundancy"`
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
	s.redundancy = tmp.Redundancy.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *Switch) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"Enabled",
		"IndicatorLED",
		"IsManaged",
		"LocationIndicatorActive",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSwitch will get a Switch instance from the service.
func GetSwitch(c Client, uri string) (*Switch, error) {
	return GetObject[Switch](c, uri)
}

// ListReferencedSwitchs gets the collection of Switch from
// a provided reference.
func ListReferencedSwitchs(c Client, link string) ([]*Switch, error) {
	return GetCollectionObjects[Switch](c, link)
}

// This action shall reset this switch.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without this parameter and can complete an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Switch) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Chassis gets the Chassis linked resource.
func (s *Switch) Chassis() (*Chassis, error) {
	if s.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](s.client, s.chassis)
}

// Endpoints gets the Endpoints linked resources.
func (s *Switch) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](s.client, s.endpoints)
}

// ManagedBy gets the ManagedBy linked resources.
func (s *Switch) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](s.client, s.managedBy)
}

// PCIeDevice gets the PCIeDevice linked resource.
func (s *Switch) PCIeDevice() (*PCIeDevice, error) {
	if s.pCIeDevice == "" {
		return nil, nil
	}
	return GetObject[PCIeDevice](s.client, s.pCIeDevice)
}

// Certificates gets the Certificates collection.
func (s *Switch) Certificates() ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](s.client, s.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (s *Switch) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if s.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](s.client, s.environmentMetrics)
}

// LogServices gets the LogServices collection.
func (s *Switch) LogServices() ([]*LogService, error) {
	if s.logServices == "" {
		return nil, nil
	}
	return GetCollectionObjects[LogService](s.client, s.logServices)
}

// Metrics gets the Metrics linked resource.
func (s *Switch) Metrics() (*SwitchMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return GetObject[SwitchMetrics](s.client, s.metrics)
}

// Ports gets the Ports collection.
func (s *Switch) Ports() ([]*Port, error) {
	if s.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](s.client, s.ports)
}

// Redundancy gets the Redundancy linked resource.
func (s *Switch) Redundancy() (*Redundancy, error) {
	if s.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](s.client, s.redundancy)
}

// CXL shall contain CXL-specific properties for a switch.
type CXL struct {
	// MaxSupportedPIDs shall contain the maximum number of port-based routing
	// (PBR) identifiers (PIDs) supported in this switch.
	//
	// Version added: v1.11.0
	MaxSupportedPIDs int
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
	// PIDTargetList shall contain the Compute Express Link Specification-defined
	// port-based routing (PBR) identifier (PID) target list for this CXL switch.
	//
	// Version added: v1.11.0
	PIDTargetList []PIDTargetList
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

// UnmarshalJSON unmarshals a CXL object from the raw JSON.
func (c *CXL) UnmarshalJSON(b []byte) error {
	type temp CXL
	var tmp struct {
		temp
		VirtualCXLSwitches Link `json:"VirtualCXLSwitches"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CXL(tmp.temp)

	// Extract the links to other entities for later
	c.virtualCXLSwitches = tmp.VirtualCXLSwitches.String()

	return nil
}

// VirtualCXLSwitches gets the VirtualCXLSwitches collection.
func (c *CXL) VirtualCXLSwitches(client Client) ([]*VirtualCXLSwitch, error) {
	if c.virtualCXLSwitches == "" {
		return nil, nil
	}
	return GetCollectionObjects[VirtualCXLSwitch](client, c.virtualCXLSwitches)
}

// PIDTargetList shall contain the Compute Express Link Specification-defined
// port-based routing (PBR) identifier (PID) target list properties for a CXL
// switch.
type PIDTargetList struct {
	// InstanceID shall contain the index of the PID for targets that support
	// multiple PIDs.
	//
	// Version added: v1.11.0
	InstanceID int
	// PID shall contain the Compute Express Link Specification-defined PID
	// assigned to this port.
	//
	// Version added: v1.11.0
	PID int
	// PortID shall contain the physical port identifier of the target port.
	//
	// Version added: v1.11.0
	PortID int
	// TargetID shall contain the identifier of the PID target for use in Compute
	// Express Link Specification-defined 'Configure PID Assignment'.
	//
	// Version added: v1.11.0
	TargetID int
	// TargetType shall contain the port type for the assigned PID.
	//
	// Version added: v1.11.0
	TargetType TargetType
	// VcsID shall contain the Compute Express Link Specification-defined virtual
	// CXL switch identifier of the target port.
	//
	// Version added: v1.11.0
	VcsID int
}

// VCSSwitch shall contain Virtual CXL Switch (VCS) properties for a switch.
type VCSSwitch struct {
	// HDMDecoders shall contain the number of Host Device Memory (HDM) Decoders
	// supported by this switch.
	//
	// Version added: v1.9.0
	HDMDecoders int
}
