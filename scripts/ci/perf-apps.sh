#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
invocation_dir="$(pwd)"
cd "$root_dir"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/ci/perf-apps.sh [options]

Options:
  --update-baseline         Regenerate scripts/ci/perf/app-profile-baseline.json from current ratios.
  --keep-work               Keep build work directory under .cache/perf-apps/work.
  -h, --help                Show this help.

Environment:
  HAXE_BIN                  Haxe binary (default: haxe)
  GO_BIN                    Go binary (default: go)
  GO_APP_PERF_CACHE_DIR     Cache/output root (default: .cache/perf-apps)
  GO_APP_PERF_BASELINE_FILE Baseline JSON path (default: scripts/ci/perf/app-profile-baseline.json)
  GO_APP_PERF_WORKLOAD_RUNS Scripted run count for latency distribution (default: 20)
  GO_APP_PERF_STARTUP_ITERS Startup loop count for help-mode startup timing (default: 50)
  GO_APP_PERF_BENCH_TIME    go test bench time (default: 200ms)
  GO_APP_PERF_BENCH_COUNT   Optional go test bench count (default: unset)
  GO_APP_PERF_THROUGHPUT_WARN_PCT  Warn if throughput ratio regresses beyond this pct (default: 12)
  GO_APP_PERF_LATENCY_WARN_PCT     Warn if latency ratio regresses beyond this pct (default: 12)
  GO_APP_PERF_ALLOC_WARN_PCT       Warn if alloc ratio regresses beyond this pct (default: 15)
  GO_APP_PERF_MEMORY_WARN_PCT      Warn if rss ratio regresses beyond this pct (default: 12)
  GO_APP_PERF_STARTUP_WARN_PCT     Warn if startup ratio regresses beyond this pct (default: 15)
  GO_APP_PERF_SIZE_WARN_PCT        Warn if size ratio regresses beyond this pct (default: 8)
  GO_APP_PERF_ENFORCE_METAL_BUDGET Fail on metal hard-budget regressions (default: 0/off)
  GO_APP_PERF_METAL_THROUGHPUT_FAIL_PCT  Hard fail if metal throughput ratio drops beyond this pct (default: 25)
  GO_APP_PERF_METAL_LATENCY_FAIL_PCT     Hard fail if metal latency ratio rises beyond this pct (default: 25)
  GO_APP_PERF_METAL_ALLOC_FAIL_PCT       Hard fail if metal alloc ratios rise beyond this pct (default: 30)
  GO_APP_PERF_METAL_MEMORY_FAIL_PCT      Hard fail if metal rss ratio rises beyond this pct (default: 20)
  GO_APP_PERF_METAL_STARTUP_FAIL_PCT     Hard fail if metal startup ratio rises beyond this pct (default: 30)
  GO_APP_PERF_METAL_SIZE_FAIL_PCT        Hard fail if metal size ratios rise beyond this pct (default: 15)
USAGE
}

log() {
  printf '[app-perf] %s\n' "$*"
}

fail() {
  printf '[app-perf] error: %s\n' "$*" >&2
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

is_truthy() {
  local value="${1:-}"
  case "$value" in
    1|true|TRUE|yes|YES|on|ON)
      return 0
      ;;
    *)
      return 1
      ;;
  esac
}

filesize_bytes() {
  local file="$1"
  if stat -f%z "$file" >/dev/null 2>&1; then
    stat -f%z "$file"
  else
    stat -c%s "$file"
  fi
}

stripped_size_bytes() {
  local file="$1"
  local tmp="${file}.app-perf-strip.tmp"
  cp "$file" "$tmp"
  if strip -x "$tmp" >/dev/null 2>&1; then
    :
  elif strip --strip-unneeded "$tmp" >/dev/null 2>&1; then
    :
  elif strip "$tmp" >/dev/null 2>&1; then
    :
  fi
  local out
  out="$(filesize_bytes "$tmp")"
  rm -f "$tmp"
  printf '%s\n' "$out"
}

go_quote() {
  local value="$1"
  value="${value//\\/\\\\}"
  value="${value//\"/\\\"}"
  printf '"%s"' "$value"
}

measure_avg_ms() {
  local iterations="$1"
  shift
  python3 - "$iterations" "$@" <<'PY'
import subprocess
import sys
import time

iters = int(sys.argv[1])
cmd = sys.argv[2:]
if iters <= 0:
    sys.stderr.write("iterations must be > 0\n")
    sys.exit(2)
if not cmd:
    sys.stderr.write("missing command\n")
    sys.exit(2)

values = []
for _ in range(iters):
    start = time.perf_counter()
    proc = subprocess.run(cmd, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    if proc.returncode != 0:
        sys.stderr.write(f"command failed for startup timing: {' '.join(cmd)}\n")
        sys.exit(2)
    values.append((time.perf_counter() - start) * 1000.0)

print(f"{sum(values) / len(values):.6f}")
PY
}

measure_workload_stats() {
  local workload_key="$1"
  local runs="$2"
  shift 2
  python3 - "$workload_key" "$runs" "$@" <<'PY'
import math
import subprocess
import sys
import time

workload_key = sys.argv[1]
runs = int(sys.argv[2])
cmd = sys.argv[3:]
if runs <= 0:
    sys.stderr.write("workload runs must be > 0\n")
    sys.exit(2)
if not cmd:
    sys.stderr.write("missing command\n")
    sys.exit(2)

def percentile(values, p):
    ordered = sorted(values)
    if not ordered:
        return 0.0
    if len(ordered) == 1:
        return ordered[0]
    rank = (len(ordered) - 1) * p
    lo = int(math.floor(rank))
    hi = int(math.ceil(rank))
    if lo == hi:
        return ordered[lo]
    return ordered[lo] + (ordered[hi] - ordered[lo]) * (rank - lo)

latencies = []
first_output = None
for _ in range(runs):
    start = time.perf_counter()
    proc = subprocess.run(cmd, capture_output=True, text=True)
    elapsed_ms = (time.perf_counter() - start) * 1000.0
    if proc.returncode != 0:
        sys.stderr.write(f"workload command failed: {' '.join(cmd)}\n")
        sys.stderr.write(proc.stderr)
        sys.exit(2)
    latencies.append(elapsed_ms)
    if first_output is None:
        first_output = proc.stdout

count = None
prefix = workload_key + "="
for line in (first_output or "").splitlines():
    if line.startswith(prefix):
        raw = line[len(prefix):].strip()
        try:
            count = float(raw)
        except ValueError:
            sys.stderr.write(f"failed to parse numeric workload count for key {workload_key}: {raw}\n")
            sys.exit(2)
        break

if count is None:
    sys.stderr.write(f"workload output missing key: {workload_key}\n")
    sys.stderr.write(first_output or "")
    sys.exit(2)

avg_ms = sum(latencies) / len(latencies)
p95_ms = percentile(latencies, 0.95)
p99_ms = percentile(latencies, 0.99)
throughput = 0.0
if avg_ms > 0:
    throughput = count * 1000.0 / avg_ms

print(f"{avg_ms:.6f}\t{p95_ms:.6f}\t{p99_ms:.6f}\t{int(round(count))}\t{throughput:.6f}")
PY
}

measure_max_rss_kb() {
  local rss_log="$1"
  shift
  if [[ "$time_mode" == "bsd" ]]; then
    /usr/bin/time -l "$@" >/dev/null 2>"$rss_log"
  else
    /usr/bin/time -v "$@" >/dev/null 2>"$rss_log"
  fi

  local rss
  rss="$(awk '
BEGIN { IGNORECASE=1 }
/maximum resident set size/ {
  for (i = 1; i <= NF; i++) {
    token = $i
    gsub(/[^0-9]/, "", token)
    if (token != "") {
      print token
      exit
    }
  }
}
' "$rss_log")"
  if [[ -z "${rss:-}" ]]; then
    fail "failed to parse max RSS from $(display_path "$rss_log")"
  fi
  if [[ "$time_mode" == "bsd" ]]; then
    rss="$(awk -v value="$rss" 'BEGIN { printf "%d\n", int((value + 1023) / 1024) }')"
  fi
  printf '%s\n' "$rss"
}

app_interface_type() {
  local app="$1"
  case "$app" in
    pulseforge)
      printf 'app.runtime.PulseRuntime\n'
      ;;
    fluxproxy)
      printf 'app.runtime.FluxRuntime\n'
      ;;
    *)
      fail "unknown app: $app"
      ;;
  esac
}

app_workload_key() {
  local app="$1"
  case "$app" in
    pulseforge)
      printf 'ingest.accepted\n'
      ;;
    fluxproxy)
      printf 'proxy.responses\n'
      ;;
    *)
      fail "unknown app: $app"
      ;;
  esac
}

compile_haxe_lane() {
  local app="$1"
  local profile="$2"
  local variant="$3"
  local out_dir="$4"
  local iface
  iface="$(app_interface_type "$app")"

  mkdir -p "$out_dir"
  (
    cd "$root_dir/examples/$app"
    "$haxe_bin" \
      -cp . \
      -cp ../../src \
      -cp ../../vendor/reflaxe/src \
      --macro reflaxe.go.CompilerBootstrap.Start\(\) \
      --macro reflaxe.go.CompilerInit.Start\(\) \
      -D "go_output=$out_dir" \
      -D "go_module=perf_${app}_${profile}_${variant}" \
      -D "reflaxe_go_profile=$profile" \
      -D reflaxe_go_strict_examples \
      -D "reflaxe_go_auto_empty_ctor_interfaces=$iface" \
      -D "${app}_profile_${profile}" \
      -D "${app}_variant_${variant}" \
      -main Main >/dev/null
  )
}

write_perf_bench_test() {
  local target_file="$1"
  shift
  local args_literal='"app"'
  local arg
  for arg in "$@"; do
    args_literal="${args_literal}, $(go_quote "$arg")"
  done

  cat > "$target_file" <<EOF
package main

import (
  "os"
  "testing"
)

func BenchmarkPerfScripted(b *testing.B) {
  oldArgs := os.Args
  defer func() { os.Args = oldArgs }()

  devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
  if err != nil {
    b.Fatalf("open devnull: %v", err)
  }
  defer devNull.Close()

  oldStdout := os.Stdout
  oldStderr := os.Stderr
  defer func() {
    os.Stdout = oldStdout
    os.Stderr = oldStderr
  }()
  os.Stdout = devNull
  os.Stderr = devNull

  b.ReportAllocs()
  for i := 0; i < b.N; i++ {
    os.Args = []string{$args_literal}
    main()
  }
}
EOF
}

bench_alloc_metrics() {
  local module_dir="$1"
  local bench_log="$2"

  local -a bench_cmd=("$go_bin" test -run '^$' -bench '^BenchmarkPerfScripted$' -benchmem)
  if [[ -n "${bench_time:-}" ]]; then
    bench_cmd+=("-benchtime=$bench_time")
  fi
  if [[ -n "${bench_count:-}" ]]; then
    bench_cmd+=("-count=$bench_count")
  fi

  (
    cd "$module_dir"
    "${bench_cmd[@]}" >"$bench_log" 2>&1
  )

  python3 - "$bench_log" <<'PY'
import re
import sys

log_path = sys.argv[1]
text = open(log_path, "r", encoding="utf-8").read()
match = re.search(
    r'BenchmarkPerfScripted(?:-\d+)?\s+\d+\s+[\d.]+\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op',
    text,
)
if not match:
    sys.stderr.write(f"failed to parse BenchmarkPerfScripted output from {log_path}\n")
    sys.exit(2)
bytes_per_op = match.group(1)
allocs_per_op = match.group(2)
print(f"{bytes_per_op}\t{allocs_per_op}")
PY
}

collect_lane_metrics() {
  local app="$1"
  local kind="$2"
  local profile="$3"
  local variant="$4"
  local module_dir="$5"
  local bin_path="$6"
  local workload_key="$7"
  shift 7
  local workload_args=("$@")

  local lane_id="${app}_${kind}_${profile}_${variant}"
  local lane_dir="$work_dir/lanes/$lane_id"
  mkdir -p "$lane_dir"

  log "metrics lane=$lane_id"

  local startup_ms
  startup_ms="$(measure_avg_ms "$startup_iters" "$bin_path" help)"

  local workload_stats
  workload_stats="$(measure_workload_stats "$workload_key" "$workload_runs" "$bin_path" "${workload_args[@]}")"
  local latency_avg_ms latency_p95_ms latency_p99_ms workload_count throughput_ops_per_sec
  IFS=$'\t' read -r latency_avg_ms latency_p95_ms latency_p99_ms workload_count throughput_ops_per_sec <<<"$workload_stats"

  local rss_max_kb
  rss_max_kb="$(measure_max_rss_kb "$lane_dir/rss.time" "$bin_path" "${workload_args[@]}")"

  write_perf_bench_test "$module_dir/perf_bench_test.go" "${workload_args[@]}"
  local alloc_metrics
  alloc_metrics="$(bench_alloc_metrics "$module_dir" "$lane_dir/bench.log")"
  local alloc_bytes_per_op allocs_per_op
  IFS=$'\t' read -r alloc_bytes_per_op allocs_per_op <<<"$alloc_metrics"

  local binary_bytes stripped_bytes
  binary_bytes="$(filesize_bytes "$bin_path")"
  stripped_bytes="$(stripped_size_bytes "$bin_path")"

  printf "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n" \
    "$lane_id" "$app" "$kind" "$profile" "$variant" \
    "$workload_count" "$throughput_ops_per_sec" \
    "$latency_avg_ms" "$latency_p95_ms" "$latency_p99_ms" \
    "$alloc_bytes_per_op" "$allocs_per_op" \
    "$rss_max_kb" "$startup_ms" "$binary_bytes" "$stripped_bytes" >> "$metrics_tsv"
}

update_baseline=0
keep_work=0
while [[ $# -gt 0 ]]; do
  case "$1" in
    --update-baseline)
      update_baseline=1
      shift
      ;;
    --keep-work)
      keep_work=1
      shift
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

haxe_bin="${HAXE_BIN:-haxe}"
go_bin="${GO_BIN:-go}"
cache_root="${GO_APP_PERF_CACHE_DIR:-$root_dir/.cache/perf-apps}"
baseline_file="${GO_APP_PERF_BASELINE_FILE:-$root_dir/scripts/ci/perf/app-profile-baseline.json}"
workload_runs="${GO_APP_PERF_WORKLOAD_RUNS:-20}"
startup_iters="${GO_APP_PERF_STARTUP_ITERS:-50}"
bench_time="${GO_APP_PERF_BENCH_TIME:-200ms}"
bench_count="${GO_APP_PERF_BENCH_COUNT:-}"
throughput_warn_pct="${GO_APP_PERF_THROUGHPUT_WARN_PCT:-12}"
latency_warn_pct="${GO_APP_PERF_LATENCY_WARN_PCT:-12}"
alloc_warn_pct="${GO_APP_PERF_ALLOC_WARN_PCT:-15}"
memory_warn_pct="${GO_APP_PERF_MEMORY_WARN_PCT:-12}"
startup_warn_pct="${GO_APP_PERF_STARTUP_WARN_PCT:-15}"
size_warn_pct="${GO_APP_PERF_SIZE_WARN_PCT:-8}"
enforce_metal_budget="${GO_APP_PERF_ENFORCE_METAL_BUDGET:-0}"
metal_throughput_fail_pct="${GO_APP_PERF_METAL_THROUGHPUT_FAIL_PCT:-25}"
metal_latency_fail_pct="${GO_APP_PERF_METAL_LATENCY_FAIL_PCT:-25}"
metal_alloc_fail_pct="${GO_APP_PERF_METAL_ALLOC_FAIL_PCT:-30}"
metal_memory_fail_pct="${GO_APP_PERF_METAL_MEMORY_FAIL_PCT:-20}"
metal_startup_fail_pct="${GO_APP_PERF_METAL_STARTUP_FAIL_PCT:-30}"
metal_size_fail_pct="${GO_APP_PERF_METAL_SIZE_FAIL_PCT:-15}"
baseline_display="$(display_path "$baseline_file")"

require_command "$haxe_bin"
require_command "$go_bin"
require_command python3

if [[ -x /usr/bin/time ]]; then
  if /usr/bin/time -l true >/dev/null 2>/dev/null; then
    time_mode="bsd"
  elif /usr/bin/time -v true >/dev/null 2>/dev/null; then
    time_mode="gnu"
  else
    fail "unsupported /usr/bin/time variant (requires -l or -v)"
  fi
else
  fail "required timing command not found: /usr/bin/time"
fi

work_dir="$cache_root/work"
results_dir="$cache_root/results"
metrics_tsv="$results_dir/raw_metrics.tsv"
current_json="$results_dir/current.json"
comparison_json="$results_dir/comparison.json"
summary_md="$results_dir/summary.md"
warnings_txt="$results_dir/warnings.txt"
hard_failures_txt="$results_dir/hard_failures.txt"

cleanup() {
  local original_exit="${1:-0}"
  if [[ "$keep_work" -eq 1 ]] || is_truthy "${KEEP_ARTIFACTS:-0}"; then
    log "keeping work dir: $(display_path "$work_dir")"
    return "$original_exit"
  fi
  rm -rf "$work_dir"
  return "$original_exit"
}
trap 'cleanup $?' EXIT

rm -rf "$work_dir"
mkdir -p "$work_dir" "$results_dir" "$(dirname "$baseline_file")"

printf "id\tapp\tkind\tprofile\tvariant\tworkload_count\tthroughput_ops_per_sec\tlatency_avg_ms\tlatency_p95_ms\tlatency_p99_ms\talloc_bytes_per_op\tallocs_per_op\trss_max_kb\tstartup_avg_ms\tbinary_bytes\tstripped_bytes\n" > "$metrics_tsv"

declare -a apps=(pulseforge fluxproxy)
declare -a profiles=(portable metal)
declare -a variants=(core go_native)

for app in "${apps[@]}"; do
  workload_key="$(app_workload_key "$app")"

  for profile in "${profiles[@]}"; do
    for variant in "${variants[@]}"; do
      lane_dir="$work_dir/build/$app/haxe/$profile/$variant"
      out_dir="$lane_dir/out"
      bin_path="$lane_dir/${app}_haxe_${profile}_${variant}"
      mkdir -p "$lane_dir"

      log "build app=$app kind=haxe profile=$profile variant=$variant"
      compile_haxe_lane "$app" "$profile" "$variant" "$out_dir"
      (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

      collect_lane_metrics "$app" "haxe" "$profile" "$variant" "$out_dir" "$bin_path" "$workload_key" --scripted
    done
  done

  for variant in "${variants[@]}"; do
    lane_dir="$work_dir/build/$app/pure_go/$variant"
    module_dir="$lane_dir/module"
    bin_path="$lane_dir/${app}_pure_go_${variant}"
    mkdir -p "$module_dir"

    log "build app=$app kind=pure_go profile=pure variant=$variant"
    cp -R "$root_dir/benchmarks/pure_go/$app/." "$module_dir/"
    (cd "$module_dir" && "$go_bin" build -o "$bin_path" .)

    collect_lane_metrics "$app" "pure_go" "pure" "$variant" "$module_dir" "$bin_path" "$workload_key" --scripted --variant "$variant"
  done
done

haxe_version="$($haxe_bin --version 2>/dev/null | tr -d '\r' | head -n 1 || true)"
go_version="$($go_bin version 2>/dev/null | tr -d '\r' | head -n 1 || true)"

GO_APP_PERF_METRICS_TSV="$metrics_tsv" \
GO_APP_PERF_CURRENT_JSON="$current_json" \
GO_APP_PERF_COMPARISON_JSON="$comparison_json" \
GO_APP_PERF_SUMMARY_MD="$summary_md" \
GO_APP_PERF_WARNINGS_TXT="$warnings_txt" \
GO_APP_PERF_HARD_FAILURES_TXT="$hard_failures_txt" \
GO_APP_PERF_BASELINE_FILE="$baseline_file" \
GO_APP_PERF_BASELINE_DISPLAY="$baseline_display" \
GO_APP_PERF_UPDATE_BASELINE="$update_baseline" \
GO_APP_PERF_HAXE_VERSION="$haxe_version" \
GO_APP_PERF_GO_VERSION="$go_version" \
GO_APP_PERF_WORKLOAD_RUNS="$workload_runs" \
GO_APP_PERF_STARTUP_ITERS="$startup_iters" \
GO_APP_PERF_BENCH_TIME="$bench_time" \
GO_APP_PERF_BENCH_COUNT="$bench_count" \
GO_APP_PERF_THROUGHPUT_WARN_PCT="$throughput_warn_pct" \
GO_APP_PERF_LATENCY_WARN_PCT="$latency_warn_pct" \
GO_APP_PERF_ALLOC_WARN_PCT="$alloc_warn_pct" \
GO_APP_PERF_MEMORY_WARN_PCT="$memory_warn_pct" \
GO_APP_PERF_STARTUP_WARN_PCT="$startup_warn_pct" \
GO_APP_PERF_SIZE_WARN_PCT="$size_warn_pct" \
GO_APP_PERF_ENFORCE_METAL_BUDGET="$enforce_metal_budget" \
GO_APP_PERF_METAL_THROUGHPUT_FAIL_PCT="$metal_throughput_fail_pct" \
GO_APP_PERF_METAL_LATENCY_FAIL_PCT="$metal_latency_fail_pct" \
GO_APP_PERF_METAL_ALLOC_FAIL_PCT="$metal_alloc_fail_pct" \
GO_APP_PERF_METAL_MEMORY_FAIL_PCT="$metal_memory_fail_pct" \
GO_APP_PERF_METAL_STARTUP_FAIL_PCT="$metal_startup_fail_pct" \
GO_APP_PERF_METAL_SIZE_FAIL_PCT="$metal_size_fail_pct" \
python3 <<'PY'
import csv
import datetime as dt
import json
import os
import re
from typing import Dict, List

metrics_path = os.environ["GO_APP_PERF_METRICS_TSV"]
current_json_path = os.environ["GO_APP_PERF_CURRENT_JSON"]
comparison_json_path = os.environ["GO_APP_PERF_COMPARISON_JSON"]
summary_md_path = os.environ["GO_APP_PERF_SUMMARY_MD"]
warnings_txt_path = os.environ["GO_APP_PERF_WARNINGS_TXT"]
hard_failures_txt_path = os.environ["GO_APP_PERF_HARD_FAILURES_TXT"]
baseline_path = os.environ["GO_APP_PERF_BASELINE_FILE"]
baseline_display = os.environ.get("GO_APP_PERF_BASELINE_DISPLAY", baseline_path)
update_baseline = os.environ.get("GO_APP_PERF_UPDATE_BASELINE", "0") == "1"
haxe_version = os.environ.get("GO_APP_PERF_HAXE_VERSION", "")
go_version = os.environ.get("GO_APP_PERF_GO_VERSION", "")
workload_runs = int(os.environ.get("GO_APP_PERF_WORKLOAD_RUNS", "20"))
startup_iters = int(os.environ.get("GO_APP_PERF_STARTUP_ITERS", "50"))
bench_time = os.environ.get("GO_APP_PERF_BENCH_TIME", "200ms")
bench_count = os.environ.get("GO_APP_PERF_BENCH_COUNT", "")
throughput_warn_pct = float(os.environ.get("GO_APP_PERF_THROUGHPUT_WARN_PCT", "12"))
latency_warn_pct = float(os.environ.get("GO_APP_PERF_LATENCY_WARN_PCT", "12"))
alloc_warn_pct = float(os.environ.get("GO_APP_PERF_ALLOC_WARN_PCT", "15"))
memory_warn_pct = float(os.environ.get("GO_APP_PERF_MEMORY_WARN_PCT", "12"))
startup_warn_pct = float(os.environ.get("GO_APP_PERF_STARTUP_WARN_PCT", "15"))
size_warn_pct = float(os.environ.get("GO_APP_PERF_SIZE_WARN_PCT", "8"))
enforce_metal_budget = bool(re.match(r"^(1|true|yes|on)$", os.environ.get("GO_APP_PERF_ENFORCE_METAL_BUDGET", "0"), re.IGNORECASE))
metal_throughput_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_THROUGHPUT_FAIL_PCT", "25"))
metal_latency_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_LATENCY_FAIL_PCT", "25"))
metal_alloc_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_ALLOC_FAIL_PCT", "30"))
metal_memory_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_MEMORY_FAIL_PCT", "20"))
metal_startup_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_STARTUP_FAIL_PCT", "30"))
metal_size_fail_pct = float(os.environ.get("GO_APP_PERF_METAL_SIZE_FAIL_PCT", "15"))

numeric_int_fields = {
    "workload_count",
    "alloc_bytes_per_op",
    "allocs_per_op",
    "rss_max_kb",
    "binary_bytes",
    "stripped_bytes",
}
numeric_float_fields = {
    "throughput_ops_per_sec",
    "latency_avg_ms",
    "latency_p95_ms",
    "latency_p99_ms",
    "startup_avg_ms",
}

with open(metrics_path, "r", encoding="utf-8") as handle:
    reader = csv.DictReader(handle, delimiter="\t")
    metrics: List[Dict[str, object]] = []
    for row in reader:
        parsed: Dict[str, object] = {}
        for key, value in row.items():
            if key in numeric_int_fields:
                parsed[key] = int(float(value))
            elif key in numeric_float_fields:
                parsed[key] = float(value)
            else:
                parsed[key] = value
        metrics.append(parsed)

if not metrics:
    raise SystemExit("no app metrics captured")

order_kind = {"haxe": 0, "pure_go": 1}
order_profile = {"portable": 0, "metal": 1, "pure": 2}
order_variant = {"core": 0, "go_native": 1}
metrics.sort(
    key=lambda item: (
        item["app"],
        order_kind.get(item["kind"], 99),
        order_profile.get(item["profile"], 99),
        order_variant.get(item["variant"], 99),
    )
)

def ratio(current: float, baseline: float) -> float:
    if baseline == 0:
        return 0.0
    return current / baseline

pure_index: Dict[str, Dict[str, object]] = {}
for metric in metrics:
    if metric["kind"] == "pure_go":
        pure_index[f"{metric['app']}::{metric['variant']}"] = metric

haxe_vs_pure: List[Dict[str, object]] = []
for metric in metrics:
    if metric["kind"] != "haxe":
        continue
    pure = pure_index.get(f"{metric['app']}::{metric['variant']}")
    if pure is None:
        continue
    haxe_vs_pure.append(
        {
            "app": metric["app"],
            "variant": metric["variant"],
            "profile": metric["profile"],
            "throughput_ratio_vs_pure": ratio(metric["throughput_ops_per_sec"], pure["throughput_ops_per_sec"]),
            "latency_avg_ratio_vs_pure": ratio(metric["latency_avg_ms"], pure["latency_avg_ms"]),
            "alloc_bytes_ratio_vs_pure": ratio(metric["alloc_bytes_per_op"], pure["alloc_bytes_per_op"]),
            "allocs_ratio_vs_pure": ratio(metric["allocs_per_op"], pure["allocs_per_op"]),
            "rss_ratio_vs_pure": ratio(metric["rss_max_kb"], pure["rss_max_kb"]),
            "startup_ratio_vs_pure": ratio(metric["startup_avg_ms"], pure["startup_avg_ms"]),
            "binary_ratio_vs_pure": ratio(metric["binary_bytes"], pure["binary_bytes"]),
            "stripped_ratio_vs_pure": ratio(metric["stripped_bytes"], pure["stripped_bytes"]),
        }
    )
haxe_vs_pure.sort(key=lambda item: (item["app"], item["variant"], item["profile"]))

warning_thresholds = {
    "throughputWarnPct": throughput_warn_pct,
    "latencyWarnPct": latency_warn_pct,
    "allocWarnPct": alloc_warn_pct,
    "memoryWarnPct": memory_warn_pct,
    "startupWarnPct": startup_warn_pct,
    "sizeWarnPct": size_warn_pct,
}
metal_fail_thresholds = {
    "throughputFailPct": metal_throughput_fail_pct,
    "latencyFailPct": metal_latency_fail_pct,
    "allocFailPct": metal_alloc_fail_pct,
    "memoryFailPct": metal_memory_fail_pct,
    "startupFailPct": metal_startup_fail_pct,
    "sizeFailPct": metal_size_fail_pct,
}

now = dt.datetime.now(dt.timezone.utc).isoformat()
current_payload = {
    "schemaVersion": 1,
    "generatedAt": now,
    "toolchain": {
        "haxe": haxe_version,
        "go": go_version,
    },
    "params": {
        "workloadRuns": workload_runs,
        "startupIterations": startup_iters,
        "benchTime": bench_time,
        "benchCount": bench_count if bench_count else None,
        "warningThresholds": warning_thresholds,
        "metalFailThresholds": metal_fail_thresholds,
        "enforceMetalBudget": enforce_metal_budget,
    },
    "metrics": metrics,
    "derived": {
        "haxeVsPure": haxe_vs_pure,
    },
}

os.makedirs(os.path.dirname(current_json_path), exist_ok=True)
with open(current_json_path, "w", encoding="utf-8") as handle:
    json.dump(current_payload, handle, indent=2)
    handle.write("\n")

baseline_payload = {
    "schemaVersion": 1,
    "generatedAt": now,
    "thresholds": warning_thresholds,
    "derivedBaseline": {
        "haxeVsPure": haxe_vs_pure,
    },
}
if update_baseline:
    os.makedirs(os.path.dirname(baseline_path), exist_ok=True)
    with open(baseline_path, "w", encoding="utf-8") as handle:
        json.dump(baseline_payload, handle, indent=2)
        handle.write("\n")

warnings: List[str] = []
hard_failures: List[str] = []
baseline_loaded = None

current_ratio_index = {
    f"{row['app']}::{row['variant']}::{row['profile']}": row
    for row in haxe_vs_pure
}

if not update_baseline:
    if not os.path.exists(baseline_path):
        warnings.append(f"baseline file not found: {baseline_display}")
    else:
        with open(baseline_path, "r", encoding="utf-8") as handle:
            baseline_loaded = json.load(handle)
        baseline_rows = baseline_loaded.get("derivedBaseline", {}).get("haxeVsPure", [])
        baseline_ratio_index = {
            f"{row['app']}::{row['variant']}::{row['profile']}": row
            for row in baseline_rows
        }

        def compare_higher_is_better(key: str, label: str, current_value: float, baseline_value: float, warn_pct: float, fail_pct: float, allow_hard: bool) -> None:
            if baseline_value <= 0:
                return
            warn_floor = baseline_value * (1 - warn_pct / 100.0)
            if current_value < warn_floor:
                drop_pct = (1 - (current_value / baseline_value)) * 100.0
                warnings.append(
                    f"{key}.{label} dropped {drop_pct:.2f}% (current={current_value:.6f}, baseline={baseline_value:.6f}, budget={warn_pct:.2f}%)"
                )
            if allow_hard:
                fail_floor = baseline_value * (1 - fail_pct / 100.0)
                if current_value < fail_floor:
                    drop_pct = (1 - (current_value / baseline_value)) * 100.0
                    hard_failures.append(
                        f"{key}.{label} dropped {drop_pct:.2f}% (current={current_value:.6f}, baseline={baseline_value:.6f}, metal budget={fail_pct:.2f}%)"
                    )

        def compare_lower_is_better(key: str, label: str, current_value: float, baseline_value: float, warn_pct: float, fail_pct: float, allow_hard: bool) -> None:
            if baseline_value <= 0:
                return
            warn_ceiling = baseline_value * (1 + warn_pct / 100.0)
            if current_value > warn_ceiling:
                rise_pct = ((current_value / baseline_value) - 1) * 100.0
                warnings.append(
                    f"{key}.{label} rose {rise_pct:.2f}% (current={current_value:.6f}, baseline={baseline_value:.6f}, budget={warn_pct:.2f}%)"
                )
            if allow_hard:
                fail_ceiling = baseline_value * (1 + fail_pct / 100.0)
                if current_value > fail_ceiling:
                    rise_pct = ((current_value / baseline_value) - 1) * 100.0
                    hard_failures.append(
                        f"{key}.{label} rose {rise_pct:.2f}% (current={current_value:.6f}, baseline={baseline_value:.6f}, metal budget={fail_pct:.2f}%)"
                    )

        for key, current_row in current_ratio_index.items():
            baseline_row = baseline_ratio_index.get(key)
            if baseline_row is None:
                warnings.append(f"baseline entry missing: {key}")
                continue

            is_metal = current_row["profile"] == "metal"
            compare_higher_is_better(
                key,
                "throughput_ratio_vs_pure",
                float(current_row["throughput_ratio_vs_pure"]),
                float(baseline_row["throughput_ratio_vs_pure"]),
                throughput_warn_pct,
                metal_throughput_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "latency_avg_ratio_vs_pure",
                float(current_row["latency_avg_ratio_vs_pure"]),
                float(baseline_row["latency_avg_ratio_vs_pure"]),
                latency_warn_pct,
                metal_latency_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "alloc_bytes_ratio_vs_pure",
                float(current_row["alloc_bytes_ratio_vs_pure"]),
                float(baseline_row["alloc_bytes_ratio_vs_pure"]),
                alloc_warn_pct,
                metal_alloc_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "allocs_ratio_vs_pure",
                float(current_row["allocs_ratio_vs_pure"]),
                float(baseline_row["allocs_ratio_vs_pure"]),
                alloc_warn_pct,
                metal_alloc_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "rss_ratio_vs_pure",
                float(current_row["rss_ratio_vs_pure"]),
                float(baseline_row["rss_ratio_vs_pure"]),
                memory_warn_pct,
                metal_memory_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "startup_ratio_vs_pure",
                float(current_row["startup_ratio_vs_pure"]),
                float(baseline_row["startup_ratio_vs_pure"]),
                startup_warn_pct,
                metal_startup_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "binary_ratio_vs_pure",
                float(current_row["binary_ratio_vs_pure"]),
                float(baseline_row["binary_ratio_vs_pure"]),
                size_warn_pct,
                metal_size_fail_pct,
                is_metal,
            )
            compare_lower_is_better(
                key,
                "stripped_ratio_vs_pure",
                float(current_row["stripped_ratio_vs_pure"]),
                float(baseline_row["stripped_ratio_vs_pure"]),
                size_warn_pct,
                metal_size_fail_pct,
                is_metal,
            )

def fmt_ratio(value: float) -> str:
    return f"{value:.3f}x"

lines: List[str] = []
lines.append("## Flagship App Performance Harness")
lines.append("")
lines.append(f"- Mode: `{'update-baseline' if update_baseline else 'compare'}`")
lines.append(f"- Baseline: `{baseline_display}`")
lines.append("- Lanes: `haxe/portable`, `haxe/metal`, `pure_go` across `core` and `go_native` variants")
lines.append(f"- Workload runs: `{workload_runs}`")
lines.append(f"- Startup iterations: `{startup_iters}`")
lines.append(f"- Bench time: `{bench_time}`")
if bench_count:
    lines.append(f"- Bench count: `{bench_count}`")
if haxe_version or go_version:
    lines.append(f"- Toolchain: {haxe_version or 'haxe:unknown'} | {go_version or 'go:unknown'}")
lines.append("")
lines.append("### Raw Metrics")
lines.append("| App | Kind | Profile | Variant | Throughput ops/s | Lat avg ms | Lat p95 ms | Lat p99 ms | B/op | allocs/op | RSS KB | Startup ms | Binary KB | Stripped KB |")
lines.append("| --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |")
for metric in metrics:
    lines.append(
        f"| {metric['app']} | {metric['kind']} | {metric['profile']} | {metric['variant']} | "
        f"{metric['throughput_ops_per_sec']:.2f} | {metric['latency_avg_ms']:.3f} | {metric['latency_p95_ms']:.3f} | {metric['latency_p99_ms']:.3f} | "
        f"{metric['alloc_bytes_per_op']} | {metric['allocs_per_op']} | {metric['rss_max_kb']} | {metric['startup_avg_ms']:.3f} | "
        f"{metric['binary_bytes'] / 1024.0:.1f} | {metric['stripped_bytes'] / 1024.0:.1f} |"
    )
lines.append("")

lines.append("### Haxe vs Pure-Go Ratios (same app + variant)")
lines.append("| App | Variant | Profile | Throughput | Latency | B/op | allocs/op | RSS | Startup | Binary | Stripped |")
lines.append("| --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |")
for row in sorted(haxe_vs_pure, key=lambda item: (item["app"], item["variant"], item["profile"])):
    lines.append(
        f"| {row['app']} | {row['variant']} | {row['profile']} | "
        f"{fmt_ratio(row['throughput_ratio_vs_pure'])} | {fmt_ratio(row['latency_avg_ratio_vs_pure'])} | "
        f"{fmt_ratio(row['alloc_bytes_ratio_vs_pure'])} | {fmt_ratio(row['allocs_ratio_vs_pure'])} | "
        f"{fmt_ratio(row['rss_ratio_vs_pure'])} | {fmt_ratio(row['startup_ratio_vs_pure'])} | "
        f"{fmt_ratio(row['binary_ratio_vs_pure'])} | {fmt_ratio(row['stripped_ratio_vs_pure'])} |"
    )
lines.append("")

lines.append("### Budget Warnings")
if warnings:
    for warning in warnings:
        lines.append(f"- {warning}")
else:
    lines.append("- none")
lines.append("")

lines.append("### Metal Hard-Fail Candidates")
if hard_failures:
    for failure in hard_failures:
        lines.append(f"- {failure}")
else:
    lines.append("- none")
lines.append("")

with open(summary_md_path, "w", encoding="utf-8") as handle:
    handle.write("\n".join(lines) + "\n")

comparison_payload = {
    "schemaVersion": 1,
    "generatedAt": now,
    "mode": "update-baseline" if update_baseline else "compare",
    "baselinePath": baseline_display,
    "baselineAvailable": update_baseline or baseline_loaded is not None,
    "enforceMetalBudget": enforce_metal_budget,
    "warningCount": len(warnings),
    "metalWarningCount": len([warning for warning in warnings if ("::metal::" in warning or "::metal." in warning)]),
    "hardFailureCount": len(hard_failures),
    "warnings": warnings,
    "hardFailures": hard_failures,
}
with open(comparison_json_path, "w", encoding="utf-8") as handle:
    json.dump(comparison_payload, handle, indent=2)
    handle.write("\n")

with open(warnings_txt_path, "w", encoding="utf-8") as handle:
    if warnings:
        handle.write("\n".join(warnings) + "\n")

with open(hard_failures_txt_path, "w", encoding="utf-8") as handle:
    if hard_failures:
        handle.write("\n".join(hard_failures) + "\n")
PY

warning_count=0
baseline_warning_count=0
hard_failure_count=0
if [[ -s "$warnings_txt" ]]; then
  while IFS= read -r warning; do
    [[ -n "$warning" ]] || continue
    warning_count=$((warning_count + 1))
    if [[ "$warning" == baseline* ]]; then
      baseline_warning_count=$((baseline_warning_count + 1))
    fi
    echo "::warning::[app-perf] $warning"
  done < "$warnings_txt"
fi

if [[ -s "$hard_failures_txt" ]]; then
  while IFS= read -r hard_failure; do
    [[ -n "$hard_failure" ]] || continue
    hard_failure_count=$((hard_failure_count + 1))
    if is_truthy "$enforce_metal_budget"; then
      echo "::error::[app-perf] $hard_failure"
    else
      echo "::warning::[app-perf][metal-hard-candidate] $hard_failure"
    fi
  done < "$hard_failures_txt"
fi

if [[ -n "${GITHUB_STEP_SUMMARY:-}" && -f "$summary_md" ]]; then
  {
    echo ""
    cat "$summary_md"
    echo ""
  } >> "$GITHUB_STEP_SUMMARY"
fi

if is_truthy "$enforce_metal_budget"; then
  if [[ "$hard_failure_count" -gt 0 || "$baseline_warning_count" -gt 0 ]]; then
    echo "::error::[app-perf] metal budget enforcement failed (hard_failures=$hard_failure_count baseline_warnings=$baseline_warning_count)"
    exit 1
  fi
fi

log "done"
log "metrics: $(display_path "$current_json")"
log "comparison: $(display_path "$comparison_json")"
log "summary: $(display_path "$summary_md")"
log "raw metrics: $(display_path "$metrics_tsv")"
if [[ "$update_baseline" -eq 1 ]]; then
  log "baseline updated: $(display_path "$baseline_file")"
fi
