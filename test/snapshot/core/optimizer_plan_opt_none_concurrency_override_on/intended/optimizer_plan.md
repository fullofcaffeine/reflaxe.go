# Optimizer Plan Report

- schema version: `2`
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

## go ast passes
- `normalize_names`
- `rewrite_string_ops`
- `rewrite_virtual_calls`
- `insert_runtime_prelude`
- `elide_blank_identifier_guards`
- `collect_imports`
