//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// ElectricalContext is the combination of current-carrying conductors.
type ElectricalContext string

const (
	// The circuits that share the L1 current-carrying conductor.
	Line1ElectricalContext ElectricalContext = "Line1"
	// The circuit formed by L1 and L2 current-carrying conductors.
	Line1ToLine2ElectricalContext ElectricalContext = "Line1ToLine2"
	// The circuit formed by L1 and neutral current-carrying conductors.
	Line1ToNeutralElectricalContext ElectricalContext = "Line1ToNeutral"
	// The circuit formed by L1, L2, and neutral current-carrying conductors.
	Line1ToNeutralAndL1L2ElectricalContext ElectricalContext = "Line1ToNeutralAndL1L2"
	// The circuits that share the L2 current-carrying conductor.
	Line2ElectricalContext ElectricalContext = "Line2"
	// The circuit formed by L2 and L3 current-carrying conductors.
	Line2ToLine3ElectricalContext ElectricalContext = "Line2ToLine3"
	// The circuit formed by L2 and neutral current-carrying conductors.
	Line2ToNeutralElectricalContext ElectricalContext = "Line2ToNeutral"
	// The circuit formed by L1, L2, and Neutral current-carrying conductors.
	Line2ToNeutralAndL1L2ElectricalContext ElectricalContext = "Line2ToNeutralAndL1L2"
	// The circuits formed by L2, L3, and neutral current-carrying conductors.
	Line2ToNeutralAndL2L3ElectricalContext ElectricalContext = "Line2ToNeutralAndL2L3"
	// The circuits that share the L3 current-carrying conductor.
	Line3ElectricalContext ElectricalContext = "Line3"
	// The circuit formed by L3 and L1 current-carrying conductors.
	Line3ToLine1ElectricalContext ElectricalContext = "Line3ToLine1"
	// The circuit formed by L3 and neutral current-carrying conductors.
	Line3ToNeutralElectricalContext ElectricalContext = "Line3ToNeutral"
	// The circuit formed by L3, L1, and neutral current-carrying conductors.
	Line3ToNeutralAndL3L1ElectricalContext ElectricalContext = "Line3ToNeutralAndL3L1"
	// The circuit formed by two current-carrying conductors.
	LineToLineElectricalContext ElectricalContext = "LineToLine"
	// The circuit formed by a line and neutral current-carrying conductor.
	LineToNeutralElectricalContext ElectricalContext = "LineToNeutral"
	// The grounded current-carrying return circuit of current-carrying conductors.
	NeutralElectricalContext ElectricalContext = "Neutral"
	// The circuit formed by all current-carrying conductors.
	TotalElectricalContext ElectricalContext = "Total"
)
