//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #ServiceRoot.v1_20_0.ServiceRoot

package gofish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/schemas"
)

// Service shall represent the root of the Redfish service.
type Service struct {
	schemas.Entity
	// AccountService shall contain a link to a resource of type 'AccountService'.
	accountService string
	// AggregationService shall contain a link to a resource of type
	// 'AggregationService'.
	//
	// Version added: v1.8.0
	aggregationService string
	// AutomationNodes shall contain a link to a resource collection of type
	// 'AutomationNodeCollection'.
	//
	// Version added: v1.19.0
	automationNodes string
	// Cables shall contain a link to a resource collection of type
	// 'CableCollection'.
	//
	// Version added: v1.11.0
	cables string
	// CertificateService shall contain a link to a resource of type
	// 'CertificateService'.
	//
	// Version added: v1.5.0
	certificateService string
	// Chassis shall contain a link to a resource collection of type
	// 'ChassisCollection'.
	chassis string
	// ComponentIntegrity shall contain a link to a resource collection of type
	// 'ComponentIntegrityCollection'.
	//
	// Version added: v1.13.0
	componentIntegrity string
	// CompositionService shall contain a link to a resource of type
	// 'CompositionService'.
	//
	// Version added: v1.2.0
	compositionService string
	// EventService shall contain a link to a resource of type 'EventService'.
	eventService string
	// Fabrics shall contain a link to a resource collection of type
	// 'FabricCollection'.
	//
	// Version added: v1.1.0
	fabrics string
	// Facilities shall contain a link to a resource collection of type
	// 'FacilityCollection'.
	//
	// Version added: v1.6.0
	facilities string
	// JSONSchemas shall contain a link to a resource collection of type
	// 'JsonSchemaFileCollection'.
	jSONSchemas string
	// JobService shall contain a link to a resource of type 'JobService'.
	//
	// Version added: v1.4.0
	jobService string
	// KeyService shall contain a link to a resource of type 'KeyService'.
	//
	// Version added: v1.11.0
	keyService string
	// LicenseService shall contain a link to a resource of type 'LicenseService'.
	//
	// Version added: v1.12.0
	licenseService string
	// Managers shall contain a link to a resource collection of type
	// 'ManagerCollection'.
	managers string
	// NVMeDomains shall contain a link to a resource collection of type
	// 'NVMeDomainCollection'.
	//
	// Version added: v1.10.0
	nVMeDomains string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerEquipment shall contain a link to a resource of type 'PowerEquipment'.
	//
	// Version added: v1.6.0
	powerEquipment string
	// Product shall include the name of the product represented by this Redfish
	// service.
	//
	// Version added: v1.3.0
	Product string
	// ProtocolFeaturesSupported shall contain information about protocol features
	// that the service supports.
	//
	// Version added: v1.3.0
	ProtocolFeaturesSupported ProtocolFeaturesSupported
	// RedfishVersion shall represent the Redfish protocol version, as specified in
	// the 'Protocol version' clause of the Redfish Specification, to which this
	// service conforms.
	RedfishVersion string
	// RegisteredClients shall contain a link to a resource collection of type
	// 'RegisteredClientCollection'.
	//
	// Version added: v1.13.0
	registeredClients string
	// Registries shall contain a link to a resource collection of type
	// 'MessageRegistryFileCollection'.
	registries string
	// ResourceBlocks shall contain a link to a resource collection of type
	// 'ResourceBlockCollection'.
	//
	// Version added: v1.5.0
	resourceBlocks string
	// ServiceConditions shall contain a link to a resource of type
	// 'ServiceConditions'.
	//
	// Version added: v1.13.0
	serviceConditions string
	// ServiceIdentification shall contain a vendor-provided or user-provided value
	// that identifies and associates a discovered Redfish service with a
	// particular product instance. The value of the property shall contain the
	// value of the 'ServiceIdentification' property in the 'Manager' resource
	// providing the Redfish service root resource. The value of this property is
	// used in conjunction with the 'Product' and 'Vendor' properties to match user
	// credentials or other a priori product instance information necessary for
	// initial deployment to the correct, matching Redfish service. This property
	// shall not be present if the value of the 'ServiceIdentification' property in
	// the 'Manager' resource providing the Redfish service root resource is an
	// empty string or 'null'.
	//
	// Version added: v1.14.0
	ServiceIdentification string
	// ServiceUseNotification shall contain the usage notification message for this
	// service. The value of the property shall contain the value of the
	// 'ServiceUseNotification' property in the 'Manager' resource providing the
	// Redfish service root resource. This property shall not be present if the
	// value of the 'ServiceUseNotification' property in the 'Manager' resource
	// providing the Redfish service root resource is an empty string or 'null'.
	//
	// Version added: v1.20.0
	ServiceUseNotification string
	// SessionService shall contain a link to a resource of type 'SessionService'.
	sessionService string
	// Storage shall contain a link to a resource collection of type
	// 'StorageCollection'.
	//
	// Version added: v1.9.0
	storage string
	// StorageServices shall contain a link to a resource collection of type
	// 'StorageServiceCollection'.
	//
	// Version added: v1.1.0
	storageServices string
	// StorageSystems shall contain a link to a resource collection of type
	// 'StorageSystemCollection'. This collection shall contain computer systems
	// that act as storage servers. The 'HostingRoles' property of each such
	// computer system shall contain a 'StorageServer' entry.
	//
	// Version added: v1.1.0
	storageSystems string
	// Systems shall contain a link to a resource collection of type
	// 'ComputerSystemCollection'.
	systems string
	// Tasks shall contain a link to a resource of type 'TaskService'.
	tasks string
	// TelemetryService shall contain a link to a resource of type
	// 'TelemetryService'.
	//
	// Version added: v1.4.0
	telemetryService string
	// ThermalEquipment shall contain a link to a resource of type
	// 'ThermalEquipment'.
	//
	// Version added: v1.16.0
	thermalEquipment string
	// UUID shall contain the identifier of the Redfish service instance. If SSDP
	// is used, this value shall contain the same UUID returned in an HTTP '200 OK'
	// response from an SSDP 'M-SEARCH' request during discovery. RFC4122 describes
	// methods to use to create a UUID value. The value should be considered to be
	// opaque. Client software should only treat the overall value as a universally
	// unique identifier and should not interpret any subfields within the UUID.
	UUID string
	// UpdateService shall contain a link to a resource of type 'UpdateService'.
	//
	// Version added: v1.1.0
	updateService string
	// Vendor shall include the name of the manufacturer or vendor represented by
	// this Redfish service. If this property is supported, the vendor name shall
	// not be included in the 'Product' property value.
	//
	// Version added: v1.5.0
	Vendor string
	// managerProvidingService is the URI for ManagerProvidingService.
	managerProvidingService string
	// sessions is the URI for Sessions.
	sessions string
}

// UnmarshalJSON unmarshals a Service object from the raw JSON.
func (s *Service) UnmarshalJSON(b []byte) error {
	type temp Service
	type sLinks struct {
		ManagerProvidingService schemas.Link `json:"ManagerProvidingService"`
		Sessions                schemas.Link `json:"Sessions"`
	}
	var tmp struct {
		temp
		Links              sLinks
		AccountService     schemas.Link `json:"AccountService"`
		AggregationService schemas.Link `json:"AggregationService"`
		AutomationNodes    schemas.Link `json:"AutomationNodes"`
		Cables             schemas.Link `json:"Cables"`
		CertificateService schemas.Link `json:"CertificateService"`
		Chassis            schemas.Link `json:"Chassis"`
		ComponentIntegrity schemas.Link `json:"ComponentIntegrity"`
		CompositionService schemas.Link `json:"CompositionService"`
		EventService       schemas.Link `json:"EventService"`
		Fabrics            schemas.Link `json:"Fabrics"`
		Facilities         schemas.Link `json:"Facilities"`
		JSONSchemas        schemas.Link `json:"JsonSchemas"`
		JobService         schemas.Link `json:"JobService"`
		KeyService         schemas.Link `json:"KeyService"`
		LicenseService     schemas.Link `json:"LicenseService"`
		Managers           schemas.Link `json:"Managers"`
		NVMeDomains        schemas.Link `json:"NVMeDomains"`
		PowerEquipment     schemas.Link `json:"PowerEquipment"`
		RegisteredClients  schemas.Link `json:"RegisteredClients"`
		Registries         schemas.Link `json:"Registries"`
		ResourceBlocks     schemas.Link `json:"ResourceBlocks"`
		ServiceConditions  schemas.Link `json:"ServiceConditions"`
		SessionService     schemas.Link `json:"SessionService"`
		Storage            schemas.Link `json:"Storage"`
		StorageServices    schemas.Link `json:"StorageServices"`
		StorageSystems     schemas.Link `json:"StorageSystems"`
		Systems            schemas.Link `json:"Systems"`
		Tasks              schemas.Link `json:"Tasks"`
		TelemetryService   schemas.Link `json:"TelemetryService"`
		ThermalEquipment   schemas.Link `json:"ThermalEquipment"`
		UpdateService      schemas.Link `json:"UpdateService"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Service(tmp.temp)

	// Extract the links to other entities for later
	s.managerProvidingService = tmp.Links.ManagerProvidingService.String()
	s.sessions = tmp.Links.Sessions.String()
	s.accountService = tmp.AccountService.String()
	s.aggregationService = tmp.AggregationService.String()
	s.automationNodes = tmp.AutomationNodes.String()
	s.cables = tmp.Cables.String()
	s.certificateService = tmp.CertificateService.String()
	s.chassis = tmp.Chassis.String()
	s.componentIntegrity = tmp.ComponentIntegrity.String()
	s.compositionService = tmp.CompositionService.String()
	s.eventService = tmp.EventService.String()
	s.fabrics = tmp.Fabrics.String()
	s.facilities = tmp.Facilities.String()
	s.jSONSchemas = tmp.JSONSchemas.String()
	s.jobService = tmp.JobService.String()
	s.keyService = tmp.KeyService.String()
	s.licenseService = tmp.LicenseService.String()
	s.managers = tmp.Managers.String()
	s.nVMeDomains = tmp.NVMeDomains.String()
	s.powerEquipment = tmp.PowerEquipment.String()
	s.registeredClients = tmp.RegisteredClients.String()
	s.registries = tmp.Registries.String()
	s.resourceBlocks = tmp.ResourceBlocks.String()
	s.serviceConditions = tmp.ServiceConditions.String()
	s.sessionService = tmp.SessionService.String()
	s.storage = tmp.Storage.String()
	s.storageServices = tmp.StorageServices.String()
	s.storageSystems = tmp.StorageSystems.String()
	s.systems = tmp.Systems.String()
	s.tasks = tmp.Tasks.String()
	s.telemetryService = tmp.TelemetryService.String()
	s.thermalEquipment = tmp.ThermalEquipment.String()
	s.updateService = tmp.UpdateService.String()

	return nil
}

// ServiceRoot will get a Service instance from the service.
func ServiceRoot(c schemas.Client) (*Service, error) {
	resp, err := c.Get(schemas.DefaultServiceRoot)
	defer schemas.DeferredCleanupHTTPResponse(resp)
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

// ManagerProvidingService gets the ManagerProvidingService linked resource.
func (s *Service) ManagerProvidingService() (*schemas.Manager, error) {
	if s.managerProvidingService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.Manager](s.GetClient(), s.managerProvidingService)
}

// Sessions gets the Sessions linked resource.
func (s *Service) Sessions() (*schemas.Session, error) {
	if s.sessions == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.Session](s.GetClient(), s.sessions)
}

// AccountService gets the AccountService linked resource.
func (s *Service) AccountService() (*schemas.AccountService, error) {
	if s.accountService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.AccountService](s.GetClient(), s.accountService)
}

// AggregationService gets the AggregationService linked resource.
func (s *Service) AggregationService() (*schemas.AggregationService, error) {
	if s.aggregationService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.AggregationService](s.GetClient(), s.aggregationService)
}

// AutomationNodes gets the AutomationNodes collection.
func (s *Service) AutomationNodes() ([]*schemas.AutomationNode, error) {
	if s.automationNodes == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.AutomationNode](s.GetClient(), s.automationNodes)
}

// Cables gets the Cables collection.
func (s *Service) Cables() ([]*schemas.Cable, error) {
	if s.cables == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Cable](s.GetClient(), s.cables)
}

// CertificateService gets the CertificateService linked resource.
func (s *Service) CertificateService() (*schemas.CertificateService, error) {
	if s.certificateService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.CertificateService](s.GetClient(), s.certificateService)
}

// Chassis gets the Chassis collection.
func (s *Service) Chassis() ([]*schemas.Chassis, error) {
	if s.chassis == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Chassis](s.GetClient(), s.chassis)
}

// ComponentIntegrity gets the ComponentIntegrity collection.
func (s *Service) ComponentIntegrity() ([]*schemas.ComponentIntegrity, error) {
	if s.componentIntegrity == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.ComponentIntegrity](s.GetClient(), s.componentIntegrity)
}

// CompositionService gets the CompositionService linked resource.
func (s *Service) CompositionService() (*schemas.CompositionService, error) {
	if s.compositionService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.CompositionService](s.GetClient(), s.compositionService)
}

// EventService gets the EventService linked resource.
func (s *Service) EventService() (*schemas.EventService, error) {
	if s.eventService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.EventService](s.GetClient(), s.eventService)
}

// Fabrics gets the Fabrics collection.
func (s *Service) Fabrics() ([]*schemas.Fabric, error) {
	if s.fabrics == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Fabric](s.GetClient(), s.fabrics)
}

// Facilities gets the Facilities collection.
func (s *Service) Facilities() ([]*schemas.Facility, error) {
	if s.facilities == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Facility](s.GetClient(), s.facilities)
}

// JSONSchemas gets the JSONSchemas collection.
func (s *Service) JSONSchemas() ([]*schemas.JSONSchemaFile, error) {
	if s.jSONSchemas == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.JSONSchemaFile](s.GetClient(), s.jSONSchemas)
}

// JobService gets the JobService linked resource.
func (s *Service) JobService() (*schemas.JobService, error) {
	if s.jobService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.JobService](s.GetClient(), s.jobService)
}

// KeyService gets the KeyService linked resource.
func (s *Service) KeyService() (*schemas.KeyService, error) {
	if s.keyService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.KeyService](s.GetClient(), s.keyService)
}

// LicenseService gets the LicenseService linked resource.
func (s *Service) LicenseService() (*schemas.LicenseService, error) {
	if s.licenseService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.LicenseService](s.GetClient(), s.licenseService)
}

// Managers gets the Managers collection.
func (s *Service) Managers() ([]*schemas.Manager, error) {
	if s.managers == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Manager](s.GetClient(), s.managers)
}

// NVMeDomains gets the NVMeDomains collection.
func (s *Service) NVMeDomains() ([]*schemas.NVMeDomain, error) {
	if s.nVMeDomains == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.NVMeDomain](s.GetClient(), s.nVMeDomains)
}

// PowerEquipment gets the PowerEquipment linked resource.
func (s *Service) PowerEquipment() (*schemas.PowerEquipment, error) {
	if s.powerEquipment == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.PowerEquipment](s.GetClient(), s.powerEquipment)
}

// RegisteredClients gets the RegisteredClients collection.
func (s *Service) RegisteredClients() ([]*schemas.RegisteredClient, error) {
	if s.registeredClients == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.RegisteredClient](s.GetClient(), s.registeredClients)
}

// Registries gets the Registries collection.
func (s *Service) Registries() ([]*schemas.MessageRegistryFile, error) {
	if s.registries == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.MessageRegistryFile](s.GetClient(), s.registries)
}

// ResourceBlocks gets the ResourceBlocks collection.
func (s *Service) ResourceBlocks() ([]*schemas.ResourceBlock, error) {
	if s.resourceBlocks == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.ResourceBlock](s.GetClient(), s.resourceBlocks)
}

// ServiceConditions gets the ServiceConditions linked resource.
func (s *Service) ServiceConditions() (*schemas.ServiceConditions, error) {
	if s.serviceConditions == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.ServiceConditions](s.GetClient(), s.serviceConditions)
}

// SessionService gets the SessionService linked resource.
func (s *Service) SessionService() (*schemas.SessionService, error) {
	if s.sessionService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.SessionService](s.GetClient(), s.sessionService)
}

// Storage gets the Storage collection.
func (s *Service) Storage() ([]*schemas.Storage, error) {
	if s.storage == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.Storage](s.GetClient(), s.storage)
}

// StorageServices gets the StorageServices collection.
func (s *Service) StorageServices() ([]*schemas.StorageService, error) {
	if s.storageServices == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.StorageService](s.GetClient(), s.storageServices)
}

// StorageSystems gets the StorageSystems collection.
func (s *Service) StorageSystems() ([]*schemas.ComputerSystem, error) {
	if s.storageSystems == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.ComputerSystem](s.GetClient(), s.storageSystems)
}

// Systems gets the Systems collection.
func (s *Service) Systems() ([]*schemas.ComputerSystem, error) {
	if s.systems == "" {
		return nil, nil
	}
	return schemas.GetCollectionObjects[schemas.ComputerSystem](s.GetClient(), s.systems)
}

// Tasks gets the Tasks linked resource.
func (s *Service) Tasks() (*schemas.TaskService, error) {
	if s.tasks == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.TaskService](s.GetClient(), s.tasks)
}

// TelemetryService gets the TelemetryService linked resource.
func (s *Service) TelemetryService() (*schemas.TelemetryService, error) {
	if s.telemetryService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.TelemetryService](s.GetClient(), s.telemetryService)
}

// ThermalEquipment gets the ThermalEquipment linked resource.
func (s *Service) ThermalEquipment() (*schemas.ThermalEquipment, error) {
	if s.thermalEquipment == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.ThermalEquipment](s.GetClient(), s.thermalEquipment)
}

// UpdateService gets the UpdateService linked resource.
func (s *Service) UpdateService() (*schemas.UpdateService, error) {
	if s.updateService == "" {
		return nil, nil
	}
	return schemas.GetObject[schemas.UpdateService](s.GetClient(), s.updateService)
}

// CreateSession creates a new session and returns the token and id
func (s *Service) CreateSession(username, password string) (*schemas.AuthToken, error) {
	return schemas.CreateSession(s.GetClient(), s.sessions, username, password)
}

// DeleteSession logout the specified session
func (s *Service) DeleteSession(url string) error {
	return schemas.DeleteSession(s.GetClient(), url)
}

// DeepOperations shall contain information about deep operations that the
// service supports.
type DeepOperations struct {
	// DeepPATCH shall indicate whether this service supports the Redfish
	// Specification-defined deep 'PATCH' operation.
	//
	// Version added: v1.7.0
	DeepPATCH bool
	// DeepPOST shall indicate whether this service supports the Redfish
	// Specification-defined deep 'POST' operation.
	//
	// Version added: v1.7.0
	DeepPOST bool
	// MaxLevels shall contain the maximum levels of resources allowed in deep
	// operations.
	//
	// Version added: v1.7.0
	MaxLevels uint
}

// Expand shall contain information about the support of the '$expand' query
// parameter by the service.
type Expand struct {
	// ExpandAll shall indicate whether this service supports the asterisk ('*')
	// option of the '$expand' query parameter.
	//
	// Version added: v1.3.0
	ExpandAll bool
	// Levels shall indicate whether the service supports the '$levels' option of
	// the '$expand' query parameter.
	//
	// Version added: v1.3.0
	Levels bool
	// Links shall be a boolean indicating whether this service supports the use
	// of tilde (expand only entries in the Links section) as a value for the
	// $expand query parameter as described by the specification.
	Links bool
	// MaxLevels shall contain the maximum '$levels' option value in the '$expand'
	// query parameter. This property shall be present if the 'Levels' property
	// contains 'true'.
	//
	// Version added: v1.3.0
	MaxLevels uint
	// NoLinks shall indicate whether the service supports the period ('.') option
	// of the '$expand' query parameter.
	//
	// Version added: v1.3.0
	NoLinks bool
	// managerProvidingService is the URI for ManagerProvidingService.
	ManagerProvidingService string
	// sessions is the URI for Sessions.
	Sessions string
}

// ProtocolFeaturesSupported shall contain information about protocol features
// that the service supports.
type ProtocolFeaturesSupported struct {
	// DeepOperations shall contain information about deep operations that the
	// service supports.
	//
	// Version added: v1.7.0
	DeepOperations DeepOperations
	// ExcerptQuery shall indicate whether this service supports the 'excerpt'
	// query parameter.
	//
	// Version added: v1.4.0
	ExcerptQuery bool
	// ExpandQuery shall contain information about the support of the '$expand'
	// query parameter by the service.
	//
	// Version added: v1.3.0
	ExpandQuery Expand
	// FilterQuery shall indicate whether this service supports the '$filter' query
	// parameter.
	//
	// Version added: v1.3.0
	FilterQuery bool
	// FilterQueryComparisonOperations shall indicate whether the service supports
	// the 'eq', 'ge', 'gt', 'le', 'lt', and 'ne' options for the '$filter' query
	// parameter. This property shall not be present if 'FilterQuery' contains
	// 'false'.
	//
	// Version added: v1.17.0
	FilterQueryComparisonOperations bool
	// FilterQueryCompoundOperations shall indicate whether the service supports
	// the Redfish Specification-defined grouping operators '()', 'and', 'not', and
	// 'or' options for the '$filter' query parameter. This property shall not be
	// present if 'FilterQuery' contains 'false'.
	//
	// Version added: v1.17.0
	FilterQueryCompoundOperations bool
	// IncludeOriginOfConditionQuery shall indicate whether the service supports
	// the 'includeoriginofcondition' query parameter.
	//
	// Version added: v1.18.0
	IncludeOriginOfConditionQuery bool
	// MultipleHTTPRequests shall indicate whether this service supports multiple
	// outstanding HTTP requests.
	//
	// Version added: v1.14.0
	MultipleHTTPRequests bool
	// OnlyMemberQuery shall indicate whether this service supports the 'only'
	// query parameter.
	//
	// Version added: v1.4.0
	OnlyMemberQuery bool
	// SelectQuery shall indicate whether this service supports the '$select' query
	// parameter.
	//
	// Version added: v1.3.0
	SelectQuery bool
	// TopSkipQuery shall indicate whether this service supports both the '$top'
	// and '$skip' query parameters.
	//
	// Version added: v1.17.0
	TopSkipQuery bool
}
