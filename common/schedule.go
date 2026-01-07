//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// DayOfWeek is Days of the Week.
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

// MonthOfYear is Months of the year.
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

// Schedule shall be used to Schedule a series of occurrences.
type Schedule struct {
	// EnabledDaysOfMonth shall contain the days of the month when scheduled
	// occurrences are enabled, for enabled days of week and months of year. If the
	// array contains a single value of '0', or if the property is not present, all
	// days of the month shall be enabled.
	EnabledDaysOfMonth []*int
	// EnabledDaysOfWeek shall be enabled.
	EnabledDaysOfWeek []DayOfWeek
	// EnabledIntervals shall be an ISO 8601 conformant interval specifying when
	// occurrences are enabled.
	//
	// Version added: v1.1.0
	EnabledIntervals []string
	// EnabledMonthsOfYear shall contain the months of the year when scheduled
	// occurrences are enabled, for enabled days of week and days of month. If not
	// present, all months of the year shall be enabled.
	EnabledMonthsOfYear []MonthOfYear
	// InitialStartTime shall be an ISO 8601 conformant time of day on which the
	// initial occurrence is scheduled to occur.
	InitialStartTime string
	// Lifetime shall be an ISO 8601 conformant duration describing the time after
	// provisioning when the schedule expires.
	Lifetime string
	// MaxOccurrences Maximum number of scheduled occurrences.
	MaxOccurrences *int `json:",omitempty"`
	// Name is the name of the resource or array element.
	Name string
	// RecurrenceInterval shall be an ISO 8601 conformant duration describing the
	// time until the next occurrence.
	RecurrenceInterval string
}
