package hxrt

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
	"time"
)

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

func ThreadCurrentId() int {
	return int(currentGoroutineID())
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

func ThreadLockWaitTimeoutAny(handle *LockHandle, timeout any) bool {
	return ThreadLockWaitTimeout(handle, threadTimeoutSeconds(timeout))
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

func ThreadSemaphoreTryAcquireTimeoutAny(handle *SemaphoreHandle, timeout any) bool {
	return ThreadSemaphoreTryAcquireTimeout(handle, threadTimeoutSeconds(timeout))
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

func threadTimeoutSeconds(timeout any) float64 {
	switch value := timeout.(type) {
	case float64:
		return value
	case float32:
		return float64(value)
	case int:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	default:
		return 0
	}
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
