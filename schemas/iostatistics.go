//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0.5 - #IOStatistics.v1_0_4

package schemas

// IOStatistics shall be used to represent the IO statistics of the requested
// object.
type IOStatistics struct {
	// NonIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing non IO requests from the time of
	// last reset or wrap.
	NonIORequestTime string
	// NonIORequests shall represent the total count from the time of last reset or
	// wrap of non IO requests.
	NonIORequests *int `json:",omitempty"`
	// ReadHitIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from memory.
	ReadHitIORequests *int `json:",omitempty"`
	// ReadIOKiBytes shall represent the total number of kibibytes read from the
	// time of last reset or wrap.
	ReadIOKiBytes *int `json:",omitempty"`
	// ReadIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing read requests from the time of
	// last reset or wrap.
	ReadIORequestTime string
	// ReadIORequests shall represent the total count from the time of last reset
	// or wrap of read IO requests satisfied from either media or memory (i.e. from
	// a storage device or from a cache).
	ReadIORequests *int `json:",omitempty"`
	// WriteHitIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests coalesced into memory.
	WriteHitIORequests *int `json:",omitempty"`
	// WriteIOKiBytes shall represent the total number of kibibytes written from
	// the time of last reset or wrap.
	WriteIOKiBytes *int `json:",omitempty"`
	// WriteIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing write requests from the time of
	// last reset or wrap.
	WriteIORequestTime string
	// WriteIORequests shall represent the total count from the time of last reset
	// or wrap of write IO requests.
	WriteIORequests *int `json:",omitempty"`
}
