//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import "fmt"

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

	// apply client settings first, followed by override settings
	opts = append(c.GetSettings().DefaultQueryOptions, opts...)

	for _, opt := range opts {
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
		url = url + "?" + queryOpts
	}

	return url
}
