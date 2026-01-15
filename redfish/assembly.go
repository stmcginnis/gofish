//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #Assembly.v1_6_0.Assembly

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// Assembly shall represent an assembly for a Redfish implementation. Assembly
// information contains details about a device, such as part number, serial
// number, manufacturer, and production date. It also provides access to the
// original data for the assembly.
type Assembly struct {
	common.Entity
	// Assemblies shall define assembly records for a Redfish implementation.
	Assemblies []AssemblyData
	// Assemblies@odata.count
	AssembliesCount int `json:"Assemblies@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Assembly object from the raw JSON.
func (a *Assembly) UnmarshalJSON(b []byte) error {
	type temp Assembly
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = Assembly(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *Assembly) Update() error {
	readWriteFields := []string{
		"Assemblies",
		"Assemblies@odata.count",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetAssembly will get a Assembly instance from the service.
func GetAssembly(c common.Client, uri string) (*Assembly, error) {
	return common.GetObject[Assembly](c, uri)
}

// ListReferencedAssemblys gets the collection of Assembly from
// a provided reference.
func ListReferencedAssemblys(c common.Client, link string) ([]*Assembly, error) {
	return common.GetCollectionObjects[Assembly](c, link)
}

// AssemblyData represents the AssemblyData type.
type AssemblyData struct {
	// BinaryDataURI shall contain the URI at which to access an image of the
	// assembly information, using the Redfish protocol and authentication methods.
	// The service provides this URI for the download of the OEM-specific binary
	// image of the assembly data. An HTTP 'GET' from this URI shall return a
	// response payload of MIME type 'application/octet-stream'. If the service
	// supports it, an HTTP 'PUT' to this URI shall replace the binary image of the
	// assembly.
	BinaryDataURI string
	// Description provides a description of this resource.
	Description string
	// EngineeringChangeLevel shall contain the engineering change level or
	// revision of the assembly.
	EngineeringChangeLevel string
	// ISOCountryCodeOfOrigin shall contain the ISO 3166-1-defined alpha-2 or
	// alpha-3 country code that reflects the manufacturing country of origin.
	//
	// Version added: v1.5.0
	ISOCountryCodeOfOrigin string
	// Location shall contain the location information of the associated assembly.
	//
	// Version added: v1.3.0
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	//
	// Version added: v1.3.0
	LocationIndicatorActive bool
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the name by which the manufacturer generally refers to
	// the assembly.
	Model string
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number of the assembly.
	PartNumber string
	// PhysicalContext shall contain a description of the physical context for the
	// assembly data.
	//
	// Version added: v1.2.0
	PhysicalContext PhysicalContext
	// Producer shall contain the name of the company that produced or manufactured
	// the assembly. This value shall be equal to the 'Manufacturer' field value in
	// a PLDM FRU structure, if applicable, for the assembly.
	Producer string
	// ProductionDate shall contain the date of production or manufacture for the
	// assembly. The time of day portion of the property shall be '00:00:00Z', if
	// the time of day is unknown.
	ProductionDate string
	// ReadyToRemove shall indicate whether the assembly is ready for removal.
	// Setting the value to 'true' shall cause the service to perform appropriate
	// actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	//
	// Version added: v1.6.0
	ReadyToRemove bool
	// Replaceable shall indicate whether the component associated this assembly
	// can be independently replaced as allowed by the vendor's replacement policy.
	// A value of 'false' indicates the component needs to be replaced by policy as
	// part of another component. If the 'LocationType' property of this assembly
	// contains 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.4.0
	Replaceable bool
	// SKU shall contain the SKU of the assembly.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the assembly.
	//
	// Version added: v1.2.0
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the assembly.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.1.0
	Status common.Status
	// Vendor shall contain the name of the company that provides the final product
	// that includes this assembly. This value shall be equal to the 'Vendor' field
	// value in a PLDM FRU structure, if applicable, for the assembly.
	Vendor string
	// Version shall contain the hardware version of the assembly as determined by
	// the vendor or supplier.
	Version string
}
