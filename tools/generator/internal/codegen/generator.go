//
// SPDX-License-Identifier: BSD-3-Clause
//

package codegen

import (
	_ "embed"
	"fmt"
	"os/exec"
	"strings"
	"text/template"

	"github.com/stmcginnis/gofish/tools/generator/internal/parser"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

//go:embed source.tmpl
var sourceTemplate string

// TemplateData holds the data for template rendering
type TemplateData struct {
	Package      string
	NeedsReflect bool
	Enums        []*EnumData
	Structs      []*StructData
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
	Name        string
	JSONName    string
	Description string
	TargetField string
	Parameters  []*ActionParameterData
}

// ActionParameterData represents an action parameter for the template
type ActionParameterData struct {
	Name        string
	Type        string
	Description string
	Required    bool
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
func (g *Generator) Generate(objectName string, packageType schema.PackageType, definitions []*schema.Definition) (string, error) {
	g.packageType = packageType
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

// buildEnumData converts a Definition to EnumData
func (g *Generator) buildEnumData(def *schema.Definition) *EnumData {
	ed := &EnumData{
		Name:         def.Name,
		Description:  def.Description,
		EnumTypeName: def.Name,
	}

	for _, ev := range def.EnumValues {
		evd := &EnumValueData{
			Name:        ev.Name,
			Value:       ev.Value,
			Description: ev.Description,
		}
		ed.Values = append(ed.Values, evd)
	}

	return ed
}

// buildStructData converts a Definition to StructData
func (g *Generator) buildStructData(def *schema.Definition, isMainType bool) *StructData {
	sd := &StructData{
		Name:                def.Name,
		Description:         def.Description,
		IsEntity:            def.IsEntity,
		IsMainType:          isMainType,
		ReceiverName:        g.getReceiverName(def.Name),
		ReadWriteProperties: def.ReadWriteProperties,
	}

	for _, prop := range def.Properties {
		typeStr := g.buildTypeString(prop)
		linkType := ""

		// Link properties store URIs as strings, but need the target type for getters
		if prop.IsLink {
			linkType = prop.Type

			// Determine package prefix for link target
			skipGetter := false
			switch g.packageType {
			case schema.PackageCommon:
				// Strip common. prefix if present
				if after, ok := strings.CutPrefix(linkType, "common."); ok {
					linkType = after
				} else if !strings.Contains(linkType, ".") {
					// No package prefix - need to determine where this type lives
					// Check if it's a known common type
					if !parser.IsCommonType(linkType) {
						// Cross-package link from common to redfish/swordfish
						// Cannot generate typed getter due to circular dependency
						skipGetter = true
						prop.IsPrivate = false
						prop.GetterMethod = ""
					}
				} else {
					// Has a package prefix and it's not common.
					// This is a cross-package reference - skip getter
					skipGetter = true
					prop.IsPrivate = false
					prop.GetterMethod = ""
				}
			case schema.PackageRedfish, schema.PackageSwordfish:
				// Add common. prefix to common types
				if parser.IsCommonType(linkType) && !strings.Contains(linkType, ".") {
					linkType = "common." + linkType
				}
				// Strip our own package prefix
				ownPrefix := string(g.packageType) + "."
				if after, ok := strings.CutPrefix(linkType, ownPrefix); ok {
					linkType = after
				}
			}

			if skipGetter {
				linkType = ""
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

	// Build action data
	for _, action := range def.Actions {
		ad := &ActionData{
			Name:        action.Name,
			JSONName:    action.JSONName,
			Description: action.Description,
			TargetField: strings.ToLower(action.Name[:1]) + action.Name[1:] + "Target",
		}

		for _, param := range action.Parameters {
			pd := &ActionParameterData{
				Name:        param.Name,
				Type:        param.Type,
				Description: param.Description,
				Required:    param.Required,
			}
			ad.Parameters = append(ad.Parameters, pd)
		}

		sd.Actions = append(sd.Actions, ad)
	}
	sd.HasActions = len(sd.Actions) > 0

	// Build link data
	for _, link := range def.Links {
		linkType := link.Type

		// Map idRef to Entity (generic reference type)
		if linkType == "idRef" {
			linkType = "Entity"
		}

		// Handle package prefixes for link types
		switch g.packageType {
		case schema.PackageCommon:
			// Strip common. prefix if present
			if after, ok := strings.CutPrefix(linkType, "common."); ok {
				linkType = after
			}
		case schema.PackageRedfish, schema.PackageSwordfish:
			// Add common. prefix to common types
			if parser.IsCommonType(linkType) && !strings.Contains(linkType, ".") {
				linkType = "common." + linkType
			}
			// Strip our own package prefix if present
			ownPrefix := string(g.packageType) + "."
			if after, ok := strings.CutPrefix(linkType, ownPrefix); ok {
				linkType = after
			}
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
		return string(name[0])
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
