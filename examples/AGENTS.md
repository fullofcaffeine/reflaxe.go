# Agent Instructions for `examples/`

- Examples are treated as real app code: avoid raw `__go__()` injections in example sources.
- Keep examples deterministic and CI-friendly. If behavior is interactive, provide scripted harness flows.
- Example `.hxml` files should include `-D reflaxe_go_strict_examples`.
- Prefer one source tree per example with `compile.<profile>.hxml` variants rather than duplicating folders.
- Keep baseline semantics comparable across profiles; profile-specific behavior should be additive and documented.
