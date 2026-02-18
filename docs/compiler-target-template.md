# Compiler Target Template (Reflaxe-first)

Use this checklist when building new `reflaxe.<target>` compilers.

## 1) Bootstrap

- Provide `CompilerBootstrap.Start()` and call it from `extraParams.hxml`.
- Inject vendored Reflaxe (`vendor/reflaxe/src`) when present.
- Inject target std overrides only when target build is active.
- Inject classpaths at the front when override precedence matters.
- Include robust target detection (defines + args + nested hxml fallback).

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
