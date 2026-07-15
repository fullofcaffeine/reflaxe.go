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

This parity program is therefore about semantic closure and coverage promotion,
not about the compatibility-preset lifecycle.

## Architecture Rules

1. Portable contract code must not depend on `go.*` APIs.
2. Native facade code must be explicit (`go.*`) and documented as non-portable.
3. Compiler-owned shims remain only where compiler context is required (for example metadata-dependent lowering).
4. Library-expressible behavior should migrate to ordinary source under the
   canonical `std/go/_std` override root; `.cross.hx` is generated only in the
   staged package.
5. Upstream sync must be provenance-tracked and boundary-checked in CI.
6. `go.*` core authority is singular: `src/go/*` owns `go.Go`/`go.Chan` with target-conditional behavior; `std/go/*` stays focused on package extern facades (`Fmt`, `Time`, `ContextPkg`, `Http`, ...).

Ownership is chosen by the rubric in `docs/ownership-rubric.md`.
This document tracks closure state and blocker families, not the full decision tree.

## Provenance And Boundary Workflow

Governance artifacts:

- `docs/stdlib-provenance-ledger.json`: baseline upstream tag plus per-file
  provenance, migration ownership, exact destination, and compiler-shim audit
  for every tracked std/support source.
- `scripts/ci/upstream-stdlib-boundary-check.js`: prevents tracked upstream
  vendor roots and enforces the approved ordinary-source layout (`std/*.hx`,
  `std/haxe/**`, `std/sys/**`, `std/go/**`, `std/hxrt/**`).
- `scripts/ci/stdlib-provenance-ledger-check.js`: validates the ledger schema
  and exact tracked-file coverage.
- `test/test_stdlib_migration_ledger_contract.py`: locks ownership and target
  paths for the canonical migration and cross-checks the compiler-shim audit.

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

Automated closure workflow (`compile-only -> snapshot -> semantic-diff`, or explicit policy lock) is produced by:

- `python3 test/run-portable-parity-closure.py`
- `test/portable_parity_promotions.json` (deterministic promotion registry consumed by inventory generation)

The summary artifact includes module-level `next_step`, `closure_policy`, and `actionable` guidance for every remaining non-semantic-diff surface.
`actionable: true` means there is still hidden parity work to do.
`actionable: false` means the module is intentionally held at snapshot coverage or explicit exclusion policy.
For every remaining `compile-only` module, the generated inventory and closure summary must also carry:

- `blocker_issue`
- `blocker_family`
- `closure_target`

That metadata is the repo-level proof that no portable-eligible module is still implicit/unknown.

Current non-semantic-diff surfaces are generated from `test/portable_stdlib_inventory.json` and summarized in:

- `test/.test-cache/portable_parity_closure_summary.json`
- `test/.test-cache/portable_parity_closure_summary.md`

That generated summary is the authoritative live list. This document should describe the workflow and closure rules, not duplicate per-module blocker bookkeeping by hand.

Closed root-surface follow-up:

- `haxe.go-14as.21` promoted `Xml` to semantic-diff coverage through `root_xml_contract` and `stdlib/xml_root_dom_basic`.
- `haxe.go-14as.24` closes parsed-CDATA node-type preservation for `Xml.parse()`, so the root `Xml` DOM subset no longer carries that documented caveat.
- `haxe.go-14as.12` closed the generic `haxe.misc` tranche triage by promoting `haxe.Http` from existing semantic-diff evidence and splitting the remaining modules into `haxe.go-14as.25` to `haxe.go-14as.29`.
- `haxe.go-14as.25` promoted direct `haxe.Log`, `haxe.Resource`, and `haxe.SysTools` usage to semantic-diff coverage, then split the remaining direct blockers into `haxe.go-14as.38` (`haxe.Template`) and `haxe.go-14as.39` (`haxe.ValueException`).
- `haxe.go-14as.39` closed the direct `haxe.ValueException` string-payload parity blocker through `haxe_value_exception_contract` and `stdlib/haxe_value_exception_basic`, using the existing hxrt exception carrier instead of keeping the direct constructor hard-failed.
- `haxe.go-14as.38` promoted direct `haxe.Template` constructor/execute usage to semantic-diff coverage through `haxe_template_contract` and `stdlib/haxe_template_basic`, using a staged `std/go/_std/haxe/Template.hx` override instead of forcing the untouched upstream module through unsupported source-owned assumptions.
- `haxe.go-14as.43` promoted direct `haxe.exceptions.PosException`, `haxe.exceptions.ArgumentException`, and `haxe.exceptions.NotImplementedException` to semantic-diff coverage through `haxe_exceptions_direct_contract` and `stdlib/haxe_exceptions_direct`, using staged exception overrides plus the hxrt exception carrier bridge instead of keeping the direct constructor path compile-only.
- `haxe.go-14as.46` promoted direct `haxe.ds.BalancedTree` and `haxe.ds.GenericStack` runtime use to semantic-diff coverage through `haxe_ds_source_owned_collections_contract` and `stdlib/haxe_ds_source_owned_collections`, using staged collection overrides to keep ownership in std code while leaving broader iterator parity to the later collection audit.
- `haxe.go-14as.47` closed the `haxe.ds.WeakMap` stance blocker by promoting the honest upstream platform contract instead of inventing fake weak references. On Go, direct `new haxe.ds.WeakMap()` now compiles and throws `haxe.exceptions.NotImplementedException` at runtime, matching the generic Haxe stdlib behavior. Evidence: `haxe_ds_weakmap_contract` and `stdlib/haxe_ds_weakmap_platform`.
- `haxe.go-14as.14` closed the direct `haxe.http` half of the old combined HTTP/RTTI blocker. `haxe.http.HttpBase` now has staged-stdlib baseline support through `std/go/_std/haxe/http/HttpBase.hx`, with parity evidence in `haxe_http_base_contract` and `stdlib/haxe_http_base_direct`. Direct `haxe.http.HttpJs` and `haxe.http.HttpNodeJs` are now classified honestly as explicit unsupported target-conditional modules on Go through `negative/direct_haxe_httpjs_unsupported` and `negative/direct_haxe_httpnodejs_unsupported`.
- `haxe.go-14as.57` closed the remaining direct `haxe.rtti.*` reflection tranche after the HTTP half was split out. `haxe.rtti.Meta`, `haxe.rtti.Rtti`, `haxe.rtti.CType`, and `haxe.rtti.XmlParser` now have parity evidence through `haxe_rtti_direct_contract` and `stdlib/haxe_rtti_direct`. Ownership stays mixed on purpose: public parser/type logic lives in `std/go/_std/haxe/rtti/*.hx`, while the backend still owns the narrow class-token `__meta__` / `__rtti` contract plus anonymous-record array-field mutation lowering under that staged surface.
- `haxe.go-14as.15` closed the direct `haxe.io` misc tranche. `haxe.io.BufferInput`, `haxe.io.BytesData`, `haxe.io.Encoding`, `haxe.io.Eof`, `haxe.io.Error`, `haxe.io.FPHelper`, `haxe.io.Mime`, `haxe.io.Scheme`, and `haxe.io.StringInput` now have parity evidence through `haxe_io_misc_contract` and `stdlib/haxe_io_misc_direct`. Ownership is explicit instead of hand-waved as one blob: `FPHelper` moved to staged std (`std/go/_std/haxe/io/FPHelper.hx`), `Mime` / `Scheme` stay upstream source-owned abstractions, and the base IO hierarchy still keeps `StringInput`, `BufferInput`, `Encoding`, `Eof`, and `Error` in compiler-owned shims because their type shapes and inherited helper wiring remain representation-sensitive on Go.
- `haxe.go-14as.17` closed the original direct `sys.db` + `sys.io` parity tranche through `sys_db_io_contract` and `stdlib/sys_db_io_direct`. `haxe_go-vfp.8.7.5` later replaced its interim compiler-owned File wrappers: `sys.io.File`, `FileInput`, `FileOutput`, and `FileSeek` are canonical staged std over typed `runtime/hxrt/file.go` capabilities, while `sys.db` remains upstream/public source ownership.
- `haxe.go-14as.30` closed the remaining `haxe.Resource` embedding gap by wiring compiler resources into the backend-owned `haxe.Resource.content` table, with snapshot coverage in `stdlib/haxe_resource_embedded_basic`.
- `haxe.go-14as.13` promoted `haxe.ds.Either` to semantic-diff coverage through `haxe_ds_either_contract` and `stdlib/haxe_ds_either_direct`, then split the remaining collection/exception debt into `haxe.go-14as.43` to `haxe.go-14as.47` so each backend problem is tracked explicitly.
- `haxe.go-14as.44` promoted direct `haxe.ds.HashMap` to semantic-diff coverage through `haxe_ds_hashmap_contract` and `stdlib/haxe_ds_hashmap_direct`, closing the lowercase `hashCode()` parity gap without requiring target-specific `HashCode` aliases.
- `haxe.go-14as.45` promoted direct `haxe.ds.ArraySort` and `haxe.ds.ListSort` to semantic-diff coverage through `haxe_ds_sort_helpers_contract` and `stdlib/haxe_ds_sort_helpers_direct`, using narrow call-site adapters instead of widening compiler-owned generic lowering.
- `haxe.go-14as.27` promoted `haxe.EnumFlags` and `haxe.EnumTools` to semantic-diff coverage via `haxe_enum_helpers_contract` and `stdlib/haxe_enum_helpers_direct`, closing the enum-helper tranche without adding a target-owned std override.
- `haxe.go-14as.28` closed the stack-fallback half of the old stack/main-loop tranche. `haxe.CallStack` and `haxe.NativeStackTrace` stay under explicit target-sensitive snapshot coverage through `stdlib/haxe_stack_loop_target_sensitive`.
- `haxe.go-14as.40` kept native stack capture out of the portable semantic baseline and reduced it to an explicit Go diagnostic capability. `haxe.go-14as.76` implemented that capability behind `-D reflaxe_go_native_stack_trace` with target-sensitive snapshot/runtime coverage. See `docs/spikes/native-stack-capture-contract.md`.
- `haxe.go-14as.29` closed the legacy text tranche. `haxe.Utf8` now lives in staged std through `std/go/_std/haxe/Utf8.hx` with semantic-diff coverage in `haxe_utf8_contract` plus snapshot coverage in `stdlib/haxe_utf8_basic`, while `haxe.Ucs2` stays under explicit target-sensitive snapshot coverage through `stdlib/haxe_ucs2_platform_exclusion`.
- `haxe.go-14as.42` closed the deprecated `haxe.Utf8` optional-size constructor follow-up. `new haxe.Utf8(size)` now compiles through the staged std override, keeps the constructor parameter typed as `Int`, and ignores the deprecated capacity hint because it has no visible runtime semantics.
- `haxe.go-14as.69` promoted direct `haxe.EntryPoint`, `haxe.MainLoop`, and `haxe.Timer` to snapshot/runtime smoke coverage through `stdlib/haxe_main_loop_runtime_direct`. Ownership is mixed: public APIs live in ordinary staged source under `std/go/_std/haxe/*.hx`, while `runtime/hxrt/thread.go` continues to own the scheduler beneath `sys.thread.EventLoop`.

Update sequence when std override files change:

1. Update tracked std/support files under the approved staged layout
   (`std/*.hx`, `std/haxe/**`, `std/sys/**`, `std/go/**`, `std/hxrt/**`).
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
5. live non-semantic-diff surfaces are generated from the inventory instead of hand-maintained in this roadmap.
   - No compile-only portable blocker families remain.
   - Snapshot-only and explicitly unsupported surfaces stay visible through `test/.test-cache/portable_parity_closure_summary.md`.

Recently closed direct portable tranches:

- `haxe.go-14as.16` promoted the `haxe.io` typed-array family (`ArrayBufferView`, `UInt8Array`, `UInt16Array`,
  `UInt32Array`, `Int32Array`, `Float32Array`, `Float64Array`) to semantic-diff coverage. Public typed-array
  behavior now lives in staged overrides under `std/go/_std/haxe/io/*.hx`, while storage still rides on the
  compiler-owned `haxe.io.Bytes` / `ArrayBufferViewImpl` carrier.

Direct `haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer` usage now has limited direct support.
The support level is snapshot/runtime smoke coverage, not semantic-diff coverage, because
event-loop scheduling depends on target runtime timing. The public API is staged std, and
the scheduler is the existing `sys.thread.EventLoop` / `runtime/hxrt/thread.go` runtime contract.

`haxe.go-14as.19` is now fully closed. Direct `sys.thread` primitives
(`Condition`, `Deque`, `IThreadPool`, `Lock`, `Mutex`, `NoEventLoopException`,
`Semaphore`, `ThreadPoolException`, and `Tls`) keep parity evidence in
`semantic_diff/sys_thread_primitives_contract` and
`snapshot/stdlib/sys_thread_primitives_direct`. The second-wave runtime
surfaces (`Thread`, `EventLoop`, `ElasticThreadPool`, `FixedThreadPool`) now
have matching evidence in `semantic_diff/sys_thread_runtime_contract` and
`snapshot/stdlib/sys_thread_runtime_direct`.

The old `haxe.go-cgk.*` planning work is historical context now, not the active execution tracker.

Language hard-fail proof is already locked:

- `haxe.go-14as.56` closed the remaining lowering-guard tranche by making the invariant inventory explicit and release-checked.
- The release contract now treats those guards as invariant-only proof points, not as open supported-language gaps. See `docs/known-gaps.md`, `docs/feature-support-matrix.md`, and `test/test_language_hard_fail_inventory_contract.py`.

## Definition of Done

1. Portable mode supports the complete portable-eligible Haxe stdlib surface for the pinned baseline.
2. Remaining deltas (if any) are explicit, tested, and release-blocking.
3. Docs, CI artifacts, and issue tracker all agree on parity status.
