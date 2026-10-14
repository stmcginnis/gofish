//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// This action shall control the power state of the outlet group.
// powerState - This parameter shall contain the desired power state of the
// outlet group.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (o *OutletGroup) PowerControl(powerState OutletPowerState) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PowerState"] = powerState
	resp, taskInfo, err := PostWithTask(o.client,
		o.powerControlTarget, payload, o.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
