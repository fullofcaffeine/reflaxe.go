#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
invocation_dir="$(pwd)"
cd "$root_dir"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/ci/perf-go-profiles.sh [options]

Options:
  --update-baseline         Regenerate scripts/ci/perf/go-profile-baseline.json from current metrics.
  --keep-work               Keep build work directory under .cache/perf-go/work.
  -h, --help                Show this help.

Environment:
  HAXE_BIN                  Haxe binary (default: haxe)
  GO_BIN                    Go binary (default: go)
  GO_PERF_CACHE_DIR         Cache/output root (default: .cache/perf-go)
  GO_PERF_BASELINE_FILE     Baseline JSON path (default: scripts/ci/perf/go-profile-baseline.json)
  GO_PERF_SIZE_WARN_PCT     Soft warning threshold for size ratios (default: 5)
  GO_PERF_RUNTIME_WARN_PCT  Soft warning threshold for startup ratios (default: 10)
  GO_PERF_ENFORCE_METAL_BUDGET
                            Fail when metal profile exceeds budget (default: 0/off).
  GO_PERF_METAL_SIZE_FAIL_PCT
                            Hard-fail threshold for metal size ratios (default: 25).
  GO_PERF_METAL_RUNTIME_FAIL_PCT
                            Hard-fail threshold for metal startup ratios (default: 100).
  GO_PERF_HELLO_ITERS       Startup loop count for hello case (default: 300)
  GO_PERF_ARRAY_ITERS       Startup loop count for array case (default: 300)
  GO_PERF_TUI_ITERS         Startup loop count for tui case (default: 30)
USAGE
}

log() {
  printf '[go-perf] %s\n' "$*"
}

fail() {
  printf '[go-perf] error: %s\n' "$*" >&2
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

require_command() {
  local cmd="$1"
  if ! command -v "$cmd" >/dev/null 2>&1; then
    fail "required command not found: $cmd"
  fi
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
  local tmp="${file}.go-perf-strip.tmp"
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

measure_startup_ms() {
  local bin="$1"
  local iterations="$2"
  local timing_log="$3"

  ITER="$iterations" BIN="$bin" "$time_bin" -p bash -c '
    i=0
    while [ "$i" -lt "$ITER" ]; do
      "$BIN" >/dev/null 2>&1 || exit 1
      i=$((i + 1))
    done
  ' >/dev/null 2>"$timing_log"

  local real_seconds
  real_seconds="$(awk '/^real[[:space:]]+/ { print $2; exit }' "$timing_log")"
  if [[ -z "${real_seconds:-}" ]]; then
    fail "failed to parse startup timing from $(display_path "$timing_log")"
  fi

  awk -v real="$real_seconds" -v count="$iterations" 'BEGIN { printf "%.6f\n", (real * 1000.0) / count }'
}

write_haxe_hello_case() {
  local dir="$1"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<'EOF'
class Main {
  static function main():Void {
    Sys.println("hi");
  }
}
EOF
}

write_haxe_array_case() {
  local dir="$1"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<'EOF'
class Main {
  static function main():Void {
    var xs = [1, 2, 3];
    var sum = 0;
    for (x in xs) {
      sum += x;
    }
    Sys.println(sum);
  }
}
EOF
}

compile_haxe_case() {
  local src_dir="$1"
  local out_dir="$2"
  local profile="$3"

  mkdir -p "$out_dir"
  (
    cd "$src_dir"
    "$haxe_bin" \
      -cp . \
      -cp "$root_dir/src" \
      --macro reflaxe.go.CompilerBootstrap.Start\(\) \
      --macro reflaxe.go.CompilerInit.Start\(\) \
      -D "go_output=$out_dir" \
      -D "reflaxe_go_profile=$profile" \
      -D go_no_build \
      -D reflaxe.dont_output_metadata_id \
      -D no-traces \
      -D no_traces \
      -main Main >/dev/null
  )
}

write_pure_hello_module() {
  local dir="$1"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_hello

go 1.22
EOF
  cat > "$dir/main.go" <<'EOF'
package main

import "fmt"

func main() {
  fmt.Println("hi")
}
EOF
}

write_pure_array_module() {
  local dir="$1"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_array

go 1.22
EOF
  cat > "$dir/main.go" <<'EOF'
package main

import "fmt"

func main() {
  xs := []int{1, 2, 3}
  sum := 0
  for _, x := range xs {
    sum += x
  }
  fmt.Println(sum)
}
EOF
}

record_metric() {
  local id="$1"
  local case_name="$2"
  local profile="$3"
  local kind="$4"
  local bin_path="$5"
  local iterations="$6"
  local timing_log="$7"

  if [[ ! -f "$bin_path" ]]; then
    fail "binary not found: $(display_path "$bin_path")"
  fi

  local startup_ms
  startup_ms="$(measure_startup_ms "$bin_path" "$iterations" "$timing_log")"
  local bin_bytes
  bin_bytes="$(filesize_bytes "$bin_path")"
  local stripped_bytes
  stripped_bytes="$(stripped_size_bytes "$bin_path")"

  printf "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n" \
    "$id" "$case_name" "$profile" "$kind" \
    "$bin_bytes" "$stripped_bytes" "$startup_ms" "$iterations" >> "$metrics_tsv"
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
cache_root="${GO_PERF_CACHE_DIR:-$root_dir/.cache/perf-go}"
baseline_file="${GO_PERF_BASELINE_FILE:-$root_dir/scripts/ci/perf/go-profile-baseline.json}"
baseline_display="$(display_path "$baseline_file")"
size_warn_pct="${GO_PERF_SIZE_WARN_PCT:-5}"
runtime_warn_pct="${GO_PERF_RUNTIME_WARN_PCT:-10}"
enforce_metal_budget="${GO_PERF_ENFORCE_METAL_BUDGET:-0}"
metal_size_fail_pct="${GO_PERF_METAL_SIZE_FAIL_PCT:-25}"
metal_runtime_fail_pct="${GO_PERF_METAL_RUNTIME_FAIL_PCT:-100}"
hello_iters="${GO_PERF_HELLO_ITERS:-300}"
array_iters="${GO_PERF_ARRAY_ITERS:-300}"
tui_iters="${GO_PERF_TUI_ITERS:-30}"

if [[ -x /usr/bin/time ]]; then
  time_bin="/usr/bin/time"
else
  fail "required timing command not found: /usr/bin/time"
fi

require_command "$haxe_bin"
require_command "$go_bin"
require_command node

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
mkdir -p "$work_dir" "$results_dir"
mkdir -p "$(dirname "$baseline_file")"

printf "id\tcase\tprofile\tkind\tbinary_bytes\tstripped_bytes\tstartup_avg_ms\tstartup_iterations\n" > "$metrics_tsv"

log "collecting metrics (results: $(display_path "$results_dir"))"

declare -a profiles=(portable gopher metal)

hello_src="$work_dir/haxe_cases/hello"
write_haxe_hello_case "$hello_src"

for profile in "${profiles[@]}"; do
  log "hello case ($profile)"
  case_dir="$work_dir/hello/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/hello_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$hello_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "hello_haxe_${profile}" "hello" "$profile" "haxe" \
    "$bin_path" "$hello_iters" "$case_dir/startup.time"
done

log "hello pure Go baseline"
hello_pure_dir="$work_dir/hello/pure"
hello_pure_bin="$hello_pure_dir/pure_hello"
write_pure_hello_module "$hello_pure_dir"
(cd "$hello_pure_dir" && "$go_bin" build -o "$hello_pure_bin" .)
record_metric "hello_pure_go" "hello" "pure" "pure_go" \
  "$hello_pure_bin" "$hello_iters" "$hello_pure_dir/startup.time"

array_src="$work_dir/haxe_cases/array"
write_haxe_array_case "$array_src"

for profile in "${profiles[@]}"; do
  log "array case ($profile)"
  case_dir="$work_dir/array/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/array_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$array_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "array_haxe_${profile}" "array" "$profile" "haxe" \
    "$bin_path" "$array_iters" "$case_dir/startup.time"
done

log "array pure Go baseline"
array_pure_dir="$work_dir/array/pure"
array_pure_bin="$array_pure_dir/pure_array"
write_pure_array_module "$array_pure_dir"
(cd "$array_pure_dir" && "$go_bin" build -o "$array_pure_bin" .)
record_metric "array_pure_go" "array" "pure" "pure_go" \
  "$array_pure_bin" "$array_iters" "$array_pure_dir/startup.time"

for profile in "${profiles[@]}"; do
  log "tui case ($profile)"
  case_dir="$work_dir/tui/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/tui_haxe_${profile}"
  mkdir -p "$case_dir"

  (
    cd "$root_dir/examples/tui_todo"
    "$haxe_bin" "compile.${profile}.ci.hxml" -D "go_output=$out_dir" -D go_no_build >/dev/null
  )
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "tui_haxe_${profile}" "tui" "$profile" "haxe" \
    "$bin_path" "$tui_iters" "$case_dir/startup.time"
done

haxe_version="$($haxe_bin --version 2>/dev/null | tr -d '\r' | head -n 1 || true)"
go_version="$($go_bin version 2>/dev/null | tr -d '\r' | head -n 1 || true)"

GO_PERF_METRICS_TSV="$metrics_tsv" \
GO_PERF_CURRENT_JSON="$current_json" \
GO_PERF_COMPARISON_JSON="$comparison_json" \
GO_PERF_SUMMARY_MD="$summary_md" \
GO_PERF_WARNINGS_TXT="$warnings_txt" \
GO_PERF_HARD_FAILURES_TXT="$hard_failures_txt" \
GO_PERF_BASELINE_FILE="$baseline_file" \
GO_PERF_BASELINE_DISPLAY="$baseline_display" \
GO_PERF_UPDATE_BASELINE="$update_baseline" \
GO_PERF_SIZE_WARN_PCT="$size_warn_pct" \
GO_PERF_RUNTIME_WARN_PCT="$runtime_warn_pct" \
GO_PERF_ENFORCE_METAL_BUDGET="$enforce_metal_budget" \
GO_PERF_METAL_SIZE_FAIL_PCT="$metal_size_fail_pct" \
GO_PERF_METAL_RUNTIME_FAIL_PCT="$metal_runtime_fail_pct" \
GO_PERF_HELLO_ITERS="$hello_iters" \
GO_PERF_ARRAY_ITERS="$array_iters" \
GO_PERF_TUI_ITERS="$tui_iters" \
GO_PERF_HAXE_VERSION="$haxe_version" \
GO_PERF_GO_VERSION="$go_version" \
node <<'NODE'
const fs = require("fs");
const path = require("path");

const metricsPath = process.env.GO_PERF_METRICS_TSV;
const currentJsonPath = process.env.GO_PERF_CURRENT_JSON;
const comparisonJsonPath = process.env.GO_PERF_COMPARISON_JSON;
const summaryPath = process.env.GO_PERF_SUMMARY_MD;
const warningsPath = process.env.GO_PERF_WARNINGS_TXT;
const hardFailuresPath = process.env.GO_PERF_HARD_FAILURES_TXT;
const baselinePath = process.env.GO_PERF_BASELINE_FILE;
const baselineDisplay = process.env.GO_PERF_BASELINE_DISPLAY || baselinePath;
const updateBaseline = process.env.GO_PERF_UPDATE_BASELINE === "1";
const sizeWarnPct = Number(process.env.GO_PERF_SIZE_WARN_PCT || "5");
const runtimeWarnPct = Number(process.env.GO_PERF_RUNTIME_WARN_PCT || "10");
const enforceMetalBudget = /^(1|true|yes|on)$/i.test(process.env.GO_PERF_ENFORCE_METAL_BUDGET || "0");
const metalSizeFailPct = Number(process.env.GO_PERF_METAL_SIZE_FAIL_PCT || "25");
const metalRuntimeFailPct = Number(process.env.GO_PERF_METAL_RUNTIME_FAIL_PCT || "100");
const helloIters = Number(process.env.GO_PERF_HELLO_ITERS || "300");
const arrayIters = Number(process.env.GO_PERF_ARRAY_ITERS || "300");
const tuiIters = Number(process.env.GO_PERF_TUI_ITERS || "30");
const haxeVersion = process.env.GO_PERF_HAXE_VERSION || "";
const goVersion = process.env.GO_PERF_GO_VERSION || "";

const profiles = ["portable", "gopher", "metal"];

function parseMetrics(tsvPath) {
  const raw = fs.readFileSync(tsvPath, "utf8").trim();
  const lines = raw.split(/\r?\n/);
  const header = lines.shift();
  const cols = header.split("\t");
  return lines
    .filter((line) => line.trim().length > 0)
    .map((line) => {
      const fields = line.split("\t");
      const entry = {};
      cols.forEach((col, index) => {
        entry[col] = fields[index] ?? "";
      });
      return {
        id: entry.id,
        case: entry.case,
        profile: entry.profile,
        kind: entry.kind,
        binary_bytes: Number(entry.binary_bytes),
        stripped_bytes: Number(entry.stripped_bytes),
        startup_avg_ms: Number(entry.startup_avg_ms),
        startup_iterations: Number(entry.startup_iterations),
      };
    });
}

const metrics = parseMetrics(metricsPath);
const byId = Object.fromEntries(metrics.map((metric) => [metric.id, metric]));

function requireMetric(id) {
  const found = byId[id];
  if (!found) {
    throw new Error(`Missing metric: ${id}`);
  }
  return found;
}

function ratio(current, base) {
  if (base === 0) {
    return 0;
  }
  return current / base;
}

function buildCaseOverhead(caseName) {
  const pure = requireMetric(`${caseName}_pure_go`);
  const out = {};
  for (const profile of profiles) {
    const metric = requireMetric(`${caseName}_haxe_${profile}`);
    out[profile] = {
      binaryRatio: ratio(metric.binary_bytes, pure.binary_bytes),
      strippedRatio: ratio(metric.stripped_bytes, pure.stripped_bytes),
      startupRatio: ratio(metric.startup_avg_ms, pure.startup_avg_ms),
    };
  }
  return out;
}

const helloOverheadRatios = buildCaseOverhead("hello");
const arrayOverheadRatios = buildCaseOverhead("array");

const tuiMetrics = Object.fromEntries(
  profiles.map((profile) => [profile, requireMetric(`tui_haxe_${profile}`)])
);
const tuiMin = {
  binary_bytes: Math.min(...profiles.map((profile) => tuiMetrics[profile].binary_bytes)),
  stripped_bytes: Math.min(...profiles.map((profile) => tuiMetrics[profile].stripped_bytes)),
  startup_avg_ms: Math.min(...profiles.map((profile) => tuiMetrics[profile].startup_avg_ms)),
};
const tuiRelativeToMin = {};
for (const profile of profiles) {
  const metric = tuiMetrics[profile];
  tuiRelativeToMin[profile] = {
    binaryRatio: ratio(metric.binary_bytes, tuiMin.binary_bytes),
    strippedRatio: ratio(metric.stripped_bytes, tuiMin.stripped_bytes),
    startupRatio: ratio(metric.startup_avg_ms, tuiMin.startup_avg_ms),
  };
}

const current = {
  schemaVersion: 1,
  generatedAt: new Date().toISOString(),
  toolchain: {
    haxe: haxeVersion,
    go: goVersion,
  },
  thresholds: {
    sizeWarnPct,
    runtimeWarnPct,
  },
  startupLoops: {
    hello: helloIters,
    array: arrayIters,
    tui: tuiIters,
  },
  metrics,
  derived: {
    helloOverheadRatios,
    arrayOverheadRatios,
    tuiRelativeToMin,
  },
};

fs.mkdirSync(path.dirname(currentJsonPath), { recursive: true });
fs.writeFileSync(currentJsonPath, `${JSON.stringify(current, null, 2)}\n`);

const baselinePayload = {
  schemaVersion: 1,
  generatedAt: current.generatedAt,
  thresholds: current.thresholds,
  startupLoops: current.startupLoops,
  derivedBaseline: current.derived,
};

if (updateBaseline) {
  fs.mkdirSync(path.dirname(baselinePath), { recursive: true });
  fs.writeFileSync(baselinePath, `${JSON.stringify(baselinePayload, null, 2)}\n`);
}

const warnings = [];
const hardFailures = [];

function compareGroup(groupLabel, currentGroup, baselineGroup) {
  if (!baselineGroup) {
    warnings.push(`${groupLabel}: missing baseline group`);
    return;
  }

  const specs = [
    { key: "binaryRatio", label: "binary ratio", warnPct: sizeWarnPct },
    { key: "strippedRatio", label: "stripped ratio", warnPct: sizeWarnPct },
    { key: "startupRatio", label: "startup ratio", warnPct: runtimeWarnPct },
  ];

  for (const profile of profiles) {
    const currentProfile = currentGroup[profile];
    const baselineProfile = baselineGroup[profile];
    if (!currentProfile || !baselineProfile) {
      warnings.push(`${groupLabel}.${profile}: missing data in current/baseline`);
      continue;
    }

    for (const spec of specs) {
      const currentValue = Number(currentProfile[spec.key]);
      const baselineValue = Number(baselineProfile[spec.key]);
      if (!Number.isFinite(currentValue) || !Number.isFinite(baselineValue) || baselineValue <= 0) {
        continue;
      }
      const maxAllowed = baselineValue * (1 + spec.warnPct / 100);
      if (currentValue > maxAllowed) {
        const increasePct = ((currentValue / baselineValue) - 1) * 100;
        warnings.push(
          `${groupLabel}.${profile}.${spec.label} +${increasePct.toFixed(2)}% ` +
            `(current=${currentValue.toFixed(6)}, baseline=${baselineValue.toFixed(6)}, budget=+${spec.warnPct.toFixed(2)}%)`
        );
      }
    }
  }
}

function compareMetalHard(groupLabel, currentGroup, baselineGroup) {
  if (!baselineGroup) {
    return;
  }

  const profile = "metal";
  const currentProfile = currentGroup[profile];
  const baselineProfile = baselineGroup[profile];
  if (!currentProfile || !baselineProfile) {
    return;
  }

  const specs = [
    { key: "binaryRatio", label: "binary ratio", failPct: metalSizeFailPct },
    { key: "strippedRatio", label: "stripped ratio", failPct: metalSizeFailPct },
    { key: "startupRatio", label: "startup ratio", failPct: metalRuntimeFailPct },
  ];

  for (const spec of specs) {
    const currentValue = Number(currentProfile[spec.key]);
    const baselineValue = Number(baselineProfile[spec.key]);
    if (!Number.isFinite(currentValue) || !Number.isFinite(baselineValue) || baselineValue <= 0) {
      continue;
    }
    const maxAllowed = baselineValue * (1 + spec.failPct / 100);
    if (currentValue > maxAllowed) {
      const increasePct = ((currentValue / baselineValue) - 1) * 100;
      hardFailures.push(
        `${groupLabel}.${profile}.${spec.label} +${increasePct.toFixed(2)}% ` +
          `(current=${currentValue.toFixed(6)}, baseline=${baselineValue.toFixed(6)}, budget=+${spec.failPct.toFixed(2)}%)`
      );
    }
  }
}

let baselineLoaded = null;
if (!updateBaseline) {
  if (!fs.existsSync(baselinePath)) {
    warnings.push(`baseline file not found: ${baselineDisplay}`);
  } else {
    baselineLoaded = JSON.parse(fs.readFileSync(baselinePath, "utf8"));
    const baselineDerived = baselineLoaded.derivedBaseline || {};
    compareGroup("hello_overhead", current.derived.helloOverheadRatios, baselineDerived.helloOverheadRatios);
    compareGroup("array_overhead", current.derived.arrayOverheadRatios, baselineDerived.arrayOverheadRatios);
    compareGroup("tui_relative", current.derived.tuiRelativeToMin, baselineDerived.tuiRelativeToMin);
    compareMetalHard("hello_overhead", current.derived.helloOverheadRatios, baselineDerived.helloOverheadRatios);
    compareMetalHard("array_overhead", current.derived.arrayOverheadRatios, baselineDerived.arrayOverheadRatios);
    compareMetalHard("tui_relative", current.derived.tuiRelativeToMin, baselineDerived.tuiRelativeToMin);
  }
}

const comparison = {
  schemaVersion: 1,
  generatedAt: current.generatedAt,
  mode: updateBaseline ? "update-baseline" : "compare",
  baselinePath: baselineDisplay,
  baselineAvailable: baselineLoaded != null || updateBaseline,
  enforceMetalBudget,
  metalHardFailureCount: hardFailures.length,
  metalHardFailureBudgets: {
    sizeFailPct: metalSizeFailPct,
    runtimeFailPct: metalRuntimeFailPct,
  },
  metalWarningCount: warnings.filter((warning) => warning.includes(".metal.")).length,
  warningCount: warnings.length,
  warnings,
  hardFailures,
};
fs.writeFileSync(comparisonJsonPath, `${JSON.stringify(comparison, null, 2)}\n`);
fs.writeFileSync(warningsPath, warnings.length > 0 ? `${warnings.join("\n")}\n` : "");
fs.writeFileSync(hardFailuresPath, hardFailures.length > 0 ? `${hardFailures.join("\n")}\n` : "");

function formatRatio(v) {
  return Number(v).toFixed(3);
}

function ratioTable(title, ratioGroup) {
  const lines = [];
  lines.push(`### ${title}`);
  lines.push("| Profile | Binary x | Stripped x | Startup x |\n| --- | ---: | ---: | ---: |");
  for (const profile of profiles) {
    const row = ratioGroup[profile];
    lines.push(
      `| ${profile} | ${formatRatio(row.binaryRatio)} | ${formatRatio(row.strippedRatio)} | ${formatRatio(row.startupRatio)} |`
    );
  }
  lines.push("");
  return lines.join("\n");
}

const summaryLines = [];
summaryLines.push("## Go Profile Performance Benchmarks");
summaryLines.push("");
summaryLines.push(`- Mode: \`${comparison.mode}\``);
summaryLines.push(`- Size budget: \`+${sizeWarnPct}%\``);
summaryLines.push(`- Runtime budget: \`+${runtimeWarnPct}%\``);
summaryLines.push(`- Metal enforcement: \`${enforceMetalBudget ? "on" : "off"}\``);
summaryLines.push(`- Metal hard budgets: size=\`+${metalSizeFailPct}%\`, runtime=\`+${metalRuntimeFailPct}%\``);
summaryLines.push(`- Startup loops: hello=${helloIters}, array=${arrayIters}, tui=${tuiIters}`);
if (haxeVersion.length > 0 || goVersion.length > 0) {
  summaryLines.push(`- Toolchain: ${haxeVersion || "haxe:unknown"} | ${goVersion || "go:unknown"}`);
}
summaryLines.push("");
summaryLines.push(ratioTable("Hello Overhead (x vs pure Go hello)", current.derived.helloOverheadRatios));
summaryLines.push(ratioTable("Array Overhead (x vs pure Go array loop)", current.derived.arrayOverheadRatios));
summaryLines.push(ratioTable("TUI Profile Spread (x vs fastest/smallest profile in this run)", current.derived.tuiRelativeToMin));

if (warnings.length > 0) {
  summaryLines.push("### Soft Budget Warnings");
  for (const warning of warnings) {
    summaryLines.push(`- ${warning}`);
  }
} else {
  summaryLines.push("### Soft Budget Warnings");
  summaryLines.push("- none");
}
summaryLines.push("");

if (hardFailures.length > 0) {
  summaryLines.push("### Metal Hard-Fail Candidates");
  for (const hardFailure of hardFailures) {
    summaryLines.push(`- ${hardFailure}`);
  }
} else {
  summaryLines.push("### Metal Hard-Fail Candidates");
  summaryLines.push("- none");
}
summaryLines.push("");

fs.writeFileSync(summaryPath, `${summaryLines.join("\n")}\n`);

console.log(`[go-perf] mode=${comparison.mode} warnings=${warnings.length}`);
NODE

warning_count=0
metal_warning_count=0
baseline_warning_count=0
hard_failure_count=0
if [[ -s "$warnings_txt" ]]; then
  while IFS= read -r warning; do
    [[ -n "$warning" ]] || continue
    warning_count=$((warning_count + 1))
    if [[ "$warning" == *".metal."* ]]; then
      metal_warning_count=$((metal_warning_count + 1))
    fi
    if [[ "$warning" == baseline\ file\ not\ found:* ]]; then
      baseline_warning_count=$((baseline_warning_count + 1))
    fi
    echo "::warning::[go-perf] $warning"
  done < "$warnings_txt"
fi

if [[ -s "$hard_failures_txt" ]]; then
  while IFS= read -r hard_failure; do
    [[ -n "$hard_failure" ]] || continue
    hard_failure_count=$((hard_failure_count + 1))
    if is_truthy "$enforce_metal_budget"; then
      echo "::error::[go-perf] $hard_failure"
    else
      echo "::warning::[go-perf][metal-hard-candidate] $hard_failure"
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

if [[ -f "$baseline_file" ]]; then
  cp "$baseline_file" "$results_dir/baseline_used.json"
fi

if is_truthy "$enforce_metal_budget"; then
  if [[ "$hard_failure_count" -gt 0 || "$baseline_warning_count" -gt 0 ]]; then
    echo "::error::[go-perf] metal budget enforcement failed (hard_failures=$hard_failure_count baseline_warnings=$baseline_warning_count)"
    log "failing due to GO_PERF_ENFORCE_METAL_BUDGET with budget regressions"
    log "metrics: $(display_path "$current_json")"
    log "comparison: $(display_path "$comparison_json")"
    log "summary: $(display_path "$summary_md")"
    exit 1
  fi
fi

log "done (warnings=$warning_count, metal_warnings=$metal_warning_count, metal_hard_failures=$hard_failure_count)"
log "metrics: $(display_path "$current_json")"
log "comparison: $(display_path "$comparison_json")"
log "summary: $(display_path "$summary_md")"
