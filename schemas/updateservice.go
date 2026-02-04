//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/UpdateService.v1_17_0.json
// 2025.3 - #UpdateService.v1_17_0.UpdateService

package schemas

import (
	"encoding/json"
)

type UpdateServiceApplyTime string

const (
	// ImmediateUpdateServiceApplyTime shall indicate the 'HttpPushUri'-provided software is
	// applied immediately.
	ImmediateUpdateServiceApplyTime UpdateServiceApplyTime = "Immediate"
	// OnResetUpdateServiceApplyTime shall indicate the 'HttpPushUri'-provided software is
	// applied when the system or service is reset.
	OnResetUpdateServiceApplyTime UpdateServiceApplyTime = "OnReset"
	// AtMaintenanceWindowStartUpdateServiceApplyTime shall indicate the 'HttpPushUri'-provided
	// software is applied during the maintenance window specified by the
	// 'MaintenanceWindowStartTime' and 'MaintenanceWindowDurationInSeconds'
	// properties. A service may perform resets during this maintenance window.
	AtMaintenanceWindowStartUpdateServiceApplyTime UpdateServiceApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetUpdateServiceApplyTime shall indicate the
	// 'HttpPushUri'-provided software is applied during the maintenance window
	// specified by the 'MaintenanceWindowStartTime' and
	// 'MaintenanceWindowDurationInSeconds' properties, and if a reset occurs
	// within the maintenance window.
	InMaintenanceWindowOnResetUpdateServiceApplyTime UpdateServiceApplyTime = "InMaintenanceWindowOnReset"
	// OnStartUpdateRequestUpdateServiceApplyTime shall indicate the 'HttpPushUri'-provided
	// software is applied when the 'StartUpdate' action of the update service is
	// invoked.
	OnStartUpdateRequestUpdateServiceApplyTime UpdateServiceApplyTime = "OnStartUpdateRequest"
	// OnTargetResetUpdateServiceApplyTime shall indicate the 'HttpPushUri'-provided software is
	// applied when the target is reset.
	OnTargetResetUpdateServiceApplyTime UpdateServiceApplyTime = "OnTargetReset"
)

type SupportedUpdateImageFormatType string

const (
	// PLDMv10SupportedUpdateImageFormatType shall indicate an image that conforms
	// to the v1.0 image format as defined in DMTF DSP0267.
	PLDMv10SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_0"
	// PLDMv11SupportedUpdateImageFormatType shall indicate an image that conforms
	// to the v1.1 image format as defined in DMTF DSP0267.
	PLDMv11SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_1"
	// PLDMv12SupportedUpdateImageFormatType shall indicate an image that conforms
	// to the v1.2 image format as defined in DMTF DSP0267.
	PLDMv12SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_2"
	// PLDMv13SupportedUpdateImageFormatType shall indicate an image that conforms
	// to the v1.3 image format as defined in DMTF DSP0267.
	PLDMv13SupportedUpdateImageFormatType SupportedUpdateImageFormatType = "PLDMv1_3"
	// UEFICapsuleSupportedUpdateImageFormatType shall indicate an image that
	// conforms to the UEFI capsule format as defined in the UEFI Specification.
	UEFICapsuleSupportedUpdateImageFormatType SupportedUpdateImageFormatType = "UEFICapsule"
	// VendorDefinedSupportedUpdateImageFormatType shall indicate a vendor-defined
	// format.
	VendorDefinedSupportedUpdateImageFormatType SupportedUpdateImageFormatType = "VendorDefined"
)

type TransferProtocolType string

const (
	// CIFSTransferProtocolType Common Internet File System (CIFS).
	CIFSTransferProtocolType TransferProtocolType = "CIFS"
	// FTPTransferProtocolType File Transfer Protocol (FTP).
	FTPTransferProtocolType TransferProtocolType = "FTP"
	// SFTPTransferProtocolType SSH File Transfer Protocol (SFTP).
	SFTPTransferProtocolType TransferProtocolType = "SFTP"
	// HTTPTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPTransferProtocolType TransferProtocolType = "HTTP"
	// HTTPSTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSTransferProtocolType TransferProtocolType = "HTTPS"
	// NSFTransferProtocolType Network File System (NFS).
	NSFTransferProtocolType TransferProtocolType = "NSF"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType is a manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
)

// UpdateService shall represent an update service and the properties that
// affect the service itself for a Redfish implementation.
type UpdateService struct {
	Entity
	// ClientCertificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that represents the client identity certificates
	// that are provided to the server referenced by the 'ImageURI' parameter in
	// 'SimpleUpdate' as part of TLS handshaking.
	//
	// Version added: v1.10.0
	clientCertificates string
	// FirmwareInventory shall contain a link to a resource collection of type
	// 'SoftwareInventoryCollection'. The resource collection should contain the
	// set of software components generally referred to as platform firmware or
	// that does not execute within a host operating system. Software in this
	// collection is generally updated using platform-specific methods or
	// utilities.
	firmwareInventory string
	// HTTPPushURI shall contain a URI at which the update service supports an HTTP
	// or HTTPS 'POST' of a software image for the purpose of installing software
	// contained within the image. Access to this URI shall require the same
	// privilege as access to the update service. If the service requires the
	// 'Content-Length' header for 'POST' requests to this URI, the service should
	// return HTTP '411 Length Required' status code if the client does not include
	// this header in the 'POST' request. The value of this property should not
	// contain a URI of a Redfish resource. See the 'Redfish-defined URIs and
	// relative reference rules' clause in the Redfish Specification.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.15.0
	// This property has been deprecated in favor of the 'MultipartHttpPushUri'
	// property.
	HTTPPushURI string `json:"HttpPushUri"`
	// HTTPPushURIOptions shall contain options and requirements of the service for
	// 'HttpPushUri'-provided software updates.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.15.0
	// This property has been deprecated in favor of the update parameters used
	// with 'MultipartHttpPushUri'-provided software updates.
	HTTPPushURIOptions HTTPPushURIOptions `json:"HttpPushUriOptions"`
	// HTTPPushURIOptionsBusy shall indicate whether a client uses the
	// 'HttpPushUriOptions' properties for software updates. When a client uses any
	// 'HttpPushUriOptions' properties for software updates, it should set this
	// property to 'true'. When a client no longer uses 'HttpPushUriOptions'
	// properties for software updates, it should set this property to 'false'.
	// This property can provide multiple clients a way to negotiate ownership of
	// 'HttpPushUriOptions' properties. Clients can use this property to determine
	// whether another client uses 'HttpPushUriOptions' properties for software
	// updates. This property has no functional requirements for the service.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.15.0
	// This property has been deprecated in favor of
	// 'MultipartHttpPushUri'-provided software updates.
	HTTPPushURIOptionsBusy bool `json:"HttpPushUriOptionsBusy"`
	// HTTPPushURITargets shall contain zero or more URIs that indicate where to
	// apply the update image when using the URI specified by the 'HttpPushUri'
	// property to push a software image. These targets should correspond to
	// software inventory instances or their related items. If this property is not
	// present or contains no targets, the service shall apply the software image
	// to all applicable targets, as determined by the service. If the target
	// specifies a device resource, the software image file shall be applied to the
	// specified device. If the target specifies a resource collection, the
	// software image shall be applied to each applicable member of the specified
	// collection. If the target resource specifies an 'Aggregate' resource, the
	// software image file shall be applied to each applicable element of the
	// specified aggregate. If the target resource specifies a 'ComputerSystem'
	// resource, the software image file shall be applied to the applicable
	// components within the specified computer system.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.15.0
	// This property has been deprecated in favor of the update parameters used
	// with 'MultipartHttpPushUri'-provided software updates.
	HTTPPushURITargets []string `json:"HttpPushUriTargets"`
	// HTTPPushURITargetsBusy shall indicate whether any client has reserved the
	// 'HttpPushUriTargets' property for firmware updates. A client should set this
	// property to 'true' when it uses 'HttpPushUriTargets' for firmware updates. A
	// client should set it to 'false' when it no longer uses 'HttpPushUriTargets'
	// for updates. The property can provide multiple clients a way to negotiate
	// ownership of 'HttpPushUriTargets' and helps clients determine whether
	// another client is using 'HttpPushUriTargets' to make firmware updates. This
	// property has no functional requirements for the service.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.15.0
	// This property has been deprecated in favor of
	// 'MultipartHttpPushUri'-provided software updates.
	HTTPPushURITargetsBusy bool `json:"HttpPushUriTargetsBusy"`
	// LocalImageStore shall contain a link to a resource collection of type
	// 'SoftwareInventoryCollection'. The resource collection shall contain the set
	// of images that are stored locally by the service. This may include operating
	// system images, drivers, or other software components. These images may be
	// referenced by other features of the service, such as virtual media.
	//
	// Version added: v1.17.0
	localImageStore string
	// LocalImageStoreAvailableCapacityBytes shall contain the number of bytes
	// currently available in the local image store.
	//
	// Version added: v1.17.0
	LocalImageStoreAvailableCapacityBytes *int `json:",omitempty"`
	// LocalImageStoreTotalCapacityBytes shall contain the total number of bytes
	// available to the local image store.
	//
	// Version added: v1.17.0
	LocalImageStoreTotalCapacityBytes *int `json:",omitempty"`
	// MaxImageSizeBytes shall indicate the maximum size of the software update
	// image that clients can send to this update service.
	//
	// Version added: v1.5.0
	MaxImageSizeBytes *int `json:",omitempty"`
	// MultipartHTTPPushURI shall contain a URI used to perform a Redfish
	// Specification-defined multipart HTTP or HTTPS 'POST' of a software image for
	// the purpose of installing software contained within the image. The value of
	// this property should not contain a URI of a Redfish resource. See the
	// 'Redfish-defined URIs and relative reference rules' clause in the Redfish
	// Specification.
	//
	// Version added: v1.6.0
	MultipartHTTPPushURI string `json:"MultipartHttpPushUri"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PublicIdentitySSHKey shall contain a link to a resource of type 'Key' that
	// represents the public key that is used with the 'SimpleUpdate' action for
	// the key-based authentication. This property shall not be present if a
	// key-pair is not available.
	//
	// Version added: v1.13.0
	publicIdentitySSHKey string
	// RemoteServerCertificates shall contain a link to a resource collection of
	// type 'CertificateCollection' that represents the server certificates for the
	// server referenced by the 'ImageURI' parameter in 'SimpleUpdate'. If
	// 'VerifyRemoteServerCertificate' is 'true', services shall compare the
	// certificates in this collection with the certificate obtained during
	// handshaking with the image server in order to verify the identity of the
	// image server prior to transferring the image. If the server cannot be
	// verified, the service shall not send the transfer request. If
	// 'VerifyRemoteServerCertificate' is 'false', the service shall not perform
	// certificate verification with certificates in this collection. Regardless of
	// the contents of this collection, services may perform additional
	// verification based on other factors, such as the configuration of the
	// 'SecurityPolicy' resource.
	//
	// Version added: v1.9.0
	remoteServerCertificates string
	// RemoteServerSSHKeys shall contain a link to a resource collection of type
	// 'KeyCollection' that represents the server SSH keys for the server
	// referenced by the 'ImageURI' Parameter in 'SimpleUpdate'. If
	// 'VerifyRemoteServerSSHKey' is 'true', services shall compare the keys in
	// this collection with the key obtained during handshaking with the image
	// server in order to verify the identity of the image server prior to
	// transferring the image. If the server cannot be verified, the service shall
	// not send the transfer request. If 'VerifyRemoteServerSSHKey' is 'false', the
	// service shall not perform key verification with keys in this collection.
	//
	// Version added: v1.12.0
	remoteServerSSHKeys string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// SoftwareInventory shall contain a link to a resource collection of type
	// 'SoftwareInventoryCollection'. The resource collection should contain the
	// set of software components executed in the context of a host operating
	// system. This can include device drivers, applications, or offload workloads.
	// Software in this collection is generally updated using operating
	// system-centric methods.
	softwareInventory string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedUpdateImageFormats shall contain the image format types supported
	// by the service.
	//
	// Version added: v1.13.0
	SupportedUpdateImageFormats []SupportedUpdateImageFormatType
	// UpdateServiceCapabilities shall contain a link to a resource of type
	// 'UpdateServiceCapabilities' that contains the capabilities of the update
	// service.
	//
	// Version added: v1.16.0
	updateServiceCapabilities string
	// VerifyRemoteServerCertificate shall indicate whether the service will verify
	// the certificate of the server referenced by the 'ImageURI' parameter in
	// 'SimpleUpdate' prior to sending the transfer request with the certificates
	// found in the collection referenced by the 'RemoteServerCertificates'
	// property. If this property is not supported by the service, it shall be
	// assumed to be 'false'. This property should default to 'false' in order to
	// maintain compatibility with older clients. Regardless of the value of this
	// property, services may perform additional verification based on other
	// factors, such as the configuration of the 'SecurityPolicy' resource.
	//
	// Version added: v1.9.0
	VerifyRemoteServerCertificate bool
	// VerifyRemoteServerSSHKey shall indicate whether the service will verify the
	// SSH key of the server referenced by the 'ImageURI' parameter in
	// 'SimpleUpdate' prior to sending the transfer request with the keys found in
	// the collection referenced by the 'RemoteServerSSHKeys' property. If this
	// property is not supported by the service, it shall be assumed to be 'false'.
	// This property should default to 'false' in order to maintain compatibility
	// with older clients.
	//
	// Version added: v1.12.0
	VerifyRemoteServerSSHKey bool
	// activateTarget is the URL to send Activate requests.
	activateTarget string
	// generateSSHIdentityKeyPairTarget is the URL to send GenerateSSHIdentityKeyPair requests.
	generateSSHIdentityKeyPairTarget string
	// removeSSHIdentityKeyPairTarget is the URL to send RemoveSSHIdentityKeyPair requests.
	removeSSHIdentityKeyPairTarget string
	// simpleUpdateTarget is the URL to send SimpleUpdate requests.
	simpleUpdateTarget string
	// startUpdateTarget is the URL to send StartUpdate requests.
	startUpdateTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a UpdateService object from the raw JSON.
func (u *UpdateService) UnmarshalJSON(b []byte) error {
	type temp UpdateService
	type uActions struct {
		Activate                   ActionTarget `json:"#UpdateService.Activate"`
		GenerateSSHIdentityKeyPair ActionTarget `json:"#UpdateService.GenerateSSHIdentityKeyPair"`
		RemoveSSHIdentityKeyPair   ActionTarget `json:"#UpdateService.RemoveSSHIdentityKeyPair"`
		SimpleUpdate               ActionTarget `json:"#UpdateService.SimpleUpdate"`
		StartUpdate                ActionTarget `json:"#UpdateService.StartUpdate"`
	}
	var tmp struct {
		temp
		Actions                   uActions
		ClientCertificates        Link `json:"ClientCertificates"`
		FirmwareInventory         Link `json:"FirmwareInventory"`
		LocalImageStore           Link `json:"LocalImageStore"`
		PublicIdentitySSHKey      Link `json:"PublicIdentitySSHKey"`
		RemoteServerCertificates  Link `json:"RemoteServerCertificates"`
		RemoteServerSSHKeys       Link `json:"RemoteServerSSHKeys"`
		SoftwareInventory         Link `json:"SoftwareInventory"`
		UpdateServiceCapabilities Link `json:"UpdateServiceCapabilities"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*u = UpdateService(tmp.temp)

	// Extract the links to other entities for later
	u.activateTarget = tmp.Actions.Activate.Target
	u.generateSSHIdentityKeyPairTarget = tmp.Actions.GenerateSSHIdentityKeyPair.Target
	u.removeSSHIdentityKeyPairTarget = tmp.Actions.RemoveSSHIdentityKeyPair.Target
	u.simpleUpdateTarget = tmp.Actions.SimpleUpdate.Target
	u.startUpdateTarget = tmp.Actions.StartUpdate.Target
	u.clientCertificates = tmp.ClientCertificates.String()
	u.firmwareInventory = tmp.FirmwareInventory.String()
	u.localImageStore = tmp.LocalImageStore.String()
	u.publicIdentitySSHKey = tmp.PublicIdentitySSHKey.String()
	u.remoteServerCertificates = tmp.RemoteServerCertificates.String()
	u.remoteServerSSHKeys = tmp.RemoteServerSSHKeys.String()
	u.softwareInventory = tmp.SoftwareInventory.String()
	u.updateServiceCapabilities = tmp.UpdateServiceCapabilities.String()

	// This is a read/write object, so we need to save the raw object data for later
	u.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (u *UpdateService) Update() error {
	readWriteFields := []string{
		"HttpPushUriOptionsBusy",
		"HttpPushUriTargets",
		"HttpPushUriTargetsBusy",
		"ServiceEnabled",
		"VerifyRemoteServerCertificate",
		"VerifyRemoteServerSSHKey",
	}

	return u.UpdateFromRawData(u, u.RawData, readWriteFields)
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(c Client, uri string) (*UpdateService, error) {
	return GetObject[UpdateService](c, uri)
}

// ListReferencedUpdateServices gets the collection of UpdateService from
// a provided reference.
func ListReferencedUpdateServices(c Client, link string) ([]*UpdateService, error) {
	return GetCollectionObjects[UpdateService](c, link)
}

// This action shall activate specified software inventory instance.
// targets - This parameter shall contain an array of target software inventory
// instances to activate on staged devices. The service shall activate each
// software image on all applicable devices.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (u *UpdateService) Activate(targets []string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Targets"] = targets
	resp, taskInfo, err := PostWithTask(u.client,
		u.activateTarget, payload, u.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall generate a new SSH identity key-pair to be used with the
// 'UpdateService' resource. The service shall store the generated public key
// in the 'Key' resource referenced by the 'PublicIdentitySSHKey' property. If
// the 'UpdateService' resource already has an associated SSH identity
// key-pair, the service shall delete the key-pair and replace it with the new
// key-pair.
// curve - This parameter shall contain the curve to use with the SSH key. This
// parameter shall be required if the 'KeyType' parameter contains 'ECDSA' and
// shall be rejected for other values.
// keyLength - This parameter shall contain the length of the SSH key, in bits.
// This parameter shall be required if the 'KeyType' parameter contains 'RSA'
// and shall be rejected for other values.
// keyType - This parameter shall contain the type of SSH key.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (u *UpdateService) GenerateSSHIdentityKeyPair(curve ECDSACurveType, keyLength int, keyType SSHAlgoKeyType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Curve"] = curve
	payload["KeyLength"] = keyLength
	payload["KeyType"] = keyType
	resp, taskInfo, err := PostWithTask(u.client,
		u.generateSSHIdentityKeyPairTarget, payload, u.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall remove the private SSH identity key-pair used with the
// 'UpdateService' resource.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (u *UpdateService) RemoveSSHIdentityKeyPair() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(u.client,
		u.removeSSHIdentityKeyPairTarget, payload, u.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// UpdateServiceSimpleUpdateParameters holds the parameters for the SimpleUpdate action.
type UpdateServiceSimpleUpdateParameters struct {
	// ExcludeTargets shall contain zero or more URIs that indicate where not to
	// apply the update image. This parameter shall be ignored if the 'Targets'
	// parameter is provided and contains at least one member. These excluded
	// targets should correspond to software inventory instances or their related
	// items. If this parameter is not present or contains no targets, the service
	// shall apply the software image to all applicable targets, as determined by
	// the service. If an excluded target specifies a device resource, the software
	// image file shall not be applied to that specified device. If an excluded
	// target specifies a resource collection, the software image shall not be
	// applied to each applicable member of the specified collection. If an
	// excluded target resource specifies an 'Aggregate' resource, the software
	// image file shall not be applied to each applicable element of the specified
	// aggregate. If an excluded target resource specifies a 'ComputerSystem'
	// resource, the software image file shall not be applied to the applicable
	// components within the specified computer system.
	ExcludeTargets []string `json:"ExcludeTargets,omitempty"`
	// ForceUpdate shall indicate whether the service should bypass update policies
	// when applying the provided image, such as allowing a component to be
	// downgraded. Services may contain update policies that are never bypassed,
	// such as minimum version enforcement. If the client does not provide this
	// parameter, the service shall default this value to 'false'.
	ForceUpdate bool `json:"ForceUpdate,omitempty"`
	// ImageURI shall contain an RFC3986-defined URI that links to a software image
	// that the update service retrieves to install software in that image. This
	// URI should contain a scheme that describes the transfer protocol. If the
	// 'TransferProtocol' parameter is absent or not supported, and a transfer
	// protocol is not specified by a scheme contained within this URI, the service
	// shall use HTTP to get the image.
	ImageURI string `json:"ImageURI,omitempty"`
	// LocalImage shall indicate whether the service adds the image to the resource
	// collection referenced by the 'LocalImageStore' property. If the client does
	// not provide this parameter, the service shall default this value to 'false'.
	LocalImage bool `json:"LocalImage,omitempty"`
	// Password shall contain the password to access the URI specified by the
	// 'ImageURI' parameter.
	Password string `json:"Password,omitempty"`
	// Stage shall indicate whether the service stages the image on target devices
	// for a client to activate at a later time with the 'Activate' action on the
	// 'SoftwareInventory' resource. If the client does not provide this parameter,
	// the service shall default this value to 'false'.
	Stage bool `json:"Stage,omitempty"`
	// Targets shall contain zero or more URIs that indicate where to apply the
	// update image. These targets should correspond to software inventory
	// instances or their related items. If this parameter is not present or
	// contains no targets, the service shall apply the software image to all
	// applicable targets, as determined by the service. If the target specifies a
	// device resource, the software image file shall be applied to the specified
	// device. If the target specifies a resource collection, the software image
	// shall be applied to each applicable member of the specified collection. If
	// the target resource specifies an 'Aggregate' resource, the software image
	// file shall be applied to each applicable element of the specified aggregate.
	// If the target resource specifies a 'ComputerSystem' resource, the software
	// image file shall be applied to the applicable components within the
	// specified computer system.
	Targets []string `json:"Targets,omitempty"`
	// TransferProtocol shall contain the network protocol that the update service
	// shall use to retrieve the software image located at the URI specified by the
	// 'ImageURI' parameter. Services should ignore this parameter if the URI
	// specified by the 'ImageURI' parameter contains a scheme. If this parameter
	// is not provided (or supported), and a transfer protocol is not specified by
	// a scheme contained within this URI, the service shall use HTTP to retrieve
	// the image.
	TransferProtocol TransferProtocolType `json:"TransferProtocol,omitempty"`
	// Username shall contain the username to access the URI specified by the
	// 'ImageURI' parameter.
	Username string `json:"Username,omitempty"`
}

// This action shall update installed software components in a software image
// file located at an 'ImageURI' parameter-specified URI.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (u *UpdateService) SimpleUpdate(params *UpdateServiceSimpleUpdateParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(u.client,
		u.simpleUpdateTarget, params, u.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall start an update of software component that have been
// scheduled with the 'OperationApplyTime' value of 'OnStartUpdateRequest'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (u *UpdateService) StartUpdate() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(u.client,
		u.startUpdateTarget, payload, u.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ClientCertificates gets the ClientCertificates collection.
func (u *UpdateService) ClientCertificates() ([]*Certificate, error) {
	if u.clientCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](u.client, u.clientCertificates)
}

// FirmwareInventory gets the FirmwareInventory collection.
func (u *UpdateService) FirmwareInventory() ([]*SoftwareInventory, error) {
	if u.firmwareInventory == "" {
		return nil, nil
	}
	return GetCollectionObjects[SoftwareInventory](u.client, u.firmwareInventory)
}

// LocalImageStore gets the LocalImageStore collection.
func (u *UpdateService) LocalImageStore() ([]*SoftwareInventory, error) {
	if u.localImageStore == "" {
		return nil, nil
	}
	return GetCollectionObjects[SoftwareInventory](u.client, u.localImageStore)
}

// PublicIdentitySSHKey gets the PublicIdentitySSHKey linked resource.
func (u *UpdateService) PublicIdentitySSHKey() (*Key, error) {
	if u.publicIdentitySSHKey == "" {
		return nil, nil
	}
	return GetObject[Key](u.client, u.publicIdentitySSHKey)
}

// RemoteServerCertificates gets the RemoteServerCertificates collection.
func (u *UpdateService) RemoteServerCertificates() ([]*Certificate, error) {
	if u.remoteServerCertificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](u.client, u.remoteServerCertificates)
}

// RemoteServerSSHKeys gets the RemoteServerSSHKeys collection.
func (u *UpdateService) RemoteServerSSHKeys() ([]*Key, error) {
	if u.remoteServerSSHKeys == "" {
		return nil, nil
	}
	return GetCollectionObjects[Key](u.client, u.remoteServerSSHKeys)
}

// SoftwareInventory gets the SoftwareInventory collection.
func (u *UpdateService) SoftwareInventory() ([]*SoftwareInventory, error) {
	if u.softwareInventory == "" {
		return nil, nil
	}
	return GetCollectionObjects[SoftwareInventory](u.client, u.softwareInventory)
}

// UpdateServiceCapabilities gets the UpdateServiceCapabilities linked resource.
func (u *UpdateService) UpdateServiceCapabilities() (*UpdateServiceCapabilities, error) {
	if u.updateServiceCapabilities == "" {
		return nil, nil
	}
	return GetObject[UpdateServiceCapabilities](u.client, u.updateServiceCapabilities)
}

// HTTPPushURIApplyTime shall contain settings for when to apply
// 'HttpPushUri'-provided software.
type HTTPPushURIApplyTime struct {
	// ApplyTime shall indicate the time when to apply the 'HttpPushUri'-provided
	// software update.
	//
	// Version added: v1.4.0
	ApplyTime UpdateServiceApplyTime
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance
	// window as the number of seconds after the time specified by the
	// 'MaintenanceWindowStartTime' property. This property shall be required if
	// the 'HttpPushUriApplyTime' property value is 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.4.0
	MaintenanceWindowDurationInSeconds uint
	// MaintenanceWindowStartTime shall indicate the date and time when the service
	// can start to apply the 'HttpPushUri'-provided software as part of a
	// maintenance window. This property shall be required if the
	// 'HttpPushUriApplyTime' property value is 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	//
	// Version added: v1.4.0
	MaintenanceWindowStartTime string
}

// HTTPPushURIOptions shall contain settings and requirements of the service for
// 'HttpPushUri'-provided software updates.
type HTTPPushURIOptions struct {
	// ForceUpdate shall indicate whether the service should bypass update policies
	// when applying the 'HttpPushUri'-provided image, such as allowing a component
	// to be downgraded. Services may contain update policies that are never
	// bypassed, such as minimum version enforcement. If this property is not
	// present, it shall be assumed to be 'false'.
	//
	// Version added: v1.11.0
	ForceUpdate bool
	// HTTPPushURIApplyTime shall contain settings for when to apply
	// 'HttpPushUri'-provided firmware.
	//
	// Version added: v1.4.0
	HTTPPushURIApplyTime HTTPPushURIApplyTime `json:"HttpPushUriApplyTime"`
}

// UpdateParameters shall contain the update parameters when passing the update
// image when using the URI specified by the 'MultipartHttpPushUri' property to
// push a software image.
type UpdateParameters struct {
	// ExcludeTargets shall contain zero or more URIs that indicate where not to
	// apply the update image. This property shall be ignored if the 'Targets'
	// property is provided and contains at least one member. These excluded
	// targets should correspond to software inventory instances or their related
	// items. If this parameter is not present or contains no targets, the service
	// shall apply the software image to all applicable targets, as determined by
	// the service. If an excluded target specifies a device resource, the software
	// image file shall not be applied to that specified device. If an excluded
	// target specifies a resource collection, the software image shall not be
	// applied to each applicable member of the specified collection. If an
	// excluded target resource specifies an 'Aggregate' resource, the software
	// image file shall not be applied to each applicable element of the specified
	// aggregate. If an excluded target resource specifies a 'ComputerSystem'
	// resource, the software image file shall not be applied to the applicable
	// components within the specified computer system.
	//
	// Version added: v1.17.0
	ExcludeTargets []string
	// ForceUpdate shall indicate whether the service should bypass update policies
	// when applying the provided image, such as allowing a component to be
	// downgraded. Services may contain update policies that are never bypassed,
	// such as minimum version enforcement. If the client does not provide this
	// parameter, the service shall default this value to 'false'.
	//
	// Version added: v1.11.0
	ForceUpdate bool
	// LocalImage shall indicate whether the service adds the image to the resource
	// collection referenced by the 'LocalImageStore' property. If the client does
	// not provide this parameter, the service shall default this value to 'false'.
	//
	// Version added: v1.17.0
	LocalImage bool
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.8.0
	OEM json.RawMessage `json:"Oem"`
	// Stage shall indicate whether the service stages the image on target devices
	// for a client to activate at a later time with the 'Activate' action on the
	// 'SoftwareInventory' resource. If the client does not provide this parameter,
	// the service shall default this value to 'false'.
	//
	// Version added: v1.16.0
	Stage bool
	// Targets shall contain zero or more URIs that indicate where to apply the
	// update image when using the URI specified by the 'MultipartHttpPushUri'
	// property to push a software image. These targets should correspond to
	// software inventory instances or their related items. If this property is not
	// present or contains no targets, the service shall apply the software image
	// to all applicable targets, as determined by the service. If the target
	// specifies a device resource, the software image file shall be applied to the
	// specified device. If the target specifies a resource collection, the
	// software image shall be applied to each applicable member of the specified
	// collection. If the target resource specifies an 'Aggregate' resource, the
	// software image file shall be applied to each applicable element of the
	// specified aggregate. If the target resource specifies a 'ComputerSystem'
	// resource, the software image file shall be applied to the applicable
	// components within the specified computer system.
	//
	// Version added: v1.6.0
	Targets []string
}
