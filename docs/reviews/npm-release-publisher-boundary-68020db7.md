# npm Release Publisher Boundary Review

Date: 2026-07-26

Bead: `haxe_go-vfp.4.14`

Reviewed commits: `61af4e3e`, `68020db7`

Reviewer: fresh cross-model `gpt-5.6-terra` pass at xhigh reasoning

## Decision

Pass. Haxe.go does not configure semantic-release's npm publisher, so the
unconditional upstream dependency now resolves to a small private local
sentinel. This removes the unused npm CLI and its bundled vulnerability surface
without weakening the complete development-dependency audit. Every publisher
lifecycle hook throws if configuration drift tries to activate the sentinel.

Oracle review was not needed. Local tracing left one defensible design—remove
the unused subtree behind a fail-closed boundary—and exact package-resolution,
release-policy, install-integrity, and vulnerability-scan contracts decide its
correctness.

## Findings And Disposition

The first pass found two required issues:

1. The isolated vulnerability audit copied only `package.json` and
   `package-lock.json`. npm could therefore report a clean audit while the new
   local package link was dangling. Commit `68020db7` now safely copies every
   root-declared in-repository `file:` dependency and requires `npm ls --all`
   before scanning. A regression reproduced the malformed install before the
   fix.
2. The initial decision trail called npm 11.16.0 the CI npm version. CI selects
   Node 24 and currently uses its bundled npm; npm 11.16.0 is the
   repository-declared lock-generator version used for the deterministic local
   refresh check. The trail now states that precisely.

The reviewer also suggested importing the sentinel through its installed
package name. That hardening was accepted, so the hook test proves the
lock-selected package resolves and fails closed rather than testing only its
source path.

The final pass reproduced a clean isolated install under npm 11.16.0, confirmed
the sentinel resolved without a real npm publisher or npm CLI subtree, and
reported no remaining required findings.

## Verification

- `npx --yes npm@11.16.0 install --package-lock-only ...`: no lockfile change
- `npx --yes npm@11.16.0 ci ...`: 260 packages installed
- `npx --yes npm@11.16.0 audit --include=dev --audit-level=high`: zero findings
- `GOTOOLCHAIN=go1.25.12 npm run security:deps`: npm and reachable Go scans clean
- `npm run security:supply-chain`: pass
- `npm run test:release-contracts`: pass
- `npm run test:changed`: pass
- `npm test`: 301/301 snapshots pass

## Follow-up

`haxe_go-vfp.4.15` tracks whether every CI npm invocation should execute the
exact `packageManager` version instead of the npm bundled with the selected
Node 24 image. That policy improvement is separate from this incident: the
committed lock installs structurally and audits cleanly in current CI.
