//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/CertificateLocations.v1_0_4.json
// 2018.3 - #CertificateLocations.v1_0_4.CertificateLocations

package schemas

import (
	"encoding/json"
)

// CertificateLocations shall represent the certificate location properties for
// a Redfish implementation.
type CertificateLocations struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// certificates are the URIs for Certificates.
	certificates []string
}

// UnmarshalJSON unmarshals a CertificateLocations object from the raw JSON.
func (c *CertificateLocations) UnmarshalJSON(b []byte) error {
	type temp CertificateLocations
	type cLinks struct {
		Certificates Links `json:"Certificates"`
	}
	var tmp struct {
		temp
		Links cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CertificateLocations(tmp.temp)

	// Extract the links to other entities for later
	c.certificates = tmp.Links.Certificates.ToStrings()

	return nil
}

// GetCertificateLocations will get a CertificateLocations instance from the service.
func GetCertificateLocations(c Client, uri string) (*CertificateLocations, error) {
	return GetObject[CertificateLocations](c, uri)
}

// ListReferencedCertificateLocations gets the collection of CertificateLocations from
// a provided reference.
func ListReferencedCertificateLocations(c Client, link string) ([]*CertificateLocations, error) {
	return GetCollectionObjects[CertificateLocations](c, link)
}

// Certificates gets the Certificates linked resources.
func (c *CertificateLocations) Certificates() ([]*Certificate, error) {
	return GetObjects[Certificate](c.client, c.certificates)
}
