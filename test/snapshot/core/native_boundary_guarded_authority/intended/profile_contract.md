# Contract Report

- schema version: `8`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- native authority policy: `guarded` (source `reflaxe_go_native_authority`)
- native specialization policy: `proven` (source `policy_preset`)
- native fallback policy: `allow` (source `policy_preset`)
- auto lowering mode: `off`
- strict examples: `no`
- strict user boundary policy: `auto`
- strict user boundaries: `no`
- metal fallback allowed: `no`
- metal contract hard error: `no`
- emit line directives: `no`
- raw native mode: `interp`
- hxrt selective enabled: `no`
- hxrt force full copy: `no`
- hxrt no feature infer: `no`
- native fallback events: `0`
- native fallback boundary events: `0`
- native fallback non-boundary events: `0`
- metal fallback violations: `0`
- metal fallback lane violations: `0`
- metal fallback non-lane violations: `0`
- portable native import scan mode: `typed`
- portable native import hits: `0`
- portable native import typed hits: `0`
- portable native import scanner hits: `0`
- contract diagnostics: `0`
- lowering decisions: `10` (attempts `6`, success `4`, fallback `0`)

## hxrt manual features
- none

## native boundary modules
- `Main`

## metal lane modules
- `Main`

## portable native import hits
- none

## portable native import typed hits
- none

## portable native import scanner hits
- none

## contract diagnostics
- none

## lowering decisions
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `attempted` | `Main:22` | Attempt typed go.Chan method specialization.
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `succeeded` | `Main:22` | Applied typed go.Chan method specialization (element type: int).
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_recv` | `attempted` | `Main:21` | Attempt typed go.Chan method specialization.
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_recv` | `succeeded` | `Main:21` | Applied typed go.Chan method specialization (element type: int).
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_send` | `attempted` | `Main:20` | Attempt typed go.Chan method specialization.
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_method_send` | `succeeded` | `Main:20` | Applied typed go.Chan method specialization (element type: int).
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `Main:19` | Attempt typed go.Go.newChan specialization.
- `Main` (native-boundary) | `go.concurrency.typed` | `go_chan_new` | `succeeded` | `Main:19` | Applied typed go.Go.newChan specialization (element type: int).
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_method___hx_setBuffer` | `attempted` | `go.Go:19` | Attempt typed go.Chan method specialization.
- `go.Go` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `go.Go:17` | Attempt typed go.Chan constructor specialization.

## native fallback event summary by module
- none

## native fallback events
- none

## metal fallback violation summary by module
- none

## metal fallback violations
- none
