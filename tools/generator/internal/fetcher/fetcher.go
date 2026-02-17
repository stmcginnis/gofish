//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
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

// FetchRedfish fetches Redfish schemas by shallow cloning the repository at the latest release tag.
func (f *Fetcher) FetchRedfish() (string, error) {
	return f.cloneAndFindSchemas("redfish", RedfishGitRepo, func(tag string) bool { return true })
}

// FetchSwordfish fetches Swordfish schemas by shallow cloning the repository at the latest official release tag.
func (f *Fetcher) FetchSwordfish() (string, error) {
	return f.cloneAndFindSchemas("swordfish", SwordfishGitRepo, func(tag string) bool {
		return strings.HasSuffix(tag, "_Release")
	})
}

// cloneAndFindSchemas finds the latest tag satisfying tagFilter, shallow-clones
// the repository at that tag, and returns the path to the json-schema directory.
func (f *Fetcher) cloneAndFindSchemas(name, repoURL string, tagFilter func(string) bool) (string, error) {
	tag, err := latestTag(repoURL, tagFilter)
	if err != nil {
		return "", fmt.Errorf("failed to determine latest tag for %s: %w", name, err)
	}

	cloneDir := filepath.Join(f.tempDir, name)

	cmd := exec.Command("git", "clone", "--depth", "1", "--branch", tag, repoURL, cloneDir)
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

// latestTag returns the lexicographically last tag from the remote repository
// that satisfies filter. Tags are listed via git ls-remote without requiring a
// local clone.
func latestTag(repoURL string, filter func(string) bool) (string, error) {
	cmd := exec.Command("git", "ls-remote", "--tags", "--refs", repoURL)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to list remote tags: %w", err)
	}

	var tags []string
	for line := range strings.SplitSeq(strings.TrimSpace(string(output)), "\n") {
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			continue
		}
		tag := strings.TrimPrefix(parts[1], "refs/tags/")
		if filter(tag) {
			tags = append(tags, tag)
		}
	}

	if len(tags) == 0 {
		return "", fmt.Errorf("no matching tags found in %s", repoURL)
	}

	sort.Strings(tags)
	return tags[len(tags)-1], nil
}

