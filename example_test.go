//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish_test

import (
	"fmt"
	"log"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/schemas"
)

// ExampleConnect demonstrates the basic pattern for connecting to a Redfish
// service and ensuring the session is cleaned up on exit.
func ExampleConnect() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	fmt.Println(c.Service.RedfishVersion)
}

// ExampleConnect_basicAuth shows how to use HTTP Basic Auth instead of
// session-based authentication. Basic Auth avoids creating a server-side
// session but sends credentials with every request.
func ExampleConnect_basicAuth() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint:  "https://bmc-ip",
		Username:  "my-username",
		Password:  "my-password",
		Insecure:  true,
		BasicAuth: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	fmt.Println(c.Service.RedfishVersion)
}

// ExampleConnect_reuseSession shows how to save a session token from one
// connection and reuse it in a later connection, avoiding the overhead of
// creating a new session each time.
func ExampleConnect_reuseSession() {
	// Initial connection that creates a session.
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Save the session token for later reuse.
	session, err := c.GetSession()
	if err != nil {
		log.Fatal(err)
	}

	// ... save session.ID and session.Token to persistent storage ...

	// Reconnect later using the saved session instead of credentials.
	c2, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Insecure: true,
		Session: &gofish.Session{
			ID:    session.ID,
			Token: session.Token,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c2.Logout()

	fmt.Println(c2.Service.RedfishVersion)
}

// Example_querySystemInventory demonstrates how to retrieve processor and
// memory information from all systems registered with the service.
func Example_querySystemInventory() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	systems, err := c.Service.Systems()
	if err != nil {
		log.Fatal(err)
	}

	for _, system := range systems {
		fmt.Printf("System: %s\n", system.Name)

		processors, err := system.Processors()
		if err != nil {
			log.Printf("error getting processors: %v", err)
		}
		for _, p := range processors {
			fmt.Printf("  CPU: %s %s\n", p.Manufacturer, p.Model)
			fmt.Printf("    Cores: %d  Threads: %d\n",
				gofish.Deref(p.TotalCores), gofish.Deref(p.TotalThreads))
			fmt.Printf("    Max speed: %d MHz\n", gofish.Deref(p.MaxSpeedMHz))
		}

		memory, err := system.Memory()
		if err != nil {
			log.Printf("error getting memory: %v", err)
		}
		for _, dimm := range memory {
			fmt.Printf("  DIMM: %s  %d MiB  %d MHz\n",
				dimm.Name,
				gofish.Deref(dimm.CapacityMiB),
				gofish.Deref(dimm.OperatingSpeedMhz))
		}
	}
}

// Example_queryThermal shows how to read temperature sensor and fan speed
// readings from all chassis managed by the service.
func Example_queryThermal() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	chassis, err := c.Service.Chassis()
	if err != nil {
		log.Fatal(err)
	}

	for _, ch := range chassis {
		thermal, err := ch.Thermal()
		if err != nil {
			log.Printf("error getting thermal for chassis %s: %v", ch.Name, err)
			continue
		}

		fmt.Printf("Chassis: %s\n", ch.Name)

		for _, temp := range thermal.Temperatures {
			if temp.ReadingCelsius == nil {
				continue
			}
			fmt.Printf("  Temp: %-30s  %.1f °C\n", temp.Name, *temp.ReadingCelsius)
		}

		for _, fan := range thermal.Fans {
			if fan.Reading == nil {
				continue
			}
			fmt.Printf("  Fan:  %-30s  %d %s\n", fan.Name, *fan.Reading, fan.ReadingUnits)
		}
	}
}

// Example_queryManagers shows how to list BMC/management controller details,
// such as firmware version and network protocol configuration.
func Example_queryManagers() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	managers, err := c.Service.Managers()
	if err != nil {
		log.Fatal(err)
	}

	for _, mgr := range managers {
		fmt.Printf("Manager: %s  type=%s  firmware=%s\n",
			mgr.Name, mgr.ManagerType, mgr.FirmwareVersion)

		netProto, err := mgr.NetworkProtocol()
		if err != nil {
			continue
		}
		fmt.Printf("  Hostname: %s\n", netProto.HostName)
	}
}

// Example_firmwareInventory shows how to list all firmware components
// reported by the UpdateService, which is useful for auditing installed
// firmware versions across the managed system.
func Example_firmwareInventory() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	updateService, err := c.Service.UpdateService()
	if err != nil {
		log.Fatal(err)
	}

	inventory, err := updateService.FirmwareInventory()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range inventory {
		fmt.Printf("%-40s  version=%-20s  updateable=%v\n",
			item.Name, item.Version, item.Updateable)
	}
}

// Example_firmwareUpdate demonstrates a simple firmware update using an
// image URI. The update is submitted to the service and a task monitor is
// returned for polling completion status.
func Example_firmwareUpdate() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	updateService, err := c.Service.UpdateService()
	if err != nil {
		log.Fatal(err)
	}

	taskInfo, err := updateService.SimpleUpdate(&schemas.UpdateServiceSimpleUpdateParameters{
		ImageURI:         "https://firmware-server/bmc-firmware-2.0.bin",
		TransferProtocol: schemas.HTTPSTransferProtocolType,
	})
	if err != nil {
		log.Fatal(err)
	}

	if taskInfo != nil {
		fmt.Printf("Update task submitted: %s\n", taskInfo.TaskMonitor)
	}
}

// Example_readEventLog shows how to read entries from the system event log
// via the Manager's LogService. This is useful for diagnostics and auditing.
func Example_readEventLog() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	managers, err := c.Service.Managers()
	if err != nil {
		log.Fatal(err)
	}

	for _, mgr := range managers {
		logServices, err := mgr.LogServices()
		if err != nil {
			continue
		}

		for _, ls := range logServices {
			entries, err := ls.Entries()
			if err != nil {
				continue
			}

			fmt.Printf("Log: %s\n", ls.Name)
			for _, entry := range entries {
				fmt.Printf("  [%s] %s: %s\n",
					entry.Severity, entry.Created, entry.Message)
			}
		}
	}
}

// Example_subscribeEvents shows how to register a webhook endpoint to
// receive Redfish event notifications. The returned subscription URI can
// be used later to modify or delete the subscription.
func Example_subscribeEvents() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	eventService, err := c.Service.EventService()
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe using registry prefixes (Redfish v1.5+).
	subscriptionURI, err := eventService.CreateEventSubscriptionInstance(
		"https://my-event-receiver/redfish/events", // destination
		[]string{"Alert", "ResourceEvent"},          // registry prefixes
		[]string{},                                   // all resource types
		nil,                                          // no custom HTTP headers
		schemas.RedfishEventDestinationProtocol,
		"my-monitoring-service", // client-supplied context string
		schemas.RetryForeverDeliveryRetryPolicy,
		nil, // no OEM data
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Subscription created: %s\n", subscriptionURI)

	// Delete the subscription when no longer needed.
	if err := eventService.DeleteEventSubscription(subscriptionURI); err != nil {
		log.Fatal(err)
	}
}

// Example_updateBiosAttributes shows how to read and modify BIOS settings.
// Changes are applied at the next reboot unless an immediate apply time
// is specified.
func Example_updateBiosAttributes() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	systems, err := c.Service.Systems()
	if err != nil {
		log.Fatal(err)
	}

	for _, system := range systems {
		bios, err := system.Bios()
		if err != nil {
			log.Printf("error getting BIOS for %s: %v", system.Name, err)
			continue
		}

		// Print current attribute values.
		for name, value := range bios.Attributes {
			fmt.Printf("  %s = %v\n", name, value)
		}

		// Update an attribute to take effect on next reboot.
		if err := bios.UpdateBiosAttributesApplyAt(
			schemas.SettingsAttributes{"NumaGroupSizeOpt": "Flat"},
			schemas.OnResetSettingsApplyTime,
		); err != nil {
			log.Printf("error updating BIOS: %v", err)
		}
	}
}

// Example_mountVirtualMedia demonstrates how to mount a remote ISO image as
// virtual media on a manager, enabling virtual CD/DVD boot without physical
// media.
func Example_mountVirtualMedia() {
	c, err := gofish.Connect(gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	managers, err := c.Service.Managers()
	if err != nil {
		log.Fatal(err)
	}

	for _, mgr := range managers {
		virtualMedia, err := mgr.VirtualMedia()
		if err != nil {
			continue
		}

		for _, vm := range virtualMedia {
			// Find the CD/DVD slot.
			if vm.MediaTypes == nil {
				continue
			}
			for _, mt := range vm.MediaTypes {
				if mt != schemas.CDVirtualMediaType && mt != schemas.DVDVirtualMediaType {
					continue
				}

				_, err := vm.InsertMedia(&schemas.VirtualMediaInsertMediaParameters{
					Image:                "https://file-server/os-installer.iso",
					TransferProtocolType: gofish.ToRef(schemas.HTTPSTransferProtocolType),
					Inserted:             gofish.ToRef(true),
					WriteProtected:       gofish.ToRef(true),
				})
				if err != nil {
					log.Printf("error inserting media: %v", err)
					continue
				}

				fmt.Printf("Mounted ISO on %s\n", vm.Name)
			}
		}
	}
}
