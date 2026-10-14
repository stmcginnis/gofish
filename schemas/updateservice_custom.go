//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

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
