package main

import (
	"fmt"
	"os"
	s "strings"

	gofish "github.com/stmcginnis/gofish/school"
)

func main() {
	queryObject := ""
	if len(os.Args) > 1 {
		queryObject = s.ToLower(os.Args[1])
	}
	c := gofish.APIClient("localhost", 5000, false)
	service, _ := gofish.ServiceRoot(c)

	switch queryObject {
	case "chassis":
		objs, _ := service.Chassis()
		for _, obj := range objs {
			fmt.Printf("Chassis: %#v\n\n", obj)
		}
	case "managers":
		objs, _ := service.Managers()
		for _, obj := range objs {
			fmt.Printf("Manager: %#v\n\n", obj)
		}
	case "tasks":
		objs, _ := service.Tasks()
		for _, obj := range objs {
			fmt.Printf("Task: %#v\n\n", obj)
		}
	case "sessions":
		objs, _ := service.Sessions()
		for _, obj := range objs {
			fmt.Printf("Session: %#v\n\n", obj)
		}
	case "storageservices":
		objs, _ := service.StorageServices()
		for _, obj := range objs {
			fmt.Printf("Storage service: %#v\n\n", obj)
		}
	case "storagesystems":
		objs, _ := service.StorageSystems()
		for _, obj := range objs {
			fmt.Printf("Storage system: %#v\n\n", obj)
		}
	case "accounts":
		obj, _ := service.AccountService()
		fmt.Printf("Account service: %#v\n\n", obj)
	case "events":
		obj, _ := service.EventService()
		fmt.Printf("Event service: %#v\n\n", obj)
	case "systems":
		objs, _ := service.Systems()
		for _, obj := range objs {
			fmt.Printf("System: %#v\n\n", obj)
		}
	default:
		fmt.Printf("ServiceRoot: %#v\n\n", service)
	}
}
