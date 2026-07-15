# Security Dependency Audit

This page explains what the dependency audit proves, why Go standard-library
findings can originate in the selected toolchain, and how CI handles every exit
path.

## What It Checks

The audit has two independent parts:

- `npm audit` checks the complete locked Node tree, including dev-scoped CI and
  release tooling, at high severity or above.
- `govulncheck` uses call-graph analysis to check `runtime/hxrt` for
  reachable entries in the Go vulnerability database.

The scanner version is pinned for reproducibility and runs on the supported Go
toolchain from [the toolchain policy](toolchain-policy.md). CI uses
`govulncheck` text format with full traces because the official command exit
code is nonzero when reachable vulnerabilities are found. JSON, SARIF, and VEX
formats are useful interchange artifacts but are not used as this gate because
those formats can exit successfully even when they contain findings.

All repository Node dependencies are currently declared under
`devDependencies`, but that npm installation category is not the security
boundary. Tools that analyze commits, choose a release version, render release
notes, or publish a GitHub release execute with repository or release
authority. The isolated install therefore uses `--include=dev` explicitly so
an inherited `NODE_ENV=production` cannot silently remove the operational
tooling from the audit.

The initial full-tree finding inventory, configured-plugin reachability map,
and all 19 remediations are recorded in
[the operational npm dependency audit](reviews/npm-operational-dependency-audit-vfp-4.12.md).

## Why SSL And Network Findings Can Appear

`runtime/hxrt` implements Haxe standard-library SSL and network support.
Those helpers intentionally call Go packages including
`crypto/tls`, `crypto/x509`, `net`, `encoding/pem`, and
`encoding/asn1`.

If a selected Go patch contains a known vulnerable implementation in one of
those packages and an `hxrt` call path reaches the affected symbol,
`govulncheck` reports a Go standard-library vulnerability. That is a
toolchain vulnerability with project release impact: it is not automatically a
logic defect in the Haxe.Go helper, but the project must not ship a green
release gate while the vulnerable path remains reachable.

This is not a dependency install failure. A project defect, such as discarding
an error or mishandling TLS state, is a separate root cause. Both kinds of
problem can block a release.

## How CI Fails Closed

The CI gate fails closed: only a completed scan with zero reachable findings
passes.

The audit has three terminal states:

| Result | CI outcome | Required response |
| --- | --- | --- |
| No reachable finding | pass | Retain the report as evidence for the exact Go patch. |
| One or more reachable findings | fail | Upgrade to a fixed supported Go patch, remove the reachable path without breaking Haxe semantics, or keep the release blocked. |
| Scanner install, load, database, or execution error | fail | Repair the audit; absence of a completed scan is not evidence of safety. |

`SKIP_GOVULNCHECK=1` is rejected in CI. A local-only skip or explicitly
enabled local install soft-fail can help offline development, but neither path
is available to a release job.

The report is sanitized before printing so raw workflow commands and
source-position problem matchers do not create misleading annotations. The
upstream trace, gate result, configured scanner version, scanner exit code, and
Go version are retained under `.cache/security/dependency-audit`. Both
dependency-audit workflows upload `.cache/security` with `if: always()`, so
reports are uploaded even when findings or tool errors fail the job.

## Interpreting A Failure

Use the trace and metadata in this order:

1. Confirm the exact Go patch and pinned `govulncheck` version.
2. Identify whether the affected module is `stdlib` or a third-party module.
3. For `stdlib`, check whether a newer patch in the supported Go line fixes
   the advisory.
4. If no fixed supported patch exists, determine whether the affected
   capability can be removed from the packaged runtime and proven unreachable.
5. For a third-party module, update or remove the dependency and rerun the same
   call-graph scan.
6. If the scanner itself failed, repair the tool or database access and rerun;
   never classify the missing result as clean.

Do not downgrade a reachable finding to a warning because the path is expected.
Expected reachability explains the trace; it does not make vulnerable code safe.
