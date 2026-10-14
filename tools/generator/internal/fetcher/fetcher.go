//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
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
	if err := cloneSchemas(repoURL, tag, cloneDir); err != nil {
		return "", "", fmt.Errorf("failed to clone %s repository: %w", name, err)
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

// cloneAttempts and cloneRetryDelay bound the retries of a failed clone. The
// publication repositories are large, and a transfer from a CI runner
// occasionally dies with "early EOF".
var (
	cloneAttempts   = 3
	cloneRetryDelay = 10 * time.Second
)

// cloneSchemas shallow-clones repoURL at tag into dir, checking out only the
// json-schema directory, and retries a failed clone.
func cloneSchemas(repoURL, tag, dir string) error {
	var err error
	for attempt := 1; attempt <= cloneAttempts; attempt++ {
		if attempt > 1 {
			log.Printf("Clone of %s failed, retrying (%d/%d): %v", repoURL, attempt, cloneAttempts, err)
			time.Sleep(cloneRetryDelay * time.Duration(attempt-1))
		}
		if err = os.RemoveAll(dir); err != nil {
			return err
		}
		if err = run("git", "-c", "advice.detachedHead=false", "clone", "--quiet", "--depth", "1",
			"--branch", tag, "--filter=blob:none", "--sparse", repoURL, dir); err != nil {
			continue
		}
		if err = run("git", "-C", dir, "sparse-checkout", "set", "--no-cone", "json-schema/"); err == nil {
			return nil
		}
	}
	return err
}

// run runs a command and includes its output in the returned error.
func run(name string, args ...string) error {
	output, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s: %w\n%s", name, strings.Join(args, " "), err, output)
	}
	return nil
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
