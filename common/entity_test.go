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
	Field string `json:",omitempty"`
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
		originalEntity any
		updatedEntity  any
		expected       map[string]any
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
			expected: map[string]any{},
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
			expected: map[string]any{
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
			expected: map[string]any{
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
			expected: map[string]any{
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
			expected: map[string]any{
				"C": map[string]string{
					"D": "d",
					"E": "e",
					"F": "f",
					"G": "g",
				},
				"E": map[string]any{
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
			expected: map[string]any{
				"Field": "fuuuu", // Обновленное значение
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

type User struct {
	Name   string `json:"username"`
	Age    int
	Email  *string
	Roles  []string
	Data   Data
	Hidden string `json:"-"`
}

type Data struct {
	DarkMode bool
	Profile  Profile
}

type Profile struct {
	Avatar string
}

func TestGetPatchPayloadFromUpdateWithPointer(t *testing.T) {
	type testCase struct {
		name     string
		original any
		updated  any
		expected map[string]any
	}

	strPtr := func(s string) *string { return &s }

	tests := []testCase{
		{
			name: "No changes",
			original: User{
				Name: "Alice",
				Age:  30,
			},
			updated: User{
				Name: "Alice",
				Age:  30,
			},
			expected: map[string]any{},
		},
		{
			name: "Simple field change",
			original: User{
				Name: "Alice",
			},
			updated: User{
				Name: "Bob",
			},
			expected: map[string]any{
				"username": "Bob",
			},
		},
		{
			name: "Pointer field change",
			original: User{
				Email: strPtr("old@test.com"),
			},
			updated: User{
				Email: strPtr("new@test.com"),
			},
			expected: map[string]any{
				"Email": "new@test.com",
			},
		},
		{
			name: "Slice change",
			original: User{
				Roles: []string{"admin"},
			},
			updated: User{
				Roles: []string{"admin", "user"},
			},
			expected: map[string]any{
				"Roles": []string{"admin", "user"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := reflect.ValueOf(tt.original)
			updated := reflect.ValueOf(tt.updated)
			result := getPatchPayloadFromUpdate(original, updated)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s:\nExpected: %+v\nGot:      %+v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestNestedStructChanges(t *testing.T) {
	tests := []struct {
		name     string
		original User
		updated  User
		expected map[string]any
	}{
		{
			name: "Nested struct change",
			original: User{
				Data: Data{
					DarkMode: true,
					Profile: Profile{
						Avatar: "old.png",
					},
				},
			},
			updated: User{
				Data: Data{
					DarkMode: false,
					Profile: Profile{
						Avatar: "new.png",
					},
				},
			},
			expected: map[string]any{
				"Data": map[string]any{
					"DarkMode": false,
					"Profile": map[string]any{
						"Avatar": "new.png",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPatchPayloadFromUpdate(reflect.ValueOf(tt.original), reflect.ValueOf(tt.updated))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s:\nExpected: %+v\nGot:      %+v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestIgnoredFields(t *testing.T) {
	t.Run("Ignored field", func(t *testing.T) {
		original := User{Hidden: "secret"}
		updated := User{Hidden: "new-secret"}
		result := getPatchPayloadFromUpdate(reflect.ValueOf(original), reflect.ValueOf(updated))

		if len(result) != 0 {
			t.Errorf("Expected empty map for ignored field, got %v", result)
		}
	})
}
func TestEdgeCases(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		result := getPatchPayloadFromUpdate(reflect.ValueOf(nil), reflect.ValueOf(nil))
		if len(result) != 0 {
			t.Errorf("Expected empty map for nil input, got %v", result)
		}
	})

	t.Run("Different types", func(t *testing.T) {
		result := getPatchPayloadFromUpdate(
			reflect.ValueOf(User{}),
			reflect.ValueOf(struct{}{}),
		)
		if len(result) != 0 {
			t.Errorf("Expected empty map for different types, got %v", result)
		}
	})
}
func TestGetPatchPayloadFromUpdate_EdgeCases(t *testing.T) {
	t.Run("Nil values", func(t *testing.T) {
		var original *User
		var updated *User
		result := getPatchPayloadFromUpdate(reflect.ValueOf(original), reflect.ValueOf(updated))
		if len(result) != 0 {
			t.Errorf("Expected empty map for nil values, got %v", result)
		}
	})

	t.Run("Non-struct types", func(t *testing.T) {
		original := "string"
		updated := "another string"
		result := getPatchPayloadFromUpdate(reflect.ValueOf(original), reflect.ValueOf(updated))
		if len(result) != 0 {
			t.Errorf("Expected empty map for non-struct types, got %v", result)
		}
	})

	t.Run("Different types", func(t *testing.T) {
		original := User{Name: "Alice"}
		updated := struct{ Name string }{Name: "Bob"}
		result := getPatchPayloadFromUpdate(reflect.ValueOf(original), reflect.ValueOf(updated))
		if len(result) != 0 {
			t.Errorf("Expected empty map for different types, got %v", result)
		}
	})
}
