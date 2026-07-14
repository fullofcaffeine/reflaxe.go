# Contract Report

- schema version: `8`
- contract: `metal`
- policy preset: `metal_compatibility`
- semantic boundary source: `typed_api_or_module`
- native authority policy: `explicit` (source `policy_preset`)
- native specialization policy: `eager` (source `policy_preset`)
- native fallback policy: `allow` (source `reflaxe_go_native_fallback`)
- auto lowering mode: `off`
- strict examples: `no`
- strict user boundary policy: `auto`
- strict user boundaries: `yes`
- metal fallback allowed: `yes`
- metal contract hard error: `no`
- emit line directives: `no`
- raw native mode: `interp`
- hxrt selective enabled: `no`
- hxrt force full copy: `no`
- hxrt no feature infer: `no`
- native fallback events: `4`
- native fallback boundary events: `1`
- native fallback non-boundary events: `3`
- metal fallback violations: `4`
- metal fallback lane violations: `1`
- metal fallback non-lane violations: `3`
- portable native import scan mode: `typed`
- portable native import hits: `0`
- portable native import typed hits: `0`
- portable native import scanner hits: `0`
- contract diagnostics: `0`
- lowering decisions: `8` (attempts `4`, success `0`, fallback `4`)

## hxrt manual features
- none

## native boundary modules
- `LaneWorker`

## metal lane modules
- `LaneWorker`

## portable native import hits
- none

## portable native import typed hits
- none

## portable native import scanner hits
- none

## contract diagnostics
- none

## lowering decisions
- `LaneWorker` (native-boundary) | `go.result.typed` | `go_result_static_failure` | `attempted` | `LaneWorker:4` | Attempt typed go.Result.failure specialization.
- `LaneWorker` (native-boundary) | `go.result.typed` | `go_result_static_failure_unmorphable` | `fallback` | `LaneWorker:4` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `NonLaneWorker` (non-boundary) | `go.result.typed` | `go_result_static_failure` | `attempted` | `NonLaneWorker:3` | Attempt typed go.Result.failure specialization.
- `NonLaneWorker` (non-boundary) | `go.result.typed` | `go_result_static_failure_unmorphable` | `fallback` | `NonLaneWorker:3` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_method___hx_setBuffer` | `attempted` | `go.Go:19` | Attempt typed go.Chan method specialization.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_method_unmorphable` | `fallback` | `go.Go:19` | Could not monomorphize go.Chan method call for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `go.Go:17` | Attempt typed go.Chan constructor specialization.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_new_unmorphable` | `fallback` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.

## native fallback event summary by module
- `LaneWorker` (native-boundary): `1`
- `NonLaneWorker` (non-boundary): `1`
- `go.Go` (non-boundary): `2`

## native fallback events
- `LaneWorker` (native-boundary) | `go_result_static_failure_unmorphable` | `LaneWorker:4` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `NonLaneWorker` (non-boundary) | `go_result_static_failure_unmorphable` | `NonLaneWorker:3` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go_chan_method_unmorphable` | `go.Go:19` | Could not monomorphize go.Chan method call for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go_chan_new_unmorphable` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.

## metal fallback violation summary by module
- `LaneWorker` (lane): `1`
- `NonLaneWorker` (non-lane): `1`
- `go.Go` (non-lane): `2`

## metal fallback violations
- `LaneWorker` (lane) | `go_result_static_failure_unmorphable` | `LaneWorker:4` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `NonLaneWorker` (non-lane) | `go_result_static_failure_unmorphable` | `NonLaneWorker:3` | Could not monomorphize go.Result<T>.failure return type for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-lane) | `go_chan_method_unmorphable` | `go.Go:19` | Could not monomorphize go.Chan method call for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-lane) | `go_chan_new_unmorphable` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
