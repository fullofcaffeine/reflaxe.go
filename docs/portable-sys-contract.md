# Portable root `Sys` contract

## What this contract covers

Root `Sys` is Haxe's process-level API for arguments, environment, clocks,
working-directory state, program paths, commands, and standard streams. On
`haxe.go`, these methods have one profile-independent semantic contract:
`portable` and `metal` are policy presets, so neither preset changes what a
root `Sys` call means.

The implementation is deliberately split:

1. Haxe 4.3.7's mainstream `Sys` extern remains the source-facing API.
2. `lowerSysStdlibShimDecls` emits typed target symbols and translates Go
   errors/EOF into Haxe exceptions.
3. `runtime/hxrt/sys.go` owns root OS calls, clocks, paths, and blocking;
   `runtime/hxrt/file.go` owns standard-stream handles.

The root adapters construct canonical staged `sys.io.FileInput` / `FileOutput`
classes from typed opaque handles. File stream shape and behavior therefore stay
in `std/go/_std/sys/io`, rather than turning compiler-owned `GoRaw` into a
second standard library.

## Method matrix

| Root method | Portable status | Runtime or boundary rule |
| --- | --- | --- |
| `print`, `println` | Supported | `hxrt.Print` / `hxrt.Println` use Haxe display conversion; the unavoidable `Dynamic` value is confined to this public formatting boundary. |
| `args` | Supported | Returns arguments after the program name. |
| `getEnv`, `environment` | Supported | Missing `getEnv` keys remain `null`; the environment map is a call-time snapshot. |
| `putEnv` | Supported | The portable adapter intentionally discards native mutation errors for Haxe 4.3.7 eval parity; typed hxrt still retains the error for native APIs. |
| `sleep` | Supported | Converts Haxe seconds to Go `time.Duration`; non-positive and NaN inputs return immediately. |
| `setTimeLocale` | Explicitly unavailable | Returns `false`. Go has no process-global C time-locale switch, so success must not be reported. |
| `getCwd`, `setCwd` | Supported | `setCwd` preserves `os.Chdir` failure as a catchable Haxe value. |
| `systemName` | Supported | Maps admitted Go hosts to Haxe's `Windows`, `Linux`, `BSD`, or `Mac` names. |
| `command`, `exit` | Supported | Child streams are inherited; command exit status and process exit remain runtime-owned. |
| `time` | Supported | Returns wall-clock Unix epoch time in fractional seconds. It is not the process-relative thread clock. |
| `cpuTime` | Compile-time unsupported | Compilation fails with an actionable diagnostic. Wall-clock substitution would violate the process/thread CPU-time contract. |
| `programPath` | Supported | Uses `os.Executable` and converts lookup failures through the Haxe exception boundary. |
| `executablePath` | Supported deprecated alias | Calls the same `programPath` adapter, matching upstream Haxe's migration path. |
| `stdin`, `stdout`, `stderr` | Supported | Reuse the existing Haxe IO carriers over non-owning process streams. Closing a wrapper detaches that wrapper but does not close the process-wide descriptor; a later `Sys.stdout()`/`stderr()`/`stdin()` call returns a fresh usable wrapper. Stdout/stderr flush is a successful no-op because Go file writes are unbuffered and `Sync` is invalid for some pipes/terminals. |
| `getChar` | Supported byte-stream core | Reads one byte, reports EOF through `haxe.io.Eof`, and writes that byte once when `echo` is true. Raw/canonical terminal-mode control is the platform-specific follow-up `haxe_go-vfp.8.7.3`; redirected-stream behavior is the admitted deterministic contract. |

Platform-specific process CPU accounting or terminal control belongs behind an
explicit typed Go API or `@:goNative` module boundary. Selecting the `metal`
compatibility preset does not silently change these root `Sys` semantics.

## Why `cpuTime` fails at compile time

Haxe defines `Sys.cpuTime()` as time actually consumed on the CPU. Go's
standard library exposes wall-clock time but no portable process CPU clock.
Returning `Sys.time()` would make sleeping look like CPU work and falsely mark
the API as implemented. A runtime panic would delay the same fact until
production. The compiler therefore rejects any direct call or first-class
reference to `Sys.cpuTime` and points platform-specific users to an explicit
native boundary.

## Standard-stream lifetime and failure behavior

Ordinary `sys.io.File` handles own their `os.File`: `close()` releases the file
and detaches the generated carrier. The wrappers returned by `Sys.stdin()`,
`Sys.stdout()`, and `Sys.stderr()` are non-owning. Their `close()` detaches the
individual Haxe wrapper while leaving the process descriptor open, so code can
reacquire a fresh wrapper without allowing a library to close a descriptor
needed by the rest of the process.

Reads distinguish EOF from other OS errors. The staged `FileInput` translates
the runtime EOF sentinel to `haxe.io.Eof`; typed runtime capabilities translate
other failures through the Haxe exception boundary. Writes and cwd/path
failures likewise never become apparent success.

## Sibling-target comparison

The local sibling audit found one strong precedent and two non-precedents:

- `haxe.rust` owns the full root class in staged Haxe and delegates OS behavior
  to typed `hxrt.sys.NativeSys` helpers. Its standard streams are staged
  `haxe.io.Input` / `Output` subclasses, `setTimeLocale` returns `false`, and
  `cpuTime` is explicitly rejected instead of being replaced with wall time.
- `haxe.go` follows the same semantic ownership rule. `sys.io.File*` is now
  canonical staged source like the sibling target; only root `Sys` and Process
  retain thin generated adapters while their separate migrations remain open.
- The audited `haxe.elixir` tree has no production root `Sys` override to copy.
  The `haxe.ruby` root surface is placeholder/incomplete and uses no-op or zero
  results, so it is not acceptable parity evidence.

This is a bounded `thinking:high` ownership decision with a single honest
design after local tracing, so it does not require Oracle. By contrast, any
future deprecation of the global `metal` compatibility selector remains the
separate `thinking:xhigh` decision in `haxe_go-vfp.6.6` and still requires its
independent review, usage evidence, and SemVer migration plan.

## Evidence

- Snapshot and generated-symbol contract:
  `test/snapshot/sys/root_sys_portable`
- Compile-time unsupported contract:
  `test/snapshot/negative/sys_cpu_time_unsupported`
- Haxe 4.3.7 eval differential contract:
  `test/semantic_diff/root_sys_portable_contract`
- Direct runtime and standard-stream lifetime tests:
  `runtime/hxrt/sys_test.go`
- Existing adjacent contracts:
  `test/semantic_diff/sys_sleep_contract`,
  `test/semantic_diff/sys_command_contract`, and the file/process semantic
  contracts listed in the feature matrix
