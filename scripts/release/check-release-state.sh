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
  if ! grep -Fq "$pattern" "$path"; then
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

TAG_REGEX='^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$'
SEMVERS="$(git tag --merged HEAD | grep -E "$TAG_REGEX" || true)"
if [[ -z "$SEMVERS" ]]; then
  fail "no semver tag is reachable from current HEAD; semantic-release may treat this as an initial release"
fi

LATEST_TAG="$(printf '%s\n' "$SEMVERS" | sort -V | tail -n 1)"
TAG_COUNT="$(printf '%s\n' "$SEMVERS" | sed '/^$/d' | wc -l | tr -d ' ')"
log "reachable semver tags: $TAG_COUNT"
log "latest reachable semver tag: $LATEST_TAG"

PACKAGE_VERSION="$(node -p "require('./package.json').version")"
HAXELIB_VERSION="$(node -p "require('./haxelib.json').version")"
EXPECTED_TAG="v${PACKAGE_VERSION}"

if [[ "$PACKAGE_VERSION" != "$HAXELIB_VERSION" ]]; then
  fail "package.json (${PACKAGE_VERSION}) and haxelib.json (${HAXELIB_VERSION}) versions differ"
fi
log "package/haxelib version parity: ${PACKAGE_VERSION}"

if [[ "$LATEST_TAG" != "$EXPECTED_TAG" ]]; then
  warn "latest tag (${LATEST_TAG}) differs from package version tag (${EXPECTED_TAG}); this can be normal between releases"
fi

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
log "ci harness semantic-release wiring: OK"

if ORIGIN_URL="$(git remote get-url origin 2>/dev/null || true)" && [[ -n "$ORIGIN_URL" ]]; then
  if REPO_SLUG="$(parse_repo_slug "$ORIGIN_URL" 2>/dev/null)"; then
    log "origin repository: ${REPO_SLUG}"
    if command -v gh >/dev/null 2>&1; then
      if RELEASE_JSON="$(gh release view "$LATEST_TAG" --repo "$REPO_SLUG" --json tagName,isDraft,isPrerelease,url,publishedAt 2>/dev/null)" && [[ -n "$RELEASE_JSON" ]]; then
        RELEASE_URL="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); d.url || ''")"
        RELEASE_DRAFT="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(!!d.isDraft)")"
        RELEASE_PRERELEASE="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); String(!!d.isPrerelease)")"
        RELEASE_PUBLISHED_AT="$(printf '%s\n' "$RELEASE_JSON" | node -p "const d=JSON.parse(require('fs').readFileSync(0,'utf8')); d.publishedAt || ''")"
        log "remote release visibility: tag=${LATEST_TAG} draft=${RELEASE_DRAFT} prerelease=${RELEASE_PRERELEASE} published_at=${RELEASE_PUBLISHED_AT}"
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
