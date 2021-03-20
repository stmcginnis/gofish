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
	// CertificateService shall be a link to the CertificateService.
	certificateService string
	// Chassis shall only contain a reference to a collection of resources that
	// comply to the Chassis schema.
	chassis string
	// CompositionService shall only contain a reference to a resource that
	// complies to the CompositionService schema.
	compositionService string
	// Description provides a description of this resource.
	Description string
	// EventService shall only contain a reference to a resource that complies
	// to the EventService schema.
	eventService string
	// Fabrics shall contain references to all Fabric instances.
	fabrics string
	// JobService shall only contain a reference to a resource that conforms to
	// the JobService schema.
	jobService string
	// JsonSchemas shall only contain a reference to a collection of resources
	// that comply to the SchemaFile schema where the files are Json-Schema
	// files.
	jsonSchemas string
	// Managers shall only contain a reference to a collection of resources that
	// comply to the Managers schema.
	managers string
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
	// Registries shall contain a reference to Message Registry.
	registries string
	// ResourceBlocks shall contain references to all Resource Block instances.
	resourceBlocks string
	// SessionService shall only contain a reference to a resource that complies
	// to the SessionService schema.
	sessionService string
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
}

// UnmarshalJSON unmarshals a Service object from the raw JSON.
func (serviceroot *Service) UnmarshalJSON(b []byte) error {
	type temp Service
	var t struct {
		temp
		CertificateService common.Link
		Chassis            common.Link
		Managers           common.Link
		Tasks              common.Link
		StorageServices    common.Link
		StorageSystems     common.Link
		AccountService     common.Link
		EventService       common.Link
		Registries         common.Link
		Systems            common.Link
		CompositionService common.Link
		Fabrics            common.Link
		JobService         common.Link
		JSONSchemas        common.Link `json:"JsonSchemas"`
		ResourceBlocks     common.Link
		SessionService     common.Link
		TelemetryService   common.Link
		UpdateService      common.Link
		Links              struct {
			Sessions common.Link
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*serviceroot = Service(t.temp)
	serviceroot.certificateService = string(t.CertificateService)
	serviceroot.chassis = string(t.Chassis)
	serviceroot.managers = string(t.Managers)
	serviceroot.tasks = string(t.Tasks)
	serviceroot.sessions = string(t.Links.Sessions)
	serviceroot.storageServices = string(t.StorageServices)
	serviceroot.storageSystems = string(t.StorageSystems)
	serviceroot.accountService = string(t.AccountService)
	serviceroot.eventService = string(t.EventService)
	serviceroot.registries = string(t.Registries)
	serviceroot.systems = string(t.Systems)
	serviceroot.compositionService = string(t.CompositionService)
	serviceroot.fabrics = string(t.Fabrics)
	serviceroot.jobService = string(t.JobService)
	serviceroot.jsonSchemas = string(t.JSONSchemas)
	serviceroot.resourceBlocks = string(t.ResourceBlocks)
	serviceroot.sessionService = string(t.SessionService)
	serviceroot.telemetryService = string(t.TelemetryService)
	serviceroot.updateService = string(t.UpdateService)

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

// Chassis gets the chassis instances managed by this service.
func (serviceroot *Service) Chassis() ([]*redfish.Chassis, error) {
	return redfish.ListReferencedChassis(serviceroot.Client, serviceroot.chassis)
}

// Managers gets the manager instances of this service.
func (serviceroot *Service) Managers() ([]*redfish.Manager, error) {
	return redfish.ListReferencedManagers(serviceroot.Client, serviceroot.managers)
}

// StorageSystems gets the storage system instances managed by this service.
func (serviceroot *Service) StorageSystems() ([]*swordfish.StorageSystem, error) {
	return swordfish.ListReferencedStorageSystems(serviceroot.Client, serviceroot.storageSystems)
}

// StorageServices gets the Swordfish storage services
func (serviceroot *Service) StorageServices() ([]*swordfish.StorageService, error) {
	return swordfish.ListReferencedStorageServices(serviceroot.Client, serviceroot.storageServices)
}

// Tasks gets the system's tasks
func (serviceroot *Service) Tasks() ([]*redfish.Task, error) {
	return redfish.ListReferencedTasks(serviceroot.Client, serviceroot.tasks)
}

// CreateSession creates a new session and returns the token and id
func (serviceroot *Service) CreateSession(username, password string) (*redfish.AuthToken, error) {
	return redfish.CreateSession(serviceroot.Client, serviceroot.sessions, username, password)
}

// Sessions gets the system's active sessions
func (serviceroot *Service) Sessions() ([]*redfish.Session, error) {
	return redfish.ListReferencedSessions(serviceroot.Client, serviceroot.sessions)
}

// DeleteSession logout the specified session
func (serviceroot *Service) DeleteSession(url string) error {
	return redfish.DeleteSession(serviceroot.Client, url)
}

// AccountService gets the Redfish AccountService
func (serviceroot *Service) AccountService() (*redfish.AccountService, error) {
	return redfish.GetAccountService(serviceroot.Client, serviceroot.accountService)
}

// EventService gets the Redfish EventService
func (serviceroot *Service) EventService() (*redfish.EventService, error) {
	return redfish.GetEventService(serviceroot.Client, serviceroot.eventService)
}

// Registries gets the Redfish Registries
func (serviceroot *Service) Registries() ([]*redfish.MessageRegistryFile, error) {
	return redfish.ListReferencedMessageRegistryFiles(serviceroot.Client, serviceroot.registries)
}

// MessageRegistries gets all the available message registries in all languages
func (serviceroot *Service) MessageRegistries() ([]*redfish.MessageRegistry, error) {
	return redfish.ListReferencedMessageRegistries(serviceroot.Client, serviceroot.registries)
}

// MessageRegistry gets a specific message registry.
// uri is the uri for the message registry
func (serviceroot *Service) MessageRegistry(uri string) (*redfish.MessageRegistry, error) {
	return redfish.GetMessageRegistry(serviceroot.Client, uri)
}

// MessageRegistriesByLanguage gets the message registries by language.
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistriesByLanguage(language string) ([]*redfish.MessageRegistry, error) {
	return redfish.ListReferencedMessageRegistriesByLanguage(serviceroot.Client, serviceroot.registries, language)
}

// MessageRegistryByLanguage gets a specific message registry by language.
// registry is used to identify the correct Message Registry file and it shall
// contain the Message Registry name and it major and minor versions, as defined
// by the Redfish Specification, for example: "Alert.1.0.0".
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageRegistryByLanguage(registry, language string) (*redfish.MessageRegistry, error) {
	return redfish.GetMessageRegistryByLanguage(serviceroot.Client, serviceroot.registries, registry, language)
}

// MessageByLanguage tries to find and get the message in the correct language from the informed messageID.
// messageID is the key used to find the registry, version and message, for example: "Alert.1.0.LanDisconnect"
//  - The segment before the 1st period is the Registry Name (Registry Prefix): Alert
//  - The segment between the 1st and 2nd period is the major version: 1
//  - The segment between the 2nd and 3rd period is the minor version: 0
//  - The segment after the 3rd period is the Message Identifier in the Registry: LanDisconnect
// language is the RFC5646-conformant language code for the message registry, for example: "en".
func (serviceroot *Service) MessageByLanguage(messageID, language string) (*redfish.MessageRegistryMessage, error) {
	return redfish.GetMessageFromMessageRegistryByLanguage(serviceroot.Client, serviceroot.registries, messageID, language)
}

// Systems get the system instances from the service
func (serviceroot *Service) Systems() ([]*redfish.ComputerSystem, error) {
	return redfish.ListReferencedComputerSystems(serviceroot.Client, serviceroot.systems)
}

// CompositionService gets the composition service instance
func (serviceroot *Service) CompositionService() (*redfish.CompositionService, error) {
	return redfish.GetCompositionService(serviceroot.Client, serviceroot.compositionService)
}

// UpdateService gets the update service instance
func (serviceroot *Service) UpdateService() (*redfish.UpdateService, error) {
	return redfish.GetUpdateService(serviceroot.Client, serviceroot.updateService)
}
