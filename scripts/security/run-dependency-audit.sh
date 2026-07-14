#!/usr/bin/env bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT"

# This explicit tool version makes audit output reproducible. It does not define
# supported Go build lines; docs/toolchain-policy.md is authoritative for those.
# The fail-closed vulnerability-policy slice owns future govulncheck upgrades.
govulncheck_version="${GOVULNCHECK_VERSION:-v1.6.0}"
govulncheck_bin="${GOVULNCHECK_BIN:-}"
govulncheck_install_attempts="${GOVULNCHECK_INSTALL_ATTEMPTS:-3}"
govulncheck_retry_delay_sec="${GOVULNCHECK_INSTALL_RETRY_DELAY_SEC:-2}"
govulncheck_allow_install_failure="${GOVULNCHECK_ALLOW_INSTALL_FAILURE:-0}"
govulncheck_report_dir="${GOVULNCHECK_REPORT_DIR:-$ROOT/.cache/security/dependency-audit}"
govulncheck_report="$govulncheck_report_dir/govulncheck.txt"
govulncheck_metadata="$govulncheck_report_dir/metadata.txt"
npm_audit_report="$govulncheck_report_dir/npm-audit.txt"

mkdir -p "$govulncheck_report_dir"
: >"$govulncheck_report"
: >"$npm_audit_report"

detected_go_version="unavailable"
if command -v go >/dev/null 2>&1; then
  detected_go_version="$(go version 2>/dev/null || printf 'unavailable')"
fi

write_govulncheck_metadata() {
  local result="$1"
  local exit_code="$2"
  {
    printf 'result=%s\n' "$result"
    printf 'govulncheck_exit=%s\n' "$exit_code"
    printf 'govulncheck_version=%s\n' "$govulncheck_version"
    printf 'go_version=%s\n' "$detected_go_version"
  } >"$govulncheck_metadata"
}

write_govulncheck_metadata "not_run" "-"

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
  ) 2>&1 | tee "$npm_audit_report"
fi

if [[ "${SKIP_GOVULNCHECK:-0}" == "1" ]]; then
  if [[ "${CI:-}" == "true" ]]; then
    write_govulncheck_metadata "ci_skip_rejected" "1"
    echo "[deps] error: SKIP_GOVULNCHECK=1 is not permitted in CI" >&2
    exit 1
  fi
  write_govulncheck_metadata "locally_skipped" "0"
  echo "[deps] SKIP_GOVULNCHECK=1, skipping govulncheck"
  echo "[deps] dependency audit passed"
  exit 0
fi

if ! command -v go >/dev/null 2>&1; then
  write_govulncheck_metadata "tool_error" "1"
  echo "[deps] error: go toolchain is required for govulncheck" >&2
  echo "[deps] hint: install Go or run local-only bypass with SKIP_GOVULNCHECK=1" >&2
  exit 1
fi

ensure_govulncheck() {
  if [[ -n "$govulncheck_bin" ]]; then
    if [[ "$govulncheck_bin" == */* ]]; then
      [[ -x "$govulncheck_bin" ]]
      return
    fi
    if command -v "$govulncheck_bin" >/dev/null 2>&1; then
      govulncheck_bin="$(command -v "$govulncheck_bin")"
      return 0
    fi
    return 1
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

  govulncheck_bin="$(go env GOPATH)/bin/govulncheck"
  [[ -x "$govulncheck_bin" ]]
}

print_sanitized_govulncheck_report() {
  local log_file="$1"

  # Keep the upstream report visible, but avoid GitHub problem matchers and raw
  # workflow commands turning report traces into unowned annotations.
  sed -E \
    -e 's/^::/: :/' \
    -e 's#([[:alnum:]_./-]+\.go):([0-9]+):([0-9]+):#\1 line \2 col \3:#g' \
    "$log_file"
}

if ! ensure_govulncheck; then
  write_govulncheck_metadata "install_error" "1"
  if [[ "${CI:-}" == "true" ]]; then
    echo "[deps] error: govulncheck install failed after $govulncheck_install_attempts attempts (CI mode)" >&2
    exit 1
  fi

  if [[ "$govulncheck_allow_install_failure" == "1" ]]; then
    write_govulncheck_metadata "local_install_soft_fail" "0"
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

go 1.22
EOF

govuln_log="$govuln_tmp_dir/govulncheck.log"
set +e
(
  cd "$govuln_tmp_dir"
  env -u GITHUB_ACTIONS "$govulncheck_bin" -show traces -format=text ./...
) >"$govuln_log" 2>&1
govuln_status=$?
set -e

print_sanitized_govulncheck_report "$govuln_log" | tee "$govulncheck_report"

if [[ "$govuln_status" -ne 0 ]]; then
  if grep -Eq '^Vulnerability #|^Your code is affected by' "$govuln_log"; then
    write_govulncheck_metadata "reachable_vulnerabilities" "$govuln_status"
    echo "[deps] error: govulncheck reported reachable vulnerabilities; reachable vulnerabilities are release-blocking" >&2
    echo "[deps] details: docs/security-dependency-audit.md" >&2
    if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
      echo "::error title=Reachable Go vulnerabilities::govulncheck found reachable vulnerabilities in runtime/hxrt; see the uploaded dependency-audit report and docs/security-dependency-audit.md."
    fi
    rm -rf "$govuln_tmp_dir"
    exit 1
  fi

  write_govulncheck_metadata "tool_error" "$govuln_status"
  echo "[deps] error: govulncheck failed before producing a clean scan; see $govulncheck_report" >&2
  rm -rf "$govuln_tmp_dir"
  exit "$govuln_status"
fi

write_govulncheck_metadata "clean" "0"
rm -rf "$govuln_tmp_dir"

echo "[deps] dependency audit passed"
