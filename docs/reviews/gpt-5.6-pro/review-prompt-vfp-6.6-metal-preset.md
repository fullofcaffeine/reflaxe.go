# Independent Review Prompt: Global Metal Preset Retention

You are the independent xhigh reviewer for `haxe.go` Bead
`haxe_go-vfp.6.6`. Review the current `HEAD` commit in this repository. Do not
edit files, create commits, or merely restate the proposal. Use read-only source
tracing and tests where useful, and report the exact 40-character reviewed
commit from `git rev-parse HEAD`.

The requested route is `gpt-5.6-sol` with xhigh reasoning. Begin your response
with the actual model route you were invoked as. If you cannot establish that
you are that exact route, stop and say the provenance is insufficient; do not
claim GPT-5.6 on behalf of an alias or older fallback.

## Decision under review

The draft in `docs/metal-preset-retention-decision.md` chooses to retain the
public `reflaxe_go_profile=metal` selector without a deprecation warning. It
defines `metal` as a convenience and compatibility policy preset, not a second
semantic product. It rejects rename, deprecate-now, and remove-now. No compiler
behavior change or deprecation implementation is proposed.

The proposed rationale is:

1. Direct `GoProfile.Metal` references are confined to input resolution, policy
   mapping, and compatibility reports. Semantic lowering does not branch on the
   profile name.
2. Portable plus explicit authority/eager specialization/error fallback/strict
   checking reproduces the effective metal policy and generated Go.
3. The one-selector spelling retains ergonomic and compatibility value even
   though its behavior is decomposed.
4. Public code-search absence is only a lower bound and cannot justify breaking
   private or generated configurations.
5. Sibling targets support target-specific decisions: Rust retains a real
   semantic profile, Elixir carries native intent at source/API boundaries, and
   Ruby keeps existing public profile compatibility where removal churn has no
   payoff.

## Required repository evidence

Read and independently check at least:

- `AGENTS.md`
- `docs/metal-preset-retention-decision.md`
- `docs/native-policy-presets.md`
- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/reviews/gpt-5.6-pro/metal-preset-usage-evidence-vfp-6.6.json`
- `test/test_metal_preset_retention_contract.py`
- `test/snapshot/go_native/metal_preset_equivalence/`
- `test/snapshot/go_native/explicit_policy_equivalence/`
- `src/reflaxe/go/ProfileResolver.hx`
- `src/reflaxe/go/compiler/GoPolicyPreset.hx`
- `src/reflaxe/go/GoReflaxeCompiler.hx`
- `src/reflaxe/go/GoCompiler.hx`

Run or reproduce the focused contract if useful. Distinguish generated-Go
equivalence from intentionally different selector, preset, and provenance
labels in reports.

For sibling precedent, inspect committed state if the adjacent repositories are
readable. Do not use their dirty worktrees as evidence:

- `../haxe.rust` commit
  `c1c95fbe7debccac68975ac9b5d75c115894675f`, especially `README.md`,
  `docs/architecture-capability.md`, and `docs/defines-reference.md`;
- `../haxe.elixir.codex` commit
  `a4897cd3106f916c26f813388215f69699e742cc`, especially
  `docs/02-user-guide/AUTHORING_STYLES_PORTABLE_VS_ELIXIR_FIRST.md`;
- `../haxe.ruby` commit
  `a74818e996b68e467621a76cfaae520f553c6960`, especially
  `docs/profiles.md` and `docs/public-contract.md`.

The usage-evidence JSON contains time-sensitive GitHub search and repository
metadata captured separately. You may assess whether its interpretation is
sound, but do not pretend to have independently reproduced network facts unless
you actually did. In particular, decide whether public-search absence can
support retention, deprecation, or removal.

## Questions to adjudicate

1. Is there any remaining profile-name branch that can alter semantic lowering,
   generated application behavior, or runtime ownership?
2. Does the paired fixture prove the proposed explicit-axis replacement, and
   are any meaningful policy or report dimensions omitted?
3. Does retaining a shorthand impose enough public or implementation cost to
   outweigh compatibility and ergonomics?
4. Would renaming improve the product, or merely add another alias and migration
   period?
5. Is current observed usage evidence sufficient to deprecate or remove a
   documented input?
6. Is the sibling-target interpretation accurate and relevant rather than
   cargo-culted?
7. Are the proposed 0.x minimum window, stable major-only rule, reopen criteria,
   and rollback triggers concrete enough without preempting the general SemVer
   Bead `haxe_go-vfp.6.4`?
8. Should this decision retain, warn, rename, deprecate, or remove now?

## Required response format

Produce a self-contained Markdown review with these exact sections:

1. `# Metal Preset Retention Independent Review`
2. `## Provenance` — actual model route, reasoning effort, reviewed commit, and
   whether sibling commits and network facts were independently verified.
3. `## Decision verdict` — exactly one of `ACCEPT RETENTION`,
   `RETAIN WITH FOLLOW-UPS`, or `BLOCK RETENTION`.
4. `## Selector-deprecation verdict` — exactly one of `RETAIN`,
   `READY TO DEPRECATE`, or `INSUFFICIENT EVIDENCE`, followed by a concise
   rationale.
5. `## Findings` — each finding gets an ID, severity (`blocker`, `major`,
   `minor`, or `note`), confidence, repository-relative evidence, impact, and
   required disposition. Say `No findings` if there are none.
6. `## Alternatives adjudication` — separately judge retain, rename,
   deprecate, and remove.
7. `## SemVer and migration assessment`
8. `## Residual risks and reopen triggers`
9. `## Evidence checked` — commands/files actually checked and any limitation.

Do not require a universal compiler IR merely to replace a policy shortcut.
Do not accept a second semantic product unless source tracing proves one exists.
Conversely, block the proposal if the evidence hides real preset-only behavior,
if the replacement proof is incomplete, or if the migration posture would make
the public contract misleading.
