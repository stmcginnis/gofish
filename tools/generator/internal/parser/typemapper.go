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

	// Check if this should be a Link type (lowercase start and not odata)
	if len(propName) > 0 && propName[0] >= 'a' && propName[0] <= 'z' && !strings.Contains(strings.ToLower(propName), "odata") {
		return "common.Link", false, false
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
				return "Entity", false, false
			}
		}

		if slices.Contains(stringTypes, refType) {
			return "string", false, false
		}

		return refType, false, false
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
					return "Entity", false, true
				}
			}

			if slices.Contains(stringTypes, itemType) {
				return "string", false, false
			}

			return itemType, false, true
		}
		if len(prop.Items.AnyOf) > 0 {
			for _, anyOfItem := range prop.Items.AnyOf {
				if ref, ok := anyOfItem["$ref"]; ok {
					itemType := extractRefType(ref)
					return itemType, false, true
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
						return "Entity", false, false
					}
				}

				if slices.Contains(stringTypes, refType) {
					return "string", false, false
				}

				return refType, false, false
			}
		}
	}

	// Handle type as string or array of strings
	typeStr, nullable := extractTypeAndNullable(prop.Type)

	switch typeStr {
	case "object":
		// Custom type
		return propName, false, false

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

// extractRefType extracts the type name from a $ref URL
func extractRefType(ref string) string {
	// Extract last part after /
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		// Handle cases where it ends with "Collection"
		// Collections should resolve to the element type, not the collection type
		if strings.HasSuffix(lastPart, "Collection") {
			return strings.TrimSuffix(lastPart, "Collection")
		}
		return lastPart
	}
	return ""
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

	// Check if description indicates this is a link to another resource
	desc := strings.ToLower(prop.Description)
	longDesc := strings.ToLower(prop.LongDescription)
	linkPhrases := []string{
		"link to a resource",
		"link to the resource",
		"link to an instance",
		"link to a collection",
		"shall contain a link",
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
