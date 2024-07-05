//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
	"github.com/stmcginnis/gofish/swordfish"
)

// DeepOperations shall contain information about deep operations that the service supports.
type DeepOperations struct {
	// DeepPATCH shall indicate whether this service supports the Redfish Specification-defined deep PATCH operation.
	DeepPATCH bool
	// DeepPOST shall indicate whether this service supports the Redfish Specification-defined deep POST operation.
	DeepPOST bool
	// MaxLevels shall contain the maximum levels of resources allowed in deep operations.
	MaxLevels int
}

// Expand shall contain information about the support of the $expand query
// parameter by the service.
type Expand struct {
	// ExpandAll shall be a boolean indicating whether this service supports the
	// use of asterisk (expand all entries) as a value for the $expand query
	// parameter as described by the specification.
	ExpandAll bool
	// Levels shall be a boolean indicating whether this service supports the
	// use of $levels as a value for the $expand query parameter as described by
	// the specification.
	Levels bool
	// Links shall be a boolean indicating whether this service supports the use
	// of tilde (expand only entries in the Links section) as a value for the
	// $expand query parameter as described by the specification.
	Links bool
	// MaxLevels shall be the maximum value of the $levels qualifier supported
	// by the service and shall only be included if the value of the Levels
	// property is true.
	MaxLevels int
	// NoLinks shall be a boolean indicating whether this service supports the
	// use of period (expand only entries not in the Links section) as a value
	// for the $expand query parameter as described by the specification.
	NoLinks bool
}

// ProtocolFeaturesSupported contains information about protocol features
// supported by the service.
type ProtocolFeaturesSupported struct {
	// ExcerptQuery shall be a boolean indicating whether this service supports
	// the use of the 'excerpt' query parameter as described by the
	// specification.
	ExcerptQuery bool
	// ExpandQuery shall contain information about the support of the $expand
	// query parameter by the service.
	ExpandQuery Expand
	// FilterQuery shall be a boolean indicating whether this service supports
	// the use of the $filter query parameter as described by the specification.
	FilterQuery bool
	// OnlyMemberQuery shall be a boolean indicating whether this service
	// supports the use of the 'only' query parameter as described by the
	// specification.
	OnlyMemberQuery bool
	// SelectQuery shall be a boolean indicating whether this service supports
	// the use of the $select query parameter as described by the specification.
	SelectQuery bool
}

// Service represents the root Redfish service. All values for resources
// described by this schema shall comply to the requirements as described in the
// Redfish specification.
type Service struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountService shall only contain a reference to a resource that complies
	// to the AccountService schema.
	accountService string
	// AggregationService shall contain a link to a resource of type AggregationService.
	aggregationService string
	// Cables shall contain a link to a resource collection of type CableCollection.
	cables string
	// CertificateService shall be a link to the CertificateService.
	certificateService string
	// Chassis shall only contain a reference to a collection of resources that
	// comply to the Chassis schema.
	chassis string
	// ComponentIntegrity shall contain a link to a resource collection of type ComponentIntegrityCollection.
	componentIntegrity string
	// CompositionService shall only contain a reference to a resource that
	// complies to the CompositionService schema.
	compositionService string
	// Description provides a description of this resource.
	Description string
	// EventService shall contain a link to a resource of type EventService.
	eventService string
	// Fabrics shall contain references to all Fabric instances.
	fabrics string
	// Facilities shall contain a link to a resource collection of type FacilityCollection.
	facilities string
	// JobService shall only contain a reference to a resource that conforms to
	// the JobService schema.
	jobService string
	// JsonSchemas shall only contain a reference to a collection of resources
	// that comply to the SchemaFile schema where the files are Json-Schema
	// files.
	jsonSchemas string
	// KeyService shall contain a link to a resource of type KeyService.
	keyService string
	// LicenseService shall contain a link to a resource of type LicenseService.
	licenseService string
	// Managers shall only contain a reference to a collection of resources that
	// comply to the Managers schema.
	managers string
	// NVMeDomains shall contain a link to a resource collection of type NVMeDomainCollection.
	nvmeDomains string
	// Oem contains all the vendor specific actions. It is vendor responsibility to parse
	// this field accordingly
	Oem json.RawMessage
	// (v1.6+) PowerEquipment shall only contain a reference to a collection of resources that
	// comply to the PowerEquipment schema.
	powerEquipment string
	// Product shall include the name of the product represented by this Redfish
	// service.
	Product string
	// ProtocolFeaturesSupported contains information about protocol features
	// supported by the service.
	ProtocolFeaturesSupported ProtocolFeaturesSupported
	// RedfishVersion shall represent the version of the Redfish service. The
	// format of this string shall be of the format
	// majorversion.minorversion.errata in compliance with Protocol Version
	// section of the Redfish specification.
	RedfishVersion string
	// RegisteredClients shall contain a link to a resource collection of type RegisteredClientCollection.
	registeredClients string
	// Registries shall contain a reference to Message Registry.
	registries string
	// ResourceBlocks shall contain references to all Resource Block instances.
	resourceBlocks string
	// ServiceConditions shall contain a link to a resource of type ServiceConditions.
	serviceConditions string
	// ServiceIdentification shall contain a vendor-provided or user-provided value that identifies and associates a
	// discovered Redfish service with a particular product instance. The value of the property shall contain the value
	// of the ServiceIdentification property in the Manager resource providing the Redfish service root resource. The
	// value of this property is used in conjunction with the Product and Vendor properties to match user credentials
	// or other a priori product instance information necessary for initial deployment to the correct, matching Redfish
	// service. This property shall not be present if its value is an empty string or 'null'.
	ServiceIdentification string
	// SessionService shall only contain a reference to a resource that complies
	// to the SessionService schema.
	sessionService string
	// Storage shall contain a link to a resource collection of type StorageCollection.
	storage string
	// StorageServices shall contain references to all StorageService instances.
	storageServices string
	// StorageSystems shall contain computer systems that act as storage
	// servers. The HostingRoles attribute of each such computer system shall
	// have an entry for StorageServer.
	storageSystems string
	// Systems shall only contain a reference to a collection of resources that
	// comply to the Systems schema.
	systems string
	// Tasks shall only contain a reference to a resource that complies to the
	// TaskService schema.
	tasks string
	// TelemetryService shall be a link to the TelemetryService.
	telemetryService string
	// ThermalEquipment shall contain a link to a resource of type ThermalEquipment.
	thermalEquipment string
	// UUID shall be an exact match of the UUID value returned in a 200OK from
	// an SSDP M-SEARCH request during discovery. RFC4122 describes methods that
	// can be used to create a UUID value. The value should be considered to be
	// opaque. Client software should only treat the overall value as a
	// universally unique identifier and should not interpret any sub-fields
	// within the UUID.
	UUID string
	// UpdateService shall only contain a reference to a resource that complies
	// to the UpdateService schema.
	updateService string
	// Vendor shall include the name of the manufacturer or vendor represented
	// by this Redfish service. If this property is supported, the vendor name
	// shall not be included in the value of the Product property.
	Vendor string

	// Sessions shall contain the link to a collection of Sessions.
	sessions string
	// ManagerProvidingService shall contain a link to a resource of type Manager that represents the manager providing
	// this Redfish service.
	managerProvidingService string
}

// UnmarshalJSON unmarshals a Service object from the raw JSON.
func (serviceroot *Service) UnmarshalJSON(b []byte) error {
	type temp Service
	var t struct {
		temp
		AccountService     common.Link
		AggregationService common.Link
		Cables             common.Link
		CertificateService common.Link
		Chassis            common.Link
		ComponentIntegrity common.Link
		CompositionService common.Link
		EventService       common.Link
		Fabrics            common.Link
		Facilities         common.Link
		JobService         common.Link
		JSONSchemas        common.Link
		KeyService         common.Link
		LicenseService     common.Link
		Managers           common.Link
		NVMeDomains        common.Link
		PowerEquipment     common.Link
		Registries         common.Link
		RegisteredClients  common.Link
		ResourceBlocks     common.Link
		ServiceConditions  common.Link
		SessionService     common.Link
		Storage            common.Link
		StorageServices    common.Link
		StorageSystems     common.Link
		Systems            common.Link
		Tasks              common.Link
		TelemetryService   common.Link
		ThermalEquipment   common.Link
		UpdateService      common.Link
		Links              struct {
			ManagerProvidingService common.Link
			Sessions                common.Link
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*serviceroot = Service(t.temp)
	serviceroot.accountService = t.AccountService.String()
	serviceroot.aggregationService = t.AggregationService.String()
	serviceroot.cables = t.Cables.String()
	serviceroot.certificateService = t.CertificateService.String()
	serviceroot.chassis = t.Chassis.String()
	serviceroot.componentIntegrity = t.ComponentIntegrity.String()
	serviceroot.compositionService = t.CompositionService.String()
	serviceroot.eventService = t.EventService.String()
	serviceroot.fabrics = t.Fabrics.String()
	serviceroot.facilities = t.Facilities.String()
	serviceroot.jobService = t.JobService.String()
	serviceroot.jsonSchemas = t.JSONSchemas.String()
	serviceroot.keyService = t.KeyService.String()
	serviceroot.licenseService = t.LicenseService.String()
	serviceroot.managers = t.Managers.String()
	serviceroot.nvmeDomains = t.NVMeDomains.String()
	serviceroot.powerEquipment = t.PowerEquipment.String()
	serviceroot.registeredClients = t.RegisteredClients.String()
	serviceroot.registries = t.Registries.String()
	serviceroot.resourceBlocks = t.ResourceBlocks.String()
	serviceroot.serviceConditions = t.ServiceConditions.String()
	serviceroot.sessionService = t.SessionService.String()
	serviceroot.storage = t.Storage.String()
	serviceroot.storageServices = t.StorageServices.String()
	serviceroot.storageSystems = t.StorageSystems.String()
	serviceroot.systems = t.Systems.String()
	serviceroot.tasks = t.Tasks.String()
	serviceroot.telemetryService = t.TelemetryService.String()
	serviceroot.thermalEquipment = t.ThermalEquipment.String()
	serviceroot.updateService = t.UpdateService.String()

	serviceroot.sessions = t.Links.Sessions.String()
	serviceroot.managerProvidingService = t.Links.ManagerProvidingService.String()

	return nil
}

// ServiceRoot will get a Service instance from the service.
func ServiceRoot(c common.Client) (*Service, error) {
	resp, err := c.Get(common.DefaultServiceRoot)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var serviceroot Service
	err = json.NewDecoder(resp.Body).Decode(&serviceroot)
	if err != nil {
		return nil, err
	}

	serviceroot.SetClient(c)
	return &serviceroot, nil
}

// AccountService gets the Redfish AccountService
func (serviceroot *Service) AccountService() (*redfish.AccountService, error) {
	return redfish.GetAccountService(serviceroot.GetClient(), serviceroot.accountService)
}

// AggregationService gets the aggregation service.
func (serviceroot *Service) AggregationService() (*redfish.AggregationService, error) {
	if serviceroot.aggregationService == "" {
		return nil, nil
	}
	return redfish.GetAggregationService(serviceroot.GetClient(), serviceroot.aggregationService)
}

// Cables gets a collection of cables.
func (serviceroot *Service) Cables() ([]*redfish.Cable, error) {
	return redfish.ListReferencedCables(serviceroot.GetClient(), serviceroot.cables)
}

// CertificateService gets the certificate service.
func (serviceroot *Service) CertificateService() (*redfish.CertificateService, error) {
	if serviceroot.certificateService == "" {
		return nil, nil
	}
	return redfish.GetCertificateService(serviceroot.GetClient(), serviceroot.certificateService)
}

// Chassis gets the chassis instances managed by this service.
func (serviceroot *Service) Chassis() ([]*redfish.Chassis, error) {
	return redfish.ListReferencedChassis(serviceroot.GetClient(), serviceroot.chassis)
}

// ComponentIntegrity gets a collection of cables.
func (serviceroot *Service) ComponentIntegrity() ([]*redfish.ComponentIntegrity, error) {
	return redfish.ListReferencedComponentIntegritys(serviceroot.GetClient(), serviceroot.componentIntegrity)
}

// CompositionService gets the composition service.
func (serviceroot *Service) CompositionService() (*redfish.CompositionService, error) {
	if serviceroot.compositionService == "" {
		return nil, nil
	}
	return redfish.GetCompositionService(serviceroot.GetClient(), serviceroot.compositionService)
}

// EventService gets the Redfish EventService
func (serviceroot *Service) EventService() (*redfish.EventService, error) {
	return redfish.GetEventService(serviceroot.GetClient(), serviceroot.eventService)
}

// Fabrics gets a collection of fabrics.
func (serviceroot *Service) Fabrics() ([]*redfish.Fabric, error) {
	return redfish.ListReferencedFabrics(serviceroot.GetClient(), serviceroot.fabrics)
}

// Facilities gets a collection of facilities.
func (serviceroot *Service) Facilities() ([]*redfish.Facility, error) {
	return redfish.ListReferencedFacilities(serviceroot.GetClient(), serviceroot.facilities)
}

// JobService gets the job service instance
func (serviceroot *Service) JobService() (*redfish.JobService, error) {
	return redfish.GetJobService(serviceroot.GetClient(), serviceroot.jobService)
}

// KeyService gets the key service.
func (serviceroot *Service) KeyService() (*redfish.KeyService, error) {
	if serviceroot.keyService == "" {
		return nil, nil
	}
	return redfish.GetKeyService(serviceroot.GetClient(), serviceroot.keyService)
}

// LicenseService gets the license service.
func (serviceroot *Service) LicenseService() (*redfish.LicenseService, error) {
	if serviceroot.licenseService == "" {
		return nil, nil
	}
	return redfish.GetLicenseService(serviceroot.GetClient(), serviceroot.licenseService)
}

// Managers gets the manager instances of this service.
func (serviceroot *Service) Managers() ([]*redfish.Manager, error) {
	return redfish.ListReferencedManagers(serviceroot.GetClient(), serviceroot.managers)
}

// // NVMeDomains gets a collection of Swordfish NVMe domains.
// func (serviceroot *Service) NVMeDomains() ([]*swordfish.NVMeDomain, error) {
// 	var result []*swordfish.NVMeDomain

// 	collectionError := common.NewCollectionError()
// 	for _, uri := range serviceroot.nvmeDomains {
// 		item, err := swordfish.GetNVMeDomain(serviceroot.GetClient(), uri)
// 		if err != nil {
// 			collectionError.Failures[uri] = err
// 		} else {
// 			result = append(result, item)
// 		}
// 	}

// 	if collectionError.Empty() {
// 		return result, nil
// 	}

// 	return result, collectionError
// }

// RegisteredClients gets a collection of registered clients.
func (serviceroot *Service) RegisteredClients() ([]*redfish.RegisteredClient, error) {
	return redfish.ListReferencedRegisteredClients(serviceroot.GetClient(), serviceroot.registeredClients)
}

// Registries gets the Redfish Registries
func (serviceroot *Service) Registries() ([]*redfish.MessageRegistryFile, error) {
	return redfish.ListReferencedMessageRegistryFiles(serviceroot.GetClient(), serviceroot.registries)
}

// ResourceBlocks gets a collection of resource blocks.
func (serviceroot *Service) ResourceBlocks() ([]*redfish.ResourceBlock, error) {
	return redfish.ListReferencedResourceBlocks(serviceroot.GetClient(), serviceroot.resourceBlocks)
}

// ServiceConditions gets the service conditions.
func (serviceroot *Service) ServiceConditions() (*redfish.ServiceConditions, error) {
	if serviceroot.serviceConditions == "" {
		return nil, nil
	}
	return redfish.GetServiceConditions(serviceroot.GetClient(), serviceroot.serviceConditions)
}

// SessionService gets the session service.
func (serviceroot *Service) SessionService() (*redfish.SessionService, error) {
	if serviceroot.sessionService == "" {
		return nil, nil
	}
	return redfish.GetSessionService(serviceroot.GetClient(), serviceroot.sessionService)
}

// Storage gets a collection of storage objects.
func (serviceroot *Service) Storage() ([]*redfish.Storage, error) {
	return redfish.ListReferencedStorages(serviceroot.GetClient(), serviceroot.storage)
}

// StorageServices gets the Swordfish storage services
func (serviceroot *Service) StorageServices() ([]*swordfish.StorageService, error) {
	return swordfish.ListReferencedStorageServices(serviceroot.GetClient(), serviceroot.storageServices)
}

// StorageSystems gets the storage system instances managed by this service.
func (serviceroot *Service) StorageSystems() ([]*swordfish.StorageSystem, error) {
	return swordfish.ListReferencedStorageSystems(serviceroot.GetClient(), serviceroot.storageSystems)
}

// Tasks gets the system's tasks
func (serviceroot *Service) Tasks() ([]*redfish.Task, error) {
	ts, err := redfish.GetTaskService(serviceroot.GetClient(), serviceroot.tasks)
	if err != nil {
		return nil, err
	}

	return ts.Tasks()
}

// TaskService gets the task service instance
func (serviceroot *Service) TaskService() (*redfish.TaskService, error) {
	return redfish.GetTaskService(serviceroot.GetClient(), serviceroot.tasks)
}

// CreateSession creates a new session and returns the token and id
func (serviceroot *Service) CreateSession(username, password string) (*redfish.AuthToken, error) {
	return redfish.CreateSession(serviceroot.GetClient(), serviceroot.sessions, username, password)
}

// ManagerProvidingService gets the manager for this Redfish service.
func (serviceroot *Service) ManagerProvidingService() (*redfish.Manager, error) {
	if serviceroot.managerProvidingService == "" {
		return nil, nil
	}
	return redfish.GetManager(serviceroot.GetClient(), serviceroot.managerProvidingService)
}

// PowerEquipment gets the powerEquipment instances of this service.
func (serviceroot *Service) PowerEquipment() (*redfish.PowerEquipment, error) {
	return redfish.GetPowerEquipment(serviceroot.GetClient(), serviceroot.powerEquipment)
}

// Sessions gets the system's active sessions
func (serviceroot *Service) Sessions() ([]*redfish.Session, error) {
	return redfish.ListReferencedSessions(serviceroot.GetClient(), serviceroot.sessions)
}

// DeleteSession logout the specified session
func (serviceroot *Service) DeleteSession(url string) error {
	return redfish.DeleteSession(serviceroot.GetClient(), url)
}

// MessageRegistries gets all the available message registries in all languages
func (serviceroot *Service) MessageRegistries() ([]*redfish.MessageRegistry, error) {
	return redfish.ListReferencedMessageRegistries(serviceroot.GetClient(), serviceroot.registries)
}

// MessageRegistry gets a specific message registry.
// uri is the uri for the message registry
func (serviceroot *Service) MessageRegistry(uri string) (*redfish.MessageRegistry, error) {
	return redfish.GetMessageRegistry(serviceroot.GetClient(), uri)
}

// MessageRegistriesByLanguage gets the message registries by language.
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistriesByLanguage(language string) ([]*redfish.MessageRegistry, error) {
	return redfish.ListReferencedMessageRegistriesByLanguage(serviceroot.GetClient(), serviceroot.registries, language)
}

// MessageRegistryByLanguage gets a specific message registry by language.
// registry is used to identify the correct Message Registry file and it shall
// contain the Message Registry name and it major and minor versions, as defined
// by the Redfish Specification, for example: "Alert.1.0.0".
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistryByLanguage(registry, language string) (*redfish.MessageRegistry, error) {
	return redfish.GetMessageRegistryByLanguage(serviceroot.GetClient(), serviceroot.registries, registry, language)
}

// MessageByLanguage tries to find and get the message in the correct language from the informed messageID.
// messageID is the key used to find the registry, version and message, for example: "Alert.1.0.LanDisconnect"
//
//   - The segment before the 1st period is the Registry Name (Registry Prefix): Alert
//   - The segment between the 1st and 2nd period is the major version: 1
//   - The segment between the 2nd and 3rd period is the minor version: 0
//   - The segment after the 3rd period is the Message Identifier in the Registry: LanDisconnect
//
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageByLanguage(messageID, language string) (*redfish.MessageRegistryMessage, error) {
	return redfish.GetMessageFromMessageRegistryByLanguage(serviceroot.GetClient(), serviceroot.registries, messageID, language)
}

// Systems get the system instances from the service
func (serviceroot *Service) Systems() ([]*redfish.ComputerSystem, error) {
	return redfish.ListReferencedComputerSystems(serviceroot.GetClient(), serviceroot.systems)
}

// TelemetryService gets the telemetry service instance.
func (serviceroot *Service) TelemetryService() (*redfish.TelemetryService, error) {
	return redfish.GetTelemetryService(serviceroot.GetClient(), serviceroot.telemetryService)
}

// ThermalEquipment gets the thermal equipment instance.
func (serviceroot *Service) ThermalEquipment() (*redfish.ThermalEquipment, error) {
	return redfish.GetThermalEquipment(serviceroot.GetClient(), serviceroot.thermalEquipment)
}

// UpdateService gets the update service instance
func (serviceroot *Service) UpdateService() (*redfish.UpdateService, error) {
	return redfish.GetUpdateService(serviceroot.GetClient(), serviceroot.updateService)
}
