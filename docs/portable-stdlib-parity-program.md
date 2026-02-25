# Portable Stdlib Parity Program (100% Target)

This document is the active standard-library execution plan for `reflaxe.go`.

## Decision

`reflaxe.go` ships both surfaces, with distinct contracts:

1. `portable` contract surface (canonical, cross-target-oriented)
2. `go.*` native facade surface (explicit target-native opt-in)

These are not two competing stdlibs. They are two lanes with different guarantees.

## Portable Contract Objective

Portable mode targets full parity for the portable-eligible Haxe stdlib surface.

That means every portable-eligible upstream Haxe stdlib module for the pinned baseline (`4.3.7`) must end in one of these states:

1. supported with semantic-diff/runtime contract coverage, or
2. explicitly tracked as a temporary gap with a dated closure plan.

No portable-eligible module should remain implicit/unknown.

Contract artifact:

- `test/portable_allowlist.json` (tiered canonical portable module set, validated in CI)
- `docs/portable-semantics-v1.md` (versioned portable semantics contract for high-risk cross-target behavior)

Excluded from this parity objective:

- target-specific stdlib namespaces such as `cpp.*`, `java.*`, `cs.*`, `hl.*`, `lua.*`, `php.*`, `python.*`, `js.*`, and similar target-bound modules.
- those remain target-native surfaces, outside portable contract guarantees.

## Why This Shape

1. Portable-first is the only way to keep cross-target Haxe promises credible.
2. A typed native facade is the only maintainable way to expose Go-native power without raw `__go__` sprawl.
3. Runtime/shim/slicing controls are orthogonal capabilities; they should not redefine the semantic contract.

## Architecture Rules

1. Portable contract code must not depend on `go.*` APIs.
2. Native facade code must be explicit (`go.*`) and documented as non-portable.
3. Compiler-owned shims remain only where compiler context is required (for example metadata-dependent lowering).
4. Library-expressible behavior should migrate to staged stdlib sources (`.cross.hx` and approved override paths).
5. Upstream sync must be provenance-tracked and boundary-checked in CI.
6. `go.*` core authority is singular: `src/go/*` owns `go.Go`/`go.Chan` with target-conditional behavior; `std/go/*` stays focused on package extern facades (`Fmt`, `Time`, `ContextPkg`, `Http`, ...).

## Provenance And Boundary Workflow

Governance artifacts:

- `docs/stdlib-provenance-ledger.json`: baseline upstream tag + per-file provenance records for tracked std override files.
- `scripts/ci/upstream-stdlib-boundary-check.js`: prevents tracked upstream vendor roots and enforces approved std override roots.
- `scripts/ci/stdlib-provenance-ledger-check.js`: validates ledger schema and ensures ledger coverage exactly matches tracked std override files.

Required commands:

```bash
npm run test:portable-allowlist
npm run test:portable-conformance
npm run test:stdlib:governance
```

Update sequence when std override files change:

1. Update tracked std override files under approved roots (`std/go/**`, `std/_std/**`).
2. Add/update matching entries in `docs/stdlib-provenance-ledger.json`.
3. Run `npm run test:stdlib:governance`.
4. Run `python3 test/run-ci.py` (or `npm run test:ci`) before merging.

## Portable Native-Import Policy

Portable programmers may temporarily import native modules during migration, but the compiler should make that explicit.

Recommended policy:

1. default local mode: warning on native-surface imports in portable contract code;
2. CI/release mode: error on native-surface imports in portable contract code;
3. explicit override: allow only for sanctioned adapter modules.

This keeps migration practical while preserving a strict portability gate for release quality.

## Program Phases (Beads)

Epic: `haxe.go-cgk`

1. `haxe.go-cgk.1` - Canonical stdlib contract/layering spec
2. `haxe.go-cgk.2` - Portable contract import/boundary gate
3. `haxe.go-cgk.3` - Full stdlib inventory + parity ledger
4. `haxe.go-cgk.4` - Upstream provenance ledger + boundary CI
5. `haxe.go-cgk.5` - Staged stdlib source migration (`.cross.hx` / `std/_std`)
6. `haxe.go-cgk.6` - Consolidate duplicated `go.*` authority
7. `haxe.go-cgk.7` - Parity closure harness (all modules)
8. `haxe.go-cgk.8` - Family shared portable stdlib extraction spike
9. `haxe.go-cgk.9` - Rename lane metadata to `@:goMetal` only
10. `haxe.go-cgk.10` - Portable allowlist artifact (`portable_allowlist.json`, tiered contract set)
11. `haxe.go-cgk.11` - Portable semantics spec v1 (family-grade)
12. `haxe.go-cgk.12` - Tier1 portable conformance suite seed (repo-local)
13. `haxe.go-cgk.13` - Module mapping contract doc (Haxe code vs runtime binding vs intrinsic)
14. `haxe.go-cgk.14` - Family extraction execution prep (post-spike checklist)

## Definition of Done

1. Portable mode supports the complete portable-eligible Haxe stdlib surface for the pinned baseline.
2. Remaining deltas (if any) are explicit, tested, and release-blocking.
3. Docs, CI artifacts, and issue tracker all agree on parity status.
