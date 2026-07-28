# Runtime Plan Report

- schema version: `4`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- mode: `full_copy`
- selective enabled: `no`
- full copy: `yes`
- inference disabled: `no`
- manifest authority: `typed_usage_plus_surface_plan_runtime_manifest`
- surface plan authority: `go_build_context_plus_typed_registry_decision`
- surface plan decisions: `5`

## manual features
- none

## inferred features
- `array`
- `core`
- `exception`
- `print`
- `string`

## selected features
- `core`
- `array`
- `string`
- `equality`
- `print`
- `exception`
- `json`
- `sys`
- `file_io`
- `filesystem`
- `process`
- `bytes`
- `ssl`
- `thread`
- `enum_value`
- `map_int`
- `map_string`
- `map_object`
- `atomic_int`
- `atomic_object`

## runtime files
- `array.go`
- `atomic_int.go`
- `atomic_object.go`
- `bytes.go`
- `core.go`
- `enum_value.go`
- `equality.go`
- `exception.go`
- `file.go`
- `filesystem.go`
- `hxrt.go`
- `json.go`
- `map_int.go`
- `map_object.go`
- `map_string.go`
- `print.go`
- `process.go`
- `ssl.go`
- `string.go`
- `sys.go`
- `thread.go`

## capability manifest
- `core` -> `core.go, hxrt.go`
  - `baseline` (`compiler_baseline`)
  - `compatibility_contract` (`default_full_copy`)
- `array` -> `array.go`
  - `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `string` -> `string.go`
  - `baseline` (`compiler_baseline`)
  - `compatibility_contract` (`default_full_copy`)
  - `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `equality` -> `equality.go`
  - `compatibility_contract` (`default_full_copy`)
- `print` -> `print.go`
  - `baseline` (`compiler_baseline`)
  - `compatibility_contract` (`default_full_copy`)
- `exception` -> `exception.go`
  - `baseline` (`compiler_baseline`)
  - `compatibility_contract` (`default_full_copy`)
- `json` -> `json.go`
  - `compatibility_contract` (`default_full_copy`)
- `sys` -> `sys.go`
  - `compatibility_contract` (`default_full_copy`)
- `file_io` -> `file.go`
  - `compatibility_contract` (`default_full_copy`)
- `filesystem` -> `filesystem.go`
  - `compatibility_contract` (`default_full_copy`)
- `process` -> `process.go`
  - `compatibility_contract` (`default_full_copy`)
- `bytes` -> `bytes.go`
  - `compatibility_contract` (`default_full_copy`)
- `ssl` -> `ssl.go`
  - `compatibility_contract` (`default_full_copy`)
- `thread` -> `thread.go`
  - `compatibility_contract` (`default_full_copy`)
- `enum_value` -> `enum_value.go`
  - `compatibility_contract` (`default_full_copy`)
- `map_int` -> `map_int.go`
  - `compatibility_contract` (`default_full_copy`)
- `map_string` -> `map_string.go`
  - `compatibility_contract` (`default_full_copy`)
- `map_object` -> `map_object.go`
  - `compatibility_contract` (`default_full_copy`)
- `atomic_int` -> `atomic_int.go`
  - `compatibility_contract` (`default_full_copy`)
- `atomic_object` -> `atomic_object.go`
  - `compatibility_contract` (`default_full_copy`)

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `core` <- `compatibility_contract` (`default_full_copy`)
- `array` <- `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `string` <- `baseline` (`compiler_baseline`)
- `string` <- `compatibility_contract` (`default_full_copy`)
- `string` <- `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `equality` <- `compatibility_contract` (`default_full_copy`)
- `print` <- `baseline` (`compiler_baseline`)
- `print` <- `compatibility_contract` (`default_full_copy`)
- `exception` <- `baseline` (`compiler_baseline`)
- `exception` <- `compatibility_contract` (`default_full_copy`)
- `json` <- `compatibility_contract` (`default_full_copy`)
- `sys` <- `compatibility_contract` (`default_full_copy`)
- `file_io` <- `compatibility_contract` (`default_full_copy`)
- `filesystem` <- `compatibility_contract` (`default_full_copy`)
- `process` <- `compatibility_contract` (`default_full_copy`)
- `bytes` <- `compatibility_contract` (`default_full_copy`)
- `ssl` <- `compatibility_contract` (`default_full_copy`)
- `thread` <- `compatibility_contract` (`default_full_copy`)
- `enum_value` <- `compatibility_contract` (`default_full_copy`)
- `map_int` <- `compatibility_contract` (`default_full_copy`)
- `map_string` <- `compatibility_contract` (`default_full_copy`)
- `map_object` <- `compatibility_contract` (`default_full_copy`)
- `atomic_int` <- `compatibility_contract` (`default_full_copy`)
- `atomic_object` <- `compatibility_contract` (`default_full_copy`)

## portable surface plan consequences

- required imports: `none`
- required runtime features: `array, string`

## portable surface decisions

- `Main` | location `Main:1` | usage `expression` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `rejected:shape_mismatch` | eligibility detail `Observed type shape does not match the contract pattern.` | selection `fallback:registry_rejected` | representation `hxrt_array` | fallback `shape_mismatch: Observed type shape does not match the contract pattern.` | imports `none` | runtime `array` | no-hxrt contract `conditional` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"v","optional":false,"shape":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]}},{"name":"infos","optional":true,"shape":{"kind":"abstract","path":"StdTypes.Null","parameters":[{"kind":"typedef","path":"haxe.PosInfos","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string` | no-hxrt contract `ineligible` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `function_declaration` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
