# Optimizer Plan Report

- schema version: `4`
- contract: `portable`
- auto lowering mode: `off`
- optimization preset: `portable_fast`
- portable string fastpath enabled: `yes`
- portable concurrency fastpath enabled: `yes`
- string instance typed lowerings: `4`
- string instance legacy lowerings: `0`
- string length field typed lowerings: `2`
- string length field legacy lowerings: `0`
- portable concurrency typed fastpath hits: `0`
- portable concurrency typed fastpath fallbacks: `0`
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
- `normalize_names` | `planner` | No planner reason recorded; pass selected by registry ordering.
- `rewrite_string_ops` | `planner` | No planner reason recorded; pass selected by registry ordering.
- `rewrite_virtual_calls` | `planner` | No planner reason recorded; pass selected by registry ordering.
- `insert_runtime_prelude` | `planner` | No planner reason recorded; pass selected by registry ordering.
- `elide_blank_identifier_guards` | `planner` | No planner reason recorded; pass selected by registry ordering.
- `collect_imports` | `planner` | No planner reason recorded; pass selected by registry ordering.
