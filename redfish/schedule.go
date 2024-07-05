//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DayOfWeek is the days of the week.
type DayOfWeek string

const (
	// MondayDayOfWeek Monday.
	MondayDayOfWeek DayOfWeek = "Monday"
	// TuesdayDayOfWeek Tuesday.
	TuesdayDayOfWeek DayOfWeek = "Tuesday"
	// WednesdayDayOfWeek Wednesday.
	WednesdayDayOfWeek DayOfWeek = "Wednesday"
	// ThursdayDayOfWeek Thursday.
	ThursdayDayOfWeek DayOfWeek = "Thursday"
	// FridayDayOfWeek Friday.
	FridayDayOfWeek DayOfWeek = "Friday"
	// SaturdayDayOfWeek Saturday.
	SaturdayDayOfWeek DayOfWeek = "Saturday"
	// SundayDayOfWeek Sunday.
	SundayDayOfWeek DayOfWeek = "Sunday"
	// EveryDayOfWeek shall be the only member in the array.
	EveryDayOfWeek DayOfWeek = "Every"
)

// MonthOfYear is the months of the year.
type MonthOfYear string

const (
	// JanuaryMonthOfYear January.
	JanuaryMonthOfYear MonthOfYear = "January"
	// FebruaryMonthOfYear February.
	FebruaryMonthOfYear MonthOfYear = "February"
	// MarchMonthOfYear March.
	MarchMonthOfYear MonthOfYear = "March"
	// AprilMonthOfYear April.
	AprilMonthOfYear MonthOfYear = "April"
	// MayMonthOfYear May.
	MayMonthOfYear MonthOfYear = "May"
	// JuneMonthOfYear June.
	JuneMonthOfYear MonthOfYear = "June"
	// JulyMonthOfYear July.
	JulyMonthOfYear MonthOfYear = "July"
	// AugustMonthOfYear August.
	AugustMonthOfYear MonthOfYear = "August"
	// SeptemberMonthOfYear September.
	SeptemberMonthOfYear MonthOfYear = "September"
	// OctoberMonthOfYear October.
	OctoberMonthOfYear MonthOfYear = "October"
	// NovemberMonthOfYear November.
	NovemberMonthOfYear MonthOfYear = "November"
	// DecemberMonthOfYear December.
	DecemberMonthOfYear MonthOfYear = "December"
	// EveryMonthOfYear shall be the only member in the array.
	EveryMonthOfYear MonthOfYear = "Every"
)

// Schedule shall schedule a series of occurrences.
type Schedule struct {
	common.Entity
	// EnabledDaysOfMonth shall contain the days of the month when scheduled occurrences are enabled, for enabled days
	// of week and months of year. If the array contains a single value of '0', or if the property is not present, all
	// days of the month shall be enabled.
	EnabledDaysOfMonth []int
	// EnabledDaysOfWeek shall be enabled.
	EnabledDaysOfWeek []DayOfWeek
	// EnabledIntervals shall be an ISO 8601 conformant interval specifying when occurrences are enabled.
	EnabledIntervals []string
	// EnabledMonthsOfYear shall contain the months of the year when scheduled occurrences are enabled, for enabled
	// days of week and days of month. If not present, all months of the year shall be enabled.
	EnabledMonthsOfYear []MonthOfYear
	// InitialStartTime shall contain the date and time when the initial occurrence is scheduled to occur.
	InitialStartTime string
	// Lifetime shall contain a Redfish Duration that describes the time after provisioning when the schedule expires.
	Lifetime string
	// MaxOccurrences shall contain the maximum number of scheduled occurrences.
	MaxOccurrences int
	// RecurrenceInterval shall contain the duration between consecutive occurrences.
	RecurrenceInterval string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Schedule object from the raw JSON.
func (schedule *Schedule) UnmarshalJSON(b []byte) error {
	type temp Schedule
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*schedule = Schedule(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	schedule.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (schedule *Schedule) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Schedule)
	original.UnmarshalJSON(schedule.rawData)

	readWriteFields := []string{
		"EnabledDaysOfMonth",
		"EnabledDaysOfWeek",
		"EnabledIntervals",
		"EnabledMonthsOfYear",
		"InitialStartTime",
		"Lifetime",
		"MaxOccurrences",
		"RecurrenceInterval",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(schedule).Elem()

	return schedule.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSchedule will get a Schedule instance from the service.
func GetSchedule(c common.Client, uri string) (*Schedule, error) {
	return common.GetObject[Schedule](c, uri)
}

// ListReferencedSchedules gets the collection of Schedule from
// a provided reference.
func ListReferencedSchedules(c common.Client, link string) ([]*Schedule, error) {
	return common.GetCollectionObjects[Schedule](c, link)
}
