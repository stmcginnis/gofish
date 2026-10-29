//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"path/filepath"
	"slices"
	"strings"
	"sync"

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
	// schemaName is the schema being parsed; local "#/definitions/" refs resolve against it.
	schemaName string
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

	// HttpBootUri is nullable in the schema, and callers need to be able to PATCH an
	// explicit empty value to clear it - a plain string can't be distinguished from
	// "absent" under omitempty, so this is a pointer despite being a string type.
	if propName == "HttpBootUri" {
		return "string", true, false
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
		typeStr != "any" &&
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
				return tm.qualifyRefType("Entity", prop.Ref), false, false
			}
		}

		if slices.Contains(stringTypes, refType) {
			return "string", false, false
		}

		return tm.qualifyRefType(refType, prop.Ref), false, false
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
					return tm.qualifyRefType("Entity", prop.Items.Ref), false, true
				}
			}

			if slices.Contains(stringTypes, itemType) {
				return "string", false, true
			}

			return tm.qualifyRefType(itemType, prop.Items.Ref), false, true
		}
		if len(prop.Items.AnyOf) > 0 {
			for _, anyOfItem := range prop.Items.AnyOf {
				if ref, ok := anyOfItem["$ref"]; ok {
					itemType := extractRefType(ref)
					return tm.qualifyRefType(itemType, ref), false, true
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
						return tm.qualifyRefType("Entity", ref), false, false
					}
				}

				if slices.Contains(stringTypes, refType) {
					return "string", false, false
				}

				return tm.qualifyRefType(refType, ref), false, false
			}
		}
	}

	switch typeStr {
	case "any":
		return "any", false, false

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

// qualifyRefType returns the Go type name for a definition referenced by ref,
// applying any configured rename for the schema the definition lives in.
func (tm *TypeMapper) qualifyRefType(typeName, ref string) string {
	schemaName := extractSchemaNameFromRef(ref)
	if schemaName == "" {
		schemaName = tm.schemaName
	}
	return goTypeName(schemaName, typeName)
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

// extractTypeAndNullable handles type which can be a string or array of strings.
// When multiple distinct non-null types are present (e.g. ["string", "boolean", "number", "null"]),
// it returns "any" to avoid silently picking one type and losing data at unmarshal time.
func extractTypeAndNullable(typeVal any) (string, bool) {
	switch t := typeVal.(type) {
	case string:
		return t, false
	case []any:
		nullable := false
		var firstType string
		multiType := false
		for _, item := range t {
			str, ok := item.(string)
			if !ok {
				continue
			}
			if str == "null" {
				nullable = true
				continue
			}
			if firstType == "" {
				firstType = str
			} else if str != firstType {
				multiType = true
			}
		}
		if multiType {
			return "any", nullable
		}
		if firstType != "" {
			return firstType, nullable
		}
		return "string", nullable
	default:
		return "string", false
	}
}

// SchemaDirs lists every schema directory being generated, so references can be
// resolved across the Redfish and Swordfish bundles.
var SchemaDirs []string

var (
	resourceCacheMu sync.Mutex
	resourceCache   = map[string]bool{}
)

// isResource reports whether name can be referenced by @odata.id: its base
// schema lists idRef as an alternative. Common objects such as Schedule or
// Message share their schema's name but are always embedded, not linked.
func isResource(name string) bool {
	resourceCacheMu.Lock()
	defer resourceCacheMu.Unlock()
	if v, ok := resourceCache[name]; ok {
		return v
	}
	v := lookupResource(name)
	resourceCache[name] = v
	return v
}

func lookupResource(name string) bool {
	for _, dir := range SchemaDirs {
		raw, err := schema.ReadAndParseSchema(filepath.Join(dir, name+".json"))
		if err != nil {
			continue
		}
		defs, _ := raw["definitions"].(map[string]any)
		def, _ := defs[name].(map[string]any)
		anyOf, _ := def["anyOf"].([]any)
		hasIDRef := false
		for _, item := range anyOf {
			if m, ok := item.(map[string]any); ok {
				if ref, _ := m["$ref"].(string); strings.Contains(ref, "idRef") {
					hasIDRef = true
				}
			}
		}
		return hasIDRef
	}
	// Unknown schemas keep the previous behavior.
	return true
}

// isResourceRef reports whether ref points at a standalone resource (a schema whose file
// name matches the definition name) or at a resource collection: a reference to either is a
// link carrying the resource's @odata.id.
func isResourceRef(ref string) bool {
	if ref == "" || strings.Contains(ref, "odata-v4.json") || strings.Contains(ref, "Resource.json") {
		return false
	}
	if strings.Contains(ref, "Collection") {
		return true
	}
	schemaName := extractSchemaNameFromRef(ref)
	typeName := extractRefType(ref)
	return schemaName != "" && schemaName == typeName && isResource(typeName)
}

// IsResourceLinkProperty reports whether a property or action parameter is a reference to a
// resource or resource collection (or an array of them), as opposed to a URI-valued string. Such
// a value is exchanged as a link object ({"@odata.id": ...}), not as a bare string.
func IsResourceLinkProperty(prop *schema.JSONProperty) bool {
	if prop.Ref != "" {
		return isResourceRef(prop.Ref)
	}
	return prop.Items != nil && isResourceRef(prop.Items.Ref)
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
				return isResource(typeName)
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
				return isResource(typeName)
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
