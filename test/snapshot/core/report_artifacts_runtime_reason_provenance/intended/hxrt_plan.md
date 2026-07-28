# Runtime Plan Report

- schema version: `4`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- mode: `selective`
- selective enabled: `yes`
- full copy: `no`
- inference disabled: `no`
- manifest authority: `typed_usage_plus_surface_plan_runtime_manifest`
- surface plan authority: `go_build_context_plus_typed_registry_decision`
- surface plan decisions: `4`

## manual features
- `process`

## inferred features
- `array`
- `core`
- `exception`
- `json`
- `print`
- `string`

## selected features
- `core`
- `array`
- `string`
- `print`
- `exception`
- `json`
- `process`

## runtime files
- `array.go`
- `core.go`
- `exception.go`
- `hxrt.go`
- `json.go`
- `print.go`
- `process.go`
- `string.go`

## capability manifest
- `core` -> `core.go, hxrt.go`
  - `baseline` (`compiler_baseline`)
- `array` -> `array.go`
  - `dependency_edge` (`json->array`)
- `string` -> `string.go`
  - `baseline` (`compiler_baseline`)
  - `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `print` -> `print.go`
  - `baseline` (`compiler_baseline`)
- `exception` -> `exception.go`
  - `baseline` (`compiler_baseline`)
- `json` -> `json.go`
  - `class_usage` (`haxe.Json`)
- `process` -> `process.go`
  - `manual_define` (`reflaxe_go_hxrt_features`)

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `array` <- `dependency_edge` (`json->array`)
- `string` <- `baseline` (`compiler_baseline`)
- `string` <- `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `print` <- `baseline` (`compiler_baseline`)
- `exception` <- `baseline` (`compiler_baseline`)
- `json` <- `class_usage` (`haxe.Json`)
- `process` <- `manual_define` (`reflaxe_go_hxrt_features`)

## portable surface plan consequences

- required imports: `none`
- required runtime features: `string`

## portable surface decisions

- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"source","optional":false,"shape":{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string` | no-hxrt contract `ineligible` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `function_declaration` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
