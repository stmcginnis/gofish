//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"testing"
	"time"
)

func TestParseRetryAfter(t *testing.T) {
	testCases := map[string]struct {
		input     string
		output    time.Time
		expectErr string
	}{
		"date value": {
			input:  "Wed, 21 Oct 2015 07:28:00 GMT",
			output: time.Date(2015, 10, 21, 7, 28, 0, 0, time.UTC),
		},
		"seconds": {
			input:  "120",
			output: time.Now().Add(120 * time.Second),
		},
		"blank input": {
			input:     "",
			expectErr: ErrIsEmpty.Error(),
		},
		"negative seconds": {
			input:     "-120",
			expectErr: ErrIsNegative.Error(),
		},
		"junk": {
			input:     "not a date",
			expectErr: "neither",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase // to support go 1.21
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res, err := ParseRetryAfter(testCase.input)
			if testCase.expectErr != "" {
				RequireErrorContains(t, err, testCase.expectErr)
			} else {
				RequireNoError(t, err)
				expectedTime := testCase.output.UnixMilli()
				actualTime := res.UnixMilli()

				AssertInDelta(t, expectedTime, actualTime, 100)
			}
		})
	}
}
