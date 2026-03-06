# Known Gaps And Production Caveats

This page is the blunt status view for current limitations so teams can plan migrations with clear tradeoffs.

Before reading:

- `portable` and `metal` are compiler profiles (contracts), not app variants. See `docs/profiles.md`.
- `go.*` is the Go-native facade surface. It is intentionally outside the cross-target portable contract.
- `semantic-diff` is the runtime parity harness against Haxe `--interp`. See `docs/semantic-diff-guide.md`.

Current architecture status:

- `GoBuildContext` and `GoBuildContextResolver` are already in place for centralized contract/capability resolution.
- Deterministic contract/runtime/optimizer reports (`profile_contract`, `hxrt_plan`, `optimizer_plan`) are already emitted when enabled.
- Remaining work is primarily language/stdlib parity closure, not profile-model replacement.

## Compiler/output caveats

- Output remains a single Go package (multi-file, single package); multi-package emission is not implemented yet.
- Multi-package output is currently deferred as non-blocking for production GA; explicit boundary conditions for re-opening are documented in `docs/multi-package-output-evaluation.md`.
- Some unsupported typed-expression guards still exist by design (`docs/feature-support-matrix.md` inventory section). Current compiler hard-fails:
  - `Unsupported assignment target` (`lowerLValue`)
  - `Unsupported postfix unary operator` (`lowerExpr` / `lowerExprWithPrefix`)
  - `Unsupported expression` (catch-all `lowerExpr` fallback)
  - `Std.isOfType` still has conservative fallback behavior for unresolved runtime-value abstract targets (documented as partial support, not a hard-fail)
- Invariant fixture strategy for surviving hard-fail paths (`haxe.go-14as.8`):
  - `Unsupported assignment target`: `test/snapshot/negative/non_lvalue_assignment_invariant`
  - `Unsupported postfix unary operator`: `test/snapshot/negative/postfix_non_inc_dec_invariant`
  - `Unsupported expression` catch-all: closure-by-node-family via `test/semantic_diff/type_expr_contract`, `test/semantic_diff/throw_expr_contract`, `test/snapshot/core/untyped_ident_nil`, `test/snapshot/core/const_kinds_contract`
  - `Std.isOfType` fallback behavior: `test/semantic_diff/std_is_of_type_contract`, `test/semantic_diff/std_is_of_type_runtime_core_abstract_contract`, and `test/snapshot/core/std_is_of_type_basic`, `test/snapshot/core/std_is_of_type_dynamic`, `test/snapshot/core/type_switch_no_binding_std_is_of_type`
- `go.*` APIs are target-specific. They compile to real Go behavior on this target, but they are not portability-safe across non-Go Haxe targets.
- Direct `haxe.Template` usage is intentionally blocked for now: the source-owned std inclusion path still needs module-local enum emission to support the upstream implementation cleanly. See `test/snapshot/negative/direct_haxe_template_unsupported` and `haxe.go-14as.38`.
- Direct `haxe.ValueException` usage is intentionally blocked for now: string-payload message parity still depends on unresolved `Any`/string boxing semantics. See `test/snapshot/negative/direct_haxe_value_exception_unsupported` and `haxe.go-14as.39`.

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
- Shim ownership decisions: `docs/stdlib-shim-rationale.md`
- Phase roadmap: `docs/phase2-roadmap.md`
