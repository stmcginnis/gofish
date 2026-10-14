//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "net/url"

// AuthToken contains the authentication and session information.
type AuthToken struct {
	Token     string
	Session   string
	Username  string
	Password  string
	BasicAuth bool
}

type authPayload struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

// CreateSession creates a new session and returns the token and id
func CreateSession(c Client, uri, username, password string) (auth *AuthToken, err error) {
	a := &authPayload{
		UserName: username,
		Password: password,
	}

	resp, err := c.Post(uri, a)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return auth, err
	}

	auth = &AuthToken{}
	auth.Token = resp.Header.Get("X-Auth-Token")
	auth.Session = resp.Header.Get("Location")

	if urlParser, err := url.ParseRequestURI(auth.Session); err == nil {
		auth.Session = urlParser.RequestURI()
	}

	return auth, err
}

// DeleteSession deletes a session using the location as argument
func DeleteSession(c Client, sessionURL string) (err error) {
	resp, err := c.Delete(sessionURL)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}
	return nil
}
