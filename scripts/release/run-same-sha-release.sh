#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT_DIR"

log() {
  printf '[same-sha-release] %s\n' "$*"
}

fail() {
  printf '[same-sha-release] error: %s\n' "$*" >&2
  exit 1
}

require_command() {
  local command="$1"
  if ! command -v "$command" >/dev/null 2>&1; then
    fail "required command not found: $command"
  fi
}

require_command git
require_command comm
require_command sort

TESTED_SHA="${RELEASE_TESTED_SHA:-}"
if [[ "${CI:-false}" == "true" && -z "$TESTED_SHA" ]]; then
  fail "RELEASE_TESTED_SHA is required in CI"
fi
if [[ -z "$TESTED_SHA" ]]; then
  TESTED_SHA="$(git rev-parse HEAD)"
  log "local run defaulted RELEASE_TESTED_SHA to HEAD"
fi
if [[ ! "$TESTED_SHA" =~ ^[0-9a-f]{40}$ ]]; then
  fail "RELEASE_TESTED_SHA must be exactly 40 lowercase hexadecimal characters"
fi

HEAD_BEFORE="$(git rev-parse HEAD)"
if [[ "$HEAD_BEFORE" != "$TESTED_SHA" ]]; then
  fail "current HEAD $HEAD_BEFORE does not match RELEASE_TESTED_SHA $TESTED_SHA"
fi
if [[ -n "$(git status --porcelain=v1 --untracked-files=no)" ]]; then
  fail "tracked checkout is not clean before semantic-release"
fi

RELEASE_BIN="${SEMANTIC_RELEASE_BIN:-$ROOT_DIR/node_modules/.bin/semantic-release}"
if [[ ! -x "$RELEASE_BIN" ]]; then
  fail "semantic-release executable not found or not executable: $RELEASE_BIN"
fi

BEFORE_TAGS="$(mktemp)"
AFTER_TAGS="$(mktemp)"
NEW_TAGS="$(mktemp)"
REMOVED_TAGS="$(mktemp)"
trap 'rm -f "$BEFORE_TAGS" "$AFTER_TAGS" "$NEW_TAGS" "$REMOVED_TAGS"' EXIT

git tag --list 'v*' | LC_ALL=C sort >"$BEFORE_TAGS"
"$RELEASE_BIN" "$@"

HEAD_AFTER="$(git rev-parse HEAD)"
if [[ "$HEAD_AFTER" != "$TESTED_SHA" ]]; then
  fail "semantic-release moved HEAD from tested SHA $TESTED_SHA to $HEAD_AFTER"
fi
if [[ -n "$(git status --porcelain=v1 --untracked-files=no)" ]]; then
  fail "tracked checkout changed during semantic-release"
fi

git tag --list 'v*' | LC_ALL=C sort >"$AFTER_TAGS"
comm -13 "$BEFORE_TAGS" "$AFTER_TAGS" >"$NEW_TAGS"
comm -23 "$BEFORE_TAGS" "$AFTER_TAGS" >"$REMOVED_TAGS"
if [[ -s "$REMOVED_TAGS" ]]; then
  fail "semantic-release removed existing version tags: $(tr '\n' ' ' <"$REMOVED_TAGS")"
fi

if [[ ! -s "$NEW_TAGS" ]]; then
  log "same-tested-SHA release contract: OK (no new release tag)"
  exit 0
fi

NEW_TAG_COUNT="$(wc -l <"$NEW_TAGS" | tr -d ' ')"
if [[ "$NEW_TAG_COUNT" != "1" ]]; then
  fail "semantic-release created $NEW_TAG_COUNT version tags; expected exactly one"
fi

NEW_TAG="$(sed -n '1p' "$NEW_TAGS")"
if [[ "$NEW_TAG" == "v0.0.0" ]]; then
  fail "the v0.0.0 development sentinel cannot become a release tag"
fi
if [[ ! "$NEW_TAG" =~ ^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]; then
  fail "new release tag is not canonical stable SemVer: $NEW_TAG"
fi
TAG_COMMIT="$(git rev-parse "$NEW_TAG^{commit}")"
if [[ "$TAG_COMMIT" != "$TESTED_SHA" ]]; then
  fail "new tag $NEW_TAG resolves to $TAG_COMMIT and does not match tested SHA $TESTED_SHA"
fi

if [[ "${CI:-false}" == "true" ]]; then
  REMOTE_REFS="$(git ls-remote --tags origin "refs/tags/$NEW_TAG" "refs/tags/$NEW_TAG^{}")" \
    || fail "cannot query origin for new tag $NEW_TAG"
  REMOTE_COMMIT=""
  while IFS=$'\t' read -r object_id ref_name; do
    if [[ "$ref_name" == "refs/tags/$NEW_TAG" && -z "$REMOTE_COMMIT" ]]; then
      REMOTE_COMMIT="$object_id"
    elif [[ "$ref_name" == "refs/tags/$NEW_TAG^{}" ]]; then
      REMOTE_COMMIT="$object_id"
    fi
  done <<<"$REMOTE_REFS"
  if [[ "$REMOTE_COMMIT" != "$TESTED_SHA" ]]; then
    fail "origin tag $NEW_TAG resolves to ${REMOTE_COMMIT:-missing}, not tested SHA $TESTED_SHA"
  fi
fi

log "same-tested-SHA release contract: OK ($NEW_TAG -> $TESTED_SHA)"
