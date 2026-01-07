//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/stmcginnis/gofish/schemas"
)

type ApplyUpdate string

const (
	ApplyUpdateTrue  ApplyUpdate = "True"
	ApplyUpdateFalse ApplyUpdate = "False"
)

type IgnoreCertWarning string

const (
	IgnoreCertWarningOn  IgnoreCertWarning = "On"
	IgnoreCertWarningOff IgnoreCertWarning = "Off"
)

type ProxyType string

const (
	ProxyTypeHTTP  ProxyType = "HTTP"
	ProxyTypeSOCKS ProxyType = "SOCKS"
)

type ProxySupport string

const (
	ProxySupportDefault         ProxySupport = "DefaultProxy"
	ProxySupportOff             ProxySupport = "Off"
	ProxySupportParametersProxy ProxySupport = "ParametersProxy"
)

type IFRShareType string

const (
	LocalIFRShareType IFRShareType = "LOCAL"
	NFSIFRShareType   IFRShareType = "NFS"
	CIFSIFRShareType  IFRShareType = "CIFS"
	HTTPIFRShareType  IFRShareType = "HTTP"
	HTTPSIFRShareType IFRShareType = "HTTPS"
)

type InstallFromRepoBody struct {
	// If ApplyUpdate is set to True, the updatable packages from Catalog XML are staged. If it is set to False, no updates are applied. The list of updatable packages can be seen by invoking the GetRepoBasedUpdateList. Default value is True.
	ApplyUpdate ApplyUpdate `json:",omitempty"`
	// Name of the catalog file on the repository. Default is Catalog.xml.
	CatalogFile string `json:",omitempty"`
	// IP address for the remote share.
	IPAddress string
	// Specifies if certificate warning should be ignored when HTTPS is used. If IgnoreCertWarning is On, warnings are ignored. Default is 2 (On).
	IgnoreCertWarning IgnoreCertWarning `json:",omitempty"`
	// The local directory where the share should be mounted. This is applicable for CIFS.
	MountPoint string `json:",omitempty"`
	// Password for the remote share. This parameter must be provided for CIFS.
	Password string `json:",omitempty"`
	// The password for the proxy server.
	ProxyPasswd string `json:",omitempty"`
	// Port for the proxy server. Default is set to 80.
	ProxyPort int `json:",omitempty"`
	// The IP address of the proxy server.
	ProxyServer string `json:",omitempty"`
	// Specifies if a proxy should be used. Default is Off.
	ProxySupport ProxySupport `json:",omitempty"`
	// The proxy type of the proxy server. Default is HTTP.
	ProxyType ProxyType `json:",omitempty"`
	// The user name for the proxyserver.
	ProxyUname string `json:",omitempty"`
	// This property indicates if a reboot should be performed. True indicates that the system (host) is rebooted during the update process. False indicates that the updates take effect after the system is rebooted the next time. Default value is set to False.
	RebootNeeded bool `json:",omitempty"`
	// Name of the CIFS share or full path to the NFS share. Optional for HTTP/HTTPS share, this may be treated as the path of the directory containing the file.
	ShareName string `json:",omitempty"`
	// Type of the network share. Default value is NFS.
	ShareType IFRShareType
	// User name for the remote share. This parameter must be provided for CIFS.
	UserName string `json:",omitempty"`
	// Workgroup for the CIFS share - optional.
	Workgroup string `json:",omitempty"`
}

type UpdateList []UpdateListFirmware
type UpdateListFirmware struct {
	// Name of package
	Name string
	// Importance of update:
	// 	- "1" = Warning
	// 	- "2" = Critical
	// 	- "3" = Recommended
	Criticality string
	// Display name of component
	DisplayName string
	// Prefix of upgrade path
	BaseLocation string
	// Path to upgrade file (includes PackageName)
	PackagePath string
	// Name of firmware package
	PackageName string
	// Version of firmware in catalog
	PackageVersion string
	// Reboot type required by package:
	// 	- "NONE" = No reboot needed
	//	- "IDRAC" = iDRAC reboot required (non disruptive)
	// 	- "HOST" = full host power cycle needed
	RebootType string
	JobID      string
	// FQDD of component targeted by update
	Target        string
	ComponentID   string
	ComponentType string
	// Semicolon separated list of component ID types
	ComponentInfoName string
	// Semicolon separated list of component IDs
	ComponentInfoValue string
	// Currently installed firmware version
	InstalledVersion string
}

type SoftwareInstallationService struct {
	schemas.Entity

	Actions struct {
		InstallFromRepository struct {
			Target string `json:"target"`
		} `json:"#DellSoftwareInstallationService.InstallFromRepository"`
		GetRepoBasedUpdateList struct {
			Target string `json:"target"`
		} `json:"#DellSoftwareInstallationService.GetRepoBasedUpdateList"`
	}
}

type xmlUpdateList struct {
	Packages []struct {
		PackageAttributes []struct {
			Name  string `xml:"NAME,attr"`
			Value string `xml:"VALUE"`
		} `xml:"PROPERTY"`
		PackageProperties []struct {
			Name  string `xml:"NAME,attr"`
			Value string `xml:"VALUE.ARRAY>VALUE"`
		} `xml:"PROPERTY.ARRAY"`
	} `xml:"MESSAGE>SIMPLEREQ>VALUE.NAMEDINSTANCE>INSTANCENAME"`
}

func (xul *xmlUpdateList) parseFromXML() UpdateList {
	var updateList UpdateList

	for _, p := range xul.Packages {
		var f UpdateListFirmware

		for _, pa := range p.PackageAttributes {
			switch pa.Name {
			case "Criticality":
				f.Criticality = pa.Value
			case "DisplayName":
				f.DisplayName = pa.Value
			case "BaseLocation":
				f.BaseLocation = pa.Value
			case "PackagePath":
				f.PackagePath = pa.Value
			case "PackageName":
				f.PackageName = pa.Value
			case "PackageVersion":
				f.PackageVersion = pa.Value
			case "RebootType":
				f.RebootType = pa.Value
			case "JobID":
				f.JobID = pa.Value
			case "Target":
				f.Target = pa.Value
			case "ComponentID":
				f.ComponentID = pa.Value
			case "ComponentType":
				f.ComponentType = pa.Value
			}
		}

		for _, pp := range p.PackageProperties {
			switch pp.Name {
			case "ComponentInfoName":
				f.ComponentInfoName = pp.Value
			case "ComponentInfoValue":
				f.ComponentInfoValue = pp.Value
			case "ComponentInstalledVersion":
				f.InstalledVersion = pp.Value
			}
		}

		updateList = append(updateList, f)
	}

	return updateList
}

// validateInstallFromRepoBody validates required fields in InstallFromRepoBody
func validateInstallFromRepoBody(b *InstallFromRepoBody) error {
	if b.IPAddress == "" {
		return errors.New("IPAddress is required")
	}
	if b.ShareName == "" {
		return errors.New("ShareName is required")
	}
	if b.ShareType == "" {
		return errors.New("ShareType is required")
	}

	// Validate ShareType is one of the allowed values
	validShareTypes := []IFRShareType{
		LocalIFRShareType,
		NFSIFRShareType,
		CIFSIFRShareType,
		HTTPIFRShareType,
		HTTPSIFRShareType,
	}

	valid := false
	for _, validType := range validShareTypes {
		if b.ShareType == validType {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid ShareType: %s, must be one of LOCAL, NFS, CIFS, HTTP, HTTPS", b.ShareType)
	}

	// Username is required for CIFS shares
	if b.ShareType == CIFSIFRShareType && b.UserName == "" {
		return errors.New("UserName is required for CIFS shares")
	}

	return nil
}

// Simple way to upgrade server firmware packages. Uses a Dell update catalog to compare FW versions and get download links for each package.
//
// Returns a Dell OEM Job
func (sis *SoftwareInstallationService) InstallFromRepository(b *InstallFromRepoBody) (*Job, error) {
	if err := validateInstallFromRepoBody(b); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	res, err := sis.PostWithResponse(sis.Actions.InstallFromRepository.Target, b)
	defer schemas.DeferredCleanupHTTPResponse(res)
	if err != nil {
		return nil, err
	}

	return GetJob(sis.GetClient(), res.Header.Get("Location"))
}

// Queries BMC for package list of available updates
//
// Must be called after "InstallFromRepository" with ApplyUpdate = False.
// To install the firmware, call "InstallFromRepository" again with ApplyUpdate = True
//
// On success, returns a struct with the firmware upgrade details
// On failure to get the catalog OR if all firmware is current, returns a schemas.Error error with an extended schemas.error message.
func (sis *SoftwareInstallationService) GetRepoBasedUpdateList() (*UpdateList, error) {
	var b struct{}
	res, err := sis.PostWithResponse(sis.Actions.GetRepoBasedUpdateList.Target, b)
	defer schemas.DeferredCleanupHTTPResponse(res)
	if err != nil {
		return nil, err
	}

	var pl struct {
		PackageList string
	}

	if err := json.NewDecoder(res.Body).Decode(&pl); err != nil {
		return nil, err
	}

	var t xmlUpdateList

	err = xml.Unmarshal([]byte(pl.PackageList), &t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML package list: %w", err)
	}

	ul := t.parseFromXML()

	// Validate that we got some packages
	if len(ul) == 0 {
		return nil, errors.New("no firmware packages found in update list - catalog may be empty or invalid")
	}

	return &ul, nil
}

func GetSoftwareInstallationService(c schemas.Client, uri string) (*SoftwareInstallationService, error) {
	return schemas.GetObject[SoftwareInstallationService](c, uri)
}
