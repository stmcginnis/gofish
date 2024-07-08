//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	// OEMConnectionMethodType shall indicate the connection method is OEM. The ConnectionMethodVariant property shall
	// contain further identification information.
	OEMConnectionMethodType ConnectionMethodType = "OEM"
)

type TunnelingProtocolType string

const (
	// SSHTunnelingProtocolType shall indicate that the tunneling protocol is SSH.
	SSHTunnelingProtocolType TunnelingProtocolType = "SSH"
	// OEMTunnelingProtocolType shall indicate that the tunneling protocol is OEM-specific.
	OEMTunnelingProtocolType TunnelingProtocolType = "OEM"
)

// ConnectionMethod shall represent a connection method for a Redfish implementation.
type ConnectionMethod struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConnectionMethodType shall contain an identifier of the connection method.
	ConnectionMethodType ConnectionMethodType
	// ConnectionMethodVariant shall contain an additional identifier of the connection method. This property shall be
	// present if ConnectionMethodType is 'OEM'.
	ConnectionMethodVariant string
	// Description provides a description of this resource.
	Description string
	// TunnelingProtocol shall contain the tunneling protocol used for this connection method.
	TunnelingProtocol TunnelingProtocolType

	aggregationSources []string
	// AggregationSourcesCount is the number of AggregationSource that are using this connection method.
	AggregationSourcesCount int
}

// UnmarshalJSON unmarshals a ConnectionMethod object from the raw JSON.
func (connectionmethod *ConnectionMethod) UnmarshalJSON(b []byte) error {
	type temp ConnectionMethod
	type Links struct {
		// AggregationSources shall contain an array of links to resources of type AggregationSource that are using this
		// connection method.
		AggregationSources common.Links
		// AggregationSources@odata.count
		AggregationSourcesCount int `json:"AggregationSources@odata.count"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connectionmethod = ConnectionMethod(t.temp)

	// Extract the links to other entities for later
	connectionmethod.aggregationSources = t.Links.AggregationSources.ToStrings()
	connectionmethod.AggregationSourcesCount = t.Links.AggregationSourcesCount

	return nil
}

// GetConnectionMethod will get a ConnectionMethod instance from the service.
func GetConnectionMethod(c common.Client, uri string) (*ConnectionMethod, error) {
	return common.GetObject[ConnectionMethod](c, uri)
}

// ListReferencedConnectionMethods gets the collection of ConnectionMethod from
// a provided reference.
func ListReferencedConnectionMethods(c common.Client, link string) ([]*ConnectionMethod, error) {
	return common.GetCollectionObjects[ConnectionMethod](c, link)
}

// AggregationSources gets the access points using this connection method.
func (connectionmethod *ConnectionMethod) AggregationSources() ([]*AggregationSource, error) {
	return common.GetObjects[AggregationSource](connectionmethod.GetClient(), connectionmethod.aggregationSources)
}
