# Compiler Target Template (Reflaxe-first)

Use this checklist when building new `reflaxe.<target>` compilers.

## 1) Bootstrap

- Declare compiler source, target support, and target `_std` roots in the
  library HXML before any bootstrap/init macro. Put the target `_std` root
  after ordinary std so it has effective Haxe override precedence.
- Resolve source-checkout paths through `${SCOPE_DIR}` so nested project and
  example builds use the same initial configuration.
- Expose vendored Reflaxe (`vendor/reflaxe/src`) through a companion library
  HXML instead of discovering it through compiler internals.
- Provide `CompilerBootstrap.Start()` through `extraParams.hxml` only for
  typed validation or a non-conflicting vendored-framework fallback.
- Never reorder classpaths through `Dynamic`/`Reflect` access to
  `Compiler.getConfiguration()`.
- Keep robust target detection (defines + args + nested HXML fallback) in
  compiler initialization and policy macros, not in std override selection.

## 2) Init + registration

- Keep one registration entrypoint: `CompilerInit.Start()`.
- Resolve profile exactly once in init.
- Initialize strict/boundary macro policy here.
- Prefer Reflaxe-native registration flow (`ReflectCompiler.Start` + `AddCompiler`) unless there is a documented blocker.

## 3) AST pipeline

- Keep explicit three-stage flow: build -> transform passes -> printer.
- Use pass registry with validation:
  - duplicate names
  - missing dependencies
  - cycle detection
- Keep lean bundle default with optional granular bundle.

## 3.5) Stdlib ownership model

- Start hybrid unless there is strong evidence for a pure approach:
  - runtime package helpers
  - compiler shims for compile-time-context-sensitive behavior
  - staged stdlib migration path
- Treat this as target-agnostic architecture, then document target-specific pressure points.
- For each surface, choose ownership by evidence:
  - runtime when behavior is reusable target-runtime logic
  - compiler when behavior depends on typed metadata/profile lowering
  - staged stdlib when parity is proven and maintenance cost drops
- Add migration criteria up front (tests/perf/complexity thresholds) before moving ownership.

## 4) Boundary policy

- Enforce strict examples policy in repo examples/snapshots.
- Enforce user strict mode in app sources.
- Implement enforcers on typed AST, not raw file scanning.
- If experimental low-level profile exists, allow only framework-owned typed facades.

## 5) Docs + tests parity

- Every public profile/define must be documented.
- Every profile/define must be validated by snapshot tests.
- Add positive and negative cases for profile conflicts and invalid values.
- Keep an examples/snapshot matrix that proves behavioral contract.

## 6) Decision record

Every major architecture or policy decision should include:

- date
- decision statement
- alternatives considered
- acceptance criteria
- rollback/follow-up trigger

## 7) Tooling split (recommended pattern)

- Use `npm` scripts as the cross-target workflow surface (setup, test harness, dev wrappers).
- Use the target-native toolchain for generated output lifecycle.
- Add one root `scripts/dev/<target>-hx.sh` wrapper that:
  - selects `compile*.hxml` via `--project`, `--profile`, `--ci`
  - runs Haxe compile
  - resolves output directory from target define (`go_output`, `rust_output`, etc.)
  - runs target action (`run`, `build`, `test`, etc.)
- Mirror this wrapper pattern in `templates/basic` for consumer projects.

## 8) Profile documentation contract

- Ship a dedicated profile semantics guide (`portable` vs `metal`).
- Explicitly document:
  - semantic guarantees
  - codegen expectations
  - migration rules (`portable` -> `metal`, `metal` -> `portable`)
  - cross-target interoperability guidance
- Keep this guide linked from `README`, `start-here`, and `profiles` reference docs.
