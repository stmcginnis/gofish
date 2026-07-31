//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
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
func GetSSLCert(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*SSLCert, error) {
	return GetSSLCertWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetSSLCertWithContext will get the SSLCert instance from the Redfish
// service.
func GetSSLCertWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*SSLCert, error) {
	return common.GetObjectWithContext[SSLCert](ctx, c, uri, queryOpts...)
}

// Upload will update the SSL certificate on the BMC with the provided certificate and key.
func (cert *SSLCert) Upload(certFile, keyFile io.Reader) error {
	return cert.UploadWithContext(common.ContextOf(cert.GetClient()), certFile, keyFile)
}

// UploadWithContext will update the SSL certificate on the BMC with the provided certificate and key.
func (cert *SSLCert) UploadWithContext(ctx context.Context, certFile, keyFile io.Reader) error {
	if cert.uploadTarget == "" {
		return errors.New("upload is not supported by this system")
	}

	payload := make(map[string]io.Reader)
	payload["cert_file"] = certFile
	payload["key_file"] = keyFile

	resp, err := cert.GetClient().PostMultipartWithContext(ctx, cert.uploadTarget, payload)
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
func GetIPMIConfig(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IPMIConfig, error) {
	return GetIPMIConfigWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetIPMIConfigWithContext will get the IPMIConfig instance from the Redfish
// service.
func GetIPMIConfigWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*IPMIConfig, error) {
	return common.GetObjectWithContext[IPMIConfig](ctx, c, uri, queryOpts...)
}

// Upload restores a saved IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) Upload() error {
	return ipmi.UploadWithContext(common.ContextOf(ipmi.GetClient()))
}

// UploadWithContext restores a saved IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) UploadWithContext(ctx context.Context) error {
	if ipmi.uploadTarget == "" {
		return errors.New("upload is not supported by this system")
	}

	return ipmi.PostWithContext(ctx, ipmi.uploadTarget, nil)
}

// Download saves the current IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) Download() error {
	return ipmi.DownloadWithContext(common.ContextOf(ipmi.GetClient()))
}

// DownloadWithContext saves the current IPMI configuration.
// NOTE: This is probably not correct. The jsonschema reported by SMC does not
// include any parameters for this action. That seems very unlikely, so expect
// this to fail.
func (ipmi *IPMIConfig) DownloadWithContext(ctx context.Context) error {
	if ipmi.downloadTarget == "" {
		return errors.New("download is not supported by this system")
	}

	return ipmi.PostWithContext(ctx, ipmi.downloadTarget, nil)
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
func GetUpdateService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*UpdateService, error) {
	return GetUpdateServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetUpdateServiceWithContext will get a UpdateService instance from the service.
func GetUpdateServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*UpdateService, error) {
	return common.GetObjectWithContext[UpdateService](ctx, c, uri, queryOpts...)
}

// Install performs the installation of an update.
func (us *UpdateService) Install(targets, installOptions []string) error {
	return us.InstallWithContext(common.ContextOf(us.GetClient()), targets, installOptions)
}

// InstallWithContext performs the installation of an update.
func (us *UpdateService) InstallWithContext(ctx context.Context, targets, installOptions []string) error {
	if us.installTarget == "" {
		return errors.New("install is not supported by this system")
	}

	return us.PostWithContext(ctx, us.installTarget, map[string]any{
		"Targets":        targets,
		"InstallOptions": installOptions,
	})
}

// SSLCert will get the SSLCert information from the service.
func (us *UpdateService) SSLCert(queryOpts ...common.QueryGroupOption) (*SSLCert, error) {
	return us.SSLCertWithContext(common.ContextOf(us.GetClient()), queryOpts...)
}

// SSLCertWithContext will get the SSLCert information from the service.
func (us *UpdateService) SSLCertWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*SSLCert, error) {
	return GetSSLCertWithContext(ctx, us.GetClient(), us.sslCert, queryOpts...)
}

// IPMIConfig will get the IPMIConfig information from the service.
func (us *UpdateService) IPMIConfig(queryOpts ...common.QueryGroupOption) (*IPMIConfig, error) {
	return us.IPMIConfigWithContext(common.ContextOf(us.GetClient()), queryOpts...)
}

// IPMIConfigWithContext will get the IPMIConfig information from the service.
func (us *UpdateService) IPMIConfigWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*IPMIConfig, error) {
	return GetIPMIConfigWithContext(ctx, us.GetClient(), us.ipmiConfig, queryOpts...)
}
