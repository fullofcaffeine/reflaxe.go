# PRD: reflaxe.go — Haxe 4.3.7 → Go target (Reflaxe-based)

This PRD is optimized for LLM-agent implementation: it specifies a deterministic **test harness** (snapshot + build + runtime smoke), a modular compiler architecture (AST-first), and clear milestones with acceptance criteria.

## 0) What I extracted from your shipped targets (patterns worth copying)

From **reflaxe.elixir** and **reflaxe.rust** (repomix bundles), the “winning” patterns are:

### A. Bootstrapping macros: *early* and *front-of-classpath*
- **CompilerBootstrap** runs *first* from `extraParams.hxml`.
- It detects “is this build actually for our target?” robustly (Haxe 4 often compiles under `cross`).
- It injects:
  - vendored Reflaxe (`vendor/reflaxe/src`)
  - target std overrides (`std/`, and optionally staged `_std/`)
- Elixir’s bootstrap additionally injects **at the front** of the classpath list to ensure overrides *shadow* canonical stdlib. This is a major correctness/CI-stability factor.

**Action for Go:** implement **Elixir-grade bootstrap** (not the simpler Rust-grade one), because Go will need std overrides and will be used from consumer projects where ordering matters.

### B. CompilerInit macro: single registration point + policy defines
- `CompilerInit.Start()` gates on a target signal (e.g., `-D rust_output` / `-D elixir_output`).
- It calls `ReflectCompiler.Start()` and registers the target via `ReflectCompiler.AddCompiler(...)`.
- It defines policy toggles at compile-time (profiles, strict mode, threaded support, string policy).

**Action for Go:** do the same and make it the only place where:
- profile is resolved (`portable|idiomatic|gopher|metal` concept)
- strict/boundary macros are enabled
- std path injection is belt-and-suspenders (bootstrap first, init second)

### C. AST-first pipeline with a real pass registry
- Elixir: explicitly 3-phase: **Builder → Transformer passes → Printer**.
- Pass registry is *modularized* into groups and supports:
  - “lean” bundle registry (default)
  - “granular” registry (opt-in via define)
  - registry validation: unique pass names, missing deps, cycle detection

Rust is “AST-first” too (Rust AST + transformer + printer), even if much is in one file.

**Action for Go:** replicate Elixir’s registry structure from day one (even with just a few passes), because Go will need many semantic/idiom passes and you’ll want ordering clarity + determinism.

### D. “Purity boundary” macros (hugely valuable for a portable target)
Rust has:
- `BoundaryEnforcer` (used by snapshots/examples) forbids raw `__rust__()` usage in tests/examples.
- `StrictModeEnforcer` (opt-in for user projects; metal enables it by default).

**Action for Go:** implement the same for `__go__()` and use it in snapshot `.hxml` to keep the test suite honest (forces missing features into std/runtime rather than app-local escape hatches).

### E. Snapshot harness patterns (this is the big one)
You have two strong harness styles:

**Elixir**
- `test/Makefile` drives snapshot compilation with:
  - per-test timeout
  - parallelism that’s resilient
  - intended vs out diff (with explicit excludes)
  - negative tests that must fail
  - runtime smoke execution target
  - “update intended” target
  - deterministic run-id temp result files

**Rust**
- Snapshot layout is strict:
  - `test/snapshot/<case>/compile.hxml`
  - `intended/` golden output
  - `out/` generated each run
- Harness (described in `test/README.md`) does:
  - `haxe ... -D rust_no_build` (codegen-only)
  - `cargo fmt -- --check`
  - `cargo build -q`
  - diff intended vs out
  - `--update` mode

**Action for Go:** create a single **LLM-friendly deterministic harness script** that merges the best of both:
- snapshot compile
- gofmt
- `go test ./...` (build)
- diff intended/out
- optional runtime smoke (stdout compare)
- negative tests
- update intended
- `--case`, `--category`, `--failed`, `--changed`, `--chunk` selection
- timeouts
- clean artifacts unless `KEEP_ARTIFACTS=1`

---

## 1) Executive summary

Build a Reflaxe-based Go backend that supports:

1) **Portable Haxe mode (default)**  
   Prioritizes Haxe semantics + cross-target portability and supports meaningful Haxe stdlib coverage. Output is Go that compiles cleanly and is readable, but may use runtime wrappers where needed.

2) **Go-native mode (opt-in, “sweet spot”)**  
   Adds a Haxe-facing `std/go/**` API layer (macros + abstracts + externs) that exposes native Go idioms:
   - `error`-style results
   - slices/maps/channels
   - goroutines + select
   - context cancellation
   - go modules/import hints
   while still allowing portable code to compile unchanged.

Support **Haxe 4.3.7**.

---

## 2) Goals

### G1: Correctness + portability first
- Portable Haxe code (no `#if go`) should compile and behave consistently.
- Haxe stdlib modules required by common programs should work (incremental parity ladder).

### G2: Idiomatic Go output where it doesn’t compromise semantics
- Use Go naming conventions where feasible.
- Prefer Go constructs (`switch`, `range`, slices, maps) when semantics match.
- Always produce gofmt-able output (and ideally stable after gofmt).

### G3: A Go-native Haxe API surface that feels good
- `std/go/**` provides typed, ergonomic Haxe APIs for Go’s strengths.
- Raw injection `__go__` exists but is discouraged and can be banned in app code.

### G4: Maintainable compiler code
- AST-first (abstract syntax tree) pipeline.
- Pass registry with ordering rules + validation.
- Minimal global mutable state; use a `CompilationContext`.

### G5: Deterministic, agent-friendly harness
- One command to run snapshots deterministically.
- Clear diffs on failure.
- Easy to run one case.
- Easy to update intended outputs intentionally.
- Records failing tests.

---

## 3) Non-goals (initially)

- NG1: “Perfect zero-cost Go” for all features (wrappers are fine).
- NG2: Full Haxe stdlib parity immediately (we’ll ladder it).
- NG3: Modeling *all* Go features directly in portable Haxe (some are opt-in).
- NG4: Generating multi-package Go outputs with perfect import graph (start single-package to avoid cycles; revisit later).

---

## 4) Hard architecture decisions (make these explicit now)

These are the friction points agents will hit unless the PRD nails them.

### 4.1 Output packaging: single Go package (initially)
**Decision:** Generate **one Go package** per build (all files in one folder / one `package`).

**Why**
- Go import cycles are easy to create if we map Haxe packages → Go packages.
- Single package avoids cyclic imports and simplifies early correctness.

**Consequence**
- We must prevent name collisions across fully-qualified Haxe names.
- We’ll use a deterministic **name mangling scheme** for all top-level identifiers.

### 4.2 Deterministic naming/mangling scheme
**Decision:** Every emitted top-level identifier must include enough namespace to be unique.

Recommended scheme:
- Type `haxe.ds.IntMap` → `Haxe_Ds_IntMap`
- Function `haxe.Log.trace` (if emitted) → `Haxe_Log_trace`
- Enum constructor `haxe.ds.Option.Some` → `Haxe_Ds_Option_Some`

Rules:
- Segment separator: `_`
- Segment casing:
  - package segments: PascalCase (`haxe` → `Haxe`, `ds` → `Ds`)
  - type segments: original PascalCase
- Escape Go keywords by suffixing `_` (`type` → `type_`) *after mangling*.
- Keep the mangling algorithm in **one file**: `src/reflaxe/go/naming/GoNaming.hx`.

### 4.3 Class inheritance + virtual dispatch in Go
Go has no inheritance, and embedding alone breaks “base method calls derived override”.

**Decision:** Use a **hybrid “self-dispatch”** pattern (sweet spot between correctness and Go-ness):

- Each class `B` has a Go struct `B_Impl` for storage (fields).
- `B_Impl` includes a field `__hx_this IB` where `IB` is an interface for B’s instance methods that need virtual dispatch (initially: all instance methods).
- Derived `C_Impl` embeds `B_Impl` and in constructor sets `b.__hx_this = c` (as `IB`), so base methods can dispatch to derived overrides.

Call rules:
- Call on a receiver of static type `*B_Impl`: emit `b.__hx_this.Method(...)` (virtual dispatch).
- Call on static type `*C_Impl`: direct call `c.Method(...)` is fine (but can also use `__hx_this` uniformly if present).

**Why**
- Preserves virtual dispatch semantics from within base methods.
- Avoids rewriting every field access into interface getter/setter calls.
- Keeps storage as real structs (Go-friendly).

**Edge cases to call out explicitly**
- Constructors must initialize `__hx_this` for every base in the chain.
- `super.method()` must call the base implementation directly (not through `__hx_this`).

This decision should be implemented behind a small runtime helper pattern so we can later switch models if needed.

### 4.4 String nullability (Haxe String can be null; Go string cannot)
**Decision:** In `portable` and `idiomatic` profiles, represent Haxe `String` as `*string` (pointer) and implement helpers in runtime:
- `hx.StringFromLiteral("x") -> *string`
- `hx.StringEq(a,b) bool` with nil semantics
- `hx.StringConcat(a,b) *string` with Haxe semantics (`null + "x" == "nullx"`)

In `gopher/metal`, optionally allow a define to treat strings as non-nullable `string` (with strict mode forbidding `null` string usage), mirroring Rust’s string policy pattern.

### 4.5 Snapshot determinism knobs
Always compile snapshots with:
- `-D reflaxe.dont_output_metadata_id` (prevents nondeterministic metadata IDs)
- `--no-traces -D no_traces` (reduces noise)
- `HAXE_NO_SERVER=1` (avoids server flakiness under parallelism)

---

## 5) Repository layout (scaffold)

reflaxe.go/
.haxerc
haxelib.json
extraParams.hxml
package.json
vendor/reflaxe/src/ # vendored Reflaxe (optional but recommended)
src/
reflaxe/go/
CompilerBootstrap.hx
CompilerInit.hx
GoCompiler.hx
GoOutputIterator.hx
CompilationContext.hx
GoProfile.hx
ProfileResolver.hx
naming/GoNaming.hx
ast/GoAST.hx
ast/GoASTPrinter.hx
ast/GoASTTransformer.hx
ast/transformers/registry/...
macros/GoInjection.hx
macros/BoundaryEnforcer.hx
macros/StrictModeEnforcer.hx
macros/GoModMetaRegistry.hx # later
macros/GoExtraSrcRegistry.hx # later
runtime/
hxrt/ # Go runtime package emitted/copied into output
hxrt.go
string.go
array.go
dynamic.go
exception.go
...
std/
AGENTS.md
go/
Option.hx
Result.hx
Chan.hx
Go.hx
...
haxe/
... cross overrides as needed
templates/
basic/
compile.hxml
src/Main.hx
test/
README.md
run-snapshots.py
snapshot/
core/...
stdlib/...
regression/...
negative/...
upstream_std_modules.txt


---

## 6) Toolchain pinning (Haxe 4.3.7)

Use the same pattern as reflaxe.rust:

**`.haxerc`**
```json
{
  "version": "4.3.7",
  "resolveLibs": "scoped"
}

package.json

    dev dependency on lix

    postinstall: lix download

    npm test runs harness script

7) Build invocation contract (developer UX)
7.1 HXML defines

    -lib reflaxe.go

    -D go_output=out (required to activate the backend)

    -D reflaxe_go_profile=portable|idiomatic|gopher|metal (optional, default portable)

    -D reflaxe_go_strict (optional)

    -D reflaxe_go_strict_examples (used by repo snapshots/examples)

    -D reflaxe.dont_output_metadata_id (recommended for determinism)

7.2 Output

Given -D go_output=out, compiler writes:

out/
  go.mod
  *.go              # generated files
  hxrt/             # runtime package (copy)

No network access required.
8) Compiler architecture
8.1 CompilerBootstrap (macro)

Copy Elixir’s robustness:

    Detect Go build even under Haxe 4 Cross platform and when defines aren’t visible yet.

    Inject vendor/reflaxe/src.

    If Go build: inject std/ (and optional staged std/_std if you choose to use it).

    Inject at the front of classpath list so overrides win.

8.2 CompilerInit (macro)

    Gate on Go build:

        target.name == "go" OR -D go_output present OR platform is Cross + args contain go_output define.

    Call ReflectCompiler.Start().

    Resolve profile via ProfileResolver.

    Enable boundary macros:

        BoundaryEnforcer.init() (when reflaxe_go_strict_examples)

        StrictModeEnforcer.init() (when reflaxe_go_strict; metal enables by default)

    Register compiler:

ReflectCompiler.AddCompiler(new GoCompiler(), {
  fileOutputExtension: ".go",
  outputDirDefineName: "go_output",
  fileOutputType: FilePerModule,
  targetCodeInjectionName: "__go__",
  ignoreBodilessFunctions: false,
  ignoreExterns: true,
  trackUsedTypes: true,
  expressionPreprocessors: prepasses
});

Start with minimal preprocessors, then add more as needed (mirroring Elixir’s approach).
8.3 GoCompiler (Reflaxe GenericCompiler)

    Generates a Go AST (abstract syntax tree), not strings.

    Stores compiled outputs in classes/enums/typedefs/abstracts arrays as DataAndFileInfo<GoFileAst>.

    Keeps state in CompilationContext:

        profile

        module name

        string policy

        import registry

        type id registry

        runtime features used

8.4 GoOutputIterator

Like Rust:

    For each generated AST:

        GoASTTransformer.transform(ast, context)

        GoASTPrinter.printFile(ast)

        returns DataAndFileInfo<StringOrBytes>

8.5 AST types (GoAST)

Keep it minimal but explicit:

    GoFile (package name, imports, decls)

    GoDecl (type, var, const, func)

    GoStmt (if, for, switch, return, assign, block, etc)

    GoExpr (ident, call, selector, index, literal, unary, binary, composite literal)

8.6 Pass registry (Elixir-style)

Structure:

    ast/transformers/registry/GoASTPassRegistry.hx

    registry/RegistryCore.hx (validate)

    registry/groups/*.hx (group pass bundles)

Support:

    -D go_granular_pass_registry to run granular list (debugging)

    default “lean bundle” list for maintainability

Start passes (minimum):

    NormalizeNames (mangling, keyword escaping)

    CollectImports (build import list)

    RewriteStringOps (pointer-string helpers)

    RewriteVirtualCalls (use __hx_this rules)

    InsertRuntimePrelude (ensure runtime referenced types are imported)

9) Runtime + stdlib strategy
9.1 runtime/hxrt (Go code)

This is where “portable semantics” live.

Minimum runtime modules:

    hxrt/string.go — pointer-string helpers, string concat/eq/compare, Std.string semantics

    hxrt/array.go — Haxe-like Array wrapper or helpers

    hxrt/dynamic.go — Dynamic (any) helpers, type tests

    hxrt/exception.go — throw/catch mapping via panic/recover wrappers

    hxrt/reflect.go — Type/Reflect minimal subset as needed

    hxrt/sys.go — incremental sys support (later)

The compiler should copy runtime/hxrt into out/hxrt (like rust emits hxrt crate).
9.2 std/ overrides (Haxe code)

    Use Haxe modules to:

        match stdlib signatures exactly where required

        hide __go__ injections behind typed APIs

        provide Go-native APIs under std/go/**

Add std/AGENTS.md like rust:

    injection allowed only in std, not in app code

    avoid inline wrappers that leak injections

    keep signatures compatible with eval std coreApi externs

10) Go-native Haxe API surface (opt-in)

Create std/go/** modules that are only available on Go target (#if go or #if reflaxe_go).

Initial set:

    go.Go — go(fn), defer(fn) as macro sugar

    go.Chan<T> — channel wrapper: make, send, recv, select helpers

    go.Result<T> — idiomatic error-return patterns (maps to (T, error) or error plus value struct)

    go.Error — extern mapping to Go error

    go.Slice<T> — wrappers to/from Array<T>

    go.Map<K,V> — wrappers to/from haxe.ds.* as appropriate

Rule: portable code shouldn’t need these; they’re for #if go enhancements.
11) Test harness (deterministic + agent-friendly)

This is the core deliverable for LLM productivity.
11.1 Snapshot layout (required)

test/snapshot/<category>/<case>/
  compile.hxml
  Main.hx (+ other .hx)
  intended/          # golden
  out/               # generated (gitignored)
  expected.stdout    # optional runtime expectation
  expected.error.txt # optional negative expectation

Categories:

    core — language basics

    stdlib — stdlib coverage

    regression — bug repros

    go_native — std/go API cases

    negative — expected compile failures

11.2 Snapshot compile.hxml conventions

Every positive snapshot should include:

    -lib reflaxe.go

    -D go_output=out

    -D reflaxe_go_strict_examples

    -D reflaxe.dont_output_metadata_id

    --no-traces -D no_traces

    -main Main

Example:

-cp .
-lib reflaxe.go
-D go_output=out
-D reflaxe_go_strict_examples
-D reflaxe.dont_output_metadata_id
--no-traces
-D no_traces
-main Main

Negative snapshots are the same but live under test/snapshot/negative/....
11.3 Harness script: test/run-snapshots.py

Implement in Python (best DX for agents + best error reporting). It should:

Discovery

    Find all compile.hxml under test/snapshot/** excluding **/_archive/**.

    Sort cases deterministically.

Commands / flags

    --list

    --category core,stdlib

    --case stdlib/bytes_basic (category/case)

    --pattern <regex>

    --jobs N (optional parallel)

    --timeout SECONDS (default 120)

    --update (copy out → intended on success)

    --runtime (run cases with expected.stdout)

    --failed (re-run last failing set)

    --changed (git diff selection; optional if git exists)

    --chunk i/n (deterministic sharding; hash case path)

Per-case pipeline

    Delete out/

    Run haxe compile.hxml with env HAXE_NO_SERVER=1

    Run gofmt -w on all out/**/*.go

    Run go test ./... inside out/ (no network; fail if tries)

    Diff out/ vs intended/ (exclude go.sum, .cache, _GeneratedFiles.json)

    If --runtime and expected.stdout exists:

        go run . (or go run main.go if generated) with a timeout

        compare stdout exactly (normalize line endings)

Failure reporting

    On Haxe compile fail: print stderr/stdout

    On go test fail: print output

    On diff mismatch: print unified diff for first N files + paths

State files

    .test-cache/last_failed.txt — list of failing case paths

    .test-cache/last_run.json — structured results (passed/failed, durations)

Environment

    Respect KEEP_ARTIFACTS=1 to keep out/ on failure for debugging

    Otherwise delete out/ for passed cases (or keep for speed; but be explicit)

11.4 Minimal test/README.md

Explain:

    how to run all

    how to run one case

    how to update intended

    what strict_examples means

    troubleshooting

11.5 Upstream stdlib sweep (later but planned now)

Like rust:

    test/upstream_std_modules.txt curated list

    test/run-upstream-stdlib-sweep.py compiles each module via a generated dummy Main.hx that imports it

    goal: catch missing std overrides early

12) Milestones (each milestone is “mergeable” + harness-gated)

Each milestone lists:

    deliverables

    required tests to add

    acceptance criteria (what npm test must prove)

Milestone 0 — Repo scaffold + harness foundation (NO real codegen yet)

Deliverables

    .haxerc pinned to 4.3.7

    haxelib.json with reflaxe metadata (name: "Go", abbv: "go", stdPaths: ["std"])

    extraParams.hxml calling:

        --macro reflaxe.go.CompilerBootstrap.Start()

        --macro reflaxe.go.CompilerInit.Start()

    CompilerBootstrap.hx (Elixir-grade front injection)

    CompilerInit.hx (gates on go_output, registers compiler)

    Harness:

        test/run-snapshots.py

        test/snapshot/core/hello/ with compile.hxml and Main.hx

        test/.gitignore ignoring out/

Acceptance criteria

    npm test discovers and attempts snapshot runs.

    It fails with a clear error indicating the compiler has no implementation yet (expected at this milestone), but:

        harness itself works (list, selection flags, timeout wrapper).

    npm run test:list shows the case.

(Why allow failure? Because at Milestone 0 you’re validating harness plumbing; you can mark the single case as “expected fail” or use a --allow-fail flag until Milestone 1.)
Milestone 1 — Minimal Go emission: one file builds

Deliverables

    GoAST + GoASTPrinter that can print:

        package decl

        imports

        func main()

    GoCompiler emits main.go for Main.main with trace("hi") lowered to a runtime print.

    runtime/hxrt minimal:

        hxrt.Println(s *string) or similar

        hxrt.StringFromLiteral

    Compiler copies runtime/hxrt into out/hxrt.

    Harness now:

        gofmt

        go test ./...

        diff intended vs out

Tests to add

    test/snapshot/core/hello_trace/ with expected output.

Acceptance

    npm test passes.

    Generated code gofmt’s clean and go test succeeds.

Milestone 2 — Expressions + primitives (portable semantics)

Deliverables

    Primitives mapping:

        Int (decide int32)

        Float (float64)

        Bool

        String (*string in portable)

    Binary ops:

        arithmetic on Int/Float

        boolean ops

        string concat via runtime helper with Haxe semantics

    Local vars, assignments

    if/else, return

    trace and Std.string basics

Tests

    core/arithmetic

    core/if_else

    core/string_concat_null_semantics

    core/locals_assign

Acceptance

    Snapshots match + go test passes.

    Runtime smoke (optional) for selected tests if expected.stdout exists.

Milestone 3 — Arrays and loops

Deliverables

    Array<T> representation strategy (choose one):

        hxrt.Array[T] wrapper OR raw slices + helper fns

    Implement:

        array literal, indexing, length

        push/pop

    Loops:

        while

        for (i in 0...n)

        for (x in array) (range) when semantics match

    Ensure loop semantics match Haxe (inclusive/exclusive ranges, mutation safety)

Tests

    core/array_basic

    core/array_push_pop

    core/loops_range

    core/loops_array_iter

Milestone 4 — Functions + closures

Deliverables

    Static functions and local functions

    Closures capturing locals

    Default args (if supported by Haxe)

    Varargs mapping

Tests

    core/closures_capture

    core/function_values

    core/varargs

Milestone 5 — Classes (no inheritance yet) + field access

Deliverables

    Class to struct mapping

    Constructor New_<Type> style

    Instance fields

    Instance methods

    Static fields/methods (prefixed names)

    Name mangling for package collisions

Tests

    core/class_fields_methods

    core/static_fields_methods

Milestone 6 — Inheritance + virtual dispatch (the hard one)

Deliverables

    Implement the self-dispatch pattern:

        base struct has __hx_this interface

        constructor wiring sets it correctly

        method calls on base-typed refs dispatch through __hx_this

        super calls bypass dispatch

    Downcast / Std.isOfType minimal

    Interfaces (Haxe interfaces) mapping (can be later if too much)

Tests

    core/inheritance_override_dispatch

    core/super_calls

    core/base_calls_virtual

Acceptance

    These tests must include cases where:

        base method calls an overridden method

        variable typed as base holds derived instance

Milestone 7 — Enums + switch (pattern matching)

Deliverables

    Enum representation (recommend: interface + constructor structs)

    Switch lowering:

        switch on Int/bool/string → Go switch

        switch on enums → type switch or tag switch

    haxe.ds.Option / haxe.functional.Result handled as first-class (optionally via metadata like Elixir did to prevent “tag-only” erasure)

Tests

    core/enum_basic

    core/enum_switch_bindings

    stdlib/option_basic

    stdlib/result_basic

Milestone 8 — Exceptions (throw/try/catch)

Deliverables

    Map Haxe throw/catch to panic/recover wrapper functions in runtime:

        compiler wraps try blocks in a func() { defer ... }()

        catch filters by type where possible

    haxe.Exception subset

Tests

    core/try_catch_basic

    core/throw_custom

    stdlib/haxe_exception_smoke

Milestone 9 — Stdlib ladder (high-signal subset)

Deliverables
Prioritize the same kinds of modules Rust target lists in upstream_std_modules.txt:

    haxe.Json

    haxe.io.Bytes, BytesBuffer, Input, Output

    haxe.ds.IntMap, StringMap, ObjectMap, EnumValueMap, List

    Sys, sys.io.File, sys.io.Process (incremental)

Add test/upstream_std_modules.txt and sweep script.

Tests

    stdlib/json_parse_stringify

    stdlib/bytes_basic

    stdlib/intmap_basic

    sys/file_read_write_smoke (if sys is started)

Milestone 10 — Go-native surface (profiles)

Deliverables

    Add profiles (Rust-style):

        portable default

        idiomatic same semantics, cleaner output

        gopher Go-first helpers allowed

        metal strict injection boundary + typed “interop facade”

    Implement ProfileResolver with conflict diagnostics

    Add std/go/** modules:

        channels + goroutines

        error-style result

    Add strict mode default for metal

Tests

    go_native/channel_basic

    go_native/goroutine_smoke

    go_native/error_result_mapping

Milestone 11 — Polishing: CI stability + docs for agents

Deliverables

    AGENTS.md at repo root (Rust-style guardrails)

    std/AGENTS.md (Rust-style injection rules)

    Add “registry validation” tests for pass registry uniqueness/deps

    Add chunked harness mode for CI parallelism

Acceptance

    npm test stable

    npm run test:update works

    npm run test -- --failed works

    npm run test -- --chunk 0/4 works deterministically

13) Agent instructions (files to include verbatim)
Root AGENTS.md (template)

Include:

    “AST-first pipeline: Builder → Passes → Printer”

    “No raw Dynamic unless unavoidable; contain it”

    “No absolute paths in output”

    “Always add a regression test when fixing a bug”

    “How to run snapshots / update intended”

    “Policy: apps/examples must not use __go__”

std/AGENTS.md (template)

Include:

    injections allowed only in std/runtime

    avoid inline injection wrappers leaking

    keep signatures compatible with std coreApi externs

14) Open questions (explicitly defer; don’t block Milestone 1–3)

    Whether to use generics heavily in runtime (requires Go 1.18+). Recommendation: keep runtime mostly non-generic early; add generics later.

    Whether to generate go.mod always or only in tests. Recommendation: always generate minimal go.mod deterministically, but exclude go.sum from snapshots.

    Whether to emit multi-package output. Recommendation: postpone until after stdlib ladder is stable.


If you want one extra “agent superpower” baked into the harness: add a `--bless` mode that updates `intended/` *only for files that changed*, and prints a short “why it changed” checklist (gofmt changes? naming changes? runtime changes?). That reduces accidental snapshot churn and makes LLM commits much cleaner.
::contentReference[oaicite:0]{index=0}
