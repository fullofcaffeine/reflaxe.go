#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
output="${1:-dist/reflaxe.go.zip}"
source_root="${2:-$root_dir}"

if [[ "$output" != /* ]]; then
  output="$root_dir/$output"
fi
if [[ "$source_root" != /* ]]; then
  source_root="$root_dir/$source_root"
fi

if ! command -v haxe >/dev/null 2>&1; then
  echo "[package] ERROR: haxe is required" >&2
  exit 2
fi
if ! command -v python3 >/dev/null 2>&1; then
  echo "[package] ERROR: python3 is required" >&2
  exit 2
fi

temporary_root="$(mktemp -d "${TMPDIR:-/tmp}/reflaxe-go-haxelib.XXXXXX")"
cleanup() {
  rm -rf "$temporary_root"
}
trap cleanup EXIT

package_root="$temporary_root/reflaxe.go"

(
  cd "$root_dir"
  haxe --run Run build "$package_root" --source-root "$source_root" --clean
)

python3 "$root_dir/scripts/ci/canonical_stdlib_layout_check.py" \
  --source-root "$source_root" \
  --package-root "$package_root"
python3 "$root_dir/scripts/release/deterministic-zip.py" "$package_root" "$output"
echo "[package] wrote $output"
