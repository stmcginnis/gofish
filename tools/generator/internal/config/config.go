//
// SPDX-License-Identifier: BSD-3-Clause
//

package config

import "strings"

// CommonNameChanges maps JSON property names to Go field names
var CommonNameChanges = map[string]string{
	"Oem": "OEM",
	"Id":  "ID",
}

// CommonDescriptions provides standard descriptions for common properties
var CommonDescriptions = map[string]string{
	"Description":    "Description provides a description of this resource.",
	"Id":             "ID uniquely identifies the resource.",
	"Name":           "Name is the name of the resource or array element.",
	"@odata.context": "ODataContext is the odata context.",
	"@odata.etag":    "ODataEtag is the odata etag.",
	"@odata.id":      "ODataID is the odata identifier.",
	"@odata.type":    "ODataType is the odata type.",
	"Identifier":     "Identifier shall be unique within the managed ecosystem.",
}

// CommonTypes maps property names to Go types.
// These are all in the same schemas package, so no package prefix is needed
// (except for json.RawMessage which is a stdlib type).
var CommonTypes = map[string]string{
	"Condition":         "Condition",
	"ElectricalContext": "ElectricalContext",
	"EventType":         "EventType",
	"Health":            "Health",
	"Identifier":        "Identifier",
	"IndicatorLED":      "IndicatorLED",
	"Location":          "Location",
	"LogicalContext":    "LogicalContext",
	"Oem":               "json.RawMessage",
	"OperationType":     "OperationType",
	"PowerState":        "PowerState",
	"PrivilegeType":     "PrivilegeType",
	"Protocol":          "Protocol",
	"Redundancy":        "Redundancy",
	"RedundantGroup":    "RedundantGroup",
	"ResetType":         "ResetType",
	"Schedule":          "Schedule",
	"Status":            "Status",
}

// EntityProperties are properties that indicate a type should embed Entity
// These properties should be skipped when generating structs that embed Entity
var EntityProperties = []string{
	"Name",
	"Id",
	"Description",
	"@odata.id",
	"@odata.etag",
	"@Message.ExtendedInfo",
}

// ExcludedDefinitions are definition names to skip during generation
// These are manually-maintained infrastructure types, not schema-generated
var ExcludedDefinitions = []string{
	"Actions",
	"OemActions",
	"Entity",             // Manually maintained in entity.go
	"Resource",           // Manually maintained in types.go
	"ResourceCollection", // Manually maintained in types.go
	"Item",               // Base type, manually maintained
}

var ExcludeReadWriteProperties = []string{
	"Oem",
	"OEM",
	"ODataContext",
	"ODataType",
	"odata.context",
	"odata.type",
}

// PropertyHasTarget checks if a definition has 'target' and 'title' properties (action definitions)
func IsActionDefinition(properties map[string]bool) bool {
	hasTarget := properties["target"]
	hasTitle := properties["title"]
	return hasTarget && hasTitle
}

// GetGoFieldName converts a JSON property name to a Go field name
func GetGoFieldName(jsonName string) string {
	if override, ok := CommonNameChanges[jsonName]; ok {
		return override
	}

	// Handle @odata.* properties (including Property@odata.count pattern)
	if strings.Contains(jsonName, "@") {
		return convertODataName(jsonName)
	}

	// Handle properties with dots or hyphens - convert to camelCase
	if strings.ContainsAny(jsonName, ".-") {
		return convertToCamelCase(jsonName)
	}

	return jsonName
}

// convertToCamelCase converts names with dots and hyphens to camelCase
// Examples: "some.property" -> "SomeProperty", "some-property" -> "SomeProperty"
func convertToCamelCase(name string) string {
	// Split on dots and hyphens
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '.' || r == '-'
	})

	// Capitalize first letter of each part
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return strings.Join(parts, "")
}

// convertODataName converts @odata.* names to Go field names
func convertODataName(name string) string {
	// @odata.id -> ODataID
	// @odata.type -> ODataType
	// @odata.context -> ODataContext
	// @odata.etag -> ODataEtag
	// Members@odata.count -> MembersCount

	if name == "@odata.id" {
		return "ODataID"
	}
	if name == "@odata.type" {
		return "ODataType"
	}
	if name == "@odata.context" {
		return "ODataContext"
	}
	if name == "@odata.etag" {
		return "ODataEtag"
	}

	// Handle property@odata.count pattern -> PropertyCount
	if len(name) > 12 && name[len(name)-12:] == "@odata.count" {
		prefix := name[:len(name)-12]
		return prefix + "Count"
	}

	return "OData" + capitalizeFirst(name[7:])
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

// TypeRenames maps "Schema.Definition" to the Go type name to emit for that
// definition. All definitions share one Go package, so a definition whose name
// collides with a different type from another schema gets a schema-specific
// name here. References to the definition from any schema use the new name.
var TypeRenames = map[string]string{
	"ActionInfo.Parameters":                          "ActionInfoParameter",
	"AddressPool.Ethernet":                           "APEthernet",
	"AddressPool.GenZ":                               "APGenZ",
	"Cable.ConnectorType":                            "CableConnectorType",
	"Certificate.Identifier":                         "CertificateIdentifier",
	"ComputerSystem.VirtualMediaConfig":              "ComputerVirtualMediaConfig",
	"Control.ImplementationType":                     "ControlImplementationType",
	"CXLLogicalDevice.QoSTelemetryCapabilities":      "CXLDeviceQoSTelemetryCapabilities",
	"DataSecurityLoSCapabilities.AuthenticationType": "SFAuthenticationType",
	"FabricAdapter.GenZ":                             "FabricAdapterGenZ",
	"Job.Payload":                                    "JobPayload",
	"JsonSchemaFile.Location":                        "SchemaFileLocation",
	"Key.NVMeoF":                                     "KeyNVMeoF",
	"Key.SSHKeyType":                                 "SSHAlgoKeyType",
	"LogService.OverWritePolicy":                     "LogServiceOverWritePolicy",
	"LogService.TransferProtocolType":                "LogServiceTransferProtocolType",
	"ManagerNetworkProtocol.HTTPSProtocol":           "HTTPSProtocolSettings",
	"ManagerNetworkProtocol.Protocol":                "ProtocolSetting",
	"Memory.CXL":                                     "MemoryCXL",
	"Memory.HealthData":                              "MemoryHealthData",
	"MemoryMetrics.CurrentPeriod":                    "MemoryMetricsCurrentPeriod",
	"MemoryMetrics.CXL":                              "MemoryMetricsCXL",
	"MemoryMetrics.LifeTime":                         "MemoryMetricsLifeTime",
	"MemoryRegion.MemoryExtent":                      "MemoryRegionExtent",
	"MessageRegistry.Message":                        "MessageRegistryMessage",
	"MessageRegistryFile.Location":                   "MessageRegistryFileLocation",
	"MetricDefinition.ImplementationType":            "MetricDefinitionImplementationType",
	"NetworkDeviceFunctionMetrics.Ethernet":          "NetworkDeviceFunctionMetricsEthernet",
	"NetworkDeviceFunctionMetrics.FibreChannel":      "NetworkDeviceFunctionMetricsFibreChannel",
	"NetworkPort.LinkStatus":                         "NetworkPortLinkStatus",
	"NetworkPort.PortConnectionType":                 "NetworkPortConnectionType",
	"Outlet.PowerState":                              "OutletPowerState",
	"Policy.Condition":                               "PolicyCondition",
	"Policy.ConditionType":                           "PolicyConditionType",
	"Policy.Response":                                "PolicyResponse",
	"Port.CXL":                                       "CXLPort",
	"Port.GenZ":                                      "PortGenZ",
	"Port.LinkStatus":                                "PortLinkStatus",
	"Port.UALink":                                    "PortUALink",
	"PortMetrics.CXL":                                "PortMetricsCXL",
	"PortMetrics.FibreChannel":                       "PortMetricsFibreChannel",
	"PortMetrics.GenZ":                               "GenZPortMetrics",
	"PowerSupply.InputRange":                         "PowerSupplyUnitInputRange",
	"PowerSupply.PowerSupply":                        "PowerSupplyUnit",
	"PowerSupply.PowerSupplyType":                    "PowerSupplyUnitType",
	"Processor.EthernetInterface":                    "ProcessorEthernetInterface",
	"Processor.MemorySummary":                        "ProcessorMemorySummary",
	"Processor.UALink":                               "ProcessorUALink",
	"ProcessorMetrics.CurrentPeriod":                 "ProcessorMetricsCurrentPeriod",
	"ProcessorMetrics.LifeTime":                      "ProcessMetricsLifeTime",
	"SecureBootDatabase.ResetKeysType":               "SecureBootDatabaseResetKeysType",
	"Sensor.VoltageType":                             "SensorVoltageType",
	"SerialInterface.FlowControl":                    "SerialInferfaceFlowControl",
	"SerialInterface.Role":                           "InterfaceRole",
	"Settings.ApplyTime":                             "SettingsApplyTime",
	"Storage.ResetToDefaultsType":                    "StorageResetToDefaultsType",
	"Storage.StorageController":                      "StorageStorageController",
	"StoragePool.PoolType":                           "StoragePoolType",
	"StorageReplicaInfo.ConsistencyType":             "ReplicationConsistencyType",
	"SwitchMetrics.CurrentPeriod":                    "SwitchMetricsCurrentPeriod",
	"SwitchMetrics.LifeTime":                         "SwitchMetricsLifeTime",
	"TaskService.OverWritePolicy":                    "TaskServiceOverWritePolicy",
	"TelemetryService.MetricValue":                   "TelemetryMetricValue",
	"Thermal.Fan":                                    "ThermalFan",
	"Triggers.Threshold":                             "TriggerThreshold",
	"Triggers.Thresholds":                            "TriggerThresholds",
	"Triggers.Wildcard":                              "TriggerWildcard",
	"UpdateService.ApplyTime":                        "UpdateServiceApplyTime",
	"VirtualMedia.MediaType":                         "VirtualMediaType",
	"VirtualMedia.TransferProtocolType":              "VirtualMediaTransferProtocolType",
	// Keeps the unexported type name and "...vPPBStatusTypes" constant names.
	"VirtualPCI2PCIBridge.vPPBStatusTypes": "vPPBStatusTypes",
}

// TypeOwners names the schema that emits a Go type defined by more than one
// schema, when the default (the schema whose main type it is, then Resource,
// then the first schema alphabetically) is not the right one. Listing a type
// here also records that its other copies, if they differ, are deliberately
// replaced by the owner's.
var TypeOwners = map[string]string{
	"AccessCapability":               "Connection",
	"Attributes":                     "AttributeRegistry",
	"AuthenticationMethod":           "StorageGroup",
	"CalculationAlgorithmEnum":       "MetricDefinition",
	"Component":                      "ComputerSystem",
	"ConfigLockOptions":              "Storage",
	"ConfigurationLock":              "Storage",
	"DataDirection":                  "AllowDeny",
	"ExportType":                     "ComputerSystem",
	"ExternalAccountProvider":        "ExternalAccountProvider",
	"GetPersonalityNonceResponse":    "Storage",
	"GraphicalConnectTypesSupported": "ComputerSystem",
	"IndicatorLED":                   "Resource",
	"IPAddressType":                  "NetworkDeviceFunction",
	"LinkNetworkTechnology":          "Port",
	"LogDiagnosticDataTypes":         "LogEntry",
	"MemoryChunk":                    "MemoryRegion",
	"NMVeMinimumRequiredResetType":   "Storage",
	"NMVePersonalityKeyAlgorithm":    "Storage",
	"NMVeUnfreezeAuthMode":           "Storage",
	"NVMeConfigurationLockState":     "Storage",
	"NVMeoFSecureHashType":           "KeyPolicy",
	"NVMeoFSecurityProtocolType":     "KeyPolicy",
	"NVMePersonality":                "Storage",
	"PowerRestorePolicyTypes":        "ComputerSystem",
	"PowerState":                     "Resource",
	"SNMPAuthenticationProtocols":    "EventDestination",
	"SNMPEncryptionProtocols":        "EventDestination",
	"SNMPSettings":                   "EventDestination",
	"StorageAccessCapability":        "StorageAccessCapability",
	"SupportedEthernetCapabilities":  "Port",
	"SyslogFacility":                 "LogService",
	"SyslogFilter":                   "LogService",
	"SyslogSeverity":                 "LogService",
	"TargetConfigurationLockLevel":   "Storage",
	"TransferProtocolType":           "UpdateService",
	"VoltageType":                    "Outlet",
	"Wildcard":                       "MetricDefinition",
}

// SkipDefinitions lists "Schema.Definition" entries that are not generated.
var SkipDefinitions = map[string]bool{
	"Job.Parameters":                  true, // Job.Parameters is a map[string]any (FieldOverrides)
	"JobDocument.Parameter":           true,
	"MessageRegistry.MessageProperty": true, // MessageRegistry.Messages is a map (FieldOverrides)
	"Resource.Oem":                    true,
	"Resource.OemObject":              true,
	"Storage.StorageControllerLinks":  true,
}

// FieldOverrides replaces a generated struct field, keyed by "Type.field" using
// the generated Go names. Keys can name schema properties, Links fields,
// action target fields ("resetTarget") and action parameter struct fields. The value is the full replacement declaration,
// including any struct tag, or "" to drop the field. An overridden field is a
// plain field: the generator emits no link handling or getter for it, so a
// hand-written UnmarshalJSON or accessor must cover it if needed.
var FieldOverrides = map[string]string{
	"NetworkDeviceFunction.physicalPortAssignment":               "physicalPortAssignment string",
	"Settings.maintenanceWindowResource":                         "maintenanceWindowResource string",
	"Aggregate.ElementsCount":                                    "ElementsCount *uint `json:\",omitempty\"`",
	"Attributes.LowerBound":                                      "LowerBound *uint64 `json:\",omitempty\"`",
	"Attributes.UpperBound":                                      "UpperBound *uint64 `json:\",omitempty\"`",
	"Bios.Attributes":                                            "Attributes SettingsAttributes",
	"Boot.AliasBootOrder":                                        "AliasBootOrder []BootSource `json:\",omitempty\"`",
	"Boot.AutomaticRetryConfig":                                  "AutomaticRetryConfig AutomaticRetryConfig `json:\",omitempty\"`",
	"Boot.BootNext":                                              "BootNext string `json:\",omitempty\"`",
	"Boot.BootOrder":                                             "BootOrder []string `json:\",omitempty\"`",
	"Boot.BootOrderPropertySelection":                            "BootOrderPropertySelection BootOrderTypes `json:\",omitempty\"`",
	"Boot.BootSourceOverrideEnabled":                             "BootSourceOverrideEnabled BootSourceOverrideEnabled `json:\",omitempty\"`",
	"Boot.BootSourceOverrideMode":                                "BootSourceOverrideMode BootSourceOverrideMode `json:\",omitempty\"`",
	"Boot.BootSourceOverrideTarget":                              "BootSourceOverrideTarget BootSource `json:\",omitempty\"`",
	"Boot.StopBootOnFault":                                       "StopBootOnFault StopBootOnFault `json:\",omitempty\"`",
	"Boot.TrustedModuleRequiredToBoot":                           "TrustedModuleRequiredToBoot TrustedModuleRequiredToBoot `json:\",omitempty\"`",
	"Boot.UefiTargetBootSourceOverride":                          "UefiTargetBootSourceOverride string `json:\",omitempty\"`",
	"BootProgress.OemLastState":                                  "OEMLastState string",
	"Chassis.LocationIndicatorActive":                            "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"HostedServices.storageServices":                             "storageServices string",
	"Chassis.drives":                                             "drives string",
	"Circuit.LocationIndicatorActive":                            "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"CollectTelemetryDataResponse.TelemetryData":                 "telemetryData []string",
	"ComputerSystem.LocationIndicatorActive":                     "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"CoolantConnector.pH":                                        "PH SensorExcerpt `json:\"pH\"`",
	"DataProtectionLineOfService.replicaClassOfService":          "ReplicaClassOfService ClassOfService",
	"DataSecurityLineOfService.HostAuthenticationType":           "HostAuthenticationType AuthenticationType",
	"DataSecurityLineOfService.UserAuthenticationType":           "UserAuthenticationType AuthenticationType",
	"Dependency.MapToValue":                                      "MapToValue any `json:\",omitempty\"`",
	"Drive.Location":                                             "Location []Location",
	"Drive.LocationIndicatorActive":                              "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"EthernetInterface.MACAddress":                               "MACAddress string",
	"EthernetInterface.PermanentMACAddress":                      "PermanentMACAddress string",
	"EventDestination.OriginResources":                           "OriginResources []string",
	"EventDestination.Protocol":                                  "Protocol EventDestinationProtocol",
	"EventDestination.metricReportDefinitions":                   "MetricReportDefinitions []MetricReportDefinition",
	"EventService.submitTestEventTarget":                         "SubmitTestEventTarget string",
	"EventService.subscriptions":                                 "",
	"FileSystem.ImportedShares":                                  "ImportedShares []any",
	"IOConnectivityLoSCapabilities.supportedLinesOfService":      "SupportedLinesOfService []IOConnectivityLineOfService",
	"IPv4Address.SubnetMask":                                     "SubnetMask string",
	"IPv6Address.PrefixLength":                                   "PrefixLength uint8",
	"IPv6GatewayStaticAddress.PrefixLength":                      "PrefixLength uint8",
	"IPv6StaticAddress.PrefixLength":                             "PrefixLength uint8",
	"JSONSchemaFile.Location":                                    "Location SchemaFileLocation",
	"Job.Parameters":                                             "Parameters map[string]any",
	"JobDocumentSubmitJobParameters.Parameters":                  "Parameters []any `json:\"Parameters,omitempty\"`",
	"Limits.CPUCount":                                            "CPUCount *float32 `json:\",omitempty\"`",
	"LogEntry.diagnosticDataType":                                "DiagnosticDataType LogDiagnosticDataTypes",
	"LogServiceCollectDiagnosticDataParameters.TransferProtocol": "TransferProtocol LogServiceTransferProtocolType `json:\"TransferProtocol,omitempty\"`",
	"Manager.ServiceEntryPointUUID":                              "ServiceEntryPointUUID string",
	"ManagerNetworkProtocol.mDNS":                                "MDNS ProtocolSetting",
	"MapFrom.MapFromValue":                                       "MapFromValue any `json:\",omitempty\"`",
	"MappedVolume.volume":                                        "volume string",
	"Memory.FirmwareApiVersion":                                  "FirmwareAPIVersion string",
	"MessageRegistry.Messages":                                   "Messages map[string]MessageRegistryMessage",
	"MessageRegistryFile.Location":                               "Location []MessageRegistryFileLocation",
	"MessageRegistryMessage.MessageSeverity":                     "MessageSeverity string",
	"MetricReport.MetricReportDefinition":                        "metricReportDefinition ResolvedLink[MetricReportDefinition]",
	"NetworkDeviceFunction.iSCSIBoot":                            "ISCSIBoot ISCSIBoot",
	"Outlet.LocationIndicatorActive":                             "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"Outlet.PowerState":                                          "PowerState OutletPowerState",
	"OutletGroup.PowerState":                                     "PowerState OutletPowerState",
	"PostalAddress.Location":                                     "",
	"Power.redundancy":                                           "Redundancy []Redundancy",
	"PowerControl.PowerAllocatedWatts":                           "PowerAllocatedWatts *float32 `json:\",omitempty\"`",
	"PowerControl.PowerAvailableWatts":                           "PowerAvailableWatts *float32 `json:\",omitempty\"`",
	"PowerControl.PowerCapacityWatts":                            "PowerCapacityWatts *float32 `json:\",omitempty\"`",
	"PowerControl.PowerConsumedWatts":                            "PowerConsumedWatts *float32 `json:\",omitempty\"`",
	"PowerMetric.AverageConsumedWatts":                           "AverageConsumedWatts *float32 `json:\",omitempty\"`",
	"PowerMetric.MaxConsumedWatts":                               "MaxConsumedWatts *float32 `json:\",omitempty\"`",
	"PowerMetric.MinConsumedWatts":                               "MinConsumedWatts *float32 `json:\",omitempty\"`",
	"PowerSupply.LastPowerOutputWatts":                           "LastPowerOutputWatts *float32 `json:\",omitempty\"`",
	"PowerSupply.LineInputVoltage":                               "LineInputVoltage *float32 `json:\",omitempty\"`",
	"PowerSupply.PowerCapacityWatts":                             "PowerCapacityWatts *float32 `json:\",omitempty\"`",
	"PowerSupply.PowerInputWatts":                                "PowerInputWatts *float32 `json:\",omitempty\"`",
	"PowerSupply.PowerOutputWatts":                               "PowerOutputWatts *float32 `json:\",omitempty\"`",
	"PowerSupply.redundancy":                                     "redundancy []string",
	"Processor.metrics":                                          "metrics string",
	"ProcessorMemorySummary.metrics":                             "metrics string",
	"Redundancy.RedundancySet":                                   "redundancySet []string",
	"SSEFilterPropertiesSupported.EventType":                     "EventType bool",
	"Settings.settingsObject":                                    "SettingsObject string",
	"Storage.StorageControllers":                                 "StorageControllers []StorageController",
	"StoragePool.defaultClassOfService":                          "DefaultClassOfService ClassOfService",
	"Switch.LocationIndicatorActive":                             "LocationIndicatorActive *bool `json:\",omitempty\"`",
	"TelemetryMetricValue.metricDefinition":                      "metricDefinition string",
	"Temperature.LowerThresholdUser":                             "LowerThresholdUser *float32 `json:\",omitempty\"`",
	"Temperature.UpperThresholdUser":                             "UpperThresholdUser *float32 `json:\",omitempty\"`",
	"ThermalFan.redundancy":                                      "redundancy string",
	"TrustedComponent.tPMClearTarget":                            "tpmClearTarget string",
	"TrustedComponent.tPMGetEventLogTarget":                      "tpmGetEventLogTarget string",
	"VLAN.VLANID":                                                "VLANID uint16 `json:\"VLANId\"`",
	"VLAN.VLANPriority":                                          "VLANPriority uint8",
	"VLanNetworkInterface.VLANID":                                "VLANID uint16 `json:\"VLANId\"`",
	"VLanNetworkInterface.VLANPriority":                          "VLANPriority uint8",
	"VirtualMedia.Inserted":                                      "Inserted *bool `json:\",omitempty\"`",
	"VirtualMedia.WriteProtected":                                "WriteProtected *bool `json:\",omitempty\"`",
	"VirtualMedia.ejectMediaTarget":                              "",
	"VirtualMedia.insertMediaTarget":                             "",
	"VirtualMediaInsertMediaParameters.Inserted":                 "Inserted *bool `json:\"Inserted,omitempty\"`",
	"VirtualMediaInsertMediaParameters.Password":                 "Password *string `json:\"Password,omitempty\"`",
	"VirtualMediaInsertMediaParameters.TransferMethod":           "TransferMethod *TransferMethod `json:\"TransferMethod,omitempty\"`",
	"VirtualMediaInsertMediaParameters.TransferProtocolType":     "TransferProtocolType *TransferProtocolType `json:\"TransferProtocolType,omitempty\"`",
	"VirtualMediaInsertMediaParameters.UserName":                 "UserName *string `json:\"UserName,omitempty\"`",
	"VirtualMediaInsertMediaParameters.WriteProtected":           "WriteProtected *bool `json:\"WriteProtected,omitempty\"`",
	"Voltage.LowerThresholdCritical":                             "LowerThresholdCritical *float32 `json:\",omitempty\"`",
	"Voltage.LowerThresholdFatal":                                "LowerThresholdFatal *float32 `json:\",omitempty\"`",
	"Voltage.LowerThresholdNonCritical":                          "LowerThresholdNonCritical *float32 `json:\",omitempty\"`",
	"Voltage.MaxReadingRange":                                    "MaxReadingRange *float32 `json:\",omitempty\"`",
	"Voltage.MinReadingRange":                                    "MinReadingRange *float32 `json:\",omitempty\"`",
	"Voltage.ReadingVolts":                                       "ReadingVolts *float32 `json:\",omitempty\"`",
	"Voltage.UpperThresholdCritical":                             "UpperThresholdCritical *float32 `json:\",omitempty\"`",
	"Voltage.UpperThresholdFatal":                                "UpperThresholdFatal *float32 `json:\",omitempty\"`",
	"Voltage.UpperThresholdNonCritical":                          "UpperThresholdNonCritical *float32 `json:\",omitempty\"`",
}

// ExtraFields are hand-maintained field declarations appended to a generated
// struct, keyed by Go type name. Hand-written code in the same package (often a
// custom UnmarshalJSON) is responsible for populating them.
var ExtraFields = map[string][]string{
	"Certificate": {
		"// issuer is the URI for Issuer.\n\tissuer string",
	},
	// Not part of the schema and never populated; kept for API compatibility.
	"Expand": {
		"ManagerProvidingService string",
		"Sessions string",
	},
	"Bios": {
		"// settingsTarget is the URL to send settings update requests to.\n\tsettingsTarget string",
		"// settingsApplyTimes is a set of allowed settings update apply times. If none\n\t// are specified, then the system does not provide that information.\n\tsettingsApplyTimes []SettingsApplyTime",
		"// rawData holds the original serialized JSON so we can compare updates.\n\trawData []byte",
	},
	"Boot": {
		"AllowableBootSourceOverrideTargetValues []BootSource `json:\"BootSourceOverrideTarget@Redfish.AllowableValues,omitempty\"`",
		"AllowableUefiTargetBootSourceOverrideValues []string `json:\"UefiTargetBootSourceOverride@Redfish.AllowableValues,omitempty\"`",
	},
	"CXLLogicalDevice": {
		"// PCIeFunctionsCount is the number of PCIe functions assigned to this CXL logical device.\n\tPCIeFunctionsCount int",
	},
	"Chassis": {
		"// linkedDrives shall contain an array of links to resources\n\t// of type Drive that are in this chassis.\n\tlinkedDrives []string",
		"// DrivesCount is the number of drives attached to this chassis.\n\tDrivesCount int `json:\"Drives@odata.count\"`",
		"linkedPCIeDevices []string",
		"// resetActionInfoTarget is the URL to check what values are supported\n\tresetActionInfoTarget string",
		"// SupportedResetTypes, if provided, is the reset types this chassis supports.\n\tSupportedResetTypes []ResetType",
	},
	"ComputerSystem": {
		"// resetActionInfoTarget is the URL to check what values are supported\n\tresetActionInfoTarget string",
		"// settingsApplyTimes is a set of allowed settings update apply times. If none\n\t// are specified, then the system does not provide that information.\n\tSettingsApplyTimes []SettingsApplyTime",
		"settingsTarget string",
		"// SupportedResetTypes, if provided, is the reset types this system supports.\n\tsupportedResetTypes []ResetType",
	},
	"EventService": {
		"// Subscriptions shall contain the link to a resource collection of type\n\t// 'EventDestinationCollection'.\n\tSubscriptionsLink string",
	},
	"LogService": {
		"// collectDiagnosticInfoTarget is the URL to get ActionInfo about the CollectDiagnosticData action.\n\tcollectDiagnosticInfoTarget string",
	},
	"Manager": {
		"// OEMLinks are all OEM data under the 'Links' section.\n\tOEMLinks json.RawMessage",
		"// resetInfo contains URI for an ActionInfo Resource that describes this action.\n\tresetActionInfo string",
		"// SupportedResetTypes, if provided, is the reset types this system supports.\n\tSupportedResetTypes []ResetType",
		"resetToDefaultsActionInfoTarget string",
		"SupportedResetToDefaultsTypes []ResetToDefaultsType",
	},
	"PCIeDevice": {
		"// pCIeFunctions are the URIs for PCIeFunctions.\n\tpCIeFunctionsLink []string",
	},
	"VirtualMedia": {
		"// MediaType indicates which MediaTypes is used (non-standard)\n\tMediaType VirtualMediaType",
		"// ejectMediaTarget is the URL to send EjectMedia requests.\n\tejectMedia ActionTarget",
		"// SupportsMediaEject indicates if ejecting virtual media is supported.\n\tSupportsMediaEject bool",
		"// insertMediaTarget is the URL to send InsertMedia requests.\n\tinsertMedia ActionTarget",
		"// SupportsMediaInsert indicates if inserting virtual media is supported.\n\tSupportsMediaInsert bool",
	},
	"VirtualMediaInsertMediaParameters": {
		"// MediaType is not part of the current schema, but some vendors require it.\n\tMediaType VirtualMediaType `json:\",omitempty\"`",
	},
}

// PluralNames overrides the default "<Type>s" plural used in ListReferenced<Plural>.
var PluralNames = map[string]string{
	"CertificateLocations": "CertificateLocations",
	"MessageRegistry":      "MessageRegistries",
	"PowerSupply":          "PowerSupplies",
	"Redundancy":           "Redundancies",
	"SoftwareInventory":    "SoftwareInventories",
}
