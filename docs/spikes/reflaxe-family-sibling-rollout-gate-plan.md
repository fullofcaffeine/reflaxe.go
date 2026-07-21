# `reflaxe.family.std` Sibling Adoption Gate Plan

Owner issue: `haxe.go-cgk.18`; corrected by `haxe_go-1h8p`

Scope: evidence-based, opt-in adoption after the family core has one reproducible identity.

## Goal

Offer a genuinely shared portable-contract core to sibling compilers without
breaking each backend's documented semantics or misrepresenting target-local
artifacts as a common package.

This plan explains:

1. what each sibling already uses;
2. whether shared-core adoption would add value;
3. per-repository prerequisites and semantic differences;
4. mandatory identity and evidence gates for any adoption pull request.

## Current State (2026-07-20)

| Repository | Family directory | Existing governance | Current conclusion |
| --- | --- | --- | --- |
| `haxe.rust` | Present | Local family mirror plus Rust conformance, profile, runtime-plan, and generated-output checks | Comparison partner, not yet a consumer of one shared release |
| `genes` | Absent | Deterministic compatibility reports and semantic-differential evidence | Local governance is already credible; adopt only a useful shared core |
| `haxe.elixir.codex` | Absent | Stdlib inventory/parity guards, upstream `unitstd`, ExUnit runtime evidence, snapshots, and an authoring-profile contract | Does not need a new global profile selector; adoption is optional |
| `hxhx` / `reflaxe.ocaml` | Not reassessed in this comparison | Earlier plan evidence must be refreshed before adoption | No rollout claim yet |

Rust and Go currently label different payloads as
`reflaxe.family.std@0.1.0-bootstrap.1`. Therefore Rust's existing directory is
not evidence that it consumes the same package. It is evidence that both
repositories found local mirror verification useful.

## Smallest Useful Model

```text
one immutable family core
    + one target-qualified adapter overlay
    + that target's compiler/runtime tests
    = reproducible cross-repository contract evidence
```

The **family core** contains only rules and identifiers that have the same
meaning for every adopting compiler. A **target adapter overlay** maps those
rules to a particular backend's modules, fixtures, deviations, and evidence.
The overlay may differ; the core payload for a pinned core version may not.

## Adoption Policy

There is no mandatory repository order. The first adoption pair should be the
two repositories that can prove they consume the exact same core payload while
keeping target differences in separate overlays. Go and Rust are the natural
comparison pair because both already have bootstrap machinery, but neither is
an external-package consumer yet.

Genes, Elixir, and OCaml should adopt only when the shared core removes real
duplicated work or supplies reusable fixtures. Their absence must not block Go
or Rust release readiness.

## Global Prerequisites

Before any repository claims to consume a shared family release:

1. the family core has an explicit file boundary and immutable content digest;
2. at least two repositories verify the same core version and digest;
3. target-specific allowlists, module mappings, fixture bindings, deviations,
   and implementation evidence live in target-qualified overlays;
4. `haxe.go` local mirror checks remain green:
   - `npm run test:family-stdlib-sync`
   - `npm run test:family-stdlib-bootstrap`
5. `haxe.go` portable parity reports are stable and reviewed:
   - `test/.test-cache/portable_parity_closure_summary.json`
   - `test/.test-cache/family_std_dual_run_report.json`
6. the pin identifies the family core version, source, and content digest;
7. no extraction blocker remains in
   `docs/spikes/reflaxe-family-stdlib-execution-checklist.md`.

## Per-Repository Gates

### `haxe.rust` (first comparison candidate)

Prerequisites:

1. Compare every proposed core file with Go and resolve differences explicitly.
2. Keep Rust's product/profile contract and runtime policies local.
3. Keep existing runtime feature inference and reporting operational while
   wiring shared core artifacts.

Known target differences:

1. Authority, specialization, fallback, strictness, and representation choices
   remain target policy; a legacy profile name is not shared semantic authority.
2. `rust_no_hxrt` and no-runtime lanes remain Rust-specific.

Required evidence:

1. Rust and Go pins contain the same core version and content digest.
2. Shared Tier1 case identifiers bind to Rust fixtures through a
   Rust-qualified overlay, and those fixtures pass.
3. Existing Rust contract and boundary tests stay green.
4. Runtime-plan reports remain deterministic and identify both core and overlay
   provenance.

### `genes` (optional adopter)

Prerequisites:

1. Identify shared semantic facts or fixture identifiers that materially reduce
   duplication beyond Genes' existing compatibility report.
2. Preserve the Haxe JavaScript target as Genes' documented runtime semantic
   baseline unless a separately reviewed cross-target contract applies.

Required evidence:

1. Existing semantic-differential, compatibility-report, strict-diagnostic, and
   downstream-contract checks remain green.
2. Any claimed portable subset is explicit and backed by cross-target runtime
   evidence; source appearance or successful compilation alone is insufficient.
3. The family dependency is optional for ordinary Genes development unless it
   supplies an actual build-time contract.

### `haxe.elixir.codex` (optional adopter)

Prerequisites:

1. Preserve its existing portable-stdlib-first and typed-Elixir-first authoring
   contract; do not require a `portable|metal` selector.
2. Keep strict boundary enforcement available for applications and examples,
   and preserve explicit typed BEAM-native boundaries.

Known target differences:

1. BEAM-oriented interop and idiomatic transforms may not map one-to-one to
   portable contract defaults.
2. Upstream `unitstd` and local ExUnit evidence may bind to shared case
   identifiers differently from Go snapshots.

Required evidence:

1. Any adopted shared Tier1 cases pass through Elixir's existing runtime
   evidence lanes.
2. Stdlib parity, API inventory, upstream `unitstd`, snapshots, and
   authoring-profile checks stay green.
3. Shared-core adoption does not replace or weaken typed BEAM-native boundary
   tests.

### `hxhx` / `reflaxe.ocaml` (requires refreshed assessment)

Prerequisites:

1. Reinspect the current repository before relying on the older profile and
   runtime-selection assumptions in this plan.
2. Preserve its documented product contract and native boundaries.
3. Maintain its runtime module-selection flow while introducing any shared core
   assets.

Required evidence:

1. Adopted shared cases pass with the OCaml backend.
2. Existing boundary and fixture-harness tests remain green.
3. Runtime-selection evidence is deterministic and target-qualified.

## Adoption Pull Request Checklist

1. Pin the family core version, source, and content digest.
2. Pin or identify the target-qualified adapter overlay separately.
3. Verify the pinned core identity before comparing target evidence.
4. Wire the applicable portable semantics, conformance, and provenance gates.
5. Preserve the compiler's existing local regression and native-boundary gates.
6. Document intentional target differences in the adapter overlay.
7. Publish deterministic reports for reviewer inspection.

## Rollback Rule

If adoption fails a mandatory gate after merge, revert to local canonical
governance and reopen the rollout task with blocker details. Do not publish a
replacement payload under the same family-core version.
