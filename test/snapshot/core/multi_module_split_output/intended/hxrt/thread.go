package hxrt

import (
	"bytes"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type ThreadState struct {
	queueMu   sync.Mutex
	queueCond *sync.Cond
	queue     []any
	eventLoop *EventLoopHandle
}

type EventLoopHandle struct {
	mu        sync.Mutex
	cond      *sync.Cond
	oneTime   []func()
	promised  int
	regular   []*RegularEvent
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
	mu         sync.Mutex
	mutexCond  *sync.Cond
	signalCond *sync.Cond
	owner      int64
	depth      int
	waiters    int
	signaled   int
}

type SemaphoreHandle struct {
	mu    sync.Mutex
	cond  *sync.Cond
	count int
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
	mainState.eventLoop = newEventLoopHandle()
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

func newEventLoopHandle() *EventLoopHandle {
	h := &EventLoopHandle{
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
	state.eventLoop = eventLoop
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
				ThreadEventLoopLoop(state.eventLoop)
			})
		})
	}()
	return id
}

func ThreadHasEventLoop(threadID int) bool {
	available := false
	withThreadState(threadID, func(state *ThreadState) {
		available = state.eventLoop != nil
	})
	return available
}

func ThreadEvents(threadID int) *EventLoopHandle {
	var handle *EventLoopHandle
	withThreadState(threadID, func(state *ThreadState) {
		handle = state.eventLoop
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
	if state.eventLoop != nil {
		job()
		return
	}
	loop := newEventLoopHandle()
	state.eventLoop = loop
	defer func() {
		state.eventLoop = nil
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
	handle.cancelled[eventID] = true
	filtered := handle.regular[:0]
	for _, event := range handle.regular {
		if event.ID != eventID {
			filtered = append(filtered, event)
		}
	}
	handle.regular = filtered
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
	var timedOut bool
	timer := time.AfterFunc(time.Duration(timeout*float64(time.Second)), func() {
		handle.mu.Lock()
		timedOut = true
		handle.cond.Broadcast()
		handle.mu.Unlock()
	})
	defer timer.Stop()
	for !timedOut {
		handle.cond.Wait()
		if eventLoopHasReadyLocked(handle) {
			return true
		}
	}
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
	h := &ConditionHandle{}
	h.mutexCond = sync.NewCond(&h.mu)
	h.signalCond = sync.NewCond(&h.mu)
	return h
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
	handle.waiters++
	handle.mutexCond.Signal()
	for handle.signaled == 0 {
		handle.signalCond.Wait()
	}
	handle.signaled--
	handle.waiters--
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
	if handle.waiters > 0 {
		handle.signaled++
		handle.signalCond.Signal()
	}
	handle.mu.Unlock()
}

func ThreadConditionBroadcast(handle *ConditionHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	if handle.waiters > 0 {
		handle.signaled += handle.waiters
		handle.signalCond.Broadcast()
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

func currentGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	line := bytes.TrimPrefix(buf[:n], []byte("goroutine "))
	idx := bytes.IndexByte(line, ' ')
	if idx < 0 {
		return 0
	}
	id, err := strconv.ParseInt(string(line[:idx]), 10, 64)
	if err != nil {
		return 0
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
