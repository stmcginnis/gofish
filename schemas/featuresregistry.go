//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/FeaturesRegistry.v1_2_1.json
// 1.2.7 - #FeaturesRegistry.v1_2_1.FeaturesRegistry

package schemas

import (
	"encoding/json"
)

// FeaturesRegistry shall be used to represent a Feature registry for a Redfish
// implementation.
type FeaturesRegistry struct {
	Entity
	// FeatureMappings shall represent the suffix to be used in the FeatureId and
	// shall be unique within this message registry. This may contain both standard
	// and OEM-defined features.
	//
	// Version added: v1.2.0
	FeatureMappings []FeatureMap
	// Features shall represent the suffix to be used in the FeatureId and shall be
	// unique within this message registry.
	//
	// Deprecated: v1.2.0
	// This property is deprecated in favor of the FeatureMappings property.
	Features []SupportedFeature
	// FeaturesUsed shall contain an array of all the standard feature names
	// defined in the registry. OEM feature names shall be defined on the
	// OEMFeaturesUsed property.
	//
	// Version added: v1.2.0
	FeaturesUsed []string
	// Language shall be a string consisting of an RFC 5646 language code.
	Language string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMFeaturesUsed shall contain an array of all the OEM feature names defined
	// in the registry.
	//
	// Version added: v1.2.0
	OEMFeaturesUsed []string
	// OwningEntity shall be a string that represents the publisher of this
	// registry.
	OwningEntity string
	// RegistryPrefix shall be the prefix used in IDs which uniquely identifies all
	// of the Features in this registry as belonging to this registry.
	RegistryPrefix string
	// RegistryVersion shall be the version of this message registry. The format of
	// this string shall be of the format majorversion.minorversion.errata.
	RegistryVersion string
}

// GetFeaturesRegistry will get a FeaturesRegistry instance from the service.
func GetFeaturesRegistry(c Client, uri string) (*FeaturesRegistry, error) {
	return GetObject[FeaturesRegistry](c, uri)
}

// ListReferencedFeaturesRegistrys gets the collection of FeaturesRegistry from
// a provided reference.
func ListReferencedFeaturesRegistrys(c Client, link string) ([]*FeaturesRegistry, error) {
	return GetCollectionObjects[FeaturesRegistry](c, link)
}

// FeatureMap shall contain the defined name of a feature and the corresponding
// resources for which the feature applies. For globally applied features,
// implementations should not implement the 'Resources' property for the
// specified feature.
type FeatureMap struct {
	// CorrespondingProfileDefinition shall define a profile definition that
	// contains the named profile declaration.
	//
	// Version added: v1.2.0
	CorrespondingProfileDefinition string
	// Description provides a description of this resource.
	//
	// Version added: v1.2.0
	Description string
	// FeatureName shall be the unique name of the feature prefixed by the defining
	// organization separated by a period (e.g. 'vendor.feature').
	//
	// Version added: v1.2.0
	FeatureName string
	// Resources shall contain an array of resources in the service containing the
	// set that are related to support for the specified feature. For globally
	// applied features, implementations should not implement the 'Resources'
	// property for the specified feature.
	//
	// Version added: v1.2.0
	resources []string
	// ResourcesCount
	ResourcesCount int `json:"Resources@odata.count"`
	// Version shall uniquely identify the version of the feature, using the
	// major.minor.errata format.
	//
	// Version added: v1.2.0
	Version string
}

// UnmarshalJSON unmarshals a FeatureMap object from the raw JSON.
func (f *FeatureMap) UnmarshalJSON(b []byte) error {
	type temp FeatureMap
	var tmp struct {
		temp
		Resources Links `json:"Resources"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FeatureMap(tmp.temp)

	// Extract the links to other entities for later
	f.resources = tmp.Resources.ToStrings()

	return nil
}

// Resources gets the Resources linked resources.
func (f *FeatureMap) Resources(client Client) ([]*Resource, error) {
	return GetObjects[Resource](client, f.resources)
}

// FeaturesRegistryProperty shall represent the suffix to be used in the Feature
// and shall be unique within this registry.
type FeaturesRegistryProperty struct {
}

// SupportedFeature shall name a feature.
type SupportedFeature struct {
	// CorrespondingProfileDefinition shall define a profile definition that
	// contains the named profile declaration.
	CorrespondingProfileDefinition string
	// Description provides a description of this resource.
	Description string
	// FeatureName shall be the unique name of the feature prefixed by the defining
	// organization separated by a period (e.g. 'vendor.feature').
	FeatureName string
	// Version shall uniquely identify the version of the feature, using the
	// major.minor.errata format.
	Version string
}
