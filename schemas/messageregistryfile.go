//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2017.1 - #MessageRegistryFile.v1_1_5.MessageRegistryFile

package schemas

import (
	"encoding/json"
)

// MessageRegistryFile shall represent the registry file locator for a Redfish
// implementation.
type MessageRegistryFile struct {
	Entity
	// Languages This property contains a set of RFC5646-conformant language codes.
	Languages []string
	// Location shall contain the location information for this registry file.
	Location []MessageRegistryFileLocation
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Registry shall contain the registry name and it major and minor versions, as
	// defined by the Redfish Specification. This registry can be any type of
	// registry, such as message registry, privilege registry, or attribute
	// registry.
	Registry string
}

// GetMessageRegistryFile will get a MessageRegistryFile instance from the service.
func GetMessageRegistryFile(c Client, uri string) (*MessageRegistryFile, error) {
	return GetObject[MessageRegistryFile](c, uri)
}

// ListReferencedMessageRegistryFiles gets the collection of MessageRegistryFile from
// a provided reference.
func ListReferencedMessageRegistryFiles(c Client, link string) ([]*MessageRegistryFile, error) {
	return GetCollectionObjects[MessageRegistryFile](c, link)
}

// MessageRegistryFileLocation shall contain the location information for a registry file.
type MessageRegistryFileLocation struct {
	// ArchiveFile shall contain the file name of the individual registry file
	// within the archive file specified by the 'ArchiveUri' property. The file
	// name shall conform to the Redfish Specification-specified syntax.
	ArchiveFile string
	// ArchiveURI shall contain a URI that is colocated with the Redfish service
	// that specifies the location of the registry file, which can be retrieved
	// using the Redfish protocol and authentication methods. This property shall
	// be used for only ZIP or other archive files. The 'ArchiveFile' property
	// shall contain the file name of the individual registry file within the
	// archive file.
	ArchiveURI string `json:"ArchiveUri"`
	// Language shall contain an RFC5646-conformant language code or 'default'.
	Language string
	// PublicationURI shall contain a URI not colocated with the Redfish service
	// that specifies the canonical location of the registry file. This property
	// shall be used for only individual registry files.
	PublicationURI string `json:"PublicationUri"`
	// URI shall contain a URI colocated with the Redfish service that specifies
	// the location of the registry file, which can be retrieved using the Redfish
	// protocol and authentication methods. This property shall be used for only
	// individual registry files. The file name portion of the URI shall conform to
	// Redfish Specification-specified syntax.
	URI string `json:"Uri"`
}
