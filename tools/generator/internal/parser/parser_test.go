//
// SPDX-License-Identifier: BSD-3-Clause
//

package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func TestExtractJSONKeyOrder(t *testing.T) {
	// Sample JSON with a specific key order for action parameters
	rawJSON := []byte(`{
  "definitions": {
    "Reset": {
      "parameters": {
        "ResetType": {
          "type": "string",
          "description": "The type of reset"
        },
        "Force": {
          "type": "boolean",
          "description": "Force the reset"
        },
        "Delay": {
          "type": "integer",
          "description": "Delay in seconds"
        }
      }
    }
  }
}`)

	keys := extractJSONKeyOrder(rawJSON, "definitions.Reset.parameters")

	expectedOrder := []string{"ResetType", "Force", "Delay"}

	if len(keys) != len(expectedOrder) {
		t.Errorf("Expected %d keys, got %d. Keys: %v", len(expectedOrder), len(keys), keys)
		return
	}

	for i, expected := range expectedOrder {
		if keys[i] != expected {
			t.Errorf("Expected key at position %d to be %s, got %s", i, expected, keys[i])
		}
	}
}

func TestExtractJSONKeyOrder_DifferentOrder(t *testing.T) {
	// Test with different order to ensure it's not just alphabetical
	rawJSON := []byte(`{
  "definitions": {
    "PowerControl": {
      "parameters": {
        "ZValue": {
          "type": "string"
        },
        "AValue": {
          "type": "string"
        },
        "MValue": {
          "type": "string"
        }
      }
    }
  }
}`)

	keys := extractJSONKeyOrder(rawJSON, "definitions.PowerControl.parameters")

	// Should maintain schema order, not alphabetical
	expectedOrder := []string{"ZValue", "AValue", "MValue"}

	if len(keys) != len(expectedOrder) {
		t.Errorf("Expected %d keys, got %d. Keys: %v", len(expectedOrder), len(keys), keys)
		return
	}

	for i, expected := range expectedOrder {
		if keys[i] != expected {
			t.Errorf("Expected key at position %d to be %s, got %s", i, expected, keys[i])
		}
	}
}

func TestParseActionResponse(t *testing.T) {
	// Create a test schema with an action that has an actionResponse
	schemaJSON := `{
  "definitions": {
    "Actions": {
      "type": "object",
      "properties": {
        "#TestResource.TestAction": {
          "$ref": "#/definitions/TestAction"
        }
      }
    },
    "TestAction": {
      "type": "object",
      "description": "This is a test action",
      "actionResponse": {
        "$ref": "#/definitions/TestActionResponse"
      },
      "parameters": {
        "TestParam": {
          "type": "string"
        }
      }
    },
    "TestActionResponse": {
      "type": "object",
      "properties": {
        "Result": {
          "type": "string"
        }
      }
    },
    "TestResource": {
      "type": "object",
      "properties": {
        "Id": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "@odata.id": {
          "type": "string"
        },
        "@odata.type": {
          "type": "string"
        },
        "Actions": {
          "$ref": "#/definitions/Actions"
        }
      }
    }
  }
}`

	// Write the test schema to a temporary file
	tmpFile := "/tmp/test-schema.json"
	if err := os.WriteFile(tmpFile, []byte(schemaJSON), 0644); err != nil {
		t.Fatalf("Failed to write test schema: %v", err)
	}
	defer os.Remove(tmpFile)

	// Parse the schema
	p := NewParser("/tmp")
	defs, err := p.ParseSchema(tmpFile)
	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	// Find the TestResource definition
	var testResourceDef *schema.Definition
	for _, def := range defs {
		if def.Name == "TestResource" {
			testResourceDef = def
			break
		}
	}

	if testResourceDef == nil {
		t.Fatal("TestResource definition not found")
	}

	// Verify the action has the response type
	if len(testResourceDef.Actions) != 1 {
		t.Fatalf("Expected 1 action, got %d", len(testResourceDef.Actions))
	}

	action := testResourceDef.Actions[0]
	if action.Name != "TestAction" {
		t.Errorf("Expected action name 'TestAction', got '%s'", action.Name)
	}

	if action.ResponseType != "TestActionResponse" {
		t.Errorf("Expected response type 'TestActionResponse', got '%s'", action.ResponseType)
	}
}

func TestResolveLatestVersionNoAnyOf(t *testing.T) {
	schemaJSON := `{
  "definitions": {
    "Example": {
      "type": "object",
      "properties": {
        "Id": { "type": "string" }
      }
    }
  }
}`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "Example.json")
	if err := os.WriteFile(tmpFile, []byte(schemaJSON), 0644); err != nil {
		t.Fatalf("Failed to write test schema: %v", err)
	}

	versionFile, err := ResolveLatestVersion(tmpFile, tmpDir)
	if err != nil {
		t.Fatalf("ResolveLatestVersion failed: %v", err)
	}

	if versionFile != tmpFile {
		t.Fatalf("Expected base file, got %s", versionFile)
	}
}

func TestParseActionParametersFallbackOrder(t *testing.T) {
	p := NewParser("/tmp")
	p.rawSchema = []byte(`{}`)

	paramsMap := map[string]any{
		"Zeta":  map[string]any{"type": "string"},
		"Alpha": map[string]any{"type": "string"},
	}

	params := p.parseActionParameters("Missing", paramsMap)
	if len(params) != 2 {
		t.Fatalf("Expected 2 params, got %d", len(params))
	}

	if params[0].OriginalName != "Alpha" || params[1].OriginalName != "Zeta" {
		t.Fatalf("Expected alphabetical order, got %s then %s", params[0].OriginalName, params[1].OriginalName)
	}
}

func TestParseActionParametersResourceLinks(t *testing.T) {
	// Capacity.json defines CapacitySource, which can be referenced by
	// @odata.id, and the embedded CapacityInfo object.
	dir := t.TempDir()
	capacityJSON := `{"definitions": {
		"CapacitySource": {"anyOf": [
			{"$ref": "http://redfish.dmtf.org/schemas/v1/odata-v4.json#/definitions/idRef"},
			{"$ref": "http://redfish.dmtf.org/schemas/swordfish/v1/Capacity.v1_2_3.json#/definitions/CapacitySource"}
		]},
		"CapacityInfo": {"$ref": "http://redfish.dmtf.org/schemas/swordfish/v1/Capacity.v1_2_3.json#/definitions/CapacityInfo"}
	}}`
	if err := os.WriteFile(filepath.Join(dir, "Capacity.json"), []byte(capacityJSON), 0o600); err != nil {
		t.Fatalf("Failed to write schema: %v", err)
	}
	oldDirs := SchemaDirs
	SchemaDirs = []string{dir}
	t.Cleanup(func() { SchemaDirs = oldDirs })

	p := NewParser("/tmp")
	p.rawSchema = []byte(`{}`)

	paramsMap := map[string]any{
		// A reference to a resource defined in another resource's schema is a link.
		"CapacitySource": map[string]any{
			"$ref": "http://redfish.dmtf.org/schemas/swordfish/v1/Capacity.json#/definitions/CapacitySource",
		},
		// A reference to an embedded object is not.
		"CapacityInfo": map[string]any{
			"$ref": "http://redfish.dmtf.org/schemas/swordfish/v1/Capacity.json#/definitions/CapacityInfo",
		},
		// A reference to a resource collection is sent as a link object.
		"CertificateCollection": map[string]any{
			"$ref": "http://redfish.dmtf.org/schemas/v1/CertificateCollection.json#/definitions/CertificateCollection",
		},
		// A reference to a standalone resource is a link too, even though the name ends in Uri.
		"CertificateUri": map[string]any{
			"$ref": "http://redfish.dmtf.org/schemas/v1/Certificate.json#/definitions/Certificate",
		},
		// An array of resource references is an array of links.
		"Targets": map[string]any{
			"type":  "array",
			"items": map[string]any{"$ref": "http://redfish.dmtf.org/schemas/v1/SoftwareInventory.json#/definitions/SoftwareInventory"},
		},
		// A URI-valued string stays a string.
		"ImageURI": map[string]any{
			"type":   "string",
			"format": "uri-reference",
		},
		// An enum reference keeps its type.
		"ResetType": map[string]any{
			"$ref": "http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/ResetType",
		},
	}

	params := p.parseActionParameters("Missing", paramsMap)

	got := make(map[string]*schema.ActionParameter, len(params))
	for _, param := range params {
		got[param.OriginalName] = param
	}

	want := map[string]struct {
		typ    string
		isLink bool
	}{
		"CapacitySource":        {"string", true},
		"CapacityInfo":          {"CapacityInfo", false},
		"CertificateCollection": {"string", true},
		"CertificateUri":        {"string", true},
		"Targets":               {"[]string", true},
		"ImageURI":              {"string", false},
		"ResetType":             {"ResetType", false},
	}
	for name, w := range want {
		p := got[name]
		if p == nil {
			t.Fatalf("parameter %s missing", name)
		}
		if p.Type != w.typ || p.IsLink != w.isLink {
			t.Errorf("parameter %s: type %q link %v, want %q %v", name, p.Type, p.IsLink, w.typ, w.isLink)
		}
	}
	if !got["Targets"].IsArray || got["CertificateCollection"].IsArray {
		t.Errorf("IsArray wrong: Targets=%v CertificateCollection=%v", got["Targets"].IsArray, got["CertificateCollection"].IsArray)
	}
}

func TestParseLinksAppliesSchemaRenames(t *testing.T) {
	for _, tt := range []struct {
		name       string
		schemaName string
		ref        string
		want       string
	}{
		{"collection", "Chassis", "http://redfish.dmtf.org/schemas/v1/PowerSupplyCollection.json#/definitions/PowerSupplyCollection", "PowerSupplyUnit"},
		{"resource", "Outlet", "http://redfish.dmtf.org/schemas/v1/PowerSupply.json#/definitions/PowerSupply", "PowerSupplyUnit"},
		{"local renamed resource", "PowerSupply", "#/definitions/PowerSupply", "PowerSupplyUnit"},
		{"local legacy resource", "Power", "#/definitions/PowerSupply", "PowerSupply"},
		{"local collection", "PowerSupplyCollection", "#/definitions/PowerSupplyCollection", "PowerSupplyUnit"},
		{"generic OData reference", "Chassis", "http://redfish.dmtf.org/schemas/v1/odata-v4.json#/definitions/idRef", "Entity"},
		{"generic link", "Chassis", "http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/Link", "Target"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(t.TempDir())
			p.schemaName = tt.schemaName
			p.typeMapper.schemaName = tt.schemaName
			for _, shape := range []struct {
				name    string
				prop    map[string]any
				isArray bool
			}{
				{"direct", map[string]any{"$ref": tt.ref}, false},
				{"nullable", map[string]any{"anyOf": []any{map[string]any{"$ref": tt.ref}, map[string]any{"type": "null"}}}, false},
				{"array", map[string]any{"items": map[string]any{"$ref": tt.ref}}, true},
			} {
				t.Run(shape.name, func(t *testing.T) {
					links := p.parseLinks(map[string]any{"properties": map[string]any{"Targets": shape.prop}})
					if len(links) != 1 {
						t.Fatalf("got %d links, want 1", len(links))
					}
					if links[0].Type != tt.want || links[0].IsArray != shape.isArray {
						t.Errorf("link = (%q, %v), want (%q, %v)", links[0].Type, links[0].IsArray, tt.want, shape.isArray)
					}
				})
			}
		})
	}
}
