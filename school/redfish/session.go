// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// DefaultSessionPath is the default URI for SessionService collections.
const DefaultSessionPath = "/redfish/v1/Sessions"

// Session describes a single connection (session) between a client and a
// Redfish service instance.
type Session struct {
	common.Entity
	Description string
	Modified    string
	UserName    string
}

type AuthToken struct {
	Token   string
	Session string
}

type authPayload struct {
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

// CreateSession creates a new session and returns the token and id
func CreateSession(c common.Client, username string, password string) (auth *AuthToken, err error) {
	a := &authPayload{
		Username: username,
		Password: password,
	}

	payload, err := json.Marshal(a)
	if err != nil {
		return auth, err
	}

	resp, err := c.Post(DefaultSessionPath, payload)
	if err != nil {
		return auth, err
	}

	auth = &AuthToken{}
	auth.Token = resp.Header.Get("X-Auth-Token")
	auth.Session = resp.Header.Get("Location")

	return auth, err
}

// DeleteSession deletes a session using the location as argument
func DeleteSession(c common.Client, url string) (err error) {
	return c.Delete(url)
}

// GetSession will get a Session instance from the Redfish service.
func GetSession(c common.Client, uri string) (*Session, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t Session
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedSessions gets the collection of Sessions
func ListReferencedSessions(c common.Client, link string) ([]*Session, error) {
	var result []*Session
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, sLink := range links.ItemLinks {
		s, err := GetSession(c, sLink)
		if err != nil {
			return result, err
		}
		result = append(result, s)
	}

	return result, nil
}

// ListSessions gets all Session in the system
func ListSessions(c common.Client) ([]*Session, error) {
	return ListReferencedSessions(c, DefaultSessionPath)
}
