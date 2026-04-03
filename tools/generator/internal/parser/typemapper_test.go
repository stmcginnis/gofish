//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func TestExtractTypeAndNullable(t *testing.T) {
	tests := []struct {
		name         string
		typeVal      any
		wantType     string
		wantNullable bool
	}{
		{
			name:         "simple string type",
			typeVal:      "string",
			wantType:     "string",
			wantNullable: false,
		},
		{
			name:         "simple integer type",
			typeVal:      "integer",
			wantType:     "integer",
			wantNullable: false,
		},
		{
			name:         "simple boolean type",
			typeVal:      "boolean",
			wantType:     "boolean",
			wantNullable: false,
		},
		{
			name:         "nullable string",
			typeVal:      []any{"string", "null"},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "nullable integer",
			typeVal:      []any{"integer", "null"},
			wantType:     "integer",
			wantNullable: true,
		},
		{
			name:         "null first then string",
			typeVal:      []any{"null", "string"},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "single non-null type in array",
			typeVal:      []any{"boolean"},
			wantType:     "boolean",
			wantNullable: false,
		},
		{
			name:         "multi-type returns any",
			typeVal:      []any{"string", "boolean", "number", "null"},
			wantType:     "any",
			wantNullable: true,
		},
		{
			name:         "multi-type without null",
			typeVal:      []any{"string", "integer"},
			wantType:     "any",
			wantNullable: false,
		},
		{
			name:         "multi-type two types with null",
			typeVal:      []any{"string", "number", "null"},
			wantType:     "any",
			wantNullable: true,
		},
		{
			name:         "all non-null types",
			typeVal:      []any{"string", "boolean", "number", "integer"},
			wantType:     "any",
			wantNullable: false,
		},
		{
			name:         "null only",
			typeVal:      []any{"null"},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "empty array",
			typeVal:      []any{},
			wantType:     "string",
			wantNullable: false,
		},
		{
			name:         "nil type",
			typeVal:      nil,
			wantType:     "string",
			wantNullable: false,
		},
		{
			name:         "non-string type in array ignored",
			typeVal:      []any{"string", 42},
			wantType:     "string",
			wantNullable: false,
		},
		{
			name:         "unexpected type defaults to string",
			typeVal:      42,
			wantType:     "string",
			wantNullable: false,
		},
		{
			name:         "duplicate non-null types collapse to single",
			typeVal:      []any{"string", "string", "null"},
			wantType:     "string",
			wantNullable: true,
		},
		{
			name:         "duplicate then distinct widens to any",
			typeVal:      []any{"string", "string", "integer"},
			wantType:     "any",
			wantNullable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotNullable := extractTypeAndNullable(tt.typeVal)
			if gotType != tt.wantType {
				t.Errorf("extractTypeAndNullable() type = %q, want %q", gotType, tt.wantType)
			}
			if gotNullable != tt.wantNullable {
				t.Errorf("extractTypeAndNullable() nullable = %v, want %v", gotNullable, tt.wantNullable)
			}
		})
	}
}

func TestMapType_MultiType(t *testing.T) {
	tm := NewTypeMapper()

	tests := []struct {
		name      string
		propName  string
		prop      *schema.JSONProperty
		wantType  string
		wantPtr   bool
		wantArray bool
	}{
		{
			name:     "AttributeRegistry CurrentValue multi-type",
			propName: "CurrentValue",
			prop: &schema.JSONProperty{
				Type: []any{"string", "boolean", "number", "null"},
			},
			wantType:  "any",
			wantPtr:   false,
			wantArray: false,
		},
		{
			name:     "AttributeRegistry DefaultValue multi-type",
			propName: "DefaultValue",
			prop: &schema.JSONProperty{
				Type: []any{"string", "boolean", "number", "null"},
			},
			wantType:  "any",
			wantPtr:   false,
			wantArray: false,
		},
		{
			name:     "single nullable string",
			propName: "SomeValue",
			prop: &schema.JSONProperty{
				Type: []any{"string", "null"},
			},
			wantType:  "string",
			wantPtr:   false,
			wantArray: false,
		},
		{
			name:     "single type boolean",
			propName: "Enabled",
			prop: &schema.JSONProperty{
				Type: "boolean",
			},
			wantType:  "bool",
			wantPtr:   false,
			wantArray: false,
		},
		{
			name:     "multi-type without null not pointer",
			propName: "MixedField",
			prop: &schema.JSONProperty{
				Type: []any{"string", "integer"},
			},
			wantType:  "any",
			wantPtr:   false,
			wantArray: false,
		},
		{
			name:     "nullable integer maps to pointer int",
			propName: "Count",
			prop: &schema.JSONProperty{
				Type: []any{"integer", "null"},
			},
			wantType:  "int",
			wantPtr:   true,
			wantArray: false,
		},
		{
			name:     "lowercase multi-type does not become Link",
			propName: "currentValue",
			prop: &schema.JSONProperty{
				Type: []any{"string", "boolean", "number", "null"},
			},
			wantType:  "any",
			wantPtr:   false,
			wantArray: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotPtr, gotArray := tm.MapType(tt.propName, tt.prop)
			if gotType != tt.wantType {
				t.Errorf("MapType() type = %q, want %q", gotType, tt.wantType)
			}
			if gotPtr != tt.wantPtr {
				t.Errorf("MapType() pointer = %v, want %v", gotPtr, tt.wantPtr)
			}
			if gotArray != tt.wantArray {
				t.Errorf("MapType() array = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}
