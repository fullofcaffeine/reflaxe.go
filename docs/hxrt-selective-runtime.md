# Selective `hxrt` Runtime Plan

## Decision Summary

Selective `hxrt` is a runtime-packaging policy. It is orthogonal to source
semantics and to the compatible `portable|metal` policy presets.

- Portable semantics remain the default product contract.
- Typed APIs/externs and `@:goNative` declare Go-native source boundaries.
- Selective `hxrt` minimizes copied runtime support.

These concerns should evolve together without collapsing into one toggle.

## Why runtime slicing is separate

`metal + selective hxrt` reduces emitted runtime footprint, but it does not answer:

- what semantic compatibility baseline we guarantee,
- what strict-boundary behavior is enforced by default,
- what optimization/interoperability behaviors are opt-in vs guaranteed.

Portable Haxe remains the semantic baseline. Typed native APIs and modules
remain the interoperability boundary. `metal` remains a supported compatibility
preset for stricter/eager defaults.

## Goals

1. Keep Haxe source portability guarantees explicit (`portable` baseline).
2. Make generated output as close as possible to handwritten Go wherever the
   selected source semantics permit it.
3. Trim runtime footprint via deterministic feature inference + override controls.

## What the runtime manifest is

The runtime manifest is the compiler's single, typed answer to four questions:

1. Which closed `hxrt` capabilities does this program need?
2. What typed usage, surface contract, dependency, define, or compatibility
   contract selected each capability?
3. Which runtime files belong directly to each selected capability?
4. What exact file set must the generated project contain?

`GoRuntimeCapabilityManifest` builds that immutable answer after typed lowering.
Both the file copier and `hxrt_plan.json` schema v4 consume it. The report's
`capabilities` entries therefore explain the files that were actually copied;
the report does not independently infer them.

Selective mode includes only evidenced capabilities and their dependencies.
Compatibility full-copy mode remains supported, but broad capabilities are
explicitly attributed to the `default_full_copy` compatibility contract.
Footprint-explicit capabilities such as sockets, terminal access, reflection,
and native stack capture still require typed use or an explicit define.

No-`hxrt` facts come from the portable surface registry. Each reported surface
decision keeps the reviewed registry status and separately states whether the
representation selected for that use has no runtime requirement. An empty file
or import list cannot grant eligibility by itself.

The runtime review also removed the old exported `IntMapSnapshot`,
`StringMapSnapshot`, and `ObjectMapSnapshot` bridges. They had no staged Haxe
binding or current generated caller. A retired compiler serializer emitter did
call them before commit `abd42e87` moved serialization into staged source;
staged Serializer now walks the typed map `Keys` and `Get` APIs instead. With
that former owner gone and no admitted public compatibility consumer, keeping
the helpers would have treated an unowned migration bridge as a compatibility
contract. The ordinary map capabilities and their public Haxe behavior are
unchanged.

## Implementation Tracks

1. Runtime slice split:
   - Split `runtime/hxrt/hxrt.go` into `runtime/hxrt/*.go` feature groups.
   - Keep package/API stable.
2. Feature inference:
   - Infer required runtime features from used module/type surfaces and compiler shim groups.
   - Keep ordering deterministic and dependency-complete.
3. Selective runtime emit:
   - Copy only required runtime files when selective mode is enabled.
   - Preserve full-copy behavior for compatibility and fallback.

## Define Matrix

- `reflaxe_go_hxrt_default_features`
  - Force full runtime copy (compat mode).
  - Takes precedence over selective runtime flags.
- `reflaxe_go_hxrt_features=core,json,sys,terminal,file_io,filesystem,http,socket,ssl,socket_ssl,...`
  - Enables selective runtime mode and adds manual feature list.
  - Use empty value (`-D reflaxe_go_hxrt_features=`) to enable selective mode with inferred-only features.

`sys.FileSystem` usage infers the dedicated `filesystem` feature, which copies
`runtime/hxrt/filesystem.go`. Keeping it separate prevents unrelated `sys.*`
programs from inheriting native filesystem support in selective mode.

Portable root `Array<T>` usage selects the dedicated `array` feature and copies
`runtime/hxrt/array.go`. Runtime features that inspect or construct a portable
Array, including staged `haxe.Template` and `haxe.Json` support, depend on that
feature explicitly so selective output never references a carrier definition it
did not copy.

`sys.io.File`, its stream classes, and root `Sys` standard streams infer the
dedicated `file_io` feature, which copies `runtime/hxrt/file.go`. Other staged
root `Sys` capabilities infer `sys`, while typed `hxrt.process.NativeProcess`
usage under staged `sys.io.Process` infers `process`. Direct File, root Sys, and
Process use therefore do not pull one another's native slices. `core/runtime_hxrt_infer_sys` and
`core/runtime_hxrt_infer_process` lock both positive and negative file sets.

`Sys.getChar` use selects the dedicated `terminal` feature through typed
`hxrt.sys.NativeTerminal` authority. Its six platform/build-tag files remain
footprint-explicit even in ordinary full-copy mode so unrelated output never
acquires the POSIX unsafe boundary. Explicitly disabling feature inference
retains the traditional all-files full-copy escape.

`sys.net.Host`, `sys.net.Socket`, `sys.net.UdpSocket`, and their typed
`hxrt.net` bindings infer `socket`, which copies `runtime/hxrt/socket.go`, its
build-tagged `socket_broadcast_*.go` option adapters,
`socket_listener_*.go` bind/listen adapters, and the `string` and `exception`
dependencies. Plain programs and SSL leaf-only programs keep the OS networking
capability out, including ordinary full-copy mode.
`core/runtime_hxrt_infer_socket` locks that positive and negative shape, while
`test/test_socket_runtime_cross_build.py` compiles the runtime for POSIX and
Windows so descriptor-type differences cannot regress silently.

`sys.Http` and typed `hxrt.http` bindings infer `http`, which copies
footprint-explicit `runtime/hxrt/http.go` plus `string`, `bytes`, and `socket`.
The socket dependency is intentional: `customRequest` accepts a typed
`sys.net.Socket`, so the HTTP capability must compile against the same opaque
`SocketHandle` even when a particular request does not supply one.
`core/runtime_hxrt_infer_http` locks the positive file set, while unrelated
selective-runtime cases remain free of `http.go`. Ordinary full-copy mode also
keeps `http.go` footprint-explicit unless typed use, a manual `http` feature, or
disabled inference requests it.

`sys.ssl.Certificate`, `Digest`, and `Key` infer `ssl` without networking.
`sys.ssl.Socket` and `hxrt.ssl.NativeSocket` infer `socket_ssl`; that feature
depends on both `socket` and `ssl` and additionally copies
`runtime/hxrt/socket_ssl.go`. `core/runtime_hxrt_infer_ssl` and
`core/runtime_hxrt_infer_socket_ssl` prove the leaf/transport split.

`EReg` and typed `hxrt.regex` bindings infer `regex`, which copies
`runtime/hxrt/regex.go` plus its string/exception dependencies. `haxe.Serializer`,
`haxe.Unserializer`, and typed `hxrt.serialization` bindings infer
`serialization`, which copies `runtime/hxrt/serialization.go` plus string and
equality support. Their staged `Reflect` calls also select the shared,
memory-safe `runtime/hxrt/reflect.go` helper. That helper provides the dynamic
object and anonymous-structure fallback; generated class fields and methods
still use compiler-emitted typed metadata. The two capabilities do not depend on
one another:
`core/runtime_hxrt_infer_regex` and
`core/runtime_hxrt_infer_serialization` lock both positive and negative file
sets. Both native capability files remain footprint-explicit in broad full-copy
mode unless typed use, manual feature selection, or disabled inference requests
them; that keeps RE2 and serialization float-parsing support out of unrelated
programs. Generated private-field access for serialization reuses the ordinary
typed Reflect metadata already selected by the staged calls; it does not add an
unsafe runtime slice.
- `reflaxe_go_hxrt_no_feature_infer`
  - Enables selective runtime mode and disables inference (use core + manual only).

## Rollout Policy

Phase 1:
- Keep full runtime copy as default.
- Add selective mode behind defines + tests (`reflaxe_go_hxrt_features` and/or `reflaxe_go_hxrt_no_feature_infer`).

Phase 2:
- Promote selective mode based on coverage/perf evidence, independently of the
  selected compatibility preset.
- Keep explicit full-copy fallback for debugging and migrations.

## Perf/Size Harness

Run selective-vs-full runtime footprint metrics:

```bash
bash scripts/ci/perf-hxrt-selective.sh
```

Optional hard budget enforcement (used in CI):

```bash
GO_HXRT_SLICE_ENFORCE=1 bash scripts/ci/perf-hxrt-selective.sh
```

Regenerate baseline:

```bash
bash scripts/ci/perf-hxrt-selective.sh --update-baseline
```

Artifacts:

- `.cache/perf-hxrt-selective/results/current.json`
- `.cache/perf-hxrt-selective/results/comparison.json`
- `.cache/perf-hxrt-selective/results/summary.md`
- `scripts/ci/perf/hxrt-selective-baseline.json`

The harness includes a serialization-specific footprint case. It compiles and
links a private-field class round trip, so both the small float parser and the
shared Reflect helper are counted in source-file and binary-size budgets. Every
lane enables `hxrt_plan.json`, verifies that its manifest authority and
per-capability reasons are present, checks the reported files against the
generated directory, and measures source bytes from that manifest.

Interpretation:

- `file_delta`: selective runtime files minus full runtime files (should be `<= 0`).
- `source_delta_pct`: selective runtime source bytes change vs full runtime (should be `<= 0` in normal cases).
- `binary_delta_pct`: selective binary bytes change vs full runtime (expected near 0; a small positive drift can happen by case/toolchain).
- `drift` columns in `summary.md`: delta vs `scripts/ci/perf/hxrt-selective-baseline.json` to track trend over time.

## Related Beads

- Epic: `haxe.go-e73`
- Canonical sequencing is tracked in bead dependencies under that epic.
