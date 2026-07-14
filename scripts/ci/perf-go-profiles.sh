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
  GO_PERF_PORTABLE_CONCURRENCY_FASTPATH
                            Sets -D reflaxe_go_opt_go_concurrency_fastpath for portable builds in this harness
                            (default: 1/on).
  GO_PERF_HXRT_FEATURES     Base selective hxrt feature list for microbench builds
                            (default: core,string,print; inference still adds case-specific features).
  GO_PERF_DELTA_WARN_PCT    Soft warning threshold for portable-vs-metal startup delta drift vs baseline
                            for selected cases (default: 15).
  GO_PERF_ENFORCE_DELTA_BUDGET
                            Fail when selected portable-vs-metal startup deltas exceed hard budget
                            (default: 0/off).
  GO_PERF_DELTA_FAIL_PCT    Hard-fail threshold for portable-vs-metal startup delta drift vs baseline
                            when GO_PERF_ENFORCE_DELTA_BUDGET=1 (default: 25).
  GO_PERF_DELTA_CASES       Comma-separated microcases for delta budget checks
                            (default: string,string_instance,select,channel).
  GO_PERF_STARTUP_SAMPLES   Startup timing samples per binary; when greater than 1, ratios use
                            the median sample to reduce single-run jitter. Keep this aligned with
                            the checked-in baseline methodology (default: 3).
  GO_PERF_HELLO_ITERS       Startup loop count for hello case (default: 300)
  GO_PERF_ARRAY_ITERS       Startup loop count for array case (default: 300)
  GO_PERF_ATOMIC_ITERS      Startup loop count for atomic case (default: 120)
  GO_PERF_ATOMIC_WORK       Atomic add operations per process run (default: 200000)
  GO_PERF_TUI_ITERS         Startup loop count for tui case (default: 30)
  GO_PERF_CHANNEL_ITERS     Startup loop count for channel case (default: 100)
  GO_PERF_CHANNEL_WORK      Channel send/recv operations per process run (default: 40000)
  GO_PERF_MAP_ITERS         Startup loop count for map case (default: 100)
  GO_PERF_MAP_WORK          Map set/get operations per process run (default: 40000)
  GO_PERF_GENERIC_ITERS     Startup loop count for generic case (default: 100)
  GO_PERF_GENERIC_WORK      Generic push/get operations per process run (default: 50000)
  GO_PERF_STRING_ITERS      Startup loop count for string case (default: 80)
  GO_PERF_STRING_WORK       String concat operations per process run (default: 12000)
  GO_PERF_STRING_INSTANCE_ITERS
                            Startup loop count for string_instance case (default: 60)
  GO_PERF_STRING_INSTANCE_WORK
                            String instance operations per process run (default: 6000)
  GO_PERF_VIRTUAL_ITERS     Startup loop count for virtual dispatch case (default: 100)
  GO_PERF_VIRTUAL_WORK      Virtual dispatch operations per process run (default: 100000)
  GO_PERF_SELECT_ITERS      Startup loop count for select case (default: 100)
  GO_PERF_SELECT_WORK       Select helper operations per process run (default: 40000)
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

measure_startup_ms_once() {
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

measure_startup_stats() {
  local bin="$1"
  local iterations="$2"
  local timing_log="$3"
  local samples="$4"
  local samples_file="${timing_log}.samples"

  : > "$timing_log"
  : > "$samples_file"

  local sample
  for sample in $(seq 1 "$samples"); do
    local sample_log="${timing_log}.${sample}"
    local sample_ms
    sample_ms="$(measure_startup_ms_once "$bin" "$iterations" "$sample_log")"
    printf "%s\n" "$sample_ms" >> "$samples_file"
    {
      printf "sample %s: %s ms/run\n" "$sample" "$sample_ms"
      cat "$sample_log"
      printf "\n"
    } >> "$timing_log"
  done

  python3 - "$samples_file" <<'PY'
import math
import statistics
import sys

with open(sys.argv[1], "r", encoding="utf-8") as handle:
    values = sorted(float(line.strip()) for line in handle if line.strip())

if not values:
    raise SystemExit("no startup timing samples recorded")

p95_index = max(0, math.ceil(len(values) * 0.95) - 1)
avg = sum(values) / len(values)
median = statistics.median(values)
print(
    f"{avg:.6f}\t{median:.6f}\t{values[0]:.6f}\t"
    f"{values[-1]:.6f}\t{values[p95_index]:.6f}"
)
PY
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

write_haxe_atomic_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
import haxe.atomic.AtomicInt;

class Main {
  static function main():Void {
    var atom = new AtomicInt(0);
    var i = 0;
    while (i < ${work}) {
      atom.add(1);
      i++;
    }
    Sys.println(atom.load());
  }
}
EOF
}

write_haxe_channel_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
import go.Chan;
import go.Go;

class Main {
  static function main():Void {
    var channel:Chan<Int> = Go.newChan(${work});
    var i = 0;
    while (i < ${work}) {
      channel.send(i);
      i++;
    }

    var last = 0;
    i = 0;
    while (i < ${work}) {
      last = channel.recvOr(0);
      i++;
    }
    channel.close();
    Sys.println(last);
  }
}
EOF
}

write_haxe_map_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
import go.Map;

class Main {
  static function main():Void {
    var values = new Map<Int, Int>();
    var i = 0;
    while (i < ${work}) {
      values.set(i, i + 3);
      i++;
    }

    var found = 0;
    i = 0;
    while (i < ${work}) {
      if (values.exists(i)) {
        found++;
      }
      i++;
    }
    Sys.println(found);
  }
}
EOF
}

write_haxe_generic_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
import go.Slice;

class Main {
  static function main():Void {
    var values = new Slice<Int>();
    var i = 0;
    while (i < ${work}) {
      values.push(i);
      i++;
    }

    var hits = 0;
    i = 0;
    while (i < ${work}) {
      values.get(i);
      hits++;
      i++;
    }
    var view = values.toArray();
    Sys.println(hits + view.length);
  }
}
EOF
}

write_haxe_string_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
class Main {
  static function main():Void {
    var out = "";
    var i = 0;
    while (i < ${work}) {
      out = out + Std.string(i & 15);
      i++;
    }
    Sys.println(out.length);
  }
}
EOF
}

write_haxe_string_instance_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
class Main {
  static function main():Void {
    var seed = "héllo_";
    var out = "";
    var i = 0;
    while (i < ${work}) {
      var text = seed + Std.string(i & 1023);
      var len = text.length;
      out = out + text.charAt(i % len);
      var code = text.charCodeAt((i + 1) % len);
      if (code != null) {
        out = out + Std.string(code);
      }
      out = out + text.substring(0, i % len);
      out = out + text.substr(-2);
      i++;
    }
    Sys.println(out.length);
  }
}
EOF
}

write_haxe_virtual_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
interface Worker {
  public function step(value:Int):Int;
}

class AddWorker implements Worker {
  public function new() {}

  public function step(value:Int):Int {
    return value + 1;
  }
}

class Main {
  static function main():Void {
    var workers:Array<Worker> = [new AddWorker(), new AddWorker(), new AddWorker(), new AddWorker()];
    var total = 0;
    var i = 0;
    while (i < ${work}) {
      var worker = workers[i & 3];
      total += worker.step(i);
      i++;
    }
    Sys.println(total);
  }
}
EOF
}

write_haxe_select_case() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/Main.hx" <<EOF
import go.Chan;
import go.Go;
import go.Select;

class Main {
  static function main():Void {
    var gate:Chan<Int> = Go.newChan(1);
    var left:Chan<Int> = Go.newChan(1);
    var right:Chan<Int> = Go.newChan(1);
    var total = 0;
    var i = 0;

    while (i < ${work}) {
      total += switch (Select.send(gate, i)) {
        case Sent: 1;
        case Defaulted: 0;
      };

      total += switch (Select.recv(gate)) {
        case Received(value): value;
        case Defaulted: 0;
      };

      if ((i & 1) == 0) {
        left.send(i);
      } else {
        right.send(i);
      }
      total += switch (Select.recv2(left, right)) {
        case First(value): value;
        case Second(value): value;
        case Defaulted: 0;
      };

      total += switch (Select.send2(left, i + 1, right, i + 2)) {
        case FirstSent: 1;
        case SecondSent: 2;
        case Defaulted: 0;
      };
      total += switch (Select.recv2(left, right)) {
        case First(value): value;
        case Second(value): value;
        case Defaulted: 0;
      };

      i++;
    }

    gate.close();
    left.close();
    right.close();
    Sys.println(total);
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
    local -a haxe_args=(
      -cp .
      -cp "$root_dir/src"
      -cp "$root_dir/vendor/reflaxe/src"
      -cp "$root_dir/std"
      -cp "$root_dir/std/_std"
      -cp "$root_dir/std/go/_std"
      --macro "reflaxe.go.CompilerBootstrap.Start()"
      --macro "reflaxe.go.CompilerInit.Start()"
      -D "go_output=$out_dir"
      -D "reflaxe_go_profile=$profile"
      -D go_no_build
      -D reflaxe.dont_output_metadata_id
      -D no-traces
      -D no_traces
      -D "reflaxe_go_hxrt_features=$hxrt_features"
      -main Main
    )
    if [[ "$profile" == "portable" ]]; then
      haxe_args+=(-D "reflaxe_go_opt_go_concurrency_fastpath=$portable_concurrency_fastpath_bool")
    fi
    "$haxe_bin" "${haxe_args[@]}" >/dev/null
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

write_pure_atomic_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_atomic

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import (
  "fmt"
  "sync/atomic"
)

func main() {
  var cell atomic.Int64
  for i := 0; i < ${work}; i++ {
    cell.Add(1)
  }
  fmt.Println(cell.Load())
}
EOF
}

write_pure_channel_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_channel

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import "fmt"

func main() {
  channel := make(chan int, ${work})
  for i := 0; i < ${work}; i++ {
    channel <- i
  }

  last := 0
  for i := 0; i < ${work}; i++ {
    last = <-channel
  }
  close(channel)
  fmt.Println(last)
}
EOF
}

write_pure_map_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_map

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import "fmt"

func main() {
  values := make(map[int]int, ${work})
  for i := 0; i < ${work}; i++ {
    values[i] = i + 3
  }

  found := 0
  for i := 0; i < ${work}; i++ {
    if _, ok := values[i]; ok {
      found++
    }
  }
  fmt.Println(found)
}
EOF
}

write_pure_generic_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_generic

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import "fmt"

type IntBag[T ~int] struct {
  values []T
}

func (b *IntBag[T]) Add(value T) {
  b.values = append(b.values, value)
}

func (b *IntBag[T]) Get(index int) T {
  return b.values[index]
}

func (b *IntBag[T]) Len() int {
  return len(b.values)
}

func main() {
  var bag IntBag[int]
  for i := 0; i < ${work}; i++ {
    bag.Add(i)
  }

  hits := 0
  for i := 0; i < ${work}; i++ {
    _ = bag.Get(i)
    hits++
  }
  fmt.Println(hits + bag.Len())
}
EOF
}

write_pure_string_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_string

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import (
  "fmt"
  "strconv"
)

func main() {
  out := ""
  for i := 0; i < ${work}; i++ {
    out += strconv.Itoa(i & 15)
  }
  fmt.Println(len(out))
}
EOF
}

write_pure_string_instance_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_string_instance

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import (
  "fmt"
  "strconv"
)

func main() {
  seed := "héllo_"
  out := ""

  for i := 0; i < ${work}; i++ {
    text := seed + strconv.Itoa(i&1023)
    runes := []rune(text)
    length := len(runes)

    out += string(runes[i%length])
    out += strconv.Itoa(int(runes[(i+1)%length]))
    out += string(runes[:i%length])
    start := length - 2
    if start < 0 {
      start = 0
    }
    out += string(runes[start:])
  }

  fmt.Println(len([]rune(out)))
}
EOF
}

write_pure_virtual_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_virtual

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import "fmt"

type Worker interface {
  Step(value int) int
}

type AddWorker struct{}

func (AddWorker) Step(value int) int {
  return value + 1
}

func main() {
  workers := []Worker{AddWorker{}, AddWorker{}, AddWorker{}, AddWorker{}}
  total := 0
  for i := 0; i < ${work}; i++ {
    worker := workers[i&3]
    total += worker.Step(i)
  }
  fmt.Println(total)
}
EOF
}

write_pure_select_module() {
  local dir="$1"
  local work="$2"
  mkdir -p "$dir"
  cat > "$dir/go.mod" <<'EOF'
module pure_select

go 1.22
EOF
  cat > "$dir/main.go" <<EOF
package main

import "fmt"

func main() {
  gate := make(chan int, 1)
  left := make(chan int, 1)
  right := make(chan int, 1)
  total := 0

  for i := 0; i < ${work}; i++ {
    select {
    case gate <- i:
      total += 1
    default:
    }

    select {
    case value := <-gate:
      total += value
    default:
    }

    if i%2 == 0 {
      left <- i
    } else {
      right <- i
    }
    select {
    case value := <-left:
      total += value
    case value := <-right:
      total += value
    default:
    }

    select {
    case left <- (i + 1):
      total += 1
    case right <- (i + 2):
      total += 2
    default:
    }
    select {
    case value := <-left:
      total += value
    case value := <-right:
      total += value
    default:
    }
  }

  close(gate)
  close(left)
  close(right)
  fmt.Println(total)
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

  local startup_stats
  startup_stats="$(measure_startup_stats "$bin_path" "$iterations" "$timing_log" "$startup_samples")"
  local startup_avg_ms
  local startup_median_ms
  local startup_min_ms
  local startup_max_ms
  local startup_p95_ms
  IFS=$'\t' read -r startup_avg_ms startup_median_ms startup_min_ms startup_max_ms startup_p95_ms <<< "$startup_stats"
  local bin_bytes
  bin_bytes="$(filesize_bytes "$bin_path")"
  local stripped_bytes
  stripped_bytes="$(stripped_size_bytes "$bin_path")"

  printf "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n" \
    "$id" "$case_name" "$profile" "$kind" \
    "$bin_bytes" "$stripped_bytes" "$startup_avg_ms" "$startup_median_ms" \
    "$startup_min_ms" "$startup_max_ms" "$startup_p95_ms" "$iterations" \
    "$startup_samples" >> "$metrics_tsv"
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
portable_concurrency_fastpath="${GO_PERF_PORTABLE_CONCURRENCY_FASTPATH:-1}"
hxrt_features="${GO_PERF_HXRT_FEATURES:-core,string,print}"
delta_warn_pct="${GO_PERF_DELTA_WARN_PCT:-15}"
enforce_delta_budget="${GO_PERF_ENFORCE_DELTA_BUDGET:-0}"
delta_fail_pct="${GO_PERF_DELTA_FAIL_PCT:-25}"
delta_cases="${GO_PERF_DELTA_CASES:-string,string_instance,select,channel}"
startup_samples="${GO_PERF_STARTUP_SAMPLES:-3}"
hello_iters="${GO_PERF_HELLO_ITERS:-300}"
array_iters="${GO_PERF_ARRAY_ITERS:-300}"
atomic_iters="${GO_PERF_ATOMIC_ITERS:-120}"
atomic_work="${GO_PERF_ATOMIC_WORK:-200000}"
tui_iters="${GO_PERF_TUI_ITERS:-30}"
channel_iters="${GO_PERF_CHANNEL_ITERS:-100}"
channel_work="${GO_PERF_CHANNEL_WORK:-40000}"
map_iters="${GO_PERF_MAP_ITERS:-100}"
map_work="${GO_PERF_MAP_WORK:-40000}"
generic_iters="${GO_PERF_GENERIC_ITERS:-100}"
generic_work="${GO_PERF_GENERIC_WORK:-50000}"
string_iters="${GO_PERF_STRING_ITERS:-80}"
string_work="${GO_PERF_STRING_WORK:-12000}"
string_instance_iters="${GO_PERF_STRING_INSTANCE_ITERS:-60}"
string_instance_work="${GO_PERF_STRING_INSTANCE_WORK:-6000}"
virtual_iters="${GO_PERF_VIRTUAL_ITERS:-100}"
virtual_work="${GO_PERF_VIRTUAL_WORK:-100000}"
select_iters="${GO_PERF_SELECT_ITERS:-100}"
select_work="${GO_PERF_SELECT_WORK:-40000}"

portable_concurrency_fastpath_bool="0"
if is_truthy "$portable_concurrency_fastpath"; then
  portable_concurrency_fastpath_bool="1"
fi

if [[ -x /usr/bin/time ]]; then
  time_bin="/usr/bin/time"
else
  fail "required timing command not found: /usr/bin/time"
fi

require_command "$haxe_bin"
require_command "$go_bin"
require_command node
require_command python3

if ! [[ "$startup_samples" =~ ^[0-9]+$ ]] || [[ "$startup_samples" -lt 1 ]]; then
  fail "GO_PERF_STARTUP_SAMPLES must be a positive integer"
fi

work_dir="$cache_root/work"
results_dir="$cache_root/results"
metrics_tsv="$results_dir/raw_metrics.tsv"
current_json="$results_dir/current.json"
comparison_json="$results_dir/comparison.json"
summary_md="$results_dir/summary.md"
warnings_txt="$results_dir/warnings.txt"
hard_failures_txt="$results_dir/hard_failures.txt"
warning_history_json="$results_dir/warning_history.json"
warning_history_md="$results_dir/warning_history.md"
delta_dry_run_json="$results_dir/delta_hard_gate_dry_run.json"
delta_dry_run_md="$results_dir/delta_hard_gate_dry_run.md"

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

printf "id\tcase\tprofile\tkind\tbinary_bytes\tstripped_bytes\tstartup_avg_ms\tstartup_median_ms\tstartup_min_ms\tstartup_max_ms\tstartup_p95_ms\tstartup_iterations\tstartup_sample_count\n" > "$metrics_tsv"

log "collecting metrics (results: $(display_path "$results_dir"))"

declare -a profiles=(portable metal)

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

atomic_src="$work_dir/haxe_cases/atomic"
write_haxe_atomic_case "$atomic_src" "$atomic_work"

for profile in "${profiles[@]}"; do
  log "atomic case ($profile)"
  case_dir="$work_dir/atomic/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/atomic_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$atomic_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "atomic_haxe_${profile}" "atomic" "$profile" "haxe" \
    "$bin_path" "$atomic_iters" "$case_dir/startup.time"
done

log "atomic pure Go baseline"
atomic_pure_dir="$work_dir/atomic/pure"
atomic_pure_bin="$atomic_pure_dir/pure_atomic"
write_pure_atomic_module "$atomic_pure_dir" "$atomic_work"
(cd "$atomic_pure_dir" && "$go_bin" build -o "$atomic_pure_bin" .)
record_metric "atomic_pure_go" "atomic" "pure" "pure_go" \
  "$atomic_pure_bin" "$atomic_iters" "$atomic_pure_dir/startup.time"

channel_src="$work_dir/haxe_cases/channel"
write_haxe_channel_case "$channel_src" "$channel_work"

for profile in "${profiles[@]}"; do
  log "channel case ($profile)"
  case_dir="$work_dir/channel/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/channel_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$channel_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "channel_haxe_${profile}" "channel" "$profile" "haxe" \
    "$bin_path" "$channel_iters" "$case_dir/startup.time"
done

log "channel pure Go baseline"
channel_pure_dir="$work_dir/channel/pure"
channel_pure_bin="$channel_pure_dir/pure_channel"
write_pure_channel_module "$channel_pure_dir" "$channel_work"
(cd "$channel_pure_dir" && "$go_bin" build -o "$channel_pure_bin" .)
record_metric "channel_pure_go" "channel" "pure" "pure_go" \
  "$channel_pure_bin" "$channel_iters" "$channel_pure_dir/startup.time"

map_src="$work_dir/haxe_cases/map"
write_haxe_map_case "$map_src" "$map_work"

for profile in "${profiles[@]}"; do
  log "map case ($profile)"
  case_dir="$work_dir/map/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/map_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$map_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "map_haxe_${profile}" "map" "$profile" "haxe" \
    "$bin_path" "$map_iters" "$case_dir/startup.time"
done

log "map pure Go baseline"
map_pure_dir="$work_dir/map/pure"
map_pure_bin="$map_pure_dir/pure_map"
write_pure_map_module "$map_pure_dir" "$map_work"
(cd "$map_pure_dir" && "$go_bin" build -o "$map_pure_bin" .)
record_metric "map_pure_go" "map" "pure" "pure_go" \
  "$map_pure_bin" "$map_iters" "$map_pure_dir/startup.time"

generic_src="$work_dir/haxe_cases/generic"
write_haxe_generic_case "$generic_src" "$generic_work"

for profile in "${profiles[@]}"; do
  log "generic case ($profile)"
  case_dir="$work_dir/generic/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/generic_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$generic_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "generic_haxe_${profile}" "generic" "$profile" "haxe" \
    "$bin_path" "$generic_iters" "$case_dir/startup.time"
done

log "generic pure Go baseline"
generic_pure_dir="$work_dir/generic/pure"
generic_pure_bin="$generic_pure_dir/pure_generic"
write_pure_generic_module "$generic_pure_dir" "$generic_work"
(cd "$generic_pure_dir" && "$go_bin" build -o "$generic_pure_bin" .)
record_metric "generic_pure_go" "generic" "pure" "pure_go" \
  "$generic_pure_bin" "$generic_iters" "$generic_pure_dir/startup.time"

string_src="$work_dir/haxe_cases/string"
write_haxe_string_case "$string_src" "$string_work"

for profile in "${profiles[@]}"; do
  log "string case ($profile)"
  case_dir="$work_dir/string/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/string_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$string_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "string_haxe_${profile}" "string" "$profile" "haxe" \
    "$bin_path" "$string_iters" "$case_dir/startup.time"
done

log "string pure Go baseline"
string_pure_dir="$work_dir/string/pure"
string_pure_bin="$string_pure_dir/pure_string"
write_pure_string_module "$string_pure_dir" "$string_work"
(cd "$string_pure_dir" && "$go_bin" build -o "$string_pure_bin" .)
record_metric "string_pure_go" "string" "pure" "pure_go" \
  "$string_pure_bin" "$string_iters" "$string_pure_dir/startup.time"

string_instance_src="$work_dir/haxe_cases/string_instance"
write_haxe_string_instance_case "$string_instance_src" "$string_instance_work"

for profile in "${profiles[@]}"; do
  log "string_instance case ($profile)"
  case_dir="$work_dir/string_instance/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/string_instance_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$string_instance_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "string_instance_haxe_${profile}" "string_instance" "$profile" "haxe" \
    "$bin_path" "$string_instance_iters" "$case_dir/startup.time"
done

log "string_instance pure Go baseline"
string_instance_pure_dir="$work_dir/string_instance/pure"
string_instance_pure_bin="$string_instance_pure_dir/pure_string_instance"
write_pure_string_instance_module "$string_instance_pure_dir" "$string_instance_work"
(cd "$string_instance_pure_dir" && "$go_bin" build -o "$string_instance_pure_bin" .)
record_metric "string_instance_pure_go" "string_instance" "pure" "pure_go" \
  "$string_instance_pure_bin" "$string_instance_iters" "$string_instance_pure_dir/startup.time"

virtual_src="$work_dir/haxe_cases/virtual"
write_haxe_virtual_case "$virtual_src" "$virtual_work"

for profile in "${profiles[@]}"; do
  log "virtual case ($profile)"
  case_dir="$work_dir/virtual/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/virtual_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$virtual_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "virtual_haxe_${profile}" "virtual" "$profile" "haxe" \
    "$bin_path" "$virtual_iters" "$case_dir/startup.time"
done

log "virtual pure Go baseline"
virtual_pure_dir="$work_dir/virtual/pure"
virtual_pure_bin="$virtual_pure_dir/pure_virtual"
write_pure_virtual_module "$virtual_pure_dir" "$virtual_work"
(cd "$virtual_pure_dir" && "$go_bin" build -o "$virtual_pure_bin" .)
record_metric "virtual_pure_go" "virtual" "pure" "pure_go" \
  "$virtual_pure_bin" "$virtual_iters" "$virtual_pure_dir/startup.time"

select_src="$work_dir/haxe_cases/select"
write_haxe_select_case "$select_src" "$select_work"

for profile in "${profiles[@]}"; do
  log "select case ($profile)"
  case_dir="$work_dir/select/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/select_haxe_${profile}"
  mkdir -p "$case_dir"

  compile_haxe_case "$select_src" "$out_dir" "$profile"
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "select_haxe_${profile}" "select" "$profile" "haxe" \
    "$bin_path" "$select_iters" "$case_dir/startup.time"
done

log "select pure Go baseline"
select_pure_dir="$work_dir/select/pure"
select_pure_bin="$select_pure_dir/pure_select"
write_pure_select_module "$select_pure_dir" "$select_work"
(cd "$select_pure_dir" && "$go_bin" build -o "$select_pure_bin" .)
record_metric "select_pure_go" "select" "pure" "pure_go" \
  "$select_pure_bin" "$select_iters" "$select_pure_dir/startup.time"

tui_profiles_collected=0
for profile in "${profiles[@]}"; do
  log "tui case ($profile)"
  case_dir="$work_dir/tui/$profile"
  out_dir="$case_dir/out"
  bin_path="$case_dir/tui_haxe_${profile}"
  compile_file="$root_dir/examples/tui_todo/compile.${profile}.ci.hxml"

  if [[ ! -f "$compile_file" ]]; then
    log "tui case ($profile) skipped: compile.${profile}.ci.hxml not found"
    continue
  fi

  mkdir -p "$case_dir"
  (
    cd "$root_dir/examples/tui_todo"
    "$haxe_bin" "compile.${profile}.ci.hxml" -D "go_output=$out_dir" -D go_no_build >/dev/null
  )
  (cd "$out_dir" && "$go_bin" build -o "$bin_path" .)

  record_metric "tui_haxe_${profile}" "tui" "$profile" "haxe" \
    "$bin_path" "$tui_iters" "$case_dir/startup.time"
  tui_profiles_collected=$((tui_profiles_collected + 1))
done

if [[ "$tui_profiles_collected" -eq 0 ]]; then
  fail "tui case: no profile compile files found (expected at least one compile.<profile>.ci.hxml)"
fi

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
GO_PERF_PORTABLE_CONCURRENCY_FASTPATH="$portable_concurrency_fastpath_bool" \
GO_PERF_HXRT_FEATURES="$hxrt_features" \
GO_PERF_DELTA_WARN_PCT="$delta_warn_pct" \
GO_PERF_ENFORCE_DELTA_BUDGET="$enforce_delta_budget" \
GO_PERF_DELTA_FAIL_PCT="$delta_fail_pct" \
GO_PERF_DELTA_CASES="$delta_cases" \
GO_PERF_STARTUP_SAMPLES="$startup_samples" \
GO_PERF_HELLO_ITERS="$hello_iters" \
GO_PERF_ARRAY_ITERS="$array_iters" \
GO_PERF_ATOMIC_ITERS="$atomic_iters" \
GO_PERF_ATOMIC_WORK="$atomic_work" \
GO_PERF_TUI_ITERS="$tui_iters" \
GO_PERF_CHANNEL_ITERS="$channel_iters" \
GO_PERF_CHANNEL_WORK="$channel_work" \
GO_PERF_MAP_ITERS="$map_iters" \
GO_PERF_MAP_WORK="$map_work" \
GO_PERF_GENERIC_ITERS="$generic_iters" \
GO_PERF_GENERIC_WORK="$generic_work" \
GO_PERF_STRING_ITERS="$string_iters" \
GO_PERF_STRING_WORK="$string_work" \
GO_PERF_STRING_INSTANCE_ITERS="$string_instance_iters" \
GO_PERF_STRING_INSTANCE_WORK="$string_instance_work" \
GO_PERF_VIRTUAL_ITERS="$virtual_iters" \
GO_PERF_VIRTUAL_WORK="$virtual_work" \
GO_PERF_SELECT_ITERS="$select_iters" \
GO_PERF_SELECT_WORK="$select_work" \
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
const portableConcurrencyFastpathEnabled = /^(1|true|yes|on)$/i.test(process.env.GO_PERF_PORTABLE_CONCURRENCY_FASTPATH || "1");
const hxrtFeatures = process.env.GO_PERF_HXRT_FEATURES || "core,string,print";
const deltaWarnPct = Number(process.env.GO_PERF_DELTA_WARN_PCT || "15");
const enforceDeltaBudget = /^(1|true|yes|on)$/i.test(process.env.GO_PERF_ENFORCE_DELTA_BUDGET || "0");
const deltaFailPct = Number(process.env.GO_PERF_DELTA_FAIL_PCT || "25");
const deltaCases = (process.env.GO_PERF_DELTA_CASES || "string,string_instance,select,channel")
  .split(",")
  .map((value) => value.trim().toLowerCase())
  .filter((value) => value.length > 0);
const uniqueDeltaCases = [...new Set(deltaCases)];
const startupSamples = Number(process.env.GO_PERF_STARTUP_SAMPLES || "3");
const helloIters = Number(process.env.GO_PERF_HELLO_ITERS || "300");
const arrayIters = Number(process.env.GO_PERF_ARRAY_ITERS || "300");
const atomicIters = Number(process.env.GO_PERF_ATOMIC_ITERS || "120");
const atomicWork = Number(process.env.GO_PERF_ATOMIC_WORK || "200000");
const tuiIters = Number(process.env.GO_PERF_TUI_ITERS || "30");
const channelIters = Number(process.env.GO_PERF_CHANNEL_ITERS || "100");
const channelWork = Number(process.env.GO_PERF_CHANNEL_WORK || "40000");
const mapIters = Number(process.env.GO_PERF_MAP_ITERS || "100");
const mapWork = Number(process.env.GO_PERF_MAP_WORK || "40000");
const genericIters = Number(process.env.GO_PERF_GENERIC_ITERS || "100");
const genericWork = Number(process.env.GO_PERF_GENERIC_WORK || "50000");
const stringIters = Number(process.env.GO_PERF_STRING_ITERS || "80");
const stringWork = Number(process.env.GO_PERF_STRING_WORK || "12000");
const stringInstanceIters = Number(process.env.GO_PERF_STRING_INSTANCE_ITERS || "60");
const stringInstanceWork = Number(process.env.GO_PERF_STRING_INSTANCE_WORK || "6000");
const virtualIters = Number(process.env.GO_PERF_VIRTUAL_ITERS || "100");
const virtualWork = Number(process.env.GO_PERF_VIRTUAL_WORK || "100000");
const selectIters = Number(process.env.GO_PERF_SELECT_ITERS || "100");
const selectWork = Number(process.env.GO_PERF_SELECT_WORK || "40000");
const haxeVersion = process.env.GO_PERF_HAXE_VERSION || "";
const goVersion = process.env.GO_PERF_GO_VERSION || "";

const profiles = ["portable", "metal"];

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
        startup_median_ms: Number(entry.startup_median_ms || entry.startup_avg_ms),
        startup_min_ms: Number(entry.startup_min_ms || entry.startup_avg_ms),
        startup_max_ms: Number(entry.startup_max_ms || entry.startup_avg_ms),
        startup_p95_ms: Number(entry.startup_p95_ms || entry.startup_avg_ms),
        startup_iterations: Number(entry.startup_iterations),
        startup_sample_count: Number(entry.startup_sample_count || "1"),
      };
    });
}

const metrics = parseMetrics(metricsPath);
const byId = Object.fromEntries(metrics.map((metric) => [metric.id, metric]));
const tuiProfiles = [...new Set(
  metrics
    .filter((metric) => metric.case === "tui" && metric.kind === "haxe" && profiles.includes(metric.profile))
    .map((metric) => metric.profile)
)]
  .sort((a, b) => profiles.indexOf(a) - profiles.indexOf(b));
if (tuiProfiles.length === 0) {
  throw new Error("Missing metric: tui_haxe_<profile> (expected at least one haxe tui metric)");
}

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

function startupMeasurementMs(metric) {
  const median = Number(metric.startup_median_ms);
  if (Number.isFinite(median) && median > 0) {
    return median;
  }
  return metric.startup_avg_ms;
}

function buildCaseOverhead(caseName) {
  const pure = requireMetric(`${caseName}_pure_go`);
  const out = {};
  for (const profile of profiles) {
    const metric = requireMetric(`${caseName}_haxe_${profile}`);
    out[profile] = {
      binaryRatio: ratio(metric.binary_bytes, pure.binary_bytes),
      strippedRatio: ratio(metric.stripped_bytes, pure.stripped_bytes),
      startupRatio: ratio(startupMeasurementMs(metric), startupMeasurementMs(pure)),
    };
  }
  return out;
}

const helloOverheadRatios = buildCaseOverhead("hello");
const arrayOverheadRatios = buildCaseOverhead("array");
const atomicOverheadRatios = buildCaseOverhead("atomic");
const channelOverheadRatios = buildCaseOverhead("channel");
const mapOverheadRatios = buildCaseOverhead("map");
const genericOverheadRatios = buildCaseOverhead("generic");
const stringOverheadRatios = buildCaseOverhead("string");
const stringInstanceOverheadRatios = buildCaseOverhead("string_instance");
const virtualOverheadRatios = buildCaseOverhead("virtual");
const selectOverheadRatios = buildCaseOverhead("select");
const caseOverheadByName = {
  hello: helloOverheadRatios,
  array: arrayOverheadRatios,
  atomic: atomicOverheadRatios,
  channel: channelOverheadRatios,
  map: mapOverheadRatios,
  generic: genericOverheadRatios,
  string: stringOverheadRatios,
  string_instance: stringInstanceOverheadRatios,
  virtual: virtualOverheadRatios,
  select: selectOverheadRatios,
};

function buildPortableMetalDeltaRatios(overheadByCase) {
  const out = {};
  for (const [caseName, ratioGroup] of Object.entries(overheadByCase || {})) {
    const portable = ratioGroup?.portable;
    const metal = ratioGroup?.metal;
    if (!portable || !metal) {
      continue;
    }
    out[caseName] = {
      binaryRatio: ratio(portable.binaryRatio, metal.binaryRatio),
      strippedRatio: ratio(portable.strippedRatio, metal.strippedRatio),
      startupRatio: ratio(portable.startupRatio, metal.startupRatio),
    };
  }
  return out;
}

function deriveCaseOverheadByName(derived) {
  return {
    hello: derived?.helloOverheadRatios,
    array: derived?.arrayOverheadRatios,
    atomic: derived?.atomicOverheadRatios,
    channel: derived?.channelOverheadRatios,
    map: derived?.mapOverheadRatios,
    generic: derived?.genericOverheadRatios,
    string: derived?.stringOverheadRatios,
    string_instance: derived?.stringInstanceOverheadRatios,
    virtual: derived?.virtualOverheadRatios,
    select: derived?.selectOverheadRatios,
  };
}

function derivePortableMetalDeltaRatios(derived) {
  if (derived?.portableMetalDeltaRatios && typeof derived.portableMetalDeltaRatios === "object") {
    return derived.portableMetalDeltaRatios;
  }
  return buildPortableMetalDeltaRatios(deriveCaseOverheadByName(derived));
}

const portableMetalDeltaRatios = buildPortableMetalDeltaRatios(caseOverheadByName);

const tuiMetrics = Object.fromEntries(
  tuiProfiles.map((profile) => [profile, requireMetric(`tui_haxe_${profile}`)])
);
const tuiMin = {
  binary_bytes: Math.min(...tuiProfiles.map((profile) => tuiMetrics[profile].binary_bytes)),
  stripped_bytes: Math.min(...tuiProfiles.map((profile) => tuiMetrics[profile].stripped_bytes)),
  startup_measurement_ms: Math.min(...tuiProfiles.map((profile) => startupMeasurementMs(tuiMetrics[profile]))),
};
const tuiRelativeToMin = {};
for (const profile of tuiProfiles) {
  const metric = tuiMetrics[profile];
  tuiRelativeToMin[profile] = {
    binaryRatio: ratio(metric.binary_bytes, tuiMin.binary_bytes),
    strippedRatio: ratio(metric.stripped_bytes, tuiMin.stripped_bytes),
    startupRatio: ratio(startupMeasurementMs(metric), tuiMin.startup_measurement_ms),
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
    deltaWarnPct,
    deltaFailPct,
  },
  startupLoops: {
    hello: helloIters,
    array: arrayIters,
    atomic: atomicIters,
    channel: channelIters,
    map: mapIters,
    generic: genericIters,
    string: stringIters,
    string_instance: stringInstanceIters,
    virtual: virtualIters,
    select: selectIters,
    tui: tuiIters,
  },
  caseParams: {
    atomicWork,
    channelWork,
    mapWork,
    genericWork,
    stringWork,
    stringInstanceWork,
    virtualWork,
    selectWork,
  },
  metrics,
  options: {
    portableConcurrencyFastpathEnabled,
    hxrtFeatures,
    enforceMetalBudget,
    enforceDeltaBudget,
    deltaCases: uniqueDeltaCases,
    tuiProfiles,
    startupSamples,
  },
  derived: {
    helloOverheadRatios,
    arrayOverheadRatios,
    atomicOverheadRatios,
    channelOverheadRatios,
    mapOverheadRatios,
    genericOverheadRatios,
    stringOverheadRatios,
    stringInstanceOverheadRatios,
    virtualOverheadRatios,
    selectOverheadRatios,
    portableMetalDeltaRatios,
    tuiRelativeToMin,
  },
};

fs.mkdirSync(path.dirname(currentJsonPath), { recursive: true });
fs.writeFileSync(currentJsonPath, `${JSON.stringify(current, null, 2)}\n`);

const baselinePayload = {
  schemaVersion: 1,
  generatedAt: current.generatedAt,
  toolchain: current.toolchain,
  thresholds: current.thresholds,
  startupLoops: current.startupLoops,
  caseParams: current.caseParams,
  options: current.options,
  derivedBaseline: current.derived,
};

if (updateBaseline) {
  fs.mkdirSync(path.dirname(baselinePath), { recursive: true });
  fs.writeFileSync(baselinePath, `${JSON.stringify(baselinePayload, null, 2)}\n`);
}

const warnings = [];
const hardFailures = [];

function compareGroup(groupLabel, currentGroup, baselineGroup, profileList = profiles) {
  if (!baselineGroup) {
    warnings.push(`${groupLabel}: missing baseline group`);
    return;
  }

  const specs = [
    { key: "binaryRatio", label: "binary ratio", warnPct: sizeWarnPct },
    { key: "strippedRatio", label: "stripped ratio", warnPct: sizeWarnPct },
    { key: "startupRatio", label: "startup ratio", warnPct: runtimeWarnPct },
  ];

  for (const profile of profileList) {
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

function comparePortableMetalDelta(currentDeltaRatios, baselineDeltaRatios) {
  if (!baselineDeltaRatios || typeof baselineDeltaRatios !== "object") {
    warnings.push("delta.portable_metal: missing baseline group");
    return;
  }

  for (const caseName of uniqueDeltaCases) {
    const currentCase = currentDeltaRatios?.[caseName];
    const baselineCase = baselineDeltaRatios?.[caseName];
    if (!currentCase || !baselineCase) {
      warnings.push(`delta.${caseName}.startup ratio missing data in current/baseline`);
      continue;
    }
    const currentValue = Number(currentCase.startupRatio);
    const baselineValue = Number(baselineCase.startupRatio);
    if (!Number.isFinite(currentValue) || !Number.isFinite(baselineValue) || baselineValue <= 0) {
      continue;
    }

    const warnAllowed = baselineValue * (1 + deltaWarnPct / 100);
    if (currentValue > warnAllowed) {
      const increasePct = ((currentValue / baselineValue) - 1) * 100;
      warnings.push(
        `delta.${caseName}.startup ratio +${increasePct.toFixed(2)}% ` +
          `(current=${currentValue.toFixed(6)}, baseline=${baselineValue.toFixed(6)}, budget=+${deltaWarnPct.toFixed(2)}%)`
      );
    }

    const failAllowed = baselineValue * (1 + deltaFailPct / 100);
    if (currentValue > failAllowed) {
      const increasePct = ((currentValue / baselineValue) - 1) * 100;
      hardFailures.push(
        `delta.${caseName}.startup ratio +${increasePct.toFixed(2)}% ` +
          `(current=${currentValue.toFixed(6)}, baseline=${baselineValue.toFixed(6)}, budget=+${deltaFailPct.toFixed(2)}%)`
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
    const baselinePortableMetalDeltaRatios = derivePortableMetalDeltaRatios(baselineDerived);
    compareGroup("hello_overhead", current.derived.helloOverheadRatios, baselineDerived.helloOverheadRatios);
    compareGroup("array_overhead", current.derived.arrayOverheadRatios, baselineDerived.arrayOverheadRatios);
    compareGroup("atomic_overhead", current.derived.atomicOverheadRatios, baselineDerived.atomicOverheadRatios);
    compareGroup("channel_overhead", current.derived.channelOverheadRatios, baselineDerived.channelOverheadRatios);
    compareGroup("map_overhead", current.derived.mapOverheadRatios, baselineDerived.mapOverheadRatios);
    compareGroup("generic_overhead", current.derived.genericOverheadRatios, baselineDerived.genericOverheadRatios);
    compareGroup("string_overhead", current.derived.stringOverheadRatios, baselineDerived.stringOverheadRatios);
    compareGroup("string_instance_overhead", current.derived.stringInstanceOverheadRatios, baselineDerived.stringInstanceOverheadRatios);
    compareGroup("virtual_overhead", current.derived.virtualOverheadRatios, baselineDerived.virtualOverheadRatios);
    compareGroup("select_overhead", current.derived.selectOverheadRatios, baselineDerived.selectOverheadRatios);
    compareGroup("tui_relative", current.derived.tuiRelativeToMin, baselineDerived.tuiRelativeToMin, tuiProfiles);
    compareMetalHard("hello_overhead", current.derived.helloOverheadRatios, baselineDerived.helloOverheadRatios);
    compareMetalHard("array_overhead", current.derived.arrayOverheadRatios, baselineDerived.arrayOverheadRatios);
    compareMetalHard("atomic_overhead", current.derived.atomicOverheadRatios, baselineDerived.atomicOverheadRatios);
    compareMetalHard("channel_overhead", current.derived.channelOverheadRatios, baselineDerived.channelOverheadRatios);
    compareMetalHard("map_overhead", current.derived.mapOverheadRatios, baselineDerived.mapOverheadRatios);
    compareMetalHard("generic_overhead", current.derived.genericOverheadRatios, baselineDerived.genericOverheadRatios);
    compareMetalHard("string_overhead", current.derived.stringOverheadRatios, baselineDerived.stringOverheadRatios);
    compareMetalHard("string_instance_overhead", current.derived.stringInstanceOverheadRatios, baselineDerived.stringInstanceOverheadRatios);
    compareMetalHard("virtual_overhead", current.derived.virtualOverheadRatios, baselineDerived.virtualOverheadRatios);
    compareMetalHard("select_overhead", current.derived.selectOverheadRatios, baselineDerived.selectOverheadRatios);
    comparePortableMetalDelta(current.derived.portableMetalDeltaRatios, baselinePortableMetalDeltaRatios);
  }
}

const deltaWarnings = warnings.filter((warning) => warning.startsWith("delta."));
const metalWarnings = warnings.filter((warning) => warning.includes(".metal."));
const deltaHardFailures = hardFailures.filter((failure) => failure.startsWith("delta."));
const metalHardFailures = hardFailures.filter((failure) => !failure.startsWith("delta."));

const comparison = {
  schemaVersion: 1,
  generatedAt: current.generatedAt,
  mode: updateBaseline ? "update-baseline" : "compare",
  baselinePath: baselineDisplay,
  baselineAvailable: baselineLoaded != null || updateBaseline,
  portableConcurrencyFastpathEnabled,
  enforceMetalBudget,
  enforceDeltaBudget,
  deltaCases: uniqueDeltaCases,
  metalHardFailureCount: metalHardFailures.length,
  deltaHardFailureCount: deltaHardFailures.length,
  metalHardFailureBudgets: {
    sizeFailPct: metalSizeFailPct,
    runtimeFailPct: metalRuntimeFailPct,
  },
  deltaHardFailureBudgets: {
    startupFailPct: deltaFailPct,
  },
  metalWarningCount: metalWarnings.length,
  deltaWarningCount: deltaWarnings.length,
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

function ratioTable(title, ratioGroup, profileList = profiles) {
  const lines = [];
  lines.push(`### ${title}`);
  lines.push("| Profile | Binary x | Stripped x | Startup x |\n| --- | ---: | ---: | ---: |");
  for (const profile of profileList) {
    const row = ratioGroup[profile];
    if (!row) {
      lines.push(`| ${profile} | - | - | - |`);
      continue;
    }
    lines.push(
      `| ${profile} | ${formatRatio(row.binaryRatio)} | ${formatRatio(row.strippedRatio)} | ${formatRatio(row.startupRatio)} |`
    );
  }
  lines.push("");
  return lines.join("\n");
}

function portableMetalDeltaTable(title, deltaRatioGroup) {
  const lines = [];
  lines.push(`### ${title}`);
  lines.push("| Case | Binary x | Stripped x | Startup x |\n| --- | ---: | ---: | ---: |");
  for (const caseName of Object.keys(deltaRatioGroup).sort()) {
    const row = deltaRatioGroup[caseName];
    lines.push(
      `| ${caseName} | ${formatRatio(row.binaryRatio)} | ${formatRatio(row.strippedRatio)} | ${formatRatio(row.startupRatio)} |`
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
summaryLines.push(`- Portable-vs-metal delta warn budget: \`+${deltaWarnPct}%\` (startup ratio drift vs baseline)`);
summaryLines.push(`- Metal enforcement: \`${enforceMetalBudget ? "on" : "off"}\``);
summaryLines.push(`- Delta enforcement: \`${enforceDeltaBudget ? "on" : "off"}\``);
summaryLines.push(`- Metal hard budgets: size=\`+${metalSizeFailPct}%\`, runtime=\`+${metalRuntimeFailPct}%\``);
summaryLines.push(`- Delta hard budget: startup=\`+${deltaFailPct}%\` for cases=\`${uniqueDeltaCases.join(",") || "none"}\``);
summaryLines.push(`- Portable concurrency fastpath: \`${portableConcurrencyFastpathEnabled ? "on" : "off"}\``);
summaryLines.push(`- Microbench hxrt features: \`${hxrtFeatures}\``);
summaryLines.push(`- Startup samples: \`${startupSamples}\` (startup ratios use the median sample)`);
summaryLines.push(`- Startup loops: hello=${helloIters}, array=${arrayIters}, atomic=${atomicIters}, channel=${channelIters}, map=${mapIters}, generic=${genericIters}, string=${stringIters}, string_instance=${stringInstanceIters}, virtual=${virtualIters}, select=${selectIters}, tui=${tuiIters}`);
summaryLines.push(`- Workload params: atomic_ops=${atomicWork}, channel_ops=${channelWork}, map_ops=${mapWork}, generic_ops=${genericWork}, string_ops=${stringWork}, string_instance_ops=${stringInstanceWork}, virtual_ops=${virtualWork}, select_ops=${selectWork}`);
if (haxeVersion.length > 0 || goVersion.length > 0) {
  summaryLines.push(`- Toolchain: ${haxeVersion || "haxe:unknown"} | ${goVersion || "go:unknown"}`);
}
summaryLines.push("");
summaryLines.push(ratioTable("Hello Overhead (x vs pure Go hello)", current.derived.helloOverheadRatios));
summaryLines.push(ratioTable("Array Overhead (x vs pure Go array loop)", current.derived.arrayOverheadRatios));
summaryLines.push(ratioTable("Atomic Overhead (x vs pure Go atomic loop)", current.derived.atomicOverheadRatios));
summaryLines.push(ratioTable("Channel Overhead (x vs pure Go buffered channel loop)", current.derived.channelOverheadRatios));
summaryLines.push(ratioTable("Map Overhead (x vs pure Go map set/get loop)", current.derived.mapOverheadRatios));
summaryLines.push(ratioTable("Generic Overhead (x vs pure Go generic bag loop)", current.derived.genericOverheadRatios));
summaryLines.push(ratioTable("String Overhead (x vs pure Go concat loop)", current.derived.stringOverheadRatios));
summaryLines.push(ratioTable("String Instance Overhead (x vs pure Go string instance loop)", current.derived.stringInstanceOverheadRatios));
summaryLines.push(ratioTable("Virtual Overhead (x vs pure Go interface dispatch loop)", current.derived.virtualOverheadRatios));
summaryLines.push(ratioTable("Select Overhead (x vs pure Go select helper loop)", current.derived.selectOverheadRatios));
summaryLines.push(portableMetalDeltaTable("Portable-vs-metal Delta (portable ratio / metal ratio)", current.derived.portableMetalDeltaRatios));
summaryLines.push(ratioTable("TUI Profile Spread (x vs fastest/smallest profile in this run)", current.derived.tuiRelativeToMin, tuiProfiles));

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

if (metalHardFailures.length > 0) {
  summaryLines.push("### Metal Hard-Fail Candidates");
  for (const hardFailure of metalHardFailures) {
    summaryLines.push(`- ${hardFailure}`);
  }
} else {
  summaryLines.push("### Metal Hard-Fail Candidates");
  summaryLines.push("- none");
}
summaryLines.push("");

if (deltaHardFailures.length > 0) {
  summaryLines.push("### Delta Hard-Fail Candidates");
  for (const hardFailure of deltaHardFailures) {
    summaryLines.push(`- ${hardFailure}`);
  }
} else {
  summaryLines.push("### Delta Hard-Fail Candidates");
  summaryLines.push("- none");
}
summaryLines.push("");

fs.writeFileSync(summaryPath, `${summaryLines.join("\n")}\n`);

console.log(`[go-perf] mode=${comparison.mode} warnings=${warnings.length}`);
NODE

python3 scripts/ci/perf-warning-summary.py \
  --harness go-profile \
  --comparison "$comparison_json" \
  --out-json "$warning_history_json" \
  --out-md "$warning_history_md"

python3 scripts/ci/perf-delta-dry-run.py \
  --harness go-profile \
  --comparison "$comparison_json" \
  --out-json "$delta_dry_run_json" \
  --out-md "$delta_dry_run_md"

warning_count=0
metal_warning_count=0
delta_warning_count=0
baseline_warning_count=0
hard_failure_count=0
metal_hard_failure_count=0
delta_hard_failure_count=0
if [[ -s "$warnings_txt" ]]; then
  while IFS= read -r warning; do
    [[ -n "$warning" ]] || continue
    warning_count=$((warning_count + 1))
    if [[ "$warning" == delta.* ]]; then
      delta_warning_count=$((delta_warning_count + 1))
    fi
    if [[ "$warning" == *".metal."* ]]; then
      metal_warning_count=$((metal_warning_count + 1))
    fi
    if [[ "$warning" == baseline\ file\ not\ found:* ]]; then
      baseline_warning_count=$((baseline_warning_count + 1))
    fi
    echo "::warning::[go-perf][soft-budget-signal] $warning (warning-only; see docs/performance-budget-policy.md and warning_history artifacts)"
  done < "$warnings_txt"
fi

if [[ -s "$hard_failures_txt" ]]; then
  while IFS= read -r hard_failure; do
    [[ -n "$hard_failure" ]] || continue
    hard_failure_count=$((hard_failure_count + 1))
    if [[ "$hard_failure" == delta.* ]]; then
      delta_hard_failure_count=$((delta_hard_failure_count + 1))
      if is_truthy "$enforce_delta_budget"; then
        echo "::error::[go-perf] $hard_failure"
      else
        echo "::warning::[go-perf][delta-hard-candidate][not-enforced] $hard_failure (hard gate is disabled; see docs/performance-budget-policy.md and delta_hard_gate_dry_run artifacts)"
      fi
    else
      metal_hard_failure_count=$((metal_hard_failure_count + 1))
      if is_truthy "$enforce_metal_budget"; then
        echo "::error::[go-perf] $hard_failure"
      else
        echo "::warning::[go-perf][metal-hard-candidate][not-enforced] $hard_failure (hard gate is disabled; see docs/performance-budget-policy.md and delta_hard_gate_dry_run artifacts)"
      fi
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
  if [[ "$metal_hard_failure_count" -gt 0 || "$baseline_warning_count" -gt 0 ]]; then
    echo "::error::[go-perf] metal budget enforcement failed (hard_failures=$metal_hard_failure_count baseline_warnings=$baseline_warning_count)"
    log "failing due to GO_PERF_ENFORCE_METAL_BUDGET with budget regressions"
    log "metrics: $(display_path "$current_json")"
    log "comparison: $(display_path "$comparison_json")"
    log "summary: $(display_path "$summary_md")"
    exit 1
  fi
fi

if is_truthy "$enforce_delta_budget"; then
  if [[ "$delta_hard_failure_count" -gt 0 || "$baseline_warning_count" -gt 0 ]]; then
    echo "::error::[go-perf] delta budget enforcement failed (hard_failures=$delta_hard_failure_count baseline_warnings=$baseline_warning_count)"
    log "failing due to GO_PERF_ENFORCE_DELTA_BUDGET with budget regressions"
    log "metrics: $(display_path "$current_json")"
    log "comparison: $(display_path "$comparison_json")"
    log "summary: $(display_path "$summary_md")"
    exit 1
  fi
fi

log "done (warnings=$warning_count, metal_warnings=$metal_warning_count, delta_warnings=$delta_warning_count, metal_hard_failures=$metal_hard_failure_count, delta_hard_failures=$delta_hard_failure_count)"
log "metrics: $(display_path "$current_json")"
log "comparison: $(display_path "$comparison_json")"
log "summary: $(display_path "$summary_md")"
