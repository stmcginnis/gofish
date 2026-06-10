//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"encoding/json"
	"fmt"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
	"github.com/coreweave/gofish/swordfish"
)

var ErrSessionNotSupported = fmt.Errorf("session functionality is not supported")

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
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountService shall only contain a reference to a resource that complies
	// to the AccountService schema.
	AccountServiceLink common.Link `json:"AccountService"`
	// AggregationService shall contain a link to a resource of type AggregationService.
	AggregationServiceLink common.Link `json:"AggregationService"`
	// Cables shall contain a link to a resource collection of type CableCollection.
	CablesLink common.Link `json:"Cables"`
	// CertificateService shall be a link to the CertificateService.
	CertificateServiceLink common.Link `json:"CertificateService"`
	// Chassis shall only contain a reference to a collection of resources that
	// comply to the Chassis schema.
	ChassisLink common.Link `json:"Chassis"`
	// ComponentIntegrity shall contain a link to a resource collection of type ComponentIntegrityCollection.
	ComponentIntegrityLink common.Link `json:"ComponentIntegrity"`
	// CompositionService shall only contain a reference to a resource that
	// complies to the CompositionService schema.
	CompositionServiceLink common.Link `json:"CompositionService"`
	// Description provides a description of this resource.
	Description string
	// EventService shall contain a link to a resource of type EventService.
	EventServiceLink common.Link `json:"EventService"`
	// Fabrics shall contain references to all Fabric instances.
	FabricsLink common.Link `json:"Fabrics"`
	// Facilities shall contain a link to a resource collection of type FacilityCollection.
	FacilitiesLink common.Link `json:"Facilities"`
	// JobService shall only contain a reference to a resource that conforms to
	// the JobService schema.
	JobServiceLink common.Link `json:"JobService"`
	// JsonSchemas shall only contain a reference to a collection of resources
	// that comply to the SchemaFile schema where the files are Json-Schema
	// files.
	JSONSchemasLink common.Link `json:"JSONSchemas"`
	// KeyService shall contain a link to a resource of type KeyService.
	KeyServiceLink common.Link `json:"KeyService"`
	// LicenseService shall contain a link to a resource of type LicenseService.
	LicenseServiceLink common.Link `json:"LicenseService"`
	// Managers shall only contain a reference to a collection of resources that
	// comply to the Managers schema.
	ManagersLink common.Link `json:"Managers"`
	// NVMeDomains shall contain a link to a resource collection of type NVMeDomainCollection.
	NvmeDomainsLink common.Link `json:"NVMeDomains"`
	// Oem contains all the vendor specific actions. It is vendor responsibility to parse
	// this field accordingly
	Oem json.RawMessage
	// (v1.6+) PowerEquipment shall only contain a reference to a collection of resources that
	// comply to the PowerEquipment schema.
	PowerEquipmentLink common.Link `json:"PowerEquipment"`
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
	RegisteredClientsLink common.Link `json:"RegisteredClients"`
	// Registries shall contain a reference to Message Registry.
	RegistriesLink common.Link `json:"Registries"`
	// ResourceBlocks shall contain references to all Resource Block instances.
	ResourceBlocksLink common.Link `json:"ResourceBlocks"`
	// ServiceConditions shall contain a link to a resource of type ServiceConditions.
	ServiceConditionsLink common.Link `json:"ServiceConditions"`
	// ServiceIdentification shall contain a vendor-provided or user-provided value that identifies and associates a
	// discovered Redfish service with a particular product instance. The value of the property shall contain the value
	// of the ServiceIdentification property in the Manager resource providing the Redfish service root resource. The
	// value of this property is used in conjunction with the Product and Vendor properties to match user credentials
	// or other a priori product instance information necessary for initial deployment to the correct, matching Redfish
	// service. This property shall not be present if its value is an empty string or 'null'.
	ServiceIdentification string
	// SessionService shall only contain a reference to a resource that complies
	// to the SessionService schema.
	SessionServiceLink common.Link `json:"SessionService"`
	// Storage shall contain a link to a resource collection of type StorageCollection.
	StorageLink common.Link `json:"Storage"`
	// StorageServices shall contain references to all StorageService instances.
	StorageServicesLink common.Link `json:"StorageServices"`
	// StorageSystems shall contain computer systems that act as storage
	// servers. The HostingRoles attribute of each such computer system shall
	// have an entry for StorageServer.
	StorageSystemsLink common.Link `json:"StorageSystems"`
	// Systems shall only contain a reference to a collection of resources that
	// comply to the Systems schema.
	SystemsLink common.Link `json:"Systems"`
	// Tasks shall only contain a reference to a resource that complies to the
	// TaskService schema.
	TasksLink common.Link `json:"Tasks"`
	// TelemetryService shall be a link to the TelemetryService.
	TelemetryServiceLink common.Link `json:"TelemetryService"`
	// ThermalEquipment shall contain a link to a resource of type ThermalEquipment.
	ThermalEquipmentLink common.Link `json:"ThermalEquipment"`
	// UUID shall be an exact match of the UUID value returned in a 200OK from
	// an SSDP M-SEARCH request during discovery. RFC4122 describes methods that
	// can be used to create a UUID value. The value should be considered to be
	// opaque. Client software should only treat the overall value as a
	// universally unique identifier and should not interpret any sub-fields
	// within the UUID.
	UUID string
	// UpdateService shall only contain a reference to a resource that complies
	// to the UpdateService schema.
	UpdateServiceLink common.Link `json:"UpdateService"`
	// Vendor shall include the name of the manufacturer or vendor represented
	// by this Redfish service. If this property is supported, the vendor name
	// shall not be included in the value of the Product property.
	Vendor string

	Links struct {
		// ManagerProvidingService shall contain a link to a resource of type Manager that represents the manager providing
		// this Redfish service.

		ManagerProvidingService common.Link `json:"ManagerProvidingService"`
		// Sessions shall contain the link to a collection of Sessions.
		Sessions common.Link `json:"Sessions"`
	} `json:"Links"`
}

// ServiceRoot will get a Service instance from the service.
func ServiceRoot(c common.Client) (*Service, error) {
	resp, err := c.Get(common.DefaultServiceRoot)
	defer common.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

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
	if serviceroot.AccountServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetAccountService(serviceroot.GetClient(), serviceroot.AccountServiceLink.String())
}

// AggregationService gets the aggregation service.
func (serviceroot *Service) AggregationService() (*redfish.AggregationService, error) {
	if serviceroot.AggregationServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetAggregationService(serviceroot.GetClient(), serviceroot.AggregationServiceLink.String())
}

// Cables gets a collection of cables.
func (serviceroot *Service) Cables() ([]*redfish.Cable, error) {
	if serviceroot.CablesLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedCables(serviceroot.GetClient(), serviceroot.CablesLink.String())
}

// CertificateService gets the certificate service.
func (serviceroot *Service) CertificateService() (*redfish.CertificateService, error) {
	if serviceroot.CertificateServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetCertificateService(serviceroot.GetClient(), serviceroot.CertificateServiceLink.String())
}

// Chassis gets the chassis instances managed by this service.
func (serviceroot *Service) Chassis() ([]*redfish.Chassis, error) {
	if serviceroot.ChassisLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedChassis(serviceroot.GetClient(), serviceroot.ChassisLink.String())
}

// ComponentIntegrity gets a collection of cables.
func (serviceroot *Service) ComponentIntegrity() ([]*redfish.ComponentIntegrity, error) {
	if serviceroot.ComponentIntegrityLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedComponentIntegritys(serviceroot.GetClient(), serviceroot.ComponentIntegrityLink.String())
}

// CompositionService gets the composition service.
func (serviceroot *Service) CompositionService() (*redfish.CompositionService, error) {
	if serviceroot.CompositionServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetCompositionService(serviceroot.GetClient(), serviceroot.CompositionServiceLink.String())
}

// EventService gets the Redfish EventService
func (serviceroot *Service) EventService() (*redfish.EventService, error) {
	if serviceroot.EventServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetEventService(serviceroot.GetClient(), serviceroot.EventServiceLink.String())
}

// Fabrics gets a collection of fabrics.
func (serviceroot *Service) Fabrics() ([]*redfish.Fabric, error) {
	if serviceroot.FabricsLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedFabrics(serviceroot.GetClient(), serviceroot.FabricsLink.String())
}

// Facilities gets a collection of facilities.
func (serviceroot *Service) Facilities() ([]*redfish.Facility, error) {
	if serviceroot.FacilitiesLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedFacilities(serviceroot.GetClient(), serviceroot.FacilitiesLink.String())
}

// JobService gets the job service instance
func (serviceroot *Service) JobService() (*redfish.JobService, error) {
	if serviceroot.JobServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetJobService(serviceroot.GetClient(), serviceroot.JobServiceLink.String())
}

// KeyService gets the key service.
func (serviceroot *Service) KeyService() (*redfish.KeyService, error) {
	if serviceroot.KeyServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetKeyService(serviceroot.GetClient(), serviceroot.KeyServiceLink.String())
}

// LicenseService gets the license service.
func (serviceroot *Service) LicenseService() (*redfish.LicenseService, error) {
	if serviceroot.LicenseServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetLicenseService(serviceroot.GetClient(), serviceroot.LicenseServiceLink.String())
}

// Managers gets the manager instances of this service.
func (serviceroot *Service) Managers() ([]*redfish.Manager, error) {
	if serviceroot.ManagersLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedManagers(serviceroot.GetClient(), serviceroot.ManagersLink.String())
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
	if serviceroot.RegisteredClientsLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedRegisteredClients(serviceroot.GetClient(), serviceroot.RegisteredClientsLink.String())
}

// Registries gets the Redfish Registries
func (serviceroot *Service) Registries() ([]*redfish.MessageRegistryFile, error) {
	if serviceroot.RegistriesLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedMessageRegistryFiles(serviceroot.GetClient(), serviceroot.RegistriesLink.String())
}

// ResourceBlocks gets a collection of resource blocks.
func (serviceroot *Service) ResourceBlocks() ([]*redfish.ResourceBlock, error) {
	if serviceroot.ResourceBlocksLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedResourceBlocks(serviceroot.GetClient(), serviceroot.ResourceBlocksLink.String())
}

// ServiceConditions gets the service conditions.
func (serviceroot *Service) ServiceConditions() (*redfish.ServiceConditions, error) {
	if serviceroot.ServiceConditionsLink.IsZero() {
		return nil, nil
	}
	return redfish.GetServiceConditions(serviceroot.GetClient(), serviceroot.ServiceConditionsLink.String())
}

// SessionService gets the session service.
func (serviceroot *Service) SessionService() (*redfish.SessionService, error) {
	if serviceroot.SessionServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetSessionService(serviceroot.GetClient(), serviceroot.SessionServiceLink.String())
}

// Storage gets a collection of storage objects.
func (serviceroot *Service) Storage() ([]*redfish.Storage, error) {
	if serviceroot.StorageLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedStorages(serviceroot.GetClient(), serviceroot.StorageLink.String())
}

// StorageServices gets the Swordfish storage services
func (serviceroot *Service) StorageServices() ([]*swordfish.StorageService, error) {
	if serviceroot.StorageServicesLink.IsZero() {
		return nil, nil
	}
	return swordfish.ListReferencedStorageServices(serviceroot.GetClient(), serviceroot.StorageServicesLink.String())
}

// StorageSystems gets the storage system instances managed by this service.
func (serviceroot *Service) StorageSystems() ([]*swordfish.StorageSystem, error) {
	if serviceroot.StorageSystemsLink.IsZero() {
		return nil, nil
	}
	return swordfish.ListReferencedStorageSystems(serviceroot.GetClient(), serviceroot.StorageSystemsLink.String())
}

// Tasks gets the system's tasks
func (serviceroot *Service) Tasks() ([]*redfish.Task, error) {
	if serviceroot.TasksLink.IsZero() {
		return nil, nil
	}
	ts, err := redfish.GetTaskService(serviceroot.GetClient(), serviceroot.TasksLink.String())
	if err != nil {
		return nil, err
	}

	return ts.Tasks()
}

// TaskService gets the task service instance
func (serviceroot *Service) TaskService() (*redfish.TaskService, error) {
	if serviceroot.TasksLink.IsZero() {
		return nil, nil
	}
	return redfish.GetTaskService(serviceroot.GetClient(), serviceroot.TasksLink.String())
}

// CreateSession creates a new session and returns the token and id
func (serviceroot *Service) CreateSession(username, password string) (*redfish.AuthToken, error) {
	if serviceroot.Links.Sessions.IsZero() {
		return nil, ErrSessionNotSupported
	}

	return redfish.CreateSession(serviceroot.GetClient(), serviceroot.Links.Sessions.String(), username, password)
}

// ManagerProvidingService gets the manager for this Redfish service.
func (serviceroot *Service) ManagerProvidingService() (*redfish.Manager, error) {
	if serviceroot.Links.ManagerProvidingService.IsZero() {
		return nil, nil
	}
	return redfish.GetManager(serviceroot.GetClient(), serviceroot.Links.ManagerProvidingService.String())
}

// PowerEquipment gets the powerEquipment instances of this service.
func (serviceroot *Service) PowerEquipment() (*redfish.PowerEquipment, error) {
	if serviceroot.PowerEquipmentLink.IsZero() {
		return nil, nil
	}
	return redfish.GetPowerEquipment(serviceroot.GetClient(), serviceroot.PowerEquipmentLink.String())
}

// Sessions gets the system's active sessions
func (serviceroot *Service) Sessions() ([]*redfish.Session, error) {
	if serviceroot.Links.Sessions.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedSessions(serviceroot.GetClient(), serviceroot.Links.Sessions.String())
}

// DeleteSession logout the specified session
func (serviceroot *Service) DeleteSession(url string) error {
	return redfish.DeleteSession(serviceroot.GetClient(), url)
}

// MessageRegistries gets all the available message registries in all languages
func (serviceroot *Service) MessageRegistries() ([]*redfish.MessageRegistry, error) {
	if serviceroot.RegistriesLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedMessageRegistries(serviceroot.GetClient(), serviceroot.RegistriesLink.String())
}

// MessageRegistry gets a specific message registry.
// uri is the uri for the message registry
func (serviceroot *Service) MessageRegistry(uri string) (*redfish.MessageRegistry, error) {
	return redfish.GetMessageRegistry(serviceroot.GetClient(), uri)
}

// MessageRegistriesByLanguage gets the message registries by language.
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistriesByLanguage(language string) ([]*redfish.MessageRegistry, error) {
	if serviceroot.RegistriesLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedMessageRegistriesByLanguage(serviceroot.GetClient(), serviceroot.RegistriesLink.String(), language)
}

// MessageRegistryByLanguage gets a specific message registry by language.
// registry is used to identify the correct Message Registry file and it shall
// contain the Message Registry name and it major and minor versions, as defined
// by the Redfish Specification, for example: "Alert.1.0.0".
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistryByLanguage(registry, language string) (*redfish.MessageRegistry, error) {
	if serviceroot.RegistriesLink.IsZero() {
		return nil, nil
	}
	return redfish.GetMessageRegistryByLanguage(serviceroot.GetClient(), serviceroot.RegistriesLink.String(), registry, language)
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
	if serviceroot.RegistriesLink.IsZero() {
		return nil, nil
	}
	return redfish.GetMessageFromMessageRegistryByLanguage(serviceroot.GetClient(), serviceroot.RegistriesLink.String(), messageID, language)
}

// Systems get the system instances from the service
func (serviceroot *Service) Systems() ([]*redfish.ComputerSystem, error) {
	if serviceroot.SystemsLink.IsZero() {
		return nil, nil
	}
	return redfish.ListReferencedComputerSystems(serviceroot.GetClient(), serviceroot.SystemsLink.String())
}

// TelemetryService gets the telemetry service instance.
func (serviceroot *Service) TelemetryService() (*redfish.TelemetryService, error) {
	if serviceroot.TelemetryServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetTelemetryService(serviceroot.GetClient(), serviceroot.TelemetryServiceLink.String())
}

// ThermalEquipment gets the thermal equipment instance.
func (serviceroot *Service) ThermalEquipment() (*redfish.ThermalEquipment, error) {
	if serviceroot.ThermalEquipmentLink.IsZero() {
		return nil, nil
	}
	return redfish.GetThermalEquipment(serviceroot.GetClient(), serviceroot.ThermalEquipmentLink.String())
}

// UpdateService gets the update service instance
func (serviceroot *Service) UpdateService() (*redfish.UpdateService, error) {
	if serviceroot.UpdateServiceLink.IsZero() {
		return nil, nil
	}
	return redfish.GetUpdateService(serviceroot.GetClient(), serviceroot.UpdateServiceLink.String())
}
