//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/redfish"
)

// AccountService is a Supermicro OEM instance of an AccountService.
type AccountService struct {
	redfish.AccountService
	SMCLDAP struct {
		StartTLSEnabled bool
	}
	SMCActiveDirectory struct {
		DNSLookupEnable        bool
		Prefix                 string
		Port                   int
		UserDomainNames        []string
		DynamicServerAddresses []string
	}
}

// FromAccountService converts a standard AccountService object to the OEM implementation.
func FromAccountService(accountService *redfish.AccountService) (*AccountService, error) {
	as := AccountService{
		AccountService: *accountService,
	}

	var t struct {
		Oem struct {
			Supermicro struct {
				LDAP struct {
					StartTLSEnabled bool
				} `json:"LDAP"`
				ActiveDirectory struct {
					DNSLookupEnable        bool
					Prefix                 string
					Port                   int
					UserDomainNames        []string
					DynamicServerAddresses []string
				} `json:"ActiveDirectory"`
			} `json:"Supermicro"`
		} `json:"Oem"`
	}

	err := json.Unmarshal(accountService.RawData, &t)
	if err != nil {
		return nil, err
	}

	as.SMCLDAP.StartTLSEnabled = t.Oem.Supermicro.LDAP.StartTLSEnabled
	as.SMCActiveDirectory.DNSLookupEnable = t.Oem.Supermicro.ActiveDirectory.DNSLookupEnable
	as.SMCActiveDirectory.Prefix = t.Oem.Supermicro.ActiveDirectory.Prefix
	as.SMCActiveDirectory.Port = t.Oem.Supermicro.ActiveDirectory.Port
	as.SMCActiveDirectory.UserDomainNames = t.Oem.Supermicro.ActiveDirectory.UserDomainNames
	as.SMCActiveDirectory.DynamicServerAddresses = t.Oem.Supermicro.ActiveDirectory.DynamicServerAddresses

	return &as, nil
}
