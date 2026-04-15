//
// SPDX-License-Identifier: BSD-3-Clause
//

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/stmcginnis/gofish/tools/generator/internal/analyzer"
	"github.com/stmcginnis/gofish/tools/generator/internal/codegen"
	"github.com/stmcginnis/gofish/tools/generator/internal/parser"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

// BatchProcessor handles batch generation of multiple schemas
type BatchProcessor struct {
	schemaDirs []string
	outputDir  string
	verbose    bool
	gen        *codegen.Generator
	analyzer   *analyzer.DependencyAnalyzer
}

// NewBatchProcessor creates a new BatchProcessor
func NewBatchProcessor(schemaDirs []string, outputDir string, verbose bool) (*BatchProcessor, error) {
	gen, err := codegen.NewGenerator()
	if err != nil {
		return nil, fmt.Errorf("failed to create generator: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create package subdirectories
	for _, pkg := range []string{"common", "redfish", "swordfish"} {
		pkgDir := filepath.Join(outputDir, pkg)
		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create package directory %s: %w", pkg, err)
		}
	}

	return &BatchProcessor{
		schemaDirs: schemaDirs,
		outputDir:  outputDir,
		verbose:    verbose,
		gen:        gen,
	}, nil
}

// ProcessAll processes all schemas in the directories using two-pass approach
func (bp *BatchProcessor) ProcessAll(parallel bool) error {
	// Collect all base schema files from all directories
	baseFiles := bp.collectBaseFiles()

	if bp.verbose {
		log.Printf("Found %d schemas to process across %d directories", len(baseFiles), len(bp.schemaDirs))
	}

	// PASS 1: Analyze all schemas to determine dependencies
	if bp.verbose {
		log.Printf("Pass 1: Analyzing schema dependencies...")
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

	// Compute which Swordfish types are used by Redfish (transitive closure)
	bp.analyzer.ComputeDependencies()

	if bp.verbose {
		// Log which Swordfish definitions will go to common
		for defName := range bp.analyzer.SwordfishDefsUsedByRedfish {
			log.Printf("  Swordfish definition %s will be placed in common (used by Redfish)", defName)
		}
	}

	// PASS 2: Generate code with proper package assignments
	if bp.verbose {
		log.Printf("Pass 2: Generating code...")
	}

	if parallel {
		return bp.processParallel(baseFiles)
	}
	return bp.processSequential(baseFiles)
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

			// Skip odata and redfish-schema meta files
			if strings.HasPrefix(basename, "odata") || basename == "redfish-schema" {
				continue
			}

			baseFiles = append(baseFiles, file)
		}
	}

	return baseFiles
}

// processSequential processes schemas one at a time
func (bp *BatchProcessor) processSequential(files []string) error {
	success := 0
	failed := 0

	for _, file := range files {
		objectName := strings.TrimSuffix(filepath.Base(file), ".json")

		if err := bp.processOne(objectName, file); err != nil {
			if bp.verbose {
				log.Printf("ERROR processing %s: %v", objectName, err)
			}
			failed++
		} else {
			success++
		}
	}

	if bp.verbose {
		log.Printf("Completed: %d successful, %d failed", success, failed)
	}
	return nil
}

// processParallel processes schemas in parallel
func (bp *BatchProcessor) processParallel(files []string) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	success := 0
	failed := 0
	maxWorkers := 8

	semaphore := make(chan struct{}, maxWorkers)

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()

			semaphore <- struct{}{} // Acquire
			defer func() { <-semaphore }() // Release

			objectName := strings.TrimSuffix(filepath.Base(f), ".json")

			if err := bp.processOne(objectName, f); err != nil {
				if bp.verbose {
					log.Printf("ERROR processing %s: %v", objectName, err)
				}
				mu.Lock()
				failed++
				mu.Unlock()
			} else {
				mu.Lock()
				success++
				mu.Unlock()
			}
		}(file)
	}

	wg.Wait()

	if bp.verbose {
		log.Printf("Completed: %d successful, %d failed", success, failed)
	}
	return nil
}

// processOne processes a single schema
func (bp *BatchProcessor) processOne(objectName, baseFile string) error {
	if bp.verbose {
		log.Printf("Processing %s...", objectName)
	}

	// The baseFile's directory is the schema directory for this file
	schemaDir := filepath.Dir(baseFile)

	// Resolve latest version
	versionFile, err := parser.ResolveLatestVersion(baseFile, schemaDir)
	if err != nil {
		return fmt.Errorf("failed to resolve version: %w", err)
	}

	// Parse schema (both base and versioned to merge definitions)
	p := parser.NewParser(schemaDir)
	definitions, err := p.ParseSchemaWithBase(baseFile, versionFile)
	if err != nil {
		return fmt.Errorf("failed to parse schema: %w", err)
	}

	// Special case: ServiceRoot goes to root with "gofish" package
	if objectName == "ServiceRoot" {
		for _, def := range definitions {
			def.Package = "gofish"
		}
		code, err := bp.gen.Generate(objectName, "gofish", definitions, false)
		if err != nil {
			return fmt.Errorf("failed to generate code: %w", err)
		}
		outputFile := filepath.Join(bp.outputDir, strings.ToLower(objectName)+".go")
		if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
		if bp.verbose {
			log.Printf("Generated %s", filepath.Base(outputFile))
		}
		return nil
	}

	// Check if this schema needs split generation (definitions going to different packages)
	if bp.analyzer != nil && bp.analyzer.NeedsSplitGeneration(objectName) {
		return bp.processSplitSchema(objectName, definitions)
	}

	// Single package generation (normal case)
	var pkgType schema.PackageType
	var isSwordfishInCommon bool

	if bp.analyzer != nil {
		pkgType = bp.analyzer.GetPackageType(objectName)
		isSwordfishInCommon = bp.analyzer.IsSwordfishTypeInCommon(objectName)
	} else {
		pkgType = parser.CategorizeSchema(baseFile, objectName)
	}

	// Set package type on all definitions
	for _, def := range definitions {
		def.Package = pkgType
	}

	// Generate code
	code, err := bp.gen.Generate(objectName, pkgType, definitions, isSwordfishInCommon)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write output
	outputDir := filepath.Join(bp.outputDir, string(pkgType))
	outputFile := filepath.Join(outputDir, strings.ToLower(objectName)+".go")

	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if bp.verbose {
		log.Printf("Generated %s/%s", string(pkgType), filepath.Base(outputFile))
	}

	return nil
}

// processSplitSchema handles schemas where definitions go to different packages
func (bp *BatchProcessor) processSplitSchema(objectName string, allDefinitions []*schema.Definition) error {
	if bp.verbose {
		log.Printf("  Schema %s requires split generation", objectName)
	}

	// Get definitions grouped by target package
	defsByPackage := bp.analyzer.GetDefinitionsByPackage(objectName)

	// Create a map of definition names for quick lookup
	defMap := schema.BuildDefinitionMap(allDefinitions, true)

	// Process each package that has definitions
	for pkgType, defNames := range defsByPackage {
		// Filter definitions for this package
		var pkgDefinitions []*schema.Definition
		for _, defName := range defNames {
			if def, ok := defMap[defName]; ok {
				// Clone the definition to avoid modifying the original
				defCopy := *def
				defCopy.Package = pkgType
				pkgDefinitions = append(pkgDefinitions, &defCopy)
			}
		}

		if len(pkgDefinitions) == 0 {
			continue
		}

		// Determine the main type for this package subset
		mainTypeName := bp.determineMainType(objectName, pkgDefinitions)

		// Determine if this is a Swordfish type in common
		isSwordfishInCommon := pkgType == schema.PackageCommon &&
			bp.analyzer.GetSchemaOrigin(objectName) == schema.OriginSwordfish

		// Generate code for this subset
		code, err := bp.gen.Generate(mainTypeName, pkgType, pkgDefinitions, isSwordfishInCommon)
		if err != nil {
			return fmt.Errorf("failed to generate code for %s in %s: %w", mainTypeName, pkgType, err)
		}

		// Determine output filename
		outputDir := filepath.Join(bp.outputDir, string(pkgType))
		outputFile := filepath.Join(outputDir, strings.ToLower(mainTypeName)+".go")

		if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		if bp.verbose {
			log.Printf("  Generated %s/%s (%d definitions)", string(pkgType), filepath.Base(outputFile), len(pkgDefinitions))
		}
	}

	return nil
}

// determineMainType finds the main type name for a subset of definitions
// If the subset contains the schema's main type (entity), use the schema name
// Otherwise, use the most prominent type name (prefer entity types, then largest struct)
func (bp *BatchProcessor) determineMainType(schemaName string, definitions []*schema.Definition) string {
	// First check if any definition matches the schema name (is the main entity)
	for _, def := range definitions {
		if def.OriginalName == schemaName || def.Name == schemaName {
			return schemaName
		}
	}

	// Find the best candidate among the definitions
	var bestCandidate *schema.Definition
	for _, def := range definitions {
		if def.IsEnum {
			continue // Skip enums as main type
		}
		if def.IsEntity {
			// Prefer entity types
			return def.OriginalName
		}
		if bestCandidate == nil || len(def.Properties) > len(bestCandidate.Properties) {
			bestCandidate = def
		}
	}

	if bestCandidate != nil {
		return bestCandidate.OriginalName
	}

	// Fallback: use first non-enum definition
	for _, def := range definitions {
		if !def.IsEnum {
			return def.OriginalName
		}
	}

	// Last resort: use first definition
	if len(definitions) > 0 {
		return definitions[0].OriginalName
	}

	return schemaName
}
