# pulseforge

Scaffold for the flagship observability-stream app.

This milestone intentionally keeps scope small while proving the execution model:

- one shared Haxe codebase
- all profiles compile (`portable`, `metal`)
- explicit variant plumbing (`core`, `go_native`)
- deterministic scripted output for CI

## Compile

Core variant (default compile files):

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

Go-native variant (CI compile files):

```bash
haxe compile.portable.ci.hxml
haxe compile.metal.ci.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_metal && go run .)
```

## Variant strategy

- `core` variant:
  - deterministic loop-based runtime path (`runtime.capability=core_loop`)
- `go_native` variant:
  - typed channel/select path (`runtime.capability=chan_select`)

Current profile behavior differs by build identity and generated code shape, while keeping the same contract-level domain metrics.

## Matrix expectation

- `*.stdout` files represent `core` variant output.
- `*.ci.stdout` files represent `go_native` variant output.
