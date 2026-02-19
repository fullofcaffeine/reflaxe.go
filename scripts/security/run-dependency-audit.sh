#!/usr/bin/env bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT"

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

if ! command -v govulncheck >/dev/null 2>&1; then
  echo "[deps] Installing govulncheck"
  go install golang.org/x/vuln/cmd/govulncheck@latest
  export PATH="$(go env GOPATH)/bin:$PATH"
fi

echo "[deps] govulncheck (runtime/hxrt package)"
GO111MODULE=off govulncheck ./runtime/hxrt

echo "[deps] dependency audit passed"
