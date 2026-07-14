# 0. Review integrity and evidence accounting

## Artifact identity

The canonical artifact `dist/review/haxe-go-gpt56-evidence-cd79624f.zip` matched exactly:

| Check | Result |
|---|---|
| Filename | `haxe-go-gpt56-evidence-cd79624f.zip` |
| Size | `13,502,320` bytes — matched |
| SHA-256 | `ab4b0a1097229ad2202ca7da6c092b2b85cba537522951fb625f6b1312c4b511` — matched |
| ZIP integrity | `unzip -t` passed |
| Internal hashes | All 379 `SHA256SUMS` entries verified; zero failures |
| Builder file | SHA-256 `447be705cb2a4ad73116f83256dcac125d2b4e745d195584b5819a8dd4272a78` matched |
| Source identity | Authoritative Git archive records `cd79624f855521dbf320ac2b7524d889ca388c0e` |
| Payload security checks | Gitleaks passed; all 375 checked UTF-8 payloads passed workspace-path validation |

I did not read either untracked `repomix-output-haxe.go.xml` file. The eight Repomix scanner exclusions were accounted for through `primary/repomix-security-exclusions.json`, their raw companion files, and the authoritative Git archive.

## Evidence actually inspected

I inspected:

- Bundle root `MANIFEST.json`, `README.md`, `SHA256SUMS`.
- `primary/source-inventory.json`, the Git archive, the line-numbered XML navigation view, and all exclusion metadata.
- Product/profile/compatibility documentation, including `README.md`, `docs/profiles.md`, `docs/portable-semantics-v1.md`, `docs/feature-support-matrix.md`, `docs/known-gaps.md`, release/security/performance documents.
- Package and release configuration: `package.json`, `package-lock.json`, `haxelib.json`, `.releaserc.json`, release scripts and workflows.
- Compiler architecture and representative raw clusters: `GoCompiler.hx`, `GoReflaxeCompiler.hx`, AST/printer, pass registry, bootstrap, runtime-feature analyzer, selected stdlib emitters, typed Go facades.
- Runtime implementations for exceptions, filesystem, process, threading, event loops, synchronization and pools.
- Representative generated portable/metal/Go-native output for maps, slices, results, channels, exceptions, applications and snapshots.
- Tests around filesystem, process, thread pools, condition primitives, channels, results, profiles, strictness, semantic differential behavior and examples.
- Exact-SHA metadata and logs for all five cited GitHub Actions runs.
- Frozen release, tag, repository, security-setting and host-control captures.
- Performance methodology and committed app/microbenchmark baselines.
- Roadmap root and nine relevant workstream epics.
- Pinned Rust typed-usage/registry/runtime/release precedents, Ruby SemVer/release repair precedents, and Elixir `_std`/package/release precedents.

## Evidence limitations

- This was a prioritized architecture review, not a line-by-line audit of all 5,368 tracked files or all 5,108 raw-node occurrences.
- I did not execute the compiler or modify the reviewed project. Executable conclusions rely on exact-SHA CI logs and committed outputs; several proposed defects therefore require the focused regressions specified below.
- CI artifact bytes such as uploaded benchmark and examples ZIPs were not present in the evidence bundle; only their logs and metadata were available.
- The roadmap reports 663 issues, but its detailed object graph contains only the root and 12 immediate issues. Child-level duplicate checking is therefore impossible from the bundle.
- No fresh project-host check was performed. Release assets, rulesets, protection and immutability conclusions use the frozen capture only.
- The archive has no `.git`; identity depends on the recorded Git-archive construction, remote containment, hashes and exact-SHA CI metadata as documented.
- Windows was cross-built, not natively runtime-tested. macOS ran the quality suite, but the strongest deployment evidence remains Linux.
- No comprehensive memory-unsafe or malicious-input audit was possible. I found no demonstrated checkptr/unsafe P0, but the absence of such evidence is not proof of safety.

## Separately dated external checks

On 2026-07-13 I checked only official upstream sources:

- Go 1.26.5 and 1.25.12 were current patch releases; Go supports a major line only until two newer majors exist. Thus the frozen 1.22/1.23 matrix is unsupported now. [Official Go release policy and history](https://go.dev/doc/devel/release).
- The official Node schedule records Node 20 EOL as 2026-04-30. [Official Node release schedule](https://raw.githubusercontent.com/nodejs/Release/main/schedule.json).
- Haxe 4.3.7 officially provides `Compiler.getConfiguration`, `Context.onAfterInitMacros`, `withOptions`, `makeMonomorph`, and typed metadata/define registration. [Haxe 4.3.7 Compiler API](https://api.haxe.org/haxe/macro/Compiler.html), [Context API](https://api.haxe.org/haxe/macro/Context.html).
- Haxe 4.3.7 explicitly requires `File.getContent` and `File.saveContent` to throw on read/write failure. [Official Haxe 4.3.7 `sys.io.File`](https://raw.githubusercontent.com/HaxeFoundation/haxe/4.3.7/std/sys/io/File.hx).

No files, Beads, commits or remote state were changed. The current tracked worktree was clean and `master` matched `origin/master`; pre-existing untracked Repomix files and scratch files were left untouched.

# 1. Executive disposition

| Judgment | Verdict | Confidence | Evidence limit |
|---|---|---:|---|
| Bounded production use now | `READY_WITH_BOUNDED_SCOPE` | MEDIUM | Credible only for pinned, application-revalidated, portable, non-sys/non-concurrent workloads; no trustworthy release artifact represents the reviewed SHA. |
| Stable 1.x compatibility promise | `NOT_READY` | HIGH | Version authority, public-unit inventory, operation-level semantic admission, package layout, deprecation policy and release identity are incomplete. |
| Compiler architecture | `READY_WITH_BOUNDED_SCOPE` | MEDIUM | The outer AST-first design is sound, but raw/string authority, a 14.8k-line ownership concentration and false-success build handling create concrete change risk. |
| Portable-specialization direction | `READY_WITH_BOUNDED_SCOPE` | HIGH | The direction is correct, but broad native representation still lacks typed usage closure, a versioned registry and per-surface proof. |
| Go-native authoring experience | `NOT_READY` | HIGH | Current `go.*` is a small experimental wrapper layer and does not support ordinary Go architecture without important semantic compromises. |
| Generated-Go output quality | `READY_WITH_BOUNDED_SCOPE` | MEDIUM | Output is deterministic, gofmt-clean and Go-buildable, but representative native output remains `any`/reflection-heavy and materially behind equivalent Go baselines. |
| Release and distribution integrity | `NOT_READY` | HIGH | No reviewed-SHA Haxelib artifact, no same-tested-SHA transaction, no deterministic package/recovery path, mutable releases and no repository rulesets/protection. |

# 2. Exact supported-product statements

## 2.1 Truthful statement today

> Haxe.Go is a pre-1.0 beta compiler for Haxe 4.3.7. Its portable profile can support pinned, application-qualified workloads within the explicitly tested subset. Go-native APIs, concurrency, process, network/TLS and distribution artifacts remain experimental or separately qualified.

Do not use “beta-stable.” It conflates lifecycle maturity with compatibility stability. If retained in historical material, define it only as “a beta release line whose APIs may still break”; the better choice is to retire the phrase.

## 2.2 Current bounded-production conditions

A defensible workload is deterministic Linux application or service logic that:

- Pins source commit `cd79624f…`; it must not claim that `v0.53.1` contains that source.
- Uses Haxe 4.3.7 and a currently supported, application-tested Go patch release.
- Uses `portable`, enables `reflaxe_go_strict`, and sets `reflaxe_go_portable_native_policy=error`.
- Keeps `reflaxe_go_auto=off`; conservative adopters should also use `reflaxe_go_opt=none`.
- Uses no raw application injection and no `go.*` facade.
- Excludes `sys.io.File` text helpers, `sys.io.Process`, `sys.thread`, event loops, thread pools, network, HTTP/TLS/socket behavior and cross-goroutine exceptions.
- Runs explicit `go test ./...`, `go vet ./...`, the deployed binary, and application semantic tests after generation; the backend’s default `go build` result cannot be trusted.
- Adds application-specific differential tests for business-critical Haxe semantics, race testing even if concurrency is believed absent, load tests, failure injection and deployment smoke.
- Obtains qualified licensing review for copied `hxrt` code before distribution.

That is a useful but narrow lane: in-memory/core domain logic, deterministic CLI/batch computation, or services whose filesystem/network/process infrastructure is supplied through separately audited typed adapters.

## 2.3 Statement after validated-beta blockers close

> Haxe.Go is a validated pre-1.0 beta for the documented portable subset and explicitly listed native lanes on supported toolchains and platforms. Each release binds a deterministic Haxelib artifact to the exact tested commit, and every admitted runtime surface has semantic, failure-path and race evidence. APIs may still change before 1.0 with documented migration guidance.

## 2.4 Minimum stable-1.x statement

> Haxe.Go 1.x guarantees its versioned portable semantics subset, admitted public Haxe APIs, documented profile and policy behavior, selected typed Go interop metadata, package/install contract, versioned reports, diagnostic categories, supported toolchain/platform policy and exact-source release provenance. Unlisted Go-native APIs and generated implementation details remain experimental; generated module and runtime boundaries are stable only where the compatibility manifest explicitly says so.

# 3. Known-fact and hypothesis adjudication

## Known facts

| # | Disposition | Evidence and correction | Confidence | Roadmap effect |
|---:|---|---|---|---|
| 1 | Confirmed | `primary/source-inventory.json` reports all stated counts. | HIGH | None. |
| 2 | Confirmed | Raw archive `src/reflaxe/go/GoCompiler.hx` is 14,825 lines. | HIGH | Keep `.8`. |
| 3 | Confirmed, qualified | Counts are accurate debt locators; raw reflection/serialization fixtures are not automatically defects. | HIGH | Add category budgets, not a blanket ban. |
| 4 | Confirmed | Profiles plus separate policy controls are implemented and documented. | HIGH | `.6`; clarify coupled defaults. |
| 5 | Confirmed | `README.md:12-21`, `docs/profiles.md:1-16`. | HIGH | Preserve wording direction. |
| 6 | Confirmed | Sixty-eight source `.cross.hx` files and noncanonical `std/_std`; package generation is a separate concern. | HIGH | `.5` remains required for package release. |
| 7 | Confirmed | `src/go/**`, `std/go/**`, metadata fixtures and the thin current facade. | HIGH | `.9`, but not a portable-beta blocker. |
| 8 | Confirmed | `src/go/Select.hx:25-71`, `docs/known-gaps.md:115-119`. | HIGH | Rename/experimentalize in `.9`. |
| 9 | Confirmed | Current representations and generated bridges are distinct and sometimes lose Go error identity. | HIGH | `.6`, `.9`, `.10`. |
| 10 | Confirmed | `package.json:4`, `haxelib.json:12`, `.releaserc.json:25-42`, frozen zero-assets release. | HIGH | `.4`. |
| 11 | Confirmed as frozen fact; now stale | CI used Go 1.22/1.23 and Node 20; official upstream policy now makes those unsupported/EOL. | HIGH | Raise `.4` urgency. |
| 12 | Confirmed mechanically, narrowed | Logs prove 231 snapshots, 55 strict sweep, 126 semantic cases, six optimizer cases, three lane cases and 175 accounted modules. The 175 full sweep is compile-only because CI omits `--go-test`; snapshots do not run programs unless `--runtime` is set. | HIGH | Correct evidence taxonomy in `.6/.12`. |
| 13 | Confirmed only at summary level | Stats show 663 issues and zero cycles; detailed child issues are absent from the bundle. | HIGH | Child duplicate review still required. |

## Open hypotheses

| # | Disposition | Reason | Confidence | Scope effect |
|---:|---|---|---|---|
| 1 | Confirmed, narrower | A source-pinned portable core is usable; “stable” without “beta” is indefensible. Current sys/concurrency exclusions must be explicit. | HIGH | Amend `.6/.12`. |
| 2 | Partly confirmed | Rust is the closest pinned contract precedent for typed systems output, but Go’s GC, interfaces, errors, runtime and tooling support a materially simpler design. | HIGH | Avoid Rust-shaped machinery without Go need. |
| 3 | Confirmed | Typed usage, surface contracts, runtime requirements and proof gates transfer. Ownership/borrowing, `Send`/`Sync`, Cargo and `no_hxrt` do not. | HIGH | `.7` should adapt concepts, not code. |
| 4 | Partly confirmed | Canonical ordinary `.hx` source under `std/go/_std` is correct. Not every current file is an upstream override; support/runtime/public facade ownership must be classified first. Source `.cross.hx` need not remain. | HIGH | `.5` acceptance is directionally right. |
| 5 | Confirmed | Raw barriers already force conservative optimization, and compiler-owned std/runtime blocks enlarge regression radius. Typed seams and ratchets are the right response. | HIGH | `.8`; no cosmetic split. |
| 6 | Confirmed as direction | Broad portable specialization is not yet proven. | HIGH | `.7` stable/default prerequisite, not portable-beta prerequisite if disabled. |
| 7 | Confirmed | Current facades cannot express ordinary Go architecture. | HIGH | `.9` product program, experimental meanwhile. |
| 8 | Partly confirmed | A handwritten-Go performance corpus exists and reveals meaningful gaps. Missing pieces are readability review, race/debug/failure evidence and repeated current-toolchain measurements. | HIGH | Narrow `.11`; reject “no comparison exists.” |
| 9 | Confirmed | Current semantic-release path can tag an untested prepare commit and hosts no verified Haxelib artifact or recovery contract. | HIGH | `.4` blocks validated release. |
| 10 | Partly confirmed | Haxe 4.3.7 can reduce reflection and dynamic typing, but canonical packaging—not reflective mutation through `CompilerConfiguration` internals—is the correct bootstrap fix. | MEDIUM | Targeted modernization under `.5/.8`. |

# 4. Product and stable-admission matrix

| Contract/surface | Current status | Stable disposition | Qualification; protected/excluded units | Existing / missing evidence | Breaking implications; milestone |
|---|---|---|---|---|---|
| `portable` profile | Beta subset | `Admit-qualified` | Protect profile name, portable semantic invariance and explicit subset; exclude unlisted operations. | Strong fixture set; missing operation-level manifest and current-toolchain run. | Default or semantic changes are breaking; `.6/.12`. |
| `metal` profile | Experimental Go-first preset | `Keep experimental` | Protect only opt-in/fail-fast meaning until native surface is admitted. | Profile negatives and examples; native API incomplete. | Defaults and fallback policy can break users; `.6/.9`. |
| Strict/native boundary policies | Implemented | `Admit` | Protect define grammar, defaults and diagnostic category. Note profile-dependent defaults. | Good negative fixtures. | Changing default acceptance is breaking; `.6`. |
| Optimizer/auto/runtime controls | Implemented, partly coupled | `Admit-qualified` | Stable controls may alter representation only under semantic invariance; experimental optimizer capabilities remain versioned. | Six optimizer and three lane cases; no central proof registry. | New default specialization can be breaking semantically/operationally; `.7`. |
| Portable language semantics | Partial normative contract | `Needs evidence` | Admit rules by operation/type shape, not by module compilation. | Strong null/string/bytes/numeric/exception subset; spec has only five broad families. | Expanding is additive; changing admitted behavior needs major contract version; `.6/.7`. |
| Public Haxe std APIs | Mixed | `Needs evidence` | Protect only enumerated members/operations. Exclude current File/Process/thread/network failures until fixed. | 126 semantic cases and 55-member sweep; failure paths sparse. | Transitive type/member changes can break; `.5/.6/.10`. |
| `go.*` facades | Thin, target-specific | `Keep experimental` | `Map`, `Slice`, `Result`, `Chan`, `Select` must not be called complete Go abstractions. | Snapshots and smokes; semantic gaps below. | Method sets, carriers and generated shapes can break; `.9`. |
| Typed extern metadata | Useful partial surface | `Admit-qualified` | Admit validated grammar for `@:go.import`, `@:go.name`, `@:go.receiver`; keep advanced callbacks/generics/unsafe experimental. | Positive snapshots; needs exhaustive malformed-input and collision tests. | Metadata grammar/default changes are breaking; `.6/.9`. |
| Native multi-results/errors | Partial bridges | `Keep experimental` | Keep Haxe result, `go.Result`, `(T,error)` and tuple results distinct. | Shape tests; error identity/wrapping not preserved. | Carrier or method-set changes are breaking; `.9/.10`. |
| Defines and metadata catalogue | Public but stringly | `Admit-qualified` | Protect names, accepted values and defaults for admitted controls; internal counters excluded. | Define reference and negatives; no generated compatibility manifest. | New enum-like values may break exhaustive consumers; `.6`. |
| Diagnostics | Unversioned | `Needs evidence` | Stabilize codes/categories and source anchoring, not exact prose. | Negative fixtures; build errors incorrectly warn. | Severity/exit-status changes can break automation; `.6/.8`. |
| JSON/Markdown reports | Opt-in snapshots | `Keep experimental` until schema-versioned | Protect filenames/schema versions only after inventory; Markdown prose excluded. | Snapshot coverage; no stable compatibility policy. | Removing/renaming fields is breaking once admitted; `.6/.7`. |
| Generated package/file layout | Deterministic current shape | `Admit-qualified` | Stabilize module/runtime/import boundaries required by consumers; exclude internal filenames, temporaries and formatting. | 3,624 intended files and gofmt/go-test. | Import path, package boundary and build-tag changes can break; `.6/.11`. |
| `_std` and installed package layout | Noncanonical | `Needs evidence` | Ordinary source in `std/go/_std`; package build may generate `.cross.hx`; public facades separate. | Sibling precedent; no isolated deterministic Haxelib package. | Migration is pre-1.0 breaking and should precede stable; `.5`. |
| `hxrt` boundary | Functional but broad | `Admit-qualified` | Stabilize only versioned import/capability contract; implementation internal. Exclude unsafe/thread/network capabilities until admitted. | Runtime smokes and slicing harness; no race/checkptr gate. | Runtime import/module changes can break generated projects; `.10`. |
| Toolchains/platforms | Frozen CI policy obsolete | `Needs evidence` | Initially admit Haxe 4.3.7, supported Go lines, Node LTS and Linux amd64; other OSes explicitly qualified. | Old Linux/macOS CI, Windows cross-build. | Raising floors is compatibility-significant; `.4/.6`. |
| Version/release/provenance | Incoherent | `Needs evidence` | One version authority, exact tested SHA, immutable artifacts and same-tag recovery. | Tags and CI identity captured; no Haxelib asset. | Any source/artifact mismatch violates release identity; `.4`. |
| Licensing/notices | Ambiguous for generated runtime | `Needs evidence` | Compiler GPL-3.0-only; copied `hxrt` and generated-output licensing require explicit policy and notices. | Manifest license and MIT upstream headers; no generated-runtime statement. | License change is adoption-significant; `.4/.5`. |

# 5. Severity-ranked findings

## HG-REL-01 — Default Go build failure can report compiler success

- **Finding ID:** `HG-REL-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `src/reflaxe/go/GoReflaxeCompiler.hx:233-282`
- **Supporting evidence:** README says the backend runs `go build` by default at `README.md:156`; snapshots use `go_no_build` and separately run Go tooling.
- **Affected contract or public claim:** Successful compilation and default post-codegen build.
- **Violated invariant:** A failed mandatory build step must make the compiler invocation fail.
- **Concrete failure scenario:** `go build` exits 7, but only a warning is emitted and the Haxe invocation can exit successfully, allowing broken artifacts downstream.
- **Root cause:** Nonzero status and exceptions are converted to `Context.warning`.
- **Minimal root-cause fix or scope disposition:** Restore cwd with a guaranteed cleanup path and emit a source-anchored fatal compilation error unless the user selected codegen-only mode.
- **Exact regression or empirical evidence required:** Fixture `-D go_cmd=<script-that-exits-7>`; assert nonzero Haxe exit, restored cwd, one diagnostic, and no false success.
- **Production blocking scope:** Blocks validated beta closure and any workflow trusting default build success.
- **Existing Bead or milestone placement:** `haxe_go-vfp.8`, closure gate in `.12`.
- **Proposed adjudication:** accepted.

## HG-SYS-01 — File text helpers silently lose write/read failures

- **Finding ID:** `HG-SYS-01`
- **Severity and confidence:** Blocker / HIGH
- **Repository-relative file and exact line range:** `runtime/hxrt/sys.go:110-120`; compiler bridge `src/reflaxe/go/GoCompiler.hx:5208-5225`; happy-path-only test `test/semantic_diff/file_read_write_contract/Main.hx:1-13`
- **Supporting evidence:** Haxe 4.3.7 requires exceptions on read/write failure; the implementation discards `os.WriteFile` errors and returns `""` for any `os.ReadFile` error.
- **Affected contract or public claim:** `sys.io.File` is marked semantic-diff supported at `docs/feature-support-matrix.md:61-64`.
- **Violated invariant:** Portable std behavior must preserve normal failure, not turn it into success.
- **Concrete failure scenario:** Saving to a directory, full filesystem or unwritable path reports success; reading a missing/denied file is indistinguishable from a valid empty file.
- **Root cause:** Convenience wrappers collapse Go `error` into `Void`/empty string.
- **Minimal root-cause fix or scope disposition:** Throw the appropriate Haxe-visible failure while preserving message/cause where promised; until fixed, exclude these methods from production and supported status.
- **Exact regression or empirical evidence required:** Differential tests for missing file, directory-as-file, permission failure and unwritable destination; assert both profiles throw and never return an empty success.
- **Production blocking scope:** Blocks any production use of these methods and any beta claiming `sys.io.File`; qualification removes it from the narrow core lane.
- **Existing Bead or milestone placement:** `haxe_go-vfp.10`.
- **Proposed adjudication:** accepted.

## HG-SYS-02 — `sys.io.Process` conceals startup, stream and lifecycle failures

- **Finding ID:** `HG-SYS-02`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `runtime/hxrt/process.go:8-59`; `test/semantic_diff/process_echo_contract/Main.hx:1-17`
- **Supporting evidence:** Pipe/start errors produce a partially initialized object; EOF, scanner error and empty line collapse to `""`; close kills and ignores all errors. The only semantic test runs `haxe --version`.
- **Affected contract or public claim:** `sys.io.Process` is marked supported at `docs/feature-support-matrix.md:61`.
- **Violated invariant:** Startup failure, EOF, empty data, nonzero exit and cleanup must remain distinct.
- **Concrete failure scenario:** Missing executable looks like empty stdout; a scanner token over 64 KiB terminates silently; `close()` destroys the process without exposing status.
- **Root cause:** An underspecified wrapper omits most of the Haxe `Process` surface and error contract.
- **Minimal root-cause fix or scope disposition:** Implement typed start/pipe/exit/stderr/stdin lifecycle or mark the entire surface experimental/unsupported.
- **Exact regression or empirical evidence required:** Nonexistent executable, empty line vs EOF, >64 KiB stdout, stderr, nonzero exit, blocking/nonblocking exit code, kill/close, partial read and 1,000-process FD/goroutine leak tests.
- **Production blocking scope:** Blocks process use and validated beta admission of the surface.
- **Existing Bead or milestone placement:** `haxe_go-vfp.10`.
- **Proposed adjudication:** accepted.

## HG-CONC-01 — Thread-pool run/shutdown races can lose accepted work

- **Finding ID:** `HG-CONC-01`
- **Severity and confidence:** Blocker / HIGH
- **Repository-relative file and exact line range:** `std/sys/thread/FixedThreadPool.cross.hx:46-64`; `std/sys/thread/ElasticThreadPool.cross.hx:44-99`; `std/sys/thread/ElasticThreadPoolWorker.cross.hx:7-75`
- **Supporting evidence:** `_isShutdown` is read/written without synchronization; fixed-pool work can be enqueued after shutdown sentinels; elastic worker fields use inconsistent locks. Current test only does orderly run then shutdown at `test/semantic_diff/sys_thread_runtime_contract/Main.hx:87-109`.
- **Affected contract or public claim:** Thread pool semantics and “all previously submitted tasks execute.”
- **Violated invariant:** Every accepted task executes exactly once or submission fails deterministically; shared state must be race-free.
- **Concrete failure scenario:** `run()` observes false, shutdown enqueues termination, then `run()` enqueues a task behind all sentinels; every worker exits and the accepted task is lost.
- **Root cause:** Nonatomic check/enqueue/shutdown transition copied from generic upstream source without a Go-memory-model-safe ownership protocol.
- **Minimal root-cause fix or scope disposition:** Serialize acceptance and shutdown with one state mutex and queue protocol; until fixed, exclude thread pools.
- **Exact regression or empirical evidence required:** `go test -race` with 10,000 concurrent run/shutdown schedules at `GOMAXPROCS=1,2,8`, unique task IDs, exactly-once accounting, deterministic rejection and bounded shutdown; repeat 100 times.
- **Production blocking scope:** Blocks concurrency production and validated beta admission of thread pools.
- **Existing Bead or milestone placement:** `haxe_go-vfp.10`.
- **Proposed adjudication:** accepted.

## HG-CONC-02 — Condition/event-loop/thread lifecycle has deadlock, race and leak hazards

- **Finding ID:** `HG-CONC-02`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `runtime/hxrt/thread.go:13-18,143-176,192-217,221-256,348-362,397-407,677-726,792-805`
- **Supporting evidence:** `eventLoop` is read/written without a field lock; broadcast uses globally consumable tokens; cancellation markers remain for pending events; goroutine identity parses `runtime.Stack`; spawned jobs have no recover policy. Condition tests cover only reentrant locking at `test/semantic_diff/sys_thread_primitives_contract/Main.hx:62-67`.
- **Affected contract or public claim:** Threads, conditions, event loops, background failure and cleanup.
- **Violated invariant:** Broadcast applies only to waiters present at broadcast; lifecycle state is race-free and bounded; background failure policy is explicit.
- **Concrete failure scenario:** A late waiter consumes a broadcast token before an already-broadcast waiter reacquires the mutex, leaving the original waiter blocked indefinitely.
- **Root cause:** Token accounting substitutes for waiter generations; thread state has no single ownership model.
- **Minimal root-cause fix or scope disposition:** Use generation/per-waiter condition state, synchronize event-loop ownership, delete cancellation state on all terminal paths and define background panic propagation.
- **Exact regression or empirical evidence required:** Barrier-controlled old/late-waiter broadcast test; race test for events/runWithEventLoop; cancel 100,000 pending timers and assert bounded state; background panic contract; goroutine/state leak test.
- **Production blocking scope:** Blocks concurrency admission; otherwise stable-only qualification.
- **Existing Bead or milestone placement:** `haxe_go-vfp.10`.
- **Proposed adjudication:** accepted.

## HG-SEC-01 — Release gates use unsupported toolchains and intentionally pass reachable vulnerabilities

- **Finding ID:** `HG-SEC-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `.github/workflows/ci-quality.yml:17-59`; `.github/workflows/ci-harness.yml:19-54`; `scripts/security/run-dependency-audit.sh:139-161`; `docs/security-dependency-audit.md:44-56`
- **Supporting evidence:** Exact-SHA log `live/github/runs/29294117737.log:2472-2668` reports 18 reachable Go-standard-library vulnerabilities and then passes.
- **Affected contract or public claim:** Supported production toolchain, network/TLS security and validated release gate.
- **Violated invariant:** An admitted production surface must be tested on supported patched toolchains and fail closed on reachable unwaived vulnerabilities.
- **Concrete failure scenario:** A release passes while compiled/tested with known reachable TLS/X.509/network vulnerabilities.
- **Root cause:** Stale Go/Node pins and an explicit fail-open classification policy.
- **Minimal root-cause fix or scope disposition:** Test supported Go lines and Node LTS; fail on reachable findings unless a time-bounded, surface-specific waiver excludes the affected capability.
- **Exact regression or empirical evidence required:** Full CI, race and `govulncheck` on current supported Go patches; inject a known reachable test advisory and assert failure; inject an unreachable advisory and assert classified nonblocking result.
- **Production blocking scope:** Blocks validated beta; network/TLS is excluded from present bounded production.
- **Existing Bead or milestone placement:** `haxe_go-vfp.4`.
- **Proposed adjudication:** accepted.

## HG-REL-02 — Current release path cannot establish one tested source identity

- **Finding ID:** `HG-REL-02`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `.releaserc.json:1-43`; `.github/workflows/ci-harness.yml:241-273`; `.github/workflows/examples-artifacts.yml:35-42,54-114`; `package.json:4`; `haxelib.json:12`
- **Supporting evidence:** Frozen `v0.53.1` release has zero assets and `immutable:false`; tag points to `4adb9ae…`, not reviewed source; no rulesets or branch protection. Semantic-release prepares a changelog/release commit after gates.
- **Affected contract or public claim:** Version, tag, Haxelib package, GitHub Release and provenance identity.
- **Violated invariant:** Tag, tested commit and hosted deterministic artifact must identify the same source.
- **Concrete failure scenario:** Gates pass commit A; release preparation creates commit B; tag identifies B although B was not gated, and no deterministic Haxelib bytes prove what was shipped.
- **Root cause:** Release mutation is mixed with publication and no package transaction/recovery design exists.
- **Minimal root-cause fix or scope disposition:** Use a tested release PR or stage metadata outside the checkout; build deterministic package twice; isolated-install test; tag exact tested SHA; attach ZIP/checksum/manifest/provenance; immutable same-tag repair.
- **Exact regression or empirical evidence required:** Two byte-identical builds; isolated portable/metal install smoke; tag/SHA assertion; interrupt after tag but before assets and rerun without new version, moved tag or changed bytes.
- **Production blocking scope:** Does not block source-pinned private use; blocks the next validated public beta and all stable releases.
- **Existing Bead or milestone placement:** `haxe_go-vfp.4`, final gate `.12`.
- **Proposed adjudication:** accepted.

## HG-CONTRACT-01 — “Supported” currently means too little for sys/native/concurrency claims

- **Finding ID:** `HG-CONTRACT-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `docs/feature-support-matrix.md:15-45,49-80,202-209`; `docs/known-gaps.md:99-128`
- **Supporting evidence:** Support can be established through snapshot/harness/go-test without failure-path, race or semantic completeness; File, Process, channels/results/maps/slices are labeled supported despite concrete gaps.
- **Affected contract or public claim:** Public support matrix and README production caveat.
- **Violated invariant:** “Supported” must identify the admitted operation and its evidence strength.
- **Concrete failure scenario:** A user treats `sys.io.File` or `go.Map<K,V>` as production-supported because a happy-path semantic fixture and snapshot exist.
- **Root cause:** Evidence-tier and product-admission concepts are conflated.
- **Minimal root-cause fix or scope disposition:** Separate compile-covered, shape-covered, runtime-smoke, semantic, race/failure and release-installed states; admit operations, not whole modules.
- **Exact regression or empirical evidence required:** Machine-readable support manifest lint requiring operation/member, profile, platform, evidence IDs, exclusions and failure/race status before `Supported`.
- **Production blocking scope:** Blocks truthful validated-beta wording; current bounded wording can qualify it.
- **Existing Bead or milestone placement:** `haxe_go-vfp.6`, `.12`.
- **Proposed adjudication:** accepted.

## HG-NATIVE-01 — Current Go-native containers/errors/channels are not faithful Go abstractions

- **Finding ID:** `HG-NATIVE-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `src/go/Go.hx:3-41`; `src/go/Map.hx:3-20`; `src/go/Slice.hx:3-30`; `src/go/Result.hx:3-40`; `src/go/Chan.hx:3-104`; `src/go/Select.hx:25-71`
- **Supporting evidence:** Map stringifies all keys; channel handle is `Dynamic`; `tryRecv` lacks closed-channel `ok`; `Result` stores message wrappers; `Select` polls first branch first. Generated map/slice/result remain `any`-based.
- **Affected contract or public claim:** “metal-ready” list and Go-developer authoring north star.
- **Violated invariant:** Typed Go facades must preserve Go type/identity/concurrency semantics or carry an explicit different name.
- **Concrete failure scenario:** Distinct keys with equal `Std.string` collide; a closed channel yields a successful zero value indistinguishable from a sent zero; wrapped error identity is destroyed.
- **Root cause:** Haxe wrapper implementations are specialized cosmetically without first-class Go types/results in IR and facade contracts.
- **Minimal root-cause fix or scope disposition:** Keep these experimental; build true comparable maps, typed slices/channels, comma-ok/multi-result carriers and error-preserving adapters. Rename `Select` polling helpers to `PriorityPoll`.
- **Exact regression or empirical evidence required:** Key-collision/comparability negatives, closed-channel comma-ok, nil channel, send-on-closed, wrapped/sentinel error identity, evaluation-order and native-select fairness/shape tests.
- **Production blocking scope:** Blocks Go-native product claim and stable admission; not the portable core beta.
- **Existing Bead or milestone placement:** `haxe_go-vfp.9`, with error/concurrency parts in `.10`.
- **Proposed adjudication:** accepted.

## HG-ARCH-01 — AST-first shell is sound, but raw/string authority creates optimization and ownership risk

- **Finding ID:** `HG-ARCH-01`
- **Severity and confidence:** Medium / HIGH
- **Repository-relative file and exact line range:** `src/reflaxe/go/ast/GoAST.hx:3-87`; `GoASTPrinter.hx:162-184,365-415`; `RewriteVirtualCallsPass.hx:64-120`; `RegistryCore.hx:9-76`; `GoReflaxeCompiler.hx:175-181`
- **Supporting evidence:** Types, operators, imports and result lists are strings; raw text passes verbatim; optimization clears knowledge at raw statements; pass dependencies exist but invariants do not; compiler generic parameters use `Dynamic`.
- **Affected contract or public claim:** AST-first architecture, maintainability and deterministic optimization.
- **Violated invariant:** Typed semantic information should remain authoritative until printing.
- **Concrete failure scenario:** A raw composite or multi-result hides aliasing/call effects, forcing whole regions into conservative fallback or allowing malformed Go only at printer time.
- **Root cause:** IR grew around a minimal printer while std/runtime lowering accumulated in compiler-owned emitters.
- **Minimal root-cause fix or scope disposition:** Introduce typed Go type/import/operator/composite/multiassign/error nodes capability by capability and declare pass pre/postconditions.
- **Exact regression or empirical evidence required:** Characterization snapshots, malformed-node negative tests, raw-category budget, and per-capability semantic tests before moving ownership.
- **Production blocking scope:** Broad work is stable/maintainability scope; the false-success build slice blocks beta.
- **Existing Bead or milestone placement:** `haxe_go-vfp.8`.
- **Proposed adjudication:** accepted.

## HG-STD-01 — Source/package layout is not yet an isolated installable Reflaxe target

- **Finding ID:** `HG-STD-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `haxelib.json:17-25`; `src/reflaxe/go/CompilerBootstrap.hx:20-31,45-95`; all 68 tracked `*.cross.hx`
- **Supporting evidence:** Only `std` is declared; source uses `std/_std`; classpath order is reflectively mutated; no deterministic isolated Haxelib package was supplied.
- **Affected contract or public claim:** Installation, target std selection, package provenance and modern Reflaxe conventions.
- **Violated invariant:** Source ownership and installed selection must be explicit and reproducible.
- **Concrete failure scenario:** Source checkout passes because bootstrap mutates internal classpath state, while an installed package selects a different std module or fails under compilation-server reuse.
- **Root cause:** Historical `.cross.hx` packaging artifacts became source authority.
- **Minimal root-cause fix or scope disposition:** Classify ownership, migrate overrides to ordinary `.hx` under `std/go/_std`, declare std paths, generate `.cross.hx` only when packaging, and isolated-install test.
- **Exact regression or empirical evidence required:** Source/package selection shadow test; package build twice; isolated haxelib install; portable and metal compile/go-test/run; assert no source `.cross.hx`.
- **Production blocking scope:** Blocks validated distribution and stable; not source-checkout bounded use.
- **Existing Bead or milestone placement:** `haxe_go-vfp.5`.
- **Proposed adjudication:** accepted.

## HG-COMPAT-01 — No mechanically enforceable stable public contract or version authority

- **Finding ID:** `HG-COMPAT-01`
- **Severity and confidence:** High / HIGH
- **Repository-relative file and exact line range:** `package.json:4`; `haxelib.json:12`; `.releaserc.json:5-42`; `docs/portable-semantics-v1.md:27-148`
- **Supporting evidence:** Two manifests say `0.1.0`, latest tag says `0.53.1`; semantic contract covers five risk families but not complete admitted units; report/diagnostic/generated boundaries lack stable classification.
- **Affected contract or public claim:** SemVer and stable 1.x.
- **Violated invariant:** Every stable consumer-visible unit must have one owner, version and change rule.
- **Concrete failure scenario:** A metadata default, enum case, report field, generated path or transitive public type changes under a minor release because it was absent from compatibility checks.
- **Root cause:** Snapshot inventory substitutes for an API/contract manifest.
- **Minimal root-cause fix or scope disposition:** One version authority plus generated manifests for public Haxe units, metadata, defines, diagnostics, reports, package/runtime boundaries and qualified generated layout.
- **Exact regression or empirical evidence required:** Baseline-vs-candidate compatibility diff with fixtures for removal, changed default, enum case, report field, diagnostic severity and toolchain floor.
- **Production blocking scope:** Stable 1.x only, apart from fixing public beta version truth.
- **Existing Bead or milestone placement:** `haxe_go-vfp.6`.
- **Proposed adjudication:** accepted.

## HG-OUT-01 — Output is buildable but not yet close to ordinary Go on admitted native lanes

- **Finding ID:** `HG-OUT-01`
- **Severity and confidence:** Medium / HIGH
- **Repository-relative file and exact line range:** `test/snapshot/go_native/slice_map_metal_monomorph/intended/module_go_map.go:5-38`; `module_go_slice.go:3-44`; `main.go:8-57`; `scripts/ci/perf/app-profile-baseline.json:51-156`
- **Supporting evidence:** Generated native containers use `any`, StringMap, assertions, reflection and IIFEs. Baseline shows roughly 0.54–0.59× throughput, 1.69–1.83× latency, 2.34–3.83× allocation bytes and 7.35–7.85× allocation count versus scripted pure Go.
- **Affected contract or public claim:** “Generate readable Go,” metal/native quality and performance.
- **Violated invariant:** Native-lane abstractions should lower recognizably and avoid unnecessary boxing when type proof exists.
- **Concrete failure scenario:** A Go engineer profiling a nominally typed metal container finds generic `any` carriers and assertion-heavy stacks with little portable/metal difference.
- **Root cause:** Wrapper specialization occurs around erased backing representations.
- **Minimal root-cause fix or scope disposition:** Keep current claim at “inspectable, gofmt-clean Go”; fix representations through typed IR/registry, not peephole text rewriting.
- **Exact regression or empirical evidence required:** Three steady-state runs on current toolchains, `benchstat`, escape/allocation reports, committed handwritten/generated corpus and human readability rubric.
- **Production blocking scope:** Neither for narrow portable beta; stable native admission and stronger performance claims only.
- **Existing Bead or milestone placement:** `haxe_go-vfp.11`.
- **Proposed adjudication:** accepted.

## HG-HXRT-01 — Runtime slicing is coarse and panic/error ownership is underdefined

- **Finding ID:** `HG-HXRT-01`
- **Severity and confidence:** Medium / HIGH
- **Repository-relative file and exact line range:** `src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx:60-143`; `runtime/hxrt/exception.go:16-61`; `GoReflaxeCompiler.hx:308-355`
- **Supporting evidence:** Core/string/print/exception are unconditional; any `sys.*` adds broad sys; several shim groups also add process; full runtime is default. `TryCatch` recovers every panic, including compiler/runtime panics.
- **Affected contract or public claim:** Minimal runtime, failure distinction and runtime reports.
- **Violated invariant:** Capability selection should follow typed reachable use, and Haxe exceptions should not silently absorb arbitrary backend faults.
- **Concrete failure scenario:** A simple program copies and compiles unused process/TLS/thread code; a generated type-assertion panic becomes a normal Haxe catch.
- **Root cause:** Coarse class/shim inference and one universal panic carrier.
- **Minimal root-cause fix or scope disposition:** Typed capability ledger; distinguish explicit Haxe exception carrier from foreign/backend panic; full-copy remains supported but honestly reported.
- **Exact regression or empirical evidence required:** Absence tests per capability; malformed compiler panic must escape; explicit Haxe throw must catch; report exactly explains every copied file/import.
- **Production blocking scope:** Stable/runtime-quality scope; security-sensitive unused surfaces strengthen beta toolchain requirements.
- **Existing Bead or milestone placement:** `haxe_go-vfp.7`, `.10`.
- **Proposed adjudication:** accepted.

## HG-SUPPLY-01 — Vendored Reflaxe provenance and output confinement are insufficiently bounded

- **Finding ID:** `HG-SUPPLY-01`
- **Severity and confidence:** Medium / MEDIUM
- **Repository-relative file and exact line range:** `vendor/reflaxe/haxelib.json:1-11`; `vendor/reflaxe/PATCHES.md:1-11,154-193`; `vendor/reflaxe/src/reflaxe/output/OutputManager.hx:327-440`
- **Supporting evidence:** Version is `4.0.0-beta` without an upstream commit/tree digest; patch documentation is inherited from Elixir. `saveFile` permits selected absolute prefixes and does not visibly enforce canonical output-root containment.
- **Affected contract or public claim:** Framework provenance, malicious-input safety and package reproducibility.
- **Violated invariant:** Vendored code must have exact upstream identity and generated writes must remain within the chosen output root unless explicitly authorized.
- **Concrete failure scenario:** A future vendor refresh cannot prove its base; a crafted path or symlink may escape the output directory if an untrusted route reaches `saveFile`.
- **Root cause:** Version-label vendoring and ad hoc path sanitization instead of canonical-path confinement.
- **Minimal root-cause fix or scope disposition:** Record upstream commit/tree hash and patch series; enforce realpath-relative output ownership. Reachability of malicious paths must be established before calling it a security defect.
- **Exact regression or empirical evidence required:** Try `../escape.go`, absolute paths and an output symlink to an external directory; assert no external write. Reconstruct vendor tree from pinned upstream plus patches and compare hashes.
- **Production blocking scope:** Provenance blocks validated package release; path risk is currently an evidence gap.
- **Existing Bead or milestone placement:** `haxe_go-vfp.4`, framework portion of `.5/.8`.
- **Proposed adjudication:** evidence-gap.

## HG-LIC-01 — Copied runtime licensing is not explained to generated-code consumers

- **Finding ID:** `HG-LIC-01`
- **Severity and confidence:** Medium / HIGH
- **Repository-relative file and exact line range:** `haxelib.json:4`; `LICENSE:1-30`; generated-output copying in `src/reflaxe/go/GoReflaxeCompiler.hx:308-355`; `vendor/reflaxe/haxelib.json:4`
- **Supporting evidence:** Project declares GPL-3.0-only; `hxrt` is copied into generated projects; no reviewed document states the license of generated compiler output versus copied runtime or required notices.
- **Affected contract or public claim:** Commercial/production adoption and distributed artifacts.
- **Violated invariant:** Consumers must be able to determine the licensing and notice obligations of shipped generated artifacts.
- **Concrete failure scenario:** An adopter distributes an application containing copied `hxrt` without knowing whether GPL terms apply to that runtime or broader work.
- **Root cause:** Compiler, runtime, generated output and imported MIT sources are not separated in licensing documentation.
- **Minimal root-cause fix or scope disposition:** Obtain qualified legal review, declare licenses per component, ship required notices and state generated-output policy. This is not legal advice.
- **Exact regression or empirical evidence required:** Package-content license/notice linter and an isolated generated artifact inventory reviewed by qualified counsel.
- **Production blocking scope:** Requires adopter-specific review now; blocks a generally advertised validated distribution until clarified.
- **Existing Bead or milestone placement:** `haxe_go-vfp.4/.5`.
- **Proposed adjudication:** accepted.

## HG-AUDIT-01 — Three first-pass claims are false positives

- **Finding ID:** `HG-AUDIT-01`
- **Severity and confidence:** Low / HIGH
- **Repository-relative file and exact line range:** `package-lock.json`; `primary/repomix-security-exclusions.json`; `src/reflaxe/go/ast/GoAST.hx:35-87`
- **Supporting evidence:** A lockfile exists; all eight XML omissions exist in raw/archive form; typed AST nodes already cover select, go, defer, send, range and type assertions.
- **Affected contract or public claim:** Roadmap scope and architectural debt framing.
- **Violated invariant:** Audit findings must identify the actual missing invariant.
- **Concrete failure scenario:** Work recreates a lockfile, restores files that were never absent, or replaces legitimate typed/raw leaf nodes solely to reduce a count.
- **Root cause:** Inventory observations were treated as feature absence.
- **Minimal root-cause fix or scope disposition:** Amend `.4` from “lacks lockfile” to “does not enforce lockfile”; reject missing-file and blanket-raw claims.
- **Exact regression or empirical evidence required:** Use `npm ci`; keep exclusion/hash verification; maintain category-based raw budget.
- **Production blocking scope:** Neither.
- **Existing Bead or milestone placement:** `haxe_go-vfp.3/.4/.8`.
- **Proposed adjudication:** rejected.

# 6. Cross-cutting architecture assessment

What should be preserved:

- Portable-default/Go-native-opt-in is coherent.
- Central build-context resolution is a useful control point.
- The outer typed-AST → Go AST → passes → printer direction is right.
- Pass dependency validation is deterministic.
- Profiles, strictness, examples, semantic differentials and generated snapshots form a strong governance base.
- Examples operate as executable contracts rather than decorative snippets.

Accidental complexity:

- String-encoded types/operators/imports/results inside the IR.
- Raw blocks that combine compile-time shape decisions with whole library implementations.
- Reflective classpath mutation.
- Profile checks directly controlling representation without a semantic registry.
- Full/coarse runtime copying.
- Vendored Elixir-origin patch commentary and ad hoc output-path logic.
- Public wrapper types that look native but retain erased backing storage.

Dependency-oriented decomposition should proceed as:

1. Build/package/report driver out of `GoReflaxeCompiler`, including fatal post-build behavior.
2. Module, symbol, package, import and Go-type planning.
3. Typed Haxe-expression/statement lowering.
4. Portable specialization admission and decision reporting.
5. Standard-library bridge registry, grouped by owner.
6. Runtime-call/capability abstraction.
7. Pure printer with no semantic decisions.

`GoCompiler.hx` should not first be split by line count. Each extraction needs characterization snapshots and an invariant such as “all multi-result construction is typed” or “File semantics no longer originate in compiler raw code.”

Representative ownership:

- `EReg`: staged `std/go/_std/EReg.hx` plus a thin regex runtime helper; metadata-driven serializer tables can remain compiler-context-generated but should use typed switch/struct nodes.
- File/Process/thread public semantics: staged target std; actual OS/state operations in typed `hxrt`.
- Reflection/RTTI tables: compile-context emitter, typed outer IR, dynamic payload only where Haxe reflection requires it.
- Go packages, interfaces, methods, results and channels: typed Go-specific IR/facades.
- Imports: typed planner, never incidental raw text.

Reflaxe should provide type usage, generic compiler plumbing, package output and target `_std` conventions. Go-specific abstractions remain justified for multi-results, method sets, interfaces, channels/select, errors, package imports and build constraints.

Do not build:

- A second Haxe compiler in macros.
- A complete hand-maintained Go standard-library facade before a binding path exists.
- Rust ownership/borrow machinery.
- Multi-package output merely to close beta.
- Zero runtime as an absolute goal.
- A global raw-node ban.
- Syntax churn or macros for ordinary callable APIs.

Useful Haxe 4.3.7 modernization:

- Use `Compiler.getConfiguration` only for its supported typed fields; remove internal classpath mutation through canonical packaging.
- Move typer-dependent initialization to `Context.onAfterInitMacros`.
- Use typed metadata/define registration for diagnostics and completion.
- Replace profile/metadata/report strings with enum abstracts or exhaustive typed values.
- Replace `GenericCompiler<Bool,Bool,Dynamic,Dynamic,Dynamic>` with concrete target node types.
- Apply null safety incrementally to compiler/planner packages.
- Use `Context.withOptions`, `withImports` and `makeMonomorph` only where they remove real state/reflection or improve typing.

All 16 stated architectural principles are sound. None requires replacement.

# 7. Sibling transfer matrix

| Precedent | Classification | Go-specific disposition |
|---|---|---|
| Rust typed usage ledger | Adapt to Go | Track operations/type shapes, not ownership/borrow facts. |
| Rust surface contract registry | Adapt to Go | Keep version/source semantics/native/fallback/import/runtime fields; use Go comparability, nil, aliasing and error rules. |
| Rust runtime requirement analysis | Adapt to Go | Directly useful for `hxrt`, but Go linker/runtime behavior permits a simpler capability model. |
| Rust portable-specialization proof gate | Use as contract/evidence precedent | Do not copy first-pass implementation blindly. |
| Rust deterministic release architecture | Transfer directly | Same tested SHA, reproducible package, manifest/checksum/provenance and recovery are target-independent. |
| Rust ownership, borrow, lifetimes, `Send`/`Sync` | Reject as target-specific | Go has GC and a different concurrency contract; race freedom still needs Go-native evidence. |
| Cargo/no-runtime mechanics | Reject as target-specific | Use `go.mod`, package imports and typed `hxrt` capabilities. |
| Ruby pre-major SemVer policy | Transfer directly | Breaking pre-1.0 changes need explicit bump/migration behavior. |
| Ruby release lineage and same-tag repair | Adapt to Go | Haxelib ZIP and generated Go install smoke replace Ruby-specific package steps. |
| Ruby runtime/package conventions | Reject as target-specific | Not relevant to Go output/runtime. |
| Elixir ordinary target `_std` source layout | Transfer directly | Use `std/go/_std/**/*.hx`; `.cross.hx` is a package artifact. |
| Elixir deterministic Haxelib archive tests | Adapt to Go | Add Go portable/metal install/build/run and Go module checks. |
| Elixir BEAM/runtime design | Reject as target-specific | Go runtime, goroutines and packages require independent design. |

Haxe.Go is closest to Haxe.Rust among the supplied references only in the limited sense of typed systems-language output, native representations and runtime minimization. That statement is useful for contract architecture; it is distorting if it drives Rust ownership, Cargo or runtime complexity into Go.

# 8. Empirical-verification list

## Established by frozen hashes/source

- Exact source identity, inventory, file counts and exclusions.
- Compiler/runtime source behavior described above.
- Current release/tag/host capture.
- Version/configuration mismatch.
- Presence of the lockfile.
- Generated snapshots and performance baselines.

## Established by executable evidence in the bundle

- Exact-SHA CI success.
- 231 snapshots compile/gofmt/go-test/diff.
- 55 strict upstream modules with Go tests.
- 175 full modules compile, but not all with Go tests.
- 126 semantic differential cases, six optimizer cases, three lane cases.
- 12 example profile cases compile/go-test/run/stdout-match.
- Linux/macOS quality runs and cross-build matrix.
- Gitleaks and dependency-audit output.
- Perf harness completion and current warning/hard-gate policy.

## Inferred, not established

- Snapshot readability and idiomaticity.
- Semantic completeness of whole modules.
- Production safety of sys/thread/network surfaces.
- Stable compatibility.
- Race freedom and checkptr safety.
- Release reproducibility.
- Native-select semantics.
- Go-developer usability.

## Required fresh verification

| Invariant | Exact verification |
|---|---|
| Build failure propagates | Run a snapshot with `go_cmd` script exiting 7; expected Haxe exit nonzero and cwd unchanged. |
| File failures preserve Haxe behavior | Differential missing/permission/directory/full-disk fixtures; assert exceptions, not empty success. |
| Process lifecycle | Missing executable, empty line, EOF, 100 KiB line, stderr, exit 9, close/kill and 1,000-child leak fixture. |
| Thread-pool exactly-once | 10,000 run/shutdown schedules at `GOMAXPROCS=1,2,8`; `go test -race -count=100`; accepted IDs equal executed IDs exactly once. |
| Condition broadcast | Barrier N old waiters, broadcast, insert late waiter before reacquisition; all old return, late remains pending. Repeat 10,000 times. |
| Runtime race/checkptr | `go test -race -count=100 ./...`; `go test -gcflags=all=-d=checkptr=2 -count=20 ./...` on generated runtime fixtures. |
| Timer cancellation boundedness | Cancel 100,000 pending events; white-box runtime test asserts cancellation state returns to baseline. |
| Supported toolchains | Full CI on Go 1.25.12 and 1.26.5, Node 22 and 24, Haxe 4.3.7; Linux and macOS native. |
| Windows qualification | Native Windows compile/go-test/run for representative portable/sys examples; cross-build is insufficient. |
| Vulnerability fail-closed | Inject reachable and unreachable advisory fixtures; reachable unwaived case must fail. |
| Output confinement | Attempt `../escape.go`, absolute path and output-root symlink escape; assert no write outside canonical root. |
| Release reproducibility | Build Haxelib ZIP twice in fresh checkouts; `cmp` and SHA-256 must match. |
| Isolated install | Install ZIP in empty Haxelib repo; compile/go-test/run one portable and one qualified metal case. |
| Release recovery | Terminate after immutable tag but before assets; rerun against same tag; no new version/tag and identical artifact bytes. |
| Performance | Warm once, run `npm run test:perf:apps` three times on a pinned runner, compare medians with `benchstat`; record current toolchain. |
| Generated quality | `go vet ./...`, pinned staticcheck, race, escape analysis and committed handwritten/generated diff for admitted corpus. |
| Native channel/error semantics | Closed/open comma-ok, nil channel, send-on-closed, wrapped/sentinel error identity, evaluation order and source positions. |
| Archive safety | Package fixtures with `..`, absolute entries, duplicate paths, symlinks and case collisions; reject before extraction. |

# 9. Dependency-ordered programs

## 9.1 Minimal validated beta-baseline release

1. **Define and enforce the beta boundary.**
   - Necessary: prevents current supported-surface overclaim.
   - Qualification can remove File/Process/thread/network/Go-native work from beta.
   - Prerequisite: this review.
   - Evidence: machine-readable admitted-operation manifest and exact README wording.
   - Owner: `.6`, `.12`.
   - Independent second pass: yes.

2. **Move CI to supported toolchains and fail-closed security/dependency resolution.**
   - Necessary: old green CI is not current release evidence.
   - Cannot be qualified away for a public release.
   - Prerequisite: supported-toolchain policy.
   - Evidence: full CI/race/security on supported Go and Node; `npm ci`.
   - Owner: `.4`.
   - Independent second pass: yes.

3. **Fix red-contract blockers inside the admitted beta surface.**
   - At minimum: fatal post-build status; File/Process failures if admitted; thread/pool/condition issues only if concurrency is admitted.
   - Qualification/exclusion can remove sys/concurrency implementation from this milestone.
   - Evidence: exact regressions in section 8.
   - Owner: `.8`, `.10`.
   - Independent second pass: runtime/concurrency yes.

4. **Ship canonical package layout.**
   - Necessary for trustworthy Haxelib installation.
   - Cannot be skipped if the milestone publishes Haxelib.
   - Prerequisites: ownership inventory and beta public manifest.
   - Evidence: no source `.cross.hx`, deterministic package, isolated install.
   - Owner: `.5`.
   - Independent second pass: yes.

5. **Replace the release transaction and enable minimal host controls.**
   - Necessary: bind tested source to immutable bytes.
   - Prerequisite: canonical package and supported gates.
   - Evidence: double build, isolated test, immutable tag/release, hosted ZIP/checksum/manifest/provenance, repair drill.
   - Owner: `.4`.
   - Independent second pass: yes.

6. **Run closure evidence and publish.**
   - Evidence: admitted examples, application-level smoke, final support manifest, resolved findings and second review.
   - Owner: `.12`.
   - Independent second pass: mandatory.

Broad portable specialization, complete Go-native P0/P1, multi-package output and extensive optimization should not block this portable-first beta.

## 9.2 Additional minimum work for stable 1.x

1. One version authority and mechanically generated public-contract manifest.
2. Operation-level stable portable semantics with explicit exclusions and deprecation policy.
3. Versioned reports, diagnostic codes/categories and qualified generated-layout contract.
4. Typed-usage registry for every native representation enabled by default.
5. Decompose the admitted compiler/runtime seams enough to make compatibility changes reviewable.
6. Admit only the Go-native subset with complete error/channel/container semantics; keep everything else experimental.
7. Current supported platform/toolchain policy with floor-change rules.
8. Stable release repair, host protection and artifact provenance policy.
9. License/notice policy for installed and generated artifacts.
10. Independent stable-admission review against an actual release candidate.

# 10. Beads adjudication proposals

The supplied epics already cover the findings. Do not create new broad epics. Detailed child duplication must be checked against live Beads before creating child tasks.

| Existing Bead / amendment | Type, priority, thinking | Dependencies | Why / What / How and acceptance | Evidence / duplicate check / second pass |
|---|---|---|---|---|
| Amend root `haxe_go-vfp`: split “validated portable beta” from full native/stable program | epic, P0, `thinking:xhigh` | Review `.3` | Beta acceptance must not require every `.7/.9/.11` expansion. Require only admitted-surface evidence and truthful exclusions. | Roadmap root acceptance currently overbundles scope. Existing root, no duplicate. Independent review required. |
| Close review checkpoint `haxe_go-vfp.3` only after findings are recorded | epic, P0, `thinking:xhigh` | None | Attach this disposition; map every accepted/rejected finding; record evidence limitations. | Exact purpose of `.3`; no new issue. Independent review is this deliverable. |
| Amend `haxe_go-vfp.4` | epic, P0, `thinking:xhigh` | `.3`, package phase depends `.5/.6` | Correct “lacks lockfile” to “workflows do not enforce the lockfile.” Add current Go/Node matrix, fail-closed reachability, action SHA pinning, release transaction, recovery, vendor provenance, host controls and licensing. | Files: workflows, release/security scripts, package manifests, vendor metadata. Existing epic cleanly owns it. Oracle/second pass required. |
| Amend `haxe_go-vfp.5` | epic, P0, `thinking:xhigh` | `.3`, contract inventory `.6` | Classify each of 68 files before migration; ordinary `std/go/_std` source; support/public/runtime separation; generated `.cross.hx` only in package; isolated install. | Existing acceptance is correct. No new epic. Second pass required. |
| Reprioritize `haxe_go-vfp.6` from P1 to P0 | epic, P0, `thinking:xhigh` | `.3` | Publish exact beta boundary first. Add operation-level evidence states, coupled profile defaults, diagnostic/report/generated-boundary classification and stable manifest. | Owns `HG-CONTRACT-01` and `HG-COMPAT-01`. Existing epic. Second pass required. |
| Split implementation slices inside `haxe_go-vfp.10` | epic, P0 for admitted slices, `thinking:xhigh` | `.6`; ownership portions `.5` | Children should separately own File errors, Process lifecycle, pool run/shutdown, condition generations, event-loop state, panic/error policy and race/checkptr gates. Qualification may defer excluded children. | Existing epic is correct but too broad for one acceptance path. Child duplicate check required. Second pass required. |
| Amend `haxe_go-vfp.8` | epic, P1; build-failure child P0, `thinking:high` | `.3`; broad decomposition after `.5/.7` | First child: fatal post-build contract. Then typed Go type/import/composite/multi-result seams and category raw budgets. | Do not make wholesale monolith split a beta blocker. Existing epic. No Oracle unless scope becomes xhigh. |
| Rephase `haxe_go-vfp.7` | epic, P1 stable/default specialization; P2 for beta, `thinking:xhigh` | `.6`, typed seams `.8` | Required before a new native representation becomes default, not before a conservative portable beta. Registry acceptance must include typed usage, fallback, imports/runtime and semantic proof. | Existing epic owns it. Second pass required before default admission. |
| Rephase `haxe_go-vfp.9` | epic, P1 product roadmap; P2 for portable beta, `thinking:xhigh` | `.6`, `.8`, error/runtime `.10` | Keep current facades experimental. Define minimal typed P0/P1 capability program and rename polling `Select`. | Existing epic, no duplicate. Second pass required before public native admission. |
| Narrow beta dependency on `haxe_go-vfp.11` | epic, P1, `thinking:high` | `.8/.7` for deeper optimization | Existing corpus disproves “no comparison.” Beta needs current-toolchain representative evidence and disclosed thresholds; broad optimization remains optional. | Amend description, retain corpus/readability work. No Oracle normally. |
| Amend `haxe_go-vfp.12` | epic, P0, `thinking:xhigh` | `.4/.5/.6` plus only admitted slices of `.8/.10` | Closure must consume the machine-readable admitted surface, package provenance, exact release bytes and final review. It must state that beta closure does not imply native or stable completion. | Existing final milestone. Independent second pass mandatory. |

# 11. Direct answers

1. **Suitable for production today?** Yes, only for source-pinned, strict portable, non-sys/non-concurrent core workloads with explicit Go build/test/run and application qualification described in section 2.
2. **Is “beta-stable” accurate?** No. Use “pre-1.0 beta for pinned, application-qualified portable workloads.”
3. **Ready for stable 1.x?** No.
4. **Minimum stable contract?** Versioned portable subset, admitted public Haxe APIs, policy/metadata/define grammar, diagnostic categories, report schemas, package/install/runtime boundaries, supported toolchains and exact release provenance; keep broad native/generated detail experimental.
5. **Misleading claims?** Whole-module “Supported,” “metal-ready” containers/results, broad “generate readable Go,” “usable with caveats,” and “zero actionable blockers” without explaining that the latter is inventory policy rather than production evidence.
6. **Portable-default/native-opt-in coherent?** Yes.
7. **Axes truly orthogonal?** Separately configurable, but no: strict defaults depend on profile and specialization checks explicitly branch on profile.
8. **Can portable converge safely?** Yes, with typed operation usage, shape closure, versioned surface contract, explicit fallback/runtime/import accounting and semantic differential/property evidence.
9. **Is Rust closest?** At the contract level, yes. Transfer ledgers/registries/runtime analysis/release evidence; reject ownership, borrow, Cargo and `Send`/`Sync` designs.
10. **Is `_std` migration correct?** Yes. Upstream overrides become ordinary `std/go/_std` source; target support and public Go facades remain separate; package build may generate `.cross.hx`. No source `.cross.hx` currently needs to remain authoritative.
11. **Is Reflaxe leveraged appropriately?** Partly. Generic compiler/output/type usage are used, but canonical std packaging, typed usage and upstreamable output confinement are bypassed or duplicated.
12. **Useful Haxe 4.3.7 features?** Typed configuration reads, `onAfterInitMacros`, typed metadata/define registration, `withOptions`, `withImports`, `makeMonomorph`, enum abstracts and package null safety.
13. **Does the 14,825-line compiler create risk?** Yes: std/runtime ownership and raw/string IR enlarge regression radius. Decompose by build, planning, typed lowering, specialization, std bridges, runtime capabilities and printer.
14. **Legitimate raw clusters?** Irreducible syntax leaves, compile-context reflection/serializer data and explicit test fixtures. Missing typed IR includes composites/multiassign/error paths; misplaced ownership includes EReg, socket, sys/file/process and other library behavior.
15. **Is `hxrt` narrow enough?** Its subsystems are mostly genuine runtime work, but selection is coarse and default full-copy is broader than necessary.
16. **Credible semantic defects?** Yes in File/Process failure behavior, channel closed-state/error identity, thread pools and conditions. No additional P1 in null/object/generic/closure/string semantics was proven from the reviewed evidence.
17. **Reachable races/deadlocks/leaks/panics?** Credible races and task loss in pools, `eventLoop` race, condition lost wake, cancellation-state leak, background panic crash and goroutine-ID dependence. No demonstrated checkptr/unsafe P0.
18. **Are panic/catch/error mappings coherent?** Not fully. Explicit Haxe throws work in covered cases, but arbitrary panics are caught, background goroutine panics are unowned, and normal Go errors are sometimes stringified or swallowed.
19. **Are result contracts distinct?** Documentationally yes; generated bridges still lose Go error identity, so typed conversion is incomplete.
20. **Does `go.Select` mislead?** Yes. Rename current helpers to `PriorityPoll`; implement true native select only as a typed macro/lowering for syntax Haxe cannot express.
21. **Minimum P0/P1 native capability set?** Typed values/pointers/nil, real slices/maps, structs/tags/interfaces/method sets, multi-results and comma-ok, preserved errors, channel directions/range/close/native select, `defer`/`go`, context/sync/time baseline, imports/aliases/build constraints and binding validation.
22. **Surface ownership?** Ordinary facades for packages/types/functions; compiler lowering for Go representations; generated bindings for external packages; macros only for select/defer/go/destructuring/tags/build constraints; unsafe/CGo experimental; broad ecosystem facades optional.
23. **Can a Go developer write recognizable Go today?** Not plausibly beyond small interop examples. Container semantics, errors, methods/interfaces, context/sync and syntax gaps block the claim.
24. **Generated Go production bar?** Portable bounded lane: qualified yes. Current native/metal north-star lane: no. Sys/concurrency/network lanes: no until failure/race/toolchain evidence.
25. **Quality causes?** Haxe null/dynamic/exception machinery is sometimes semantic necessity; boxing from erased native wrappers is avoidable lowering debt; native specialization and runtime slicing are missing optimization.
26. **Are perf baselines meaningful?** Yes, workloads are scripted equivalently and expose real gaps. They are one old-toolchain baseline and current app gates are mostly nonblocking, so they do not establish a production SLA.
27. **Toolchain/platform/security policy sufficient?** No: unsupported Go/Node lines, no race/checkptr/staticcheck gate, reachable vulnerabilities pass, Windows lacks runtime evidence.
28. **SemVer model correct?** Pre-major Conventional Commit intent is plausible; actual version authority is not.
29. **Release architecture correct?** No.
30. **Required host controls?** Before validated beta: protected immutable tags/releases, required exact gates, no force-delete/move, least-privilege release workflow, pinned actions and controlled publication environment. Before stable: durable branch/ruleset enforcement, reviewed protection changes and documented recovery governance. Native secret scanning is desirable but not a substitute for fail-closed Gitleaks.
31. **Missing evidence by milestone?** Bounded production: application semantic/failure/race/load evidence on supported Go. Validated beta: supported toolchains, admitted-surface failure/race gates, canonical installable package and exact release provenance. Stable: compatibility manifest, deprecation policy, versioned schemas, proof registry and stable-admission review.
32. **False positives/optional gaps?** Lockfile absence is false; XML omissions are not missing files; raw count is not a bug count. Full Go stdlib, multi-package output, advanced generics, CGo and ecosystem breadth should not delay portable beta.
33. **Beads changes?** Amend/reprioritize `.4/.5/.6/.8/.10/.12`; rephase `.7/.9/.11`; split the root beta milestone from the full north-star program.
34. **Five most consequential actions?** Define admitted beta scope; move to supported fail-closed toolchains; fix/qualify build/File/Process/concurrency blockers; ship canonical deterministic package/release; run independent closure and publish.

# 12. Final go/no-go criteria

- **Present bounded-production verdict:** Go, with the narrow source-pinned portable conditions and exclusions in section 2.
- **Present beta-stability wording:** No-go for “beta-stable.” Use “pre-1.0 beta for pinned, application-qualified portable workloads.”
- **Present stable-1.x verdict:** No-go.
- **What flips bounded scope to broader production:** Fatal build failures, correct admitted sys error semantics, race-safe admitted concurrency, supported patched toolchains, current application evidence and clarified runtime licensing.
- **What flips validated beta:** Canonical isolated package, exact tested-SHA deterministic release, fail-closed security, truthful machine-readable admission and second-pass closure.
- **What flips stable 1.x:** One version authority, enforceable compatibility manifest, versioned semantic/report/diagnostic contracts, stable package/runtime boundaries, proof for default specializations and an actual release-candidate audit.
- **May remain deferred after validated beta:** Broad Go stdlib facades, full native P0/P1, multi-package output, CGo/unsafe breadth, advanced generics, wider platform promises and deep optimization.
- **May remain deferred after stable 1.x:** Experimental Go-native expansion, ecosystem bindings, advanced generics/CGo and generated implementation details outside the explicit stable boundary.
- **Next single action:** Amend `haxe_go-vfp` and `haxe_go-vfp.6` to publish the exact admitted beta surface and remove File/Process/thread/network/Go-native from “supported” until their specific gates pass.

The five consequential actions, in dependency order:

1. Freeze the truthful beta contract and exclusions.
2. Establish supported fail-closed toolchain/security/race gates.
3. Add red tests and fix or exclude build, filesystem, process and concurrency failures.
4. Complete canonical `_std`, deterministic package and same-tested-SHA release/recovery.
5. Run the xhigh second-pass closure and publish the validated beta.

# 13. Optional longer-term improvements

- Broader generated third-party bindings and standard-library facades.
- Wider native platform CI after Linux policy is credible.
- Native representations for more portable strings, collections, options and results.
- Advanced Go generics and constraints.
- Multi-package output when measured import-cycle/build/operability needs justify it.
- CGo and `unsafe` boundary design.
- Richer Haxe-to-Go debugging, stack mapping, profiling and incident tooling.
- Deeper allocation, startup, binary-size and readability optimization.
- Ecosystem integrations beyond the flagship examples.
