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
	analyzer   *analyzer.DependencyAnalyzer
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

	return &BatchProcessor{
		schemaDirs: schemaDirs,
		outputDir:  outputDir,
		verbose:    verbose,
	}, nil
}

// ProcessAll processes all schemas in the directories using two-pass approach
func (bp *BatchProcessor) ProcessAll(parallel bool) error {
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

	// PASS 2: Generate code
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

			semaphore <- struct{}{}        // Acquire
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

// isSwordfishBundle returns true if the file path indicates it's from the Swordfish bundle.
// This is based on directory path, not content analysis.
func isSwordfishBundle(filePath string) bool {
	return strings.Contains(strings.ToLower(filePath), "swordfish")
}

// processOne processes a single schema
func (bp *BatchProcessor) processOne(objectName, baseFile string) error {
	if bp.verbose {
		log.Printf("Processing %s...", objectName)
	}

	// Determine the content origin from this specific file
	contentOrigin := parser.GetSchemaOrigin(baseFile, objectName)

	// For skip decisions, use bundle path (Swordfish directory means skip duplicates)
	// This handles cases where SNIA schemas are included in the Redfish bundle
	bundleIsSwordfish := isSwordfishBundle(baseFile)
	if bundleIsSwordfish {
		action, hasConflict := schema.SchemaConflicts[objectName]
		if hasConflict && action == schema.ConflictSkipSwordfish {
			if bp.verbose {
				log.Printf("  Skipping %s (Swordfish bundle duplicate)", objectName)
			}
			return nil
		}
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

	// Each call gets its own Generator to avoid race conditions during parallel processing
	gen, err := codegen.NewGenerator()
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}

	// Special case: ServiceRoot goes to root with "gofish" package, renamed to "Service"
	if objectName == "ServiceRoot" {
		for _, def := range definitions {
			def.Package = "gofish"
		}
		typeRenames := map[string]string{"ServiceRoot": "Service"}
		code, err := gen.Generate(objectName, "gofish", definitions, false, typeRenames)
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

	// Determine if this schema needs SF prefix based on actual file content origin
	sfPrefix := schema.NeedsSFPrefix(objectName, contentOrigin)

	// All other schemas go to the schemas package
	pkgType := schema.PackageSchemas

	// Set package type on all definitions
	for _, def := range definitions {
		def.Package = pkgType
	}

	// Generate code
	code, err := gen.Generate(objectName, pkgType, definitions, sfPrefix, nil)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write output to schemas/ directory
	// SF-prefixed schemas use filename sf<name>.go
	outputDir := filepath.Join(bp.outputDir, "schemas")
	filename := strings.ToLower(objectName) + ".go"
	if sfPrefix {
		filename = "sf" + filename
	}
	outputFile := filepath.Join(outputDir, filename)

	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if bp.verbose {
		log.Printf("Generated schemas/%s", filepath.Base(outputFile))
	}

	return nil
}
