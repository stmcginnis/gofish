//
// SPDX-License-Identifier: BSD-3-Clause
//

package config

import "strings"

// CommonNameChanges maps JSON property names to Go field names
var CommonNameChanges = map[string]string{
	"Oem": "OEM",
	"Id":  "ID",
}

// CommonDescriptions provides standard descriptions for common properties
var CommonDescriptions = map[string]string{
	"Description":    "Description provides a description of this resource.",
	"Id":             "ID uniquely identifies the resource.",
	"Name":           "Name is the name of the resource or array element.",
	"@odata.context": "ODataContext is the odata context.",
	"@odata.etag":    "ODataEtag is the odata etag.",
	"@odata.id":      "ODataID is the odata identifier.",
	"@odata.type":    "ODataType is the odata type.",
	"Identifier":     "Identifier shall be unique within the managed ecosystem.",
}

// CommonTypes maps property names to common Go types
var CommonTypes = map[string]string{
	"Status":             "common.Status",
	"Identifier":         "common.Identifier",
	"Location":           "common.Location",
	"PhysicalContext":    "common.PhysicalContext",
	"PhysicalSubContext": "common.PhysicalSubContext",
	"LogicalContext":     "common.LogicalContext",
	"ElectricalContext":  "common.ElectricalContext",
	"Protocol":           "common.Protocol",
	"Oem":                "json.RawMessage",
}

// EntityProperties are properties that indicate a type should embed common.Entity
// These properties should be skipped when generating structs that embed Entity
var EntityProperties = []string{
	"Name",
	"Id",
	"Description",
	"@odata.id",
	"@odata.etag",
	"@Message.ExtendedInfo",
}

// ExcludedDefinitions are definition names to skip during generation
// These are manually-maintained infrastructure types, not schema-generated
var ExcludedDefinitions = []string{
	"Actions",
	"OemActions",
	"Entity",             // Manually maintained in entity.go
	"Resource",           // Manually maintained in types.go
	"ResourceCollection", // Manually maintained in types.go
	"Item",               // Base type, manually maintained
}

// PropertyHasTarget checks if a definition has 'target' and 'title' properties (action definitions)
func IsActionDefinition(properties map[string]bool) bool {
	hasTarget := properties["target"]
	hasTitle := properties["title"]
	return hasTarget && hasTitle
}

// GetGoFieldName converts a JSON property name to a Go field name
func GetGoFieldName(jsonName string) string {
	if override, ok := CommonNameChanges[jsonName]; ok {
		return override
	}

	// Handle @odata.* properties (including Property@odata.count pattern)
	if strings.Contains(jsonName, "@") {
		return convertODataName(jsonName)
	}

	// Handle properties with dots or hyphens - convert to camelCase
	if strings.ContainsAny(jsonName, ".-") {
		return convertToCamelCase(jsonName)
	}

	if strings.HasSuffix(jsonName, "Id") {
		// Convert properties ending in "Id" to end with "ID"
		// Example: "ProcessorId" -> "ProcessorID"
		return jsonName[:len(jsonName)-1] + "D"
	}

	return jsonName
}

// convertToCamelCase converts names with dots and hyphens to camelCase
// Examples: "some.property" -> "SomeProperty", "some-property" -> "SomeProperty"
func convertToCamelCase(name string) string {
	// Split on dots and hyphens
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '.' || r == '-'
	})

	// Capitalize first letter of each part
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return strings.Join(parts, "")
}

// convertODataName converts @odata.* names to Go field names
func convertODataName(name string) string {
	// @odata.id -> ODataID
	// @odata.type -> ODataType
	// @odata.context -> ODataContext
	// @odata.etag -> ODataEtag
	// Members@odata.count -> MembersCount

	if name == "@odata.id" {
		return "ODataID"
	}
	if name == "@odata.type" {
		return "ODataType"
	}
	if name == "@odata.context" {
		return "ODataContext"
	}
	if name == "@odata.etag" {
		return "ODataEtag"
	}

	// Handle property@odata.count pattern -> PropertyCount
	if len(name) > 12 && name[len(name)-12:] == "@odata.count" {
		prefix := name[:len(name)-12]
		return prefix + "Count"
	}

	return "OData" + capitalizeFirst(name[7:])
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
