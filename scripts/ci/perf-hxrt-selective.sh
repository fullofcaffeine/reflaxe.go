#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
invocation_dir="$(pwd)"
cd "$root_dir"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/ci/perf-hxrt-selective.sh [options]

Options:
  --update-baseline   Regenerate scripts/ci/perf/hxrt-selective-baseline.json from current run.
  --keep-work         Keep generated work tree under .cache/perf-hxrt-selective/work.
  -h, --help          Show this help.

Environment:
  HAXE_BIN                        Haxe binary (default: haxe)
  GO_BIN                          Go binary (default: go)
  GO_HXRT_SLICE_CACHE_DIR         Cache/output root (default: .cache/perf-hxrt-selective)
  GO_HXRT_SLICE_BASELINE_FILE     Baseline JSON path (default: scripts/ci/perf/hxrt-selective-baseline.json)
  GO_HXRT_SLICE_ENFORCE           Fail if selective metrics regress over full metrics (default: 0/off)
  GO_HXRT_SLICE_MAX_SOURCE_PCT    Max selective runtime source bytes increase over full (default: 5)
  GO_HXRT_SLICE_MAX_BINARY_PCT    Max selective binary size increase over full (default: 10)
USAGE
}

log() {
  printf '[hxrt-slice] %s\n' "$*"
}

fail() {
  printf '[hxrt-slice] error: %s\n' "$*" >&2
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

sum_runtime_source_bytes() {
  local runtime_dir="$1"
  local sum=0
  while IFS= read -r file; do
    local size
    size="$(filesize_bytes "$file")"
    sum=$((sum + size))
  done < <(find "$runtime_dir" -maxdepth 1 -type f | sort)
  printf '%s\n' "$sum"
}

write_case_main() {
  local case_name="$1"
  local target="$2"
  mkdir -p "$target"
  case "$case_name" in
    core)
      cat > "$target/Main.hx" <<'EOF'
class Main {
  static function main():Void {
    Sys.println("ok");
  }
}
EOF
      ;;
    json)
      cat > "$target/Main.hx" <<'EOF'
class Main {
  static function main():Void {
    var value:Dynamic = haxe.Json.parse('{"n":1}');
    Sys.println(Std.string(value != null));
  }
}
EOF
      ;;
    serialization)
      cat > "$target/Main.hx" <<'EOF'
class SerializationValue {
  private var amount:Float;

  public function new(amount:Float) {
    this.amount = amount;
  }
}

class Main {
  static function main():Void {
    var encoded = haxe.Serializer.run(new SerializationValue(3));
    var decoded:SerializationValue = cast haxe.Unserializer.run(encoded);
    Sys.println(decoded != null);
  }
}
EOF
      ;;
    sys_process)
      cat > "$target/Main.hx" <<'EOF'
class Main {
  static function main():Void {
    var cwd = Sys.getCwd();
    var process = new sys.io.Process("echo", ["ok"]);
    process.close();
    Sys.println(cwd);
  }
}
EOF
      ;;
    metal_go_native)
      cat > "$target/Main.hx" <<'EOF'
import go.Chan;
import go.Go;

class Main {
  static function main():Void {
    var channel:Chan<Int> = Go.newChan();
    Go.spawn(function() channel.send(1));
    Sys.println(channel.recv());
  }
}
EOF
      ;;
    *)
      fail "unknown case: $case_name"
      ;;
  esac
}

compile_case() {
  local src_dir="$1"
  local out_dir="$2"
  local profile="$3"
  local mode="$4"

  mkdir -p "$out_dir"
  local args=(
    -cp .
    -cp "$root_dir/src"
    -cp "$root_dir/vendor/reflaxe/src"
    -cp "$root_dir/std"
    -cp "$root_dir/std/go/_std"
    --macro reflaxe.go.CompilerBootstrap.Start\(\)
    --macro reflaxe.go.CompilerInit.Start\(\)
    -D "go_output=$out_dir"
    -D "reflaxe_go_profile=$profile"
    -D go_no_build
    -D reflaxe.dont_output_metadata_id
    -D no-traces
    -D no_traces
    -main Main
  )

  if [[ "$mode" == "selective" ]]; then
    args+=(-D "reflaxe_go_hxrt_features=")
  fi

  (
    cd "$src_dir"
    "$haxe_bin" "${args[@]}" >/dev/null
  )
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
cache_root="${GO_HXRT_SLICE_CACHE_DIR:-$root_dir/.cache/perf-hxrt-selective}"
baseline_file="${GO_HXRT_SLICE_BASELINE_FILE:-$root_dir/scripts/ci/perf/hxrt-selective-baseline.json}"
enforce_budget="${GO_HXRT_SLICE_ENFORCE:-0}"
max_source_pct="${GO_HXRT_SLICE_MAX_SOURCE_PCT:-5}"
max_binary_pct="${GO_HXRT_SLICE_MAX_BINARY_PCT:-10}"

require_command "$haxe_bin"
require_command "$go_bin"
require_command node

work_dir="$cache_root/work"
results_dir="$cache_root/results"
metrics_tsv="$results_dir/raw_metrics.tsv"
current_json="$results_dir/current.json"
comparison_json="$results_dir/comparison.json"
summary_md="$results_dir/summary.md"

cleanup() {
  local exit_code="${1:-0}"
  if [[ "$keep_work" -eq 1 ]] || is_truthy "${KEEP_ARTIFACTS:-0}"; then
    log "keeping work dir: $(display_path "$work_dir")"
    return "$exit_code"
  fi
  rm -rf "$work_dir"
  return "$exit_code"
}
trap 'cleanup $?' EXIT

rm -rf "$work_dir"
mkdir -p "$work_dir" "$results_dir" "$(dirname "$baseline_file")"
printf "id\tcase\tprofile\tmode\truntime_file_count\truntime_source_bytes\tbinary_bytes\n" > "$metrics_tsv"

declare -a case_specs=(
  "core:portable"
  "json:portable"
  "serialization:portable"
  "sys_process:portable"
  "metal_go_native:metal"
)

for spec in "${case_specs[@]}"; do
  case_name="${spec%%:*}"
  profile="${spec##*:}"
  src_dir="$work_dir/src/$case_name"
  write_case_main "$case_name" "$src_dir"

  for mode in full selective; do
    log "case=$case_name profile=$profile mode=$mode"
    case_dir="$work_dir/$case_name/$profile/$mode"
    out_dir="$case_dir/out"
    bin_path="$case_dir/app"
    mkdir -p "$case_dir"

    compile_case "$src_dir" "$out_dir" "$profile" "$mode"
    (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

    runtime_dir="$out_dir/hxrt"
    runtime_files="$(find "$runtime_dir" -maxdepth 1 -type f | wc -l | tr -d ' ')"
    runtime_bytes="$(sum_runtime_source_bytes "$runtime_dir")"
    binary_bytes="$(filesize_bytes "$bin_path")"

    printf "%s\t%s\t%s\t%s\t%s\t%s\t%s\n" \
      "${case_name}_${profile}_${mode}" "$case_name" "$profile" "$mode" \
      "$runtime_files" "$runtime_bytes" "$binary_bytes" >> "$metrics_tsv"
  done
done

GO_HXRT_SLICE_METRICS="$metrics_tsv" \
GO_HXRT_SLICE_CURRENT="$current_json" \
GO_HXRT_SLICE_COMPARISON="$comparison_json" \
GO_HXRT_SLICE_SUMMARY="$summary_md" \
GO_HXRT_SLICE_BASELINE="$baseline_file" \
GO_HXRT_SLICE_UPDATE_BASELINE="$update_baseline" \
GO_HXRT_SLICE_ENFORCE="$enforce_budget" \
GO_HXRT_SLICE_MAX_SOURCE_PCT="$max_source_pct" \
GO_HXRT_SLICE_MAX_BINARY_PCT="$max_binary_pct" \
node <<'NODE'
const fs = require("fs");

const metricsPath = process.env.GO_HXRT_SLICE_METRICS;
const currentPath = process.env.GO_HXRT_SLICE_CURRENT;
const comparisonPath = process.env.GO_HXRT_SLICE_COMPARISON;
const summaryPath = process.env.GO_HXRT_SLICE_SUMMARY;
const baselinePath = process.env.GO_HXRT_SLICE_BASELINE;
const updateBaseline = process.env.GO_HXRT_SLICE_UPDATE_BASELINE === "1";
const enforce = /^(1|true|yes|on)$/i.test(process.env.GO_HXRT_SLICE_ENFORCE || "0");
const maxSourcePct = Number(process.env.GO_HXRT_SLICE_MAX_SOURCE_PCT || "5");
const maxBinaryPct = Number(process.env.GO_HXRT_SLICE_MAX_BINARY_PCT || "10");

function parseMetrics(tsvPath) {
  const raw = fs.readFileSync(tsvPath, "utf8").trim();
  const lines = raw.split(/\r?\n/);
  const header = lines.shift().split("\t");
  return lines.filter(Boolean).map((line) => {
    const cells = line.split("\t");
    const entry = {};
    header.forEach((key, i) => (entry[key] = cells[i] ?? ""));
    return {
      id: entry.id,
      case: entry.case,
      profile: entry.profile,
      mode: entry.mode,
      runtime_file_count: Number(entry.runtime_file_count),
      runtime_source_bytes: Number(entry.runtime_source_bytes),
      binary_bytes: Number(entry.binary_bytes),
    };
  });
}

function ratio(a, b) {
  if (!b) return 0;
  return a / b;
}

const metrics = parseMetrics(metricsPath);
const current = {
  generated_at: new Date().toISOString(),
  metrics,
};
fs.writeFileSync(currentPath, JSON.stringify(current, null, 2));

const groups = {};
for (const row of metrics) {
  const key = `${row.case}:${row.profile}`;
  if (!groups[key]) groups[key] = {};
  groups[key][row.mode] = row;
}

const diffs = [];
const failures = [];
for (const [key, rows] of Object.entries(groups)) {
  const full = rows.full;
  const selective = rows.selective;
  if (!full || !selective) continue;
  const [caseName, profile] = key.split(":");
  const sourceRatio = ratio(selective.runtime_source_bytes, full.runtime_source_bytes);
  const binaryRatio = ratio(selective.binary_bytes, full.binary_bytes);
  const sourceDeltaPct = (sourceRatio - 1) * 100;
  const binaryDeltaPct = (binaryRatio - 1) * 100;
  const fileDelta = selective.runtime_file_count - full.runtime_file_count;
  diffs.push({
    case: caseName,
    profile,
    full,
    selective,
    source_ratio: sourceRatio,
    binary_ratio: binaryRatio,
    source_delta_pct: sourceDeltaPct,
    binary_delta_pct: binaryDeltaPct,
    file_delta: fileDelta,
  });

  if (enforce) {
    if (fileDelta > 0) {
      failures.push(`${caseName}/${profile}: selective runtime file count increased (${full.runtime_file_count} -> ${selective.runtime_file_count})`);
    }
    if (sourceDeltaPct > maxSourcePct) {
      failures.push(`${caseName}/${profile}: selective runtime source bytes +${sourceDeltaPct.toFixed(2)}% exceeds ${maxSourcePct}%`);
    }
    if (binaryDeltaPct > maxBinaryPct) {
      failures.push(`${caseName}/${profile}: selective binary bytes +${binaryDeltaPct.toFixed(2)}% exceeds ${maxBinaryPct}%`);
    }
  }
}

let baseline = null;
if (fs.existsSync(baselinePath)) {
  baseline = JSON.parse(fs.readFileSync(baselinePath, "utf8"));
}
if (updateBaseline) {
  fs.writeFileSync(baselinePath, JSON.stringify({ generated_at: current.generated_at, diffs }, null, 2));
}

let baselineIndex = null;
if (baseline && Array.isArray(baseline.diffs)) {
  baselineIndex = Object.create(null);
  for (const row of baseline.diffs) {
    baselineIndex[`${row.case}:${row.profile}`] = row;
  }
  for (const row of diffs) {
    const baselineRow = baselineIndex[`${row.case}:${row.profile}`];
    if (!baselineRow) continue;
    row.baseline_source_delta_pct = baselineRow.source_delta_pct;
    row.baseline_binary_delta_pct = baselineRow.binary_delta_pct;
    row.baseline_file_delta = baselineRow.file_delta;
    row.source_delta_drift_pct = row.source_delta_pct - baselineRow.source_delta_pct;
    row.binary_delta_drift_pct = row.binary_delta_pct - baselineRow.binary_delta_pct;
    row.file_delta_drift = row.file_delta - baselineRow.file_delta;
  }
}

fs.writeFileSync(comparisonPath, JSON.stringify({
  generated_at: current.generated_at,
  baseline_present: !!baseline,
  baseline_generated_at: baseline ? baseline.generated_at || null : null,
  diffs,
  failures,
}, null, 2));

const lines = [];
lines.push("# HXRT Selective Runtime Metrics");
lines.push("");
lines.push("| Case | Profile | Full files | Selective files | File delta | Source delta % | Binary delta % | Drift source pct | Drift binary pct |");
lines.push("| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |");
for (const entry of diffs) {
  const sourceDrift = Number.isFinite(entry.source_delta_drift_pct) ? entry.source_delta_drift_pct.toFixed(2) : "n/a";
  const binaryDrift = Number.isFinite(entry.binary_delta_drift_pct) ? entry.binary_delta_drift_pct.toFixed(2) : "n/a";
  lines.push(`| ${entry.case} | ${entry.profile} | ${entry.full.runtime_file_count} | ${entry.selective.runtime_file_count} | ${entry.file_delta} | ${entry.source_delta_pct.toFixed(2)} | ${entry.binary_delta_pct.toFixed(2)} | ${sourceDrift} | ${binaryDrift} |`);
}
lines.push("");
if (failures.length > 0) {
  lines.push("## Failures");
  lines.push("");
  for (const failure of failures) {
    lines.push(`- ${failure}`);
  }
}
fs.writeFileSync(summaryPath, lines.join("\n") + "\n");

if (failures.length > 0) {
  process.exit(2);
}
NODE

log "current metrics: $(display_path "$current_json")"
log "comparison report: $(display_path "$comparison_json")"
log "summary: $(display_path "$summary_md")"
if [[ "$update_baseline" -eq 1 ]]; then
  log "baseline updated: $(display_path "$baseline_file")"
fi
