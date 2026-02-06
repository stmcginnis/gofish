# Query Session

This is an example of querying the current active session information from the
service.

```go
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
	config := gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
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

	// Query the session service
	sessionService, err := service.SessionService()
	if err != nil {
		panic(err)
	}

	// Get the active sessions
	sessions, err := sessionService.Sessions()
	if err != nil {
		panic(err)
	}

	for _, session := range sessions {
		fmt.Printf("Session: %#v\n\n", session)
	}
}
```
