package hxrt

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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

// expectFileSystemThrow verifies that a runtime filesystem capability preserves
// the catchable Haxe exception boundary for invalid native operations.
func expectFileSystemThrow(t *testing.T, run func()) {
	t.Helper()
	didThrow := false
	func() {
		defer func() {
			if _, ok := recover().(HaxeException); ok {
				didThrow = true
			}
		}()
		run()
	}()
	if !didThrow {
		t.Fatal("filesystem operation did not throw a HaxeException")
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

func TestTypedFileCapabilitiesPreserveBinaryStreamsAndSeek(t *testing.T) {
	path := filepath.Join(t.TempDir(), "typed-file.bin")
	pathValue := StringFromLiteral(path)
	values := []int{0, 255, 128, 65}
	FileWriteByteValues(pathValue, values)
	if got := FileReadByteValues(pathValue); !slices.Equal(got, values) {
		t.Fatalf("FileReadByteValues() = %v, want %v", got, values)
	}

	output := FileOpenUpdate(pathValue)
	if pos := FileOutputTell(output); pos != 0 {
		t.Fatalf("FileOutputTell() = %d, want 0", pos)
	}
	FileOutputSeek(output, 1, io.SeekStart)
	if written := FileOutputWriteValues(output, []int{7, 8, 9}, 1, 2); written != 2 {
		t.Fatalf("FileOutputWriteValues() = %d, want 2", written)
	}
	FileOutputFlush(output)
	FileOutputClose(output)

	input := FileOpenInput(pathValue)
	if value := FileInputReadByteValue(input); value != 0 {
		t.Fatalf("FileInputReadByteValue() = %d, want 0", value)
	}
	if got := FileInputReadValues(input, 8); !slices.Equal(got, []int{8, 9, 65}) {
		t.Fatalf("FileInputReadValues() = %v, want [8 9 65]", got)
	}
	if !FileInputEof(input) {
		t.Fatal("FileInputEof() = false at end of file")
	}
	FileInputSeek(input, 1, io.SeekStart)
	if pos := FileInputTell(input); pos != 1 {
		t.Fatalf("FileInputTell() = %d, want 1", pos)
	}
	FileInputClose(input)
}

func TestFileSystemCapabilitiesPreservePathsMetadataAndMutation(t *testing.T) {
	root := filepath.Join(t.TempDir(), "nested", "filesystem")
	rootPtr := StringFromLiteral(root)
	if FileSystemExists(rootPtr) {
		t.Fatalf("FileSystemExists(%q) = true before creation", root)
	}

	FileSystemCreateDirectory(rootPtr)
	if !FileSystemExists(rootPtr) || !FileSystemIsDirectory(rootPtr) {
		t.Fatalf("created directory %q is not visible as a directory", root)
	}

	missing := filepath.Join(root, "missing", "child.txt")
	if FileSystemIsDirectory(StringFromLiteral(missing)) {
		t.Fatalf("FileSystemIsDirectory(%q) = true for a missing path", missing)
	}
	absolute := FileSystemAbsolutePath(StringFromLiteral(missing))
	if absolute == nil || !filepath.IsAbs(*absolute) {
		t.Fatalf("FileSystemAbsolutePath(%q) = %v, want an absolute path", missing, absolute)
	}

	canonical := FileSystemFullPath(rootPtr)
	if canonical == nil || !filepath.IsAbs(*canonical) {
		t.Fatalf("FileSystemFullPath(%q) = %v, want an absolute path", root, canonical)
	}
	if runtime.GOOS != "windows" {
		link := filepath.Join(filepath.Dir(root), "filesystem-link")
		if err := os.Symlink(root, link); err != nil {
			t.Fatal(err)
		}
		resolvedLink := FileSystemFullPath(StringFromLiteral(link))
		expectedRoot, err := filepath.EvalSymlinks(root)
		if err != nil {
			t.Fatal(err)
		}
		if resolvedLink == nil || filepath.Clean(*resolvedLink) != filepath.Clean(expectedRoot) {
			t.Fatalf("FileSystemFullPath(%q) = %v, want %q", link, resolvedLink, expectedRoot)
		}
	}

	from := filepath.Join(root, "from.txt")
	to := filepath.Join(root, "to.txt")
	if err := os.WriteFile(from, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	FileSystemRename(StringFromLiteral(from), StringFromLiteral(to))
	if FileSystemExists(StringFromLiteral(from)) || !FileSystemExists(StringFromLiteral(to)) {
		t.Fatal("FileSystemRename did not move the file")
	}

	stat := FileSystemStatPath(StringFromLiteral(to))
	if stat == nil || stat.Size != 5 || stat.Nlink < 1 || stat.Mode == 0 || stat.MtimeMs <= 0 {
		t.Fatalf("FileSystemStatPath(%q) = %#v", to, stat)
	}

	entries := FileSystemReadDirectory(rootPtr)
	if len(entries) != 1 || entries[0] == nil || *entries[0] != "to.txt" {
		t.Fatalf("FileSystemReadDirectory(%q) = %v", root, entries)
	}

	wrongDirectory := filepath.Join(root, "not-a-file")
	FileSystemCreateDirectory(StringFromLiteral(wrongDirectory))
	expectFileSystemThrow(t, func() {
		FileSystemDeleteFile(StringFromLiteral(wrongDirectory))
	})
	if !FileSystemIsDirectory(StringFromLiteral(wrongDirectory)) {
		t.Fatal("FileSystemDeleteFile removed a directory")
	}
	FileSystemDeleteDirectory(StringFromLiteral(wrongDirectory))

	expectFileSystemThrow(t, func() {
		FileSystemDeleteDirectory(StringFromLiteral(to))
	})
	if !FileSystemExists(StringFromLiteral(to)) {
		t.Fatal("FileSystemDeleteDirectory removed a file")
	}

	FileSystemDeleteFile(StringFromLiteral(to))
	FileSystemDeleteDirectory(rootPtr)
	if FileSystemExists(rootPtr) {
		t.Fatalf("FileSystemDeleteDirectory(%q) left the directory behind", root)
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

func TestStagedSysCapabilitiesPreserveEnvironmentPathsAndShellCommands(t *testing.T) {
	key := StringFromLiteral("HAXE_GO_STAGED_SYS_TEST")
	value := StringFromLiteral("typed")
	SysSetEnvironment(key, value)
	t.Cleanup(func() { SysSetEnvironment(key, nil) })
	if got := SysGetEnv(key); got == nil || *got != "typed" {
		t.Fatalf("SysGetEnv() = %v, want typed", got)
	}
	found := false
	for _, entry := range SysEnvironmentEntries() {
		if entry != nil && entry.Key != nil && *entry.Key == *key {
			found = entry.Value != nil && *entry.Value == "typed"
		}
	}
	if !found {
		t.Fatal("SysEnvironmentEntries() omitted the staged environment value")
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
	SysChangeCwd(StringFromLiteral(temporary))
	if cwd := SysGetCwd(); cwd == nil || *cwd == "" {
		t.Fatalf("SysGetCwd() = %v after SysChangeCwd", cwd)
	}

	programPath := SysCurrentProgramPath()
	if programPath == nil || *programPath == "" || !filepath.IsAbs(*programPath) {
		t.Fatalf("SysCurrentProgramPath() = %v, want an absolute path", programPath)
	}

	shellCommand := "exit 7"
	if runtime.GOOS == "windows" {
		shellCommand = "exit /B 7"
	}
	if code := SysCommand(StringFromLiteral(shellCommand), nil); code != 7 {
		t.Fatalf("SysCommand(%q, nil) = %d, want 7", shellCommand, code)
	}
}

func TestPortableSysClockCwdAndProgramPath(t *testing.T) {
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
		second := SysReadCharValue()
		if second != 'B' {
			t.Fatalf("SysReadCharValue() = %d, want 'B'", second)
		}
		if err := SysStdout().WriteByte(second); err != nil {
			t.Fatalf("echo redirected character: %v", err)
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
		if value := SysReadCharValue(); value != -1 {
			t.Fatalf("SysReadCharValue() = %d, want -1", value)
		}
	})
}

func TestPortableSysReadCharValuePreservesRedirectedInput(t *testing.T) {
	withStandardFiles(t, "Q", func(_ string, _ string) {
		if value := SysReadCharValue(); value != 'Q' {
			t.Fatalf("SysReadCharValue() = %d, want %d", value, 'Q')
		}
		if value := SysReadCharValue(); value != -1 {
			t.Fatalf("SysReadCharValue() at EOF = %d, want -1", value)
		}
	})
}
