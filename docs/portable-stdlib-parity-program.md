# Portable Stdlib Parity Program (100% Target)

This document is the active standard-library execution plan for `reflaxe.go`.

## Decision

`reflaxe.go` ships both surfaces, with distinct contracts:

1. `portable` contract surface (canonical, cross-target-oriented)
2. `go.*` native facade surface (explicit target-native opt-in)

These are not two competing stdlibs. They are two lanes with different guarantees.

In family terms, this program uses explicit `contracts + capabilities + planner + lanes`, not inferred semantics.

## Portable Contract Objective

Portable mode targets full parity for the portable-eligible Haxe stdlib surface.

That means every portable-eligible upstream Haxe stdlib module for the pinned baseline (`4.3.7`) must end in one of these states:

1. supported with semantic-diff/runtime contract coverage, or
2. explicitly tracked as a temporary gap with a dated closure plan.

No portable-eligible module should remain implicit/unknown.

Contract artifact:

- `test/portable_allowlist.json` (tiered canonical portable module set, validated in CI)
- `docs/portable-semantics-v1.md` (versioned portable semantics contract for high-risk cross-target behavior)
- `docs/portable-module-mapping-contract.md` (Tier1 ownership map: Haxe source vs runtime binding vs compiler intrinsic)
- `docs/ownership-rubric.md` (canonical rule for deciding where behavior is allowed to live)

Excluded from this parity objective:

- target-specific stdlib namespaces such as `cpp.*`, `java.*`, `cs.*`, `hl.*`, `lua.*`, `php.*`, `python.*`, `js.*`, and similar target-bound modules.
- those remain target-native surfaces, outside portable contract guarantees.

## Why This Shape

1. Portable-first is the only way to keep cross-target Haxe promises credible.
2. A typed native facade is the only maintainable way to expose Go-native power without raw `__go__` sprawl.
3. Runtime/shim/slicing controls are orthogonal capabilities; they should not redefine the semantic contract.

## Compiler Resolution Spine (current)

- `GoBuildContextResolver.resolve()` centralizes contract, boundary defaults, runtime/planner flags, and report toggles.
- `GoReflaxeCompiler` consumes that context during compile start/end and runtime plan emission.
- Optional report artifacts (`profile_contract.*`, `hxrt_plan.*`, `optimizer_plan.*`) provide deterministic audit output for contract/runtime/planner decisions.

This parity program is therefore about semantic closure and coverage promotion, not about replacing the profile architecture.

## Architecture Rules

1. Portable contract code must not depend on `go.*` APIs.
2. Native facade code must be explicit (`go.*`) and documented as non-portable.
3. Compiler-owned shims remain only where compiler context is required (for example metadata-dependent lowering).
4. Library-expressible behavior should migrate to staged stdlib sources (`.cross.hx` and approved override paths).
5. Upstream sync must be provenance-tracked and boundary-checked in CI.
6. `go.*` core authority is singular: `src/go/*` owns `go.Go`/`go.Chan` with target-conditional behavior; `std/go/*` stays focused on package extern facades (`Fmt`, `Time`, `ContextPkg`, `Http`, ...).

Ownership is chosen by the rubric in `docs/ownership-rubric.md`.
This document tracks closure state and blocker families, not the full decision tree.

## Provenance And Boundary Workflow

Governance artifacts:

- `docs/stdlib-provenance-ledger.json`: baseline upstream tag + per-file provenance records for tracked std override files.
- `scripts/ci/upstream-stdlib-boundary-check.js`: prevents tracked upstream vendor roots and enforces the approved staged std layout (`std/*.hx`, `std/*.cross.hx`, `std/haxe/**`, `std/go/**`, `std/_std/**`).
- `scripts/ci/stdlib-provenance-ledger-check.js`: validates ledger schema and ensures ledger coverage exactly matches tracked std override files.

Required commands:

```bash
npm run test:portable-allowlist
npm run test:portable-conformance
npm run test:portable-parity-closure
npm run test:stdlib-sweep:full
npm run test:family-stdlib-sync
npm run test:family-stdlib-bootstrap
npm run test:stdlib:governance
```

Full CI runs (`python3 test/run-ci.py` with no focused flags) include:

1. curated strict stdlib sweep (`test/upstream_std_modules.txt`),
2. full portable-eligible sweep (`test/upstream_std_modules_full.txt`),
3. portable parity closure summary generation (`test/.test-cache/portable_parity_closure_summary.*`).

Automated promotion workflow (`compile-only -> snapshot -> semantic-diff`) is produced by:

- `python3 test/run-portable-parity-closure.py`
- `test/portable_parity_promotions.json` (deterministic promotion registry consumed by inventory generation)

The summary artifact includes module-level `next_step` guidance for every remaining blocker.
For every remaining `compile-only` module, the generated inventory and closure summary must also carry:

- `blocker_issue`
- `blocker_family`
- `closure_target`

That metadata is the repo-level proof that no portable-eligible module is still implicit/unknown.

Current blocker families are generated from `test/portable_stdlib_inventory.json` and summarized in:

- `test/.test-cache/portable_parity_closure_summary.json`
- `test/.test-cache/portable_parity_closure_summary.md`

That generated summary is the authoritative live blocker list. This document should describe the workflow and closure rules, not duplicate per-module blocker bookkeeping by hand.

Closed root-surface follow-up:

- `haxe.go-14as.21` promoted `Xml` to semantic-diff coverage through `root_xml_contract` and `stdlib/xml_root_dom_basic`.
- `haxe.go-14as.24` closes parsed-CDATA node-type preservation for `Xml.parse()`, so the root `Xml` DOM subset no longer carries that documented caveat.
- `haxe.go-14as.12` closed the generic `haxe.misc` tranche triage by promoting `haxe.Http` from existing semantic-diff evidence and splitting the remaining modules into `haxe.go-14as.25` to `haxe.go-14as.29`.
- `haxe.go-14as.25` promoted direct `haxe.Log`, `haxe.Resource`, and `haxe.SysTools` usage to semantic-diff coverage, then split the remaining direct blockers into `haxe.go-14as.38` (`haxe.Template`) and `haxe.go-14as.39` (`haxe.ValueException`).
- `haxe.go-14as.39` closed the direct `haxe.ValueException` string-payload parity blocker through `haxe_value_exception_contract` and `stdlib/haxe_value_exception_basic`, using the existing hxrt exception carrier instead of keeping the direct constructor hard-failed.
- `haxe.go-14as.38` promoted direct `haxe.Template` constructor/execute usage to semantic-diff coverage through `haxe_template_contract` and `stdlib/haxe_template_basic`, using a staged `std/haxe/Template.cross.hx` override instead of forcing the untouched upstream module through unsupported source-owned assumptions.
- `haxe.go-14as.43` promoted direct `haxe.exceptions.PosException`, `haxe.exceptions.ArgumentException`, and `haxe.exceptions.NotImplementedException` to semantic-diff coverage through `haxe_exceptions_direct_contract` and `stdlib/haxe_exceptions_direct`, using staged exception overrides plus the hxrt exception carrier bridge instead of keeping the direct constructor path compile-only.
- `haxe.go-14as.46` promoted direct `haxe.ds.BalancedTree` and `haxe.ds.GenericStack` runtime use to semantic-diff coverage through `haxe_ds_source_owned_collections_contract` and `stdlib/haxe_ds_source_owned_collections`, using staged collection overrides to keep ownership in std code while leaving broader iterator parity to the later collection audit.
- `haxe.go-14as.30` closed the remaining `haxe.Resource` embedding gap by wiring compiler resources into the backend-owned `haxe.Resource.content` table, with snapshot coverage in `stdlib/haxe_resource_embedded_basic`.
- `haxe.go-14as.13` promoted `haxe.ds.Either` to semantic-diff coverage through `haxe_ds_either_contract` and `stdlib/haxe_ds_either_direct`, then split the remaining collection/exception debt into `haxe.go-14as.43` to `haxe.go-14as.47` so each backend problem is tracked explicitly.
- `haxe.go-14as.44` promoted direct `haxe.ds.HashMap` to semantic-diff coverage through `haxe_ds_hashmap_contract` and `stdlib/haxe_ds_hashmap_direct`, closing the lowercase `hashCode()` parity gap without requiring target-specific `HashCode` aliases.
- `haxe.go-14as.45` promoted direct `haxe.ds.ArraySort` and `haxe.ds.ListSort` to semantic-diff coverage through `haxe_ds_sort_helpers_contract` and `stdlib/haxe_ds_sort_helpers_direct`, using narrow call-site adapters instead of widening compiler-owned generic lowering.
- `haxe.go-14as.27` promoted `haxe.EnumFlags` and `haxe.EnumTools` to semantic-diff coverage via `haxe_enum_helpers_contract` and `stdlib/haxe_enum_helpers_direct`, closing the enum-helper tranche without adding a target-owned std override.
- `haxe.go-14as.28` closed the stack-fallback half of the old stack/main-loop tranche. `haxe.CallStack` and `haxe.NativeStackTrace` stay under explicit target-sensitive snapshot coverage through `stdlib/haxe_stack_loop_target_sensitive`.
- `haxe.go-14as.29` closed the legacy text tranche. `haxe.Utf8` now lives in staged std through `std/haxe/Utf8.cross.hx` with semantic-diff coverage in `haxe_utf8_contract` plus snapshot coverage in `stdlib/haxe_utf8_basic`, while `haxe.Ucs2` stays under explicit target-sensitive snapshot coverage through `stdlib/haxe_ucs2_platform_exclusion`.
- `haxe.go-dt4s` re-opened direct `haxe.EntryPoint`, `haxe.MainLoop`, and `haxe.Timer` support after CI exposed broken generated Go from the old source-owned inclusion path. Those modules now remain explicit compile-only blockers until the event-loop surface is real.

Update sequence when std override files change:

1. Update tracked std override files under the approved staged layout (`std/*.hx`, `std/*.cross.hx`, `std/haxe/**`, `std/sys/**`, `std/go/**`, `std/_std/**`).
2. Add/update matching entries in `docs/stdlib-provenance-ledger.json`.
3. Run `npm run test:stdlib:governance`.
4. Run `python3 test/run-ci.py` (or `npm run test:ci`) before merging.

Portable parity closure artifacts:

- `test/.test-cache/portable_parity_closure_summary.json`
- `test/.test-cache/portable_parity_closure_summary.md`
- `test/.test-cache/family_std_dual_run_report.json`
- `test/.test-cache/family_std_dual_run_report.md`
- `docs/spikes/reflaxe-family-sibling-rollout-gate-plan.md`

## Portable Native-Import Policy

Portable programmers may temporarily import native modules during migration, but the compiler should make that explicit.

Recommended policy:

1. default local mode: warning on native-surface imports in portable contract code;
2. CI/release mode: error on native-surface imports in portable contract code;
3. explicit override: allow only for sanctioned adapter modules.

This keeps migration practical while preserving a strict portability gate for release quality.

## Program Phases (Beads)

Primary epic: `haxe.go-14as`

Current execution order:

1. `haxe.go-14as.55` - architecture lock for `portable + metal` plus ownership boundaries
2. `haxe.go-14as.55.1` - canonical ownership rubric (`compiler` vs `staged std` vs `hxrt` vs `go.*`)
3. `haxe.go-14as.55.2` - split compiler-owned stdlib emitters out of `GoCompiler.hx`
4. `haxe.go-14as.55.4` - make ownership/parity governance release-blocking
5. remaining portable stdlib blocker families in owner-driven order:
   - `haxe.go-14as.47` (`haxe.ds.WeakMap` stance)
   - `haxe.go-dt4s` (direct event-loop surfaces)
   - `haxe.go-14as.14` (`haxe.http` + `haxe.rtti`)
   - `haxe.go-14as.15` (`haxe.io` misc)
   - `haxe.go-14as.16` (`haxe.io` typed arrays)
   - `haxe.go-14as.17` (`sys.db` + `sys.io`)
   - `haxe.go-14as.18` (`sys.net` + `sys.ssl`)
   - `haxe.go-14as.19` (`sys.thread`)
   - `haxe.go-14as.42` (`haxe.Utf8` deprecated optional-size constructor)

The old `haxe.go-cgk.*` planning work is historical context now, not the active execution tracker.

Language hard-fail proof is already locked:

- `haxe.go-14as.56` closed the remaining lowering-guard tranche by making the invariant inventory explicit and release-checked.
- The release contract now treats those guards as invariant-only proof points, not as open supported-language gaps. See `docs/known-gaps.md`, `docs/feature-support-matrix.md`, and `test/test_language_hard_fail_inventory_contract.py`.

## Definition of Done

1. Portable mode supports the complete portable-eligible Haxe stdlib surface for the pinned baseline.
2. Remaining deltas (if any) are explicit, tested, and release-blocking.
3. Docs, CI artifacts, and issue tracker all agree on parity status.
