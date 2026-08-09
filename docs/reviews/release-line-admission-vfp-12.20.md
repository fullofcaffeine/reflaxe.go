# Release-line admission second pass

## Decision

The original portable-beta review is a historical beta baseline. It approves
the named pre-1.0 release line. It does not approve each later commit.

Each routine release proves its own code through current CI, current governed
support data, current blocker data, current security results, and verified
package bytes. This matches the release-line policy in the sibling compiler
repositories.

## Checks

The second pass challenged these failure paths:

- A moved or renamed baseline tag must fail.
- A release outside the baseline Git history must fail.
- A changed historical review SHA must fail.
- A missing review trigger must fail.
- A routine candidate can differ from the baseline only when all current
  exact-SHA evidence passes.
- The original Oracle SHA must remain separate from the local correction SHA.
- An open applicable P0 or P1 Bead must continue to block release.
- A failed publication must not move, erase, or reuse a public version tag.

The focused mutation tests cover the first six checks. The existing release
reconciliation tests cover immutable tag and hosted-asset behavior.

## Review boundary

The machine gate can require the complete review-trigger list. It cannot infer
the meaning of every source change. The repository work rules and the Bead
risk label own that classification. Critical changes require a fresh review as
part of their implementation task.

No Oracle request was used for this correction. The sibling repositories give
one clear precedent, the failure has one deterministic cause, and the mutation
tests distinguish the old rule from the new rule. A local xhigh second pass is
the lowest sufficient independent check under the project policy.

## Unpublished tag

The failed workflow reserved `v0.54.1` at `e276af1b`. It created no GitHub
Release and uploaded no assets. The tag stays immutable and unpublished. The
project will publish the next eligible patch version from the corrected commit.

## Result

`PASS` for the historical-baseline model, with one explicit limitation: CI
does not classify conceptual risk. The repository work rules and Beads own that
decision before the automated release gate runs.
