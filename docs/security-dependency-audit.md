# Security Dependency Audit

This page explains the dependency audit job in CI and why it can show Go standard-library vulnerability warnings even when the job succeeds.

## What It Checks

The audit has two parts:

- `npm audit` checks production Node dependencies at high severity or above.
- `govulncheck` checks the Go runtime package under `runtime/hxrt` for reachable
  Go vulnerability database entries.

`govulncheck` means "Go vulnerability check". It follows actual call paths in Go
code and reports vulnerabilities that are reachable from the package being
scanned.

## Why SSL And Network Warnings Appear

`runtime/hxrt` implements Haxe standard-library features such as SSL and network support. Those helpers intentionally call Go packages like `crypto/tls`,
`crypto/x509`, `net`, `encoding/pem`, and `encoding/asn1`.

When the CI Go toolchain has known standard-library vulnerabilities in those
packages, `govulncheck` reports the call paths. That is useful classified audit
evidence: it tells us which runtime helpers could reach the vulnerable standard
library code.

It is not a dependency install failure, and it is not the same thing as a broken
npm dependency. It means the runtime uses Go standard-library functionality that
must be reviewed against the active Go toolchain and the affected generated-code
surface.

## How CI Reports It

The audit script disables raw `govulncheck` GitHub error annotations and emits a
repo-owned annotation instead:

- `[deps][govulncheck-stdlib-reachability]`: reachable Go standard-library
  vulnerability reports were found in `runtime/hxrt`.

The full upstream `govulncheck` report is still printed in the job log. The
single classified annotation exists so CI readers can tell that the finding is
tracked security evidence, not an unexplained failing dependency step.

## What To Do With A Finding

Review the reported call paths and decide which response fits:

1. Upgrade the CI Go version when the vulnerability is fixed by a newer supported
   Go release.
2. Patch or narrow the `hxrt` helper if the generated-code surface can avoid the
   vulnerable path without breaking Haxe semantics.
3. Keep the finding as classified audit evidence when the call path is expected
   and the project accepts the current toolchain risk until the next Go update.

Do not delete the warning just to make CI quieter. Security warnings are allowed
to be non-blocking only when they remain visible and classified.
