//
// SPDX-License-Identifier: BSD-3-Clause
//

package main

import (
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func enumDef(name string, values ...string) *schema.Definition {
	def := &schema.Definition{Name: name, OriginalName: name, IsEnum: true}
	for _, v := range values {
		def.EnumValues = append(def.EnumValues, &schema.EnumValue{Name: v + name, Value: v})
	}
	return def
}

func TestAssignOwners(t *testing.T) {
	// Shared enum defined by two schemas: Resource owns it, and its values are
	// the union of both copies.
	resource := &schemaUnit{objectName: "Resource", definitions: []*schema.Definition{enumDef("PowerState", "On", "Off")}}
	circuit := &schemaUnit{objectName: "Circuit", definitions: []*schema.Definition{
		{Name: "Circuit", OriginalName: "Circuit"},
		enumDef("PowerState", "On", "Off", "PowerCycle"),
	}}
	// A schema's main type always stays in that schema.
	account := &schemaUnit{objectName: "AccountService", definitions: []*schema.Definition{{Name: "ExternalAccountProvider", OriginalName: "ExternalAccountProvider"}}}
	provider := &schemaUnit{objectName: "ExternalAccountProvider", definitions: []*schema.Definition{{Name: "ExternalAccountProvider", OriginalName: "ExternalAccountProvider"}}}

	assignOwners([]*schemaUnit{account, circuit, provider, resource})

	if len(circuit.definitions) != 1 || circuit.definitions[0].Name != "Circuit" {
		t.Errorf("Circuit kept %v, want only Circuit", circuit.definitions)
	}
	var values []string
	for _, v := range resource.definitions[0].EnumValues {
		values = append(values, v.Value)
	}
	if len(values) != 3 || values[2] != "PowerCycle" {
		t.Errorf("PowerState values = %v, want On, Off, PowerCycle", values)
	}
	if len(account.definitions) != 0 || len(provider.definitions) != 1 {
		t.Errorf("ExternalAccountProvider should be emitted only by its own schema")
	}
}
