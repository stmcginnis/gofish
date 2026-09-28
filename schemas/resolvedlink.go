//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"sync/atomic"
)

// ResolvedLink holds a singleton link's URI plus, if the service may have
// inlined real content for it (e.g. via $expand), the raw JSON of that
// content decodes it and decides whether it is genuinely an expanded object.
// The zero value (a bare, never-set link) behaves like an empty URI with
// nothing inlined.
type ResolvedLink[T any] struct {
	uri     string
	inlined *atomic.Pointer[json.RawMessage]
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
// URI is always populated when present, and the raw JSON is retained only
// when it carries something beyond "@odata.id"/"href" and so might be an
// $expand-inlined object.
//
// This check is deliberately permissive: a decorated link (an "@odata.type"
// or "@odata.etag" alongside the URI) carries no real content but still
// gets retained here. ResolveOrGet makes the authoritative call, using the
// same test resolveMember applies to collection members, and falls back to
// a live fetch for anything that is not a fully populated object.
//
// Exported so that generated types in package gofish, rather than package schemas,
// can call this as schemas.ResolveLink.
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
			// encoding/json may reuse raw's backing array after this
			// returns, so keep our own copy.
			buf := make(json.RawMessage, len(raw))
			copy(buf, raw)

			link.inlined = new(atomic.Pointer[json.RawMessage])
			link.inlined.Store(&buf)

			break
		}
	}

	return link
}
