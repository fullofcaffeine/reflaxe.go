# Optimizer Plan Report

- schema version: `4`
- contract: `portable`
- auto lowering mode: `off`
- optimization preset: `portable_fast`
- portable string fastpath enabled: `yes`
- portable concurrency fastpath enabled: `yes`
- string instance typed lowerings: `0`
- string instance legacy lowerings: `0`
- string length field typed lowerings: `0`
- string length field legacy lowerings: `0`
- portable concurrency typed fastpath hits: `0`
- portable concurrency typed fastpath fallbacks: `0`
- lowering fallback lane count: `0`
- lowering fallback non-lane count: `0`
- go ast pass selection source: `legacy_granular_bundle`

## go ast passes
- `normalize_names`
- `rewrite_string_ops`
- `rewrite_virtual_calls`
- `insert_runtime_prelude`
- `elide_blank_identifier_guards`
- `collect_imports`

## go ast pass selection reasons
- `normalize_names` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
- `rewrite_string_ops` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
- `rewrite_virtual_calls` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
- `insert_runtime_prelude` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
- `elide_blank_identifier_guards` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
- `collect_imports` | `legacy_granular_bundle` | Selected by compatibility define `-D go_granular_pass_registry`.
