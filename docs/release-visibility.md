# Release Visibility Checks

## Quick check

```bash
npm run release:status
```

This runs `scripts/release/check-release-state.sh` and verifies:

- a reachable semver baseline tag exists from current `HEAD`
- `package.json` and `haxelib.json` versions match
- semantic-release `tagFormat` is `v${version}`
- CI workflow release wiring exists (`semantic-release` in `ci-harness`)
- examples tag-release asset paths are normalized to deterministic staging paths
- optional remote GitHub release visibility for the latest tag (when `gh` can access the repo)

## Why this exists

Release automation previously failed because workflow assumptions drifted in two places:

1. tag/release baseline expectations (initial-release semantics)
2. artifact path assumptions between `upload-artifact` and `download-artifact`

The status script makes those assumptions explicit and machine-checkable before release jobs run.

## Related files

- `.github/workflows/ci-harness.yml`
- `.github/workflows/examples-artifacts.yml`
- `.releaserc.json`
- `scripts/release/check-release-state.sh`
