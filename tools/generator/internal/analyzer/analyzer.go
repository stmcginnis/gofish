//
// SPDX-License-Identifier: BSD-3-Clause
//

package analyzer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// SchemaInfo holds metadata about a parsed schema
type SchemaInfo struct {
	Name       string
	Origin     schema.SchemaOrigin
	SchemaFile string
}

// DefinitionInfo holds metadata about a single definition within a schema
type DefinitionInfo struct {
	Name          string              // Definition name (e.g., "ReplicaInfo")
	SchemaName    string              // Parent schema (e.g., "StorageReplicaInfo")
	SchemaOrigin  schema.SchemaOrigin // Redfish or Swordfish
	IsEntity      bool                // Is this a main entity type (has Id, Name, @odata.id)?
	Dependencies  []string            // Other definitions this one references
	TargetPackage schema.PackageType  // Where this definition should go
}

// DependencyAnalyzer performs cross-schema analysis at definition level
type DependencyAnalyzer struct {
	Schemas                    map[string]*SchemaInfo
	Definitions                map[string]*DefinitionInfo // All definitions across all schemas
	SwordfishDefsUsedByRedfish map[string]bool            // Definition-level tracking
	verbose                    bool
}

// NewDependencyAnalyzer creates a new analyzer
func NewDependencyAnalyzer(verbose bool) *DependencyAnalyzer {
	return &DependencyAnalyzer{
		Schemas:                    make(map[string]*SchemaInfo),
		Definitions:                make(map[string]*DefinitionInfo),
		SwordfishDefsUsedByRedfish: make(map[string]bool),
		verbose:                    verbose,
	}
}

// AnalyzeSchema analyzes a single schema file
func (da *DependencyAnalyzer) AnalyzeSchema(schemaFile, objectName string) error {
	rawSchema, err := schema.ReadAndParseSchema(schemaFile)
	if err != nil {
		return err
	}

	origin := da.determineOrigin(rawSchema)

	// Check if a schema with this name already exists
	// Prefer Redfish schemas over Swordfish when there's a name collision
	// This handles cases where the same schema exists in both repos
	if existing, ok := da.Schemas[objectName]; ok {
		if existing.Origin == schema.OriginRedfish && origin == schema.OriginSwordfish {
			// Redfish schema already exists, skip Swordfish version
			if da.verbose {
				fmt.Printf("  Skipping Swordfish schema %s (Redfish version already analyzed)\n", objectName)
			}
			return nil
		}
	}

	info := &SchemaInfo{
		Name:       objectName,
		Origin:     origin,
		SchemaFile: schemaFile,
	}
	da.Schemas[objectName] = info

	// Extract all definitions and their dependencies
	da.extractDefinitions(rawSchema, schemaFile, objectName, origin)

	return nil
}

// extractDefinitions extracts all definitions from a schema and their dependencies
func (da *DependencyAnalyzer) extractDefinitions(rawSchema map[string]any, schemaFile, objectName string, origin schema.SchemaOrigin) {
	schemaDir := filepath.Dir(schemaFile)
	baseName := strings.TrimSuffix(filepath.Base(schemaFile), ".json")

	// Collect definitions from both base and versioned schemas
	allDefs := make(map[string]map[string]any)

	// Get definitions from base schema
	if defs, ok := rawSchema["definitions"].(map[string]any); ok {
		for defName, defData := range defs {
			if defMap, ok := defData.(map[string]any); ok {
				allDefs[defName] = defMap
			}
		}
	}

	// Also check versioned schema file for additional/updated definitions
	versionedFile := da.findLatestVersionFile(schemaDir, baseName)
	if versionedFile != "" {
		if versionedData, err := os.ReadFile(versionedFile); err == nil {
			var versionedSchema map[string]any
			if json.Unmarshal(versionedData, &versionedSchema) == nil {
				if defs, ok := versionedSchema["definitions"].(map[string]any); ok {
					for defName, defData := range defs {
						if defMap, ok := defData.(map[string]any); ok {
							allDefs[defName] = defMap
						}
					}
				}
			}
		}
	}

	// Process each definition
	for defName, defMap := range allDefs {
		// Use qualified key (SchemaName.DefName) to avoid collisions
		// Many schemas have definitions with the same name (e.g., "Actions")
		qualifiedName := objectName + "." + defName

		defInfo := &DefinitionInfo{
			Name:         defName,
			SchemaName:   objectName,
			SchemaOrigin: origin,
			IsEntity:     da.isEntityDefinition(defMap),
		}

		// Extract dependencies (other definitions this one references)
		// Pass the schema name so internal refs can be qualified
		defInfo.Dependencies = da.extractDefinitionDependencies(defMap, objectName)

		da.Definitions[qualifiedName] = defInfo
	}
}

// isEntityDefinition checks if a definition is an entity type (has Id, Name, @odata.id)
func (da *DependencyAnalyzer) isEntityDefinition(defMap map[string]any) bool {
	props, ok := defMap["properties"].(map[string]any)
	if !ok {
		return false
	}

	hasID := false
	hasName := false
	hasODataID := false

	for propName := range props {
		switch propName {
		case "Id":
			hasID = true
		case "Name":
			hasName = true
		case "@odata.id":
			hasODataID = true
		}
	}

	return hasID && hasName && hasODataID
}

// extractDefinitionDependencies extracts all type references from a single definition
// schemaName is used to qualify internal references
func (da *DependencyAnalyzer) extractDefinitionDependencies(defMap map[string]any, schemaName string) []string {
	deps := make(map[string]bool)
	da.extractRefsFromDefinition(defMap, deps, schemaName)

	result := make([]string, 0, len(deps))
	for dep := range deps {
		result = append(result, dep)
	}
	return result
}

// extractRefsFromDefinition recursively extracts $ref values from a definition
// schemaName is used to qualify internal references
func (da *DependencyAnalyzer) extractRefsFromDefinition(m map[string]any, deps map[string]bool, schemaName string) {
	for key, value := range m {
		if key == "$ref" {
			if ref, ok := value.(string); ok {
				typeName := da.extractTypeNameFromRef(ref, schemaName)
				if typeName != "" {
					deps[typeName] = true
				}
			}
		} else {
			switch v := value.(type) {
			case map[string]any:
				da.extractRefsFromDefinition(v, deps, schemaName)
			case []any:
				for _, item := range v {
					if itemMap, ok := item.(map[string]any); ok {
						da.extractRefsFromDefinition(itemMap, deps, schemaName)
					}
				}
			}
		}
	}
}

// extractTypeNameFromRef extracts the type name from a $ref URL
// Returns qualified name (SchemaName.DefName) to avoid collisions
// currentSchema is used for internal references
func (da *DependencyAnalyzer) extractTypeNameFromRef(ref string, currentSchema string) string {
	// Skip odata references
	if strings.Contains(ref, "odata") {
		return ""
	}

	// Handle internal references like "#/definitions/ReplicaInfo"
	// These refer to definitions within the same schema
	if strings.HasPrefix(ref, "#/definitions/") {
		defName := strings.TrimPrefix(ref, "#/definitions/")
		// Skip Links (URI references) and OemActions (vendor-specific)
		// Note: We track Actions to ensure action parameter types are included in dependency analysis
		if defName == "Links" || defName == "OemActions" {
			return ""
		}
		// Qualify with current schema name
		return currentSchema + "." + defName
	}

	// Extract from external URLs like:
	// "http://redfish.dmtf.org/schemas/v1/Volume.json#/definitions/Volume"
	// "http://redfish.dmtf.org/schemas/swordfish/v1/Capacity.json#/definitions/Capacity"
	if strings.Contains(ref, "#/definitions/") {
		parts := strings.Split(ref, "#/definitions/")
		if len(parts) == 2 {
			defName := parts[1]
			// Handle Collection suffix
			if strings.HasSuffix(defName, "Collection") {
				defName = strings.TrimSuffix(defName, "Collection")
			}

			// Extract schema name from the URL
			// URL looks like: http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/ResetType
			urlPart := parts[0]
			schemaName := extractSchemaNameFromURL(urlPart)
			if schemaName != "" {
				return schemaName + "." + defName
			}

			// Fallback: use definition name without qualification
			// This may cause issues but is better than nothing
			return defName
		}
	}

	return ""
}

// extractSchemaNameFromURL extracts the schema name from a URL
// e.g., "http://redfish.dmtf.org/schemas/v1/Resource.json" -> "Resource"
func extractSchemaNameFromURL(url string) string {
	// Find the last path component before .json
	if idx := strings.LastIndex(url, "/"); idx >= 0 {
		filename := url[idx+1:]
		// Remove .json extension and any version suffix
		if dotIdx := strings.Index(filename, "."); dotIdx >= 0 {
			return filename[:dotIdx]
		}
		return filename
	}
	return ""
}

// determineOrigin determines if a schema is Redfish or Swordfish origin
func (da *DependencyAnalyzer) determineOrigin(rawSchema map[string]any) schema.SchemaOrigin {
	return schema.DetermineSchemaOrigin(rawSchema)
}

// findLatestVersionFile finds the latest versioned schema file
func (da *DependencyAnalyzer) findLatestVersionFile(schemaDir, baseName string) string {
	pattern := filepath.Join(schemaDir, baseName+".v*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return ""
	}

	// Find the latest version
	var latestFile string
	var latestMajor, latestMinor, latestPatch int

	re := regexp.MustCompile(`\.v(\d+)_(\d+)_(\d+)\.json$`)
	for _, match := range matches {
		if m := re.FindStringSubmatch(match); m != nil {
			var major, minor, patch int
			fmt.Sscanf(m[1], "%d", &major)
			fmt.Sscanf(m[2], "%d", &minor)
			fmt.Sscanf(m[3], "%d", &patch)

			if schema.CompareVersions(major, minor, patch, latestMajor, latestMinor, latestPatch) {
				latestMajor, latestMinor, latestPatch = major, minor, patch
				latestFile = match
			}
		}
	}

	return latestFile
}

// ComputeDependencies computes which Swordfish definitions are used by Redfish
// Uses transitive closure at the definition level
func (da *DependencyAnalyzer) ComputeDependencies() {
	// Step 1: Find Swordfish definitions directly referenced by Redfish definitions
	for _, defInfo := range da.Definitions {
		if defInfo.SchemaOrigin == schema.OriginRedfish {
			for _, dep := range defInfo.Dependencies {
				if depInfo, ok := da.Definitions[dep]; ok {
					if depInfo.SchemaOrigin == schema.OriginSwordfish {
						da.SwordfishDefsUsedByRedfish[dep] = true
					}
				}
			}
		}
	}

	// Step 2: Transitive closure - find all definitions reachable from those needed in common.
	// Since common cannot import redfish or swordfish, all transitive dependencies
	// of common-bound types must also be placed in common.
	changed := true
	for changed {
		changed = false
		for defName := range da.SwordfishDefsUsedByRedfish {
			if defInfo, ok := da.Definitions[defName]; ok {
				for _, dep := range defInfo.Dependencies {
					if _, ok := da.Definitions[dep]; ok {
						if !da.SwordfishDefsUsedByRedfish[dep] {
							da.SwordfishDefsUsedByRedfish[dep] = true
							changed = true
						}
					}
				}
			}
		}
	}

	// Step 3: Assign target packages to all definitions
	for defName, defInfo := range da.Definitions {
		defInfo.TargetPackage = da.computeDefinitionPackage(defName, defInfo)
	}

	if da.verbose {
		// Log Swordfish definitions going to common
		for defName := range da.SwordfishDefsUsedByRedfish {
			if defInfo, ok := da.Definitions[defName]; ok {
				fmt.Printf("  Swordfish definition %s (from %s) will be placed in common\n",
					defName, defInfo.SchemaName)
			}
		}
	}
}

// computeDefinitionPackage determines the target package for a definition
// defName is the qualified name (SchemaName.DefName)
func (da *DependencyAnalyzer) computeDefinitionPackage(defName string, defInfo *DefinitionInfo) schema.PackageType {
	// Schemas explicitly placed in common take precedence.
	if schema.IsCommonSchema(defInfo.SchemaName) {
		return schema.PackageCommon
	}

	// Infrastructure types always go to common (use simple name from defInfo)
	if schema.InfrastructureTypes[defInfo.Name] {
		return schema.PackageCommon
	}

	// Swordfish definitions used by Redfish go to common (use qualified name)
	if da.SwordfishDefsUsedByRedfish[defName] {
		return schema.PackageCommon
	}

	// Other Swordfish definitions go to swordfish
	if defInfo.SchemaOrigin == schema.OriginSwordfish {
		return schema.PackageSwordfish
	}

	// Redfish definitions go to redfish
	return schema.PackageRedfish
}

// GetPackageType returns the appropriate package for a schema (for backward compatibility)
func (da *DependencyAnalyzer) GetPackageType(schemaName string) schema.PackageType {
	// Schemas explicitly placed in common take precedence.
	if schema.IsCommonSchema(schemaName) {
		return schema.PackageCommon
	}

	// Infrastructure types always go to common
	if schema.InfrastructureTypes[schemaName] {
		return schema.PackageCommon
	}

	// Check if this schema has definitions going to common vs its native package
	hasCommonDefs := false
	hasNativePackageDefs := false

	for _, defInfo := range da.Definitions {
		if defInfo.SchemaName == schemaName {
			if defInfo.TargetPackage == schema.PackageCommon {
				hasCommonDefs = true
			} else {
				hasNativePackageDefs = true
			}
		}
	}

	// If all definitions go to common, the whole schema goes to common
	if hasCommonDefs && !hasNativePackageDefs {
		return schema.PackageCommon
	}

	if info, ok := da.Schemas[schemaName]; ok {
		if info.Origin == schema.OriginSwordfish {
			return schema.PackageSwordfish
		}
	}

	return schema.PackageRedfish
}

// GetDefinitionPackage returns the target package for a specific definition
func (da *DependencyAnalyzer) GetDefinitionPackage(defName string) schema.PackageType {
	if defInfo, ok := da.Definitions[defName]; ok {
		return defInfo.TargetPackage
	}

	// Default to redfish if not found
	return schema.PackageRedfish
}

// GetSchemaOrigin returns the origin of a schema
func (da *DependencyAnalyzer) GetSchemaOrigin(schemaName string) schema.SchemaOrigin {
	if info, ok := da.Schemas[schemaName]; ok {
		return info.Origin
	}
	return schema.OriginRedfish
}

// IsSwordfishTypeInCommon checks if a schema has definitions that are Swordfish types placed in common
func (da *DependencyAnalyzer) IsSwordfishTypeInCommon(schemaName string) bool {
	for _, defInfo := range da.Definitions {
		if defInfo.SchemaName == schemaName {
			if defInfo.SchemaOrigin == schema.OriginSwordfish && defInfo.TargetPackage == schema.PackageCommon {
				return true
			}
		}
	}
	return false
}

// GetDefinitionsForSchema returns all definitions for a given schema
func (da *DependencyAnalyzer) GetDefinitionsForSchema(schemaName string) []*DefinitionInfo {
	var defs []*DefinitionInfo
	for _, defInfo := range da.Definitions {
		if defInfo.SchemaName == schemaName {
			defs = append(defs, defInfo)
		}
	}
	return defs
}

// GetDefinitionsByPackage returns definitions for a schema grouped by target package
// Returns simple definition names (not qualified) for lookup in parsed definitions
func (da *DependencyAnalyzer) GetDefinitionsByPackage(schemaName string) map[schema.PackageType][]string {
	result := make(map[schema.PackageType][]string)

	for _, defInfo := range da.Definitions {
		if defInfo.SchemaName == schemaName {
			// Return simple name (not qualified) for compatibility with parser's definition map
			result[defInfo.TargetPackage] = append(result[defInfo.TargetPackage], defInfo.Name)
		}
	}

	return result
}

// NeedsSplitGeneration checks if a schema needs to be split across packages
func (da *DependencyAnalyzer) NeedsSplitGeneration(schemaName string) bool {
	packages := da.GetDefinitionsByPackage(schemaName)
	return len(packages) > 1
}
