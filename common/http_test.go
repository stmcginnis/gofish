package common_test

import (
	"testing"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			expectErr: common.ErrIsEmpty.Error(),
		},
		"negative seconds": {
			input:     "-120",
			expectErr: common.ErrIsNegative.Error(),
		},
		"junk": {
			input:     "not a date",
			expectErr: "neither",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res, err := common.ParseRetryAfter(testCase.input)
			if testCase.expectErr != "" {
				require.ErrorContains(t, err, testCase.expectErr)
			} else {
				require.NoError(t, err)
				expectedTime := testCase.output.UnixMilli()
				actualTime := res.UnixMilli()

				assert.InDelta(t, expectedTime, actualTime, 100)
			}
		})
	}
}
