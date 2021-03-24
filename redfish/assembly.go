//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Assembly is used to represent an assembly information resource for a
// Redfish implementation.
type Assembly struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assemblies shall be the definition for assembly records for a Redfish
	// implementation.
	Assemblies []AssemblyData
	// Assemblies@odata.count is
	AssembliesCount int `json:"Assemblies@odata.count"`
	// Description provides a description of this resource.
	Description string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Assembly object from the raw JSON.
func (assembly *Assembly) UnmarshalJSON(b []byte) error {
	type temp Assembly
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*assembly = Assembly(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	assembly.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (assembly *Assembly) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Assembly)
	err := original.UnmarshalJSON(assembly.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Assemblies",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(assembly).Elem()

	return assembly.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAssembly will get a Assembly instance from the service.
func GetAssembly(c common.Client, uri string) (*Assembly, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assembly Assembly
	err = json.NewDecoder(resp.Body).Decode(&assembly)
	if err != nil {
		return nil, err
	}

	assembly.SetClient(c)
	return &assembly, nil
}

// ListReferencedAssemblys gets the collection of Assembly from
// a provided reference.
func ListReferencedAssemblys(c common.Client, link string) ([]*Assembly, error) {
	var result []*Assembly
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, assemblyLink := range links.ItemLinks {
		assembly, err := GetAssembly(c, assemblyLink)
		if err != nil {
			return result, err
		}
		result = append(result, assembly)
	}

	return result, nil
}

// AssemblyData is information about an assembly.
type AssemblyData struct {
	// BinaryDataURI shall be a URI at which the Service provides for the
	// download of the OEM-specific binary image of the assembly data. An HTTP
	// GET from this URI shall return a response payload of MIME time
	// application/octet-stream. An HTTP PUT to this URI, if supported by the
	// Service, shall replace the binary image of the assembly.
	BinaryDataURI string
	// Description provides a description of this resource.
	Description string
	// EngineeringChangeLevel shall be the Engineering Change Level (ECL) or
	// revision of the assembly.
	EngineeringChangeLevel string
	// MemberID shall uniquely identify the member within the collection.
	MemberID string
	// Model shall be the name by which the manufacturer generally refers to the
	// assembly.
	Model string
	// Name provides the name of the resource.
	Name string
	// PartNumber shall be the part number of the assembly.
	PartNumber string
	// PhysicalContext shall be a description of the physical context for this
	// assembly data.
	PhysicalContext string
	// Producer shall be the name of the company which supplied or manufactured
	// this assembly. This value shall be equal to the 'Manufacturer' field in a
	// PLDM FRU structure, if applicable, for this assembly.
	Producer string
	// ProductionDate shall be the date of production or manufacture for this
	// assembly. The time of day portion of the property shall be '00:00:00Z' if
	// the time of day is unknown.
	ProductionDate string
	// SKU shall be the name of the assembly.
	SKU string
	// SerialNumber is used to identify the assembly.
	SerialNumber string
	// SparePartNumber shall be the name of the assembly.
	SparePartNumber string
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status
	// Vendor shall be the name of the company which provides the final product
	// that includes this assembly. This value shall be equal to the 'Vendor'
	// field in a PLDM FRU structure, if applicable, for this assembly.
	Vendor string
	// Version shall be the version of the assembly as determined by the vendor
	// or supplier.
	Version string
}
