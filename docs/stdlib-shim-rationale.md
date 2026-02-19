# Stdlib Strategy and Shim Rationale

## Context

`reflaxe.go` supports Haxe stdlib behavior through a mix of:

- generated Go runtime support (`runtime/hxrt/hxrt.go`)
- compiler-emitted compatibility shims (`src/reflaxe/go/GoCompiler.hx`)
- optional staged stdlib classpath support (`std/_std`) wired by `src/reflaxe/go/CompilerBootstrap.hx`

This document explains why shims currently live in compiler core, and what simpler alternatives were considered.

## Alternatives Evaluated

### 1. Externs + external Go package

Approach:

- model stdlib surfaces as `extern` Haxe APIs
- implement behavior in a separately maintained Go package

Why this is not enough right now:

- Externs only provide type-level contracts; they do not provide behavior by themselves.
- The current compiler config intentionally ignores extern declarations during code emission (`ignoreExterns: true` in `src/reflaxe/go/CompilerInit.hx`).
- A separate Go package would duplicate compatibility logic outside compiler lowering/profile policy and add another versioned surface to keep in sync.

### 2. Raw `__go__` injection in Haxe std/app code

Approach:

- use target-code injection (`__go__`) directly in Haxe classes for edge behavior

Why this is not the default path:

- Strict policy intentionally forbids this in app/examples lanes (`src/reflaxe/go/macros/StrictModeEnforcer.hx`, `src/reflaxe/go/macros/BoundaryEnforcer.hx`).
- It weakens portability/readability and pushes target logic into user/app code paths.
- It is harder to reason about and test as a reusable compiler contract.

### 3. Vendored stdlib-only (no core shims)

Approach:

- rely only on transpiling vendored stdlib sources under `std/_std`

Status:

- This is a long-term direction and is already classpath-wired.
- In practice, some behavior-heavy surfaces still need explicit target glue and policy-aware lowering (regex/serializer/socket/http contracts and deterministic runtime tradeoffs).

## Current Decision

Use compiler-core shims for behavior-critical stdlib surfaces, with these constraints:

- Inject only when referenced (`requiredStdlibShimGroups` in `src/reflaxe/go/GoCompiler.hx`).
- Keep behavior under CI contract coverage (snapshot + semantic diff + stdlib sweep).
- Prefer minimal, deterministic shims and remove them when a cleaner source-level path is proven.

## Revisit Triggers

We should retire/move shims when one of these is true:

1. Vendored stdlib (`std/_std`) can provide equivalent behavior with equal or better CI parity.
2. A shared external runtime package can provide stable semantics without fragmenting profile/lowering policy.
3. A shim is only forwarding behavior and no longer needs compiler-context decisions.

Until then, shims remain the lowest-risk way to keep Haxe contracts explicit, testable, and profile-aware.
