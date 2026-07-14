# Contract Report

- schema version: `8`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- native authority policy: `guarded` (source `policy_preset`)
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
- portable native import scan mode: `scanner`
- portable native import hits: `0`
- portable native import typed hits: `1`
- portable native import scanner hits: `0`
- contract diagnostics: `0`
- lowering decisions: `6` (attempts `4`, success `2`, fallback `0`)

## hxrt manual features
- none

## native boundary modules
- none

## metal lane modules
- none

## portable native import hits
- none

## portable native import typed hits
- `Main`

## portable native import scanner hits
- none

## contract diagnostics
- none

## lowering decisions
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `attempted` | `Main:4` | Attempt typed go.Chan method specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_method_close` | `succeeded` | `Main:4` | Applied typed go.Chan method specialization (element type: int).
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `Main:3` | Attempt typed go.Go.newChan specialization.
- `Main` (non-boundary) | `go.concurrency.typed` | `go_chan_new` | `succeeded` | `Main:3` | Applied typed go.Go.newChan specialization (element type: int).
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
