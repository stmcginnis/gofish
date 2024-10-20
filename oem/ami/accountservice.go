//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// PAMOrder is the PAM modules used for authentication.
type PAMOrder string

const (
	// IPMIPAMOrder specifies IPMI authentication.
	IPMIPAMOrder PAMOrder = "IPMI"
	// LDAPPAMOrder specifies LDAP authentication.
	LDAPPAMOrder PAMOrder = "LDAP"
	// ACTIVEDIRECTORYPAMOrder specifies ACTIVE DIRECTORY authentication.
	ACTIVEDIRECTORYPAMOrder PAMOrder = "ACTIVE DIRECTORY"
	// RADIUSPAMOrder specifies RADIUS authentication.
	RADIUSPAMOrder PAMOrder = "RADIUS"
)

// AccountServiceConfigurations allows additional configuring of the AMI AccountService.
type AccountServiceConfigurations struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// PAMEnabled indicates whether or not PAM authentication should be used when authenticating Redfish requests.
	PAMEnabled bool
	// PAMOrder is an array that represents the order the PAM modules will be checked for authentication.
	PAMOrder []PAMOrder
}

// GetAccountServiceConfigurations will get an AccountServiceConfigurations instance from the Redfish
// service.
func GetAccountServiceConfigurations(c common.Client, uri string) (*AccountServiceConfigurations, error) {
	return common.GetObject[AccountServiceConfigurations](c, uri)
}

// AccountService is an AMI OEM instance of an AccountService.
type AccountService struct {
	redfish.AccountService

	configuration string
}

// FromAccountService converts a standard AccountService object to the OEM implementation.
func FromAccountService(accountService *redfish.AccountService) (*AccountService, error) {
	as := AccountService{
		AccountService: *accountService,
	}

	var t struct {
		Oem struct {
			AMI struct {
				Configurtion common.Link `json:"Configuration"`
			} `json:"AMI"`
		} `json:"Oem"`
	}

	err := json.Unmarshal(accountService.RawData, &t)
	if err != nil {
		return nil, err
	}

	as.configuration = t.Oem.AMI.Configurtion.String()
	as.SetClient(accountService.GetClient())

	return &as, nil
}

// Configuration will get the AccountServiceConfigurations for this AccountService.
func (as *AccountService) Configuration() (*AccountServiceConfigurations, error) {
	return GetAccountServiceConfigurations(as.GetClient(), as.configuration)
}
