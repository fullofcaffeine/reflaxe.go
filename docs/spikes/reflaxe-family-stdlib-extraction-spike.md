# Spike: `reflaxe.family.std` Extraction Plan (Portable Contract Surfaces)

Status: proposed  
Source program: `haxe.go-cgk.8`  
Prepared from current `haxe.go` portable parity assets (`allowlist`, `semantics v1`, `tier1 conformance`, provenance guards, parity closure harness).

## Objective

Design a shared family package/repo (`reflaxe.family.std`) that centralizes portable-contract governance artifacts while keeping backend-specific runtime/lowering ownership local to each compiler.

Primary consumers:

1. `haxe.go` (first adopter)
2. `haxe.rust`
3. `haxe.elixir.codex`
4. `hxhx`/`reflaxe.ocaml` (phase-gated)

## Non-goals (Phase-1)

1. Not extracting backend runtime implementations (`hxrt`, rust runtime crate, ocaml runtime modules, etc.).
2. Not forcing a unified lowering architecture across targets.
3. Not claiming full semantic-diff parity across all compilers on day one.
4. Not extracting target-native surfaces (`go.*`, `rust.*`, `ocaml.*`, `elixir.*`).

## Proposed Repository Shape

```text
reflaxe.family.std/
  contracts/
    portable-semantics/
      v1.md
  allowlists/
    portable_allowlist.v1.json
  conformance/
    tier1/
      portable_conformance_tier1.v1.json
      semantic_diff_case_index.v1.json
  provenance/
    stdlib-provenance-ledger.schema.json
    upstream-boundary-policy.v1.json
  docs/
    module-mapping-contract.v1.md
    migration-playbook.md
  tools/
    verify_allowlist.py
    verify_parity_closure.py
    verify_provenance.py
  fixtures/
    semantic_diff/
      ...shared cases...
```

## Ownership Model (Cut Line)

Portable shared package owns:

1. portable contract semantics spec (`portable-semantics-v1`)
2. portable allowlist schema + canonical allowlist instance
3. Tier1 conformance mapping schema + canonical mapping instance
4. provenance schema and boundary policy spec
5. shared conformance fixture indexes (and eventually shared fixtures where practical)

Target compiler repos keep ownership of:

1. runtime implementation (`runtime/hxrt/*.go`, rust runtime crate, etc.)
2. compiler lowering/intrinsics
3. target-native facade namespaces and strict-boundary policies
4. backend-specific fixture extensions and perf harnesses

## Versioning and Compatibility Policy

Use semver on `reflaxe.family.std` tags.

- `MAJOR`: breaking contract schema/semantics changes
- `MINOR`: additive modules/cases/rules, no breaking schema change
- `PATCH`: clarifications/fixes with stable schemas

Contract version tuple in consumers:

- family package version (`family_std_version`)
- portable semantics version (`portable-semantics-v1`, later `v2`, ...)
- pinned upstream Haxe baseline version (`4.3.7` initially)

Consumer compilers must pin explicit versions in-repo (not floating `main`).

## CI Federation Model

Each compiler repo runs two classes of checks:

1. Local implementation checks (existing snapshots, runtime tests, target-specific suites).
2. Family-contract checks (imported/shared):
   - allowlist validation
   - conformance mapping validation
   - parity closure summary schema validation
   - provenance schema validation

Planned federation workflow:

1. Family repo publishes release tag.
2. Consumer repo updates pinned family version.
3. Consumer CI runs shared verification tools against local implementation artifacts.
4. Migration PR is blocked unless all shared contract checks pass.

## Migration Sequence (Proposed)

### Phase 0: Snapshot export from `haxe.go` (no external repo switch yet)

1. keep current local artifacts as source of truth;
2. add export packaging script for:
   - `docs/portable-semantics-v1.md`
   - `test/portable_allowlist.json`
   - `test/portable_conformance_tier1.json`
   - `docs/portable-module-mapping-contract.md`
3. verify export payload determinism in CI.

### Phase 1: Bootstrap `reflaxe.family.std`

1. create family repo skeleton (structure above);
2. import Phase-0 payload as `v1` assets;
3. wire schema checks + release pipeline in family repo.

### Phase 2: Consume from `haxe.go`

1. pin family package version in `haxe.go`;
2. add sync/verify tooling in `haxe.go` CI;
3. keep local mirror copies while dual-running comparison checks;
4. remove local canonical ownership only after dual-run stability window.

### Phase 3: Roll into sibling compilers

1. apply same pin/sync/verify flow in `haxe.rust`;
2. repeat for `haxe.elixir.codex` and `hxhx` with target-specific adapters;
3. track documented deviations explicitly until closed.

## Hard Blockers

1. Fixture portability mismatch:
   - some semantic-diff fixtures are backend/runtime-specific.
2. Schema drift risk:
   - local docs/scripts may diverge before family repo pinning is fully enforced.
3. Ownership ambiguity:
   - modules marked `mixed` need per-target mapping clarity before central governance hardens.
4. Tooling heterogeneity:
   - consumer repos use different harness runners and language toolchains.

## Phase-1 Cut Lines (Required to start extraction)

Must be true:

1. `portable-semantics-v1` is stable and versioned.
2. Tier1 allowlist + conformance mapping are deterministic and CI-gated.
3. parity closure summary artifact is generated in full CI.
4. module ownership contract exists for Tier1.
5. provenance policy + boundary checks are present and green.

Deferred to later phases:

1. full fixture sharing across all compilers
2. automatic sync bots for all repos
3. unified release cadence across all compiler repos

## Approval Output

Post-spike execution checklist:

- `docs/spikes/reflaxe-family-stdlib-execution-checklist.md`

When approved, execute follow-up tasks:

1. bootstrap family package skeleton and schemas
2. add export/import synchronization tooling
3. pin and validate family package in `haxe.go`
4. define sibling compiler rollout order and compatibility gates
