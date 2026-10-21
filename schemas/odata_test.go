//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "testing"

func TestBuildQuery(t *testing.T) {
	expandAll := WithCollectionQueryOpts(WithExpand(ExpandOptionAsterisk))
	for _, tc := range []struct {
		name, uri, want string
		opts            []QueryGroupOption
		defaults        []QueryGroupOption
	}{
		{"no query", "/redfish/v1/Chassis", "/redfish/v1/Chassis?$expand=*", []QueryGroupOption{expandAll}, nil},
		{"existing query", "/logs?$skip=50", "/logs?$skip=50&$expand=*", []QueryGroupOption{expandAll}, nil},
		{
			"existing expand", "/logs?$skip=50&$expand=*", "/logs?$skip=50&$expand=*",
			[]QueryGroupOption{WithCollectionQueryOpts(WithExpand(ExpandOptionPeriod), WithExpandLevel(2))}, nil,
		},
		{"encoded expand", "/logs?%24expand=*", "/logs?%24expand=*", []QueryGroupOption{expandAll}, nil},
		{"no expand", "/logs?$skip=50", "/logs?$skip=50", nil, nil},
		{"default expand", "/logs?$skip=50", "/logs?$skip=50&$expand=*", nil, []QueryGroupOption{expandAll}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := &TestClient{}
			c.Settings.DefaultQueryOptions = tc.defaults
			if got := BuildQuery(c, tc.uri, true, tc.opts...); got != tc.want {
				t.Errorf("expected %q, got %q", tc.want, got)
			}
		})
	}
}
