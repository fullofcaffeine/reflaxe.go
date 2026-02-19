#!/usr/bin/env bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$project_root"

haxe_version="${HAXE_VERSION:-4.3.7}"
reflaxe_go_source="${REFLAXE_GO_SOURCE:-github:fullofcaffeine/reflaxe.go}"
mode="all"

usage() {
	cat <<'EOUSAGE'
Usage: scripts/dev/setup.sh [--haxe-only | --target-only]

Bootstraps lix scope, reflaxe.go dependency, and Haxe toolchain.

Environment:
  HAXE_VERSION        Default: 4.3.7
  REFLAXE_GO_SOURCE   Default: github:fullofcaffeine/reflaxe.go
EOUSAGE
}

while [[ $# -gt 0 ]]; do
	case "$1" in
	--haxe-only)
		mode="haxe"
		shift
		;;
	--target-only)
		mode="target"
		shift
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		echo "[setup] error: unknown arg: $1" >&2
		usage >&2
		exit 2
		;;
	esac
done

if ! command -v npx >/dev/null 2>&1; then
	echo "[setup] error: npx is required (install Node.js first)." >&2
	exit 1
fi

if [[ ! -d "$project_root/haxe_libraries" ]]; then
	echo "[setup] creating lix scope"
	npx lix scope create
fi

if [[ "$mode" == "all" || "$mode" == "target" ]]; then
	echo "[setup] installing reflaxe.go from: $reflaxe_go_source"
	npx lix install "$reflaxe_go_source"
fi

if [[ "$mode" == "all" || "$mode" == "haxe" ]]; then
	echo "[setup] ensuring Haxe $haxe_version in lix scope"
	npx lix download haxe "$haxe_version"
	npx lix use haxe "$haxe_version"
fi

echo "[setup] downloading lix-managed toolchain dependencies"
npx lix download

if command -v go >/dev/null 2>&1; then
	echo "[setup] go: $(go version)"
else
	echo "[setup] warning: Go toolchain not found in PATH. Install Go before running hx:run/hx:build."
fi

echo "[setup] done"
