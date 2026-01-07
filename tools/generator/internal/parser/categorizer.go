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

// CategorizeSchema determines the package type for a schema.
// All generated schemas go to the schemas package.
func CategorizeSchema(schemaFile, objectName string) schema.PackageType {
	return schema.PackageSchemas
}
