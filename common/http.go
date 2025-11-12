package common

import (
	"errors"
	"io"
	"net/http"
)

// CleanupHttpResponse MUST be called for any HTTP response to ensure that it is properly closed.
// This function can safely be called even if the HTTP client returned an error.
func CleanupHttpResponse(response *http.Response) error {
	if response == nil || response.Body == nil {
		return nil
	}

	_, errFlush := io.Copy(io.Discard, response.Body) // EOF doesn't trigger an error
	errClose := response.Body.Close()

	return errors.Join(errFlush, errClose)
}
