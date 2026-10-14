//
// SPDX-License-Identifier: BSD-3-Clause
//

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func TestHandwrittenDeclarationsAreSkipped(t *testing.T) {
	dir := t.TempDir()
	hand := `package schemas

type Custom struct{}

func (w *Widget) UnmarshalJSON(b []byte) error { return nil }

func (g *Gadget) postUnmarshal(b []byte) error { return nil }
`
	generated := "package schemas\n\n" + GeneratedMarker + "\n\ntype Ignored struct{}\n"
	if err := os.WriteFile(filepath.Join(dir, "widget_custom.go"), []byte(hand), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "widget.go"), []byte(generated), 0o600); err != nil {
		t.Fatal(err)
	}

	names, err := ScanHandwritten(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !names["Custom"] || !names["Widget.UnmarshalJSON"] || names["Ignored"] {
		t.Fatalf("ScanHandwritten = %v", names)
	}

	gen, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	gen.SetHandwritten(names)
	defs := []*schema.Definition{
		{
			Name: "Widget", OriginalName: "Widget", IsEntity: true,
			Links: []*schema.Link{{Name: "Chassis", JSONName: "Chassis", Type: "Chassis"}},
		},
		{Name: "Custom", OriginalName: "Custom"},
	}
	code, err := gen.Generate("Widget", schema.PackageSchemas, defs, false, nil)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(code, "func (w *Widget) UnmarshalJSON") {
		t.Error("generated UnmarshalJSON despite a hand-written one")
	}
	if strings.Contains(code, "type Custom struct") {
		t.Error("generated Custom despite a hand-written one")
	}
	if !strings.Contains(code, "func (w *Widget) Chassis()") || !strings.Contains(code, GeneratedMarker) {
		t.Error("expected the Chassis getter and the generated-file marker")
	}

	// A hand-written postUnmarshal hook is called at the end of the generated
	// UnmarshalJSON, which is generated even for a type with no links.
	code, err = gen.Generate("Gadget", schema.PackageSchemas, []*schema.Definition{
		{Name: "Gadget", OriginalName: "Gadget", Properties: []*schema.Property{{Name: "Size", JSONName: "Size", Type: "int"}}},
	}, false, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(code, "func (g *Gadget) UnmarshalJSON") || !strings.Contains(code, "return g.postUnmarshal(b)") {
		t.Errorf("expected UnmarshalJSON calling postUnmarshal, got:\n%s", code)
	}
}
