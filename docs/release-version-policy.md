# Release Version and Source-Identity Policy

Git tags are the only version authority for haxe.go. Tracked manifests identify
a development checkout; they do not pretend to be the latest published
release. The release workflow selects its next version from the canonical Git
tag lineage and can create a new tag only at the exact CI-tested SHA.

## What, why, and how

The release identity flow is:

1. Normal pushes, pull requests, and scheduled runs execute validation but
   cannot publish. Publication starts only from an explicit manual release request
   on `master` with the `publish_release` input enabled.
2. `package.json`, `package-lock.json`, and
   `haxelib.json` carry the 0.0.0 development sentinel.
3. semantic-release reads the newest canonical `v<SemVer>` tag
   reachable from the tested commit.
4. `scripts/release/analyze-commits.mjs` maps Conventional Commits
   to the next version under the reviewed major-zero policy.
5. `scripts/release/run-same-sha-release.sh` requires the current
   commit to equal `RELEASE_TESTED_SHA`, requires a clean tracked
   checkout, and runs semantic-release.
6. The wrapper verifies that semantic-release changed neither `HEAD`
   nor tracked files. If a tag was created, that tag must resolve to the tested
   commit; CI also verifies the tag at `origin`.
7. When versioned package metadata is needed,
   `scripts/release/stage-release-metadata.py` writes it into a new
   output directory and binds the version, tag, source commit, and metadata
   hashes in `release-identity.json`. It never edits the source
   manifests.

The sentinel prevents the old ambiguity where source manifests said one
version while the Git history had advanced much further. Temporary staging
keeps released metadata truthful without introducing an untested version or
changelog commit after the quality gates. The sentinel itself is never a valid
release tag or staged release version.

## Conventional Commit mapping

The custom analyzer delegates parsing to the installed official
`@semantic-release/commit-analyzer`. It supports both a
`type!:` header and a `BREAKING CHANGE:` footer, then
applies these haxe.go rules:

| Commit set | Current lineage | Result |
| --- | --- | --- |
| `fix:` or `perf:` | supported major | patch |
| `feat:` | supported major | minor |
| breaking change | `0.x` without major 1 approval | minor |
| breaking change | `0.x` with major 1 approval | `1.0.0` |
| breaking change | stable N with N+1 approved | `(N+1).0.0` |
| docs, tests, or chores only | supported major | no release |

A breaking 0.x change advances the minor version. Major zero already
communicates that the public API is in initial development; a breaking commit
must not promote the project to `1.0.0` accidentally.

## Stable-major approval

The analyzer option `approvedStableMajors` is a reviewed, contiguous
list. It is empty today. Graduation to major 1 requires changing it to
`[1]` in a tested review. A later major 2 requires `[1, 2]`.
Skipped, duplicated, unknown, prerelease, and noncanonical lineages fail
closed.

## Mutation-free publication

The semantic-release configuration contains only the policy analyzer and the
GitHub publisher. Release automation never rewrites or commits CHANGELOG.md,
`package.json`, or `haxelib.json`. Hosted release notes are
generated from commits between the exact previous and next Git tags.

The final workflow job exists only for a manual run whose
`publish_release` input is true on `master`. It waits for
every release-blocking job, checks out the workflow run's explicit commit,
receives that same value as `RELEASE_TESTED_SHA`, and has only
`contents: write`. Commenting, labeling, and issue mutation are
disabled because they are unrelated to publishing the tested source. A routine
master push cannot publish accumulated Conventional Commits.

## Staged metadata

Create versioned metadata for a known tag and source commit with:

```bash
npm run release:stage-metadata -- \
  --version 0.54.0 \
  --source-sha aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa \
  --output-dir /tmp/haxe-go-release-metadata
```

The output directory must not already exist. Versions must be canonical stable
SemVer and source identities must be full lowercase 40-character commit
hashes. The source checkout remains byte-identical.

## Deterministic Haxelib artifacts

Build release-ready package evidence from an exact tested commit with:

```bash
npm run release:build-haxelib -- \
  --version 0.54.0 \
  --tag v0.54.0 \
  --source-sha aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa \
  --output-dir /tmp/haxe-go-haxelib-artifact
```

The builder makes two independent `git archive` exports of the supplied
commit. In each fresh tree it stages versioned metadata, runs the canonical
Reflaxe package adapter, and independently verifies package paths, hashes,
fixed ZIP metadata, generated `.cross.hx` ownership, and the absence of local
paths or development debris. The two ZIPs and their embedded manifests must be
byte-identical before any output directory is created.

The output directory contains exactly:

- `reflaxe.go-<version>.zip`
- `reflaxe.go-<version>.zip.sha256`
- `reflaxe.go-<version>.manifest.json`

The JSON manifest binds the artifact digest and complete embedded
source-to-package map to the version, proposed tag, and source commit. Artifact
construction happens before publication, so the proposed tag may not exist
yet. If an existing tag has that name, the builder fails unless it already
resolves to the supplied source SHA. The later publication transaction remains
responsible for creating or verifying that same tag without moving it.
Successful construction is not publication approval by itself: isolated
package execution, licensing/notices, hosted provenance, and release-state
checks remain separate fail-closed gates.

Verify a candidate ZIP independently with:

```bash
npm run release:verify-haxelib -- \
  --zip /tmp/haxe-go-haxelib-artifact/reflaxe.go-0.54.0.zip \
  --version 0.54.0 \
  --tag v0.54.0 \
  --source-sha aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
```

## Executable evidence

Run the focused version and source-identity contracts with:

```bash
npm run test:release-version-policy
python3 test/test_release_identity_contract.py
python3 test/test_same_sha_release_wrapper.py
npm run release:policy
```

The JavaScript test drives both the custom analyzer and the installed
semantic-release engine against temporary Git history. The Python tests prove
metadata staging, checkout immutability, no-release behavior, correct-tag
behavior, and rejection of mismatched commits or mutated tracked files.

`npm run release:status` composes this policy with the supported
toolchain, supply-chain, tag visibility, and release asset-path checks. Artifact
construction, checksums, provenance publication, idempotence, and repair are
separate release-protocol gates and must preserve this same source identity.
