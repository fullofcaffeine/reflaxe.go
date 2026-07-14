# Independent Review Prompt: Metal as a Compatibility Policy Preset

Review Haxe.Go implementation commit
`6b570a7fba95d33882c89eb34544398eeee3cb0a` at maximum reasoning depth. Treat
the Git commit as the source authority; do not review an uncommitted working
tree.

## Provenance requirement

This prompt is intended for a genuine GPT-5.6-class independent reviewer. Name
the exact model and reasoning setting in the response. If that model is not
actually available, stop and report the mismatch. Do not silently substitute
`gpt-5-pro`, another older model, or a local written review while claiming
GPT-5.6 provenance.

## Scope

The commit retains the public `reflaxe_go_profile=portable|metal` selector but
redefines it as a compatibility policy-preset selector, not a choice between
semantic products or backend engines. Typed APIs, externs, and `@:goNative`
modules carry Go-native source authority. Canonical authority, specialization,
fallback, strictness, optimizer, and runtime axes remain independently
overrideable.

Review two decisions separately:

1. Is the compatibility-preserving refactor safe and coherent as implemented?
2. What additional evidence would be required before considering deprecation,
   renaming, or removal of the global selector?

Do not infer that approval of the first decision approves the second.

## Required evidence

Read at least:

- `docs/native-policy-presets.md`;
- `docs/profiles.md` and `docs/profile-semantics-guide.md`;
- `docs/defines-reference.md`;
- `src/reflaxe/go/compiler/GoBuildContext.hx`;
- `src/reflaxe/go/compiler/GoBuildContextResolver.hx`;
- `src/reflaxe/go/compiler/GoPolicyPreset.hx` and all
  `GoNative*Policy.hx` files;
- `src/reflaxe/go/analyze/GoNativeBoundaryAnalyzer.hx`;
- `src/reflaxe/go/macros/NativeAuthorityGate.hx` and
  `src/reflaxe/go/macros/NativeBoundaryEnforcer.hx`;
- native lowering/fallback branches and report construction in
  `src/reflaxe/go/GoCompiler.hx` and
  `src/reflaxe/go/GoReflaxeCompiler.hx`;
- the report, positive, and negative fixtures listed under “Evidence anchors”
  in `docs/native-policy-presets.md`;
- `test/test_examples_qa_contract.py` and
  `test/test_metal_graduation_contract.py`;
- `docs/reviews/gpt-5.6-pro/evidence-6b570a7f.md` as a local evidence index,
  while independently checking its claims against the commit.

## Questions to adjudicate

1. Does any behavior-bearing code still branch directly on `GoProfile` or the
   metal preset where it should branch on typed policy or API/module authority?
2. Are portable Haxe semantics invariant when only the compatibility preset
   changes, except for the documented policy defaults and diagnostics?
3. Is precedence unambiguous for canonical native axes, the legacy metal
   fallback alias, and preset defaults? Identify contradictory or
   under-specified combinations.
4. Can `@:goNative`, typed `go.*` surfaces, and typed extern metadata express
   every current Go-native semantic boundary under either preset?
5. Are raw-injection, guarded-authority, strictness, eager specialization, and
   fail-fast fallback checks attached to the correct local or orthogonal axis?
6. Are report schema v8, runtime schema v2, and optimizer schema v6 complete
   enough to prove effective policy and provenance? Could a compatibility field
   be mistaken for canonical authority?
7. Do the positive and negative fixtures prove both override directions, invalid
   values, conflicts, legacy aliases, canonical metadata, and boundary-scoped
   enforcement?
8. Do examples and docs teach portable-by-default without presenting portable
   as a slow/basic mode or metal as the only real Go mode?
9. Does the change preserve existing users' accepted selectors, aliases,
   defaults, diagnostics, reports, and generated behavior closely enough for the
   current pre-1.0 compatibility policy?
10. Is the prior concurrency lifecycle fix still correctly API-scoped rather
    than dependent on a global preset?

## Sibling precedent

Compare only committed evidence at these SHAs:

- haxe.rust `f0d098bec23aaaccff66ddec90f17b0e071e20ea`;
- haxe.ruby `3617c0911a94f58ae42a9afe073d526d435cfadd`;
- Reflaxe.Elixir `fac97172f1009782d56b226a95e9dca50417b58e`.

Determine whether Go should follow Elixir's API/module-scoped intent, Ruby's
one-pipeline compatibility guardrails, Rust's genuine portable/metal contract,
or a Go-specific combination. Do not transfer Rust ownership, borrowing,
nullability, or no-runtime constraints unless Go has an equivalent irreducible
semantic need.

## Deprecation gate

For the future selector decision, require concrete answers for:

- remaining behavior unique to the preset after explicit axes are applied;
- observed user/configuration usage and migration impact;
- replacement completeness for mixed portable/native applications;
- SemVer version, warning start, alias lifetime, removal window, and rollback
  criteria;
- diagnostic and report-schema migration;
- whether retaining the preset has lower public cost than deprecating it.

The acceptable outcome may be “retain indefinitely.” Do not assume deprecation
is the goal.

## Output format

Start with one of:

- `ACCEPT CURRENT REFACTOR`;
- `ACCEPT WITH FOLLOW-UPS`;
- `BLOCK CURRENT REFACTOR`.

Then provide:

1. exact reviewer model/provenance;
2. findings ordered by severity, each with commit-relative file/line evidence,
   failure scenario, and smallest root-cause fix;
3. answers to all ten implementation questions;
4. a sibling-precedent judgment;
5. a separate selector-deprecation verdict: `NOT READY`, `READY TO DECIDE`, or
   `RETAIN`, with missing evidence;
6. tests or report contracts needed for every proposed change.

Do not fabricate review output when the requested reviewer is unavailable.
