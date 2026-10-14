//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCloneSchemas(t *testing.T) {
	src := t.TempDir()
	git := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", src}, args...)...)
		cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@example.com",
			"GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@example.com")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	for path, body := range map[string]string{
		"json-schema/Chassis.json": "{}",
		"mockups/big.json":         "{}",
	} {
		full := filepath.Join(src, path)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	git("init", "--quiet")
	git("config", "uploadpack.allowFilter", "true")
	git("add", ".")
	git("commit", "--quiet", "-m", "schemas")
	git("tag", "2026.2")

	dir := filepath.Join(t.TempDir(), "redfish")
	if err := cloneSchemas("file://"+src, "2026.2", dir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "json-schema", "Chassis.json")); err != nil {
		t.Errorf("json-schema not checked out: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "mockups")); err == nil {
		t.Error("mockups checked out; only json-schema should be")
	}

	cloneRetryDelay = 0
	if err := cloneSchemas("file://"+src, "no-such-tag", filepath.Join(t.TempDir(), "x")); err == nil {
		t.Error("expected an error for a missing tag")
	}
}
