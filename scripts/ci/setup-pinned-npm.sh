#!/usr/bin/env bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT"

# What: Activate the exact npm version declared by package.json.
# Why: Node ships npm, but Node's Corepack documentation states that the
# packageManager field does not replace the global npm executable.
# How: Use the bundled npm only as a bootstrap for the exact, script-disabled
# package, then fail unless the npm found on PATH reports that version.
package_manager="$(
  node -e '
    const value = require("./package.json").packageManager;
    if (typeof value !== "string" || !/^npm@[0-9]+\.[0-9]+\.[0-9]+$/.test(value)) {
      process.stderr.write("package.json packageManager must be an exact npm version\n");
      process.exit(1);
    }
    process.stdout.write(value);
  '
)"
expected_version="${package_manager#npm@}"

echo "[npm-bootstrap] installing ${package_manager}"
npm install --global --ignore-scripts --no-audit --no-fund "$package_manager"
hash -r

actual_version="$(npm --version)"
if [[ "$actual_version" != "$expected_version" ]]; then
  echo "[npm-bootstrap] error: expected npm ${expected_version}, got ${actual_version}" >&2
  exit 1
fi

echo "[npm-bootstrap] npm ${actual_version} active"
