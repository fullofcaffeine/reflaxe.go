# `portable` and `metal` Policy Presets

Use `portable` by default. Use explicit Go-native APIs or `@:goNative` when a
module intentionally owns Go semantics.

```bash
-D reflaxe_go_profile=portable|metal
```

The selector is retained for compatibility. It chooses defaults for independent
policies; it does not choose a second compiler backend or a second semantic
product.

Compatibility support is narrower than release admission. The current
[compatibility matrix](compatibility-support-matrix.md) retains `metal` as a
`compatibility-only` input and admits only its named portable operation/member
surface for the pre-1.0 beta claim.

The canonical detailed contract is [Native policy presets and semantic
boundaries](native-policy-presets.md).

## Terms

- [policy preset](glossary.md#policy-preset): a compatibility selector that
  supplies defaults for independent compiler policies.
- [Go-native](glossary.md#go-native): source that intentionally adopts Go API
  or runtime semantics.
- [native boundary](glossary.md#native-boundary): a module-level declaration of
  Go-native authority, normally written as `@:goNative`.

The product rule is: portable Haxe semantics are the default product path;
typed `go.*`/extern APIs and `@:goNative` modules are explicit Go-native source
boundaries. The `metal` selector is defined as a convenience policy preset
rather than a second semantic product. Portable by default, Go-native by
explicit source boundary.

## Short model

Semantics come from source surfaces:

- ordinary Haxe, staged stdlib, `haxe.*`, and `sys.*` keep portable Haxe
  semantics;
- typed `go.*`, typed Go externs, and `@:goNative` declare Go-native intent;
- compiler optimization may use Go-shaped representations whenever it proves
  the selected source semantics are preserved.

The build-wide selector only supplies policy defaults:

| Selector | Internal preset | Authority | Specialization | Fallback | Strict raw boundary (`auto`) |
| --- | --- | --- | --- | --- | --- |
| `portable` (default) | `portable_default` | guarded | proven | allow | off |
| `metal` | `metal_compatibility` | explicit | eager | error | on |

`metal` remains supported without a deprecation warning. Its possible future
deprecation is a separate SemVer and independent-review decision.

## Independent policy axes

```bash
-D reflaxe_go_native_authority=guarded|explicit
-D reflaxe_go_native_specialization=proven|eager
-D reflaxe_go_native_fallback=allow|error
-D reflaxe_go_strict_policy=auto|on|off
```

Canonical defines override preset defaults. The legacy
`reflaxe_go_metal_allow_fallback` alias remains accepted when the canonical
fallback axis is absent. Contradictory `allow`/`error` inputs fail compilation.

Planner, optimizer, runtime packaging, and diagnostic controls remain
orthogonal:

- `reflaxe_go_auto=off|auto|auto_strict`;
- `reflaxe_go_opt=portable_fast|none`;
- selective `hxrt` feature defines;
- native stack capture and other target-sensitive capabilities.

## Explicit native modules

Use canonical `@:goNative` to declare an owning module as a Go-native boundary:

```haxe
@:goNative
class Worker {
  public static function run():Void {
    var channel:go.Chan<Int> = go.Go.newChan(1);
    channel.send(1);
  }
}
```

The boundary works under either preset. It is exempt from guarded `go.*`
diagnostics, rejects raw `__go__`, participates in `auto_strict` fallback checks,
and is visible in reports.

`@:goMetal` is a silent compatibility alias. New source and docs should use
`@:goNative`. `@:haxeMetal` is removed and remains an error.

## Raw `__go__` policy

Raw injection is not normal application interop.

- `reflaxe_go_strict` or `reflaxe_go_strict_policy=on` rejects app-side raw
  injection.
- The `metal_compatibility` preset turns strictness on when strict policy is
  `auto`.
- `@:goNative` always rejects raw injection, regardless of preset.
- `@:goAllowRaw` is reserved for narrow framework-owned std/runtime abstraction
  modules and cannot bypass a `@:goNative` boundary.

Prefer typed extern metadata (`@:go.import`, `@:go.name`, `@:go.receiver`) and
framework facades. Raw snippets do not infer Go package imports.

## Guarded native usage

When authority is `guarded`, these compatibility-named controls govern native
usage outside `@:goNative` modules:

```bash
-D reflaxe_go_portable_native_policy=warn|error|off
-D reflaxe_go_portable_native_scan_mode=typed|scanner|hybrid
-D reflaxe_go_portable_native_allow=<csv>
```

For CI, `native_authority=guarded` plus
`reflaxe_go_portable_native_policy=error` is the clearest way to prevent
accidental Go coupling.

## Typed fallback

When typed specialization cannot use a concrete safe Go representation:

- `native_fallback=allow` uses the semantics-safe representation and records a
  fallback event;
- `native_fallback=error` rejects user-owned fallback sites.

`auto_strict` additionally rejects fallback inside `@:goNative` modules.

Do not select `metal` merely to request performance. Keep source portable when
its semantics are portable, measure the output, and improve proven lowering.
Select typed Go APIs or an explicit native module only when the source contract
itself is Go-specific.

## Reports

```bash
-D reflaxe_go_contract_report
-D reflaxe_go_runtime_plan_report
-D reflaxe_go_optimizer_plan_report
```

Reports expose the preset, every native policy and provenance, native-boundary
modules, lowering decisions, fallback events, runtime selection, and optimizer
selection. Historical metal/lane fields remain compatibility aliases; use the
canonical `native*` fields for new automation.

## Recommended configurations

Normal portable application:

```bash
-D reflaxe_go_profile=portable
```

Portable application with explicit native adapters:

```bash
-D reflaxe_go_profile=portable
-D reflaxe_go_native_authority=guarded
-D reflaxe_go_portable_native_policy=error
```

Mark each approved adapter `@:goNative`.

Legacy Go-first build with preserved defaults:

```bash
-D reflaxe_go_profile=metal
```

Equivalent explicit policy bundle:

```bash
-D reflaxe_go_profile=portable
-D reflaxe_go_native_authority=explicit
-D reflaxe_go_native_specialization=eager
-D reflaxe_go_native_fallback=error
-D reflaxe_go_strict_policy=on
```

## Validation

```bash
npm test
npm run test:semantic-diff
npm run test:examples
```

## Related docs

- [Native policy presets and semantic boundaries](native-policy-presets.md)
- [Profile semantics and migration guide](profile-semantics-guide.md)
- [Defines reference](defines-reference.md)
- [Portable canonical contract](portable-canonical-contract.md)
- [Go concurrency and interop](go-concurrency-interop-guide.md)
- [Glossary](glossary.md)
