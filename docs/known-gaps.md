# Known Gaps And Production Caveats

This page is the blunt status view for current limitations so teams can plan migrations with clear tradeoffs.

The current release claim is bounded by the generated
[compatibility and support matrix](compatibility-support-matrix.md) and its
[machine manifest](compatibility-support-manifest.json). Implementation or test
evidence in this page does not expand that scope. In particular, concurrency,
networking, Go-native APIs, non-canonical platforms, and untrusted-source
compilation are not admitted by the current pre-1.0 beta claim.

Generated files are now protected by the
[output-confinement contract](generated-output-confinement.md), including stale
metadata deletion and package-clean symlink cases. That defense does not admit
adversarial source: Haxe macros, dependencies, plugins, and explicitly selected
child tools execute with compiler-process authority. `reflaxe.go` is not a
sandbox for untrusted Haxe programs.

Before reading:

- `portable` and `metal` are compatible policy presets, not app variants or
  separate semantic products. See [native policy presets](native-policy-presets.md).
- `go.*` is the Go-native facade surface. It is intentionally outside the cross-target portable contract.
- `semantic-diff` is the runtime parity harness against Haxe `--interp`. See [semantic diff guide](semantic-diff-guide.md).

Current architecture status:

- `GoBuildContext` and `GoBuildContextResolver` are already in place for centralized contract/capability resolution.
- Deterministic contract/runtime/optimizer reports (`profile_contract`, `hxrt_plan`, `optimizer_plan`) are already emitted when enabled.
- The preset model is now decomposed into typed authority, specialization,
  fallback, strictness, planner, optimizer, and runtime axes. Whether the global
  `metal` selector should ever be deprecated remains a separately reviewed
  compatibility decision.

## Production hardening scoreboard

This table separates "safe to ship with a known caveat" from "needs new engineering before we should promise more."

Terms:

- **Owner** means the layer responsible for the next decision or implementation.
- **Current decision** says whether the caveat is acceptable for a production release, needs watch-only tracking, or should be reopened.
- **Evidence** is the command, doc, or tracker that proves the decision is not just an opinion.
- **Reopen trigger** is the concrete signal that should create or unblock a Bead.

| Caveat class | Owner | Current decision | Evidence | Reopen trigger |
| --- | --- | --- | --- | --- |
| Multi-package output | compiler/output | Ship as single Go package for GA; reopen only if production-scale projects hit concrete Go tooling limits. | [`docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers`](multi-package-output-evaluation.md#measurable-production-reopen-triggers), `python3 test/run-ci.py`, generated-output telemetry artifacts. | Generated file size, compile time, package-private boundary needs, or Go tooling limits become measurable user blockers. If this happens, file a new scoped Bead with the trigger evidence. |
| Advanced Go extern interop | typed native facade | Ship `@:go.import` / `@:go.name` / `@:go.receiver`, single `(T, error)` support through `go.Result<T>`, and generated tuple carrier wrappers for supported multi-return APIs. Keep complex/unsupported result shapes behind typed facade wrappers. | [Go extern generator](goextern.md), `tools/goextern/main_test.go`, `test/run-goextern-fixtures.py`, `test/snapshot/go_native/extern_tuple_return`, `--dynamic-report` output. | Users hit common Go APIs whose callback/generic/unsafe or cross-package result shapes still fall back to `Dynamic`; file a new scoped Bead for a typed generator rule or documented facade pattern, not raw `__go__` in app code. |
| Target-sensitive parity surfaces | harness/runtime | Ship with snapshot/runtime evidence where deterministic interpreter-vs-Go comparison would be misleading. | `python3 test/run-portable-parity-closure.py --list-blockers`, `test/.test-cache/portable_parity_closure_summary.json`, [event-loop semantic-diff spike](spikes/event-loop-semantic-diff-spike.md), [SSL/UDP semantic-diff spike](spikes/ssl-udp-semantic-diff-spike.md). | A stable async/network/stack comparison harness becomes possible, or a snapshot-only surface starts hiding user-visible behavior drift. If this happens, file a new scoped Bead for that surface. |
| Performance budget drift | perf/CI | Ship with perf visibility gates and warning annotations; decide separately which warning-only drifts should become release-blocking. | [`docs/performance-budget-policy.md`](performance-budget-policy.md), `npm run test:perf:go`, `npm run test:perf:hxrt-selective`, `npm run test:perf:apps`, warning-history artifacts, delta hard-gate dry-run artifacts. | A warning repeats across stable CI runs, affects a flagship app, or hides a portable-vs-metal regression users would notice. If this happens, file a new scoped Bead with the run history. |
| Strict production boundary policy | profiles/governance | Ship with strict mode recommended for production; do not allow app-side raw `__go__` to become the default escape hatch. | [Profiles](profiles.md), [defines reference](defines-reference.md), `npm run test:release-contracts`. | Examples, docs, or generated reports make raw app-side injection look normal instead of exceptional. If this happens, file a new scoped Bead for the drift. |

## Target-sensitive parity policy

Some APIs depend on the operating system, wall-clock time, Go runtime stack
frames, socket scheduling, or TLS handshakes. In this document,
**target-sensitive means** the generated Go behavior is real and tested on Go,
but the Haxe interpreter is not a stable enough oracle for byte-for-byte
runtime comparison.

That is why these surfaces use **snapshot/runtime** evidence. A snapshot/runtime
test proves that the generated Go compiles, runs, and preserves a stable Go-side
contract. It does not claim that every detail can be compared to Haxe
`--interp` through `semantic-diff`.

Important rule: a target-sensitive surface **must not be promoted to `semantic-diff`**
unless the comparison normalizes away the target-specific details that would
otherwise make the test flaky or dishonest.

Current decisions:

| Surface | Current evidence | Decision | Why this is not semantic-diff yet | Reopen trigger |
| --- | --- | --- | --- | --- |
| `haxe.EntryPoint` | `stdlib/haxe_main_loop_runtime_direct` | Keep snapshot/runtime. | Main-thread scheduling depends on Go runtime wakeups and timer behavior. | A deterministic async harness can compare only logical callback order, not wall-clock timing. |
| `haxe.MainLoop` | `stdlib/haxe_main_loop_runtime_direct` | Keep snapshot/runtime. | Repeating events and queued callbacks are runtime-scheduled. | A stable event-loop comparison harness exists and proves logical behavior without depending on timing jitter. |
| `haxe.Timer` | `stdlib/haxe_main_loop_runtime_direct` | Keep snapshot/runtime. | Timer delivery depends on elapsed time and scheduler behavior. | A normalized harness can compare "callback eventually ran" semantics without comparing exact timing. |
| `haxe.CallStack` | `stdlib/haxe_stack_loop_target_sensitive`, `stdlib/haxe_native_stack_trace_opt_in` | Keep snapshot/runtime. | Default portable behavior is deterministic empty stacks; opt-in native stacks expose Go runtime frames. | A future normalized frame format can compare Haxe-level frames without raw Go paths/function names. |
| `haxe.NativeStackTrace` | `stdlib/haxe_stack_loop_target_sensitive`, `stdlib/haxe_native_stack_trace_opt_in` | Keep snapshot/runtime. | Native stack capture is a Go diagnostic feature, not portable semantic parity. | A stable normalized stack-carrier contract exists across targets. |
| `sys.net.UdpSocket` | `stdlib/sys_net_udp_socket_direct` | Keep snapshot/runtime. | UDP loopback is stable enough for Go runtime smoke, but packet delivery, broadcast behavior, and socket options are OS/network-policy sensitive. | A deterministic local-network harness exists that avoids LAN broadcast and OS-specific socket-option drift. |
| `sys.ssl.Certificate` | `stdlib/sys_ssl_leaf_direct` | Keep snapshot/runtime. | Certificate parsing is Go-backed and tied to Go's TLS/x509 behavior. | A semantic contract is narrowed to deterministic fields also guaranteed by the interpreter. |
| `sys.ssl.Digest` | `stdlib/sys_ssl_leaf_direct` | Keep snapshot/runtime. | Digest/sign/verify behavior touches Go crypto APIs and key parsing. | A deterministic cross-target crypto fixture proves the same supported algorithms and error behavior. |
| `sys.ssl.Key` | `stdlib/sys_ssl_leaf_direct` | Keep snapshot/runtime. | PEM parsing and private-key handling are Go crypto/runtime behavior. | A normalized key contract can compare only stable public behavior, not backend carrier details. |
| `sys.ssl.Socket` | `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct` | Keep snapshot/runtime. | TLS socket connect/accept/handshake behavior depends on Go networking and TLS state machines. | A deterministic loopback TLS harness can compare portable observable behavior without Go-specific handshake internals. |

The grouped `sys.ssl.Certificate` / `sys.ssl.Digest` / `sys.ssl.Key` leaf
surface is still useful and production-relevant. It simply remains
target-sensitive until the test oracle can compare only the portable part of the
behavior. A focused spike in `docs/spikes/ssl-udp-semantic-diff-spike.md` found
that `sys.ssl.Digest.make` is not available as an interpreter reference and that
`sys.net.UdpSocket` is not available on Haxe `--interp`, so the current honest
contract remains snapshot/runtime evidence for SSL leaf and UDP socket behavior.

The grouped `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` event-loop
surface follows the same rule: it is supported on Go, but runtime scheduling is
not the same thing as interpreter-vs-Go semantic parity. A focused spike in
`docs/spikes/event-loop-semantic-diff-spike.md` tried a narrow
`EntryPoint.runInMainThread` probe and found that Haxe `--interp` does not execute `EntryPoint.runInMainThread` callbacks in that direct comparison, so the
current honest contract remains snapshot/runtime evidence.

## Compiler/output caveats

- Output remains a single Go package (multi-file, single package); multi-package emission is not implemented yet.
- Multi-package output is currently deferred as non-blocking for production GA; explicit boundary conditions for re-opening are documented in [`docs/multi-package-output-evaluation.md`](multi-package-output-evaluation.md).
- These remaining lowering guards are invariant checks, not open supported-language gaps.
- No currently supported Haxe source construct is expected to hit them in normal typed lowering.
- Current invariant inventory is summarized in the [feature support matrix](feature-support-matrix.md):
  - `Unsupported assignment target` (`lowerLValue`)
  - `Unsupported postfix unary operator` (`lowerExpr` / `lowerExprWithPrefix`)
  - `Unsupported expression` (catch-all `lowerExpr` fallback)
  - `Std.isOfType` still has conservative fallback behavior for unresolved runtime-value abstract targets (documented as partial support, not a hard-fail)
- Invariant fixture strategy for the remaining lowering guards:
  - `Unsupported assignment target`: `test/snapshot/negative/non_lvalue_assignment_invariant`
  - `Unsupported postfix unary operator`: `test/snapshot/negative/postfix_non_inc_dec_invariant`
  - `Unsupported expression` catch-all: closure-by-node-family via `test/semantic_diff/type_expr_contract`, `test/semantic_diff/throw_expr_contract`, `test/snapshot/core/untyped_ident_nil`, `test/snapshot/core/const_kinds_contract`
  - `Std.isOfType` fallback behavior: `test/semantic_diff/std_is_of_type_contract`, `test/semantic_diff/std_is_of_type_runtime_core_abstract_contract`, and `test/snapshot/core/std_is_of_type_basic`, `test/snapshot/core/std_is_of_type_dynamic`, `test/snapshot/core/type_switch_no_binding_std_is_of_type`
- `go.*` APIs are target-specific. They compile to real Go behavior on this target, but they are not portability-safe across non-Go Haxe targets.
- Direct `haxe.ValueException` constructor/message/value parity is covered.
- Direct `haxe.exceptions` subclass construction for `PosException`, `ArgumentException`, and `NotImplementedException` is covered too.
- There is currently no compile-only portable stdlib debt in the generated inventory.
- The portable parity closure summary reports `0 actionable portable stdlib blockers`; remaining non-semantic-diff surfaces are policy-locked as target-sensitive snapshots or explicit target-conditional exclusions.
- Direct `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` usage now has snapshot/runtime smoke coverage through `stdlib/haxe_main_loop_runtime_direct`.
- The implementation is intentionally staged std over the runtime-backed `sys.thread.EventLoop` contract: public Haxe APIs live in `std/go/_std/haxe/*.hx`, while `runtime/hxrt/thread.go` owns scheduling, timers, and main-thread wakeups.
- This is not yet semantic-diff coverage because event-loop timing is target-sensitive; keep using runtime snapshots for this surface until the harness has a stable asynchronous comparison mode.
- `haxe.CallStack` and `haxe.NativeStackTrace` use deterministic empty-stack fallbacks by default. Native Go stack capture is available only with `-D reflaxe_go_native_stack_trace` as an explicit target-sensitive diagnostic capability, not as portable semantic-diff behavior. See [native stack capture contract](spikes/native-stack-capture-contract.md).

## Interop caveats

- Typed extern metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) is supported, but advanced Go signatures may still require façade wrappers.
- Single `(T, error)` extern surfaces can use `@:go.valueError` + `go.Result<T>`, and supported multi-return APIs can use generated tuple carriers. Complex callback/generic/unsafe shapes may still need façade wrappers.
- Keep strict policy enabled in production (`reflaxe_go_strict` / `reflaxe_go_strict_examples`) to avoid drifting into raw app-side `__go__`.

## Concurrency caveats

- `go.Go.spawn` and `go.Chan<T>` map to real goroutine/channel behavior on Go output.
- `go.Go.spawn` is intentionally not a portable foreground-thread API: generated
  `main` does not wait for those goroutines, and an uncaught panic is fatal.
- `sys.thread.Thread.create` and `createWithEventLoop` use the portable contract:
  generated `main` drains nested foreground workers, uncaught Haxe throws are
  reported and terminate only their worker, and foreign Go panics remain fatal.
- `go.Select` exposes typed deterministic helpers (`recv`, `recv2`, `send`, `send2`) built on non-blocking channel operations.
- Multi-branch helper priority is explicit and deterministic (`first` branch checked before `second`); it does not model Go runtime pseudo-random ready-case selection.
- `go.Chan<T>` now distinguishes `"empty"` from drained `"closed"` in
  `tryRecv`; send-after-close and invalid close retain native Go panics. Channel
  direction types, range syntax, and native multi-case select remain broader
  Go-native capability work. See [the concurrency contract](concurrency-contract.md).
- Re-entrant portable synchronization uses a bounded 64-byte `runtime.Stack`
  header parse because Go has no supported goroutine-local identity API. The
  parser is strict, zero-allocation in steady state, and lifecycle-tested; a Go
  runtime header-format change fails closed rather than corrupting ownership.
- `sys.thread.Tls` get/set/clear and worker isolation have semantic-diff
  coverage. Completed portable threads and compiler-owned `go.Go.spawn`
  callbacks release their identity and TLS state on return or panic unwind.
  Arbitrary goroutines created by handwritten Go, typed extern callbacks, or
  other foreign code still have no automatic detach signal; using
  `Thread.current()` or `Tls` there remains outside the admitted boundary. See
  the [compatibility matrix](compatibility-support-matrix.md#portable-concurrency).

## Native specialization caveats

- Current high-value typed native surfaces:
  - `go.Chan<T>`
  - `go.Slice<T>`
  - `go.Map<K,V>`
  - `go.Result<T>`
- Use portable Haxe as the semantic baseline. Use typed Go APIs or
  `@:goNative` when source is intentionally Go-specific. Select eager
  specialization only when its output/performance evidence justifies it; do not
  use a preset name as the evidence.

## Performance caveats

- Generated code aims for predictable shape first, then optimization under harness gates.
- Shim-heavy paths can still carry conversion overhead versus direct handwritten Go.
- Current warning-vs-hard-gate decisions are documented in [`docs/performance-budget-policy.md`](performance-budget-policy.md).
- Track real costs with:
  - `npm run test:perf:go`
  - `npm run test:perf:stdlib-shims`

## Source of truth links

- Release admission authority: [compatibility support manifest](compatibility-support-manifest.json)
- Human release scope: [compatibility and support matrix](compatibility-support-matrix.md)
- Generated release wording: [compatibility release status](compatibility-release-status.md)
- Implementation evidence inventory: [feature support matrix](feature-support-matrix.md)
- Ownership decision rule: [ownership rubric](ownership-rubric.md)
- Shim ownership decisions: [stdlib shim rationale](stdlib-shim-rationale.md)
- Phase roadmap: [phase 2 roadmap](phase2-roadmap.md)
