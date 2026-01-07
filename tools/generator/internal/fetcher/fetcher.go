//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	RedfishGitRepo      = "https://github.com/DMTF/Redfish-Publications.git"
	SwordfishSchemaBase = "https://redfish.dmtf.org/schemas/swordfish/v1/"
)

// Fetcher handles downloading schemas
type Fetcher struct {
	tempDir string
}

// NewFetcher creates a new Fetcher
func NewFetcher() (*Fetcher, error) {
	tempDir, err := os.MkdirTemp("", "gofish-schemas-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	return &Fetcher{
		tempDir: tempDir,
	}, nil
}

// Cleanup removes temporary directories
func (f *Fetcher) Cleanup() error {
	if f.tempDir != "" {
		return os.RemoveAll(f.tempDir)
	}
	return nil
}

// GetTempDir returns the temporary directory path
func (f *Fetcher) GetTempDir() string {
	return f.tempDir
}

// FetchRedfish fetches Redfish schemas by shallow cloning the repository
func (f *Fetcher) FetchRedfish() (string, error) {
	redfishDir := filepath.Join(f.tempDir, "redfish")

	// Shallow clone the repository
	cmd := exec.Command("git", "clone", "--depth", "1", RedfishGitRepo, redfishDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to clone Redfish repository: %w\n%s", err, string(output))
	}

	// Find the json-schema directory
	// The structure is usually DMTF/Redfish-Publications/json-schema/
	schemaDir := filepath.Join(redfishDir, "json-schema")
	if _, err := os.Stat(schemaDir); err != nil {
		// Try to find it in subdirectories
		matches, _ := filepath.Glob(filepath.Join(redfishDir, "*/json-schema"))
		if len(matches) > 0 {
			schemaDir = matches[0]
		} else {
			return "", fmt.Errorf("json-schema directory not found in cloned repository")
		}
	}

	return schemaDir, nil
}

// FetchSwordfish fetches Swordfish schemas by downloading files
func (f *Fetcher) FetchSwordfish() (string, error) {
	swordfishDir := filepath.Join(f.tempDir, "swordfish")
	if err := os.MkdirAll(swordfishDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create swordfish directory: %w", err)
	}

	// Get the index page to find all JSON files
	resp, err := http.Get(SwordfishSchemaBase)
	if err != nil {
		return "", fmt.Errorf("failed to fetch swordfish index: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Parse HTML to find .json files (simple approach)
	content := string(body)
	files := extractJSONFiles(content)

	// Download each file
	for _, file := range files {
		if err := f.downloadFile(SwordfishSchemaBase+file, filepath.Join(swordfishDir, file)); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to download %s: %v\n", file, err)
		}
	}

	return swordfishDir, nil
}

// downloadFile downloads a file from a URL to a local path
func (f *Fetcher) downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// extractJSONFiles extracts JSON file names from HTML content
func extractJSONFiles(html string) []string {
	var files []string

	// Simple parsing - look for .json files in href attributes
	lines := strings.SplitSeq(html, "\n")
	for line := range lines {
		if strings.Contains(line, ".json") {
			// Extract href="filename.json"
			start := strings.Index(line, "href=\"")
			if start == -1 {
				continue
			}
			start += 6
			end := strings.Index(line[start:], "\"")
			if end == -1 {
				continue
			}
			filename := line[start : start+end]
			if strings.HasSuffix(filename, ".json") && !strings.Contains(filename, "/") {
				files = append(files, filename)
			}
		}
	}

	return files
}

