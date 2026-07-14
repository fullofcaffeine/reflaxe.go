# GPT-5.6 Pro deep architecture and production-readiness review

Prepared for the independent Haxe.Go architecture, bounded-production, generated-Go, Go-native authoring, release, and stable-1.x audit of source commit `cd79624f855521dbf320ac2b7524d889ca388c0e`.

Paste the complete **Review prompt** section into GPT-5.6 Pro and attach the exact artifact listed under **Upload checklist**. Do not substitute a working-tree archive or one of the older untracked Repomix files.

## Review prompt

### Deep architecture, production-readiness, and stable-1.x audit: Haxe.Go (`reflaxe.go`)

Act as an independent senior compiler architect and production-readiness reviewer with expertise in:

- Haxe 4.3.7 compiler semantics, macros, typed AST transformations, and standard-library contracts;
- Reflaxe compiler architecture and target standard-library conventions;
- Go language semantics, type design, interfaces, generics, packages, runtime behavior, and tooling;
- the Go memory model, goroutines, channels, `select`, synchronization, cancellation, races, and resource lifecycle;
- generated-code readability, performance, allocation behavior, debuggability, and operational use;
- language-runtime and target-standard-library implementation;
- typed foreign-function interfaces, code generation, binding generation, and macro DSL design;
- SemVer, release engineering, artifact provenance, GitHub Releases, and software supply-chain security;
- compatibility policy and production adoption of cross-language compilers.

This is an adversarial, evidence-first review. Do not implement changes. Inspect the supplied exact source archive, line-numbered source view, generated outputs, tests, reports, workflows, live-state captures, roadmap, and pinned sibling precedents. Return an actionable disposition that can be adjudicated into Beads.

Do not merely restate repository documentation or the first-pass audit. Try to disprove both. Identify false positives, missing evidence, competing designs, and cases where qualification or exclusion is better than implementation work.

#### Primary objective

Make seven separate judgments about whether Haxe.Go has:

1. credible **bounded production use now**, including whether “beta-stable” is defensible and under exactly what conditions;
2. enough contract definition and evidence for a **stable 1.x compatibility promise**;
3. a sound and maintainable **compiler architecture**;
4. a sound **portable-specialization direction** that preserves Haxe portability while lowering to native Go representations when proven safe;
5. a sufficiently complete and coherent **Go-native authoring experience** for a Go developer who chooses the typed Go layer;
6. credible **generated-Go output quality** relative to readable, idiomatic, hand-written Go with equivalent semantics;
7. trustworthy **release and distribution integrity**, including SemVer, tested-commit identity, artifacts, provenance, recovery, and GitHub host controls.

These judgments are intentionally independent. Do not average them into one readiness score. A compiler can be useful for bounded production workloads while not being ready for stable 1.x, and it can have sound product direction while still having inadequate release machinery or native authoring coverage.

Stable 1.0 is not urgent. A `NOT_READY` stable-1.x verdict is acceptable. Conversely, do not require universal Haxe parity, every Go package, every Go syntax form, every operating system, or perfect generated Go before permitting a carefully bounded beta or stable contract.

For every gap, decide whether it must be:

- fixed before any credible bounded production use;
- fixed for the validated beta-baseline milestone;
- fixed only if the affected surface is admitted to stable 1.x;
- resolved by explicitly qualifying, deprecating, or excluding that surface;
- placed in optional longer-term product work;
- rejected as a false positive or already-covered duplicate.

Avoid turning the audit into a feature wishlist or an aesthetic rewrite plan.

#### Evidence authority and exact identity

Primary repository and product:

- Product: Haxe.Go
- Public repository: `fullofcaffeine/reflaxe.go`
- Repository URL: https://github.com/fullofcaffeine/reflaxe.go
- Reviewed source commit: `cd79624f855521dbf320ac2b7524d889ca388c0e`
- Reviewed branch containment at bundle construction: `origin/master`
- Target: Haxe 4.3.7 to Go through Reflaxe

Canonical review artifact:

- Filename: `haxe-go-gpt56-evidence-cd79624f.zip`
- Size: 13,502,320 bytes
- SHA-256: `ab4b0a1097229ad2202ca7da6c092b2b85cba537522951fb625f6b1312c4b511`
- Builder commit: `531e5bc1368c2b72a41cd5b46fc2449e7ed90393`
- Builder-file SHA-256: `447be705cb2a4ad73116f83256dcac125d2b4e745d195584b5819a8dd4272a78`

The Git archive is the source authority. `primary/haxe.go-source-cd79624f.tar` is an exact `git archive` of the reviewed commit after the two documented operational Beads files were excluded. `primary/haxe.go-source-cd79624f.xml` is a line-numbered Repomix navigation view, not a second source authority.

Repomix's security scanner omitted eight tracked TLS/proxy fixtures from its XML view. The paths are listed in `primary/repomix-security-exclusions.json`; the exact raw files are present under `primary/repomix-security-exclusions/` and in the authoritative Git archive. The whole assembled payload passed Gitleaks. Do not misdiagnose those eight XML omissions as missing repository files.

The bundle was built twice with byte-identical output. `unzip -t` passed, 379 internal `SHA256SUMS` entries verified, the full bundle passed Gitleaks, and 375 UTF-8 payload files passed the developer-machine workspace-path check.

The archive does not contain `.git`. Commit identity is therefore supplied by the recorded builder checks, archive inventory, remote-branch containment, cryptographic payload hashes, and exact-SHA CI metadata. If you independently browse live GitHub state, label that as separately dated live verification rather than silently mixing it with the frozen bundle.

#### Current CI, release, and host evidence

The following GitHub Actions runs all report `success` at the exact reviewed source commit:

- Security - Static Analysis: https://github.com/fullofcaffeine/reflaxe.go/actions/runs/29294117737
- Examples Artifacts: https://github.com/fullofcaffeine/reflaxe.go/actions/runs/29294117742
- CI - Quality: https://github.com/fullofcaffeine/reflaxe.go/actions/runs/29294117743
- Security - Gitleaks: https://github.com/fullofcaffeine/reflaxe.go/actions/runs/29294117745
- CI Harness: https://github.com/fullofcaffeine/reflaxe.go/actions/runs/29294117785

Treat green CI as evidence of the exercised contracts, not proof of correctness, production readiness, release readiness, or coverage completeness.

Live GitHub state captured during bundle construction reports:

- current release tag: `v0.53.1`;
- tag commit: `4adb9ae8209c24850495110e79ba6c5a8e1fa2bd`;
- GitHub release assets: zero;
- GitHub release immutability: `false`;
- repository rulesets: none;
- default branch: `master`;
- default-branch protection API: HTTP 404, branch not protected.

The release tag predates the reviewed source commit. Do not treat `v0.53.1` as a shipped artifact of `cd79624f…`; no such release binding is claimed.

## Known facts

The following are frozen facts or mechanically reported observations in the supplied evidence. They are not, by themselves, readiness conclusions:

1. The authoritative source inventory contains 5,368 included tracked files, 712 Haxe source files, 68 checked-in `.cross.hx` files, 142 committed generated-metal files, 195 committed generated-portable files, 3,624 intended snapshot files, and five workflow files.
2. `src/reflaxe/go/GoCompiler.hx` is 14,825 lines at the reviewed source commit.
3. A mechanical textual inventory reports roughly 5,108 `GoRaw` occurrences under `src/reflaxe/go/**`; approximately 2,279 are in `GoCompiler.hx`, with large additional clusters in specialized emitters. This is a debt locator, not proof that every occurrence is wrong.
4. The repository documents and tests two principal profiles named `portable` and `metal`, plus orthogonal strictness, native-boundary, optimization, and runtime-policy controls.
5. The default profile contract is portable. Metal is an explicit Go-native build contract; it is not documented as the only route to good or idiomatic Go.
6. Current source contains staged-standard-library material, target support types, compiler-owned std shims, and 68 checked-in `.cross.hx` files, but it is not yet in the canonical planned `std/go/_std` ordinary-`.hx` layout.
7. The current typed Go surface includes some `go.*` types and typed extern metadata such as `@:go.import`, `@:go.name`, and `@:go.receiver`; it does not claim complete typed coverage of Go language and standard-library features.
8. The documented `go.Select` helpers are deterministic, branch-priority polling helpers. They are not a typed lowering of Go runtime pseudo-random ready-case `select` semantics.
9. Portable result semantics, native `go.Result<T>`, and Go `(T, error)`/multi-result behavior are distinct contracts and must not be silently conflated.
10. Package manifests currently say `0.1.0`, while the latest GitHub tag is `v0.53.1`. The current semantic-release configuration mutates tracked changelog/version files through a release commit and publishes no verified Haxelib package asset in the captured release.
11. The reviewed workflows exercise Haxe 4.3.7, Node 20, and Go 1.22/1.23-era lanes. The dependency audit intentionally classifies reachable Go standard-library vulnerability reports as visible but non-blocking.
12. The first-pass audit reports 231/231 snapshots, a strict upstream-stdlib sweep of 55/55, a semantic report of 126/126 plus 3/3 lane cases, and 175 portable modules accounted for. Verify what each count proves; compile coverage and snapshot agreement are not semantic parity.
13. The roadmap snapshot contains 663 issues rooted at `haxe_go-vfp`, at Dolt commit `d4n7bp04tc9k1srifgn0ipbb8thkgld9`, with zero dependency cycles at capture time.
14. The roadmap already contains concrete work for release, `_std`, compatibility, portable specialization, typed IR, native Go DX, runtime/concurrency, output quality, examples, and beta closure. Closed or open Beads are planning evidence, not proof that a finding is resolved.

## Open hypotheses

The first pass proposed the following hypotheses. Confirm, narrow, or reject each one with direct evidence:

1. Haxe.Go is a credible bounded-production 0.x beta for pinned, tested workloads, but “stable” without the word “beta” would currently overstate distribution, compatibility, and native-DX maturity.
2. Haxe.Go is structurally closer to Haxe.Rust than to Haxe.Ruby or Haxe.Elixir because Go and Rust both benefit from typed native representations, runtime minimization, portable/metal boundaries, and generated systems-language output. This may still be misleading because Go has garbage collection, interfaces, structural conventions, a different error model, and a different concurrency/runtime contract.
3. Haxe.Rust's typed usage ledger, surface-contract registry, runtime requirement analysis, and portable-specialization gates are transferable design precedents, while Rust ownership, borrow, `Send`/`Sync`, Cargo, and `no_hxrt` mechanics are not transferable designs by default.
4. The 68 checked-in `.cross.hx` files and mixed std/support layout should migrate to canonical `std/go/_std` ordinary `.hx` overrides plus separately owned target support, with deterministic `.cross.hx` generated only where a modern Reflaxe packaging boundary truly requires it.
5. The size of `GoCompiler.hx` and the `GoRaw` concentration create concrete correctness and change-risk, but the correct response is typed ownership seams and debt ratchets rather than a cosmetic file split or a ban on every raw leaf.
6. Portable lowering should become progressively more Go-native through typed usage plus a versioned semantic registry, not through profile selection or optimistic global representation substitution.
7. The current `go.*` facade is too thin for the north-star claim that a Go developer can write Haxe as if writing Go with better Haxe ergonomics and macros.
8. Current generated output is well covered by snapshots but insufficiently compared with equivalent hand-written Go for readability, allocations, binary size, runtime, race behavior, and operational debugging.
9. Current release automation is not a trustworthy same-tested-SHA Haxelib/GitHub release system because version authority, deterministic package construction, isolated package testing, hosted assets, immutable release state, repair semantics, and host enforcement are incomplete.
10. Haxe 4.3.7 typed APIs and macro lifecycle features can reduce reflection, `Dynamic`/`Any`, boilerplate, and stringly metadata handling without turning the project into a syntax-modernization exercise.

For each hypothesis, say whether it is `confirmed`, `partly confirmed`, `rejected`, or `insufficient evidence`, and explain why.

#### Intended product boundary

The product north star is: Haxe.Go should be the best way to write Go when not writing raw Go, while remaining a first-class Haxe target for code that must also compile to Rust, Ruby, Elixir, JavaScript, or other supported Haxe targets.

The governing profile rule is:

> Portable is the default product path; Go-native by opt-in; metal-like generated Go whenever the compiler can prove that the lowering preserves portable Haxe semantics.

Interpret that rule precisely:

- Portable is for normal Haxe and Haxe-standard-library authoring with an evidence-backed supported subset. It is not the “slow/basic” mode.
- Metal is an explicit Go-native authoring and fail-fast boundary contract. It is not the only “real Go” mode.
- A source tree may keep domain/shared code portable and isolate Go-native infrastructure behind typed adapters or `@:goMetal` lanes.
- One compilation selects one principal profile, but profile choice must remain orthogonal to strictness, optimization, runtime slicing, native-import policy, and evidence level.
- Portable output should use native slices, maps, strings, closures, errors, or other Go shapes whenever typed proof establishes Haxe-semantic preservation and a safe fallback exists.
- Go-native users should get typed access to Go language and library capabilities without raw string injection for ordinary work.
- Haxe developers should be able to remain portable without understanding Go internals for every operation.
- Go developers should be able to recognize generated code, diagnose it with standard Go tools, and opt into Go-native APIs with Haxe type checking and macro ergonomics.

The desired authoring layers are:

1. portable Haxe and supported upstream Haxe stdlib semantics;
2. portable source with compiler-proven Go-native representations;
3. narrow typed Go adapters inside otherwise portable applications;
4. explicit metal-profile Go-native applications;
5. controlled framework/runtime escape islands only where typed Haxe or externs cannot express the boundary.

Assess whether this layering is coherent, teachable, mechanically enforceable, and actually reflected in compiler behavior and examples.

#### Non-negotiable architectural principles to assess

Do not recommend violating these merely to make a test pass:

1. Well-typed Haxe is the source-language contract.
2. Compiler design is AST-first: typed builder/lowering, transform passes with declared invariants, then printer/output.
3. Prefer compile-time lowering when typed AST, metadata, types, or literals provide a closed answer.
4. Prefer Reflaxe framework capabilities to duplicated local machinery when the framework abstraction actually fits the Go target.
5. Put library-expressible Haxe standard-library behavior in staged target `_std`; use thin `hxrt` helpers for real Go-side runtime support; keep compiler-owned std shims only for compile-context-sensitive behavior.
6. Do not grow large stdlib reimplementations as `GoStmt.GoRaw` or `GoExpr.GoRaw` blocks in the compiler when staged source or a narrow runtime helper is the correct owner.
7. Avoid `Dynamic`/`Any` in compiler, macros, facades, and runtime boundaries whenever an explicit type is possible. Localize semantically unavoidable dynamic values and justify them.
8. Portable native representation requires compiler-observed typed usage, a versioned semantic contract, eligible shapes, explicit fallback, runtime/import accounting, and semantic-difference proof. Profile selection alone is not proof.
9. Generated Go is a first-class output: `gofmt`-clean, understandable, reasonably idiomatic, tool-friendly, deterministic, and close to hand-written Go performance when Haxe semantics allow.
10. The default Go-native integration mechanism is typed externs/facades. Use macros only for genuine Haxe syntax gaps such as native `select` clauses, `defer`/`go` statements, multi-result destructuring, comma-ok forms, struct tags, or build constraints.
11. Use macros only when needed: macros must not hide ordinary APIs, weaken types, emit uninspectable strings, or create a second compiler inside the facade layer.
12. Application, test, and example code must not be taught to use raw `__go__`, generated-file edits, or `Dynamic` workarounds for compiler defects.
13. Framework-owned `__go__`/`@:goAllowRaw` islands may exist only as narrow, audited target layers after typed externs are insufficient; imports remain typed and explicit.
14. Portable Haxe result semantics, native `go.Result<T>`, and Go multi-result/`error` conventions stay distinct unless an explicit typed conversion is requested.
15. Examples are executable QA and layered contract documentation, not decorative snippets.
16. Fix compiler/runtime root causes and add focused regressions; do not recommend application-specific source contortions.

If one of these principles is unsound, provide a concrete counterexample and the smallest replacement principle. Do not reject a principle merely because the current implementation does not yet satisfy it.

#### Sibling precedent and transfer discipline

The evidence bundle includes narrow, commit-pinned sibling slices:

- Haxe.Rust: `5b8c9416f963e541229e633a2bb655a93e3e9c16` — portable specialization, runtime requirement analysis, typed surface contracts, `_std`, and release architecture;
- Haxe.Ruby: `08faba040457165b883ae5327315581979ea07db` — SemVer transition and release mechanics;
- Haxe.Elixir: `68625fa91ffff48c5ffb269bff01c6f3e716128c` — canonical target `_std` source/package layout and release mechanics.

These are precedent, not proof that a sibling design belongs in Haxe.Go. They are not source dependencies and do not establish Haxe.Go correctness. Use them to answer:

- which patterns are universal Reflaxe/backend architecture;
- which are target-family conventions;
- which are Go-specific and should be designed independently;
- whether the “Haxe.Go is closest to Haxe.Rust” hypothesis is useful or distorting;
- whether a simpler Go design is possible because the Go runtime, GC, interfaces, and tooling solve different problems.

Do not penalize Haxe.Go for failing to reproduce a Rust-specific ownership mechanism, Ruby packaging convention, or Elixir runtime model.

#### Compatibility and SemVer contract to inspect

Review the entire consumer-visible contract, not only importable Haxe classes:

- public Haxe types and public members;
- `go.*` facades, typed externs, generated bindings, macros, and transitive public types;
- metadata names, argument grammar, and defaults;
- compiler defines and profile-selection behavior;
- diagnostics, source positions, warning/error policy, and exit behavior;
- JSON/Markdown report schemas and filenames;
- generated package/module/file layout, symbol naming, build tags, and `go.mod` behavior;
- `hxrt` import path and promised runtime slicing/no-runtime behavior;
- packaged class paths, target std paths, vendored Reflaxe, executable entrypoints, and installation flow;
- supported Haxe, Go, Node, OS, and tool versions;
- deprecation and migration rules;
- release tags, Haxelib version, GitHub release identity, and artifact/provenance formats.

Determine which units deserve one of these dispositions:

- Admit to stable 1.x;
- Admit with explicit qualification;
- Keep experimental during 1.x;
- Exclude/internalize;
- Require more evidence before deciding.

Do not freeze generated Go implementation detail merely because snapshots exist. Conversely, do not call an importable, relied-upon surface “internal” solely because maintainers wish users had not imported it.

Review pre-1.0 SemVer conventions separately from stable-1.x admission. Check breaking changes hidden as “additive,” including enum/abstract cases, metadata fields, report fields, generated paths, changed defaults, diagnostics, Go interface method sets, and transitive types.

#### Release architecture to evaluate

The desired release invariant is:

```text
one version authority
→ exact commit passes release gates
→ deterministic Haxelib ZIP built twice from that commit
→ isolated install compiles, go-tests, and runs portable and metal representatives
→ tag points to that exact tested commit
→ GitHub Release is immutable and contains ZIP, checksum, manifest, and provenance
→ hosted bytes and source identity are verified
```

An interrupted release repair path should accept an existing immutable version tag, recover or deterministically rebuild only the artifact for that tag, and complete missing hosted state. It must not derive a new version, tag an arbitrary branch, move a tag, overwrite different bytes, or silently publish a different source identity.

Assess whether the current semantic-release design, tracked release commit, version files, Haxelib conventions, GitHub assets, release notes, permissions, action pinning, tag protection, branch protection, rulesets, and recovery behavior satisfy that invariant or should be replaced. Prefer the smallest correct release architecture; do not add ceremony without a threat or failure model.

#### Evidence bundle map and files to prioritize

Start with:

- `MANIFEST.json`, `README.md`, and `SHA256SUMS` at the bundle root;
- `primary/source-inventory.json` and `primary/repomix-security-exclusions.json`;
- `roadmap/haxe-go-next.json`;
- `live/github/host-controls.json`, release/tag metadata, repository metadata, and exact-SHA CI summaries/logs;
- `references/reference-manifest.json` and the line-numbered sibling reference view.

In the primary source, prioritize:

- `AGENTS.md` and `README.md`;
- `package.json`, `package-lock.json`, `haxelib.json`, and `.releaserc.json`;
- `.github/workflows/**` and `.github/actions/**`;
- release, security, CI, package, and artifact scripts;
- `docs/profiles.md`, `docs/profile-semantics-guide.md`, `docs/portable-canonical-contract.md`, and `docs/portable-semantics-v1.md`;
- `docs/feature-support-matrix.md`, `docs/known-gaps.md`, `docs/ownership-rubric.md`, and stdlib provenance/migration documents;
- `docs/release-readiness-checklist.md`, `docs/release-visibility.md`, and `docs/security-dependency-audit.md`;
- `src/reflaxe/go/GoCompiler.hx` and `src/reflaxe/go/GoReflaxeCompiler.hx`;
- `src/reflaxe/go/ast/**`, `compiler/**`, `analyze/**`, `macros/**`, and bootstrap/profile code;
- `src/go/**` and `std/go/**` typed Go surfaces;
- `std/**`, especially `.cross.hx`, `_std`, `sys`, thread, network, SSL, and IO ownership;
- `runtime/hxrt/**`;
- `vendor/reflaxe/**`, including patch/provenance notes;
- semantic-difference fixtures and reports;
- compiler snapshots and representative intended Go;
- portable, metal, mixed-boundary, concurrency, interop, network, and flagship examples;
- performance baselines and generated-output telemetry;
- raw-injection, strictness, profile, runtime-slicing, and ownership gates.

Documentation states intended contracts; it is not proof. Generated snapshots establish exact current shape, not automatically semantic correctness or idiomatic quality. Prefer implementation and executable evidence when they disagree.

#### Required audit dimensions

##### 1. Product definition, beta claim, and public wording

Determine exactly what can be built safely today and whether “beta-stable,” “production-capable beta,” “experimental,” or another label is truthful.

Separate:

- compile-covered behavior;
- snapshot-covered output;
- semantically compared behavior;
- runtime-tested behavior;
- race/static/security-tested behavior;
- release-installed behavior;
- public stable-contract behavior;
- aspirational north-star behavior.

Identify overly broad or unnecessarily weak claims. Propose exact replacement wording for the README and release description.

##### 2. Haxe semantic correctness and portability

Inspect likely semantic boundaries for:

- nullability, optional arguments, default values, zero values, and typed nil;
- integer overflow, unsigned shifts, floats, `Int64`, and target numeric differences;
- strings, Unicode/UTF-16 expectations, bytes, encodings, paths, and OS strings;
- arrays, maps, lists, iterators, mutation, aliasing, and copy/reference semantics;
- classes, constructors, inheritance, interfaces, overriding, `super`, accessors, and dispatch;
- closures, captures, loop variables, evaluation order, temporaries, and side effects;
- anonymous structures, reflection, `Dynamic`, `Any`, and runtime casts;
- enums, enum abstracts, abstracts, generic constraints, monomorphization, and pattern matching;
- exceptions, typed catches, nested catches, rethrow, stack behavior, and cross-goroutine propagation;
- resource embedding, RTTI, serialization, templates, JSON, regex, XML, and crypto;
- Haxe stdlib and `sys.*` behavior;
- portable behavior across sibling Haxe targets.

Distinguish a deliberate, documented target difference from an accidental semantic regression. Distinguish module compile coverage from member/operation parity.

##### 3. Compiler and Reflaxe architecture

Assess:

- AST-first discipline and the authority of raw text;
- typed IR coverage for Go expressions, statements, declarations, types, imports, packages, interfaces, errors, and multi-results;
- lowering ownership and type propagation;
- pass responsibilities, ordering, prerequisites, determinism, and observability;
- Reflaxe API use versus copied or reflective local machinery;
- compiler bootstrap and class-path mutation;
- module/path/name/import planning and collision handling;
- report and generated-artifact ownership;
- diagnostic anchoring and error behavior;
- hidden global state or ordering dependence;
- decomposition seams and change-risk concentration.

Classify `GoRaw` occurrences by purpose: unavoidable leaf syntax, missing typed IR, misplaced stdlib/runtime behavior, test fixture, printer escape, or untrusted/stringly input. Do not equate a raw-node count with a bug count.

Determine whether the 14,825-line `GoCompiler.hx` is merely inconvenient or a concrete correctness/change-risk problem. If decomposition is warranted, propose dependency-oriented seams protected by characterization tests, not a cosmetic file split.

##### 4. Canonical Reflaxe `_std`, package layout, and ownership

Review whether the target follows modern Reflaxe conventions and `reflaxe new` expectations:

- canonical `_std` selection and target path;
- ordinary `.hx` staged source versus generated `.cross.hx` artifacts;
- Haxelib `classPath`, `reflaxe.stdPaths`, install layout, and package runner;
- bootstrap configuration without reflective class-path mutation;
- source ownership and provenance;
- separation of upstream semantic overrides, target support, public Go facades, and runtime internals;
- isolated installed-package compilation;
- sibling-family layout consistency where useful.

For representative stdlib helpers, decide whether the correct owner is upstream Haxe source unchanged, staged target `_std`, thin `hxrt`, compiler context, typed Go extern/facade, or a narrow framework-owned raw island.

Propose a safe migration order that preserves selection semantics and package behavior. Identify any `.cross.hx` that should remain generated and explain the Reflaxe contract that requires it.

##### 5. Portable specialization and profile semantics

Assess whether the portable/metal model is coherent and whether the compiler can safely converge portable output toward metal-like/native Go.

Review:

- typed usage tracking;
- eligibility analysis and type-shape closure;
- a versioned Go surface-contract registry;
- fallback representation and rollback;
- runtime and import consequences;
- deterministic reports explaining decisions;
- semantic-difference proof;
- strings, bytes, arrays/slices, maps, iterators, closures, `Option`, and result/error surfaces;
- profile/native-policy/strictness/optimizer/runtime orthogonality;
- accidental behavior changes hidden behind optimizer presets.

Judge whether the pinned Haxe.Rust pattern is the right precedent, a partial precedent, or unnecessary complexity for Go.

##### 6. Go-native typed authoring experience

Evaluate whether a Go developer can express ordinary Go designs through Haxe types and recognizable APIs without abandoning Haxe safety.

Inspect current and missing support for:

- packages, imports, aliases, visibility, receivers, methods, and method sets;
- values, pointers, zero/nil, structs, embedding, tags, interfaces, type assertions, and type switches;
- slices, arrays, maps, `make`, `new`, capacity, append/copy/delete, and comparability;
- Go generics and constraints where Haxe can model them coherently;
- `(T, error)`, multiple results, comma-ok, blank identifiers, and typed destructuring;
- `defer`, `go`, native `select`, channel directions, close, range, and cancellation;
- `context.Context`, `sync`, atomics, timers, and lifecycle patterns;
- standard-library facades and third-party package binding generation;
- build constraints, platform files, module paths, and CGo/unsafe boundaries;
- documentation and examples for portable, mixed, and metal applications.

Assess the current `go.Result`, channel, goroutine, and priority-polling APIs. Determine what should be a normal typed API, a macro DSL, generated extern binding, compiler syntax lowering, or deliberately unsupported surface.

Macros only for genuine Haxe syntax gaps. Require every proposed DSL to preserve types, evaluation order, source positions, import ownership, deterministic expansion, and readable generated Go.

##### 7. Go runtime semantics, memory model, and concurrency

Audit compiler output and `hxrt` for:

- Go memory model compliance and happens-before relationships;
- races, atomicity, visibility, and unsafe publication;
- mutex/condition/semaphore correctness, lost wakeups, reentrancy, and lock ordering;
- goroutine identity, lifecycle, leaks, panic propagation, and cleanup;
- channels, closed-channel behavior, send-on-closed behavior, receive comma-ok, blocking, and fairness;
- native `select` versus deterministic polling semantics;
- event loop and thread-pool shutdown, wakeups, backpressure, and task accounting;
- timers, deadlines, cancellation, contexts, and resource release;
- cross-goroutine exceptions and recover boundaries;
- `unsafe.Pointer`, reflection, interface representation, and checkptr-sensitive behavior;
- maps, concurrent mutation, shared slices, aliasing, and object identity;
- sockets, TLS, processes, files, and cleanup on error.

Use race detector, `-race`, `checkptr`, stress, schedule perturbation, and targeted failure injection where proportionate. Do not label every `panic` or unsafe operation a defect; determine reachability through admitted surfaces and the intended Haxe/Go failure contract.

Explicitly identify any credible P0/P1 data race, deadlock, data loss, resource leak, process crash, security compromise, or undefined/unsafe behavior.

##### 8. Errors, exceptions, panic/recover, and failure contracts

Review whether normal Go errors, Haxe exceptions, process exits, EOF, timeouts, cancellation, partial IO, and programmer invariants are kept distinct.

Assess:

- conversion between Go `error` and Haxe-visible failure;
- `go.Result<T>` and multi-result extern metadata;
- preservation of original error identity/message/cause where promised;
- `panic/recover` ownership and typed catch behavior;
- panics in background goroutines;
- double panic or panic during cleanup;
- defer ordering and return-value interactions;
- network/process/filesystem wrappers that swallow or reshape errors;
- sentinel zero/empty returns that conceal failure;
- generated code paths that can panic on ordinary user input.

Recommend exact stable qualifications or compile-time exclusions when a surface is not ready.

##### 9. Generated-Go shape, idiom, performance, and operability

Inspect representative portable, metal, mixed-boundary, nested-module, no/minimal-runtime, generic, exception, dynamic, collection, concurrency, network, and interop outputs.

Assess:

- `gofmt` cleanliness and normal Go formatting;
- `go test`, `go vet`, staticcheck, race detector, CodeQL, and build-diagnostic cleanliness;
- readable control flow and names;
- package organization, imports, initialization, visibility, and import-cycle risk;
- idiomatic errors, interfaces, zero values, contexts, channels, and resource ownership;
- excessive closures/IIFEs, temporaries, type assertions, `any`, pointers, reflection, and `hxrt` calls;
- unnecessary allocations, copies, boxing, map-backed objects, and escape-to-heap behavior;
- string/byte conversions and Unicode costs;
- binary size, startup, throughput, latency, memory, and generated-code size;
- source positions, stack traces, Haxe-to-Go debugging, profiling, and incident diagnosis;
- whether a Go engineer can safely inspect and patch the source-level cause rather than editing generated files.

Compare equivalent semantics with a committed hand-written-Go corpus. Do not demand identical syntax or benchmark non-equivalent behavior. Identify root causes in lowering, representation, planning, or runtime design.

##### 10. Runtime scope and capability slicing

For each material `hxrt` subsystem, ask:

- is it genuine dynamic/stateful/platform runtime behavior;
- could the typed compiler or staged std have emitted the answer directly;
- is the API narrowly typed and internal;
- are dependencies, imports, and copied files minimal and deterministic;
- does typed usage actually drive capability selection;
- can unused runtime features be proven absent;
- do portable and metal reports truthfully describe runtime consequences;
- has `hxrt` become a convenience library or second standard library.

The goal is not “zero runtime at all costs.” The goal is the smallest runtime consistent with Haxe semantics and useful Go output.

##### 11. Standard library, systems, network, and platform behavior

Review admitted behavior for:

- files, directories, paths, environment, commands, and processes;
- partial reads/writes, seek, EOF, buffering, encodings, and descriptor cleanup;
- sockets, DNS, HTTP, proxies, redirects, TLS, certificates, SNI, deadlines, and cancellation;
- threads, mutexes, conditions, semaphores, event loops, and pools;
- atomic values and shared object state;
- DB compile/runtime boundaries;
- platform differences and build tags;
- Go failure conversion to Haxe semantics;
- deterministic dependency and import selection.

Do not require broad network/TLS/DB support if the qualification is truthful. Flag misleading breadth or missing production-critical behavior inside the declared lane.

##### 12. Interop, raw authority, binding generation, and input safety

Review:

- typed extern metadata grammar and validation;
- package import/name/receiver mapping;
- multi-result and `error` mapping;
- raw `__go__`, `GoInjection`, and `@:goAllowRaw` containment;
- macro-generated syntax and imports;
- external-package binding generation and version identity;
- Go module ownership and dependency conflicts;
- extra sources, symlinks, path traversal, package collisions, build tags, and generated filenames;
- untrusted metadata, struct tags, raw fragments, package paths, JSON reports, and archives;
- public access to helpers claimed to be internal;
- CGo, `unsafe`, and platform-specific escape boundaries.

Prefer typed, inspectable contracts. Bounded raw authority can be a legitimate internal or experimental escape hatch, but ordinary application and example code must not require it.

##### 13. Haxe 4.3.7 and macro/compiler modernization

Determine whether the code uses modern Haxe 4.3.7 capabilities proportionately:

- typed macro APIs and supported lifecycle hooks;
- typed initial compiler/Reflaxe configuration;
- enum abstracts or typed values for metadata, defines, diagnostics, profiles, surface IDs, and report enums;
- package-scoped null-safety opportunities;
- typed expression helpers and exhaustive matches;
- warning cleanup and a warning ratchet;
- elimination of avoidable reflection and `Dynamic`/`Any`;
- reusable macro helpers without over-abstraction.

Prioritize correctness, types, and less code. Do not propose syntax churn or clever macros without a concrete maintenance or safety benefit.

##### 14. Stable API, SemVer, and migration policy

Assess whether the project can mechanically inventory and classify its public contract, protect it against drift, and make truthful pre-1.0 and stable-1.x promises.

Review:

- one version authority across tag, Haxelib, package metadata, changelog, and reports;
- Conventional Commits and pre-major breaking-change behavior;
- deprecation windows and aliases;
- compatibility manifest coverage of public/transitive units;
- experimental surface containment;
- stable generated-artifact and report schemas;
- changed profile defaults or optimizer behavior;
- toolchain/platform floor changes;
- package layout and class-path changes;
- exact migration guidance and automated checks.

Propose the minimum coherent and useful stable-1.x contract. A small contract is acceptable only if it is useful and does not hide publicly relied-upon behavior.

##### 15. Toolchains, platforms, dependencies, security, and licensing

Review:

- Haxe 4.3.7 pin and update policy;
- supported Go minimum/current policy and fresh dependency resolution;
- Node/npm policy and lockfile use;
- Linux/macOS/Windows coverage and public wording;
- Go standard-library vulnerability policy and `govulncheck` fail-open behavior;
- npm audit, CodeQL, Gitleaks, dependency review, and action pinning;
- workflow permissions, untrusted forks/PRs, caches, artifacts, and credentials;
- compiler/runtime denial-of-service and malicious input paths;
- vendored Reflaxe provenance, patch drift, upstreaming, and package contents;
- licenses and notices shipped with compiler, runtime, generated output, and vendor material.

Do not give legal advice. Name concrete licensing/provenance questions that need qualified review.

##### 16. Packaging, release, provenance, recovery, and GitHub conventions

Verify or identify the absence of:

- tag-to-commit identity;
- release-gate-to-tag identity;
- deterministic Haxelib package construction;
- isolated install and portable/metal smoke behavior;
- complete archive allowlist/denylist validation;
- artifact manifest, checksum, and source provenance;
- hosted asset digest binding;
- immutable tag/release behavior;
- idempotent same-tag partial-publication recovery;
- least-privilege workflow permissions;
- commit-pinned actions and tool versions;
- branch/tag protection or repository rulesets;
- release-note and changelog correctness;
- rollback and repair documentation;
- absence of release-time source identity ambiguity.

Distinguish source-verifiable behavior from live host settings. The captured zero assets, mutable release, absent rulesets, and unprotected default branch are evidence inputs, not details to smooth over.

##### 17. Test strategy and evidence quality

Assess whether tests find semantic and production defects rather than only preserving current output.

Review:

- red-green contract-first negative tests;
- semantic-difference oracles against Haxe behavior;
- generated snapshots and intentional review discipline;
- representative `go test`/run behavior;
- determinism and repeatability;
- exact generated-artifact and report tests;
- property, fuzz, mutation, stress, race, and failure-injection opportunities;
- concurrency flake detection and repeated scheduling;
- toolchain and platform matrices;
- package-install and release lifecycle tests;
- equivalent hand-written-Go comparisons;
- evidence freshness and exact-source binding.

Do not write “add more tests” as a recommendation. Name the exact invariant, fixture, oracle, mutation, environment, schedule, or failure injection required.

##### 18. Developer experience, operability, maintainability, and governance

Assess:

- clean installation and first compile;
- Haxe-developer onboarding and portability guidance;
- Go-developer discoverability and recognizable APIs;
- layered examples and third-party package interop;
- diagnostics, source positions, and actionable failures;
- build speed and incremental workflow;
- generated-Go debugging, stack traces, profiling, and observability;
- upgrade, migration, rollback, and release repair;
- architecture ownership, complexity concentration, and bus-factor risk;
- specification/code/test/report synchronization;
- public inventory and evidence drift guards;
- vendored framework governance;
- whether Beads closure or green CI can conceal unresolved product questions.

A production compiler must be operable and supportable, not merely compilable.

#### Questions requiring direct answers

Answer every question explicitly:

1. Is Haxe.Go suitable for any production use today? If yes, state the exact bounded workload, exclusions, pinning, and application-level evidence required.
2. Is “beta-stable” accurate? What exact public label and wording should replace it if not?
3. Is Haxe.Go ready for a stable 1.x compatibility promise today?
4. What is the minimum coherent and useful stable-1.x contract?
5. Which public claims are materially misleading, too weak, or unsupported?
6. Is the portable-default / Go-native-opt-in product model coherent and enforceable?
7. Are profile, native-boundary, strictness, optimizer, and runtime policies truly orthogonal?
8. Can portable code safely converge toward native Go representations through typed proof? What proof system is minimally sufficient?
9. Is the Haxe.Rust relationship genuinely the closest useful precedent? Which exact patterns transfer, and which must not?
10. Is the canonical `_std` migration direction correct? Which current `.cross.hx` classes should become ordinary staged source, generated artifacts, public support, or runtime code?
11. Is Reflaxe being leveraged appropriately, and where is local code duplicating or bypassing it?
12. Which Haxe 4.3.7 features would materially improve correctness, types, or code size?
13. Does the 14,825-line `GoCompiler.hx` create concrete risk? If so, what dependency-oriented decomposition should precede beta closure or stable admission?
14. Which `GoRaw` clusters are legitimate, and which indicate missing typed IR or misplaced ownership?
15. Is `hxrt` acceptably narrow? Where is compiler-known behavior still deferred to runtime?
16. Are there credible semantic defects in nullability, objects, interfaces, generics, closures, evaluation order, strings, collections, reflection, exceptions, or std/sys behavior?
17. Are there reachable races, deadlocks, lost wakeups, goroutine leaks, unsafe/checkptr hazards, resource leaks, data loss, or avoidable process panics?
18. Are `panic/recover`, typed Haxe catches, goroutine failures, and normal Go errors mapped coherently?
19. Are `go.Result`, portable result semantics, and Go multi-results/errors kept sufficiently distinct?
20. Does current `go.Select` naming/behavior mislead users? What typed native-select design is warranted, if any?
21. What is the minimum P0/P1 Go-native capability set required for the stated north star?
22. Which Go surfaces should be ordinary typed facades, compiler lowerings, generated bindings, macro DSLs, experimental escapes, or out of scope?
23. Can a Go developer plausibly pick up Haxe.Go today and write recognizable Go-style architecture with better Haxe ergonomics? What blocks that claim?
24. Does generated Go meet a credible production bar for each admitted lane?
25. Which generated-code quality problems are semantic necessities, which are missing optimization, and which are avoidable lowering debt?
26. Do current performance baselines compare equivalent semantics and detect meaningful regressions?
27. Are current toolchain, platform, security, race, static-analysis, and vulnerability policies sufficient for a validated beta release?
28. Is the current SemVer/version-authority model correct for pre-1.0 and future stable 1.x?
29. Is the release architecture correct, deterministic, same-tested-SHA, recoverable, and proportionate?
30. Which GitHub host controls and supply-chain measures are required before the next validated beta, versus before stable 1.x?
31. What exact evidence is missing for bounded production, beta-baseline closure, and stable-1.x admission respectively?
32. Which first-pass gaps are false positives, duplicates, already covered, or optional work that should not delay the beta milestone?
33. Which existing roadmap Beads should be amended, split, reprioritized, rejected, or supplemented?
34. What are the five most consequential actions in dependency order?

#### Required output format

##### 0. Review integrity and evidence accounting

State:

- whether the artifact filename, size, and SHA-256 matched;
- whether internal hashes were verified;
- which primary files, generated outputs, CI logs, reports, and sibling references were actually inspected;
- any unreadable, missing, truncated, or ambiguous evidence;
- every live fact independently checked outside the frozen bundle, with date and source;
- material limitations of the review.

Do not claim comprehensive inspection if context limits prevented it.

##### 1. Executive disposition

Provide a table with a separate verdict and confidence for all seven judgments:

- bounded production use now;
- stable 1.x compatibility promise;
- compiler architecture;
- portable-specialization direction;
- Go-native authoring experience;
- generated-Go output quality;
- release and distribution integrity.

Use only:

- `READY`
- `READY_WITH_BOUNDED_SCOPE`
- `NOT_READY`

Confidence must be `HIGH`, `MEDIUM`, or `LOW`, with one-sentence evidence limits. Do not average the seven verdicts.

##### 2. Exact supported-product statements

Write:

1. the exact short public statement Haxe.Go can truthfully publish today;
2. the exact bounded-production conditions and exclusions behind that statement;
3. the exact stronger validated-beta statement that becomes defensible after beta blockers close;
4. the minimum stable-1.x statement that becomes defensible after stable-admission blockers close.

Say explicitly whether and how the phrase “beta-stable” should be used.

##### 3. Known-fact and hypothesis adjudication

For every numbered item under **Known facts** and **Open hypotheses**, report:

- disposition;
- supporting or contradicting evidence;
- confidence;
- any correction to the wording;
- whether it changes roadmap scope.

Do not silently accept the first-pass framing.

##### 4. Product and stable-admission matrix

For each significant contract family, provide:

- contract/surface;
- current product status;
- stable-1.x disposition;
- exact qualification;
- protected units;
- exclusions/internal units;
- existing evidence;
- missing evidence;
- breaking-change implications;
- recommended owning milestone.

Use the stable dispositions:

- `Admit`
- `Admit-qualified`
- `Keep experimental`
- `Exclude/internalize`
- `Needs evidence`

Cover profiles and policies, Haxe APIs, Go-native APIs, metadata, defines, reports, diagnostics, generated layout, package layout, runtime boundary, toolchains/platforms, and release/provenance contracts.

##### 5. Severity-ranked findings

Use a table or repeated structured records containing exactly:

- finding ID;
- severity and confidence;
- repository-relative file and exact line range;
- supporting test, output, workflow, or live-state evidence;
- affected contract or public claim;
- violated invariant;
- concrete failure scenario;
- root cause;
- minimal root-cause fix or scope disposition;
- exact regression or empirical evidence required;
- whether it blocks any production, validated beta closure, only stable 1.x, or neither;
- existing Bead or milestone placement;
- proposed adjudication: accepted / rejected / duplicate / deferred / evidence-gap.

If a first-pass claim is a false positive, include it with the contradicting evidence instead of omitting it.

Use these severity meanings:

- `Blocker`: credible data corruption, memory/unsafe violation, security compromise, unrecoverable publication identity failure, or a foundational contradiction that makes the supported product untruthful.
- `High`: likely production failure, serious race/deadlock/resource/process-failure risk, stable-contract violation, or missing evidence essential to an admitted beta/stable surface.
- `Medium`: important semantic qualification, generated-quality, maintainability, operability, or incomplete governance issue.
- `Low`: bounded cleanup or optional hardening.

Do not inflate severity because a feature is incomplete or aesthetically non-idiomatic.

##### 6. Cross-cutting architecture assessment

Explain:

- what is sound and should be preserved;
- what is accidental complexity;
- what should be typed or moved to a different owner;
- what should be decomposed and along which dependency seams;
- where Reflaxe should replace local machinery;
- where Reflaxe is insufficient and a Go-specific abstraction is justified;
- what should explicitly not be built;
- how Haxe 4.3.7 can simplify the design without macro overreach.

Favor the smallest architecture that preserves the stated contracts.

##### 7. Sibling transfer matrix

For each material pinned Haxe.Rust, Haxe.Ruby, and Haxe.Elixir precedent, classify it as:

- transfer directly;
- adapt to Go;
- use only as a contract/evidence precedent;
- reject as target-specific;
- insufficient evidence.

Give the Go-specific reason. Directly answer whether Haxe.Go is closest to Haxe.Rust and whether that statement is architecturally useful.

##### 8. Empirical-verification list

Separate:

- facts established by frozen source and payload hashes;
- facts established by executable evidence in the bundle;
- claims inferred from snapshots or documentation;
- host, platform, toolchain, race, performance, or release facts requiring fresh empirical verification.

For every needed verification, give the exact command, fixture, environment, repetition count, mutation, schedule, or failure injection and the expected invariant.

##### 9. Dependency-ordered programs

Produce two separate dependency graphs or ordered lists:

1. the minimal work required to publish the validated beta-baseline release;
2. the additional minimal work required for a coherent stable-1.x promise.

For each item state:

- why it is necessary;
- whether qualification/exclusion can remove it;
- prerequisites;
- completion evidence;
- existing roadmap owner;
- whether an independent second pass is needed.

Keep optional feature expansion out of both critical paths unless it is foundational to the stated contract.

##### 10. Beads adjudication proposals

Map findings first to the existing roadmap epics:

- `haxe_go-vfp.4` — toolchains, security, and deterministic releases;
- `haxe_go-vfp.5` — canonical Reflaxe `_std` and package layout;
- `haxe_go-vfp.6` — compatibility, profiles, SemVer, and public API;
- `haxe_go-vfp.7` — typed usage and portable surface registry;
- `haxe_go-vfp.8` — typed Go IR and compiler ownership;
- `haxe_go-vfp.9` — Go-native typed authoring layer;
- `haxe_go-vfp.10` — runtime semantics, concurrency, errors, and unsafe boundaries;
- `haxe_go-vfp.11` — idiomatic-Go output and performance evidence;
- `haxe_go-vfp.12` — examples, closure evidence, beta release, and later stable admission.

For every proposed new or amended Bead provide:

- title;
- issue type;
- priority P0-P3;
- `thinking:low|medium|high|xhigh`;
- dependencies;
- concise Why / What / How description;
- measurable acceptance criteria;
- relevant files and evidence;
- duplicate check against the supplied roadmap;
- whether Oracle/independent second-pass review is required.

Do not create a duplicate issue for work already represented. Recommend amending an existing Bead when its acceptance criteria can own the finding cleanly.

##### 11. Direct answers

Answer all 34 numbered questions directly and concisely, even when earlier sections already discuss them.

##### 12. Final go/no-go criteria

End with:

- present bounded-production verdict;
- present beta-stability wording verdict;
- present stable-1.x verdict;
- exact conditions that would flip each negative verdict;
- what may safely remain deferred after the validated beta and after stable 1.x;
- the next single action recommended;
- the five most consequential actions in dependency order.

##### 13. Optional longer-term improvements

Keep broad Go standard-library facades, additional third-party binding polish, wider platform CI, deeper performance optimization, multi-package output, advanced generics, CGo breadth, extensive ecosystem integrations, and other non-blocking opportunities here unless direct evidence makes one foundational.

#### Evidence discipline

- Cite repository-relative files and exact line ranges. Use the line-numbered Repomix views for navigation and confirm important claims against the authoritative raw files where possible.
- Prefer code, executable tests, generated output, workflow behavior, artifact contents, and captured API state over prose.
- Clearly label inference, frozen evidence, reported evidence, and fresh live verification.
- Do not treat a green test, snapshot, closed issue, compatibility table, or documentation claim as proof by itself.
- Do not require every experimental or unsupported feature to become stable.
- Do not require perfection before bounded production use.
- Do not accept “more testing” without the exact invariant and test strategy.
- Recommend root-cause compiler/runtime/contract changes, not generated-file edits or application workarounds.
- Distinguish Haxe semantic requirements from Go idiom preferences.
- Distinguish Go-native capability gaps from portable Haxe correctness gaps.
- Distinguish missing audit evidence from a missing repository feature.
- Treat sibling code as precedent, not proof, and never as automatic implementation authority.
- Treat the captured release as evidence of current failure/success state, not proof of all release paths.
- Do not let release mechanics crowd out compiler semantics and DX, but do not minimize clear provenance blockers.
- Do not invent missing release assets or assume unpublished Haxelib bytes exist.
- Do not implement changes or edit the supplied project.

## Upload checklist

### Required

#### 1. Exact canonical evidence ZIP

Attach only:

`haxe-go-gpt56-evidence-cd79624f.zip`

Verify before upload:

```text
bytes   13502320
sha256  ab4b0a1097229ad2202ca7da6c092b2b85cba537522951fb625f6b1312c4b511
```

The ZIP already contains:

- the exact source Git archive at `cd79624f855521dbf320ac2b7524d889ca388c0e`;
- the line-numbered primary Repomix view;
- the eight recorded Repomix scanner exclusions and raw companion files;
- source/blob inventory and evidence index;
- committed portable, metal, example, and snapshot output;
- exact-SHA GitHub Actions logs and metadata;
- live release, tag, repository, and host-control captures;
- the Haxe.Go Next Beads/Dolt roadmap snapshot;
- pinned Haxe.Rust, Haxe.Ruby, and Haxe.Elixir reference slices and a line-numbered combined view;
- Gitleaks and local-path validation results;
- internal hashes and reproduction documentation.

Do not also attach the old untracked files named `repomix-output-haxe.go.xml` or `repomix-output-haxe.go.xml.zip`. Their commit identity and omission set are not authoritative.

#### 2. This prompt

Paste the complete **Review prompt** section. The prompt was authored after the frozen source commit and is intentionally not part of the compiler source archive under review.

### Optional only when a disputed question requires it

#### 3. Official Haxe 4.3.7 sources

Use the exact official 4.3.7 tag for disputed stdlib, macro, typed-AST, or target-selection semantics. Do not use current Haxe `main` as if it described 4.3.7.

#### 4. Upstream Reflaxe reference

Use a commit-pinned upstream Reflaxe slice only for a concrete framework-divergence or upstreamability dispute. The primary bundle already contains the actually shipped vendored Reflaxe implementation.

#### 5. Fresh live host evidence

Only if necessary, re-check release assets, release immutability, tag identity, rulesets, branch/tag protection, or workflow runs. Record the exact time and URLs, and keep those facts separate from the frozen capture.

No separate `v0.53.1` Haxelib release ZIP is available in the captured GitHub Release. Its absence is itself evidence and must not be filled with a locally invented artifact.
