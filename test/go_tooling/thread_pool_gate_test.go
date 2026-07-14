package main

import (
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"snapshot/hxrt"
)

type generatedThreadPool interface {
	get_isShutdown() bool
	run(func())
	shutdown()
}

func submitGeneratedTask(pool generatedThreadPool, task func()) (accepted bool) {
	accepted = true
	hxrt.TryCatch(func() {
		pool.run(task)
	}, func(caught any) {
		if _, rejected := caught.(*sys__thread__ThreadPoolException); rejected {
			accepted = false
			return
		}
		hxrt.Throw(caught)
	})
	return accepted
}

func waitForPoolBlockers(t *testing.T, started <-chan struct{}, count int) {
	t.Helper()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for range count {
		select {
		case <-started:
		case <-timer.C:
			t.Fatalf("only part of the worker barrier started: want %d workers", count)
		}
	}
}

func assertConcurrentRunShutdownContract(t *testing.T, pool generatedThreadPool, workerCount int) {
	t.Helper()
	const submissions = 10_000
	const submitters = 64
	const shutdownCallers = 8

	blockerStarted := make(chan struct{}, workerCount)
	releaseBlockers := make(chan struct{})
	for range workerCount {
		if !submitGeneratedTask(pool, func() {
			blockerStarted <- struct{}{}
			<-releaseBlockers
		}) {
			t.Fatal("fresh pool rejected a worker-barrier task")
		}
	}
	waitForPoolBlockers(t, blockerStarted, workerCount)

	accepted := make([]atomic.Bool, submissions)
	executions := make([]atomic.Int32, submissions)
	var duplicate atomic.Bool
	start := make(chan struct{})
	var calls sync.WaitGroup
	calls.Add(submitters + shutdownCallers)

	for submitter := 0; submitter < submitters; submitter++ {
		submitter := submitter
		go func() {
			defer calls.Done()
			<-start
			for taskID := submitter; taskID < submissions; taskID += submitters {
				taskID := taskID
				if submitGeneratedTask(pool, func() {
					if executions[taskID].Add(1) != 1 {
						duplicate.Store(true)
					}
				}) {
					accepted[taskID].Store(true)
				}
			}
		}()
	}
	for range shutdownCallers {
		go func() {
			defer calls.Done()
			<-start
			pool.shutdown()
		}()
	}

	close(start)
	calls.Wait()
	close(releaseBlockers)
	hxrt.ThreadWaitForAll()

	if !pool.get_isShutdown() {
		t.Fatal("pool did not publish shutdown state")
	}
	if submitGeneratedTask(pool, func() {}) {
		t.Fatal("pool accepted a task after shutdown completed")
	}
	if duplicate.Load() {
		t.Fatal("an accepted task executed more than once")
	}

	acceptedCount := 0
	executedCount := 0
	for taskID := range submissions {
		if accepted[taskID].Load() {
			acceptedCount++
		}
		executed := executions[taskID].Load()
		if executed != 0 {
			executedCount++
		}
		want := int32(0)
		if accepted[taskID].Load() {
			want = 1
		}
		if executed != want {
			t.Fatalf("task %d executions = %d, want %d from its acceptance result", taskID, executed, want)
		}
	}
	if executedCount != acceptedCount {
		t.Fatalf("executed tasks = %d, accepted tasks = %d", executedCount, acceptedCount)
	}
}

func TestGeneratedThreadPoolsConcurrentRunShutdownExactlyOnce(t *testing.T) {
	for _, gomaxprocs := range []int{1, 2, 8} {
		name := strconv.Itoa(gomaxprocs)
		t.Run("fixed/gomaxprocs="+name, func(t *testing.T) {
			previous := runtime.GOMAXPROCS(gomaxprocs)
			defer runtime.GOMAXPROCS(previous)
			assertConcurrentRunShutdownContract(t, New_sys__thread__FixedThreadPool(4), 4)
		})
		t.Run("elastic/gomaxprocs="+name, func(t *testing.T) {
			previous := runtime.GOMAXPROCS(gomaxprocs)
			defer runtime.GOMAXPROCS(previous)
			assertConcurrentRunShutdownContract(t, New_sys__thread__ElasticThreadPool(4, 0.01), 4)
		})
	}
}

func assertGeneratedPoolReplacesFailedWorker(t *testing.T, pool generatedThreadPool) {
	t.Helper()
	var completed atomic.Int32
	if !submitGeneratedTask(pool, func() {
		hxrt.Throw(hxrt.StringFromLiteral("expected worker failure"))
	}) {
		t.Fatal("fresh pool rejected the throwing task")
	}
	if !submitGeneratedTask(pool, func() {
		completed.Add(1)
	}) {
		t.Fatal("pool rejected work queued after the throwing task")
	}
	pool.shutdown()
	hxrt.ThreadWaitForAll()
	if got := completed.Load(); got != 1 {
		t.Fatalf("work after a throwing task executed %d times, want exactly once", got)
	}
}

func TestGeneratedThreadPoolsReplaceWorkerAfterHaxeThrow(t *testing.T) {
	t.Run("fixed", func(t *testing.T) {
		assertGeneratedPoolReplacesFailedWorker(t, New_sys__thread__FixedThreadPool(1))
	})
	t.Run("elastic", func(t *testing.T) {
		assertGeneratedPoolReplacesFailedWorker(t, New_sys__thread__ElasticThreadPool(1, 0.01))
	})
}
