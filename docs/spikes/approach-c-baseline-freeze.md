# Approach C Baseline Freeze (M0)

Date: 2026-02-25

This snapshot captures the current contract/capability behavior before the Approach C refactor (`contracts + capabilities + lanes`).

## 1. Contract selection (current)

- Contract selector is still profile-shaped: `portable|metal`.
- Resolution entrypoint: `ProfileResolver.resolve()`.
- Default: `portable` when no selector is set.
- Removed selectors hard-error: `gopher`, `idiomatic`, alias defines for both.

References:
- `src/reflaxe/go/ProfileResolver.hx:8`
- `src/reflaxe/go/ProfileResolver.hx:17`
- `src/reflaxe/go/ProfileResolver.hx:48`
- `src/reflaxe/go/ProfileResolver.hx:64`

## 2. Strict boundary defaults (current)

- `CompilerInit.Start()` resolves profile and computes strict default inline.
- In `metal`, strict mode is enabled by default unless `-D reflaxe_go_metal_allow_fallback` is set.
- `reflaxe_go_strict_examples` and `reflaxe_go_strict` are independent toggles.

References:
- `src/reflaxe/go/CompilerInit.hx:16`
- `src/reflaxe/go/CompilerInit.hx:28`
- `src/reflaxe/go/CompilerInit.hx:33`
- `src/reflaxe/go/CompilerInit.hx:34`

## 3. Runtime slicing (current)

- Runtime copy planning lives in `GoReflaxeCompiler.resolveRuntimeCopyPlan`.
- Full copy is default unless selective mode is explicitly enabled.
- Selective mode combines manual features with inferred features from `CompilationContext`.
- Feature defines are orthogonal to profile selection.

References:
- `src/reflaxe/go/GoReflaxeCompiler.hx:28`
- `src/reflaxe/go/GoReflaxeCompiler.hx:160`
- `src/reflaxe/go/GoReflaxeCompiler.hx:205`
- `src/reflaxe/go/GoReflaxeCompiler.hx:215`
- `docs/hxrt-selective-runtime.md:5`
- `docs/hxrt-selective-runtime.md:41`

## 4. Build configuration ownership (current)

Current effective configuration is scattered:

- Contract/profile resolution:
  - `src/reflaxe/go/ProfileResolver.hx`
- Raw native mode:
  - `src/reflaxe/go/RawNativeModeResolver.hx`
- Strict/fallback defaults:
  - `src/reflaxe/go/CompilerInit.hx`
- Runtime plan calculation:
  - `src/reflaxe/go/GoReflaxeCompiler.hx`
- Compile context container:
  - `src/reflaxe/go/CompilationContext.hx`

There is no single typed build context object holding effective contract + capabilities + enforcement policy.

## 5. Metal specialization + fallback behavior (current)

Metal specialization exists for typed go surfaces, but fallback enforcement is not centralized:

- `go.Chan` constructor path in metal falls back when monomorphization cannot produce a concrete type.
- `lowerMetalGoChanCall` returns `null` in several cases; caller falls back to generic lowering path.
- Monomorphizable check currently treats `any` as non-monomorphizable.

References:
- `src/reflaxe/go/GoCompiler.hx:9801`
- `src/reflaxe/go/GoCompiler.hx:9810`
- `src/reflaxe/go/GoCompiler.hx:10612`
- `src/reflaxe/go/GoCompiler.hx:10619`
- `src/reflaxe/go/GoCompiler.hx:10637`
- `src/reflaxe/go/GoCompiler.hx:11481`

## 6. Strict enforcer behavior scope (current)

- Strict mode enforcer blocks raw `__go__` in project sources.
- In metal strict mode it allows framework-owned typed injection callsites from compiler-owned paths.
- The allowlist currently references `std/go/metal` file paths.

References:
- `src/reflaxe/go/macros/StrictModeEnforcer.hx:29`
- `src/reflaxe/go/macros/StrictModeEnforcer.hx:68`
- `src/reflaxe/go/macros/StrictModeEnforcer.hx:122`
- `src/reflaxe/go/macros/StrictModeEnforcer.hx:137`

## 7. Known architecture gaps this freeze preserves

1. No unified `GoBuildContext` with resolved effective contract/capability data.
2. No module-level metal lanes metadata analyzer/enforcement (`@:haxeMetal` pending).
3. Metal fallback policy is not generalized across typed-lowering fallbacks.
4. No deterministic contract/runtime plan reports emitted as build artifacts.
5. Semantic diff harness is portable-only by default; no lane-specific coverage yet.

## 8. Baseline docs in current repo

- Profile overview: `docs/profiles.md`
- Defines matrix: `docs/defines-reference.md`
- Semantics guide: `docs/profile-semantics-guide.md`
- Portable contract: `docs/portable-canonical-contract.md`
- Runtime slicing: `docs/hxrt-selective-runtime.md`
