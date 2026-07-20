#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT_DIR"

log() {
  printf '[release-status] %s\n' "$*"
}

warn() {
  printf '[release-status] warning: %s\n' "$*" >&2
}

fail() {
  printf '[release-status] error: %s\n' "$*" >&2
  exit 1
}

require_file() {
  local path="$1"
  if [[ ! -f "$path" ]]; then
    fail "required file not found: $path"
  fi
}

require_contains() {
  local path="$1"
  local pattern="$2"
  local label="$3"
  if ! grep -Fq -- "$pattern" "$path"; then
    fail "missing ${label} in ${path}: ${pattern}"
  fi
}

require_command() {
  local command="$1"
  if ! command -v "$command" >/dev/null 2>&1; then
    fail "required command not found: $command"
  fi
}

parse_repo_slug() {
  local remote_url="$1"
  local without_suffix
  without_suffix="${remote_url%.git}"
  if [[ "$without_suffix" =~ ^git@github\.com:(.+)/(.+)$ ]]; then
    printf '%s/%s\n' "${BASH_REMATCH[1]}" "${BASH_REMATCH[2]}"
    return 0
  fi
  if [[ "$without_suffix" =~ ^https://github\.com/(.+)/(.+)$ ]]; then
    printf '%s/%s\n' "${BASH_REMATCH[1]}" "${BASH_REMATCH[2]}"
    return 0
  fi
  return 1
}

require_command git
require_command node
require_command python3

require_file "docs/compatibility-support-source.json"
require_file "docs/compatibility-support-manifest.json"
require_file "docs/compatibility-support-matrix.md"
require_file "docs/compatibility-release-status.md"
python3 scripts/compatibility/generate_support_manifest.py --check
COMPATIBILITY_RELEASE_CLAIM="$(node -p "require('./docs/compatibility-support-manifest.json').release_claim.statement")"
COMPATIBILITY_PRESET="$(node -p "require('./docs/compatibility-support-manifest.json').release_claim.admitted_preset")"
COMPATIBILITY_PLATFORM="$(node -p "require('./docs/compatibility-support-manifest.json').release_claim.admitted_platform")"
log "compatibility support manifest: preset=${COMPATIBILITY_PRESET}, platform=${COMPATIBILITY_PLATFORM}"
log "compatibility release claim: ${COMPATIBILITY_RELEASE_CLAIM}"

python3 scripts/security/verify-supply-chain.py
log "supply-chain provenance: OK"

python3 scripts/release/verify-release-policy.py
log "release identity policy: OK"

for release_file in \
  scripts/release/build-haxelib-artifact.py \
  scripts/release/verify-haxelib-artifact.py \
  scripts/release/verify-release-assets.py \
  scripts/release/reconcile-github-release.mjs; do
  require_file "$release_file"
done
require_contains "scripts/release/run-same-sha-release.sh" "release-assets.json" "verified asset handoff"
require_contains "scripts/release/run-same-sha-release.sh" "verify-release-assets.py" "independent asset verification"
require_contains "scripts/release/run-same-sha-release.sh" "reconcile-github-release.mjs" "same-workflow hosted reconciliation"
log "complete release asset pipeline: ZIP, checksum, manifest, provenance, hosted reconciliation"

python3 scripts/release/verify-license-policy.py --mode audit
log "license inventory policy: audited (approval is checked only by publication entrypoints)"

require_file "docs/toolchain-policy.json"
POLICY_HAXE_SELECTOR="$(node -p "require('./docs/toolchain-policy.json').haxe.ci_selector")"
POLICY_GO_FLOOR="$(node -p "require('./docs/toolchain-policy.json').go.generated_language_floor")"
POLICY_GO_VERSIONS="$(node -p "require('./docs/toolchain-policy.json').go.ci_versions.join(' ')")"
POLICY_GO_MATRIX="$(node -p "'[' + require('./docs/toolchain-policy.json').go.ci_versions.map((version) => JSON.stringify(version)).join(', ') + ']'")"
POLICY_GO_RECOMMENDED_VERSION="$(node -p "require('./docs/toolchain-policy.json').go.recommended_build_version")"
POLICY_NODE_SELECTOR="$(node -p "require('./docs/toolchain-policy.json').node.ci_selector")"

require_file ".github/workflows/ci-quality.yml"
require_contains ".github/workflows/ci-quality.yml" "HAXE_VERSION: \"${POLICY_HAXE_SELECTOR}\"" "supported Haxe selector"
require_contains ".github/workflows/ci-quality.yml" "NODE_VERSION: \"${POLICY_NODE_SELECTOR}\"" "supported Node selector"
for version in $POLICY_GO_VERSIONS; do
  require_contains ".github/workflows/ci-quality.yml" "go: \"${version}\"" "supported Go matrix version ${version}"
done

for workflow in .github/workflows/ci-harness.yml .github/workflows/security-static-analysis.yml; do
  require_file "$workflow"
  require_contains "$workflow" "NODE_VERSION: \"${POLICY_NODE_SELECTOR}\"" "recommended Node selector"
  require_contains "$workflow" "GO_VERSION: \"${POLICY_GO_RECOMMENDED_VERSION}\"" "recommended Go version"
  require_contains "$workflow" "go: ${POLICY_GO_MATRIX}" "supported Go matrix versions"
done

require_file ".github/workflows/examples-artifacts.yml"
require_contains ".github/workflows/examples-artifacts.yml" "GO_VERSION: \"${POLICY_GO_RECOMMENDED_VERSION}\"" "recommended example-build Go version"
require_contains "src/reflaxe/go/GoReflaxeCompiler.hx" "\"go ${POLICY_GO_FLOOR}\"" "generated Go language floor"
require_contains "src/reflaxe/go/GoOutputIterator.hx" "\"go ${POLICY_GO_FLOOR}\"" "iterator Go language floor"
log "toolchain policy wiring: Haxe ${POLICY_HAXE_SELECTOR}, Go ${POLICY_GO_VERSIONS} (recommended ${POLICY_GO_RECOMMENDED_VERSION}), Node ${POLICY_NODE_SELECTOR}, generated floor ${POLICY_GO_FLOOR}"

SEMVERS="$(git tag --merged HEAD | node scripts/release/reconcile-github-release.mjs --list-stable-tags)"
if [[ -z "$SEMVERS" ]]; then
  fail "no canonical stable SemVer tag is reachable from current HEAD; semantic-release may treat this as an initial release"
fi

LATEST_TAG="$(printf '%s\n' "$SEMVERS" | sed -n '1p')"
TAG_COUNT="$(printf '%s\n' "$SEMVERS" | sed '/^$/d' | wc -l | tr -d ' ')"
log "reachable semver tags: $TAG_COUNT"
log "latest reachable semver tag: $LATEST_TAG"

PACKAGE_VERSION="$(node -p "require('./package.json').version")"
HAXELIB_VERSION="$(node -p "require('./haxelib.json').version")"
HAXELIB_URL="$(node -p "require('./haxelib.json').url || ''")"
DEVELOPMENT_VERSION="0.0.0"

if [[ "$PACKAGE_VERSION" != "$HAXELIB_VERSION" ]]; then
  fail "package.json (${PACKAGE_VERSION}) and haxelib.json (${HAXELIB_VERSION}) versions differ"
fi
if [[ "$PACKAGE_VERSION" != "$DEVELOPMENT_VERSION" ]]; then
  fail "source manifests must use the ${DEVELOPMENT_VERSION} development sentinel (package=${PACKAGE_VERSION}, haxelib=${HAXELIB_VERSION})"
fi
log "package/haxelib development sentinel parity: ${DEVELOPMENT_VERSION}; Git tags own released versions"

RELEASE_TAG_FORMAT="$(node -p "(require('./.releaserc.json').tagFormat || '')")"
if [[ "$RELEASE_TAG_FORMAT" != "v\${version}" ]]; then
  fail "unexpected semantic-release tagFormat in .releaserc.json: ${RELEASE_TAG_FORMAT}"
fi
log "semantic-release tagFormat: ${RELEASE_TAG_FORMAT}"

require_file ".github/workflows/examples-artifacts.yml"
require_contains ".github/workflows/examples-artifacts.yml" "dist-upload/release-files/checksums.txt" "release checksums path"
require_contains ".github/workflows/examples-artifacts.yml" "dist-upload/release-files/manifest.json" "release manifest path"
require_contains ".github/workflows/examples-artifacts.yml" "dist-upload/release-files/examples-\${{ github.ref_name }}.tar.gz" "release archive path"
require_contains ".github/workflows/examples-artifacts.yml" "dist-upload/release-files/examples-\${{ github.ref_name }}.tar.gz.sha256" "release checksum path"
log "examples release asset path normalization wiring: OK"

require_file ".github/workflows/ci-harness.yml"
require_contains ".github/workflows/ci-harness.yml" "semantic-release:" "semantic-release job declaration"
require_contains ".github/workflows/ci-harness.yml" "npm run release" "semantic-release publish command"
require_contains ".github/workflows/ci-harness.yml" "- go-tooling" "release-blocking Go tooling gate"
log "ci harness same-workflow release wiring: OK"

if ORIGIN_URL="$(git remote get-url origin 2>/dev/null || true)" && [[ -n "$ORIGIN_URL" ]]; then
  if REPO_SLUG="$(parse_repo_slug "$ORIGIN_URL" 2>/dev/null)"; then
    log "origin repository: ${REPO_SLUG}"
    EXPECTED_HAXELIB_URL="https://github.com/${REPO_SLUG}"
    if [[ "$HAXELIB_URL" != "$EXPECTED_HAXELIB_URL" ]]; then
      fail "haxelib.json url (${HAXELIB_URL}) does not match origin repository (${EXPECTED_HAXELIB_URL})"
    fi
    log "haxelib url: ${HAXELIB_URL}"
    if command -v gh >/dev/null 2>&1; then
      if RELEASE_JSON="$(gh release view "$LATEST_TAG" --repo "$REPO_SLUG" --json tagName,isDraft,isImmutable,isPrerelease,url,publishedAt,assets 2>/dev/null)" && [[ -n "$RELEASE_JSON" ]]; then
        RELEASE_URL="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); d.url || ''")"
        RELEASE_DRAFT="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(!!d.isDraft)")"
        RELEASE_IMMUTABLE="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(!!d.isImmutable)")"
        RELEASE_PRERELEASE="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(!!d.isPrerelease)")"
        RELEASE_PUBLISHED_AT="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); d.publishedAt || ''")"
        RELEASE_ASSET_COUNT="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(Array.isArray(d.assets) ? d.assets.length : 0)")"
        log "remote release visibility (GitHub API): tag=${LATEST_TAG} draft=${RELEASE_DRAFT} immutable=${RELEASE_IMMUTABLE} prerelease=${RELEASE_PRERELEASE} published_at=${RELEASE_PUBLISHED_AT} assets=${RELEASE_ASSET_COUNT}"
        if [[ -n "$RELEASE_URL" ]]; then
          log "remote release URL: ${RELEASE_URL}"
        fi
      else
        warn "no GitHub release found for ${LATEST_TAG} (tag exists); publish may still be pending"
      fi
    fi
  fi
fi

log "OK"
