# std Agent Instructions

This directory hosts target-facing standard library layers and compatibility shims.

## Injection Rules

- `__go__` usage is allowed only inside std/runtime implementation boundaries.
- Do not introduce raw injection into app-level examples or snapshot fixture code.
- Keep injected snippets small, explicit, and easy to replace with typed lowering over time.

## API Compatibility

- Keep public signatures stable and compatible with expected Haxe std/coreApi contracts.
- Preserve nullable/string/error semantics expected by portable mode unless explicitly profile-gated.
- Prefer additive changes over breaking signature rewrites.

## Implementation Guidance

- Avoid leaking target-only wrapper types through portable APIs unless the module is explicitly `go_native`.
- Keep behavioral differences documented and covered by snapshot tests.
- Add/update regression snapshots for every bug fix or shim behavior change.
