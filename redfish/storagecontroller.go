//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ANAAccessState string

const (
	// OptimizedANAAccessState Commands processed by a controller provide optimized access to any namespace in the ANA
	// group.
	OptimizedANAAccessState ANAAccessState = "Optimized"
	// NonOptimizedANAAccessState Commands processed by a controller that reports this state for an ANA group provide
	// non-optimized access characteristics, such as lower performance or non-optimal use of subsystem resources, to
	// any namespace in the ANA group.
	NonOptimizedANAAccessState ANAAccessState = "NonOptimized"
	// InaccessibleANAAccessState Namespaces in this group are inaccessible. Commands are not able to access user data
	// of namespaces in the ANA group.
	InaccessibleANAAccessState ANAAccessState = "Inaccessible"
	// PersistentLossANAAccessState The group is persistently inaccessible. Commands are persistently not able to
	// access user data of namespaces in the ANA group.
	PersistentLossANAAccessState ANAAccessState = "PersistentLoss"
)

type NVMeControllerType string

const (
	// AdminNVMeController NVMe controller is an admin controller.
	AdminNVMeControllerType NVMeControllerType = "Admin"
	// DiscoveryNVMeController NVMe controller is a discovery controller.
	DiscoveryNVMeControllerType NVMeControllerType = "Discovery"
	// IONVMeController NVMe controller is an I/O controller.
	IONVMeControllerType NVMeControllerType = "IO"
)

// ANACharacteristics shall contain the ANA characteristics and volume information for a storage controller.
type ANACharacteristics struct {
	// AccessState shall contain the reported ANA access state.
	AccessState ANAAccessState
	// Volume shall contain a link to a resource of type Volume.
	Volume string
}

// UnmarshalJSON unmarshals a ANACharacteristics object from the raw JSON.
func (anacharacteristics *ANACharacteristics) UnmarshalJSON(b []byte) error {
	type temp ANACharacteristics
	var t struct {
		temp
		Volume common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*anacharacteristics = ANACharacteristics(t.temp)

	// Extract the links to other entities for later
	anacharacteristics.Volume = t.Volume.String()

	return nil
}

// AttachDetachNamespacesResponse shall contain the properties found in the response body for the AttachNamespaces
// and DetachNamespaces actions.
type AttachDetachNamespacesResponse struct {
	// AttachedVolumes shall contain an array of links to resources of type Volume that are attached to this instance
	// of storage controller.
	AttachedVolumes []string
	// AttachedVolumesCount is the number of attached volumes.
	AttachedVolumesCount int `json:"AttachedVolumes@odata.count"`
}

// UnmarshalJSON unmarshals a AttachDetachNamespacesResponse object from the raw JSON.
func (attachdetachnamespacesresponse *AttachDetachNamespacesResponse) UnmarshalJSON(b []byte) error {
	type temp AttachDetachNamespacesResponse
	var t struct {
		temp
		AttachedVolumes common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*attachdetachnamespacesresponse = AttachDetachNamespacesResponse(t.temp)

	// Extract the links to other entities for later
	attachdetachnamespacesresponse.AttachedVolumes = t.AttachedVolumes.ToStrings()

	return nil
}

// NVMeControllerAttributes shall contain NVMe controller attributes for a storage controller.
type NVMeControllerAttributes struct {
	// ReportsNamespaceGranularity shall indicate whether or not the controller supports reporting of Namespace
	// Granularity.
	ReportsNamespaceGranularity bool
	// ReportsUUIDList shall indicate whether or not the controller supports reporting of a UUID list.
	ReportsUUIDList bool
	// Supports128BitHostID shall indicate whether or not the controller supports a 128-bit Host Identifier.
	Supports128BitHostID bool
	// SupportsEnduranceGroups shall indicate whether or not the controller supports Endurance Groups.
	SupportsEnduranceGroups bool
	// SupportsExceedingPowerOfNonOperationalState shall indicate whether or not the controller supports exceeding
	// Power of Non-Operational State in order to execute controller-initiated background operations in a non-
	// operational power state.
	SupportsExceedingPowerOfNonOperationalState bool
	// SupportsNVMSets shall indicate whether or not the controller supports NVM Sets.
	SupportsNVMSets bool
	// SupportsPredictableLatencyMode shall indicate whether or not the controller supports Predictable Latency Mode.
	SupportsPredictableLatencyMode bool
	// SupportsReadRecoveryLevels shall indicate whether or not the controller supports Read Recovery Levels.
	SupportsReadRecoveryLevels bool
	// SupportsReservations shall indicate if the controller supports reservations.
	SupportsReservations bool
	// SupportsSQAssociations shall indicate whether or not the controller supports SQ Associations.
	SupportsSQAssociations bool
	// SupportsTrafficBasedKeepAlive shall indicate whether or not the controller supports restarting the Keep Alive
	// Timer if traffic is processed from an admin command or I/O during a Keep Alive Timeout interval.
	SupportsTrafficBasedKeepAlive bool
}

// NVMeControllerProperties shall contain NVMe-related properties for a storage controller.
type NVMeControllerProperties struct {
	// ANACharacteristics shall contain the ANA characteristics and volume information.
	ANACharacteristics []ANACharacteristics
	// AllocatedCompletionQueues shall contain the number of I/O completion queues allocated to this NVMe I/O
	// controller.
	AllocatedCompletionQueues int
	// AllocatedSubmissionQueues shall contain the number of I/O submission queues allocated to this NVMe I/O
	// controller.
	AllocatedSubmissionQueues int
	// ControllerType shall contain the type of NVMe controller.
	ControllerType NVMeControllerType
	// MaxQueueSize shall contain the maximum individual queue entry size supported per queue. This is a zero-based
	// value, where the minimum value is one, indicating two entries. For PCIe, this applies to both submission and
	// completion queues. For NVMe-oF, this applies only to submission queues.
	MaxQueueSize int
	// NVMeControllerAttributes shall contain NVMe controller attributes.
	NVMeControllerAttributes NVMeControllerAttributes
	// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for this storage controller. This
	// property can contain possible triggers for the predictive drive failure warning for the corresponding drive.
	NVMeSMARTCriticalWarnings NVMeSMARTCriticalWarnings
	// NVMeVersion shall contain the version of the NVMe Base Specification supported.
	NVMeVersion string
}

// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for a storage controller.
type NVMeSMARTCriticalWarnings struct {
	// MediaInReadOnly shall indicate the media has been placed in read-only mode. This is not set when the read-only
	// condition of the media is a result of a change in the write protection state of a namespace.
	MediaInReadOnly bool
	// OverallSubsystemDegraded shall indicate that the NVM subsystem reliability has been compromised.
	OverallSubsystemDegraded bool
	// PMRUnreliable shall indicate that the Persistent Memory Region has become unreliable. PCIe memory reads can
	// return invalid data or generate poisoned PCIe TLP(s). Persistent Memory Region memory writes might not update
	// memory or might update memory with undefined data. The Persistent Memory Region might also have become non-
	// persistent.
	PMRUnreliable bool
	// PowerBackupFailed shall indicate that the volatile memory backup device has failed.
	PowerBackupFailed bool
	// SpareCapacityWornOut shall indicate that the available spare capacity has fallen below the threshold.
	SpareCapacityWornOut bool
}

// SecurityReceiveResponse shall contain the security data transferred from a controller.
type SecurityReceiveResponse struct {
	// Data shall contain a Base64-encoded string of the security protocol data transferred from a controller.
	Data string
}

// StorageController is used to represent a resource that represents a
// storage controller in the Redfish specification.
type StorageController struct {
	common.Entity

	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the storage controller for inventory
	// purposes.
	AssetTag string
	// CacheSummary shall contain properties which describe the cache memory for
	// the current resource.
	CacheSummary CacheSummary
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates string
	// ControllerRates shall contain all the rate settings available on the controller.
	ControllerRates Rates
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this storage controller.
	environmentMetrics string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated storage controller.
	FirmwareVersion string
	// Identifiers shall contain a list of all known durable names for the
	// associated storage controller.
	Identifiers []common.Identifier
	// Location shall contain location information of the
	// associated storage controller.
	Location common.Location
	// Manufacturer shall be the name of the organization responsible for
	// producing the storage controller. This organization might be the entity
	// from whom the storage controller is purchased, but this is not
	// necessarily true.
	Manufacturer string
	// Metrics shall contain a link to the metrics associated with this storage controller.
	metrics string
	// Model shall be the name by which the manufacturer generally refers to the
	// storage controller.
	Model string
	// NVMeControllerProperties shall contain NVMe-related properties for this storage controller.
	NVMeControllerProperties NVMeControllerProperties
	// PCIeInterface is used to connect this PCIe-based controller to its host.
	PCIeInterface PCIeInterface
	// PartNumber shall be a part number assigned by the organization that is
	// responsible for producing or manufacturing the storage controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// SKU shall be the stock-keeping unit number for this storage storage
	// controller.
	SKU string
	// SerialNumber is used to identify the storage controller.
	SerialNumber string
	// SpeedGbps shall represent the maximum supported speed of the Storage bus
	// interface (in Gigabits per second). The interface specified connects the
	// controller to the storage devices, not the controller to a host (e.g. SAS
	// bus, not PCIe host bus).
	SpeedGbps float32
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedControllerProtocols shall be the set of protocols this storage
	// controller can be communicated to.
	SupportedControllerProtocols []common.Protocol
	// SupportedDeviceProtocols shall be the set of protocols this storage
	// controller can use to communicate with attached devices.
	SupportedDeviceProtocols []common.Protocol
	// SupportedRAIDTypes shall contain all the RAIDType values supported by the
	// current resource.
	SupportedRAIDTypes []RAIDType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	attachedVolumes []string
	// AttachedVolumesCount is the number of attached volumes.
	AttachedVolumesCount int
	batteries            []string
	// BatteriesCount is the number of connected batteries.
	BatteriesCount int
	endpoints      []string
	// EndpointsCount is the number of enclosures.
	EndpointsCount           int
	nvmeDiscoveredSubsystems []string
	// NVMeDiscoveredSubsystemsCount is the number of discovered NVMe subsystems.
	NVMeDiscoveredSubsystemsCount int
	networkDeviceFunctions        []string
	// NetworkDeviceFunctionsCount is the number of network device functions.
	NetworkDeviceFunctionsCount int
	pcieFunctions               []string
	// PCIeFunctionsCount is the number of PCIeFunctions for this storage controller.
	PCIeFunctionsCount int
	// StorageServices shall be a reference to the resources that this
	// controller is associated with and shall reference a resource of type
	// StorageService.
	// This property has been deprecated in favor of StorageServices within the Links property at the root level.
	storageServices []string
	// StorageServicesCount is the number of storage services.
	// This property has been deprecated in favor of StorageServices within the Links property at the root level.
	StorageServicesCount int
}

// UnmarshalJSON unmarshals a StorageController object from the raw JSON.
func (storagecontroller *StorageController) UnmarshalJSON(b []byte) error {
	type temp StorageController
	type links struct {
		AttachedVolumes               common.Links
		AttachedVolumesCount          int `json:"AttachedVolumes@odata.count"`
		Batteries                     common.Links
		BatteriesCount                int `json:"Batteries@odata.count"`
		Endpoints                     common.Links
		EndpointsCount                int `json:"Endpoints@odata.count"`
		NVMeDiscoveredSubsystems      common.Links
		NVMeDiscoveredSubsystemsCount int `json:"NVMeDiscoveredSubsystems@odata.count"`
		NetworkDeviceFunctions        common.Links
		NetworkDeviceFunctionsCount   int `json:"NetworkDeviceFunctions@odata.count"`
		PCIeFunctions                 common.Links
		PCIeFunctionsCount            int `json:"PCIeFunctions@odata.count"`
		StorageServices               common.Links
		StorageServicesCount          int `json:"StorageServices@odata.count"`
	}
	var t struct {
		temp
		Assembly           common.Link
		Certificates       common.Link
		EnvironmentMetrics common.Link
		Metrics            common.Link
		Ports              common.Link
		Links              links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontroller = StorageController(t.temp)

	// Extract the links to other entities for later
	storagecontroller.assembly = t.Assembly.String()
	storagecontroller.certificates = t.Certificates.String()
	storagecontroller.environmentMetrics = t.EnvironmentMetrics.String()
	storagecontroller.metrics = t.Metrics.String()
	storagecontroller.ports = t.Ports.String()

	storagecontroller.attachedVolumes = t.Links.AttachedVolumes.ToStrings()
	storagecontroller.AttachedVolumesCount = t.Links.AttachedVolumesCount
	storagecontroller.batteries = t.Links.Batteries.ToStrings()
	storagecontroller.BatteriesCount = t.Links.BatteriesCount
	storagecontroller.endpoints = t.Links.Endpoints.ToStrings()
	storagecontroller.EndpointsCount = t.Links.EndpointsCount
	storagecontroller.nvmeDiscoveredSubsystems = t.Links.NVMeDiscoveredSubsystems.ToStrings()
	storagecontroller.NVMeDiscoveredSubsystemsCount = t.Links.NVMeDiscoveredSubsystemsCount
	storagecontroller.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	storagecontroller.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	storagecontroller.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	storagecontroller.PCIeFunctionsCount = t.Links.PCIeFunctionsCount
	storagecontroller.storageServices = t.Links.StorageServices.ToStrings()
	storagecontroller.StorageServicesCount = t.Links.StorageServicesCount

	// This is a read/write object, so we need to save the raw object data for later
	storagecontroller.rawData = b

	return nil
}

// Assembly gets the storage controller's assembly.
func (storagecontroller *StorageController) Assembly() (*Assembly, error) {
	if storagecontroller.assembly == "" {
		return nil, nil
	}
	return GetAssembly(storagecontroller.GetClient(), storagecontroller.assembly)
}

// Certificates gets the storage controller's certificates.
func (storagecontroller *StorageController) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(storagecontroller.GetClient(), storagecontroller.certificates)
}

// EnvironmentMetrics gets the environment metrics for this storage controller.
func (storagecontroller *StorageController) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if storagecontroller.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(storagecontroller.GetClient(), storagecontroller.environmentMetrics)
}

// Metrics gets the metrics associated with this storage controller.
func (storagecontroller *StorageController) Metrics() (*StorageControllerMetrics, error) {
	if storagecontroller.metrics == "" {
		return nil, nil
	}
	return GetStorageControllerMetrics(storagecontroller.GetClient(), storagecontroller.metrics)
}

// Ports gets the ports that exist on the storage controller.
func (storagecontroller *StorageController) Ports() ([]*Port, error) {
	return ListReferencedPorts(storagecontroller.GetClient(), storagecontroller.ports)
}

// // AttachedVolumes gets the volumes that are attached to this instance of storage controller.
// func (storagecontroller *StorageController) AttachedVolumes() ([]*swordfish.Volume, error) {
//  	return common.GetObjects[Volume](storagecontroller.GetClient(), storagecontroller.attachedVolumes)
// }

// Batteries gets the batteries that provide power to this storage controller during a power-loss event,
// such as with battery-backed RAID controllers. This property shall not be present if the batteries
// power the containing chassis as a whole rather than the individual storage controller.
func (storagecontroller *StorageController) Batteries() ([]*Battery, error) {
	return common.GetObjects[Battery](storagecontroller.GetClient(), storagecontroller.batteries)
}

// Endpoints gets the storage controller's endpoints.
func (storagecontroller *StorageController) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](storagecontroller.GetClient(), storagecontroller.endpoints)
}

// NVMeDiscoveredSubsystems gets the storage that represent the NVMe subsystems discovered by
// this discovery controller. This property shall only be present if ControllerType in
// NVMeControllerProperties contains 'Discovery'.
func (storagecontroller *StorageController) NVMeDiscoveredSubsystems() ([]*Storage, error) {
	return common.GetObjects[Storage](storagecontroller.GetClient(), storagecontroller.nvmeDiscoveredSubsystems)
}

// NetworkDeviceFunctions the network device functions that provide connectivity to this controller.
func (storagecontroller *StorageController) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return common.GetObjects[NetworkDeviceFunction](storagecontroller.GetClient(), storagecontroller.networkDeviceFunctions)
}

// PCIeFunctions gets the the PCIe functions that the storage controller produces.
func (storagecontroller *StorageController) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](storagecontroller.GetClient(), storagecontroller.pcieFunctions)
}

// Update commits updates to this object's properties to the running system.
func (storagecontroller *StorageController) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StorageController)
	err := original.UnmarshalJSON(storagecontroller.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagecontroller).Elem()

	return storagecontroller.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetStorageController will get a Storage controller instance from the service.
func GetStorageController(c common.Client, uri string) (*StorageController, error) {
	return common.GetObject[StorageController](c, uri)
}

// ListReferencedStorageControllers gets the collection of StorageControllers
// from a provided reference.
func ListReferencedStorageControllers(c common.Client, link string) ([]*StorageController, error) {
	return common.GetCollectionObjects[StorageController](c, link)
}
