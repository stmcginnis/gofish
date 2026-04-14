# Query Chassis

This is an example of querying the chassis information.

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

	// Query the chassis data using the session token
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	for _, ch := range chassis {
		fmt.Printf("Chassis: %s\n", ch.Name)
		fmt.Printf("\tType:         %s\n", ch.ChassisType)
		fmt.Printf("\tManufacturer: %s\n", ch.Manufacturer)
		fmt.Printf("\tSerial:       %s\n", ch.SerialNumber)
		fmt.Printf("\tPower state:  %s\n", ch.PowerState)
	}
}
```
