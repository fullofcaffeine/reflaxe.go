# Defines Reference (`-D ...`)

## Core output

- `go_output=<dir>`
  - Required signal for Go generation.
  - Compiler-managed source, reports, metadata, and copied runtime files are
    confined below this canonical root. See
    [generated-output confinement](generated-output-confinement.md).
- `go_module=<module/path>`
  - Optional Go module path for generated `go.mod` and runtime imports (`<module/path>/hxrt`).
  - Defaults to `snapshot` when omitted.
  - Runtime details: `docs/hxrt-runtime.md`
- `reflaxe.dont_output_metadata_id`
  - Recommended for deterministic snapshots.

## Build controls

- `go_no_build` or `go_codegen_only`
  - Explicit codegen-only mode; skip the backend-owned `go build` step.
  - The caller must run its own Go build/test gate before treating the output
    as successful.
- `go_cmd=<binary>`
  - Override Go CLI used by backend build step (default: `go`).
- `go_build_output=<path>`
  - Optional output path passed as `go build -o <path>`.
  - This is an explicit caller-authorized binary sink, not a compiler-managed
    generated-source path; it may intentionally be outside `go_output`.
- Without a codegen-only define, failure to launch `go build` or any nonzero
  exit is a fatal Haxe compiler outcome. Child stderr remains visible, while
  the compiler diagnostic omits generated-output and machine-local command
  paths.
- `reflaxe_go_line_directives`
  - Emit `//line <module>.hx:<line>` directives for user functions/constructors so panics and traces resolve to Haxe source locations.

## Profiles

- `reflaxe_go_profile=portable|metal`
  - Compatibility selector for the `portable_default` or `metal_compatibility`
    policy preset.
  - It does not choose a second semantic backend. Typed APIs/externs and
    `@:goNative` define Go-native source boundaries.
- `reflaxe_go_portable`
  - Alias selector for portable.
- `reflaxe_go_metal`
  - Alias selector for metal.

Design note:

- `metal` remains supported without a deprecation warning. See the
  [retention decision](metal-preset-retention-decision.md) for evidence and
  future reopen criteria.
- `auto` is a planner mode, not a semantic profile.
- Canonical contract: `docs/native-policy-presets.md`.

## Native policy axes

- `reflaxe_go_native_authority=guarded|explicit`
  - `guarded` applies `reflaxe_go_portable_native_policy` outside explicit
    `@:goNative` modules.
  - `explicit` accepts typed native APIs without that diagnostic gate.
  - Preset defaults: portable `guarded`; metal compatibility `explicit`.
- `reflaxe_go_native_specialization=proven|eager`
  - `proven` attempts typed representations only through enabled capability
    paths.
  - `eager` attempts every supported typed `go.Chan`, `go.Slice`, `go.Map`, and
    `go.Result` specialization.
  - Preset defaults: portable `proven`; metal compatibility `eager`.
- `reflaxe_go_native_fallback=allow|error`
  - `allow` uses and reports a semantics-safe fallback.
  - `error` rejects user-owned native specialization fallback sites.
  - Preset defaults: portable `allow`; metal compatibility `error`.

Resolution precedence is canonical axis define, then the compatible metal
fallback alias where applicable, then the preset default. Invalid values and
contradictory fallback inputs are compile errors. Effective values and sources
are present in report artifacts.

## IO encoding

- `reflaxe_go_raw_native_mode=interp|utf16le`
  - Controls `haxe.io.Encoding.RawNative` conversion strategy in generated IO/Bytes shims.
  - `interp` (default): match Haxe `--interp` behavior (RawNative treated like UTF-8 conversion path).
  - `utf16le`: opt-in compatibility mode that encodes/decodes RawNative as UTF-16LE bytes (useful when aligning with Java/C#-style RawNative expectations).

## Diagnostics

- `reflaxe_go_native_stack_trace`
  - Enables Go-native stack capture for `haxe.CallStack` / `haxe.NativeStackTrace`.
  - This is an explicit target-sensitive diagnostic capability, not portable semantic-diff parity.
  - Default behavior remains deterministic empty stacks.
  - Runtime details: `docs/spikes/native-stack-capture-contract.md`

## Runtime slicing

- `reflaxe_go_hxrt_default_features`
  - Force legacy full runtime copy (`runtime/hxrt/**`) into output.
  - Footprint-explicit diagnostic files, such as native stack capture, still
    require their own define before they are copied.
  - Takes precedence over selective runtime defines.
- `reflaxe_go_hxrt_features=<csv>`
  - Enable selective runtime mode and add manual feature names (for example `core,string,sys` for root OS capabilities, `core,string,file_io` for `sys.io.File`, `core,string,process` for `sys.io.Process`, `core,string,filesystem` for `sys.FileSystem`, `core,string,bytes,exception,socket,http` for `sys.Http`, `core,string,exception,socket` for DNS/TCP/UDP, `core,string,exception,bytes,ssl` for SSL leaf helpers, or `core,string,exception,bytes,socket,ssl,socket_ssl` for TLS sockets).
  - Unknown feature names fail compilation.
  - Empty value is allowed (`-D reflaxe_go_hxrt_features=`) and means "manual list empty; use inference unless disabled."
- `reflaxe_go_hxrt_no_feature_infer`
  - Enable selective runtime mode and disable inferred features from compiler analysis.
  - Resulting runtime set is `core` plus any manual `reflaxe_go_hxrt_features`.

## Optimizer controls

- `reflaxe_go_auto=off|auto|auto_strict`
  - Explicit auto-lowering planner mode.
  - This is additive and does **not** switch source semantics or policy preset.
  - Default: `off`.
  - `auto`: allow deterministic planner-driven lowering attempts with normal fallback behavior.
  - `auto_strict`: same planner attempts, plus fail-fast fallback checks inside
    explicit native-boundary modules.
  - Under `native_specialization=proven`, `auto|auto_strict` enable typed
    specialization attempts for `go.Slice` / `go.Map` / `go.Result`; outcomes
    are recorded in contract lowering ledgers.
  - Under `native_specialization=proven`, `go.Chan` typed specialization remains
    controlled by `reflaxe_go_opt_go_concurrency_fastpath`.
- `reflaxe_go_opt=portable_fast|none`
  - Additive optimizer preset (not a semantic switch).
  - Default: `portable_fast`.
  - `portable_fast` enables semantics-preserving portable convergence fastpaths.
  - `none` disables preset-driven fastpaths.
- `reflaxe_go_opt_go_concurrency_fastpath[=0|1|false|true|off|on]`
  - Capability toggle for the proven typed go-concurrency fastpath (`go.Chan`
    specialization path); eager specialization does not need this toggle.
  - Defaults to `on` when `reflaxe_go_opt=portable_fast`, otherwise `off`.
  - Example disable: `-D reflaxe_go_opt_go_concurrency_fastpath=0`.
  - Example explicit override with preset off: `-D reflaxe_go_opt=none -D reflaxe_go_opt_go_concurrency_fastpath=1`.

## Strictness

- `reflaxe_go_strict`
  - Enforce strict no raw `__go__` policy in app project sources.
- `reflaxe_go_strict_policy=auto|on|off`
  - Explicit app-boundary strictness policy axis.
  - `auto` (default): `metal_compatibility` -> strict on,
    `portable_default` -> strict off.
  - `on`: strict on regardless of preset.
  - `off`: strict off regardless of preset.
  - `reflaxe_go_strict` remains a compatibility alias for forcing strict `on`.
- `reflaxe_go_strict_examples`
  - Enforce strict no raw `__go__` policy for repo examples/snapshots.
- `reflaxe_go_metal_allow_fallback`
  - Compatibility alias selecting native fallback `allow` when
    `reflaxe_go_native_fallback` is absent.
  - Canonical `allow` may be supplied alongside it and keeps canonical report
    provenance; canonical `error` conflicts and fails compilation.
  - Does not change strict boundary policy.
  - Does not affect strict examples enforcement in snapshots/examples.

## Portable native-import gate

- `reflaxe_go_portable_native_policy=warn|error|off`
  - Compatibility-named policy for target-native `go.*` usage whenever native
    authority is `guarded`.
  - `warn` (default): emit warnings outside `@:goNative` modules.
  - `error`: fail compilation outside `@:goNative` modules (recommended for
    guarded CI/release builds).
  - `off`: disable portable native-import checks.
- `reflaxe_go_portable_native_scan_mode=typed|scanner|hybrid`
  - Detection mode for portable native-import policy:
    - `typed` (default): typed analyzer traversal.
    - `scanner`: explicit source scanner (`import`/`using`) for deterministic low-noise module hits.
    - `hybrid`: union of `typed` + `scanner` hits.
- `reflaxe_go_portable_native_allow=<csv>`
  - Compatibility-named comma-separated module-prefix allowlist used when
    native authority is `guarded` (for sanctioned adapter islands).
  - Example: `-D reflaxe_go_portable_native_allow=app.adapters.go,app.platform`.

## Native-boundary metadata

- `@:goNative`
  - Canonical module/type/field metadata for explicit Go-native module
    boundaries under either preset.
  - The owning module:
    - is exempt from guarded unscoped-`go.*` diagnostics;
    - raw `__go__` is disallowed.
    - with `-D reflaxe_go_auto=auto_strict`, typed fallback paths for `go.Chan`/`go.Slice`/`go.Map`/`go.Result` are disallowed (for example non-monomorphizable `Dynamic`/`Any` native-boundary usage).
    - typed-specialization eligibility is centralized in `src/reflaxe/go/compiler/GoNativeTypeEligibility.hx` (concrete non-`any` type required; nullable primitive dynamic-path types excluded; `go.Map` keys must be comparable).
    - with `reflaxe_go_auto=off|auto`, typed fallback paths are allowed and recorded in lowering/report artifacts.
- `@:goMetal`
  - Silent compatibility alias for `@:goNative`; new source should use
    `@:goNative`.
  - Validation commands:
    - `python3 test/run-snapshots.py --case core/native_boundary_guarded_authority`
    - `python3 test/run-snapshots.py --case negative/go_metal_lane_injection`
    - `python3 test/run-semantic-diff.py --suite lanes`
    - `npm run test:semantic-diff:lanes`

Removed:

- `@:haxeMetal` -> compile error, use `@:goNative`.

- `@:goAllowRaw`
  - Module/type metadata for framework-owned low-level abstraction islands that need raw
    `__go__` even when strict boundary enforcement is enabled.
  - Use this sparingly and document the boundary with `Why / What / How` HaxeDoc where the
    abstraction is declared.
  - Intended scope:
    - staged std overrides,
    - runtime bindings,
    - narrow target-owned helper modules.
  - Not intended scope:
    - application business logic,
    - examples/snapshots as user-facing coding style.
  - Does not bypass `@:goNative` boundary restrictions.
  - `__go__` still does not infer package imports; if the raw snippet needs external Go
    packages, carry imports through typed extern metadata (`@:go.import`, `@:go.name`,
    `@:go.receiver`) or existing framework-owned wrappers.

## Reports

- `reflaxe_go_contract_report`
  - Emit `profile_contract.json` and `profile_contract.md` into output root with effective contract/capability state.
  - `profile_contract.json` schema v8 adds `policyPreset`,
    `semanticBoundarySource`, every `native*Policy` value and source,
    `nativeBoundaryModules`, `nativeFallbackEvent*`, and `inNativeBoundary`.
  - Historical `contract`, `metalLaneModules`, `metalFallbackViolation*`, and
    `inMetalLane` fields remain compatibility aliases.
- `reflaxe_go_runtime_plan_report`
  - Emit `hxrt_plan.json` and `hxrt_plan.md` into output root with selected runtime features/files and selection reasons.
  - `hxrt_plan.json` schema v2 adds `policyPreset` and
    `semanticBoundarySource`.
  - `hxrt_plan` reasons are deterministic per-feature provenance entries (`baseline`, `class_usage`, `enum_usage`, `shim_group`, `io_helper_surface`, `manual_define`, `dependency_edge`) so runtime selection remains auditable in CI.
- `reflaxe_go_optimizer_plan_report`
  - Emit `optimizer_plan.json` and `optimizer_plan.md` into output root with effective optimizer preset/capabilities and applied lowering counters.
  - `optimizer_plan.json` schema v6 adds `policyPreset`, native
    specialization value/provenance, and canonical boundary/non-boundary
    fallback counters. Historical lane counters remain aliases.
- `reflaxe_go_type_usage_report`
  - Emit `type_usage.json` and `type_usage.md` with deterministic
    compiler-observed generic/function/anonymous shapes, typed member/call
    locations, metadata and resolved native import paths, and resulting
    runtime-capability reasons.
  - The schema uses normalized `Module:line` locations and contains no macro
    objects or absolute source paths.
  - See [Typed usage ledger](typed-usage-ledger.md).

## Constructor devex

- `reflaxe_go_auto_empty_ctor_interfaces=<csv>`
  - Auto-inject `public function new() {}` into classes implementing listed interfaces.
  - Value is a comma-separated list of fully-qualified interface paths (for example `app.runtime.PulseRuntime`).
  - Keeps app classes free from repeated empty constructor boilerplate while preserving explicit opt-in behavior.

## Pass registry

- `go_granular_pass_registry`
  - Compatibility fallback: use legacy fixed granular pass bundle.
- `reflaxe_go_legacy_pass_bundle`
  - Compatibility fallback: use legacy fixed lean pass bundle.
- (default, no compatibility define)
  - Planner-driven deterministic pass selection keyed by contract/auto/optimizer build context.
- `reflaxe_go_test_registry_case=<duplicate|missing_dep|cycle>`
  - Test-only define used by negative snapshot cases for registry validation.
