//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"reflect"
	"testing"
)

type nestedStruct struct {
	X bool `json:"x"`
	Y []string
}
type F struct {
	Field string
}

type testStruct struct {
	A string `json:"a"`
	B []string
	C map[string]string
	D int
	E nestedStruct
	F
}

func TestGetPatchPayloadFromUpdate(t *testing.T) { //nolint:funlen
	testcases := []struct {
		name           string
		originalEntity interface{}
		updatedEntity  interface{}
		expected       map[string]interface{}
	}{
		{
			name: "no updates were made",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			updatedEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			expected: map[string]interface{}{},
		},
		{
			name: "string field updated",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			updatedEntity: &testStruct{
				A: "defg",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			expected: map[string]interface{}{
				"a": "defg",
			},
		},
		{
			name: "string and slice field updated",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			updatedEntity: &testStruct{
				A: "defg",
				B: []string{"D", "E", "F", "G"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			expected: map[string]interface{}{
				"a": "defg",
				"B": []string{"D", "E", "F", "G"},
			},
		},
		{
			name: "map field updated",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
			},
			updatedEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"D": "d",
					"E": "e",
					"F": "f",
					"G": "g",
				},
				D: 1,
			},
			expected: map[string]interface{}{
				"C": map[string]string{
					"D": "d",
					"E": "e",
					"F": "f",
					"G": "g",
				},
			},
		},
		{
			name: "nested struct field updated",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
				E: nestedStruct{
					X: true,
					Y: []string{"X", "Y"},
				},
			},
			updatedEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"D": "d",
					"E": "e",
					"F": "f",
					"G": "g",
				},
				D: 1,
				E: nestedStruct{
					X: false,
					Y: []string{"Z"},
				},
			},
			expected: map[string]interface{}{
				"C": map[string]string{
					"D": "d",
					"E": "e",
					"F": "f",
					"G": "g",
				},
				"E": map[string]interface{}{
					"x": false,
					"Y": []string{"Z"},
				},
			},
		},
		{
			name: "embedded struct updated",
			originalEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
				E: nestedStruct{
					X: true,
					Y: []string{"X", "Y"},
				},
				F: F{Field: "f"},
			},
			updatedEntity: &testStruct{
				A: "abcd",
				B: []string{"A", "B", "C", "D"},
				C: map[string]string{
					"A": "a",
					"B": "b",
					"C": "c",
					"D": "d",
				},
				D: 1,
				E: nestedStruct{
					X: true,
					Y: []string{"X", "Y"},
				},
				F: F{Field: "fuuuu"},
			},
			expected: map[string]interface{}{
				"Field": "fuuuu",
			},
		},
	}

	for idx := range testcases {
		tc := testcases[idx]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			payload := getPatchPayloadFromUpdate(
				reflect.ValueOf(tc.originalEntity).Elem(),
				reflect.ValueOf(tc.updatedEntity).Elem(),
			)
			if !reflect.DeepEqual(payload, tc.expected) {
				t.Errorf("expected %+v payload, got %+v", tc.expected, payload)
			}
		})
	}
}
