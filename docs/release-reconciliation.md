# Release Retry and Reconciliation

Release reconciliation is the safe answer to an interrupted publication. It is
not a second workflow, a second version decision, or an application deployment.
It is the final, rerunnable part of the same release command that owns the exact
CI-tested source commit.

In simple terms: the version tag is the sealed label on a package. If a network
failure happens after that label exists but before every file reaches GitHub, a
rerun finishes filling the same box. It does not print a new label, move the old
label to different code, or replace a file whose bytes disagree.

## What owns each decision

- `semantic-release` examines Conventional Commits and may create one new
  `vMAJOR.MINOR.PATCH` tag at the tested commit.
- `scripts/release/reconcile-github-release.mjs` does not analyze commits or
  select a new version. Given an existing tag, tested source SHA, and locally
  verified asset manifest, it reconciles or verifies only that hosted release.
- GitHub's API is authoritative for the remote tag target, Release visibility,
  draft/prerelease/immutable state, and hosted asset name, upload state, size,
  and SHA-256 digest.
- The deterministic artifact builder is authoritative for the expected local
  files. `haxe_go-vfp.4.8` owns production of the final ZIP, checksum, manifest,
  and provenance/attestation plus wiring this reconciler into the release job.

That last boundary is intentional. This state machine must not invent empty or
placeholder assets while the complete artifact contract is still being built.
The future artifact step hands it a schema-versioned JSON control document with
the tag, source SHA, and each asset's relative path, byte count, and
`sha256:<digest>`. This hand-off is not the hosted content manifest and can list
that separate manifest as one of its assets without a self-hash. The reconciler
verifies the local bytes again before any GitHub mutation.

## State contract

| GitHub state | Reconciliation result |
| --- | --- |
| Remote tag missing or points elsewhere | Fail. The command never creates, moves, or deletes a Git tag. |
| Tag exists, Release absent | Create one draft with `gh release create --verify-tag`, then query the API and recheck the tag. A lost create response is resolved by querying GitHub again. |
| Empty or partial matching draft | Preserve matching bytes and upload only missing assets. |
| Draft has a duplicate, unexpected, incomplete, wrong-size, or wrong-digest asset | Fail before the first mutation. Never delete or replace the conflicting bytes. |
| Complete matching draft | Publish it, then re-query until GitHub reports the exact complete Release immutable. |
| Complete immutable published Release | Read-only success. This makes an exact rerun idempotent. |
| Published but mutable, incomplete, prerelease, or conflicting Release | Fail as a release incident. Publish a corrective version; do not rewrite history. |

The read-only `verify` mode applies the same rules but cannot create a draft,
upload an asset, or publish. Reconciliation uploads only missing assets; it has
no delete-or-replace operation. `npm run release:status` uses the locked `semver`
library for reachable-tag selection and reports GitHub API visibility and asset
count without pretending that the asset-less `v0.53.1` historical release
satisfies the future `haxe_go-vfp.4.8` artifact set.

Tag identity is checked before draft creation, before every asset upload,
immediately before publication, and once more after immutable verification. If
another actor publishes the draft between checks, reconciliation stops mutating
and accepts only a complete, digest-matching immutable Release. Repository tag
protection remains the host-side defense against a tag changing in the tiny
interval between two API operations.

## Why this is one path

A GitHub Actions environment turns a protected job into a GitHub Deployment in
the product UI. That model is useful for shipping an application to staging or
production, but this repository publishes an immutable library package. A
separate repair workflow would duplicate permissions, input validation, and
source-identity logic while making operators choose between two release paths.

The intended same-job flow is:

1. Check out the exact SHA whose required gates passed.
2. Run `semantic-release` once for version selection and tag creation.
3. If that exact SHA now has one canonical stable tag—whether created now or by
   an interrupted prior attempt—rebuild and verify the approved assets.
4. Invoke the same reconciliation command for that tag and manifest.
5. Recheck the tag before publication and after the final hosted verification.
6. Treat immutable, digest-matching GitHub API state as the only success.

There is therefore no `.github/workflows/release-repair.yml`, no
`deployments: write` permission, and no release environment approval object.
Human approval can still exist at the manual `publish_release` input and
repository/ruleset level without creating a second publication engine.

## Sibling compiler comparison

The comparison below uses committed sibling checkouts, not their working-tree
changes.

| Compiler | Current pattern | What haxe.go learns |
| --- | --- | --- |
| `haxe.elixir.codex` (`a842290c`) | `complete-release.js` runs inside the same rerunnable CI Release job. It verifies an existing exact tag, preserves matching draft assets, fills missing assets, and makes a completed immutable rerun read-only. Its former separate repair environment was retired. | Adopt the topology and the fail-closed asset plan. This is the primary precedent. |
| `haxe.ruby` (`bb2c9263`) | A well-tested shared hosting state machine is also exposed through a separate `release-repair.yml`; draft repair may delete and replace a conflicting expected asset. | Reuse the explicit state fixtures and authenticated draft lookup, but reject both the second workflow and replacement of conflicting bytes. |
| `haxe.rust` (`85067736`) | Recovery is a separate repair workflow with its own operator path. | Keep its identity and immutability lessons, but not its topology. |

This is not a claim that one workflow design fits every repository. It is the
smallest design that matches haxe.go's threat model: a library release can be
partially published, but it has no independent deployment target to repair.

## Commands and hand-off format

The state fixtures are local and do not contact GitHub:

```bash
npm run test:release-reconciliation
npm run test:release-contracts
```

The production command is intentionally explicit:

```bash
npm run release:reconcile -- \
  --repository OWNER/REPOSITORY \
  --assets /absolute/or/workspace/relative/release-assets.json \
  --notes-file /path/to/release-notes.md
```

Add `--verify-only` for a read-only audit. The asset manifest has this shape;
all paths are relative to the manifest and may not escape its directory:

```json
{
  "schemaVersion": 1,
  "tag": "v0.54.0",
  "sourceSha": "0123456789abcdef0123456789abcdef01234567",
  "assets": [
    {
      "name": "reflaxe.go-0.54.0.zip",
      "path": "reflaxe.go-0.54.0.zip",
      "size": 12345,
      "digest": "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
    }
  ]
}
```

`GH_TOKEN` or `GITHUB_TOKEN` supplies job-scoped GitHub API authority. The
reconciliation command itself never requests or uses GitHub Deployments.
