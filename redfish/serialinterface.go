//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// The receive and transmit rate of data flow, typically in bits per second (bit/s), over the serial connection.
type BitRate string

const (
	// A bit rate of 1200 bit/s.
	BitRate1200 BitRate = "1200"
	// A bit rate of 2400 bit/s.
	BitRate2400 BitRate = "2400"
	// A bit rate of 4800 bit/s.
	BitRate4800 BitRate = "4800"
	// A bit rate of 9600 bit/s.
	BitRate9600 BitRate = "9600"
	// A bit rate of 19200 bit/s.
	BitRate19200 BitRate = "19200"
	// A bit rate of 38400 bit/s.
	BitRate38400 BitRate = "38400"
	// A bit rate of 57600 bit/s.
	BitRate57600 BitRate = "57600"
	// A bit rate of 115200 bit/s.
	BitRate115200 BitRate = "115200"
	// A bit rate of 230400 bit/s.
	BitRate230400 BitRate = "230400"
)

// The type of connector used for this interface.
type ConnectorType string

const (
	// A DB25 Female connector.
	DB25FemaleConnector ConnectorType = "DB25 Female"
	// A DB25 Male connector.
	DB25MaleConnector ConnectorType = "DB25 Male"
	// A DB9 Female connector.
	DB9FemaleConnector ConnectorType = "DB9 Female"
	// A DB9 Male connector.
	DB9MaleConnector ConnectorType = "DB9 Male"
	// A mUSB connector.
	MUSBConnector ConnectorType = "mUSB"
	// An RJ11 connector.
	RJ11Connector ConnectorType = "RJ11"
	// An RJ45 connector.
	RJ45Connector ConnectorType = "RJ45"
	// A USB connector.
	USBConnector ConnectorType = "USB"
	// A uUSB connector.
	UUSBConnector ConnectorType = "uUSB"
)

// The number of data bits that follow the start bit over the serial connection.
type DataBits string

const (
	// Five bits of data following the start bit.
	DataBits5 DataBits = "5"
	// Six bits of data following the start bit.
	DataBits6 DataBits = "6"
	// Seven bits of data following the start bit.
	DataBits7 DataBits = "7"
	// Eight bits of data following the start bit.
	DataBits8 DataBits = "8"
)

// The period of time before the next start bit is transmitted.
type StopBits string

const (
	// One stop bit following the data bits.
	StopBits1 DataBits = "1"
	// Two stop bits following the data bits.
	StopBits2 DataBits = "2"
)

// The type of flow control, if any, that is imposed on the serial connection.
type SerialConnectionFlowControl string

const (
	// Out-of-band flow control imposed.
	HardwareSerialConnectionFlowControl SerialConnectionFlowControl = "Hardware"
	// No flow control imposed.
	NoneSerialConnectionFlowControl SerialConnectionFlowControl = "None"
	// XON/XOFF in-band flow control imposed.
	SoftwareSerialConnectionFlowControl SerialConnectionFlowControl = "Software"
)

// The type of parity used by the sender and receiver to detect errors over the serial connection.
type Parity string

const (
	// An even parity bit.
	EvenParityBit Parity = "Even"
	// A mark parity bit.
	MarkParityBit Parity = "Mark"
	// No parity bit.
	NoneParityBit Parity = "None"
	// An odd parity bit.
	OddParityBit Parity = "Odd"
	// A space parity bit.
	SpaceParityBit Parity = "Space"
)

// The physical pinout configuration for a serial connector.
type PinOutConfiguration string

const (
	// The Cisco pinout configuration.
	CiscoPinOutConfiguration PinOutConfiguration = "Cisco"
	// The Cyclades pinout configuration.
	CycladesPinOutConfiguration PinOutConfiguration = "Cyclades"
	// The Digi pinout configuration.
	DigiPinOutConfiguration PinOutConfiguration = "Digi"
)

// The type of signal used for the communication connection.
type SignalType string

const (
	// The serial interface follows RS232.
	Rs232SignalType SignalType = "Rs232"
	// The serial interface follows RS485.
	Rs485SignalType SignalType = "Rs485"
)

// SerialInterface is used to represent Serial Interface resources as part of
// the Redfish specification.
type SerialInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	BitRate          BitRate
	ConnectorType    ConnectorType
	DataBits         DataBits
	FlowControl      SerialConnectionFlowControl
	InterfaceEnabled bool
	Parity           Parity
	PinOut           PinOutConfiguration
	SignalType       SignalType
	StopBits         StopBits
	Oem              json.RawMessage

	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
	rawData    []byte
}

// UnmarshalJSON unmarshals a SerialInterface object from the raw JSON.
func (serialInterface *SerialInterface) UnmarshalJSON(b []byte) error {
	type temp SerialInterface
	type actions struct {
		Oem json.RawMessage // OEM actions will be stored here
	}

	var t struct {
		temp
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*serialInterface = SerialInterface(t.temp)

	// Extract the links to other entities for later
	serialInterface.OemActions = t.Actions.Oem

	// This is a read/write object, so we need to save the raw object data for later
	serialInterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (serialInterface *SerialInterface) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SerialInterface)
	if err := original.UnmarshalJSON(serialInterface.rawData); err != nil {
		return err
	}

	readWriteFields := []string{
		"BitRate",
		"DataBits",
		"FlowControl",
		"InterfaceEnabled",
		"Parity",
		"StopBits",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(serialInterface).Elem()

	return serialInterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSerialInterface will get a SerialInterface instance from the service.
func GetSerialInterface(c common.Client, uri string) (*SerialInterface, error) {
	var serialInterface SerialInterface
	return &serialInterface, serialInterface.Get(c, uri, &serialInterface)
}

// ListReferencedSerialInterfaces gets the collection of SerialInterface from
// a provided reference.
func ListReferencedSerialInterfaces(c common.Client, link string) ([]*SerialInterface, error) { //nolint:dupl
	var result []*SerialInterface
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *SerialInterface
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		serialInterface, err := GetSerialInterface(c, link)
		ch <- GetResult{Item: serialInterface, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
