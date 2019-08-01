// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// EventService contains properties for managing event subcriptions and
// generates the events sent to subscribers.  The resource has links to the
// actual collection of subscriptions (called Event Destinations).
type EventService struct {
	common.Entity
	Status                       common.Status
	ServiceEnabled               string
	DeliveryRetryAttempts        int
	DeliveryRetryIntervalSeconds int
	EventTypesForSubscriptions   []string
	subscriptions                string
}

// UnmarshalJSON unmarshals EventService object from the raw JSON.
func (es *EventService) UnmarshalJSON(b []byte) error {
	type temp EventService
	var t struct {
		temp
		Subscriptions common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*es = EventService(t.temp)

	// Extract the links to other entities for later
	es.subscriptions = string(t.Subscriptions)

	return nil
}

// GetEventService will get a Event instance from the Redfish service.
func GetEventService(c common.Client, uri string) (*EventService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t EventService
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
