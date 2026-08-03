# Independent final review prompt: Haxe.Go portable beta

Paste this prompt into a genuine GPT-5.6-class reviewer and attach the exact final-review packet named in the upload checklist. This is the one independent review required by `haxe_go-vfp.12.5`; it is not a request for another implementation pass or a generic “look for anything” review.

## Review prompt

Act as an independent senior compiler and release-admission reviewer. Review the exact Haxe.Go candidate and frozen evidence described below at the deepest reasoning setting available in your session.

Do not implement changes. Do not rewrite the roadmap. Do not treat green CI as proof by itself. Try to falsify the candidate’s narrow claim, verify its provenance, and return an actionable admission decision that can be reconciled into Beads.

### Review model and provenance

State the exact review model and reasoning setting you actually used. If the serving route is not exposed, say so instead of guessing. If a genuine GPT-5.6-class model is unavailable, stop and report the mismatch; do not silently substitute an older model while claiming the requested provenance.

Distinguish clearly between:

1. source, policy, and artifact facts you personally verified;
2. tests or reproductions you personally reran;
3. implementing-agent reports preserved in the bundle;
4. supplied exact-SHA CI and artifact evidence.

### Primary task

Decide whether the exact candidate supports the stated **Portable beta admission** and whether `haxe_go-vfp.12.5` may close. This is a release-scope and evidence decision, not a stable-1.x, Go-native, public-Internet, every-platform, or full-Haxe-parity review.

Do not require optional roadmap expansion to manufacture a broad readiness claim. A finding blocks this beta only when it applies to an admitted operation/member, declared platform/toolchain, release/security/licensing boundary, or the truthfulness of the published claim.

### Exact source and artifact identity

- Repository: `fullofcaffeine/reflaxe.go`
- Exact source commit: `c22af0ea82e5e481e23277e513ed5b7c6b5c770b`
- Source tree: `fbdb80c6ce39ba8d89c029334452b151a6cce4a3`
- Source parent: `cc40b388779f6bcc265e74e235ed9c929c6bf77c`
- Remote branch recorded at construction: `origin/master`
- Canonical inner evidence ZIP: `haxe-go-portable-beta-c22af0ea.zip`
- ZIP size: 27,529,201 bytes
- ZIP SHA-256: `7d76936576bb9e7e654f73ac010a474305e23b2c08ec2fe2f406441af633eb66`
- Internal payload count: 383
- Internal checksum entries: 384

The authoritative source is `primary/haxe.go-source-c22af0ea.tar` inside the inner ZIP. It is a Git archive whose commit identity must resolve to the exact source commit above. `primary/haxe.go-source-c22af0ea.xml` is a line-numbered Repomix navigation view, not a second source authority.

The outer handoff packet also contains:

- `portable-beta-candidate-c22af0ea.json`, the frozen machine-readable evidence index;
- `portable-beta-candidate-c22af0ea.md`, the friendly explanation of the claim;
- this prompt;
- an outer `SHA256SUMS` file.

The evidence index was committed after the candidate so it could record completed hosted runs and the final ZIP digest. Treat it as an index to verify against the inner artifact, not as candidate source.

### Integrity checks to perform first

Before reaching a verdict:

1. Verify the outer packet checksums.
2. Verify the inner ZIP SHA-256 and size.
3. Run or independently inspect `unzip -t` for the inner ZIP.
4. Verify every inner `SHA256SUMS` entry.
5. Confirm `git get-tar-commit-id` for the source archive returns the exact candidate commit.
6. Confirm `MANIFEST.json` names the same source, five CI runs, sibling commits, roadmap snapshot, and builder hash as the evidence index.
7. Confirm `validation/gitleaks.log` and `validation/local-paths.log` passed.
8. Treat the nine paths listed as Repomix security exclusions as XML-view omissions only. Confirm their raw companion files and authoritative Git-archive copies exist before calling anything missing.

If any identity or checksum fails, stop with `NOT_READY: PROVENANCE FAILURE` and explain the mismatch.

### Narrow public claim under review

The proposed claim is:

> Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads on the admitted toolchain, platform, and operation/member surface.

Its exact boundaries are:

- preset/scorecard: `portable`;
- platform: Linux/amd64;
- Haxe: 4.3.7;
- Go: 1.25.12 or 1.26.5;
- admitted operations: 38;
- admitted named symbols: 173;
- operation/member surface SHA-256: `99625a5bcb401561a8393ddcef5675ba552c1a84ab288ea4ed0e1cc950bac0d0`;
- trust model: `trusted-source-only`;
- default rule: every unlisted module operation, member, platform, architecture, native API, error path, concurrency pattern, or trust model is excluded.

The compatibility manifest—not prose—is authoritative for which operation/member subsets are admitted.

### Explicit non-admissions

- **Go-native remains excluded** behind `haxe_go-vfp.9.1`.
- **HTTP remains excluded** behind `haxe_go-vfp.10.8`, including `haxe.Http`/`sys.Http` admission.
- Stable 1.x remains excluded behind `haxe_go-vfp.12.7`.
- Untrusted or adversarial source compilation is excluded; Haxe.Go is not a sandbox for macros, plugins, dependencies, metadata, or selected child tools.
- Windows, Darwin, other operating systems, non-amd64 architectures, IPv6, public-network/hostile-peer guarantees, and every unlisted operation remain excluded.
- `metal` is a compatibility policy preset, not a second semantic product. Portable and metal results are separate scorecards; metal success cannot advance the portable claim.

Some named file, process, thread, terminal, TCP, UDP, and TLS operation subsets are admitted. Do not replace the precise manifest with a blanket “all sys/network is excluded” statement, and do not widen those subsets into whole-module support.

### Exact hosted evidence

All selected workflows completed successfully on the exact candidate source commit:

- Examples Artifacts: run `30743182325`;
- Security - Gitleaks: run `30743182328`;
- Security - Static Analysis: run `30743182367`;
- CI Harness: run `30743182369`;
- CI - Quality: run `30743182373`.

The inner bundle preserves sanitized logs, job graphs, and metadata for 20 hosted artifacts, including artifact names, sizes, SHA-256 digests, and the common expiry `2026-10-31T10:09:21Z`.

Check what each job actually ran. In particular, verify that the evidence supports:

- 313 compiler/generated-shape snapshot owners;
- 155 portable semantic-diff fixtures;
- the complete applicable official Haxe target inventory of 55 strict upstream modules on both supported Go versions;
- representative cold Haxe-through-custom-backend compilation, strict generated-Go build/test, execution, and observed output;
- 8 maintained example program directories, 13 profile cases, 1 release-claim-bearing case, and independent portable/metal scorecards;
- generated runtime, race, checkptr, resource, terminal, file, process, thread, socket, UDP, and TLS contracts applicable to admitted members;
- fatal post-generation Go build failures;
- deterministic Haxelib packaging and isolated package installation;
- supported-Go tooling, dependency/vulnerability audit, static analysis, Gitleaks, and raw-injection hygiene;
- license/notice, public API, compatibility, SemVer, release-readiness, same-SHA, retry/idempotence/repair, and provenance contracts.

Do not count raw test files, snapshots, or compiled modules without checking their owner and claim level. Cold proof remains required; compiler-server/warm behavior is only an accelerator.

### Independent oracles and test sensitivity

Assess whether material expectations come from an independent authority: Haxe 4.3.7 behavior, manually authored expected output, pinned upstream inventory, Go/runtime behavior, an invariant/metamorphic property, reviewed golden provenance, or real downstream execution.

Flag any case where:

- the implementation generated its own expected value;
- generated text inspection is used to claim runtime behavior;
- a metal/native success is laundered into portable evidence;
- one platform advances another platform’s claim;
- a high-level regression has no faithful focused owner even though one is feasible;
- selector/backstop logic can omit claim-bearing evidence without a cold full lane detecting it;
- retry or quarantine policy can turn a product failure into apparent success.

### Prior-review reconciliation

Every applicable prior finding must be traced to one of:

- fixed and covered by exact-candidate evidence;
- still excluded by the current compatibility contract with a durable owner;
- intentionally deferred because it does not apply to the admitted beta;
- false positive or superseded, with evidence.

At minimum inspect and reconcile:

- `docs/reviews/gpt-5.6-pro/review-cd79624f.md` and its provenance record;
- `docs/reviews/network-admission-oracle-disposition-vfp-10.4.md`;
- `docs/reviews/gpt-5.6-pro/review-21acb7eb-socket-admission.md` and its provenance record;
- the exact compatibility source/manifest, release status, known gaps, public contract, licensing policy, release-readiness policy, roadmap snapshot, and relevant closed/open Beads in the bundle.

Do not infer closure from a closed Bead alone. Check the named code, policy, test, and exact-SHA evidence. Conversely, do not reopen an old finding merely because its original wording was broad if the affected operation is now truthfully excluded or was split into precisely admitted members with decisive correction evidence.

### Release, licensing, and provenance questions

Verify:

1. the license policy is approved by Marcelo Serpa under the recorded project-copyright-owner authority and scope digest `56e4dc13b3676562cd2ed6ac23f1706141a2098fdd67f408ceba026d84848c8a`;
2. generated user programs may use their own chosen license while required MIT notices for compiler-emitted/runtime/stdlib-derived material are emitted and packaged;
3. package contents, checksums, manifest, provenance statement, isolated install, and hosted-asset reconciliation fail closed;
4. the same-SHA workflow does not publish a partially verified release and has bounded retry/idempotence/repair semantics;
5. the evidence task itself made no publication and does not claim the old `v0.53.1` tag is this candidate;
6. current host controls and release/tag metadata are treated as dated live evidence, not silently inferred from source;
7. no unresolved P0 or P1 release, security, licensing, or semantic issue applies to the admitted beta.

### Architecture scope

This is not a request for a new compiler architecture program. Check only whether a current architectural defect makes the admitted beta unsafe or unmaintainable enough to block it.

The intended direction is:

- staged Haxe owns Haxe-visible stdlib behavior;
- typed `hxrt` owns Go/OS resources and narrow native capabilities;
- compiler lowering remains AST-first;
- portable semantics are default;
- typed `go.*`, externs, and `@:goNative` boundaries carry native authority;
- `metal` remains a convenience policy preset;
- targeted typed seams are preferred over a universal compiler IR.

Do not demand a universal IR, a portable/metal semantic fork, full Go syntax coverage, every standard-library module, or all operating systems unless you demonstrate a concrete admitted-beta failure that cannot be fixed at the existing typed seam.

### Findings and severity

For each finding provide:

- severity: `BLOCKER`, `HIGH`, `MEDIUM`, or `LOW` relative to this exact beta claim;
- exact source/policy/test/log anchor;
- concrete failure or misleading claim;
- why existing evidence would not catch it;
- smallest root-cause correction or narrow exclusion;
- decisive verification;
- whether it blocks portable beta, bounded production, stable 1.x, or only an excluded/optional surface.

A P0 or P1 issue blocks this beta only if it applies to the admitted contract. Do not convert optional Go-native, stable, performance, output-quality, or platform expansion into a beta blocker.

### Required output format

Return the following sections in order.

1. **Review model and provenance**
   - exact model, reasoning setting, serving route if actually exposed;
   - exact source and artifact identity verified;
   - what you personally inspected/reran versus supplied evidence.

2. **Integrity verdict**
   - `PASS` or `FAIL`, with checksum/source/run mismatches if any.

3. **Portable beta admission**
   - exactly one: `READY`, `READY_AFTER_BOUNDED_FIXES`, or `NOT_READY`.

4. **Bounded-production verdict**
   - exactly one: `READY_WITH_LIMITS` or `NOT_READY`;
   - state the concrete acceptable workload and boundaries in plain language.

5. **Stable-1.x verdict**
   - exactly one: `READY` or `NOT_READY`;
   - do not block a truthful 0.x beta merely because stable 1.x is `NOT_READY`.

6. **Architecture verdict**
   - `SOUND_FOR_ADMITTED_BETA`, `SOUND_WITH_FOLLOW_UPS`, or `BLOCKING_DEFECT`;
   - say whether any compiler-wide seam change is actually required.

7. **Findings**
   - ordered by severity with all fields above.

8. **Prior-finding disposition table**
   - Every applicable prior finding, source review, current state, evidence/owner, and whether it affects beta.

9. **Operation and policy audit**
   - confirm or correct 38 operations, 173 symbols, surface digest, toolchains, platform, trust boundary, exclusions, and portable/metal separation.

10. **Residual exclusions**
    - exact list that must remain outside the beta promise.

11. **Beads disposition**
    - whether `haxe_go-vfp.12.5` may close;
    - bounded follow-up Beads only for genuinely new material gaps, without inventing duplicate IDs;
    - whether publication task `haxe_go-vfp.12.6` may proceed.

12. **Final checklist**
    - integrity, evidence, claim, findings, and next action.

Do not return `READY` solely because CI is green. Do not return `NOT_READY` solely because optional roadmap work remains. The decision must follow the exact admitted claim and its evidence.

## Upload checklist

Attach exactly one outer packet containing:

- this prompt;
- `portable-beta-candidate-c22af0ea.json`;
- `portable-beta-candidate-c22af0ea.md`;
- `haxe-go-portable-beta-c22af0ea.zip` with SHA-256 `7d76936576bb9e7e654f73ac010a474305e23b2c08ec2fe2f406441af633eb66`;
- outer `SHA256SUMS`.

Do not substitute a working-tree archive, an older Repomix output, a previous network-only review packet, or a ZIP built from a different basename/input set.
