//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	username := "my-username"
	config := gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: username,
		Password: "my-password",
		Insecure: true,
	}
	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Retrieve the service root
	service := c.Service

	// Query the chassis data using the session token
	accts, err := service.AccountService()
	if err != nil {
		panic(err)
	}
	accs, err := accts.Accounts()
	if err != nil {
		panic(err)
	}

	for _, acc := range accs {
		if acc.UserName == username {
			urlE := "/redfish/v1/AccountService/Accounts/"+acc.ID
			payload := make(map[string]string)
			payload["UserName"] = "new-username"
			payload["Password"] = "new-password"
			res, err := c.Patch(urlE, payload)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v\n", res)
		}
	}
}
