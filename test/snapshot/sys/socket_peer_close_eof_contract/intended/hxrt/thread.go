package hxrt

import (
	"bytes"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type ThreadState struct {
	queueMu           sync.Mutex
	queueCond         *sync.Cond
	queue             []any
	eventLoopMu       sync.RWMutex
	eventLoop         *EventLoopHandle
	threadLocalMu     sync.RWMutex
	threadLocalValues map[uint64]any
}

type EventLoopHandle struct {
	mu        sync.Mutex
	cond      *sync.Cond
	oneTime   []func()
	promised  int
	regular   []*RegularEvent
	running   map[int]bool
	cancelled map[int]bool
}

type RegularEvent struct {
	ID       int
	NextRun  float64
	Interval float64
	Run      func()
}

type EventLoopProgress struct {
	Kind int
	Time float64
}

type LockHandle struct {
	mu    sync.Mutex
	cond  *sync.Cond
	count int
}

type MutexHandle struct {
	mu    sync.Mutex
	cond  *sync.Cond
	owner int64
	depth int
}

type ConditionHandle struct {
	mu           sync.Mutex
	mutexCond    *sync.Cond
	owner        int64
	depth        int
	nextWaiterID uint64
	waiters      map[uint64]chan struct{}
}

type SemaphoreHandle struct {
	mu    sync.Mutex
	cond  *sync.Cond
	count int
}

// ThreadLocalHandle identifies one staged sys.thread.Tls instance without
// exposing raw slot identifiers to generated Haxe code.
type ThreadLocalHandle struct {
	id uint64
}

const (
	threadEventLoopNow = iota
	threadEventLoopNever
	threadEventLoopAnyTime
	threadEventLoopAt
)

var (
	threadRuntimeOnce     sync.Once
	threadRuntimeMu       sync.Mutex
	threadRuntimeCond     *sync.Cond
	threadStates          map[int]*ThreadState
	goroutineThreadIDs    map[int64]int
	activePortableThreads int
	nextThreadID          int
	nextEventID           atomic.Int64
	nextThreadLocalID     atomic.Uint64
	threadStartTime       time.Time
)

func initThreadRuntime() {
	threadStates = make(map[int]*ThreadState)
	goroutineThreadIDs = make(map[int64]int)
	threadRuntimeCond = sync.NewCond(&threadRuntimeMu)
	activePortableThreads = 0
	nextThreadID = 1
	threadStartTime = time.Now()
	mainState := newThreadState()
	setThreadEventLoop(mainState, newEventLoopHandle())
	threadStates[0] = mainState
	goroutineThreadIDs[currentGoroutineID()] = 0
}

func ensureThreadRuntime() {
	threadRuntimeOnce.Do(initThreadRuntime)
}

func newThreadState() *ThreadState {
	state := &ThreadState{}
	state.queueCond = sync.NewCond(&state.queueMu)
	return state
}

// threadEventLoop reads the loop ownership of a logical thread.
//
// What: it is the single read path for ThreadState.eventLoop.
// Why: loop ownership is observable from other goroutines while the owning
// thread can install or remove a temporary loop.
// How: a dedicated RWMutex keeps that small ownership field race-free without
// serializing message-queue or event-loop callback work.
func threadEventLoop(state *ThreadState) *EventLoopHandle {
	state.eventLoopMu.RLock()
	handle := state.eventLoop
	state.eventLoopMu.RUnlock()
	return handle
}

// setThreadEventLoop publishes loop ownership. Callers that need a
// compare-and-install transition must hold eventLoopMu directly instead.
func setThreadEventLoop(state *ThreadState, handle *EventLoopHandle) {
	state.eventLoopMu.Lock()
	state.eventLoop = handle
	state.eventLoopMu.Unlock()
}

func newEventLoopHandle() *EventLoopHandle {
	h := &EventLoopHandle{
		running:   make(map[int]bool),
		cancelled: make(map[int]bool),
	}
	h.cond = sync.NewCond(&h.mu)
	return h
}

func allocateThreadState(eventLoop *EventLoopHandle) int {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	defer threadRuntimeMu.Unlock()
	id := nextThreadID
	nextThreadID++
	state := newThreadState()
	setThreadEventLoop(state, eventLoop)
	threadStates[id] = state
	activePortableThreads++
	return id
}

func registerCurrentGoroutineThreadID(id int) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	goroutineThreadIDs[currentGoroutineID()] = id
	threadRuntimeMu.Unlock()
}

func unregisterCurrentGoroutineThreadID(id int, completed bool) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	delete(goroutineThreadIDs, currentGoroutineID())
	if _, exists := threadStates[id]; exists {
		delete(threadStates, id)
		if completed {
			activePortableThreads--
			if activePortableThreads == 0 {
				threadRuntimeCond.Broadcast()
			}
		}
	}
	threadRuntimeMu.Unlock()
}

func currentLogicalThreadID() int {
	ensureThreadRuntime()
	gid := currentGoroutineID()
	threadRuntimeMu.Lock()
	defer threadRuntimeMu.Unlock()
	if id, ok := goroutineThreadIDs[gid]; ok {
		return id
	}
	id := nextThreadID
	nextThreadID++
	goroutineThreadIDs[gid] = id
	threadStates[id] = newThreadState()
	return id
}

func withThreadState(threadID int, fn func(*ThreadState)) bool {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	state, ok := threadStates[threadID]
	threadRuntimeMu.Unlock()
	if !ok || state == nil {
		return false
	}
	fn(state)
	return true
}

func currentThreadState() *ThreadState {
	id := currentLogicalThreadID()
	var state *ThreadState
	withThreadState(id, func(found *ThreadState) {
		state = found
	})
	return state
}

func threadNowSeconds() float64 {
	ensureThreadRuntime()
	return time.Since(threadStartTime).Seconds()
}

func ThreadNowSeconds() float64 {
	return threadNowSeconds()
}

func ThreadCurrentId() int {
	return currentLogicalThreadID()
}

// ThreadLocalNew allocates a process-unique TLS slot handle.
//
// What: the handle names one sys.thread.Tls instance.
// Why: values belong to ThreadState so deleting a completed thread state drops
// every value it owned, while a monotonically increasing handle prevents one
// Tls instance from observing another instance's values.
// How: zero is reserved and wraparound fails natively rather than reusing a
// live or historical slot identifier.
func ThreadLocalNew() *ThreadLocalHandle {
	id := nextThreadLocalID.Add(1)
	if id == 0 {
		panic("hxrt: thread-local handle space exhausted")
	}
	return &ThreadLocalHandle{id: id}
}

// ThreadLocalGet reads the current logical thread's value for one TLS handle.
//
// What: it returns the value stored in the caller's ThreadState.
// Why: sys.thread.Tls<T> must isolate values without retaining them after the
// owning supported thread lifecycle ends.
// How: the ThreadState lock protects the slot map. The any boundary is
// intentional because T can be every Haxe value; staged Tls restores the type.
func ThreadLocalGet(handle *ThreadLocalHandle) any {
	if handle == nil {
		return nil
	}
	state := currentThreadState()
	if state == nil {
		return nil
	}
	state.threadLocalMu.RLock()
	value := state.threadLocalValues[handle.id]
	state.threadLocalMu.RUnlock()
	return value
}

// ThreadLocalSet writes or clears the current logical thread's TLS slot.
//
// What: non-null values populate the caller's ThreadState; null deletes a slot.
// Why: TLS payloads must share the lifecycle of their logical thread instead of
// remaining in a global or per-Tls registry after that thread completes.
// How: the ThreadState lock protects mutation, and clearing the last value also
// releases the map allocation.
func ThreadLocalSet(handle *ThreadLocalHandle, value any) {
	if handle == nil {
		return
	}
	state := currentThreadState()
	if state == nil {
		return
	}
	state.threadLocalMu.Lock()
	if AnyEqualsNull(value) {
		delete(state.threadLocalValues, handle.id)
		if len(state.threadLocalValues) == 0 {
			state.threadLocalValues = nil
		}
	} else {
		if state.threadLocalValues == nil {
			state.threadLocalValues = make(map[uint64]any)
		}
		state.threadLocalValues[handle.id] = value
	}
	state.threadLocalMu.Unlock()
}

// ThreadWaitForAll drains foreground portable threads, including descendants
// created by a running portable thread. Explicit go.Go.spawn goroutines are not
// counted and retain Go's native process-shutdown behavior.
func ThreadWaitForAll() {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	for activePortableThreads > 0 {
		threadRuntimeCond.Wait()
	}
	threadRuntimeMu.Unlock()
}

// clearDetachedThreadState removes only lazily-created identity owned by a
// compiler-scoped detached goroutine. Detached goroutines never contribute to
// activePortableThreads, so this transition must not touch the foreground
// counter.
func clearDetachedThreadState(goroutineID int64) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	if id, exists := goroutineThreadIDs[goroutineID]; exists && id != 0 {
		delete(goroutineThreadIDs, goroutineID)
		delete(threadStates, id)
	}
	threadRuntimeMu.Unlock()
}

// runDetachedThreadJob owns identity cleanup for one fresh Go goroutine.
// Capturing the runtime identity before the callback keeps the deferred cleanup
// non-panicking during native panic unwinding; the panic itself is never
// recovered or converted into a Haxe exception.
func runDetachedThreadJob(job func()) {
	goroutineID := currentGoroutineID()
	clearDetachedThreadState(goroutineID)
	defer clearDetachedThreadState(goroutineID)
	job()
}

// ThreadSpawnDetached starts a compiler-owned go.Go.spawn callback.
//
// What: it launches one unjoined native goroutine with an identity cleanup
// scope around its callback.
// Why: synchronous initialization keeps reserved logical identity zero owned by
// the caller even when this is the program's first hxrt thread operation.
// How: initialize before go, then let runDetachedThreadJob remove callback state
// during return or panic unwind. A nil callback is deliberately invoked so its
// native panic is preserved rather than silently ignored.
func ThreadSpawnDetached(job func()) {
	ensureThreadRuntime()
	go runDetachedThreadJob(job)
}

func runPortableThreadJob(id int, job func()) {
	registerCurrentGoroutineThreadID(id)
	defer func() {
		if recovered := recover(); recovered != nil {
			value, ok := unwrapHaxeException(recovered)
			if !ok {
				// Remove lookup state but keep the foreground count nonzero until
				// the re-panic terminates the process. Waking generated main here
				// would let normal process exit race with the fatal native panic.
				unregisterCurrentGoroutineThreadID(id, false)
				panic(recovered)
			}
			ReportUncaughtException(value)
		}
		unregisterCurrentGoroutineThreadID(id, true)
	}()
	job()
}

func ThreadSpawn(job func()) int {
	if job == nil {
		return 0
	}
	id := allocateThreadState(nil)
	go func() {
		runPortableThreadJob(id, job)
	}()
	return id
}

func ThreadSpawnWithEventLoop(job func()) int {
	if job == nil {
		return 0
	}
	id := allocateThreadState(newEventLoopHandle())
	go func() {
		runPortableThreadJob(id, func() {
			job()
			withThreadState(id, func(state *ThreadState) {
				ThreadEventLoopLoop(threadEventLoop(state))
			})
		})
	}()
	return id
}

func ThreadHasEventLoop(threadID int) bool {
	available := false
	withThreadState(threadID, func(state *ThreadState) {
		available = threadEventLoop(state) != nil
	})
	return available
}

func ThreadEvents(threadID int) *EventLoopHandle {
	var handle *EventLoopHandle
	withThreadState(threadID, func(state *ThreadState) {
		handle = threadEventLoop(state)
	})
	return handle
}

func ThreadRunWithEventLoop(job func()) {
	if job == nil {
		return
	}
	state := currentThreadState()
	if state == nil {
		job()
		return
	}
	state.eventLoopMu.Lock()
	if state.eventLoop != nil {
		state.eventLoopMu.Unlock()
		job()
		return
	}
	loop := newEventLoopHandle()
	state.eventLoop = loop
	state.eventLoopMu.Unlock()
	defer func() {
		state.eventLoopMu.Lock()
		if state.eventLoop == loop {
			state.eventLoop = nil
		}
		state.eventLoopMu.Unlock()
	}()
	job()
	ThreadEventLoopLoop(loop)
}

func ThreadSendMessage(threadID int, message any) {
	withThreadState(threadID, func(state *ThreadState) {
		state.queueMu.Lock()
		state.queue = append(state.queue, message)
		state.queueCond.Signal()
		state.queueMu.Unlock()
	})
}

func ThreadReadMessage(block bool) any {
	state := currentThreadState()
	if state == nil {
		return nil
	}
	state.queueMu.Lock()
	defer state.queueMu.Unlock()
	for len(state.queue) == 0 {
		if !block {
			return nil
		}
		state.queueCond.Wait()
	}
	value := state.queue[0]
	state.queue[0] = nil
	state.queue = state.queue[1:]
	return value
}

func ThreadEventLoopNew() *EventLoopHandle {
	return newEventLoopHandle()
}

func ThreadEventLoopPromise(handle *EventLoopHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	handle.promised++
	handle.mu.Unlock()
}

func ThreadEventLoopRun(handle *EventLoopHandle, event func()) {
	if handle == nil || event == nil {
		return
	}
	handle.mu.Lock()
	handle.oneTime = append(handle.oneTime, event)
	handle.cond.Signal()
	handle.mu.Unlock()
}

func ThreadEventLoopRunPromised(handle *EventLoopHandle, event func()) {
	if handle == nil || event == nil {
		return
	}
	handle.mu.Lock()
	handle.oneTime = append(handle.oneTime, event)
	if handle.promised > 0 {
		handle.promised--
	}
	handle.cond.Signal()
	handle.mu.Unlock()
}

func ThreadEventLoopRepeat(handle *EventLoopHandle, event func(), intervalMs int) int {
	if handle == nil || event == nil {
		return 0
	}
	if intervalMs < 1 {
		intervalMs = 1
	}
	id := int(nextEventID.Add(1))
	interval := float64(intervalMs) / 1000.0
	regular := &RegularEvent{
		ID:       id,
		NextRun:  threadNowSeconds() + interval,
		Interval: interval,
		Run:      event,
	}
	handle.mu.Lock()
	handle.regular = append(handle.regular, regular)
	sort.Slice(handle.regular, func(i, j int) bool {
		return handle.regular[i].NextRun < handle.regular[j].NextRun
	})
	handle.cond.Signal()
	handle.mu.Unlock()
	return id
}

func ThreadEventLoopCancel(handle *EventLoopHandle, eventID int) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	removed := false
	filtered := handle.regular[:0]
	for _, event := range handle.regular {
		if event.ID == eventID {
			removed = true
			continue
		}
		filtered = append(filtered, event)
	}
	for index := len(filtered); index < len(handle.regular); index++ {
		handle.regular[index] = nil
	}
	handle.regular = filtered
	if handle.running[eventID] {
		// A repeating callback is run without the event-loop lock. Remember a
		// cancellation only for that bounded interval so progress can suppress
		// its reschedule. Queued and unknown identifiers need no tombstone.
		handle.cancelled[eventID] = true
	}
	if removed || handle.running[eventID] {
		handle.cond.Broadcast()
	}
	handle.mu.Unlock()
}

func ThreadEventLoopProgress(handle *EventLoopHandle) *EventLoopProgress {
	if handle == nil {
		return &EventLoopProgress{Kind: threadEventLoopNever, Time: -1}
	}
	now := threadNowSeconds()
	var callbacks []func()
	var reschedule []*RegularEvent
	nextAt := -1.0
	promised := 0

	handle.mu.Lock()
	for len(handle.regular) > 0 && handle.regular[0].NextRun <= now {
		event := handle.regular[0]
		handle.regular = handle.regular[1:]
		handle.running[event.ID] = true
		callbacks = append(callbacks, event.Run)
		reschedule = append(reschedule, event)
	}
	if len(handle.oneTime) > 0 {
		callbacks = append(callbacks, handle.oneTime...)
		handle.oneTime = handle.oneTime[:0]
	}
	promised = handle.promised
	if len(handle.regular) > 0 {
		nextAt = handle.regular[0].NextRun
	}
	handle.mu.Unlock()

	for _, callback := range callbacks {
		if callback != nil {
			callback()
		}
	}

	if len(reschedule) > 0 {
		handle.mu.Lock()
		rescheduleAt := threadNowSeconds()
		for _, event := range reschedule {
			delete(handle.running, event.ID)
			if handle.cancelled[event.ID] {
				delete(handle.cancelled, event.ID)
				continue
			}
			event.NextRun = rescheduleAt + event.Interval
			handle.regular = append(handle.regular, event)
		}
		sort.Slice(handle.regular, func(i, j int) bool {
			return handle.regular[i].NextRun < handle.regular[j].NextRun
		})
		if len(handle.regular) > 0 {
			nextAt = handle.regular[0].NextRun
		} else {
			nextAt = -1.0
		}
		handle.mu.Unlock()
	}

	if len(callbacks) > 0 {
		return &EventLoopProgress{Kind: threadEventLoopNow, Time: -1}
	}
	if promised > 0 {
		if nextAt >= 0 {
			return &EventLoopProgress{Kind: threadEventLoopAnyTime, Time: nextAt}
		}
		return &EventLoopProgress{Kind: threadEventLoopAnyTime, Time: -1}
	}
	if nextAt >= 0 {
		return &EventLoopProgress{Kind: threadEventLoopAt, Time: nextAt}
	}
	return &EventLoopProgress{Kind: threadEventLoopNever, Time: -1}
}

func ThreadEventLoopWait(handle *EventLoopHandle) bool {
	return ThreadEventLoopWaitTimeout(handle, -1)
}

// ThreadEventLoopWaitTimeout blocks for one scheduler-state transition.
//
// What: it returns whether work remains after an event insertion, cancellation,
// or timeout wakes the loop.
// Why: waiting until the old deadline after a wake would delay a newly inserted
// earlier timer and would strand a loop whose final timer was cancelled.
// How: every producer signals handle.cond, and a timed wait returns after that
// single signal so ThreadEventLoopLoop can recompute the complete schedule.
func ThreadEventLoopWaitTimeout(handle *EventLoopHandle, timeout float64) bool {
	if handle == nil {
		return false
	}
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if eventLoopHasPendingLocked(handle) == false {
		return false
	}
	if eventLoopHasReadyLocked(handle) {
		return true
	}
	if timeout < 0 {
		handle.cond.Wait()
		return eventLoopHasPendingLocked(handle)
	}
	if timeout == 0 {
		return eventLoopHasPendingLocked(handle)
	}
	timer := time.AfterFunc(time.Duration(timeout*float64(time.Second)), func() {
		handle.mu.Lock()
		handle.cond.Broadcast()
		handle.mu.Unlock()
	})
	defer timer.Stop()
	handle.cond.Wait()
	return eventLoopHasPendingLocked(handle)
}

func ThreadEventLoopLoop(handle *EventLoopHandle) {
	if handle == nil {
		return
	}
	for {
		progress := ThreadEventLoopProgress(handle)
		switch progress.Kind {
		case threadEventLoopNow:
			continue
		case threadEventLoopNever:
			return
		case threadEventLoopAnyTime:
			if progress.Time >= 0 {
				timeout := progress.Time - threadNowSeconds()
				if timeout < 0 {
					timeout = 0
				}
				ThreadEventLoopWaitTimeout(handle, timeout)
			} else {
				ThreadEventLoopWait(handle)
			}
		case threadEventLoopAt:
			timeout := progress.Time - threadNowSeconds()
			if timeout < 0 {
				timeout = 0
			}
			ThreadEventLoopWaitTimeout(handle, timeout)
		}
	}
}

func eventLoopHasPendingLocked(handle *EventLoopHandle) bool {
	return len(handle.oneTime) > 0 || len(handle.regular) > 0 || handle.promised > 0
}

func eventLoopHasReadyLocked(handle *EventLoopHandle) bool {
	if len(handle.oneTime) > 0 {
		return true
	}
	if len(handle.regular) > 0 && handle.regular[0].NextRun <= threadNowSeconds() {
		return true
	}
	return false
}

func ThreadLockNew() *LockHandle {
	h := &LockHandle{}
	h.cond = sync.NewCond(&h.mu)
	return h
}

func ThreadLockWait(handle *LockHandle) bool {
	if handle == nil {
		return false
	}
	handle.mu.Lock()
	for handle.count == 0 {
		handle.cond.Wait()
	}
	handle.count--
	handle.mu.Unlock()
	return true
}

func ThreadLockWaitTimeout(handle *LockHandle, timeout float64) bool {
	if handle == nil {
		return false
	}
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if handle.count > 0 {
		handle.count--
		return true
	}
	if timeout <= 0 {
		return false
	}
	return waitForConditionWithTimeout(&handle.mu, handle.cond, timeout, func() bool {
		if handle.count > 0 {
			handle.count--
			return true
		}
		return false
	})
}

func ThreadLockRelease(handle *LockHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	handle.count++
	handle.cond.Signal()
	handle.mu.Unlock()
}

func ThreadMutexNew() *MutexHandle {
	h := &MutexHandle{}
	h.cond = sync.NewCond(&h.mu)
	return h
}

func ThreadMutexAcquire(handle *MutexHandle) {
	if handle == nil {
		return
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	for handle.owner != 0 && handle.owner != gid {
		handle.cond.Wait()
	}
	handle.owner = gid
	handle.depth++
	handle.mu.Unlock()
}

func ThreadMutexTryAcquire(handle *MutexHandle) bool {
	if handle == nil {
		return false
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if handle.owner != 0 && handle.owner != gid {
		return false
	}
	handle.owner = gid
	handle.depth++
	return true
}

func ThreadMutexRelease(handle *MutexHandle) {
	if handle == nil {
		return
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	if handle.owner != gid || handle.depth == 0 {
		handle.mu.Unlock()
		return
	}
	handle.depth--
	if handle.depth == 0 {
		handle.owner = 0
		handle.cond.Signal()
	}
	handle.mu.Unlock()
}

func ThreadConditionNew() *ConditionHandle {
	h := &ConditionHandle{waiters: make(map[uint64]chan struct{})}
	h.mutexCond = sync.NewCond(&h.mu)
	return h
}

// registerConditionWaiterLocked creates one non-transferable wakeup target.
//
// What: each active waiter owns a distinct channel while registered.
// Why: a shared signal-credit counter lets late waiters consume an earlier
// broadcast and lets duplicate signals leak into a later generation.
// How: signal/broadcast close only channels that already exist; a late waiter
// receives a new open channel and therefore cannot inherit an earlier wakeup.
// The caller must hold handle.mu.
func registerConditionWaiterLocked(handle *ConditionHandle) uint64 {
	for {
		handle.nextWaiterID++
		waiterID := handle.nextWaiterID
		if _, exists := handle.waiters[waiterID]; exists {
			continue
		}
		handle.waiters[waiterID] = make(chan struct{})
		return waiterID
	}
}

// conditionWaiterSignaledLocked reports whether one registered waiter's
// channel is closed. The caller must hold handle.mu.
func conditionWaiterSignaledLocked(handle *ConditionHandle, waiterID uint64) bool {
	waiter, exists := handle.waiters[waiterID]
	if !exists {
		return false
	}
	select {
	case <-waiter:
		return true
	default:
		return false
	}
}

// signalConditionWaiterLocked closes an open waiter channel exactly once. The
// caller must hold handle.mu.
func signalConditionWaiterLocked(handle *ConditionHandle, waiterID uint64) bool {
	if conditionWaiterSignaledLocked(handle, waiterID) {
		return false
	}
	waiter, exists := handle.waiters[waiterID]
	if !exists {
		return false
	}
	close(waiter)
	return true
}

// removeConditionWaiterLocked retires a wakeup target after its waiter has
// observed the signal. The caller must hold handle.mu.
func removeConditionWaiterLocked(handle *ConditionHandle, waiterID uint64) {
	delete(handle.waiters, waiterID)
}

func ThreadConditionAcquire(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	for handle.owner != 0 && handle.owner != gid {
		handle.mutexCond.Wait()
	}
	handle.owner = gid
	handle.depth++
	handle.mu.Unlock()
}

func ThreadConditionTryAcquire(handle *ConditionHandle) bool {
	if handle == nil {
		return false
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if handle.owner != 0 && handle.owner != gid {
		return false
	}
	handle.owner = gid
	handle.depth++
	return true
}

func ThreadConditionRelease(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	if handle.owner != gid || handle.depth == 0 {
		handle.mu.Unlock()
		return
	}
	handle.depth--
	if handle.depth == 0 {
		handle.owner = 0
		handle.mutexCond.Signal()
	}
	handle.mu.Unlock()
}

func ThreadConditionWait(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	gid := currentGoroutineID()
	handle.mu.Lock()
	if handle.owner != gid || handle.depth == 0 {
		handle.mu.Unlock()
		return
	}
	savedDepth := handle.depth
	handle.owner = 0
	handle.depth = 0
	waiterID := registerConditionWaiterLocked(handle)
	waiter := handle.waiters[waiterID]
	handle.mutexCond.Broadcast()
	handle.mu.Unlock()

	<-waiter

	handle.mu.Lock()
	removeConditionWaiterLocked(handle, waiterID)
	handle.mutexCond.Broadcast()
	for handle.owner != 0 && handle.owner != gid {
		handle.mutexCond.Wait()
	}
	handle.owner = gid
	handle.depth = savedDepth
	handle.mu.Unlock()
}

func ThreadConditionSignal(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	var selected uint64
	found := false
	for waiterID := range handle.waiters {
		if conditionWaiterSignaledLocked(handle, waiterID) {
			continue
		}
		if !found || waiterID < selected {
			selected = waiterID
			found = true
		}
	}
	if found {
		signalConditionWaiterLocked(handle, selected)
	}
	handle.mu.Unlock()
}

func ThreadConditionBroadcast(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	for waiterID := range handle.waiters {
		signalConditionWaiterLocked(handle, waiterID)
	}
	handle.mu.Unlock()
}

func ThreadSemaphoreNew(value int) *SemaphoreHandle {
	h := &SemaphoreHandle{count: value}
	h.cond = sync.NewCond(&h.mu)
	return h
}

func ThreadSemaphoreAcquire(handle *SemaphoreHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	for handle.count == 0 {
		handle.cond.Wait()
	}
	handle.count--
	handle.mu.Unlock()
}

func ThreadSemaphoreTryAcquire(handle *SemaphoreHandle) bool {
	if handle == nil {
		return false
	}
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if handle.count == 0 {
		return false
	}
	handle.count--
	return true
}

func ThreadSemaphoreTryAcquireTimeout(handle *SemaphoreHandle, timeout float64) bool {
	if handle == nil {
		return false
	}
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if handle.count > 0 {
		handle.count--
		return true
	}
	if timeout <= 0 {
		return false
	}
	return waitForConditionWithTimeout(&handle.mu, handle.cond, timeout, func() bool {
		if handle.count > 0 {
			handle.count--
			return true
		}
		return false
	})
}

func ThreadSemaphoreRelease(handle *SemaphoreHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	handle.count++
	handle.cond.Signal()
	handle.mu.Unlock()
}

const goroutineIDStackBytes = 64
const maximumGoroutineID = int64(1<<63 - 1)

var goroutineIDStackPrefix = []byte("goroutine ")
var goroutineIDStackBuffers = sync.Pool{
	New: func() any {
		return new([goroutineIDStackBytes]byte)
	},
}

// parseGoroutineID parses only the bounded runtime.Stack header used for
// logical-thread and re-entrant-mutex ownership.
//
// What: it accepts the leading "goroutine <positive-id> " header and rejects
// every malformed, reserved-zero, negative, or overflowing identity.
// Why: Go has no supported goroutine-local key, but Haxe's re-entrant Mutex,
// Condition, Tls, and Thread.current contracts require stable caller identity.
// How: callers provide a fixed-size stack prefix; this function never scans or
// retains a complete stack trace.
func parseGoroutineID(stackPrefix []byte) (int64, bool) {
	if !bytes.HasPrefix(stackPrefix, goroutineIDStackPrefix) {
		return 0, false
	}
	identifier := stackPrefix[len(goroutineIDStackPrefix):]
	delimiter := bytes.IndexByte(identifier, ' ')
	if delimiter <= 0 {
		return 0, false
	}
	var id int64
	for _, digitByte := range identifier[:delimiter] {
		if digitByte < '0' || digitByte > '9' {
			return 0, false
		}
		digit := int64(digitByte - '0')
		if id > (maximumGoroutineID-digit)/10 {
			return 0, false
		}
		id = id*10 + digit
	}
	if id == 0 {
		return 0, false
	}
	return id, true
}

// currentGoroutineID returns a stable positive ownership key for the current
// goroutine. A runtime header-format mismatch is fatal: silently returning the
// reserved zero value would make a locked re-entrant mutex appear unlocked.
func currentGoroutineID() int64 {
	stackPrefix := goroutineIDStackBuffers.Get().(*[goroutineIDStackBytes]byte)
	written := runtime.Stack(stackPrefix[:], false)
	id, ok := parseGoroutineID(stackPrefix[:written])
	goroutineIDStackBuffers.Put(stackPrefix)
	if !ok {
		panic("hxrt: unable to determine current goroutine identity from runtime.Stack")
	}
	return id
}

func waitForConditionWithTimeout(mu *sync.Mutex, cond *sync.Cond, timeoutSeconds float64, consumeReady func() bool) bool {
	if consumeReady() {
		return true
	}
	var timedOut bool
	timer := time.AfterFunc(time.Duration(timeoutSeconds*float64(time.Second)), func() {
		mu.Lock()
		timedOut = true
		cond.Broadcast()
		mu.Unlock()
	})
	defer timer.Stop()
	for !timedOut {
		cond.Wait()
		if consumeReady() {
			return true
		}
	}
	return false
}
