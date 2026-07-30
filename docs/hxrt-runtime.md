# `hxrt` Runtime: Why It Exists, How It Works, and What It Does

## Terms

- `hxrt`: the runtime helper package copied into each generated Go module (`<go_module>/hxrt`).
- `compiler shim`: helper declarations emitted by the compiler at build time (not copied from `runtime/hxrt` files).
- `semantic contract`: the source behavior rules we keep stable, especially
  portable Haxe semantics and explicit typed native API contracts.

## What `hxrt` is

`hxrt` is the Go runtime support package emitted with every `reflaxe.go` build output.

- Source of truth: `runtime/hxrt/*.go`
- Generated location: `<go_output>/hxrt/*.go`
- Import path in generated code: `<go_module>/hxrt`

`hxrt` is not a separate dependency you install from the network. It is copied into the generated module so output stays self-contained and reproducible.

Selective runtime strategy and profile interaction are documented in `docs/hxrt-selective-runtime.md`.

## Why `hxrt` is needed

`hxrt` exists to bridge semantic and representation gaps between Haxe and Go in a deterministic, reusable way.

This is not a Go-only idea; many compiler targets use a runtime package. In
`reflaxe.go`, the runtime is paired with staged stdlib source. Remaining
compiler shims are measured compatibility debt during that migration, except
for exact registered metadata or representation primitives.

1. Haxe semantics do not map 1:1 to native Go primitives.
   - String behavior and nullability need helper semantics (`Std.string` shape, null-safe concat/equality, rune-aware length/indexing).
   - Haxe exception flow (`throw`/`try`/`catch`) needs a controlled panic/recover boundary.
   - Some stdlib/runtime contracts require target-specific behavior (`Sys`, file/process wrappers, atomic cell behavior).

2. Centralizing helpers avoids re-emitting large behavior blocks into every generated file.
   - Compiler lowers calls to stable helpers instead of duplicating logic.
   - Generated output remains smaller and easier to inspect.

3. It supports strict boundary policy.
   - Project code does not need raw `__go__` for basic runtime glue.
   - Compiler/runtime contract is explicit and testable.

For broader shim ownership tradeoffs, see `docs/stdlib-shim-rationale.md`.

## How it works

Compilation wiring:

1. `go_module` is resolved (default `snapshot`).
2. Runtime import path is computed as `<module>/hxrt` in `src/reflaxe/go/CompilationContext.hx`.
3. Compiler emits generated Go that imports and calls `hxrt` helpers.
4. On output, backend writes:
   - `go.mod`
   - generated `.go` files
   - copied runtime directory `hxrt/` from `runtime/hxrt`
   Every destination is validated by the
   [generated-output confinement boundary](generated-output-confinement.md)
   immediately before the managed write; an existing or broken output symlink
   is rejected rather than followed.
5. Backend runs `go build` by default and fails compilation if it cannot
   launch or exits nonzero. `-D go_no_build` and `-D go_codegen_only` are
   explicit opt-outs for callers that own a separate Go build/test gate.

Key implementation points:

- Runtime copy/write: `src/reflaxe/go/GoReflaxeCompiler.hx`
- Runtime copy helper for iterator flows: `src/reflaxe/go/GoOutputIterator.hx`
- Runtime source: `runtime/hxrt/*.go`

## What `hxrt` currently does

`hxrt` currently owns helper functions in these groups:

- String/runtime conversion helpers:
  - `StringFromLiteral`, `StdString`, `StringSlice`
  - `StringConcatAny`, `StringEqualAny`
  - `StringConcatStringPtr`, `StringEqualStringPtr`
  - `StringLength`, `StringCharAt`, `StringCharCodeAt`, `StringSubstring`
- Numeric helpers:
  - `FloatMod`, `Int32Wrap`
- Portable Array identity (`runtime/hxrt/array.go`):
  - one shared mutable carrier for root Haxe `Array<T>` values
  - null-safe indexed reads, sparse growth, length changes, and alias-preserving mutation
  - localized `[]any` storage; typed values are recovered by compiler-generated operations
- Atomic runtime cells:
  - `AtomicInt*` helpers
  - `AtomicObject*` helpers
- Exception bridging:
  - `Throw`, `TryCatch`, `UnwrapException`, `ReportUncaughtException`
  - `ExceptionCaught`, `ExceptionThrown`, `ExceptionMessage`
- Portable thread lifecycle:
  - `ThreadSpawn`, `ThreadSpawnWithEventLoop`, `ThreadWaitForAll`
  - `ThreadLocalNew`, `ThreadLocalGet`, `ThreadLocalSet`
  - `ThreadSpawnDetached` for lifecycle-only scoping of compiler-owned
    `go.Go.spawn` callbacks when portable thread identity is reachable
  - logical thread identity, thread-local values, message queues,
    synchronization, and event-loop state
- JSON wrappers:
  - `JsonParse`, `JsonStringify`
- System wrappers (`runtime/hxrt/sys.go`):
  - `SysGetCwd`, `SysChangeCwd`, `SysArgs`, `SysGetEnv`, `SysPutEnv`,
    `SysSetEnvironment`, typed environment entries, `SysSleep`, `SysTime`,
    `SysCurrentProgramPath`, `SysCommand`, and `SysExit`
- File capabilities (`runtime/hxrt/file.go`):
  - typed `FileReadContent`, `FileWriteContent`, `FileReadByteValues`, `FileWriteByteValues`, and `FileCopyContents`
  - opaque `FileInput` / `FileOutput` handles plus `FileOpen*`, read/write, seek/tell/eof, flush, and close operations
  - non-owning `SysStdin`, `SysStdout`, and `SysStderr` handles used by the staged file-stream classes
- Process wrappers (`runtime/hxrt/process.go`):
  - native `NewProcess` handles plus the typed `ProcessCreate`, pipe, byte-transfer, PID, status, kill, and close capabilities consumed by `std/hxrt/process`
- HTTP transport (`runtime/hxrt/http.go`):
  - opaque request/response handles consumed only through typed
    `std/hxrt/http` bindings;
  - URL/query/form construction, proxy configuration, bounded synchronous
    `net/http` execution, deterministic indexed response headers, body closure,
    idle-transport cleanup, and optional typed `SocketHandle` consumption;
  - multipart files are pulled from a typed chunk callback into Go's request
    writer, so partial source reads are preserved without buffering the whole
    declared upload; premature EOF or source failure aborts the exchange;
  - request selection, data-URL behavior, multipart policy, public maps,
    callbacks, and status/error classification remain in staged `sys.Http`.
- Network capabilities (`runtime/hxrt/socket.go` plus build-tagged
  `runtime/hxrt/socket_broadcast_*.go` and
  `runtime/hxrt/socket_listener_*.go` /
  `runtime/hxrt/socket_readiness_*.go` adapters):
  - one opaque, synchronized `SocketHandle` shared by TCP and UDP;
  - typed eager DNS/IPv4, connect/bind/listen/accept, byte transfer, deadline,
    blocking-policy, readiness, shutdown, address, broadcast, and datagram operations;
  - `new Host(name)` and `Host.reverse()` complete before any socket timeout
    can participate; the exact release exclusion is documented in the
    [socket DNS and timeout boundary](socket-dns-boundary.md);
  - one snapshotted dial policy applies the staged timeout to TCP connection
    establishment and, through TLS composition, to the TLS handshake;
  - connection installation is transactional: if a reset or another native
    error prevents deadline or fast-send policy from being applied, the new
    connection is detached and closed before the public operation fails;
  - typed pre-listen state preserves `bind` before `listen`; the later call
    passes its nonnegative backlog to the OS and then converts the descriptor
    into Go's pollable listener. The exact lifecycle and release boundary are
    documented in [socket server lifecycle and backlog](socket-server-lifecycle.md);
  - typed native-readiness snapshots preserve Haxe-owned buffered bytes and
    map POSIX `select` results back to caller indexes. Each raw descriptor is
    duplicated close-on-exec only inside Go's valid `Control` callback and
    closed after its bounded poll slice, so concurrent close cannot leave a
    reused number in the readiness set or leak the temporary capability to a
    child process. Linux and Darwin use explicit descriptor adapters;
    Windows and other unreviewed platforms fail explicitly instead of
    inventing readiness. The exact member and release boundary is documented in
    [socket readiness and nonblocking controls](socket-readiness-nonblocking.md);
  - protocol-aware shutdown sends TLS `close_notify` for write-only shutdown,
    rejects TLS read-only shutdown explicitly, and preserves plain TCP
    half-close. Fast-send follows only typed `NetConn` links to apply
    `TCP_NODELAY` to a TLS-wrapped TCP transport. The exact behavior and
    release boundary is documented in
    [socket shutdown and fast-send controls](socket-tls-controls.md);
  - a repeated native lifecycle matrix covers TCP success, timeout, reset,
    stalled TLS handshake, UDP, listener/readiness, concurrent close, and TLS
    close-notify. It requires active connections to reach zero and goroutines
    plus Linux file descriptors to return to bounded post-warm-up levels; see
    [socket and TLS resource convergence](socket-resource-convergence.md);
  - POSIX and Windows keep their native descriptor types behind separate
    build-tagged broadcast and listener helpers, with explicit
    unsupported-platform errors;
  - concrete result carriers keep byte progress, EOF/blocked state, peer addresses,
    accepted handles, and readiness indexes explicit instead of using `Dynamic`.
- TLS socket composition (`runtime/hxrt/socket_ssl.go`):
  - typed client setup, deferred server-listener wrapping, handshake,
    peer-certificate access, and SNI certificate selection over the shared
    `SocketHandle`;
  - client dialing uses the socket handle's timeout-aware `net.Dialer`, so a
    peer that accepts TCP but stalls the handshake cannot ignore `setTimeout`;
  - typed certificate/key primitives remain in `runtime/hxrt/ssl.go`, so SSL
    digest/certificate users do not select network transport automatically.
- Regex execution (`runtime/hxrt/regex.go`):
  - typed compiled RE2 handles, match snapshots, and quoting;
  - UTF-8 byte indexes are converted to the code-point offsets expected by the Haxe string contract before staged `EReg` sees them;
  - match state, capture validation, split/map traversal, Haxe replacement-template expansion, and global policy remain in `std/go/_std/EReg.hx`.
- Serialization representation (`runtime/hxrt/serialization.go`):
  - bounded host float parsing is the only serialization-specific native capability;
  - staged `haxe.Serializer` / `haxe.Unserializer` own token streams, reference caches, recursion, resolver policy, field traversal, assignment, and custom-hook sequencing through ordinary `Reflect` calls;
  - existing typed same-package Reflect metadata discovers and accesses generated private and inherited fields and methods without `unsafe`; erased calls still use the shared safe `Reflect.callMethod` helper, while Type metadata constructs empty instances with initialized superclass carriers and valid virtual-dispatch self bindings.

These helpers preserve native failures at the runtime boundary. Canonical staged file and Process wrappers translate bounds, EOF, nullable exit availability, and public lifecycle policy in Haxe source; process startup and non-EOF read failures remain distinct from normal EOF and child exit codes. Portable `Sys.putEnv` is the intentional exception: staged `Sys.hx` calls the non-throwing `SysSetEnvironment` capability to match the upstream Haxe 4.3.7 eval contract, while `SysPutEnv` retains the native error for typed Go-native bindings.

- Byte representation capabilities (`runtime/hxrt/bytes.go`):
  - an opaque immutable `ByteView` crosses typed `std/hxrt/io` bindings without
    exposing generated `haxe.io.Bytes` layout;
  - allocation, `[]int`/`[]byte` view conversion and validation, UTF-8/UTF-16LE
    conversion, overlap-safe copy, cloning, and growable-slice append primitives;
  - scalar IEEE-754 bit reinterpretation used by staged `haxe.io.FPHelper`.

Public bounds, hex algorithms, encoding selection, stream behavior, alias
observation, and cache invalidation remain in canonical staged Haxe. Base64,
digest, and HTTP body consumers reuse the same opaque view, so they do not copy
through `go.NativeSlice<Int>` or depend on generated `haxe.io.Bytes` fields.

## Exception and concurrency boundaries

Portable Haxe exceptions and native Go panics are deliberately different
failure domains:

1. `throw value` creates an `hxrt.HaxeException` carrier.
   The carrier and `Throw` live in the `core` runtime feature so even a
   core-only selective runtime can represent portable validation failures;
   catch/message helpers remain in the additive `exception` feature.
2. Generated Haxe `try`/`catch` unwraps only that carrier. A panic originating
   in a typed Go extern, the Go runtime, or malformed generated code continues
   unwinding as a native panic and cannot be mistaken for a Haxe value.
   Portable runtime validation failures, such as a failed dynamic-to-`Int`
   conversion, must call `Throw` and therefore remain Haxe-catchable; they must
   never use a raw Go panic as a shortcut.
3. An uncaught Haxe throw in `sys.thread.Thread.create` or
   `createWithEventLoop` ends that worker, writes
   `Uncaught exception <message>` to stderr, and does not crash the process.
4. Portable Haxe workers are foreground threads: a foreground thread keeps the
   generated program alive until it finishes. When the inferred runtime plan
   includes `thread`, generated `main` defers `ThreadWaitForAll`, which also
   drains portable workers created by other portable workers.
5. `go.Go.spawn` remains an explicit Go-native boundary. It is not included in
   the portable foreground count and preserves Go's normal fatal-panic and
   process-shutdown behavior. If `sys.thread` is reachable, the compiler uses
   `ThreadSpawnDetached` to initialize the runtime on the caller before launch
   and release only callback-owned logical identity and TLS state on return or
   panic unwind; it neither joins nor recovers the goroutine, and a nil callback
   retains native panic behavior.
6. Arbitrary goroutines created outside compiler-owned `Thread` and
   `go.Go.spawn` boundaries have no observable exit hook. Calling
   `Thread.current()` or `Tls` from such a foreign goroutine is therefore
   outside the automatic lifecycle-reclamation contract.

Synchronization, pool linearization, event-loop cancellation, condition
generations, channel close behavior, and the bounded `runtime.Stack` identity
fallback are specified in [`docs/concurrency-contract.md`](concurrency-contract.md).

The split is covered by `stdlib/sys_thread_uncaught_exception`,
`go_native/native_panic_not_haxe_catch`,
`go_native/goroutine_native_panic`, and
`go_native/goroutine_native_shutdown`, plus direct race-capable runtime tests
in `runtime/hxrt/exception_test.go` and `runtime/hxrt/thread_test.go`.

## What `hxrt` does not own

`hxrt` is not the whole Haxe stdlib implementation. Public library semantics
belong in upstream or canonical staged Haxe source. Exact admitted compiler
intrinsics are listed in `docs/compiler-stdlib-intrinsics.json`. No
`migration_required` compiler-stdlib group remains; adding one fails the
closeout gate instead of creating a new compatibility exception.

`sys.FileSystem` follows the preferred split: `std/go/_std/sys/FileSystem.hx`
owns its public API and `FileStat` construction, while typed `std/hxrt/fs`
bindings reach the selectively copied native capabilities in
`runtime/hxrt/filesystem.go`.

Root `Sys` follows the same rule: `std/go/_std/Sys.hx` owns the public Haxe
contract, typed `std/hxrt/sys` and `std/hxrt/fs` externs expose narrow native
capabilities, and `runtime/hxrt/sys.go` / `file.go` own only OS state and
handles. The typed `std/hxrt/sys/NativeTerminal.hx` binding selects build-tagged
`runtime/hxrt/terminal*.go` for terminal state and one-byte input. `hxrt` does
not construct the public environment map, aliases, fallbacks, Haxe stream
wrappers, `haxe.io.Eof`, or requested character echo.

Networking follows it too. `std/go/_std/sys/net/{Host,Socket,UdpSocket}.hx`,
`std/sys/net/_SocketIO.hx`, and `std/go/_std/sys/ssl/Socket.hx` own the public
objects, byte bounds/copies, Haxe EOF and blocked translation, address/result
construction, select identity, TLS configuration, and accepted SSL object
identity. Typed bindings under `std/hxrt/net` and `std/hxrt/ssl` expose only
native resources and concrete result carriers. `hxrt` never constructs a
generated `sys.net.Socket`, `Host`, `Address`, `Bytes`, or Haxe exception.

Regex and serialization use the same split. Typed bindings under
`std/hxrt/regex` and `std/hxrt/serialization` expose only native execution
capabilities. Serialization reuses the same generated Reflect field/method
metadata as ordinary `Reflect.field`, `Reflect.fields`, `Reflect.setField`, and
`Reflect.callMethod`; it has no private bridge, unsafe field lift, or duplicate
metadata registry. The staged calls intentionally select the shared safe
`runtime/hxrt/reflect.go` helper for dynamic objects and erased invocation.

Examples that are currently compiler-owned migration debt (not approved
`hxrt` ownership):

- `sys.Http`
- most base `haxe.io` shim declarations

## Change guidelines

Use `hxrt` when a helper is:

- target-runtime behavior (not just AST rewriting),
- reusable across many lowering sites,
- easier to verify once than duplicated per generated file.

Keep only an exact primitive in the compiler when correctness needs compile-time
metadata, policy, or a representation fact that staged source and typed `hxrt`
cannot express. A large generated type shape is a reason to design a typed
boundary, not by itself a reason to keep public behavior in a compiler shim.

When changing `hxrt`, update evidence:

- snapshots: `python3 test/run-snapshots.py`
- semantic diff: `python3 test/run-semantic-diff.py`
- full CI harness: `python3 test/run-ci.py`

And if ownership boundaries move, update:

- `docs/stdlib-shim-rationale.md`
- `docs/feature-support-matrix.md`

## Related docs

- `docs/start-here.md` - first-run commands and how generated output is produced.
- `docs/hxrt-selective-runtime.md` - runtime feature slicing policy and controls.
- `docs/stdlib-shim-rationale.md` - ownership split between runtime, compiler shims, and staged stdlib overrides.
- `docs/native-policy-presets.md` - policy presets and native-boundary behavior.
- `docs/concurrency-contract.md` - portable and Go-native concurrency invariants and evidence.
- `docs/portable-canonical-contract.md` - portable semantics baseline used by tests.
- `docs/glossary.md` - shared definitions used across docs.
