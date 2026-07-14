# Profile and Preset Admission Criteria

A semantic profile should exist only when it has a real, testable semantic
contract. A policy preset has a lower bar, but must be described honestly as a
bundle of defaults rather than a second product.

## Admission rules

All rules below must be true before adding or retaining a semantic profile:

1. An observable, irreducible semantic difference exists in compiler/runtime
   code.
2. Explicit APIs, metadata, and orthogonal policy axes cannot express the
   difference more accurately.
3. The difference is covered by semantic and generated-output tests.
4. Public docs explain who should choose it and which behavior changes.
5. Compatibility and migration behavior is defined for deprecations/removals.

## Non-admission indicators

Do not add (or keep) a profile when:

- It is naming-only and behavior-equivalent to another profile.
- Its behavior decomposes into authority, strictness, optimization, fallback,
  runtime, or diagnostic policy.
- It has no dedicated tests.
- It increases cognitive load without clear product benefit.
- It is documented as different but implemented identically.

## Policy preset rules

A compatibility preset may remain when it:

- preserves a documented public selector;
- expands deterministically to named, independently overridable axes;
- has no hidden semantic branches;
- reports every effective value and its provenance;
- has tested defaults, overrides, invalid values, and compatibility aliases.

In `haxe.go`, `portable_default` and `metal_compatibility` currently satisfy this
role. The source semantic boundary comes from typed APIs/externs and
`@:goNative`, not the preset.

## Maintenance rules

- Keep profile set minimal by default.
- Do not remove or deprecate a public selector without usage evidence, an
  explicit SemVer plan, and independent review.
- Experimental profiles must be labeled experimental and include strict boundary policy if they expose low-level interop.

See [Native policy presets and semantic boundaries](native-policy-presets.md).
