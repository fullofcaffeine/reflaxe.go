# Release Visibility Checks

## Quick check

```bash
npm run release:status
```

This runs `scripts/release/check-release-state.sh` and verifies:

- the governed compatibility source, machine manifest, human matrix, and
  release-status paragraph are current and agree;
- supported Haxe, Go, and Node workflow wiring matches
  [`toolchain-policy.json`](toolchain-policy.json)
- a reachable semver baseline tag exists from current `HEAD`
- source manifests and the npm lock use the `0.0.0` development
  sentinel
- semantic-release `tagFormat` is `v${version}`
- the release configuration has no tracked-checkout mutator
- the final workflow checks out and publishes the exact tested SHA
- publication requires an explicit manual `publish_release` request
  on `master`; ordinary pushes only validate
- CI workflow release wiring exists (`semantic-release` in `ci-harness`)
- examples tag-release asset paths are normalized to deterministic staging paths
- optional remote GitHub release visibility for the latest tag (when `gh` can access the repo)

## Why this exists

Release automation previously failed because workflow assumptions drifted in two places:

1. tag/release baseline expectations (initial-release semantics)
2. artifact path assumptions between `upload-artifact` and `download-artifact`

The status script makes those assumptions explicit and machine-checkable before
release jobs run. The detailed version and source-identity contract is in
[Release Version and Source-Identity Policy](release-version-policy.md).

Compatibility truth has its own fail-closed chain:

1. Humans edit `docs/compatibility-support-source.json`.
2. `npm run compatibility:generate` derives
   `docs/compatibility-support-manifest.json`,
   `docs/compatibility-support-matrix.md`, and
   `docs/compatibility-release-status.md` from that source plus the toolchain
   policy and portable stdlib inventory.
3. `npm run compatibility:verify` rejects stale outputs, unknown evidence
   states, implicit operation/member entries, or missing evidence references.
4. `npm run release:status` consumes the generated manifest and prints the
   bounded release claim, admitted preset, and admitted platform.

The [generated release status](compatibility-release-status.md) is the
compatibility paragraph for release notes. The
[generated support matrix](compatibility-support-matrix.md) explains the human
scope; the [JSON manifest](compatibility-support-manifest.json) is authoritative.

## Related files

- `.github/workflows/ci-harness.yml`
- `.github/workflows/examples-artifacts.yml`
- `docs/compatibility-support-source.json`
- `docs/compatibility-support-manifest.json`
- `docs/compatibility-support-matrix.md`
- `docs/compatibility-release-status.md`
- `docs/toolchain-policy.json`
- `.releaserc.json`
- `scripts/release/verify-release-policy.py`
- `scripts/release/run-same-sha-release.sh`
- `scripts/release/check-release-state.sh`
- `scripts/compatibility/generate_support_manifest.py`
