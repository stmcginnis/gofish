//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #SerialInterface.v1_3_0.SerialInterface

package redfish

import (
	"encoding/json"

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

// SerialConnectionFlowControl is the type of flow control, if any, that is
// imposed on the serial connection.
type SerialConnectionFlowControl string

const (
	// Out-of-band flow control imposed.
	HardwareSerialConnectionFlowControl SerialConnectionFlowControl = "Hardware"
	// No flow control imposed.
	NoneSerialConnectionFlowControl SerialConnectionFlowControl = "None"
	// XON/XOFF in-band flow control imposed.
	SoftwareSerialConnectionFlowControl SerialConnectionFlowControl = "Software"
)

type Parity string

const (
	// NoneParity No parity bit.
	NoneParity Parity = "None"
	// EvenParity is an even parity bit.
	EvenParity Parity = "Even"
	// OddParity is an odd parity bit.
	OddParity Parity = "Odd"
	// MarkParity is a mark parity bit.
	MarkParity Parity = "Mark"
	// SpaceParity is a space parity bit.
	SpaceParity Parity = "Space"
)

type PinOut string

const (
	// CiscoPinOut The Cisco pinout configuration.
	CiscoPinOut PinOut = "Cisco"
	// CycladesPinOut The Cyclades pinout configuration.
	CycladesPinOut PinOut = "Cyclades"
	// DigiPinOut The Digi pinout configuration.
	DigiPinOut PinOut = "Digi"
	// Modbus2WirePinOut shall indicate the MODBUS over Serial Line Specification
	// and Implementation guide-defined 2W-MODBUS pinout. This value shall only be
	// used if the 'SignalType' property contains 'Rs485'.
	Modbus2WirePinOut PinOut = "Modbus2Wire"
	// Modbus4WirePinOut shall indicate the MODBUS over Serial Line Specification
	// and Implementation guide-defined 4W-MODBUS pinout. This value shall only be
	// used if the 'SignalType' property contains 'Rs485'.
	Modbus4WirePinOut PinOut = "Modbus4Wire"
	// ModbusRs232PinOut shall indicate the MODBUS over Serial Line Specification
	// and Implementation guide-defined RS232-MODBUS pinout. This value shall only
	// be used if the 'SignalType' property contains 'Rs232'.
	ModbusRs232PinOut PinOut = "ModbusRs232"
)

type ModbusRole string

const (
	// ClientModbusRole The serial interface is a client and connects to one or more
	// servers across a serial bus.
	ClientModbusRole ModbusRole = "Client"
	// ServerModbusRole The serial interface is a server and allows one or more clients
	// to connect across a serial bus.
	ServerModbusRole ModbusRole = "Server"
)

type SignalType string

const (
	// Rs232SignalType The serial interface follows RS232.
	Rs232SignalType SignalType = "Rs232"
	// Rs485SignalType The serial interface follows RS485.
	Rs485SignalType SignalType = "Rs485"
)

// The period of time before the next start bit is transmitted.
type StopBits string

const (
	// One stop bit following the data bits.
	StopBits1 DataBits = "1"
	// Two stop bits following the data bits.
	StopBits2 DataBits = "2"
)

// SerialInterface shall represent a serial interface as part of the Redfish
// Specification.
type SerialInterface struct {
	common.Entity
	// BitRate shall indicate the transmit and receive speed of the serial
	// connection.
	BitRate BitRate
	// ConnectorType shall indicate the type of physical connector used for this
	// serial connection.
	ConnectorType ConnectorType
	// DataBits shall indicate number of data bits for the serial connection.
	DataBits DataBits
	// FlowControl shall indicate the flow control mechanism for the serial
	// connection.
	FlowControl SerialConnectionFlowControl
	// InterfaceEnabled shall indicate whether this interface is enabled.
	InterfaceEnabled bool
	// Modbus shall contain the Modbus settings for this serial interface.
	//
	// Version added: v1.3.0
	Modbus Modbus
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Parity shall indicate parity information for a serial connection.
	Parity Parity
	// PinOut shall indicate the physical pinout for the serial connector.
	PinOut PinOut
	// SignalType shall contain the type of serial signaling in use for the serial
	// connection.
	SignalType SignalType
	// StopBits shall indicate the stop bits for the serial connection.
	StopBits StopBits
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SerialInterface object from the raw JSON.
func (s *SerialInterface) UnmarshalJSON(b []byte) error {
	type temp SerialInterface
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SerialInterface(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SerialInterface) Update() error {
	readWriteFields := []string{
		"BitRate",
		"DataBits",
		"FlowControl",
		"InterfaceEnabled",
		"Modbus",
		"Parity",
		"StopBits",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSerialInterface will get a SerialInterface instance from the service.
func GetSerialInterface(c common.Client, uri string) (*SerialInterface, error) {
	return common.GetObject[SerialInterface](c, uri)
}

// ListReferencedSerialInterfaces gets the collection of SerialInterface from
// a provided reference.
func ListReferencedSerialInterfaces(c common.Client, link string) ([]*SerialInterface, error) {
	return common.GetCollectionObjects[SerialInterface](c, link)
}

// Modbus shall contain the Modbus settings for this serial interface.
type Modbus struct {
	// ReadOnly shall indicate if the serial interface is read only. If 'true', the
	// serial service on this manager shall reject or ignore requests that modify
	// data. This property shall only be present if the 'Role' property contains
	// 'Server'.
	//
	// Version added: v1.3.0
	ReadOnly bool
	// Role shall contain the role of this serial interface.
	//
	// Version added: v1.3.0
	Role ModbusRole
	// ServerId shall contain the MODBUS over Serial Line Specification and
	// Implementation guide-defined address that identifies this Modbus RTU server.
	// This property shall only be present if the 'Role' property contains
	// 'Server'.
	//
	// Version added: v1.3.0
	ServerID *uint `json:"ServerId,omitempty"`
}
