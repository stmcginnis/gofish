//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #SerialInterface.v1_3_0.SerialInterface

package schemas

import (
	"encoding/json"
)

type BitRate string

const (
	// 1200BitRate is a bit rate of 1200 bit/s.
	V1200BitRate BitRate = "1200"
	// 2400BitRate is a bit rate of 2400 bit/s.
	V2400BitRate BitRate = "2400"
	// 4800BitRate is a bit rate of 4800 bit/s.
	V4800BitRate BitRate = "4800"
	// 9600BitRate is a bit rate of 9600 bit/s.
	V9600BitRate BitRate = "9600"
	// 19200BitRate is a bit rate of 19200 bit/s.
	V19200BitRate BitRate = "19200"
	// 38400BitRate is a bit rate of 38400 bit/s.
	V38400BitRate BitRate = "38400"
	// 57600BitRate is a bit rate of 57600 bit/s.
	V57600BitRate BitRate = "57600"
	// 115200BitRate is a bit rate of 115200 bit/s.
	V115200BitRate BitRate = "115200"
	// 230400BitRate is a bit rate of 230400 bit/s.
	V230400BitRate BitRate = "230400"
)

type ConnectorType string

const (
	// RJ45ConnectorType is an RJ45 connector.
	RJ45ConnectorType ConnectorType = "RJ45"
	// RJ11ConnectorType is an RJ11 connector.
	RJ11ConnectorType ConnectorType = "RJ11"
	// DB9FemaleConnectorType is a DB9 Female connector.
	DB9FemaleConnectorType ConnectorType = "DB9 Female"
	// DB9MaleConnectorType is a DB9 Male connector.
	DB9MaleConnectorType ConnectorType = "DB9 Male"
	// DB25FemaleConnectorType is a DB25 Female connector.
	DB25FemaleConnectorType ConnectorType = "DB25 Female"
	// DB25MaleConnectorType is a DB25 Male connector.
	DB25MaleConnectorType ConnectorType = "DB25 Male"
	// USBConnectorType is a USB connector.
	USBConnectorType ConnectorType = "USB"
	// mUSBConnectorType is a mUSB connector.
	MUSBConnectorType ConnectorType = "mUSB"
	// uUSBConnectorType is a uUSB connector.
	UUSBConnectorType ConnectorType = "uUSB"
)

type DataBits string

const (
	// 5DataBits Five bits of data following the start bit.
	V5DataBits DataBits = "5"
	// 6DataBits Six bits of data following the start bit.
	V6DataBits DataBits = "6"
	// 7DataBits Seven bits of data following the start bit.
	V7DataBits DataBits = "7"
	// 8DataBits Eight bits of data following the start bit.
	V8DataBits DataBits = "8"
)

type SerialInferfaceFlowControl string

const (
	// NoneSerialInferfaceFlowControl No flow control imposed.
	NoneSerialInferfaceFlowControl SerialInferfaceFlowControl = "None"
	// SoftwareSerialInferfaceFlowControl XON/XOFF in-band flow control imposed.
	SoftwareSerialInferfaceFlowControl SerialInferfaceFlowControl = "Software"
	// HardwareSerialInferfaceFlowControl Out-of-band flow control imposed.
	HardwareSerialInferfaceFlowControl SerialInferfaceFlowControl = "Hardware"
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

type InterfaceRole string

const (
	// ClientInterfaceRole The serial interface is a client and connects to one or more
	// servers across a serial bus.
	ClientInterfaceRole InterfaceRole = "Client"
	// ServerInterfaceRole The serial interface is a server and allows one or more clients
	// to connect across a serial bus.
	ServerInterfaceRole InterfaceRole = "Server"
)

type SignalType string

const (
	// Rs232SignalType The serial interface follows RS232.
	Rs232SignalType SignalType = "Rs232"
	// Rs485SignalType The serial interface follows RS485.
	Rs485SignalType SignalType = "Rs485"
)

type StopBits string

const (
	// 1StopBits One stop bit following the data bits.
	V1StopBits StopBits = "1"
	// 2StopBits Two stop bits following the data bits.
	V2StopBits StopBits = "2"
)

// SerialInterface shall represent a serial interface as part of the Redfish
// Specification.
type SerialInterface struct {
	Entity
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
	FlowControl SerialInferfaceFlowControl
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
	// OEM shall contain the OEM extensions. All values for properties that this
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
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
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SerialInterface) Update() error {
	readWriteFields := []string{
		"BitRate",
		"DataBits",
		"FlowControl",
		"InterfaceEnabled",
		"Parity",
		"StopBits",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSerialInterface will get a SerialInterface instance from the service.
func GetSerialInterface(c Client, uri string) (*SerialInterface, error) {
	return GetObject[SerialInterface](c, uri)
}

// ListReferencedSerialInterfaces gets the collection of SerialInterface from
// a provided reference.
func ListReferencedSerialInterfaces(c Client, link string) ([]*SerialInterface, error) {
	return GetCollectionObjects[SerialInterface](c, link)
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
	Role InterfaceRole
	// ServerID shall contain the MODBUS over Serial Line Specification and
	// Implementation guide-defined address that identifies this Modbus RTU server.
	// This property shall only be present if the 'Role' property contains
	// 'Server'.
	//
	// Version added: v1.3.0
	ServerID *uint `json:"ServerId,omitempty"`
}
