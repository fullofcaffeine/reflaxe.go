package hxrt

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// SysEnvironmentEntry is one typed environment key/value pair for staged Sys.
//
// What: Carries process environment data across the typed Haxe/runtime boundary.
// Why: A native Go map is not the public Haxe Map representation, and constructing
// generated StringMap internals in hxrt would couple the runtime to compiler output.
// How: Export pointer-backed strings; staged Sys builds the public StringMap.
type SysEnvironmentEntry struct {
	Key   *string
	Value *string
}

func SysGetCwd() *string {
	cwd, err := os.Getwd()
	if err != nil {
		return StringFromLiteral("")
	}
	return StringFromLiteral(cwd)
}

func SysArgs() []*string {
	args := os.Args
	if len(args) <= 1 {
		return []*string{}
	}
	out := make([]*string, 0, len(args)-1)
	for _, arg := range args[1:] {
		out = append(out, StringFromLiteral(arg))
	}
	return out
}

func SysGetEnv(key *string) *string {
	if key == nil {
		return nil
	}
	value, ok := os.LookupEnv(*key)
	if !ok {
		return nil
	}
	return StringFromLiteral(value)
}

// SysPutEnv changes one environment entry and reports the native OS failure.
//
// Portable Sys.putEnv deliberately ignores this error because the upstream
// Haxe 4.3.7 eval contract exposes a Void operation that does not throw for an
// invalid key. Keeping the error in this helper lets native facades preserve Go
// os.Setenv/os.Unsetenv behavior instead of inheriting that portable contract.
func SysPutEnv(key *string, value *string) error {
	if key == nil {
		return nil
	}
	if value == nil {
		return os.Unsetenv(*key)
	}
	return os.Setenv(*key, *value)
}

// SysSetEnvironment applies the non-throwing Haxe 4.3.7 putEnv contract.
//
// What: Sets or removes one environment variable for staged Sys.putEnv.
// Why: The eval reference surface exposes Void and ignores malformed-key errors,
// while SysPutEnv intentionally retains the native error for Go-native callers.
// How: Delegate to SysPutEnv and deliberately discard its result at this boundary.
func SysSetEnvironment(key *string, value *string) {
	_ = SysPutEnv(key, value)
}

func SysCommand(command *string, args []*string) int {
	if command == nil {
		return -1
	}
	var cmd *exec.Cmd
	if args == nil {
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/C", *StdString(command))
		} else {
			cmd = exec.Command("/bin/sh", "-c", *StdString(command))
		}
	} else {
		cmd = exec.Command(*StdString(command), StringSlice(args)...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		return -1
	}
	return 0
}

func SysExit(code int) {
	os.Exit(code)
}

// SysEnvironmentEntries snapshots the process environment as typed pairs.
func SysEnvironmentEntries() []*SysEnvironmentEntry {
	out := make([]*SysEnvironmentEntry, 0)
	for _, entry := range os.Environ() {
		key, value, ok := strings.Cut(entry, "=")
		if !ok {
			out = append(out, &SysEnvironmentEntry{
				Key:   StringFromLiteral(entry),
				Value: StringFromLiteral(""),
			})
			continue
		}
		out = append(out, &SysEnvironmentEntry{
			Key:   StringFromLiteral(key),
			Value: StringFromLiteral(value),
		})
	}
	return out
}

func SysSystemName() *string {
	switch runtime.GOOS {
	case "darwin":
		return StringFromLiteral("Mac")
	case "linux":
		return StringFromLiteral("Linux")
	case "windows":
		return StringFromLiteral("Windows")
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		return StringFromLiteral("BSD")
	default:
		return StringFromLiteral(runtime.GOOS)
	}
}

// SysSleep suspends the current goroutine for a Haxe seconds-based duration.
//
// What: Implements the portable Sys.sleep(Float) blocking contract.
// Why: The mainstream Haxe declaration lowers to a target-owned Sys_sleep
// symbol, while Go's time.Sleep accepts a nanosecond Duration instead of
// seconds. Keeping conversion here avoids library behavior in compiler GoRaw.
// How: Clamp non-positive and NaN inputs to an immediate return, convert
// positive seconds through time.Second, and delegate scheduling to time.Sleep.
func SysSleep(seconds float64) {
	if !(seconds > 0) {
		return
	}
	time.Sleep(time.Duration(seconds * float64(time.Second)))
}

// SysSetCwd changes the process working directory and preserves native errors.
//
// What: Implements Sys.setCwd through os.Chdir.
// Why: Missing and inaccessible directories are recoverable Haxe failures, not
// successful no-ops or Go panics.
// How: Return the typed Go error so the generated adapter can throw it through
// the Haxe exception boundary.
func SysSetCwd(path *string) error {
	return os.Chdir(*StdString(path))
}

// SysChangeCwd is the Haxe-shaped working-directory capability.
func SysChangeCwd(path *string) {
	if err := SysSetCwd(path); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

// SysTime returns wall-clock epoch time in fractional seconds.
//
// What: Implements the precise timestamp contract of Sys.time.
// Why: The thread runtime clock is process-relative and therefore cannot stand
// in for the system timestamp required by the root Sys API.
// How: Convert Unix nanoseconds to seconds at the typed runtime boundary.
func SysTime() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Second)
}

// SysProgramPath returns the path of the executable running this program.
//
// What: Implements Sys.programPath for compiled Go programs.
// Why: The upstream method describes the current program artifact; os.Args[0]
// may be relative, PATH-resolved, or unrelated after argument mutation.
// How: Delegate to os.Executable and retain its error for Haxe translation.
func SysProgramPath() (*string, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	return StringFromLiteral(path), nil
}

// SysCurrentProgramPath is the Haxe-shaped executable-path capability.
func SysCurrentProgramPath() *string {
	path, err := SysProgramPath()
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return StringFromLiteral("")
	}
	return path
}
