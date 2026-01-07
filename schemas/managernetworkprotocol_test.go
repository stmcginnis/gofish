//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var networkProtocolBody = `{
    "SNMP": {
        "EnableSNMPv3": false,
        "EngineId": {
            "PrivateEnterpriseId": "20 10 af 68",
            "ArchitectureId": "24 18 43 43 2F 37 5A 37 31 2D 4A 39 30 30 36 50 32 31"
        },
        "Port": 161
    },
    "SSDP": {
        "NotifyMulticastIntervalSeconds": 60,
        "NotifyTTL": 2,
        "Port": 1900,
        "NotifyIPv6Scope": "Organization",
        "ProtocolEnabled": true
    },
    "Description": "The resource is used to represent the network service settings for the manager for a Redfish implementation.",
    "Id": "NetworkProtocol",
    "Name": "Manager Network Protocol",
    "@odata.type": "#ManagerNetworkProtocol.v1_8_0.ManagerNetworkProtocol",
    "NTP": {
        "NTPServers": [
            "0.africa.pool.ntp.org"
        ],
        "ProtocolEnabled": true
    },
    "@odata.context": "/redfish/v1/$metadata#ManagerNetworkProtocol.ManagerNetworkProtocol",
    "HTTP": {
        "Port": 80,
        "ProtocolEnabled": true
    },
    "KVMIP": {
        "Port": 3900,
        "ProtocolEnabled": true
    },
    "Status": {
        "Health": "OK",
        "State": "Enabled"
    },
    "HostName": "ZCC-7372-J8205P20",
    "FQDN": "ZCC-7372-J8205P20.localdomain",
    "SSH": {
        "Port": 22,
        "ProtocolEnabled": true
    },
    "HTTPS": {
        "Port": 443,
        "ProtocolEnabled": true,
        "Certificates": {
            "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates"
        }
    },
    "Oem": {
        "Lenovo": {
            "SLP": {
                "Port": 427,
                "AddressType": "Multicast",
                "MulticastAddress": "239.255.255.253",
                "ProtocolEnabled": true
            },
            "LDAPClient": {
                "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/Oem/Lenovo/LDAPClient"
            },
            "CimOverHTTPS": {
                "Port": 5989,
                "BackendEnabled": false,
                "ProtocolEnabled": false
            },
            "@odata.type": "#LenovoManagerNetworkProtocol.v1_0_0.LenovoManagerNetworkProtocolProperties",
            "OpenPorts": [
                "22",
                "68",
                "80",
                "427",
                "443",
                "546",
                "1900",
                "3900"
            ],
            "DNS": {
                "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/Oem/Lenovo/DNS"
            },
            "SFTP": {
                "Port": 115,
                "ProtocolEnabled": false
            },
            "SMTPClient": {
                "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/Oem/Lenovo/SMTPClient"
            },
            "SNMP": {
                "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/Oem/Lenovo/SNMP"
            },
            "WebOverHTTPS": {
                "ProtocolEnabled": true
            }
        }
    },
    "@odata.etag": "\"e267adf3b95a31d8c4162\"",
    "IPMI": {
        "Port": 623,
        "ProtocolEnabled": false
    },
    "VirtualMedia": {
        "Port": 3900,
        "ProtocolEnabled": true
    },
    "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol",
    "DHCP": {
        "ProtocolEnabled": true
    },
    "DHCPv6": {
        "ProtocolEnabled": true
    }
}`

var managerNetworkProtocolBody = `{
    "@odata.type": "#ManagerNetworkProtocol.v1_5_0.ManagerNetworkProtocol",
    "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol",
    "Id": "NetworkProtocol",
    "Name": "Manager Network Protocol",
    "Description": "Manager Network Service Status",
    "Status": {
      "State": "Enabled",
      "Health": "OK"
    },
    "HostName": "",
    "FQDN": "",
    "HTTP": {
      "ProtocolEnabled": true,
      "Port": 80
    },
    "HTTPS": {
      "ProtocolEnabled": true,
      "Port": 443,
      "Certificates": {
        "@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates"
      }
    },
    "SNMP": {
      "AuthenticationProtocol": null,
      "CommunityAccessMode": null,
      "CommunityStrings": [],
      "EnableSNMPv2c": false,
      "EnableSNMPv3": false,
      "EncryptionProtocol": null,
      "EngineId": {
        "PrivateEnterpriseId": "00005345"
      },
      "HideCommunityStrings": null,
      "ProtocolEnabled": false,
      "Port": 161
    },
    "IPMI": {
      "ProtocolEnabled": true,
      "Port": 623
    },
    "SSH": {
      "ProtocolEnabled": true,
      "Port": 22
    },
    "VirtualMedia": {
      "ProtocolEnabled": true,
      "Port": 623
    },
    "KVMIP": {
      "ProtocolEnabled": true,
      "Port": 5900
    },
    "NTP": {
      "ProtocolEnabled": false,
      "NTPServers": [
        "localhost",
        "127.0.0.1"
      ],
      "NTPServers@odata.count": 2
    },
    "SSDP": {
      "ProtocolEnabled": true,
      "Port": 1900,
      "NotifyTTL": 2,
      "NotifyIPv6Scope": "Link"
    }
}`

func TestNetworkProtocol(t *testing.T) {
	var result ManagerNetworkProtocol
	err := json.NewDecoder(strings.NewReader(networkProtocolBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if !result.SSH.ProtocolEnabled {
		t.Errorf("SSH Protocol should be enabled")
	}

	if result.NTP.NTPServers[0] != "0.africa.pool.ntp.org" {
		t.Errorf("NTPServers should be set")
	}

	if result.SNMP.EngineID.PrivateEnterpriseID != "20 10 af 68" {
		t.Errorf("Wrong SNMP EngineID PrivateEnterpriseId property")
	}
}

func TestNetworkProtocolManager(t *testing.T) {
	var result ManagerNetworkProtocol
	err := json.NewDecoder(strings.NewReader(managerNetworkProtocolBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if !result.SSH.ProtocolEnabled {
		t.Error("SSH Protocol should be enabled")
	}

	if len(result.NTP.NTPServers) != 2 {
		t.Error("Should be 2 NTP servers defined")
	}

	if result.NTP.NTPServers[1] != "127.0.0.1" {
		t.Error("NTPServers should be set")
	}

	if result.SNMP.EngineID.PrivateEnterpriseID != "00005345" {
		t.Error("Wrong SNMP EngineID PrivateEnterpriseId property")
	}

	if result.HTTPS.certificates != "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates" {
		t.Errorf("Wrong HTTPS certificates link: %s", result.HTTPS.certificates)
	}
}
