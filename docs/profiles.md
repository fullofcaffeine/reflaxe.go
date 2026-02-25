# Profiles (`-D reflaxe_go_profile=...`)

This target supports two profiles:

```bash
-D reflaxe_go_profile=portable|metal
```

## Matrix

| Profile | Best for | Behavior contract |
| --- | --- | --- |
| `portable` (default) | Haxe-first and cross-target code | Stable Haxe-oriented semantics, portability-first output (includes former gopher-safe optimizations) |
| `metal` (experimental) | Teams needing typed low-level interop lane | `portable` + strict default app-boundary policy + typed framework interop façade |

## Why two profiles

- `portable` is the semantic baseline.
- `metal` is the explicit low-level performance/interop lane.
- Canonical portable semantics are documented in `docs/portable-canonical-contract.md`.
- World-class profile strategy, semantics, and migration guidance:
  `docs/profile-semantics-guide.md`.

Selective `hxrt` runtime inference is complementary and does not replace profile contracts.
See `docs/hxrt-selective-runtime.md`.

`auto` spike decision (why `portable+metal` remains canonical and how a future additive planner would work):
`docs/profile-auto-spike.md`.

## Removed selectors

- `-D reflaxe_go_profile=gopher` is removed; use `portable`.
- `-D reflaxe_go_gopher` is removed; use `reflaxe_go_profile=portable`.
- `-D reflaxe_go_profile=idiomatic` is removed; use `portable`.
- `-D reflaxe_go_idiomatic` is removed; use `reflaxe_go_profile=portable`.

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

Framework-owned typed facades are allowed in `metal` strict mode; raw app-side injection remains disallowed.

## `@:haxeMetal` lanes (portable builds)

`@:haxeMetal` marks module islands that must obey metal-clean restrictions even when the build contract is `portable`.

- Current enforced rule: raw `__go__` is disallowed in `@:haxeMetal` modules under portable contract.
- Snapshot coverage: `test/snapshot/negative/haxe_metal_lane_injection`.

Lane module discovery is deterministic and emitted in profile contract reports.

## Contract/runtime reports

Opt-in report defines:

- `-D reflaxe_go_contract_report` -> `profile_contract.json`, `profile_contract.md`
- `-D reflaxe_go_runtime_plan_report` -> `hxrt_plan.json`, `hxrt_plan.md`

Snapshot coverage: `test/snapshot/core/report_artifacts_basic`.

## Example references

- Cross-profile micro app: `examples/profile_storyboard`
- Cross-profile complex app: `examples/tui_todo`
- Worker pool/select-style concurrency app: `examples/worker_pool_select`
- Coverage + artifact matrix: `docs/examples-matrix.md`
- Production caveats: `docs/known-gaps.md`
- Canonical profile playbook: `docs/profile-semantics-guide.md`
