//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Cable.v1_2_4.json
// 2021.4 - #Cable.v1_2_4.Cable

package schemas

import (
	"encoding/json"
)

type CableClass string

const (
	// PowerCableClass This cable is used for connecting to a power system.
	PowerCableClass CableClass = "Power"
	// NetworkCableClass This cable is used for connecting to a networking system.
	NetworkCableClass CableClass = "Network"
	// StorageCableClass This cable is used for connecting to a storage system.
	StorageCableClass CableClass = "Storage"
	// FanCableClass This cable is used for connecting to a fan system.
	FanCableClass CableClass = "Fan"
	// PCIeCableClass This cable is used for connecting to a PCIe endpoint.
	PCIeCableClass CableClass = "PCIe"
	// USBCableClass This cable is used for connecting to a USB endpoint.
	USBCableClass CableClass = "USB"
	// VideoCableClass This cable is used for connecting to a video system.
	VideoCableClass CableClass = "Video"
	// FabricCableClass This cable is used for connecting to a fabric.
	FabricCableClass CableClass = "Fabric"
	// SerialCableClass This cable is used for connecting to a serial endpoint.
	SerialCableClass CableClass = "Serial"
	// GeneralCableClass This cable is used for providing general connectivity.
	GeneralCableClass CableClass = "General"
)

type CableStatus string

const (
	// NormalCableStatus shall indicate the cable is operating normally. The
	// 'State' property in 'Status' shall contain the value 'Enabled' and the
	// 'Health' property in 'Status' shall contain the value 'OK'.
	NormalCableStatus CableStatus = "Normal"
	// DegradedCableStatus shall indicate the cable is degraded. The 'State'
	// property in 'Status' shall contain the value 'Enabled' and the 'Health'
	// property in 'Status' shall contain the value 'Warning'.
	DegradedCableStatus CableStatus = "Degraded"
	// FailedCableStatus shall indicate the cable has failed. The 'State' property
	// in 'Status' shall contain the value 'Enabled' and the 'Health' property in
	// 'Status' shall contain the value 'Critical'.
	FailedCableStatus CableStatus = "Failed"
	// TestingCableStatus shall indicate the cable is under test. The 'State'
	// property in 'Status' shall contain the value 'InTest'.
	TestingCableStatus CableStatus = "Testing"
	// DisabledCableStatus shall indicate the cable is disabled. The 'State'
	// property in 'Status' shall contain the value 'Disabled'.
	DisabledCableStatus CableStatus = "Disabled"
	// SetByServiceCableStatus shall indicate the status for the cable is not
	// defined by the user. If implemented, the service shall determine the value
	// of the 'State' and 'Health' properties in 'Status'.
	SetByServiceCableStatus CableStatus = "SetByService"
)

type CableConnectorType string

const (
	// ACPowerCableConnectorType This cable connects to an AC power connector.
	ACPowerCableConnectorType CableConnectorType = "ACPower"
	// DB9CableConnectorType This cable connects to a DB9 connector.
	DB9CableConnectorType CableConnectorType = "DB9"
	// DCPowerCableConnectorType This cable connects to a DC power connector.
	DCPowerCableConnectorType CableConnectorType = "DCPower"
	// DisplayPortCableConnectorType This cable connects to a DisplayPort power
	// connector.
	DisplayPortCableConnectorType CableConnectorType = "DisplayPort"
	// HDMICableConnectorType This cable connects to an HDMI connector.
	HDMICableConnectorType CableConnectorType = "HDMI"
	// ICICableConnectorType This cable connects to an ICI connector.
	ICICableConnectorType CableConnectorType = "ICI"
	// IPASSCableConnectorType This cable connects to an IPASS connector.
	IPASSCableConnectorType CableConnectorType = "IPASS"
	// PCIeCableConnectorType This cable connects to a PCIe connector.
	PCIeCableConnectorType CableConnectorType = "PCIe"
	// ProprietaryCableConnectorType This cable connects to a proprietary connector.
	ProprietaryCableConnectorType CableConnectorType = "Proprietary"
	// RJ45CableConnectorType This cable connects to an RJ45 connector.
	RJ45CableConnectorType CableConnectorType = "RJ45"
	// SATACableConnectorType This cable connects to a SATA connector.
	SATACableConnectorType CableConnectorType = "SATA"
	// SCSICableConnectorType This cable connects to a SCSI connector.
	SCSICableConnectorType CableConnectorType = "SCSI"
	// SlimSASCableConnectorType This cable connects to a SlimSAS connector.
	SlimSASCableConnectorType CableConnectorType = "SlimSAS"
	// SFPCableConnectorType This cable connects to an SFP connector.
	SFPCableConnectorType CableConnectorType = "SFP"
	// SFPPlusCableConnectorType This cable connects to an SFPPlus connector.
	SFPPlusCableConnectorType CableConnectorType = "SFPPlus"
	// USBACableConnectorType This cable connects to a USB-A connector.
	USBACableConnectorType CableConnectorType = "USBA"
	// USBCCableConnectorType This cable connects to a USB-C connector.
	USBCCableConnectorType CableConnectorType = "USBC"
	// QSFPCableConnectorType This cable connects to a QSFP connector.
	QSFPCableConnectorType CableConnectorType = "QSFP"
	// CDFPCableConnectorType This cable connects to a CDFP connector.
	CDFPCableConnectorType CableConnectorType = "CDFP"
	// OSFPCableConnectorType This cable connects to an OSFP connector.
	OSFPCableConnectorType CableConnectorType = "OSFP"
)

// Cable This resource contains a simple cable for a Redfish implementation.
type Cable struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// AssetTag shall track the cable for inventory purposes.
	AssetTag string
	// CableClass shall contain the cable class for this cable.
	CableClass CableClass
	// CableStatus shall contain the user-reported status of this resource.
	CableStatus CableStatus
	// CableType shall contain a user-defined type for this cable.
	CableType string
	// DownstreamConnectorTypes shall contain an array of connector types this
	// cable supports.
	DownstreamConnectorTypes []CableConnectorType
	// DownstreamName shall contain any identifier for a downstream resource.
	DownstreamName string
	// LengthMeters shall contain the length of the cable in meters.
	LengthMeters *float64 `json:",omitempty"`
	// Location shall contain the location information of the associated assembly.
	Location Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the cable. This organization might be the entity from whom the
	// cable is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to
	// the cable.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number assigned by the organization that
	// is responsible for producing or manufacturing the cable.
	PartNumber string
	// SKU shall contain the stock-keeping unit (SKU) number for this cable.
	SKU string
	// SerialNumber shall contain the manufacturer-allocated number that identifies
	// the cable.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UpstreamConnectorTypes shall contain an array of connector types this cable
	// supports.
	UpstreamConnectorTypes []CableConnectorType
	// UpstreamName shall contain any identifier for an upstream resource.
	UpstreamName string
	// UserDescription shall contain a user-defined description for this cable.
	UserDescription string
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.1.0
	UserLabel string
	// Vendor shall contain the name of the company that provides the final product
	// that includes this cable.
	Vendor string
	// downstreamChassis are the URIs for DownstreamChassis.
	downstreamChassis []string
	// downstreamPorts are the URIs for DownstreamPorts.
	downstreamPorts []string
	// downstreamResources are the URIs for DownstreamResources.
	downstreamResources []string
	// upstreamChassis are the URIs for UpstreamChassis.
	upstreamChassis []string
	// upstreamPorts are the URIs for UpstreamPorts.
	upstreamPorts []string
	// upstreamResources are the URIs for UpstreamResources.
	upstreamResources []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Cable object from the raw JSON.
func (c *Cable) UnmarshalJSON(b []byte) error {
	type temp Cable
	type cLinks struct {
		DownstreamChassis   Links `json:"DownstreamChassis"`
		DownstreamPorts     Links `json:"DownstreamPorts"`
		DownstreamResources Links `json:"DownstreamResources"`
		UpstreamChassis     Links `json:"UpstreamChassis"`
		UpstreamPorts       Links `json:"UpstreamPorts"`
		UpstreamResources   Links `json:"UpstreamResources"`
	}
	var tmp struct {
		temp
		Links    cLinks
		Assembly Link `json:"Assembly"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Cable(tmp.temp)

	// Extract the links to other entities for later
	c.downstreamChassis = tmp.Links.DownstreamChassis.ToStrings()
	c.downstreamPorts = tmp.Links.DownstreamPorts.ToStrings()
	c.downstreamResources = tmp.Links.DownstreamResources.ToStrings()
	c.upstreamChassis = tmp.Links.UpstreamChassis.ToStrings()
	c.upstreamPorts = tmp.Links.UpstreamPorts.ToStrings()
	c.upstreamResources = tmp.Links.UpstreamResources.ToStrings()
	c.assembly = tmp.Assembly.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Cable) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"CableClass",
		"CableStatus",
		"CableType",
		"DownstreamConnectorTypes",
		"DownstreamName",
		"LengthMeters",
		"Manufacturer",
		"Model",
		"PartNumber",
		"SKU",
		"SerialNumber",
		"UpstreamConnectorTypes",
		"UpstreamName",
		"UserDescription",
		"UserLabel",
		"Vendor",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCable will get a Cable instance from the service.
func GetCable(c Client, uri string) (*Cable, error) {
	return GetObject[Cable](c, uri)
}

// ListReferencedCables gets the collection of Cable from
// a provided reference.
func ListReferencedCables(c Client, link string) ([]*Cable, error) {
	return GetCollectionObjects[Cable](c, link)
}

// DownstreamChassis gets the DownstreamChassis linked resources.
func (c *Cable) DownstreamChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.downstreamChassis)
}

// DownstreamPorts gets the DownstreamPorts linked resources.
func (c *Cable) DownstreamPorts() ([]*Port, error) {
	return GetObjects[Port](c.client, c.downstreamPorts)
}

// DownstreamResources gets the DownstreamResources linked resources.
func (c *Cable) DownstreamResources() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.downstreamResources)
}

// UpstreamChassis gets the UpstreamChassis linked resources.
func (c *Cable) UpstreamChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.upstreamChassis)
}

// UpstreamPorts gets the UpstreamPorts linked resources.
func (c *Cable) UpstreamPorts() ([]*Port, error) {
	return GetObjects[Port](c.client, c.upstreamPorts)
}

// UpstreamResources gets the UpstreamResources linked resources.
func (c *Cable) UpstreamResources() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.upstreamResources)
}

// Assembly gets the Assembly linked resource.
func (c *Cable) Assembly() (*Assembly, error) {
	if c.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](c.client, c.assembly)
}
