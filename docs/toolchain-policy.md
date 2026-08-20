# Supported Toolchain Policy

Effective date: 2026-08-19

Machine-readable source: [toolchain-policy.json](toolchain-policy.json)

The generated [compatibility support manifest](compatibility-support-manifest.json)
embeds this policy verbatim and combines it with the admitted platform and
operation/member scope. Toolchain support alone does not admit a workload.

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

| Role | Supported | Recommended | CI version | Meaning |
| --- | --- | --- | --- | --- |
| Haxe compiler | `4.3.7` | `4.3.7` | `4.3.7` | Exact current stable Haxe compiler used by the target and semantic baselines. |
| Generated Go language floor | `1.22` | n/a | n/a | Language and module semantics permitted in generated `go.mod`; not a patched build-toolchain claim. |
| Go build and test toolchain | `1.25.13` and `1.26.6` | `1.26.6` | `1.25.13`, `1.26.6` | Exact approved patches from the two current upstream-supported Go release lines. |
| Node repository tooling | latest patch of `24` | latest patch of `24` | `24` | Active LTS line for npm and release tooling. |
| npm lock and CI executor | exact `packageManager` value | exact `packageManager` value | `npm@11.16.0` | One reviewed npm version generates and consumes the lock; `package.json` owns the exact value. |

The upstream authorities checked for this policy are:

- [Go release history and support policy](https://go.dev/doc/devel/release):
  each major release remains supported until two newer major releases exist;
- [Node.js release schedule](https://github.com/nodejs/Release): Node 24 is
  Active LTS on the effective date; and
- [Haxe download list](https://haxe.org/download/list/): Haxe 4.3.7 is the
  current stable release.

Go versions are exact pins, not `.x` wildcards. This is deliberate: a wildcard
can be satisfied by a stale patch already present in a CI runner's tool cache,
so it does not prove that CI used the latest patched toolchain. The
fail-closed dependency audit exposed this on 2026-07-14 when `1.25.x` resolved
to 1.25.11 and `1.26.x` resolved to 1.26.4 even though Go 1.25.12 and 1.26.5
were available. Both older patches were affected by
[GO-2026-5856](https://pkg.go.dev/vuln/GO-2026-5856).

Schema version 2 therefore replaces `go.ci_selectors` with exact
`go.ci_versions` and adds `go.recommended_build_version`. The separate
`supported_build_lines` and `recommended_build_line` fields still express the
long-lived support policy; the exact-version fields identify the toolchains
that produced current CI and release evidence. A later patch is adopted by an
intentional policy change, not silently during an unrelated build.

The 2026-08-19 refresh moved both supported lines to 1.25.13 and 1.26.6.
The previous patches had reachable standard-library findings in `net/url`,
`crypto/tls`, `encoding/asn1`, and `net/http`. The fail-closed dependency audit
rejected them before release evidence could be accepted.

Setting `check-latest: true` on a wildcard was considered and rejected for
release evidence. It would avoid preferring an older matching cache entry, but
it would also let the concrete toolchain change without a repository change.
[The `setup-go` resolution contract](https://github.com/actions/setup-go#usage)
checks the local tool cache first by default, then its version manifest, then
the official Go distribution. An exact pin keeps that fallback path while
allowing only the reviewed patch to satisfy the request.

## How CI Applies It

1. The quality matrix runs Linux on the exact approved Go 1.25 and 1.26
   patches.
2. The recommended lane runs Linux and macOS on the exact approved Go 1.26
   patch.
3. Harness, security, performance, example-artifact, and release jobs use the
   recommended supported line unless a separate matrix is the subject of the
   test.
4. Haxe is installed as exactly 4.3.7. Node repository jobs use Node 24.
5. Every Node job then runs `scripts/ci/setup-pinned-npm.sh`, which activates
   and verifies the exact npm version declared by `package.json` before any npm
   install, audit, test, performance, or release command.
6. Release status checks read the machine policy and supply-chain contract and
   fail when workflow wiring drifts from it.

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

1. Check the official Go, Node, and Haxe sources above. For Go, verify the
   concrete patches in the release history and review current vulnerability
   advisories.
2. Add or update a failing contract for the intended transition.
3. Change `ci_versions`, `recommended_build_version`, and the effective date in
   the machine-readable policy. Change the line fields only when support moves
   to a different release line.
4. Align workflow matrices, release-status wiring, and this explanation.
5. Regenerate the compatibility support manifest and its generated docs.
6. Run the targeted policy contract, release contracts, security gates,
   snapshots, and full CI. Preserve `go version` output in release evidence.
7. Remove an old lane or label it explicitly as compatibility-only with
   `production_supported=false` and `security_supported=false`.

Update npm independently through the exact `packageManager` value in
`package.json`. Regenerate the lock with that version, inspect the operational
dependency audit, and keep every workflow behind the bootstrap contract. Do not
duplicate the npm version in workflow environment variables.

Do not retain an end-of-life production lane merely to preserve a green
historical baseline. Preserve the baseline as fixture evidence and run it with
an explicitly bounded purpose instead.
