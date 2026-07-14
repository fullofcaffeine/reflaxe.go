# Metal Compatibility-Preset Refactor Evidence: `6b570a7f`

This record pins the `thinking:xhigh` second-pass review of the implementation
commit `6b570a7fba95d33882c89eb34544398eeee3cb0a` (`refactor: make metal a
compatibility policy preset`). It supports Bead `haxe_go-vfp.6.2`. It does not
approve deprecating, renaming, or removing the public `metal` selector; that
decision remains isolated in `haxe_go-vfp.6.6`.

## Decision under review

`haxe.go` retains `reflaxe_go_profile=portable|metal` and its existing aliases
as compatibility-stable inputs. The selector now resolves only a bundle of
policy defaults:

| Selector | Preset | Native authority | Specialization | Fallback | Strict raw boundary (`auto`) |
| --- | --- | --- | --- | --- | --- |
| `portable` | `portable_default` | `guarded` | `proven` | `allow` | off |
| `metal` | `metal_compatibility` | `explicit` | `eager` | `error` | on |

Typed APIs, externs, and `@:goNative` modules determine Go-native source
authority under either preset. `@:goMetal` remains a silent compatibility
alias. Portable Haxe semantics remain the default product contract; generated
Go may use native shapes whenever typed proof preserves those semantics.

## Commit-pinned implementation evidence

- `src/reflaxe/go/compiler/GoPolicyPreset.hx` maps the legacy selector to a
  named default bundle.
- `src/reflaxe/go/compiler/GoBuildContextResolver.hx` resolves typed authority,
  specialization, and fallback axes with deterministic provenance.
- `src/reflaxe/go/compiler/GoBuildContext.hx` carries the resolved typed axes;
  profile-shaped accessors are compatibility-only.
- `src/reflaxe/go/analyze/GoNativeBoundaryAnalyzer.hx` recognizes
  `@:goNative` and the legacy `@:goMetal` spelling.
- `src/reflaxe/go/macros/NativeAuthorityGate.hx` and
  `src/reflaxe/go/macros/NativeBoundaryEnforcer.hx` enforce API/module-scoped
  authority independently of the selected preset.
- `src/reflaxe/go/compiler/GoNativeTypeEligibility.hx` owns native
  representation eligibility without a profile dependency.
- `src/reflaxe/go/GoCompiler.hx` branches on typed specialization and fallback
  policies rather than `GoProfile`.
- `src/reflaxe/go/GoReflaxeCompiler.hx` emits canonical policy, provenance,
  boundary, lowering, and fallback report fields. Historical metal/lane fields
  remain documented compatibility projections.
- `docs/native-policy-presets.md` is the normative user contract. The README,
  profile, define, semantics, examples, runtime, ownership, and glossary docs
  link or align with it.

The source audit found no behavior-bearing `GoProfile` branch in lowering or
boundary enforcement. Remaining profile comparisons are selector mapping and
legacy report labels. `GoBuildContext.isMetalContract()` has no behavior call
sites and is explicitly documented as a compatibility predicate.

## Policy precedence reviewed

Each canonical axis define overrides the selected preset. The legacy
`reflaxe_go_metal_allow_fallback` define is then honored only where it does not
contradict an explicit canonical fallback value. Preset defaults fill any axis
that remains unspecified.

The reviewed edge cases are:

- canonical `allow` plus the legacy allow alias is accepted and canonical
  provenance remains visible;
- canonical `error` plus the legacy allow alias fails with a conflict;
- portable plus explicit/eager/error reproduces the metal preset's native
  policy without changing semantic authority;
- metal plus guarded/proven overrides the preset defaults;
- invalid axis values fail before lowering.

These cases are covered by the new report and negative snapshots named in
`docs/native-policy-presets.md`.

## Test-first and validation evidence

The change was developed through failing contracts before implementation:

- canonical report fields and metadata recognition failed before the typed
  policy model existed;
- precedence tests exposed legacy provenance incorrectly winning over an
  explicit canonical `allow` value;
- the example-boundary contract identified four modules that owned typed Go
  APIs without declaring `@:goNative`;
- reverse-direction coverage initially showed the metal preset still attempting
  eager lowering when canonical `proven` was selected;
- package inventory and legacy diagnostic/report snapshots failed until their
  intentional compatibility changes were recorded.

The review closeout reran the broader release-contract suite after the first
commit draft and found a second stale exact-inventory assertion (`229/230`
instead of the canonical `240/241`). That assertion was corrected and folded
into the implementation commit before this SHA was pinned; the complete suite
then passed.

Final local validation for the implementation commit's tree:

- `npm test`: 248 passed, 0 failed;
- `npm run test:semantic-diff`: 129/129 passed;
- `npm run test:examples`: 12/12 passed;
- `npm run test:stdlib-sweep:go-test`: 55/55 passed;
- changed snapshots, example QA, raw-injection hygiene, report-schema,
  release-contract, package-install, and canonical-layout contracts passed;
- `npm run test:perf:go`, `npm run test:perf:hxrt-selective`, and
  `npm run test:perf:apps` passed with warning-only baseline noise and zero hard
  failures;
- Go race, checkptr, vet, and staticcheck lanes passed;
- supply-chain and Gitleaks gates passed;
- the dependency audit correctly failed under the stale host Go 1.25.6 standard
  library, then passed with zero reachable vulnerabilities under the patched
  supported toolchain selected by `GOTOOLCHAIN=go1.25.12`;
- `git diff --check` passed.

The pre-commit hook subsequently formatted staged Haxe sources and reran the
local-path and Gitleaks guards successfully. Formatting did not alter behavior.

## Sibling-target comparison

The comparison uses committed sibling sources, not their uncommitted working
trees:

- haxe.rust `f0d098bec23aaaccff66ddec90f17b0e071e20ea`: its
  `docs/portable-vs-metal-authoring.md`, `docs/metal-profile.md`, and
  `docs/defines-reference.md` keep a real portable/metal contract because
  string representation, nullability, ownership/borrowing, Rust-native APIs,
  and `rust_no_hxrt` can impose irreducible target semantics. Even there, typed
  surfaces and `@:rustMetal` islands communicate local native intent.
- haxe.ruby `3617c0911a94f58ae42a9afe073d526d435cfadd`:
  `docs/profiles.md` keeps one compiler pipeline, treats `portable|ruby_first`
  as semantic/lint guardrails, rejects a generic metal/performance profile, and
  requires profiles to keep earning their public cost through tested behavior.
- Reflaxe.Elixir `fac97172f1009782d56b226a95e9dca50417b58e`:
  `docs/02-user-guide/AUTHORING_STYLES_PORTABLE_VS_ELIXIR_FIRST.md` has one
  pipeline, no application-wide authoring-profile backend switch, typed
  APIs/imports and local metadata for native intent, and orthogonal strictness.

The resulting Go direction is family-consistent without forcing identical
models: Elixir is the closest precedent, Ruby supports retaining a compatibility
selector as a guardrail, and Rust demonstrates when a true global semantic
contract is justified. No equivalent irreducible Go-wide representation or
runtime semantic choice was found in this change.

## Written second-pass findings

1. **One semantic pipeline is preserved.** Profile selection no longer controls
   representation eligibility, native-boundary discovery, or fallback behavior
   directly.
2. **Compatibility is explicit rather than accidental.** Public selectors,
   aliases, legacy metadata, report fields, and historical environment knobs
   remain supported and are labeled as compatibility surfaces.
3. **The native boundary is local and typed.** `@:goNative` and typed `go.*` or
   extern APIs carry source authority under either preset. Raw application
   injection remains rejected.
4. **Overrides are genuinely orthogonal.** Both portable-with-metal-policies and
   metal-with-portable-policies have executable contracts, preventing the preset
   from regaining hidden semantic authority.
5. **Reports are auditable.** Effective values, resolution provenance, boundary
   membership, lowering attempts, and fallback events are canonical fields;
   legacy fields are projections only.
6. **Examples teach the boundary.** Modules that own typed Go APIs now declare
   `@:goNative`, while portable domain examples remain preset-neutral.
7. **The concurrency correction remains valid.** Its behavior is attached to
   typed `go.Go.spawn` and portable `sys.thread` contracts, not to a global
   profile branch.

No closure-blocking defect was found in the compatibility-preserving refactor.
The residual risk is public migration impact: usage prevalence, alias lifetime,
warning policy, and SemVer timing are not established. That risk belongs to
`haxe_go-vfp.6.6` and is why this record does not authorize deprecation.

## Independent review provenance

A genuine independent GPT-5.6 solver review at maximum reasoning was attempted
for this work, but the Codex account usage limit was reached before any review
output was produced. The reported reset is July 19, 2026 at 12:58 PM. The
installed Oracle fallback was rejected because its “5.6 Pro” selection maps to
the older `gpt-5-pro`; using that output would falsify reviewer provenance.

Therefore:

- there is no independent model finding to quote or adjudicate for this commit;
- this file is the explicit written second-pass fallback allowed for landing the
  compatibility-preserving `thinking:xhigh` implementation;
- the prepared independent prompt is
  `docs/reviews/gpt-5.6-pro/review-prompt-6b570a7f.md`;
- `haxe_go-vfp.6.6` requires a genuine independent review before any deprecation
  decision and remains open.

## Verdict

Accept commit `6b570a7fba95d33882c89eb34544398eeee3cb0a` as the implementation of
`haxe_go-vfp.6.2`. Retain `metal` without deprecation warnings as a compatibility
policy preset. Do not approve deprecation, renaming, or removal of the global
selector until `haxe_go-vfp.6.6` satisfies its independent-review and SemVer
gates.
