//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// CleanupHTTPResponse MUST be called for any HTTP response to ensure that it is properly closed.
// This function can safely be called even if the HTTP client returned an error.
func CleanupHTTPResponse(response *http.Response) error {
	if response == nil || response.Body == nil {
		return nil
	}

	_, errFlush := io.Copy(io.Discard, response.Body) // EOF doesn't trigger an error
	errClose := response.Body.Close()

	return errors.Join(errFlush, errClose)
}

// DeferredCleanupHTTPResponse is the same as CleanupHTTPResponse but intended to be used when we don't care about the error.
func DeferredCleanupHTTPResponse(response *http.Response) {
	_ = CleanupHTTPResponse(response)
}

func ParseRetryAfter(retryAfterStr string) (time.Time, error) {
	if retryAfterStr == "" {
		return time.Time{}, ErrIsEmpty
	}

	parsedTime, parseTimeErr := http.ParseTime(retryAfterStr)
	if parseTimeErr == nil {
		return parsedTime, nil
	}

	parsedSeconds, err := strconv.ParseInt(retryAfterStr, 10, 64)
	if err != nil {
		return time.Time{},
			fmt.Errorf("Retry-After value [%s] is neither a time [%w] or duration [%w]",
				retryAfterStr, parseTimeErr, err)
	}

	if parsedSeconds < 0 {
		return time.Time{}, ErrIsNegative
	}

	return time.Now().Add(time.Duration(parsedSeconds) * time.Second), nil
}
