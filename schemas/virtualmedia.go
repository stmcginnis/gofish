//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/VirtualMedia.v1_6_5.json
// 2022.3 - #VirtualMedia.v1_6_5.VirtualMedia

package schemas

import (
	"encoding/json"
	"errors"
)

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

type EjectPolicy string

const (
	// OnPowerOffEjectPolicy The virtual media ejection occurs during a system
	// power or reset event.
	OnPowerOffEjectPolicy EjectPolicy = "OnPowerOff"
	// SessionEjectPolicy The virtual media ejection occurs when a session is
	// terminated. The session might be outside the Redfish service.
	SessionEjectPolicy EjectPolicy = "Session"
	// TimedEjectPolicy The virtual media ejection occurs when a timer configured
	// by the 'EjectTimeout' property expires.
	TimedEjectPolicy EjectPolicy = "Timed"
	// AfterUseEjectPolicy The virtual media ejection occurs after the media is
	// used.
	AfterUseEjectPolicy EjectPolicy = "AfterUse"
	// PersistentEjectPolicy The virtual media mount information persists
	// indefinitely.
	PersistentEjectPolicy EjectPolicy = "Persistent"
)

type VirtualMediaType string

const (
	// CDVirtualMediaType is a CD-ROM format (ISO) image.
	CDVirtualMediaType VirtualMediaType = "CD"
	// FloppyVirtualMediaType is a floppy disk image.
	FloppyVirtualMediaType VirtualMediaType = "Floppy"
	// USBStickVirtualMediaType is an emulation of a USB storage device.
	USBStickVirtualMediaType VirtualMediaType = "USBStick"
	// DVDVirtualMediaType is a DVD-ROM format image.
	DVDVirtualMediaType VirtualMediaType = "DVD"
)

type TransferMethod string

const (
	// StreamTransferMethod Stream image file data from the source URI.
	StreamTransferMethod TransferMethod = "Stream"
	// UploadTransferMethod Upload the entire image file from the source URI to the
	// service.
	UploadTransferMethod TransferMethod = "Upload"
)

type VirtualMediaTransferProtocolType string

const (
	// CIFSVirtualMediaTransferProtocolType Common Internet File System (CIFS).
	CIFSVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "CIFS"
	// FTPVirtualMediaTransferProtocolType File Transfer Protocol (FTP).
	FTPVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "FTP"
	// SFTPVirtualMediaTransferProtocolType SSH File Transfer Protocol (SFTP).
	SFTPVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "SFTP"
	// HTTPVirtualMediaTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "HTTP"
	// HTTPSVirtualMediaTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "HTTPS"
	// NFSVirtualMediaTransferProtocolType Network File System (NFS).
	NFSVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "NFS"
	// SCPVirtualMediaTransferProtocolType Secure Copy Protocol (SCP).
	SCPVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "SCP"
	// TFTPVirtualMediaTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "TFTP"
	// OEMVirtualMediaTransferProtocolType is a manufacturer-defined protocol.
	OEMVirtualMediaTransferProtocolType VirtualMediaTransferProtocolType = "OEM"
)

// VirtualMedia shall represent a virtual media service for a Redfish
// implementation.
type VirtualMedia struct {
	Entity
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the server certificates for the
	// server referenced by the 'Image' property. If 'VerifyCertificate' is 'true',
	// services shall compare the certificates in this collection with the
	// certificate obtained during handshaking with the image server in order to
	// verify the identity of the image server prior to completing the remote media
	// connection. If the server cannot be verified, the service shall not complete
	// the remote media connection. If 'VerifyCertificate' is 'false', the service
	// shall not perform certificate verification with certificates in this
	// collection. Regardless of the contents of this collection, services may
	// perform additional verification based on other factors, such as the
	// configuration of the 'SecurityPolicy' resource.
	//
	// Version added: v1.4.0
	certificates string
	// ClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates
	// that are provided to the server referenced by the 'Image' property as part
	// of TLS handshaking.
	//
	// Version added: v1.5.0
	clientCertificates string
	// ConnectedVia shall contain the current connection method from a client to
	// the virtual media that this resource represents.
	ConnectedVia ConnectedVia
	// EjectPolicy shall contain the ejection policy for the virtual media.
	//
	// Version added: v1.6.0
	EjectPolicy EjectPolicy
	// EjectTimeout shall indicate the amount of time before virtual media is
	// automatically ejected when 'EjectPolicy' contains 'Timed'.
	//
	// Version added: v1.6.0
	EjectTimeout string
	// Image shall contain the URI of the media attached to the virtual media. This
	// value may specify an absolute URI to remote media or a relative URI to media
	// local to the implementation. A service may allow a relative URI to reference
	// a SoftwareInventory resource. The value 'null' shall indicate no image
	// connection. If the URI specified by this property is not accessible,
	// services should populate the 'Conditions' property in 'Status' with messages
	// that describe the condition.
	Image string
	// ImageName shall contain the name of the image.
	ImageName string
	// Inserted shall indicate whether the media is mounted and visible to the host
	// system.
	Inserted *bool `json:",omitempty"`
	// MediaTypes shall contain an array of the supported media types for this
	// connection.
	MediaTypes []VirtualMediaType
	// MediaType indicates which MediaTypes is used (non-standard)
	MediaType VirtualMediaType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password to access the URI specified by the
	// 'Image' property. The value shall be 'null' in responses.
	//
	// Version added: v1.3.0
	Password string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.4.0
	Status Status
	// TransferMethod shall describe how the image transfer occurs.
	//
	// Version added: v1.3.0
	TransferMethod TransferMethod
	// TransferProtocolType shall contain network protocol to use with the URI
	// specified by the 'Image' property.
	//
	// Version added: v1.3.0
	TransferProtocolType VirtualMediaTransferProtocolType
	// UserName shall contain the username to access the URI specified by the
	// 'Image' property.
	//
	// Version added: v1.3.0
	UserName string
	// VerifyCertificate shall indicate whether the service will verify the
	// certificate of the server referenced by the 'Image' property prior to
	// completing the remote media connection with the certificates found in the
	// collection referenced by the 'Certificates' property. If this property is
	// not supported by the service, it shall be assumed to be 'false'. This
	// property should default to 'false' in order to maintain compatibility with
	// older clients. Regardless of the value of this property, services may
	// perform additional verification based on other factors, such as the
	// configuration of the SecurityPolicy resource.
	//
	// Version added: v1.4.0
	VerifyCertificate bool
	// WriteProtected shall indicate whether the remote device media prevents
	// writing to that media.
	WriteProtected *bool `json:",omitempty"`

	// ejectMediaTarget is the URL to send EjectMedia requests.
	ejectMedia ActionTarget
	// SupportsMediaEject indicates if ejecting virtual media is supported.
	SupportsMediaEject bool

	// insertMediaTarget is the URL to send InsertMedia requests.
	insertMedia ActionTarget
	// SupportsMediaInsert indicates if inserting virtual media is supported.
	SupportsMediaInsert bool

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (v *VirtualMedia) UnmarshalJSON(b []byte) error {
	type temp VirtualMedia
	type vActions struct {
		EjectMedia  ActionTarget `json:"#VirtualMedia.EjectMedia"`
		InsertMedia ActionTarget `json:"#VirtualMedia.InsertMedia"`
	}
	var tmp struct {
		temp
		Actions            vActions
		Certificates       Link `json:"Certificates"`
		ClientCertificates Link `json:"ClientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VirtualMedia(tmp.temp)

	// Extract the links to other entities for later
	v.certificates = tmp.Certificates.String()
	v.clientCertificates = tmp.ClientCertificates.String()

	v.ejectMedia = tmp.Actions.EjectMedia
	v.insertMedia = tmp.Actions.InsertMedia

	v.SupportsMediaEject = v.ejectMedia.Target != ""
	v.SupportsMediaInsert = v.insertMedia.Target != ""

	// This is a read/write object, so we need to save the raw object data for later
	v.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *VirtualMedia) Update() error {
	readWriteFields := []string{
		"EjectPolicy",
		"EjectTimeout",
		"Image",
		"Inserted",
		"Password",
		"TransferMethod",
		"TransferProtocolType",
		"UserName",
		"VerifyCertificate",
		"WriteProtected",
	}

	return v.UpdateFromRawData(v, v.RawData, readWriteFields)
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c Client, uri string) (*VirtualMedia, error) {
	return GetObject[VirtualMedia](c, uri)
}

// ListReferencedVirtualMedias gets the collection of VirtualMedia from
// a provided reference.
func ListReferencedVirtualMedias(c Client, link string) ([]*VirtualMedia, error) {
	return GetCollectionObjects[VirtualMedia](c, link)
}

// EjectMediaActionInfo provides the ActionInfo, if supported, for an EjectMedia Action
func (v *VirtualMedia) EjectMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia EjectMedia ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// This action shall detach the remote media from the virtual media. At the
// completion of the operation, inserted shall be set to 'false' and the image
// name shall be cleared.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (v *VirtualMedia) EjectMedia() (*TaskMonitorInfo, error) {
	if !v.SupportsMediaEject {
		return nil, errors.New("redfish service does not support VirtualMedia.EjectMedia calls")
	}

	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(v.client,
		v.ejectMedia.Target, payload, v.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// VirtualMediaInsertMediaParameters holds the parameters for the InsertMedia action.
type VirtualMediaInsertMediaParameters struct {
	// Image shall contain the URI of the media to be attached to the virtual
	// media. This parameter may specify an absolute URI to remote media or a
	// relative URI to media local to the implementation. A service may allow a
	// relative URI to reference a 'SoftwareInventory' resource. Services may
	// reject this action if the URI is not accessible.
	Image string `json:"Image,omitempty"`
	// Inserted shall contain whether the image is treated as mounted and visible
	// to the host system upon completion of the action. If the client does not
	// provide this parameter, the service shall default this value to 'true'.
	Inserted *bool `json:"Inserted,omitempty"`
	// Password shall contain the password to access the URI specified by the
	// 'Image' parameter.
	Password *string `json:"Password,omitempty"`
	// TransferMethod shall contain the transfer method to use with the specified
	// image URI.
	TransferMethod *TransferMethod `json:"TransferMethod,omitempty"`
	// TransferProtocolType shall contain the network protocol to use with the URI
	// specified by the 'Image' parameter.
	TransferProtocolType *TransferProtocolType `json:"TransferProtocolType,omitempty"`
	// UserName shall contain the username to access the URI specified by the
	// 'Image' parameter.
	UserName *string `json:"UserName,omitempty"`
	// WriteProtected shall contain whether the remote media is treated as
	// write-protected. If the client does not provide this parameter, the service
	// shall default this value to 'true'.
	WriteProtected *bool `json:"WriteProtected,omitempty"`
	// MediaType is not part of the current schema, but some vendors require it.
	MediaType VirtualMediaType `json:",omitempty"`
}

// InsertMediaActionInfo provides the ActionInfo, if supported, for an InsertMedia Action
func (v *VirtualMedia) InsertMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia InsertMedia ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// This action shall attach remote media to virtual media. Service should
// reject the request if the 'Image' property in the resource does not contain
// 'null'. Users are expected to eject media prior to inserting new media.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (v *VirtualMedia) InsertMedia(params *VirtualMediaInsertMediaParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(v.client,
		v.insertMedia.Target, params, v.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Certificates gets the Certificates collection.
func (v *VirtualMedia) Certificates() ([]*Certificate, error) {
	if v.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](v.client, v.certificates)
}

// ClientCertificates gets the ClientCertificates collection.
func (v *VirtualMedia) ClientCertificates() ([]*Certificate, error) {
	if v.clientCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](v.client, v.clientCertificates)
}
