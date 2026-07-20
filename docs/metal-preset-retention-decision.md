# Metal Preset Retention Decision

Bead: `haxe_go-vfp.6.6`

Decision: retain the global `metal` selector as the supported compatibility and
convenience preset it already is. It remains supported without a deprecation
warning. This record does not authorize deprecation, renaming, or removal, and
it does not create a second semantic product.

The selector earns its current place as one short spelling for a coherent
policy bundle. Source APIs still choose semantics: portable Haxe surfaces keep
portable Haxe behavior, while typed `go.*` APIs, typed Go externs, and
`@:goNative` modules explicitly choose Go-native behavior. The compiler may
emit Go-shaped output under either preset when it can preserve the source
contract.

## Decision and alternatives

| Alternative | Verdict | Reason |
| --- | --- | --- |
| Retain `metal` | Chosen | It is a low-complexity shortcut for four independent defaults and preserves existing build, report, example, and private-consumer compatibility. |
| Rename it | Rejected | A new name would require another long-lived alias and migration story without removing policy complexity or improving source semantics. |
| Deprecate it now | Rejected | Explicit axes prove technical replacement, but not equivalent convenience. Public code search cannot observe private or generated build configuration. |
| Remove it now | Rejected | Removal would break a documented input while providing no compiler-architecture benefit; semantic lowering already ignores the profile name. |

This decision concerns the global `reflaxe_go_profile=metal` policy preset and
its define aliases. `@:goMetal` is a separate source-metadata compatibility
alias for `@:goNative`; no new warning or removal decision for that alias is
made here.

## Remaining preset-only behavior

No separate metal backend, AST, runtime, or semantic lowering remains. Direct
`GoProfile.Metal` references are confined to input resolution, preset mapping,
and legacy report projections:

- `src/reflaxe/go/ProfileResolver.hx` parses accepted input;
- `src/reflaxe/go/compiler/GoPolicyPreset.hx` maps it to typed defaults;
- `src/reflaxe/go/GoReflaxeCompiler.hx` emits compatibility labels and report
  projections.

`GoCompiler.hx` does not branch on `GoProfile.Metal` or a metal-compatibility
predicate. New behavior must continue to branch on typed authority,
specialization, fallback, strictness, runtime, planner, or optimizer policy.

The selector still provides three user-visible conveniences:

1. one stable spelling for `explicit` authority, `eager` specialization,
   `error` fallback, and strict raw-boundary checking;
2. stable configuration for existing HXML, CI, templates, and examples;
3. compatibility report and generated-example lane names used by tooling and
   historical comparisons.

Those are policy and compatibility effects, not a second semantic product.

## Observed usage and migration impact

The commit-pinned [usage evidence](reviews/gpt-5.6-pro/metal-preset-usage-evidence-vfp-6.6.json)
records both repository and public-search evidence.

After adding the focused equivalence contract, the repository contains 28
tracked HXML files selecting `metal`: 10 example configurations, 17 snapshot
configurations, and one project template. It also contains five committed
`generated/metal` example trees, 10 `metal*.stdout` example contracts, 15 files
using the legacy `@:goMetal` metadata alias, and 15 files mentioning the legacy
fallback define. Even a mechanical internal migration would therefore create
meaningful configuration, fixture, and review churn.

GitHub public code search found no indexed external HXML use of the selector.
The repository currently has three stars, no forks, and releases without asset
download telemetry. That is weak evidence of a small public migration surface,
not evidence of zero users: private repositories, local HXML, generated command
lines, shell scripts, and unindexed source are invisible. The search result is
therefore explicitly a lower bound and does not authorize removal.

## Explicit-axis replacement

The compatibility preset resolves to this canonical bundle:

```hxml
-D reflaxe_go_profile=portable
-D reflaxe_go_native_authority=explicit
-D reflaxe_go_native_specialization=eager
-D reflaxe_go_native_fallback=error
-D reflaxe_go_strict_policy=on
```

The paired snapshot contracts
`go_native/metal_preset_equivalence` and
`go_native/explicit_policy_equivalence` compile identical typed `go.Chan<Int>`
source. They require identical generated Go, `go.mod`, runtime output, effective
canonical policy fields, and lowering decisions. The reports may retain
different selector, preset, resolution-source, strict-policy-input,
`metalContractHardError`, legacy `metalFallbackViolation*`, and planner-reason
labels because those fields explain how the same effective behavior was
selected and preserve compatibility projections. Explicit-axis replacement is
behaviorally equivalent; it is intentionally not report-byte-identical.

This proves replacement completeness for behavior. It does not make four
defines as ergonomic as one selector, and therefore does not itself justify a
deprecation.

## Sibling-compiler precedent

The comparison uses committed sibling state, not their dirty worktrees:

| Compiler and commit | Current contract | Consequence for Go |
| --- | --- | --- |
| `haxe.rust` at `c1c95fbe7debccac68975ac9b5d75c115894675f` | A true portable/metal contract still governs ownership, nullability, Rust-native APIs, and fail-closed `rust_no_hxrt` eligibility. | Keep a global profile only when it carries irreducible target-wide semantics. Rust currently does; Go does not. |
| `haxe.elixir.codex` at `a4897cd3106f916c26f813388215f69699e742cc` | One pipeline; typed APIs, imports, and local metadata carry Elixir-first intent, with strictness orthogonal. | Strong precedent for Go's API/module-scoped semantic authority. |
| `haxe.ruby` at `a74818e996b68e467621a76cfaae520f553c6960` | One pipeline retains the existing public `portable|ruby_first` contract because removing or renaming it has no current payoff; generic metal/performance policy is rejected. | Existing public inputs should keep earning their cost, but absence of semantic branching alone is not enough reason to churn them. |

Go therefore follows the Elixir-style source boundary internally and the Ruby
compatibility posture publicly. Rust remains the justified exception rather
than a template requiring Go to invent a second semantic product.

## SemVer and migration timeline

No warning or removal release is scheduled. During the current 0.x line,
`metal` remains an accepted, tested compatibility input and new documentation
continues to describe it as a preset rather than a product mode.

The `haxe_go-vfp.6.3` public contract places the selector and its aliases in the
federated boundary. Bead `haxe_go-vfp.6.4` owns the general pre-1.0 and stable
deprecation policy. At 1.0 admission, the project must explicitly decide whether
the still-supported selector enters the stable public contract; if it does, it
remains supported throughout 1.x and removal is major-only.

If later evidence reopens this decision during 0.x, the minimum migration floor
is:

1. approve a separate xhigh decision and implementation bead with current usage
   evidence and exact output/report equivalence;
2. introduce a warning in one minor release while keeping the selector fully
   functional and documenting a copyable replacement;
3. keep it functional through at least the next minor release;
4. remove it no earlier than the following minor release, and only if the final
   `haxe_go-vfp.6.4` policy permits that 0.x break.

Patch releases must not introduce the warning or removal. Renaming would use
the same minimum window and keep `metal` as an alias for its documented
lifetime. Any implementation remains separate from this decision record.

## Reopen and rollback criteria

Reopen only after all of these exist:

- a published and validated package or beta with usable consumer feedback or
  telemetry, not just public-code-search absence;
- the `haxe_go-vfp.6.3` public contract and `haxe_go-vfp.6.4` SemVer policy;
- a replacement that is both behaviorally complete and demonstrably easier to
  teach or automate than the retained shortcut;
- commit-pinned generated-output, report, runtime, example, and migration
  evidence;
- a new independent xhigh review and a separate approved implementation bead.

Roll back any future warning before removal if users report blocked migrations,
if canonical policy or lowering equivalence fails, if a supported tool still
consumes selector-specific report fields, or if the replacement remains
materially less ergonomic. The safe rollback is to keep accepting `metal`; it
does not require restoring a semantic backend because none exists.

## Independent review

A direct, ephemeral, read-only `gpt-5.6-sol` review ran at xhigh against commit
`4c0d8eb48d336e3149c26efa55bb4ddce4f7fbb5`. The Codex CLI 0.144.6 invocation
recorded the exact route; no Oracle alias or older-model fallback was used. See
the [review](reviews/gpt-5.6-pro/review-vfp-6.6-metal-preset.md) and its
[provenance](reviews/gpt-5.6-pro/review-vfp-6.6-metal-preset.provenance.json).

The decision verdict was `RETAIN WITH FOLLOW-UPS`; the selector-deprecation
verdict was `RETAIN`. The reviewer independently confirmed that no profile-name
branch changes semantic lowering, generated behavior, or runtime ownership,
and that no universal IR or second backend is needed.

Findings were adjudicated as follows:

| Finding | Disposition |
| --- | --- |
| F-1, the usage counts were not reproducible from their original base-commit pin | Accepted and fixed. Network-capture context, repository-inventory commit, and inventory tree are now separate, and the contract reproduces counts from the pinned commit. |
| F-2, the source guard checked only literal `GoProfile.Metal` uses | Accepted and fixed. The contract now inventories preset literals, compatibility predicates, `policyPreset` consumers, the report-only planner tag, and forbidden `GoCompiler` uses. |
| F-3, explicit-axis output is behaviorally but not report-byte equivalent | Accepted as an intentional compatibility fact. The exact differing report families are documented above and become required migration evidence if deprecation is ever reconsidered. |

The review could not rerun Go in its read-only sandbox because the Go tool needs
a writable temporary directory. The primary session ran both paired snapshot
and runtime cases successfully before review and reruns them in final validation.

## Validation contract

```bash
python3 test/test_metal_preset_retention_contract.py
python3 test/run-snapshots.py \
  --case go_native/metal_preset_equivalence \
  --case go_native/explicit_policy_equivalence \
  --runtime
npm run test:changed
npm test
```

Because this decision does not change compiler, runtime, staged std, profile
resolution, or generated application behavior, semantic-diff, stdlib sweep,
examples, and performance gates are unchanged rather than newly exercised.
