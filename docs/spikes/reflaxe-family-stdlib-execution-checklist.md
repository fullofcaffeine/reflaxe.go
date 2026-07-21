# `reflaxe.family.std` Execution Checklist (Post-Spike)

Status: local bootstrap complete; external cutover blocked on shared-content identity

Source spike: `docs/spikes/reflaxe-family-stdlib-extraction-spike.md`  
Tracking task: `haxe.go-cgk.14`

This checklist converts the approved spike shape into an execution-ready plan with hard gates.

## Decision Checklist (Must Be Explicit)

| Decision area | Decision | Status | Evidence |
| --- | --- | --- | --- |
| Contract baseline | Portable semantics contract remains explicit and versioned (`portable-semantics-v1`) | Locked | `docs/portable-semantics-v1.md` |
| Portable module scope | Portable-eligible module scope remains allowlist-driven | Locked | `test/portable_allowlist.json` |
| Tier1 conformance model | Tier1 stays deterministic module->semantic-diff mapping | Locked | `test/portable_conformance_tier1.json`, `test/run-portable-conformance.py` |
| Ownership mapping | Tier1 ownership split (`haxe_source`/`runtime_binding`/`compiler_intrinsic`/`mixed`) | Locked | `docs/portable-module-mapping-contract.md` |
| Full parity visibility | Full portable-eligible sweep + parity closure summary required in full CI runs | Locked | `test/run-ci.py`, `test/run-portable-parity-closure.py` |
| Provenance governance | Upstream boundary + provenance ledger checks remain mandatory | Locked | `scripts/ci/upstream-stdlib-boundary-check.js`, `scripts/ci/stdlib-provenance-ledger-check.js` |
| Family package shape | Local mirror shape is implemented; immutable shared core plus target-overlay cut line is required before external extraction | Reassessment required | `docs/spikes/reflaxe-family-stdlib-extraction-spike.md` |

## 2026-07-20 Reassessment

The completed Phase-1 work is a useful **local mirror check**: CI proves that
`haxe.go`'s canonical portable artifacts and its bundled copy agree. It is not a
cross-repository package-consumption check.

`haxe.rust` independently contains the same bootstrap version and mirror
machinery, but several payload files differ from Go. `genes` and
`haxe.elixir.codex` have no family directory and instead enforce compatibility
through their own reports, inventories, differential tests, runtime tests, and
snapshots. Elixir already has a documented portable-versus-Elixir-first
authoring contract; it does not need a global `portable|metal` selector merely
to qualify for shared semantic fixtures.

Before external extraction or a cutover can be approved:

1. define the immutable family-core file set;
2. reconcile the Go and Rust copies of that core;
3. record and verify one content digest for each family-core version;
4. move target-specific allowlists, mappings, fixture bindings, and evidence to
   target-qualified overlays;
5. prove at least two repositories consume the exact same core payload;
6. treat Genes, Elixir, and other sibling adoption as opt-in based on a concrete
   reduction in duplicated governance—not as evidence needed to validate the
   local Go compiler.

Until those conditions hold, `source: in-repo-bootstrap` and `mode: dual-run`
mean “compare this repository with its own mirror,” not “consume a released
family authority.”

## Local Bootstrap Verification Gates

The implemented local bootstrap remains valid only while these gates pass:

1. `npm run test:ci`
2. `npm run test:portable-conformance`
3. `npm run test:portable-parity-closure`
4. `npm run test:stdlib-sweep:full`
5. `npm run test:stdlib:governance`

Active dual-run pin and report artifacts:

- `family/family_std_pin.json`
- `test/.test-cache/family_std_dual_run_report.json`
- `test/.test-cache/family_std_dual_run_report.md`

## Completed Local Bootstrap Sequence

The original spike produced these now-completed tasks:

1. `haxe.go-cgk.15` - bootstrap `reflaxe.family.std` skeleton + schema pack
2. `haxe.go-cgk.16` - implement export/import sync tooling for `haxe.go`
3. `haxe.go-cgk.17` - pin family package in `haxe.go` with dual-run checks
4. `haxe.go-cgk.18` - sibling compiler rollout gate plan

Sibling rollout plan artifact:

- `docs/spikes/reflaxe-family-sibling-rollout-gate-plan.md`

Their closure proves that the local skeleton, synchronization tooling, pin, and
rollout-plan artifact were created. It does not prove that an external family
release exists or that another repository consumes identical contents.

## Historical Bootstrap Order

1. Phase-1 bootstrap (`cgk.15`)
2. `haxe.go` sync tooling (`cgk.16`)
3. `haxe.go` pinned consumption + dual-run (`cgk.17`)
4. sibling adoption planning (`cgk.18`)

## Cut Lines (Do Not Cross in Phase-1)

1. No backend runtime extraction.
2. No target-native facade extraction (`go.*`, `rust.*`, `ocaml.*`, `elixir.*` remain local).
3. No semantic-contract inference replacing explicit profile/contract selection.
4. No cross-repo rollout before `haxe.go` dual-run stability is demonstrated.

## External Extraction Approval

The original approval authorized the local bootstrap sequence. Before external
extraction or sibling cutover is approved, record:

1. approver
2. approval date
3. the reviewed immutable core file set and digest policy
4. the target-overlay identity and ownership policy
5. evidence that two consumers verify the same core payload
6. any narrowed scope or cut-line updates

Append that record here. Do not infer external approval from the closed local
bootstrap tasks.
