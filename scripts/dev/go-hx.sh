#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
invocation_dir="$(pwd)"
cd "$root_dir"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/dev/go-hx.sh [options] [-- <go args...>]

Options:
  --project <path>          Optional. Project directory containing compile*.hxml.
                            Default: current working directory.
  --profile <name>          Optional. portable|gopher|metal.
  --hxml <path>             Optional. Explicit hxml file (relative to --project by default).
  --ci                      Prefer compile*.ci.hxml variants.
  --action <name>           Action: compile|run|build|test|vet|fmt. Default: run.
  --out <dir>               Override go_output directory (passes -D go_output=<dir>).
  --binary <path>           Build output path for --action build.
  --haxe-bin <path>         Haxe binary. Default: $HAXE_BIN or haxe.
  --go-bin <path>           Go binary. Default: $GO_BIN or go.
  --define <k[=v]>          Extra Haxe -D define (repeatable).
  -h, --help                Show this help.

Examples:
  bash scripts/dev/go-hx.sh --project examples/tui_todo --profile portable --action run
  bash scripts/dev/go-hx.sh --project examples/profile_storyboard --profile gopher --ci --action test
  bash scripts/dev/go-hx.sh --project ./my_haxe_go_app --action build --binary bin/my_hx_app
USAGE
}

fail() {
  echo "error: $*" >&2
  exit 2
}

display_path() {
  local input="$1"
  if [[ "$input" == "$invocation_dir" ]]; then
    printf ".\n"
  elif [[ "$input" == "$invocation_dir/"* ]]; then
    printf ".%s\n" "${input#"$invocation_dir"}"
  elif [[ "$input" == "$root_dir" ]]; then
    printf ".\n"
  elif [[ "$input" == "$root_dir/"* ]]; then
    printf "%s\n" "${input#"$root_dir/"}"
  else
    printf "[external:%s]\n" "$(basename "$input")"
  fi
}

normalize_existing_dir() {
  local input="$1"
  if [[ ! -d "$input" ]]; then
    fail "project directory not found: $(display_path "$input")"
  fi
  (cd "$input" && pwd)
}

resolve_path_from_base() {
  local input="$1"
  local base="$2"
  if [[ "$input" == /* ]]; then
    printf "%s\n" "$input"
  else
    printf "%s/%s\n" "$base" "$input"
  fi
}

extract_go_output() {
  local hxml_path="$1"
  awk '
    function trim(v) {
      sub(/^[ \t]+/, "", v)
      sub(/[ \t]+$/, "", v)
      return v
    }
    function dequote(v) {
      if (length(v) >= 2 && substr(v, 1, 1) == "\"" && substr(v, length(v), 1) == "\"") {
        return substr(v, 2, length(v) - 2)
      }
      if (length(v) >= 2 && substr(v, 1, 1) == "'"'"'" && substr(v, length(v), 1) == "'"'"'") {
        return substr(v, 2, length(v) - 2)
      }
      return v
    }
    {
      line = $0
      sub(/[ \t]*#.*/, "", line)
      line = trim(line)
      if (line == "") {
        next
      }

      if (line ~ /^-D[ \t]+go_output=/) {
        sub(/^-D[ \t]+go_output=/, "", line)
        print dequote(trim(line))
        exit
      }

      if (line ~ /^-D[ \t]+go_output[ \t]+/) {
        sub(/^-D[ \t]+go_output[ \t]+/, "", line)
        print dequote(trim(line))
        exit
      }

      if (line ~ /^--define[ \t]+go_output=/) {
        sub(/^--define[ \t]+go_output=/, "", line)
        print dequote(trim(line))
        exit
      }

      if (line ~ /^--define[ \t]+go_output[ \t]+/) {
        sub(/^--define[ \t]+go_output[ \t]+/, "", line)
        print dequote(trim(line))
        exit
      }
    }
  ' "$hxml_path"
}

project_arg="$invocation_dir"
profile=""
hxml_arg=""
action="run"
ci=0
out_arg=""
binary_arg=""
haxe_bin="${HAXE_BIN:-haxe}"
go_bin="${GO_BIN:-go}"
extra_defines=()
extra_go_args=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --project)
      [[ $# -ge 2 ]] || fail "--project requires a value"
      project_arg="$2"
      shift 2
      ;;
    --profile)
      [[ $# -ge 2 ]] || fail "--profile requires a value"
      profile="$2"
      shift 2
      ;;
    --hxml)
      [[ $# -ge 2 ]] || fail "--hxml requires a value"
      hxml_arg="$2"
      shift 2
      ;;
    --action)
      [[ $# -ge 2 ]] || fail "--action requires a value"
      action="$2"
      shift 2
      ;;
    --ci)
      ci=1
      shift
      ;;
    --out)
      [[ $# -ge 2 ]] || fail "--out requires a value"
      out_arg="$2"
      shift 2
      ;;
    --binary)
      [[ $# -ge 2 ]] || fail "--binary requires a value"
      binary_arg="$2"
      shift 2
      ;;
    --haxe-bin)
      [[ $# -ge 2 ]] || fail "--haxe-bin requires a value"
      haxe_bin="$2"
      shift 2
      ;;
    --go-bin)
      [[ $# -ge 2 ]] || fail "--go-bin requires a value"
      go_bin="$2"
      shift 2
      ;;
    --define)
      [[ $# -ge 2 ]] || fail "--define requires a value"
      extra_defines+=("$2")
      shift 2
      ;;
    --)
      shift
      extra_go_args=("$@")
      break
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      fail "unknown argument: $1"
      ;;
  esac
done

case "$action" in
  compile|run|build|test|vet|fmt) ;;
  *) fail "invalid --action '$action' (expected: compile, run, build, test, vet, or fmt)" ;;
esac

if [[ -n "$profile" ]]; then
  case "$profile" in
    portable|gopher|metal) ;;
    *) fail "invalid --profile '$profile' (expected: portable, gopher, or metal)" ;;
  esac
fi

project_abs="$(resolve_path_from_base "$project_arg" "$invocation_dir")"
project_dir="$(normalize_existing_dir "$project_abs")"

if [[ -n "$hxml_arg" && ( -n "$profile" || "$ci" -eq 1 ) ]]; then
  fail "--hxml cannot be combined with --profile/--ci"
fi

selected_hxml_arg=""
selected_hxml_abs=""

if [[ -n "$hxml_arg" ]]; then
  selected_hxml_abs="$(resolve_path_from_base "$hxml_arg" "$project_dir")"
  [[ -f "$selected_hxml_abs" ]] || fail "hxml not found: $(display_path "$selected_hxml_abs")"
  if [[ "$selected_hxml_abs" == "$project_dir/"* ]]; then
    selected_hxml_arg="${selected_hxml_abs#"$project_dir/"}"
  else
    selected_hxml_arg="$selected_hxml_abs"
  fi
else
  declare -a candidates=()
  if [[ "$ci" -eq 1 ]]; then
    if [[ -n "$profile" ]]; then
      candidates+=("compile.${profile}.ci.hxml")
      candidates+=("compile.${profile}.hxml")
    else
      candidates+=(
        "compile.ci.hxml"
        "compile.portable.ci.hxml"
        "compile.gopher.ci.hxml"
        "compile.metal.ci.hxml"
      )
    fi
    candidates+=("compile.hxml")
  else
    if [[ -n "$profile" ]]; then
      candidates+=("compile.${profile}.hxml")
    fi
    candidates+=("compile.hxml")
  fi

  if [[ -z "$profile" ]]; then
    candidates+=("compile.portable.hxml" "compile.gopher.hxml" "compile.metal.hxml")
  fi

  for candidate in "${candidates[@]}"; do
    if [[ -f "$project_dir/$candidate" ]]; then
      selected_hxml_arg="$candidate"
      selected_hxml_abs="$project_dir/$candidate"
      break
    fi
  done

  if [[ -z "$selected_hxml_arg" ]]; then
    available="$(cd "$project_dir" && ls compile*.hxml 2>/dev/null | tr '\n' ' ' || true)"
    fail "no matching hxml in $(display_path "$project_dir") (tried: ${candidates[*]}). Available: ${available:-<none>}"
  fi
fi

if ! command -v "$haxe_bin" >/dev/null 2>&1; then
  fail "haxe binary not found: $haxe_bin"
fi

haxe_args=("$selected_hxml_arg")
if [[ -n "$out_arg" ]]; then
  haxe_args+=("-D" "go_output=$out_arg")
fi
if [[ "${#extra_defines[@]}" -gt 0 ]]; then
  for define_item in "${extra_defines[@]}"; do
    haxe_args+=("-D" "$define_item")
  done
fi

echo "[hx-go] project=$(display_path "$project_dir") profile=${profile:-auto} ci=$ci action=$action"
echo "[hx-go] hxml=$selected_hxml_arg"
(cd "$project_dir" && "$haxe_bin" "${haxe_args[@]}")

go_output_rel="$out_arg"
if [[ -z "$go_output_rel" ]]; then
  go_output_rel="$(extract_go_output "$selected_hxml_abs" || true)"
fi
[[ -n "$go_output_rel" ]] || fail "missing '-D go_output=...' in $(display_path "$selected_hxml_abs") (or pass --out)"

go_output_abs="$(resolve_path_from_base "$go_output_rel" "$project_dir")"
if [[ ! -d "$go_output_abs" ]]; then
  fail "go output directory not found after Haxe compile: $(display_path "$go_output_abs")"
fi

if [[ "$action" == "compile" ]]; then
  echo "[hx-go] compile complete: $(display_path "$go_output_abs")"
  exit 0
fi

if ! command -v "$go_bin" >/dev/null 2>&1; then
  fail "go binary not found: $go_bin"
fi

echo "[hx-go] go output=$(display_path "$go_output_abs")"

case "$action" in
  run)
    if (( ${#extra_go_args[@]} > 0 )); then
      (cd "$go_output_abs" && "$go_bin" run . "${extra_go_args[@]}")
    else
      (cd "$go_output_abs" && "$go_bin" run .)
    fi
    ;;
  test)
    if (( ${#extra_go_args[@]} > 0 )); then
      (cd "$go_output_abs" && "$go_bin" test ./... "${extra_go_args[@]}")
    else
      (cd "$go_output_abs" && "$go_bin" test ./...)
    fi
    ;;
  vet)
    if (( ${#extra_go_args[@]} > 0 )); then
      (cd "$go_output_abs" && "$go_bin" vet ./... "${extra_go_args[@]}")
    else
      (cd "$go_output_abs" && "$go_bin" vet ./...)
    fi
    ;;
  fmt)
    if (( ${#extra_go_args[@]} > 0 )); then
      (cd "$go_output_abs" && "$go_bin" fmt ./... "${extra_go_args[@]}")
    else
      (cd "$go_output_abs" && "$go_bin" fmt ./...)
    fi
    ;;
  build)
    if [[ -n "$binary_arg" ]]; then
      binary_abs="$(resolve_path_from_base "$binary_arg" "$project_dir")"
      mkdir -p "$(dirname "$binary_abs")"
      if (( ${#extra_go_args[@]} > 0 )); then
        (cd "$go_output_abs" && "$go_bin" build "${extra_go_args[@]}" -o "$binary_abs" .)
      else
        (cd "$go_output_abs" && "$go_bin" build -o "$binary_abs" .)
      fi
      echo "[hx-go] built $(display_path "$binary_abs")"
    else
      if (( ${#extra_go_args[@]} > 0 )); then
        (cd "$go_output_abs" && "$go_bin" build "${extra_go_args[@]}" .)
      else
        (cd "$go_output_abs" && "$go_bin" build .)
      fi
    fi
    ;;
esac
