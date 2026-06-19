# Examples QA Contract

Examples are part of the compiler test surface. They are not loose demos.

This contract keeps examples useful for external users and useful for compiler
maintenance: if compiler, runtime, stdlib, or profile behavior changes, the
examples should tell us whether the user-facing story still compiles and still
behaves as documented.

## What the harness checks

The examples harness is:

```bash
python3 test/run-examples.py
```

For every discovered example profile lane, it checks:

1. The CI HXML compiles with `-D go_no_build`.
2. The generated CI Go module passes `go test ./...`.
3. The generated CI Go module runs with `go run .` and matches
   `expected/*.ci.stdout`.
4. The default HXML compiles with `-D go_no_build`.
5. The generated default Go module passes `go test ./...`.
6. The generated default Go module runs with `go run .` and matches
   `expected/*.stdout`.
7. The default generated tree matches the committed `generated/<profile>` tree.

That means the harness covers both compile validity and behavior validity. The
behavior check is intentionally plain text because examples are user-facing
programs; expected output is easy to inspect in review.

## Example lane requirements

Every example profile lane must compile through the harness.

A profile lane is discovered when these files exist together:

- `compile.<profile>.hxml`
- `compile.<profile>.ci.hxml`

For each discovered lane, the example must also provide:

- `expected/<profile>.stdout`
- `expected/<profile>.ci.stdout`
- `generated/<profile>/`

If an example has a `README.md`, it must have at least one discovered profile
lane. This prevents new example directories from being accidentally invisible to
CI.

## Runtime behavior tests

Every runnable example should have an expected-output contract unless there is a
clear reason not to run it in CI.

Use:

- `run.args` for default profile-lane runtime arguments.
- `run.ci.args` for CI-lane runtime arguments when the CI lane needs different
  inputs.
- `expected/*.stdout` for the exact output users should see.

If an example demonstrates behavior that the compiler cannot catch, encode that
behavior in the expected output. Examples include:

- command dispatch and CLI output,
- HTTP loopback request/response behavior,
- profile/variant selection,
- typed Go interop behavior,
- worker/channel/select behavior.

## Compile-only mode

Use compile-only mode only when a caller intentionally wants a faster smoke:

```bash
python3 test/run-examples.py --compile-only
```

Compile-only mode still compiles generated Go and runs `go test ./...`; it only
skips the expected-output `go run` checks. Full CI should use the full examples
harness, not compile-only mode.

## When to run it

Run `npm run test:examples` when changing:

- compiler lowering,
- runtime `hxrt` behavior,
- staged std overrides,
- profile/strictness behavior,
- example code, README files, HXML files, generated trees, or expected output.

Run `npm run test:examples:changed` for narrow example-only edits. For compiler,
runtime, stdlib, or profile changes, prefer the full example matrix because any
example can be affected.

## Updating examples intentionally

When behavior intentionally changes:

1. Update the Haxe source or HXML.
2. Run `python3 test/run-examples.py` and inspect failures.
3. Update `expected/*.stdout` only when the new behavior is intentional.
4. Run `python3 test/run-examples.py --bless-generated` to refresh committed
   generated trees.
5. Re-run `python3 test/run-examples.py`.
6. Explain the behavior change in the commit or PR notes.

Do not update expected output just to make CI green. Expected output is the
user-visible contract for the example.

## CI ownership

`python3 test/run-ci.py` runs examples during full CI runs by default. GitHub
Actions runs that stable CI surface in `.github/workflows/ci-harness.yml`.

The release contract suite also checks that example directories cannot silently
fall out of the examples harness.
