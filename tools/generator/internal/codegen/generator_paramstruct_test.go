//
// SPDX-License-Identifier: BSD-3-Clause
//

package codegen

import (
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/tools/generator/internal/schema"
)

func TestParamStructGeneration(t *testing.T) {
	gen, err := NewGenerator()
	if err != nil {
		t.Fatalf("NewGenerator: %v", err)
	}

	defs := []*schema.Definition{
		{
			Name:         "CertificateService",
			OriginalName: "CertificateService",
			IsEntity:     true,
			Actions: []*schema.Action{
				{
					Name:         "GenerateCSR",
					JSONName:     "#CertificateService.GenerateCSR",
					Description:  "Generate a CSR",
					ResponseType: "GenerateCSRResponse",
					Parameters: []*schema.ActionParameter{
						{Name: "alternativeNames", Type: "[]string", OriginalName: "AlternativeNames", FieldName: "AlternativeNames", FieldDescription: "\t// AlternativeNames are the additional host names.", Required: false, Ordinal: 0},
						{Name: "certificateCollection", Type: "string", OriginalName: "CertificateCollection", FieldName: "CertificateCollection", FieldDescription: "\t// CertificateCollection is a link to a resource collection.", Required: true, Ordinal: 1},
						{Name: "city", Type: "string", OriginalName: "City", FieldName: "City", FieldDescription: "\t// City is the city of the organization.", Required: true, Ordinal: 2},
						{Name: "commonName", Type: "string", OriginalName: "CommonName", FieldName: "CommonName", FieldDescription: "\t// CommonName is the fully qualified domain name.", Required: true, Ordinal: 3},
					},
				},
				{
					Name:        "SmallAction",
					JSONName:    "#CertificateService.SmallAction",
					Description: "A small action",
					Parameters: []*schema.ActionParameter{
						{Name: "param1", Type: "string", OriginalName: "Param1", FieldName: "Param1", FieldDescription: "\t// Param1 is the first parameter.", Required: true, Ordinal: 0},
						{Name: "param2", Type: "int", OriginalName: "Param2", FieldName: "Param2", FieldDescription: "\t// Param2 is the second parameter.", Required: false, Ordinal: 1},
					},
				},
				{
					Name:        "BigTaskAction",
					JSONName:    "#CertificateService.BigTaskAction",
					Description: "A big task action",
					Parameters: []*schema.ActionParameter{
						{Name: "a", Type: "string", OriginalName: "A", FieldName: "A", FieldDescription: "\t// A is param a.", Required: true, Ordinal: 0},
						{Name: "b", Type: "string", OriginalName: "B", FieldName: "B", FieldDescription: "\t// B is param b.", Required: true, Ordinal: 1},
						{Name: "c", Type: "string", OriginalName: "C", FieldName: "C", FieldDescription: "\t// C is param c.", Required: false, Ordinal: 2},
						{Name: "d", Type: "int", OriginalName: "D", FieldName: "D", FieldDescription: "\t// D is param d.", Required: false, Ordinal: 3},
					},
				},
			},
		},
	}

	output, err := gen.Generate("CertificateService", schema.PackageSchemas, defs, false, nil)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	// Verify >3 param action (GenerateCSR) gets a param struct
	if !strings.Contains(output, "CertificateServiceGenerateCSRParameters") {
		t.Error("Expected CertificateServiceGenerateCSRParameters struct to be generated")
	}
	if !strings.Contains(output, "type CertificateServiceGenerateCSRParameters struct") {
		t.Error("Expected param struct type definition")
	}
	if !strings.Contains(output, "func (c *CertificateService) GenerateCSR(params *CertificateServiceGenerateCSRParameters)") {
		t.Error("Expected method signature with params struct")
	}
	// Required field should not have omitempty
	if !strings.Contains(output, "`json:\"CertificateCollection\"`") {
		t.Error("Expected required field without omitempty")
	}
	// Optional field should have omitempty
	if !strings.Contains(output, "`json:\"AlternativeNames,omitempty\"`") {
		t.Error("Expected optional field with omitempty")
	}

	// Verify <=3 param action (SmallAction) keeps positional args
	if strings.Contains(output, "CertificateServiceSmallActionParameters") {
		t.Error("SmallAction should NOT have a param struct (only 2 params)")
	}
	if !strings.Contains(output, "func (c *CertificateService) SmallAction(param1 string, param2 int)") {
		t.Error("SmallAction should use positional parameters")
	}

	// Verify >3 param task action (BigTaskAction) gets param struct
	if !strings.Contains(output, "CertificateServiceBigTaskActionParameters") {
		t.Error("Expected CertificateServiceBigTaskActionParameters struct")
	}
	if !strings.Contains(output, "func (c *CertificateService) BigTaskAction(params *CertificateServiceBigTaskActionParameters)") {
		t.Error("Expected BigTaskAction with params struct")
	}

	// Verify param struct passes params directly (no map construction)
	// GenerateCSR uses ResponseType, so it should call PostWithResponse with params
	if strings.Contains(output, "PostWithResponse(c.generateCSRTarget, payload)") {
		t.Error("GenerateCSR should pass params, not payload map")
	}

	// BigTaskAction is task-type (no ResponseType), should call PostWithTask with params
	if strings.Contains(output, "PostWithTask(c.client,\n\t\tc.bigTaskActionTarget, payload") {
		t.Error("BigTaskAction should pass params, not payload map")
	}

	// SmallAction should still use payload map
	if !strings.Contains(output, "PostWithTask(c.client,\n\t\tc.smallActionTarget, payload") {
		t.Error("SmallAction should still use payload map")
	}

	t.Log("Generated output:\n" + output)
}
