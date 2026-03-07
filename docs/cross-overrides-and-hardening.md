# `.cross.hx`, `_std`, and Family Hardening Notes

This document records how `reflaxe.go` currently uses `.cross.hx` and `_std`, and what that means for coexistence with sibling Reflaxe targets.

## Current model in this repo

`reflaxe.go` uses `.cross.hx` broadly.

That includes both:

- plain `std/*.cross.hx` files such as `std/StringTools.cross.hx`
- staged `_std/*.cross.hx` files such as `std/_std/haxe/Json.cross.hx`

This repo does **not** currently have an early `src/haxe/*.cross.hx` set.

That matters because it lowers the risk of early module-path collisions compared with `reflaxe.ocaml` and `reflaxe.elixir`.

## Quick matrix

| Question | Answer for this repo |
| --- | --- |
| Main override style | broad `.cross.hx` usage, including `_std/*.cross.hx` |
| Is `_std` used? | yes |
| Is `.cross.hx` used broadly? | yes |
| Does this repo own early `src/haxe/*` modules? | no |
| Bootstrap activation currently keys off raw Haxe 4 `Cross`? | no |
| Same-compilation sibling-target coexistence safe today? | not guaranteed |
| Highest-priority hardening item | add mixed-target fail-fast while preserving narrow target detection |

## What `.cross.hx` means here

In this repo, `.cross.hx` is mostly the normal target-conditional stdlib ownership mechanism.

It is not mainly an early-bootstrap exception mechanism.

That matches the current bootstrap flow:

- inject `std`, `std/_std`, and vendored Reflaxe only when `BuildDetection.isGoBuild()` says the build is really a Go build
- do not activate on raw generic `Cross` alone

That narrower activation model is a good property and should be preserved.

## What `_std` means here

`_std` is still useful, but it is not the only override lane.

This repo combines:

- target-conditional `.cross.hx` selection
- staged `_std` ownership
- runtime/compiler intrinsic lowering where needed

So the key rule here is not "always prefer `_std`".

The key rule is:

- keep bootstrap gating narrow,
- keep ownership explicit,
- and avoid same-module collisions becoming accidental.

## Current coexistence risk

The risk here is lower than in `reflaxe.elixir`, because:

- bootstrap activation is target-specific
- there is no early `src/haxe/*.cross.hx` ownership set

But the risk is not zero.

This repo still shares module names with siblings under `std/**/*.cross.hx`, including overlaps such as:

- `DateTools`
- `StringTools`
- `haxe.CallStack`
- `haxe.Constraints`
- `haxe.NativeStackTrace`

The most important one from a sibling-collision perspective is `haxe.NativeStackTrace`, because `reflaxe.ocaml` currently owns an early `src/haxe/NativeStackTrace.cross.hx`.

If both libraries are loaded into one `cross` compilation, the early sibling file can win resolution first.

## Risk level

Current status:

- default one-target-at-a-time use: acceptable
- same-compilation multi-target coexistence: still risky enough to harden

This repo is not the highest-risk member of the family, but it should still fail clearly instead of relying on classpath luck.

## Hardening direction

Recommended next steps:

1. Keep target detection narrow; do not broaden it to raw generic `Cross`.
2. Add explicit mixed-target detection/fail-fast behavior when conflicting sibling target libraries are active.
3. Document which `std/**/*.cross.hx` modules overlap with sibling targets and why that overlap is currently acceptable only in one-target-at-a-time builds.

## Local sibling references

Workspace-local companion docs:

- `../haxe.ocaml/docs/02-user-guide/CROSS_AND_STAGED_STDLIB_GUIDE.md`
- `../haxe.ocaml/docs/00-project/REFLAXE_FAMILY_CROSS_OVERRIDE_AUDIT.md`
- `../haxe.elixir.codex/docs/05-architecture/CROSS_OVERRIDES_AND_MULTI_TARGET_HARDENING.md`
- `../haxe.rust/docs/cross-overrides-and-hardening.md`

These sibling-relative paths are intended for local multi-repo work, not for a single published docs site.

## Absolute-path protection

This repo already has staged local-path leak protection in pre-commit via:

- `scripts/hooks/pre-commit`
- `scripts/lint/local_path_guard_staged.sh`

So the hardening gap here is mixed-target clarity, not hook absence.
