//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/StorageController.v1_11_0.json
// 2025.4 - #StorageController.v1_11_0.StorageController

package schemas

import (
	"encoding/json"
)

type ANAAccessState string

const (
	// OptimizedANAAccessState Commands processed by a controller provide optimized
	// access to any namespace in the ANA group.
	OptimizedANAAccessState ANAAccessState = "Optimized"
	// NonOptimizedANAAccessState Commands processed by a controller that reports
	// this state for an ANA group provide non-optimized access characteristics,
	// such as lower performance or non-optimal use of subsystem resources, to any
	// namespace in the ANA group.
	NonOptimizedANAAccessState ANAAccessState = "NonOptimized"
	// InaccessibleANAAccessState Namespaces in this group are inaccessible.
	// Commands are not able to access user data of namespaces in the ANA group.
	InaccessibleANAAccessState ANAAccessState = "Inaccessible"
	// PersistentLossANAAccessState The group is persistently inaccessible.
	// Commands are persistently not able to access user data of namespaces in the
	// ANA group.
	PersistentLossANAAccessState ANAAccessState = "PersistentLoss"
)

type NVMeControllerType string

const (
	// AdminNVMeControllerType The NVMe controller is an admin controller.
	AdminNVMeControllerType NVMeControllerType = "Admin"
	// DiscoveryNVMeControllerType The NVMe controller is a discovery controller.
	DiscoveryNVMeControllerType NVMeControllerType = "Discovery"
	// IONVMeControllerType The NVMe controller is an I/O controller.
	IONVMeControllerType NVMeControllerType = "IO"
)

// StorageController shall represent a storage controller in the Redfish
// Specification.
type StorageController struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// AssetTag shall track the storage controller for inventory purposes.
	AssetTag string
	// CacheSummary shall contain properties that describe the cache memory for
	// this resource.
	CacheSummary CacheSummary
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.1.0
	certificates string
	// ControllerRates shall contain all the rate settings available on the
	// controller.
	ControllerRates Rates
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that represents the environment metrics for this
	// storage controller.
	//
	// Version added: v1.2.0
	environmentMetrics string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated storage controller.
	FirmwareVersion string
	// Identifiers shall contain a list of all known durable names for the
	// associated storage controller.
	Identifiers []Identifier
	// IsLogical shall indicate whether this is a logical storage controller.
	//
	// Version added: v1.10.0
	IsLogical bool
	// Location shall contain the location information of the associated storage
	// controller.
	Location Location
	// MPFProperties shall contain the physical function-related properties for
	// this storage controller within a multiple physical function storage
	// controller. This property should only be present if 'IsLogical' contains
	// 'true'.
	//
	// Version added: v1.10.0
	MPFProperties MPFProperties
	// Manufacturer shall contain the name of the organization responsible for
	// producing the storage controller. This organization may be the entity from
	// which the storage controller is purchased, but this is not necessarily true.
	Manufacturer string
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// Metrics shall contain a link to the metrics associated with this storage
	// controller.
	//
	// Version added: v1.7.0
	metrics string
	// Model shall contain the name by which the manufacturer generally refers to
	// the storage controller.
	Model string
	// NVMeControllerProperties shall contain NVMe-related properties for this
	// storage controller.
	NVMeControllerProperties NVMeControllerProperties
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this
	// PCIe-based controller to its host.
	PCIeInterface PCIeInterface
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the storage controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// SKU shall contain the stock-keeping unit number for this storage controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the storage controller.
	SerialNumber string
	// SpeedGbps shall represent the maximum supported speed of the storage bus
	// interface, in Gbit/s. The specified interface connects the controller to the
	// storage devices, not the controller to a host. For example, SAS bus not PCIe
	// host bus.
	SpeedGbps *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedControllerProtocols shall contain the supported set of protocols
	// for communicating with this storage controller.
	SupportedControllerProtocols []Protocol
	// SupportedDeviceProtocols shall contain the set of protocols this storage
	// controller can use to communicate with attached devices.
	SupportedDeviceProtocols []Protocol
	// SupportedRAIDTypes shall contain an array of all the RAID types supported by
	// this controller.
	SupportedRAIDTypes []RAIDType
	// attachNamespacesTarget is the URL to send AttachNamespaces requests.
	attachNamespacesTarget string
	// detachNamespacesTarget is the URL to send DetachNamespaces requests.
	detachNamespacesTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// securityReceiveTarget is the URL to send SecurityReceive requests.
	securityReceiveTarget string
	// securitySendTarget is the URL to send SecuritySend requests.
	securitySendTarget string
	// attachedVolumes are the URIs for AttachedVolumes.
	attachedVolumes []string
	// batteries are the URIs for Batteries.
	batteries []string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// nVMeDiscoveredSubsystems are the URIs for NVMeDiscoveredSubsystems.
	nVMeDiscoveredSubsystems []string
	// networkDeviceFunctions are the URIs for NetworkDeviceFunctions.
	networkDeviceFunctions []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a StorageController object from the raw JSON.
func (s *StorageController) UnmarshalJSON(b []byte) error {
	type temp StorageController
	type sActions struct {
		AttachNamespaces ActionTarget `json:"#StorageController.AttachNamespaces"`
		DetachNamespaces ActionTarget `json:"#StorageController.DetachNamespaces"`
		Reset            ActionTarget `json:"#StorageController.Reset"`
		SecurityReceive  ActionTarget `json:"#StorageController.SecurityReceive"`
		SecuritySend     ActionTarget `json:"#StorageController.SecuritySend"`
	}
	type sLinks struct {
		AttachedVolumes          Links `json:"AttachedVolumes"`
		Batteries                Links `json:"Batteries"`
		Endpoints                Links `json:"Endpoints"`
		NVMeDiscoveredSubsystems Links `json:"NVMeDiscoveredSubsystems"`
		NetworkDeviceFunctions   Links `json:"NetworkDeviceFunctions"`
		PCIeFunctions            Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Actions            sActions
		Links              sLinks
		Assembly           Link `json:"Assembly"`
		Certificates       Link `json:"Certificates"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		Metrics            Link `json:"Metrics"`
		Ports              Link `json:"Ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageController(tmp.temp)

	// Extract the links to other entities for later
	s.attachNamespacesTarget = tmp.Actions.AttachNamespaces.Target
	s.detachNamespacesTarget = tmp.Actions.DetachNamespaces.Target
	s.resetTarget = tmp.Actions.Reset.Target
	s.securityReceiveTarget = tmp.Actions.SecurityReceive.Target
	s.securitySendTarget = tmp.Actions.SecuritySend.Target
	s.attachedVolumes = tmp.Links.AttachedVolumes.ToStrings()
	s.batteries = tmp.Links.Batteries.ToStrings()
	s.endpoints = tmp.Links.Endpoints.ToStrings()
	s.nVMeDiscoveredSubsystems = tmp.Links.NVMeDiscoveredSubsystems.ToStrings()
	s.networkDeviceFunctions = tmp.Links.NetworkDeviceFunctions.ToStrings()
	s.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	s.assembly = tmp.Assembly.String()
	s.certificates = tmp.Certificates.String()
	s.environmentMetrics = tmp.EnvironmentMetrics.String()
	s.metrics = tmp.Metrics.String()
	s.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageController) Update() error {
	readWriteFields := []string{
		"AssetTag",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetStorageController will get a StorageController instance from the service.
func GetStorageController(c Client, uri string) (*StorageController, error) {
	return GetObject[StorageController](c, uri)
}

// ListReferencedStorageControllers gets the collection of StorageController from
// a provided reference.
func ListReferencedStorageControllers(c Client, link string) ([]*StorageController, error) {
	return GetCollectionObjects[StorageController](c, link)
}

// This action shall attach referenced namespaces or volumes to the storage
// controller. For NVMe, namespaces are modeled as 'Volume' resources. Services
// shall add the attached namespaces and volumes to the 'AttachedVolumes'
// property in 'Links'.
// namespaces - This parameter shall contain an array of links to resources of
// type 'Volume' that represent the namespaces or volumes to attach to the
// storage controller.
func (s *StorageController) AttachNamespaces(namespaces []string) (*AttachDetachNamespacesResponse, error) {
	payload := make(map[string]any)
	payload["Namespaces"] = namespaces

	resp, err := s.PostWithResponse(s.attachNamespacesTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result AttachDetachNamespacesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall detach referenced namespaces or volumes from the storage
// controller. Services shall remove the detached namespaces and volumes from
// the 'AttachedVolumes' property in 'Links'. For NVMe, namespaces are modeled
// as 'Volume' resources.
// namespaces - This parameter shall contain an array of links to resources of
// type 'Volume' that represent the namespaces or volumes to detach from the
// storage controller.
func (s *StorageController) DetachNamespaces(namespaces []string) (*AttachDetachNamespacesResponse, error) {
	payload := make(map[string]any)
	payload["Namespaces"] = namespaces

	resp, err := s.PostWithResponse(s.detachNamespacesTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result AttachDetachNamespacesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall reset this storage controller.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without this parameter and can complete an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageController) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall transfer security protocol data from the controller. The
// contents of the request are specified by the 'SECURITY PROTOCOL IN command'
// section of the SPC-5 Specification.
// allocationLength - This parameter shall contain the allocated size of the
// received data, which imposes a maximum length of the data. The response may
// contain padding to meet this length.
// securityProtocol - This parameter shall contain the numeric identifier of
// the security protocol, as defined by the 'SECURITY PROTOCOL field in
// SECURITY PROTOCOL IN command' table of the SPC-5 Specification, and possibly
// extended by transport-specific standards. Services shall only accept the
// values '0', '1', or '2'.
// securityProtocolSpecific - This parameter shall contain the security
// protocol-specific data for the transfer operation. The value is defined by
// the protocol specified by the SecurityProtocolSpecific parameter.
func (s *StorageController) SecurityReceive(allocationLength int, securityProtocol int, securityProtocolSpecific int) (*SecurityReceiveResponse, error) {
	payload := make(map[string]any)
	payload["AllocationLength"] = allocationLength
	payload["SecurityProtocol"] = securityProtocol
	payload["SecurityProtocolSpecific"] = securityProtocolSpecific

	resp, err := s.PostWithResponse(s.securityReceiveTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result SecurityReceiveResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall transfer security protocol data to the controller. The
// contents of the request are specified by the 'SECURITY PROTOCOL OUT command'
// section of the SPC-5 Specification.
// data - This parameter shall contain a Base64-encoded string, with padding
// characters, of the security protocol data to transfer.
// securityProtocol - This parameter shall contain the numeric identifier of
// the security protocol, as defined by the 'SECURITY PROTOCOL field in
// SECURITY PROTOCOL OUT command' table of the SPC-5 Specification, and
// possibly extended by transport-specific standards. Services shall only
// accept the values '1' or '2'.
// securityProtocolSpecific - This parameter shall contain the security
// protocol-specific data for the transfer operation. The value is defined by
// the protocol specified by the SecurityProtocolSpecific parameter.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageController) SecuritySend(data string, securityProtocol int, securityProtocolSpecific int) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Data"] = data
	payload["SecurityProtocol"] = securityProtocol
	payload["SecurityProtocolSpecific"] = securityProtocolSpecific
	resp, taskInfo, err := PostWithTask(s.client,
		s.securitySendTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AttachedVolumes gets the AttachedVolumes linked resources.
func (s *StorageController) AttachedVolumes() ([]*Volume, error) {
	return GetObjects[Volume](s.client, s.attachedVolumes)
}

// Batteries gets the Batteries linked resources.
func (s *StorageController) Batteries() ([]*Battery, error) {
	return GetObjects[Battery](s.client, s.batteries)
}

// Endpoints gets the Endpoints linked resources.
func (s *StorageController) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](s.client, s.endpoints)
}

// NVMeDiscoveredSubsystems gets the NVMeDiscoveredSubsystems linked resources.
func (s *StorageController) NVMeDiscoveredSubsystems() ([]*Storage, error) {
	return GetObjects[Storage](s.client, s.nVMeDiscoveredSubsystems)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions linked resources.
func (s *StorageController) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](s.client, s.networkDeviceFunctions)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (s *StorageController) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](s.client, s.pCIeFunctions)
}

// Assembly gets the Assembly linked resource.
func (s *StorageController) Assembly() (*Assembly, error) {
	if s.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](s.client, s.assembly)
}

// Certificates gets the Certificates collection.
func (s *StorageController) Certificates() ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](s.client, s.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (s *StorageController) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if s.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](s.client, s.environmentMetrics)
}

// Metrics gets the Metrics linked resource.
func (s *StorageController) Metrics() (*StorageControllerMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return GetObject[StorageControllerMetrics](s.client, s.metrics)
}

// Ports gets the Ports collection.
func (s *StorageController) Ports() ([]*Port, error) {
	if s.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](s.client, s.ports)
}

// ANACharacteristics shall contain the ANA characteristics and volume
// information for a storage controller.
type ANACharacteristics struct {
	// AccessState shall contain the reported ANA access state.
	AccessState ANAAccessState
	// Volume shall contain a link to a resource of type 'Volume'.
	volume string
}

// UnmarshalJSON unmarshals a ANACharacteristics object from the raw JSON.
func (a *ANACharacteristics) UnmarshalJSON(b []byte) error {
	type temp ANACharacteristics
	var tmp struct {
		temp
		Volume Link `json:"Volume"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = ANACharacteristics(tmp.temp)

	// Extract the links to other entities for later
	a.volume = tmp.Volume.String()

	return nil
}

// Volume gets the Volume linked resource.
func (a *ANACharacteristics) Volume(client Client) (*Volume, error) {
	if a.volume == "" {
		return nil, nil
	}
	return GetObject[Volume](client, a.volume)
}

// AttachDetachNamespacesResponse shall contain the properties found in the
// response body for the 'AttachNamespaces' and 'DetachNamespaces' actions.
type AttachDetachNamespacesResponse struct {
	// AttachedVolumes shall contain an array of links to resources of type
	// 'Volume' that are attached to this instance of storage controller.
	//
	// Version added: v1.7.0
	attachedVolumes []string
	// AttachedVolumesCount
	AttachedVolumesCount int `json:"AttachedVolumes@odata.count"`
}

// UnmarshalJSON unmarshals a AttachDetachNamespacesResponse object from the raw JSON.
func (a *AttachDetachNamespacesResponse) UnmarshalJSON(b []byte) error {
	type temp AttachDetachNamespacesResponse
	var tmp struct {
		temp
		AttachedVolumes Links `json:"AttachedVolumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AttachDetachNamespacesResponse(tmp.temp)

	// Extract the links to other entities for later
	a.attachedVolumes = tmp.AttachedVolumes.ToStrings()

	return nil
}

// AttachedVolumes gets the AttachedVolumes linked resources.
func (a *AttachDetachNamespacesResponse) AttachedVolumes(client Client) ([]*Volume, error) {
	return GetObjects[Volume](client, a.attachedVolumes)
}

// MPFProperties shall contain the physical function-related properties for a
// storage controller within a multiple physical function storage controller.
type MPFProperties struct {
	// IsSupervisor shall indicate whether this physical function is the
	// supervisor, which manages the other physical functions, manages the
	// configuration of the controller, and controls management access.
	//
	// Version added: v1.10.0
	IsSupervisor bool
}

// NVMeControllerAttributes shall contain NVMe controller attributes for a
// storage controller.
type NVMeControllerAttributes struct {
	// ReportsNamespaceGranularity shall indicate whether or not the controller
	// supports reporting of Namespace Granularity.
	ReportsNamespaceGranularity bool
	// ReportsUUIDList shall indicate whether or not the controller supports
	// reporting of a UUID list.
	ReportsUUIDList bool
	// Supports128BitHostID shall indicate whether or not the controller supports a
	// 128-bit Host Identifier.
	Supports128BitHostID bool `json:"Supports128BitHostId"`
	// SupportsEnduranceGroups shall indicate whether or not the controller
	// supports Endurance Groups.
	SupportsEnduranceGroups bool
	// SupportsExceedingPowerOfNonOperationalState shall indicate whether or not
	// the controller supports exceeding Power of Non-Operational State in order to
	// execute controller-initiated background operations in a non-operational
	// power state.
	SupportsExceedingPowerOfNonOperationalState bool
	// SupportsNVMSets shall indicate whether or not the controller supports NVM
	// Sets.
	SupportsNVMSets bool
	// SupportsPredictableLatencyMode shall indicate whether or not the controller
	// supports Predictable Latency Mode.
	SupportsPredictableLatencyMode bool
	// SupportsReadRecoveryLevels shall indicate whether or not the controller
	// supports Read Recovery Levels.
	SupportsReadRecoveryLevels bool
	// SupportsReservations shall indicate if the controller supports reservations.
	//
	// Version added: v1.2.0
	SupportsReservations bool
	// SupportsSQAssociations shall indicate whether or not the controller supports
	// SQ Associations.
	SupportsSQAssociations bool
	// SupportsTrafficBasedKeepAlive shall indicate whether or not the controller
	// supports restarting the Keep Alive Timer if traffic is processed from an
	// admin command or I/O during a Keep Alive Timeout interval.
	SupportsTrafficBasedKeepAlive bool
}

// NVMeControllerProperties shall contain NVMe-related properties for a storage
// controller.
type NVMeControllerProperties struct {
	// ANACharacteristics shall contain the ANA characteristics and volume
	// information.
	ANACharacteristics []ANACharacteristics
	// AllocatedCompletionQueues shall contain the number of I/O completion queues
	// allocated to this NVMe I/O controller.
	//
	// Version added: v1.4.0
	AllocatedCompletionQueues *int `json:",omitempty"`
	// AllocatedSubmissionQueues shall contain the number of I/O submission queues
	// allocated to this NVMe I/O controller.
	//
	// Version added: v1.4.0
	AllocatedSubmissionQueues *int `json:",omitempty"`
	// ControllerType shall contain the type of NVMe controller.
	ControllerType NVMeControllerType
	// DiscoveryTransportServiceID shall contain the NVMe discovery transport
	// service identifier for the discovery controller. This property shall only be
	// present if 'ControllerType' contains 'Discovery'. For NVMe/TCP, the default
	// value should be '8009'.
	//
	// Version added: v1.9.0
	DiscoveryTransportServiceID *int `json:"DiscoveryTransportServiceId,omitempty"`
	// MaxAttachedNamespaces shall contain the maximum number of attached
	// namespaces allowed by this NVMe I/O controller.
	//
	// Version added: v1.8.0
	MaxAttachedNamespaces *uint `json:",omitempty"`
	// MaxQueueSize shall contain the maximum individual queue entry size supported
	// per queue. This is a zero-based value, where the minimum value is one,
	// indicating two entries. For PCIe, this applies to both submission and
	// completion queues. For NVMe-oF, this applies only to submission queues.
	MaxQueueSize *int `json:",omitempty"`
	// NVMeControllerAttributes shall contain NVMe controller attributes.
	NVMeControllerAttributes NVMeControllerAttributes
	// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for
	// this storage controller. This property can contain possible triggers for the
	// predictive drive failure warning for the corresponding drive.
	NVMeSMARTCriticalWarnings NVMeSMARTCriticalWarnings
	// NVMeVersion shall contain the version of the NVMe Base Specification
	// supported.
	NVMeVersion string
}

// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for
// a storage controller.
type NVMeSMARTCriticalWarnings struct {
	// MediaInReadOnly shall indicate the media has been placed in read-only mode.
	// This is not set when the read-only condition of the media is a result of a
	// change in the write protection state of a namespace.
	MediaInReadOnly bool
	// OverallSubsystemDegraded shall indicate that the NVM subsystem reliability
	// has been compromised.
	OverallSubsystemDegraded bool
	// PMRUnreliable shall indicate that the Persistent Memory Region has become
	// unreliable. PCIe memory reads can return invalid data or generate poisoned
	// PCIe TLP(s). Persistent Memory Region memory writes might not update memory
	// or might update memory with undefined data. The Persistent Memory Region
	// might also have become non-persistent.
	PMRUnreliable bool
	// PowerBackupFailed shall indicate that the volatile memory backup device has
	// failed.
	PowerBackupFailed bool
	// SpareCapacityWornOut shall indicate that the available spare capacity has
	// fallen below the threshold.
	SpareCapacityWornOut bool
}

// SecurityReceiveResponse shall contain the security data transferred from a
// controller.
type SecurityReceiveResponse struct {
	// Data shall contain a Base64-encoded string, with padding characters, of the
	// security protocol data transferred from a controller.
	//
	// Version added: v1.7.0
	Data string
}
