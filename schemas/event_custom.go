//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// IsValidEventType will check if it is a valid EventType.
// Should remove and leave it to the service to decide if it's valid, but since
// this is deprecated leaving it in for now.
func (e EventType) IsValidEventType() bool {
	switch e {
	case AlertEventType, ResourceAddedEventType,
		ResourceRemovedEventType, ResourceUpdatedEventType,
		StatusChangeEventType, MetricReportEventType, OtherEventType:
		return true
	}
	return false
}

var (
	SupportedEventTypes = map[string]EventType{
		"Alert":                    AlertEventType,
		"ResourceAdded":            ResourceAddedEventType,
		"ResourceRemovedEventType": ResourceRemovedEventType,
		"ResourceUpdated":          ResourceUpdatedEventType,
		"StatusChange":             StatusChangeEventType,
	}
)
