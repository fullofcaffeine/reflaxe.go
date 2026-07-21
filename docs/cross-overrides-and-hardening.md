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

No `.cross.hx` files are checked in. Repo-authored target support remains
ordinary `.hx` source under `std/haxe/**` and `std/sys/**`; typed bindings for
real Go runtime APIs live under `std/hxrt/**`. Narrow source-callable adapters
for compiler-owned closed-world metadata live under the exact internal root
`std/reflaxe/go/internal/**`; this root is not a public facade or a general
extension point. Exact ownership and canonical paths live in
`docs/stdlib-provenance-ledger.json`.

## Quick Matrix

| Question | Answer for this repo |
| --- | --- |
| Canonical override source | ordinary `.hx` under `std/go/_std/**` |
| Packaged override shape | flattened `src/**/*.cross.hx`, generated during package staging |
| Checked-in upstream override `.cross.hx` files | none |
| Checked-in support/runtime `.cross.hx` files | none |
| Public Go facades | ordinary modules under `std/go/**`, outside `_std` |
| Internal compiler-metadata adapters | ordinary modules under the narrow `std/reflaxe/go/internal/**` root |
| Does this repo own early `src/haxe/*` modules? | no |
| Bootstrap activation keys off raw Haxe 4 `Cross`? | no |
| Same-compilation sibling-target coexistence safe today? | rejected before application typing |

## What `.cross.hx` Means Here

A `.cross.hx` file is a package artifact that lets Haxe select a
target-specific replacement for an upstream module. It is not the source
authority for an upstream override.

Keeping the two shapes separate matters:

1. source review, HaxeDoc, and migration history stay on ordinary Haxe modules;
2. only declared canonical overrides become `.cross.hx`;
3. support modules and public native facades retain ordinary `.hx` paths;
4. package generation can prove a deterministic source-to-artifact manifest.

Support/runtime modules never enter that conversion because they do not
replace upstream Haxe modules.

## What `_std` Means Here

The canonical target root is `std/go/_std`. Its directory structure mirrors
upstream Haxe module paths, so `std/go/_std/haxe/Json.hx` owns the Go-target
replacement for `haxe.Json`.

Source builds declare `src`, ordinary `std`, and canonical `std/go/_std` in
`haxe_libraries/reflaxe.go.hxml` before any macro is typed. The canonical root
comes last, which gives it effective Haxe override precedence. Ordinary `std`
exposes support, typed `hxrt` bindings, and public `go.*` facades without an
extra override root. A companion `haxe_libraries/reflaxe.hxml` supplies
vendored Reflaxe at the same initial configuration stage.

`CompilerBootstrap` no longer changes classpath order. It only provides a
typed, non-conflicting vendored-Reflaxe fallback for direct
`extraParams.hxml` consumers and diagnoses an invalid source/package layout.

## Current Coexistence Policy

The risk is lower than in a layout that places target modules directly under
`src/haxe`:

- activation remains Go-target-specific;
- canonical source is isolated under `std/go/_std`;
- the ledger distinguishes upstream overrides from support and public facades.

The risk is not zero. Canonical overrides still share logical module names
with sibling targets, including `DateTools`, `StringTools`,
`haxe.CallStack`, `haxe.Constraints`, and `haxe.NativeStackTrace`.
If multiple target libraries mutate or declare competing classpaths in one
compilation, selection could otherwise depend on ordering.

Current status:

- default one-target-at-a-time use: acceptable;
- same-compilation multi-target coexistence: rejected with a deterministic
  diagnostic rather than left to classpath luck.

## Fail-fast contract

`SiblingTargetConflictGuard` runs only after `BuildDetection` has established
that this is a Go build. It checks twice: once at the start of Go compiler
initialization and once after all initialization macros. The second check
catches a sibling library expanded later in the same Haxe command, but both
checks still happen before ordinary application modules are typed.

The guard accepts two narrow kinds of evidence:

- **source-checkout signal**: an active classpath is exactly a canonical
  sibling `std/<target>/_std` root;
- **packaged signal**: a supported sibling library/version define such as
  `reflaxe.rust` or `reflaxe.elixir` is active. Normal `-lib` package loading
  supplies this identity even though Reflaxe has flattened the package's
  `_std` files into `src/**/*.cross.hx`.

It recognizes `genes`, `reflaxe.c`, `reflaxe.elixir`, `reflaxe.ocaml`,
`reflaxe.ruby`, and `reflaxe.rust`. The diagnostic sorts those stable names
and never prints machine-local paths. It does not reject arbitrary additional
classpaths, a source directory merely named after a language, or non-Go use of
the `reflaxe.go` bootstrap.

For example, activating both a Rust source `_std` root and a packaged Ruby
target during a Go build fails with:

```text
Reflaxe.Go cannot compile with competing sibling targets: reflaxe.ruby, reflaxe.rust.
```

This is a safety boundary, not a multi-backend feature: it **does not make
simultaneous multi-target compilation supported**. Compile each target in its
own Haxe invocation so its initial classpath, stdlib ownership, compiler
registration, and output policy remain internally consistent.

## Hardening Status

Package generation and isolated installed-package selection are now enforced:

1. `.cross.hx` files are generated only while staging a package;
2. source and installed-package imports are behavior-tested against the same
   declared module ownership;
3. the installed ZIP compiles and runs in a fresh local Haxelib repository
   without checkout classpaths.

Mixed-target clarity is now enforced by source-layout, package-layout,
late-initialization, harmless-classpath, and non-Go regression fixtures. The
guard is part of `npm test`, `npm run test:changed`, and the release-contract
suite. The inventory and installed-package closeout evidence remain recorded
in `docs/canonical-std-migration-closeout.md`.

## Sibling comparison

This comparison uses read-only committed sibling snapshots, not a claim that
their current policies must match Go's. Their target identities and override
layouts are useful inputs; the fail-fast implementation remains owned by this
repository.

| Sibling evidence | Observed policy at the inspected commit | What haxe.go uses |
| --- | --- | --- |
| reflaxe.rust at `85067736d0b929dfc67d6684d59b7e2bd3bae6ea` | Its hardening document explicitly calls mixed-target fail-fast a next step; no equivalent guard was found. | Recognize `reflaxe.rust` and `std/rust/_std`. |
| reflaxe.elixir at `17f1c66ae4c6bcae3c15cf694c16e63f27f2d9aa` | Its hardening document also calls for fail-fast and identifies broad Haxe 4 `Cross` activation as an additional local risk. | Recognize `reflaxe.elixir` and `std/elixir/_std` without copying its bootstrap behavior. |
| reflaxe.ruby at `d32a4adeaece13a86e768b405ee51f80a58e996c` | Bootstrap injects its target `_std` root for Ruby builds; no mixed-target guard was found. | Recognize `reflaxe.ruby` and `std/ruby/_std`. |
| genes at `2b4b71b00528fb376f7f0f8527237cf336b0f36b` | It has selected early `src/haxe/ds` modules rather than the same `_std` layout; no general mixed-target guard was found. | Recognize its target defines, not every `src/haxe` classpath. |
| reflaxe.c at `6167c3e6d66472a181a2c694b3136d8174e11803` | Its bootstrap keeps target activation narrow and does not expose an equivalent family conflict guard. | Recognize `reflaxe.c` and the canonical `std/c/_std` shape when present. |
| reflaxe.ocaml at `faaffaa1aa582dc1815245548c6f5d23c670e24e` | Its scoped library configuration declares `std/ocaml/_std`; no equivalent guard was found. | Recognize `reflaxe.ocaml` and `std/ocaml/_std`. |

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

Mixed-target clarity and path-leak prevention are separate gates: the former
protects module selection, while the latter protects committed/generated
artifacts.
