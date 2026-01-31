//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"github.com/stmcginnis/gofish/schemas"
)

// Snooping is an instance of a Snooping object.
type Snooping struct {
	schemas.Entity

	PostCode string
}

// GetSnooping will get a Snooping instance from the service.
func GetSnooping(c schemas.Client, uri string) (*Snooping, error) {
	return schemas.GetObject[Snooping](c, uri)
}
