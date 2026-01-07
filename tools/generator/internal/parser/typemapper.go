//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"slices"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/config"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

var NonLinkTypes = []string{
	"Capacity",
	"IOStatistics",
	"PhysicalContext",
	"ProvidedCapacity",
	"ProvidedClassOfService",
	"Manifest",
	"Scheduler",
	"ResolutionStep",
	"MetricReportDefinition",
}

// TypeMapper handles conversion from JSON Schema types to Go types
type TypeMapper struct {
	commonTypes map[string]string
}

// NewTypeMapper creates a new TypeMapper
func NewTypeMapper() *TypeMapper {
	return &TypeMapper{
		commonTypes: config.CommonTypes,
	}
}

// MapType converts a JSON schema property to a Go type
func (tm *TypeMapper) MapType(propName string, prop *schema.JSONProperty) (goType string, isPointer bool, isArray bool) {
	// Check if this is a common type
	if commonType, ok := tm.commonTypes[propName]; ok {
		return commonType, false, false
	}

	// Handle special cases
	if propName == "UUID" {
		return "string", false, false
	}

	stringTypes := []string{
		"ID",
		"Id",
		"Description",
		"Name",
	}

	if slices.Contains(stringTypes, propName) {
		return "string", false, false
	}

	// Handle @odata.* properties - always map to string
	if strings.HasPrefix(propName, "@odata.") {
		return "string", false, false
	}

	// Handle type as string or array of strings
	typeStr, nullable := extractTypeAndNullable(prop.Type)

	// Check if this should be a Link type (lowercase start and not odata)
	if len(propName) > 0 &&
		typeStr != "string" &&
		propName[0] >= 'a' &&
		propName[0] <= 'z' &&
		!strings.Contains(strings.ToLower(propName), "odata") {
		return "Link", false, false
	}

	// Handle $ref references
	if prop.Ref != "" {
		refType := extractRefType(prop.Ref)

		// Map OData primitive types to Go types
		if strings.Contains(prop.Ref, "odata-v4.json") {
			switch refType {
			case "count":
				return "int", false, false
			case "idRef":
				return qualifyRefType("Entity", prop.Ref), false, false
			}
		}

		if slices.Contains(stringTypes, refType) {
			return "string", false, false
		}

		return qualifyRefType(refType, prop.Ref), false, false
	}

	// Handle items (arrays)
	if prop.Items != nil {
		if prop.Items.Ref != "" {
			itemType := extractRefType(prop.Items.Ref)

			// Map OData primitive types to Go types
			if strings.Contains(prop.Items.Ref, "odata-v4.json") {
				switch itemType {
				case "count":
					return "int", false, true
				case "idRef":
					return qualifyRefType("Entity", prop.Items.Ref), false, true
				}
			}

			if slices.Contains(stringTypes, itemType) {
				return "string", false, true
			}

			return qualifyRefType(itemType, prop.Items.Ref), false, true
		}
		if len(prop.Items.AnyOf) > 0 {
			for _, anyOfItem := range prop.Items.AnyOf {
				if ref, ok := anyOfItem["$ref"]; ok {
					itemType := extractRefType(ref)
					return qualifyRefType(itemType, ref), false, true
				}
			}
		}
	}

	// Handle anyOf
	if len(prop.AnyOf) > 0 {
		for _, anyOfItem := range prop.AnyOf {
			if ref, ok := anyOfItem["$ref"]; ok {
				refType := extractRefType(ref)

				// Map OData primitive types to Go types
				if strings.Contains(ref, "odata-v4.json") {
					switch refType {
					case "count":
						return "int", false, false
					case "idRef":
						return qualifyRefType("Entity", ref), false, false
					}
				}

				if slices.Contains(stringTypes, refType) {
					return "string", false, false
				}

				return qualifyRefType(refType, ref), false, false
			}
		}
	}

	switch typeStr {
	case "object":
		// Custom type
		return cleanIdentifier(propName), false, false

	case "integer":
		// Use uint if minimum is >= 0
		if prop.Minimum != nil && *prop.Minimum >= 0 {
			return "uint", nullable, false
		}
		return "int", nullable, false

	case "number", "numeric":
		// Check if this looks like an integer field
		if strings.HasSuffix(strings.ToLower(propName), "count") {
			return "int", nullable, false
		}
		return "float64", nullable, false

	case "boolean":
		return "bool", false, false

	case "array":
		// Determine array element type
		if prop.Items != nil {
			elemType, elemPointer, _ := tm.MapType(propName, prop.Items)
			if elemPointer {
				return "*" + elemType, false, true
			}
			return elemType, false, true
		}
		return "string", false, true

	default:
		// Default to string
		return "string", false, false
	}
}

func qualifyRefType(typeName, _ string) string {
	return cleanIdentifier(typeName)
}

// extractSchemaNameFromRef extracts the schema name from a $ref URL.
func extractSchemaNameFromRef(ref string) string {
	if strings.HasPrefix(ref, "#/definitions/") {
		return ""
	}
	if strings.Contains(ref, "#/definitions/") {
		parts := strings.Split(ref, "#/definitions/")
		if len(parts) == 2 {
			return extractSchemaNameFromURL(parts[0])
		}
	}
	return ""
}

// extractSchemaNameFromURL extracts the schema name from a URL.
func extractSchemaNameFromURL(url string) string {
	if idx := strings.LastIndex(url, "/"); idx >= 0 {
		filename := url[idx+1:]
		if dotIdx := strings.Index(filename, "."); dotIdx >= 0 {
			return filename[:dotIdx]
		}
		return filename
	}
	return ""
}

// extractRefType extracts the type name from a $ref URL
func extractRefType(ref string) string {
	return schema.ExtractTypeFromRef(ref, true)
}

// extractTypeAndNullable handles type which can be a string or array of strings
func extractTypeAndNullable(typeVal any) (string, bool) {
	switch t := typeVal.(type) {
	case string:
		return t, false
	case []any:
		// Array of types, check for null
		nullable := false
		actualType := "string"
		for _, item := range t {
			if str, ok := item.(string); ok {
				if str == "null" {
					nullable = true
				} else {
					actualType = str
				}
			}
		}
		return actualType, nullable
	default:
		return "string", false
	}
}

// IsLinkProperty checks if a property is a link to another resource
func IsLinkProperty(propName string, prop *schema.JSONProperty) bool {
	if slices.Contains(NonLinkTypes, propName) {
		return false
	}

	// Check for URIs or @odata.id pattern
	if strings.HasSuffix(propName, "URI") || strings.HasSuffix(propName, "Uri") {
		return true
	}

	// Check format
	if prop.Format == "uri-reference" || prop.Format == "uri" {
		return true
	}

	// Check for collection references
	if prop.Ref != "" && strings.Contains(prop.Ref, "Collection") {
		return true
	}

	// Check for idRef - OData generic reference type
	if prop.Ref != "" && strings.Contains(prop.Ref, "idRef") {
		return true
	}

	// Check for direct reference to a standalone resource (non-array, non-collection)
	// If the $ref points to a schema where the filename matches the definition name,
	// it's a standalone entity with its own URI - a reference to it is a link
	if prop.Ref != "" && !strings.Contains(prop.Ref, "Collection") {
		// Skip odata primitive types and Resource.json definitions
		if !strings.Contains(prop.Ref, "odata-v4.json") && !strings.Contains(prop.Ref, "Resource.json") {
			schemaName := extractSchemaNameFromRef(prop.Ref)
			typeName := extractRefType(prop.Ref)
			if schemaName != "" && schemaName == typeName {
				return true
			}
		}
	}

	// Check for array items that reference another resource type
	// These are typically links stored as arrays of URIs
	if prop.Items != nil && prop.Items.Ref != "" {
		// Skip odata primitive types
		if !strings.Contains(prop.Items.Ref, "odata-v4.json") {
			// If the schema filename matches the definition name, it's a standalone
			// entity with its own schema file and URIs — references to it are links
			schemaName := extractSchemaNameFromRef(prop.Items.Ref)
			typeName := extractRefType(prop.Items.Ref)
			if schemaName != "" && schemaName == typeName {
				return true
			}

			// Check if the description indicates these are links
			desc := strings.ToLower(prop.Description)
			longDesc := strings.ToLower(prop.LongDescription)
			if strings.Contains(longDesc, "links to") ||
				strings.Contains(longDesc, "array of links") ||
				strings.Contains(desc, "links to") ||
				strings.Contains(desc, "array of links") {
				return true
			}
		}
	}

	// Check if description indicates this is a link to another resource
	desc := strings.ToLower(prop.Description)
	longDesc := strings.ToLower(prop.LongDescription)
	linkPhrases := []string{
		"link to a resource",
		"link to the resource",
		"link to an instance",
		"link to a collection",
		"shall contain a link",
		"shall contain links",
		"an array of links",
	}

	for _, phrase := range linkPhrases {
		if strings.Contains(longDesc, phrase) || strings.Contains(desc, phrase) {
			return true
		}
	}

	return false
}

// IsCollectionProperty checks if a property is a collection link
func IsCollectionProperty(prop *schema.JSONProperty) bool {
	if prop.Ref != "" && strings.Contains(prop.Ref, "Collection") {
		return true
	}
	return false
}
