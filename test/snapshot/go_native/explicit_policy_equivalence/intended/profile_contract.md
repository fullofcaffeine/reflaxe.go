# Contract Report

- schema version: `8`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- native authority policy: `explicit` (source `reflaxe_go_native_authority`)
- native specialization policy: `eager` (source `reflaxe_go_native_specialization`)
- native fallback policy: `error` (source `reflaxe_go_native_fallback`)
- auto lowering mode: `off`
- strict examples: `no`
- strict user boundary policy: `on`
- strict user boundaries: `yes`
- metal fallback allowed: `no`
- metal contract hard error: `no`
- emit line directives: `no`
- raw native mode: `interp`
- hxrt selective enabled: `yes`
- hxrt force full copy: `no`
- hxrt no feature infer: `no`
- native fallback events: `2`
- native fallback boundary events: `0`
- native fallback non-boundary events: `2`
- metal fallback violations: `0`
- metal fallback lane violations: `0`
- metal fallback non-lane violations: `0`
- portable native import scan mode: `typed`
- portable native import hits: `0`
- portable native import typed hits: `0`
- portable native import scanner hits: `0`
- contract diagnostics: `0`
- lowering decisions: `12` (attempts `6`, success `4`, fallback `2`)

## hxrt manual features
- none

## native boundary modules
- none

## metal lane modules
- none

## portable native import hits
- none

## portable native import typed hits
- none

## portable native import scanner hits
- none

## contract diagnostics
- none

## lowering decisions
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `attempted` | `Main:9` | Attempt typed go.Chan method specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `succeeded` | `Main:9` | Applied typed go.Chan method specialization (element type: int).
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_recv` | `attempted` | `Main:8` | Attempt typed go.Chan method specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_recv` | `succeeded` | `Main:8` | Applied typed go.Chan method specialization (element type: int).
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_send` | `attempted` | `Main:7` | Attempt typed go.Chan method specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_send` | `succeeded` | `Main:7` | Applied typed go.Chan method specialization (element type: int).
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `Main:6` | Attempt typed go.Go.newChan specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `succeeded` | `Main:6` | Applied typed go.Go.newChan specialization (element type: int).
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_method___hx_setBuffer` | `attempted` | `go.Go:19` | Attempt typed go.Chan method specialization.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_method_unmorphable` | `fallback` | `go.Go:19` | Could not monomorphize go.Chan method call for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `go.Go:17` | Attempt typed go.Chan constructor specialization.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_new_unmorphable` | `fallback` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.

## native fallback event summary by module
- `go.Go` (non-boundary): `2`

## native fallback events
- `go.Go` (non-boundary) | `go_chan_method_unmorphable` | `go.Go:19` | Could not monomorphize go.Chan method call for native specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.
- `go.Go` (non-boundary) | `go_chan_new_unmorphable` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization: Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.

## metal fallback violation summary by module
- none

## metal fallback violations
- none
