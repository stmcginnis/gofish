//
// SPDX-License-Identifier: BSD-3-Clause
//

package schema

// PackageType represents the Go package category
type PackageType string

const (
	PackageCommon    PackageType = "common"
	PackageRedfish   PackageType = "redfish"
	PackageSwordfish PackageType = "swordfish"
)

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
}

// JSONSchema represents a parsed JSON Schema structure
type JSONSchema struct {
	ID          string                 `json:"$id"`
	Schema      string                 `json:"$schema"`
	Title       string                 `json:"title"`
	Type        string                 `json:"type"`
	Definitions map[string]*Definition `json:"definitions"`
	Properties  map[string]*JSONProperty
	Required    []string
}

// JSONProperty represents a property in the JSON Schema
type JSONProperty struct {
	Type                 interface{}              `json:"type"` // can be string or []string
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
	AdditionalInfo       map[string]interface{}   `json:"Redfish.Revisions"`
}
