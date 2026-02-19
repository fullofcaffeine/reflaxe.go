#!/usr/bin/env bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT"

govulncheck_version="${GOVULNCHECK_VERSION:-latest}"
govulncheck_install_attempts="${GOVULNCHECK_INSTALL_ATTEMPTS:-3}"
govulncheck_retry_delay_sec="${GOVULNCHECK_INSTALL_RETRY_DELAY_SEC:-2}"
govulncheck_allow_install_failure="${GOVULNCHECK_ALLOW_INSTALL_FAILURE:-0}"

if ! [[ "$govulncheck_install_attempts" =~ ^[1-9][0-9]*$ ]]; then
  echo "[deps] error: GOVULNCHECK_INSTALL_ATTEMPTS must be a positive integer" >&2
  exit 2
fi

if ! [[ "$govulncheck_retry_delay_sec" =~ ^[0-9]+$ ]]; then
  echo "[deps] error: GOVULNCHECK_INSTALL_RETRY_DELAY_SEC must be an integer >= 0" >&2
  exit 2
fi

if [[ -f package.json ]]; then
  echo "[deps] npm audit (production dependencies, high+ severity)"
  npm_audit_tmp_dir="$(mktemp -d)"
  trap 'rm -rf "$npm_audit_tmp_dir"' EXIT
  cp package.json "$npm_audit_tmp_dir/package.json"
  if [[ -f .npmrc ]]; then
    cp .npmrc "$npm_audit_tmp_dir/.npmrc"
  fi
  (
    cd "$npm_audit_tmp_dir"
    npm install --ignore-scripts --package-lock-only --no-audit --no-fund
    npm audit --omit=dev --audit-level=high
  )
fi

if [[ "${SKIP_GOVULNCHECK:-0}" == "1" ]]; then
  echo "[deps] SKIP_GOVULNCHECK=1, skipping govulncheck"
  echo "[deps] dependency audit passed"
  exit 0
fi

if ! command -v go >/dev/null 2>&1; then
  echo "[deps] error: go toolchain is required for govulncheck" >&2
  echo "[deps] hint: install Go or run local-only bypass with SKIP_GOVULNCHECK=1" >&2
  exit 1
fi

ensure_govulncheck() {
  if command -v govulncheck >/dev/null 2>&1; then
    return 0
  fi

  local install_ref="golang.org/x/vuln/cmd/govulncheck@$govulncheck_version"
  local attempt=1
  while (( attempt <= govulncheck_install_attempts )); do
    echo "[deps] Installing govulncheck ($install_ref), attempt $attempt/$govulncheck_install_attempts"
    if go install "$install_ref"; then
      break
    fi

    if (( attempt < govulncheck_install_attempts )); then
      local backoff=$((govulncheck_retry_delay_sec * attempt))
      echo "[deps] govulncheck install failed; retrying in ${backoff}s"
      sleep "$backoff"
    fi
    attempt=$((attempt + 1))
  done

  export PATH="$(go env GOPATH)/bin:$PATH"
  command -v govulncheck >/dev/null 2>&1
}

if ! ensure_govulncheck; then
  if [[ "${CI:-}" == "true" ]]; then
    echo "[deps] error: govulncheck install failed after $govulncheck_install_attempts attempts (CI mode)" >&2
    exit 1
  fi

  if [[ "$govulncheck_allow_install_failure" == "1" ]]; then
    echo "[deps] warning: govulncheck unavailable after $govulncheck_install_attempts attempts; continuing (local soft-fail enabled)"
    echo "[deps] warning: set GOVULNCHECK_ALLOW_INSTALL_FAILURE=0 to enforce hard-fail locally"
    echo "[deps] dependency audit passed (partial: npm audit only)"
    exit 0
  fi

  echo "[deps] error: govulncheck install failed after $govulncheck_install_attempts attempts" >&2
  echo "[deps] hint: verify network/proxy access to proxy.golang.org and rerun" >&2
  echo "[deps] hint: local-only bypass: SKIP_GOVULNCHECK=1 npm run security:deps" >&2
  echo "[deps] hint: local soft-fail: GOVULNCHECK_ALLOW_INSTALL_FAILURE=1 npm run security:deps" >&2
  exit 1
fi

echo "[deps] govulncheck (runtime/hxrt package)"
govuln_tmp_dir="$(mktemp -d)"
cp -R runtime/hxrt/. "$govuln_tmp_dir/"
cat >"$govuln_tmp_dir/go.mod" <<'EOF'
module reflaxe_go_hxrt_audit

go 1.23
EOF

if ! (
  cd "$govuln_tmp_dir"
  govulncheck ./...
); then
  govuln_status=$?
  rm -rf "$govuln_tmp_dir"
  exit "$govuln_status"
fi
rm -rf "$govuln_tmp_dir"

echo "[deps] dependency audit passed"
