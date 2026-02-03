//
// SPDX-License-Identifier: BSD-3-Clause
//

package analyzer

import (
	"fmt"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// SchemaInfo holds metadata about a parsed schema
type SchemaInfo struct {
	Origin schema.SchemaOrigin
}

// DependencyAnalyzer tracks schema origins for conflict detection.
type DependencyAnalyzer struct {
	Schemas map[string]*SchemaInfo
	verbose bool
}

// NewDependencyAnalyzer creates a new analyzer
func NewDependencyAnalyzer(verbose bool) *DependencyAnalyzer {
	return &DependencyAnalyzer{
		Schemas: make(map[string]*SchemaInfo),
		verbose: verbose,
	}
}

// AnalyzeSchema records a schema's origin for later conflict detection.
func (da *DependencyAnalyzer) AnalyzeSchema(schemaFile, objectName string) error {
	rawSchema, err := schema.ReadAndParseSchema(schemaFile)
	if err != nil {
		return err
	}

	origin := schema.DetermineSchemaOrigin(rawSchema)

	// If both Redfish and Swordfish define the same schema name,
	// keep the Redfish version as the primary entry.
	if existing, ok := da.Schemas[objectName]; ok {
		if existing.Origin == schema.OriginRedfish && origin == schema.OriginSwordfish {
			if da.verbose {
				fmt.Printf("  Skipping Swordfish schema %s (Redfish version already analyzed)\n", objectName)
			}
			return nil
		}
	}

	da.Schemas[objectName] = &SchemaInfo{
		Origin: origin,
	}

	return nil
}

// GetSchemaOrigin returns the origin of a schema
func (da *DependencyAnalyzer) GetSchemaOrigin(schemaName string) schema.SchemaOrigin {
	if info, ok := da.Schemas[schemaName]; ok {
		return info.Origin
	}
	return schema.OriginRedfish
}

// ShouldSkip returns true if this schema should be skipped (Swordfish duplicate).
func (da *DependencyAnalyzer) ShouldSkip(objectName string) bool {
	origin := da.GetSchemaOrigin(objectName)
	return schema.ShouldSkip(objectName, origin)
}

// NeedsSFPrefix returns true if this schema needs an "SF" prefix (Swordfish conflict).
func (da *DependencyAnalyzer) NeedsSFPrefix(objectName string) bool {
	origin := da.GetSchemaOrigin(objectName)
	return schema.NeedsSFPrefix(objectName, origin)
}

