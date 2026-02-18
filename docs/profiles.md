# Profiles (`-D reflaxe_go_profile=...`)

This target supports three profiles:

```bash
-D reflaxe_go_profile=portable|gopher|metal
```

## Matrix

| Profile | Best for | Behavior contract |
| --- | --- | --- |
| `portable` (default) | Haxe-first and cross-target code | Stable Haxe-oriented semantics, portability-first output |
| `gopher` | Go-aware teams wanting cleaner Go-first output style | Go-first API/lowering preferences without forcing semantic drift (includes safe compile-time folding for literal string helper ops, typed string helper lowering for `String`+`String`, and leaf-receiver devirtualization for `self`, inline constructor targets, tracked constructor-local aliases, and known leaf-returning call targets) |
| `metal` (experimental) | Teams needing typed low-level interop lane | `gopher` + strict default app-boundary policy + typed framework interop façade |

## What changed

- `idiomatic` was removed.
- `-D reflaxe_go_profile=idiomatic` fails fast and should be replaced with `gopher`.
- `-D reflaxe_go_idiomatic` alias also fails fast.

## Boundary policy

- `reflaxe_go_strict_examples`: forbids raw `__go__` in repo examples/snapshots.
- `reflaxe_go_strict`: forbids raw `__go__` in app project sources.
- `metal` enables strict mode by default for app-side raw injection.

Framework-owned typed facades are allowed in `metal` strict mode; raw app-side injection remains disallowed.
