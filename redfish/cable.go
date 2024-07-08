//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type CableClass string

const (
	// PowerCableClass is used for connecting to a power system.
	PowerCableClass CableClass = "Power"
	// NetworkCableClass is used for connecting to a networking system.
	NetworkCableClass CableClass = "Network"
	// StorageCableClass is used for connecting to a storage system.
	StorageCableClass CableClass = "Storage"
	// FanCableClass is used for connecting to a fan system.
	FanCableClass CableClass = "Fan"
	// PCIeCableClass is used for connecting to a PCIe endpoint.
	PCIeCableClass CableClass = "PCIe"
	// USBCableClass is used for connecting to a USB endpoint.
	USBCableClass CableClass = "USB"
	// VideoCableClass is used for connecting to a video system.
	VideoCableClass CableClass = "Video"
	// FabricCableClass is used for connecting to a fabric.
	FabricCableClass CableClass = "Fabric"
	// SerialCableClass is used for connecting to a serial endpoint.
	SerialCableClass CableClass = "Serial"
	// GeneralCableClass is used for providing general connectivity.
	GeneralCableClass CableClass = "General"
)

type CableStatus string

const (
	// NormalCableStatus shall indicate the cable is operating normally. The State property in Status shall contain the
	// value 'Enabled' and the Health property in Status shall contain the value 'OK'.
	NormalCableStatus CableStatus = "Normal"
	// DegradedCableStatus shall indicate the cable is degraded. The State property in Status shall contain the value
	// 'Enabled' and the Health property in Status shall contain the value 'Warning'.
	DegradedCableStatus CableStatus = "Degraded"
	// FailedCableStatus shall indicate the cable has failed. The State property in Status shall contain the value
	// 'Enabled' and the Health property in Status shall contain the value 'Critical'.
	FailedCableStatus CableStatus = "Failed"
	// TestingCableStatus shall indicate the cable is under test. The State property in Status shall contain the value
	// 'InTest'.
	TestingCableStatus CableStatus = "Testing"
	// DisabledCableStatus shall indicate the cable is disabled. The State property in Status shall contain the value
	// 'Disabled'.
	DisabledCableStatus CableStatus = "Disabled"
	// SetByServiceCableStatus shall indicate the status for the cable is not defined by the user. If implemented, the
	// service shall determine the value of the State and Health properties in Status.
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
	// DisplayPortCableConnectorType This cable connects to a DisplayPort power connector.
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

// Cable contains a simple cable for a Redfish implementation.
type Cable struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// AssetTag shall track the cable for inventory purposes.
	AssetTag string
	// CableClass shall contain the cable class for this cable.
	CableClass CableClass
	// CableStatus shall contain the user-reported status of this resource.
	CableStatus CableStatus
	// CableType shall contain a user-defined type for this cable.
	CableType string
	// Description provides a description of this resource.
	Description string
	// DownstreamConnectorTypes shall contain an array of connector types this cable supports.
	DownstreamConnectorTypes []CableConnectorType
	// DownstreamName shall contain any identifier for a downstream resource.
	DownstreamName string
	// LengthMeters shall contain the length of the cable in meters.
	LengthMeters float64
	// Location shall contain the location information of the associated assembly.
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for producing the cable. This organization
	// might be the entity from whom the cable is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the cable.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number assigned by the organization that is responsible for producing or
	// manufacturing the cable.
	PartNumber string
	// SKU shall contain the stock-keeping unit (SKU) number for this cable.
	SKU string
	// SerialNumber shall contain the manufacturer-allocated number that identifies the cable.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpstreamConnectorTypes shall contain an array of connector types this cable supports.
	UpstreamConnectorTypes []CableConnectorType
	// UpstreamName shall contain any identifier for an upstream resource.
	UpstreamName string
	// UserDescription shall contain a user-defined description for this cable.
	UserDescription string
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Vendor shall contain the name of the company that provides the final product that includes this cable.
	Vendor string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	downstreamChassis []string
	// DownstreamChassisCount is the number of physical downstream containers connected to this cable.
	DownstreamChassisCount int
	downstreamPorts        []string
	// DownstreamPortsCount is the number of physical downstream connections connected to this cable.
	DownstreamPortsCount int
	downstreamResources  []string
	// DownstreamResourcesCount is the number of physical downstream resources connected to this cable.
	DownstreamResourcesCount int
	upstreamChassis          []string
	// UpstreamChassisCount is the number of physical upstream containers connected to this cable.
	UpstreamChassisCount int
	upstreamPorts        []string
	// UpstreamPortsCount is the number of physical upstream connections connected to this cable.
	UpstreamPortsCount int
	upstreamResources  []string
	// UpstreamResourcesCount is the number of physical upstream resources connected to this cable.
	UpstreamResourcesCount int
}

// UnmarshalJSON unmarshals a Cable object from the raw JSON.
func (cable *Cable) UnmarshalJSON(b []byte) error {
	type temp Cable
	type Links struct {
		DownstreamChassis        common.Links
		DownstreamChassisCount   int `json:"DownstreamChassis@odata.count"`
		DownstreamPorts          common.Links
		DownstreamPortsCount     int `json:"DownstreamPorts@odata.count"`
		DownstreamResources      common.Links
		DownstreamResourcesCount int             `json:"DownstreamResources@odata.count"`
		OEM                      json.RawMessage `json:"Oem"`
		UpstreamChassis          common.Links
		UpstreamChassisCount     int `json:"UpstreamChassis@odata.count"`
		UpstreamPorts            common.Links
		UpstreamPortsCount       int `json:"UpstreamPorts@odata.count"`
		UpstreamResources        common.Links
		UpstreamResourcesCount   int `json:"UpstreamResources@odata.count"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cable = Cable(t.temp)

	// Extract the links to other entities for later
	cable.downstreamChassis = t.Links.DownstreamChassis.ToStrings()
	cable.downstreamPorts = t.Links.DownstreamPorts.ToStrings()
	cable.downstreamResources = t.Links.DownstreamResources.ToStrings()
	cable.upstreamChassis = t.Links.UpstreamChassis.ToStrings()
	cable.upstreamPorts = t.Links.UpstreamPorts.ToStrings()
	cable.upstreamResources = t.Links.UpstreamResources.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	cable.rawData = b

	return nil
}

// DownstreamChassis gets the physical downstream containers connected to this cable.
func (cable *Cable) DownstreamChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](cable.GetClient(), cable.downstreamChassis)
}

// DownstreamPorts gets the physical downstream connections connected to this cable.
func (cable *Cable) DownstreamPorts() ([]*Port, error) {
	return common.GetObjects[Port](cable.GetClient(), cable.downstreamPorts)
}

// UpstreamChassis gets the physical upstream containers connected to this cable.
func (cable *Cable) UpstreamChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](cable.GetClient(), cable.upstreamChassis)
}

// UpstreamPorts gets the physical upstream connections connected to this cable.
func (cable *Cable) UptreamPorts() ([]*Port, error) {
	return common.GetObjects[Port](cable.GetClient(), cable.upstreamPorts)
}

// Update commits updates to this object's properties to the running system.
func (cable *Cable) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Cable)
	original.UnmarshalJSON(cable.rawData)

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

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(cable).Elem()

	return cable.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCable will get a Cable instance from the service.
func GetCable(c common.Client, uri string) (*Cable, error) {
	return common.GetObject[Cable](c, uri)
}

// ListReferencedCables gets the collection of Cable from
// a provided reference.
func ListReferencedCables(c common.Client, link string) ([]*Cable, error) {
	return common.GetCollectionObjects[Cable](c, link)
}
