package main

import (
	"fmt"

	gofish "github.com/stmcginnis/gofish/school"
)

func main() {
	c := gofish.APIClient("localhost", 5000, false)
	service, _ := gofish.ServiceRoot(c)

	chassis, _ := service.StorageSystems()
	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}
}
