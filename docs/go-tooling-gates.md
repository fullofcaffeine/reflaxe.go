# Go Tooling Release Gates

## What It Is

The Go tooling lane runs dynamic and static checks against the copied `hxrt`
runtime, two generated concurrency contract modules, and four committed
flagship application modules:

- `runtime/hxrt`, with a native concurrency and reflection fixture;
- generated `sys.thread` output, with fixed/elastic pool admission and shutdown stress;
- generated `go.Chan<T>` output, with close/nil/native-panic contracts;
- PulseForge portable and metal generated output; and
- FluxProxy portable and metal generated output.

Each flagship fixture invokes its real `main()` function with `--scripted`.
That makes `go test -race` execute application behavior instead of merely
compiling a package with no tests.

Run the same gate locally from the repository root:

```bash
npm run security:go-tooling
```

Reports are written to `.cache/security/go-tooling/`. `manifest.json` records
the resolved Go and analyzer versions, every command result, duration, and
report path. `summary.md` is the short human-readable index. CI uploads that
directory even when a command fails.

## Why It Exists

Generated applications and `hxrt` use goroutines, synchronization, atomics,
reflection, and target-runtime adapters that ordinary snapshot comparisons do
not validate. A successful compile is not evidence of race freedom, safe
pointer behavior, or clean analyzer results. These gates turn those properties
into release-blocking evidence on every supported Go line.

## How It Works

For each declared scope, the runner creates an isolated temporary copy and
runs exactly these commands:

1. `go test -race -count=1 -timeout=5m ./...`;
2. `go test -gcflags=all=-d=checkptr=2 -count=1 -timeout=5m ./...`;
3. `go vet -stdmethods=false ./...`; and
4. `staticcheck -checks=SA* ./...`.

Staticcheck 2026.1 (`v0.7.0`) is installed from an exact module version. The
runner rejects a binary that reports another version. `SA*` is the complete
Staticcheck correctness family. Style and simplification checks are outside
this lane because generated identifiers and mechanical code shape are owned by
the compiler and would create churn without adding a correctness signal.

The one disabled vet analyzer is `stdmethods`. Haxe's public `ReadByte` and
`WriteByte` methods intentionally have Haxe-compatible signatures rather than
the signatures of Go's similarly named `io` interfaces. All other default vet
analyzers remain enabled. This is a named scope decision, not suppression of
individual findings.

The `hxrt` fixture concurrently exercises logical thread creation and message
delivery, integer and reference atomics, and the reflection branches for maps,
slices, functions, channels, pointers, and `unsafe.Pointer`. The checkptr run
executes those same paths with strict pointer checking enabled.

The one production `unsafe.Pointer` is the POSIX terminal ioctl bridge. Its
behavioral evidence is intentionally separate from the headless Go-tooling
fixture: `test/test_sys_get_char_terminal.py` builds the generated program with
`-gcflags=all=-d=checkptr=2`, then drives the actual terminal-mode transition,
one-byte read, echo policy, and restoration through a real pseudo-terminal.

The generated thread-pool fixture races 10,000 submissions against repeated
shutdown at `GOMAXPROCS=1,2,8` and requires each accepted task to execute
exactly once. It also requires fixed and elastic pools to replace a worker after
a task's Haxe throw so later accepted work still drains. The generated channel
fixture covers both reflective and typed helpers, including buffered drain,
closed versus empty receive, nil channels, send-after-close, and double-close.
Neither contract uses retries.

## Failure And Flake Policy

Any nonzero exit, analyzer finding, setup failure, or timeout fails the job and
therefore blocks semantic release. There are no test retries. Deterministic
race, checkptr, vet, or Staticcheck failures cannot be downgraded to warnings.
The matrix uses `fail-fast: false` only so evidence from both supported Go lines
is preserved; it does not change either lane's result.

The default per-command budget is six-minute timeout enforcement, while each
Go test also has its own five-minute timeout. Each supported-toolchain matrix
job has a 30-minute hard ceiling. A normal clean run should remain well below
that ceiling. Network or analyzer setup failures are not retried inside the
gate; rerunning a workflow is an explicit operator action and does not replace
the failed evidence.

## Updating The Lane

Add a failing contract before changing scopes, analyzer families, tool pins,
timeouts, or release dependencies. An analyzer upgrade is intentional policy
work: review its new checks, update the exact pin, run both supported Go lines,
and commit the contract and documentation change together.
