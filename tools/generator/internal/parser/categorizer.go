//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// GetSchemaOrigin determines if a schema is Redfish or Swordfish origin
func GetSchemaOrigin(schemaFile, objectName string) schema.SchemaOrigin {
	rawSchema, err := schema.ReadAndParseSchema(schemaFile)
	if err != nil {
		return schema.OriginRedfish
	}

	return schema.DetermineSchemaOrigin(rawSchema)
}

// CategorizeSchema determines the package type for a schema
// This is a simplified version for single-file mode without full dependency analysis
// For batch mode, use the analyzer package instead
func CategorizeSchema(schemaFile, objectName string) schema.PackageType {
	// Check if it's an infrastructure type
	if schema.IsInfrastructureType(objectName) {
		return schema.PackageCommon
	}

	// Schemas explicitly placed in common take precedence.
	if schema.IsCommonSchema(objectName) {
		return schema.PackageCommon
	}

	// Determine origin and use that as the package
	origin := GetSchemaOrigin(schemaFile, objectName)
	if origin == schema.OriginSwordfish {
		return schema.PackageSwordfish
	}

	return schema.PackageRedfish
}
