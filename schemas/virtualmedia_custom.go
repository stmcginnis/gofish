//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
)

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (v *VirtualMedia) UnmarshalJSON(b []byte) error {
	type temp VirtualMedia
	type vActions struct {
		EjectMedia  ActionTarget `json:"#VirtualMedia.EjectMedia"`
		InsertMedia ActionTarget `json:"#VirtualMedia.InsertMedia"`
	}
	var tmp struct {
		temp
		Actions            vActions
		Certificates       Link `json:"Certificates"`
		ClientCertificates Link `json:"ClientCertificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VirtualMedia(tmp.temp)

	// Extract the links to other entities for later
	v.certificates = tmp.Certificates.String()
	v.clientCertificates = tmp.ClientCertificates.String()

	v.ejectMedia = tmp.Actions.EjectMedia
	v.insertMedia = tmp.Actions.InsertMedia

	v.SupportsMediaEject = v.ejectMedia.Target != ""
	v.SupportsMediaInsert = v.insertMedia.Target != ""

	// This is a read/write object, so we need to save the raw object data for later
	v.RawData = b

	return nil
}

// EjectMediaActionInfo provides the ActionInfo, if supported, for an EjectMedia Action
func (v *VirtualMedia) EjectMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia EjectMedia ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// This action shall detach the remote media from the virtual media. At the
// completion of the operation, inserted shall be set to 'false' and the image
// name shall be cleared.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (v *VirtualMedia) EjectMedia() (*TaskMonitorInfo, error) {
	if !v.SupportsMediaEject {
		// Some implementations (e.g. Lenovo XCC) do not expose the standard
		// VirtualMedia.EjectMedia action; fall back to PATCHing the resource.
		payload := struct {
			Inserted       bool
			WriteProtected bool
		}{}
		return nil, v.Patch(v.ODataID, payload)
	}

	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(v.client,
		v.ejectMedia.Target, payload, v.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// InsertMediaActionInfo provides the ActionInfo, if supported, for an InsertMedia Action
func (v *VirtualMedia) InsertMediaActionInfo() (*ActionInfo, error) {
	if v.insertMedia.ActionInfoTarget == "" {
		return nil, errors.New("VirtualMedia InsertMedia ActionInfo not supported by this service")
	}

	return GetObject[ActionInfo](v.GetClient(), v.insertMedia.ActionInfoTarget)
}

// This action shall attach remote media to virtual media. Service should
// reject the request if the 'Image' property in the resource does not contain
// 'null'. Users are expected to eject media prior to inserting new media.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (v *VirtualMedia) InsertMedia(params *VirtualMediaInsertMediaParameters) (*TaskMonitorInfo, error) {
	if !v.SupportsMediaInsert {
		// Some implementations (e.g. Lenovo XCC) do not expose the standard
		// VirtualMedia.InsertMedia action; fall back to PATCHing the resource.
		return nil, v.Patch(v.ODataID, params)
	}

	resp, taskInfo, err := PostWithTask(v.client,
		v.insertMedia.Target, params, v.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
