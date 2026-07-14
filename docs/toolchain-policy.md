# Supported Toolchain Policy

Effective date: 2026-07-14

Machine-readable source: [toolchain-policy.json](toolchain-policy.json)

## What It Is

This policy separates three things that are easy to conflate:

1. the Haxe compiler version used to compile Haxe source;
2. the Go language floor written into generated `go.mod` files; and
3. the patched Go and Node toolchains that Haxe.Go supports for builds, tests,
   security evidence, and releases.

The machine-readable JSON is authoritative. This page explains why each value
exists and how CI applies it.

## Why It Exists

A generated module can declare an older Go language version while still being
built with a current, security-supported Go toolchain. Treating that language
floor as a production toolchain pin would leave users on unsupported standard
libraries. Similarly, the Node runtime used by GitHub Actions is not the same
as the Node version used by repository tooling.

Release evidence is valid only when it records the exact resolved patch
versions from supported lines. A green compatibility fixture on an older line
does not establish production or security support.

## Supported Versions

| Role | Supported | Recommended | CI selector | Meaning |
| --- | --- | --- | --- | --- |
| Haxe compiler | `4.3.7` | `4.3.7` | `4.3.7` | Exact current stable Haxe compiler used by the target and semantic baselines. |
| Generated Go language floor | `1.22` | n/a | n/a | Language and module semantics permitted in generated `go.mod`; not a patched build-toolchain claim. |
| Go build and test toolchain | latest patch of `1.25` and `1.26` | latest patch of `1.26` | `1.25.x`, `1.26.x` | The two current upstream-supported Go release lines. |
| Node repository tooling | latest patch of `24` | latest patch of `24` | `24` | Active LTS line for npm and release tooling. |

The upstream authorities checked for this policy are:

- [Go release history and support policy](https://go.dev/doc/devel/release):
  each major release remains supported until two newer major releases exist;
- [Node.js release schedule](https://github.com/nodejs/Release): Node 24 is
  Active LTS on the effective date; and
- [Haxe download list](https://haxe.org/download/list/): Haxe 4.3.7 is the
  current stable release.

Selectors ending in `.x` intentionally resolve the latest patch available in
that supported Go line. Release evidence must record the concrete version
reported by `go version`; the selector alone is not sufficient provenance.

## How CI Applies It

1. The quality matrix runs Linux on the latest Go 1.25 and 1.26 patches.
2. The recommended lane runs Linux and macOS on the latest Go 1.26 patch.
3. Harness, security, performance, example-artifact, and release jobs use the
   recommended supported line unless a separate matrix is the subject of the
   test.
4. Haxe is installed as exactly 4.3.7. Node repository jobs use Node 24.
5. Release status checks read the machine policy and fail when workflow wiring
   drifts from it.

The compiler continues to emit `go 1.22`. Raising CI versions does not raise
that directive. A language-floor change requires a separate red-first compiler
contract showing that emitted code or a shipped dependency needs newer Go
language semantics.

## Compatibility-Only Evidence

The committed `goextern` fixtures remain pinned to Go 1.23 so generator drift
can be reviewed against a stable standard-library snapshot. That is a
non-production compatibility fixture and does not establish security support.
Supported CI toolchains run the current-toolchain smoke path instead.

Snapshots containing `go 1.22` also prove generated shape and the declared
language floor. They are not evidence that a Go 1.22 toolchain is safe or
supported for production.

## Updating The Policy

1. Check the official Go, Node, and Haxe sources above.
2. Change the machine-readable policy first.
3. Add or update a failing contract for the intended transition.
4. Align workflow matrices, release-status wiring, and this explanation.
5. Run the targeted policy contract, release contracts, snapshots, and full CI.
6. Remove an old lane or label it explicitly as compatibility-only with
   `production_supported=false` and `security_supported=false`.

Do not retain an end-of-life production lane merely to preserve a green
historical baseline. Preserve the baseline as fixture evidence and run it with
an explicitly bounded purpose instead.
