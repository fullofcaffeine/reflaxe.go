#!/usr/bin/env bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$project_root"

usage() {
	cat <<'EOUSAGE'
Usage: scripts/dev/hx-go.sh <compile|run|build|test> [--profile portable|gopher|metal] [--hxml <path>] [--out <dir>] [--binary <path>]

Notes:
  - compile uses backend defaults (includes auto go build unless disabled by define).
  - run/build/test force -D go_no_build and run Go commands explicitly.

Examples:
  bash scripts/dev/hx-go.sh run
  bash scripts/dev/hx-go.sh build --profile gopher
  bash scripts/dev/hx-go.sh test --hxml compile.custom.hxml --out out_custom
EOUSAGE
}

if [[ $# -lt 1 ]]; then
	usage >&2
	exit 2
fi

if [[ "$1" == "-h" || "$1" == "--help" ]]; then
	usage
	exit 0
fi

action="$1"
shift

profile="portable"
hxml=""
out_dir=""
binary=""

while [[ $# -gt 0 ]]; do
	case "$1" in
	--profile)
		profile="${2:-}"
		shift 2
		;;
	--hxml)
		hxml="${2:-}"
		shift 2
		;;
	--out)
		out_dir="${2:-}"
		shift 2
		;;
	--binary)
		binary="${2:-}"
		shift 2
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		echo "[hx-go] error: unknown arg: $1" >&2
		usage >&2
		exit 2
		;;
	esac
done

case "$profile" in
portable | gopher | metal) ;;
*)
	echo "[hx-go] error: invalid profile '$profile' (expected portable|gopher|metal)" >&2
	exit 2
	;;
esac

if [[ -z "$hxml" ]]; then
	case "$profile" in
	portable) hxml="compile.hxml" ;;
	gopher) hxml="compile.gopher.hxml" ;;
	metal) hxml="compile.metal.hxml" ;;
	esac
fi

if [[ -z "$out_dir" ]]; then
	case "$profile" in
	portable) out_dir="out" ;;
	gopher) out_dir="out_gopher" ;;
	metal) out_dir="out_metal" ;;
	esac
fi

if [[ -z "$binary" ]]; then
	case "$profile" in
	portable) binary="bin/hx_app" ;;
	gopher) binary="bin/hx_app_gopher" ;;
	metal) binary="bin/hx_app_metal" ;;
	esac
fi

if [[ ! -f "$hxml" ]]; then
	echo "[hx-go] error: missing hxml file: $hxml" >&2
	exit 1
fi

if ! command -v haxe >/dev/null 2>&1; then
	echo "[hx-go] error: haxe command not found. Run 'npm run setup' first." >&2
	exit 1
fi

compile() {
	local no_build="${1:-0}"
	echo "[hx-go] compiling via $hxml"
	if [[ "$no_build" == "1" ]]; then
		haxe "$hxml" -D go_no_build
	else
		haxe "$hxml"
	fi
}

run_generated() {
	if [[ ! -d "$out_dir" ]]; then
		echo "[hx-go] error: missing output directory: $out_dir" >&2
		exit 1
	fi
	(
		cd "$out_dir"
		go run .
	)
}

test_generated() {
	if [[ ! -d "$out_dir" ]]; then
		echo "[hx-go] error: missing output directory: $out_dir" >&2
		exit 1
	fi
	(
		cd "$out_dir"
		go test ./...
	)
}

build_generated() {
	if [[ ! -d "$out_dir" ]]; then
		echo "[hx-go] error: missing output directory: $out_dir" >&2
		exit 1
	fi
	local binary_abs="$binary"
	if [[ "$binary_abs" != /* ]]; then
		binary_abs="$project_root/$binary_abs"
	fi
	mkdir -p "$(dirname "$binary_abs")"
	(
		cd "$out_dir"
		go build -o "$binary_abs" .
	)
	echo "[hx-go] built $binary_abs"
}

case "$action" in
compile)
	compile 0
	;;
run)
	compile 1
	run_generated
	;;
test)
	compile 1
	test_generated
	;;
build)
	compile 1
	build_generated
	;;
*)
	echo "[hx-go] error: unknown action '$action' (expected compile|run|build|test)" >&2
	exit 2
	;;
esac
