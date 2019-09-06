//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"fmt"

	"crypto/tls"
	"net/http"

	"github.com/stmcginnis/gofish"
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

	// Make sure to delete the session and logout
	defer service.DeleteSession(auth.Session)

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
}
