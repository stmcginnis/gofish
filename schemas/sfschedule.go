//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// DayOfWeek is Days of the Week.
type SFDayOfWeek string

const (
	// MondayDayOfWeek Monday.
	MondaySFDayOfWeek SFDayOfWeek = "Monday"
	// TuesdayDayOfWeek Tuesday.
	TuesdaySFDayOfWeek SFDayOfWeek = "Tuesday"
	// WednesdayDayOfWeek Wednesday.
	WednesdaySFDayOfWeek SFDayOfWeek = "Wednesday"
	// ThursdayDayOfWeek Thursday.
	ThursdaySFDayOfWeek SFDayOfWeek = "Thursday"
	// FridayDayOfWeek Friday.
	FridaySFDayOfWeek SFDayOfWeek = "Friday"
	// SaturdayDayOfWeek Saturday.
	SaturdaySFDayOfWeek SFDayOfWeek = "Saturday"
	// SundayDayOfWeek Sunday.
	SundaySFDayOfWeek SFDayOfWeek = "Sunday"
)

// MonthOfYear is Months of the year.
type SFMonthOfYear string

const (
	// JanuaryMonthOfYear January.
	JanuarySFMonthOfYear SFMonthOfYear = "January"
	// FebruaryMonthOfYear February.
	FebruarySFMonthOfYear SFMonthOfYear = "February"
	// MarchMonthOfYear March.
	MarchSFMonthOfYear SFMonthOfYear = "March"
	// AprilMonthOfYear April.
	AprilSFMonthOfYear SFMonthOfYear = "April"
	// MayMonthOfYear May.
	MaySFMonthOfYear SFMonthOfYear = "May"
	// JuneMonthOfYear June.
	JuneSFMonthOfYear SFMonthOfYear = "June"
	// JulyMonthOfYear July.
	JulySFMonthOfYear SFMonthOfYear = "July"
	// AugustMonthOfYear August.
	AugustSFMonthOfYear SFMonthOfYear = "August"
	// SeptemberMonthOfYear September.
	SeptemberSFMonthOfYear SFMonthOfYear = "September"
	// OctoberMonthOfYear October.
	OctoberSFMonthOfYear SFMonthOfYear = "October"
	// NovemberMonthOfYear November.
	NovemberSFMonthOfYear SFMonthOfYear = "November"
	// DecemberMonthOfYear December.
	DecemberSFMonthOfYear SFMonthOfYear = "December"
)

// Schedule shall be used to Schedule a series of occurrences.
type SFSchedule struct {
	// EnabledDaysOfMonth shall be enabled.
	EnabledDaysOfMonth []*float64
	// EnabledDaysOfWeek shall be enabled.
	EnabledDaysOfWeek []DayOfWeek
	// EnabledIntervals shall be an ISO 8601 conformant interval specifying when
	// occurrences are enabled.
	EnabledIntervals []string
	// EnabledMonthsOfYear shall be enabled.
	EnabledMonthsOfYear []MonthOfYear
	// InitialStartTime shall be an ISO 8601 conformant time of day on which the
	// initial occurrence is scheduled to occur.
	InitialStartTime string
	// Lifetime shall be an ISO 8601 conformant duration describing the time after
	// provisioning when the schedule expires.
	Lifetime string
	// MaxOccurrences Maximum number of scheduled occurrences.
	MaxOccurrences *float64 `json:",omitempty"`
	// Name is the name of the resource or array element.
	Name string
	// RecurrenceInterval shall be an ISO 8601 conformant duration describing the
	// time until the next occurrence.
	RecurrenceInterval string
}
