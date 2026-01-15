//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2017.1 - #JsonSchemaFile.v1_1_5.JsonSchemaFile

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// JSONSchemaFile shall represent the schema file locator resource for a Redfish
// implementation.
type JSONSchemaFile struct {
	common.Entity
	// Languages This property contains a set of RFC5646-conformant language codes.
	Languages []string
	// Location shall contain the location information for this schema file.
	Location common.Location
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Schema shall contain the '@odata.type' property value for that schema and
	// shall conform to the Redfish Specification-specified syntax for the 'Type'
	// property.
	Schema string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a JSONSchemaFile object from the raw JSON.
func (j *JSONSchemaFile) UnmarshalJSON(b []byte) error {
	type temp JSONSchemaFile
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*j = JSONSchemaFile(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	j.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (j *JSONSchemaFile) Update() error {
	readWriteFields := []string{
		"Location",
	}

	return j.UpdateFromRawData(j, j.rawData, readWriteFields)
}

// GetJSONSchemaFile will get a JSONSchemaFile instance from the service.
func GetJSONSchemaFile(c common.Client, uri string) (*JSONSchemaFile, error) {
	return common.GetObject[JSONSchemaFile](c, uri)
}

// ListReferencedJSONSchemaFiles gets the collection of JSONSchemaFile from
// a provided reference.
func ListReferencedJSONSchemaFiles(c common.Client, link string) ([]*JSONSchemaFile, error) {
	return common.GetCollectionObjects[JSONSchemaFile](c, link)
}

// Location shall describe location information for a schema file.
type Location struct {
	// ArchiveFile shall contain the file name of the individual schema file within
	// the archive file that the 'ArchiveUri' property specifies. The file name
	// shall conform to the Redfish Specification-described format.
	ArchiveFile string
	// ArchiveUri shall contain a URI colocated with the Redfish service that
	// specifies the location of the schema file, which can be retrieved using the
	// Redfish protocol and authentication methods. This property shall be used for
	// only archive files, in zip or other formats. The 'ArchiveFile' value shall
	// be the individual schema file name within the archive file.
	ArchiveURI string
	// Language shall contain an RFC5646-conformant language code or the 'default'
	// string.
	Language string
	// PublicationUri shall contain a URI not colocated with the Redfish service
	// that specifies the canonical location of the schema file. This property
	// shall be used for only individual schema files.
	PublicationURI string
	// Uri shall contain a URI colocated with the Redfish service that specifies
	// the location of the schema file, which can be retrieved using the Redfish
	// protocol and authentication methods. This property shall be used for only
	// individual schema files. The file name portion of the URI shall conform to
	// the format specified in the Redfish Specification.
	URI string
}
