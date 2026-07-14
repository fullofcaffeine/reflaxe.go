package hxrt

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const threadFailureProbeEnv = "HAXE_GO_THREAD_FAILURE_PROBE"

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
