package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

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

	// Attached the client to service root
	service := c.Service

	// Query the update service
	update, err := service.UpdateService()
	if err != nil {
		panic(err)
	}

	// Fetch the firmware inventory
	inventory, err := update.FirmwareInventory()
	if err != nil {
		panic(err)
	}

	// Fetch the inventory members
	firmwares, err := inventory.Members()
	if err != nil {
		panic(err)
	}

	// Loop over the inventory, printing the details
	for _, f := range firmwares {
		fmt.Printf("%v@%v: %v", f.Name, f.Version, f.Description)
	}

	// Get the session token
	session, err := c.GetSession()
	if err != nil {
		// Didn't have a token. Upgrade to a session
		c, err = c.CloneWithSession()
		if err != nil {
			panic(err)
		}
		defer c.Logout()
	}

	// Open the firmware binary
	firmwareReader, err := os.Open("firmware.bin")
	if err != nil {
		panic(err)
	}
	defer firmwareReader.Close()

	if service.Vendor == "HPE" {
		// For HPE systems, need a "parameters" JSON blob
		parameters := map[string]interface{}{
			"UpdateRepository": true,
			"UpdateTarget":     true,
			"ETag":             "sometag",
			"Section":          0,
		}

		// Marshal that to a reader
		parameterBytes, err := json.Marshal(parameters)
		if err != nil {
			panic(err)
		}
		payloadBuffer := bytes.NewReader(parameterBytes)

		// Also need the session key and file contents.
		// Wrap strings in readers
		values := map[string]io.Reader{
			"sessionKey": strings.NewReader(session.Token),
			"parameters": payloadBuffer,
			"file":       firmwareReader,
		}

		// Post the upload to the specified URL
		response, err := c.PostMultipart(update.HTTPPushURI, values)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
	} else {
		panic(fmt.Sprintf("I don't know how to upload firmware to %s", service.Vendor))
	}
}
