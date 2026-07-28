# Runtime Plan Report

- schema version: `3`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- mode: `full_copy`
- selective enabled: `no`
- full copy: `yes`
- inference disabled: `no`
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
- none

## runtime files
- full copy (`runtime/hxrt/**`, excluding footprint-explicit diagnostic/capability files unless their typed use or define is enabled)

## selection reasons
- none

## portable surface plan consequences

- required imports: `none`
- required runtime features: `array, string`

## portable surface decisions

- `Main` | location `Main:1` | usage `expression` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `rejected:shape_mismatch` | eligibility detail `Observed type shape does not match the contract pattern.` | selection `fallback:registry_rejected` | representation `hxrt_array` | fallback `shape_mismatch: Observed type shape does not match the contract pattern.` | imports `none` | runtime `array`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"v","optional":false,"shape":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]}},{"name":"infos","optional":true,"shape":{"kind":"abstract","path":"StdTypes.Null","parameters":[{"kind":"typedef","path":"haxe.PosInfos","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:1` | usage `expression` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string`
- `Main` | location `Main:1` | usage `function_declaration` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
