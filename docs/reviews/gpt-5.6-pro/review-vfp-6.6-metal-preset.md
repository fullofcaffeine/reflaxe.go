# Metal Preset Retention Independent Review

## Provenance

- Model route: `gpt-5.6-sol (recorded by the direct invocation wrapper)`
- Reasoning effort: `xhigh (recorded by the direct invocation wrapper)`
- Reviewed commit: `4c0d8eb48d336e3149c26efa55bb4ddce4f7fbb5`
- Sibling commits: independently verified from Git objects using the specified full SHAs and `git show`; dirty sibling worktrees were not used.
- Network facts: not independently reproduced. The captured GitHub search and repository metadata were assessed only as recorded evidence.

## Decision verdict

RETAIN WITH FOLLOW-UPS

Retention is technically and product-wise sound. No remaining profile-name branch alters semantic lowering, generated application behavior, or runtime ownership. The selector resolves four policy defaults; downstream behavior consumes the resolved authority, specialization, fallback, and strictness values.

`GoCompiler.hx` branches on typed policy helpers at `src/reflaxe/go/GoCompiler.hx:8282` and `src/reflaxe/go/GoCompiler.hx:8395`, not on the profile. Runtime planning is profile-independent. The additional `policyPreset` use in `GoASTPassRegistry.hx` only annotates planner-report reasons; the selected pass list is identical.

No universal compiler IR or second backend is needed to replace this shortcut, and source tracing does not support treating `metal` as a second semantic product.

## Selector-deprecation verdict

RETAIN

The explicit axes are behaviorally sufficient but are less convenient and not report-label compatible. Public-search absence is a lower bound, not evidence that private HXML, CI, templates, generated configuration, or report consumers do not exist. With negligible architectural savings from removal and no demonstrated replacement UX benefit, warning or deprecation is not justified now.

## Findings

### F-1 — Usage evidence is not reproducible from its stated source commit

- Severity: `minor`
- Confidence: high
- Evidence: `docs/reviews/gpt-5.6-pro/metal-preset-usage-evidence-vfp-6.6.json:4-14` names `93a1fd11d80f40092d23b7cec5e9846b856f46e0` while reporting 28 metal HXML files, including the new equivalence fixture. That commit contains 27 such files and neither paired fixture. The contract at `test/test_metal_preset_retention_contract.py:105-147` validates the current worktree inventory but not consistency with `sourceCommit`.
- Impact: The counts are correct at reviewed `HEAD`, but the document’s “commit-pinned” characterization is not fully auditable from its declared pin.
- Required disposition: Record separate network-capture and repository-inventory commits/tree identities, or repin the inventory to a tree containing the fixtures. Add a consistency assertion for the inventory pin. This does not block retention.

### F-2 — The profile-confinement regression guard is narrower than the architectural rule

- Severity: `minor`
- Confidence: high
- Evidence: `test/test_metal_preset_retention_contract.py:84-103` inventories only literal `GoProfile.Metal` references and checks two strings in `GoCompiler.hx`. It would not detect a future semantic branch through `GoPolicyPreset.MetalCompatibility`, `policyPreset`, or `isMetalContract()` elsewhere. Current indirect uses are legitimate mapping/report uses, including `GoBuildContextResolver.hx:43`, `GoBuildContextResolver.hx:202-213`, and the report-only planner tag at `GoASTPassRegistry.hx:61-97`.
- Impact: Current HEAD satisfies the rule, but the executable guard could permit future profile-shaped behavior outside `GoCompiler.hx`.
- Required disposition: Broaden the contract to inventory preset predicates and constrain non-resolution uses to report-only locations, or add a wider policy-equivalence matrix. No compiler behavior change is required.

### F-3 — Explicit-axis replacement is behavioral, not report-identical

- Severity: `note`
- Confidence: high
- Evidence: Every committed `.go` file and `go.mod` is byte-identical between the paired fixtures, and effective policies plus lowering decisions match. Reports intentionally differ in `contract`, `policyPreset`, policy provenance, `strictUserBoundaryPolicy`, `metalContractHardError`, legacy `metalFallbackViolation*` projections, and optimizer planner-reason source labels.
- Impact: A consumer relying on legacy selector-specific fields cannot replace `metal` with explicit axes and expect byte-identical reports. This strengthens the compatibility rationale for retention.
- Required disposition: Preserve the draft’s distinction between generated-behavior equivalence and selector/provenance compatibility. Any future deprecation must include report-consumer migration evidence.

## Alternatives adjudication

- Retain: Accept. It preserves one-selector ergonomics, existing configuration, fixture lanes, and report compatibility at low implementation cost.
- Rename: Reject now. A clearer new name would still require retaining `metal` as an alias and operating a migration period. It would add surface area without changing source semantics or reducing policy complexity.
- Deprecate: Reject now. Technical decomposition is established, but ergonomic and report compatibility are not replaced, and no meaningful maintenance burden or user benefit has been demonstrated.
- Remove: Reject. It would break a documented input without removing a backend, lowering path, runtime, or semantic branch.

The sibling interpretation is relevant rather than cargo-culted:

- Rust genuinely retains profile-level semantics, including nullability defaults, metal restrictions, native-only capabilities, and fail-closed `rust_no_hxrt`.
- Elixir uses one pipeline with intent carried by APIs, imports, and local metadata; strictness is orthogonal.
- Ruby preserves documented profile inputs and aliases while rejecting a generic performance-only profile.

One caveat is that the pinned Ruby document itself characterizes Go’s `metal` as a meaningful native-performance contract. That sentence is inconsistent with the current Go proposal, but the Go decision does not rely on it; it relies on Ruby’s compatibility and anti-churn reasoning.

## SemVer and migration assessment

The proposed lifecycle is concrete enough while remaining subordinate to `haxe_go-vfp.6.4`:

- No warning or removal is scheduled now.
- During 0.x, a future warning must begin in a minor release.
- The selector remains functional through at least the following minor.
- Removal cannot occur before the subsequent minor and only if the final general policy permits that break.
- Patch releases cannot introduce the warning or removal.
- At 1.0, admission to the stable manifest is an explicit decision. If admitted, support continues throughout 1.x and removal is major-only.

This establishes a minimum safety floor without deciding the project-wide deprecation policy on behalf of `haxe_go-vfp.6.4`.

## Residual risks and reopen triggers

Residual risks are:

- unknown private or generated configuration usage;
- tooling dependence on selector-specific report fields;
- future accidental preset branching not caught by the current narrow guard;
- users continuing to interpret “metal” as the only performant or genuinely Go-shaped mode;
- drift if new policy axes are added without updating the preset and equivalence evidence.

The proposed reopen criteria are adequate: published consumer evidence, the public API manifest, the general SemVer policy, a behaviorally complete and easier replacement, commit-pinned output/report/runtime evidence, independent xhigh review, and a separate implementation issue.

Public-search absence alone must not reopen deprecation or removal. The rollback triggers are also actionable: failed policy/lowering equivalence, report-consumer breakage, blocked migrations, or persistently worse replacement ergonomics all lead safely back to continuing acceptance of `metal`.

## Evidence checked

Checked from current `HEAD` with `git show HEAD:<path>`:

- `AGENTS.md`
- `docs/metal-preset-retention-decision.md`
- `docs/native-policy-presets.md`
- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/reviews/gpt-5.6-pro/metal-preset-usage-evidence-vfp-6.6.json`
- `test/test_metal_preset_retention_contract.py`
- both paired snapshot trees
- `ProfileResolver.hx`, `GoPolicyPreset.hx`, `GoBuildContext*.hx`
- `GoReflaxeCompiler.hx`, `GoCompiler.hx`
- macro enforcement and AST pass registry consumers
- the documented override, fallback, authority, and conflict fixtures
- `docs/release-version-policy.md` and `docs/profile-admission-criteria.md`

Commands included `git rev-parse HEAD`, repository-wide `git grep` for profile/preset predicates, committed-file hash comparison, JSON report diffs, and inventory reproduction at both `HEAD` and the evidence file’s stated source commit.

`python3 test/test_metal_preset_retention_contract.py` ran four of five tests successfully. The remaining test failed because the independent review and provenance artifacts do not yet exist at reviewed `HEAD`, which is expected for this invocation.

The requested Go runtime executions could not run because the read-only sandbox prohibited creation of Go’s temporary build directory. Runtime behavior was therefore not independently re-executed; committed generated output and expected stdout were inspected instead.

Pinned sibling commits checked:

- `haxe.rust` at `c1c95fbe7debccac68975ac9b5d75c115894675f`
- `haxe.elixir.codex` at `a4897cd3106f916c26f813388215f69699e742cc`
- `haxe.ruby` at `a74818e996b68e467621a76cfaae520f553c6960`

No files were edited or created.
