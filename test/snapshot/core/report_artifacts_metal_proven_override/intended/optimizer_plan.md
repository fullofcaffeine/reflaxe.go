# Optimizer Plan Report

- schema version: `6`
- contract: `metal`
- policy preset: `metal_compatibility`
- native specialization policy: `proven` (source `reflaxe_go_native_specialization`)
- auto lowering mode: `off`
- optimization preset: `none`
- portable string fastpath enabled: `no`
- portable concurrency fastpath enabled: `no`
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
- `normalize_names` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Canonicalize generated identifiers before rewrite passes.
- `rewrite_string_ops` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Apply planner-selected string rewrite/folding pass for deterministic code shape.
- `rewrite_virtual_calls` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Apply planner-selected safe virtual-call rewrite pass.
- `insert_runtime_prelude` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Inject runtime prelude declarations before cleanup/import collection.
- `elide_blank_identifier_guards` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Remove redundant blank-identifier consume guards after lowering.
- `collect_imports` | `planner(preset=metal_compatibility, auto=off, opt=none)` | Collect final deterministic import set after all rewrites.
