# Release Visibility Checks

## Quick check

```bash
npm run release:status
```

This runs `scripts/release/check-release-state.sh` and verifies:

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

## Related files

- `.github/workflows/ci-harness.yml`
- `.github/workflows/examples-artifacts.yml`
- `docs/toolchain-policy.json`
- `.releaserc.json`
- `scripts/release/verify-release-policy.py`
- `scripts/release/run-same-sha-release.sh`
- `scripts/release/check-release-state.sh`
