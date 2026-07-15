package hxrt

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	threadFailureProbeEnv    = "HAXE_GO_THREAD_FAILURE_PROBE"
	detachedFirstUseProbeEnv = "HAXE_GO_DETACHED_FIRST_USE_PROBE"
)

type notifyingLocker struct {
	sync.Locker
	once     sync.Once
	unlocked chan struct{}
}

func (locker *notifyingLocker) Unlock() {
	locker.once.Do(func() { close(locker.unlocked) })
	locker.Locker.Unlock()
}

func TestPortableThreadFailureProcessContract(t *testing.T) {
	switch os.Getenv(threadFailureProbeEnv) {
	case "haxe":
		ThreadSpawn(func() {
			Throw(StringFromLiteral("child-failure"))
		})
		ThreadWaitForAll()
		fmt.Fprintln(os.Stdout, "main-survived")
		os.Exit(0)
	case "native":
		ThreadSpawn(func() {
			panic("native-thread-panic")
		})
		ThreadWaitForAll()
		os.Exit(0)
	}

	for _, probe := range []struct {
		name         string
		wantSuccess  bool
		wantStdout   string
		wantStderr   string
		rejectStderr string
	}{
		{
			name:         "haxe",
			wantSuccess:  true,
			wantStdout:   "main-survived\n",
			wantStderr:   "Uncaught exception child-failure\n",
			rejectStderr: "panic:",
		},
		{
			name:         "native",
			wantSuccess:  false,
			wantStderr:   "panic: native-thread-panic",
			rejectStderr: "Uncaught exception",
		},
	} {
		t.Run(probe.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestPortableThreadFailureProcessContract$")
			cmd.Env = append(os.Environ(), threadFailureProbeEnv+"="+probe.name)
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			if ctx.Err() != nil {
				t.Fatalf("thread failure probe timed out: %v", ctx.Err())
			}
			if probe.wantSuccess && err != nil {
				t.Fatalf("thread failure probe failed: %v\nstderr:\n%s", err, stderr.String())
			}
			if !probe.wantSuccess && err == nil {
				t.Fatal("native panic probe returned success")
			}
			if stdout.String() != probe.wantStdout {
				t.Fatalf("stdout = %q, want %q", stdout.String(), probe.wantStdout)
			}
			if !strings.Contains(stderr.String(), probe.wantStderr) {
				t.Fatalf("stderr = %q, want substring %q", stderr.String(), probe.wantStderr)
			}
			if strings.Contains(stderr.String(), probe.rejectStderr) {
				t.Fatalf("stderr = %q, reject substring %q", stderr.String(), probe.rejectStderr)
			}
		})
	}
}

func TestThreadWaitForAllDrainsNestedJobsAndState(t *testing.T) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	baselineStates := len(threadStates)
	baselineMappings := len(goroutineThreadIDs)
	threadRuntimeMu.Unlock()

	const roots = 32
	const childrenPerRoot = 8
	var completed atomic.Int64
	for root := 0; root < roots; root++ {
		ThreadSpawn(func() {
			completed.Add(1)
			for child := 0; child < childrenPerRoot; child++ {
				ThreadSpawn(func() {
					completed.Add(1)
					ThreadSpawn(func() {
						completed.Add(1)
					})
				})
			}
		})
	}
	ThreadSpawnWithEventLoop(func() {
		completed.Add(1)
	})

	ThreadWaitForAll()
	wantCompleted := int64(roots*(1+childrenPerRoot*2) + 1)
	if got := completed.Load(); got != wantCompleted {
		t.Fatalf("completed jobs = %d, want %d", got, wantCompleted)
	}

	threadRuntimeMu.Lock()
	gotStates := len(threadStates)
	gotMappings := len(goroutineThreadIDs)
	gotActive := activePortableThreads
	threadRuntimeMu.Unlock()
	if gotActive != 0 {
		t.Fatalf("active portable threads = %d, want 0", gotActive)
	}
	if gotStates != baselineStates || gotMappings != baselineMappings {
		t.Fatalf(
			"thread runtime state after drain = (%d states, %d mappings), want (%d, %d)",
			gotStates,
			gotMappings,
			baselineStates,
			baselineMappings,
		)
	}
}

func TestThreadEventLoopCancellationStateReturnsToBaseline(t *testing.T) {
	handle := newEventLoopHandle()
	const cancellations = 100_000

	for range cancellations {
		eventID := ThreadEventLoopRepeat(handle, func() {}, 60_000)
		ThreadEventLoopCancel(handle, eventID)
	}

	// Callers may race event completion and cannot know whether an event is
	// still registered when they request cancellation. An unknown identifier
	// therefore has to remain a bounded no-op.
	ThreadEventLoopCancel(handle, cancellations+1)

	handle.mu.Lock()
	defer handle.mu.Unlock()
	if got := len(handle.regular); got != 0 {
		t.Fatalf("cancelled regular events still registered: got %d, want 0", got)
	}
	if got := len(handle.cancelled); got != 0 {
		t.Fatalf("cancellation markers retained after events were removed: got %d, want 0", got)
	}
}

func TestThreadEventLoopCancelWhileCallbackRunsSuppressesReschedule(t *testing.T) {
	handle := newEventLoopHandle()
	started := make(chan struct{})
	release := make(chan struct{})
	progressDone := make(chan struct{})

	eventID := ThreadEventLoopRepeat(handle, func() {
		close(started)
		<-release
	}, 60_000)
	handle.mu.Lock()
	handle.regular[0].NextRun = threadNowSeconds() - 1
	handle.mu.Unlock()

	go func() {
		defer close(progressDone)
		ThreadEventLoopProgress(handle)
	}()

	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("repeating callback did not start")
	}
	ThreadEventLoopCancel(handle, eventID)

	handle.mu.Lock()
	if !handle.running[eventID] || !handle.cancelled[eventID] {
		handle.mu.Unlock()
		t.Fatal("running cancellation was not retained until callback completion")
	}
	handle.mu.Unlock()

	close(release)
	select {
	case <-progressDone:
	case <-time.After(time.Second):
		t.Fatal("event-loop progress did not finish")
	}

	handle.mu.Lock()
	defer handle.mu.Unlock()
	if got := len(handle.regular); got != 0 {
		t.Fatalf("cancelled callback was rescheduled: got %d events, want 0", got)
	}
	if got := len(handle.running); got != 0 {
		t.Fatalf("running event state retained after callback: got %d, want 0", got)
	}
	if got := len(handle.cancelled); got != 0 {
		t.Fatalf("cancellation state retained after callback: got %d, want 0", got)
	}
}

func TestThreadEventLoopCancellationWakesTimedWait(t *testing.T) {
	handle := newEventLoopHandle()
	eventID := ThreadEventLoopRepeat(handle, func() {}, 60_000)
	waiting := make(chan struct{})
	handle.cond = sync.NewCond(&notifyingLocker{Locker: &handle.mu, unlocked: waiting})
	waitResult := make(chan bool, 1)

	go func() {
		waitResult <- ThreadEventLoopWaitTimeout(handle, 2)
	}()
	<-waiting
	ThreadEventLoopCancel(handle, eventID)

	select {
	case pending := <-waitResult:
		if pending {
			t.Fatal("timed wait reported pending work after its final event was cancelled")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("cancelling the final event did not wake the timed event-loop wait")
	}
}

func TestThreadEventLoopEarlierTimerWakesTimedWaitForRecalculation(t *testing.T) {
	handle := newEventLoopHandle()
	ThreadEventLoopRepeat(handle, func() {}, 60_000)
	waiting := make(chan struct{})
	handle.cond = sync.NewCond(&notifyingLocker{Locker: &handle.mu, unlocked: waiting})
	waitResult := make(chan bool, 1)

	go func() {
		waitResult <- ThreadEventLoopWaitTimeout(handle, 2)
	}()
	<-waiting
	ThreadEventLoopRepeat(handle, func() {}, 30_000)

	select {
	case pending := <-waitResult:
		if !pending {
			t.Fatal("timed wait lost pending work after an earlier timer was inserted")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("an earlier timer did not wake the timed event-loop wait for deadline recalculation")
	}
}

func TestThreadConditionBroadcastDoesNotWakeLateWaiter(t *testing.T) {
	handle := ThreadConditionNew()
	handle.mu.Lock()
	first := registerConditionWaiterLocked(handle)
	second := registerConditionWaiterLocked(handle)
	handle.mu.Unlock()

	ThreadConditionBroadcast(handle)
	ThreadConditionBroadcast(handle)

	handle.mu.Lock()
	late := registerConditionWaiterLocked(handle)
	if !conditionWaiterSignaledLocked(handle, first) {
		t.Error("first waiter did not receive broadcast")
	}
	if !conditionWaiterSignaledLocked(handle, second) {
		t.Error("second waiter did not receive broadcast")
	}
	if conditionWaiterSignaledLocked(handle, late) {
		t.Error("late waiter consumed an earlier broadcast generation")
	}
	handle.mu.Unlock()

	ThreadConditionSignal(handle)
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if !conditionWaiterSignaledLocked(handle, late) {
		t.Error("late waiter did not receive its own signal")
	}
}

func TestThreadConditionSignalDoesNotLeaveDuplicateCredits(t *testing.T) {
	handle := ThreadConditionNew()
	handle.mu.Lock()
	first := registerConditionWaiterLocked(handle)
	second := registerConditionWaiterLocked(handle)
	handle.mu.Unlock()

	ThreadConditionSignal(handle)
	ThreadConditionSignal(handle)
	ThreadConditionSignal(handle)

	handle.mu.Lock()
	if !conditionWaiterSignaledLocked(handle, first) || !conditionWaiterSignaledLocked(handle, second) {
		handle.mu.Unlock()
		t.Fatal("two signals did not mark the two existing waiters")
	}
	removeConditionWaiterLocked(handle, first)
	removeConditionWaiterLocked(handle, second)
	late := registerConditionWaiterLocked(handle)
	defer handle.mu.Unlock()
	if conditionWaiterSignaledLocked(handle, late) {
		t.Error("duplicate signal credit leaked to a later waiter")
	}
}

func TestThreadEventLoopOwnershipIsRaceSafe(t *testing.T) {
	const iterations = 1_000
	begin := make(chan struct{})
	loopReady := make(chan struct{})
	advance := make(chan struct{})
	readerStarted := make(chan struct{})
	stopReader := make(chan struct{})
	readerDone := make(chan struct{})

	workerID := ThreadSpawn(func() {
		<-begin
		for range iterations {
			ThreadRunWithEventLoop(func() {
				loopReady <- struct{}{}
				<-advance
			})
		}
	})

	go func() {
		defer close(readerDone)
		ThreadHasEventLoop(workerID)
		ThreadEvents(workerID)
		close(readerStarted)
		for {
			select {
			case <-stopReader:
				return
			default:
				ThreadHasEventLoop(workerID)
				ThreadEvents(workerID)
			}
		}
	}()

	<-readerStarted
	close(begin)
	for range iterations {
		<-loopReady
		if !ThreadHasEventLoop(workerID) {
			t.Fatal("worker did not publish its active event loop")
		}
		if ThreadEvents(workerID) == nil {
			t.Fatal("worker event-loop handle was nil while active")
		}
		advance <- struct{}{}
	}
	ThreadWaitForAll()
	close(stopReader)
	<-readerDone
}

func waitForConditionWaiterCount(t *testing.T, handle *ConditionHandle, want int) {
	t.Helper()
	timedOut := false
	timer := time.AfterFunc(2*time.Second, func() {
		handle.mu.Lock()
		timedOut = true
		handle.mutexCond.Broadcast()
		handle.mu.Unlock()
	})

	handle.mu.Lock()
	for len(handle.waiters) != want && !timedOut {
		handle.mutexCond.Wait()
	}
	got := len(handle.waiters)
	handle.mu.Unlock()
	timer.Stop()
	if got != want {
		t.Fatalf("registered condition waiters = %d, want %d", got, want)
	}
}

func receiveConditionWakeups(t *testing.T, woke <-chan int, count int) {
	t.Helper()
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	for range count {
		select {
		case <-woke:
		case <-timer.C:
			t.Fatalf("condition woke fewer than %d expected waiters", count)
		}
	}
}

func TestThreadConditionSignalBroadcastAndLateWaiter(t *testing.T) {
	for _, gomaxprocs := range []int{1, 2, 8} {
		t.Run(fmt.Sprintf("gomaxprocs=%d", gomaxprocs), func(t *testing.T) {
			previous := runtime.GOMAXPROCS(gomaxprocs)
			defer runtime.GOMAXPROCS(previous)

			handle := ThreadConditionNew()
			const originalWaiters = 32
			woke := make(chan int, originalWaiters+1)
			startWaiter := func(id int) {
				go func() {
					ThreadConditionAcquire(handle)
					ThreadConditionWait(handle)
					ThreadConditionRelease(handle)
					woke <- id
				}()
			}

			for id := 0; id < originalWaiters; id++ {
				startWaiter(id)
			}
			waitForConditionWaiterCount(t, handle, originalWaiters)

			ThreadConditionSignal(handle)
			receiveConditionWakeups(t, woke, 1)
			waitForConditionWaiterCount(t, handle, originalWaiters-1)

			ThreadConditionBroadcast(handle)
			receiveConditionWakeups(t, woke, originalWaiters-1)
			waitForConditionWaiterCount(t, handle, 0)

			startWaiter(originalWaiters)
			waitForConditionWaiterCount(t, handle, 1)
			handle.mu.Lock()
			for waiterID := range handle.waiters {
				if conditionWaiterSignaledLocked(handle, waiterID) {
					handle.mu.Unlock()
					t.Fatal("late waiter inherited the earlier broadcast")
				}
			}
			handle.mu.Unlock()

			ThreadConditionSignal(handle)
			receiveConditionWakeups(t, woke, 1)
			waitForConditionWaiterCount(t, handle, 0)
		})
	}
}

func TestParseGoroutineID(t *testing.T) {
	for _, testCase := range []struct {
		name  string
		stack string
		want  int64
		ok    bool
	}{
		{name: "ordinary", stack: "goroutine 42 [running]:\n", want: 42, ok: true},
		{name: "largest", stack: "goroutine 9223372036854775807 [runnable]:\n", want: 9223372036854775807, ok: true},
		{name: "missing prefix", stack: "thread 42 [running]:\n"},
		{name: "missing delimiter", stack: "goroutine 42"},
		{name: "zero reserved", stack: "goroutine 0 [running]:\n"},
		{name: "negative", stack: "goroutine -1 [running]:\n"},
		{name: "not numeric", stack: "goroutine unknown [running]:\n"},
		{name: "overflow", stack: "goroutine 9223372036854775808 [running]:\n"},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			got, ok := parseGoroutineID([]byte(testCase.stack))
			if got != testCase.want || ok != testCase.ok {
				t.Fatalf("parseGoroutineID(%q) = (%d, %t), want (%d, %t)", testCase.stack, got, ok, testCase.want, testCase.ok)
			}
		})
	}
}

func TestCurrentGoroutineIDIsPositiveAndStable(t *testing.T) {
	first := currentGoroutineID()
	if first <= 0 {
		t.Fatalf("current goroutine ID = %d, want a positive identity", first)
	}
	for range 1_000 {
		if got := currentGoroutineID(); got != first {
			t.Fatalf("current goroutine ID changed from %d to %d", first, got)
		}
	}
}

func TestPortableThreadIdentityStateReturnsToBaselineAfterChurn(t *testing.T) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	baselineStates := len(threadStates)
	baselineMappings := len(goroutineThreadIDs)
	threadRuntimeMu.Unlock()

	const batches = 10
	const threadsPerBatch = 1_000
	var completed atomic.Int64
	for range batches {
		for range threadsPerBatch {
			ThreadSpawn(func() {
				completed.Add(1)
			})
		}
		ThreadWaitForAll()
		threadRuntimeMu.Lock()
		gotStates := len(threadStates)
		gotMappings := len(goroutineThreadIDs)
		threadRuntimeMu.Unlock()
		if gotStates != baselineStates || gotMappings != baselineMappings {
			t.Fatalf(
				"identity state after drain = (%d states, %d mappings), want baseline (%d, %d)",
				gotStates,
				gotMappings,
				baselineStates,
				baselineMappings,
			)
		}
	}

	if got, want := completed.Load(), int64(batches*threadsPerBatch); got != want {
		t.Fatalf("completed portable threads = %d, want %d", got, want)
	}
}

func threadRuntimeStateCounts() (states int, mappings int, localValues int, active int) {
	ensureThreadRuntime()
	threadRuntimeMu.Lock()
	stateSnapshot := make([]*ThreadState, 0, len(threadStates))
	for _, state := range threadStates {
		stateSnapshot = append(stateSnapshot, state)
	}
	states = len(threadStates)
	mappings = len(goroutineThreadIDs)
	active = activePortableThreads
	threadRuntimeMu.Unlock()

	for _, state := range stateSnapshot {
		state.threadLocalMu.RLock()
		localValues += len(state.threadLocalValues)
		state.threadLocalMu.RUnlock()
	}
	return states, mappings, localValues, active
}

func TestDetachedFirstUseInitializesRuntimeOnCaller(t *testing.T) {
	if os.Getenv(detachedFirstUseProbeEnv) == "1" {
		callerGoroutineID := currentGoroutineID()
		handle := ThreadLocalNew()
		callbackGoroutineID := make(chan int64, 1)
		ThreadSpawnDetached(func() {
			callbackGoroutineID <- currentGoroutineID()
			ThreadLocalSet(handle, "detached-first-use")
		})
		detachedGoroutineID := <-callbackGoroutineID

		deadline := time.Now().Add(5 * time.Second)
		for time.Now().Before(deadline) {
			states, mappings, values, active := threadRuntimeStateCounts()
			threadRuntimeMu.Lock()
			callerID, callerExists := goroutineThreadIDs[callerGoroutineID]
			_, detachedExists := goroutineThreadIDs[detachedGoroutineID]
			threadRuntimeMu.Unlock()
			if callerExists && callerID == 0 && !detachedExists && states == 1 && mappings == 1 && values == 0 && active == 0 {
				return
			}
			runtime.Gosched()
		}

		states, mappings, values, active := threadRuntimeStateCounts()
		threadRuntimeMu.Lock()
		callerID, callerExists := goroutineThreadIDs[callerGoroutineID]
		_, detachedExists := goroutineThreadIDs[detachedGoroutineID]
		threadRuntimeMu.Unlock()
		t.Fatalf(
			"detached first-use state = (%d states, %d mappings, %d values, %d active, caller=%d/%t, detached=%t), want (1, 1, 0, 0, caller=0/true, detached=false)",
			states,
			mappings,
			values,
			active,
			callerID,
			callerExists,
			detachedExists,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestDetachedFirstUseInitializesRuntimeOnCaller$")
	cmd.Env = append(os.Environ(), detachedFirstUseProbeEnv+"=1")
	output, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		t.Fatalf("detached first-use probe timed out: %v", ctx.Err())
	}
	if err != nil {
		t.Fatalf("detached first-use probe failed: %v\n%s", err, output)
	}
}

func TestDetachedNilJobRetainsNativePanic(t *testing.T) {
	baselineStates, baselineMappings, baselineValues, baselineActive := threadRuntimeStateCounts()
	recovered := make(chan any, 1)

	go func() {
		defer func() {
			recovered <- recover()
		}()
		runDetachedThreadJob(nil)
	}()

	select {
	case got := <-recovered:
		if got == nil {
			t.Fatal("nil detached callback returned instead of preserving native panic behavior")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("nil detached callback panic timed out")
	}

	gotStates, gotMappings, gotValues, gotActive := threadRuntimeStateCounts()
	if gotStates != baselineStates || gotMappings != baselineMappings || gotValues != baselineValues || gotActive != baselineActive {
		t.Fatalf(
			"thread-local runtime state after nil callback panic = (%d states, %d mappings, %d values, %d active), want baseline (%d, %d, %d, %d)",
			gotStates,
			gotMappings,
			gotValues,
			gotActive,
			baselineStates,
			baselineMappings,
			baselineValues,
			baselineActive,
		)
	}
}

func TestThreadLocalStateReturnsToBaselineAfterPortableThreadChurn(t *testing.T) {
	baselineStates, baselineMappings, baselineValues, baselineActive := threadRuntimeStateCounts()
	handle := ThreadLocalNew()

	const workers = 1_000
	failures := make(chan string, workers*2)
	for worker := 0; worker < workers; worker++ {
		worker := worker
		ThreadSpawn(func() {
			if got := ThreadLocalGet(handle); got != nil {
				failures <- fmt.Sprintf("worker %d inherited thread-local value %v", worker, got)
			}
			value := fmt.Sprintf("portable-%d", worker)
			ThreadLocalSet(handle, value)
			if got := ThreadLocalGet(handle); got != value {
				failures <- fmt.Sprintf("worker %d thread-local value = %v, want %q", worker, got, value)
			}
		})
	}
	ThreadWaitForAll()
	close(failures)
	for failure := range failures {
		t.Error(failure)
	}

	gotStates, gotMappings, gotValues, gotActive := threadRuntimeStateCounts()
	if gotStates != baselineStates || gotMappings != baselineMappings || gotValues != baselineValues || gotActive != baselineActive {
		t.Fatalf(
			"thread-local runtime state after portable churn = (%d states, %d mappings, %d values, %d active), want baseline (%d, %d, %d, %d)",
			gotStates,
			gotMappings,
			gotValues,
			gotActive,
			baselineStates,
			baselineMappings,
			baselineValues,
			baselineActive,
		)
	}
}

func TestDetachedThreadIdentityAndLocalStateReturnsToBaseline(t *testing.T) {
	baselineStates, baselineMappings, baselineValues, baselineActive := threadRuntimeStateCounts()
	handle := ThreadLocalNew()

	const workers = 1_000
	var completed sync.WaitGroup
	completed.Add(workers)
	failures := make(chan string, workers*2)
	for worker := 0; worker < workers; worker++ {
		worker := worker
		go func() {
			defer completed.Done()
			runDetachedThreadJob(func() {
				if id := ThreadCurrentId(); id == 0 {
					failures <- fmt.Sprintf("worker %d received reserved thread id", worker)
				}
				if got := ThreadLocalGet(handle); got != nil {
					failures <- fmt.Sprintf("worker %d inherited detached thread-local value %v", worker, got)
				}
				ThreadLocalSet(handle, worker)
				if got := ThreadLocalGet(handle); got != worker {
					failures <- fmt.Sprintf("worker %d detached thread-local value = %v", worker, got)
				}
			})
		}()
	}
	completed.Wait()
	close(failures)
	for failure := range failures {
		t.Error(failure)
	}

	gotStates, gotMappings, gotValues, gotActive := threadRuntimeStateCounts()
	if gotStates != baselineStates || gotMappings != baselineMappings || gotValues != baselineValues || gotActive != baselineActive {
		t.Fatalf(
			"thread-local runtime state after detached churn = (%d states, %d mappings, %d values, %d active), want baseline (%d, %d, %d, %d)",
			gotStates,
			gotMappings,
			gotValues,
			gotActive,
			baselineStates,
			baselineMappings,
			baselineValues,
			baselineActive,
		)
	}
}

func TestDetachedThreadCleanupRunsDuringNativePanicUnwind(t *testing.T) {
	baselineStates, baselineMappings, baselineValues, baselineActive := threadRuntimeStateCounts()
	handle := ThreadLocalNew()
	panicMarker := &struct{}{}
	recovered := make(chan any, 1)

	go func() {
		defer func() {
			recovered <- recover()
		}()
		runDetachedThreadJob(func() {
			ThreadLocalSet(handle, panicMarker)
			panic(panicMarker)
		})
	}()

	select {
	case got := <-recovered:
		if got != panicMarker {
			t.Fatalf("detached native panic = %v, want original marker", got)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("detached native panic cleanup timed out")
	}

	gotStates, gotMappings, gotValues, gotActive := threadRuntimeStateCounts()
	if gotStates != baselineStates || gotMappings != baselineMappings || gotValues != baselineValues || gotActive != baselineActive {
		t.Fatalf(
			"thread-local runtime state after native panic = (%d states, %d mappings, %d values, %d active), want baseline (%d, %d, %d, %d)",
			gotStates,
			gotMappings,
			gotValues,
			gotActive,
			baselineStates,
			baselineMappings,
			baselineValues,
			baselineActive,
		)
	}
}

func BenchmarkCurrentGoroutineID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = currentGoroutineID()
	}
}
