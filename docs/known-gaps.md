# Known Gaps And Production Caveats

This page is the blunt status view for current limitations so teams can plan migrations with clear tradeoffs.

## Compiler/output caveats

- Output remains a single Go package (multi-file, single package); multi-package emission is not implemented yet.
- Multi-package output is currently deferred as non-blocking for production GA; explicit boundary conditions for re-opening are documented in `docs/multi-package-output-evaluation.md`.
- Some unsupported typed-expression guards still exist by design (`docs/feature-support-matrix.md` inventory section). Current compiler hard-fails:
  - `Unsupported assignment target` (`lowerLValue`)
  - `Unsupported postfix unary operator` (`lowerExpr` / `lowerExprWithPrefix`)
  - `Unsupported expression` (catch-all `lowerExpr` fallback)
  - `Std.isOfType` still has conservative fallback behavior for unresolved runtime-value abstract targets (documented as partial support, not a hard-fail)
- `go.*` APIs are target-specific. They compile to real Go behavior on this target, but they are not portability-safe across non-Go Haxe targets.

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
