# Profiles (`-D reflaxe_go_profile=...`)

Use profiles to choose the semantic contract for your build.

```bash
-D reflaxe_go_profile=portable|metal
```

## Terms

- [portable](/docs/glossary.md#portable-profile): portability-first profile.
- [metal](/docs/glossary.md#metal-profile): Go-first profile with stricter defaults.
- [lane](/docs/glossary.md#lane): scoped enforcement zone (for example `@:goMetal` modules).
- [fallback](/docs/glossary.md#fallback): safe path when strict typed lowering cannot apply.

## Model (`contracts + capabilities + planner + lanes`)

`reflaxe.go` treats profile selection as one axis, not the whole build policy:

- contract axis: `portable|metal` (semantic contract)
- boundary axis: strict policy + portable native-import policy
- runtime axis: full copy vs selective feature copy
- planner axis: `reflaxe_go_auto=off|auto|auto_strict`
- lane axis: `@:goMetal` scoped enforcement inside portable builds

Current implementation resolves these once in `GoBuildContextResolver.resolve()`, then `GoReflaxeCompiler` consumes that context at compile start/end and emits deterministic report artifacts when enabled.

## Matrix

| Profile | Best for | Practical behavior |
| --- | --- | --- |
| `portable` (default) | Cross-target-friendly Haxe code | Keeps portability semantics first; native usage can be warned/blocked by policy. |
| `metal` | Go-first lanes and strict native policy | Enables stricter defaults and stronger typed specialization pressure in supported native surfaces. |

Profile does not implicitly select runtime slicing or planner mode.

## Practical policy difference

### Raw `__go__` policy

- `portable`: allowed unless strict policy is enabled.
- `metal`: strict by default (`auto` policy), so raw app-side `__go__` is rejected.

Control flags:

- `-D reflaxe_go_strict`
- `-D reflaxe_go_strict_policy=auto|on|off`
- `-D reflaxe_go_strict_examples`

### Native facade usage (`go.*`) in portable

Portable builds can warn or error when `go.*` is used:

- `-D reflaxe_go_portable_native_policy=warn|error|off`
- `-D reflaxe_go_portable_native_scan_mode=typed|scanner|hybrid`
- `-D reflaxe_go_portable_native_allow=<csv>`

## `@:goMetal` lanes (portable builds)

`@:goMetal` marks modules that must obey metal-clean restrictions even in a portable build.

Current enforced rules:

1. Raw `__go__` is disallowed in `@:goMetal` modules.
2. Under `-D reflaxe_go_auto=auto_strict`, typed-lowering fallback in these modules is disallowed for:
   - `go.Chan`
   - `go.Slice`
   - `go.Map`
   - `go.Result`

This supports incremental migration: you can keep the full app portable while hardening specific modules.

## Runtime policy (additive, not a semantic switch)

Runtime copy planning is orthogonal to profile selection:

- default: full runtime copy
- selective: `-D reflaxe_go_hxrt_features=<csv>` (manual + inferred feature set)
- selective locked: `-D reflaxe_go_hxrt_features=<csv> -D reflaxe_go_hxrt_no_feature_infer` (manual-only, inference off)

Planning is resolved in `GoReflaxeCompiler.resolveRuntimeCopyPlan` using the already-resolved `GoBuildContext`.

## Planner and optimizer controls (additive, not profile switches)

- `-D reflaxe_go_auto=off|auto|auto_strict`
- `-D reflaxe_go_opt=portable_fast|none`
- `-D reflaxe_go_opt_go_concurrency_fastpath=0|1|off|on|false|true`

These flags tune lowering behavior but do not silently change profile semantics.

## Report artifacts

Optional report defines:

- `-D reflaxe_go_contract_report` -> `profile_contract.json`, `profile_contract.md`
- `-D reflaxe_go_runtime_plan_report` -> `hxrt_plan.json`, `hxrt_plan.md`
- `-D reflaxe_go_optimizer_plan_report` -> `optimizer_plan.json`, `optimizer_plan.md`

Use these reports to audit:

- active profile and strictness
- fallback counts and reasons
- lane vs non-lane fallback attribution
- optimizer capability outcomes

## Recommended defaults

For most teams:

1. Start with `portable`.
2. Set `reflaxe_go_portable_native_policy=error` in CI/release.
3. Promote only performance-critical or native-heavy modules toward metal lanes.
4. Use `metal` for Go-first deployments that need stricter native policy.

## Validation commands

```bash
python3 test/run-snapshots.py
python3 test/run-semantic-diff.py
python3 test/run-semantic-diff.py --suite lanes
python3 test/run-ci.py
```

## Related docs

- Docs map: [docs/index.md](index.md)
- Glossary: [docs/glossary.md](glossary.md)
- Profile semantics deep guide: [docs/profile-semantics-guide.md](profile-semantics-guide.md)
- Portable contract: [docs/portable-canonical-contract.md](portable-canonical-contract.md)
- Versioned semantics spec: [docs/portable-semantics-v1.md](portable-semantics-v1.md)
- Defines reference: [docs/defines-reference.md](defines-reference.md)
- Examples matrix: [docs/examples-matrix.md](examples-matrix.md)
