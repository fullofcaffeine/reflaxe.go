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

Removed:

- `reflaxe_go_profile=gopher` -> compile error, use `portable`.
- `reflaxe_go_gopher` -> compile error, use `reflaxe_go_profile=portable`.
- `reflaxe_go_profile=idiomatic` -> compile error, use `portable`.
- `reflaxe_go_idiomatic` -> compile error, use `reflaxe_go_profile=portable`.

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

## Strictness

- `reflaxe_go_strict`
  - Enforce strict no raw `__go__` policy in app project sources.
- `reflaxe_go_strict_examples`
  - Enforce strict no raw `__go__` policy for repo examples/snapshots.
- `reflaxe_go_metal_allow_fallback`
  - Opt-in escape hatch for `metal` profile to allow raw `__go__` in app sources.
  - Does not affect `reflaxe_go_strict` (explicit strict mode still forbids raw injection).
  - Does not affect strict examples enforcement in snapshots/examples.

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
