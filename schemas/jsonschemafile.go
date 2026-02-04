//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/JsonSchemaFile.v1_1_5.json
// 2017.1 - #JsonSchemaFile.v1_1_5.JsonSchemaFile

package schemas

import (
	"encoding/json"
)

// JSONSchemaFile shall represent the schema file locator resource for a Redfish
// implementation.
type JSONSchemaFile struct {
	Entity
	// Languages This property contains a set of RFC5646-conformant language codes.
	Languages []string
	// Location shall contain the location information for this schema file.
	Location SchemaFileLocation
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Schema shall contain the '@odata.type' property value for that schema and
	// shall conform to the Redfish Specification-specified syntax for the 'Type'
	// property.
	Schema string
}

// GetJSONSchemaFile will get a JSONSchemaFile instance from the service.
func GetJSONSchemaFile(c Client, uri string) (*JSONSchemaFile, error) {
	return GetObject[JSONSchemaFile](c, uri)
}

// ListReferencedJSONSchemaFiles gets the collection of JSONSchemaFile from
// a provided reference.
func ListReferencedJSONSchemaFiles(c Client, link string) ([]*JSONSchemaFile, error) {
	return GetCollectionObjects[JSONSchemaFile](c, link)
}

// SchemaFileLocation shall describe location information for a schema file.
type SchemaFileLocation struct {
	// ArchiveFile shall contain the file name of the individual schema file within
	// the archive file that the 'ArchiveUri' property specifies. The file name
	// shall conform to the Redfish Specification-described format.
	ArchiveFile string
	// ArchiveURI shall contain a URI colocated with the Redfish service that
	// specifies the location of the schema file, which can be retrieved using the
	// Redfish protocol and authentication methods. This property shall be used for
	// only archive files, in zip or other formats. The 'ArchiveFile' value shall
	// be the individual schema file name within the archive file.
	ArchiveURI string `json:"ArchiveUri"`
	// Language shall contain an RFC5646-conformant language code or the 'default'
	// string.
	Language string
	// PublicationURI shall contain a URI not colocated with the Redfish service
	// that specifies the canonical location of the schema file. This property
	// shall be used for only individual schema files.
	PublicationURI string `json:"PublicationUri"`
	// URI shall contain a URI colocated with the Redfish service that specifies
	// the location of the schema file, which can be retrieved using the Redfish
	// protocol and authentication methods. This property shall be used for only
	// individual schema files. The file name portion of the URI shall conform to
	// the format specified in the Redfish Specification.
	URI string `json:"Uri"`
}
