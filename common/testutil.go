//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"reflect"
	"strings"
	"testing"
)

func RequireErrorContains(t *testing.T, err error, expectedMsg string) {
	if err == nil {
		t.Fatal("expecting error but got nil")
		return
	}

	msg := err.Error()
	if !strings.Contains(msg, expectedMsg) {
		t.Fatalf("expecting error with msg [%s], found [%s]", expectedMsg, msg)
	}
}

func RequireNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("found an unexpected error: %v", err)
	}
}

func AssertInDelta[T int64 | float32 | float64](t *testing.T, a, b, delta T) {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}

	if diff > delta {
		t.Errorf("expecting [%v] to be within [%v] of [%v]", a, delta, b)
	}
}

func AssertEqual[T any](t *testing.T, expected, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected [%v] not equal to actual [%v]", expected, actual)
	}
}

func AssertEqualMsg[T any](t *testing.T, expected, actual T, msg string, args ...any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(msg, args...)
	}
}

func RequireLessMsg[T int](t *testing.T, a, b T, msg string, args ...any) {
	if a >= b {
		t.Fatalf(msg, args...)
	}
}
