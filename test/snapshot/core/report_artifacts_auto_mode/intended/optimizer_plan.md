# Optimizer Plan Report

- schema version: `7`
- contract: `portable`
- policy preset: `portable_default`
- native specialization policy: `proven` (source `policy_preset`)
- auto lowering mode: `auto`
- optimization preset: `portable_fast`
- portable string fastpath enabled: `yes`
- portable concurrency fastpath enabled: `yes`
- string instance typed lowerings: `0`
- string instance legacy lowerings: `0`
- string length field typed lowerings: `0`
- string length field legacy lowerings: `0`
- portable concurrency typed fastpath hits: `0`
- portable concurrency typed fastpath fallbacks: `0`
- go collections typed lowerings: `0`
- go collections typed fallbacks: `0`
- go result typed lowerings: `0`
- go result typed fallbacks: `0`
- lowering fallback boundary count: `0`
- lowering fallback non-boundary count: `0`
- lowering fallback lane count: `0`
- lowering fallback non-lane count: `0`
- surface plan authority: `go_build_context_plus_typed_registry_decision`
- surface plan decisions: `5`
- go ast pass selection source: `planner`

## auto lowering capabilities
- none

## go ast passes
- `normalize_names`
- `rewrite_string_ops`
- `rewrite_virtual_calls`
- `insert_runtime_prelude`
- `elide_blank_identifier_guards`
- `collect_imports`

## go ast pass selection reasons
- `normalize_names` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Canonicalize generated identifiers before rewrite passes.
- `rewrite_string_ops` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Apply planner-selected string rewrite/folding pass for deterministic code shape.
- `rewrite_virtual_calls` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Apply planner-selected safe virtual-call rewrite pass.
- `insert_runtime_prelude` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Inject runtime prelude declarations before cleanup/import collection.
- `elide_blank_identifier_guards` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Remove redundant blank-identifier consume guards after lowering.
- `collect_imports` | `planner(preset=portable_default, auto=auto, opt=portable_fast)` | Collect final deterministic import set after all rewrites.

## portable surface plan consequences

- required imports: `none`
- required runtime features: `array, string`

## portable surface decisions

- `Main` | location `Main:1` | usage `expression` | `haxe.Array` v1 | used type `{"kind":"class","path":"Array","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `rejected:shape_mismatch` | eligibility detail `Observed type shape does not match the contract pattern.` | selection `fallback:registry_rejected` | representation `hxrt_array` | fallback `shape_mismatch: Observed type shape does not match the contract pattern.` | imports `none` | runtime `array` | no-hxrt contract `conditional` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[{"name":"v","optional":false,"shape":{"kind":"dynamic","path":"","parameters":[],"arguments":[],"returnType":null,"fields":[]}},{"name":"infos","optional":true,"shape":{"kind":"abstract","path":"StdTypes.Null","parameters":[{"kind":"typedef","path":"haxe.PosInfos","parameters":[],"arguments":[],"returnType":null,"fields":[]}],"arguments":[],"returnType":null,"fields":[]}}],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `expression` | `haxe.String` v1 | used type `{"kind":"class","path":"String","parameters":[],"arguments":[],"returnType":null,"fields":[]}` | eligibility `admitted:contract_admitted` | eligibility detail `Contract admitted this exact typed shape.` | selection `native:registry_admitted` | representation `go_string` | fallback `none` | imports `none` | runtime `string` | no-hxrt contract `ineligible` | selected no-hxrt eligible `no`
- `Main` | location `Main:1` | usage `function_declaration` | `haxe.Function` v0 | used type `{"kind":"function","path":"","parameters":[],"arguments":[],"returnType":{"kind":"abstract","path":"StdTypes.Void","parameters":[],"arguments":[],"returnType":null,"fields":[]},"fields":[]}` | eligibility `rejected:contract_missing` | eligibility detail `Known portable surface has no admitted contract.` | selection `existing:no_registered_fallback` | representation `none` | fallback `contract_missing: Known portable surface has no admitted contract.` | imports `none` | runtime `none` | no-hxrt contract `none` | selected no-hxrt eligible `no`
