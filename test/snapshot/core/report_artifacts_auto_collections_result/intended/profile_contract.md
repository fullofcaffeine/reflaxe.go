# Contract Report

- schema version: `7`
- contract: `portable`
- auto lowering mode: `auto`
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
- portable native import scan mode: `typed`
- portable native import hits: `1`
- portable native import typed hits: `1`
- portable native import scanner hits: `0`
- contract diagnostics: `1`
- lowering decisions: `22` (attempts `12`, success `10`, fallback `0`)

## hxrt manual features
- none

## metal lane modules
- none

## portable native import hits
- `Main`

## portable native import typed hits
- `Main`

## portable native import scanner hits
- none

## contract diagnostics
- `Main` | `portable_native_import` | `warning` | `Main:1` | PortableNativeImportGate: module `Main` uses target-native `go.*` surfaces while `reflaxe_go_profile=portable` is active. Move native usage behind adapters, or use `-D reflaxe_go_portable_native_policy=off|warn|error`.

## lowering decisions
- `Main` (non-lane) | `go.collections.typed` | `go_map_method_exists` | `attempted` | `Main:9` | Attempt typed go.Map method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_map_method_exists` | `succeeded` | `Main:9` | Applied typed go.Map method specialization (key: *string, value: int).
- `Main` (non-lane) | `go.collections.typed` | `go_map_method_set` | `attempted` | `Main:8` | Attempt typed go.Map method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_map_method_set` | `succeeded` | `Main:8` | Applied typed go.Map method specialization (key: *string, value: int).
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get` | `attempted` | `Main:11` | Attempt typed go.Slice method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get` | `attempted` | `Main:8` | Attempt typed go.Slice method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get` | `succeeded` | `Main:11` | Applied typed go.Slice method specialization (element type: int).
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get` | `succeeded` | `Main:8` | Applied typed go.Slice method specialization (element type: int).
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get_length` | `attempted` | `Main:14` | Attempt typed go.Slice method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_get_length` | `succeeded` | `Main:14` | Applied typed go.Slice method specialization (element type: int).
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_push` | `attempted` | `Main:4` | Attempt typed go.Slice method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_push` | `attempted` | `Main:5` | Attempt typed go.Slice method specialization.
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_push` | `succeeded` | `Main:4` | Applied typed go.Slice method specialization (element type: int).
- `Main` (non-lane) | `go.collections.typed` | `go_slice_method_push` | `succeeded` | `Main:5` | Applied typed go.Slice method specialization (element type: int).
- `Main` (non-lane) | `go.result.typed` | `go_result_method_isOk` | `attempted` | `Main:12` | Attempt typed go.Result method specialization.
- `Main` (non-lane) | `go.result.typed` | `go_result_method_isOk` | `succeeded` | `Main:12` | Applied typed go.Result method specialization (element type: int).
- `Main` (non-lane) | `go.result.typed` | `go_result_method_unwrap` | `attempted` | `Main:12` | Attempt typed go.Result method specialization.
- `Main` (non-lane) | `go.result.typed` | `go_result_method_unwrap` | `succeeded` | `Main:12` | Applied typed go.Result method specialization (element type: int).
- `Main` (non-lane) | `go.result.typed` | `go_result_static_ok` | `attempted` | `Main:11` | Attempt typed go.Result.ok specialization.
- `Main` (non-lane) | `go.result.typed` | `go_result_static_ok` | `succeeded` | `Main:11` | Applied typed go.Result.ok specialization (element type: int).
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_method___hx_setBuffer` | `attempted` | `go.Go:19` | Attempt typed go.Chan method specialization.
- `go.Go` (non-lane) | `go.concurrency.typed` | `go_chan_new` | `attempted` | `go.Go:17` | Attempt typed go.Chan constructor specialization.

## metal fallback violation summary by module
- none

## metal fallback violations
- none
