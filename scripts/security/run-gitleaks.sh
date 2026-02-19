#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(git rev-parse --show-toplevel)"
MODE="full"

if [ "${1:-}" = "--staged" ]; then
  MODE="staged"
fi

if ! command -v gitleaks >/dev/null 2>&1; then
  echo "[gitleaks] ERROR: gitleaks is required but not installed." >&2
  echo "[gitleaks] Install: https://github.com/gitleaks/gitleaks#installing" >&2
  exit 1
fi

GITLEAKS_HELP="$(gitleaks --help 2>&1 || true)"

if [ "$MODE" = "staged" ]; then
  echo "[gitleaks] Scanning staged changes"
  if printf '%s' "$GITLEAKS_HELP" | grep -q '\<protect\>'; then
    gitleaks protect --staged --redact --config "$ROOT_DIR/.gitleaks.toml"
  elif printf '%s' "$GITLEAKS_HELP" | grep -q '\<git\>'; then
    gitleaks git --staged --redact --config "$ROOT_DIR/.gitleaks.toml"
  else
    echo "[gitleaks] ERROR: unsupported gitleaks CLI; expected 'protect' or 'git' command." >&2
    exit 1
  fi
  exit 0
fi

echo "[gitleaks] Scanning repository history"
if printf '%s' "$GITLEAKS_HELP" | grep -q '\<detect\>'; then
  gitleaks detect --source "$ROOT_DIR" --redact --config "$ROOT_DIR/.gitleaks.toml"
elif printf '%s' "$GITLEAKS_HELP" | grep -q '\<git\>'; then
  gitleaks git --redact --config "$ROOT_DIR/.gitleaks.toml"
else
  echo "[gitleaks] ERROR: unsupported gitleaks CLI; expected 'detect' or 'git' command." >&2
  exit 1
fi
