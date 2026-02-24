# PulseForge Pure-Go Baseline

Pure-Go parity baseline for `examples/pulseforge`.

This module mirrors the same baseline workload and output contract used by the Haxe PulseForge harness so we can compare Haxe->Go profiles against a handwritten Go implementation.

## Run

```bash
cd benchmarks/pure_go/pulseforge
go run . --scripted
go run . --scripted --variant core
go run . --scripted --variant go_native
```

## Validate

```bash
cd benchmarks/pure_go/pulseforge
go test ./...
```

## Benchmark

```bash
cd benchmarks/pure_go/pulseforge
go test -run '^$' -bench . -benchmem
```
