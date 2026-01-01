//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// CommonTypes lists schemas that should always be in the common package
var CommonTypes = map[string]bool{
	"Status":              true,
	"Location":            true,
	"Identifier":          true,
	"Health":              true,
	"State":               true,
	"PowerState":          true,
	"IndicatorLED":        true,
	"ResetType":           true,
	"Link":                true,
	"Links":               true,
	"Resource":            true,
	"ReferenceableMember": true,
	"ResourceCollection":  true,
	"Redundancy":          true,
	"Settings":            true,
	"Message":             true,
	"MessageRegistry":     true,
	"Privileges":          true,
	"PrivilegeRegistry":   true,
	"PhysicalContext":     true,
	"Protocol":            true,
	"IPAddresses":         true,
	"Schedule":            true,
	"ResolutionStep":      true,
	"LogEntry":            true,
	"Entity":              true,
	// Storage-related types referenced by Redfish
	"Volume":              true,
	"Capacity":            true,
	"IOStatistics":        true,
	"StorageReplicaInfo":  true,
}

// SwordfishTypes lists schemas that are Swordfish-specific
// Note: Types referenced by Redfish must be in CommonTypes instead
var SwordfishTypes = map[string]bool{
	"ClassOfService":                  true,
	"ConsistencyGroup":                true,
	"DataProtectionLineOfService":     true,
	"DataProtectionLoSCapabilities":   true,
	"DataSecurityLineOfService":       true,
	"DataSecurityLoSCapabilities":     true,
	"DataStorageLineOfService":        true,
	"DataStorageLoSCapabilities":      true,
	"EndpointGroup":                   true,
	"FileShare":                       true,
	"FileSystem":                      true,
	"IOConnectivityLineOfService":     true,
	"IOConnectivityLoSCapabilities":   true,
	"IOPerformanceLineOfService":      true,
	"IOPerformanceLoSCapabilities":    true,
	"LineOfService":                   true,
	"SpareResourceSet":                true,
	"StorageGroup":                    true,
	"StoragePool":                     true,
	"StorageService":                  true,
}

// CategorizeSchema determines the package type for a schema
func CategorizeSchema(schemaFile, objectName string) schema.PackageType {
	// Check if it's a known common type
	if CommonTypes[objectName] {
		return schema.PackageCommon
	}

	// Check if it's a known Swordfish type
	if SwordfishTypes[objectName] {
		return schema.PackageSwordfish
	}

	// Try to determine from schema content
	data, err := os.ReadFile(schemaFile)
	if err != nil {
		// Default to redfish if we can't read the file
		return schema.PackageRedfish
	}

	var rawSchema map[string]interface{}
	if err := json.Unmarshal(data, &rawSchema); err != nil {
		return schema.PackageRedfish
	}

	// Check the $id or copyright for Swordfish indicators
	if id, ok := rawSchema["$id"].(string); ok {
		if strings.Contains(id, "swordfish") {
			return schema.PackageSwordfish
		}
	}

	if copyright, ok := rawSchema["copyright"].(string); ok {
		if strings.Contains(strings.ToLower(copyright), "snia") {
			return schema.PackageSwordfish
		}
	}

	// Check owningEntity
	if owner, ok := rawSchema["owningEntity"].(string); ok {
		if owner == "SNIA" {
			return schema.PackageSwordfish
		}
	}

	// Look at URIs to determine if it's common or redfish
	defsMap, ok := rawSchema["definitions"].(map[string]interface{})
	if ok {
		if def, ok := defsMap[objectName].(map[string]interface{}); ok {
			// Check if this has URIs (only full resources have URIs, not common types)
			if uris, ok := def["uris"].([]interface{}); ok && len(uris) > 0 {
				// Has URIs, so it's a full resource (redfish)
				return schema.PackageRedfish
			}

			// Check if it's marked as abstract (common types are often abstract)
			if insertable, ok := def["insertable"].(bool); ok && !insertable {
				if deletable, ok := def["deletable"].(bool); ok && !deletable {
					if updatable, ok := def["updatable"].(bool); ok && !updatable {
						// Not insertable, deletable, or updatable - likely a common enum/type
						return schema.PackageCommon
					}
				}
			}
		}
	}

	// Default to redfish for everything else
	return schema.PackageRedfish
}

// IsCommonType checks if a schema should be in the common package
func IsCommonType(objectName string) bool {
	return CommonTypes[objectName]
}

// IsSwordfishType checks if a schema should be in the swordfish package
func IsSwordfishType(objectName string) bool {
	return SwordfishTypes[objectName]
}
