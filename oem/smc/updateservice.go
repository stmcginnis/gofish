//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type SSLCert struct {
	common.Entity

	// GoodThrough is the certificate expiration date.
	GoodThrough string `json:"GoodTHRU"`
	// ValidFrom is the certificate start date. It's misspelled as VaildFrom in the schema.
	ValidFrom string `json:"VaildFrom"`

	// uploadTarget is the URL to upload certificates to.
	uploadTarget string
}

// UnmarshalJSON unmarshals a SSLCert object from the raw JSON.
func (cert *SSLCert) UnmarshalJSON(b []byte) error {
	type temp SSLCert
	var t struct {
		temp
		Actions struct {
			Upload common.ActionTarget `json:"#SmcSSLCert.Upload"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cert = SSLCert(t.temp)
	cert.uploadTarget = t.Actions.Upload.Target

	return nil
}

// GetSSLCert will get the SSLCert instance from the Redfish
// service.
func GetSSLCert(c common.Client, uri string) (*SSLCert, error) {
	return common.GetObject[SSLCert](c, uri)
}

// Upload will update the SSL certificate on the BMC with the provided certificate and key.
func (cert *SSLCert) Upload(certFile, keyFile io.Reader) error {
	if cert.uploadTarget == "" {
		return errors.New("upload is not supported by this system")
	}

	payload := make(map[string]io.Reader)
	payload["cert_file"] = certFile
	payload["key_file"] = keyFile

	resp, err := cert.GetClient().PostMultipart(cert.uploadTarget, payload)
	if err != nil {
		return err
	}

	return common.CleanupHTTPResponse(resp)
}

type IPMIConfig struct {
	common.Entity

	uploadTarget   string
	downloadTarget string
}

// UnmarshalJSON unmarshals a UpdateService object from the raw JSON.
func (ipmi *IPMIConfig) UnmarshalJSON(b []byte) error {
	type temp IPMIConfig
	var t struct {
		temp
		Actions struct {
			Upload   common.ActionTarget `json:"#SmcIPMIConfig.Upload"`
			Download common.ActionTarget `json:"#SmcIPMIConfig.Download"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ipmi = IPMIConfig(t.temp)
	ipmi.uploadTarget = t.Actions.Upload.Target
	ipmi.downloadTarget = t.Actions.Download.Target

	return nil
}

// GetIPMIConfig will get the IPMIConfig instance from the Redfish
// service.
func GetIPMIConfig(c common.Client, uri string) (*IPMIConfig, error) {
	return common.GetObject[IPMIConfig](c, uri)
}

// Upload restores a saved IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) Upload() error {
	if ipmi.uploadTarget == "" {
		return errors.New("upload is not supported by this system")
	}

	return ipmi.Post(ipmi.uploadTarget, nil)
}

// Download saves the current IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) Download() error {
	if ipmi.downloadTarget == "" {
		return errors.New("download is not supported by this system")
	}

	return ipmi.Post(ipmi.downloadTarget, nil)
}

// UpdateService is the update service instance associated with the system.
type UpdateService struct {
	redfish.UpdateService

	sslCert    string
	ipmiConfig string

	installTarget string
}

// FromUpdateService gets the OEM instance of the UpdateService.
func FromUpdateService(updateService *redfish.UpdateService) (*UpdateService, error) {
	us := UpdateService{
		UpdateService: *updateService,
	}

	var t struct {
		Actions struct {
			Oem struct {
				Install common.ActionTarget `json:"#SmcUpdateService.Install"`
			}
		}
		Oem struct {
			Supermicro struct {
				SSLCert    common.Link
				IPMIConfig common.Link
			}
		}
	}

	err := json.Unmarshal(updateService.RawData, &t)
	if err != nil {
		return nil, err
	}

	us.sslCert = t.Oem.Supermicro.SSLCert.String()
	us.ipmiConfig = t.Oem.Supermicro.IPMIConfig.String()

	us.installTarget = t.Actions.Oem.Install.Target

	return &us, nil
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(c common.Client, uri string) (*UpdateService, error) {
	return common.GetObject[UpdateService](c, uri)
}

// Install performs the installation of an update.
func (us *UpdateService) Install(targets, installOptions []string) error {
	if us.installTarget == "" {
		return errors.New("install is not supported by this system")
	}

	return us.Post(us.installTarget, map[string]any{
		"Targets":        targets,
		"InstallOptions": installOptions,
	})
}

// SSLCert will get the SSLCert information from the service.
func (us *UpdateService) SSLCert() (*SSLCert, error) {
	return GetSSLCert(us.GetClient(), us.sslCert)
}

// IPMIConfig will get the IPMIConfig information from the service.
func (us *UpdateService) IPMIConfig() (*IPMIConfig, error) {
	return GetIPMIConfig(us.GetClient(), us.ipmiConfig)
}
