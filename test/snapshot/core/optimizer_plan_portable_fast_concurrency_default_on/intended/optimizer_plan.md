# Optimizer Plan Report

- schema version: `3`
- contract: `portable`
- auto lowering mode: `off`
- optimization preset: `portable_fast`
- portable string fastpath enabled: `yes`
- portable concurrency fastpath enabled: `yes`
- string instance typed lowerings: `4`
- string instance legacy lowerings: `0`
- string length field typed lowerings: `2`
- string length field legacy lowerings: `0`
- portable concurrency typed fastpath hits: `3`
- portable concurrency typed fastpath fallbacks: `0`
- lowering fallback lane count: `0`
- lowering fallback non-lane count: `2`

## go ast passes
- `normalize_names`
- `rewrite_string_ops`
- `rewrite_virtual_calls`
- `insert_runtime_prelude`
- `elide_blank_identifier_guards`
- `collect_imports`
