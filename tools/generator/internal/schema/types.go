//
// SPDX-License-Identifier: BSD-3-Clause
//

package schema

import (
	"encoding/json"
	"os"
	"strings"
)

// PackageType represents the Go package category
type PackageType string

const (
	PackageSchemas PackageType = "schemas"
)

// SchemaOrigin indicates where a schema comes from (Redfish or Swordfish)
type SchemaOrigin string

const (
	OriginRedfish   SchemaOrigin = "redfish"
	OriginSwordfish SchemaOrigin = "swordfish"
)

// InfrastructureTypes lists core infrastructure types.
// These are manually maintained base types, not generated from schemas.
var InfrastructureTypes = map[string]bool{
	"Entity":              true,
	"Link":                true,
	"Links":               true,
	"Resource":            true,
	"ReferenceableMember": true,
	"ResourceCollection":  true,
}

// ConflictAction defines how to resolve a schema name conflict between Redfish and Swordfish.
type ConflictAction int

const (
	// ConflictSkipSwordfish means the Swordfish version is a true duplicate and should be skipped.
	ConflictSkipSwordfish ConflictAction = iota
	// ConflictPrefixSwordfish means the Swordfish version gets an "SF" prefix.
	ConflictPrefixSwordfish
)

// SchemaConflicts defines how to handle schemas that exist in both Redfish and Swordfish.
// Note: Volume is SNIA-owned in both bundles (Redfish includes SNIA schemas), so we skip
// the duplicate. Schedule has different DMTF vs SNIA versions that need SF prefix.
var SchemaConflicts = map[string]ConflictAction{
	"EndpointGroup": ConflictSkipSwordfish,
	"Schedule":      ConflictPrefixSwordfish,
	"Volume":        ConflictSkipSwordfish, // Both bundles have same SNIA schema
}

// IsInfrastructureType checks if a type is a core infrastructure type.
// These are manually maintained base types.
func IsInfrastructureType(typeName string) bool {
	return InfrastructureTypes[typeName]
}

// ShouldSkip returns true if this schema/origin combination should be skipped entirely.
func ShouldSkip(objectName string, origin SchemaOrigin) bool {
	if origin != OriginSwordfish {
		return false
	}
	action, ok := SchemaConflicts[objectName]
	return ok && action == ConflictSkipSwordfish
}

// NeedsSFPrefix returns true if this schema/origin combination needs an "SF" prefix.
func NeedsSFPrefix(objectName string, origin SchemaOrigin) bool {
	if origin != OriginSwordfish {
		return false
	}
	action, ok := SchemaConflicts[objectName]
	return ok && action == ConflictPrefixSwordfish
}

// DetermineSchemaOrigin determines if a parsed schema is Redfish or Swordfish origin.
// It checks owningEntity, $id, and copyright fields.
func DetermineSchemaOrigin(rawSchema map[string]any) SchemaOrigin {
	// Check owningEntity
	if owner, ok := rawSchema["owningEntity"].(string); ok {
		if owner == "SNIA" {
			return OriginSwordfish
		}
	}

	// Check $id for swordfish
	if id, ok := rawSchema["$id"].(string); ok {
		if strings.Contains(strings.ToLower(id), "swordfish") {
			return OriginSwordfish
		}
	}

	// Check copyright for SNIA
	if copyright, ok := rawSchema["copyright"].(string); ok {
		if strings.Contains(strings.ToLower(copyright), "snia") {
			return OriginSwordfish
		}
	}

	return OriginRedfish
}

// ReadAndParseSchema reads a JSON schema file and returns the parsed map.
// This is a helper to avoid duplicating the file read and unmarshal pattern.
func ReadAndParseSchema(schemaFile string) (map[string]any, error) {
	data, err := os.ReadFile(schemaFile)
	if err != nil {
		return nil, err
	}

	var rawSchema map[string]any
	if err := json.Unmarshal(data, &rawSchema); err != nil {
		return nil, err
	}

	return rawSchema, nil
}

// CompareVersions returns true if version (major, minor, patch) is greater than
// the current latest version (latestMajor, latestMinor, latestPatch).
// This is used for finding the latest versioned schema file.
func CompareVersions(major, minor, patch, latestMajor, latestMinor, latestPatch int) bool {
	return major > latestMajor ||
		(major == latestMajor && minor > latestMinor) ||
		(major == latestMajor && minor == latestMinor && patch > latestPatch)
}

// ExtractTypeFromRef extracts the type name from a $ref URL.
// ref looks like "http://redfish.dmtf.org/schemas/v1/ComputerSystem.json#/definitions/ComputerSystem"
// or "#/definitions/Link" or similar.
// If stripCollection is true, "Collection" suffix is removed from the type name.
func ExtractTypeFromRef(ref string, stripCollection bool) string {
	parts := strings.Split(ref, "/")
	if len(parts) == 0 {
		return ""
	}
	lastPart := parts[len(parts)-1]
	if stripCollection && strings.HasSuffix(lastPart, "Collection") {
		return strings.TrimSuffix(lastPart, "Collection")
	}
	return lastPart
}

// BuildDefinitionMap creates a map for quick definition lookup from a slice.
// If includeNameAlias is true, definitions are also mapped by their Name field
// (in addition to OriginalName), which is useful when names may differ.
func BuildDefinitionMap(defs []*Definition, includeNameAlias bool) map[string]*Definition {
	defMap := make(map[string]*Definition, len(defs))
	for _, def := range defs {
		defMap[def.OriginalName] = def
		if includeNameAlias && def.Name != def.OriginalName {
			defMap[def.Name] = def
		}
	}
	return defMap
}

// Definition represents a complete type definition (struct or enum)
type Definition struct {
	// Name is the Go type name
	Name string
	// OriginalName is the name from the JSON schema
	OriginalName string
	// Package is the target package (common, redfish, or swordfish)
	Package PackageType
	// Description is the type's documentation
	Description string
	// IsEntity indicates if this type embeds common.Entity
	IsEntity bool
	// IsEnum indicates if this is an enum type
	IsEnum bool
	// Properties are the struct fields (for non-enum types)
	Properties []*Property
	// EnumValues are the enum constants (for enum types)
	EnumValues []*EnumValue
	// ReadWriteProperties lists property names that can be updated
	ReadWriteProperties []string
	// SchemaFile is the source JSON schema file
	SchemaFile string
	// Version is the schema version (e.g., "v1_2_0")
	Version string
	// Actions are the available actions for this resource
	Actions []*Action
	// Links are the links to related resources
	Links []*Link
	// Release is the schema release version (e.g., "2018.3")
	Release string
	// Title is the schema title (e.g., "#CertificateService.v1_0_1.CertificateService")
	Title string
}

// Property represents a struct field
type Property struct {
	// Name is the Go field name
	Name string
	// JSONName is the original JSON property name
	JSONName string
	// Type is the Go type string
	Type string
	// Description is the field documentation
	Description string
	// IsPointer indicates if this should be a pointer type (nullable)
	IsPointer bool
	// IsArray indicates if this is an array type
	IsArray bool
	// IsReadOnly indicates if the property is read-only
	IsReadOnly bool
	// JSONTag is the struct tag for JSON serialization
	JSONTag string
	// IsPrivate indicates if this field should be unexported
	IsPrivate bool
	// GetterMethod is the name of the getter method (if IsPrivate)
	GetterMethod string
	// VersionAdded is the version when this property was added
	VersionAdded string
	// IsLink indicates if this is a reference to another resource
	IsLink bool
	// IsCollection indicates if this is a collection link
	IsCollection bool
	// IsDeprecated indicates if this property is deprecated
	IsDeprecated bool
}

// EnumValue represents an enum constant
type EnumValue struct {
	// Name is the Go constant name
	Name string
	// Value is the string value
	Value string
	// Description is the constant documentation
	Description string
}

// Action represents an action that can be performed on a resource
type Action struct {
	// Name is the action name (e.g., "Reset")
	Name string
	// JSONName is the full JSON action name (e.g., "#Chassis.Reset")
	JSONName string
	// Description is the action documentation
	Description string
	// Parameters are the action parameters
	Parameters []*ActionParameter
	// ResponseType is the Go type name for the action response (e.g., "GenerateCSRResponse")
	ResponseType string
}

// ActionParameter represents a parameter for an action
type ActionParameter struct {
	// Name is the parameter name
	Name string
	// Type is the parameter Go type
	Type string
	// Description is the parameter documentation
	Description string
	// Required indicates if the parameter is required
	Required bool
	// AllowableValues are the allowed values (for enums)
	AllowableValues []string
	// Ordinal is the order position of this parameter
	Ordinal int
	// OriginalName is the unmodified name of the JSON parameter definition
	OriginalName string
	// FieldName is the PascalCase cleaned name for parameter struct fields
	FieldName string
	// FieldDescription is the pre-formatted godoc comment for parameter struct fields
	FieldDescription string
}

// Link represents a link to a related resource
type Link struct {
	// Name is the link name (e.g., "ComputerSystems")
	Name string
	// JSONName is the JSON property name
	JSONName string
	// Type is the target resource type (e.g., "ComputerSystem")
	Type string
	// Description is the link documentation
	Description string
	// IsArray indicates if this is a link to multiple resources
	IsArray bool
	// Deprecated indicates if this link is deprecated
	Deprecated bool
}

// JSONProperty represents a property in the JSON Schema
type JSONProperty struct {
	Type                 any                      `json:"type"` // can be string or []string
	Description          string                   `json:"description"`
	LongDescription      string                   `json:"longDescription"`
	ReadOnly             bool                     `json:"readonly"`
	Deprecated           bool                     `json:"deprecated"`
	Ref                  string                   `json:"$ref"`
	AnyOf                []map[string]string      `json:"anyOf"`
	AllOf                []map[string]string      `json:"allOf"`
	Enum                 []string                 `json:"enum"`
	EnumDescriptions     map[string]string        `json:"enumDescriptions"`
	EnumLongDescriptions map[string]string        `json:"enumLongDescriptions"`
	Items                *JSONProperty            `json:"items"`
	Properties           map[string]*JSONProperty `json:"properties"`
	Minimum              *float64                 `json:"minimum"`
	Pattern              string                   `json:"pattern"`
	Format               string                   `json:"format"`
	AdditionalInfo       map[string]any           `json:"Redfish.Revisions"`
}
