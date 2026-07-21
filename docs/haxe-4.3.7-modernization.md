# Haxe 4.3.7 Compiler Modernization

## What this work is

This program uses the Haxe features already guaranteed by haxe.go's pinned
4.3.7 toolchain to make the compiler implementation easier to check and harder
to misconfigure. It is compiler-maintenance work, not a change to portable Haxe
application semantics or a second Go product mode.

The work has four bounded parts:

1. compile the target with every Haxe warning category enabled and reject new
   warnings;
2. enable null safety only for audited compiler packages that pass it;
3. replace loose metadata, define, diagnostic, and surface-name strings with
   typed identifiers where that prevents invalid states; and
4. keep unavoidable `Dynamic` values at documented macro, Reflaxe, runtime, or
   Haxe-language boundaries while removing avoidable uses.

## Warning ratchet

### What

`npm run test:haxe-warnings` compiles a fresh representative Go program with
Haxe 4.3.7 and `-w +WAll`. The test requires zero warnings from haxe.go-owned
source and enforces per-file, per-warning ceilings for the pinned Haxe standard
library and vendored Reflaxe sources.

### Why

Haxe 4.3 enables all current warning categories by default, but an explicit
gate makes that policy reviewable and prevents a future HXML option from
silently weakening it. The first audit found project-owned direct comparisons
of the recursive `Binop` enum and loop variables that shadowed the printer's
input statement. Pattern matching and unambiguous local names remove those
warnings without changing generated Go.

Warnings in the pinned Haxe compiler standard library and vendored Reflaxe are
not hidden or treated as haxe.go source. They have separate ceilings because
this repository does not own those files. A new warning kind, a warning in a
new external file, or an increase above a reviewed ceiling fails the gate.

### How

The checked-in authority is `test/haxe_warning_policy.json`. A lower count
passes so upstream cleanup is immediately safe, but the policy should be
regenerated downward in the same update so the improvement becomes permanent.
Do not raise a ceiling to silence a failure: first identify the owner, explain
why the warning cannot be fixed here, and review the exact changed entry.

The warning ratchet runs in the changed, full, and release-contract suites.
It complements the compiler-debt ratchet: compiler debt measures intentionally
dynamic or raw implementation boundaries, while this gate measures diagnostics
emitted by the pinned Haxe type checker.

## Scoped null safety

### What

Both source-checkout and installed-package configuration enable
`nullSafety("reflaxe.go")`. This opts the compiler and macro implementation into
Haxe null-safety analysis without changing application packages, staged Haxe
standard-library APIs, or the public `go.*` surface.

### Why

The compiler owns many optional typed-AST and configuration values. Checking
that package catches accidental dereferences while preserving intentional
nullable state such as optional expressions, delayed macro initialization, and
lookups that can legitimately fail. Applying null safety repository-wide would
incorrectly include public Haxe `Dynamic` semantics and target shims that need
separate ownership audits.

### How

The source library declares the macro in `haxe_libraries/reflaxe.go.hxml`; the
published Haxelib package declares the same macro in `extraParams.hxml`. The
warning contract verifies both entry points, and the ordinary source/package,
snapshot, semantic, and example builds exercise the audited package under the
same setting.

## Typed compiler identifiers

### What

`GoCompilerDefine`, `GoMetadataName`, `GoContractDiagnosticCode`,
`GoContractDiagnosticSeverity`, and `GoHxrtFeatureId` are closed enum abstracts
for strings that cross compiler subsystems. They convert to `String` when the
Haxe macro API or a report needs text, but arbitrary strings do not implicitly
convert back into trusted identifiers.

### Why

These names are small protocol fields: one typo can disable a compiler option,
miss a source annotation, omit a runtime file, or create a report value that no
consumer recognizes. Keeping the spelling in one typed registry makes the
accepted vocabulary searchable and lets Haxe reject accidental values.

### How

Generic helpers still accept `String` when their purpose is to inspect a
caller-supplied target or define. At haxe.go-owned call sites, compatibility
constants point to the typed registry. Metadata matching also owns Haxe's
optional leading colon in one `matches(...)` helper instead of repeating two
string comparisons throughout the compiler.

`CompilerBootstrap` has one documented exception: it keeps the `reflaxe`
library define as a local typed `String` constant because that bootstrap must
compile in isolation before the shared haxe.go compiler package or vendored
Reflaxe sources are available on the classpath.

## Dynamic boundary audit

### What

`npm run test:haxe-dynamic-boundaries` reports compiler implementation sites
that use `Dynamic` as a Haxe type. The reviewed policy permits only three sites
in `GoPostBuildRunner`: two catch variables and the helper that immediately
turns an arbitrary thrown value into a redacted string.

### Why

The audit removed four placeholder `Dynamic` types from Reflaxe's legacy
per-node output hooks. Those hooks do not emit haxe.go output, so
`GoReflaxeStagedOutput` now describes the slot explicitly. It also narrowed
three optional `Context.getType(...)`/`Context.getModule(...)` catches to the
documented `String` exception type. The remaining process-launch catches stay
dynamic because Haxe permits arbitrary thrown values at that language boundary.

Generated Go `any`, typed-AST variants such as `TDynamic`, and source-language
`Dynamic` behavior are not compiler implementation debt and are deliberately
outside this narrow count.

### How

`test/haxe_dynamic_boundary_policy.json` is an exact allowlist by source file.
A new site or a moved site fails changed, full, and release-contract testing.
Do not expand it without documenting why a more precise Haxe type cannot model
the boundary.

## Macro lifecycle and compilation-server safety

### What

Initialization code uses Haxe 4.3.7's supported typed phase callbacks:
`onAfterInitMacros` for the final sibling-target configuration check and
`onAfterTyping` for checks that require typed application modules.

### Why

Initialization macros run before ordinary typing, so a typed query made too
early can be incomplete or unsafe when the Haxe compilation server reuses
cached modules. Haxe 4.3.7 marks `onMacroContextReused` deprecated and its
implementation throws; it is not a lifecycle reset mechanism.

### How

Small `initialized`/`bootstrapped` sentinels prevent duplicate registration
inside one macro context. They intentionally do not use `@:persistent`, so the
compilation server does not retain them as cross-build state. The lifecycle
contract rejects the unsupported reuse hook and persistent initialization
sentinels while requiring the phase-appropriate callbacks.

See Haxe's [macro `Context` API](https://api.haxe.org/v/4.3.7/haxe/macro/Context.html)
and [persistent-variable documentation](https://haxe.org/manual/macro-persistent-variables.html)
for the upstream contracts.

## Sibling precedent

The inspected sibling projects do not yet share a family-wide Haxe warning
policy. `haxe.c` already opts its compiler package into null safety. The Rust,
Elixir, Ruby, and Genes checkouts inspected for this work did not expose an
equivalent release-blocking `-w +WAll` ratchet. haxe.go therefore keeps this
policy local and evidence-backed instead of presenting it as an existing
cross-compiler standard.
