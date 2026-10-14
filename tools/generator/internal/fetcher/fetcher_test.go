//
// SPDX-License-Identifier: BSD-3-Clause
//

package fetcher

import (
	"sort"
	"testing"
)

func TestVersionLess(t *testing.T) {
	tags := []string{"v1.2.10_Release", "v1.2.9_Release", "v1.1.0a_Release", "v1.2.1c_Release"}
	sort.Slice(tags, func(i, j int) bool { return versionLess(tags[i], tags[j]) })
	want := []string{"v1.1.0a_Release", "v1.2.1c_Release", "v1.2.9_Release", "v1.2.10_Release"}
	for i := range want {
		if tags[i] != want[i] {
			t.Fatalf("sorted = %v, want %v", tags, want)
		}
	}

	if !versionLess("2026.2", "2026.10") {
		t.Error("2026.2 should sort before 2026.10")
	}
}
