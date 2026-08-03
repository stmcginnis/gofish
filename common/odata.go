//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"fmt"
	"strings"
)

type Query struct {
	expand         ExpandOption
	expandLevel    int
	expandFallback bool
}

type QueryGroup struct {
	QueryCollection Query // options for collections
	QueryResource   Query // options for resources
}

type QueryOption func(*Query)

type QueryGroupOption func(*QueryGroup)

type ExpandOption string

const (
	ExpandNone           ExpandOption = ""
	ExpandOptionAsterisk ExpandOption = "*"
	ExpandOptionTilde    ExpandOption = "~"
	ExpandOptionPeriod   ExpandOption = "."
)

func WithExpand(expandValue ExpandOption) func(*Query) {
	return func(q *Query) {
		q.expand = expandValue
	}
}

func WithExpandFallback(enable bool) func(*Query) {
	return func(q *Query) {
		q.expandFallback = enable
	}
}

func WithExpandLevel(expandLevel int) func(*Query) {
	return func(q *Query) {
		q.expandLevel = expandLevel
	}
}

// WithoutExpand disables $expand for a single call, on both collection and
// resource requests, overriding any default set on the client.
func WithoutExpand() func(*QueryGroup) {
	return func(q *QueryGroup) {
		q.QueryCollection.expand = ExpandNone
		q.QueryCollection.expandLevel = 0
		q.QueryResource.expand = ExpandNone
		q.QueryResource.expandLevel = 0
	}
}

func WithResourceQueryOpts(queryOpts ...QueryOption) func(*QueryGroup) {
	return func(q *QueryGroup) {
		for _, queryOpt := range queryOpts {
			queryOpt(&q.QueryResource)
		}
	}
}

func WithCollectionQueryOpts(queryOpts ...QueryOption) func(*QueryGroup) {
	return func(q *QueryGroup) {
		for _, queryOpt := range queryOpts {
			queryOpt(&q.QueryCollection)
		}
	}
}

func BuildQueryGroup(c Client, opts ...QueryGroupOption) *QueryGroup {
	queryGroup := &QueryGroup{}

	// apply client settings first, followed by override settings.
	// this must not append directly to the client's slice: any spare capacity
	// there would be written to by concurrent callers, letting one call's
	// options leak into another's.
	defaults := c.GetSettings().DefaultQueryOptions
	all := make([]QueryGroupOption, 0, len(defaults)+len(opts))
	all = append(all, defaults...)
	all = append(all, opts...)

	for _, opt := range all {
		opt(queryGroup)
	}

	return queryGroup
}

func BuildQuery(c Client, url string, collection bool, opts ...QueryGroupOption) string {
	queryGroup := BuildQueryGroup(c, opts...)

	q := queryGroup.QueryResource
	if collection {
		q = queryGroup.QueryCollection
	}

	queryOpts := ""
	if q.expand != ExpandNone {
		queryOpts += fmt.Sprintf("$expand=%s", string(q.expand))
		if q.expandLevel > 0 {
			queryOpts += fmt.Sprintf("($levels=%d)", q.expandLevel)
		}
	}

	if queryOpts != "" {
		// the url may already carry a query string. Members@odata.nextLink,
		// for example, encodes the page offset in one.
		separator := "?"
		if strings.Contains(url, "?") {
			separator = "&"
		}
		url = url + separator + queryOpts
	}

	return url
}
