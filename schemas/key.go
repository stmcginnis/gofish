//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Key.v1_4_1.json
// 2023.3 - #Key.v1_4_1.Key

package schemas

import (
	"encoding/json"
)

type ECDSACurveType string

const (
	// NISTP256ECDSACurveType shall indicate the 'nistp256' curve in RFC5656.
	NISTP256ECDSACurveType ECDSACurveType = "NISTP256"
	// NISTP384ECDSACurveType shall indicate the 'nistp384' curve in RFC5656.
	NISTP384ECDSACurveType ECDSACurveType = "NISTP384"
	// NISTP521ECDSACurveType shall indicate the 'nistp521' curve in RFC5656.
	NISTP521ECDSACurveType ECDSACurveType = "NISTP521"
	// NISTK163ECDSACurveType shall indicate the 'nistk163' curve in RFC5656.
	NISTK163ECDSACurveType ECDSACurveType = "NISTK163"
	// NISTP192ECDSACurveType shall indicate the 'nistp192' curve in RFC5656.
	NISTP192ECDSACurveType ECDSACurveType = "NISTP192"
	// NISTP224ECDSACurveType shall indicate the 'nistp224' curve in RFC5656.
	NISTP224ECDSACurveType ECDSACurveType = "NISTP224"
	// NISTK233ECDSACurveType shall indicate the 'nistk233' curve in RFC5656.
	NISTK233ECDSACurveType ECDSACurveType = "NISTK233"
	// NISTB233ECDSACurveType shall indicate the 'nistb233' curve in RFC5656.
	NISTB233ECDSACurveType ECDSACurveType = "NISTB233"
	// NISTK283ECDSACurveType shall indicate the 'nistk283' curve in RFC5656.
	NISTK283ECDSACurveType ECDSACurveType = "NISTK283"
	// NISTK409ECDSACurveType shall indicate the 'nistk409' curve in RFC5656.
	NISTK409ECDSACurveType ECDSACurveType = "NISTK409"
	// NISTB409ECDSACurveType shall indicate the 'nistb409' curve in RFC5656.
	NISTB409ECDSACurveType ECDSACurveType = "NISTB409"
	// NISTT571ECDSACurveType shall indicate the 'nistt571' curve in RFC5656.
	NISTT571ECDSACurveType ECDSACurveType = "NISTT571"
)

type KeyType string

const (
	// NVMeoFKeyType shall indicate the format of the key is defined by one of the
	// NVMe specifications.
	NVMeoFKeyType KeyType = "NVMeoF"
	// SSHAlgoKeyType shall indicate the format of the key is defined by one of the SSH
	// public key formats as defined in, but not limited to, RFC4253, RFC4716, or
	// RFC8709.
	SSHKeyType KeyType = "SSH"
)

type SSHAlgoKeyType string

const (
	// RSASSHAlgoKeyType shall indicate an RFC4253-defined 'ssh-rsa' key type.
	RSASSHAlgoKeyType SSHAlgoKeyType = "RSA"
	// DSASSHAlgoKeyType shall indicate an RFC4253-defined 'ssh-dss' key type.
	DSASSHAlgoKeyType SSHAlgoKeyType = "DSA"
	// ECDSASSHAlgoKeyType shall indicate an RFC5656-defined ECDSA key type.
	ECDSASSHAlgoKeyType SSHAlgoKeyType = "ECDSA"
	// Ed25519SSHAlgoKeyType shall indicate an RFC8709-defined 'ssh-ed25519' key type.
	Ed25519SSHAlgoKeyType SSHAlgoKeyType = "Ed25519"
)

// Key shall represent a key for a Redfish implementation.
type Key struct {
	Entity
	// KeyString shall contain the key, and the format shall follow the
	// requirements specified by the 'KeyType' property value.
	KeyString string
	// KeyType shall contain the format type for the key.
	KeyType KeyType
	// NVMeoF shall contain NVMe-oF specific properties for this key. This property
	// shall be present if 'KeyType' contains the value 'NVMeoF'.
	NVMeoF KeyNVMeoF
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SSH shall contain SSH specific properties for this key. This property shall
	// be present if 'KeyType' contains the value 'SSH'.
	//
	// Version added: v1.2.0
	SSH SSHType
	// UserDescription shall contain a user-provided string that describes the key.
	//
	// Version added: v1.1.0
	UserDescription string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Key object from the raw JSON.
func (k *Key) UnmarshalJSON(b []byte) error {
	type temp Key
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*k = Key(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	k.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (k *Key) Update() error {
	readWriteFields := []string{
		"UserDescription",
	}

	return k.UpdateFromRawData(k, k.RawData, readWriteFields)
}

// GetKey will get a Key instance from the service.
func GetKey(c Client, uri string) (*Key, error) {
	return GetObject[Key](c, uri)
}

// ListReferencedKeys gets the collection of Key from
// a provided reference.
func ListReferencedKeys(c Client, link string) ([]*Key, error) {
	return GetCollectionObjects[Key](c, link)
}

// KeyNVMeoF shall contain NVMe-oF specific properties for a key.
type KeyNVMeoF struct {
	// HostKeyID shall contain the value of the 'Id' property of the 'Key' resource
	// representing the host key paired with this target key. An empty string shall
	// indicate the key is not paired. This property shall be absent for host keys.
	HostKeyID string `json:"HostKeyId"`
	// NQN shall contain the NVMe Qualified Name (NQN) of the host or target
	// subsystem associated with this key. The value of this property shall follow
	// the NQN format defined by the NVMe Base Specification.
	NQN string
	// OEMSecurityProtocolType shall contain the OEM-defined security protocol that
	// this key uses. The value shall be derived from the contents of the
	// 'KeyString' property. This property shall be present if
	// 'SecurityProtocolType' contains the value 'OEM'.
	OEMSecurityProtocolType string
	// SecureHashAllowList shall contain the secure hash algorithms allowed with
	// the usage of this key. An empty list or the absence of this property shall
	// indicate any secure hash algorithms are allowed with this key.
	SecureHashAllowList []NVMeoFSecureHashType
	// SecurityProtocolType shall contain the security protocol that this key uses.
	// The value shall be derived from the contents of the 'KeyString' property.
	SecurityProtocolType NVMeoFSecurityProtocolType
}

// SSHType shall contain SSH specific properties for a key.
type SSHType struct {
	// Comment shall contain the user-specified comment associated with this key,
	// which typically contains the client's username and host name.
	//
	// Version added: v1.4.0
	Comment string
	// Fingerprint shall contain the fingerprint of the key.
	//
	// Version added: v1.2.0
	Fingerprint string
	// RemoteServerHostName shall contain the host name of the remote server
	// associated with this key.
	//
	// Version added: v1.3.0
	RemoteServerHostName string
}
