//
// SPDX-License-Identifier: BSD-3-Clause
//

package codegen

import (
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/config"
	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func TestLenientFields(t *testing.T) {
	config.LenientFields["Gauge.Reading"] = "toFloat32"
	defer delete(config.LenientFields, "Gauge.Reading")

	gen, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	code, err := gen.Generate("Gauge", schema.PackageSchemas, []*schema.Definition{{
		Name: "Gauge", OriginalName: "Gauge",
		Properties: []*schema.Property{{Name: "Reading", JSONName: "Reading", Type: "float32", IsPointer: true}},
	}}, false, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"func (g *Gauge) UnmarshalJSON",
		"LenientReading any `json:\"Reading\"`",
		"g.Reading = toFloat32(tmp.LenientReading)",
	} {
		if !strings.Contains(code, want) {
			t.Errorf("missing %q in:\n%s", want, code)
		}
	}
}
