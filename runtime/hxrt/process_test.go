package hxrt

import (
	"os"
	"runtime"
	"testing"
	"time"
)

func processTestArgs(values ...string) []*string {
	out := make([]*string, 0, len(values))
	for _, value := range values {
		value := value
		out = append(out, &value)
	}
	return out
}

func TestNewProcessPreservesStartupFailure(t *testing.T) {
	command := "__haxe_go_missing_process__"
	process, err := NewProcess(&command, []*string{})
	if err == nil {
		t.Fatal("expected missing executable to return a startup error")
	}
	if process != nil {
		t.Fatal("startup failure returned a partially initialized process")
	}
}

func TestProcessOutputSeparatesDataFromEOF(t *testing.T) {
	command := "sh"
	process, err := NewProcess(&command, processTestArgs("-c", "printf '\\n'"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := process.Close(); err != nil {
			t.Error(err)
		}
	}()

	value, eof, err := process.Stdout().ReadByte()
	if err != nil || eof || value != '\n' {
		t.Fatalf("first byte = (%d, %t, %v), want newline data", value, eof, err)
	}
	_, eof, err = process.Stdout().ReadByte()
	if err != nil || !eof {
		t.Fatalf("second read = (eof %t, %v), want normal EOF", eof, err)
	}
}

func TestProcessExitCodeNonBlocking(t *testing.T) {
	command := "sh"
	process, err := NewProcess(&command, processTestArgs("-c", "sleep 0.05; exit 7"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := process.Close(); err != nil {
			t.Error(err)
		}
	}()

	if _, available, err := process.ExitCode(false); err != nil || available {
		t.Fatalf("initial nonblocking status = (available %t, %v), want running", available, err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for {
		code, available, err := process.ExitCode(false)
		if err != nil {
			t.Fatal(err)
		}
		if available {
			if code != 7 {
				t.Fatalf("exit code = %d, want 7", code)
			}
			return
		}
		if time.Now().After(deadline) {
			t.Fatal("process did not become available before deadline")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestProcessCloseDoesNotKillChild(t *testing.T) {
	marker := t.TempDir() + "/closed-process-finished"
	command := "sh"
	process, err := NewProcess(&command, processTestArgs("-c", "sleep 0.05; printf done > \"$1\"", "sh", marker))
	if err != nil {
		t.Fatal(err)
	}
	if err := process.Close(); err != nil {
		t.Fatal(err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for {
		if _, err := os.Stat(marker); err == nil {
			return
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		if time.Now().After(deadline) {
			t.Fatal("closing the process prevented the child from completing")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestProcessLifecycleDoesNotLeakGoroutines(t *testing.T) {
	baseline := runtime.NumGoroutine()
	command := "sh"
	for i := 0; i < 1000; i++ {
		process, err := NewProcess(&command, processTestArgs("-c", "exit 0"))
		if err != nil {
			t.Fatalf("start process %d: %v", i, err)
		}
		code, available, err := process.ExitCode(true)
		if err != nil || !available || code != 0 {
			t.Fatalf("wait process %d = (%d, %t, %v)", i, code, available, err)
		}
		if err := process.Close(); err != nil {
			t.Fatalf("close process %d: %v", i, err)
		}
	}

	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > baseline+2 {
		if time.Now().After(deadline) {
			t.Fatalf("goroutines after lifecycle stress = %d, baseline = %d", runtime.NumGoroutine(), baseline)
		}
		runtime.Gosched()
	}
}
