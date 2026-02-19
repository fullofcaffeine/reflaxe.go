#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
template_dir="$root_dir/templates/basic"

usage() {
	cat <<'EOUSAGE'
Usage: scripts/dev/new-project.sh <target-dir> [--force]

Scaffolds a new Haxe -> Go project from templates/basic.

Options:
  --force   Replace target directory if it already exists.
  -h, --help
EOUSAGE
}

to_display_path() {
	local path="$1"
	local cwd
	cwd="$(pwd)"
	if [[ "$path" == "$cwd/"* ]]; then
		printf ".%s\n" "${path#$cwd}"
		return 0
	fi
	if [[ "$path" == "$root_dir/"* ]]; then
		printf ".%s\n" "${path#$root_dir}"
		return 0
	fi
	printf "<external-path>\n"
}

to_abs_path() {
	local path="$1"
	if [[ "$path" == /* ]]; then
		printf "%s\n" "$path"
		return 0
	fi
	printf "%s/%s\n" "$(pwd)" "$path"
}

sanitize_module_name() {
	local raw="$1"
	local out
	out="$(printf '%s' "$raw" | tr '[:upper:]' '[:lower:]')"
	out="$(printf '%s' "$out" | sed -E 's/[^a-z0-9]+/_/g; s/^_+//; s/_+$//')"
	if [[ -z "$out" ]]; then
		out="hx_app"
	fi
	if [[ "$out" =~ ^[0-9] ]]; then
		out="hx_${out}"
	fi
	printf "%s\n" "$out"
}

rewrite_token() {
	local file="$1"
	local module_name="$2"
	local tmp
	tmp="$(mktemp "${TMPDIR:-/tmp}/reflaxe-go-new-project.XXXXXX")"
	sed "s/hx_app/${module_name}/g" "$file" > "$tmp"
	mv "$tmp" "$file"
}

target_arg=""
force=0
while [[ $# -gt 0 ]]; do
	case "$1" in
	--force)
		force=1
		shift
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		if [[ -n "$target_arg" ]]; then
			echo "[new-project] error: unexpected extra arg: $1" >&2
			usage >&2
			exit 2
		fi
		target_arg="$1"
		shift
		;;
	esac
done

if [[ -z "$target_arg" ]]; then
	echo "[new-project] error: missing <target-dir>" >&2
	usage >&2
	exit 2
fi

if [[ ! -d "$template_dir" ]]; then
	echo "[new-project] error: missing template directory: $template_dir" >&2
	exit 1
fi

target_dir="$(to_abs_path "$target_arg")"
target_display="$(to_display_path "$target_dir")"

if [[ -e "$target_dir" ]]; then
	if [[ "$force" -eq 1 ]]; then
		rm -rf "$target_dir"
	else
		echo "[new-project] error: target already exists: $target_display (use --force to replace)" >&2
		exit 2
	fi
fi

mkdir -p "$target_dir"
cp -R "$template_dir"/. "$target_dir"/

project_name="$(basename "$target_dir")"
module_name="$(sanitize_module_name "$project_name")"

while IFS= read -r -d '' file; do
	if grep -q "hx_app" "$file"; then
		rewrite_token "$file" "$module_name"
	fi
done < <(find "$target_dir" -type f -print0)

echo "[new-project] created $target_display"
echo "[new-project] go module ${module_name}"
echo "[new-project] next:"
echo "  cd ${target_display}"
echo "  npm install"
echo "  npm run setup"
echo "  npm run hx:run"
