# Written second-pass review: stdlib ownership closeout

## Review identity

- Bead: `haxe_go-vfp.8.7`
- Baseline: `a48b0628`
- Implementation: `32fc526f368f2ca295d09ab60b0d4d9c55a45d8a`
- Review hardening: `cd9bd583ae550fdd4035ce3038c74eaecaf7d7ce`
- Final reviewed range: `a48b0628..cd9bd583`
- Fresh-eyes reviewer: GPT-5.6 Terra, read-only review
- Review form: explicit written second pass, not Oracle and not GPT-5.6 Pro

Oracle escalation was not necessary. Local tracing left one defensible bounded
design, and exact generated-output and runtime contracts can decide whether the
selector still behaves correctly. Oracle remains appropriate when deeper local
work leaves competing designs unresolved; that condition did not occur here.

## What was reviewed

The parent migration moved Haxe-visible standard-library behavior to upstream
Haxe or canonical target source under `std/go/_std`, with narrow typed `hxrt`
capabilities for native runtime work. This pass checked the final compiler
state rather than trusting that all child Beads being closed proved the parent.

The remaining shared selector reaches:

- portable Type metadata;
- portable Reflect metadata;
- the same-package serialization invocation bridge;
- explicit `go.*` concurrency declarations;
- explicit `go.*` collection declarations; and
- explicit `go.*` result declarations.

The first three consume final compiler metadata or generated-package access
that ordinary staged source and the separate runtime package cannot reproduce.
The last three implement explicit Go-native APIs, not portable Haxe stdlib
behavior. None of the six dispatch branches contains a library algorithm.

## Findings and disposition

### 1. Missing commit-pinned closeout evidence — blocking, resolved

The first review correctly rejected closure while the decision trail still said
“review pending” and the Bead had no comment tied to the implementation commit.
This document adjudicates the review against the final range. The same commit
identities and validation results are recorded in the Bead before closure.

### 2. Validation claims were not yet recorded in Beads — medium, resolved

The implementation commit passed:

- `npm run test:changed`;
- `npm test`: 301/301 snapshots;
- `npm run test:semantic-diff`: 140/140 portable behavior contracts;
- `npm run test:stdlib-sweep:go-test`: 55/55 strict stdlib modules;
- `npm run test:examples`: 12/12 runnable profile lanes;
- `npm run test:release-contracts`;
- `npm run test:perf:hxrt-selective`; and
- focused registry, migration-ledger, compiler-debt, and Markdown checks.

After review hardening changed the gate, the seven registry tests,
compiler-debt gate, Markdown-link contract, and full release-contract suite
passed again. The parent Bead comment records these results rather than relying
only on prose in the migration log.

### 3. The 10-to-9 change is a metric correction — medium, resolved

No declaration emitter was deleted by the closeout commit. The old
`lowerStdlibShimDecls` function was structural selection plumbing, so the
compiler-shim metric no longer counts it as though it emitted a tenth shim.
All nine actual declaration emitters remain independently registered,
classified, and ratcheted. The rationale, migration log, and decision trail now
state this explicitly.

### 4. Pair comparison could miss unrelated dispatcher work — medium, resolved

The original registry test compared the six recognized key/emitter pairs but
could ignore an unrelated statement inserted beside them. A new test first
proved that gap by failing against an injected unregistered declaration. The
parser now accepts only:

1. creation of the declaration accumulator;
2. registered capability-key checks;
3. concatenation of registered `lower*ShimDecls` emitters; and
4. return of the accumulator.

Any extra statement, unrecognized body shape, duplicate key, missing
registration, or extra registration fails the normal, changed, governance, and
release paths.

### 5. Runtime documentation described old migration entries — low, resolved

`docs/hxrt-runtime.md` now states that no `migration_required` group remains
and that adding one fails the closeout gate. It no longer describes historical
migration debt as current registry state.

## Acceptance review

- Public stdlib behavior is owned by upstream or staged Haxe source: pass.
- Runtime-dependent work crosses narrow typed `hxrt` boundaries: pass.
- Remaining compiler intrinsics are exact, registered, documented, and tested:
  pass.
- Explicit Go-native emitters are separated from portable stdlib ownership:
  pass.
- No `migration_required` registry entry or migration-debt exception remains:
  pass.
- The dispatcher is structural and fail-closed, while actual emitters remain
  measured: pass.
- Snapshots, semantic parity, strict stdlib, examples, release, and relevant
  selective-runtime gates pass: pass.

## Non-blocking follow-up

Internal names such as `requiredStdlibShimGroups` still reflect migration-era
terminology. Renaming or replacing that request-local state should happen with
the typed ownership extraction in `haxe_go-vfp.8.6`, where module APIs and
dependency direction can change together. A broad terminology rewrite here
would add churn without strengthening the ownership boundary.

## Verdict

Pass after the review hardening in `cd9bd583`. No unadjudicated finding blocks
closure of `haxe_go-vfp.8.7`.
