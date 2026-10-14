//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// This action shall be used to add a drive, or set of drives, to an underlying
// capacity source for the storage pool.
// capacitySource - This parameter shall contain the target capacity source for
// the drive(s). This property does not need to be specified if the storage
// pool only contains one capacity source, or if the implementation is capable
// of automatically selecting the appropriate capacity source.
// drives - This parameter shall contain the Uri to the existing drive or
// drives to be added to a capacity source of the storage pool. The
// implementation may impose restrictions on the number of drives added
// simultaneously.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StoragePool) AddDrives(capacitySource string, drives []string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["CapacitySource"] = capacitySource
	payload["Drives"] = drives
	resp, taskInfo, err := PostWithTask(s.client,
		s.addDrivesTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
