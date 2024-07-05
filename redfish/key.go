//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ECDSACurveType string

const (
	// NISTB233ECDSACurveType NIST B-233.
	NISTB233ECDSACurveType ECDSACurveType = "NISTB233"
	// NISTB409ECDSACurveType NIST B-409.
	NISTB409ECDSACurveType ECDSACurveType = "NISTB409"
	// NISTK163ECDSACurveType NIST K-163.
	NISTK163ECDSACurveType ECDSACurveType = "NISTK163"
	// NISTK233ECDSACurveType NIST K-233.
	NISTK233ECDSACurveType ECDSACurveType = "NISTK233"
	// NISTK283ECDSACurveType NIST K-283.
	NISTK283ECDSACurveType ECDSACurveType = "NISTK283"
	// NISTK409ECDSACurveType NIST K-409.
	NISTK409ECDSACurveType ECDSACurveType = "NISTK409"
	// NISTP192ECDSACurveType NIST P-192.
	NISTP192ECDSACurveType ECDSACurveType = "NISTP192"
	// NISTP224ECDSACurveType NIST P-224.
	NISTP224ECDSACurveType ECDSACurveType = "NISTP224"
	// NISTP256ECDSACurveType NIST P-256.
	NISTP256ECDSACurveType ECDSACurveType = "NISTP256"
	// NISTP384ECDSACurveType NIST P-384.
	NISTP384ECDSACurveType ECDSACurveType = "NISTP384"
	// NISTP521ECDSACurveType NIST P-521.
	NISTP521ECDSACurveType ECDSACurveType = "NISTP521"
	// NISTT571ECDSACurveType NIST T-571.
	NISTT571ECDSACurveType ECDSACurveType = "NISTT571"
)

type SSHAlgoKeyType string

const (
	// DSASSHKeyType is DSA.
	DSASSHAlgoKeyType SSHAlgoKeyType = "DSA"
	// ECDSASSHKeyType is ECDSA.
	ECDSASSHAlgoKeyType SSHAlgoKeyType = "ECDSA"
	// ED25519SSHKeyType is Ed25519.
	ED25519SSHAlgoKeyType SSHAlgoKeyType = "Ed25519"
	// RSASSHKeyType is RSA.
	RSASSHAlgoKeyType SSHAlgoKeyType = "RSA"
)

type KeyType string

const (
	// NVMeoFKeyType shall indicate the format of the key is defined by one of the NVMe specifications.
	NVMeoFKeyType KeyType = "NVMeoF"
	// SSHKeyType shall indicate the format of the key is defined by one of the SSH public key formats as defined in,
	// but not limited to, RFC4253, RFC4716, or RFC8709.
	SSHKeyType KeyType = "SSH"
)

// Key shall represent a key for a Redfish implementation.
type Key struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// KeyString shall contain the key, and the format shall follow the requirements specified by the KeyType property
	// value.
	KeyString string
	// KeyType shall contain the format type for the key.
	KeyType KeyType
	// NVMeoF shall contain NVMe-oF specific properties for this key. This property shall be present if KeyType
	// contains the value 'NVMeoF'.
	NVMeoF KeyNVMeoF
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SSH shall contain SSH specific properties for this key. This property shall be present if KeyType contains the
	// value 'SSH'.
	SSH SSHType
	// UserDescription shall contain a user-provided string that describes the key.
	UserDescription string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Key object from the raw JSON.
func (key *Key) UnmarshalJSON(b []byte) error {
	type temp Key
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*key = Key(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	key.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (key *Key) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Key)
	original.UnmarshalJSON(key.rawData)

	readWriteFields := []string{
		"UserDescription",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(key).Elem()

	return key.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetKey will get a Key instance from the service.
func GetKey(c common.Client, uri string) (*Key, error) {
	return common.GetObject[Key](c, uri)
}

// ListReferencedKeys gets the collection of Key from
// a provided reference.
func ListReferencedKeys(c common.Client, link string) ([]*Key, error) {
	return common.GetCollectionObjects[Key](c, link)
}

// KeyNVMeoF shall contain NVMe-oF specific properties for a key.
type KeyNVMeoF struct {
	// HostKeyID shall contain the value of the ID property of the Key resource representing the host key paired with
	// this target key. An empty string shall indicate the key is not paired. This property shall be absent for host
	// keys.
	HostKeyID string
	// NQN shall contain the NVMe Qualified Name (NQN) of the host or target subsystem associated with this key. The
	// value of this property shall follow the NQN format defined by the NVMe Base Specification.
	NQN string
	// OEMSecurityProtocolType shall contain the OEM-defined security protocol that this key uses. The value shall be
	// derived from the contents of the KeyString property. This property shall be present if SecurityProtocolType
	// contains the value 'OEM'.
	OEMSecurityProtocolType string
	// SecureHashAllowList shall contain the secure hash algorithms allowed with the usage of this key. An empty list
	// or the absence of this property shall indicate any secure hash algorithms are allowed with this key.
	SecureHashAllowList []NVMeoFSecureHashType
	// SecurityProtocolType shall contain the security protocol that this key uses. The value shall be derived from the
	// contents of the KeyString property.
	SecurityProtocolType NVMeoFSecurityProtocolType
}

// SSHType shall contain SSH specific properties for a key.
type SSHType struct {
	// Comment shall contain the user-specified comment associated with this key, which typically contains the client's
	// username and host name.
	Comment string
	// Fingerprint shall contain the fingerprint of the key.
	Fingerprint string
	// RemoteServerHostName shall contain the host name of the remote server associated with this key.
	RemoteServerHostName string
}
