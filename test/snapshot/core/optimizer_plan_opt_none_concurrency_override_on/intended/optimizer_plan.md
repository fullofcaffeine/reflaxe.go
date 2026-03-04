# Optimizer Plan Report

- schema version: `4`
- contract: `portable`
- auto lowering mode: `off`
- optimization preset: `none`
- portable string fastpath enabled: `no`
- portable concurrency fastpath enabled: `yes`
- string instance typed lowerings: `0`
- string instance legacy lowerings: `4`
- string length field typed lowerings: `0`
- string length field legacy lowerings: `2`
- portable concurrency typed fastpath hits: `3`
- portable concurrency typed fastpath fallbacks: `0`
- go collections typed lowerings: `0`
- go collections typed fallbacks: `0`
- go result typed lowerings: `0`
- go result typed fallbacks: `0`
- lowering fallback lane count: `0`
- lowering fallback non-lane count: `0`
- go ast pass selection source: `planner`

## go ast passes
- `normalize_names`
- `rewrite_string_ops`
- `rewrite_virtual_calls`
- `insert_runtime_prelude`
- `elide_blank_identifier_guards`
- `collect_imports`

## go ast pass selection reasons
- `normalize_names` | `planner(contract=portable, auto=off, opt=none)` | Canonicalize generated identifiers before rewrite passes.
- `rewrite_string_ops` | `planner(contract=portable, auto=off, opt=none)` | Apply planner-selected string rewrite/folding pass for deterministic code shape.
- `rewrite_virtual_calls` | `planner(contract=portable, auto=off, opt=none)` | Apply planner-selected safe virtual-call rewrite pass.
- `insert_runtime_prelude` | `planner(contract=portable, auto=off, opt=none)` | Inject runtime prelude declarations before cleanup/import collection.
- `elide_blank_identifier_guards` | `planner(contract=portable, auto=off, opt=none)` | Remove redundant blank-identifier consume guards after lowering.
- `collect_imports` | `planner(contract=portable, auto=off, opt=none)` | Collect final deterministic import set after all rewrites.
