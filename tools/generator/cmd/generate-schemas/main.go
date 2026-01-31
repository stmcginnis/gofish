//
// SPDX-License-Identifier: BSD-3-Clause
//

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/stmcginnis/gofish/tools/generator/internal/codegen"
	"github.com/stmcginnis/gofish/tools/generator/internal/fetcher"
	"github.com/stmcginnis/gofish/tools/generator/internal/parser"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func main() {
	var (
		objectName string
		localPath  string
		outputDir  string
		verbose    bool
	)

	flag.StringVar(&objectName, "object", "", "Generate specific schema object only (optional)")
	flag.StringVar(&localPath, "local", "", "Use local schema files instead of downloading (for testing)")
	flag.StringVar(&outputDir, "output-dir", ".", "Base output directory (creates schemas subdir)")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.Parse()

	// Set up logging
	if !verbose {
		log.SetOutput(os.Stderr)
	}

	// Determine schema directories
	var schemaDirs []string

	if localPath != "" {
		// Use local schemas - support comma-separated paths for multiple directories
		schemaDirs = strings.Split(localPath, ",")
		if verbose {
			log.Printf("Using local schemas from: %s", localPath)
		}
	} else {
		// Download schemas
		if verbose {
			log.Printf("Downloading schemas from DMTF repositories...")
		}

		f, err := fetcher.NewFetcher()
		if err != nil {
			log.Fatalf("Failed to create fetcher: %v", err)
		}
		defer f.Cleanup()

		// Fetch Redfish schemas
		redfishDir, err := f.FetchRedfish()
		if err != nil {
			log.Fatalf("Failed to fetch Redfish schemas: %v", err)
		}
		schemaDirs = append(schemaDirs, redfishDir)

		if verbose {
			log.Printf("Downloaded Redfish schemas to: %s", redfishDir)
		}

		// Fetch Swordfish schemas
		swordfishDir, err := f.FetchSwordfish()
		if err != nil {
			log.Printf("Warning: Failed to fetch Swordfish schemas: %v", err)
		} else {
			schemaDirs = append(schemaDirs, swordfishDir)
			if verbose {
				log.Printf("Downloaded Swordfish schemas to: %s", swordfishDir)
			}
		}
	}

	// Single object mode
	if objectName != "" {
		// For single object mode, try to find the schema in any of the directories
		var schemaDir string
		for _, dir := range schemaDirs {
			baseFile := filepath.Join(dir, objectName+".json")
			if _, err := os.Stat(baseFile); err == nil {
				schemaDir = dir
				break
			}
		}
		if schemaDir == "" {
			log.Fatalf("Schema file for %s not found in any schema directory", objectName)
		}

		if err := generateSingleObject(objectName, schemaDir, outputDir, verbose); err != nil {
			log.Fatalf("Failed to generate %s: %v", objectName, err)
		}
		return
	}

	if verbose {
		log.Printf("Generating all schemas...")
	}

	processor, err := NewBatchProcessor(schemaDirs, outputDir, verbose)
	if err != nil {
		log.Fatalf("Failed to create batch processor: %v", err)
	}

	if err := processor.ProcessAll(true); err != nil {
		log.Fatalf("Batch processing failed: %v", err)
	}

	if verbose {
		log.Printf("Generation complete!")
	}
}

// generateSingleObject generates a single schema object
func generateSingleObject(objectName, schemaDir, outputBaseDir string, verbose bool) error {
	if verbose {
		log.Printf("Generating %s...", objectName)
	}

	// Find base schema file
	baseFile := filepath.Join(schemaDir, objectName+".json")
	if _, err := os.Stat(baseFile); err != nil {
		return fmt.Errorf("schema file not found: %s", baseFile)
	}

	// Determine package type
	pkgType := parser.CategorizeSchema(baseFile, objectName)

	// ServiceRoot uses "gofish" package
	if objectName == "ServiceRoot" {
		pkgType = "gofish"
	}

	if verbose {
		log.Printf("Categorized as: %s", pkgType)
	}

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

	// Set package type
	for _, def := range definitions {
		def.Package = pkgType
	}

	// Generate code
	gen, err := codegen.NewGenerator()
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}

	// Determine SF prefix based on schema origin
	sfPrefix := false
	if pkgType != "gofish" {
		origin := parser.GetSchemaOrigin(baseFile, objectName)
		sfPrefix = schema.NeedsSFPrefix(objectName, origin)
	}

	// Handle type renames (ServiceRoot -> Service)
	var typeRenames map[string]string
	if objectName == "ServiceRoot" {
		typeRenames = map[string]string{"ServiceRoot": "Service"}
	}

	code, err := gen.Generate(objectName, pkgType, definitions, sfPrefix, typeRenames)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write output - ServiceRoot goes to root, others to schemas/
	var outputFile string
	var logPrefix string
	switch {
	case objectName == "ServiceRoot":
		if err := os.MkdirAll(outputBaseDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		outputFile = filepath.Join(outputBaseDir, strings.ToLower(objectName)+".go")
		logPrefix = ""
	default:
		outputDir := filepath.Join(outputBaseDir, "schemas")
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		filename := strings.ToLower(objectName) + ".go"
		if sfPrefix {
			filename = "sf" + filename
		}
		outputFile = filepath.Join(outputDir, filename)
		logPrefix = "schemas/"
	}

	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if verbose {
		log.Printf("Generated: %s%s", logPrefix, filepath.Base(outputFile))
	} else {
		fmt.Printf("Generated: %s\n", outputFile)
	}

	return nil
}
