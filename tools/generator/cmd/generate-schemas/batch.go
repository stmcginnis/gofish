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

	"github.com/stmcginnis/gofish/tools/generator/internal/codegen"
	"github.com/stmcginnis/gofish/tools/generator/internal/parser"
)

// BatchProcessor handles batch generation of multiple schemas
type BatchProcessor struct {
	schemaDirs []string
	outputDir  string
	verbose    bool
	gen        *codegen.Generator
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

// ProcessAll processes all schemas in the directories
func (bp *BatchProcessor) ProcessAll(parallel bool) error {
	// Collect all base schema files from all directories
	baseFiles := []string{}

	for _, schemaDir := range bp.schemaDirs {
		// Find all base schema files (*.json, not *.v*.json)
		files, err := filepath.Glob(filepath.Join(schemaDir, "*.json"))
		if err != nil {
			return fmt.Errorf("failed to list schema files in %s: %w", schemaDir, err)
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

			// Skip odata and protocol files
			if strings.HasPrefix(basename, "odata") || basename == "redfish-schema" || basename == "Protocol.json" {
				continue
			}

			baseFiles = append(baseFiles, file)
		}
	}

	if bp.verbose {
		log.Printf("Found %d schemas to process across %d directories", len(baseFiles), len(bp.schemaDirs))
	}

	if parallel {
		return bp.processParallel(baseFiles)
	}
	return bp.processSequential(baseFiles)
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

	// Determine package type for this schema
	pkgType := parser.CategorizeSchema(baseFile, objectName)

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

	// Set package type - ServiceRoot uses "gofish" package
	if objectName == "ServiceRoot" {
		for _, def := range definitions {
			def.Package = "gofish"
		}
		pkgType = "gofish"
	} else {
		for _, def := range definitions {
			def.Package = pkgType
		}
	}

	// Generate code
	code, err := bp.gen.Generate(objectName, pkgType, definitions)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write output - ServiceRoot goes to root directory
	var outputFile string
	if objectName == "ServiceRoot" {
		outputFile = filepath.Join(bp.outputDir, strings.ToLower(objectName)+".go")
	} else {
		outputDir := filepath.Join(bp.outputDir, string(pkgType))
		outputFile = filepath.Join(outputDir, strings.ToLower(objectName)+".go")
	}

	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if bp.verbose {
		if objectName == "ServiceRoot" {
			log.Printf("Generated %s", filepath.Base(outputFile))
		} else {
			log.Printf("Generated %s/%s", string(pkgType), filepath.Base(outputFile))
		}
	}

	return nil
}
