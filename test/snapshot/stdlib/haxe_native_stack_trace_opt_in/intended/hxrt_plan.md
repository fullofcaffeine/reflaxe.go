# Runtime Plan Report

- schema version: `3`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- mode: `selective`
- selective enabled: `yes`
- full copy: `no`
- inference disabled: `no`
- surface plan authority: `go_build_context_plus_typed_registry_decision`
- surface plan decisions: `14`

## manual features
- none

## inferred features
- `array`
- `core`
- `exception`
- `print`
- `stack`
- `string`

## selected features
- `array`
- `core`
- `exception`
- `print`
- `stack`
- `string`

## runtime files
- `array.go`
- `core.go`
- `exception.go`
- `hxrt.go`
- `print.go`
- `stack.go`
- `string.go`

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `array` <- `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `string` <- `baseline` (`compiler_baseline`)
- `string` <- `surface_plan` (`go_build_context_plus_typed_registry_decision`)
- `print` <- `baseline` (`compiler_baseline`)
- `exception` <- `baseline` (`compiler_baseline`)
- `stack` <- `class_usage` (`hxrt.stack.NativeStack`)
- `stack` <- `class_usage` (`hxrt.stack.NativeStackFrame`)

## portable surface plan consequences

- required imports: `none`
- required runtime features: `array, string`

## portable surface decisions

- `Main` | location `Main:3` | usage `expression` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `rejected:shape_mismatch` | eligibility detail `Observed type shape does not match the contract pattern.` | selection `fallback:registry_rejected` | representation `hxrt_array` | fallback `shape_mismatch: Observed type shape does not match the contract pattern.` | imports `none` | runtime `array`
- `Main` | location `Main:3` | usage `expression` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[{"kind":"enum","path":"haxe.CallStack.StackItem","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `fallback:carrier_not_activated` | representation `hxrt_array` | fallback `The native carrier is admitted but not activated; this compiler keeps the semantics-safe fallback until its independent promotion gate lands.` | imports `none` | runtime `array`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"Any","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"class","path":"Array","parameters":[{"kind":"enum","path":"haxe.CallStack.StackItem","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"nativeStackTrace","optional":false,"shape":{"kind":"abstract","path":"Any","parameters":[],"arguments":[],"returnType":null,"fields":[]}},{"name":"skip","optional":true,"shape":{"kind":"abstract","path":"StdTypes.Int","parameters":[],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"class","path":"Array","parameters":[{"kind":"enum","path":"haxe.CallStack.StackItem","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"s","optional":false,"shape":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"stack","optional":false,"shape":{"kind":"abstract","path":"haxe.CallStack","parameters":[],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"value","optional":false,"shape":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `expression` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string`
- `Main` | location `Main:3` | usage `function_declaration` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none`
- `Main` | location `Main:3` | usage `variable_type` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `rejected:shape_mismatch` | eligibility detail `Observed type shape does not match the contract pattern.` | selection `fallback:registry_rejected` | representation `hxrt_array` | fallback `shape_mismatch: Observed type shape does not match the contract pattern.` | imports `none` | runtime `array`
- `Main` | location `Main:3` | usage `variable_type` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[{"kind":"enum","path":"haxe.CallStack.StackItem","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `fallback:carrier_not_activated` | representation `hxrt_array` | fallback `The native carrier is admitted but not activated; this compiler keeps the semantics-safe fallback until its independent promotion gate lands.` | imports `none` | runtime `array`
- `Main` | location `Main:3` | usage `variable_type` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string`
