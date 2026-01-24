//
// SPDX-License-Identifier: BSD-3-Clause
//

package codegen

import (
	_ "embed"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

//go:embed source.tmpl
var sourceTemplate string

// TemplateData holds the data for template rendering
type TemplateData struct {
	Package      string
	NeedsReflect bool
	ImportGroups [][]string
	Enums        []*EnumData
	Structs      []*StructData
	Release      string
	Title        string
}

// EnumData represents an enum for the template
type EnumData struct {
	Name         string
	Description  string
	Values       []*EnumValueData
	EnumTypeName string
}

// EnumValueData represents an enum value for the template
type EnumValueData struct {
	Name        string
	Value       string
	Description string
}

// StructData represents a struct for the template
type StructData struct {
	Name                string
	Description         string
	IsEntity            bool
	IsMainType          bool
	ReceiverName        string
	Properties          []*PropertyData
	ReadWriteProperties []string
	Actions             []*ActionData
	Links               []*LinkData
	HasActions          bool
	HasLinks            bool
	HasLinkProperties   bool
}

// ActionData represents an action for the template
type ActionData struct {
	Name         string
	JSONName     string
	Description  string
	TargetField  string
	Parameters   []*ActionParameterData
	ResponseType string
}

// ActionParameterData represents an action parameter for the template
type ActionParameterData struct {
	Name         string
	Type         string
	Description  string
	Required     bool
	Ordinal      int
	OriginalName string
}

// LinkData represents a link for the template
type LinkData struct {
	Name        string
	JSONName    string
	Description string
	FieldName   string
	Type        string
	IsArray     bool
}

// PropertyData represents a property for the template
type PropertyData struct {
	Name         string
	TypeString   string
	Description  string
	JSONName     string
	JSONTag      string
	IsPrivate    bool
	IsLink       bool
	LinkType     string // The target type for link properties
	GetterMethod string
	IsArray      bool
	IsCollection bool // Indicates if this is a collection link
	VersionAdded string
}

// Generator handles code generation
type Generator struct {
	tmpl                 *template.Template
	packageType          schema.PackageType
	isSwordfishInCommon  bool // True if this is a Swordfish type being placed in common
}

// packagePrefix returns the package prefix string (e.g., "redfish.")
func (g *Generator) packagePrefix() string {
	return string(g.packageType) + "."
}

// normalizeTypeForPackage normalizes a type name for the current package context.
// It strips/adds package prefixes as needed and returns whether the type is external
// (i.e., a cross-package reference that may need special handling).
func (g *Generator) normalizeTypeForPackage(typeName string) (normalized string, isExternal bool) {
	normalized = typeName

	switch g.packageType {
	case schema.PackageCommon:
		// Strip common. prefix if present
		if after, ok := strings.CutPrefix(normalized, "common."); ok {
			normalized = after
		} else if strings.Contains(normalized, ".") {
			// Has a non-common package prefix - this is a cross-package reference
			isExternal = true
		} else if !g.isSwordfishInCommon && !schema.IsInfrastructureType(normalized) {
			// No package prefix and not an infrastructure type
			// This is a cross-package link from common to redfish/swordfish
			// (skip when isSwordfishInCommon since referenced types are also in common)
			isExternal = true
		}

	case schema.PackageRedfish, schema.PackageSwordfish:
		// Add common. prefix to infrastructure types without a prefix
		if schema.IsInfrastructureType(normalized) && !strings.Contains(normalized, ".") {
			normalized = "common." + normalized
		}
		// Strip our own package prefix
		if after, ok := strings.CutPrefix(normalized, g.packagePrefix()); ok {
			normalized = after
		}
	}

	return normalized, isExternal
}

// NewGenerator creates a new Generator
func NewGenerator() (*Generator, error) {
	tmpl, err := template.New("source").Parse(sourceTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &Generator{
		tmpl: tmpl,
	}, nil
}

// Generate generates Go source code from definitions
// isSwordfishInCommon indicates if this is a Swordfish type being placed in common
// (used to skip generating methods that would create import cycles)
func (g *Generator) Generate(objectName string, packageType schema.PackageType, definitions []*schema.Definition, isSwordfishInCommon bool) (string, error) {
	g.packageType = packageType
	g.isSwordfishInCommon = isSwordfishInCommon
	data := &TemplateData{
		Package: string(packageType),
	}

	// Find the main type (matching objectName)
	var mainType *schema.Definition
	for _, def := range definitions {
		if def.OriginalName == objectName {
			mainType = def
			break
		}
	}

	// Set release and title from main type if available
	if mainType != nil {
		data.Release = mainType.Release
		data.Title = mainType.Title
	}

	// Sort definitions by name for deterministic output
	// Main type, then enums, then other structs
	sort.Slice(definitions, func(i, j int) bool {
		if definitions[i] == mainType {
			return true
		}
		if definitions[j] == mainType {
			return false
		}
		if definitions[i].IsEnum != definitions[j].IsEnum {
			return definitions[i].IsEnum
		}

		return definitions[i].Name < definitions[j].Name
	})

	// Process all definitions
	for _, def := range definitions {
		if def.IsEnum {
			enumData := g.buildEnumData(def)
			data.Enums = append(data.Enums, enumData)
		} else {
			structData := g.buildStructData(def, def == mainType)
			data.Structs = append(data.Structs, structData)

			// Check if we need reflect package
			if len(def.ReadWriteProperties) > 0 && def.IsEntity {
				data.NeedsReflect = true
			}
		}
	}

	sort.Slice(data.Enums, func(i, j int) bool {
		return data.Enums[i].Name < data.Enums[j].Name
	})

	data.ImportGroups = buildImportGroups(data)

	// Execute template
	var buf strings.Builder
	if err := g.tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	output := buf.String()

	// Format with gofmt
	formatted, err := g.formatGoCode(output)
	if err != nil {
		// Return unformatted if formatting fails
		return output, nil
	}

	return formatted, nil
}

func buildImportGroups(data *TemplateData) [][]string {
	needsJSON := false
	needsCommon := false

	for _, sd := range data.Structs {
		if (sd.IsEntity && len(sd.ReadWriteProperties) > 0) || sd.HasActions || sd.HasLinks || sd.HasLinkProperties {
			needsJSON = true
		}
		if sd.IsEntity {
			needsCommon = true
		}
		for _, prop := range sd.Properties {
			if prop.TypeString == "json.RawMessage" {
				needsJSON = true
			}
		}
	}

	var groups [][]string
	if needsJSON {
		groups = append(groups, []string{"encoding/json"})
	}
	if needsCommon && data.Package != string(schema.PackageCommon) {
		groups = append(groups, []string{"github.com/stmcginnis/gofish/common"})
	}

	return groups
}

// buildEnumData converts a Definition to EnumData
func (g *Generator) buildEnumData(def *schema.Definition) *EnumData {
	ed := &EnumData{
		Name:         def.Name,
		Description:  def.Description,
		EnumTypeName: def.Name,
	}

	for _, ev := range def.EnumValues {
		// Make sure it's public (iSCSI > ISCSI)
		name := strings.ToUpper(ev.Name[:1]) + ev.Name[1:]
		name = makeValidIdentifier(name)
		evd := &EnumValueData{
			Name:        name,
			Value:       ev.Value,
			Description: ev.Description,
		}
		ed.Values = append(ed.Values, evd)
	}

	return ed
}

func makeValidIdentifier(name string) string {
	if name == "" {
		return "Value"
	}

	first := name[0]
	if (first >= 'A' && first <= 'Z') || (first >= 'a' && first <= 'z') || first == '_' {
		return name
	}

	return "V" + name
}

// buildStructData converts a Definition to StructData
func (g *Generator) buildStructData(def *schema.Definition, isMainType bool) *StructData {
	// Sort for deterministic output
	sort.Strings(def.ReadWriteProperties)

	sd := &StructData{
		Name:                def.Name,
		Description:         def.Description,
		IsEntity:            def.IsEntity,
		IsMainType:          isMainType,
		ReceiverName:        g.getReceiverName(def.Name),
		ReadWriteProperties: def.ReadWriteProperties,
	}

	// Build a set of link names to detect duplicates
	// Properties that duplicate Links should not generate their own getter
	linkNames := make(map[string]bool)
	for _, link := range def.Links {
		linkNames[link.Name] = true
	}

	for _, prop := range def.Properties {
		// Skip link properties that duplicate entries in the Links section
		// The Links getter will handle accessing these resources
		if prop.IsLink && linkNames[prop.Name] {
			continue
		}

		typeStr := g.buildTypeString(prop)
		linkType := ""

		// Link properties store URIs as strings, but need the target type for getters
		if prop.IsLink {
			var isExternal bool
			linkType, isExternal = g.normalizeTypeForPackage(prop.Type)

			// Cross-package references cannot have typed getters due to circular dependencies
			if isExternal {
				linkType = ""
				prop.IsPrivate = false
				prop.GetterMethod = ""
			}

			// Link properties are stored as strings
			if prop.IsArray {
				typeStr = "[]string"
			} else {
				typeStr = "string"
			}
		}

		// Calculate field name after potentially modifying IsPrivate
		fieldName := prop.Name
		// Private fields start with lowercase
		if prop.IsPrivate {
			fieldName = strings.ToLower(fieldName[:1]) + fieldName[1:]
		}

		pd := &PropertyData{
			Name:         fieldName,
			TypeString:   typeStr,
			Description:  prop.Description,
			JSONName:     prop.JSONName,
			JSONTag:      prop.JSONTag,
			IsPrivate:    prop.IsPrivate,
			IsLink:       prop.IsLink,
			LinkType:     linkType,
			GetterMethod: prop.GetterMethod,
			IsArray:      prop.IsArray,
			IsCollection: prop.IsCollection,
			VersionAdded: prop.VersionAdded,
		}
		sd.Properties = append(sd.Properties, pd)

		// Check if this property is a link with a getter
		if pd.IsLink && pd.LinkType != "" {
			sd.HasLinkProperties = true
		}
	}

	// Sort actions by name for deterministic output
	sort.Slice(def.Actions, func(i, j int) bool {
		return def.Actions[i].Name < def.Actions[j].Name
	})

	// Build action data
	for _, action := range def.Actions {
		ad := &ActionData{
			Name:         action.Name,
			JSONName:     action.JSONName,
			Description:  action.Description,
			TargetField:  strings.ToLower(action.Name[:1]) + action.Name[1:] + "Target",
			ResponseType: action.ResponseType,
		}

		// Sort parameters by the order they are defined in the schema
		sort.Slice(action.Parameters, func(i, j int) bool {
			return action.Parameters[i].Ordinal < action.Parameters[j].Ordinal
		})

		for _, param := range action.Parameters {
			pd := &ActionParameterData{
				Name:         param.Name,
				Type:         param.Type,
				Description:  param.Description,
				Required:     param.Required,
				Ordinal:      param.Ordinal,
				OriginalName: param.OriginalName,
			}
			ad.Parameters = append(ad.Parameters, pd)
		}

		sd.Actions = append(sd.Actions, ad)
	}
	sd.HasActions = len(sd.Actions) > 0

	// Sort links by name for deterministic ordering
	sort.Slice(def.Links, func(i, j int) bool {
		return def.Links[i].Name < def.Links[j].Name
	})

	// Build link data
	for _, link := range def.Links {
		linkType := link.Type

		// Map idRef to Entity (generic reference type)
		if linkType == "idRef" {
			linkType = "Entity"
		}

		// Handle package prefixes for link types
		linkType, isExternal := g.normalizeTypeForPackage(linkType)
		if isExternal {
			continue
		}

		ld := &LinkData{
			Name:        link.Name,
			JSONName:    link.JSONName,
			Description: link.Description,
			FieldName:   strings.ToLower(link.Name[:1]) + link.Name[1:],
			Type:        linkType,
			IsArray:     link.IsArray,
		}
		sd.Links = append(sd.Links, ld)
	}
	sd.HasLinks = len(sd.Links) > 0

	return sd
}

// buildTypeString builds the type string for a property
func (g *Generator) buildTypeString(prop *schema.Property) string {
	typeStr := strings.ReplaceAll(prop.Type, "_", "")

	// If generating for common package, strip "common." prefix from types
	if g.packageType == schema.PackageCommon && strings.HasPrefix(typeStr, "common.") {
		typeStr = strings.TrimPrefix(typeStr, "common.")
	}

	if prop.IsArray {
		typeStr = "[]" + typeStr
	}

	if prop.IsPointer {
		typeStr = "*" + typeStr
	}

	return typeStr
}

// getReceiverName generates a receiver name from a type name
func (g *Generator) getReceiverName(typeName string) string {
	// Use lowercase first letter(s)
	if len(typeName) == 0 {
		return "x"
	}

	// Find first lowercase or end of acronym
	name := strings.ToLower(typeName)
	if len(name) > 0 {
		ret := string(name[0])
		if ret == "b" {
			// Byte parameter to unmarshaling is already "b"
			ret += string(name[1])
		}
		return ret
	}

	return "x"
}

// formatGoCode formats Go code using gofmt
func (g *Generator) formatGoCode(code string) (string, error) {
	cmd := exec.Command("gofmt", "-s")
	cmd.Stdin = strings.NewReader(code)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("gofmt failed: %w\n%s", err, string(output))
	}

	return string(output), nil
}
