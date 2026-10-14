//
// SPDX-License-Identifier: BSD-3-Clause
//

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/analyzer"
	"github.com/stmcginnis/gofish/tools/generator/internal/codegen"
	"github.com/stmcginnis/gofish/tools/generator/internal/config"
	"github.com/stmcginnis/gofish/tools/generator/internal/parser"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// BatchProcessor handles batch generation of multiple schemas
type BatchProcessor struct {
	schemaDirs []string
	outputDir  string
	verbose    bool
	analyzer   *analyzer.DependencyAnalyzer

	// Declarations provided by hand-written files in the output packages.
	schemasHandwritten map[string]bool
	rootHandwritten    map[string]bool
}

// NewBatchProcessor creates a new BatchProcessor
func NewBatchProcessor(schemaDirs []string, outputDir string, verbose bool) (*BatchProcessor, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create schemas subdirectory
	schemasDir := filepath.Join(outputDir, "schemas")
	if err := os.MkdirAll(schemasDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create schemas directory: %w", err)
	}

	schemasHandwritten, err := codegen.ScanHandwritten(schemasDir)
	if err != nil {
		return nil, err
	}
	rootHandwritten, err := codegen.ScanHandwritten(outputDir)
	if err != nil {
		return nil, err
	}

	return &BatchProcessor{
		schemaDirs:         schemaDirs,
		outputDir:          outputDir,
		verbose:            verbose,
		schemasHandwritten: schemasHandwritten,
		rootHandwritten:    rootHandwritten,
	}, nil
}

// ProcessAll processes all schemas in the directories using two-pass approach
func (bp *BatchProcessor) ProcessAll() error {
	// Collect all base schema files from all directories
	baseFiles := bp.collectBaseFiles()

	if bp.verbose {
		log.Printf("Found %d schemas to process across %d directories", len(baseFiles), len(bp.schemaDirs))
	}

	// PASS 1: Analyze all schemas to detect conflicts
	if bp.verbose {
		log.Printf("Pass 1: Analyzing schemas for conflicts...")
	}

	bp.analyzer = analyzer.NewDependencyAnalyzer(bp.verbose)
	for _, file := range baseFiles {
		objectName := strings.TrimSuffix(filepath.Base(file), ".json")
		if err := bp.analyzer.AnalyzeSchema(file, objectName); err != nil {
			if bp.verbose {
				log.Printf("Warning: failed to analyze %s: %v", objectName, err)
			}
		}
	}

	// PASS 2: Parse every schema so definitions shared between schemas can be
	// assigned to a single owning file.
	var units []*schemaUnit
	failed := 0
	for _, file := range baseFiles {
		objectName := strings.TrimSuffix(filepath.Base(file), ".json")
		unit, err := bp.parseOne(objectName, file)
		if err != nil {
			log.Printf("ERROR parsing %s: %v", objectName, err)
			failed++
			continue
		}
		if unit != nil {
			units = append(units, unit)
		}
	}
	assignOwners(units)

	// PASS 3: Generate code
	if bp.verbose {
		log.Printf("Generating code...")
	}
	for _, unit := range units {
		if err := bp.generateOne(unit); err != nil {
			log.Printf("ERROR generating %s: %v", unit.objectName, err)
			failed++
		}
	}

	if failed > 0 {
		return fmt.Errorf("%d schemas failed", failed)
	}
	return nil
}

// collectBaseFiles collects all base schema files from schema directories
func (bp *BatchProcessor) collectBaseFiles() []string {
	baseFiles := []string{}

	for _, schemaDir := range bp.schemaDirs {
		// Find all base schema files (*.json, not *.v*.json)
		files, err := filepath.Glob(filepath.Join(schemaDir, "*.json"))
		if err != nil {
			continue
		}

		// Filter out versioned files and unwanted patterns
		for _, file := range files {
			basename := filepath.Base(file)

			// Skip versioned files
			if strings.Contains(basename, ".v") {
				continue
			}

			// Skip Collection files (they're in common)
			if strings.Contains(basename, "Collection") {
				continue
			}

			// Skip odata, redfish-schema and payload-annotation meta files
			if strings.HasPrefix(basename, "odata") || strings.HasPrefix(basename, "redfish-schema") ||
				strings.HasPrefix(basename, "redfish-payload-annotations") {
				continue
			}

			baseFiles = append(baseFiles, file)
		}
	}

	return baseFiles
}

// schemaUnit is one parsed schema waiting to be written to a Go file.
type schemaUnit struct {
	objectName  string
	definitions []*schema.Definition
	sfPrefix    bool
	serviceRoot bool
}

// goName returns the Go type name a definition is emitted as.
func (u *schemaUnit) goName(def *schema.Definition) string {
	if u.sfPrefix {
		return "SF" + def.Name
	}
	return def.Name
}

// isSwordfishBundle returns true if the file path indicates it's from the Swordfish bundle.
// This is based on directory path, not content analysis.
func isSwordfishBundle(filePath string) bool {
	return strings.Contains(strings.ToLower(filePath), "swordfish")
}

// parseOne parses a single schema. It returns nil for schemas that are skipped.
func (bp *BatchProcessor) parseOne(objectName, baseFile string) (*schemaUnit, error) {
	if bp.verbose {
		log.Printf("Parsing %s...", objectName)
	}

	// Determine the content origin from this specific file
	contentOrigin := parser.GetSchemaOrigin(baseFile, objectName)

	// For skip decisions, use bundle path (Swordfish directory means skip duplicates)
	// This handles cases where SNIA schemas are included in the Redfish bundle
	if isSwordfishBundle(baseFile) {
		action, hasConflict := schema.SchemaConflicts[objectName]
		if hasConflict && action == schema.ConflictSkipSwordfish {
			if bp.verbose {
				log.Printf("  Skipping %s (Swordfish bundle duplicate)", objectName)
			}
			return nil, nil
		}
	}

	// The baseFile's directory is the schema directory for this file
	schemaDir := filepath.Dir(baseFile)

	versionFile, err := parser.ResolveLatestVersion(baseFile, schemaDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve version: %w", err)
	}

	// Parse schema (both base and versioned to merge definitions)
	p := parser.NewParser(schemaDir)
	definitions, err := p.ParseSchemaWithBase(baseFile, versionFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	kept := definitions[:0]
	for _, def := range definitions {
		if !config.SkipDefinitions[objectName+"."+def.OriginalName] {
			kept = append(kept, def)
		}
	}

	return &schemaUnit{
		objectName:  objectName,
		definitions: kept,
		sfPrefix:    schema.NeedsSFPrefix(objectName, contentOrigin),
		serviceRoot: objectName == "ServiceRoot",
	}, nil
}

// assignOwners removes definitions that more than one schema emits under the
// same Go name, keeping the copy in the owning schema. The owner is, in order:
// the TypeOwners entry, the schema whose main type it is, Resource, or the
// first schema alphabetically. Copies that differ from the owner's and have no
// TypeOwners entry are reported, because each reference now resolves to the
// owner's version.
func assignOwners(units []*schemaUnit) {
	defining := map[string][]*schemaUnit{}
	defsByUnit := map[*schemaUnit]map[string]*schema.Definition{}
	for _, u := range units {
		// ServiceRoot is generated into the gofish package, so it cannot share
		// types with the schemas package.
		if u.serviceRoot {
			continue
		}
		defsByUnit[u] = map[string]*schema.Definition{}
		for _, def := range u.definitions {
			name := u.goName(def)
			if _, dup := defsByUnit[u][name]; dup {
				continue
			}
			defsByUnit[u][name] = def
			defining[name] = append(defining[name], u)
		}
	}

	names := make([]string, 0, len(defining))
	for name := range defining {
		names = append(names, name)
	}
	sort.Strings(names)

	owners := map[string]*schemaUnit{}
	for _, name := range names {
		us := defining[name]
		if len(us) < 2 {
			continue
		}
		sort.Slice(us, func(i, j int) bool { return us[i].objectName < us[j].objectName })
		owner := pickOwner(name, us, defsByUnit)
		owners[name] = owner

		if _, configured := config.TypeOwners[name]; !configured {
			want := signature(defsByUnit[owner][name])
			for _, u := range us {
				if u != owner && signature(defsByUnit[u][name]) != want {
					log.Printf("WARNING: %s in %s differs from the copy in %s, which is used for both; "+
						"add a TypeRenames or TypeOwners entry", name, u.objectName, owner.objectName)
				}
			}
		}
		unionEnumValues(defsByUnit[owner][name], us, name, defsByUnit)
	}

	for _, u := range units {
		kept := u.definitions[:0]
		for _, def := range u.definitions {
			if owner, shared := owners[u.goName(def)]; shared && owner != u && !u.serviceRoot {
				continue
			}
			kept = append(kept, def)
		}
		u.definitions = kept
	}
}

// pickOwner chooses the schema that emits shared type name. us is sorted by
// schema name.
func pickOwner(name string, us []*schemaUnit, defsByUnit map[*schemaUnit]map[string]*schema.Definition) *schemaUnit {
	configured, hasConfig := config.TypeOwners[name]
	var resource *schemaUnit
	for _, u := range us {
		if hasConfig && u.objectName == configured {
			return u
		}
		if defsByUnit[u][name].OriginalName == u.objectName {
			return u
		}
		if u.objectName == "Resource" {
			resource = u
		}
	}
	if hasConfig {
		log.Printf("WARNING: TypeOwners[%s] = %s, which does not define it", name, configured)
	}
	if resource != nil {
		return resource
	}
	return us[0]
}

// unionEnumValues adds values that other copies of a shared enum define to
// the owner's copy, so every schema's values stay available.
func unionEnumValues(owner *schema.Definition, us []*schemaUnit, name string, defsByUnit map[*schemaUnit]map[string]*schema.Definition) {
	if !owner.IsEnum {
		return
	}
	seen := map[string]bool{}
	for _, v := range owner.EnumValues {
		seen[v.Value] = true
	}
	for _, u := range us {
		def := defsByUnit[u][name]
		if !def.IsEnum {
			continue
		}
		for _, v := range def.EnumValues {
			if !seen[v.Value] {
				seen[v.Value] = true
				owner.EnumValues = append(owner.EnumValues, v)
			}
		}
	}
}

// signature summarizes a definition's shape so duplicate copies can be compared.
func signature(def *schema.Definition) string {
	var parts []string
	for _, v := range def.EnumValues {
		parts = append(parts, v.Value)
	}
	for _, p := range def.Properties {
		parts = append(parts, p.JSONName+":"+p.Type)
	}
	sort.Strings(parts)
	return strings.Join(parts, ",")
}

// generateOne writes the Go file for a parsed schema.
func (bp *BatchProcessor) generateOne(u *schemaUnit) error {
	if bp.verbose {
		log.Printf("Generating %s...", u.objectName)
	}

	// Each call gets its own Generator to avoid shared template state
	gen, err := codegen.NewGenerator()
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}
	if u.serviceRoot {
		gen.SetHandwritten(bp.rootHandwritten)
	} else {
		gen.SetHandwritten(bp.schemasHandwritten)
	}

	// Special case: ServiceRoot goes to root with "gofish" package, renamed to "Service"
	if u.serviceRoot {
		for _, def := range u.definitions {
			def.Package = "gofish"
		}
		typeRenames := map[string]string{"ServiceRoot": "Service"}
		code, err := gen.Generate(u.objectName, "gofish", u.definitions, false, typeRenames)
		if err != nil {
			return fmt.Errorf("failed to generate code: %w", err)
		}
		outputFile := filepath.Join(bp.outputDir, strings.ToLower(u.objectName)+".go")
		return os.WriteFile(outputFile, []byte(code), 0644)
	}

	pkgType := schema.PackageSchemas
	for _, def := range u.definitions {
		def.Package = pkgType
	}

	code, err := gen.Generate(u.objectName, pkgType, u.definitions, u.sfPrefix, nil)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// SF-prefixed schemas use filename sf<name>.go
	filename := strings.ToLower(u.objectName) + ".go"
	if u.sfPrefix {
		filename = "sf" + filename
	}
	return os.WriteFile(filepath.Join(bp.outputDir, "schemas", filename), []byte(code), 0644)
}
