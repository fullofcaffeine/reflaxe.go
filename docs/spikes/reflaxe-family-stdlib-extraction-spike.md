# Spike: `reflaxe.family.std` Extraction Plan (Portable Contract Surfaces)

Status: bootstrap implemented; standalone extraction blocked on shared-content identity

Source program: `haxe.go-cgk.8`  
Prepared from current `haxe.go` portable parity assets (`allowlist`, `semantics v1`, `tier1 conformance`, provenance guards, parity closure harness).

## Cross-Repository Reality Check (2026-07-20)

The bootstrap proved that a compiler can mirror its local portable-contract
artifacts and verify that mirror in CI. It did **not** yet prove that the mirror
is one shared family package.

Tracked-file inspection found:

| Repository | `family/reflaxe.family.std` | Current semantic evidence |
| --- | --- | --- |
| `haxe.go` | Present and CI-gated | Local portable semantics, allowlist, conformance mapping, ownership mapping, semantic-diff, and stdlib sweeps |
| `haxe.rust` | Present and CI-gated | Its own local versions of the same artifact classes plus Rust-specific profile and runtime evidence |
| `genes` | Absent | Local compatibility reports, feature evidence, and semantic-differential tests against its Haxe/JavaScript baseline |
| `haxe.elixir.codex` | Absent | Local stdlib inventory/parity guards, upstream `unitstd`, Haxe-authored ExUnit tests, snapshots, and authoring-profile contract |

Go and Rust both label their bootstrap as `reflaxe.family.std`
`0.1.0-bootstrap.1`, but their semantics document, allowlist, Tier1 conformance
mapping, module mapping, boundary policy, README, and release scaffold are not
byte-identical. Each verifier compares a repository's local canonical files
with that repository's local mirror. Neither verifier establishes that the two
repositories consumed one immutable payload.

That distinction matters because a package version is an identity promise: two
consumers pinning the same version must receive the same core contents. The
current same-version/different-content state is acceptable only as an explicitly
local bootstrap. It must not be presented as a released cross-repository source
of truth.

### Revised decision

Keep the useful local CI contracts, but separate future shared and target-owned
artifacts:

1. A released **family core** may own genuinely cross-target semantics, schemas,
   stable fixture identifiers, and provenance rules. One core version must have
   one immutable content digest in every consumer.
2. A **target adapter overlay** owns backend-specific eligibility decisions,
   module ownership mappings, fixture bindings, deviations, and implementation
   evidence. Its identity must include the target and its own version or source
   revision.
3. A compiler may keep repo-local governance without adopting the family core.
   Genes and Elixir demonstrate that strong local evidence does not require this
   package shape.
4. External extraction remains blocked until the shared core can be derived
   without relabeling target-specific differences as one package release.

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
2. portable allowlist schema; a canonical cross-target baseline only after every
   claimed consumer supplies evidence for the same baseline
3. Tier1 conformance schema + stable shared case identifiers
4. provenance schema and boundary policy spec
5. shared conformance fixture indexes (and eventually shared fixtures where practical)

Target compiler repos keep ownership of:

1. runtime implementation (`runtime/hxrt/*.go`, rust runtime crate, etc.)
2. compiler lowering/intrinsics
3. target-native facade namespaces and strict-boundary policies
4. target eligibility allowlists and module ownership mappings
5. bindings from shared case identifiers to runnable target fixtures
6. backend-specific deviations, evidence, fixture extensions, and perf harnesses

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

The pin must eventually include a content digest for the immutable family core.
Target overlays need a separate target-qualified identity. A core version must
never be reused for different core payloads, even when the difference appears
reasonable for an individual backend.

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

### Phase 3: Offer the core to sibling compilers

1. compare the existing Go and Rust bootstraps and extract only their genuinely
   identical or explicitly reconciled core;
2. give Go and Rust separate target overlays and prove that the pinned core
   digest is identical in both repositories;
3. offer opt-in adoption to `genes`, `haxe.elixir.codex`, and `hxhx` only where
   the core adds value beyond their local governance;
4. preserve each compiler's existing authoring profiles and target-native
   boundaries rather than imposing one global profile selector;
5. track documented target deviations in their adapter overlays.

## Hard Blockers

1. Fixture portability mismatch:
   - some semantic-diff fixtures are backend/runtime-specific.
2. Schema drift risk:
   - local docs/scripts may diverge before family repo pinning is fully enforced.
3. Ownership ambiguity:
   - modules marked `mixed` need per-target mapping clarity before central governance hardens.
4. Tooling heterogeneity:
   - consumer repos use different harness runners and language toolchains.
5. Package identity ambiguity:
   - the Go and Rust bootstraps currently use the same version for different
     payloads, so no standalone release may be claimed from either copy.

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
