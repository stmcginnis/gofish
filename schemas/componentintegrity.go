//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.3 - #ComponentIntegrity.v1_3_2.ComponentIntegrity

package schemas

import (
	"encoding/json"
)

type ComponentIntegrityType string

const (
	// SPDMComponentIntegrityType shall indicate the integrity information is
	// obtained through the Security Protocol and Data Model (SPDM) protocol as
	// defined in DMTF DSP0274.
	SPDMComponentIntegrityType ComponentIntegrityType = "SPDM"
	// TPMComponentIntegrityType shall indicate the integrity information is
	// related to a Trusted Platform Module (TPM) as defined by the Trusted
	// Computing Group (TCG).
	TPMComponentIntegrityType ComponentIntegrityType = "TPM"
	// TCMComponentIntegrityType shall indicate the integrity information is
	// related to a Trusted Cryptography Module (TCM) as defined by the China TCM
	// Union (TCMU).
	TCMComponentIntegrityType ComponentIntegrityType = "TCM"
	// TPCMComponentIntegrityType shall indicate the integrity information is
	// related to a Trusted Platform Control Module (TPCM) as defined by the
	// Zhongguancun Trusted Computing Industry Alliance (ZTCIA).
	TPCMComponentIntegrityType ComponentIntegrityType = "TPCM"
	// OEMComponentIntegrityType shall indicate the integrity information is
	// OEM-specific and the OEM section may include additional information.
	OEMComponentIntegrityType ComponentIntegrityType = "OEM"
)

type DMTFmeasurementTypes string

const (
	// ImmutableROMDMTFmeasurementTypes Immutable ROM.
	ImmutableROMDMTFmeasurementTypes DMTFmeasurementTypes = "ImmutableROM"
	// MutableFirmwareDMTFmeasurementTypes Mutable firmware or any mutable code.
	MutableFirmwareDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmware"
	// HardwareConfigurationDMTFmeasurementTypes Hardware configuration, such as
	// straps.
	HardwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "HardwareConfiguration"
	// FirmwareConfigurationDMTFmeasurementTypes Firmware configuration, such as
	// configurable firmware policy.
	FirmwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "FirmwareConfiguration"
	// MutableFirmwareVersionDMTFmeasurementTypes Mutable firmware version.
	MutableFirmwareVersionDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareVersion"
	// MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes Mutable firmware
	// security version number.
	MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareSecurityVersionNumber"
	// MeasurementManifestDMTFmeasurementTypes Measurement Manifest.
	MeasurementManifestDMTFmeasurementTypes DMTFmeasurementTypes = "MeasurementManifest"
)

type MeasurementSpecification string

const (
	// DMTFMeasurementSpecification shall indicate the measurement specification is
	// defined by DMTF in DSP0274.
	DMTFMeasurementSpecification MeasurementSpecification = "DMTF"
)

type SPDMmeasurementSummaryType string

const (
	// TCBSPDMmeasurementSummaryType The measurement summary covers the TCB.
	TCBSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "TCB"
	// AllSPDMmeasurementSummaryType The measurement summary covers all
	// measurements in SPDM.
	AllSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "All"
)

type SecureSessionType string

const (
	// PlainSecureSessionType is a plain text session without any protection.
	PlainSecureSessionType SecureSessionType = "Plain"
	// EncryptedAuthenticatedSecureSessionType is an established session where both
	// encryption and authentication are protecting the communication.
	EncryptedAuthenticatedSecureSessionType SecureSessionType = "EncryptedAuthenticated"
	// AuthenticatedOnlySecureSessionType is an established session where only
	// authentication is protecting the communication.
	AuthenticatedOnlySecureSessionType SecureSessionType = "AuthenticatedOnly"
)

type VerificationStatus string

const (
	// SuccessVerificationStatus Successful verification.
	SuccessVerificationStatus VerificationStatus = "Success"
	// FailedVerificationStatus Unsuccessful verification.
	FailedVerificationStatus VerificationStatus = "Failed"
)

// ComponentIntegrity shall represent critical and pertinent security
// information about a specific device, system, software element, or other
// managed entity.
type ComponentIntegrity struct {
	Entity
	// ComponentIntegrityEnabled shall indicate whether security protocols are
	// enabled for the component. If 'ComponentIntegrityType' contains 'SPDM', a
	// value of 'false' shall prohibit the SPDM Requester from using SPDM to
	// communicate with the component identified by the 'TargetComponentURI'
	// property. If 'ComponentIntegrityType' contains 'TPM', a value of 'false'
	// shall disable the TPM component identified by the 'TargetComponentURI'
	// property entirely. If 'false', services shall not provide the TPM and SPDM
	// properties in response payloads for this resource. If 'false', services
	// shall reject action requests to this resource. If 'true', services shall
	// allow security protocols with the component identified by the
	// 'TargetComponentURI' property.
	ComponentIntegrityEnabled bool
	// ComponentIntegrityType shall contain the underlying security technology
	// providing integrity information for the component.
	ComponentIntegrityType ComponentIntegrityType
	// ComponentIntegrityTypeVersion shall contain the version of the security
	// technology indicated by the 'ComponentIntegrityType' property. If the
	// service has not established secure communication with the device or if
	// security protocols are disabled, this property shall contain an empty
	// string. If 'ComponentIntegrityType' contains 'SPDM', this property shall
	// contain the negotiated or selected SPDM protocol and shall follow the
	// regular expression pattern '^\d+\.\d+\.\d+$'. If 'ComponentIntegrityType'
	// contains 'TPM', this property shall contain the version of the TPM.
	ComponentIntegrityTypeVersion string
	// LastUpdated shall contain the date and time when information for the
	// component was last updated.
	LastUpdated string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SPDM shall contain integrity information about the SPDM Responder identified
	// by the 'TargetComponentURI' property as reported by an SPDM Requester. This
	// property shall be present if 'ComponentIntegrityType' contains 'SPDM' and if
	// 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property
	// shall be absent.
	SPDM SPDMinfo
	// Status shall contain any status or health properties of the resource.
	Status Status
	// TPM shall contain integrity information about the Trusted Platform Module
	// (TPM) identified by the 'TargetComponentURI' property. This property shall
	// be present if 'ComponentIntegrityType' contains 'TPM' and if
	// 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property
	// shall be absent.
	TPM TPMinfo
	// TargetComponentURI shall contain a link to the resource whose integrity
	// information is reported in this resource. If 'ComponentIntegrityType'
	// contains 'SPDM', this property shall contain a URI to the resource that
	// represents the SPDM Responder. If 'ComponentIntegrityType' contains 'TPM',
	// this property shall contain a URI with RFC6901-defined JSON fragment
	// notation to a member of the TrustedModules array in a ComputerSystem
	// resource that represents the TPM or a resource of type 'TrustedComponent'
	// that represents the TPM.
	TargetComponentURI string
	// sPDMGetSignedMeasurementsTarget is the URL to send SPDMGetSignedMeasurements requests.
	sPDMGetSignedMeasurementsTarget string
	// tPMGetSignedMeasurementsTarget is the URL to send TPMGetSignedMeasurements requests.
	tPMGetSignedMeasurementsTarget string
	// componentsProtected are the URIs for ComponentsProtected.
	componentsProtected []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a ComponentIntegrity object from the raw JSON.
func (c *ComponentIntegrity) UnmarshalJSON(b []byte) error {
	type temp ComponentIntegrity
	type cActions struct {
		SPDMGetSignedMeasurements ActionTarget `json:"#ComponentIntegrity.SPDMGetSignedMeasurements"`
		TPMGetSignedMeasurements  ActionTarget `json:"#ComponentIntegrity.TPMGetSignedMeasurements"`
	}
	type cLinks struct {
		ComponentsProtected Links `json:"ComponentsProtected"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ComponentIntegrity(tmp.temp)

	// Extract the links to other entities for later
	c.sPDMGetSignedMeasurementsTarget = tmp.Actions.SPDMGetSignedMeasurements.Target
	c.tPMGetSignedMeasurementsTarget = tmp.Actions.TPMGetSignedMeasurements.Target
	c.componentsProtected = tmp.Links.ComponentsProtected.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ComponentIntegrity) Update() error {
	readWriteFields := []string{
		"ComponentIntegrityEnabled",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetComponentIntegrity will get a ComponentIntegrity instance from the service.
func GetComponentIntegrity(c Client, uri string) (*ComponentIntegrity, error) {
	return GetObject[ComponentIntegrity](c, uri)
}

// ListReferencedComponentIntegritys gets the collection of ComponentIntegrity from
// a provided reference.
func ListReferencedComponentIntegritys(c Client, link string) ([]*ComponentIntegrity, error) {
	return GetCollectionObjects[ComponentIntegrity](c, link)
}

// This action shall generate a cryptographic signed statement over the given
// nonce and measurements corresponding to the SPDM Responder. This action
// shall not be present if the 'ComponentIntegrityType' property does not
// contain the value 'SPDM'. The SPDM Requester shall issue one or more SPDM
// 'GET_MEASUREMENTS' requests for each of the requested measurement indices to
// the SPDM Responder. When the SPDM 'GET_MEASUREMENTS' requests are made for
// version 1.2, the parameter 'RawBitStreamRequested' shall contain '0'. The
// SPDM Requester shall provide the nonce for the action to the SPDM Responder
// in the last SPDM 'GET_MEASUREMENTS' request. The SPDM Requester shall
// request a signature in the last SPDM 'GET_MEASUREMENTS' request.
// measurementIndices - This parameter shall contain an array of indices that
// identify the measurement blocks to sign. This array shall contain one or
// more unique values between '0' to '254', inclusive, or contain a single
// value of '255'. If not provided by the client, the value shall be assumed to
// be an array containing a single value of '255'.
// nonce - This parameter shall contain a 32-byte hex-encoded string that is
// signed with the measurements. If not provided by the client, the SPDM
// Requester shall generate the nonce. The value should be unique and generated
// using a random or a pseudo-random generator. The SPDM Requester shall send
// this value to the SPDM Responder in the SPDM 'GET_MEASUREMENTS' request.
// slotID - This parameter shall contain the SPDM slot identifier for the
// certificate containing the private key to generate the signature over the
// measurements. If not provided by the client, the value shall be assumed to
// be '0'. The SPDM Requester shall send this value to the SPDM Responder in
// the SPDM 'GET_MEASUREMENTS' request.
func (c *ComponentIntegrity) SPDMGetSignedMeasurements(measurementIndices []int, nonce string, slotID int) (*SPDMGetSignedMeasurementsResponse, error) {
	payload := make(map[string]any)
	payload["MeasurementIndices"] = measurementIndices
	payload["Nonce"] = nonce
	payload["SlotId"] = slotID

	resp, err := c.PostWithResponse(c.sPDMGetSignedMeasurementsTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result SPDMGetSignedMeasurementsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ComponentIntegrityTPMGetSignedMeasurementsParameters holds the parameters for the TPMGetSignedMeasurements action.
type ComponentIntegrityTPMGetSignedMeasurementsParameters struct {
	// Certificate shall contain the reference to the certificate installed on the
	// TPM that represents the TPM's attestation key for the 'signHandle' parameter
	// of the 'TPM2_Quote' command defined in the Trusted Platform Module Library
	// Specification.
	Certificate string `json:"Certificate,omitempty"`
	// Nonce shall contain a set of bytes as a hex-encoded string that is signed
	// with the measurements. Services shall reject the action request if the
	// number of bytes provided is larger than the value specified by the
	// 'NonceSizeBytesMaximum' property in the 'TPM' property. If not provided by
	// the client, the service shall generate the nonce. The value should be unique
	// and generated using a random or a pseudo-random generator. The service shall
	// send this value to the TPM in the 'qualifyingData' parameter of the
	// 'TPM2_Quote' command defined in the Trusted Platform Module Library
	// Specification.
	Nonce string `json:"Nonce,omitempty"`
	// PCRSelection shall contain a Base64-encoded string, with padding characters,
	// of the 'TPML_PCR_SELECTION' object as defined by the Trusted Platform Module
	// Library Specification, that identifies the PCRs to sign. The service shall
	// send this value to the TPM in the 'PCRselect' parameter of the 'TPM2_Quote'
	// command defined in the Trusted Platform Module Library Specification.
	PCRSelection string `json:"PCRSelection,omitempty"`
	// Scheme shall contain a Base64-encoded string, with padding characters, of
	// the 'TPMT_SIG_SCHEME' object as defined in the Trusted Platform Module
	// Library Specification, that identifies the signing scheme to use for the TPM
	// attestation key. The service shall send this value to the TPM in the
	// 'inScheme' parameter of the 'TPM2_Quote' command defined in the Trusted
	// Platform Module Library Specification.
	Scheme string `json:"Scheme,omitempty"`
}

// This action shall generate a cryptographic signed statement over the given
// nonce and PCRs of the TPM for TPM 2.0 devices. This action shall not be
// present if the 'ComponentIntegrityType' property does not contain the value
// 'TPM'.
func (c *ComponentIntegrity) TPMGetSignedMeasurements(params *ComponentIntegrityTPMGetSignedMeasurementsParameters) (*TPMGetSignedMeasurementsResponse, error) {
	resp, err := c.PostWithResponse(c.tPMGetSignedMeasurementsTarget, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result TPMGetSignedMeasurementsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ComponentsProtected gets the ComponentsProtected linked resources.
func (c *ComponentIntegrity) ComponentsProtected() ([]*Entity, error) {
	return GetObjects[Entity](c.client, c.componentsProtected)
}

// CommonAuthInfo shall contain common identity-related authentication
// information.
type CommonAuthInfo struct {
	// ComponentCertificate shall contain a link to a resource of type
	// 'Certificate' that represents the identity of the component referenced by
	// the 'TargetComponentURI' property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the
	// identity of the component referenced by the 'TargetComponentURI' property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a CommonAuthInfo object from the raw JSON.
func (c *CommonAuthInfo) UnmarshalJSON(b []byte) error {
	type temp CommonAuthInfo
	var tmp struct {
		temp
		ComponentCertificate Link `json:"ComponentCertificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CommonAuthInfo(tmp.temp)

	// Extract the links to other entities for later
	c.componentCertificate = tmp.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the ComponentCertificate linked resource.
func (c *CommonAuthInfo) ComponentCertificate(client Client) (*Certificate, error) {
	if c.componentCertificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, c.componentCertificate)
}

// CommunicationInfo shall contain information about communication between two
// components.
type CommunicationInfo struct {
	// Sessions shall contain an array of the active sessions or communication
	// channels between two components. The active sessions or communication
	// channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// SPDMGetSignedMeasurementsResponse shall contain the SPDM signed measurements
// from an SPDM Responder.
type SPDMGetSignedMeasurementsResponse struct {
	// Certificate shall contain a link to a resource of type 'Certificate' that
	// represents the certificate corresponding to the SPDM slot identifier that
	// can be used to validate the signature. This property shall not be present if
	// the SlotId parameter contains the value '15'.
	certificate string
	// HashingAlgorithm shall contain the hashing algorithm negotiated between the
	// SPDM Requester and the SPDM Responder. The allowable values for this
	// property shall be the hash algorithm names found in the 'BaseHashAlgo' field
	// of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm
	// is an extended algorithm, this property shall contain the value 'OEM'.
	HashingAlgorithm string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PublicKey shall contain a Privacy Enhanced Mail (PEM)-encoded public key, as
	// defined in section 13 of RFC7468, that can be used to validate the
	// signature. This property shall only be present when the SPDM Requester was
	// pre-provisioned with the SPDM Responder's public key and the SlotId
	// parameter contains the value '15'.
	PublicKey string
	// SignedMeasurements shall contain the cryptographic signed statement over the
	// given nonce and measurement blocks corresponding to the requested
	// measurement indices. If the SPDM version is 1.2, this value shall be a
	// concatenation of SPDM 'VCA' and 'GET_MEASUREMENTS' requests and responses
	// exchanged between the SPDM Requester and the SPDM Responder. If SPDM version
	// is 1.0 or 1.1, this value shall be a concatenation of SPDM
	// 'GET_MEASUREMENTS' requests and responses exchanged between the SPDM
	// Requester and the SPDM Responder. The last 'MEASUREMENTS' response shall
	// contain a signature generated over the 'L2' string by the SPDM Responder.
	SignedMeasurements string
	// SigningAlgorithm shall contain the asymmetric signing algorithm negotiated
	// between the SPDM Requester and the SPDM Responder. The allowable values for
	// this property shall be the asymmetric key signature algorithm names found in
	// the 'SPDM Asymmetric Signature Reference Information' table in DSP0274. If
	// the algorithm is an extended algorithm, this property shall contain the
	// value 'OEM'.
	SigningAlgorithm string
	// Version shall contain the SPDM version negotiated between the SPDM Requester
	// and the SPDM Responder to generate the cryptographic signed statement. For
	// example, '1.0', '1.1', or '1.2'.
	Version string
}

// UnmarshalJSON unmarshals a SPDMGetSignedMeasurementsResponse object from the raw JSON.
func (s *SPDMGetSignedMeasurementsResponse) UnmarshalJSON(b []byte) error {
	type temp SPDMGetSignedMeasurementsResponse
	var tmp struct {
		temp
		Certificate Link `json:"Certificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SPDMGetSignedMeasurementsResponse(tmp.temp)

	// Extract the links to other entities for later
	s.certificate = tmp.Certificate.String()

	return nil
}

// Certificate gets the Certificate linked resource.
func (s *SPDMGetSignedMeasurementsResponse) Certificate(client Client) (*Certificate, error) {
	if s.certificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, s.certificate)
}

// SPDMcommunication shall contain information about communication between two
// components.
type SPDMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication
	// channels between two components. The active sessions or communication
	// channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// SPDMidentity shall contain identity authentication information about the SPDM
// Requester and SPDM Responder.
type SPDMidentity struct {
	// RequesterAuthentication shall contain authentication information of the
	// identity of the SPDM Requester.
	RequesterAuthentication SPDMrequesterAuth
	// ResponderAuthentication shall contain authentication information of the
	// identity of the SPDM Responder.
	ResponderAuthentication SPDMresponderAuth
}

// SPDMinfo shall contain integrity information about an SPDM Responder as
// reported by an SPDM Requester.
type SPDMinfo struct {
	// ComponentCommunication shall contain information about communication between
	// the SPDM Requester and SPDM Responder.
	ComponentCommunication SPDMcommunication
	// IdentityAuthentication shall contain identity authentication information
	// about the SPDM Requester and SPDM Responder.
	IdentityAuthentication SPDMidentity
	// MeasurementSet shall contain measurement information for the SPDM Responder.
	MeasurementSet SPDMmeasurementSet
	// Requester shall contain a link to the resource representing the SPDM
	// Responder that is reporting the integrity of the SPDM Responder identified
	// by the 'TargetComponentURI' property.
	requester string
}

// UnmarshalJSON unmarshals a SPDMinfo object from the raw JSON.
func (s *SPDMinfo) UnmarshalJSON(b []byte) error {
	type temp SPDMinfo
	var tmp struct {
		temp
		Requester Link `json:"Requester"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SPDMinfo(tmp.temp)

	// Extract the links to other entities for later
	s.requester = tmp.Requester.String()

	return nil
}

// Requester gets the Requester linked resource.
func (s *SPDMinfo) Requester(client Client) (*Entity, error) {
	if s.requester == "" {
		return nil, nil
	}
	return GetObject[Entity](client, s.requester)
}

// SPDMmeasurementSet shall contain SPDM Responder measurement information.
type SPDMmeasurementSet struct {
	// MeasurementSpecification shall contain the measurement specification
	// negotiated between the SPDM Requester and SPDM Responder.
	MeasurementSpecification MeasurementSpecification
	// MeasurementSummary shall contain a Base64-encoded string, with padding
	// characters, of the measurement summary using the hash algorithm indicated by
	// the 'MeasurementSummaryHashAlgorithm' property.
	MeasurementSummary string
	// MeasurementSummaryHashAlgorithm shall contain the hash algorithm used to
	// compute the measurement summary. The allowable values for this property
	// shall be the hash algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an
	// extended algorithm, this property shall contain the value 'OEM'.
	MeasurementSummaryHashAlgorithm string
	// MeasurementSummaryType shall contain the type of measurement summary.
	MeasurementSummaryType SPDMmeasurementSummaryType
	// Measurements shall contain measurements from an SPDM Responder.
	Measurements []SPDMsingleMeasurement
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
}

// SPDMrequesterAuth shall contain authentication information of the identity of
// the SPDM Requester.
type SPDMrequesterAuth struct {
	// ProvidedCertificate shall contain a link to a resource of type 'Certificate'
	// that represents the identity of the SPDM Requester provided in mutual
	// authentication.
	providedCertificate string
}

// UnmarshalJSON unmarshals a SPDMrequesterAuth object from the raw JSON.
func (s *SPDMrequesterAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMrequesterAuth
	var tmp struct {
		temp
		ProvidedCertificate Link `json:"ProvidedCertificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SPDMrequesterAuth(tmp.temp)

	// Extract the links to other entities for later
	s.providedCertificate = tmp.ProvidedCertificate.String()

	return nil
}

// ProvidedCertificate gets the ProvidedCertificate linked resource.
func (s *SPDMrequesterAuth) ProvidedCertificate(client Client) (*Certificate, error) {
	if s.providedCertificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, s.providedCertificate)
}

// SPDMresponderAuth shall contain common identity-related authentication
// information.
type SPDMresponderAuth struct {
	// ComponentCertificate shall contain a link to a resource of type
	// 'Certificate' that represents the identity of the component referenced by
	// the 'TargetComponentURI' property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the
	// identity of the component referenced by the 'TargetComponentURI' property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a SPDMresponderAuth object from the raw JSON.
func (s *SPDMresponderAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMresponderAuth
	var tmp struct {
		temp
		ComponentCertificate Link `json:"ComponentCertificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SPDMresponderAuth(tmp.temp)

	// Extract the links to other entities for later
	s.componentCertificate = tmp.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the ComponentCertificate linked resource.
func (s *SPDMresponderAuth) ComponentCertificate(client Client) (*Certificate, error) {
	if s.componentCertificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, s.componentCertificate)
}

// SPDMsingleMeasurement shall contain a single SPDM measurement for an SPDM
// Responder.
type SPDMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the
	// measurement was last updated.
	LastUpdated string
	// Measurement shall contain a Base64-encoded string, with padding characters,
	// of the measurement using the hash algorithm indicated by the
	// 'MeasurementHashAlgorithm' property. This property shall not contain a raw
	// bit stream as a measurement. If the SPDM Responder provides a raw bit
	// stream, the SPDM Requester may apply a hash algorithm to the raw bit stream
	// in order to report the measurement.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute
	// the measurement. The allowable values for this property shall be the hash
	// algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an
	// extended algorithm, this property shall contain the value 'OEM'. This
	// property shall not be present if MeasurementSpecification does not contain
	// 'DMTF'.
	MeasurementHashAlgorithm string
	// MeasurementIndex shall contain the index of the measurement.
	MeasurementIndex *int `json:",omitempty"`
	// MeasurementType shall contain the type or characteristics of the data that
	// this measurement represents. This property shall not be present if
	// MeasurementSpecification does not contain 'DMTF'.
	MeasurementType DMTFmeasurementTypes
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartofSummaryHash shall indicate if this measurement is part of the
	// measurement summary in the 'MeasurementSummary' property. If this property
	// is not present, it shall be assumed to be 'false'.
	PartofSummaryHash bool
	// SecurityVersionNumber shall contain an 8-byte hex-encoded string of the
	// security version number the measurement represents. This property shall only
	// be present if 'MeasurementType' contains the value
	// 'MutableFirmwareSecurityVersionNumber'.
	//
	// Version added: v1.1.0
	SecurityVersionNumber string
}

// SingleSessionInfo shall contain information about a single communication
// channel or session between two components.
type SingleSessionInfo struct {
	// SessionID shall contain the unique identifier for the active session or
	// communication channel between two components.
	SessionID *int `json:"SessionId,omitempty"`
	// SessionType shall contain the type of session or communication channel
	// between two components.
	SessionType SecureSessionType
}

// TPMGetSignedMeasurementsResponse shall contain the TPM signed PCR
// measurements from a TPM.
type TPMGetSignedMeasurementsResponse struct {
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.2.0
	OEM json.RawMessage `json:"Oem"`
	// SignedMeasurements shall contain a Base64-encoded string, with padding
	// characters, of the cryptographic signed statement generated by the signer.
	// This value shall be the concatenation of the 'quoted' and 'signature'
	// response values of the 'TPM2_Quote' command defined in the Trusted Platform
	// Module Library Specification.
	//
	// Version added: v1.2.0
	SignedMeasurements string
}

// TPMauth shall contain common identity-related authentication information.
type TPMauth struct {
	// ComponentCertificate shall contain a link to a resource of type
	// 'Certificate' that represents the identity of the component referenced by
	// the 'TargetComponentURI' property.
	componentCertificate string
	// VerificationStatus shall contain the status of the verification of the
	// identity of the component referenced by the 'TargetComponentURI' property.
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a TPMauth object from the raw JSON.
func (t *TPMauth) UnmarshalJSON(b []byte) error {
	type temp TPMauth
	var tmp struct {
		temp
		ComponentCertificate Link `json:"ComponentCertificate"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TPMauth(tmp.temp)

	// Extract the links to other entities for later
	t.componentCertificate = tmp.ComponentCertificate.String()

	return nil
}

// ComponentCertificate gets the ComponentCertificate linked resource.
func (t *TPMauth) ComponentCertificate(client Client) (*Certificate, error) {
	if t.componentCertificate == "" {
		return nil, nil
	}
	return GetObject[Certificate](client, t.componentCertificate)
}

// TPMcommunication shall contain information about communication between two
// components.
type TPMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication
	// channels between two components. The active sessions or communication
	// channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// TPMinfo shall contain integrity information about a Trusted Platform Module
// (TPM).
type TPMinfo struct {
	// ComponentCommunication shall contain information about communication with
	// the TPM.
	ComponentCommunication TPMcommunication
	// IdentityAuthentication shall contain identity authentication information
	// about the TPM.
	IdentityAuthentication TPMauth
	// MeasurementSet shall contain measurement information from the TPM.
	MeasurementSet TPMmeasurementSet
	// NonceSizeBytesMaximum shall contain the maximum number of bytes that can be
	// specified in the 'Nonce' parameter of the 'TPMGetSignedMeasurements' action.
	//
	// Version added: v1.2.0
	NonceSizeBytesMaximum *uint `json:",omitempty"`
}

// TPMmeasurementSet shall contain Trusted Computing Group TPM measurement
// information.
type TPMmeasurementSet struct {
	// Measurements shall contain measurements from a TPM.
	Measurements []TPMsingleMeasurement
}

// TPMsingleMeasurement shall contain a single Trusted Computing Group TPM
// measurement.
type TPMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the
	// measurement was last updated.
	LastUpdated string
	// Measurement shall contain a Base64-encoded string, with padding characters,
	// of the PCR digest using the hashing algorithm indicated by the
	// 'MeasurementHashAlgorithm' property.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute
	// the measurement. The allowable values for this property shall be the strings
	// in the 'Algorithm Name' field of the 'TCG_ALG_ID Constants' table, formerly
	// the 'TPM_ALG_ID Constants' table, within the 'Trusted Computing Group
	// Algorithm Registry'.
	MeasurementHashAlgorithm string
	// PCR shall contain the Platform Configuration Register (PCR) bank of the
	// measurement.
	PCR *int `json:",omitempty"`
}
