//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stmcginnis/gofish/common"
)

// ConnectedVia are the ways the media may be connected.
type ConnectedVia string

const (
	// NotConnectedConnectedVia No current connection.
	NotConnectedConnectedVia ConnectedVia = "NotConnected"
	// URIConnectedVia Connected to a URI location.
	URIConnectedVia ConnectedVia = "URI"
	// AppletConnectedVia Connected to a client application.
	AppletConnectedVia ConnectedVia = "Applet"
	// OemConnectedVia Connected through an OEM-defined method.
	OemConnectedVia ConnectedVia = "Oem"
)

// EjectPolicy is the ejection policy for virtual media.
type EjectPolicy string

const (
	// OnPowerOffEjectPolicy The virtual media ejection occurs during a system power or reset event.
	OnPowerOffEjectPolicy EjectPolicy = "OnPowerOff"
	// SessionEjectPolicy The virtual media ejection occurs when a session is terminated.
	SessionEjectPolicy EjectPolicy = "Session"
	// TimedEjectPolicy The virtual media ejection occurs when a timer configured by the EjectTimeout property expires.
	TimedEjectPolicy EjectPolicy = "Timed"
	// AfterUseEjectPolicy The virtual media ejection occurs after the media is used.
	AfterUseEjectPolicy EjectPolicy = "AfterUse"
	// PersistentEjectPolicy The virtual media mount information persists indefinitely.
	PersistentEjectPolicy EjectPolicy = "Persistent"
)

// VirtualMediaType is the type of media.
type VirtualMediaType string

const (
	// CDMediaType A CD-ROM format (ISO) image.
	CDMediaType VirtualMediaType = "CD"
	// FloppyMediaType A floppy disk image.
	FloppyMediaType VirtualMediaType = "Floppy"
	// USBStickMediaType An emulation of a USB storage device.
	USBStickMediaType VirtualMediaType = "USBStick"
	// DVDMediaType A DVD-ROM format image.
	DVDMediaType VirtualMediaType = "DVD"
)

// TransferMethod is how the data is transferred.
type TransferMethod string

const (
	// StreamTransferMethod Stream image file data from the source URI.
	StreamTransferMethod TransferMethod = "Stream"
	// UploadTransferMethod Upload the entire image file from the source URI to the service.
	UploadTransferMethod TransferMethod = "Upload"
)

// TransferProtocolType is the protocol used to transfer.
type TransferProtocolType string

const (
	// CIFSTransferProtocolType Common Internet File System (CIFS).
	CIFSTransferProtocolType TransferProtocolType = "CIFS"
	// FTPTransferProtocolType File Transfer Protocol (FTP).
	FTPTransferProtocolType TransferProtocolType = "FTP"
	// SFTPTransferProtocolType Secure File Transfer Protocol (SFTP).
	SFTPTransferProtocolType TransferProtocolType = "SFTP"
	// HTTPTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPTransferProtocolType TransferProtocolType = "HTTP"
	// HTTPSTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSTransferProtocolType TransferProtocolType = "HTTPS"
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType A manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
)

// VirtualMedia shall represent a virtual media service for a Redfish implementation.
type VirtualMedia struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Certificates shall contain a link to a resource collection of type CertificateCollection.
	certificates []string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection.
	clientCertificates []string
	// ConnectedVia shall contain the current connection method.
	ConnectedVia ConnectedVia
	// Description provides a description of this resource.
	Description string
	// EjectPolicy shall contain the ejection policy for the virtual media.
	EjectPolicy *EjectPolicy
	// EjectTimeout shall indicate the amount of time before virtual media is automatically ejected.
	EjectTimeout *string
	// Image shall contain an URI. A null value indicates no image connection.
	Image *string
	// ImageName shall contain the name of the image.
	ImageName string
	// Inserted shall indicate whether media is present in the virtual media device.
	Inserted *bool
	// MediaType is the connection media type used.
	MediaType *VirtualMediaType
	// MediaTypes shall be the supported media types for this connection.
	MediaTypes []VirtualMediaType
	// Password shall represent the password to access the Image parameter-specified URI.
	Password *string
	// TransferMethod shall describe how the image transfer occurs.
	TransferMethod *TransferMethod
	// TransferProtocolType shall represent the network protocol to use.
	TransferProtocolType *TransferProtocolType
	// UserName shall represent the user name to access the Image parameter-specified URI.
	UserName *string
	// VerifyCertificate shall indicate whether the service will verify the certificate.
	VerifyCertificate *bool
	// WriteProtected shall indicate whether the remote device media prevents writing.
	WriteProtected *bool
	// Status contains the status and health of the resource.
	Status *common.Status
	// Oem contains the OEM extensions.
	Oem json.RawMessage
	// rawData holds the original serialized JSON.
	rawData []byte

	// ejectMediaTarget is the URL to send EjectMedia requests.
	ejectMedia common.ActionTarget
	// SupportsMediaEject indicates if ejecting virtual media is supported.
	SupportsMediaEject bool

	// insertMediaTarget is the URL to send InsertMedia requests.
	insertMedia common.ActionTarget
	// SupportsMediaInsert indicates if inserting virtual media is supported.
	SupportsMediaInsert bool
}

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (virtualmedia *VirtualMedia) UnmarshalJSON(b []byte) error {
	type tVirtualMedia VirtualMedia
	var t struct {
		tVirtualMedia
		Actions struct {
			EjectMedia  common.ActionTarget `json:"#VirtualMedia.EjectMedia"`
			InsertMedia common.ActionTarget `json:"#VirtualMedia.InsertMedia"`
			Oem         json.RawMessage     `json:"Oem"`
		}
		Certificates       common.LinksCollection
		ClientCertificates common.LinksCollection
		Status             *common.Status
		Oem                json.RawMessage
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualmedia = VirtualMedia(t.tVirtualMedia)
	virtualmedia.Status = t.Status
	virtualmedia.Oem = t.Oem

	// Extract the links to other entities for later
	virtualmedia.certificates = t.Certificates.ToStrings()
	virtualmedia.clientCertificates = t.ClientCertificates.ToStrings()

	virtualmedia.ejectMedia = t.Actions.EjectMedia
	virtualmedia.insertMedia = t.Actions.InsertMedia

	virtualmedia.SupportsMediaEject = virtualmedia.ejectMedia.Target != ""
	virtualmedia.SupportsMediaInsert = virtualmedia.insertMedia.Target != ""

	// This is a read/write object, so we need to save the raw object data for later
	virtualmedia.rawData = b

	return nil
}

// Certificates gets the server certificates for the server referenced by the Image property.
func (virtualmedia *VirtualMedia) Certificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](virtualmedia.GetClient(), virtualmedia.certificates)
}

// ClientCertificates gets the client identity certificates.
func (virtualmedia *VirtualMedia) ClientCertificates() ([]*Certificate, error) {
	return common.GetObjects[Certificate](virtualmedia.GetClient(), virtualmedia.clientCertificates)
}

// Update commits updates to this object's properties to the running system.
func (virtualmedia *VirtualMedia) Update() error {
	original := new(VirtualMedia)
	if err := original.UnmarshalJSON(virtualmedia.rawData); err != nil {
		return err
	}

	readWriteFields := []string{
		"EjectPolicy",
		"EjectTimeout",
		"Image",
		"Inserted",
		"MediaType",
		"Password",
		"TransferMethod",
		"TransferProtocolType",
		"UserName",
		"VerifyCertificate",
		"WriteProtected",
	}

	return virtualmedia.UpdateFromRawData(virtualmedia, virtualmedia.rawData, readWriteFields)
}

// EjectMediaActionInfo provides the ActionInfo, if supported, for an EjectMedia Action
func (virtualmedia *VirtualMedia) EjectMediaActionInfo() (*ActionInfo, error) {
	if virtualmedia.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia EjectMedia ActionInfo not supported by this service")
	}

	return common.GetObject[ActionInfo](virtualmedia.GetClient(), virtualmedia.insertMedia.ActionInfoTarget)
}

// EjectMedia sends a request to eject the media.
func (virtualmedia *VirtualMedia) EjectMedia() error {
	if !virtualmedia.SupportsMediaEject {
		return errors.New("redfish service does not support VirtualMedia.EjectMedia calls")
	}

	return virtualmedia.Post(virtualmedia.ejectMedia.Target, nil)
}

// VirtualMediaConfig is used to pass config data when inserting media.
type VirtualMediaConfig struct {
	Image                string                `json:"Image"`
	Inserted             *bool                 `json:"Inserted,omitempty"`
	MediaType            *VirtualMediaType     `json:"MediaType,omitempty"`
	Password             *string               `json:"Password,omitempty"`
	TransferMethod       *TransferMethod       `json:"TransferMethod,omitempty"`
	TransferProtocolType *TransferProtocolType `json:"TransferProtocolType,omitempty"`
	UserName             *string               `json:"UserName,omitempty"`
	WriteProtected       *bool                 `json:"WriteProtected,omitempty"`
}

// InsertMediaActionInfo provides the ActionInfo, if supported, for an InsertMedia Action
func (virtualmedia *VirtualMedia) InsertMediaActionInfo() (*ActionInfo, error) {
	if virtualmedia.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia InsertMedia ActionInfo not supported by this service")
	}

	return common.GetObject[ActionInfo](virtualmedia.GetClient(), virtualmedia.insertMedia.ActionInfoTarget)
}

// InsertMedia sends a request to insert virtual media.
func (virtualmedia *VirtualMedia) InsertMedia(image string, inserted, writeProtected bool) (*Task, error) {
	return virtualmedia.InsertMediaConfig(VirtualMediaConfig{
		Image:          image,
		Inserted:       &inserted,
		WriteProtected: &writeProtected,
	})
}

// InsertMediaConfig sends a request to insert virtual media using the VirtualMediaConfig struct.
func (virtualmedia *VirtualMedia) InsertMediaConfig(config VirtualMediaConfig) (*Task, error) {
	if !virtualmedia.SupportsMediaInsert {
		return nil, errors.New("redfish service does not support VirtualMedia.InsertMedia calls")
	}

	resp, err := virtualmedia.PostWithResponse(virtualmedia.insertMedia.Target, config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return nil, nil
	}

	if location := resp.Header["Location"]; len(location) > 0 {
		return GetTask(virtualmedia.GetClient(), location[0])
	}

	return nil, nil
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c common.Client, uri string) (*VirtualMedia, error) {
	return common.GetObject[VirtualMedia](c, uri)
}

// ListReferencedVirtualMedias gets the collection of VirtualMedia from a provided reference.
func ListReferencedVirtualMedias(c common.Client, link string) ([]*VirtualMedia, error) {
	return common.GetCollectionObjects[VirtualMedia](c, link)
}
