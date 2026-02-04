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
	ImportGroups [][]string
	Enums        []*EnumData
	Structs      []*StructData
	Release      string
	Title        string
	SchemaID     string
}

// EnumData represents an enum for the template
type EnumData struct {
	Name        string
	Description string
	Values      []*EnumValueData
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
	Name            string
	JSONName        string
	Description     string
	TargetField     string
	Parameters      []*ActionParameterData
	ResponseType    string
	UseParamStruct  bool   // true when action has >3 parameters
	ParamStructName string // e.g., "CertificateServiceGenerateCSRParameters"
}

// ActionParameterData represents an action parameter for the template
type ActionParameterData struct {
	Name             string
	Type             string
	Description      string
	Required         bool
	Ordinal          int
	OriginalName     string
	FieldName        string // PascalCase Go field name for parameter structs
	FieldDescription string // Pre-formatted godoc comment for parameter struct fields
	JSONTag          string // JSON struct tag (with omitempty for optional)
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
	tmpl        *template.Template
	packageType schema.PackageType
	sfPrefix    bool              // True if type names should be prefixed with "SF"
	typeRenames map[string]string // Map of original type name to new name
}

// applySFPrefix prepends "SF" to a name if sfPrefix is enabled.
func (g *Generator) applySFPrefix(name string) string {
	if g.sfPrefix {
		return "SF" + name
	}
	return name
}

// applyTypeRename applies any configured type renames.
func (g *Generator) applyTypeRename(name string) string {
	if g.typeRenames != nil {
		if newName, ok := g.typeRenames[name]; ok {
			return newName
		}
	}
	return name
}

// normalizeTypeForPackage normalizes a type name for the current package context.
func (g *Generator) normalizeTypeForPackage(typeName string) string {
	// For gofish package, add schemas. prefix for types
	if g.packageType == "gofish" && !strings.Contains(typeName, ".") {
		return "schemas." + typeName
	}

	// Strip any existing package prefixes for same-package types
	if after, ok := strings.CutPrefix(typeName, "common."); ok {
		return after
	}

	return typeName
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

// Generate generates Go source code from definitions.
// sfPrefix indicates if type/const/function names should be prefixed with "SF"
// (used for Swordfish schemas that conflict with Redfish names).
// typeRenames maps original type names to new names (e.g., "ServiceRoot" -> "Service").
func (g *Generator) Generate(objectName string, packageType schema.PackageType, definitions []*schema.Definition, sfPrefix bool, typeRenames map[string]string) (string, error) {
	g.packageType = packageType
	g.sfPrefix = sfPrefix
	g.typeRenames = typeRenames
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

	// Set release and title from definitions.
	// Prefer definitions that have Release set (from versioned schema) since they
	// have the most current metadata. The main type by objectName may be from the
	// base schema which lacks release info.
	for _, def := range definitions {
		// Prefer definition with Release (indicates it's from versioned schema)
		if def.Release != "" {
			data.Release = def.Release
			data.Title = def.Title
			data.SchemaID = def.SchemaID
			break
		}
	}
	// Fall back to any definition with Title if no Release found
	if data.Title == "" {
		for _, def := range definitions {
			if def.Title != "" {
				data.Title = def.Title
				break
			}
		}
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

	for _, sd := range data.Structs {
		if (sd.IsEntity && len(sd.ReadWriteProperties) > 0) || sd.HasActions || sd.HasLinks || sd.HasLinkProperties {
			needsJSON = true
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
	// gofish package needs schemas import for referenced types
	if data.Package == "gofish" {
		groups = append(groups, []string{"github.com/stmcginnis/gofish/schemas"})
	}

	return groups
}

// buildEnumData converts a Definition to EnumData
func (g *Generator) buildEnumData(def *schema.Definition) *EnumData {
	enumName := g.applyTypeRename(g.applySFPrefix(def.Name))
	ed := &EnumData{
		Name:        enumName,
		Description: def.Description,
	}

	for _, ev := range def.EnumValues {
		// Make sure it's public (iSCSI > ISCSI)
		name := strings.ToUpper(ev.Name[:1]) + ev.Name[1:]
		name = makeValidIdentifier(name)

		// When SF prefix is applied, constants have the original type name as suffix.
		// Replace the suffix with the prefixed type name.
		// e.g., "RAID10RAIDType" -> "RAID10SFRAIDType"
		if g.sfPrefix && strings.HasSuffix(name, def.Name) {
			name = strings.TrimSuffix(name, def.Name) + enumName
		}

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

// goReservedWords is the set of Go language reserved keywords.
var goReservedWords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
}

// escapeReservedWord adds an underscore suffix if the name is a Go reserved word.
func escapeReservedWord(name string) string {
	if goReservedWords[name] {
		return name + "_"
	}
	return name
}

// buildStructData converts a Definition to StructData
func (g *Generator) buildStructData(def *schema.Definition, isMainType bool) *StructData {
	// Sort for deterministic output
	sort.Strings(def.ReadWriteProperties)

	structName := g.applyTypeRename(g.applySFPrefix(def.Name))
	sd := &StructData{
		Name:                structName,
		Description:         def.Description,
		IsEntity:            def.IsEntity,
		IsMainType:          isMainType,
		ReceiverName:        g.getReceiverName(structName),
		ReadWriteProperties: def.ReadWriteProperties,
	}

	// Build a map of link names to their Link objects to detect duplicates
	// Properties that duplicate non-deprecated Links should not generate their own getter
	linksByName := make(map[string]*schema.Link)
	for _, link := range def.Links {
		linksByName[link.Name] = link
	}

	for _, prop := range def.Properties {
		// Skip link properties that duplicate entries in the Links section,
		// UNLESS the Links version is deprecated (then the property is the current location)
		if prop.IsLink {
			if link, exists := linksByName[prop.Name]; exists && !link.Deprecated {
				continue
			}
		}

		typeStr := g.buildTypeString(prop)
		linkType := ""

		// Link properties store URIs as strings, but need the target type for getters
		if prop.IsLink {
			linkType = g.normalizeTypeForPackage(prop.Type)

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
			fieldName = escapeReservedWord(fieldName)
			prop.JSONTag = ""
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
				Name:             param.Name,
				Type:             param.Type,
				Description:      param.Description,
				Required:         param.Required,
				Ordinal:          param.Ordinal,
				OriginalName:     param.OriginalName,
				FieldName:        param.FieldName,
				FieldDescription: param.FieldDescription,
			}
			ad.Parameters = append(ad.Parameters, pd)
		}

		if len(ad.Parameters) > 3 {
			ad.UseParamStruct = true
			ad.ParamStructName = structName + action.Name + "Parameters"
			for _, pd := range ad.Parameters {
				if pd.Required {
					pd.JSONTag = fmt.Sprintf("`json:%q`", pd.OriginalName)
				} else {
					pd.JSONTag = fmt.Sprintf("`json:%q`", pd.OriginalName+",omitempty")
				}
			}
		}

		sd.Actions = append(sd.Actions, ad)
	}
	sd.HasActions = len(sd.Actions) > 0

	// Sort links by name for deterministic ordering
	sort.Slice(def.Links, func(i, j int) bool {
		return def.Links[i].Name < def.Links[j].Name
	})

	// Build a set of property names that are links, to detect when a
	// deprecated Links entry has been superseded by a top-level property
	propLinkNames := make(map[string]bool)
	for _, prop := range def.Properties {
		if prop.IsLink {
			propLinkNames[prop.Name] = true
		}
	}

	// Build link data
	for _, link := range def.Links {
		// Skip deprecated links that have been superseded by a top-level property
		if link.Deprecated && propLinkNames[link.Name] {
			continue
		}

		linkType := link.Type

		// Map idRef to Entity (generic reference type)
		if linkType == "idRef" {
			linkType = "Entity"
		}

		// Handle package prefixes for link types
		linkType = g.normalizeTypeForPackage(linkType)

		fieldName := strings.ToLower(link.Name[:1]) + link.Name[1:]
		fieldName = escapeReservedWord(fieldName)

		ld := &LinkData{
			Name:        link.Name,
			JSONName:    link.JSONName,
			Description: link.Description,
			FieldName:   fieldName,
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
