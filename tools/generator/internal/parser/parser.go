//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/config"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// Parser handles parsing of JSON Schema files
type Parser struct {
	typeMapper *TypeMapper
	schemaDir  string
	rawSchema  []byte // Raw JSON for extracting key order
}

// NewParser creates a new Parser
func NewParser(schemaDir string) *Parser {
	return &Parser{
		typeMapper: NewTypeMapper(),
		schemaDir:  schemaDir,
	}
}

// ParseSchemaWithBase parses both base and versioned schema files, merging definitions
func (p *Parser) ParseSchemaWithBase(baseSchemaFile, versionedSchemaFile string) ([]*schema.Definition, error) {
	// Parse base schema to get base definitions
	var baseDefsMap map[string]*schema.Definition
	if _, err := os.Stat(baseSchemaFile); err == nil {
		baseDefs, err := p.ParseSchema(baseSchemaFile)
		if err == nil {
			baseDefsMap = schema.BuildDefinitionMap(baseDefs, false)
		}
	}
	if baseDefsMap == nil {
		baseDefsMap = make(map[string]*schema.Definition)
	}

	// Parse versioned schema
	versionedDefs, err := p.ParseSchema(versionedSchemaFile)
	if err != nil {
		return nil, err
	}

	// Build map of versioned definitions
	versionedDefsMap := schema.BuildDefinitionMap(versionedDefs, false)

	// Merge: start with all base definitions, then add/override with versioned
	result := []*schema.Definition{}

	// Add all base definitions that aren't in versioned schema
	for name, baseDef := range baseDefsMap {
		if _, inVersioned := versionedDefsMap[name]; !inVersioned {
			result = append(result, baseDef)
		}
	}

	// Add all versioned definitions (which override base)
	result = append(result, versionedDefs...)

	return result, nil
}

// ParseSchema parses a JSON schema file and returns definitions
func (p *Parser) ParseSchema(schemaFile string) ([]*schema.Definition, error) {
	data, err := os.ReadFile(schemaFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	// Store raw schema for extracting key order
	p.rawSchema = data

	var rawSchema map[string]any
	if err := json.Unmarshal(data, &rawSchema); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	definitions := []*schema.Definition{}

	// Extract definitions section
	defsMap, ok := rawSchema["definitions"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("no definitions found in schema")
	}

	// Extract version from filename
	version := extractVersion(filepath.Base(schemaFile))

	// Extract release and title from top-level schema
	release := ""
	if rel, ok := rawSchema["release"].(string); ok {
		release = rel
	}
	title := ""
	if t, ok := rawSchema["title"].(string); ok {
		title = t
	}
	schemaID := ""
	if id, ok := rawSchema["$id"].(string); ok {
		schemaID = id
	}

	for defName, defData := range defsMap {
		// Skip excluded definitions
		if p.shouldSkipDefinition(defName) {
			continue
		}

		defMap, ok := defData.(map[string]any)
		if !ok {
			continue
		}

		// Check if this is an enum
		if enumVals, hasEnum := defMap["enum"].([]any); hasEnum {
			def := p.parseEnumDefinition(defName, defMap, enumVals, version)
			def.Release = release
			def.Title = title
			def.SchemaID = schemaID
			definitions = append(definitions, def)
			continue
		}

		// Check if this is an object type (explicit "type": "object" or implicit via "properties")
		typeVal, hasType := defMap["type"]
		_, hasProps := defMap["properties"].(map[string]any)
		isObject := (hasType && typeVal == "object") || hasProps

		if isObject {
			// Skip action definitions (have target and title properties)
			if props, ok := defMap["properties"].(map[string]any); ok {
				if p.isActionDefinition(props) {
					continue
				}
			}

			def := p.parseObjectDefinition(defName, defMap, version)
			def.Release = release
			def.Title = title
			def.SchemaID = schemaID
			definitions = append(definitions, def)
		}
	}

	// Second pass: parse Actions and Links for definitions that have them
	for _, def := range definitions {
		// Get the original definition map to check for Actions/Links properties
		if defData, ok := defsMap[def.OriginalName]; ok {
			if defMap, ok := defData.(map[string]any); ok {
				if propsData, ok := defMap["properties"].(map[string]any); ok {
					// Parse Actions if this definition has an Actions property
					if _, hasActions := propsData["Actions"]; hasActions {
						if actionsDefData, ok := defsMap["Actions"]; ok {
							if actionsDefMap, ok := actionsDefData.(map[string]any); ok {
								def.Actions = p.parseActions(actionsDefMap, defsMap)
							}
						}
					}

					// Parse Links if this definition has a Links property
					if _, hasLinks := propsData["Links"]; hasLinks {
						if linksDefData, ok := defsMap["Links"]; ok {
							if linksDefMap, ok := linksDefData.(map[string]any); ok {
								def.Links = p.parseLinks(linksDefMap)
							}
						}
					}
				}
			}
		}
	}

	return definitions, nil
}

// parseEnumDefinition parses an enum type definition
func (p *Parser) parseEnumDefinition(name string, defMap map[string]any, enumVals []any, version string) *schema.Definition {
	def := &schema.Definition{
		Name:         cleanIdentifier(name),
		OriginalName: name,
		IsEnum:       true,
		Version:      version,
		Description:  p.formatEnumDescription(cleanIdentifier(name), defMap),
	}

	// Parse enum values
	enumDescs := make(map[string]string)
	if ed, ok := defMap["enumDescriptions"].(map[string]any); ok {
		for k, v := range ed {
			if str, ok := v.(string); ok {
				enumDescs[k] = str
			}
		}
	}
	if eld, ok := defMap["enumLongDescriptions"].(map[string]any); ok {
		for k, v := range eld {
			if str, ok := v.(string); ok {
				enumDescs[k] = str
			}
		}
	}

	for _, enumVal := range enumVals {
		if strVal, ok := enumVal.(string); ok {
			ev := &schema.EnumValue{
				Name:        cleanIdentifier(strVal) + def.Name,
				Value:       strVal,
				Description: p.formatEnumValueDescription(strVal, def.Name, enumDescs[strVal]),
			}
			def.EnumValues = append(def.EnumValues, ev)
		}
	}

	return def
}

// parseObjectDefinition parses an object type definition
func (p *Parser) parseObjectDefinition(name string, defMap map[string]any, version string) *schema.Definition {
	cleanName := cleanIdentifier(name)
	def := &schema.Definition{
		Name:         cleanName,
		OriginalName: name,
		IsEnum:       false,
		Version:      version,
		Description:  p.formatTypeDescription(cleanName, defMap),
	}

	// Parse properties
	propsMap, hasProps := defMap["properties"].(map[string]any)
	if !hasProps {
		return def
	}

	// Check if this is an Entity (has Id, Name, @odata.id)
	def.IsEntity = p.isEntity(propsMap)

	for propName, propData := range propsMap {
		// Skip entity base properties
		if def.IsEntity && p.isEntityProperty(propName) {
			continue
		}

		// Skip Actions and Links - they will be handled separately
		if propName == "Actions" || propName == "Links" {
			continue
		}

		// Skip OemActions - these are empty placeholders
		if propName == "OemActions" {
			continue
		}

		propMap, ok := propData.(map[string]any)
		if !ok {
			continue
		}

		prop := p.parseProperty(propName, propMap)
		if prop != nil {
			def.Properties = append(def.Properties, prop)

			// Track read-write properties
			if !prop.IsReadOnly && !slices.Contains(config.ExcludeReadWriteProperties, prop.Name) {
				def.ReadWriteProperties = append(def.ReadWriteProperties, prop.JSONName)
			}
		}
	}

	// Sort properties by name for consistency
	sort.Slice(def.Properties, func(i, j int) bool {
		return def.Properties[i].Name < def.Properties[j].Name
	})

	return def
}

// parseProperty parses a single property
func (p *Parser) parseProperty(propName string, propMap map[string]any) *schema.Property {
	// Convert to JSONProperty for type mapping
	propJSON := mapToJSONProperty(propMap)

	goType, isPointer, isArray := p.typeMapper.MapType(propName, propJSON)

	// Check if property is deprecated
	isDeprecated := false
	deprecationMsg := ""
	deprecationVersion := ""
	if deprecated, ok := propMap["deprecated"]; ok {
		// deprecated can be either a boolean or a string
		switch v := deprecated.(type) {
		case bool:
			isDeprecated = v
		case string:
			isDeprecated = true
			deprecationMsg = v
		}
	}
	// Get deprecation version if present
	if versionDep, ok := propMap["versionDeprecated"].(string); ok {
		deprecationVersion = convertVersionFormat(versionDep)
	}

	cleanName := cleanIdentifier(config.GetGoFieldName(propName))

	prop := &schema.Property{
		Name:         cleanName,
		JSONName:     propName,
		Type:         goType,
		Description:  p.formatPropertyDescription(propName, cleanName, propMap),
		IsPointer:    isPointer,
		IsArray:      isArray,
		IsReadOnly:   p.isReadOnly(propMap),
		IsLink:       goType != "string" && IsLinkProperty(propName, propJSON),
		IsCollection: IsCollectionProperty(propJSON),
		IsDeprecated: isDeprecated,
	}

	// Extract version_added if present
	if revisions, ok := propMap["Redfish.Revisions"].([]any); ok && len(revisions) > 0 {
		if rev, ok := revisions[0].(map[string]any); ok {
			if ver, ok := rev["Version"].(string); ok {
				prop.VersionAdded = convertVersionFormat(ver)
			}
		}
	}

	// More likely to have versionAdded in the schema
	if versionAdded, ok := propMap["versionAdded"].(string); ok {
		prop.VersionAdded = convertVersionFormat(versionAdded)
	}

	// Add version added to description
	if prop.VersionAdded != "" {
		prop.Description = appendCommentSection(
			prop.Description,
			fmt.Sprintf("\t// Version added: %s", prop.VersionAdded),
		)
	}

	// Add deprecation notice to description
	if isDeprecated {
		var wrappedDeprecation string

		// First line: Deprecated: v1.3.0
		versionNotice := "Deprecated"
		if deprecationVersion != "" {
			versionNotice = "Deprecated: " + deprecationVersion
		}

		if deprecationMsg != "" {
			// Version on first line, message on subsequent lines
			wrappedDeprecation = "\t// " + versionNotice + "\n"
			wrappedDeprecation += formatComment("", cleanDescription(deprecationMsg), "", false, "\t")
		} else {
			// Just the version
			wrappedDeprecation = "\t// " + versionNotice
		}

		// Append to existing description
		prop.Description = appendCommentSection(prop.Description, wrappedDeprecation)
	}

	prop.JSONTag = makeJSONTag(prop.Name, propName, isPointer)

	// Links and collections should be private with public getter
	if prop.IsLink || prop.IsCollection {
		prop.IsPrivate = true
		prop.GetterMethod = prop.Name
	}

	return prop
}

// Helper functions

func (p *Parser) shouldSkipDefinition(name string) bool {
	// Skip OemActions - these are parsed from RawData from OEM implementations
	// But don't skip enum types that end with "Actions" (e.g., WatchdogTimeoutActions)
	if name == "OemActions" {
		return true
	}

	// Skip Actions and Links definitions - they're handled specially
	if name == "Actions" || name == "Links" {
		return true
	}

	return slices.Contains(config.ExcludedDefinitions, name)
}

func (p *Parser) isActionDefinition(props map[string]any) bool {
	_, hasTarget := props["target"]
	_, hasTitle := props["title"]
	return hasTarget && hasTitle
}

func (p *Parser) isEntity(props map[string]any) bool {
	for propName := range props {
		if propName == "@odata.id" || propName == "@odata.type" {
			return true
		}
	}

	return false
}

func (p *Parser) isEntityProperty(propName string) bool {
	return slices.Contains(config.EntityProperties, propName)
}

// isReadOnly returns true unless "readonly": false is explicitly set.
// Properties without the readonly field are treated as read-only by default.
func (p *Parser) isReadOnly(propMap map[string]any) bool {
	if ro, ok := propMap["readonly"].(bool); ok {
		return ro
	}
	return true // default to read-only if not specified
}

func (p *Parser) formatTypeDescription(name string, defMap map[string]any) string {
	desc := p.getDescription(defMap)
	if desc == "" {
		return fmt.Sprintf("// %s represents the %s type.", name, name)
	}
	return formatComment(name, desc, "shall", false, "")
}

func (p *Parser) formatEnumDescription(name string, defMap map[string]any) string {
	desc := p.getDescription(defMap)
	if desc == "" {
		return ""
	}
	return formatComment(name, desc, "", true, "")
}

func (p *Parser) formatEnumValueDescription(value, enumType, desc string) string {
	constName := cleanIdentifier(value) + enumType
	if desc == "" {
		return fmt.Sprintf("\t// %s means %s", constName, value)
	}
	// Template adds indent to first line, but we need it on all lines
	return formatComment(constName, desc, "shall", false, "\t")
}

func (p *Parser) formatPropertyDescription(propName, goName string, propMap map[string]any) string {
	if stdDesc, ok := config.CommonDescriptions[propName]; ok {
		return "\t// " + stdDesc
	}

	desc := p.getDescription(propMap)
	if desc == "" {
		return fmt.Sprintf("\t// %s", goName)
	}
	return formatComment(goName, desc, "shall", false, "\t")
}

func (p *Parser) getDescription(m map[string]any) string {
	if ld, ok := m["longDescription"].(string); ok && ld != "" {
		return cleanDescription(ld)
	}
	if d, ok := m["description"].(string); ok {
		return cleanDescription(d)
	}
	return ""
}

// formatComment formats a comment with proper wrapping and indentation
// indent is the prefix to add to each line (e.g., "\t" for struct fields)
func formatComment(name, description, cutpoint string, isEnum bool, indent string) string {
	// Clean up description
	desc := cleanDescription(description)

	// Find cutpoint if specified
	if cutpoint != "" && strings.Contains(desc, cutpoint) {
		idx := strings.Index(desc, cutpoint)
		desc = desc[idx:]
	}

	prefix := name
	if isEnum {
		prefix = name + " is"
	}

	fullText := prefix + " " + desc

	// Word wrap at ~80 characters
	maxLineLength := 80 - len(indent) - 3 // 3 for "// "
	words := strings.Fields(fullText)
	lines := []string{}
	currentLine := ""

	for _, word := range words {
		if currentLine == "" {
			currentLine = word
		} else if len(currentLine)+len(word)+1 > maxLineLength {
			lines = append(lines, indent+"// "+currentLine)
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}
	if currentLine != "" {
		lines = append(lines, indent+"// "+currentLine)
	}

	return strings.Join(lines, "\n")
}

func appendCommentSection(base, section string) string {
	if base == "" {
		return section
	}
	return base + "\n\t//\n" + section
}

func cleanDescription(desc string) string {
	// Replace backticks with single quotes
	desc = strings.ReplaceAll(desc, "`", "'")
	// Collapse multiple spaces
	desc = regexp.MustCompile(`\s+`).ReplaceAllString(desc, " ")
	desc = strings.TrimSpace(desc)
	if strings.HasPrefix(desc, "A ") || strings.HasPrefix(desc, "An ") {
		desc = "is a" + desc[1:]
	}
	return desc
}

func cleanIdentifier(name string) string {
	// Remove all non-alphanumeric characters
	clean := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(name, "")

	// Replace some common things
	clean = strings.ReplaceAll(clean, "Fpga", "FPGA")
	clean = strings.ReplaceAll(clean, "Http", "HTTP")
	clean = strings.ReplaceAll(clean, "Json", "JSON")
	clean = strings.ReplaceAll(clean, "Dhcp", "DHCP")
	clean = strings.ReplaceAll(clean, "Dns", "DNS")
	clean = strings.ReplaceAll(clean, "Uri", "URI")
	if !strings.Contains(clean, "Identif") &&
		!strings.Contains(clean, "Idle") &&
		!strings.Contains(clean, "Ident") {
		clean = strings.ReplaceAll(clean, "Id", "ID")
	}

	return clean
}

func makeJSONTag(fieldName, jsonName string, isPointer bool) string {
	if fieldName == jsonName && !isPointer {
		return ""
	}

	omit := ""
	if isPointer {
		omit = ",omitempty"
	}

	name := ""
	if fieldName != jsonName {
		name = jsonName
	}

	jsonString := name + omit
	return fmt.Sprintf("`json:%q`", jsonString)
}

func extractVersion(filename string) string {
	// Extract version like v1_2_0 from filename
	re := regexp.MustCompile(`v(\d+)_(\d+)_(\d+)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) == 4 {
		return fmt.Sprintf("v%s_%s_%s", matches[1], matches[2], matches[3])
	}
	return ""
}

func convertVersionFormat(ver string) string {
	// Convert v1_3_0 to v1.3.0
	return strings.ReplaceAll(ver, "_", ".")
}

// mapToJSONProperty converts a generic map to JSONProperty
func mapToJSONProperty(m map[string]any) *schema.JSONProperty {
	data, _ := json.Marshal(m)
	var prop schema.JSONProperty
	json.Unmarshal(data, &prop)
	return &prop
}

// ResolveLatestVersion finds the latest version of a schema
func ResolveLatestVersion(baseFile, schemaDir string) (string, error) {
	data, err := os.ReadFile(baseFile)
	if err != nil {
		return "", err
	}

	var rawSchema map[string]any
	if err := json.Unmarshal(data, &rawSchema); err != nil {
		return "", err
	}

	// Get the base name (e.g., "LogService" from "LogService.json")
	baseName := strings.TrimSuffix(filepath.Base(baseFile), ".json")

	// Look in definitions
	defsMap, ok := rawSchema["definitions"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("no definitions found")
	}

	// Determine which definition to use for finding versioned files.
	// Check if root $ref points to a different definition (e.g., Capacity.json's
	// main type is CapacitySource, not Capacity).
	defName := baseName
	if rootRef, ok := rawSchema["$ref"].(string); ok {
		// Extract definition name from $ref like "#/definitions/CapacitySource"
		if strings.HasPrefix(rootRef, "#/definitions/") {
			defName = strings.TrimPrefix(rootRef, "#/definitions/")
		}
	}

	// Find the definition to check for anyOf references
	defData, ok := defsMap[defName]
	if !ok {
		// Try the base name if root $ref definition doesn't exist
		defData, ok = defsMap[baseName]
		if !ok {
			// This might be a utility schema (like IPAddresses.json) that doesn't have
			// a main definition matching the schema name. Try to find versioned files
			// by glob pattern instead.
			return resolveLatestVersionByGlob(baseName, schemaDir, baseFile)
		}
	}

	defMap, ok := defData.(map[string]any)
	if !ok {
		return "", fmt.Errorf("invalid definition format")
	}

	// Look for anyOf with $ref values
	anyOf, ok := defMap["anyOf"].([]any)
	if !ok {
		// Definition has no anyOf, try glob pattern to find versioned files
		return resolveLatestVersionByGlob(baseName, schemaDir, baseFile)
	}

	// Find all version references
	maxMajor, maxMinor, maxErrata := 0, 0, 0
	versionFile := ""

	for _, item := range anyOf {
		itemMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		ref, ok := itemMap["$ref"].(string)
		if !ok {
			continue
		}

		// Skip idRef
		if strings.Contains(ref, "idRef") {
			continue
		}

		// Extract version from ref (e.g., "LogService.v1_2_0.json#/definitions/LogService")
		re := regexp.MustCompile(`v(\d+)_(\d+)_(\d+)`)
		matches := re.FindStringSubmatch(ref)
		if len(matches) == 4 {
			major, _ := strconv.Atoi(matches[1])
			minor, _ := strconv.Atoi(matches[2])
			errata, _ := strconv.Atoi(matches[3])

			if schema.CompareVersions(major, minor, errata, maxMajor, maxMinor, maxErrata) {
				maxMajor, maxMinor, maxErrata = major, minor, errata
				// Extract filename from ref
				parts := strings.Split(ref, "#")
				if len(parts) > 0 {
					versionFile = filepath.Join(schemaDir, filepath.Base(parts[0]))
				}
			}
		}
	}

	if versionFile == "" {
		return baseFile, nil
	}

	return versionFile, nil
}

// resolveLatestVersionByGlob finds the latest versioned file by glob pattern.
// Used for utility schemas that don't have a main definition matching the schema name.
func resolveLatestVersionByGlob(baseName, schemaDir, baseFile string) (string, error) {
	// Find all versioned files matching the pattern
	pattern := filepath.Join(schemaDir, baseName+".v*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		// No versioned files found, use the base file
		return baseFile, nil
	}

	// Find the latest version
	maxMajor, maxMinor, maxErrata := 0, 0, 0
	versionFile := ""

	re := regexp.MustCompile(`\.v(\d+)_(\d+)_(\d+)\.json$`)
	for _, match := range matches {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) == 4 {
			major, _ := strconv.Atoi(submatches[1])
			minor, _ := strconv.Atoi(submatches[2])
			errata, _ := strconv.Atoi(submatches[3])

			if schema.CompareVersions(major, minor, errata, maxMajor, maxMinor, maxErrata) {
				maxMajor, maxMinor, maxErrata = major, minor, errata
				versionFile = match
			}
		}
	}

	if versionFile == "" {
		return baseFile, nil
	}

	return versionFile, nil
}

// parseActions parses the Actions definition for a resource
func (p *Parser) parseActions(actionsMap map[string]any, defsMap map[string]any) []*schema.Action {
	var actions []*schema.Action

	propsMap, ok := actionsMap["properties"].(map[string]any)
	if !ok {
		return actions
	}

	for actionName, actionData := range propsMap {
		// Skip Oem actions
		if strings.Contains(actionName, "Oem") {
			continue
		}

		// Actions are formatted like "#ResourceName.ActionName"
		if !strings.HasPrefix(actionName, "#") {
			continue
		}

		// Parse the action reference
		actionMap, ok := actionData.(map[string]any)
		if !ok {
			continue
		}

		action := &schema.Action{
			JSONName: actionName,
		}

		// Extract action name (e.g., "#Chassis.Reset" -> "Reset")
		parts := strings.Split(actionName, ".")
		if len(parts) == 2 {
			action.Name = parts[1]
		}

		// Get the action definition by following the $ref
		if ref, ok := actionMap["$ref"].(string); ok {
			// ref looks like "#/definitions/Reset"
			defName := strings.TrimPrefix(ref, "#/definitions/")
			if actionDefData, ok := defsMap[defName]; ok {
				if actionDefMap, ok := actionDefData.(map[string]any); ok {
					// Get description
					if desc := p.getDescription(actionDefMap); desc != "" {
						action.Description = formatComment("", desc, "", false, "\t")
					}

					// Parse parameters
					if paramsMap, ok := actionDefMap["parameters"].(map[string]any); ok {
						action.Parameters = p.parseActionParameters(defName, paramsMap)
					}

					// Parse action response
					if actionResp, ok := actionDefMap["actionResponse"].(map[string]any); ok {
						if ref, ok := actionResp["$ref"].(string); ok {
							// Extract response type from ref like "#/definitions/GenerateCSRResponse"
							parts := strings.Split(ref, "/")
							if len(parts) > 0 {
								action.ResponseType = cleanIdentifier(parts[len(parts)-1])
							}
						}
					}
				}
			}
		}

		actions = append(actions, action)
	}

	return actions
}

// extractJSONKeyOrder extracts the order of keys from raw JSON at a given path
// path should be dot-separated like "definitions.Reset.parameters"
func extractJSONKeyOrder(rawJSON []byte, path string) []string {
	pathParts := strings.Split(path, ".")

	// Find the position of the target object in the raw JSON
	currentJSON := string(rawJSON)

	for _, part := range pathParts {
		// Find the key in the current JSON object
		keyPattern := fmt.Sprintf(`"%s"\s*:\s*\{`, regexp.QuoteMeta(part))
		re := regexp.MustCompile(keyPattern)
		loc := re.FindStringIndex(currentJSON)
		if loc == nil {
			return nil
		}

		// Find the matching closing brace
		startIdx := loc[1] - 1 // Position of the opening brace
		braceCount := 1
		i := startIdx + 1
		inString := false

		for i < len(currentJSON) && braceCount > 0 {
			ch := currentJSON[i]
			if inString {
				if ch == '\\' {
					i += 2
					continue
				}
				if ch == '"' {
					inString = false
				}
				i++
				continue
			}

			if ch == '"' {
				inString = true
				i++
				continue
			}
			if ch == '{' {
				braceCount++
			} else if ch == '}' {
				braceCount--
			}
			i++
		}

		// Extract the object content
		currentJSON = currentJSON[startIdx+1 : i-1]
	}

	// Now extract only top-level keys from the current object
	var keys []string
	depth := 0
	i := 0

	for i < len(currentJSON) {
		ch := currentJSON[i]

		// Skip whitespace
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			i++
			continue
		}

		// At depth 0, look for keys (strings followed by colon)
		if depth == 0 && ch == '"' {
			// Extract the string content
			j := i + 1
			keyStart := j
			for j < len(currentJSON) && currentJSON[j] != '"' {
				if currentJSON[j] == '\\' {
					j++ // Skip escaped character
					if j < len(currentJSON) {
						j++ // Skip the escaped character itself
					}
				} else {
					j++
				}
			}

			if j < len(currentJSON) {
				key := currentJSON[keyStart:j]

				// Look ahead to see if this is a key (followed by colon)
				k := j + 1
				for k < len(currentJSON) && (currentJSON[k] == ' ' || currentJSON[k] == '\t' || currentJSON[k] == '\n' || currentJSON[k] == '\r') {
					k++
				}

				if k < len(currentJSON) && currentJSON[k] == ':' {
					// This is a top-level key
					keys = append(keys, key)
				}

				i = j + 1
				continue
			}
		}

		// Handle nested structures
		if ch == '"' {
			// Skip over string values
			i++
			for i < len(currentJSON) && currentJSON[i] != '"' {
				if currentJSON[i] == '\\' {
					i++ // Skip escape character
					if i < len(currentJSON) {
						i++ // Skip escaped character
					}
				} else {
					i++
				}
			}
			i++ // Skip closing quote
			continue
		} else if ch == '{' || ch == '[' {
			depth++
		} else if ch == '}' || ch == ']' {
			depth--
		}

		i++
	}

	return keys
}

// parseActionParameters parses action parameters
func (p *Parser) parseActionParameters(actionDefName string, paramsMap map[string]any) []*schema.ActionParameter {
	var params []*schema.ActionParameter

	// Extract parameter order from raw JSON
	path := fmt.Sprintf("definitions.%s.parameters", actionDefName)
	paramOrder := extractJSONKeyOrder(p.rawSchema, path)

	// Build a map of parameter name to ordinal position
	ordinalMap := make(map[string]int)
	for i, paramName := range paramOrder {
		ordinalMap[paramName] = i
	}

	paramNames := make([]string, 0, len(paramsMap))
	for paramName := range paramsMap {
		paramNames = append(paramNames, paramName)
	}

	if len(paramOrder) == 0 {
		sort.Strings(paramNames)
		for i, paramName := range paramNames {
			ordinalMap[paramName] = i
		}
	} else {
		var missing []string
		for _, paramName := range paramNames {
			if _, ok := ordinalMap[paramName]; !ok {
				missing = append(missing, paramName)
			}
		}
		sort.Strings(missing)
		for i, paramName := range missing {
			ordinalMap[paramName] = len(paramOrder) + i
		}
	}

	for paramName, paramData := range paramsMap {
		paramMap, ok := paramData.(map[string]any)
		if !ok {
			continue
		}

		param := &schema.ActionParameter{
			Name:         cleanIdentifier(strings.ToLower(paramName[:1]) + paramName[1:]),
			Type:         "string",
			Ordinal:      ordinalMap[paramName], // Set ordinal from the raw JSON key order
			OriginalName: paramName,
			FieldName:    cleanIdentifier(paramName),
		}

		// Get description
		if desc := p.getDescription(paramMap); desc != "" {
			param.Description = formatComment(param.Name+" -", desc, "", false, "\t")
		}

		// Struct-field style description (for parameter struct godoc)
		// if desc := p.getDescription(paramMap); desc != "" {
		param.FieldDescription = formatComment(param.FieldName, p.getDescription(paramMap), "shall", false, "\t")
		// } else {
		// 	param.FieldDescription = fmt.Sprintf("\t// %s", param.FieldName)
		// }

		// Check if required
		if req, ok := paramMap["required"].(bool); ok {
			param.Required = req
		}

		propJSON := mapToJSONProperty(paramMap)
		isLink := IsLinkProperty(paramName, propJSON)

		// Parse explicit type
		if _, ok := paramMap["type"]; ok {
			goType, isPointer, isArray := p.typeMapper.MapType(paramName, propJSON)

			if isLink {
				param.Type = "string"
			} else if isPointer {
				param.Type = "*" + goType
			} else {
				param.Type = goType
			}

			if isArray {
				param.Type = "[]" + param.Type
			}
		}

		// Parse type from $ref
		if ref, ok := paramMap["$ref"].(string); ok {
			if isLink {
				param.Type = "string"
			} else {
				// Extract type name from ref like "http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/ResetType"
				parts := strings.Split(ref, "/")
				if len(parts) > 0 {
					param.Type = cleanIdentifier(parts[len(parts)-1])
				}
			}
		}

		params = append(params, param)
	}

	// Sort parameters by their ordinal position to maintain schema order
	sort.Slice(params, func(i, j int) bool {
		return params[i].Ordinal < params[j].Ordinal
	})

	return params
}

// parseLinks parses the Links definition for a resource
func (p *Parser) parseLinks(linksMap map[string]any) []*schema.Link {
	var links []*schema.Link

	propsMap, ok := linksMap["properties"].(map[string]any)
	if !ok {
		return links
	}

	for linkName, linkData := range propsMap {
		// Skip Oem links
		if strings.Contains(linkName, "Oem") {
			continue
		}

		// Skip @odata.count properties
		if strings.Contains(linkName, "@odata.count") {
			continue
		}

		linkMap, ok := linkData.(map[string]any)
		if !ok {
			continue
		}

		link := &schema.Link{
			Name:     linkName,
			JSONName: linkName,
		}

		// Get description
		link.Description = p.getDescription(linkMap)

		// Check if deprecated
		if _, hasDeprecated := linkMap["deprecated"]; hasDeprecated {
			link.Deprecated = true
		}

		// Determine if it's an array by checking if it has "items"
		if _, hasItems := linkMap["items"]; hasItems {
			link.IsArray = true
		}

		// Try to determine the target type from $ref, anyOf, or items
		if ref, ok := linkMap["$ref"].(string); ok {
			link.Type = extractTypeFromRef(ref)
		} else if anyOf, ok := linkMap["anyOf"].([]any); ok {
			// Handle nullable links with anyOf: [{$ref: ...}, {type: null}]
			for _, item := range anyOf {
				if itemMap, ok := item.(map[string]any); ok {
					if ref, ok := itemMap["$ref"].(string); ok {
						link.Type = extractTypeFromRef(ref)
						break
					}
				}
			}
		} else if items, ok := linkMap["items"].(map[string]any); ok {
			if ref, ok := items["$ref"].(string); ok {
				link.Type = extractTypeFromRef(ref)
			}
		}

		// If we couldn't get type from $ref, try to infer from property name
		if link.Type == "" {
			// Remove trailing 's' if plural
			typeName := linkName
			if strings.HasSuffix(typeName, "s") && len(typeName) > 1 {
				typeName = typeName[:len(typeName)-1]
			}
			link.Type = cleanIdentifier(typeName)
		}

		links = append(links, link)
	}

	return links
}

// extractTypeFromRef extracts the resource type name from a $ref URL
func extractTypeFromRef(ref string) string {
	typeName := schema.ExtractTypeFromRef(ref, true)
	// Filter out generic Link types
	if typeName == "Link" || typeName == "Links" {
		return ""
	}
	return cleanIdentifier(typeName)
}
