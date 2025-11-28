//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import "fmt"

type Query struct {
	expand      ExpandOption
	expandLevel int
}

type QueryOption func(*Query)

type ExpandOption string

const (
	ExpandOptionAsterisk ExpandOption = "*"
	ExpandOptionTilde    ExpandOption = "~"
	ExpandOptionPeriod   ExpandOption = "."
)

func WithExpand(expandValue ExpandOption) func(*Query) {
	return func(q *Query) {
		q.expand = expandValue
	}
}

func WithExpandLevel(expandLevel int) func(*Query) {
	return func(q *Query) {
		q.expandLevel = expandLevel
	}
}

func BuildQuery(url string, opts ...QueryOption) string {
	if len(opts) == 0 {
		return url
	}

	q := &Query{}
	for _, opt := range opts {
		opt(q)
	}

	url += "?"
	if q.expand != "" {
		url += fmt.Sprintf("$expand=%s", string(q.expand))
		if q.expandLevel > 0 {
			url += fmt.Sprintf("($levels=%d)", q.expandLevel)
		}
	}

	return url
}
