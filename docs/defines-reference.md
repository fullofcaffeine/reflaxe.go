# Defines Reference (`-D ...`)

## Core output

- `go_output=<dir>`
  - Required signal for Go generation.
- `go_module=<module/path>`
  - Optional Go module path for generated `go.mod` and runtime imports (`<module/path>/hxrt`).
  - Defaults to `snapshot` when omitted.
  - Runtime details: `docs/hxrt-runtime.md`
- `reflaxe.dont_output_metadata_id`
  - Recommended for deterministic snapshots.

## Build controls

- `go_no_build` or `go_codegen_only`
  - Codegen only; skip backend `go build` step.
- `go_cmd=<binary>`
  - Override Go CLI used by backend build step (default: `go`).
- `go_build_output=<path>`
  - Optional output path passed as `go build -o <path>`.
- `reflaxe_go_line_directives`
  - Emit `//line <module>.hx:<line>` directives for user functions/constructors so panics and traces resolve to Haxe source locations.

## Profiles

- `reflaxe_go_profile=portable|metal`
  - Main profile selector.
- `reflaxe_go_portable`
  - Alias selector for portable.
- `reflaxe_go_metal`
  - Alias selector for metal.

Design note:

- `auto` is not a semantic profile in this repo.
- Spike rationale and future additive-planner shape: `docs/profile-auto-spike.md`.

## IO encoding

- `reflaxe_go_raw_native_mode=interp|utf16le`
  - Controls `haxe.io.Encoding.RawNative` conversion strategy in generated IO/Bytes shims.
  - `interp` (default): match Haxe `--interp` behavior (RawNative treated like UTF-8 conversion path).
  - `utf16le`: opt-in compatibility mode that encodes/decodes RawNative as UTF-16LE bytes (useful when aligning with Java/C#-style RawNative expectations).

## Runtime slicing

- `reflaxe_go_hxrt_default_features`
  - Force legacy full runtime copy (`runtime/hxrt/**`) into output.
  - Takes precedence over selective runtime defines.
- `reflaxe_go_hxrt_features=<csv>`
  - Enable selective runtime mode and add manual feature names (for example `core,string,sys`).
  - Unknown feature names fail compilation.
  - Empty value is allowed (`-D reflaxe_go_hxrt_features=`) and means "manual list empty; use inference unless disabled."
- `reflaxe_go_hxrt_no_feature_infer`
  - Enable selective runtime mode and disable inferred features from compiler analysis.
  - Resulting runtime set is `core` plus any manual `reflaxe_go_hxrt_features`.

## Optimizer controls

- `reflaxe_go_auto=off|auto|auto_strict`
  - Explicit auto-lowering planner mode.
  - This is additive and does **not** switch semantic contract/profile.
  - Default: `off`.
  - `auto`: allow deterministic planner-driven lowering attempts with normal fallback behavior.
  - `auto_strict`: reserved strict planner mode for fail-fast lane policies.
- `reflaxe_go_opt=portable_fast|none`
  - Additive optimizer preset (not a semantic profile switch).
  - Default: `portable_fast`.
  - `portable_fast` enables semantics-preserving portable convergence fastpaths.
  - `none` disables preset-driven fastpaths.
- `reflaxe_go_opt_go_concurrency_fastpath[=0|1|false|true|off|on]`
  - Capability toggle for typed go-concurrency fastpath (`go.Chan` specialization path) in portable builds.
  - Defaults to `on` when `reflaxe_go_opt=portable_fast`, otherwise `off`.
  - Example disable: `-D reflaxe_go_opt_go_concurrency_fastpath=0`.
  - Example explicit override with preset off: `-D reflaxe_go_opt=none -D reflaxe_go_opt_go_concurrency_fastpath=1`.

## Strictness

- `reflaxe_go_strict`
  - Enforce strict no raw `__go__` policy in app project sources.
- `reflaxe_go_strict_examples`
  - Enforce strict no raw `__go__` policy for repo examples/snapshots.
- `reflaxe_go_metal_allow_fallback`
  - Opt-in escape hatch for `metal` profile to allow typed-specialization fallback instead of hard error.
  - Also disables metal strict-by-default app-boundary enforcement.
  - Does not affect `reflaxe_go_strict` (explicit strict mode still forbids raw injection).
  - Does not affect strict examples enforcement in snapshots/examples.

## Portable native-import gate

- `reflaxe_go_portable_native_policy=warn|error|off`
  - Policy for target-native `go.*` usage when compiling with `reflaxe_go_profile=portable`.
  - `warn` (default): emit warnings for portable-contract modules using `go.*`.
  - `error`: fail compilation on portable-contract modules using `go.*` (recommended for CI/release).
  - `off`: disable portable native-import checks.
- `reflaxe_go_portable_native_allow=<csv>`
  - Optional comma-separated module-prefix allowlist for portable builds (for sanctioned adapter islands).
  - Example: `-D reflaxe_go_portable_native_allow=app.adapters.go,app.platform`.

## Lane metadata

- `@:goMetal`
  - Module/type/field metadata for metal-clean lane islands inside portable builds.
  - Portable builds enforce lane restrictions for these modules:
    - raw `__go__` is disallowed.
    - typed fallback paths for `go.Chan`/`go.Slice`/`go.Map`/`go.Result` are disallowed (for example non-monomorphizable `Dynamic`/`Any` lane usage).
  - Validation commands:
    - `python3 test/run-snapshots.py --case negative/go_metal_lane_injection --case negative/go_metal_lane_fallback_result --case negative/go_metal_lane_fallback_chan --case negative/go_metal_lane_fallback_slice --case negative/go_metal_lane_fallback_map --case core/go_metal_lane_nonlane_fallback_allowed`
    - `python3 test/run-semantic-diff.py --suite lanes`
    - `npm run test:semantic-diff:lanes`

Removed:

- `@:haxeMetal` -> compile error, use `@:goMetal`.

## Reports

- `reflaxe_go_contract_report`
  - Emit `profile_contract.json` and `profile_contract.md` into output root with effective contract/capability state.
  - `profile_contract.json` schema v5 includes `autoLoweringMode`, lowering-decision ledger fields (`loweringDecisionCount`, `loweringDecisionAttemptCount`, `loweringDecisionSuccessCount`, `loweringDecisionFallbackCount`, `loweringDecisions`), structured `metalFallbackViolations`, and deterministic lane summary fields (`metalFallbackLaneViolationCount`, `metalFallbackNonLaneViolationCount`, `metalFallbackViolationsByModule`).
- `reflaxe_go_runtime_plan_report`
  - Emit `hxrt_plan.json` and `hxrt_plan.md` into output root with selected runtime features/files and selection reasons.
- `reflaxe_go_optimizer_plan_report`
  - Emit `optimizer_plan.json` and `optimizer_plan.md` into output root with effective optimizer preset/capabilities and applied lowering counters.
  - `optimizer_plan.json` schema v2 includes `autoLoweringMode`.

## Constructor devex

- `reflaxe_go_auto_empty_ctor_interfaces=<csv>`
  - Auto-inject `public function new() {}` into classes implementing listed interfaces.
  - Value is a comma-separated list of fully-qualified interface paths (for example `app.runtime.PulseRuntime`).
  - Keeps app classes free from repeated empty constructor boilerplate while preserving explicit opt-in behavior.

## Pass registry

- `go_granular_pass_registry`
  - Use granular pass bundle.
- `reflaxe_go_test_registry_case=<duplicate|missing_dep|cycle>`
  - Test-only define used by negative snapshot cases for registry validation.
