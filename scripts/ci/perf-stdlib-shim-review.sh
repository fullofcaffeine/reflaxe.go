#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
invocation_dir="$(pwd)"
cd "$root_dir"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/ci/perf-stdlib-shim-review.sh [options]

Options:
  --keep-work   Keep temporary work directory under cache.
  -h, --help    Show this help.

Environment:
  HAXE_BIN                    Haxe binary (default: haxe)
  GO_BIN                      Go binary (default: go)
  SHIM_REVIEW_CACHE_DIR       Cache/output root (default: .cache/perf-stdlib-shim-review)
USAGE
}

log() {
  printf '[shim-review] %s\n' "$*"
}

fail() {
  printf '[shim-review] error: %s\n' "$*" >&2
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

require_command() {
  local cmd="$1"
  if ! command -v "$cmd" >/dev/null 2>&1; then
    fail "required command not found: $cmd"
  fi
}

keep_work=0
while [[ $# -gt 0 ]]; do
  case "$1" in
    --keep-work)
      keep_work=1
      shift
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      fail "unknown option: $1"
      ;;
  esac
done

haxe_bin="${HAXE_BIN:-haxe}"
go_bin="${GO_BIN:-go}"
cache_dir="${SHIM_REVIEW_CACHE_DIR:-$root_dir/.cache/perf-stdlib-shim-review}"
work_dir="$cache_dir/work"
report_json="$cache_dir/report.json"
report_md="$cache_dir/report.md"

require_command "$haxe_bin"
require_command "$go_bin"
require_command python3

mkdir -p "$cache_dir"
rm -rf "$work_dir"
mkdir -p "$work_dir/haxe_src" "$work_dir/direct"

cleanup() {
  if [[ "$keep_work" -eq 1 ]]; then
    log "keeping work directory: $(display_path "$work_dir")"
  else
    rm -rf "$work_dir"
  fi
}
trap cleanup EXIT

log "compiling Haxe shim benchmark case"
cat > "$work_dir/haxe_src/Main.hx" <<'HX'
class Main {
  static function main():Void {
    var payload = haxe.io.Bytes.ofString("bench-payload-0123456789abcdef");
    var encoded = haxe.crypto.Base64.encode(payload);
    var decoded = haxe.crypto.Base64.decode(encoded);
    Sys.println(decoded == null ? "nil" : "ok");
  }
}
HX

haxe_out="$work_dir/haxe_out"
"$haxe_bin" \
  -cp "$work_dir/haxe_src" \
  -cp "$root_dir/src" \
  --macro reflaxe.go.CompilerBootstrap.Start\(\) \
  --macro reflaxe.go.CompilerInit.Start\(\) \
  -D "go_output=$haxe_out" \
  -D reflaxe_go_profile=gopher \
  -D go_no_build \
  -D reflaxe.dont_output_metadata_id \
  -D no-traces \
  -D no_traces \
  -main Main >/dev/null

cat > "$haxe_out/main_test.go" <<'GO'
package main

import (
  "snapshot/hxrt"
  "testing"
)

var shimBenchBytes = haxe__io__Bytes_ofString(hxrt.StringFromLiteral("bench-payload-0123456789abcdef"))

func BenchmarkShimBase64Encode(b *testing.B) {
  b.ReportAllocs()
  for i := 0; i < b.N; i++ {
    _ = haxe__crypto__Base64_encode(shimBenchBytes)
  }
}
GO

log "running generated-shim benchmark"
shim_bench_log="$work_dir/shim.bench.txt"
(
  cd "$haxe_out"
  "$go_bin" test -run '^$' -bench '^BenchmarkShimBase64Encode$' -benchmem
) | tee "$shim_bench_log" >/dev/null

log "building direct Go benchmark case"
cat > "$work_dir/direct/go.mod" <<'GO'
module direct

go 1.22
GO

cat > "$work_dir/direct/base64_test.go" <<'GO'
package direct

import (
  "encoding/base64"
  "testing"
)

var directBenchBytes = []byte("bench-payload-0123456789abcdef")

func directBase64Encode(bytes []byte) string {
  return base64.StdEncoding.EncodeToString(bytes)
}

func BenchmarkDirectBase64Encode(b *testing.B) {
  b.ReportAllocs()
  for i := 0; i < b.N; i++ {
    _ = directBase64Encode(directBenchBytes)
  }
}
GO

log "running direct benchmark"
direct_bench_log="$work_dir/direct.bench.txt"
(
  cd "$work_dir/direct"
  "$go_bin" test -run '^$' -bench '^BenchmarkDirectBase64Encode$' -benchmem
) | tee "$direct_bench_log" >/dev/null

extract_bench_triplet() {
  local file="$1"
  awk '/^Benchmark/ { print $3 "\t" $5 "\t" $7; exit }' "$file"
}

read -r shim_ns shim_bytes shim_allocs <<<"$(extract_bench_triplet "$shim_bench_log")"
read -r direct_ns direct_bytes direct_allocs <<<"$(extract_bench_triplet "$direct_bench_log")"

if [[ -z "${shim_ns:-}" || -z "${direct_ns:-}" ]]; then
  fail "failed to parse benchmark output"
fi

func_loc() {
  local file="$1"
  local fn="$2"
  python3 - "$file" "$fn" <<'PY'
import re
import sys
from pathlib import Path

path = Path(sys.argv[1])
func_name = sys.argv[2]
lines = path.read_text().splitlines()
pattern = re.compile(rf'^func\s+(?:\([^\)]*\)\s+)?{re.escape(func_name)}\s*\(')
start = None
for i, line in enumerate(lines):
    if pattern.match(line):
        start = i
        break
if start is None:
    raise SystemExit(2)
brace_depth = 0
seen_open = False
for end in range(start, len(lines)):
    for ch in lines[end]:
        if ch == "{":
            brace_depth += 1
            seen_open = True
        elif ch == "}":
            brace_depth -= 1
            if seen_open and brace_depth == 0:
                print(end - start + 1)
                raise SystemExit(0)
raise SystemExit(3)
PY
}

shim_main="$haxe_out/main.go"
direct_file="$work_dir/direct/base64_test.go"
loc_shim_encode="$(func_loc "$shim_main" "haxe__crypto__Base64_encode")"
loc_shim_to_raw="$(func_loc "$shim_main" "hxrt_haxeBytesToRaw")"
loc_shim_from_raw="$(func_loc "$shim_main" "hxrt_rawToHaxeBytes")"
loc_direct_encode="$(func_loc "$direct_file" "directBase64Encode")"
loc_shim_path="$((loc_shim_encode + loc_shim_to_raw + loc_shim_from_raw))"

shim_overhead_ns_pct="$(awk -v shim="$shim_ns" -v direct="$direct_ns" 'BEGIN { printf "%.2f", ((shim / direct) - 1.0) * 100.0 }')"
shim_overhead_bytes="$(awk -v shim="$shim_bytes" -v direct="$direct_bytes" 'BEGIN { printf "%d", shim - direct }')"
shim_overhead_allocs="$(awk -v shim="$shim_allocs" -v direct="$direct_allocs" 'BEGIN { printf "%d", shim - direct }')"

run_date="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
goos="$("$go_bin" env GOOS)"
goarch="$("$go_bin" env GOARCH)"
cpu_model="$(sysctl -n machdep.cpu.brand_string 2>/dev/null || echo "unknown")"

cat > "$report_json" <<JSON
{
  "generatedAtUtc": "$run_date",
  "platform": {
    "goos": "$goos",
    "goarch": "$goarch",
    "cpu": "$cpu_model"
  },
  "surface": "haxe.crypto.Base64.encode",
  "shimPath": {
    "nsPerOp": $shim_ns,
    "bytesPerOp": $shim_bytes,
    "allocsPerOp": $shim_allocs,
    "codeShapeLoc": {
      "haxe__crypto__Base64_encode": $loc_shim_encode,
      "hxrt_haxeBytesToRaw": $loc_shim_to_raw,
      "hxrt_rawToHaxeBytes": $loc_shim_from_raw,
      "totalPathLoc": $loc_shim_path
    }
  },
  "directGoPath": {
    "nsPerOp": $direct_ns,
    "bytesPerOp": $direct_bytes,
    "allocsPerOp": $direct_allocs,
    "codeShapeLoc": {
      "directBase64Encode": $loc_direct_encode
    }
  },
  "delta": {
    "nsOverheadPct": $shim_overhead_ns_pct,
    "bytesOverhead": $shim_overhead_bytes,
    "allocsOverhead": $shim_overhead_allocs
  }
}
JSON

cat > "$report_md" <<MD
# Stdlib Shim Review Benchmark

- Run timestamp (UTC): \`$run_date\`
- Platform: \`$goos/$goarch\`
- CPU: \`$cpu_model\`
- Surface: \`haxe.crypto.Base64.encode\`

## Results

| Path | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| Generated shim path | $shim_ns | $shim_bytes | $shim_allocs |
| Direct Go path | $direct_ns | $direct_bytes | $direct_allocs |
| Delta | ${shim_overhead_ns_pct}% | +$shim_overhead_bytes | +$shim_overhead_allocs |

## Code shape

| Function path | LOC |
| --- | ---: |
| \`haxe__crypto__Base64_encode\` | $loc_shim_encode |
| \`hxrt_haxeBytesToRaw\` | $loc_shim_to_raw |
| \`hxrt_rawToHaxeBytes\` | $loc_shim_from_raw |
| Shim call-path total | $loc_shim_path |
| \`directBase64Encode\` | $loc_direct_encode |
MD

log "report written: $(display_path "$report_json")"
log "report written: $(display_path "$report_md")"
log "summary: shim ${shim_ns}ns/op vs direct ${direct_ns}ns/op (${shim_overhead_ns_pct}% overhead)"
