// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"

	"crypto/tls"
	"net/http"

	gofish "github.com/stmcginnis/gofish/school"
)

func httpclientForSelfSigned() (client *http.Client, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{
		Transport: tr,
	}

	return client, err
}

func main() {
	// Build a httpclient for bmcs with self signed certs
	httpclient, err := httpclientForSelfSigned()
	if err != nil {
		panic(err)
	}

	// Create a new instance of gofish client
	c, err := gofish.APIClient("https://bmc-ip", httpclient)
	if err != nil {
		panic(err)
	}

	// Attached the client to service root
	service, err := gofish.ServiceRoot(c)
	if err != nil {
		panic(err)
	}

	// Generates a authenticated session
	auth, err := service.CreateSession("my-username", "my-password")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", auth)

	// Assign the token back to our gofish client
	c.Token = auth.Token

	// Query the chassis data using the session token
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}

	// Delete the session and logout
	service.DeleteSession(auth.Session)
}
