# Known Gaps And Production Caveats

This page is the blunt status view for current limitations so teams can plan migrations with clear tradeoffs.

Before reading:

- `portable` and `metal` are compiler profiles (contracts), not app variants. See `docs/profiles.md`.
- `go.*` is the Go-native facade surface. It is intentionally outside the cross-target portable contract.
- `semantic-diff` is the runtime parity harness against Haxe `--interp`. See `docs/semantic-diff-guide.md`.

Current architecture status:

- `GoBuildContext` and `GoBuildContextResolver` are already in place for centralized contract/capability resolution.
- Deterministic contract/runtime/optimizer reports (`profile_contract`, `hxrt_plan`, `optimizer_plan`) are already emitted when enabled.
- Remaining work is primarily production hardening and target-sensitive parity promotion, not profile-model replacement.

## Production hardening scoreboard

This table separates "safe to ship with a known caveat" from "needs new engineering before we should promise more."

Terms:

- **Owner** means the layer responsible for the next decision or implementation.
- **Current decision** says whether the caveat is acceptable for a production release, needs watch-only tracking, or should be reopened.
- **Evidence** is the command, doc, or tracker that proves the decision is not just an opinion.
- **Reopen trigger** is the concrete signal that should create or unblock a Bead.

| Caveat class | Owner | Current decision | Evidence | Reopen trigger |
| --- | --- | --- | --- | --- |
| Multi-package output | compiler/output | Ship as single Go package for GA; reopen only if production-scale projects hit concrete Go tooling limits. | `docs/multi-package-output-evaluation.md`, `python3 test/run-ci.py` | Generated file size, compile time, package-private boundary needs, or Go tooling limits become measurable user blockers. Tracked by `haxe.go-hm3p.5`. |
| Advanced Go extern interop | typed native facade | Ship current `@:go.import` / `@:go.name` / `@:go.receiver` plus single `(T, error)` support; keep other multi-return-heavy APIs behind typed facade wrappers until generated wrappers exist. | `docs/goextern.md`, `tools/goextern/main_test.go`, `test/run-goextern-fixtures.py` | Users need direct typed bindings for multi-return-heavy Go APIs without writing facade wrappers. `haxe.go-hm3p.2` records the boundary policy; follow-up implementation should add generated wrappers, not raw `__go__` in app code. |
| Target-sensitive parity surfaces | harness/runtime | Ship with snapshot/runtime evidence where deterministic interpreter-vs-Go comparison would be misleading. | `python3 test/run-portable-parity-closure.py --list-blockers`, `test/.test-cache/portable_parity_closure_summary.json` | A stable async/network/stack comparison harness becomes possible, or a snapshot-only surface starts hiding user-visible behavior drift. Tracked by `haxe.go-hm3p.3`. |
| Performance budget drift | perf/CI | Ship with perf visibility gates and warning annotations; decide separately which warning-only drifts should become release-blocking. | `npm run test:perf:go`, `npm run test:perf:hxrt-selective`, `npm run test:perf:apps` | A warning repeats across stable CI runs, affects a flagship app, or hides a portable-vs-metal regression users would notice. Tracked by `haxe.go-hm3p.4`. |
| Strict production boundary policy | profiles/governance | Ship with strict mode recommended for production; do not allow app-side raw `__go__` to become the default escape hatch. | `docs/profiles.md`, `docs/defines-reference.md`, `npm run test:release-contracts` | Examples, docs, or generated reports make raw app-side injection look normal instead of exceptional. Tracked by `haxe.go-hm3p`. |

## Compiler/output caveats

- Output remains a single Go package (multi-file, single package); multi-package emission is not implemented yet.
- Multi-package output is currently deferred as non-blocking for production GA; explicit boundary conditions for re-opening are documented in `docs/multi-package-output-evaluation.md`.
- These remaining lowering guards are invariant checks, not open supported-language gaps.
- No currently supported Haxe source construct is expected to hit them in normal typed lowering.
- Current invariant inventory (`docs/feature-support-matrix.md`, owned by `haxe.go-14as.56`):
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
- The implementation is intentionally staged std over the runtime-backed `sys.thread.EventLoop` contract: public Haxe APIs live in `std/haxe/*.cross.hx`, while `runtime/hxrt/thread.go` owns scheduling, timers, and main-thread wakeups.
- This is not yet semantic-diff coverage because event-loop timing is target-sensitive; keep using runtime snapshots for this surface until the harness has a stable asynchronous comparison mode.
- `haxe.CallStack` and `haxe.NativeStackTrace` use deterministic empty-stack fallbacks by default. Native Go stack capture is available only with `-D reflaxe_go_native_stack_trace` as an explicit target-sensitive diagnostic capability, not as portable semantic-diff behavior. See `docs/spikes/native-stack-capture-contract.md`.

## Interop caveats

- Typed extern metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) is supported, but advanced Go signatures may still require façade wrappers.
- Single `(T, error)` extern surfaces are supported via `@:go.valueError` + `go.Result<T>`, but multi-return-heavy APIs beyond that pattern still need façade wrappers.
- Keep strict policy enabled in production (`reflaxe_go_strict` / `reflaxe_go_strict_examples`) to avoid drifting into raw app-side `__go__`.

## Concurrency caveats

- `go.Go.spawn` and `go.Chan<T>` map to real goroutine/channel behavior on Go output.
- `go.Select` exposes typed deterministic helpers (`recv`, `recv2`, `send`, `send2`) built on non-blocking channel operations.
- Multi-branch helper priority is explicit and deterministic (`first` branch checked before `second`); it does not model Go runtime pseudo-random ready-case selection.

## Metal profile caveats

- Current "metal-ready" high-value lane:
  - `go.Chan<T>`
  - `go.Slice<T>`
  - `go.Map<K,V>`
  - `go.Result<T>`
- Use `portable` as semantic baseline, then promote hot paths to `metal` only when benchmark evidence justifies it.

## Performance caveats

- Generated code aims for predictable shape first, then optimization under harness gates.
- Shim-heavy paths can still carry conversion overhead versus direct handwritten Go.
- Track real costs with:
  - `npm run test:perf:go`
  - `npm run test:perf:stdlib-shims`

## Source of truth links

- Feature/support inventory: `docs/feature-support-matrix.md`
- Ownership decision rule: `docs/ownership-rubric.md`
- Shim ownership decisions: `docs/stdlib-shim-rationale.md`
- Phase roadmap: `docs/phase2-roadmap.md`
