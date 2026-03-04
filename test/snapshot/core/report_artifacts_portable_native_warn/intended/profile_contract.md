# Contract Report

- schema version: `6`
- contract: `portable`
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
- metal fallback violations: `0`
- metal fallback lane violations: `0`
- metal fallback non-lane violations: `0`
- portable native import hits: `1`
- contract diagnostics: `1`
- lowering decisions: `8` (attempts `4`, success `2`, fallback `2`)

## hxrt manual features
- none

## metal lane modules
- none

## portable native import hits
- `Main`

## contract diagnostics
- `Main` | `portable_native_import` | `warning` | `Main:3` | PortableNativeImportGate: module `Main` uses target-native `go.*` surfaces while `reflaxe_go_profile=portable` is active. Move native usage behind adapters, or use `-D reflaxe_go_portable_native_policy=off|warn|error`.

## lowering decisions
- `Main` (non-lane) | `go.concurrency.typed` | `go_chan_method_close` | `attempted` | `Main:6` | Attempt typed go.Chan method specialization.
- `Main` (non-lane) | `go.concurrency.typed` | `go_chan_method_close` | `succeeded` | `Main:6` | Applied typed go.Chan method specialization (element type: int).
- `Main` (non-lane) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `Main:5` | Attempt typed go.Go.newChan specialization.
- `Main` (non-lane) | `go.concurrency.typed` | `go_chan_new` | `succeeded` | `Main:5` | Applied typed go.Go.newChan specialization (element type: int).
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_method___hx_setBuffer` | `attempted` | `go.Go:19` | Attempt typed go.Chan method specialization.
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_method_unmorphable` | `fallback` | `go.Go:19` | Could not monomorphize go.Chan method call (element type: any).
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `go.Go:17` | Attempt typed go.Chan constructor specialization.
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_new_unmorphable` | `fallback` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization.

## metal fallback violation summary by module
- none

## metal fallback violations
- none
