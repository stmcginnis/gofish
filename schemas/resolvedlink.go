//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"sync/atomic"
)

// ResolvedLink holds a singleton link's URI plus, if the service inlined
// real content for it (e.g. via $expand), the decoded value - consumed
// exactly once by ResolveOrGet. Built by ResolveLink; the zero value (a
// bare, never-set link) behaves like an empty URI with nothing resolved.
type ResolvedLink[T any] struct {
	uri      string
	resolved *atomic.Pointer[T]
}

// URI returns the link's URI, or "" if it was never set. Exported so
// consumers outside package schemas (and tests) can inspect it without
// going through a live ResolveOrGet fetch.
func (r ResolvedLink[T]) URI() string {
	return r.uri
}

// UnmarshalJSON lets ResolvedLink[T] be used directly as a decode-target
// field's type (like this codebase's other Link/Links/ActionTarget helpers),
// so no separate ResolveLink call or json.RawMessage intermediary is needed
// at each call site - encoding/json invokes this automatically.
func (r *ResolvedLink[T]) UnmarshalJSON(raw []byte) error {
	*r = ResolveLink[T](raw)

	return nil
}

// ResolveLink parses a Link-shaped JSON value into a ResolvedLink[T]: the
// URI is always populated when present, and resolved is set only if the raw
// JSON carried real content beyond "@odata.id"/"href".
//
// Exported (despite living next to other generated-code helpers) because
// some generated types live in package gofish, not package schemas, and call
// this as schemas.ResolveLink.
func ResolveLink[T any](raw json.RawMessage) ResolvedLink[T] {
	if len(raw) == 0 {
		return ResolvedLink[T]{}
	}

	var obj map[string]any
	if json.Unmarshal(raw, &obj) != nil {
		return ResolvedLink[T]{}
	}

	link := ResolvedLink[T]{}
	if id, ok := obj["@odata.id"].(string); ok {
		link.uri = id
	} else if href, ok := obj["href"].(string); ok {
		link.uri = href
	}

	for key := range obj {
		if key != "@odata.id" && key != "href" {
			entity := new(T)
			if json.Unmarshal(raw, entity) == nil {
				link.resolved = new(atomic.Pointer[T])
				link.resolved.Store(entity)
			}

			break
		}
	}

	return link
}
