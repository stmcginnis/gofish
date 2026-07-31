//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
)

// TestUpdateFromRawDataNilEntity verifies the nil-receiver contract: a nil
// *Entity returns an error rather than panicking in the delegating shim.
func TestUpdateFromRawDataNilEntity(t *testing.T) {
	var e *Entity
	err := e.UpdateFromRawData(&struct{}{}, []byte("{}"), nil)
	if err == nil || err.Error() != "entity is nil" {
		t.Fatalf("want \"entity is nil\" error, got %v", err)
	}
}

// TestCollectionErrorUnwrap verifies errors.Is sees through CollectionError to
// the per-member failures, in particular context cancellation.
func TestCollectionErrorUnwrap(t *testing.T) {
	ce := NewCollectionError()
	ce.Failures["/redfish/v1/x"] = fmt.Errorf("fetching member: %w", context.Canceled)
	ce.Failures["/redfish/v1/y"] = errors.New("unrelated")

	if !errors.Is(ce, context.Canceled) {
		t.Error("errors.Is(ce, context.Canceled) = false, want true")
	}
	if errors.Is(ce, context.DeadlineExceeded) {
		t.Error("errors.Is(ce, context.DeadlineExceeded) = true, want false")
	}
}

// TestTestClientContext verifies the mock client mirrors the real client's
// context behavior: per-call contexts are honored and recorded, and the
// context-free methods use the cached context.
func TestTestClientContext(t *testing.T) {
	tc := &TestClient{}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := tc.GetWithContext(cancelledCtx, "/redfish/v1/"); !errors.Is(err, context.Canceled) {
		t.Fatalf("GetWithContext with cancelled context: want context.Canceled, got %v", err)
	}

	if resp, err := tc.GetWithContext(context.Background(), "/redfish/v1/"); err != nil {
		t.Fatalf("GetWithContext with live context failed: %v", err)
	} else {
		_ = resp.Body.Close()
	}

	type key struct{}
	cached := context.WithValue(context.Background(), key{}, "cached")
	tc.SetContext(cached)
	if resp, err := tc.Get("/redfish/v1/"); err != nil {
		t.Fatalf("Get with cached context failed: %v", err)
	} else {
		_ = resp.Body.Close()
	}

	calls := tc.CapturedCalls()
	if len(calls) != 3 {
		t.Fatalf("want 3 recorded calls, got %d", len(calls))
	}
	if !errors.Is(calls[0].Context.Err(), context.Canceled) {
		t.Error("first call did not record the cancelled per-call context")
	}
	if calls[2].Context != cached {
		t.Error("context-free call did not record the cached context")
	}
}

// TestGetObjectsWithContextCancellation verifies that a dead context fails
// every member fetch fast, that all members are accounted for in the
// CollectionError, and that cancellation is visible through errors.Is.
func TestGetObjectsWithContextCancellation(t *testing.T) {
	tc := &TestClient{}
	uris := []string{"/redfish/v1/a", "/redfish/v1/b", "/redfish/v1/c", "/redfish/v1/d", "/redfish/v1/e"}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	result, err := GetObjectsWithContext[Resource](ctx, tc, uris)
	if len(result) != 0 {
		t.Errorf("want no results from a cancelled fetch, got %d", len(result))
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("want context.Canceled through the CollectionError, got %v", err)
	}
	var ce *CollectionError
	if !errors.As(err, &ce) {
		t.Fatalf("want *CollectionError, got %T", err)
	}
	if len(ce.Failures) != len(uris) {
		t.Fatalf("want %d per-member failures, got %d: %v", len(uris), len(ce.Failures), err)
	}
}

// TestCollectResourceCollectionWithContextMidCancel cancels while the first
// member is being fetched and verifies the dispatch loop stops spawning,
// flushes the remaining members synchronously, accounts for every member
// exactly once, and returns (no goroutine or WaitGroup leak).
func TestCollectResourceCollectionWithContextMidCancel(t *testing.T) {
	const members = 8
	entities := make([]*Resource, members)
	for i := range entities {
		entities[i] = &Resource{Entity: Entity{ODataID: fmt.Sprintf("/redfish/v1/x/%d", i)}}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mu sync.Mutex
	seen := map[string]int{}
	first := make(chan struct{})
	get := func(r *Resource, _ ...QueryGroupOption) {
		mu.Lock()
		seen[r.GetODataID()]++
		n := len(seen)
		mu.Unlock()
		if n == 1 {
			// Hold the first member in flight until the test cancels.
			close(first)
			<-ctx.Done()
		}
	}
	go func() {
		<-first
		cancel()
	}()

	CollectResourceCollectionWithContext(ctx, get, entities)

	mu.Lock()
	defer mu.Unlock()
	if len(seen) != members {
		t.Fatalf("want all %d members accounted for, got %d", members, len(seen))
	}
	for id, n := range seen {
		if n != 1 {
			t.Errorf("member %s dispatched %d times, want exactly once", id, n)
		}
	}
}

// TestContextOf verifies the cached-context resolution rules, including the
// nil cases.
func TestContextOf(t *testing.T) {
	if got := ContextOf(nil); got != context.Background() {
		t.Errorf("ContextOf(nil) = %v, want context.Background()", got)
	}

	var typedNil *TestClient
	if got := ContextOf(typedNil); got != context.Background() {
		t.Errorf("ContextOf(typed nil) = %v, want context.Background()", got)
	}

	tc := &TestClient{}
	type key struct{}
	cached := context.WithValue(context.Background(), key{}, "cached")
	tc.SetContext(cached)
	if got := ContextOf(tc); got != cached {
		t.Errorf("ContextOf(context client) did not return the cached context")
	}
}
