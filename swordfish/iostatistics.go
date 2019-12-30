//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

// IOStatistics is used to represent the IO statistics of the requested
// object.
type IOStatistics struct {
	// NonIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing non IO requests.
	NonIORequestTime string
	// NonIORequests shall represent the total count from the time of last reset
	// or wrap of non IO requests.
	NonIORequests int64
	// ReadHitIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from memory.
	ReadHitIORequests int64
	// ReadIOKiBytes shall represent the total number of kibibytes read from the
	// time of last reset or wrap.
	ReadIOKiBytes int64
	// ReadIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing read requests.
	ReadIORequestTime string
	// ReadIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from either media or memory
	// (i.e. from a storage device or from a cache).
	ReadIORequests int64
	// WriteHitIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests coalesced into memory.
	WriteHitIORequests int64
	// WriteIOKiBytes shall represent the total number of kibibytes written from
	// the time of last reset or wrap.
	WriteIOKiBytes int64
	// WriteIORequestTime shall be an ISO 8601 conformant duration describing
	// the time that the resource is busy processing write requests.
	WriteIORequestTime string
	// WriteIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests.
	WriteIORequests int64
}
