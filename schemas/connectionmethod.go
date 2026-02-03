//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #ConnectionMethod.v1_2_0.ConnectionMethod

package schemas

import (
	"encoding/json"
)

type ConnectionMethodType string

const (
	// RedfishConnectionMethodType shall indicate the connection method is Redfish.
	RedfishConnectionMethodType ConnectionMethodType = "Redfish"
	// SNMPConnectionMethodType shall indicate the connection method is SNMP.
	SNMPConnectionMethodType ConnectionMethodType = "SNMP"
	// IPMI15ConnectionMethodType shall indicate the connection method is IPMI 1.5.
	IPMI15ConnectionMethodType ConnectionMethodType = "IPMI15"
	// IPMI20ConnectionMethodType shall indicate the connection method is IPMI 2.0.
	IPMI20ConnectionMethodType ConnectionMethodType = "IPMI20"
	// NETCONFConnectionMethodType shall indicate the connection method is NETCONF.
	NETCONFConnectionMethodType ConnectionMethodType = "NETCONF"
	// OEMConnectionMethodType shall indicate the connection method is OEM. The
	// 'ConnectionMethodVariant' property shall contain further identification
	// information.
	OEMConnectionMethodType ConnectionMethodType = "OEM"
	// ModbusSerialConnectionMethodType shall indicate the connection method is
	// Modbus serial (RTU).
	ModbusSerialConnectionMethodType ConnectionMethodType = "ModbusSerial"
	// ModbusTCPConnectionMethodType shall indicate the connection method is Modbus
	// TCP.
	ModbusTCPConnectionMethodType ConnectionMethodType = "ModbusTCP"
)

type TunnelingProtocolType string

const (
	// SSHTunnelingProtocolType shall indicate that the tunneling protocol is SSH.
	SSHTunnelingProtocolType TunnelingProtocolType = "SSH"
	// OEMTunnelingProtocolType shall indicate that the tunneling protocol is
	// OEM-specific.
	OEMTunnelingProtocolType TunnelingProtocolType = "OEM"
)

// ConnectionMethod shall represent a connection method for a Redfish
// implementation.
type ConnectionMethod struct {
	Entity
	// ConnectionMethodType shall contain an identifier of the connection method.
	ConnectionMethodType ConnectionMethodType
	// ConnectionMethodVariant shall contain an additional identifier of the
	// connection method. This property shall be present if 'ConnectionMethodType'
	// is 'OEM'.
	ConnectionMethodVariant string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// TunnelingProtocol shall contain the tunneling protocol used for this
	// connection method.
	//
	// Version added: v1.1.0
	TunnelingProtocol TunnelingProtocolType
	// aggregationSources are the URIs for AggregationSources.
	aggregationSources []string
	// serialInterface is the URI for SerialInterface.
	serialInterface string
}

// UnmarshalJSON unmarshals a ConnectionMethod object from the raw JSON.
func (c *ConnectionMethod) UnmarshalJSON(b []byte) error {
	type temp ConnectionMethod
	type cLinks struct {
		AggregationSources Links `json:"AggregationSources"`
		SerialInterface    Link  `json:"SerialInterface"`
	}
	var tmp struct {
		temp
		Links cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ConnectionMethod(tmp.temp)

	// Extract the links to other entities for later
	c.aggregationSources = tmp.Links.AggregationSources.ToStrings()
	c.serialInterface = tmp.Links.SerialInterface.String()

	return nil
}

// GetConnectionMethod will get a ConnectionMethod instance from the service.
func GetConnectionMethod(c Client, uri string) (*ConnectionMethod, error) {
	return GetObject[ConnectionMethod](c, uri)
}

// ListReferencedConnectionMethods gets the collection of ConnectionMethod from
// a provided reference.
func ListReferencedConnectionMethods(c Client, link string) ([]*ConnectionMethod, error) {
	return GetCollectionObjects[ConnectionMethod](c, link)
}

// AggregationSources gets the AggregationSources linked resources.
func (c *ConnectionMethod) AggregationSources() ([]*AggregationSource, error) {
	return GetObjects[AggregationSource](c.client, c.aggregationSources)
}

// SerialInterface gets the SerialInterface linked resource.
func (c *ConnectionMethod) SerialInterface() (*SerialInterface, error) {
	if c.serialInterface == "" {
		return nil, nil
	}
	return GetObject[SerialInterface](c.client, c.serialInterface)
}
