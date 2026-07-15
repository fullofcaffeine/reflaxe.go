package hxrt

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func withStandardFiles(t *testing.T, stdinContent string, run func(stdoutPath string, stderrPath string)) {
	t.Helper()

	root := t.TempDir()
	stdinPath := filepath.Join(root, "stdin.txt")
	stdoutPath := filepath.Join(root, "stdout.txt")
	stderrPath := filepath.Join(root, "stderr.txt")
	if err := os.WriteFile(stdinPath, []byte(stdinContent), 0o600); err != nil {
		t.Fatal(err)
	}

	stdin, err := os.Open(stdinPath)
	if err != nil {
		t.Fatal(err)
	}
	stdout, err := os.Create(stdoutPath)
	if err != nil {
		stdin.Close()
		t.Fatal(err)
	}
	stderr, err := os.Create(stderrPath)
	if err != nil {
		stdin.Close()
		stdout.Close()
		t.Fatal(err)
	}

	originalStdin := os.Stdin
	originalStdout := os.Stdout
	originalStderr := os.Stderr
	os.Stdin = stdin
	os.Stdout = stdout
	os.Stderr = stderr
	t.Cleanup(func() {
		os.Stdin = originalStdin
		os.Stdout = originalStdout
		os.Stderr = originalStderr
		stdin.Close()
		stdout.Close()
		stderr.Close()
	})

	run(stdoutPath, stderrPath)
	if err := stdout.Sync(); err != nil {
		t.Fatal(err)
	}
	if err := stderr.Sync(); err != nil {
		t.Fatal(err)
	}
}

func TestSysSleepUsesSecondsAndReturnsWithinBound(t *testing.T) {
	started := time.Now()
	SysSleep(0.02)
	elapsed := time.Since(started)
	if elapsed < 5*time.Millisecond {
		t.Fatalf("SysSleep(0.02) elapsed %s, want at least 5ms", elapsed)
	}
	if elapsed > 5*time.Second {
		t.Fatalf("SysSleep(0.02) elapsed %s, want less than 5s", elapsed)
	}
}

func TestFileTextHelpersPreserveNativeErrors(t *testing.T) {
	root := t.TempDir()
	missing := root + "/missing.txt"
	if content, err := FileGetContent(&missing); err == nil || content != nil {
		t.Fatalf("missing read = (%v, %v), want nil content and error", content, err)
	}
	content := "not a directory replacement"
	if err := FileSaveContent(&root, &content); err == nil {
		t.Fatal("writing content to a directory returned success")
	}
}

func TestFileTextHelpersPreservePermissionErrors(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX file modes are not authoritative on Windows")
	}
	if os.Geteuid() == 0 {
		t.Skip("root can bypass the permission fixture")
	}

	path := t.TempDir() + "/locked.txt"
	if err := os.WriteFile(path, []byte("secret"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chmod(path, 0o600); err != nil {
			t.Error(err)
		}
	}()

	if _, err := FileGetContent(&path); err == nil {
		t.Fatal("permission-denied read returned success")
	}
	content := "replacement"
	if err := FileSaveContent(&path, &content); err == nil {
		t.Fatal("permission-denied write returned success")
	}
}

func TestSysPutEnvPreservesNativeError(t *testing.T) {
	key := "HAXE_GO=INVALID"
	value := "value"
	if err := SysPutEnv(&key, &value); err == nil {
		t.Fatal("invalid native environment key returned success")
	}
}

func TestPortableSysClockLocaleCwdAndProgramPath(t *testing.T) {
	if SysSetTimeLocale(StringFromLiteral("__haxe_go_missing_locale__")) {
		t.Fatal("unsupported process locale unexpectedly reported success")
	}

	started := SysTime()
	time.Sleep(2 * time.Millisecond)
	finished := SysTime()
	if started <= 0 || finished < started {
		t.Fatalf("SysTime() = (%f, %f), want positive nondecreasing epoch seconds", started, finished)
	}

	programPath, err := SysProgramPath()
	if err != nil {
		t.Fatal(err)
	}
	if programPath == nil || *programPath == "" || !filepath.IsAbs(*programPath) {
		t.Fatalf("SysProgramPath() = %v, want a non-empty absolute path", programPath)
	}

	original, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	temporary := t.TempDir()
	t.Cleanup(func() {
		if err := os.Chdir(original); err != nil {
			t.Error(err)
		}
	})
	if err := SysSetCwd(StringFromLiteral(temporary)); err != nil {
		t.Fatal(err)
	}
	current, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	resolvedTemporary, err := filepath.EvalSymlinks(temporary)
	if err != nil {
		t.Fatal(err)
	}
	if current != resolvedTemporary {
		t.Fatalf("cwd = %q, want %q", current, temporary)
	}
	missing := filepath.Join(temporary, "missing")
	if err := SysSetCwd(StringFromLiteral(missing)); err == nil {
		t.Fatal("SysSetCwd accepted a missing directory")
	}
}

func TestPortableSysPrintGetCharAndStandardStreams(t *testing.T) {
	withStandardFiles(t, "AB", func(stdoutPath string, stderrPath string) {
		Print(StringFromLiteral("print"))

		stdin := SysStdin()
		first, eof, err := stdin.ReadByte()
		if err != nil || eof || first != 'A' {
			t.Fatalf("stdin.ReadByte() = (%d, %t, %v), want ('A', false, nil)", first, eof, err)
		}
		if err := stdin.Close(); err != nil {
			t.Fatal(err)
		}
		second, eof, err := SysGetChar(true)
		if err != nil || eof || second != 'B' {
			t.Fatalf("SysGetChar(true) = (%d, %t, %v), want ('B', false, nil)", second, eof, err)
		}

		stdout := SysStdout()
		if err := stdout.WriteByte('!'); err != nil {
			t.Fatal(err)
		}
		if err := stdout.Flush(); err != nil {
			t.Fatal(err)
		}
		if err := stdout.Close(); err != nil {
			t.Fatal(err)
		}
		if err := stdout.WriteByte('?'); err == nil {
			t.Fatal("closed standard-stream wrapper remained attached")
		}
		if err := SysStdout().WriteByte('?'); err != nil {
			t.Fatalf("closing the wrapper closed the process stdout descriptor: %v", err)
		}

		stderr := SysStderr()
		if err := stderr.WriteByte('E'); err != nil {
			t.Fatal(err)
		}
		if err := stderr.Flush(); err != nil {
			t.Fatal(err)
		}

		stdoutBytes, err := os.ReadFile(stdoutPath)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(stdoutBytes); got != "printB!?" {
			t.Fatalf("stdout = %q, want %q", got, "printB!?")
		}
		stderrBytes, err := os.ReadFile(stderrPath)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(stderrBytes); got != "E" {
			t.Fatalf("stderr = %q, want %q", got, "E")
		}
	})
}

func TestPortableSysGetCharReportsEOF(t *testing.T) {
	withStandardFiles(t, "", func(_ string, _ string) {
		_, eof, err := SysGetChar(false)
		if err != nil || !eof {
			t.Fatalf("SysGetChar(false) = (eof %t, error %v), want (true, nil)", eof, err)
		}
	})
}
