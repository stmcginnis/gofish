//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	RedfishGitRepo   = "https://github.com/DMTF/Redfish-Publications.git"
	SwordfishGitRepo = "https://github.com/SNIA/Swordfish-Publications.git"
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
	return f.cloneAndFindSchemas("redfish", RedfishGitRepo)
}

// FetchSwordfish fetches Swordfish schemas by shallow cloning the repository
func (f *Fetcher) FetchSwordfish() (string, error) {
	return f.cloneAndFindSchemas("swordfish", SwordfishGitRepo)
}

// cloneAndFindSchemas shallow-clones a git repository into a named subdirectory
// of the temp dir and returns the path to the json-schema directory within it.
func (f *Fetcher) cloneAndFindSchemas(name, repoURL string) (string, error) {
	cloneDir := filepath.Join(f.tempDir, name)

	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, cloneDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to clone %s repository: %w\n%s", name, err, string(output))
	}

	schemaDir := filepath.Join(cloneDir, "json-schema")
	if _, err := os.Stat(schemaDir); err != nil {
		matches, _ := filepath.Glob(filepath.Join(cloneDir, "*/json-schema"))
		if len(matches) > 0 {
			schemaDir = matches[0]
		} else {
			return "", fmt.Errorf("json-schema directory not found in cloned repository")
		}
	}

	return schemaDir, nil
}

