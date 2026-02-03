//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// SFLocation The location of a resource. For more information see RFC4514.
type SFLocation struct {
	// Address1 Free form 1st address line for the location.
	Address1 string
	// Address2 Free form 2nd address line for the location.
	Address2 string
	// Address3 Free form 3rd address line for the location.
	Address3 string
	// Building Name of the building in which the part is installed.
	Building string
	// City Name of the town or city in which the part is installed.
	City string
	// Country The ISO 3166-1 alpha-2 ASCII country code or ISO 3166-1 numeric
	// country code of the country in which the part is installed.
	Country string
	// GPSCoords shall be expressed in the format '[-][nn]n.nnnnnn,
	// [-][nn]n.nnnnn', i.e. two numbers, either positive or negative, with six
	// decimal places of precision, comma-separated.
	GPSCoords string
	// Item Item position. If Shelf is specified, this should be the slot number
	// within the shelf, otherwise it may have a more global meaning.
	Item string
	// OtherLocationInfo Other free form text describing the item's location.
	OtherLocationInfo string
	// PostalCode Postal code (or zip code)
	PostalCode string
	// Rack Rack name or number. If Row is specified, this should be the rack
	// number within the row, otherwise it may have a more global meaning.
	Rack string
	// Room Name or number of the room in which the part is installed.
	Room string
	// Row Row name or number in which the part is installed.
	Row string
	// Shelf Shelf or unit name or number. If Rack is specified, this should be the
	// shelf number within the rack, otherwise it may have a more global meaning.
	Shelf string
	// State Name of the state in the country or territory.
	State string
	// Territory Name of the territory in the country. Not all countries use this
	// in addresses. India and China do.
	Territory string
}
