# Contract Report

- schema version: `2`
- contract: `metal`
- strict examples: `no`
- strict user boundaries: `no`
- metal fallback allowed: `yes`
- metal contract hard error: `no`
- emit line directives: `no`
- raw native mode: `interp`
- hxrt selective enabled: `no`
- hxrt force full copy: `no`
- hxrt no feature infer: `no`
- metal fallback violations: `4`

## hxrt manual features
- none

## metal lane modules
- `LaneWorker`

## metal fallback violations
- `LaneWorker` (lane) | `go_result_static_failure_unmorphable` | `LaneWorker:5` | Could not monomorphize go.Result<T>.failure return type for metal specialization.
- `NonLaneWorker` (non-lane) | `go_result_static_failure_unmorphable` | `NonLaneWorker:4` | Could not monomorphize go.Result<T>.failure return type for metal specialization.
- `go.Go` (non-lane) | `go_chan_method_unmorphable` | `go.Go:19` | Could not monomorphize go.Chan method call (element type: any).
- `go.Go` (non-lane) | `go_chan_new_unmorphable` | `go.Go:17` | Could not monomorphize go.Chan element type for constructor specialization.
