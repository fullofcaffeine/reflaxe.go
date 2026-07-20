# SemVer and Compatibility Lifecycle Policy

This policy explains what users can expect from haxe.go `0.x`, how public
features are deprecated, and what must be true before `1.0.0` is possible. The
small machine checklist in `release/policy.json` is the release analyzer's
authority for release lines and stable-major approval. The
[public contract](public-contract.md) remains the authority for which surfaces
are user-facing.

In plain language: an author or agent describes the impact of a change, tests
and review check that description, and the release tool chooses the number. An
agent cannot approve `1.0.0` by changing a commit message.

## What beta means

The compatibility manifest currently describes the bounded release claim as a
pre-1.0 beta. That manifest remains the maturity authority; this document only
explains the SemVer consequence. Beta is a product-readiness statement: named
workloads are useful under the documented toolchain, platform,
operation/member, and trust boundaries. It is not automatically a SemVer prerelease
such as `1.0.0-beta.1`.

Normal releases continue the stable `0.y.z` tag sequence on `master`. A beta
claim does not mean that all public APIs are stable, and a valid SemVer tag does
not by itself prove bounded-production readiness. Compatibility evidence and
the release number answer different questions.

## Change classes during 0.x

Major zero communicates initial development, but haxe.go still uses predictable
change classes:

| Change | Commit classification | `0.x` result |
| --- | --- | --- |
| Compatible bug, security, or performance fix | `fix:` or `perf:` | patch |
| Additive public feature or newly admitted surface | `feat:` | minor |
| Incompatible change to a documented public or compatibility surface | `!` or `BREAKING CHANGE:` | minor, never automatic `1.0.0` |
| Internal refactor with no protected behavior change | `refactor:`, `test:`, `build:`, or `chore:` | no release under current rules |
| Documentation only | `docs:` | no release under current rules |

A patch must not remove a public surface, change a documented default, add a
new deprecation warning, narrow a supported toolchain, or reinterpret metadata.
Breaking `0.x` changes require migration guidance even though their numeric bump
is minor.

## Experimental surfaces

An experimental surface is visible but is outside the admitted compatibility
promise unless the compatibility manifest explicitly promotes it. Adding,
changing, or removing documented experimental behavior therefore does not make
the stable contract false, but it is never a patch-level surprise: the minimum
change release is a minor and release notes must name the affected surface.
Use the `experimental` Conventional Commit scope—for example,
`fix(experimental):`—when a non-breaking commit changes such a surface. The
analyzer raises that otherwise-patch classification to the required minor.

After stable admission, a minor-line experimental exemption requires explicit surface proof
that:

1. the previous release classified the exact surface as experimental;
2. no admitted public surface depends on the changed behavior;
3. the compatibility manifest and executable evidence still agree; and
4. release notes provide replacement or migration guidance where applicable.

The analyzer does not infer that proof from a commit scope. Until a dedicated
machine check can establish it, a claimed stable-minor experimental exception
is rejected rather than granted a minor release. This prevents
`feat(experimental)` from becoming a loophole for silently breaking admitted
APIs.

## Deprecation windows

For an ordinary supported public surface during `0.x`, the minimum sequence is:

1. notice in minor N, with an actionable replacement and migration example;
2. keep the old behavior functional throughout minor N+1; and
3. removal no earlier than minor N+2, using a breaking commit and migration
   notes.

Patch releases do not start or finish that clock. Compatibility selectors and
aliases, including `metal` and `@:goMetal`, use the same floor unless a stricter
recorded retention decision applies.

A commit that first adds a deprecation warning uses the `deprecation` scope;
for example, `fix(deprecation):` has a minor floor. Removing the old surface
still requires `!` or a `BREAKING CHANGE:` footer. The scope records the release
floor; tests and review remain responsible for proving the N/N+1/N+2 timeline.

Experimental surfaces may change in a minor without the full N/N+1/N+2 window,
but the release must still identify the exact experimental surface and must not
affect admitted behavior. Once stable, a supported public surface may be
deprecated in a minor but remains functional for the rest of that major line;
removal waits for the next approved major.

A security or correctness emergency can shorten a window only with a reviewed
exception record, prominent release note, and concrete migration guidance. The
exception must explain why compatibility was less safe than the change.

## Profiles, metadata, and generated output

- Adding an orthogonal policy value, documented metadata, macro, or typed API is
  normally a minor feature.
- Removing a profile value or alias, changing its meaning/default, or changing
  metadata expansion in a consumer-visible way is breaking. During `0.x` it
  follows the deprecation window above; on a stable line it is major-only.
- `metal` is a compatibility convenience preset, not a second semantic
  product. That does not make its accepted selector or aliases internal.
- Exact generated whitespace, private locals, and unadvertised helper names are
  internal. Documented handwritten-Go call shapes, versioned report fields,
  package layout, and compiler/runtime compatibility are public as defined by
  the public contract.

## Release channels

The normal channel is `master` with canonical `v<SemVer>` tags and no
prerelease component. Development manifests remain at the `0.0.0` sentinel and
cannot seed release lineage. Product labels such as beta are not automatically
a SemVer prerelease.

A future `alpha`, `beta`, or release-candidate tag channel requires a separate
reviewed branch/channel design, upgrade path, and tests. Until then the analyzer
rejects prerelease lineage rather than guessing how it should advance.

## Stable 1.x admission

`release/policy.json` lists the objective requirements for stable major 1. The
current state is blocked. Before approval, all of these must be complete:

- the public contract and executable owner/evidence mapping;
- an exact stable scope and support matrix, including explicit exclusions;
- the fail-closed release-readiness gate with no applicable unresolved P0/P1
  blocker;
- a published, checksummed, provenance-bound beta baseline from the tested SHA;
- consumer upgrade and rollback rehearsal against immutable packages;
- supported-toolchain security, race, static, and licensing closure; and
- an independent stable-readiness review against a commit-pinned evidence
  bundle.

Completing those rows does not automatically authorize `1.0.0`. The stable line
also needs explicit human approval recorded with the approving person or group,
date, decision record, and reviewed source commit. When the major transition is
attempted, the analyzer verifies that the commit exists in the release
repository and is an ancestor of the release candidate. Because adding the
approval necessarily creates a later commit, only the target approval field
and Beads interaction metadata may differ after the reviewed source; any other
policy, compiler, runtime, documentation, test, workflow, or packaging change
makes the approval stale. Once that major has shipped, its historical approval
does not block ordinary maintenance releases on the admitted line. The analyzer
also rejects approval while any requirement is pending and rejects stable
lineage that lacks the recorded approval.

Stable admission is a scoped promise, not a claim that every Haxe module, Go
API, platform, or experimental feature is supported.

## Who decides

The change author—human or agent—proposes the Conventional Commit
classification by comparing the change with the public contract and this
policy. Reviewers challenge incorrect classifications. CI verifies the machine
policy, compatibility evidence, and release mechanics.

The release analyzer mechanically maps the reviewed commits to a version. It
does not decide whether a source concept should be public, and it cannot grant
stable-major approval. Only the explicit human approval record in
`release/policy.json` can unlock a stable major after every requirement passes.

## Sibling lessons

This design follows commit-pinned local sibling evidence:

- `haxe.elixir.codex` at `2030abea264dac770915dbeff427acc349ff082e`
  uses a small release-line manifest with pending requirements and an approval
  record. haxe.go adopts that fail-closed checklist pattern.
- `haxe.ruby` at `ded7f02d666612350440d2d31e52dfe48449f9b9`
  keeps readable public/deprecation policy alongside executable package and
  upgrade contracts. haxe.go keeps those human rules explicit here.
- `haxe.rust` at `85067736d0b929dfc67d6684d59b7e2bd3bae6ea`
  couples stable-candidate status and deprecation details to its much larger
  public manifest. haxe.go does not need that declaration-scale mechanism to
  enforce release-line admission.

## Validation

Run:

```bash
python3 test/test_semver_lifecycle_policy.py
npm run test:release-version-policy
npm run release:policy
npm run test:release-contracts
```
