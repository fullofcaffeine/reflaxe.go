# Behavior-first testing strategy

Haxe.Go now tells contributors two separate things clearly: whether a change
still behaves correctly, and which part of the product that result is allowed
to support. This prevents a passing example, native Go test, or generated-file
snapshot from being mistaken for proof of an unrelated compatibility promise.

## Start here

While editing, use the smallest command that exercises the behavior you changed:

```bash
# Quick real application path: Haxe -> generated Go -> Go test -> program output
npm run test:smoke

# Focused repository checks selected from your local Git changes
npm run test:changed

# Explain which broader tests own the change; this does not skip any CI work
npm run test:affected:explain
```

Before a pull request or handoff, run the complete local CI surface:

```bash
npm run test:ci
```

The practical flow is:

```text
describe the expected behavior -> make the smallest faithful test fail ->
fix it -> run a real Haxe-to-Go path -> run the complete backstop
```

The focused test makes failures faster to diagnose. The real path proves the
compiler, generated Go, Go toolchain, and runtime still work together. The
complete backstop catches mistakes in the focused-test selector.

## Terms used in this guide

- A product surface is one independently promised part of Haxe.Go, such as
  portable Haxe semantics, explicit Go-native APIs, runtime/stdlib behavior,
  diagnostics/tooling, or packaged examples.
- A scorecard is the evidence record for one product surface. It says what is
  claimed, which tests support that claim, and which risks remain.
- An oracle is an expectation that does not come from the implementation being
  tested—for example, Haxe `--interp`, a specification, reviewed output, or a
  real downstream program.
- A tracer bullet is one small end-to-end example that proves the whole path
  works before many narrower cases are added.
- Feedback rings are the R0-through-R5 test levels, ordered from one focused
  local check through full release proof. A higher ring is broader and slower;
  it does not replace the faster rings.
- An owner is the stable test, fixture, or command responsible for detecting a
  particular behavior regression. It is not necessarily a person.

## Outcome and authority

Haxe.Go has five independently claim-bearing product surfaces. One green surface cannot advance another surface's claim. In ordinary terms: a passing
Go-native channel example cannot prove portable Haxe semantics, and a portable
snapshot cannot prove native Go API behavior.

The machine-readable scorecards live in
[`test/testing-strategy.json`](../test/testing-strategy.json); this document
explains what it means. Public release admission still comes only from
[`docs/compatibility-support-manifest.json`](compatibility-support-manifest.json)
and the release-readiness policy. This strategy does not broaden any public
claim.

The updated consolidated-testing guidance was an incremental audit, not a
reason to replace the existing harnesses. The repository already had strong
snapshot, semantic-diff, staged-stdlib, native Go, package-install, example,
security, performance, and release gates. The missing pieces were explicit
surface scorecards, executable example tiers, durable behavior/oracle/red
provenance, and semantic-owner affected-selection evidence.

## Product scorecards

| Surface | Current status | What can make it greener | What cannot |
| --- | --- | --- | --- |
| `portable-compiler` | Admitted slice | Ordinary Haxe semantics, generated Go build/run, exact compatibility operation evidence, the required three-family official target smoke, the extended classified official inventory, and portable-semantic preset invariance under `metal` where no native boundary is used | Go-native APIs, example-only or compile-only evidence, another platform's result, blocked official owners, or treating inventory coverage as automatic compatibility admission |
| `go-native-metal` | Experimental | Typed `go.*`, extern, ABI, policy, Go tooling, and native runtime owners | Portable semantic diff or selecting the `metal` preset by itself |
| `runtime-stdlib` | Admitted slice | Staged-source semantics, focused `hxrt` ownership tests, generated-Haxe execution, race/checkptr, and exact platform evidence | A snapshot, cross-build, or whole-class name without member-level evidence |
| `diagnostics-tooling` | Partial | Exact negative diagnostics, stable report schemas, output confinement, and developer-command contracts | A successfully running generated program; source maps remain unclaimed |
| `package-downstream-examples` | Partial | Isolated release ZIP install and examples executed at their declared tier | Generated text alone or a cross-built binary that did not run on its target OS |

Every full scorecard records claims, inputs, outputs, profiles, focused and
vertical owners, examples, oracle/provenance, skips, quarantines, selectors,
backstops, last clean proof, and residual risks in the machine authority. The
official Haxe target-suite contract applies only to `portable-compiler`.
That scorecard lists both public presets because ordinary portable source also
runs under the compatibility `metal` preset. A `metal` result contributes only
portable preset-invariance evidence unless its lane separately declares the
`go-native-metal` surface; it never turns preset selection itself into native
API evidence. The strategy contract rejects any claim-bearing example whose
profile is absent from a declared surface's supported or tested profiles.

## Audit of the new conclusions

| Conclusion | Before this task | Incremental disposition |
| --- | --- | --- |
| Concrete behavior discovery | Partial: Beads often had strong acceptance criteria, but there was no shared scenario shape | Required fields and a real socket lifecycle scenario are now durable |
| Red -> green at the lowest faithful layer | Partial: AGENTS required TDD, but red evidence was not uniformly surfaced | Exact pre-fix command and intended failure are required; the real socket repair is the worked example |
| Independent oracle | Partial: semantic diff and Oracle reviews were strong, but example expectations had no tiered provenance record | Scorecards and the example manifest name their oracle and provenance |
| Tracer bullet first | Partial: examples and semantic diff already crossed real boundaries, but the policy was implicit | R1 and the representative workflow name the tracer before permutations |
| Lowest faithful layer and double lock | Substantially present in recent network work | Focused Go ownership tests, generated-Haxe execution, and release-policy contracts are kept as distinct locks |
| Portfolio review | Absent as a durable review | Active owner-family inventory is recorded without inventing a quota or double-counted percentage |
| Executable example tiers | Runtime execution already worked; tiers and claim ownership were implicit | `examples/qa-manifest.json` is harness-enforced; exact release-operation links prevent broader QA examples from widening the beta claim |
| R0-R5 loop and affected selection | Partial: commands and full backstops existed, but ring ownership and semantic selection did not | Rings are explicit and the selector runs in observation mode without skipping evidence |
| Separate high-risk review | Strong through thinking labels and exact-commit reviews | A stable challenge checklist now lives in the strategy authority |

## Form behavior before automating it

For meaningful behavior, record this before broad fixture expansion:

| Field | Question |
| --- | --- |
| Preconditions/input | What source, state, profile, platform, or native resource matters? |
| Action/path | Which Haxe compile, generated-Go build, package, or runtime path executes? |
| Observable result | What output, diagnostic, target form, resource effect, or user behavior is visible? |
| Error/edge result | What must fail, remain open/closed, time out, or preserve partial progress? |
| Product surface | Which one scorecard owns the result? |
| Protected claim | What exact statement becomes less trustworthy if this test disappears? |

Beads acceptance criteria, a fixture manifest, or a small scenario table are
enough. Gherkin/Cucumber is not required.

## TDD, oracle, and snapshot policy

Start at the lowest faithful owner. Record the pre-fix command and concise
failure, make it green, refactor, then run the next real boundary. A separate
red commit is optional, but the test must be red for the intended reason.

An independent oracle may be a Haxe/Go specification, a manually authored
minimal result, pinned `--interp` differential behavior, an invariant, a
reviewed golden with provenance, or a real consumer. The production algorithm
must not manufacture its own expectation. Snapshot refreshes require semantic
review; generated Go is paired with `gofmt` and `go test`, and with runtime
execution whenever runtime behavior is the claim.

## Representative behavior-first workflow

The real recent example is the socket admission repair, not a fabricated demo.

Scenario: an in-progress or failed socket operation owns a native resource.
Closing the handle, failing policy installation/handshake, or selecting with a
finite timeout beside an entered read must not publish or retain stale state.
The exact resource must close, the original typed failure must escape, and a
stale failure must not detach a replacement.

The focused tests were red for the intended reason at pre-fix commit
`21acb7eb8d0c1ce23b4b2d2bde095d1e538d2962`:

```bash
cd runtime/hxrt
go test -run 'TestSocketListenDeadlineFailureLeavesAnEmptyReusableHandle|TestSslSocketHandshakeFailureDetachesAndClosesTheFailedConnection|TestSocketSelectTimeoutIsBoundedBehindAnEnteredRead' -count=1 ./...
```

They exposed stale listener state, an attached failed TLS connection, and a 20
ms select delayed beyond 150 ms. The durable red record is
`.audit/haxe_go-vfp.10.9.8.tsv`; the independent oracle is the exact-commit
GPT-5.6 Pro review and its provenance under
`docs/reviews/gpt-5.6-pro/review-21acb7eb-socket-admission*`.

The lowest faithful owner is the focused `runtime/hxrt` failure-injection test.
The tracer bullet is `test/semantic_diff/socket_server_lifecycle_contract`:
authored Haxe -> Reflaxe.Go -> generated Go -> formatting/build -> loopback
runtime observation. Operation-level compatibility and release contracts form
the third lock. None of those results is substituted for another.

## Portfolio guardrail

The current active inventory is 315 snapshot owners, 155 portable semantic-diff
owners, one active profile/lane semantic-diff owner, and 13 executable example
profile owners. The strict curated stdlib sweep has 55 module owners and the
full portable-eligible inventory has 175. The separately pinned official target
inventory classifies 1,211 upstream owners: 607 active owners containing 650
runtime methods, 480 blocked owners, and 124 inapplicable owners. These are
active harness records and explicit classifications, not raw file counts.

This is compatible with a focused-heavy compiler portfolio and substantial
real Haxe -> Go execution, but publishing a precise 55-70/25-40 percentage now
would double-count overlapping owners. First give overlapping runtime,
semantic-diff, and example evidence stable non-overlapping identities. Static,
formatting, schema, freshness, policy, security, and provenance checks stay
outside the ratio. Review unique failure yield, escaped defects, diagnosis
time, selector misses, and claim coverage—not assertion volume.

## Executable examples

[`examples/qa-manifest.json`](../examples/qa-manifest.json) declares every
maintained example as a `flagship-application`, `capability-showcase`, or
`compile-only-snippet`, plus profiles, per-execution-lane product surfaces and
evidence modes, distinctive claim, command, and oracle. Per-lane declarations
matter because a flagship's default portable/core lane must not become native
evidence merely because its CI variant uses an explicit Go-native module. The
harness rejects manifest drift.

Ordinary claim-bearing and release-bearing are separate. Every maintained
example below proves its own documented runtime behavior. Only `portable_beta`
publishes portable-beta operation evidence, and only after the harness resolves
each declared operation against the compatibility authority and completes the
real runtime/output path. This keeps a green flagship, HTTP service, or native
interop demo from broadening a different scorecard.

All current examples are claim-bearing and therefore execute:

```text
Haxe source -> custom backend -> gofmt -> go test ./... -> go run -> reviewed stdout
```

`--compile-only` remains a fast diagnostic mode, but it cannot support runtime
claims. A metal lane is native evidence only when its source uses an explicit
native boundary; `incident_api/metal`, for example, is portable preset-
invariance evidence, not Go-native evidence.

Uploaded example telemetry records declared surfaces separately from completed
claim evidence. It activates a lane's claim only after the real runtime and
reviewed stdout oracle pass, and clears every claim field when the enclosing
case fails. Failure artifacts remain useful without becoming green scorecard
entries.

## Feedback rings

| Ring | Current command/owner | Evidence rule |
| --- | --- | --- |
| R0 focused | one snapshot, semantic-diff case, negative, or focused Go test | Active red/green owner; smallest faithful layer |
| R1 local smoke | `npm run test:smoke` | Regenerates `tui_todo`, formats, strictly Go-tests, runs, and checks output |
| R2 required PR | `npm run test:ci` plus all required workflow jobs | Clean full primary evidence remains unchanged |
| R3 affected extended | `npm run test:affected:explain` | Observation only; recommends owners/surfaces but skips nothing |
| R4 main/weekly | CI Harness plus Quality on master; weekly CI Harness subset with the supported-Go official inventory, security, tooling, performance, and example artifacts | Complete stable-graph selector-miss backstop weekly; the cross-platform Quality matrix also runs on pull requests and master pushes |
| R5 release | `npm run release:readiness`, `npm run release:dry-run`, and the supported-Go official inventory dependency | Exact SHA/artifact, clean install, supported matrices, and public claims |

The current remote graph deliberately uses cold setup and disables the Go
action cache. The managed `npm run dev` loop may reuse an owned Haxe compiler
server because that makes edits faster, but it is only an accelerator. It is not
release evidence. `dev:hx`, `npm run dev -- --once`, CI, and release builds keep
the direct cold path as the correctness baseline. The lifecycle and fallback
contract is documented in [Development watch loop](development-watch-loop.md)
and owned by completed tracker item `haxe_go-vfp.12.8`.

## Affected-selection safety

`test/run-test-plan.py` unions committed-base, unstaged, staged, renamed,
deleted, and untracked paths. It maps them to semantic owners and independent
product surfaces, prints commands and reasons, and conservatively expands an
unknown or cross-cutting path to every owner. Always-run sentinels protect the
strategy contract, real portable tracer, failure propagation, and portable
provenance.

Git discovery also fails closed. If Git is missing, a base revision is invalid,
or any discovery command fails, changed snapshots expand to every active
snapshot and the affected-plan explanation expands to every owner. An empty
selection is allowed only when discovery itself succeeded and found no changes.

This is observation mode. `npm run test:ci` still runs the complete gate, so no
claim-bearing evidence moved out of the PR path. After a useful observation
window, a separate Bead may promote selected work only with selector-miss data,
a final required aggregator, and the R4 full backstop intact.

## Flakes, retries, caches, and review

- Deterministic compiler, target-build, assertion, and runtime failures fail
  immediately; retries never erase the first failure.
- Quarantine requires an owner, Bead, reason, evidence, and expiry and never
  contributes a compatibility pass. There are no strategy-level quarantines.
- Caches accelerate only. Missing caches must pass, and R2/R4/R5 retain cold
  proof. Generated evidence artifacts are bound to exact source and toolchains.
- Compiler representation, runtime/ABI, package publication, security,
  migration, and claim changes receive a review pass distinct from
  implementation. It challenges red sensitivity, oracle independence,
  negatives, mocked boundaries, selector omissions, evidence laundering, and
  overbroad wording.

## Baseline measurements and current topology

Measurements on 2026-07-31/2026-08-01 used Haxe 4.3.7, local Go 1.25.6
darwin/arm64, Node 20.19.3, npm 10.8.2, and Python 3.14.5. They are initial
samples, not p50/p95 promises.

| Path | Before cold/warm | After cold/warm | Interpretation |
| --- | ---: | ---: | --- |
| R0 `core/hello_trace` snapshot | 5.74 / 1.28 s | 6.68 / 1.57 s | Same regenerated case; single-sample variation with no topology change |
| R1 `tui_todo/portable` tracer | 8.62 / 4.02 s | 9.97 / 4.55 s | Same full custom compiler -> Go test/run chain, now reached through the named smoke command |
| R2 exact-SHA remote required graph | 28m 44s critical path | 28m 50s quality gate | Commit `0c51420c`; [CI Harness run 30713172868](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30713172868) passed with clean jobs and caches disabled |
| R3 full portable semantic diff | 392.13 / 334.19 s | unchanged command | 155/155 real differential owners; cold sample followed by an immediate warm rerun |
| New strategy contracts | n/a | 1.21 s before this readability pass | Focused contracts cover active-inventory derivation, fail-closed Git discovery, evidence-state reporting, telemetry wiring, official-smoke policy, and first-read documentation; they run in required local/full entrypoints |
| New affected-plan explanation | n/a | 0.27 s | One runtime path; observation only and does not run or omit selected commands |
| New official Haxe target smoke | n/a | 46.13 / 22.05 s | Three active official methods, 16 live assertions, compile-directory-verified installed package, full DCE, strict Go build/run, and five deliberate failure controls, including a real generated-target timeout and an actual missing selected source |
| Extended official Haxe target inventory | n/a | 141.36 / 156.52 / 152.09 s clean-package samples | 1,211 exact owners classified; 650 active methods and 1,910 assertions execute through the installed compiler, strict Go tooling, and runtime. The 480 blocked owners include seven compile-only regressions that need a separate faithful observer and one async HTTP environment gap. The cache-warm-capable samples varied, so these are observations rather than a latency promise. This is R4/R5 evidence, not required-PR work. |
| R4 exact-SHA Quality matrix | 27m 45s Linux Go 1.25.12; 22m 31s Linux Go 1.26.5; 22m 06s macOS | 24m 38s Linux Go 1.25.12; 29m 19s Linux Go 1.26.5; 22m 55s macOS | Commit `0c51420c`; [Quality run 30713172877](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30713172877) passed all three cold hosted jobs |

The strategy and selector contracts add under two seconds locally, and the new
required official target smoke adds about 22–46 seconds on this machine. It removes
no prior evidence. The real tracer also exposed and locked three compiler
contracts that repository-authored tests had missed: packaged `target.threaded`
declaration, optional primitive function-carrier ABI consistency for null and
non-null defaults, and retained
superclass carrier emission under full DCE.

## Claims and deferred work

Claims justified by pre-existing execution remain justified, but are now
attached to the correct scorecard. No compatibility claim is broadened. The
portable claim remains the exact generated manifest slice; native, runtime,
tooling, example, OS, and release evidence remain independent.

Deferred work is tracked in Beads rather than hidden here:

- `haxe_go-vfp.12.10` establishes the classified extended inventory and its
  R4/R5 evidence without broadening compatibility admission;
- `haxe_go-vfp.12.12`: burn down the 480 exact official-inventory blockers;
  its children separately own Haxe compilation, invalid generated Go, runtime
  semantic mismatches, and typed harness-support gaps;
- `haxe_go-vfp.16`: graduate affected-test planning only after an observation
  window establishes selector value and misses;
- `haxe_go-vfp.14`: complete local changed-snapshot discovery, fixed and closed
  with this strategy change.

The managed watch-loop work in `haxe_go-vfp.12.8` is complete: it has a cold
equivalent, a real Haxe-to-Go compile/build/run tracer, failed-build recovery,
owned process/server cleanup, and installed-package delivery. Its green result
belongs to the diagnostics/tooling scorecard and does not broaden portable
compiler, runtime/stdlib, metal/native, or release claims.
