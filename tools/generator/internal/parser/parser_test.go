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
		"Zeta": map[string]any{"type": "string"},
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
