# Canonical `_std`, `.cross.hx`, and Family Hardening Notes

This document records how `reflaxe.go` separates canonical source ownership
from packaged Reflaxe overrides, and what that means for coexistence with
sibling targets.

## Current Model in This Repo

Upstream Haxe std overrides are ordinary `.hx` source under:

```text
std/go/_std/**
```

The ownership ledger currently classifies 62 files in that override lane.
They are not checked in as `.cross.hx`; the package runner owns the later,
deterministic conversion to flattened `src/**/*.cross.hx` artifacts.

Six checked-in `.cross.hx` files remain temporarily. They are classified as
target support or a typed runtime binding, not upstream overrides, and are
owned by the separate support migration:

- `std/haxe/io/GoIoHelpers.cross.hx`
- `std/sys/GoHttpHelpers.cross.hx`
- three `sys.thread` worker/sentinel companions
- `std/_std/haxe/iterators/GoStringRuntime.cross.hx`

The nine ordinary `std/_std/hxrt/**` modules are likewise typed runtime
bindings awaiting that support migration. Exact source and destination paths
live in `docs/stdlib-provenance-ledger.json`.

## Quick Matrix

| Question | Answer for this repo |
| --- | --- |
| Canonical override source | ordinary `.hx` under `std/go/_std/**` |
| Packaged override shape | flattened `src/**/*.cross.hx`, generated during package staging |
| Checked-in upstream override `.cross.hx` files | none |
| Transitional support/runtime `.cross.hx` files | six |
| Public Go facades | ordinary modules under `std/go/**`, outside `_std` |
| Does this repo own early `src/haxe/*` modules? | no |
| Bootstrap activation keys off raw Haxe 4 `Cross`? | no |
| Same-compilation sibling-target coexistence safe today? | not guaranteed |

## What `.cross.hx` Means Here

A `.cross.hx` file is a package artifact that lets Haxe select a
target-specific replacement for an upstream module. It is not the source
authority for an upstream override.

Keeping the two shapes separate matters:

1. source review, HaxeDoc, and migration history stay on ordinary Haxe modules;
2. only declared canonical overrides become `.cross.hx`;
3. support modules and public native facades retain ordinary `.hx` paths;
4. package generation can prove a deterministic source-to-artifact manifest.

The remaining checked-in support/runtime `.cross.hx` files are transitional,
not precedent for new overrides.

## What `_std` Means Here

The canonical target root is `std/go/_std`. Its directory structure mirrors
upstream Haxe module paths, so `std/go/_std/haxe/Json.hx` owns the Go-target
replacement for `haxe.Json`.

Source builds declare `src`, ordinary `std`, the transitional `std/_std`
support root, and canonical `std/go/_std` in
`haxe_libraries/reflaxe.go.hxml` before any macro is typed. The canonical root
comes last, which gives it effective Haxe override precedence. A companion
`haxe_libraries/reflaxe.hxml` supplies vendored Reflaxe at the same initial
configuration stage.

`CompilerBootstrap` no longer changes classpath order. It only provides a
typed, non-conflicting vendored-Reflaxe fallback for direct
`extraParams.hxml` consumers and diagnoses an invalid source/package layout.

## Current Coexistence Risk

The risk is lower than in a layout that places target modules directly under
`src/haxe`:

- activation remains Go-target-specific;
- canonical source is isolated under `std/go/_std`;
- the ledger distinguishes upstream overrides from support and public facades.

The risk is not zero. Canonical overrides still share logical module names
with sibling targets, including `DateTools`, `StringTools`,
`haxe.CallStack`, `haxe.Constraints`, and `haxe.NativeStackTrace`.
If multiple target libraries mutate or declare competing classpaths in one
compilation, selection can still depend on ordering.

Current status:

- default one-target-at-a-time use: acceptable;
- same-compilation multi-target coexistence: must fail clearly rather than rely
  on classpath luck.

## Hardening Direction

The remaining sequence is explicit:

1. move support and typed runtime bindings to their ledger destinations;
2. keep target classpaths declared before typing without reflective
   configuration surgery;
3. generate `.cross.hx` only while staging a package;
4. keep mixed-target detection narrow and fail fast when sibling targets
   conflict.

## Local Sibling References

Workspace-local companion docs:

- `../haxe.ocaml/docs/02-user-guide/CROSS_AND_STAGED_STDLIB_GUIDE.md`
- `../haxe.ocaml/docs/00-project/REFLAXE_FAMILY_CROSS_OVERRIDE_AUDIT.md`
- `../haxe.elixir.codex/docs/05-architecture/CROSS_OVERRIDES_AND_MULTI_TARGET_HARDENING.md`
- `../haxe.rust/docs/cross-overrides-and-hardening.md`

These sibling-relative paths are intended for local multi-repo work, not for a
single published docs site.

## Absolute-Path Protection

This repo has staged local-path leak protection in pre-commit through:

- `scripts/hooks/pre-commit`
- `scripts/lint/local_path_guard_staged.sh`

The remaining hardening gap is mixed-target clarity, not hook absence.
