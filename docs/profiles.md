# Profiles (`-D reflaxe_go_profile=...`)

This target supports two profiles:

```bash
-D reflaxe_go_profile=portable|metal
```

## Matrix

| Profile | Best for | Behavior contract |
| --- | --- | --- |
| `portable` (default) | Haxe-first and cross-target code | Stable Haxe-oriented semantics and portability-first output |
| `metal` (experimental) | Teams needing typed low-level interop lane | `portable` + strict default app-boundary policy + typed framework interop façade |

## Why two profiles

- `portable` is the semantic baseline.
- `metal` is the explicit low-level performance/interop lane.
- Canonical portable semantics are documented in `docs/portable-canonical-contract.md`.
- Versioned portable semantics rules are documented in `docs/portable-semantics-v1.md`.
- World-class profile strategy, semantics, and migration guidance:
  `docs/profile-semantics-guide.md`.

Selective `hxrt` runtime inference is complementary and does not replace profile contracts.
See `docs/hxrt-selective-runtime.md`.

`auto` spike decision (why `portable+metal` remains canonical and how a future additive planner would work):
`docs/profile-auto-spike.md`.

## Metal-ready subset (current)

Use `metal` for bounded hot paths that are already covered by typed specialization lanes:

- `go.Chan<T>`
- `go.Slice<T>`
- `go.Map<K,V>`
- `go.Result<T>`

If your code path falls outside this subset, start from `portable` and promote only after benchmark evidence.

## Boundary policy

- `reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `reflaxe_go_strict`: forbids raw `__go__` in app project sources.
- `metal` enables strict mode by default for app-side raw injection.
- `reflaxe_go_metal_allow_fallback`: opt-in escape hatch that relaxes metal hard-error fallback policy and disables metal strict-by-default boundary enforcement.
- `reflaxe_go_portable_native_policy=warn|error|off`: policy for `go.*` usage under portable contract (`warn` default, `error` recommended in CI/release).
- `reflaxe_go_portable_native_allow=<csv>`: optional portable allowlist for sanctioned native adapter modules.

Framework-owned typed facades are allowed in `metal` strict mode; raw app-side injection remains disallowed.

## `@:goMetal` lanes (portable builds)

`@:goMetal` marks module islands that must obey metal-clean restrictions even when the build contract is `portable`.

- Current enforced rule: raw `__go__` is disallowed in `@:goMetal` modules under portable contract.
- Typed fallback rule: `go.Chan` / `go.Slice` / `go.Map` / `go.Result` calls that would fall back from typed specialization (for example `Dynamic`/`Any` paths) are disallowed in `@:goMetal` modules under portable contract.
- Snapshot coverage:
  - `test/snapshot/negative/go_metal_lane_injection`
  - `test/snapshot/negative/go_metal_lane_fallback_result`
  - `test/snapshot/negative/go_metal_lane_fallback_chan`
  - `test/snapshot/negative/go_metal_lane_fallback_slice`
  - `test/snapshot/negative/go_metal_lane_fallback_map`
  - `test/snapshot/core/go_metal_lane_nonlane_fallback_allowed`

Lane module discovery is deterministic and emitted in profile contract reports.

Lane test commands:

- `python3 test/run-snapshots.py --case negative/go_metal_lane_injection --case negative/go_metal_lane_fallback_result --case negative/go_metal_lane_fallback_chan --case negative/go_metal_lane_fallback_slice --case negative/go_metal_lane_fallback_map --case core/go_metal_lane_nonlane_fallback_allowed`
- `python3 test/run-semantic-diff.py --suite lanes`
- `npm run test:semantic-diff:lanes`
- `python3 test/run-ci.py --force-semantic-diff-lanes`

## Contract/runtime reports

Opt-in report defines:

- `-D reflaxe_go_contract_report` -> `profile_contract.json`, `profile_contract.md`
- `-D reflaxe_go_runtime_plan_report` -> `hxrt_plan.json`, `hxrt_plan.md`
- `-D reflaxe_go_optimizer_plan_report` -> `optimizer_plan.json`, `optimizer_plan.md`

## Portable convergence optimizer controls

- `-D reflaxe_go_auto=off|auto|auto_strict`
  - explicit auto-lowering planner mode (additive, not a semantic profile switch).
  - default: `off`.
- `-D reflaxe_go_opt=portable_fast|none`
  - default: `portable_fast`
  - additive optimizer preset (not a semantic profile).
- `-D reflaxe_go_opt_go_concurrency_fastpath=0|1|off|on|false|true`
  - typed go-concurrency fastpath capability in portable builds.
  - defaults to `on` with `portable_fast`.

`profile_contract.json` (schema v5) includes `autoLoweringMode`, lowering-decision ledger fields (`loweringDecisionCount`, `loweringDecisionAttemptCount`, `loweringDecisionSuccessCount`, `loweringDecisionFallbackCount`, `loweringDecisions`), structured `metalFallbackViolations`, and deterministic lane summaries:
- `metalFallbackLaneViolationCount`
- `metalFallbackNonLaneViolationCount`
- `metalFallbackViolationsByModule`

Snapshot coverage:
- `test/snapshot/core/report_artifacts_basic`
- `test/snapshot/core/report_artifacts_lane_fallback`

## Example references

- Cross-profile micro app: `examples/profile_storyboard`
- Cross-profile complex app: `examples/tui_todo`
- Worker pool/select-style concurrency app: `examples/worker_pool_select`
- Coverage + artifact matrix: `docs/examples-matrix.md`
- Production caveats: `docs/known-gaps.md`
- Canonical profile playbook: `docs/profile-semantics-guide.md`
