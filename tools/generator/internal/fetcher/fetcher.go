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
	"strconv"
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

// FetchRedfish fetches Redfish schemas by shallow cloning the repository at
// tag, or at the latest release tag if tag is empty. It returns the schema
// directory and the tag used.
func (f *Fetcher) FetchRedfish(tag string) (string, string, error) {
	return f.cloneAndFindSchemas("redfish", RedfishGitRepo, tag, func(tag string) bool { return true })
}

// FetchSwordfish fetches Swordfish schemas by shallow cloning the repository at
// tag, or at the latest official release tag if tag is empty. It returns the
// schema directory and the tag used.
func (f *Fetcher) FetchSwordfish(tag string) (string, string, error) {
	return f.cloneAndFindSchemas("swordfish", SwordfishGitRepo, tag, func(tag string) bool {
		return strings.HasSuffix(tag, "_Release")
	})
}

// cloneAndFindSchemas shallow-clones the repository at tag (or the latest tag
// satisfying tagFilter when tag is empty) and returns the path to the
// json-schema directory and the tag used.
func (f *Fetcher) cloneAndFindSchemas(name, repoURL, tag string, tagFilter func(string) bool) (string, string, error) {
	if tag == "" {
		var err error
		tag, err = latestTag(repoURL, tagFilter)
		if err != nil {
			return "", "", fmt.Errorf("failed to determine latest tag for %s: %w", name, err)
		}
	}

	cloneDir := filepath.Join(f.tempDir, name)

	cmd := exec.Command("git", "clone", "--depth", "1", "--branch", tag, repoURL, cloneDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("failed to clone %s repository: %w\n%s", name, err, string(output))
	}

	schemaDir := filepath.Join(cloneDir, "json-schema")
	if _, err := os.Stat(schemaDir); err != nil {
		matches, _ := filepath.Glob(filepath.Join(cloneDir, "*/json-schema"))
		if len(matches) > 0 {
			schemaDir = matches[0]
		} else {
			return "", "", fmt.Errorf("json-schema directory not found in cloned repository")
		}
	}

	return schemaDir, tag, nil
}

// latestTag returns the highest version tag from the remote repository that
// satisfies filter. Tags are listed via git ls-remote without requiring a
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

	sort.Slice(tags, func(i, j int) bool { return versionLess(tags[i], tags[j]) })
	return tags[len(tags)-1], nil
}

// versionLess orders tags by their numeric parts, so "v1.2.10" sorts after
// "v1.2.9" and "2026.10" after "2026.2". Non-numeric parts compare as strings.
func versionLess(a, b string) bool {
	as, bs := splitVersion(a), splitVersion(b)
	for i := 0; i < len(as) && i < len(bs); i++ {
		if as[i] == bs[i] {
			continue
		}
		an, aErr := strconv.Atoi(as[i])
		bn, bErr := strconv.Atoi(bs[i])
		if aErr == nil && bErr == nil {
			return an < bn
		}
		return as[i] < bs[i]
	}
	return len(as) < len(bs)
}

// splitVersion splits s into alternating runs of digits and non-digits.
func splitVersion(s string) []string {
	var parts []string
	start := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || isDigit(s[i]) != isDigit(s[start]) {
			parts = append(parts, s[start:i])
			start = i
		}
	}
	return parts
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }
