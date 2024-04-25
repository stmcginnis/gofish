//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var securityPolicyBody = `{
	"@odata.type": "#SecurityPolicy.v1_0_1.SecurityPolicy",
	"Id": "ManagerGlobalSecurityPolicy",
	"Status": {
	  "Health": "OK",
	  "State": "Enabled"
	},
	"OverrideParentManager": true,
	"SPDM": {
	  "Enabled": true,
	  "SecureSessionEnabled": true,
	  "VerifyCertificate": true,
	  "TrustedCertificates": {
		"@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/SPDM/TrustedCertificates"
	  },
	  "RevokedCertificates": {
		"@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/SPDM/RevokedCertificates"
	  },
	  "Allowed": {
		"Versions": [
		  "ALL"
		],
		"Algorithms": {
		  "AEAD": [
			"AES-GCM-256",
			"AES-GCM-128"
		  ],
		  "BaseAsym": [
			"TPM_ALG_RSASSA_2048",
			"TPM_ALG_ECDSA_ECC_NIST_P384",
			"TPM_ALG_SM2_ECC_SM2_P256"
		  ],
		  "BaseHash": [
			"TPM_ALG_SHA_512",
			"TPM_ALG_SHA3_512"
		  ]
		}
	  },
	  "Denied": {
		"Versions": [
		  "NONE"
		],
		"Algorithms": {
		  "AEAD": [],
		  "BaseAsym": [
			"EdDSA ed25519"
		  ],
		  "BaseHash": [
			"TPM_ALG_SHA_256"
		  ]
		}
	  },
	  "AllowExtendedAlgorithms": false
	},
	"TLS": {
	  "Client": {
		"VerifyCertificate": true,
		"TrustedCertificates": {
		  "@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Server/TrustedCertificates"
		},
		"RevokedCertificates": {
		  "@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Server/RevokedCertificates"
		},
		"Allowed": {
		  "Versions": [
			"1.2",
			"1.3"
		  ],
		  "Algorithms": {
			"CipherSuites": [
			  "TLS_AES_128_GCM_SHA256",
			  "TLS_AES_128_GCM_SHA384"
			],
			"SignatureAlgorithms": []
		  }
		},
		"Denied": {
		  "Versions": [
			"1.1"
		  ],
		  "Algorithms": {
			"CipherSuites": [],
			"SignatureAlgorithms": [
			  "rsa_pkcs1_sha1",
			  "ecdsa_sha1"
			]
		  }
		}
	  },
	  "Server": {
		"VerifyCertificate": false,
		"TrustedCertificates": {
		  "@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Client/TrustedCertificates"
		},
		"RevokedCertificates": {
		  "@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Client/RevokedCertificates"
		},
		"Allowed": {
		  "Versions": [
			"1.3"
		  ],
		  "Algorithms": {
			"CipherSuites": [
			  "TLS_AES_128_GCM_SHA256",
			  "TLS_AES_128_GCM_SHA384"
			],
			"SignatureAlgorithms": []
		  }
		},
		"Denied": {
		  "Versions": [
			"1.1",
			"1.2"
		  ],
		  "Algorithms": {
			"CipherSuites": [],
			"SignatureAlgorithms": [
			  "rsa_pkcs1_sha1",
			  "ecdsa_sha1"
			]
		  }
		}
	  }
	},
	"@odata.id": "/redfish/v1/Managers/BMC/SecurityPolicy"
  }`

// TestSecurityPolicy tests the parsing of SecurityPolicy objects.
func TestSecurityPolicy(t *testing.T) {
	var result SecurityPolicy
	err := json.NewDecoder(strings.NewReader(securityPolicyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ManagerGlobalSecurityPolicy", result.ID)

	if !result.OverrideParentManager {
		t.Error("Expected OverrideParentManager to be true")
	}

	if !result.SPDM.Enabled {
		t.Error("Expected SPDM.Enabled to be true")
	}

	if !result.SPDM.SecureSessionEnabled {
		t.Error("Expected SPDM.SecureSessionEnabled to be true")
	}

	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/SPDM/TrustedCertificates", result.SPDM.trustedCertificates)
	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/SPDM/RevokedCertificates", result.SPDM.revokedCertificates)
	assertEquals(t, "ALL", result.SPDM.Allowed.Versions[0])
	assertEquals(t, "TPM_ALG_SHA3_512", result.SPDM.Allowed.Algorithms.BaseHash[1])
	assertEquals(t, "EdDSA ed25519", result.SPDM.Denied.Algorithms.BaseAsym[0])
	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Server/TrustedCertificates", result.TLS.Client.trustedCertificates)
	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Server/RevokedCertificates", result.TLS.Client.revokedCertificates)
	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Client/TrustedCertificates", result.TLS.Server.trustedCertificates)
	assertEquals(t, "/redfish/v1/Managers/BMC/SecurityPolicy/TLS/Client/RevokedCertificates", result.TLS.Server.revokedCertificates)
}
