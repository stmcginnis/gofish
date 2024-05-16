# List Drives

This is an example of listing information about the drives in a system.

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

	systems, err := service.Systems()
	if err != nil {
		panic(err)
	}

	for _, system := range systems {
		storage, err := system.Storage()
		if err != nil {
			continue
		}

		for _, ss := range storage {
			drives, err := ss.Drives()
			if err != nil {
				continue
			}

			for i, drive := range drives {
				fmt.Printf("Drive %d\n", i)
				fmt.Printf("\tManufacturer: %s\n", drive.Manufacturer)
				fmt.Printf("\tModel: %s\n", drive.Model)
				fmt.Printf("\tSize: %d GiB\n", (drive.CapacityBytes / 1024 / 1024 / 1024))
				fmt.Printf("\tSerial number: %s\n", drive.SerialNumber)
				fmt.Printf("\tPart number: %s\n", drive.PartNumber)
				fmt.Printf("\tLocation: %s %d\n", drive.PhysicalLocation.PartLocation.LocationType, drive.PhysicalLocation.PartLocation.LocationOrdinalValue)
			}
		}
	}
}
```
