# Native Policy Presets and Semantic Boundaries

This document is the canonical contract for `portable`, `metal`, and explicit
Go-native source boundaries.

## Decision

`haxe.go` has one compiler pipeline and one default semantic product: portable
Haxe semantics. Go-native semantics enter through explicit source surfaces:

- typed `go.*` APIs,
- typed Go externs and their metadata,
- or a module declared with `@:goNative`.

The public `reflaxe_go_profile=portable|metal` selector remains supported for
compatibility, but it now selects a convenience policy preset. It does not
select a second backend or silently rewrite portable Haxe APIs to different
semantics.

Here, "supported for compatibility" means existing `metal` build input remains
accepted and tested. It is not release-scope admission: the generated
[compatibility matrix](compatibility-support-matrix.md) classifies `metal` as
`compatibility-only`, while the current pre-1.0 beta claim admits only the
named portable operation/member surface.

The product rule is:

> Portable by default, Go-native by explicit source boundary, and Go-shaped
> generated output whenever the compiler can prove the lowering preserves the
> source contract.

The focused [retention decision](metal-preset-retention-decision.md) keeps
`metal` supported without a deprecation warning. It found no semantic branch to
remove, but found continuing shorthand and compatibility value. That decision
does not authorize deprecation, renaming, or removal.

## What chooses semantics

The consumed source API chooses the semantics that must be preserved.

| Source surface | Semantic contract | Expected lowering |
| --- | --- | --- |
| Haxe syntax, staged stdlib, `haxe.*`, `sys.*`, ordinary app APIs | Portable Haxe behavior | Direct Go where proven; narrow `hxrt` support where required |
| A future admitted portable facade | Its versioned portable facade contract | Native Go representation only when its admission proof permits it |
| `go.*` or a typed Go extern | Explicit Go-native API behavior | Typed, Go-shaped output where supported |
| `@:goNative` module | Explicit module-level native authority | Apply native-boundary checks and report the module as a native boundary |
| Raw `__go__` | Controlled framework escape only | Rejected in app/example code and in every `@:goNative` module |

Changing only the preset must not change the behavior of code that stays on
portable surfaces. If it does, that is a regression unless a separately
versioned portable-contract change says otherwise.

## Compatibility presets

The selector names remain `portable` and `metal` because they are public build
inputs. Internally and in report artifacts they resolve to explicit presets.

| Legacy selector | Policy preset | Native authority | Native specialization | Native fallback | Strict raw boundary (`auto`) |
| --- | --- | --- | --- | --- | --- |
| `portable` (default) | `portable_default` | `guarded` | `proven` | `allow` | off |
| `metal` | `metal_compatibility` | `explicit` | `eager` | `error` | on |

These are defaults, not an inseparable mode. Every canonical native policy axis
can override the selected preset.

### Native authority

```bash
-D reflaxe_go_native_authority=guarded|explicit
```

- `guarded` applies the configured native-usage diagnostic to `go.*` usage
  outside an explicit `@:goNative` module.
- `explicit` accepts typed native API usage without that diagnostic gate.

The existing diagnostic controls retain their compatibility names:

```bash
-D reflaxe_go_portable_native_policy=warn|error|off
-D reflaxe_go_portable_native_scan_mode=typed|scanner|hybrid
-D reflaxe_go_portable_native_allow=<csv>
```

They are consulted whenever native authority is `guarded`, even if the
compatibility selector is `metal`.

### Native specialization

```bash
-D reflaxe_go_native_specialization=proven|eager
```

- `proven` attempts typed native representations only through enabled,
  semantics-backed capability paths.
- `eager` attempts every supported typed `go.Chan`, `go.Slice`, `go.Map`, and
  `go.Result` specialization.

This axis controls when the compiler attempts a representation. It does not
grant native API authority and does not change the source contract.

### Native fallback

```bash
-D reflaxe_go_native_fallback=allow|error
```

- `allow` uses the semantics-safe fallback when typed specialization cannot be
  proven.
- `error` rejects user-owned fallback sites at compile time.

Framework-template attempts may still appear as fallback events in reports;
they are evidence about compiler capability, not application contract errors.

The compatibility alias remains accepted:

```bash
-D reflaxe_go_metal_allow_fallback
```

It selects `allow` only when the canonical fallback axis is absent. Combining
it with canonical `allow` is valid and canonical provenance wins. Combining it
with canonical `error` is a configuration error.

### Strictness, planner, optimizer, and runtime

These remain orthogonal:

- raw-boundary strictness: `reflaxe_go_strict_policy=auto|on|off`;
- planner: `reflaxe_go_auto=off|auto|auto_strict`;
- optimizer: `reflaxe_go_opt=portable_fast|none` and capability toggles;
- runtime packaging: full or selective `hxrt` copy;
- diagnostic capabilities such as native stack capture.

`auto_strict` adds native-boundary fallback enforcement; it does not create a
semantic profile. Selective `hxrt` changes packaging, not semantics.

## Resolution precedence

Each native policy is resolved once at compile start and shared by macros,
lowering, runtime planning, and reports.

1. A canonical axis define wins.
2. For fallback only, the compatible
   `reflaxe_go_metal_allow_fallback` alias applies when no canonical value is
   present.
3. The selected preset supplies any remaining default.

Invalid values and contradictory fallback inputs fail compilation. Reports
include both the effective value and its resolution source.

## `@:goNative` module boundaries

Canonical spelling:

```haxe
import go.Chan;
import go.Go;

@:goNative
class WorkerQueue {
  public static function one():Int {
    var channel:Chan<Int> = Go.newChan(1);
    channel.send(1);
    return channel.recv();
  }
}
```

`@:goNative` may be placed on a type or member. Authority is resolved to the
owning Haxe module, because one generated module cannot safely hold two
different boundary policies.

Rules:

1. The module is exempt from guarded unscoped-`go.*` diagnostics.
2. Raw `__go__` remains forbidden in the module under every preset.
3. With `reflaxe_go_auto=auto_strict`, unresolved typed native fallback in the
   module is rejected.
4. The module is recorded in `nativeBoundaryModules` and every lowering entry
   records `inNativeBoundary`.

`@:goMetal` remains a silent compatibility alias with identical behavior. New
code and documentation should use `@:goNative`. Removed `@:haxeMetal` remains a
compile error.

`@:goAllowRaw` is a separate, narrow framework authority marker. It does not
override a native-boundary rejection; a module that is also `@:goNative` still
cannot use raw `__go__`.

## Common configurations

Default portable build:

```bash
-D reflaxe_go_profile=portable
```

Portable build with one explicit native module:

```haxe
@:goNative
class GoAdapter { /* typed go.* or extern APIs */ }
```

Compatibility-equivalent metal policy without selecting `metal`:

```bash
-D reflaxe_go_profile=portable
-D reflaxe_go_native_authority=explicit
-D reflaxe_go_native_specialization=eager
-D reflaxe_go_native_fallback=error
-D reflaxe_go_strict_policy=on
```

Metal compatibility preset with guarded API admission:

```bash
-D reflaxe_go_profile=metal
-D reflaxe_go_native_authority=guarded
-D reflaxe_go_portable_native_policy=error
```

That last combination is intentional: the preset supplies defaults, and the
explicit authority axis overrides one of them.

## Report contract and compatibility aliases

Canonical report fields are:

- `policyPreset`;
- `semanticBoundarySource` (`typed_api_or_module`);
- `nativeAuthorityPolicy` and `nativeAuthorityPolicySource`;
- `nativeSpecializationPolicy` and `nativeSpecializationPolicySource`;
- `nativeFallbackPolicy` and `nativeFallbackPolicySource`;
- `nativeBoundaryModules`;
- `nativeFallbackEvent*`;
- `inNativeBoundary` on lowering and fallback entries.

Contract report schema v8, runtime plan schema v2, and optimizer plan schema v6
add these fields. The historical `contract`, `metalLaneModules`,
`metalFallbackViolation*`, and `inMetalLane` fields remain as compatibility
aliases. They must not be used as the source of new semantic decisions.

## Sibling compiler precedent

The Reflaxe family does not need one identical profile model; each target must
justify its boundary with observable semantics.

| Compiler | Current direction | Relevance to Go |
| --- | --- | --- |
| Reflaxe.Elixir | One pipeline; portable stdlib-first and typed Elixir-first intent comes from APIs/metadata, while strictness is orthogonal | Closest precedent for API-scoped native authority without a second engine |
| haxe.ruby | One pipeline with `portable` and `ruby_first` as semantic guardrails; explicitly rejects a generic metal/performance profile | Shows that output quality and performance do not justify a profile by themselves |
| haxe.rust | Keeps a true portable/metal contract because string representation, nullability, ownership/borrowing, RAII, and no-runtime gates can genuinely differ | Legitimate exception: its global contract currently carries irreducible semantic choices that Go has not demonstrated |

For `haxe.go`, the previously profile-shaped behavior is fully described by
authority, specialization, fallback, and strictness policies. That is why
`metal` is retained as a compatibility preset rather than claimed as a second
semantic product.

## Retention review and future reconsideration gate

This refactor preserves accepted selectors, defaults, and aliases while making
the underlying policy explicit. An xhigh written second pass plus the full
compiler/semantic/example/runtime/performance/security gates is sufficient to
land that compatibility-preserving step.

Bead `haxe_go-vfp.6.6` evaluated retain, rename, deprecate, and remove and chose
retention. Deprecating, renaming, or removing the global selector in the future
would be a new public API decision. It requires all of the following:

1. evidence that explicit axes and `@:goNative` replace every real use case;
2. usage and migration-impact evidence;
3. an explicit SemVer, warning, alias-lifetime, and rollback plan;
4. a commit-pinned genuine independent deep review;
5. a separate approved implementation issue.

The public contract has now landed. The general SemVer policy must also land
before that decision can reopen. Until a separate approved implementation issue
passes the full gate, `metal` remains supported without deprecation warnings.

## Evidence anchors

- `core/native_boundary_guarded_authority`: canonical local authority under the
  portable/guarded preset.
- `core/report_artifacts_native_policy_overrides`: portable preset with
  explicit/eager/error overrides.
- `core/report_artifacts_metal_proven_override`: metal preset with canonical
  proven specialization and no eager lowering attempts.
- `core/report_artifacts_lane_fallback`: legacy fallback alias provenance.
- `core/report_artifacts_lane_fallback_portable_surfaces`: canonical-over-legacy
  precedence.
- `negative/native_authority_guarded_metal`: guarded authority overriding metal.
- `negative/native_fallback_error_portable`: fail-fast fallback overriding
  portable.
- `negative/native_fallback_conflict`: contradictory configuration rejection.
- existing `@:goMetal` fixtures: compatibility-alias coverage.
