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
require_command node
require_command python3
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
if [[ "${CI:-false}" == "true" && "${RELEASE_UPSTREAM_GATES_SHA:-}" != "$TESTED_SHA" ]]; then
  fail "RELEASE_UPSTREAM_GATES_SHA must match tested SHA $TESTED_SHA before semantic-release"
fi
UPSTREAM_EVIDENCE="${RELEASE_UPSTREAM_EVIDENCE:-}"
if [[ "${CI:-false}" == "true" && ! -f "$UPSTREAM_EVIDENCE" ]]; then
  fail "RELEASE_UPSTREAM_EVIDENCE must name structured gate evidence before semantic-release"
fi
BLOCKER_EVIDENCE="${RELEASE_BLOCKER_EVIDENCE:-}"
if [[ "${CI:-false}" == "true" && ! -f "$BLOCKER_EVIDENCE" ]]; then
  fail "RELEASE_BLOCKER_EVIDENCE must name remote tracker evidence before semantic-release"
fi

python3 scripts/release/verify-license-policy.py --mode release
log "approved licensing policy: OK"

RELEASE_BIN="${SEMANTIC_RELEASE_BIN:-$ROOT_DIR/node_modules/.bin/semantic-release}"
if [[ ! -x "$RELEASE_BIN" ]]; then
  fail "semantic-release executable not found or not executable: $RELEASE_BIN"
fi

BEFORE_TAGS="$(mktemp)"
AFTER_TAGS="$(mktemp)"
NEW_TAGS="$(mktemp)"
REMOVED_TAGS="$(mktemp)"
EXACT_TAGS="$(mktemp)"
RELEASE_WORK_DIR=""
cleanup() {
  rm -f "$BEFORE_TAGS" "$AFTER_TAGS" "$NEW_TAGS" "$REMOVED_TAGS" "$EXACT_TAGS"
  if [[ -n "$RELEASE_WORK_DIR" ]]; then
    rm -rf "$RELEASE_WORK_DIR"
  fi
}
trap cleanup EXIT

DRY_RUN=false
for argument in "$@"; do
  if [[ "$argument" == "--dry-run" ]]; then
    DRY_RUN=true
  fi
done

git tag --list 'v*' | LC_ALL=C sort >"$BEFORE_TAGS"
set +e
"$RELEASE_BIN" "$@"
SEMANTIC_RELEASE_STATUS=$?
set -e

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

if [[ "$DRY_RUN" == "true" ]]; then
  if [[ "$SEMANTIC_RELEASE_STATUS" != "0" ]]; then
    fail "semantic-release dry run failed with status $SEMANTIC_RELEASE_STATUS"
  fi
  log "same-tested-SHA release contract: OK (dry run; no hosted mutation)"
  exit 0
fi

if [[ -s "$NEW_TAGS" ]]; then
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
fi

while IFS= read -r candidate; do
  if [[ "$candidate" == "v0.0.0" ]]; then
    fail "the v0.0.0 development sentinel cannot become a release tag"
  fi
  if [[ "$candidate" =~ ^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]; then
    printf '%s\n' "$candidate" >>"$EXACT_TAGS"
  fi
done < <(git tag --points-at "$TESTED_SHA" --list 'v*' | LC_ALL=C sort)

EXACT_TAG_COUNT="$(wc -l <"$EXACT_TAGS" | tr -d ' ')"
if [[ "$EXACT_TAG_COUNT" == "0" ]]; then
  if [[ "$SEMANTIC_RELEASE_STATUS" != "0" ]]; then
    fail "semantic-release failed before creating an exact release tag at the tested SHA"
  fi
  log "same-tested-SHA release contract: OK (no new release tag)"
  exit 0
fi
if [[ "$EXACT_TAG_COUNT" != "1" ]]; then
  fail "tested SHA has multiple canonical release tags; expected exactly one"
fi

RELEASE_TAG="$(sed -n '1p' "$EXACT_TAGS")"
if [[ -s "$NEW_TAGS" ]]; then
  if [[ "$RELEASE_TAG" != "$NEW_TAG" ]]; then
    fail "new release tag $NEW_TAG is not the one canonical tag at the tested SHA"
  fi
  log "completing new exact tag $RELEASE_TAG"
else
  log "completing or verifying existing exact tag $RELEASE_TAG"
fi
RELEASE_VERSION="${RELEASE_TAG#v}"

if [[ -z "${GITHUB_REPOSITORY:-}" ]]; then
  fail "GITHUB_REPOSITORY is required to complete a tagged release"
fi

if [[ "${CI:-false}" == "true" ]]; then
  REMOTE_REFS="$(git ls-remote --tags origin "refs/tags/$RELEASE_TAG" "refs/tags/$RELEASE_TAG^{}")" \
    || fail "cannot query origin for release tag $RELEASE_TAG"
  REMOTE_COMMIT=""
  while IFS=$'\t' read -r object_id ref_name; do
    if [[ "$ref_name" == "refs/tags/$RELEASE_TAG" && -z "$REMOTE_COMMIT" ]]; then
      REMOTE_COMMIT="$object_id"
    elif [[ "$ref_name" == "refs/tags/$RELEASE_TAG^{}" ]]; then
      REMOTE_COMMIT="$object_id"
    fi
  done <<<"$REMOTE_REFS"
  if [[ "$REMOTE_COMMIT" != "$TESTED_SHA" ]]; then
    fail "origin tag $RELEASE_TAG resolves to ${REMOTE_COMMIT:-missing}, not tested SHA $TESTED_SHA"
  fi
fi

RELEASE_ARTIFACT_BUILDER="${RELEASE_ARTIFACT_BUILDER:-$ROOT_DIR/scripts/release/build-haxelib-artifact.py}"
RELEASE_ASSET_VERIFIER="${RELEASE_ASSET_VERIFIER:-$ROOT_DIR/scripts/release/verify-release-assets.py}"
RELEASE_RECONCILER="${RELEASE_RECONCILER:-$ROOT_DIR/scripts/release/reconcile-github-release.mjs}"
RELEASE_READINESS_COLLECTOR="${RELEASE_READINESS_COLLECTOR:-$ROOT_DIR/scripts/release/collect-release-readiness.py}"
RELEASE_READINESS_VERIFIER="${RELEASE_READINESS_VERIFIER:-$ROOT_DIR/scripts/release/verify-release-readiness.py}"
RELEASE_READINESS_POLICY="${RELEASE_READINESS_POLICY:-$ROOT_DIR/release/readiness-policy.json}"
for required_file in \
  "$RELEASE_ARTIFACT_BUILDER" \
  "$RELEASE_ASSET_VERIFIER" \
  "$RELEASE_RECONCILER" \
  "$RELEASE_READINESS_COLLECTOR" \
  "$RELEASE_READINESS_VERIFIER" \
  "$RELEASE_READINESS_POLICY"; do
  if [[ ! -f "$required_file" ]]; then
    fail "release completion helper is missing: $required_file"
  fi
done
UPSTREAM_GATES_SHA="${RELEASE_UPSTREAM_GATES_SHA:-}"
if [[ "$UPSTREAM_GATES_SHA" != "$TESTED_SHA" ]]; then
  fail "RELEASE_UPSTREAM_GATES_SHA must match tested SHA $TESTED_SHA"
fi
if [[ ! -f "$UPSTREAM_EVIDENCE" ]]; then
  fail "RELEASE_UPSTREAM_EVIDENCE must name structured gate evidence"
fi
if [[ ! -f "$BLOCKER_EVIDENCE" ]]; then
  fail "RELEASE_BLOCKER_EVIDENCE must name remote tracker evidence"
fi

RELEASE_WORK_DIR="$(mktemp -d "${RUNNER_TEMP:-${TMPDIR:-/tmp}}/haxe-go-release.XXXXXX")"
ARTIFACT_DIR="$RELEASE_WORK_DIR/artifacts"
python3 "$RELEASE_ARTIFACT_BUILDER" \
  --version "$RELEASE_VERSION" \
  --tag "$RELEASE_TAG" \
  --source-sha "$TESTED_SHA" \
  --output-dir "$ARTIFACT_DIR"
ASSET_MANIFEST="$ARTIFACT_DIR/release-assets.json"
python3 "$RELEASE_ASSET_VERIFIER" \
  --assets "$ASSET_MANIFEST" \
  --version "$RELEASE_VERSION" \
  --tag "$RELEASE_TAG" \
  --source-sha "$TESTED_SHA"
CANDIDATE_EVIDENCE="$RELEASE_WORK_DIR/readiness-candidate.json"
python3 "$RELEASE_READINESS_COLLECTOR" \
  --phase candidate \
  --version "$RELEASE_VERSION" \
  --tag "$RELEASE_TAG" \
  --tested-sha "$TESTED_SHA" \
  --upstream-evidence "$UPSTREAM_EVIDENCE" \
  --blocker-evidence "$BLOCKER_EVIDENCE" \
  --assets "$ASSET_MANIFEST" \
  --output "$CANDIDATE_EVIDENCE"
python3 "$RELEASE_READINESS_VERIFIER" \
  --policy "$RELEASE_READINESS_POLICY" \
  --evidence "$CANDIDATE_EVIDENCE" \
  --mode live
node "$RELEASE_RECONCILER" \
  --repository "$GITHUB_REPOSITORY" \
  --assets "$ASSET_MANIFEST"
PUBLISHED_EVIDENCE="$RELEASE_WORK_DIR/readiness-published.json"
python3 "$RELEASE_READINESS_COLLECTOR" \
  --phase published \
  --version "$RELEASE_VERSION" \
  --tag "$RELEASE_TAG" \
  --tested-sha "$TESTED_SHA" \
  --upstream-evidence "$UPSTREAM_EVIDENCE" \
  --blocker-evidence "$BLOCKER_EVIDENCE" \
  --assets "$ASSET_MANIFEST" \
  --output "$PUBLISHED_EVIDENCE"
python3 "$RELEASE_READINESS_VERIFIER" \
  --policy "$RELEASE_READINESS_POLICY" \
  --evidence "$PUBLISHED_EVIDENCE" \
  --mode live

if [[ "$(git rev-parse HEAD)" != "$TESTED_SHA" ]]; then
  fail "release completion moved HEAD away from tested SHA $TESTED_SHA"
fi
if [[ -n "$(git status --porcelain=v1 --untracked-files=no)" ]]; then
  fail "tracked checkout changed during release completion"
fi

log "same-tested-SHA release and asset contract: OK ($RELEASE_TAG -> $TESTED_SHA)"
