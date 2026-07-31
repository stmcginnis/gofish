//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"

	"github.com/coreweave/gofish/common"
)

// Snooping is an instance of a Snooping object.
type Snooping struct {
	common.Entity

	PostCode string
}

// GetSnooping will get a Snooping instance from the service.
func GetSnooping(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Snooping, error) {
	return GetSnoopingWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetSnoopingWithContext will get a Snooping instance from the service.
func GetSnoopingWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Snooping, error) {
	return common.GetObjectWithContext[Snooping](ctx, c, uri, queryOpts...)
}
