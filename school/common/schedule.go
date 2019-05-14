// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// Schedule is used to Schedule a series of occurrences.
type Schedule struct {
	// EnabledDaysOfMonth is Days of month when scheduled occurrences are
	// enabled, for enabled days of week and months of year. If the array
	// contains a single value of zero, or if the property is not present,
	// all days of the month shall be enabled.
	EnabledDaysOfMonth []string
	// EnabledDaysOfWeek is Days of the week when scheduled occurrences are
	// enabled. If not present, all days of the week shall be enabled.
	EnabledDaysOfWeek []DayOfWeek
	// EnabledIntervals shall be an ISO 8601 conformant interval specifying when
	// occurences are enabled.
	EnabledIntervals []string
	// EnabledMonthsOfYear is Months of year when scheduled occurrences are
	// enabled, for enabled days of week and days of month. If not present,
	// all months of the year shall be enabled.
	EnabledMonthsOfYear []MonthOfYear
	// InitialStartTime shall be a date and time of day on which the initial
	// occurrence is scheduled to occur.
	InitialStartTime string
	// Lifetime shall be a Redfish Duration describing the time after
	// provisioning when the schedule expires.
	Lifetime string
	// MaxOccurrences is Maximum number of scheduled occurrences.
	MaxOccurrences int
	// RecurrenceInterval shall be a Redfish Duration describing the time until
	// the next occurrence.
	RecurrenceInterval string
}
