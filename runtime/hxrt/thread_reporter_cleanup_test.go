package hxrt

import (
	"testing"
	"time"
)

// A failed reporter must remove lookup state but retain foreground accounting
// until the fatal unwind terminates the process. This probe catches that unwind
// only to inspect the boundary, then restores the count for the remaining tests.
func TestPortableReporterFailureCleansLookupState(t *testing.T) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	baselineActive := activePortableThreads
	threadRuntimeMu.Unlock()
	id := allocateThreadState(nil)
	marker := &struct{}{}
	type observation struct {
		failure any
		state   bool
		mapping bool
		active  int
	}
	observed := make(chan observation, 1)
	go func() {
		defer func() {
			failure := recover()
			gid := currentGoroutineID()
			threadRuntimeMu.Lock()
			_, state := threadStates[id]
			_, mapping := goroutineThreadIDs[gid]
			active := activePortableThreads
			delete(threadStates, id)
			delete(goroutineThreadIDs, gid)
			if activePortableThreads > baselineActive {
				activePortableThreads--
			}
			if activePortableThreads == 0 {
				threadRuntimeCond.Broadcast()
			}
			threadRuntimeMu.Unlock()
			observed <- observation{failure, state, mapping, active}
		}()
		runPortableThreadJob(id, func() { Throw("original failure") }, func(any) { panic(marker) })
	}()
	select {
	case got := <-observed:
		if got.failure != marker {
			t.Fatalf("reporter panic was not preserved: %v", got.failure)
		}
		if got.state || got.mapping {
			t.Fatalf("reporter failure leaked lookup state: state=%t mapping=%t", got.state, got.mapping)
		}
		if got.active != baselineActive+1 {
			t.Fatalf("fatal unwind released foreground count: got %d, want %d", got.active, baselineActive+1)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("reporter failure probe timed out")
	}
}
