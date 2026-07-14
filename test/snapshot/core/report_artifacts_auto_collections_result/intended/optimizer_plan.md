# Optimizer Plan Report

- schema version: `6`
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
- go collections typed lowerings: `7`
- go collections typed fallbacks: `0`
- go result typed lowerings: `3`
- go result typed fallbacks: `0`
- lowering fallback boundary count: `0`
- lowering fallback non-boundary count: `0`
- lowering fallback lane count: `0`
- lowering fallback non-lane count: `0`
- go ast pass selection source: `planner`

## auto lowering capabilities
- `go.collections.typed` | attempts `7` | success `7` | fallback `0`
  fallback reasons: none
- `go.concurrency.typed` | attempts `2` | success `0` | fallback `0`
  fallback reasons: none
- `go.result.typed` | attempts `3` | success `3` | fallback `0`
  fallback reasons: none

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
