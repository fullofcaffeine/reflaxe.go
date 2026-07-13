#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"
EXPECTED_LEGACY_SHA256="0e34e32cb1ac25fdc8592aea85aa5630ca31ab59076b3e33faa6611a4e51911c"
SESSION_CLOSE=0
VERIFY_REMOTE=0

usage() {
  cat <<'EOF'
Usage: scripts/beads/check-health.sh [--session-close] [--verify-remote]

Runs read-only Beads configuration, graph, lint, orphan, archive, and local
version-control checks.

  --verify-remote  Export the active database and a disposable clone of the
                   configured Dolt remote, then require byte-identical output.
  --session-close  Also require a clean tracked Git tree and zero Git
                   ahead/behind count. Implies --verify-remote and is intended
                   to run after bd dolt push and git push.
EOF
}

while (($# > 0)); do
  case "$1" in
    --session-close)
      SESSION_CLOSE=1
      VERIFY_REMOTE=1
      ;;
    --verify-remote)
      VERIFY_REMOTE=1
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "[beads-health] unknown argument: $1" >&2
      usage >&2
      exit 2
      ;;
  esac
  shift
done

cd "$ROOT_DIR"

require_command() {
  local command_name="$1"
  if ! command -v "$command_name" >/dev/null 2>&1; then
    echo "[beads-health] required command is unavailable: $command_name" >&2
    exit 1
  fi
}

sha256_file() {
  local path="$1"
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$path" | awk '{print $1}'
    return
  fi
  shasum -a 256 "$path" | awk '{print $1}'
}

require_command bd
require_command git
require_command awk

echo "[beads-health] validating configuration and graph"
bd config validate
bd dep cycles
bd lint
bd orphans

echo "[beads-health] local Dolt status"
bd vc status --json
bd stats --no-activity

legacy_hash="$(sha256_file .beads/issues.jsonl)"
if [[ "$legacy_hash" != "$EXPECTED_LEGACY_SHA256" ]]; then
  echo "[beads-health] legacy archive hash drifted: $legacy_hash" >&2
  exit 1
fi
echo "[beads-health] legacy archive hash: $legacy_hash"

verify_remote_export() {
  local remote_url issue_prefix temp_dir local_export remote_export local_hash remote_hash
  remote_url="$(bd config get sync.remote)"
  issue_prefix="$(bd config get issue_prefix)"
  if [[ -z "$remote_url" || -z "$issue_prefix" ]]; then
    echo "[beads-health] sync.remote or issue_prefix is not configured" >&2
    exit 1
  fi

  temp_dir="$(mktemp -d "${TMPDIR:-/tmp}/haxe-go-beads-health.XXXXXX")"
  trap 'rm -rf "$temp_dir"' RETURN
  local_export="$temp_dir/local.jsonl"
  remote_export="$temp_dir/remote.jsonl"

  echo "[beads-health] exporting local Dolt state"
  bd export --include-memories -o "$local_export" >/dev/null

  mkdir -p "$temp_dir/remote"
  git -C "$temp_dir/remote" init -q
  (
    cd "$temp_dir/remote"
    bd init \
      --prefix "$issue_prefix" \
      --remote "$remote_url" \
      --skip-agents \
      --skip-hooks \
      --non-interactive \
      --quiet
  )
  bd -C "$temp_dir/remote" export --include-memories -o "$remote_export" >/dev/null

  local_hash="$(sha256_file "$local_export")"
  remote_hash="$(sha256_file "$remote_export")"
  if [[ "$local_hash" != "$remote_hash" ]] || ! cmp -s "$local_export" "$remote_export"; then
    echo "[beads-health] local and remote Dolt exports differ" >&2
    echo "[beads-health] local:  $local_hash" >&2
    echo "[beads-health] remote: $remote_hash" >&2
    exit 1
  fi
  echo "[beads-health] remote export matches: $local_hash"
}

if ((VERIFY_REMOTE)); then
  require_command cmp
  remote_ref="$(git ls-remote origin refs/dolt/data)"
  if [[ -z "$remote_ref" ]]; then
    echo "[beads-health] origin does not expose refs/dolt/data" >&2
    exit 1
  fi
  echo "[beads-health] remote Dolt ref: ${remote_ref%%[[:space:]]*}"
  verify_remote_export
fi

if ((SESSION_CLOSE)); then
  if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "[beads-health] tracked Git changes remain; commit them before session close" >&2
    exit 1
  fi

  upstream="$(git rev-parse --abbrev-ref --symbolic-full-name '@{upstream}' 2>/dev/null || true)"
  if [[ -z "$upstream" ]]; then
    echo "[beads-health] current branch has no upstream" >&2
    exit 1
  fi
  remote_name="${upstream%%/*}"
  git fetch --quiet "$remote_name"
  read -r behind ahead < <(git rev-list --left-right --count "$upstream...HEAD")
  echo "[beads-health] Git upstream: $upstream (behind=$behind ahead=$ahead)"
  if [[ "$behind" != "0" || "$ahead" != "0" ]]; then
    echo "[beads-health] Git branch is not synchronized with its upstream" >&2
    exit 1
  fi
fi

echo "[beads-health] OK"
