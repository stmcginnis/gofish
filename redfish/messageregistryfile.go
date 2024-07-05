//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// MessageRegistryFileLocation is a location
// information for the Message Registry file.
type MessageRegistryFileLocation struct {
	// ArchiveFile shall contain the file name of the individual registry file within the archive file specified by the
	// ArchiveUri property. The file name shall conform to the Redfish Specification-specified syntax.
	ArchiveFile string
	// ArchiveURI shall contain a URI that is colocated with the Redfish service that specifies the location of the
	// registry file, which can be retrieved using the Redfish protocol and authentication methods. This property shall
	// be used for only ZIP or other archive files. The ArchiveFile property shall contain the file name of the
	// individual registry file within the archive file.
	ArchiveURI string
	// Language shall contain an RFC5646-conformant language code or 'default'.
	Language string
	// PublicationURI shall contain a URI not colocated with the Redfish service that specifies the canonical location
	// of the registry file. This property shall be used for only individual registry files.
	PublicationURI string
	// Uri shall contain a URI colocated with the Redfish service that specifies the location of the registry file,
	// which can be retrieved using the Redfish protocol and authentication methods. This property shall be used for
	// only individual registry files. The file name portion of the URI shall conform to Redfish Specification-
	// specified syntax.
	URI string
}

// MessageRegistryFile describes the Message Registry file locator Resource.
type MessageRegistryFile struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Languages is the RFC5646-conformant language codes for the
	// available Message Registries.
	Languages []string
	// Location is the location information for this Message Registry file.
	Location []MessageRegistryFileLocation
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Registry shall contain the Message Registry name and it major and
	// minor versions, as defined by the Redfish Specification.
	Registry string
}

// GetMessageRegistryFile will get a MessageRegistryFile
// instance from the Redfish service.
func GetMessageRegistryFile(c common.Client, uri string) (*MessageRegistryFile, error) {
	return common.GetObject[MessageRegistryFile](c, uri)
}

// ListReferencedMessageRegistryFiles gets the collection of MessageRegistryFile.
func ListReferencedMessageRegistryFiles(c common.Client, link string) ([]*MessageRegistryFile, error) {
	return common.GetCollectionObjects[MessageRegistryFile](c, link)
}
