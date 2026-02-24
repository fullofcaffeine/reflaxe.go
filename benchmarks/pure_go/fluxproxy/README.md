# FluxProxy Pure-Go Baseline

Pure-Go parity baseline for `examples/fluxproxy`.

This module mirrors the same baseline workload and output contract used by the Haxe FluxProxy harness so we can compare Haxe->Go profiles against a handwritten Go implementation.

## Run

```bash
cd benchmarks/pure_go/fluxproxy
go run . --scripted
go run . --scripted --variant core
go run . --scripted --variant go_native
```

## Validate

```bash
cd benchmarks/pure_go/fluxproxy
go test ./...
```

## Benchmark

```bash
cd benchmarks/pure_go/fluxproxy
go test -run '^$' -bench . -benchmem
```
