# `reflaxe.family.std` Execution Checklist (Post-Spike)

Status: ready-for-approval execution checklist  
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
| Family package shape | Extraction target repo shape from spike adopted for Phase-1 | Proposed -> execute | `docs/spikes/reflaxe-family-stdlib-extraction-spike.md` |

## Execution Gates

Do not start extraction execution tasks until all gates pass:

1. `npm run test:ci`
2. `npm run test:portable-conformance`
3. `npm run test:portable-parity-closure`
4. `npm run test:stdlib-sweep:full`
5. `npm run test:stdlib:governance`

## Approved Task Sequence

Execution tasks created from spike output:

1. `haxe.go-cgk.15` - bootstrap `reflaxe.family.std` skeleton + schema pack
2. `haxe.go-cgk.16` - implement export/import sync tooling for `haxe.go`
3. `haxe.go-cgk.17` - pin family package in `haxe.go` with dual-run checks
4. `haxe.go-cgk.18` - sibling compiler rollout gate plan

These tasks are intentionally blocked by this checklist task (`haxe.go-cgk.14`) and should be unblocked only after explicit approval.

## Rollout Order

1. Phase-1 bootstrap (`cgk.15`)
2. `haxe.go` sync tooling (`cgk.16`)
3. `haxe.go` pinned consumption + dual-run (`cgk.17`)
4. sibling adoption planning/execution (`cgk.18`)

## Cut Lines (Do Not Cross in Phase-1)

1. No backend runtime extraction.
2. No target-native facade extraction (`go.*`, `rust.*`, `ocaml.*`, `elixir.*` remain local).
3. No semantic-contract inference replacing explicit profile/contract selection.
4. No cross-repo rollout before `haxe.go` dual-run stability is demonstrated.

## Approval Record

When approved, record:

1. approver
2. approval date
3. approved family package versioning policy
4. any narrowed scope/cut-line updates

This record can be appended as a short section in this file before unblocking `cgk.15`..`cgk.18`.
