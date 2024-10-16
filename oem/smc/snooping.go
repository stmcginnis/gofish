//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"github.com/stmcginnis/gofish/common"
)

// Snooping is an instance of a Snooping object.
type Snooping struct {
	common.Entity

	PostCode string
}

// GetSnooping will get a Snooping instance from the service.
func GetSnooping(c common.Client, uri string) (*Snooping, error) {
	return common.GetObject[Snooping](c, uri)
}
