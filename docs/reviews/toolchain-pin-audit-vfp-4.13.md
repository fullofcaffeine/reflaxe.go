# Exact Go Toolchain Pin Second-Pass Review

Status: accepted for landing

Date: 2026-07-14

Bead: `haxe_go-vfp.4.13`

Reviewed base: `3722f7003deb4b9b3b6e7f67b4d38313cb0c49e1`

## Review Scope

This is the explicit written second pass required for a `thinking:xhigh`
release-enforcement change. It reviews the move from wildcard Go selectors to
exact approved patch versions, the policy schema transition, workflow and
release-status enforcement, generated compatibility evidence, and the
security evidence produced by the pinned toolchains.

It does not change the generated Go language floor (`go 1.22`), compiler
semantics, portable admission, the portable/metal policy-preset contract, or
the supported Go release lines (`1.25` and `1.26`).

## Trigger And Invariant

The `Security - Static Analysis` workflow for the reviewed base failed closed:

- `1.25.x` resolved to Go 1.25.11;
- `1.26.x` resolved to Go 1.26.4; and
- both dependency-audit jobs reported reachable `GO-2026-5856` findings in
  `runtime/hxrt` call paths.

The [official Go release history](https://go.dev/doc/devel/release) identified
1.25.12 and 1.26.5 as the current patches. The
[GO-2026-5856 advisory](https://pkg.go.dev/vuln/GO-2026-5856) identifies those
versions as the fixes.
The violated repository invariant was therefore not "security should pass";
the fail-closed audit behaved correctly. The violated invariant was that the
latest approved patch in each supported line must produce CI and release
evidence. A wildcard satisfied by a stale runner cache did not enforce that
claim.

## Alternatives Adjudicated

1. Keep `.x` selectors with the default setup-go behavior — rejected. A stale
   matching tool-cache entry can satisfy the range, reproducing the incident.
2. Keep `.x` selectors and set `check-latest: true` — rejected for release
   evidence. It seeks a newer match but permits the concrete build toolchain to
   change without a reviewed repository change.
3. Pin exact patches in workflows only — rejected. It would leave the machine
   policy and generated compatibility evidence unable to explain or enforce
   the pins.
4. Pin exact patches in a versioned policy, consume them in workflow contract
   tests and release-status checks, and refresh them deliberately — accepted.
   It preserves stable support-line policy while making current evidence
   reproducible.

## Second-Pass Findings

### Accepted And Resolved

- `docs/toolchain-policy.json` schema 2 separates `supported_build_lines` from
  exact `ci_versions` and `recommended_build_version`.
- All setup-go consumers in quality, harness, security, and example-artifact
  workflows use either 1.25.12 or 1.26.5 as governed by the policy.
- Policy contracts reject supported-line `.x` selectors, validate exact
  `major.minor.patch` shape, map one CI version to each supported line, and
  require the recommended exact version to be a governed CI version.
- `check-release-state.sh` verifies both the recommended version and the exact
  two-version matrices in harness and security workflows. It no longer derives
  an unpinned selector from `recommended_build_line`.
- The compatibility manifest embeds policy schema 2 and its refreshed digest.
- The focused red phase failed on schema 1, missing exact-version fields, old
  workflow selectors, and old release-status wiring before implementation.
- The full release-contract run exposed a separate stale five-scope assertion
  after the concurrency tooling expansion. `haxe_go-vfp.4.4.1` records that
  regression; the corrected test now proves the complete seven-target by
  four-gate product rather than a magic total.

### Residual Risk And Control

Exact pins require an intentional repository update when Go publishes a new
patch. That is a feature for provenance, but it creates maintenance work. The
weekly fail-closed vulnerability job, official-source update procedure, and
policy/workflow drift tests are the control: a newly vulnerable pin blocks
release until the reviewed version transition lands.

No unresolved P0 or P1 design finding remains in this change.

## Validation Evidence

- Go 1.25.12 dependency audit: npm production audit clean; govulncheck v1.6.0
  reported no vulnerabilities.
- Go 1.26.5 dependency audit: npm production audit clean; govulncheck v1.6.0
  reported no vulnerabilities.
- Go 1.25.12 tooling gates: 7 scopes by 4 gates, 28/28 passed with no retries.
- Go 1.26.5 tooling gates: 7 scopes by 4 gates, 28/28 passed with no retries.
- Release contracts: passed after resolving the separately tracked stale-scope
  assertion.
- Release status: passed with exact policy wiring reported.
- Full snapshots: 248 passed, 0 failed.
- Compatibility generator check and diff whitespace check: passed.

## Oracle Decision

GPT-5.6 Pro remains unavailable because of the recorded Codex account limit
until 2026-07-19 12:58 PM. The installed fallback maps the requested model to
an older `gpt-5-pro`, so using it would misstate review provenance and remains
rejected.

An Oracle review is not a blocking requirement for this bounded transition.
After local tracing, only one design satisfies all three requirements at once:
fail-closed patch safety, exact release provenance, and no silent toolchain
drift. The project instructions explicitly permit this written second pass as
the `thinking:xhigh` fallback. A later genuine GPT-5.6 Pro review may audit the
broader release program, but it is not needed to re-decide these exact pins.

## Verdict

Accept. The change restores green security evidence without weakening the
gate, makes the version claim reproducible, preserves the supported-line and
generated-language-floor boundaries, and adds sufficient drift detection for
the maintenance risk it introduces.
